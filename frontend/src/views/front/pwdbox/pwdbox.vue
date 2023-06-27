<template>
  <div class="">
        <el-row class="" style="calc( 100% - 6rem );overflow:auto;padding:0rem 1rem;"
            @wheel.native="lazyLoadMore()">
          <el-col :xs="12" :sm="8" :md="6" :lg="4" :xl="3" 
                v-for="item,index in list" :key="'file_'+index">
            <el-card 
              :body-style="{ padding: '0px' }" class="card-item infinite-list-item"
                  @click="showPlatformDetail(item)"> 
                  <!-- <span>{{item.id}}</span> -->
                  
                  <div class="name"> 
                    <span class="name-text ellipse-text-line">{{item.name}}</span>
                    <span class="index">{{index}}</span>
                  </div>
                  <div class="remark ellipse-text">{{item.remark}}</div>

                  <div class="opts">
                    <el-button type="warning" plain round size="small" @click.stop="showPlatformForm(true, item)">
                      <el-icon><Edit32Filled/></el-icon>
                    </el-button>
                    <el-button type="danger" plain round size="small" @click.stop="showPlatformForm(true, item)">
                      <el-icon><Delete28Filled/></el-icon>
                    </el-button>
                  </div>
            </el-card>
          </el-col>
        </el-row>

        <div id="moreEl" style="color: green;min-height:1px">
          <template v-if="loading">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><circle cx="18" cy="12" r="0" fill="currentColor"><animate attributeName="r" begin=".67" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="12" r="0" fill="currentColor"><animate attributeName="r" begin=".33" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="6" cy="12" r="0" fill="currentColor"><animate attributeName="r" begin="0" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle></svg>
          </template>
        </div>
    
    <div class="float-buttons">
      <el-button type="primary" size="" round @click="refresh()"
        title="刷新平台">
        <el-icon><ArrowCounterclockwise28Filled /></el-icon>
      </el-button>
      <el-button type="primary" size="" round @click="showPlatformForm(false, null)"
        title="添加">
        <el-icon><AddCircle32Regular /></el-icon>
      </el-button>
      <el-button type="warning" size="" round @click="clearAesInfo()"
        title="重载">
        <el-icon><SubtractCircleArrowForward20Regular /></el-icon>
      </el-button>
    </div>

    <InitCheck ref="initCheckRef" @verified="initData"></InitCheck>
    <PlatformForm ref="platformFormComponent" @saved="addPlatform" @updated="updatePlatform"></PlatformForm>
    <PlatformDetail ref="platformDetailComponent"></PlatformDetail>
  </div>
</template>

<script src="./pwdbox.js" ></script>

<style scoped lang="scss">
.wrapper{
  //display: flex;
  //flex-direction: row;
}
.float-buttons{
  position: fixed;
  bottom: 1rem;
  right: 1rem;
  //width: 3rem;
  //text-align: right;
  display: flex;
  flex-direction: column;

  .el-button{
    width: 2.5rem;
    height: 2.5rem;
    border-radius: 50%;
    margin-top: 4px;
    margin-left: 4px !important;
    
    .el-icon{
      width: 1.8rem;
      height: 1.8rem;
      color: white;
    }
  }

}

.card-item{
  font-size: 0.9rem;
  padding: 0.1rem 0.1rem 0.6rem 0.1rem;
  margin: 0.2rem 0.1rem 1.4rem 0.1rem;
  //width: 10rem;
  //display: inline-flex;
  //width: 16rem;
  //height: 8rem;
  width: calc( 100% - 1.5rem );
  height: calc( 100% - 1.5rem );
  position: relative;
  background-color: #efeeee;

  .opts{
    position: absolute;
    bottom: 2px;
    right: 0;
    width: 100%;
    display: none;
    text-align: right;
  }

  &:hover{
    cursor: pointer;
    .name, .remark{
      color: #28853d;
    }
    .opts{
      display: unset;
    }
  }
}
.name{
  padding: 0.2rem 0rem;
  font-weight: 600;
  position: relative;
  .name-text{
    padding-left: 1rem;
    width: calc( 100% - 2rem);
  }
  .index{
    position: absolute;
    right: 0;
    top: 0;
    color: #7b77bb;
  }
}
.remark{
  -webkit-line-clamp: 2;
}
</style>