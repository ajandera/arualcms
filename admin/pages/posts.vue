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
            max-width="500px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                color="primary"
                dark
                class="mb-2"
                v-bind="attrs"
                v-on="on"
              >
                Add
              </v-btn>
            </template>
            <v-card>
              <v-form
                ref="form"
                v-model="valid"
                lazy-validation>
                <v-card-title>
                  <span class="text-h5">{{ title }}</span>
                </v-card-title>
                <v-card-text>
                  <v-container>
                    <v-row>
                      <v-col
                        cols="12"
                        sm="6"
                        md="6"
                      >
                        <v-text-field
                          v-model="editedItem.url"
                          :counter="50"
                          :rules="rules"
                          label="Test"
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
                    color="blue darken-1"
                    text
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
        <v-img :src="$config.hostname + '/storage/' + item.file" :alt="item.title[language]" height="auto" width="200px"></v-img>
      </div>
    </template>
    <template v-slot:item.title="{ item }">
      <p>{{ item.title[language] }}</p>
    </template>
    <template v-slot:item.excerpt="{ item }">
      <p>{{ item.excerpt[language] }}</p>
    </template>
    <template v-slot:item.published="{ item }">
      <p>{{ item.published.toLocaleDateString() }}</p>
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
        @click="remove(item)"
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
import Message from "~/model/Message";
import IResponseFiles from '~/model/IResponseFiles';
import IHeader from '~/model/IHeader';

@Component
export default class PostsPage extends Vue {
    @Prop() readonly languages!: string[];
    @Prop() readonly language!: string;
    posts: Post[] = [];
    post!: Post;
    title: string = "";
    editorOption: any = {
    // Some Quill options...
    };
    $axios: any;
    message: Message = {text: "", class: ""};
    content: any;
    $refs: any;
    file: any;
    headers: IHeader[] = [
      {
        text: "Image",
        align: 'start',
        sortable: false,
        value: 'file',
      },
      { text: "Title", value: 'title' },
      { text: "Excerpt", value: 'excerpt' },
      { text: "Published", value: 'published' },
      { text: "Actions", value: 'actions', sortable: false }
    ];
    editedIndex: number = -1;
    editedItem: Post = {
      body: {},
      title: {},
      excerpt: {},
      published: '',
      _id: '',
      meta: {
        title: {},
        keywords: {},
        description: {}
      }
    };
    dialog: boolean = false;
    valid: boolean = true;
    rules: any = [];

    mounted() {
        this.load();
    }

    edit(post: Post) {
      this.post = post;
      this.title = this.post.title[this.language];
    }

    remove(id: string) {
      this.$axios.delete("/post/" + id)
          .then((response: IResponsePosts) => {
            if (response.data.success) {
              this.message.text = response.data.message;
              this.message.class = "alert alert-success";
              this.load();
            } else {
              this.message.text = response.data.message;
              this.message.class = "alert alert-danger";
            }
          });
    }

    create() {
      this.post = {
        _id: '',
        title: {},
        excerpt: {},
        body: {},
        published: "",
        meta: {
          title: {},
          keywords: {},
          description: {}
        }
      };
      for (const lang of this.languages) {
        this.post.title[lang] = "";
        this.post.excerpt[lang] = "";
        this.post.body[lang] = "";
        this.post.meta.title[lang] = "";
        this.post.meta.keywords[lang] = "";
        this.post.meta.description[lang] = "";
      }
    }

    save(close: boolean) {
      if (this.post._id !== undefined) {
        this.$axios.put("/post", this.post)
            .then((response: IResponsePosts) => {
              if (response.data.success) {
                this.message.text = response.data.message;
                this.message.class = "alert alert-success";
              } else {
                this.message.text = response.data.message;
                this.message.class = "alert alert-danger";
              }
              this.load();
            });
      } else {
        this.$axios.post("/post", this.post)
            .then((response: IResponsePosts) => {
              if (response.data.success) {
                this.message.text = response.data.message;
                this.message.class = "alert alert-success";
              } else {
                this.message.text = response.data.message;
                this.message.class = "alert alert-danger";
              }
              this.load();
            });
      }
    }

    autoSave() {
      if (this.post._id !== undefined) {
        this.createSend();
      } else {
        this.editSend();
      }
    }

    createSend() {
      this.$axios.put("/post", this.post)
          .then((response: IResponsePosts) => {
            if (response.data.success) {
              this.message.text = response.data.message;
              this.message.class = "alert alert-success";
            } else {
              this.message.text = response.data.message;
              this.message.class = "alert alert-danger";
            }
          });
    }

    editSend() {
      this.$axios.post("/post", this.post)
          .then((response: IResponsePosts) => {
            if (response.data.success) {
              this.message.text = response.data.message;
              this.message.class = "alert alert-success";
            } else {
              this.message.text = response.data.message;
              this.message.class = "alert alert-danger";
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
      this.$axios.get("/post")
          .then((response: IResponsePosts) => {
            if (response.data.success) {
              for (let index in response.data.posts) {
                response.data.posts[index].published = new Date(response.data.posts[index].published);
                this.posts.push(response.data.posts[index]);
              }
            } else {
              this.message.text = response.data.error;
              this.message.class = 'danger';
            }
          });
    }

    handleFileUpload(){
      this.file = this.$refs.file.files[0];
    }

    upload() {
      let formData = new FormData();
      formData.append('file', this.file);
      this.$axios.post('/files/upload', formData)
        .then((response: IResponseFiles) => {
            if (response.data.success) {
                this.post.file = response.data.file;
                this.post.src = response.data.file;
                this.save(false);
            } else {
                this.message.text = response.data.message;
                this.message.class = "alert alert-danger";
            }
      });
    }

    close() {
      this.dialog = false;
    }

    get editor() {
      return this.$refs.myQuillEditor.quill
    }
}
</script>

