<template>
  <v-data-table
    :headers="headers"
    :items="posts"
    :items-per-page="10"
    class="elevation-1"
  >
    <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>My Posts</v-toolbar-title>
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
            <v-img v-if="post.Id !== ''" :src="$config.hostname + '/storage/' + post.Src" height="auto" width="500px"></v-img>
            <v-row v-if="post.Id !== ''">
              <v-col cols="8">
                <v-file-input
                  accept="image/*"
                  label="Cover image"
                  v-model="file"
                ></v-file-input>
              </v-col>
              <v-col cols="4">
                 <v-btn
                    color="primary"
                    dark
                    class="mb-2"
                    @click="addCover"
                 >Add Cover</v-btn>
              </v-col>
            </v-row>
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
                        v-model="post.Title[language]"
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
                      <v-menu

                        v-model="fromDateMenu"
                        :close-on-content-click="false"
                        :nudge-right="40"
                        transition="scale-transition"
                        offset-y
                        max-width="290px"
                        min-width="290px"
                      >
                        <template v-slot:activator="{ on }">
                          <v-text-field
                            label="Published"
                            prepend-icon="event"
                            readonly
                            :value="fromDateVal"
                            v-on="on"
                          ></v-text-field>
                        </template>
                        <v-date-picker
                          locale="en-us"
                          v-model="fromDateVal"
                          no-title
                          @input="fromDateMenu = false"
                        ></v-date-picker>
                      </v-menu>
                    </v-col>
                    <v-col
                      cols="12"
                    >
                      <v-textarea
                        v-model="post.Excerpt[language]"
                        :counter="300"
                        :rules="rules"
                        label="Excerpt"
                        required
                      ></v-textarea>
                      <quill-editor
                        :ref="post.Body[language]"
                        v-model="post.Body[language]"
                        :options="editorOption"
                        @blur="onEditorBlur($event)"
                        @focus="onEditorFocus($event)"
                        @ready="onEditorReady($event)"
                      />
                      <v-text-field
                        v-model="post.MetaTitle[language]"
                        :counter="100"
                        :rules="rules"
                        label="Meta title"
                        required
                      ></v-text-field>
                      <v-text-field
                        v-model="post.Description[language]"
                        :counter="150"
                        :rules="rules"
                        label="Meta description"
                        required
                      ></v-text-field>
                      <v-text-field
                        v-model="post.Keywords[language]"
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
    <template v-slot:item.file="{ item }">
      <div class="p-5">
        <v-img :src="$config.hostname + '/storage/' + item.file" :alt="item.Title[language]" height="auto"
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
import {Component, Prop, Vue} from 'nuxt-property-decorator';
import Post from "~/model/Post";
import IResponsePosts from "~/model/IResponsePosts";
import IResponseFiles from '~/model/IResponseFiles';
import IHeader from '~/model/IHeader';
import {namespace} from 'vuex-class';

const snackbar = namespace('Snackbar');

@Component
export default class PostsPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void
  @Prop() readonly languages!: string[];
  @Prop() readonly language!: string;
  posts: Post[] = [];
  title: string = "";
  editorOption: any = {
    // Some Quill options...
  };
  $axios: any;
  content: any;
  $refs: any;
  headers: IHeader[] = [
    {
      text: "Image",
      align: 'start',
      sortable: false,
      value: 'File',
    },
    {text: "Title", value: 'Title'},
    {text: "Excerpt", value: 'Excerpt'},
    {text: "Published", value: 'Published'},
    {text: "Actions", value: 'actions', sortable: false}
  ];
  post: Post = {
    Body: this.createClearTranslationObject(),
    Title: this.createClearTranslationObject(),
    Excerpt: this.createClearTranslationObject(),
    Published: '',
    Id: '',
    MetaTitle: this.createClearTranslationObject(),
    Keywords: this.createClearTranslationObject(),
    Description: this.createClearTranslationObject()
  };
  dialog: boolean = false;
  valid: boolean = true;
  rules: any = [];
  fromDateMenu: boolean = false;
  fromDateVal: string = "";
  file: File = {
    lastModified: 0, name: "", webkitRelativePath: "",
    size: 0,
    type: '',
    arrayBuffer: function (): Promise<ArrayBuffer> {
      throw new Error('Function not implemented.');
    },
    slice: function (start?: number, end?: number, contentType?: string): Blob {
      throw new Error('Function not implemented.');
    },
    stream: function (): ReadableStream<any> {
      throw new Error('Function not implemented.');
    },
    text: function (): Promise<string> {
      throw new Error('Function not implemented.');
    }
  };

  mounted() {
    this.load();
  }

  edit(post: Post) {
    this.post = post;
    this.title = this.post.Title[this.language];
    this.fromDateVal = post.Published[0].toISOString().split('T')[0].toString();
    this.dialog = true;
  }

  remove(id: string) {
    this.$axios.delete("/" + this.$route.query.siteId + "/posts/" + id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponsePosts) => {
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
    this.post = {
      Body: this.createClearTranslationObject(),
      Title: this.createClearTranslationObject(),
      Excerpt: this.createClearTranslationObject(),
      Published: '',
      Id: '',
      MetaTitle: this.createClearTranslationObject(),
      Keywords: this.createClearTranslationObject(),
      Description: this.createClearTranslationObject()
    };
  }

  save(close: boolean) {
    this.post.Published = this.fromDateVal;
    this.post.Title = JSON.stringify(this.post.Title);
    this.post.Excerpt = JSON.stringify(this.post.Excerpt);
    this.post.Body = JSON.stringify(this.post.Body);
    this.post.MetaTitle = JSON.stringify(this.post.MetaTitle);
    this.post.Keywords = JSON.stringify(this.post.Keywords);
    this.post.Description = JSON.stringify(this.post.Description);
    if (this.post.Id !== "") {
      this.$axios.put("/" + this.$route.query.siteId + "/posts", this.post, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponsePosts) => {
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
      this.post.Id = "";
      this.$axios.post("/" + this.$route.query.siteId + "/posts", this.post, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponsePosts) => {
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
    if (this.post.Id !== "") {
      this.createSend();
    } else {
      this.editSend();
    }
  }

  createSend() {
    this.$axios.put("/" + this.$route.query.siteId + "/posts", this.post, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponsePosts) => {
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
    this.$axios.post("/" + this.$route.query.siteId + "/posts", this.post, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponsePosts) => {
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

  load() {
    this.posts = [];
    this.$axios.get("/" + this.$route.query.siteId + "/posts")
      .then((response: IResponsePosts) => {
        if (response.data.success) {
          for (let index in response.data.posts) {
            response.data.posts[index].Published = [new Date(response.data.posts[index].Published)];
            response.data.posts[index].Title = JSON.parse(response.data.posts[index].Title.toString());
            response.data.posts[index].Excerpt = JSON.parse(response.data.posts[index].Excerpt.toString());
            response.data.posts[index].Body = JSON.parse(response.data.posts[index].Body.toString());
            if (response.data.posts[index].MetaTitle !== "") {
              response.data.posts[index].MetaTitle = JSON.parse(response.data.posts[index].MetaTitle.toString());
            }
            if (response.data.posts[index].Keywords !== "") {
              response.data.posts[index].Keywords = JSON.parse(response.data.posts[index].Keywords.toString());
            }
            if (response.data.posts[index].Description !== "") {
              response.data.posts[index].Description = JSON.parse(response.data.posts[index].Description.toString());
            }
            this.posts.push(response.data.posts[index]);
          }
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  addCover() {
    let formData = new FormData();
    formData.append('file', this.file);
    this.$axios.post("/" + this.$route.query.siteId + '/files/upload', formData)
      .then((response: IResponseFiles) => {
        if (response.data.success) {
          this.post.File = response.data.file;
          this.post.Src = response.data.file;
          this.save(false);
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
}
</script>

