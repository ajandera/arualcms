<template>
  <v-row justify="center" align="center">

  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import Setting from '~/model/Setting';
import Message from "~/model/Message";
import IResponseSetting from "~/model/IResponseSetting";

@Component
export default class SettingsPage extends Vue {
    setting!: Setting[];
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
      this.$axios.put("setting", setting)
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

