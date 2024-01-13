<template>
  <div class="md:flex justify-center">
    <div
      class="bg-white shadow-xl rounded-xl md:w-auto w-11/12 md:p-20 p-5 md:mt-20 m-5"
    >
      <h1 class="flex justify-center mb-5 font-medium text-3xl">
        Studule - 宿題管理アプリ
      </h1>
      <h2 class="flex justify-center mb-5 font-medium text-2xl">ログイン</h2>
      <p>{{ message }}</p>
      <form @submit.prevent="loginUser">
        <div class="my-2">
          <label for="userName">ユーザー名</label>
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
            ログイン
          </button>
        </div>
        <div class="mt-4">
          <nuxt-link class="text-blue-500" to="/register">新規登録</nuxt-link>
        </div>
        <p class="mt-4">※デモユーザーのIDとパスワードは「testuser」です。</p>
      </form>
    </div>
    <div
      class="bg-white shadow-xl rounded-xl md:w-80 w-11/12 md:pt-20 p-5 md:mt-20 md:mx-10 m-5"
    >
      <h3 class="flex justify-center mb-5 font-medium text-2xl">
        {{ cmsData.title }}
      </h3>
      <div class="prose">
        <span v-html="cmsData.contents"></span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import $axios from '@nuxtjs/axios'
import $auth from '@nuxtjs/auth-next'

type CmsData = {
  title: String
  contents: String
  createdAt: String
  publishedAt: String
  revisedAt: String
  updatedAt: String
}

export default Vue.extend({
  data() {
    return {
      user: {
        screenName: '',
        password: ''
      },
      message: '',
      cmsData: {} as CmsData
    }
  },
  head: {
    bodyAttrs: {
      class: 'bg-gray-50'
    }
  },
  mounted() {
    this.$axios
      .$get(this.$config.MICROCMS_API_URL, {
        headers: { 'X-MICROCMS-API-KEY': this.$config.MICROCMS_API_KEY }
      })
      .then((response) => (this.cmsData = response))
  },
  methods: {
    async loginUser() {
      await this.$auth.logout()
      try {
        await this.$auth.loginWith('local', {
          data: this.user
        })
        // this.message = 'ログイン成功'
      } catch (e) {
        this.message = String(e)
      }
    }
  }
})
</script>
