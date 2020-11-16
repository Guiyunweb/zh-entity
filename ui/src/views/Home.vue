<template>
  <div>
    <v-app-bar color="deep-purple accent-4" dense dark>
      <v-app-bar-nav-icon></v-app-bar-nav-icon>
      <v-toolbar-title>中文转实体类工具</v-toolbar-title>
      <v-spacer></v-spacer>
    </v-app-bar>
    <div style="margin: 50px 10%">
      <v-row>
        <v-col>
          <v-textarea
              auto-grow
              outlined
              rows="20"
              v-model="body"
              label="请输入需要转换的实体类"
          ></v-textarea>
        </v-col>
        <v-col>
          <v-textarea
              auto-grow
              outlined
              v-model="translate"
              rows="20"
              disabled
              label="请输入需要转换的实体类"
          ></v-textarea>
        </v-col>
      </v-row>
      <v-btn depressed color="primary" @click="subJava()">
        Java实体类
      </v-btn>
    </div>
  </div>
</template>

<script>
import api from '@/api'

export default {
  name: 'Home',
  data: () => ({
    body: '',
    translate: ''
  }),
  methods: {
    subJava() {
      api.SYS_USER_LOGIN({body:this.body}).then(res => {
        this.translate  = res
        console.log('res',res)
      })
    }
  }
}
</script>
