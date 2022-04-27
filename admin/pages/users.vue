<template>
  <v-data-table
    :headers="headers"
    :items="users"
    :items-per-page="50"
    class="elevation-1"
  >
  <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>Users</v-toolbar-title>
        <v-divider
          class="mx-4"
          inset
          vertical
        ></v-divider>
        <v-spacer></v-spacer>
        <v-dialog
          v-model="dialog"
          max-width="1000px"
        >
          <v-card>
            <v-card-title>
              <span class="text-h5">{{ title }}</span>
            </v-card-title>
            <v-card-text>
              <v-container>
                
              </v-container>
            </v-card-text>

            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="blue darken-1"
                text
                @click="close"
              >
                Close
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-toolbar>
    </template>
    <template v-slot:item.actions="{ item }">
      <v-icon
        small
        class="mr-2"
        @click="edit(item)"
      >
        mdi-user
      </v-icon>
    </template>
    <template v-slot:no-data>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'

@Component
export default class UsersPage extends Vue {
    isEdit: bool = false;
    alertType: string = "";
    message: string = "";
    users: User[];
    user: User = {
        username: "",
        password: ""
      };
    modalTitle: "";
    error: ""

    mounted() {
        this.load();
    }

    edit(user) {
      this.user = user;
      this.modalTitle = this.user.username;
      this.isEdit = true;
      this.show();
    }

    remove(id) {
      axios.delete("users/" + id)
          .then(response => {
            if (response.data.success) {
              this.message = response.data.message;
              this.messageClass = "alert alert-success";
              this.load();
            } else {
              this.message = response.data.message;
              this.messageClass = "alert alert-danger";
            }
            this.hide();
            setTimeout(() => {
              this.message = null;
              this.messageClass = null;
            }, 2000);
          }, (error) => {
            if (error.response.status === 401) {
              window.localStorage.removeItem("userId");
              window.localStorage.removeItem("user");
              window.localStorage.removeItem("jwt");
              this.$router.push({name: 'posts'});
            }
          });
    }

    create() {
      this.isEdit = false;
      this.user = {
        username: "",
        password: ""
      };
      this.show();
    }

    save() {
      if (this.isEdit === true) {
        axios.put(
            this.$hostname + "users",
            this.user,
            {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}}
            ).then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }
              this.load();
              this.hide();
              setTimeout(() => {
                this.message = null;
                this.messageClass = null;
              }, 2000);
            }, (error) => {
              console.log(error);
              if (error.response.status === 401) {
                window.localStorage.removeItem("userId");
                window.localStorage.removeItem("user");
                window.localStorage.removeItem("jwt");
                this.$router.push({name: 'posts'});
              }
            });
      } else {
        axios.post(this.$hostname + "users", this.user, {headers: {Authorization: "Bearer " + window.localStorage.getItem('jwt')}})
            .then(response => {
              if (response.data.success) {
                this.message = response.data.message;
                this.messageClass = "alert alert-success";
              } else {
                this.message = response.data.message;
                this.messageClass = "alert alert-danger";
              }
              this.load();
              this.hide();
              setTimeout(() => {
                this.message = null;
                this.messageClass = null;
              }, 2000);
            }, (error) => {
              if (error.response.status === 401) {
                window.localStorage.removeItem("userId");
                window.localStorage.removeItem("user");
                window.localStorage.removeItem("jwt");
                this.$router.push({name: 'posts'});
              }
            });
      }
    }

    load() {
      this.users = [];
      axios.get("users")
          .then(response => {
            if (response.data.success === true) {
              this.users = response.data.users;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }
}
</script>

