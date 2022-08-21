<template>
  <v-main>
    <v-container fluid fill-height>
      <v-layout align-center justify-center>
        <v-flex xs12 sm8 md6>
          <v-card class="elevation-12">
            <v-card-text>
              <v-col cols="12">
                <v-img class="logo"
                       src="https://user-images.githubusercontent.com/4760295/112476277-04b9b680-8d72-11eb-8433-fb53ddb9f78a.png"></v-img>
              </v-col>
              <v-form>
                <v-text-field
                  prepend-icon="lock"
                  name="password"
                  label="Password"
                  type="password"
                  v-model="password"
                ></v-text-field>
                <v-text-field
                  prepend-icon="lock"
                  name="passwordCheck"
                  label="Password Check"
                  type="password"
                  v-model="passwordCheck"
                ></v-text-field>
              </v-form>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" @click="recoveryPw()">Set Password</v-btn>
            </v-card-actions>
          </v-card>
        </v-flex>
      </v-layout>
    </v-container>
  </v-main>
</template>

<script lang="ts">
import {Component, Vue} from 'nuxt-property-decorator'
import IResponseUsers from "~/model/IResponseUsers";
import {namespace} from 'vuex-class';

const snackbar = namespace('Snackbar');

@Component
export default class RecoveryPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void
  password: string = "";
  passwordCheck: string = "";
  $route: any;
  $axios: any;

  recoveryPw() {
    const username = this.$route.params.username;
    this.$axios.post("recovery", {
      username: username,
      password: this.password
    }, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseUsers) => {
        if (response.data.success) {
          this.updateText(response.data.message);
          this.updateColor('green')
          this.updateShow(true);
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }
}
</script>

