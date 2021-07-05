<template>
  <h1>dashboard</h1>
  <button @click="test">aaaaaa</button>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import {useRouter} from 'vue-router'

export default defineComponent({
  name: '',
  setup () {
    const router = useRouter()
    const test = () => {
      router.push({path: '/login'}).catch(err => {
        console.log(err)
      })
    }
    return {
      test
    }
  }
})
</script>

