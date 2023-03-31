<template>
    <div>
      <div class="btns">
        <el-button type="primary" secondary @click="randomPictures" >换一批</el-button>
      </div>

      <div v-if="historyList.data && historyList.data.length>0">
        <span class="pre-txt">历史记录：</span>
        <el-pagination v-model:current-page="historyPage" :page-count="historyList.data.length" 
          :pager-count="7"  @current-change="showHistoryPage">
        </el-pagination>
      </div>
      
      <el-image
        class="image-pic"
        v-for="item,index in list.data" :key="'pic_'+index"
          object-fit="contain"
          :src="item.url"
      >
        <template #placeholder>
          <div style="
            width: 100px;
            height: 100px;
            display: flex;
            align-items: center;
            justify-content: center;
            background-color: #0001;
          "
        >
          Loading...
        </div>
        </template>
      </el-image>

    </div>
</template>

<style scoped lang="scss">
.btns {
  margin-top: 10px;
}

.pre-txt{
  margin-left: 1rem;
}

.image-pic{
  height: 260px;
  margin: 8px;
  align-content: center;
  justify-content: center;
}
</style>

<script src="./the-cat-api.js">
</script>
