<template>
  <v-data-table
    :headers="headers"
    :items="menus"
    :items-per-page="10"
    class="elevation-1"
  >
    <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>Menu</v-toolbar-title>
        <v-divider
          class="mx-4"
          inset
          vertical
        ></v-divider>
        <v-spacer></v-spacer>
        <v-dialog
          v-model="dialog"
          max-width="700px"
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
            <v-card-title>{{ title }}</v-card-title>
            <v-card-text>
              <v-container>
                <v-form
              ref="form"
              v-model="valid"
              lazy-validation>
                  <v-row>
                    <v-col
                      cols="12"
                      sm="6"
                      md="6"
                    >
                      <v-text-field
                        v-model="menu.Name[language]"
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
                        v-model="menu.Url"
                        :counter="50"
                        :rules="rules"
                        label="Url"
                        required
                      ></v-text-field>
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col
                      cols="6"
                    >
                    <v-select
                        v-model="menu.PageId"
                        :items="pages"
                        item-text="name"
                        item-value="id"
                        label="Page"
                        return-object
                        single-line
                      ></v-select>
                    </v-col>
                    <v-col
                      cols="6"
                    >
                      <v-select
                        v-model="menu.PostId"
                        :items="posts"
                        label="Post"
                        item-text="name"
                        item-value="id"
                        return-object
                        single-line
                      ></v-select>
                    </v-col>
                  </v-row>
              </v-form>
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
          </v-card>
        </v-dialog>
      </v-toolbar>
    </template>
    <template v-slot:item.Name="{ item }">
      <p>{{ item.Name[language] }}</p>
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
import Menu from "~/model/Menu";
import IResponseMenu from "~/model/IResponseMenu";
import IHeader from '~/model/IHeader';
import {namespace} from 'vuex-class';
import IResponsePages from '~/model/IResponsePages';
import IResponsePosts from '~/model/IResponsePosts';

const snackbar = namespace('Snackbar');

@Component
export default class MenuPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void

  @Prop() readonly languages!: string[];
  @Prop() readonly language!: string;
  @Prop() readonly defaultLanguage!: string;

  menus: Menu[] = [];
  title: string = "";
  $axios: any;
  $refs: any;
  headers: IHeader[] = [
    {text: "Name", value: 'Name'},
    {text: "Actions", value: 'actions', sortable: false}
  ];
  menu: Menu = {
    Id: "",
    Name: this.createClearTranslationObject(),
    Url: "",
    Children: [],
    PageId: "",
    PostId: "",
    ParentId: null
  };
  dialog: boolean = false;
  valid: boolean = true;
  rules: any = [];
  posts: any[] = [];
  pages: any[] = [];

  mounted() {
    this.load();
  }

  @Watch('$route.query')
  onPropertyChanged(value: string, oldValue: string) {
    this.load();
  }

  @Watch('language')
  onLangChanged(value: string, oldValue: string) {
    this.getPages();
    this.getPosts();
  }

  edit(menu: Menu) {
    this.menu = menu;
    this.title = this.menu.Name;
    this.dialog = true;
  }

  remove(id: string) {
    this.$axios.delete("/" + this.$route.query.siteId + "/menu/" + id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseMenu) => {
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
    this.title = "Add new menu item"
    this.menu = {
      Id: '',
      Name: this.createClearTranslationObject(),
      Url: '',
      Children: [],
      PageId: "",
      PostId: "",
      ParentId: null
    };
  }

  save() {
    this.menu.Name = JSON.stringify(this.menu.Name);
    if (this.menu.Id !== "") {
      this.$axios.put("/" + this.$route.query.siteId + "/menu", this.menu, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseMenu) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
          }
          this.close();
          this.load();
        });
    } else {
      this.menu.Id = "";
      this.$axios.post("/" + this.$route.query.siteId + "/menu", this.menu, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseMenu) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
          }
          this.close();
          this.load();
        });
    }
  }

  async load() {
    this.menus = [];
    this.$axios.get("/" + this.$route.query.siteId + "/menu")
      .then((response: IResponseMenu) => {
        this.menus = [];
        if (response.data.success) {
          for (let index in response.data.menu) {
            response.data.menu[index].Name = JSON.parse(response.data.menu[index].Name.toString());
            this.menus.push(response.data.menu[index]);
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

  async translate(item: Menu) {
    const name = {
      Text: item.Name[this.defaultLanguage],
      Lang: this.language
    }

    await this.$axios.post("/deepl", name, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: any) => {
        if (response.data.success === true) {
          const data = JSON.parse(response.data.text);
          item.Name[this.language] = data.translations[0].text;
        }
      });

    this.menu = item;
    await this.save();
  }

  createClearTranslationObject() {
    const response: any = {};
    for (const lang of this.languages) {
      response[lang] = "";
    }
    return response
  }

  getPages() {
    this.pages = [];
    this.$axios.get("/" + this.$route.query.siteId + "/pages")
      .then((response: IResponsePages) => {
        if (response.data.success) {
          for (let index in response.data.pages) {
            this.pages.push({
              "id": response.data.pages[index].Id,
              "name": JSON.parse(response.data.pages[index].Title.toString())[this.language]
            });
          }
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  getPosts() {
    this.posts = [];
    this.$axios.get("/" + this.$route.query.siteId + "/posts")
      .then((response: IResponsePosts) => {
        if (response.data.success) {
          for (let index in response.data.posts) {
            this.posts.push({
              "id": response.data.posts[index].Id,
              "name": JSON.parse(response.data.posts[index].Title.toString())[this.language]
            });
          }
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }
}
</script>

