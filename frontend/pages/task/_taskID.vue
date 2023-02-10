<template>
  <div class="flex justify-center">
    <div
      class="bg-white shadow-xl rounded-xl md:w-auto w-11/12 md:p-20 p-5 md:m-20 m-5"
    >
      <p>{{ $auth.user }}</p>
      <br />
      <p
        v-if="task.isArchived == true"
        class="md:w-96 w-full max-h-14 rounded-xl bg-red-200 text-center text-red-600 p-4"
      >
        完了済みの課題です
      </p>
      <input
        v-model="task.taskName"
        type="text"
        class="border border-DEFAULT md:w-96 w-full max-h-14 rounded-xl bg-transparent text-current p-5"
      />
      <br />
      <input
        v-model="task.taskDescription"
        type="text"
        class="border border-DEFAULT md:w-96 w-full max-h-14 rounded-xl bg-transparent text-current p-5"
      />
      <br />
      <input
        v-model="$dayjs(task.deadlineAt).format('YYYY-MM-DDThh:mm')"
        type="datetime-local"
        class="border border-DEFAULT md:w-96 w-full max-h-14 rounded-xl bg-transparent text-current p-5"
      />
      <br />
      <button
        class="md:w-96 w-full max-h-14 rounded-xl bg-blue-600 text-center text-white p-4"
        @click="updateTask"
      >
        更新
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import $axios from '@nuxtjs/axios'
import $auth from '@nuxtjs/auth-next'
import $dayjs from '@nuxtjs/dayjs'

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

export default Vue.extend({
  middleware: 'auth',
  data() {
    return {
      task: {} as TaskSingle
    }
  },
  created() {
    this.$axios
      .$get('/v1/task/' + this.$route.params.taskID)
      .then((response) => (this.task = response))
  },
  methods: {
    updateTask() {
      this.$axios.$post('/v1/task/' + this.$route.params.taskID, this.task)
    }
  }
})
</script>
