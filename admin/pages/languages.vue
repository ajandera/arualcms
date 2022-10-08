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
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              color="primary"
              dark
              class="mb-2"
              v-bind="attrs"
              v-on="on"
            >
              Add
            </v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="text-h5">{{ title }}</span>
            </v-card-title>
            <v-card-text>
              <v-container>
                <v-text-field
                  v-model="langObject.Key"
                  :counter="5"
                  label="Code"
                  required
                ></v-text-field>
                <v-text-field
                  v-model="langObject.Name"
                  :counter="30"
                  label="Name"
                  required
                ></v-text-field>
                <v-checkbox v-model="langObject.Default" label="Default"/>
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
              <v-btn
                color="primary"
                dark
                @click="save(langObject)"
              >
                Save
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
        mdi-pencil
      </v-icon>
      <v-icon
        small
        class="mr-2"
        @click="remove(item)"
      >
        mdi-delete
      </v-icon>
    </template>
    <template v-slot:item.Default="{ item }">
      <v-icon
        small
        class="mr-2"
        v-if="item.Default === true"
      >
        mdi-check
      </v-icon>
    </template>
    <template v-slot:no-data>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import {Component, Vue, Watch} from 'nuxt-property-decorator';
import IResponseLanguage from '~/model/IResponseLanguage';
import Language from '~/model/Language';
import IHeader from "~/model/IHeader";
import {namespace} from 'vuex-class';

const snackbar = namespace('Snackbar');

@Component
export default class LanguagesPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void

  langObject: Language = {
    Key: "",
    Name: "",
    Default: false,
    Id: "",
    SiteId: this.$route.query.siteId.toString()
  };
  title: string = 'Create language'
  $axios: any;
  languages: Language[] = [];
  dialog: boolean = false;
  headers: IHeader[] = [
    {
      text: "Name",
      align: 'start',
      sortable: true,
      value: 'Name',
    },
    {text: "Code", value: 'Key', sortable: false},
    {text: "Default", value: 'Default', sortable: false},
    {text: "Actions", value: 'actions', sortable: false}
  ];

  mounted(): void {
    this.load();
  }

  @Watch('$route.query')
  onPropertyChanged(value: string, oldValue: string) {
    this.load();
  }

  create(): void {
    this.title = "New Language";
    this.dialog = true;
  }

  save(language: Language) {
    if (language.Id !== '') {
      this.$axios.put("/" + this.$route.query.siteId + "/languages", language, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseLanguage) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
            this.load();
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
            this.load();
          }
        });
    } else {
      this.$axios.post("/" + this.$route.query.siteId + "/languages", language, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseLanguage) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
            this.load();
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
            this.load();
          }
        });
    }
  }

  async load() {
    this.dialog = false;
    this.$axios.get("/" + this.$route.query.siteId + "/languages")
      .then((response: IResponseLanguage) => {
        if (response.data.success) {
          this.languages = response.data.languages;
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  edit(langObject: Language) {
    this.title = langObject.Name;
    this.langObject = langObject;
    this.dialog = true;
  }

  remove(language: Language) {
    this.$axios.delete("/" + this.$route.query.siteId + "/languages/" + language.Id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseLanguage) => {
        if (response.data.success) {
          this.updateText(response.data.message);
          this.updateColor('green')
          this.updateShow(true);
          this.load();
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  close() {
    this.langObject = {
      Id: "",
      Key: "",
      Name: "",
      Default: false,
      SiteId: this.$route.query.siteId.toString()
    };
    this.dialog = false;
  }
}
</script>

