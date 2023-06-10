<template>
  <v-dialog
      v-model="dialog"
      width="500"
    >
      <template v-slot:activator="{ on, attrs }">
        <v-btn
          text
          class="main-btn"
          v-bind="attrs"
          v-on="on">
            Sign In
        </v-btn>
      </template>

      <v-card>
        <v-card-title class="text-h5 special white--text lighten-2">
          Sign In
        </v-card-title>

        <v-card-text v-if="forgot === false">
          <br>
          <v-form>
            <v-text-field
              prepend-icon="person"
              name="login"
              label="Login"
              type="text"
              v-model="user.Username"
            ></v-text-field>
            <v-text-field
              id="password"
              prepend-icon="lock"
              name="password"
              label="Password"
              type="password"
              v-model="user.Password"
            ></v-text-field>
          </v-form>
              <p @click="forgotForm()" class="fp">Forgot your password?</p>
        </v-card-text>
        <v-card-actions v-if="forgot === false">
          <v-spacer></v-spacer>
          <v-btn color="secondary" text @click="dialog = false">
            Close
          </v-btn>
          <v-btn color="success" @click="login()">Login</v-btn>
        </v-card-actions>
        <v-card-text v-if="forgot === true">
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
        <v-card-actions v-if="forgot === true">
          <v-spacer></v-spacer>
          <v-btn color="default" @click="forgotForm()">Back to login</v-btn>
          <v-btn color="success" @click="requestNewPw()">Request new Password</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
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

  dialog: boolean = false
  forgot: boolean = false
  forgotMail: string = ""
  user: User = {Username: "", Password: "", Id: ""}
  $axios: any;
  $router: any;
  $auth: any;
  $nuxt: any;

  async login(): Promise<void> {
    try {
      let login = {
        "username": this.user.Username,
        "password": this.user.Password
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
.special {
  background: #1976d2;
  color: #fff !important;
}
</style>
