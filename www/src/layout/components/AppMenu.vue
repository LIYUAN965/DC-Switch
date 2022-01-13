<template>
  <el-menu
    active-text-color="#ffd04b"
    background-color="#545c64"
    class="el-menu-vertical-demo"
    default-active="2"
    text-color="#fff"
    router
    :collapse="$store.state.isCollapse"
  >
    <el-menu-item index="/">
      <el-icon><IconMenu /></el-icon>
      <span>主页</span>
    </el-menu-item>
    <el-menu-item index="/switch/preparation">
      <el-icon><Document /></el-icon>
      <span>切换准备</span>
    </el-menu-item>
    <el-sub-menu index="/permission">
      <template #title>
        <el-icon><Setting /></el-icon>
        <span>切换配置</span>
      </template>
      <el-menu-item index="/permission/permission_role">
        角色
      </el-menu-item>
      <el-menu-item index="/permission/permission_admin">
        管理员
      </el-menu-item>
      <el-menu-item index="/permission/permission_rule">
        权限规则
      </el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script setup>
import {
  Location,
  Document,
  Menu as IconMenu,
  Setting
} from '@element-plus/icons'
import { } from 'vue'

</script>

<style>
.el-menu {
  border-right: null;
}
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
