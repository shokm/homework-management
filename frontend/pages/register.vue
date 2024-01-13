<template>
  <div class="flex justify-center">
    <div
      class="bg-white shadow-xl rounded-xl md:w-auto w-11/12 md:p-20 p-5 md:m-20 m-5"
    >
      <h1 class="flex justify-center mb-10 font-medium text-4xl">新規登録</h1>
      <p>{{ message }}</p>
      <form @submit.prevent="registerUser">
        <div class="my-2">
          <label for="userName">ユーザー名（半角英数字）</label>
          <br />
          <input
            v-model="user.screenName"
            type="text"
            class="border border-DEFAULT md:w-96 w-full max-h-16 shadow-inner rounded-xl bg-transparent text-current p-5"
          />
        </div>
        <div class="my-2">
          <label for="password">パスワード</label>
          <br />
          <input
            v-model="user.password"
            type="password"
            class="border border-DEFAULT md:w-96 w-full max-h-16 shadow-inner rounded-xl bg-transparent text-current p-5"
          />
        </div>
        <div class="mt-2">
          <button
            type="submit"
            class="md:w-96 w-full max-h-16 rounded-xl bg-blue-600 text-center text-white p-5"
          >
            ユーザー登録
          </button>
        </div>
        <div class="mt-4">
          <nuxt-link class="text-blue-500" to="/login">ログイン</nuxt-link>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import $axios from '@nuxtjs/axios'
import $auth from '@nuxtjs/auth-next'

export default Vue.extend({
  data() {
    return {
      user: {
        screenName: '',
        password: ''
      },
      message: ''
    }
  },
  head: {
    bodyAttrs: {
      class: 'bg-gray-50'
    }
  },
  methods: {
    async registerUser() {
      await this.$auth.logout()
      try {
        await this.$axios.$post('/v1/auth/register', this.user)
        await this.$auth.loginWith('local', {
          data: this.user
        })
        // this.message = '登録成功'
      } catch (e) {
        this.message = String(e)
      }
    }
  }
})
</script>
