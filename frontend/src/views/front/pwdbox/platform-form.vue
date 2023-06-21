<template>
  <el-dialog v-model="visiable" 
    :close-on-click-modal="false" width="30rem" center align-center
    :title="isUpdate ? '编辑平台': '新增平台'">
    <el-scrollbar height="22rem"  >
      <el-form
        ref="platformFormRef"
        :model="dataForm"
        :rules="rules"
        label-width="5rem"
        v-loading="saving"
      >
        <el-form-item label="名称" prop="name">
          <el-input v-model="dataForm.name" />
        </el-form-item>
        <el-form-item label="主页地址" prop="site">
          <el-input v-model="dataForm.site" />
        </el-form-item>
        <el-form-item label="说明" prop="remark">
          <el-input v-model="dataForm.remark" type="textarea" :rows="4"/>
        </el-form-item>
        <el-form-item label="创建时间" prop="create_time" v-if="dataForm.create_time">
          <!-- <el-input v-model="dataForm.create_time" /> -->
          {{ dataForm.create_time }}
        </el-form-item>
        <!-- <el-form-item label="图标" prop="img">
          <el-input v-model="dataForm.img" />
        </el-form-item> -->
      </el-form>
    </el-scrollbar>

    <div class="form-buttons">
      <el-button type="primary" v-if="!isUpdate" @click="save" :disabled="saving" >保存</el-button>
      <el-button type="primary" v-else  @click="update" :disabled="saving" >更新</el-button>
      <el-button type="" @click="close" :disabled="saving">取消</el-button>
    </div>
  </el-dialog>
</template>

<script setup>
import { PlatformService } from "@/../wailsjs/index"
import { reactive, ref } from 'vue';

const visiable = ref(false)
const isUpdate = ref(false)
// const data = ref(null)
const saving = ref(false)

const emptyForm = () => {
  return {
    id: null, 
    name: null,
    site: null, 
    remark: null,
    num: null,
    create_time: null,
    img: null,
  }
}

const platformFormRef = ref(null)
const dataForm = ref(emptyForm())
const rules = reactive({
  name: [{ required: true, message: '平台名称', trigger: 'blur' }]
})

const emit = defineEmits([ "saved", "updated" ])


const save = () => {
  // let rcomf = this.$refs
  platformFormRef.value.validate(valid => {
    console.log("dataForm", dataForm.value)
    if(!valid) {
      return
    }

    saving.value = true
    PlatformService.Save(dataForm.value).then(res => {
      saving.value = false
      console.log("saved", res)
      emit("saved", res) // 返回新建数据
    }, err => {
      saving.value = false
      console.error(err)
    })
  })
}
const update = () => {
  // let rcomf = this.$refs
  platformFormRef.value.validate(valid => {
    console.log("dataForm", dataForm.value)
    if(!valid) {
      return
    }

    saving.value = true
    PlatformService.Update(dataForm.value).then(res => {
      saving.value = false
      console.log("updated", res)
      emit("updated", res) // 返回更新数据
    }, err => {
      saving.value = false
      console.error(err)
    })
  })
}

const show = (isUpdateArg, dataArg) => {
  isUpdate.value = isUpdateArg
  dataForm.value = dataArg ? dataArg : emptyForm()
  visiable.value =  true
}
const close = () => {
  dataForm.value = emptyForm()
  visiable.value =  false
}

defineExpose({
  show, close, 
})
</script>

<style scoped lang="scss">
</style>