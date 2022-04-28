<template>
  <v-row justify="center" align="center">
   
  </v-row>
</template>

<script lang="ts">
import { Component, Vue } from 'nuxt-property-decorator'
import IResponse from '~/model/IResponse';
import Setting from '~/model/Setting';

@Component
export default class SettingsPage extends Vue {
    setting: Setting[];
 
    mounted() {
        this.load();
    }

    load() {
      this.$axios.get("setting")
          .then((response: IResponse) => {
            if (response.data.success === true) {
              this.setting = response.data.settings;
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }

    save(setting) {
      this.$axios.put("setting", setting)
          .then((response: IResponse) => {
            if (response.data.success) {
              this.message = response.data.message;
              this.messageClass = "alert alert-success";
            } else {
              this.message = response.data.error;
              this.messageClass = 'danger';
            }
          });
    }
}
</script>

