<template>
  <div id="editor">
    <div class="row mt-3">
      <div class="col-12">
        <input type="text" class="form-control" v-model="post.title" />
        <!-- Two-way Data-Binding -->
        <quill-editor
          ref="myQuillEditor"
          v-model="content"
          :options="editorOption"
          @blur="onEditorBlur($event)"
          @focus="onEditorFocus($event)"
          @ready="onEditorReady($event)"
        />
      </div>
    </div>
    <div class="row mt-3">
      <div class="col-6 mt-3">
        <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
      </div>
      <div class="col-6 mt-3">
        <div class="float-right">
          <button v-on:click="save" class="btn btn-lg btn-success">Save</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import axios from "axios";

import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

import { quillEditor } from 'vue-quill-editor'

export default {
  name: 'Editor',
  components: {
    quillEditor
  },
  data() {
    return {
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user"),
      post: {},
      content: '<h2>I am Example</h2>',
      editorOption: {
        // Some Quill options...
      }
    }
  },
  mounted() {
    let localApp = window.localStorage.getItem("app");
    if (localApp) {
      this.code = JSON.parse(localApp);
    }
    setInterval(() => {
      this.save();
    }, 5000);
  },
  methods: {
    save() {
      if (this.loggedUser) {
        let id = window.localStorage.getItem("userId");
        axios.post("http://localhost:9000/api/v1/app/save", {"code": this.code, "user": id, "name": this.loggedUser})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }

              setTimeout(() => {
                this.message = null;
                this.messageClass = null;
              }, 1000);
            });
      } else {
        window.localStorage.setItem("app", JSON.stringify(this.code));
        this.message = "Application saved."
        this.messageClass = "alert alert-success";
        setTimeout(() => {
          this.message = null;
          this.messageClass = null;
        }, 1000);
      }
    },
    onEditorBlur(quill) {
      console.log('editor blur!', quill)
    },
    onEditorFocus(quill) {
      console.log('editor focus!', quill)
    },
    onEditorReady(quill) {
      console.log('editor ready!', quill)
    },
    onEditorChange({ quill, html, text }) {
      console.log('editor change!', quill, html, text)
      this.content = html
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

#editor {

}

</style>
