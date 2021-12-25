<template>
  <!-- Header -->
  <header class="bg-primary py-5 mb-5">
    <div class="container h-100">
      <div class="row h-100 align-items-center">
        <div class="col-lg-12">
          <h1 class="display-4 text-white mt-5 mb-2">{{ setting.find((x) => x.key === 'title').value[language] }}</h1>
          <p class="lead mb-5 text-white-50">{{ setting.find((x) => x.key === 'subtitle').value[language] }}</p>
        </div>
      </div>
    </div>
  </header>
</template>

<script>

import axios from "axios";

export default {
  name: 'Header',
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
      this.posts = [];
      axios.get(this.$hostname + "setting")
          .then((response) => {
            if (response.data.success === true) {
              this.setting = response.data.settings;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          }, (err) => {
            console.log(err);
      });
    }
  }
}
</script>

<style lang="css" scoped>

</style>
