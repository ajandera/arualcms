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
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              color="primary"
              dark
              class="mb-2"
              v-bind="attrs"
              v-on="on"
            >
              Add
            </v-btn>
          </template>
          <v-card>
            <v-card-title>
              <span class="text-h5">{{ title }}</span>
            </v-card-title>
            <v-card-text>
              <v-container>
                <v-text-field
                  v-model="user.username"
                  :counter="30"
                  label="Username"
                  required
                ></v-text-field>
                <v-text-field
                  v-model="user.password"
                  :counter="20"
                  type="password"
                  label="Password"
                  required
                ></v-text-field>
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
              <v-btn
                color="primary"
                dark
                @click="save"
              >
                Save
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
        mdi-pencil
      </v-icon>
      <v-icon
        small
        class="mr-2"
        @click="remove(item)"
      >
        mdi-delete
      </v-icon>
    </template>
    <template v-slot:no-data>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import {Component, Vue} from 'nuxt-property-decorator'
import User from '~/model/User';
import Message from '~/model/Message';
import IResponseUsers from '~/model/IResponseUsers';
import IHeader from "~/model/IHeader";

@Component
export default class UsersPage extends Vue {
  isEdit: boolean = false;
  title: string = 'Users';
  message: Message = {text: "", class: ""};
  users: User[] = [];
  dialog: boolean = false;
  user: User = {
    username: "",
    password: "",
    _id: {
      $oid: ""
    }
  };
  headers: IHeader[] = [
    {
      text: "Username",
      align: 'start',
      sortable: true,
      value: 'username',
    },
    {text: "Actions", value: 'actions', sortable: false}
  ];
  $axios: any;

  mounted() {
    this.load();
  }

  edit(user: User) {
    this.user = user;
    this.title = this.user.username;
    this.isEdit = true;
    this.dialog = true;
  }

  remove(user: User) {
    this.$axios.delete("/users/" + user._id.$oid, {headers: {'Content-Type': "application/json;charset=utf-8"}})
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
      password: "",
      _id: {
        $oid: ""
      }
    };
  }

  save() {
    if (this.isEdit) {
      this.$axios.put("/users", this.user, {headers: {'Content-Type': "application/json;charset=utf-8"}})
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
      this.$axios.post("/users", this.user, {headers: {'Content-Type': "application/json;charset=utf-8"}})
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
        this.dialog = false;
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

  close() {
    this.dialog = false;
  }
}
</script>

