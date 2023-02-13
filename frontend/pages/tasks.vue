<template>
  <div>
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
          <TaskCard
            :tasks="tasksExpired"
            @completeTask="postCompleteTask"
          ></TaskCard>
        </div>
        <div v-if="tasksToday.length">
          <h2 class="mt-8 font-medium text-xl">今日まで</h2>
          <TaskCard
            :tasks="tasksToday"
            @completeTask="postCompleteTask"
          ></TaskCard>
        </div>
        <div v-if="tasksTomorrow.length">
          <h2 class="mt-8 font-medium text-xl">明日まで</h2>
          <TaskCard
            :tasks="tasksTomorrow"
            @completeTask="postCompleteTask"
          ></TaskCard>
        </div>
        <h2 class="mt-8 font-medium text-xl">全て</h2>
        <TaskCard
          :tasks="tasks.tasks"
          @completeTask="postCompleteTask"
        ></TaskCard>
      </div>
    </div>
    <div class="sticky flex bottom-0 bg-white h-24">
      <div class="w-5/6 flex justify-center">
        <nuxt-link
          to="/task/0"
          class="flex items-center justify-center m-4 h-12 w-full bg-blue-600 text-white shadow-lg rounded-lg"
        >
          課題の追加
        </nuxt-link>
      </div>
      <div class="w-1/6 flex justify-center">
        <nuxt-link to="/settings">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            class="m-5 w-10 h-10"
          >
            <path
              fill-rule="evenodd"
              d="M11.078 2.25c-.917 0-1.699.663-1.85 1.567L9.05 4.889c-.02.12-.115.26-.297.348a7.493 7.493 0 00-.986.57c-.166.115-.334.126-.45.083L6.3 5.508a1.875 1.875 0 00-2.282.819l-.922 1.597a1.875 1.875 0 00.432 2.385l.84.692c.095.078.17.229.154.43a7.598 7.598 0 000 1.139c.015.2-.059.352-.153.43l-.841.692a1.875 1.875 0 00-.432 2.385l.922 1.597a1.875 1.875 0 002.282.818l1.019-.382c.115-.043.283-.031.45.082.312.214.641.405.985.57.182.088.277.228.297.35l.178 1.071c.151.904.933 1.567 1.85 1.567h1.844c.916 0 1.699-.663 1.85-1.567l.178-1.072c.02-.12.114-.26.297-.349.344-.165.673-.356.985-.57.167-.114.335-.125.45-.082l1.02.382a1.875 1.875 0 002.28-.819l.923-1.597a1.875 1.875 0 00-.432-2.385l-.84-.692c-.095-.078-.17-.229-.154-.43a7.614 7.614 0 000-1.139c-.016-.2.059-.352.153-.43l.84-.692c.708-.582.891-1.59.433-2.385l-.922-1.597a1.875 1.875 0 00-2.282-.818l-1.02.382c-.114.043-.282.031-.449-.083a7.49 7.49 0 00-.985-.57c-.183-.087-.277-.227-.297-.348l-.179-1.072a1.875 1.875 0 00-1.85-1.567h-1.843zM12 15.75a3.75 3.75 0 100-7.5 3.75 3.75 0 000 7.5z"
              clip-rule="evenodd"
            />
          </svg>
        </nuxt-link>
      </div>
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
  },
  methods: {
    postCompleteTask(taskID: number) {
      // taskIDから配列の位置を検索
      const arrayNumber = this.tasks.tasks.findIndex(
        (element) => element.taskID === taskID
      )

      // JSONを更新する（画面表示も消えた状態になる）
      this.$set(this.tasks.tasks[arrayNumber], 'isArchived', true)

      // POST
      this.$axios
        .$post('/v1/task/' + taskID, this.tasks.tasks[arrayNumber])
        .catch(() =>
          // 失敗したらJSONを更新して、画面表示を戻す
          this.$set(this.tasks.tasks[arrayNumber], 'isArchived', false)
        )
    }
  }
})
</script>
