<template>
  <modal name="sign" height="auto" :clickToClose="false">
    <div class="modal-dialog">
      <!-- Modal content-->
      <div class="modal-content">
        <div class="modal-header">
          <h3 class="modal-title">
            <img src="https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png">
          </h3>
        </div>
        <div class="modal-body">
          <div class="row">
            <div class="col-12">
              <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
              <div id="signin" v-if="forgot === false">
                <label for="username">Username</label>
                <input type="email" v-model="username" class="form-control" id="username" required>
                <label for="password">Password</label>
                <input type="password" v-model="password" class="form-control" id="password" required>
                <a v-on:click="forgotForm()" href="#">Forgot your password?</a>
                <button class="float-right btn btn-success mt-3" type="submit" v-on:click="login()">Sign In</button>
              </div>
              <div id="forgot" v-if="forgot === true">
                <label for="forgotMail">Username</label>
                <input type="email" v-model="forgotMail" class="form-control" id="forgotMail" required>
                <a v-on:click="forgotForm()" href="#"> back to login</a>
                <button class="float-right btn btn-success mt-3" type="submit" v-on:click="requestNewPw()">Request new Password</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </modal>
</template>

<script>

import axios from "axios";

export default {
  name: 'SignComponent',
  props: ['language', 'languages'],
  components: {},
  data() {
    return {
      messageClass: null,
      message: null,
      forgot: false,
      forgotMail: null,
      username: null,
      password: null,
    }
  },
  mounted() {
    this.$modal.show('sign');
  },
  methods: {
    login() {
      axios.post(this.$hostname+"auth", {"username": this.username, "password": this.password})
          .then(response => {
            let res = response.data;
            this.message = res.message;
            if (res.success) {
              this.message = res.message;
              this.messageClass = "alert alert-success";
              this.username = null;
              this.password = null;
              this.loggedUser = res.username;
              window.localStorage.setItem('userId', res.id);
              window.localStorage.setItem("user", res.username)
              window.localStorage.setItem('jwt', res.jwt);
              this.hide();
              this.$router.push({name: 'posts'});
            } else {
              this.message = res.message;
              this.messageClass = "alert alert-danger";
            }
          });
    },
    hide() {
      this.$modal.hide('sign')
    },
    forgotForm() {
      this.forgot = !this.forgot;
    },
    requestNewPw() {
      const param = {
        "username": this.forgotMail,
        "domain": window.location.protocol + "//" + window.location.hostname
      };
      
      // send email
      axios.post(this.$hostname+"sendRecovery", param)
          .then(response => {
            let res = response.data;
            this.message = res.message;
            if (res.success) {
              this.message = res.message;
              this.messageClass = "alert alert-success";
              this.forgotMail = null;
            } else {
              this.message = res.message;
              this.messageClass = "alert alert-danger";
            }
          });
    }
  }
}
</script>

<style lang="css" scoped>
.modal-dialog {
  max-width: 800px;
  margin: 10px;
}
.modal-content {
  border: none;
}
.v--modal-overlay {
  background: url("/admin/images/arualbg.jpg") center center !important;
}
</style>
