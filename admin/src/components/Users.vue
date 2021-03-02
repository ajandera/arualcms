<template>
<div class="table-wrap">
    <h1>Users</h1>
    <hr>
    <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
    <table class="table table-stripped mt-3">
      <thead>
        <tr>
          <th>#</th>
          <th>Username</th>
          <th>
            <button v-on:click="create()" type="button" class="btn btn-secondary btn-success float-right">
              <i class="fa fa-plus"></i>Create</button>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="user in users" :key="user.id" v-on:click="edit(user.id)" class="actRow">
          <td>{{ user.id }}</td>
          <td>{{ user.username }}</td>
          <td class="text-right">
            <button v-on:click.stop.prevent="remove(user.id)" type="button" class="btn btn-secondary btn-danger">
              <i class="fa fa-trash"></i>Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <modal name="form" :width="800" height="auto" :scrollable="true">
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
  name: 'Users',
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
      axios.delete(this.$hostname + "user/" + id)
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
        axios.put(this.$hostname + "user/" + this.user.id, this.user)
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
            });
      } else {
        axios.post(this.$hostname + "user", this.user)
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
            });
      }
    },
    load() {
      this.users = [];
      axios.get(this.$hostname + "user")
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
