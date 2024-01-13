<template>
  <div>
    <div class="flex justify-center">
      <div class="m-4 w-full">
        <div class="flex mt-8">
          <h1 class="font-medium text-3xl">
            <span v-if="param !== 0">教科編集</span><span v-else>教科追加</span>
          </h1>
          <!-- TODO: すでにID:0のリンクにいる場合、画面遷移が効かないので暫定的に"'/subject/new' + (param + 1)"でアクセスさせる。存在しないIDのため、エラーが起こり新規作成扱いになる。 -->
          <nuxt-link
            v-if="param !== 0"
            :to="'/subject/new' + (param + 1)"
            class="ml-4 mt-1 font-medium text-2xl text-gray-400"
            >新規追加</nuxt-link
          >
        </div>
        <p v-if="subject.subjectID" class="mt-3">
          教科ID: {{ subject.subjectID }}
        </p>
        <p
          v-if="subject.isArchived == true"
          class="flex justify-center items-center mt-8 h-14 rounded-xl bg-red-200 font-medium text-red-600"
        >
          削除済みの教科です
        </p>

        <h2 class="mt-8 font-medium text-xl">教科名</h2>
        <div
          class="flex items-center mt-3 bg-white border border-gray-100 shadow-lg rounded-xl"
        >
          <div
            class="m-3 mr-0 w-2 h-12 bg-gray-600 text-white shadow-lg rounded-lg"
          ></div>
          <div class="p-3 pl-2 text-current leading-6">
            <input
              v-model="subject.subjectName"
              type="text"
              class="h-12 rounded-lg text-lg text-current p-5 pl-2"
            />
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
          to="/subjects"
          class="flex items-center justify-center m-4 mr-2 h-12 w-full bg-gray-600 text-white shadow-lg rounded-lg"
        >
          <span class="font-medium text-xl"> 戻る </span>
        </nuxt-link>
      </div>
      <div class="w-5/12 flex justify-center">
        <button
          class="flex items-center justify-center m-4 ml-2 mr-2 h-12 w-full bg-blue-600 text-white shadow-lg rounded-lg"
          @click="updateTask"
        >
          <span class="font-medium text-xl">
            <span v-if="param !== 0">保存する</span><span v-else>追加する</span>
          </span>
        </button>
      </div>

      <div class="w-2/12 flex justify-center">
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

type SubjectSingle = {
  subjectID: number
  subjectName: string
  createdAt: string
  isArchived: boolean
}

export default Vue.extend({
  middleware: 'auth',
  data() {
    return {
      subject: {} as SubjectSingle,
      param: Number(this.$route.params.subjectID)
    }
  },
  created() {
    if (Number(this.param) !== 0) {
      this.$axios
        .$get('/v1/subject/' + this.param)
        .then((response) => (this.subject = response))
        .catch(() => (this.param = 0))
    }
  },
  methods: {
    updateTask() {
      this.$axios
        .$post('/v1/subject/' + this.param, this.subject)
        .then((response) => (this.subject = response))
        .then(() => history.pushState('', '', './' + this.subject.subjectID))
        .then(() => (this.param = this.subject.subjectID))
    }
  }
})
</script>
