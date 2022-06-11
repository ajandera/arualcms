<template>
  <v-row justify="center" align="center">

  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';
import IResponseLanguage from '~/model/IResponseLanguage';
import Language from '~/model/Language';
import Message from "~/model/Message";

@Component
export default class LanguagesPage extends Vue {
    langObject?: Language;
    modalTitle: string = 'Create language'
    $axios: any;
    message: Message = {class: "", text: ""};
    languages: Language[] = [];

    mounted(): void {
        this.load();
    }

    create(): void {
      this.modalTitle = "New Language";
      this.langObject = {
        _id: "",
        key: "",
        value: "",
        default: false
      };
    }

    save(language: Language) {
      if (language._id !== '') {
        this.$axios.put("/languages", language)
            .then((response: IResponseLanguage) => {
              if (response.data.success) {
                this.message.text = response.data.message;
                this.message.class = "alert alert-success";
                this.load();
              } else {
                this.message.text = response.data.error;
                this.message.class = 'danger';
                this.load();
              }
            });
      } else {
        this.$axios.post("/languages", language)
            .then((response: IResponseLanguage) => {
              if (response.data.success) {
                this.message.text = response.data.message;
                this.message.class = "alert alert-success";
              } else {
                this.message.text = response.data.error;
                this.message.class = 'danger';
              }
            });
      }
    }

    load() {
      this.$axios.get("/languages")
          .then((response: IResponseLanguage) => {
            if (response.data.success) {
              this.languages = response.data.languages;
            } else {
              this.message.text = response.data.error;
              this.message.class = 'danger';
            }
          });
    }

    edit(langObject: Language) {
      this.modalTitle = langObject.value;
      this.langObject = langObject;
    }

    remove(id: string) {
      this.$axios.delete("/languages/" + id)
          .then((response: IResponseLanguage) => {
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
}
</script>

