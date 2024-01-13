<template>
  <div>
    <div class="flex justify-center">
      <div class="m-4 w-full">
        <div class="flex mt-8">
          <nuxt-link
            to="/tasks/"
            class="mt-1 font-medium text-2xl text-gray-400"
            >課題一覧
          </nuxt-link>
          <h1 class="ml-4 font-medium text-3xl">教科一覧</h1>
        </div>
        <div v-for="subject in subjects.subjects" :key="subject.subjectID">
          <div v-if="!subject.isArchived">
            <div
              :to="'/task/' + subject.subjectID"
              class="flex items-center mt-3 bg-white border border-gray-100 shadow-lg rounded-xl"
            >
              <button
                class="flex items-center justify-center m-3 mr-1 h-12 pl-4 pr-4 bg-red-600 text-white shadow-lg rounded-lg"
                @click="postArchiveSubject(subject.subjectID)"
              >
                削除
              </button>
              <nuxt-link
                class="flex items-center justify-center m-3 ml-1 mr-2 pl-4 pr-4 h-12 bg-gray-600 text-white shadow-lg rounded-lg"
                :to="'/subject/' + subject.subjectID"
              >
                編集
              </nuxt-link>
              <nuxt-link
                :to="'/subject/' + subject.subjectID"
                class="p-3 pl-2 text-lg text-current"
              >
                {{ subject.subjectName }}
              </nuxt-link>
            </div>
          </div>
        </div>
        <div class="h-24">
          <!-- 下のメニューバーの高さ分、下に余白を開ける -->
        </div>
      </div>
    </div>
    <div class="fixed inset-x-0 flex bottom-0 bg-white h-24">
      <div class="w-5/12 flex justify-center">
        <nuxt-link
          to="/tasks"
          class="flex items-center justify-center m-4 mr-2 h-12 w-full bg-gray-600 text-white shadow-lg rounded-lg"
        >
          <span class="font-medium text-xl">課題一覧へ</span>
        </nuxt-link>
      </div>
      <div class="w-5/12 flex justify-center">
        <nuxt-link
          to="/subject/0"
          class="flex items-center justify-center m-4 ml-2 h-12 w-full bg-blue-600 text-white shadow-lg rounded-lg"
        >
          <span class="font-medium text-xl">教科の追加</span>
        </nuxt-link>
      </div>

      <div class="w-1/6 flex justify-center">
        <nuxt-link to="/settings">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 24 24"
            fill="currentColor"
            class="m-5 ml-2 w-10 h-10"
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

type SubjectSingle = {
  subjectID: number
  subjectName: string
  createdAt: string
  isArchived: boolean
}

type SubjectsMultiple = {
  totalCount: number
  subjects: Array<SubjectSingle>
}

export default Vue.extend({
  middleware: 'auth',
  data() {
    return {
      subjects: {} as SubjectsMultiple
    }
  },
  computed: {},
  created() {
    this.$axios
      .$get('/v1/subjects')
      .then((response) => (this.subjects = response))
  },
  methods: {
    postArchiveSubject(subjectID: number) {
      // subjectIDから配列の位置を検索
      const arrayNumber = this.subjects.subjects.findIndex(
        (element) => element.subjectID === subjectID
      )

      // JSONを更新する（画面表示も消えた状態になる）
      this.$set(this.subjects.subjects[arrayNumber], 'isArchived', true)

      // POST
      this.$axios
        .$post('/v1/subject/' + subjectID, this.subjects.subjects[arrayNumber])
        .catch(() =>
          // 失敗したらJSONを更新して、画面表示を戻す
          this.$set(this.subjects.subjects[arrayNumber], 'isArchived', false)
        )
    }
  }
})
</script>
