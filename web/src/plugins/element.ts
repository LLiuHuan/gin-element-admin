// import type {App} from 'vue'
// import ElementPlus from 'element-plus'
// import 'element-plus/lib/theme-chalk/index.css'
import 'element-plus/packages/theme-chalk/src/base.scss'
// import '@/styles/element-variables.scss'
import {
    ElContainer,
    ElAside,
    ElMain,
    ElHeader,
    ElFooter,
    ElButton,
    ElMessage,
    ElMenu,
    ElMenuItem,
    ElMenuItemGroup,
    ElSubmenu,
    ElScrollbar,
} from 'element-plus'

// import {componentStore} from '@/store'

export default function setupElement(app: any) {
    app.component(ElButton.name, ElButton)
    app.component(ElContainer.name, ElContainer)
    app.component(ElAside.name, ElAside)
    app.component(ElMain.name, ElMain)
    app.component(ElHeader.name, ElHeader)
    app.component(ElFooter.name, ElFooter)
    app.component(ElMenu.name, ElMenu)
    app.component(ElMenuItem.name, ElMenuItem)
    app.component(ElMenuItemGroup.name, ElMenuItemGroup)
    app.component(ElSubmenu.name, ElSubmenu)
    app.component(ElScrollbar.name, ElScrollbar)
    app.component(ElMessage)
    app.config.globalProperties.$message = ElMessage
}
