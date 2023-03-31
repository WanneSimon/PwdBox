<template>
  <metainfo style="--wails-draggable:drag">
    <template v-slot:title="{ content }">{{ (content ? content+' - ':'') + 'munkblog' }}</template>
  </metainfo>
</template>

<script>
import { computed } from 'vue'
import { useMeta } from 'vue-meta'
export default {
  props: {
    // vue-meta 创建时用到的配置
    metaInfo: { 
      type: Object,
      default: () => {
        return {
          title: 'coco-rush', // 默认标题
          description: '',
        }
      }
    }
  },
  setup(props, context) {
    // console.log("props-metaInfo", props.metaInfo)
    // const metaInfo = props.metaInfo
    
    // const computedMeta = computed(() => metaInfo)
    // const { meta, onRemoved } = useMeta(computedMeta)

    // const { onRemoved } = useMeta(metaInfo)

    // onRemoved(() => {
      // Do something as soon as this metadata is removed,
      // eg because the component instance was destroyed
    // })
  },

  // watch: {
  //   metaInfo(newValue, oldValue) {
  //     console.log('metaInfo', oldValue, newValue)
  //   }
  // },
  created () {
    // console.log("created", this.metaInfo)
    const computedMeta = computed(() => this.metaInfo)
    // const { meta, onRemoved } = useMeta(computedMeta)
    useMeta(computedMeta)
  }

}
</script>
