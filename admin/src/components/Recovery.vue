<template>
  <modal name="recovery" height="auto" :clickToClose="false">
    <div class="modal-dialog">
      <!-- Modal content-->
      <div class="modal-content">
        <div class="modal-header">
          <h3 class="modal-title">arualCMS</h3>
        </div>
        <div class="modal-body">
          <div class="row">
            <div class="col-12">
              <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
              <div id="signin">
                <label for="password">Password {{ $route.params.username }}</label>
                <input type="password" v-model="password" class="form-control" id="password" required>
                <label for="passwordCheck">Password check <span v-if="passwordCheck.length > 0 && password !== passwordCheck" class="red">(Password not match!)</span></label>
                <input type="password" v-model="passwordCheck" class="form-control" id="passwordCheck" required>
                <button class="float-right btn btn-success mt-3" type="submit" v-on:click="recoveryPw()" :disabled="password !== passwordCheck">Set password</button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </modal>
</template>

<script>

import axios from "axios";

export default {
  name: 'Recovery',
  components: {},
  data() {
    return {
      messageClass: null,
      message: null,
      password: null,
      passwordCheck: ""
    }
  },
  mounted() {
    this.$modal.show('recovery');
  },
  methods: {
    recoveryPw() {
      const username = this.$route.params.username;
      axios.post(this.$hostname + "recovery", {username: username, password: this.password})
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
            }, 2000);
          }, (error) => {});
    }
  }
}
</script>

<style lang="css" scoped>
.modal-dialog {
  max-width: 800px;
  margin: 10px;
}
.modal-content {
  border: none;
}
.red {
  color: red;
}
</style>
