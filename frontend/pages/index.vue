<template>
  <div>
    <h2>ログイン状態:{{ $auth.loggedIn }}</h2>
    <p>{{ $auth.user }}</p>
    <button @click="$auth.logout()">Logout</button>
    <nuxt-link to="/tasks">タスク一覧</nuxt-link>
    <nuxt-link to="/task/0">タスク作成</nuxt-link>
    <nuxt-link to="/login">ログイン画面に戻る</nuxt-link>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  middleware: 'auth'
})
</script>
