<template>
<div class="table-wrap">
    <div v-if="message" v-bind:class="messageClass">{{ message }}</div>
    <table class="table table-stripped mt-3">
      <thead>
        <tr>
          <th>#</th>
          <th>Title</th>
          <th>Published</th>
          <th>
            <button v-on:click="create()" type="button" class="btn btn-secondary btn-success float-right">
              <i class="fa fa-plus"></i>Create</button>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="post in posts" :key="post.id" v-on:click="edit(post.id)" class="actRow">
          <td>{{ post.id }}</td>
          <td>{{ post.title }}</td>
          <td>{{ post.published | formatDate}}</td>
          <td class="text-right">
            <button v-on:click.stop.prevent="remove(post.id)" type="button" class="btn btn-secondary btn-danger">
              <i class="fa fa-trash"></i>Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
    <modal name="form" :width="800" :height="800" :adaptive="true">
    <div class="modal-dialog">
      <!-- Modal content-->
      <div class="modal-content">
        <div class="modal-header">
          <h3 class="modal-title">{{ modalTitle }}</h3>
          <button @click="hide" type="button" class="close" data-dismiss="modal">×</button>
        </div>
        <div class="modal-body">
          <div class="row">
            <div class="col-12">
              <div v-if="error" class="danger">{{ error }}</div>
            </div>
          </div>
          <div class="row">
            <div class="col-12">
              <Editor v-bind:post="post" />
            </div>
          </div>
        </div>
      </div>
    </div>
  </modal>
</div>
</template>

<script>

import axios from "axios";
import Editor from "@/components/Editor";

export default {
  name: 'Posts',
  components: {
    Editor
  },
  data() {
    return {
      messageClass: null,
      message: null,
      loggedUser: window.localStorage.getItem("user"),
      posts: [],
      post: {},
      modalTitle: "",
      error: ""
    }
  },
  mounted() {
    axios.get(this.$hostname + "post")
        .then(response => {
          if (response.data.success === true) {
            for (let index in response.data.posts) {
              response.data.posts[index].published = new Date(response.data.posts[index].published);
              this.posts.push(response.data.posts[index]);
            }
            console.log(this.posts);
          } else {
            this.message = response.data.error;
            this.messageClass = 'danger';
          }
        });
  },
  methods: {
    edit(id) {
      this.post = this.posts.filter(x => x.id === id)[0];
      this.modalTitle = this.post.title;
      this.show();
    },
    remove(id) {
      console.log('remove');
    },
    create() {
      this.post = {};
      this.show();
    },
    show() {
      this.$modal.show('form')
    },
    hide() {
      this.$modal.hide('form')
    }
  },
}
</script>

<style lang="css" scoped>
.table-wrap {
  width: 100%;
}
.actRow {
  cursor: pointer;
}
.actRow:hover td {
  background: #f8f8f8;
}
.modal-dialog {
  max-width: 800px;
  margin: 10px;
}
.modal-content {
  border: none;
}
</style>
