import { createApp, provide } from 'vue'
import App from './App.vue'
import { store, key } from './store'
import router from './router'
import '@/styles/index.scss'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import '@okiss/lib/style.css'
import { createPinia } from 'pinia'
import app from '@/util/app'

app.GetBaseConf().then(base => {
  window.BaseApi = base.externalController
  window.HomeDir = base.homeDir
  provide('vsPath', import.meta.env.PROD ? location.pathname + 'assets/monaco-editor/vs' : 'node_modules/monaco-editor/min/vs')
  createApp(App)
    .use(ElementPlus)
    .use(createPinia())
    .use(store, key)
    .use(router)
    .mount('#app')
})
