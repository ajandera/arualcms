<template>
  <!-- Page Content -->
  <div class="container">
    <div class="row">
      <div class="col-md-8 mb-5">
        <h2>What We Do</h2>
        <hr>
        <div v-html="texts.find((x) => x.key === 'what_we_do').value[language]"></div><br>
        <a class="btn btn-primary btn-lg mt-1" href="#">Call to Action &raquo;</a>
      </div>
      <div class="col-md-4 mb-5" id="contact">
        <h2>Contact Us</h2>
        <hr>
        <div v-html="texts.find((x) => x.key === 'contact').value[language]"></div>
      </div>
    </div>
    <!-- /.row -->
    <div class="row" id="news">
      <div class="col-md-8 mb-5">
        <h2>News from our blog</h2>
        <hr>
      </div>
    </div>
    <div class="row">
      <div class="col-md-4 mb-5" v-for="post in posts" v-bind:key="post.id">
        <div class="card h-100">
          <img class="card-img-top" src="https://placehold.it/300x200" alt="">
          <div class="card-body">
            <h4 class="card-title">{{ post.title[language] }}</h4>
            <p class="card-text">{{ post.excerpt[language] }}</p>
          </div>
          <div class="card-footer">
            <a href="#" class="btn btn-primary">Read More!</a>
          </div>
        </div>
      </div>
    </div>
    <!-- /.row -->
  </div>
  <!-- /.container -->
</template>

<script>

import axios from "axios";

export default {
  name: 'Content',
  props: ['language', 'languages'],
  components: {
  },
  data: function() {
    return {
      texts: [],
      posts: []
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    load() {
      axios.get(this.$hostname + "text")
          .then(response => {
            if (response.data.success === true) {
              this.texts = response.data.texts;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });

      axios.get(this.$hostname + "post")
          .then(response => {
            if (response.data.success === true) {
              this.posts = response.data.posts;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }
  }
}
</script>

<style lang="css" scoped>

</style>
