<script setup>
import { provide } from 'vue'
import globalStore from './components/pinia/store.js'
import ClientOnly from '@duannx/vue-client-only'
import HeadNav from '/src/views/components/head-nav/head-nav.vue'
import ElementBg from '/src/views/components/bg/element-bg.vue'

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
  <client-only>
      <ElementBg :urls="bgImgs" :interval="10000" :random="true"></ElementBg>
      <HeadNav class="head-nav"></HeadNav>
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
</style>

<script>
import InnerApp from './InnerApp.vue';

export default {
  components: { InnerApp },
  data () {
    return {
    }
  },
  mounted () {
    // document.body.removeChild(document.getElementById('appLoading'));
  },
}
</script>