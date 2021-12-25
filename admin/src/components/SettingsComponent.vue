<template>
  <div id="setting">
    <div class="row">
      <div class="col-9">
        <h1>Settings</h1>
      </div>
      <div class="col-3 text-right">
        <div class="btn-group mt-3" role="group" aria-label="Basic example">
          <button
              v-on:click="setLanguage(lang)"
              v-bind:class="{'btn btn-default': lang !== language, 'btn btn-primary': lang === language}" v-for="(lang, index) in languages" v-bind:key="index">{{ lang }}</button>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <hr>
      </div>
    </div>
    <div v-for="(item, index) in setting" class="row mt-2" v-bind:key="index">
      <div class="col-3">
        <p>{{ item.key }}</p>
      </div>
      <div class="col-8">
        <input type="text" v-model="item.value[language]" class="form-control" v-on:change="save(item)">
      </div>
    </div>
    <div class="row mt-4">
      <div class="col-11">
        <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
      </div>
    </div>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: 'SettingsComponent',
  props: ['language', 'languages'],
  components: {},
  data: function() {
    return {
      setting: [],
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user")
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
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    },
    save(setting) {
      axios.put(this.$hostname + "setting", setting, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
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
    }
  }
}
</script>
