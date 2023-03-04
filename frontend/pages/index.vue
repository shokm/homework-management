<template>
  <div></div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  middleware: 'auth',
  created() {
    if (this.$auth.loggedIn) {
      this.$router.push('/tasks')
    }
  }
})
</script>
