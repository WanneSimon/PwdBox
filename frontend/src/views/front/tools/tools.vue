<template>
  <div class="tools-wrapper">
    <meta-header :metaInfo="{ title: metaTitle }"></meta-header>
    <div v-for="item,index in tools" :key="'tool_'+index"
      class="tool-item float-block">

      
      <router-link :to="item.link" v-if="item.inner">{{item.name}}</router-link>
      <a :href="item.link" v-else>{{item.name}}</a>

      <div class="intro">
        <el-scrollbar>
          {{item.desc}}
        </el-scrollbar>
      </div>

    </div>
  </div>
</template>

<style scoped lang="scss">
.tools-wrapper{
  margin-top: 10px;
}

.tool-item{
  margin: 0.4rem 0.5rem;
  display: inline-table;
  background-color: rgba($color: #fff, $alpha: 0.4);

  a { text-decoration: none; color: #f54400b5;}
  a:visited { color: rgb(147, 72, 216); }
  a:hover { color:#cc32e0; }

  .intro{
    height: 2.2rem;
    width: 10rem;
    overflow: hidden;
    font-size: 0.8rem;
    text-align: left;
    color: #000;

    transition: height 0.3s;
    -webkit-transition: height 0.3s; /* Safari */
  }

  &:hover{
    .intro{
      height: 5rem;
    }
  }
}
</style>

<script>
import MetaHeader from '/src/components/vue-meta/meta-header.vue'
// import axios from 'axios'
import { ElScrollbar } from 'element-plus'

let baseUrl = import.meta.env.VITE_FRONT_BASE
baseUrl = baseUrl ? baseUrl : ''

export default {
  name: 'Tools',
  inject: [ 'Message' ],
  components: {
    MetaHeader,  ElScrollbar
  },
  data () {
    return {
      metaTitle: '小工具',

      tools: []
    }
  },
  created () {
    this.getToolDatas()
  },
  methods: {
    getToolDatas() {
      //const basePath = import.meta.env.VITE_FRONT_BASE
      this.tools = [
        {
          "name": "flv拉流测试",
          "link": baseUrl + "/tool/live",
          "desc": "b站开源库-flv.js，拉取视频流和直播流",
          "inner": true
        },
        {
          "name": "随机猫猫",
          "link": baseUrl + "/tool/cats",
          "desc": "thecatapi开放的随机猫猫接口",
          "inner": true
        }
      ]
    },

  }
}
</script>