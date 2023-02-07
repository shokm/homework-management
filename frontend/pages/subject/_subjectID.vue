<template>
  <div>{{ subject }}</div>
</template>

<script lang="ts">
import Vue from 'vue'
import $axios from '@nuxtjs/axios'
import $auth from '@nuxtjs/auth-next'

export default Vue.extend({
  middleware: 'auth',
  data() {
    return {
      subject: {}
    }
  },
  created() {
    this.$axios
      .$get('/v1/subject/' + this.$route.params.subjectID)
      .then((response) => (this.subject = response))
  }
})
</script>
