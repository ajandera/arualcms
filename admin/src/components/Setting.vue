<template>
  <div id="setting" class="container">
    <div v-for="(item, index) in setting" class="row mt-2" v-bind:key="index">
      <div class="col-1">
        <button v-if="index === Object.keys(setting).length - 1" v-on:click="add" class="btn btn-warning"><font-awesome-icon icon="plus" /></button>
      </div>
      <div class="col-4">
        <input type="text" v-model="item.key" class="form-control">
      </div>
      <div class="col-5">
        <input type="text" v-model="item.value" class="form-control">
      </div>
      <div class="col-2">
        <div class="float-right">
          <button v-on:click="remove(item)" class="btn btn-danger"><font-awesome-icon icon="times" /></button>
        </div>
      </div>
    </div>
    <div class="row mt-2">
      <div class="col-6">
        <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
      </div>
      <div class="col-6">
        <div class="float-right">
          <button v-on:click="save" class="btn btn-success">Save</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import axios from "axios";

export default {
  name: 'Setting',
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
    let local = window.localStorage.getItem("setting");
    if (local) {
      this.setting = JSON.parse(local);
    }
  },
  methods: {
    save() {
      if (this.loggedUser) {
        axios.post("http://loclahost:9000/api/v1/setting/save", {"setting": this.setting})
            .then(response => {
              this.message = response.data;
            });
      } else {
        window.localStorage.setItem("setting", JSON.stringify(this.setting));
        this.message = "Settings saved."
        this.messageClass = "alert alert-success";
        setTimeout(() => {
          this.message = null;
          this.messageClass = null;
        }, 1000);
      }
    },
    add() {
      this.setting.push({'key': '', 'value': ''});
    },
    remove(item) {
      var index = this.setting.indexOf(item);
      this.setting.splice(1, index);
    }
  }
}
</script>
