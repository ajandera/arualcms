<template>
  <v-data-table
    :headers="headers"
    :items="languages"
    :items-per-page="50"
    class="elevation-1"
  >
    <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>Languages</v-toolbar-title>
        <v-divider
          class="mx-4"
          inset
          vertical
        ></v-divider>
        <v-spacer></v-spacer>
        <v-dialog
          v-model="dialog"
          max-width="1000px"
        >
          <v-card>
            <v-card-title>
              <span class="text-h5">{{ title }}</span>
            </v-card-title>
            <v-card-text>
              <v-container>

              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="blue darken-1"
                text
                @click="close"
              >
                Close
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-toolbar>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon
        small
        class="mr-2"
        @click="edit(item)"
      >
        mdi-user
      </v-icon>
    </template>
    <template v-slot:no-data>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';
import IResponseLanguage from '~/model/IResponseLanguage';
import Language from '~/model/Language';
import Message from "~/model/Message";
import IHeader from "~/model/IHeader";

@Component
export default class LanguagesPage extends Vue {
    langObject?: Language;
    title: string = 'Create language'
    $axios: any;
    message: Message = {class: "", text: ""};
    languages: Language[] = [];
    dialog: boolean = false;
    headers: IHeader[] = [
      {
        text: "Name",
        align: 'start',
        sortable: true,
        value: 'value',
      },
      {text: "Code", value: 'key', sortable: false},
      {text: "Default", value: 'default', sortable: false},
      {text: "Actions", value: 'actions', sortable: false}
    ];

    mounted(): void {
        this.load();
    }

    create(): void {
      this.title = "New Language";
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
      this.title = langObject.value;
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

    close() {

    }
}
</script>

