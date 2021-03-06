<template>
<div class="table-wrap">
    <h1>Posts</h1>
    <hr>
    <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
    <table class="table table-stripped mt-3">
      <thead>
        <tr>
          <th>#</th>
          <th>Title</th>
          <th>Published</th>
          <th>
            <button v-on:click="create()" type="button" class="btn btn-secondary btn-success float-right">
              <i class="fa fa-plus"></i>Create</button>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="post in posts" :key="post.id" v-on:click="edit(post.id)" class="actRow">
          <td>{{ post.id }}</td>
          <td>{{ post.title }}</td>
          <td>{{ post.published | formatDate}}</td>
          <td class="text-right">
            <button v-on:click.stop.prevent="remove(post.id)" type="button" class="btn btn-secondary btn-danger">
              <i class="fa fa-trash"></i>Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <modal name="form" :width="800" height="auto" :scrollable="true">
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
              <div id="editor">
                <label>Title</label>
                <input type="text" class="form-control" v-model="post.title" />
                <hr>
                <label>Published</label><br>
                <date-picker v-model="post.published" :lang="lang" type="datetime" :time-picker-options="timePickerOptions"></date-picker>
                <hr>
                <!-- Two-way Data-Binding -->
                <quill-editor
                    ref="myQuillEditor"
                    v-model="post.body"
                    :options="editorOption"
                    @blur="onEditorBlur($event)"
                    @focus="onEditorFocus($event)"
                    @ready="onEditorReady($event)"
                />
                <label class="mt-4">Meta title</label>
                <input type="text" class="form-control" v-model="post.meta.title" />
                <hr>
                <label>Meta description</label>
                <input type="text" class="form-control" v-model="post.meta.description" />
                <hr>
                <label>Meta keywords</label>
                <input type="text" class="form-control" v-model="post.meta.keywords" />
                <hr>
                <div class="float-right mt-3">
                  <button v-on:click="save" class="btn btn-lg btn-success">Save</button>
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

import {quillEditor} from "vue-quill-editor";
import DatePicker from "vue2-datepicker2";

export default {
  name: 'Posts',
  components: {
    quillEditor,
    DatePicker
  },
  data() {
    return {
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user"),
      posts: [],
      post: {
        title: "",
        body: "",
        published: "",
        meta: {
          title: "",
          keywords: "",
          description: ""
        }
      },
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
    edit(id) {
      this.post = this.posts.filter(x => x.id === id)[0];
      this.modalTitle = this.post.title;
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
              this.loggedUser = false;
              window.location.reload();
            }
          });
    },
    create() {
      this.post = {
        title: "",
        body: "",
        published: "",
        meta: {
          title: "",
          keywords: "",
          description: ""
        }
      };
      this.show();
    },
    show() {
      this.$modal.show('form')
    },
    hide() {
      this.$modal.hide('form')
    },
    save() {
      if (this.post.id !== undefined) {
        axios.put(this.$hostname + "post/" + this.post.id, this.post, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
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
                    this.loggedUser = false;
                    window.location.reload();
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
                this.loggedUser = false;
                window.location.reload();
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
  max-width: 800px;
  margin: 10px;
}
.modal-content {
  border: none;
}
</style>
