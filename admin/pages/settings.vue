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
import Message from "~/model/Message";
import IResponseSetting from "~/model/IResponseSetting";

@Component
export default class SettingsPage extends Vue {
  @Prop() readonly language!: string;
  setting: Setting[] = [];
  $axios: any;
  message: Message = {class: "", text: ""};

  mounted() {
    this.load();
  }

  load() {
    this.$axios.get("/setting")
      .then((response: IResponseSetting) => {
        if (response.data.success) {
          this.setting = response.data.settings;
        } else {
          this.message.text = response.data.error;
          this.message.class = 'danger';
        }
      });
  }

  save(setting: Setting) {
    this.$axios.put("setting", setting, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseSetting) => {
        if (response.data.success) {
          this.message.text = response.data.message;
          this.message.class = "alert alert-success";
        } else {
          this.message.text = response.data.error;
          this.message.class = 'danger';
        }
      });
  }
}
</script>

