<template>
  <div id="texts">
    <div class="row">
      <div class="col-11">
        <h1>Texts</h1>
        <hr>
      </div>
    </div>
    <div v-for="(item, index) in texts" class="row mt-2" v-bind:key="item.id">
      <div class="col-1">
        <button v-if="index === Object.keys(texts).length - 1" v-on:click="add" class="btn btn-warning"><font-awesome-icon icon="plus" /></button>
      </div>
      <div class="col-3">
        <input type="text" v-model="item.key" class="form-control">
      </div>
      <div class="col-6">
        <textarea v-model="item.value" class="form-control"></textarea>
      </div>
      <div class="col-1">
        <div class="float-right">
          <button v-on:click="openEditor(item)" class="btn btn-success"><font-awesome-icon icon="edit" /></button>
          <button v-on:click="remove(item)" class="btn btn-danger"><font-awesome-icon icon="times" /></button>
        </div>
      </div>
    </div>
    <div class="row mt-4">
      <div class="col-6">
        <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
      </div>
      <div class="col-5">
        <div class="float-right">
          <button v-on:click="save" class="btn btn-success">Save</button>
        </div>
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
                <div id="editor">
                  <!-- Two-way Data-Binding -->
                  <quill-editor
                      ref="myQuillEditor"
                      v-model="text.value"
                      :options="editorOption"
                      @blur="onEditorBlur($event)"
                      @focus="onEditorFocus($event)"
                      @ready="onEditorReady($event)"
                  />
                  <hr>
                  <div class="float-right mt-3">
                    <button v-on:click="hide" class="btn btn-lg btn-default">Close</button>
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
  name: 'Texts',
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
    save() {
      axios.put(this.$hostname + "text", this.texts, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
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
      this.texts.push({'key': '', 'value': ''});
    },
    remove(item) {
      var index = this.texts.indexOf(item);
      this.texts.splice(1, index);
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
