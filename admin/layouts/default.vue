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
            <v-list-item-title v-text="item.title" />
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
      <v-toolbar-title v-text="title" />
      <v-spacer />
      <v-btn
        icon
        @click="logout"
      >
        <v-icon>mdi-logout-variant</v-icon>
      </v-btn>
    </v-app-bar>
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
    <v-footer
      :absolute="!fixed"
      app
    >
      <span>v0.9.2 arualcms.eu &copy; {{ new Date().getFullYear() }} </span>
    </v-footer>
    <v-snackbar
      v-model="message.text !== ''"
      :timeout="2000"
      :color="message.class"
    >
      {{ message.text }}

    </v-snackbar>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Language from '~/model/Language'
import Message from "~/model/Message";

@Component
export default class DefaultLayout extends Vue {
    clipped: boolean = true
    drawer: boolean = true
    fixed: boolean = true
    languages: string[] = []
    language: string = ""
    message: Message = {text: "", class: ""}
    items: Array<any> = [
        {
          icon: 'mdi-apps',
          title: 'Posts',
          to: '/'
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

    mounted() {
        this.getDefaultLanguage();
    }

    logout(): void {
      window.localStorage.removeItem("jwt");
      this.$router.push('in');
    }

    getDefaultLanguage(): void {
      this.$axios.get("/languages")
        .then((response: { data: { success: boolean; languages: Language[] } }) => {
          this.language = response.data.languages.filter(item => item.default)[0].key;
          this.languages = response.data.languages.map(item => item.key);
        });
    }

    setLanguage(lang: string): void {
      this.language = lang;
    }
}
</script>
