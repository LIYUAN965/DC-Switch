<template class="demo-type">
  <div>
    <el-icon><UserFilled /></el-icon>
  </div>
</template>

<script setup>
import {
  UserFilled
} from '@element-plus/icons-vue'

</script>

<style scoped>
.demo-type {
  display: flex;
}
.demo-type > div {
  flex: 1;
  text-align: center;
}

.demo-type > div:not(:last-child) {
  border-right: 1px solid var(--el-border-color-base);
}
</style>
