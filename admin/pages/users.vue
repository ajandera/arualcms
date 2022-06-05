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
import User from '~/model/User';
import Message from '~/model/Message';
import IResponseUsers from '~/model/IResponseUsers';

@Component
export default class UsersPage extends Vue {
    isEdit: boolean = false;
    message: Message = {text: "", class: ""};
    users!: User[];
    user: User = {
        username: "",
        password: ""
      };
    modalTitle: string = "";
    $axios: any;

    mounted() {
        this.load();
    }

    edit(user: User) {
      this.user = user;
      this.modalTitle = this.user.username;
      this.isEdit = true;
    }

    remove(id: string) {
      this.$axios.delete("/users/" + id)
          .then((response: IResponseUsers) => {
            if (response.data.success) {
              this.message = {
                text: response.data.message,
                class: "success"
              };
              this.load();
            } else {
              this.message = {
                text: response.data.message,
                class: "danger"
              };
            }
          });
    }

    create() {
      this.isEdit = false;
      this.user = {
        username: "",
        password: ""
      };
    }

    save() {
      if (this.isEdit) {
        this.$axios.put("/users", this.user)
            .then((response: IResponseUsers) => {
              if (response.data.success) {
                this.message = {
                    text: response.data.message,
                    class: "success"
                };
              } else {
                this.message = {
                    text: response.data.message,
                    class: "danger"
                };
              }
              this.load();
            });
      } else {
        this.$axios.post("/users", this.user)
            .then((response: IResponseUsers) => {
              if (response.data.success) {
                this.message = {
                    text: response.data.message,
                    class: "success"
                };
              } else {
                this.message = {
                    text: response.data.message,
                    class: "danger"
                };
              }
              this.load();
            });
      }
    }

    load() {
      this.users = [];
      this.$axios.get("/users")
          .then((response: IResponseUsers) => {
            if (response.data.success === true) {
              this.users = response.data.users;
            } else {
              this.message = {
                text: response.data.message,
                class: "danger"
              };
            }
          });
    }
}
</script>

