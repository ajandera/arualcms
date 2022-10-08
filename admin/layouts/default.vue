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
            {{ site.Name }}
          </v-btn>
        </template>

        <v-list>
          <v-list-item
            v-for="(item, i) in sites"
            :key="i"
            @click="onSiteChange(item)"
          >
            <v-list-item-title>{{ item.Name }}</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>
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
        <NuxtChild :language="language" :languages="languages" :defaultLanguage="defaultLanguage" />
      </v-container>
    </v-main>
    <v-footer
      :absolute="!fixed"
      app
    >
      <span>v0.9.3 arualcms.eu &copy; {{ new Date().getFullYear() }} </span>
    </v-footer>
    <Snackbar/>
  </v-app>
</template>

<script lang="ts">
import {Component, Vue, Watch} from 'nuxt-property-decorator'
import IResponseLanguage from "~/model/IResponseLanguage";
import Snackbar from "~/components/Snackbar.vue";
import IResponseSite from "~/model/IResponseSite";
import Site from '~/model/Site';

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
  defaultLanguage: string = ""
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
      icon: 'mdi-earth',
      title: 'Sites',
      to: '/sites'
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
  sites: Site[] = [];
  site: Site = {Id: "", Name: ""};

  mounted() {
    this.guard();
  }

  @Watch('$route')
  onPropertyChanged(value: string, oldValue: string) {
    if (!this.$route.query.siteId && this.site !== null) {
      if (this.site.Id === undefined) {
        this.$router.push({path: this.$route.path, query: {siteId: this.site}})
      } else {
        this.$router.push({path: this.$route.path, query: {siteId: this.site.Id}})
      }
    }
  }

  guard() {
    if (!this.$auth.loggedIn) {
      this.$nuxt.$options.router.push('/in');
    } else {
      this.getSites();
    }
  }

  async logout() {
    await this.$auth.logout();
    this.$nuxt.$options.router.push('/in');
  }

  async getDefaultLanguage(siteId: string): Promise<void> {
    this.$axios.get("/" + siteId +"/languages")
      .then((response: IResponseLanguage) => {
        this.language = response.data.languages.filter(item => item.Default)[0].Key;
        this.defaultLanguage = response.data.languages.filter(item => item.Default)[0].Key;
        this.languages = response.data.languages.map(item => item.Key);
      });
  }

  async getSites(): Promise<void> {
    this.$axios.get("/sites")
      .then((response: IResponseSite) => {
        this.sites = response.data.sites;
        this.site = this.sites[0];
        this.$router.push({path: this.$route.path, query: {siteId: this.site.Id}});
        this.getDefaultLanguage(this.site.Id)
      });
  }

  onLangChange(lang: string): void {
    this.language = lang;
  }

  onSiteChange(site: Site): void {
    this.site = site;
    this.$router.push({path: this.$route.path, query: {siteId: this.site.Id}})
    this.getDefaultLanguage(this.site.Id)
  }
}
</script>
