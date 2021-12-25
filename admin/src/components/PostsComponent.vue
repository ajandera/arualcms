<template>
<div class="table-wrap">
    <div class="row">
      <div class="col-9">
        <h1>Posts</h1>
      </div>
    </div>
    <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
    <table class="table table-stripped mt-3">
      <thead>
        <tr>
          <td></td>
          <th>Title</th>
          <th class="d-none d-sm-table-cell"><br>Excerpt</th>
          <th>Published</th>
          <th>
            <button v-on:click="create()" type="button" class="btn btn-secondary btn-success float-right">
              <font-awesome-icon icon="plus" /> Create</button>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="post in posts" :key="post.id" v-on:click="edit(post)" class="actRow">
          <td style="width:150px;"><img v-bind:src="post.src" class="small-cover"></td>
          <td>{{ post.title[language] }}</td>
          <td class="d-none d-sm-table-cell">{{ post.excerpt[language] }}</td>
          <td>{{ post.published | formatDate}}</td>
          <td class="text-right">
            <button v-on:click.stop.prevent="remove(post._id.$oid)" type="button" class="btn btn-secondary btn-danger">
              <font-awesome-icon icon="trash" /> Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <modal v-if="post !== null" name="form" height="auto"
           :scrollable="true" :resizable="true" :adaptive="false"
           :width="1000" >
      <div class="modal-dialog">
      <!-- Modal content-->
        <div class="modal-content">
        <div class="modal-header">
          <h3 class="modal-title">{{ modalTitle }}</h3>
          <button @click="hide" type="button" class="close" data-dismiss="modal">×</button>
        </div>
        <div class="modal-body">
          <div class="row">
            <div class="col-12">
              <div v-if="error" class="danger">{{ error }}</div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
                <img v-bind:src="post.src" class="img-fluid" /><br>
                <input type="file" id="file" ref="file" v-on:change="handleFileUpload()"/>
                <div class="float-right mt-3">
                  <button v-on:click="upload()" class="btn btn-lg btn-success">Add Cover</button>
                </div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
              <div id="editor">
                <label>Title</label>
                <input type="text" class="form-control" v-model="post.title[language]" />
                <hr>
                <label>Published</label><br>
                <date-picker v-model="post.published" :lang="lang" type="datetime" :time-picker-options="timePickerOptions"></date-picker>
                <hr>
                <label>Excerpt</label>
                <textarea class="form-control" v-model="post.excerpt[language]"></textarea>
                <hr>
                <!-- Two-way Data-Binding -->
                <quill-editor
                    ref="myQuillEditor"
                    v-model="post.body[language]"
                    :options="editorOption"
                    @blur="onEditorBlur($event)"
                    @focus="onEditorFocus($event)"
                    @ready="onEditorReady($event)"
                />
                <label class="mt-4">Meta title</label>
                <input type="text" class="form-control" v-model="post.meta.title[language]" />
                <hr>
                <label>Meta description</label>
                <input type="text" class="form-control" v-model="post.meta.description[language]" />
                <hr>
                <label>Meta keywords</label>
                <input type="text" class="form-control" v-model="post.meta.keywords[language]" />
                <hr>
                <div class="float-right mt-3">
                  <button v-on:click="save(true)" class="btn btn-lg btn-success">Save</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </modal>
</div>
</template>

<script>

import axios from "axios";

import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

import { quillEditor } from "vue-quill-editor";
import DatePicker from "vue2-datepicker2";

export default {
  name: 'PostsComponent',
  props: ['language', 'languages', 'loggedUser'],
  components: {
    quillEditor,
    DatePicker
  },
  data() {
    return {
      messageClass: null,
      message: null,
      posts: [],
      post: null,
      modalTitle: "",
      error: "",
      editorOption: {
        // Some Quill options...
      },
      // custom lang
      lang: {
        days: ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'],
        months: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
        pickers: ['next 7 days', 'next 30 days', 'previous 7 days', 'previous 30 days'],
        placeholder: {
          date: 'Select Date',
          dateRange: 'Select Date Range'
        }
      },
      timePickerOptions:{
        start: '00:00',
        step: '00:30',
        end: '23:30'
      }
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    edit(post) {
      this.post = post;
      this.modalTitle = this.post.title[this.language];
      this.show();
    },
    remove(id) {
      axios.delete(this.$hostname + "post/" + id, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}}
        )
          .then((response) => {
            if (response.data.success) {
              this.message = response.data.message;
              this.messageClass = "alert alert-success";
              this.load();
            } else {
              this.message = response.data.message;
              this.messageClass = "alert alert-danger";
            }

            this.hide();
            setTimeout(() => {
              this.message = null;
              this.messageClass = null;
            }, 2000);
          }, (error) => {
            if (error.response.status === 401) {
              window.localStorage.removeItem("userId");
              window.localStorage.removeItem("user");
              window.localStorage.removeItem("jwt");
              this.$router.push({name: 'posts'});
            }
          });
    },
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
      this.show();
    },
    show() {
      this.$modal.show('form')
    },
    hide() {
      this.$modal.hide('form')
    },
    save(close) {
      if (this.post._id !== undefined) {
        axios.put(this.$hostname + "post", this.post, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }
              this.load();
              if (close === true) {
                this.hide();
              }
              setTimeout(() => {
                this.message = null;
                this.messageClass = null;
              }, 2000);
            }, (error) => {
                  if (error.response.status === 401) {
                    window.localStorage.removeItem("userId");
                    window.localStorage.removeItem("user");
                    window.localStorage.removeItem("jwt");
                    this.$router.push({name: 'posts'});
                  }
                });
      } else {
        axios.post(this.$hostname + "post", this.post, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }

              this.load();
              this.hide();
              setTimeout(() => {
                this.message = null;
                this.messageClass = null;
              }, 2000);
            }, (error) => {
              if (error.response.status === 401) {
                window.localStorage.removeItem("userId");
                window.localStorage.removeItem("user");
                window.localStorage.removeItem("jwt");
                this.$router.push({name: 'posts'});
              }
            });
      }
    },
    onEditorBlur(quill) {

    },
    onEditorFocus(quill) {

    },
    onEditorReady(quill) {

    },
    onEditorChange({ quill, html, text }) {
      this.content = html
    },
    load() {
      this.posts = [];
      axios.get(this.$hostname + "post")
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
          }, (err) => {
            console.log(err);
      });
    },
    handleFileUpload(){
      this.file = this.$refs.file.files[0];
    },
    upload() {
      let formData = new FormData();
      formData.append('file', this.file);
      axios.post( this.$hostname + 'files/upload',
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data',
              'Authorization': "Bearer " + window.localStorage.getItem('jwt')
            }
          }
      ).then(response => {
        if (response.data.success) {
          this.post.file = response.data.file.name;
          this.post.src = response.data.file.link;
          this.save(false);
        } else {
          this.message = response.data.message;
          this.messageClass = "alert alert-danger";
        }

        setTimeout(() => {
          this.message = null;
          this.messageClass = null;
        }, 2000);
      }, (error) => {
        if (error.response.status === 401) {
          window.localStorage.removeItem("userId");
          window.localStorage.removeItem("user");
          window.localStorage.removeItem("jwt");
          this.$router.push({name: 'posts'});
        }
      });
    }
  },
  computed: {
    editor() {
      return this.$refs.myQuillEditor.quill
    }
  }
}
</script>

<style lang="css" scoped>
.small-cover {
  width: 150px;
  height: auto;
}
.table-wrap {
  width: 100%;
}
.actRow {
  cursor: pointer;
}
.actRow:hover td {
  background: #f8f8f8;
}
.modal-dialog {
  max-width: 980px;
  margin: 10px;
}
.modal-content {
  border: none;
}
</style>
