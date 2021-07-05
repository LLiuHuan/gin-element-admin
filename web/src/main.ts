import { createApp } from 'vue'
import App from './App.vue'
import {loadAllPlugins} from './plugins'
import './styles/index.scss'
import './permission'

import { store } from './store'

import router from './router'

const app = createApp(App)

loadAllPlugins(app)

app.use(router)
app.use(store)

app.mount('#app')
