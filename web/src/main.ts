import { createApp } from 'vue'
import App from './App.vue'
import {loadAllPlugins} from './plugins'

const app = createApp(App)
loadAllPlugins(app)

app.mount('#app')
