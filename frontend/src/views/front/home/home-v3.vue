<template>
    <!-- <HeadNav></HeadNav> -->
    
    <div class="main-body">
      <meta-header :metaInfo="{ title: '' }"></meta-header>
      <div class="content">
        <div v-if="sentence">{{sentence.content}}</div>
        <!-- <el-button @click="toTools">工具</el-button> -->
      </div>

    </div>
    
</template>

<style lang="scss" scoped>
@import './home-v3.scss';
</style>

<script src="./home-v3.js">
</script>