<template>
  <v-main>
    <v-container fluid fill-height>
      <v-layout align-center justify-center>
        <v-flex xs12 sm6 md4>
          <v-card v-if="forgot === false" class="elevation-12">
            <v-card-text>
              <v-col cols="12">
                <v-img class="logo" src="/logo.png"/>
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
                <v-img class="logo"
                       src="https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png"></v-img>
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
import {Component, Vue} from 'nuxt-property-decorator'
import User from '~/model/User';
import {namespace} from 'vuex-class';

const snackbar = namespace('Snackbar');

@Component({
  layout: 'sign'
})
export default class InPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void

  forgot: boolean = false
  forgotMail: string = ""
  user: User = {username: "", password: "", id: ""}
  $axios: any;
  $router: any;
  $auth: any;
  $nuxt: any;

  async login(): Promise<void> {
    try {
      let login = {
        "username": this.user.username,
        "password": this.user.password
      }
      await this.$auth.loginWith('local', { data: login })
      await this.$nuxt.$options.router.push('/');
    } catch (err) {
      this.updateText("Username and password mismatch.");
      this.updateColor('red')
      this.updateShow(true);
    }
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
    this.$axios.$post("/recovery", param)
      .then((response: { data: { message: any; success: string }; }) => {
        let res = response.data;
        if (res.success) {
          this.updateText(response.data.message);
          this.updateColor('green')
          this.updateShow(true);
          this.forgotMail = "";
        } else {
          this.updateText(res.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }
}
</script>

<style>
.signin .fp {
  cursor: pointer
}
</style>
