<template>
  <!-- Navigation -->
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top">
    <div class="container">
      <a class="navbar-brand" href="#">{{ setting.find((x) => x.key === 'title').value[language] }}</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarResponsive">
        <ul class="navbar-nav ml-auto">
          <li class="nav-item active">
            <a class="nav-link" href="/">Home
              <span class="sr-only">(current)</span>
            </a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#news">News</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#contact">Contact</a>
          </li>
          <li class="nav-item" v-for="(lang, index) in languages" v-bind:key="index">
            <button
                v-on:click="setLanguage(lang)"
                v-bind:class="{'btn btn-default': lang !== language, 'btn btn-primary': lang === language}">{{ lang }}</button>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>

import axios from "axios";

export default {
  name: 'Navigation',
  props: ['language', 'languages'],
  components: {},
  data() {
    return {
      setting: []
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    load() {
      axios.get(this.$hostname + "setting")
          .then(response => {
            if (response.data.success === true) {
              this.setting = response.data.settings;
              console.log(this.setting);
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    },
    setLanguage(lang) {
      this.language = lang;
    }
  }
}
</script>

<style lang="css" scoped>

</style>
