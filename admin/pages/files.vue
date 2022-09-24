<template>
  <v-data-table
    :headers="headers"
    :items="files"
    :items-per-page="50"
    class="elevation-1"
  >
    <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>Files</v-toolbar-title>
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
              <span class="text-h5">Upload file</span>
            </v-card-title>
            <v-card-text>
              <v-container>
                <v-file-input
                  accept="image/*"
                  label="File input"
                  v-model="file"
                ></v-file-input>
              </v-container>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="blue darken-1"
                @click="save"
              >
                Save
              </v-btn>
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
    <template v-slot:item.Src="{ item }">
      <div class="p-5">
        <v-img :src="$config.storage + item.Src" :alt="item.Name" height="auto" width="200px"></v-img>
      </div>
    </template>
    <template v-slot:item.Gallery="{ item }">
      <v-text-field
        v-model="item.Gallery"
        :counter="20"
        v-on:change="saveGallery(item.Id, item.Gallery)"
      ></v-text-field>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon
        small
        class="mr-2"
        @click="remove(item.Id)"
      >
        mdi-delete
      </v-icon>
    </template>
    <template v-slot:no-data>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import {Component, Vue, Watch} from 'nuxt-property-decorator';
import IResponseFiles from '~/model/IResponseFiles';
import IHeader from "~/model/IHeader";
import {namespace} from 'vuex-class';

const snackbar = namespace('Snackbar');

@Component
export default class FilesPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void
  files: File[] = [];
  file: File = {
    lastModified: 0, name: "", webkitRelativePath: "",
    size: 0,
    type: '',
    arrayBuffer: function (): Promise<ArrayBuffer> {
      throw new Error('Function not implemented.');
    },
    slice: function (start?: number, end?: number, contentType?: string): Blob {
      throw new Error('Function not implemented.');
    },
    stream: function (): ReadableStream<any> {
      throw new Error('Function not implemented.');
    },
    text: function (): Promise<string> {
      throw new Error('Function not implemented.');
    }
  };
  $axios: any;
  $refs: any;
  dialog: boolean = false;
  headers: IHeader[] = [
    {
      text: "Image",
      align: 'start',
      sortable: true,
      value: 'Src',
    },
    {text: "Name", value: 'Name', sortable: true},
    {text: "Gallery", value: 'Gallery', sortable: true},
    {text: "Actions", value: 'actions', sortable: false}
  ];

  mounted() {
    this.load();
  }

  @Watch('$route.query')
  onPropertyChanged(value: string, oldValue: string) {
    this.load();
  }

  remove(id: string) {
    this.$axios.delete("/" + this.$route.query.siteId + "files/" + id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseFiles) => {
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

  create() {

  }

  save() {
    let formData = new FormData();
    formData.append('file', this.file);
    this.$axios.post('/' + this.$route.query.siteId + '/files/upload',
      formData,
      {
        headers: {
          'Content-Type': 'multipart/form-data'
        }
      }
    ).then((response: IResponseFiles) => {
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

  async load() {
    this.$axios.get("/" + this.$route.query.siteId + "/files")
      .then((response: IResponseFiles) => {
        if (response.data.success) {
          this.files = response.data.files !== null ? response.data.files : [];
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  saveGallery(id: string, gallery: string) {
    this.$axios.put('/' + this.$route.query.siteId + '/files/' + id,
      {'gallery': gallery},
      {headers: {'Content-Type': "application/json;charset=utf-8"}}
    ).then((response: IResponseFiles) => {
      if (response.data.success) {
        this.updateText(response.data.message);
        this.updateColor('green')
        this.updateShow(true);
      } else {
        this.updateText(response.data.message);
        this.updateColor('red')
        this.updateShow(true);
      }
    });
  }

  close() {
    this.dialog = false;
  }
}
</script>

