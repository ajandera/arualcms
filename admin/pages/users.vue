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
                  v-model="user.Name"
                  :counter="30"
                  label="Name"
                  required
                ></v-text-field>
                <v-text-field
                  v-model="user.Username"
                  :counter="30"
                  label="Username"
                  required
                ></v-text-field>
                <v-text-field
                  v-model="user.Password"
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
        <v-dialog
          v-model="permisionDialog"
          max-width="1000px"
        >
          <v-card>
            <v-card-title>
              <span class="text-h5">{{ title }}</span>
            </v-card-title>
            <v-card-text>
              <v-container>
                <v-row v-for="(perm, index) in user.Permission" :key="index">
                  <v-col cols="6">
                    {{ sites.find(s => s.Id === perm.SiteId).Name }}
                  </v-col>
                  <v-col cols="6">
                    <v-form ref="form">
                      <v-select
                        v-model="perm.Role"
                        :items="['admin', 'author']"
                        label="Select"
                        single-line
                        @change="updatePermision(perm)"
                      ></v-select>
                    </v-form>
                  </v-col>
                </v-row>
              </v-container>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn
                color="blue darken-1"
                text
                @click="closePermissionDialog"
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
        mdi-pencil
      </v-icon>
      <v-icon
        small
        class="mr-2"
        @click="openPermission(item)"
      >
        mdi-account-lock
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
import {Component, Vue, Prop} from 'nuxt-property-decorator'
import User from '~/model/User';
import IResponseUsers from '~/model/IResponseUsers';
import Permission from '~/model/Permission';
import IHeader from "~/model/IHeader";
import {namespace} from 'vuex-class';
import Site from '~/model/Site';

const snackbar = namespace('Snackbar');

@Component
export default class UsersPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void

  @Prop() readonly sites!: Site[];
  @Prop() readonly permissions!: Permission[];

  isEdit: boolean = false;
  title?: string = 'Users';
  users: User[] = [];
  dialog: boolean = false;
  permisionDialog: boolean = false;
  user: User = {
    Username: "",
    Password: "",
    Id: "",
    ParentId: ""
  };
  headers: IHeader[] = [
    {
      text: "Name",
      align: 'start',
      sortable: true,
      value: 'Name',
    },
    {
      text: "Username",
      sortable: true,
      value: 'Username',
    },
    {text: "Actions", value: 'actions', sortable: false}
  ];
  $axios: any;
  $auth: any;
  $nuxt: any;

  mounted() {
    this.user = {
      Username: "",
      Password: "",
      Id: "",
      ParentId: this.$auth.user.Id
    };
    this.checkPermission();
    this.load();
  }

  checkPermission() {
    const siteId = this.$route.query.siteId;
    const role = this.permissions?.find((p: Permission) => p.SiteId === siteId)?.Role;
    if (role !== 'admin') {
      this.$nuxt.$options.router.push('/');
    }
  }

  edit(user: User) {
    this.user = user;
    this.title = this.user.Name;
    this.isEdit = true;
    this.dialog = true;
  }

  remove(user: User) {
    this.$axios.delete("/" + this.$route.query.siteId + "/user/" + user.Id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseUsers) => {
        if (response.data.success) {
          this.updateText(response.data.message);
          this.updateColor('green')
          this.updateShow(true);
          this.load();
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  openPermission(user: User) {
    this.user = user;
    this.title = this.user.Name;
    this.permisionDialog = true;
  }

  save() {
    if (this.isEdit) {
      this.$axios.put("/" + this.$route.query.siteId + "/users", this.user, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseUsers) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
          }
          this.load();
        });
    } else {
      this.$axios.post("/" + this.$route.query.siteId + "/users", this.user, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseUsers) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
          }
          this.load();
        });
    }
  }

  load() {
    this.users = [];
    this.$axios.get("/" + this.$route.query.siteId + "/users")
      .then((response: IResponseUsers) => {
        this.dialog = false;
        if (response.data.success === true) {
          response.data.users.map((obj: User) =>
            {
              obj.Permission = response.data.permissions.filter((item: Permission) => item.UserId === obj.Id)
            })
          this.users = response.data.users;
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  close() {
    this.dialog = false;
    this.isEdit = false;
    this.user = {
      Username: "",
      Password: "",
      Id: "",
      ParentId: this.$auth.user.Id
    };
  }

  closePermissionDialog() {
    this.permisionDialog = false;
  }

  updatePermision(permission: Permission) {
    this.$axios.put("/" + permission.SiteId + "/" + permission.SiteId + "/permission", permission, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseUsers) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
          }
        });
  }
}
</script>

