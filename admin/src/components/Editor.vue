<template>
  <div id="editor">
    <div class="row mt-3">
      <div class="col-12">
        <input type="text" class="form-control" v-model="post.title" />
        <vue-summernote ref="editer"></vue-summernote>
      </div>
    </div>
    <div class="row mt-3">
      <div class="col-6 mt-3">
        <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
      </div>
      <div class="col-6 mt-3">
        <div class="float-right">
          <button v-on:click="save" class="btn btn-lg btn-success">Save</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import axios from "axios";
import VueSummernote from 'vue-summernote'

export default {
  name: 'Editor',
  components: {
    VueSummernote
  },
  data() {
    return {
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user"),
      post: {}
    }
  },
  mounted() {
    let localApp = window.localStorage.getItem("app");
    if (localApp) {
      this.code = JSON.parse(localApp);
    }
    setInterval(() => {
      this.save();
    }, 5000);
  },
  methods: {
    save() {
      if (this.loggedUser) {
        let id = window.localStorage.getItem("userId");
        axios.post("http://localhost:9000/api/v1/app/save", {"code": this.code, "user": id, "name": this.loggedUser})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }

              setTimeout(() => {
                this.message = null;
                this.messageClass = null;
              }, 1000);
            });
      } else {
        window.localStorage.setItem("app", JSON.stringify(this.code));
        this.message = "Application saved."
        this.messageClass = "alert alert-success";
        setTimeout(() => {
          this.message = null;
          this.messageClass = null;
        }, 1000);
      }
    }
  }
}
</script>

<style lang="css" scoped>

#editor {

}

</style>
