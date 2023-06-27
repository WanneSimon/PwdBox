<template>
  <!-- 检查设置的 key 和 iv，不存在时要求输入新的 -->
  <div>
    <el-dialog v-model="visible" 
      :fullscreen="false"
      :close-on-click-modal="false" center align-center
      :width="'60%'">
        <el-scrollbar height="100%">

          <div v-if="!showForm">
            checking...
          </div>
          
          <div v-else>
            <el-form
              ref="verifyFormRef"
              :model="dataForm"
              :rules="rules"
              label-width="5rem"
              v-loading="saving || checking"
            >
              <el-form-item label="Key" prop="key">
                <el-input v-model="dataForm.key" placeholder="16位key" />
              </el-form-item>
              <el-form-item label="IV" prop="iv">
                <el-input v-model="dataForm.iv"  placeholder="16位iv"/>
              </el-form-item>
            </el-form>
            <div class="form-buttons">
              <el-button type="primary" v-if="!hasData" @click="save" :disabled="saving" >保存</el-button>
              <el-button type="primary" v-else  @click="verify" :disabled="checking" >确认</el-button>
            </div>
          </div>
  
        </el-scrollbar>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, inject, onMounted} from 'vue'
import { PwdTool } from '@/../wailsjs/index'

const visible = ref(true) // 是否可见
const hasData = ref(true) // 是否有验证数据
const checking = ref(false) // 正在进行验证
const saving = ref(false) // 表单保存
const showForm = ref(false) // 是否显示表单

const verifyFormRef = ref(null)
const dataForm = ref({
  key: null,
  iv: null,
})
const rules = ref({
  key: [{ required: true, len: 16, message: '请输入key，且长度为16位', trigger: 'blur' }],
  iv : [{ required: true, len: 16, message: '请输入iv，且长度为16位' , trigger: 'blur' }]
})

const Noti = inject("Noti")
const Message = inject("Message")

const emit = defineEmits([ "verified" ])

// 检查是否存在验证数据
const checkHasVerifyData = async () => {
  let re = await PwdTool.ExistVerifyData().then(res => res)
  // console.log("hasData", re)
  hasData.value = re

  // console.log("检查是否存在验证数据", re)
  // 如果有验证数据，则尝试加载缓存，并验证
  if(re) {
    if(loadTempData()) {
      // verify()
      innerVerify(false)
      return
    } 
  }
  
  showForm.value = true
}

// 保存key和iv
const save =  () => {
  verifyFormRef.value.validate(valid => {
    // console.log("dataForm", dataForm.value)
    if(!valid) {
      return
    }

    // save
    checking.value = true
    PwdTool.SaveAesInfo(dataForm.value.key, dataForm.value.iv).then(res => {
      checking.value = false
      if(res) {
        Noti.success({ message: "保存成功", position: 'bottom-right', duration: 2000})
        saveTempData(dataForm.value.key, dataForm.value.iv)
        verifySuccess()
      } else {
        Message.error("保存失败")
      }
    })
    
  })
}

// 验证key和iv
const verify = () => {
  verifyFormRef.value.validate(valid => {
    // console.log("dataForm", dataForm.value)
    if(!valid) {
      return
    }

    // verify
    innerVerify()
  })
}
const innerVerify = async (needNotice=true) => {
  // console.log("innerVerify", innerVerify)
  saving.value = true
  await PwdTool.VerifyAndKeepAesInfo(dataForm.value.key, dataForm.value.iv).then(res => {
    saving.value = false
    if(res) {
      if(needNotice) {
        Noti.success({ message: "验证成功", position: 'bottom-right', duration: 2000})
      }
      saveTempData(dataForm.value.key, dataForm.value.iv)
      verifySuccess()
    } else {
      showForm.value = true
      Message.error("验证失败")
    }
  }, err => {
    saving.value = false
    showForm.value = true
    Message.error(err)
  })
}

// 本次应用中保存 kv,iv
const saveTempData = (key,iv) => {
  let pwdboxJson = { key, iv }
  sessionStorage.setItem("pwdbox", JSON.stringify(pwdboxJson))
}  
// 读取 kv,iv
const loadTempData = () => {
  let jsonstr = sessionStorage.getItem("pwdbox")
  if(!jsonstr) {
    return false
  }

  let pwdboxJson = JSON.parse(jsonstr)
  dataForm.value.key = pwdboxJson.key
  dataForm.value.iv = pwdboxJson.iv
  return true
} 

const verifySuccess = () => {
  visible.value = false
  emit("verified")
}

const showAndCheck = () => {
  visible.value = true
  checkHasVerifyData()
}

onMounted(() => {
  checkHasVerifyData()
})

defineExpose({
  showAndCheck,
})
</script>

<style scoped lang="scss">
</style>
