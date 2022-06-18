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
      </v-toolbar>
    </template>
    <template v-slot:item.src="{ item }">
      <div class="p-5">
        <v-img :src="item.src" :alt="item.name" height="auto" width="200px"></v-img>
      </div>
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
    file!: File;
    message: Message = {class: "", text: ""};
    $axios: any;
    $refs: any;
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
      this.$axios.delete("files/" + id)
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
              'Content-Type': 'multipart/form-data',
              'Authorization': "Bearer " + window.localStorage.getItem('jwt')
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

    handleFileUpload(){
      this.file = this.$refs.file.files[0];
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
          {'gallery': gallery}
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
}
</script>

