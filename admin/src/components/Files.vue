<template>
<div class="table-wrap">
    <h1>Files</h1>
    <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
    <table class="table table-stripped mt-3">
      <thead>
        <tr>
          <th>#</th>
          <th>Image</th>
          <th>Name</th>
          <th>Gallery</th>
          <th>
            <button v-on:click="create()" type="button" class="btn btn-secondary btn-success float-right">
              <font-awesome-icon icon="plus" /> Create</button>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="file in files" :key="file.id" class="actRow">
          <td>{{ file.id }}</td>
          <td><img v-bind:src="file.src" class="img-thumbnail"></td>
          <td>{{ file.name }}</td>
          <td>
            <input v-model="file.gallery" type="text" v-on:change="saveGallery(file._id.$oid, file.gallery)" class="form-control" placeholder="Gallery">
          </td>
          <td class="text-right">
            <button v-on:click.stop.prevent="remove(file._id.$oid)" type="button" class="btn btn-secondary btn-danger">
              <i class="fa fa-trash"></i>Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <modal name="form" height="auto" :scrollable="true" :resizable="true" :adaptive="true">
    <div class="modal-dialog">
      <!-- Modal content-->
      <div class="modal-content">
        <div class="modal-header">
          <h3 class="modal-title">{{ modalTitle }}</h3>
          <button @click="hide" type="button" class="close" data-dismiss="modal">×</button>
        </div>
        <div class="modal-body">
          <div class="row">
            <div class="col-12">
              <div v-if="error" class="danger">{{ error }}</div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
              <div id="editor">
                <label>File
                  <input type="file" id="file" ref="file" v-on:change="handleFileUpload()"/>
                </label>
                <div class="float-right mt-3">
                  <button v-on:click="save()" class="btn btn-lg btn-success">Save</button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </modal>
</div>
</template>

<script>

import axios from "axios";

export default {
  name: 'Files',
  components: {},
  data() {
    return {
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user"),
      files: [],
      file: "",
      modalTitle: "",
      error: ""
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    remove(id) {
      axios.delete(this.$hostname + "files/" + id, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
          .then(response => {
            if (response.data.success) {
              this.message = response.data.message;
              this.messageClass = "alert alert-success";
              this.load();
            } else {
              this.message = response.data.message;
              this.messageClass = "alert alert-danger";
            }
            setTimeout(() => {
              this.message = null;
              this.messageClass = null;
            }, 2000);
          }, (error) => {
            if (error.response.status === 401) {
              window.localStorage.removeItem("userId");
              window.localStorage.removeItem("user");
              window.localStorage.removeItem("jwt");
              this.loggedUser = false;
              window.location.reload();
            }
          });
    },
    create() {
      this.show();
    },
    show() {
      this.$modal.show('form')
    },
    hide() {
      this.$modal.hide('form')
    },
    save() {
      let formData = new FormData();
      formData.append('file', this.file);
      axios.post( this.$hostname + 'files/upload',
          formData,
          {
            headers: {
              'Content-Type': 'multipart/form-data',
              'Authorization': "Bearer " + window.localStorage.getItem('jwt')
            }
          }
      ).then(response => {
        if (response.data.success) {
          this.message = response.data.message;
          this.messageClass = "alert alert-success";
          this.load();
        } else {
          this.message = response.data.message;
          this.messageClass = "alert alert-danger";
        }

        this.hide();
        setTimeout(() => {
          this.message = null;
          this.messageClass = null;
        }, 2000);
      }, (error) => {
        if (error.response.status === 401) {
          window.localStorage.removeItem("userId");
          window.localStorage.removeItem("user");
          window.localStorage.removeItem("jwt");
          this.loggedUser = false;
          window.location.reload();
        }
      });
    },
    handleFileUpload(){
      this.file = this.$refs.file.files[0];
    },
    load() {
      axios.get(this.$hostname + "files")
          .then(response => {
            if (response.data.success === true) {
              this.files = response.data.files;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    },
    saveGallery(id, gallery) {
      axios.put( this.$hostname + 'files/gallery/'+id,
          {'gallery': gallery},
          {
            headers: {
              'Authorization': "Bearer " + window.localStorage.getItem('jwt')
            }
          }
      ).then(response => {
        if (response.data.success) {
          this.message = response.data.message;
          this.messageClass = "alert alert-success";
        } else {
          this.message = response.data.message;
          this.messageClass = "alert alert-danger";
        }

        setTimeout(() => {
          this.message = null;
          this.messageClass = null;
        }, 2000);
      }, (error) => {
        if (error.response.status === 401) {
          window.localStorage.removeItem("userId");
          window.localStorage.removeItem("user");
          window.localStorage.removeItem("jwt");
          this.loggedUser = false;
          window.location.reload();
        }
      });
    }
  }
}
</script>

<style lang="css" scoped>
.table-wrap {
  width: 100%;
}
.actRow {
  cursor: pointer;
}
.actRow:hover td {
  background: #f8f8f8;
}
.modal-dialog {
  max-width: 800px;
  margin: 10px;
}
.modal-content {
  border: none;
}
.img-thumbnail {
  width: 200px;
}
</style>
