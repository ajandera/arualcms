<template>
  <v-data-table
    :headers="headers"
    :items="sites"
    :items-per-page="50"
    class="elevation-1"
  >
    <template v-slot:top>
      <v-toolbar
        flat
      >
        <v-toolbar-title>Sites</v-toolbar-title>
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
                  v-model="siteObject.Name"
                  :counter="50"
                  label="Site"
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
                @click="save(siteObject)"
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
    <template v-slot:item.default="{ item }">
      <v-icon
        small
        class="mr-2"
        v-if="item.default === true"
      >
        mdi-check
      </v-icon>
    </template>
    <template v-slot:no-data>
    </template>
  </v-data-table>
</template>

<script lang="ts">
import {Component, Prop, Vue} from 'nuxt-property-decorator';
import IHeader from "~/model/IHeader";
import {namespace} from 'vuex-class';
import Site from '~/model/Site';
import IResponseSite from '~/model/IResponseSite';
import Permission from '~/model/Permission';

const snackbar = namespace('Snackbar');

@Component
export default class SitesPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void

  @Prop() readonly permissions!: Permission[];

  siteObject: Site = {
    Id: "",
    Name: ""
  };
  title: string = 'Create site'
  $axios: any;
  sites: Site[] = [];
  dialog: boolean = false;
  headers: IHeader[] = [
    {
      text: "Name",
      align: 'start',
      sortable: true,
      value: 'Name',
    },
    {text: "Api Token", value: 'ApiToken', sortable: false},
    {text: "Actions", value: 'actions', sortable: false}
  ];
  $nuxt: any;

  mounted(): void {
    this.checkPermission();
    this.load();
  }

  checkPermission() {
    const siteId = this.$route.query.siteId;
    const role = this.permissions.find((p: Permission) => p.SiteId === siteId)?.Role;
    if (role !== 'admin') {
      this.$nuxt.$options.router.push('/');
    }
  }

  create(): void {
    this.title = "New Site";
    this.dialog = true;
  }

  save(site: Site) {
    if (site.Id !== '') {
      this.$axios.put("/sites", site, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseSite) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
            this.load();
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
            this.load();
          }
        });
    } else {
      this.$axios.post("/sites", site, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseSite) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
            this.load();
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
            this.load();
          }
        });
    }
  }

  load() {
    this.dialog = false;
    this.$axios.get("/sites")
      .then((response: IResponseSite) => {
        if (response.data.success) {
          this.sites = response.data.sites;
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  edit(siteObject: Site) {
    this.title = siteObject.Name;
    this.siteObject = siteObject;
    this.dialog = true;
  }

  remove(site: Site) {
    this.$axios.delete("/sites/" + site.Id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseSite) => {
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

  close() {
    this.siteObject = {
      Id: "",
      Name: ""
    };
    this.dialog = false;
  }
}
</script>
