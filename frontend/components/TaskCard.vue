<template>
  <div>
    <div v-for="task in tasks" :key="task.taskID">
      <div v-if="!task.isArchived">
        <div
          :to="'/task/' + task.taskID"
          class="flex items-center mt-3 bg-white border border-gray-100 shadow-lg rounded-xl"
        >
          <button
            class="flex items-center justify-center m-3 mr-2 w-12 h-12 bg-gray-600 text-white shadow-lg rounded-lg"
            @click="$emit('completeTask', task.taskID)"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              stroke-width="1.5"
              stroke="currentColor"
              class="w-6 h-6"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M4.5 12.75l6 6 9-13.5"
              />
            </svg>
          </button>
          <nuxt-link
            :to="'/task/' + task.taskID"
            class="p-3 pl-2 text-current leading-6"
          >
            <p class="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6 mr-1 text-gray-800"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 6.042A8.967 8.967 0 006 3.75c-1.052 0-2.062.18-3 .512v14.25A8.987 8.987 0 016 18c2.305 0 4.408.867 6 2.292m0-14.25a8.966 8.966 0 016-2.292c1.052 0 2.062.18 3 .512v14.25A8.987 8.987 0 0018 18a8.967 8.967 0 00-6 2.292m0-14.25v14.25"
                />
              </svg>
              {{ task.taskName }}
              (ID: {{ task.taskID }})
            </p>
            <p class="flex items-center">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6 mr-1 text-gray-800"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M4.26 10.147a60.436 60.436 0 00-.491 6.347A48.627 48.627 0 0112 20.904a48.627 48.627 0 018.232-4.41 60.46 60.46 0 00-.491-6.347m-15.482 0a50.57 50.57 0 00-2.658-.813A59.905 59.905 0 0112 3.493a59.902 59.902 0 0110.399 5.84c-.896.248-1.783.52-2.658.814m-15.482 0A50.697 50.697 0 0112 13.489a50.702 50.702 0 017.74-3.342M6.75 15a.75.75 0 100-1.5.75.75 0 000 1.5zm0 0v-3.675A55.378 55.378 0 0112 8.443m-7.007 11.55A5.981 5.981 0 006.75 15.75v-1.5"
                />
              </svg>

              <span v-if="task.subjectName">
                {{ task.subjectName }}
              </span>
              <span v-else>----</span>

              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="w-6 h-6 ml-2 mr-1 text-gray-800"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z"
                />
              </svg>

              <span
                v-if="
                  !$dayjs($dayjs(task.deadlineAt).format('YYYY-MM-DD')).isSame(
                    $dayjs('1-01-01')
                  )
                "
              >
                {{ $dayjs(task.deadlineAt).format('YYYY/MM/DD HH:mm') }}
              </span>
              <span v-else>----/--/--</span>
            </p>
          </nuxt-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
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
  props: {
    tasks: {
      type: Array,
      default: [] as Array<TaskSingle>
    } as Vue.PropOptions<Array<TaskSingle>>
  }
})
</script>
