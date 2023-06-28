<template>
  <!-- 修改密码 -->
  <div>
    <el-dialog v-model="visible" 
      :fullscreen="false"
      :close-on-click-modal="false" center align-center
      :width="'60%'">
            <el-form
              ref="modifyFormRef"
              :model="dataForm"
              :rules="rules"
              label-width="5rem"
              v-loading="saving"
            >
              <el-form-item label="新密码" prop="newPassword">
                <el-input type="password" v-model="dataForm.newPassword" placeholder="您的新密码" show-password />
              </el-form-item>
            </el-form>

            <div class="form-buttons" style="border-top:none">
              <el-button type="warning" @click="save" :disabled="saving" >保存</el-button>
              <el-button type="" @click="close" :disabled="saving" >取消</el-button>
            </div>

    </el-dialog>
  </div>
</template>

<script setup>
import { ref, inject, onMounted} from 'vue'
import { AccountService } from '@/../wailsjs/index'
import { ElMessageBox } from 'element-plus'

const visible = ref(false) // 是否可见
const accountData = ref(null) // 是否有验证数据
const saving = ref(false) // 表单保存

const modifyFormRef = ref(null)
const dataForm = ref({
  newPassword: null,
})
const rules = ref({
  newPassword: [{ required: true,  message: '请输入新密码', trigger: 'blur' }],
})

const Noti = inject("Noti")
const Message = inject("Message")
const emit = defineEmits([ "saved" ])

// 保存key和iv
const save =  () => {
  modifyFormRef.value.validate(valid => {
    // console.log("dataForm", accountData.value)
    if(!valid) {
      return
    }

    // save
    ElMessageBox.confirm('修改密码！', '确认修改账户 '+accountData.value.username + ' 的密码吗?', {
        confirmButtonText: '修改',
        cancelButtonText: '再想想',
        type: 'warning',
        center: true,
    }).then(() => {
      saving.value = true
      AccountService.UpdatePwd(accountData.value?.id, dataForm.value.newPassword).then(res => {
        saving.value = false
        if(res) {
          Noti.success({ message: "修改成功", position: 'bottom-right', duration: 2000})
          emit("saved")
          close()
        } else {
          Message.error("修改失败")
        }
      })
    })
    // .catch(() => {
      // ElMessage({
      //   type: 'info',
      //   message: 'Delete canceled',
      // })
    // })
    
  })
}

const show = (data) => {
  visible.value = true
  accountData.value = data
  dataForm.value.newPassword = null
}
const close = () => {
  visible.value = false
  accountData.value = null
}

onMounted(() => {
  // checkHasVerifyData()
})

defineExpose({
  show, close,
})
</script>

<style scoped lang="scss">
</style>
