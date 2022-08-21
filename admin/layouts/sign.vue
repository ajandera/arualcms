<template>
  <v-app dark class="signin">
    <v-main>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
    <Snackbar />
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Snackbar from "~/components/Snackbar.vue";

@Component({
  components: {
    Snackbar
  },
  layout: 'sign'
})
export default class SignLayout extends Vue {

}
</script>

<style>
.signin {
  background: url("/arualbg.jpg") center center !important;
}
</style>
