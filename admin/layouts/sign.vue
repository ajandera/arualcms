<template>
  <v-app light>
    <v-toolbar class="red darken-3">
      <v-toolbar-side-icon>
        <img src="/logosm.png" alt="arualCMS" class="logo">
      </v-toolbar-side-icon>
      <v-toolbar-title class="mx-0 white--text" v-text="title"></v-toolbar-title>
      <p>Beta</p>
      <v-spacer></v-spacer>
      <v-toolbar-items>
        <v-dialog
          v-model="dialogReg"
          width="500"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn 
              text
              v-bind="attrs"
              v-on="on"
              class="white--text">
              Sign Up
            </v-btn>
          </template>

          <v-card>
            <v-card-title class="text-h5 teal white--text lighten-2">
              Sign Up
            </v-card-title>
            <v-card-text>
              <v-form>
                <v-text-field
                  prepend-icon="person"
                  name="name"
                  label="Name"
                  type="text"
                  v-model="user.name"
                ></v-text-field>
                <v-text-field
                  prepend-icon="person"
                  name="username"
                  label="Username"
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
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="secondary" text @click="dialogReg = false">
                Close
              </v-btn>
              <v-btn color="success" @click="registerAccount()">Sign Up</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
        <Nuxt />
      </v-toolbar-items>
    </v-toolbar>
    <v-content>
      <section>
        <v-parallax :src="imageLink.sub_main" height="600">
          <v-layout column align-center justify-center class="white--text">
            <h1 class="white--text mb-2 display-1 text-xs-center" style="font-weight: 900;text-shadow: 3px 2px rgb(204, 202, 202);color: #000 !important">
              The headless CMS for epic content
            </h1>
            <div class="white--text subheading mb-3 text-xs-center" style="font-weight: 900;text-shadow: 3px 2px rgb(204, 202, 202);color: #000 !important">
              Unlesh your creativity without limitations
            </div>
            <v-btn class="red darken-3 mt-5" dark large @click="dialogReg = true">
              Get Started
            </v-btn>
          </v-layout>
        </v-parallax>
      </section>

      <section class="red darken-3 mb-5">
        <v-layout
          column
          wrap
          class="my-5"
          align-center
        >
          <v-flex xs12 sm4 class="my-3 mt-5">
            <div class="text-center">
              <h2 class="headline white--text">
                The best way to share your amazing stuff
              </h2>
              <span class="subheading">
                No more restrictions, no more limits, amazing speed
              </span>
            </div>
          </v-flex>
          <v-flex xs12>
            <v-container grid-list-xl>
              <v-layout row wrap align-center>
                <v-flex xs12 md4>
                  <v-card class="elevation-0 transparent">
                    <v-card-text class="text-center">
                      <v-icon x-large class="teal--text text--lighten-2">public</v-icon>
                    </v-card-text>
                    <v-card-title primary-title class="layout justify-center">
                      <div class="headline text-xs-center white--text">Reach the world</div>
                    </v-card-title>
                    <v-card-text class="black--text">
                      Show your stuff to the whole internet not only to your mum or your friends.
                      We love making good content. In this moment arualCMS is used by freelance developers
                       who are not famous but that want to share their works to the world.
                        Unfortunately with other CMS is creativity blocked or expensive.
                    </v-card-text>
                  </v-card>
                </v-flex>
                <v-flex xs12 md4>
                  <v-card class="elevation-0 transparent">
                    <v-card-text class="text-center">
                      <v-icon x-large class="teal--text text--lighten-2">flash_on</v-icon>
                    </v-card-text>
                    <v-card-title primary-title class="layout justify-center">
                      <div class="headline white--text">Fast communication</div>
                    </v-card-title>
                    <v-card-text class="black--text">
                      Time is important, we don't want you to waste it. So our Rest API reponse is only <strong>30ms</strong>.
                       Here you can get a massive influence on real users in minutes. Our CMS give ypu a tool with speed and easy connection.
                        And if your cntent is appreciated you won't only get positive feedback but also lovely and sincere fans
                    </v-card-text>
                  </v-card>
                </v-flex>
                <v-flex xs12 md4>
                  <v-card class="elevation-0 transparent">
                    <v-card-text class="text-center">
                      <v-icon x-large class="teal--text text--lighten-2">share</v-icon>
                    </v-card-text>
                    <v-card-title primary-title class="layout justify-center">
                      <div class="headline text-xs-center white--text">Create new consumers</div>
                    </v-card-title>
                    <v-card-text class="black--text">
                      Imagine if you can directly speak with the world's population.
                       Don't you think it would be easier to find nice people to interact with?
                        arualCMS is both local and global and help you to connect without limitations
                         with people from your city, your state and your universe! 
                    </v-card-text>
                  </v-card>
                </v-flex>
              </v-layout>
            </v-container>
          </v-flex>
        </v-layout>
      </section>

      <section>
        <v-parallax :src="imageLink.main" height="380">
          <v-layout column align-center justify-center>
            <div class="headline mb-3 text-xs-center mainHeadline">
              arualCMS is a headless CMS that allows everyone to reach a huge audience with a tap 
            </div>
            <em>With the power of arualCMS you don't need to be famous or post pics of cute cats in order to get visibility</em>
            <v-btn class="red darken-3 mt-5" dark large @click="register = true">
              Get more info
            </v-btn>
          </v-layout>
        </v-parallax>
      </section>

      <section class="red darken-3">
        <v-container grid-list-xl>
          <v-layout row wrap justify-center class="my-5">
            <v-flex xs12 sm4>
              <v-card class="elevation-0 transparent">
                <v-card-title primary-title class="layout justify-center">
                  <div class="headline white--text">Company info</div>
                </v-card-title>
                <v-card-text class="black--text">
                  We are not a company. We hate companies.
                   Just imagine us like the guys from the Silicon Valley series. We are start up. We made it with love.
                </v-card-text>
              </v-card>
            </v-flex>
            <v-flex xs12 sm4 offset-sm1>
              <v-card class="elevation-0 transparent">
                <v-card-title primary-title class="layout justify-center">
                  <div class="headline white--text">We are hiring</div>
                </v-card-title>
                <v-card-text class="black--text">
                  Are you a creative person? Do you like techy stuff? 
                  Send me an email on <a href="mailto:ales@storepredictor.com">ales@storepredictor.com</a>.
                </v-card-text>
              </v-card>
            </v-flex>
          </v-layout>
        </v-container>
      </section>

      <v-footer class="blue darken-2">
        <v-layout row wrap align-center>
          <v-flex xs12 class="text-xs-center">
            <div class="white--text ml-3">
              Made with
              <v-icon class="red--text">favorite</v-icon>
              using VueJS and Go Lang
            </div>
          </v-flex>
        </v-layout>
      </v-footer>

    </v-content>
    <Snackbar />
</v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Snackbar from "~/components/Snackbar.vue";
import ILanding from '~/model/ILanding';
import UserRegistration from '~/model/UserRegitration';
import { namespace } from 'vuex-class';

const snackbar = namespace('Snackbar');

@Component({
  components: {
    Snackbar
  },
  layout: 'sign'
})
export default class SignLayout extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void

  title: string = "arualCMS";
  register: boolean = false;
  dialogReg: boolean = false;
  imageLink: ILanding = {
    main:
      "/posts.png",
    sub_main:
      "/post-detail.png",
    social_cover:
      "/text-detail.png"
  };
  user: UserRegistration = {username: "", password: "", name: ""}
  $axios: any;

  registerAccount() {
    this.$axios.post("/registration", {
      name: this.user.name,
      username: this.user.username,
      password: this.user.password,
    })
      .then((response: { data: { message: any; success: string }; }) => {
        let res = response.data;
        if (res.success) {
          this.updateText(response.data.message);
          this.updateColor('green');
          this.updateShow(true);
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
.v-parallax__image-container {
  opacity: 0.4;
}
.logo {
  max-height: 60px;
}
.mainHeadline {
  font-weight: 900 !important;
  text-shadow: 3px 2px rgb(204, 202, 202);
  color: #000 !important;
}
</style>
