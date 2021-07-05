<!--
 * @Description: 重定向
 * @Author: LLiuHuan
 * @Date: 2021-02-18 14:16:11
 * @LastEditTime: 2021-02-18 16:58:01
 * @LastEditors: LLiuHuan
-->

<script lang="ts">
import { defineComponent } from 'vue'
import { useRoute, useRouter } from 'vue-router'

export default defineComponent({

  created() {
    const { params, query } = useRoute()
    const { path } = params
    useRouter().replace({ path: '/' + path, query }).catch((err) => {
      console.warn(err)
    })
  },

  // eslint-disable-next-line vue/require-render-return
  render() {
    // Avoid warning for missing template
  }

})
</script>
