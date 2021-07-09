<template>
  <a
      v-if="isExternal(to)"
      :href="to"
      target="_self"
      rel="noopener"
  >
    <slot />
  </a>
  <div
      v-else
      @click="push"
  >
    <slot />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { isExternal } from '../../../utils/validate'
import { useRouter } from 'vue-router'

export default defineComponent({
  name: '',
  props: {
    to: {
      type: String,
      required: true
    }
  },
  setup () {
    const router = useRouter()
    const push = () => {
      router.push(props.to).catch((err) => {
        console.log(err)
      })
    }
    return {
      isExternal,
      push
    }
  }
})
</script>
