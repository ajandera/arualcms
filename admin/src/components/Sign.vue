<template>
  <div id="sign">
    <button v-if="!loggedUser" @click="show" class="btn btn-outline-success" type="submit">Sign In / Sign Up</button>
    <table v-if="loggedUser">
      <tr>
        <td>{{ loggedUser }}</td>
        <td><button v-if="loggedUser" @click="logout" class="btn btn-outline-danger" type="submit">Sign out</button></td>
      </tr>
    </table>
    <modal name="sign" :width="600" :height="530" :adaptive="true">
      <div class="modal-dialog">
        <!-- Modal content-->
        <div class="modal-content">
          <div class="modal-header">
            <h3 class="modal-title">Sign in / Register</h3>
            <button @click="hide" type="button" class="close" data-dismiss="modal">×</button>
          </div>
          <div class="modal-body">
            <div class="row">
              <div class="col-12">
                <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
              </div>
            </div>
            <div class="row">
              <div class="col-6 borderRight">
                <label for="username">Username</label>
                <input type="text" v-model="username" class="form-control" id="username">
                <label for="username">Password</label>
                <input type="password" v-model="password" class="form-control" id="password">
                <button class="float-right btn btn-success mt-3" v-on:click="login">Sign In</button>
              </div>
              <div class="col-6">
                <p class="mb-3">For using database storing create your account.</p>
                <label for="usernameNewUser">Email</label>
                <input type="email" v-model="usernameNewUser" class="form-control" id="usernameNewUser" required>
                <label for="passwordNewUser">Password</label>
                <input type="password" v-model="passwordNewUser" class="form-control" id="passwordNewUser" required>
                <label for="passwordCheckNewUser">Password for check </label>
                <input type="password" v-model="passwordCheckNewUser" class="form-control" id="passwordCheckNewUser"
                       required>
                <div v-if="error" class="mt-1 alert alert-danger">{{ error }}</div>
                <button :disabled="this.passwordNewUser !== this.passwordCheckNewUser"
                        class="float-right btn btn-success mt-3" v-on:click="register">Sign Up
                </button>
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
  name: 'Sign',
  components: {},
  data() {
    return {
      message: null,
      messageClass: null,
      username: null,
      password: null,
      usernameNewUser: null,
      passwordNewUser: null,
      passwordCheckNewUser: null,
      error: null,
      loggedUser: window.localStorage.getItem("user")
    }
  },
  watch: {
    passwordCheckNewUser: function (val) {
      if (val !== this.passwordNewUser) {
        this.error = "Password mismatch!";
      } else {
        this.error = null;
      }
    },
  },
  methods: {
    show() {
      this.$modal.show('sign')
    },
    hide() {
      this.$modal.hide('sign')
    },
    login() {
      axios.post("http://localhost:9000/api/v1/user/login", {"username": this.username, "password": this.password})
          .then(response => {
            let res = response.data;
            this.message = res.message;
            if (res.success) {
              this.message = res.message;
              this.messageClass = "alert alert-success";
              this.username = null;
              this.password = null;
              window.localStorage.setItem('userId', res.id);
              axios.get("http://localhost:9000/api/v1/user?id="+res.id)
                  .then(response => {
                      this.loggedUser = response.data.username;
                    window.localStorage.setItem("user", response.data.username)
                  });
              this.hide();
            } else {
              this.message = res.message;
              this.messageClass = "alert alert-danger";
            }
          });
    },
    logout() {
      window.localStorage.removeItem("userId");
      window.localStorage.removeItem("user");
      this.loggedUser = null;
    },
    register() {
      axios.post(
          "http://localhost:9000/api/v1/user/register",
          {"username": this.usernameNewUser, "password": this.passwordNewUser}
          ).then(response => {
            let res = response.data;
            if (res.success) {
              this.message = res.message;
              this.messageClass = "alert alert-success";
              this.usernameNewUser = null;
              this.passwordNewUser = null;
              this.passwordCheckNewUser = null;
            } else {
              this.message = res.message;
              this.messageClass = "alert alert-danger";
              this.usernameNewUser = null;
              this.passwordNewUser = null;
              this.passwordCheckNewUser = null;
            }
        console.log(this.message);
          });
    }
  }
}
</script>
<style lang="css" scoped>

.borderRight {
  border-right: 1px solid silver;
}

.modal-content {
  border: none !important;
}

.modal-header {
  border: none !important;
}

</style>
