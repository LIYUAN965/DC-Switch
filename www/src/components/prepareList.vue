<template>
  <div>
    <el-steps :active="1" finish-status="success" simple style="margin: 0 20px 0 20px;">
      <el-step title="事件升级与预判"></el-step>
      <el-step title="切换准备"></el-step>
      <el-step title="切换演练"></el-step>
    </el-steps>
  </div>
  <div style="margin: 20px 0 0 20px;">
    <el-button type="success" :icon="Check" size="small" circle></el-button>准备完成

    <el-button type="warning" :icon="Close" size="small" circle></el-button>未准备完成
  </div>

  <div>
    <template v-for="(prepare, i) in prepareGroup" :key="i">
      <div
          class="demo-shadow"
          :style="{ boxShadow: `var(--el-box-shadow-base)` }"
      >
        <el-link href="/switch/preparation/detail/1" target="_blank" :underline="false" style="font-size: 20px">
          <el-button :type="prepare.status" :icon="prepare.icon" circle></el-button>
          {{prepare.title}}
        </el-link>
      </div>
    </template>
  </div>

  <div style="text-align: right;margin: 20px 20px 0 0;">
    <el-button type="success" size="large" round>服务台确定准备完成</el-button>
  </div>
</template>

<script setup>
import {
  Check,
  Close
} from '@element-plus/icons-vue'
</script>

<script>
import axios from "axios";

export default {
  data() {
    return {
      prepareGroup: [
        {
          title: 'MQ设置',
          icon: 'Check',
          status: 'success'
        },
        {
          title: '云桌面',
          icon: 'Check',
          status: 'success'
        },
        {
          title: '通讯录',
          icon: 'Check',
          status: 'success'
        },
        {
          title: '切换范围',
          icon: 'Check',
          status: 'success'
        },
        {
          title: 'OPS配置项',
          icon: 'Clock',
          status: 'warning'
        },
        {
          title: 'DB',
          icon: 'Clock',
          status: 'warning'
        },
        {
          title: '堡垒机',
          icon: 'Clock',
          status: 'warning'
        },
      ],
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.getAllPreparesByVersionId(1)
    },
    getAllPreparesByVersionId(versionId) {
      const url = 'http://127.0.0.1:80/switch/preparations/' + versionId
      axios.get(url).then(
          res => {
            this.prepareGroup = []
            res.data['prepares'].forEach((item) => {
              if (item['Sequence'] == 0) {
                let prepare = {}
                prepare['id'] = item['Id']
                prepare['title'] = item['Title']
                if (item['Status'] == 0) {
                  prepare['status'] = 'success'
                  prepare['icon'] = 'Check'
                } else {
                  prepare['status'] = 'waring'
                  prepare['icon'] = 'Clock'
                }
                this.prepareGroup.push(prepare)
              }
            })
          },
          error => {
            console.log(error)
          }
      )
    }
  },
}
</script>

<style scoped>
.demo-shadow {
  height: 100px;
  width: 18%;
  display: inline-block;
  margin: 20px 0 0 20px;
  line-height: 100px;
  padding-left: 5%;
}
</style>
