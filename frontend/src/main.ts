import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import router from './router/index'

// import globalStore from './components/pinia/store.js'

import metaManager from './components/vue-meta/meta-manager.js'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(metaManager)

// const store = globalStore()
app.mount('#app')
