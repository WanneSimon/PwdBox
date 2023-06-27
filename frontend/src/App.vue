<script setup lang="ts">
import { provide } from 'vue'
import globalStore from './components/pinia/store.js'
import ClientOnly from '@duannx/vue-client-only'
// import HeadNav from '/src/views/components/head-nav/head-nav.vue'
// import ElementBg from '/src/views/components/bg/element-bg.vue'
import HeadNav from '@/views/components/head-nav/head-nav.vue'
import ElementBg from '@/views/components/bg/element-bg.vue'
import InnerApp from './InnerApp.vue';

const store = globalStore()
provide("store", store)

let baseUrl = import.meta.env.VITE_FRONT_BASE
baseUrl = baseUrl ? baseUrl : ''
const bgImgs =  [
  baseUrl+'/static/images/LostEmber.jpg',
  // baseUrl+'/static/images/picture-lost.jpg',
]
</script>

<template>
  <HeadNav class="head-nav"></HeadNav>
  <client-only>
      <!-- <ElementBg :urls="bgImgs" :interval="10000" :random="true"></ElementBg> -->
      <el-scrollbar class="index" style="height: calc( 100% - 2.4rem )">
       <inner-app></inner-app>
      </el-scrollbar>
  </client-only>
</template>

<style lang="scss">
//@import './assets/css/global.scss';
@use './assets/css/global.scss' as *; 
</style>
<style scoped lang="scss">
.head-nav{
  position: sticky;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
}
.index {
  :deep(.el-scrollbar__view) {
    height: 100%;
  }
}
</style>
