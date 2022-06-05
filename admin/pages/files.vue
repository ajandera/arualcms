<template>
  <v-row justify="center" align="center">

  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator';
import IResponseFiles from '~/model/IResponseFiles';
import Message from "~/model/Message";

@Component
export default class FilesPage extends Vue {
    files!: File[];
    file!: File;
    modalTitle: string =  "";
    message: Message = {class: "", text: ""};
    $axios: any;
    $refs: any;

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

