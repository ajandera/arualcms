<template>
  <div id="texts">
    <div class="row">
      <div class="col-9">
        <h1>Texts</h1>
      </div>
      <div class="col-3 text-right">
        <div class="btn-group mt-3" role="group" aria-label="Basic example">
          <button
              v-on:click="setLanguage(lang)"
              v-bind:class="{'btn btn-default': lang !== language, 'btn btn-primary': lang === language}" v-for="(lang, index) in languages" v-bind:key="index">{{ lang }}</button>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <hr>
      </div>
    </div>
    <div v-for="(item, index) in texts" class="row mt-2" v-bind:key="index">
      <div class="col-sm-1 col-xs-12">
        <button v-if="index === Object.keys(texts).length - 1" v-on:click="add" class="btn btn-warning"><font-awesome-icon icon="plus" /></button>
      </div>
      <div class="col-sm-3 col-xs-6">
        <input type="text" v-model="item.key" class="form-control" v-on:change="save(item)">
      </div>
      <div class="col-sm-7 col-xs-6">
        <textarea v-model="item.value[language]" class="form-control" v-on:change="save(item)"></textarea>
      </div>
      <div class="col-sm-1 col-xs-12">
        <div class="float-right">
          <div class="btn-group" role="group" aria-label="Basic example">
            <button v-on:click="openEditor(item)" class="btn btn-success"><font-awesome-icon icon="edit" /></button>
            <button v-on:click="remove(item)" class="btn btn-danger"><font-awesome-icon icon="times" /></button>
          </div>
        </div>
      </div>
    </div>
    <div class="row mt-4">
      <div class="col-12">
        <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
      </div>
    </div>
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
                <div id="editor" v-if="text.value !== undefined">
                  <!-- Two-way Data-Binding -->
                  <quill-editor
                      ref="myQuillEditor"
                      v-model="text.value[language]"
                      :options="editorOption"
                      @blur="onEditorBlur($event)"
                      @focus="onEditorFocus($event)"
                      @ready="onEditorReady($event)"
                  />
                  <hr>
                  <div class="float-right mt-3">
                    <button v-on:click="save(text)" class="btn btn-lg btn-success">Save</button>
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

export default {
  name: 'TextsComponent',
  props: ['language', 'languages'],
  components: {
    quillEditor
  },
  data: function() {
    return {
      texts: [],
      text: {},
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user"),
      editorOption: {},
      modalTitle: "",
      error: ""
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    save(text) {
      if (text._id !== undefined) {
        axios.put(this.$hostname + "text", text, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.error;
                this.messageClass = 'danger';
              }
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
        axios.post(this.$hostname + "text", text, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
                this.load();
              } else {
                this.message = response.data.error;
                this.messageClass = 'danger';
              }
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
    load() {
      axios.get(this.$hostname + "text")
          .then(response => {
            if (response.data.success === true) {
              this.texts = response.data.texts;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    },
    add() {
      let text = {};
      text.key = "";
      text.value = {};
      for (const lang of this.languages) {
        text.value[lang] = "";
      }
      this.texts.push(text);
    },
    remove(item) {
      axios.delete(this.$hostname + "text/" + item._id.$oid, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
          .then(response => {
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
    openEditor(item) {
      this.text = item;
      this.modalTitle = item.key;
      this.show();
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
    show() {
      this.$modal.show('form')
    },
    hide() {
      this.$modal.hide('form')
    }
  }
}
</script>

<style lang="css" scoped>
.modal-dialog {
  max-width: 800px;
  margin: 10px;
}
.modal-content {
  border: none;
}
</style>
