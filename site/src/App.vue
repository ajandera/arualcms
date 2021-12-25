<template>
  <div>
    <Navigation :language="language" :languages="languages" />
    <Header :language="language" :languages="languages" />
    <Content :language="language" :languages="languages" />
    <Footer :language="language" :languages="languages" />
  </div>
</template>

<script>

import Navigation from "./components/Navigation";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import Content from "@/components/Content";
import axios from "axios";

export default {
  name: 'App',
  data() {
    return {
      message: null,
      messageClass: null,
      username: null,
      password: null,
      error: null,
      language: null,
      languages: []
    }
  },
  components: {
    Navigation,
    Header,
    Footer,
    Content
  },
  mounted() {
    this.getDefaultLanguage();
  },
  methods: {
    getDefaultLanguage() {
      axios.get(this.$hostname + "languages")
          .then(response => {
            if (response.data.success === true) {
              this.languages = response.data.languages.map(item => item = item.key);
              this.language = response.data.languages.find(item => item.default === true).key;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    },
  }
}
</script>

<style lang="css" scoped>
@import "https://stackpath.bootstrapcdn.com/bootstrap/4.1.2/css/bootstrap.min.css";

</style>
