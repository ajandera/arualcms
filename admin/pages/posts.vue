<template>
  <v-row justify="center" align="center">

  </v-row>
</template>

<script lang="ts">
import {Component, Prop, Vue} from 'nuxt-property-decorator';
import Post from "~/model/Post";
import IResponsePosts from "~/model/IResponsePosts";
import Message from "~/model/Message";
import IResponseFiles from '~/model/IResponseFiles';

@Component({
    components: {
        quillEditor
    }
})
export default class PostsPage extends Vue {
    @Prop readonly languages!: string[];
    @Prop readonly language!: string;
    posts!: Post[];
    post: Post;
    modalTitle: string = "";
    editorOption: any = {
    // Some Quill options...
    };
    $axios: any;
    message: Message = {text: "", class: ""};
    content: any;
    $refs: any;
    file: any;

    mounted() {
        this.load();
    }

    edit(post: Post) {
      this.post = post;
      this.modalTitle = this.post.title[this.language];
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

    get editor() {
      return this.$refs.myQuillEditor.quill
    }
}
</script>

