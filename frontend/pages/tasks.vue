<template>
  <div>
    <div v-for="task in tasks.tasks" :key="task.task_id">
      <p>
        {{ task.task_id }} {{ task.task_name }}
        <nuxt-link :to="'/task/' + task.task_id">編集</nuxt-link>
      </p>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import $axios from '@nuxtjs/axios'
import $auth from '@nuxtjs/auth-next'
import $dayjs from '@nuxtjs/dayjs'

export default Vue.extend({
  middleware: 'auth',
  data() {
    return {
      tasks: {}
    }
  },
  created() {
    this.$axios.$get('/v1/tasks').then((response) => (this.tasks = response))
  },
  methods: {}
})
</script>
