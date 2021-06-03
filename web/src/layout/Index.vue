<template>
  <el-container class="layout">
    <el-container>
      <el-aside width="200px">
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
import {defineComponent} from 'vue'

import SideBar from './components/SideBar/Index.vue'
import NavBar from './components/NavBar/Index.vue'

export default defineComponent({
  components: {
    SideBar,
    NavBar,
  },
  setup() {

    return {}
  }
})
</script>

<style lang="less">
.main {
  //height: 100%;
}
</style>