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
    <template v-slot:item.src="{ item }">
      <div class="p-5">
        <v-img :src="item.src" :alt="item.name" height="auto" width="200px"></v-img>
      </div>
    </template>
    <template v-slot:item.gallery="{ item }">
      <v-text-field
        v-model="item.gallery"
        :counter="20"
        v-on:change="saveGallery(item.id, item.gallery)"
      ></v-text-field>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon
        small
        class="mr-2"
        @click="remove(item)"
      >
        mdi-delete
      </v-icon>
    </template>
    <template v-slot:no-data>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';
import IResponseFiles from '~/model/IResponseFiles';
import Message from "~/model/Message";
import IHeader from "~/model/IHeader";

@Component
export default class FilesPage extends Vue {
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
    message: Message = {class: "", text: ""};
    $axios: any;
    $refs: any;
    dialog: boolean = false;
    headers: IHeader[] = [
      {
        text: "Image",
        align: 'start',
        sortable: true,
        value: 'src',
      },
      {text: "Name", value: 'name', sortable: true},
      {text: "Gallery", value: 'gallery', sortable: true},
      {text: "Actions", value: 'actions', sortable: false}
    ];

    mounted() {
        this.load();
    }

    remove(id:string) {
      this.$axios.delete("files/" + id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
          .then((response: IResponseFiles) => {
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

    create() {

    }

    save() {
      let formData = new FormData();
      formData.append('file', this.file.toString());
      this.$axios.post('files/upload',
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }
      ).then((response: IResponseFiles) => {
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

    load() {
      this.$axios.get("files")
          .then((response: IResponseFiles) => {
            if (response.data.success) {
              this.files = response.data.files;
            } else {
              this.message.text = response.data.error;
              this.message.class = 'danger';
            }
          });
    }

    saveGallery(id: string, gallery: string) {
      this.$axios.put( 'files/gallery/'+id,
        {'gallery': gallery},
        {headers: {'Content-Type': "application/json;charset=utf-8"}}
      ).then((response: IResponseFiles) => {
        if (response.data.success) {
          this.message.text = response.data.message;
          this.message.class = "alert alert-success";
        } else {
          this.message.text = response.data.message;
          this.message.class = "alert alert-danger";
        }
      });
    }

    close() {
      this.dialog = false;
    }
}
</script>

