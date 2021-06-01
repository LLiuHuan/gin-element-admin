// import type {App} from 'vue'
// import ElementPlus from 'element-plus'
// import 'element-plus/lib/theme-chalk/index.css'
import 'element-plus/packages/theme-chalk/src/base.scss'
// import '@/styles/element-variables.scss'
import {
    ElAside,
    ElMain,
    ElHeader,
    ElFooter,
    ElButton,
    ElMessage,
} from 'element-plus'

// import {componentStore} from '@/store'

export default function setupElement(app: any) {
    app.component(ElButton.name, ElButton)

    app.config.globalProperties.$message = ElMessage
}
