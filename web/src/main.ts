import { createApp } from 'vue'
import App from './App.vue'
import {loadAllPlugins} from './plugins'
import './styles/index.scss'

import { store } from './store'

import router from './router/index'

const app = createApp(App)

loadAllPlugins(app)

app.use(router)
app.use(store)
app.mount('#app')
