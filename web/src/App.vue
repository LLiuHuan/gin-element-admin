<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  name: 'App',
  components: {
  }
})
</script>

<style lang="scss">
.el-menu-item:hover{
  outline: 0 !important;
  color: #409EFF !important;
}

.el-submenu__title:focus, .el-submenu__title:hover{
  outline: 0 !important;
  color: #fff !important;
  background: #435EBE !important;
  border-radius: 8px !important;
}
.el-menu-item:hover{
  outline: 0 !important;
  color: #fff !important;
  background: #435EBE !important;
  border-radius: 8px !important;
}
</style>