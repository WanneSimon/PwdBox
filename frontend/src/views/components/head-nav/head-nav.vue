<template>
  <div class="nav-head float-block" style="--wails-draggable: drag">
      <label class="icon-label">
        <router-link :to="'/home'" title="Home" target="self">
          <el-button size="small" color="#626aef">Home</el-button>
        </router-link>
      </label>
      <label class="icon-label">
        <router-link :to="'/emoji'" title="Emoji">
          <el-button size="small" color="#626aef">Emoji</el-button>
        </router-link>
      </label>
      <label class="icon-label" title="小工具">
        <router-link :to="'/tools'">
          <el-button size="small" color="#626aef"> 小工具</el-button>
        </router-link>
      </label>
    
      <label class="right-pos">
        <el-button size="small" type="warning" @click="exit">退出</el-button>
        <el-button size="small" @click="goBack">返回</el-button>
      </label>
  </div>
</template>

<script>
import { Exit } from "/wailsjs/index.js"

export default {
    name: 'HeadNav',
    components: { },
    data () {
        return {
            
        }
    },

    methods: {
      goBack() {
        this.$router.back()
      },
      exit() {
        Exit()
      }
    }
}
</script>

<style scoped lang="scss">
.nav-head{
  border-radius: unset;
  padding: 0rem 0.4rem;
  height: 2.4rem;
  line-height: 2.4rem;
  position: relative;

  padding-right: 10rem;

  text-align: left;
  background-color: rgba($color: #000000, $alpha: 0.3);
  a {
    color: #606266 !important;
  }
  a:hover {
    color: #409eff !important;
  }

  .right-pos{
    position: absolute;
    right: 0.2rem;
    top: 0;
  }
}
</style>
