<template>
  <div @click="handleCollapse">
    <el-icon
      :size="20"
      style="vertical-align: middle;"
      v-show="!store.state.isCollapse"
    >
      <fold />
    </el-icon>
    <el-icon
      :size="20"
      v-show="store.state.isCollapse"
    >
      <expand />
    </el-icon>
  </div>
</template>

<script setup>
import {
  Fold,
  Expand
} from '@element-plus/icons'

import store from '@/store/index.js'

const handleCollapse = () => {
  store.commit('setIsCollapse', !store.state.isCollapse)
}
</script>

<style lang="scss" scoped>
div{
  display: inline-block;
  margin-right: 10px;
}
</style>
