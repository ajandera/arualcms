<template>
  <div>
    <div class="col-sm-1 col-xs-12" v-if="texts.length === 0">
      <v-btn
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
          v-model="item.Key"
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
          v-model="item.Value[language]"
          @change="save(item)"
        ></v-textarea>
      </div>
      <div class="col-sm-1 col-xs-12">
        <div class="float-right">
          <div class="btn-group" role="group" aria-label="Basic example">
            <v-dialog
              v-model="dialog[item.Key]"
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
                  <span class="text-h5">{{ item.Key[language] }}</span>
                </v-card-title>
                <v-card-text>
                  <v-container>
                    <quill-editor
                      :ref="item.Key['en']"
                      v-model="item.Value[language]"
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
                    color="blue darken-1"
                    text
                    @click="close(item.key)"
                  >
                    Close
                  </v-btn>
                  <v-btn
                    color="primary"
                    dark
                    @click="save(item)"
                  >
                    Save
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
import IResponseTexts from "~/model/IResponseTexts";
import {namespace} from 'vuex-class';

const snackbar = namespace('Snackbar');

@Component
export default class TextsPage extends Vue {
  @snackbar.Action
  public updateText!: (newText: string) => void

  @snackbar.Action
  public updateColor!: (newColor: string) => void

  @snackbar.Action
  public updateShow!: (newShow: boolean) => void
  @Prop() readonly languages!: string[];
  @Prop() readonly language!: string;
  texts: Text[] = [];
  text?: Text;
  editorOption: any = {};
  $axios: any;
  content: any;
  dialog: any = {};

  mounted() {
    this.load();
  }

  save(text: Text) {
    text.Value = JSON.stringify(text.Value);
    if (text.Id !== '') {
      this.$axios.put('/' + this.$route.query.siteId +"/text", text, {headers: {'Content-Type': "application/json;charset=utf-8"}})
        .then((response: IResponseTexts) => {
          if (response.data.success) {
            this.updateText(response.data.message);
            this.updateColor('green')
            this.updateShow(true);
          } else {
            this.updateText(response.data.message);
            this.updateColor('red')
            this.updateShow(true);
          }
          this.close(text.Key);
        });
    } else {
      text.Id = "";
      this.$axios.post('/' + this.$route.query.siteId +"/text", text)
        .then((response: IResponseTexts) => {
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
          this.close(text.Key);
        });
    }
  }

  load() {
    this.$axios.get('/' + this.$route.query.siteId +"/text")
      .then((response: IResponseTexts) => {
        if (response.data.success === true) {
          if (response.data.texts !== null) {
            response.data.texts.forEach((text: any) => {
              if (text.Value !== "") {
                text.Value = JSON.parse(text.Value.toString())
              }
            });
            this.texts = response.data.texts;
          }
        } else {
          this.updateText(response.data.message);
          this.updateColor('red')
          this.updateShow(true);
        }
      });
  }

  add() {
    let text: Text = {
      Key: '',
      Value: {},
      Id: ""
    };
    for (const lang of this.languages) {
      text.Value[lang] = "";
    }
    this.texts.push(text);
  }

  remove(item: Text) {
    this.$axios.delete('/' + this.$route.query.siteId +"/text/" + item.Id, {headers: {'Content-Type': "application/json;charset=utf-8"}})
      .then((response: IResponseTexts) => {
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

  onEditorBlur(quill: any) {
  }

  onEditorFocus(quill: any) {
  }

  onEditorReady(quill: any) {
  }

  onEditorChange(quill: any, html: any, text: any) {
    this.content = html
  }

  close(key: string) {
    this.dialog[key] = false;
  }
}
</script>

