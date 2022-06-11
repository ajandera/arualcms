<template>
  <v-row justify="center" align="center">

  </v-row>
</template>

<script lang="ts">
import {Component, Prop, Vue} from 'nuxt-property-decorator';
import Text from '~/model/Text';
import Message from "~/model/Message";
import IResponseTexts from "~/model/IResponseTexts";

@Component
export default class TextsPage extends Vue {
    @Prop() readonly languages!: string[];
    texts!: Text[];
    text?: Text;
    editorOption: any = {};
    modalTitle: string = "";
    message: Message = {class: "", text: ""};
    $axios: any;
    content: any;

    mounted() {
        this.load();
    }

    save(text: Text) {
      if (text._id.$oid !== '') {
        this.$axios.put("text", text)
            .then((response: IResponseTexts) => {
              if (response.data.success) {
                this.message.text = response.data.message;
                this.message.class = "alert alert-success";
              } else {
                this.message.text = response.data.error;
                this.message.class = 'danger';
              }
            });
      } else {
        this.$axios.post("text", text)
            .then((response: IResponseTexts) => {
              if (response.data.success) {
                this.message.text = response.data.message;
                this.message.class = "alert alert-success";
                this.load();
              } else {
                this.message.text = response.data.error;
                this.message.class = 'danger';
              }
            });
      }
    }

    load() {
      this.$axios.get("/text")
          .then((response: IResponseTexts) => {
            if (response.data.success === true) {
              this.texts = response.data.texts;
            } else {
              this.message.text = response.data.error;
              this.message.class = 'danger';
            }
          });
    }

    add() {
      let text: Text = {
        key: '',
        value: {},
        _id: {$oid: ""}
      };
      for (const lang of this.languages) {
        text.value[lang] = "";
      }
      this.texts.push(text);
    }

    remove(item: Text) {
      this.$axios.delete("/text/" + item._id.$oid)
          .then((response: IResponseTexts) => {
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

    openEditor(item: Text) {
      this.text = item;
      this.modalTitle = item.key;
      this.show();
    }

    show() {
        throw new Error('Method not implemented.');
    }

    onEditorBlur(quill: any) {
    }

    onEditorFocus(quill: any) {
    }
    onEditorReady(quill: any) {
    }

    onEditorChange(quill: any, html: any, text: any ) {
      this.content = html
    }
}
</script>

