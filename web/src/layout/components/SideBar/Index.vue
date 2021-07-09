<template>
  <side-logo/>
  <el-scrollbar class="scrollbar-menu">
    <el-menu
        :default-active="default_menu"
        class="el-menu-vertical-demo"
        :collapse="!isCollapse"
        :unique-opened="true"
        :router="true"
        mode="vertical"
    >
      <SideItem
          v-for="data in menus"
          :data="data"
      />
    </el-menu>
  </el-scrollbar>
</template>

<script lang="ts">
import { defineComponent, computed, reactive, ref, toRefs } from 'vue'
// import { useStore } from '..'

import SideLogo from './SideBarLogo.vue'
import {useStore} from '../../../store';
import {getMenu} from "../../../apis/login";
import SideItem from './SideBarItem.vue'
import {menuObject} from '../../../model/menuModel'

export default defineComponent({
  components: {
    SideLogo,
    SideItem
  },
  setup () {
    const store = useStore()
    const state = reactive({
      menus: ref(),
      default_menu: ref(1)
    })

    const sidebar = computed(() => {
      return store.state.app.sidebar
    })

    const isCollapse = computed(() => {
      return sidebar.value.opened
    })

    getMenu().then((res: any) => {
      state.menus = res.data.menus
      for (let i = 0; i < res.data.menus.length; i++) {
        if (res.data.menus[i].meta.default_menu) {
          state.default_menu = res.data.menus[i].ID;
        }
      }
    })

    return {
      isCollapse,
      ...toRefs(state),
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
