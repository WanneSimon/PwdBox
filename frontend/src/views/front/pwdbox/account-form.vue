<template>
  <div class="float-block-white">
    <el-page-header @back="close">
      <template #content>
        <span class="text-large font-600 mr-3"> {{ title }} </span>
      </template>
    </el-page-header>
    
    
    <div>
      <el-scrollbar height="22rem"
          class="form-wrapper">
        <el-form
            ref="accountFormRef"
            :model="dataForm"
            :rules="rules"
            label-width="5rem"
            v-loading="saving"
        >
            <el-form-item label="用户名" prop="username">
              <el-input v-model="dataForm.username" />
            </el-form-item>
            <el-form-item label="密码" prop="password" v-if="!isUpdate">
              <el-input v-model="dataForm.password" type="password" show-password />
            </el-form-item>
            <!-- <el-form-item label="密码" prop="password" v-else >
              <el-button type="warning" @click="showModifyPassword(dataForm)">密码</el-button>
            </el-form-item> -->
            <el-form-item label="电话" prop="phone">
              <el-input v-model="dataForm.phone" />
            </el-form-item>
            <el-form-item label="邮箱" prop="email">
              <el-input v-model="dataForm.email" />
            </el-form-item>
            <el-form-item label="说明" prop="remark">
              <el-input v-model="dataForm.remark" type="textarea" :rows="4"/>
            </el-form-item>
            <el-form-item label="创建时间" prop="create_time" v-if="dataForm.create_time">
              {{ dataForm.create_time }}
            </el-form-item>
        </el-form>
      </el-scrollbar>
      <div class="form-buttons">
        <el-button type="primary" v-if="!isUpdate" @click="save" :disabled="saving" >保存</el-button>
        <el-button type="primary" v-else  @click="update" :disabled="saving" >更新</el-button>
        <el-button type="" @click="close" :disabled="saving">取消</el-button>
      </div>
    </div>
  
    <!-- <ModifyPassword ref="modifyPasswordRef"></ModifyPassword> -->
  </div>
</template>

<script setup>
import { AccountService } from "@/../wailsjs/index"
import { reactive, ref, computed, inject } from 'vue'
// import ModifyPassword from "./modify-password.vue"

const title = computed(() => {
  return isUpdate.value ? '编辑' : '添加'
})
// const title = ref('')
const emit = defineEmits(["close", "saved", "updated"])
const saving = ref(false)
const accountFormRef = ref(null)
// const modifyPasswordRef =ref(null)

const emptyAccountForm = () => {
  return {
    id: null, 
    platform_id: null,
    username: null, 
    password: null,
    phone: null,
    email: null,
    remark: null,
    create_time: null,
  }
}

const isUpdate = ref(false)
const dataForm = ref(emptyAccountForm())
const rules = reactive({
  username: [{ required: true, message: '用户名', trigger: 'blur' }],
  password: [{ required: true, message: '密码', trigger: 'blur' }]
})

const Message = inject("Message")

const save = () => {
  // let rcomf = this.$refs
  accountFormRef.value.validate(valid => {
    // console.log("dataForm", dataForm.value)
    if(!valid) {
      return
    }

    saving.value = true
    AccountService.Save(dataForm.value).then(res => {
      saving.value = false
      console.log("saved", res)

      if(res) {
        emit("saved", res) // 返回新建数据
      } else {
        Message.error("保存失败")
      }
    }, err => {
      saving.value = false
      console.error(err)
    })
  })
}
const update = () => {
  // let rcomf = this.$refs
  accountFormRef.value.validate(valid => {
    console.log("dataForm", dataForm.value)
    if(!valid) {
      return
    }

    saving.value = true
    AccountService.Update(dataForm.value).then(res => {
      saving.value = false
      console.log("updated", res)
      emit("updated", res) // 返回更新数据
    }, err => {
      saving.value = false
      console.error(err)
    })
  })
}

// 必须要传入平台id
const show = (isUpdateArg, dataArg) => {
  isUpdate.value = isUpdateArg
  if(!dataArg?.platform_id || dataArg?.platform_id==0){
    console.error("miss platform")
    close()
    return
  }

  // dataForm.value = dataArg ? dataArg : emptyAccountForm()
  dataForm.value = dataArg
  // console.log("show", dataForm.value)
}

const close = () => {
  emit("close")
}

// const showModifyPassword = (account) => {
//   modifyPasswordRef.value.show(account)
// }

defineExpose({
  show,
})

</script>

<style scoped lang="scss">
.form-wrapper{
  margin-top: 1rem;
}

</style>