<template>
  <v-row justify="center" align="center">
   
  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';
import Language from '~/model/Language';
import IResponse from '~/model/IResponse';

@Component
export default class LanguagesPage extends Vue {
    langObject?: Language;
    modalTitle: string = 'Create language'

    mounted(): void {
        this.load();
    }

    create(): void {
      this.modalTitle = "New Language";
      this.langObject = {
        key: "",
        value: "",
        default: false
      };
    }

    save(language: Language) {
      if (language._id !== undefined) {
        this.$axios.put("languages", language)
            .then((response: IResponse) => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
                this.load();
              } else {
                this.message = response.data.error;
                this.messageClass = 'danger';
                this.load();
              }
            });
      } else {
        this.$axios.post("languages", language)
            .then((response: IResponse) => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.error;
                this.messageClass = 'danger';
              }
            });
      }
    }

    load() {
      axios.get("languages")
          .then((response: IResponse) => {
            if (response.data.success === true) {
              this.languages = response.data.languages;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }

    edit(langObject: Language) {
      this.modalTitle = langObject.value;
      this.langObject = langObject;
    }

    remove(id: string) {
      axios.delete("languages/" + id)
          .then((response: IResponse) => {
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
}
</script>

