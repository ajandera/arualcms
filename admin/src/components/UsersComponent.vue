<template>
<div class="table-wrap">
    <h1>Users</h1>
    <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
    <table class="table table-stripped mt-3">
      <thead>
        <tr>
          <th>Username</th>
          <th>
            <button v-on:click="create()" type="button" class="btn btn-secondary btn-success float-right">
              <font-awesome-icon icon="plus" /> Create</button>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id">
          <td>{{ user.username }}</td>
          <td class="text-right">
            <button v-on:click.stop.prevent="remove(user._id.$oid)" type="button" class="btn btn-secondary btn-danger">
              <font-awesome-icon icon="trash" /> Delete</button>
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
                <label>Username</label>
                <input type="email" class="form-control" v-model="user.username" :disabled="this.user.id !== undefined"/>
                <hr>
                <label>Password</label><br>
                <input type="password" class="form-control" v-model="user.password" />
                <div class="float-right mt-3">
                  <button v-on:click="save" class="btn btn-lg btn-success">Save</button>
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
  name: 'UsersComponent',
  props: ['language', 'languages'],
  components: {},
  data() {
    return {
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user"),
      users: [],
      user: {
        username: "",
        password: ""
      },
      modalTitle: "",
      error: ""
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    edit(id) {
      this.user = this.users.filter(x => x.id === id)[0];
      this.modalTitle = this.user.username;
      this.show();
    },
    remove(id) {
      axios.delete(this.$hostname + "users/" + id, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
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
              this.loggedUser = false;
              window.location.reload();
            }
          });
    },
    create() {
      this.user = {
        username: "",
        password: ""
      };
      this.show();
    },
    show() {
      this.$modal.show('form')
    },
    hide() {
      this.$modal.hide('form')
    },
    save() {
      if (this.user.id !== undefined) {
        axios.put(
            this.$hostname + "users/",
            this.user,
            {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}}
            ).then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }
              this.load();
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
      } else {
        axios.post(this.$hostname + "users", this.user, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }

              this.load();
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
      }
    },
    load() {
      this.users = [];
      axios.get(this.$hostname + "users")
          .then(response => {
            if (response.data.success === true) {
              this.users = response.data.users;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }
  },
  computed: {
    editor() {
      return this.$refs.myQuillEditor.quill
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
</style>
