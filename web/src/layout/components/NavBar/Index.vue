<template>
  <div class="navbar" style="height: 50px">
    <hamburger
        class="hamburger"
        :is-active="sidebar.opened"
        @toggle-click="toggleSideBar"
    />
  </div>
</template>

<script lang="ts">
import {defineComponent, computed, toRefs, reactive} from 'vue'
import Hamburger from '../../../components/Hamburger/Index.vue'
import { useStore } from '../../../store'
import {AppActionTypes} from "../../../store/app/action-types";

export default defineComponent({
  components: {
    Hamburger
  },
  setup() {
    const store = useStore()
    const sidebar = computed(() => {
      return store.state.app.sidebar
    })

    const state = reactive({
      toggleSideBar: () => {
        store.dispatch(AppActionTypes.ACTION_TOGGLE_SIDEBAR)
      }
    })
    return {
      sidebar,
      ...toRefs(state)
    }
  }
})
</script>

<style lang="scss">
.navbar {
  height: 50px;
  overflow: hidden;
  position: relative;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

  .hamburger {
    line-height: 46px;
    height: 100%;
    float: left;
    padding: 0 15px;
    cursor: pointer;
    transition: background 0.3s;
    -webkit-tap-highlight-color: transparent;

    &:hover {
      background: rgba(0, 0, 0, 0.025);
    }

    .icon {
      width: 24px;
      height: 24px;
    }
  }
}

</style>