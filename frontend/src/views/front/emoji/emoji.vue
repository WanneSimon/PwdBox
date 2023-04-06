<template>
  <div class="emoji-wrapper">
    <!-- <div style="position: sticky;margin-bottom: 0.3rem;" v-show="showConfig"> -->
    <div style="" v-if="!showConfig">
      <el-input v-model="name" class="search-input" 
          placeholder="文件名" @keydown.enter.native="searchClick" />
      <el-button @click="searchClick">搜索</el-button>
      <el-button @click="showConfig=true">设置</el-button>
      
      <div class="content ">
        
        <el-scrollbar class="images " height=" calc( 100% - 6rem ) "  
            v-loading="loading" element-loading-background="rgba(0,0,0,0.2)">
          <div v-show="files && files.length>0" style="width:100%">
          <el-card v-for="item,index in files" :key="'file_'+index"
            :body-style="{ padding: '0px' }" class="img-item"> 
                <el-image :src="item._url" :fit="'contain'" :preview-src-list="[item._url]" >
                  <template #error>{{ item._url }}</template>
                </el-image>
                
                <div class="opt">
                  <el-button class="copy" type="primary" plain @click="copyToClipboard(item)">复制</el-button>
                </div>
          </el-card>
          </div>
        </el-scrollbar>

      </div>
    </div>

    <div v-if="showConfig">
      <!-- <el-button @click="showConfig=false">返回</el-button> -->
      <EmojiConfig @saved="refreshConfig" @close="showConfig=false"></EmojiConfig>
    </div>

  </div>
</template>

<!-- <script src="./emoji.ts" lang="ts"></script> -->
<script src="./emoji.js" ></script>

<style scoped lang="scss">
.emoji-wrapper{
  padding-top: 0.4rem;
}
.search-input{
  width: 20rem;
}

.content{
  min-height: 50vh;
  display: flex;
  justify-content: center;
  margin-top: 0.4rem;
}

.center-container{
  display: flex;
  // align-items: center;
  justify-content: center;
}

.images{
  position: fixed;
  width: 100%;

  :deep(.el-card__body) {
    height: 100%;
    width: 100%;
  }
  
  .img-item{
    width: 8rem;
    height: 8rem;
    justify-content: center;
    margin: 0rem 0.1rem;
    display: inline-block;
    position: relative;

    :deep(.el-image){
      height: 100%;
      width: 100%;
    }

    &:hover{
      .opt{
        display: block;
      }
    }
  }
}

.opt{
  width: 100%;
  height: 2rem;
  position: absolute;
  display: flex;
  text-align: center;
  justify-content: center;
  // margin: 2px 0px;
  bottom: 0;
  display: none;

  .copy{
    width: 100%;
  }
}
</style>