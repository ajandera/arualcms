<template>
  <div>
    <div v-for="(item, index) in texts" class="row mt-2" v-bind:key="index">
    <div class="col-sm-1 col-xs-12">
      <v-btn
        v-if="index === Object.keys(texts).length - 1"
        @click="add"
        class="primary"
        dark
      >
        <v-icon
          small
          class="mr-2"
        >
          mdi-plus
        </v-icon>
      </v-btn>
    </div>
    <div class="col-sm-3 col-xs-6">
      <v-text-field
        v-model="item.key"
        :counter="30"
        label="Key"
        required
        @change="save(item)"
      ></v-text-field>
    </div>
    <div class="col-sm-7 col-xs-6">
      <v-textarea
        solo
        name="input-7-4"
        label="Solo textarea"
        v-model="item.value[language]"
        @change="save(item)"
      ></v-textarea>
    </div>
    <div class="col-sm-1 col-xs-12">
      <div class="float-right">
        <div class="btn-group" role="group" aria-label="Basic example">
          <v-dialog
            v-model="dialog[item.key]"
            max-width="1000px"
          >
            <template v-slot:activator="{ on, attrs }">
              <v-btn
                class="primary"
                dark
                v-bind="attrs"
                v-on="on"
              >
                <v-icon
                  small
                  class="mr-2"
                >
                  mdi-pencil
                </v-icon>
              </v-btn>
            </template>
            <v-card>
              <v-card-title>
                <span class="text-h5">{{ item.key[language] }}</span>
              </v-card-title>
              <v-card-text>
                <v-container>
                  <quill-editor
                    :ref="item.key['en']"
                    v-model="item.value[language]"
                    :options="editorOption"
                    @blur="onEditorBlur($event)"
                    @focus="onEditorFocus($event)"
                    @ready="onEditorReady($event)"
                  />
                </v-container>
              </v-card-text>

              <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn
                  color="primary"
                  dark
                  @click="save(item)"
                >
                  Save
                </v-btn>
                <v-btn
                  color="blue darken-1"
                  text
                  @click="close(item.key)"
                >
                  Close
                </v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
          <v-btn @click="remove(item)" class="error">
            <v-icon
              small
              class="mr-2"
            >
              mdi-delete
            </v-icon>
          </v-btn>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>

<script lang="ts">
import {Component, Prop, Vue} from 'nuxt-property-decorator';
import Text from '~/model/Text';
import Message from "~/model/Message";
import IResponseTexts from "~/model/IResponseTexts";

@Component
export default class TextsPage extends Vue {
    @Prop() readonly languages!: string[];
    @Prop() readonly language!: string;
    texts: Text[] = [];
    text?: Text;
    editorOption: any = {};
    message: Message = {class: "", text: ""};
    $axios: any;
    content: any;
    dialog: any = {};

    mounted() {
        this.load();
    }

    save(text: Text) {
      if (text._id.$oid !== '') {
        this.$axios.put("text", text, {headers: {'Content-Type': "application/json;charset=utf-8"}})
            .then((response: IResponseTexts) => {
              if (response.data.success) {
                this.message.text = response.data.message;
                this.message.class = "alert alert-success";
              } else {
                this.message.text = response.data.error;
                this.message.class = 'danger';
              }
            });
      } else {
        this.$axios.post("text", text)
            .then((response: IResponseTexts) => {
              if (response.data.success) {
                this.message.text = response.data.message;
                this.message.class = "alert alert-success";
                this.load();
              } else {
                this.message.text = response.data.error;
                this.message.class = 'danger';
              }
            });
      }
    }

    load() {
      this.$axios.get("/text")
          .then((response: IResponseTexts) => {
            if (response.data.success === true) {
              this.texts = response.data.texts;
              response.data.text.forEach((item: Text) => this.dialog[item.key] = false)
            } else {
              this.message.text = response.data.error;
              this.message.class = 'danger';
            }
          });
    }

    add() {
      let text: Text = {
        key: '',
        value: {},
        _id: {$oid: ""}
      };
      for (const lang of this.languages) {
        text.value[lang] = "";
      }
      this.texts.push(text);
    }

    remove(item: Text) {
      this.$axios.delete("/text/" + item._id.$oid, {headers: {'Content-Type': "application/json;charset=utf-8"}})
          .then((response: IResponseTexts) => {
            if (response.data.success) {
              this.message.text = response.data.message;
              this.message.class = "alert alert-success";
              this.load();
            } else {
              this.message.text = response.data.message;
              this.message.class = "alert alert-danger";
            }
          });
    }

    onEditorBlur(quill: any) {
    }

    onEditorFocus(quill: any) {
    }
    onEditorReady(quill: any) {
    }

    onEditorChange(quill: any, html: any, text: any ) {
      this.content = html
    }

    close(key: string) {
      this.dialog[key] = false;
    }
}
</script>

