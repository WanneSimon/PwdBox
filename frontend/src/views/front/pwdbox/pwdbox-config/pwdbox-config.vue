<template>
  <div >
    <div class="config-wrapper">
      <div v-for="item,index in emojis" :key="'emoji-config-'+index"
          class="path-wrapper">
          <div>
            <el-input v-model="emojis[index]" placeholder="请选择文件夹" readonly></el-input>
            <label class="btns">
                  <el-button type="primary" @click="selectFolder(index)">选择</el-button>
                  <el-button type="warning" @click="deleteConfig(index)">删除</el-button>
            </label>
          </div>

      </div>
    </div>

    <div class="btn-bottom">
        <el-button @click="back">返回</el-button>
        <el-button @click="pushNewConfig" type="primary"> + </el-button>
        <el-button @click="resetConfig" type="warning"> 重置  </el-button>
        <el-button @click="saveConfig" type="success"> 保存  </el-button>
    </div>

  </div>
</template>

<script src="./pwdbox-config.js"></script>

<style scoped lang="scss">

.config-wrapper{
  display: flex;
  flex-direction: column;
  margin-top: 1rem;
  align-items: center;
}


.path-wrapper{
  width: 30rem;
  // display: block;
  margin-top: 0.5rem;
  // justify-content: center;
  // align-content: center;

  >div{
    display: flex;
    text-align: center;
    justify-content: center;
  }
}



.btns{
  width: 14rem;
}
.btns>*{
  margin: 0 0.4rem 0 0;
}

.btn-bottom{
  margin-top: 1rem;
}
</style>