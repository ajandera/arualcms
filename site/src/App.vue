<template>
  <div>
    <Navigation />
    <Header />
    <Content />
    <Footer />
  </div>
</template>

<script>

import Navigation from "./components/Navigation";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
import Content from "@/components/Content";

export default {
  name: 'App',
  data() {
    return {
      message: null,
      messageClass: null,
      username: null,
      password: null,
      error: null,
    }
  },
  components: {
    Navigation,
    Header,
    Footer,
    Content
  },
  mounted() {

  },
  methods: {
  }
}
</script>

<style lang="css" scoped>
@import "https://stackpath.bootstrapcdn.com/bootstrap/4.1.2/css/bootstrap.min.css";

</style>
