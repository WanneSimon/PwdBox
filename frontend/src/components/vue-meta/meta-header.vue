<template>
  <metainfo style="--wails-draggable:drag">
    <template v-slot:title="{ content }">{{ (content ? content+' - ':'') + 'munkblog' }}</template>
  </metainfo>
</template>

<script setup lang="ts" >
import { computed } from 'vue'
import { useMeta } from 'vue-meta'

const props = defineProps({
  // vue-meta 创建时用到的配置
  metaInfo: {
    type: Object,
    default: () => {
      return {
        title: 'coco-rush', // 默认标题
        description: '',
      }
    }
  },
})

const computedMeta = computed(() => props.metaInfo)
// const { meta, onRemoved } = useMeta(computedMeta)
useMeta(computedMeta)

</script>
