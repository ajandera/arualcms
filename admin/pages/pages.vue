<template>
  <v-data-table
    :headers="headers"
    :items="pages"
    :items-per-page="10"
    class="elevation-1"
  >
    <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>My Pages</v-toolbar-title>
        <v-divider
          class="mx-4"
          inset
          vertical
        ></v-divider>
        <v-spacer></v-spacer>
        <v-dialog
          v-model="dialog"
          max-width="1000px"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              color="primary"
              dark
              class="mb-2"
              v-bind="attrs"
              v-on="on"
              @click="create"
            >
              Add
            </v-btn>
          </template>
          <v-card>
            <br>
            <v-form
              ref="form"
              v-model="valid"
              lazy-validation>
              <v-card-title></v-card-title>
              <v-card-text>
                <v-container>
                  <v-row>
                    <v-col
                      cols="12"
                      sm="6"
                      md="6"
                    >
                      <v-text-field
                        v-model="page.Title[language]"
                        :counter="50"
                        :rules="rules"
                        label="Title"
                        required
                      ></v-text-field>
                    </v-col>
                    <v-col
                      cols="12"
                      sm="6"
                      md="6"
                    >
                      <v-text-field
                        v-model="page.Key"
                        :counter="50"
                        :rules="rules"
                        label="Key"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col
                      cols="12"
                    >
                      <quill-editor
                        :ref="page.Body[language]"
                        v-model="page.Body[language]"
                        :options="editorOption"
                        @blur="onEditorBlur($event)"
                        @focus="onEditorFocus($event)"
                        @ready="onEditorReady($event)"
                      />
                      <v-text-field
                        v-model="page.MetaTitle[language]"
                        :counter="100"
                        :rules="rules"
                        label="Meta title"
                        required
                      ></v-text-field>
                      <v-text-field
                        v-model="page.Description[language]"
                        :counter="150"
                        :rules="rules"
                        label="Meta description"
                        required
                      ></v-text-field>
                      <v-text-field
                        v-model="page.Keywords[language]"
                        :counter="100"
                        :rules="rules"
                        label="Meta keywords"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-row>
                </v-container>
              </v-card-text>
              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                  color="blue darken-1"
                  text
                  @click="close"
                >
                  Cancel
                </v-btn>
                <v-btn
                  color="primary"
                  :disabled="!valid"
                  @click="save"
                >
                  Save
                </v-btn>
              </v-card-actions>
            </v-form>
          </v-card>
        </v-dialog>
      </v-toolbar>
    </template>
    <template v-slot:item.Src="{ item }">
      <div class="p-5">
        <v-img :src="$config.storage + item.Src" :alt="item.Title[language]" height="auto"
               width="200px"></v-img>
      </div>
    </template>
    <template v-slot:item.Title="{ item }">
      <p>{{ item.Title[language] }}</p>
    </template>
    <template v-slot:item.Excerpt="{ item }">
      <p>{{ item.Excerpt[language] }}</p>
    </template>
    <template v-slot:item.Published="{ item }">
      <p>{{ item.Published[0].toLocaleDateString() }}</p>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon
        small
        class="mr-2"
        @click="edit(item)"
      >
        mdi-pencil
      </v-icon>
      <v-icon
        small
        class="mr-2"
        @click="translate(item)"
      >
        mdi-flag
      </v-icon>
      <v-icon
        small
        class="mr-2"
        @click="remove(item.Id)"
      >
        mdi-delete
      </v-icon>
    </template>
    <template v-slot:no-data>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import {Component, Prop, Vue, Watch} from 'nuxt-property-decorator';
import Page from "~/model/Page";
import IResponsePages from "~/model/IResponsePages";
import IHeader from '~/model/IHeader';
import {namespace} from 'vuex-class';
import Text from "~/model/Text";

const snackbar = namespace('Snackbar');

@Component
export default class PagesPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void
  @Prop() readonly languages!: string[];
  @Prop() readonly language!: string;
  @Prop() readonly defaultLanguage!: string;

  pages: Page[] = [];
  title: string = "";
  editorOption: any = {
    // Some Quill options...
  };
  $axios: any;
  content: any;
  $refs: any;
  headers: IHeader[] = [
    {text: "Title", value: 'Title'},
    {text: "Actions", value: 'actions', sortable: false}
  ];
  page: Page = {
    Body: this.createClearTranslationObject(),
    Title: this.createClearTranslationObject(),
    Id: '',
    MetaTitle: this.createClearTranslationObject(),
    Keywords: this.createClearTranslationObject(),
    Description: this.createClearTranslationObject(),
    Key: ""
  };
  dialog: boolean = false;
  valid: boolean = true;
  rules: any = [];
  fromDateMenu: boolean = false;

  mounted() {
    this.load();
  }

  @Watch('$route.query')
  onPropertyChanged(value: string, oldValue: string) {
    this.load();
  }

  edit(page: Page) {
    this.page = page;
    this.title = this.page.Title[this.language];
    this.dialog = true;
  }

  remove(id: string) {
    this.$axios.delete("/" + this.$route.query.siteId + "/pages/" + id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponsePages) => {
        if (response.data.success) {
          this.updateText(response.data.message);
          this.updateColor('green')
          this.updateShow(true);
          this.load();
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  create() {
    this.page = {
      Body: this.createClearTranslationObject(),
      Title: this.createClearTranslationObject(),
      Id: '',
      MetaTitle: this.createClearTranslationObject(),
      Keywords: this.createClearTranslationObject(),
      Description: this.createClearTranslationObject(),
      Key: ""
    };
  }

  save(close: boolean) {
    this.page.Title = JSON.stringify(this.page.Title);
    this.page.Body = JSON.stringify(this.page.Body);
    this.page.MetaTitle = JSON.stringify(this.page.MetaTitle);
    this.page.Keywords = JSON.stringify(this.page.Keywords);
    this.page.Description = JSON.stringify(this.page.Description);
    if (this.page.Id !== "") {
      this.$axios.put("/" + this.$route.query.siteId + "/pages", this.page, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponsePages) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
          }
          this.load();
        });
    } else {
      this.page.Id = "";
      this.$axios.post("/" + this.$route.query.siteId + "/pages", this.page, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponsePages) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
          }
          this.load();
        });
    }
  }

  autoSave() {
    if (this.page.Id !== "") {
      this.createSend();
    } else {
      this.editSend();
    }
  }

  createSend() {
    this.$axios.put("/" + this.$route.query.siteId + "/pages", this.page, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponsePages) => {
        if (response.data.success) {
          this.updateText(response.data.message);
          this.updateColor('green')
          this.updateShow(true);
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  editSend() {
    this.$axios.post("/" + this.$route.query.siteId + "/pages", this.page, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponsePages) => {
        if (response.data.success) {
          this.updateText(response.data.message);
          this.updateColor('green')
          this.updateShow(true);
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  onEditorBlur(quill: any) {
  }

  onEditorFocus(quill: any) {
  }

  onEditorReady(quill: any) {
  }

  onEditorChange(quill: any, html: any, text: any) {
    this.content = html
  }

  async load() {
    this.pages = [];
    this.$axios.get("/" + this.$route.query.siteId + "/pages")
      .then((response: IResponsePages) => {
        if (response.data.success) {
          for (let index in response.data.pages) {
            response.data.pages[index].Title = JSON.parse(response.data.pages[index].Title.toString());
            response.data.pages[index].Body = JSON.parse(response.data.pages[index].Body.toString());
            if (response.data.pages[index].MetaTitle !== "") {
              response.data.pages[index].MetaTitle = JSON.parse(response.data.pages[index].MetaTitle.toString());
            }
            if (response.data.pages[index].Keywords !== "") {
              response.data.pages[index].Keywords = JSON.parse(response.data.pages[index].Keywords.toString());
            }
            if (response.data.pages[index].Description !== "") {
              response.data.pages[index].Description = JSON.parse(response.data.pages[index].Description.toString());
            }
            this.pages.push(response.data.pages[index]);
          }
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  close() {
    this.dialog = false;
  }

  createClearTranslationObject() {
    const response: any = {};
    for (const lang of this.languages) {
      response[lang] = "";
    }
    return response
  }

  get editor() {
    return this.$refs.myQuillEditor.quill
  }

  async translate(item: Page) {
    const title = {
      Text: item.Title[this.defaultLanguage],
      Lang: this.language
    }

    const body = {
      Text: item.Body[this.defaultLanguage],
      Lang: this.language
    }

    await this.$axios.post("/deepl", title, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: any) => {
        if (response.data.success === true) {
          const data = JSON.parse(response.data.text);
          item.Title[this.language] = data.translations[0].text;
        }
      });

    await this.$axios.post("/deepl", body, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: any) => {
        if (response.data.success === true) {
          const data = JSON.parse(response.data.text);
          item.Body[this.language] = data.translations[0].text;
        }
      });

    this.page = item;
    await this.save(false);
  }
}
</script>

