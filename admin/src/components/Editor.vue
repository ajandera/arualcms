<template>
  <div id="editor">
    <input type="text" class="form-control" v-model="post.title" />
    <hr>
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
    <div class="float-right mt-3">
        <button v-on:click="save" class="btn btn-lg btn-success">Save</button>
    </div>
  </div>
</template>

<script>

import axios from "axios";

import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

import { quillEditor } from 'vue-quill-editor'
import DatePicker from 'vue2-datepicker2'

export default {
  name: 'Editor',
  props: ['post'],
  components: {
    quillEditor,
    DatePicker
  },
  data() {
    return {
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
    setInterval(() => {
      //this.save();
    }, 5000);
  },
  methods: {
    save() {
      axios.post(this.$hostname + "post/" + this.post.id, this.post)
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
