<template>
  <v-content>
    <v-container fluid fill-height>
    <v-layout align-center justify-center>
      <v-flex xs12 sm8 md4>
        <v-alert
          v-if="message"
          :type="alertType"
          border="left"
          dense
          prominent
        >{{ message }}</v-alert>
        <v-card v-if="forgot === false" class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>
              arualCSM Login
            </v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-col cols="8" offset="2">
              <v-img class="logo" src="https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png"></v-img>
            </v-col>
            <v-form>
              <v-text-field
                prepend-icon="person"
                name="login"
                label="Login"
                type="text"
              ></v-text-field>
              <v-text-field
                id="password"
                prepend-icon="lock"
                name="password"
                label="Password"
                type="password"
              ></v-text-field>
            </v-form>
            <p @click="forgotForm()">Forgot your password?</p>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="login()">Login</v-btn>
          </v-card-actions>
        </v-card>
        <v-card v-if="forgot === true" class="elevation-12">
          <v-toolbar dark color="primary">
            <v-toolbar-title>
              arualCSM Forgot password
            </v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-col cols="8" offset="2">
              <v-img class="logo" src="https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png"></v-img>
            </v-col>
            <v-form>
              <v-text-field
                prepend-icon="person"
                name="login"
                label="Registered password"
                type="email"
                v-model="forgotMail"
              ></v-text-field>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="default" @click="forgotForm()">Back to login</v-btn>
            <v-btn color="primary" @click="requestNewPw()">Request new Password</v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
  </v-content>
</template>

<script lang="js">
import axios from "axios";

export default {
  name: 'InPAge',
  layout: 'sign',
  components: {},
  data() {
    return {
      alertType: null,
      message: null,
      forgot: false,
      forgotMail: null,
      username: null,
      password: null,
    }
  },
  methods: {
    login() {
      axios.post(this.$config.api+"auth", {"username": this.username, "password": this.password})
        .then(response => {
          let res = response.data;
          this.message = res.message;
          if (res.success) {
            this.message = res.message;
            this.alertType = "success";
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
            this.alertType = "error";
          }
        }).catch(error => {
        this.message = error.message;
        this.alertType = "error";
      });
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
      axios.post(this.$config.api+"sendRecovery", param)
        .then(response => {
          let res = response.data;
          this.message = res.message;
          if (res.success) {
            this.message = res.message;
            this.alertType = "success";
            this.forgotMail = null;
          } else {
            this.message = res.message;
            this.alertType = "error";
          }
        });
    }
  }
}
</script>
<style>
.v-application--wrap {
  background: url("/arualbg.jpg") center center !important;
}
.logo {
  border-radius: 20px
}
</style>
