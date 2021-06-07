<template>
  <side-logo/>
  <el-scrollbar class="scrollbar-menu">
    <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        :collapse="!isCollapse"
        :unique-opened="true"
        :router="true"
        mode="vertical"
    >
      <el-submenu index="1">
        <template #title>
          <i class="el-icon-location"></i>
          <span>导航一</span>
        </template>
        <el-menu-item-group>
          <template #title>分组一</template>
          <el-menu-item index="1-1">选项1</el-menu-item>
          <el-menu-item index="1-2">选项2</el-menu-item>
        </el-menu-item-group>
        <el-menu-item-group title="分组2">
          <el-menu-item index="1-3">选项3</el-menu-item>
        </el-menu-item-group>
        <el-submenu index="1-4">
          <template #title>选项4</template>
          <el-menu-item index="1-4-1">选项1</el-menu-item>
        </el-submenu>
      </el-submenu>
      <el-menu-item index="2">
        <i class="el-icon-menu"></i>
        <template #title>导航二</template>
      </el-menu-item>
      <el-menu-item index="3" disabled>
        <i class="el-icon-document"></i>
        <template #title>导航三</template>
      </el-menu-item>
      <el-menu-item index="4">
        <i class="el-icon-setting"></i>
        <template #title>导航四</template>
      </el-menu-item>
    </el-menu>
  </el-scrollbar>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue'
// import { useStore } from '..'

import SideLogo from './SideLogo.vue'
import {useStore} from '../../../store';

export default defineComponent({
  components: {
    SideLogo
  },
  setup () {
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

<style lang="scss">
.scrollbar-menu {
  height: 90%;
  overflow-x: hidden !important;
}

.el-scrollbar__thumb {
  position: relative;
  display: block;
  width: 0;
  height: 0;
  cursor: pointer;
  border-radius: inherit;
  background-color: rgba(144,147,153,.3);
  -webkit-transition: .3s background-color;
  transition: .3s background-color;
}
</style>
