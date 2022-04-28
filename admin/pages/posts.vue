<template>
  <v-row justify="center" align="center">
   
  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Post from '~/model/Post';
import IResponse from '~/model/IResponse';

@Component({
    components: {
        quillEditor
    }
})
export default class PostsPage extends Vue {
    posts: Post[];
    post: Post;
    modalTitle: string = "";
    editorOption: any = {
    // Some Quill options...
    };

    mounted() {
        this.load();
    }

    edit(post: Post) {
      this.post = post;
      this.modalTitle = this.post.title[this.language];
    }

    remove(id: string) {
      axios.delete("post/" + id)
          .then((response) => {
            if (response.data.success) {
              this.message = response.data.message;
              this.messageClass = "alert alert-success";
              this.load();
            } else {
              this.message = response.data.message;
              this.messageClass = "alert alert-danger";
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
        this.$axios.put("post", this.post)
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }
              this.load();
            });
      } else {
        this.$axios.post("post", this.post)
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
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
      this.$axios.put("post", this.post)
          .then(response => {
            if (response.data.success) {
              this.message = response.data.message;
              this.messageClass = "alert alert-success";
            } else {
              this.message = response.data.message;
              this.messageClass = "alert alert-danger";
            }
          });
    }

    editSend() {
      this.$axios.post("post", this.post)
          .then(response => {
            if (response.data.success) {
              this.message = response.data.message;
              this.messageClass = "alert alert-success";
            } else {
              this.message = response.data.message;
              this.messageClass = "alert alert-danger";
            }
          });
    }

    onEditorBlur(quill) {
    }

    onEditorFocus(quill) {
    }

    onEditorReady(quill) {
    }

    onEditorChange({ quill, html, text }) {
      this.content = html
    }

    load() {
      this.posts = [];
      this.$axios.get("post")
          .then((response) => {
            if (response.data.success === true) {
              for (let index in response.data.posts) {
                response.data.posts[index].published = new Date(response.data.posts[index].published);
                this.posts.push(response.data.posts[index]);
              }
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }

    handleFileUpload(){
      this.file = this.$refs.file.files[0];
    }

    upload() {
      let formData = new FormData();
      formData.append('file', this.file);
      this.$axios.post( this.$hostname + 'files/upload',formData)
        .then(response => {
            if (response.data.success) {
                this.post.file = response.data.file;
                this.post.src = response.data.file;
                this.save(false);
            } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
            }
      });
    }

    get editor() {
      return this.$refs.myQuillEditor.quill
    }
}
</script>

