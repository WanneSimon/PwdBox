import { createApp } from 'vue'
import App from './App.vue'
import { createPinia } from 'pinia'
import router from './router/index'

// import globalStore from './components/pinia/store.js'
import metaManager from './components/vue-meta/meta-manager.js'

import 'highlight.js/styles/atom-one-dark.css'
import hljs from 'highlight.js'

// 引入 Vue-quill-editor
// import 'quill/dist/quill.core.css'
// import 'quill/dist/quill.snow.css'
// import 'quill/dist/quill.bubble.css'

const app = createApp(App)
app.use(createPinia())
app.use(router)
app.use(metaManager)

// app.use(VueQuillEditor)

app.directive('highlight',function (el) {
  /*
  // 1. 多行代码块
  let blocks = el.querySelectorAll('pre code')
  blocks.forEach((block)=>{
    // 这个方法会把代码块包裹成单独的一行
    hljs.highlightBlock(block)
  })
  // 2. 行内代码块
  let innerBlocks = el.querySelectorAll('p code, li code')
  innerBlocks.forEach((block)=>{
      block.className += " inner-line-code"
  })
  */

  let blocks = el.querySelectorAll('code')
  blocks.forEach((codeEl)=>{
    let parentName = codeEl.parentNode?.tagName
    if(parentName && parentName.toLowerCase()=='pre') {
      // 这个方法会把代码块包裹成单独的一行
      hljs.highlightBlock(codeEl)
    } else {
      codeEl.classList.add('inner-line-code')
    }

  })
})


// const store = globalStore()

// the last
app.mount('#app')
