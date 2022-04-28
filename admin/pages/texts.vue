<template>
  <v-row justify="center" align="center">
   
  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';
import IResponse from '~/model/IResponse';
import Text from '~/model/Text';

@Component({
    components: {
        quillEditor
    }
})
export default class TextsPage extends Vue {
    texts: Text[];
    text?: Text; 
    editorOption: any = {};
    modalTitle: string = "";

    mounted() {
        this.load();
    }

    save(text: Text) {
      if (text._id !== undefined) {
        axios.put("text", text)
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.error;
                this.messageClass = 'danger';
              }
            });
      } else {
        axios.post("text", text)
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
                this.load();
              } else {
                this.message = response.data.error;
                this.messageClass = 'danger';
              }
            });
      }
    }

    load() {
      axios.get("text")
          .then(response => {
            if (response.data.success === true) {
              this.texts = response.data.texts;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }

    add() {
      let text = {};
      text.key = "";
      text.value = {};
      for (const lang of this.languages) {
        text.value[lang] = "";
      }
      this.texts.push(text);
    }

    remove(item) {
      axios.delete("text/" + item._id.$oid)
          .then(response => {
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

    openEditor(item) {
      this.text = item;
      this.modalTitle = item.key;
      this.show();
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
}
</script>

