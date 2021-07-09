<template>
  <div v-if="!data || !data.hidden">
    <template v-if="isExternal(data.path)">
      <a
          :href="data.path"
          target="_blank"
          rel="no-opener"
      >
        <el-menu-item :index="data.ID.toString()">
          <i :class="'el-icon-'+data.meta.icon"></i>
          <template #title>{{ data.name }}</template>
        </el-menu-item>
      </a>
    </template>
    <el-menu-item :route="data.path" :index="data.ID.toString()">
      <i :class="'el-icon-'+data.meta.icon"></i>
      <template #title>{{ data.name }}1</template>
    </el-menu-item>
  </div>
</template>

<script lang="ts">
import {defineComponent, PropType} from 'vue'
import {menuObject} from '../../../model/menuModel'
import SideBarItemLink from './SideBarItemLink.vue'
import {isExternal} from '../../../utils/validate'

export default defineComponent({
  name: '',
  props: {
    data: {
      type: Object as PropType<menuObject>,
      required: true,
    }
  },
  components: {
    SideBarItemLink
  },
  setup(prop) {
    // console.log(prop.data)
    return {
      isExternal
    }
  }
})
</script>
