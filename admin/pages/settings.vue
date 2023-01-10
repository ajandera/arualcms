<template>
  <v-container>
    <v-row v-if="setting !== null && setting.length > 0">
      <v-col
        v-for="(item, index) in setting"
        v-bind:key="index"
        cols="12"
        sm="6"
        md="6"
      >
        <v-text-field
          v-model="item.Value[language.toString()]"
          :counter="50"
          :label="item.Key"
          v-on:change="save(item)"
          required
        ></v-text-field>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import {Component, Prop, Vue, Watch} from 'nuxt-property-decorator'
import Setting from '~/model/Setting';
import IResponseSetting from "~/model/IResponseSetting";
import {namespace} from 'vuex-class';
import Permission from '~/model/Permission';

const snackbar = namespace('Snackbar');

@Component
export default class SettingsPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void
  @Prop() readonly language!: string;
  @Prop() readonly languages!: string[];
  @Prop() readonly permissions!: Permission[];

  setting: Setting[] = [];
  $axios: any;
  $nuxt: any;

  mounted() {
    this.checkPermission();
    this.load();
  }

  @Watch('$route.query')
  onPropertyChanged(value: any, oldValue: any) {
    if (value.siteId !== oldValue.siteId) {
      this.checkPermission();
      this.load();
    }
  }

  checkPermission() {
    const siteId = this.$route.query.siteId;
    const role = this.permissions.find((p: Permission) => p.SiteId === siteId)?.Role;
    if (role !== 'admin') {
      this.$nuxt.$options.router.push('/');
    }
  }

  async load() {
    this.$axios.get("/" + this.$route.query.siteId + "/setting")
      .then((response: IResponseSetting) => {
        if (response.data.success) {
          if (response.data.settings !== null) {
            response.data.settings.forEach(setting => {
              if (setting.Value.toString() !== "") {
                setting.Value = JSON.parse(setting.Value.toString())
              } else {
                setting.Value = this.createClearTranslationObject()
              }
            });
          }
          this.setting = response.data.settings;
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  save(setting: Setting) {
    this.$axios.put("/" + this.$route.query.siteId + "/setting/"+setting.Id, {
      "Id": setting.Id,
      "Value":  JSON.stringify(setting.Value)
    }, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseSetting) => {
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

  createClearTranslationObject() {
    const response: any = {};
    for (const lang of this.languages) {
      response[lang] = "";
    }
    return response
  }
}
</script>

