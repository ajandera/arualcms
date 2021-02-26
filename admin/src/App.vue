<template>
  <div id="app">
    <Header/>
    <section class="main mt-3">
      <div class="container-fluid">
        <div class="row">
          <div class="col-2">
            <Menu/>
          </div>
          <div class="col-1O">
            <component :is="selected" class="tab"></component>
          </div>
        </div>
      </div>
    </section>
    <section id="footer">
      <div class="row">
        <div class="col-12 text-center">
          <p>2021 &copy; Aleš Jandera, <a href="https://ajandera.com" target="_blank">ajandera.com</a></p>
        </div>
      </div>
    </section>
  </div>
</template>

<script>

import Editor from '@/components/Editor.vue';
import Setting from '@/components/Setting';
import Header from "@/components/Header";
import Menu from "@/components/Menu";

export default {
  name: 'App',
  data() {
    return {
      selected: "Editor",
    }
  },
  components: {
    Header,
    Editor,
    Setting,
    Menu
  }
}
</script>

<style lang="css" scoped>
@import "https://stackpath.bootstrapcdn.com/bootstrap/4.1.2/css/bootstrap.min.css";

html, body {
  height: 100%;
}

#app {
  min-height: 100%;
}

.main {
  overflow:auto;
  padding-bottom:20%; /* this needs to be bigger than footer height*/
}

#footer {
  position: relative;
  margin-top: -10%; /* negative value of footer height */
  height: 150px;
  clear:both;
  padding-top:20px;
}
</style>
