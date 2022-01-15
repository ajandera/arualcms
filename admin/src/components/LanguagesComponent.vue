<template>
  <div id="texts">
    <div class="row">
      <div class="col-11">
        <h1>Languages</h1>
      </div>
    </div>
    <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
    <table class="table table-stripped mt-3">
      <thead>
      <tr>
        <th>Code</th>
        <th>Name</th>
        <th>Default</th>
        <th>
          <button v-on:click="create()" type="button" class="btn btn-secondary btn-success float-right">
            <font-awesome-icon icon="plus" />
          </button>
        </th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="language in languages" :key="language.key" v-on:click="edit(language)" class="actRow">
        <td>{{ language.key }}</td>
        <td>{{ language.value }}</td>
        <td>
          <div v-if="language.default === true"><font-awesome-icon icon="check" /></div>
        </td>
        <td class="text-right">
          <button v-on:click.stop.prevent="remove(language._id.$oid)" type="button" class="btn btn-secondary btn-danger">
            <font-awesome-icon icon="trash" />
          </button>
        </td>
      </tr>
      </tbody>
    </table>
    <modal name="form" height="auto" v-if="langObject !== null" :scrollable="true" :resizable="true" :adaptive="true">
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
                  <label>Code</label>
                  <input type="text" class="form-control" v-model="langObject.key" />

                  <label>Name</label>
                  <input type="text" class="form-control" v-model="langObject.value" />

                  <label>default</label>
                  <input type="checkbox" class="form-control" v-model="langObject.default" />

                  <div class="float-right mt-3">
                    <button v-on:click="save(langObject)" class="btn btn-lg btn-success">Save</button>
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
  name: 'LanguagesComponent',
  props: ['loggedUser'],
  components: {},
  data: function() {
    return {
      languages: [],
      messageClass: null,
      message: null,
      error: "",
      langObject: null
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    create() {
      this.modalTitle = "New Language";
      this.langObject = {
        key: "",
        value: "",
        default: false
      };
      this.show();
    },
    save(language) {
      if (language._id !== undefined) {
        axios.put(this.$hostname + "languages", language, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
                this.load();
                this.hide();
              } else {
                this.message = response.data.error;
                this.messageClass = 'danger';
                this.load();
                this.hide();
              }
            }, (error) => {
              if (error.response.status === 401) {
                window.localStorage.removeItem("userId");
                window.localStorage.removeItem("user");
                window.localStorage.removeItem("jwt");
                this.$router.push({name: 'posts'});
              }
            });
      } else {
        axios.post(this.$hostname + "languages", language, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.error;
                this.messageClass = 'danger';
              }
            }, (error) => {
              if (error.response.status === 401) {
                window.localStorage.removeItem("userId");
                window.localStorage.removeItem("user");
                window.localStorage.removeItem("jwt");
                this.$router.push({name: 'posts'});
              }
            });
      }
    },
    load() {
      axios.get(this.$hostname + "languages")
          .then(response => {
            if (response.data.success === true) {
              this.languages = response.data.languages;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    },
    edit(langObject) {
      this.modalTitle = langObject.value;
      this.langObject = langObject;
      this.show();
    },
    remove(id) {
      axios.delete(this.$hostname + "languages/" + id, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
          .then(response => {
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
              this.$router.push({name: 'posts'});
            }
          });
    },
    show() {
      this.$modal.show('form')
    },
    hide() {
      this.$modal.hide('form')
    },
  }
}
</script>

<style lang="css" scoped>
.modal-content {
  border: none;
}
.actRow {
  cursor: pointer;
}
</style>
