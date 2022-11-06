<template>
  <v-app light>
    <v-toolbar class="white">
      <v-toolbar-side-icon>
        <img src="/logosm.png" alt="arualCMS" style="max-height: 60px">
      </v-toolbar-side-icon>
      <v-toolbar-title class="mx-0" v-text="title"></v-toolbar-title>
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
          v-on="on">
          Sign Up
        </v-btn>
      </template>

      <v-card>
        <v-card-title class="text-h5 grey lighten-2">
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
          <v-btn color="primary" @click="registerAccount()">Sign Up</v-btn>
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
            <h1 class="white--text mb-2 display-1 text-xs-center" style="font-weight: 900; text-shadow: 3px 2px #000000">
              The headless CMS for epic content
            </h1>
            <div class="white--text subheading mb-3 text-xs-center" style="font-weight: 900; text-shadow: 2px 2px #000000">
              Unlesh your creativity without limitations
            </div>
            <v-btn class="blue lighten-2 mt-5" dark large @click="register = true">
              Get Started
            </v-btn>
          </v-layout>
        </v-parallax>
      </section>

      <section>
        <v-layout
          column
          wrap
          class="my-5"
          align-center
        >
          <v-flex xs12 sm4 class="my-3">
            <div class="text-xs-center">
              <h2 class="headline">The best way to share your amazing stuff</h2>
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
                    <v-card-text class="text-xs-center">
                      <v-icon x-large class="blue--text text--lighten-2">public</v-icon>
                    </v-card-text>
                    <v-card-title primary-title class="layout justify-center">
                      <div class="headline text-xs-center">Reach the world</div>
                    </v-card-title>
                    <v-card-text>
                      Show your stuff to the whole internet not only to your mum or your friends.
                      We love making good content. In this moment arualCMS is used by freelance developers
                       who are not famous but that want to share their works to the world.
                        Unfortunately with other CMS is creativity blocked or expensive.
                    </v-card-text>
                  </v-card>
                </v-flex>
                <v-flex xs12 md4>
                  <v-card class="elevation-0 transparent">
                    <v-card-text class="text-xs-center">
                      <v-icon x-large class="blue--text text--lighten-2">flash_on</v-icon>
                    </v-card-text>
                    <v-card-title primary-title class="layout justify-center">
                      <div class="headline">Fast communication</div>
                    </v-card-title>
                    <v-card-text>
                      Time is important, we don't want you to waste it. So our Rest API reponse is only <strong>30ms</strong>.
                       Here you can get a massive influence on real users in minutes. Our CMS give ypu a tool with speed and easy connection.
                        And if your cntent is appreciated you won't only get positive feedback but also lovely and sincere fans
                    </v-card-text>
                  </v-card>
                </v-flex>
                <v-flex xs12 md4>
                  <v-card class="elevation-0 transparent">
                    <v-card-text class="text-xs-center">
                      <v-icon x-large class="blue--text text--lighten-2">share</v-icon>
                    </v-card-text>
                    <v-card-title primary-title class="layout justify-center">
                      <div class="headline text-xs-center">Create new consumers</div>
                    </v-card-title>
                    <v-card-text>
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
            <div class="headline white--text mb-3 text-xs-center">arualCMS is a headless CMS that allows everyone to reach a huge audience with a tap </div>
            <em>With the power of arualCMS you don't need to be famous or post pics of cute cats in order to get visibility</em>
            <v-btn class="blue lighten-2 mt-5" dark large @click="register = true">
              Get more info
            </v-btn>
          </v-layout>
        </v-parallax>
      </section>

       <section>
        <v-container grid-list-md>
          <v-layout row wrap>
            <v-flex xs12 text-xs-center class="mt-5">
            <div class="headline">Are you amazed? Stay tuned!</div>
            <br>
            <div>We are lunching the beta in a few time. If you want to be one of the first Endorfine users we will email you as soon as we're ready. In the beginning only few people will test before the launch. Let us know how Endorfine will help you!</div>
            </v-flex>
          </v-layout>
        </v-container>
      </section>

       <section>
        <v-parallax :src="imageLink.social_cover" height="380">
          <v-layout column align-center justify-center>
            <div class="headline white--text mb-3 text-xs-center">We are dropping cool news and opportunities on CMS</div>
          </v-layout>
          <v-layout justify-space-around justify-center>

              <v-icon x-large dark>fab fa-facebook-f</v-icon>

              <v-icon x-large dark>fab fa-twitter</v-icon>

              <v-icon x-large dark>fab fa-instagram</v-icon>

            </v-layout>
        </v-parallax>
      </section>

      <section>
        <v-container grid-list-xl>
          <v-layout row wrap justify-center class="my-5">
            <v-flex xs12 sm4>
              <v-card class="elevation-0 transparent">
                <v-card-title primary-title class="layout justify-center">
                  <div class="headline">Company info</div>
                </v-card-title>
                <v-card-text>
                  We are not a company. We hate companies.
                   Just imagine us like the guys from the Silicon Valley series. We are stat up. We made it with love.
                </v-card-text>
              </v-card>
            </v-flex>
            <v-flex xs12 sm4 offset-sm1>
              <v-card class="elevation-0 transparent">
                <v-card-title primary-title class="layout justify-center">
                  <div class="headline">We are hiring</div>
                </v-card-title>
                <v-card-text>
                  Are you a creative person? Do you like techy stuff? 
                  Send me he email on ales@storepredictor.com.
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
      "https://firebasestorage.googleapis.com/v0/b/endorfinevue.appspot.com/o/assets%2Fb13f0434-b228-11e6-8e5d-5252025056ab_web_scale_0.4666667_0.4666667__.jpg?alt=media&token=660df23e-599e-434b-9313-ba69c973eeea",
    sub_main:
      "https://firebasestorage.googleapis.com/v0/b/endorfinevue.appspot.com/o/assets%2FNight-Club-Clubbing-Jobs-Abroad2.jpg?alt=media&token=82bbda7d-5df4-430b-9217-adaf1c8485c5",
    social_cover:
      "https://firebasestorage.googleapis.com/v0/b/endorfinevue.appspot.com/o/assets%2Fo-NIGHTCLUB-facebook.jpg?alt=media&token=cefc5c4c-9714-41da-9c22-f63caf5e89a4"
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

</style>
