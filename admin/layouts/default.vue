<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title"/>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
    >
      <v-btn
        icon
        @click.stop="miniVariant = !miniVariant"
      >
        <v-icon>mdi-{{ `chevron-${miniVariant ? 'right' : 'left'}` }}</v-icon>
      </v-btn>
      <v-toolbar-title v-text="title"/>
      <v-spacer/>
      <v-menu
        bottom
        origin="center center"
        transition="scale-transition"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn
            color="secondary"
            dark
            v-bind="attrs"
            v-on="on"
          >
            {{ language }}
          </v-btn>
        </template>

        <v-list>
          <v-list-item
            v-for="(item, i) in languages"
            :key="i"
            @click="onLangChange(item)"
          >
            <v-list-item-title>{{ item }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>

      <v-btn
        icon
        @click="logout"
      >
        <v-icon>mdi-logout-variant</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <NuxtChild :language="language"/>
      </v-container>
    </v-main>
    <v-footer
      :absolute="!fixed"
      app
    >
      <span>v0.9.2 arualcms.eu &copy; {{ new Date().getFullYear() }} </span>
    </v-footer>
    <Snackbar/>
  </v-app>
</template>

<script lang="ts">
import {Component, Vue} from 'nuxt-property-decorator'
import IResponseLanguage from "~/model/IResponseLanguage";
import Snackbar from "~/components/Snackbar.vue";

@Component({
  components: {
    Snackbar
  }
})
export default class DefaultLayout extends Vue {
  clipped: boolean = true
  drawer: boolean = true
  fixed: boolean = true
  languages: string[] = []
  language: string = ""
  items: Array<any> = [
    {
      icon: 'mdi-apps',
      title: 'Posts',
      to: '/posts'
    },
    {
      icon: 'mdi-card-text',
      title: 'Texts',
      to: '/texts'
    },
    {
      icon: 'mdi-file-multiple',
      title: 'Files',
      to: '/files'
    },
    {
      icon: 'mdi-flag',
      title: 'Languages',
      to: '/languages'
    },
    {
      icon: 'mdi-cog',
      title: 'Setting',
      to: '/settings'
    },
    {
      icon: 'mdi-account',
      title: 'Users',
      to: '/users'
    }

  ]
  miniVariant: boolean = false
  title: string = 'arualCMS'
  $axios: any
  $router: any;
  $i18n: any;
  $auth: any;
  $nuxt: any;

  mounted() {
    this.guard();
  }

  guard() {
    if (!this.$auth.loggedIn) {
      this.$nuxt.$options.router.push('/in');
    } else {
      this.getDefaultLanguage();
    }
  }

  async logout() {
    await this.$auth.logout();
    this.$nuxt.$options.router.push('/in');
  }

  getDefaultLanguage(): void {
    this.$axios.get("/languages")
      .then((response: IResponseLanguage) => {
        this.language = response.data.languages.filter(item => item.Default)[0].key;
        this.languages = response.data.languages.map(item => item.Key);
      });
  }

  onLangChange(lang: string): void {
    this.language = lang;
  }
}
</script>
