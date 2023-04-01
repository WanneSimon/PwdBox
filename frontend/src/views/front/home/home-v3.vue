<template>
    <!-- <HeadNav></HeadNav> -->
    
    <div class="main-body">
      <meta-header :metaInfo="{ title: '' }"></meta-header>
      <div class="content" v-if="sentence">
        <p class="st" >{{sentence.content}}</p>
        <p class="source">
          <span v-if="sentence.name" style="text-align: right;">
             -- {{ sentence.name }}
             <span v-if="sentence.origin">（{{ sentence.origin }}）</span>
            </span>
        </p>
        <!-- <el-button @click="toTools">工具</el-button> -->
      </div>

    </div>
    
</template>

<style lang="scss" scoped>
@import './home-v3.scss';
</style>

<script src="./home-v3.js">
</script>