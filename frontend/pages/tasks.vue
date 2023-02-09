<template>
  <div class="flex justify-center">
    <div class="m-4 w-full">
      <div class="flex mt-8">
        <h1 class="font-medium text-3xl">課題一覧</h1>
        <nuxt-link
          to="/subjects/tasks/"
          class="ml-4 mt-1 font-medium text-2xl text-gray-400"
          >教科ごと</nuxt-link
        >
      </div>
      <h2 class="mt-8 font-medium text-xl">今日まで</h2>
      <div v-for="task in tasks.tasks" :key="task.task_id">
        <div
          :to="'/task/' + task.task_id"
          class="flex items-center mt-3 bg-white border border-DEFAULT shadow-lg rounded-xl"
        >
          <button
            onclick="alert('完了ボタン押下')"
            class="flex items-center justify-center m-3 mr-2 w-12 h-12 bg-gray-600 text-white shadow-lg rounded-lg"
          >
            ☑︎
          </button>
          <nuxt-link :to="'/task/' + task.task_id" class="p-3 pl-2 leading-6">
            <p>{{ task.task_id }} {{ task.task_name }}</p>
            <p>
              教科：{{ task.subject_name }} 期限：
              <span
                v-if="
                  !$dayjs($dayjs(task.deadline_at).format('YYYY-MM-DD')).isSame(
                    $dayjs('1-01-01')
                  )
                "
              >
                {{ $dayjs(task.deadline_at).format('YYYY/MM/DD hh:mm') }}
              </span>
              <span v-else>---/--/--</span>
            </p>
          </nuxt-link>
        </div>
      </div>
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
