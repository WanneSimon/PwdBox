
<template>
    <router-view></router-view>
</template>

<script setup lang="ts" >
import { provide } from 'vue';
import { ElMessage } from 'element-plus';
import { ElNotification } from 'element-plus'

// useMessage() 只能在 setup 中使用
// 向子组件注入 Message 和 Dialog， 以便能随处使用

provide("Message", ElMessage)
// const Dialog = useDialog()
// provide("Dialog", Dialog)
// provide("MunkData", MunkData)
provide("Noti", ElNotification)

// return { Message, Dialog }

</script>