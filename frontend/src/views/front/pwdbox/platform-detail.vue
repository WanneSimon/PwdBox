<template>
  <div>
  <el-dialog v-model="visible" 
    :fullscreen="false"
    :close-on-click-modal="false" center align-center
    :width="'60%'"
    :title="platformData ? platformData.name: '***'">

    <el-scrollbar height="100%"  v-loading="loadingPlatform"
        class="plat-wrapper" v-if="!isShowAccount">
        <div class="float-block-white" v-if="platformData.site"> 
          <div class="site ellipse-text-line">{{ platformData.site }}</div> 
          <div class="site-op">
            <el-icon title="打开链接" class="op-icon click-item" @click="openLink(platformData.site)"><EarthFilled/></el-icon>
            <el-icon title="复制链接" class="op-icon click-item" @click="copyText(platformData.site)"><CopyLink/></el-icon>
            <el-icon title="添加账户" class="op-icon click-item" @click="showAccount(false, null)"><AddAlt /></el-icon>
          </div>
          <div class="remark"> {{ platformData.remark }} </div>
        </div>
        <el-divider></el-divider>

        <div class="accounts float-block-white"  v-loading="loadingAccounts">

          <!-- <el-card  >   -->
            <el-row v-for="item,index in accountList" :key="'paccount=' + index" class="card-item">
              <div class="float-op">
                <el-icon class="op-icon click-item" :title="'编辑'" style="color: #abaf0a"
                  @click.stop="showAccount(true, item)"><DocumentEdit20Regular/></el-icon>
                <el-icon class="op-icon click-item" :title="'删除'" style="color: #d53a3a"
                  @click.stop="deleteAccount(item)" ><CloseOutline/></el-icon>
              </div>

              <el-col :xs="24" :sm="5" :md="5" :lg="5" :xl="5" 
                class="inline-icon click-item name" >
                <!-- 用户名 -->
                <el-icon><UserAvatar /></el-icon>
                <span @click="copyText(item.username)">{{maskName(item.username)}}</span>
              </el-col>
              <el-col :xs="24" :sm="3" :md="3" :lg="3" :xl="3" 
                class="inline-icon click-item password" >
                <!-- 密码 -->
                <el-icon><Password /></el-icon>
                <span @click="copyPassword(item.password)">***</span>
                <!-- <span>{{item.password}}</span> -->
              </el-col>
              <el-col :xs="24" :sm="6" :md="6" :lg="6" :xl="6" 
                class="inline-icon click-item phone" >
                <!-- 电话 -->
                <el-icon v-if="item.phone"><Phone /></el-icon>
                <span @click="copyText(item.phone)">{{maskPhone(item.phone)}}</span>
              </el-col>
              <el-col :xs="24" :sm="8" :md="8" :lg="6" :xl="6" 
                class="inline-icon click-item email" >
                <!-- 邮箱 -->
                <el-icon v-if="item.email"><Email /></el-icon>
                <span @click="copyText(item.email)">{{maskEmail(item.email)}}</span>
              </el-col>
              <el-col :xs="24" :sm="24" :md="24" :lg="24" :xl="24" 
                class="inline-icon remark" v-if="item.remark">
                <!-- 备注 -->
                <!-- <el-icon></el-icon> -->
                <div>{{item.remark}}</div>
              </el-col>
            </el-row>
          <!-- </el-card> -->

        </div>
    </el-scrollbar>
    
    <AccountForm ref="accountFormRef" v-else
       @saved="addAccount" @updated="updateAccount" @close="closeAccount"></AccountForm>
  </el-dialog>
  </div>
</template>

<script setup>
import { OpenUrl, PlatformService, AccountService, PwdTool } from "@/../wailsjs/index"
import { ref, reactive, inject, nextTick } from "vue"
import { DocumentEdit20Regular  } from '@vicons/fluent'
import { UserAvatar, Password, Phone, Email, CopyLink, CloseOutline, EarthFilled, AddAlt  } from '@vicons/carbon'
import { maskPhone, maskEmail, maskName } from '@/utils/tools'
import AccountForm from "./account-form.vue"
import { ElMessageBox } from 'element-plus'

const visible = ref(false)
const platformId = ref(null)
const platformData = ref({})
const loadingPlatform = ref(false)
const pageData = reactive({
  page: 1,
  size: 10,
  total: 0,
})

const accountList = ref([])
const loadingAccounts = ref(false)
const accountFormRef = ref(null)

const Noti = inject("Noti")
// const Message = inject("Message")
const isShowAccount = ref(false)

// 获取平台详细信息，保证数据实时正确
const loadPlatform = async () => {
  loadingPlatform.value = true
  let res = await PlatformService.Get(platformId.value).then(res => res)
  platformData.value = res
  loadingPlatform.value = false

  // console.log("loadPlatform", platformData.value)
}
const loadAccounts = async () => {
  if(!platformId.value) {
    return
  }
  loadingAccounts.value = true
  let allAccounts =  await AccountService.List(platformId.value, null, null, null, pageData.page, pageData.size)
    .then(res => res)
  accountList.value = allAccounts
  loadingAccounts.value = false

  // console.log("allAccounts", allAccounts)
} 
const deleteAccount = async (item) => {
  ElMessageBox.confirm('账户删除后无法找回！', '确认删除此账户 '+item.username + '?', {
      confirmButtonText: '删除',
      cancelButtonText: '再想想',
      type: 'warning',
      center: true,
  }).then(() => {
    AccountService.Delete(item.id)
    loadAccounts()
  })
  .catch(() => {
    // ElMessage({
    //   type: 'info',
    //   message: 'Delete canceled',
    // })
  })
}

const show = async (idArg) => {
  platformId.value = idArg
  visible.value =  true

  // platformData.value = {}
  // accountList.value = []
  // console.log("platformId.value", platformId.value)

  await loadPlatform()
  await loadAccounts()
}
const close = () => {
  dataForm.value = emptyForm()
  visible.value =  false
}

const openLink = (uri) => {
  OpenUrl(uri)
}
// 复制文本
const copyText = async (str) => {
  await navigator.clipboard.writeText(str).then(data => data);
  // notice
  Noti.success({ message: "复制成功", position: 'bottom-right', duration: 2000})
}
// 复制密码
const copyPassword = async (str) => {
  let decrptyStr = await PwdTool.DecryptPwd(str).then(res => res)
  // console.log("copyPassword", str, decrptyStr)
  let cpResult = await navigator.clipboard.writeText(decrptyStr).then(data => data)
  // console.log("cpResult", cpResult)
  // notice
  Noti.success({ message: "密码复制成功", position: 'bottom-right', duration: 2000})
}


const showAccount = (isUpdate, data) => {
  isShowAccount.value = true
  let cpdata = isUpdate ? Object.assign({}, data) : { platform_id: platformId.value }
  nextTick(() => {
    // console.log("show", isUpdate, cpdata)
    accountFormRef.value.show(isUpdate, cpdata)
  })
}
const closeAccount = () => {
  isShowAccount.value = false
  // accountFormRef.value.close()
}
const addAccount = (newData) => {
  accountList.value.unshift(newData)
  closeAccount()
}
const updateAccount = (newData) => {
  let arr = accountList.value
  // let index = -1;
  for(let i=0; i<arr.length; i++) {
    if(arr[i].id == newData.id) {
      // index = i
      accountList.value[i] = newData
      break;
    }
  }

  closeAccount()
}

defineExpose({
  show, close, 
})

</script>

<style scoped lang="scss">
//$--el-dialog-bg-color: #f3f3f3d1; 
:deep(.el-dialog){
  background:  #f3f3f3d1 ;
}
//.dialog-style{
//  font-size: 1rem;
//  color: red;
//  background: #f3f3f3d1;
//}

.plat-wrapper{
  font-size: 1rem;
}
.site{
  max-width: 20rem;
  margin-right: 6px;
  display: inline-block;
}
.site-op{
  display: inline-block;
  font-size: 1.1rem;
  width: 10rem;
}

.card-item{
  margin: 0.1rem 1rem 1rem 1rem;
  padding: 0.4rem;
  box-shadow: 0px 0px 12px rgba(0, 0, 0, .12);
  min-height: 2.6rem;
  align-items: center;
  position: relative;

  .float-op{
    position: absolute;
    right: 0.5rem;
    top: 0.2rem;
    font-size: 1.3rem;
    display: none;

    .el-icon{
      margin-left: 4px;
    }
  }
  &:hover{
    .float-op{
      display: unset;
    }
  }
}
.click-item{
  cursor: pointer;
  &:hover{
    transform: scale(1.1);
  }
}

.name{
  color: #005a14;
  //color: red !important;
}
.password{
  color: #098383;
}
.phone, .email{
  color: #97330e;
}

.remark{
  font-size: 0.8rem;
  color: #9b9696;
}

.op-icon{
  color: dodgerblue;
  text-decoration: none;

  &:hover{
    color: #28853d;
  }
}

</style>