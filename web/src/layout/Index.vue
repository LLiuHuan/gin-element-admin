<template>
  <el-container class="layout">
    <el-container>
      <el-aside :width="!isCollapse ? '64px' : '300px'" style="padding: 0;transition: width .38s;margin-bottom: 0;">
        <side-bar/>
      </el-aside>
      <el-container>
        <el-header>
          <nav-bar/>
        </el-header>
        <el-main class="main">
          <router-view></router-view>
        </el-main>
        <el-footer>Footer</el-footer>
      </el-container>
    </el-container>
  </el-container>
</template>

<script lang="ts">
import {computed, defineComponent} from 'vue'
import {useStore} from '../store';

import SideBar from './components/SideBar/Index.vue'
import NavBar from './components/NavBar/Index.vue'

export default defineComponent({
  components: {
    SideBar,
    NavBar,
  },
  setup() {
    const store = useStore()

    const sidebar = computed(() => {
      return store.state.app.sidebar
    })

    const isCollapse = computed(() => {
      return sidebar.value.opened
    })
    return {
      isCollapse
    }
  }
})
</script>

<style lang="less">
.main {
  //height: 100%;
}

.layout {
  height: 100%;
}

.narrow {
  width: 80px !important;
}


</style>