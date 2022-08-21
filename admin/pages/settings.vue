<template>
  <v-container>
    <v-row>
      <v-col
        v-for="(item, index) in setting"
        v-bind:key="index"
        cols="12"
        sm="6"
        md="6"
      >
        <v-text-field
          v-model="item.value[language]"
          :counter="50"
          :label="item.key"
          v-on:change="save(item)"
          required
        ></v-text-field>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import {Component, Prop, Vue} from 'nuxt-property-decorator'
import Setting from '~/model/Setting';
import IResponseSetting from "~/model/IResponseSetting";
import {namespace} from 'vuex-class';

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
  setting: Setting[] = [];
  $axios: any;

  mounted() {
    this.load();
  }

  load() {
    this.$axios.get("/setting")
      .then((response: IResponseSetting) => {
        if (response.data.success) {
          this.setting = response.data.settings;
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  save(setting: Setting) {
    this.$axios.put("setting", setting, {headers: {'Content-Type': "application/json;charset=utf-8"}})
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
}
</script>

