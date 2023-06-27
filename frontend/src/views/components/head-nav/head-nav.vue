<template>
  <div class="nav-head float-block fixed-nav-head" style="--wails-draggable: drag">
      <label class="icon-label">
        <router-link :to="'/home'" title="home home ~ sweet home!" target="self">
          <el-button size="small" color="#626aef">Home</el-button>
        </router-link>
      </label>
      <!--
      <label class="icon-label">
        <router-link :to="'/emoji'" title="Emoji">
          <el-button size="small" color="#626aef">Emoji</el-button>
        </router-link>
      </label>
      -->
      <label class="icon-label" title="password box - 保险箱">
        <router-link :to="'/pwdbox'">
          <el-button size="small" color="#626aef">pwdbox</el-button>
        </router-link>
      </label>
      <label class="icon-label" title="tools - 小工具">
        <router-link :to="'/tools'">
          <el-button size="small" color="#626aef">tools</el-button>
        </router-link>
      </label>


      <label class="right-pos">
        <!-- <el-button size="small" type="warning" @click="min" :icon="'Minus'">最小化</el-button>
        <el-button size="small" type="warning" @click="exit" :icon="'Close'">退出</el-button>
        <el-button size="small" @click="goBack" :icon="'ArrowLeftBold'">返回</el-button> -->
        <el-button size="small" @click="goBack" title="back - 返回">
            <el-icon><ArrowLeftBold /></el-icon></el-button>
        <el-button size="small" type="warning" @click="min" title="min - 最小化">
          <el-icon><Minus /></el-icon></el-button>
        <el-button size="small" type="danger" @click="exit" title="exit - 关闭">
          <el-icon><Close /></el-icon></el-button>
      </label>
  </div>
</template>

<script setup lang="ts" >
// import { Exit, Minimises } from "/wailsjs/index"
import { Exit, Minimises } from "@/../wailsjs/index"
import { Minus, Close, ArrowLeftBold } from '@element-plus/icons-vue'
import { getCurrentInstance } from "vue"
import { useRouter } from "vue-router"

const router = useRouter()

const goBack = () => {
  // this.$router.back()
  router.back()
}
const exit = () =>  {
  Exit()
}
const min= () => {
  Minimises()
}
</script>

<style scoped lang="scss">
.nav-head{
  border-radius: unset;
  padding: 0rem 0.4rem;
  height: 2.4rem;
  line-height: 2.4rem;
  position: relative;

  padding-left: 10rem;
  padding-right: 10rem;

  text-align: center;
  //background-color: rgba($color: #d6c2a8, $alpha: 0.75);
  background-color: #f3f3f3bf;
  a {
    color: #606266 !important;
  }
  a:hover {
    color: #409eff !important;
  }

  .right-pos{
    position: absolute;
    right: 0.4rem;
    top: 0;
    width: rem;
  }
}

.fixed-nav-head{
  position: fixed;
  z-index: 2003;  
}
:deep(.el-button+.el-button) {
  margin-left: 1px;
}
</style>
