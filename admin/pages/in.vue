<template>
  <v-main>
    <v-container fluid fill-height>
      <v-layout align-center justify-center>
      <v-flex xs12 sm8 md6>
        <v-alert
          v-if="message.text"
          :type="message.class"
          border="left"
          dense
          prominent
        >{{ message.text }}</v-alert>
        <v-card v-if="forgot === false" class="elevation-12">
          <v-card-text>
            <v-col cols="12">
              <v-img class="logo" src="https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png"></v-img>
            </v-col>
            <v-form>
              <v-text-field
                prepend-icon="person"
                name="login"
                label="Login"
                type="text"
                v-model="user.username"
              ></v-text-field>
              <v-text-field
                id="password"
                prepend-icon="lock"
                name="password"
                label="Password"
                type="password"
                v-model="user.password"
              ></v-text-field>
            </v-form>
            <p @click="forgotForm()" class="fp">Forgot your password?</p>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="login()">Login</v-btn>
          </v-card-actions>
        </v-card>
        <v-card v-if="forgot === true" class="elevation-12">
          <v-card-text>
            <v-col cols="12">
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
  </v-main>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import User from '~/model/User';
import Message from '~/model/Message';
import IResponseAuth from "~/model/IResponseAuth";

@Component({
    layout: "sign"
})
export default class InPage extends Vue {
    message: Message = {text: "", class: ""}
    forgot: boolean = false
    forgotMail: string = ""
    user: User = {username: "", password: ""}
    $axios: any;
    $router: any;

    login(): void {
      this.$axios.post("/auth", this.user, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseAuth) => {
          if (response.data.success) {
            window.localStorage.setItem('jwt', response.data.jwt);
            this.$router.push('/posts');
          } else {
            this.message = {
                "text": response.data.message,
                "class": "error"
            }
          }
        });
    }

    forgotForm(): void {
      this.forgot = !this.forgot;
    }

    requestNewPw(): void {
      const param = {
        "username": this.forgotMail,
        "domain": window.location.protocol + "//" + window.location.hostname
      };

      // send email
     this.$axios.$post("/sendRecovery", param)
        .then((response: { data: { message: any; success: string }; }) => {
          let res = response.data;
          this.message = res.message;
          if (res.success) {
            this.message = {
              "text": response.data.message,
              "class": "success"
            }
            this.forgotMail = "";
          } else {
            this.message = {
              "text": res.message,
              "class": "error"
            }
          }
        });
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
.fp {
    cursor: pointer
}
</style>
