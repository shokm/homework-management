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
      <div v-if="tasksExpired.length">
        <h2 class="mt-8 font-medium text-xl">期限切れ</h2>
        <TaskCard :tasks="tasksExpired"></TaskCard>
      </div>
      <div v-if="tasksToday.length">
        <h2 class="mt-8 font-medium text-xl">今日まで</h2>
        <TaskCard :tasks="tasksToday"></TaskCard>
      </div>
      <div v-if="tasksTomorrow.length">
        <h2 class="mt-8 font-medium text-xl">明日まで</h2>
        <TaskCard :tasks="tasksTomorrow"></TaskCard>
      </div>
      <h2 class="mt-8 font-medium text-xl">全て</h2>
      <TaskCard :tasks="tasks.tasks"></TaskCard>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import $axios from '@nuxtjs/axios'
import $auth from '@nuxtjs/auth-next'
import $dayjs from '@nuxtjs/dayjs'
import TaskCard from '@/components/TaskCard.vue'

type TaskSingle = {
  taskID: number
  taskName: string
  taskDescription: string
  deadlineAt: string
  createdAt: string
  isArchived: boolean
  subjectID: number
  subjectName: string
  stateID: number
}

type TasksMultiple = {
  totalCount: number
  tasks: Array<TaskSingle>
}

export default Vue.extend({
  components: {
    TaskCard
  },
  middleware: 'auth',
  data() {
    return {
      TODAY: this.$dayjs(new Date()),
      tasks: {} as TasksMultiple
    }
  },
  computed: {
    tasksExpired() {
      const resultArray: Array<TaskSingle> = []
      for (let index = 0; index < this.tasks.totalCount; index++) {
        if (
          this.$dayjs(
            this.$dayjs(this.tasks.tasks[index].deadlineAt).format(
              'YYYY-MM-DDTHH:mm:ssZ'
            )
          ).isBefore(this.TODAY.format('YYYY-MM-DDTHH:mm:ssZ'))
        ) {
          resultArray.push(this.tasks.tasks[index])
        }
      }
      return resultArray
    },
    tasksToday() {
      const resultArray: Array<TaskSingle> = []
      for (let index = 0; index < this.tasks.totalCount; index++) {
        if (
          !this.$dayjs(
            this.$dayjs(this.tasks.tasks[index].deadlineAt).format(
              'YYYY-MM-DDTHH:mm:ssZ'
            )
          ).isBefore(this.TODAY.format('YYYY-MM-DDTHH:mm:ssZ')) &&
          this.$dayjs(
            this.$dayjs(this.tasks.tasks[index].deadlineAt).format('YYYY-MM-DD')
          ).isSame(this.TODAY.format('YYYY-MM-DD'))
        ) {
          resultArray.push(this.tasks.tasks[index])
        }
      }
      return resultArray
    },
    tasksTomorrow() {
      const resultArray: Array<TaskSingle> = []
      for (let index = 0; index < this.tasks.totalCount; index++) {
        if (
          this.$dayjs(
            this.$dayjs(this.tasks.tasks[index].deadlineAt).format('YYYY-MM-DD')
          ).isSame(this.TODAY.add(1, 'days').format('YYYY-MM-DD'))
        ) {
          resultArray.push(this.tasks.tasks[index])
        }
      }
      return resultArray
    }
  },
  created() {
    this.$axios.$get('/v1/tasks').then((response) => (this.tasks = response))
  }
})
</script>
