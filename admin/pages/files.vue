<template>
  <v-row justify="center" align="center">
   
  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import File from '~/model/File';
import IResponse from '~/model/IResponse';

@Component
export default class FilesPage extends Vue {
    files: File[];
    file: File;
    modalTitle: string =  "";

    mounted() {
        this.load();
    }
  
    remove(id:string) {
      this.$axios.delete("files/" + id)
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

    create() {

    }

    save() {
      let formData = new FormData();
      formData.append('file', this.file);
      this.$axios.post('files/upload',
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data',
              'Authorization': "Bearer " + window.localStorage.getItem('jwt')
            }
          }
      ).then((response: IResponse) => {
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

    handleFileUpload(){
      this.file = this.$refs.file.files[0];
    }

    load() {
      this.$axios.get("files")
          .then(response => {
            if (response.data.success === true) {
              this.files = response.data.files;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }

    saveGallery(id: string, gallery: string) {
      this.$axios.put( 'files/gallery/'+id,
          {'gallery': gallery}
      ).then(response => {
        if (response.data.success) {
          this.message = response.data.message;
          this.messageClass = "alert alert-success";
        } else {
          this.message = response.data.message;
          this.messageClass = "alert alert-danger";
        }
      });
    }
}
</script>

