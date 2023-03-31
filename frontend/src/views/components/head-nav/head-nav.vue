<template>
  <div class="nav-head float-block" style="--wails-draggable: drag">
      <label class="icon-label">
        <router-link :to="'/home'" title="首页" target="self">
          首页
        </router-link>
      </label>
      <label class="icon-label">
        <router-link :to="'/emoji'" title="Emoji">
          Emoji
        </router-link>
      </label>
      <label class="icon-label" title="小工具">
        <router-link :to="'/tools'">
          小工具
        </router-link>
      </label>
    
      <label class="right-pos">
        <el-button size="small" @click="goBack">返回</el-button>
      </label>
  </div>
</template>

<script>
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
    color: #00ffb8 !important;
  }

  .right-pos{
    position: absolute;
    right: 0.2rem;
    top: 0;
  }
}
</style>
