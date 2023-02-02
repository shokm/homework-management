<template>
  <div>
    <div>
      <h1>ユーザ登録</h1>
      <p>{{ message }}</p>
      <form @submit.prevent="registerUser">
        <div class="form-group">
          <label for="userName">ユーザー名:</label>
          <input v-model="user.screen_name" />
        </div>
        <div class="form-group">
          <label for="password">Password:</label>
          <input v-model="user.password" type="password" />
        </div>
        <button type="submit">登録</button>
      </form>
    </div>
    <div v-if="$auth.loggedIn">
      <button @click="$auth.logout()">Logout</button>
      <nuxt-link to="/">ログイン状態をチェック</nuxt-link>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import axios from '@nuxtjs/axios'
import auth from '@nuxtjs/auth-next'

export default Vue.extend({
  data() {
    return {
      user: {
        screen_name: '',
        password: ''
      },
      message: ''
    }
  },
  methods: {
    async registerUser() {
      this.$auth.logout()
      try {
        await this.$axios.post('/v1/auth/register', this.user)
        await this.$auth.loginWith('local', {
          data: this.user
        })
        // this.message = '登録成功'
      } catch (e) {
        this.message = e
      }
    }
  }
})
</script>
