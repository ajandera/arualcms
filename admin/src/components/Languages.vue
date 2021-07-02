<template>
  <div id="texts">
    <div class="row">
      <div class="col-11">
        <h1>Languages</h1>
      </div>
    </div>
    <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
    <table class="table table-stripped mt-3">
      <thead>
      <tr>
        <th>Code</th>
        <th>Name</th>
        <th></th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="language in languages" :key="language.key" v-on:click="edit(language.id)" class="actRow">
        <td>{{ language.key }}</td>
        <td>{{ language.value }}</td>
        <td class="text-right">
          <button v-on:click.stop.prevent="remove(language.key)" type="button" class="btn btn-secondary btn-danger">
            <i class="fa fa-trash"></i>Delete</button>
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: 'Languages',
  components: {},
  data: function() {
    return {
      languages: [],
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user"),
      error: ""
    }
  },
  mounted() {
    this.load();
  },
  methods: {
    save() {
      axios.put(this.$hostname + "languages", this.languages, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
          .then(response => {
            if (response.data.success) {
              this.message = response.data.message;
              this.messageClass = "alert alert-success";
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          }, (error) => {
            if (error.response.status === 401) {
              window.localStorage.removeItem("userId");
              window.localStorage.removeItem("user");
              window.localStorage.removeItem("jwt");
              this.loggedUser = false;
              window.location.reload();
            }
          });
    },
    load() {
      axios.get(this.$hostname + "languages")
          .then(response => {
            if (response.data.success === true) {
              this.languages = response.data.languages;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    },
    add() {
      this.languages.push({'key': '', 'value': ''});
    },
    remove(item) {
      let index = this.languages.indexOf(item);
      this.languages.splice(index, 1);
    }
  }
}
</script>

<style lang="css" scoped>

</style>
