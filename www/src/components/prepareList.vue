<template>
  <div>
    <el-steps :active="1" finish-status="success" simple style="margin: 0 20px 0 20px;">
      <el-step @click="goTo(1)" title="事件升级与预判" style="cursor: pointer;"></el-step>
      <el-step @click="goTo(2)" title="切换准备" style="cursor: pointer;"></el-step>
      <el-step @click="goTo(3)" title="切换演练" style="cursor: pointer;"></el-step>
    </el-steps>
  </div>

  <div style="margin: 20px 0 0 20px;">
    <el-button type="success" :icon="Check" size="small" circle></el-button>准备完成
    <el-button type="warning" :icon="Close" size="small" circle></el-button>未准备完成
  </div>

  <div>
    <template v-for="(prepare, i) in prepareGroup" :key="i">
      <router-link :to="'/switch/preparation/detail/'+prepare.id" style="font-size: 20px; text-decoration: none;">
        <div class="demo-shadow" :style="{ boxShadow: `var(--el-box-shadow-base)` }">
          <el-button :type="prepare.status" :icon="prepare.icon" circle></el-button>
          {{prepare.title}}
        </div>
      </router-link>
    </template>
  </div>

  <div style="text-align: right;margin: 20px 20px 0 0;">
    <el-button :type="fuwutaiConfirm.status" size="large" round @click="totalConfirm">{{fuwutaiConfirm.title}}</el-button>
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
import { h } from 'vue';
import {ElMessage} from "element-plus";

export default {
  data() {
    return {
      prepareGroup: [
        {
          title: 'MQ设置',
          icon: 'Check',
          status: 'success'
        }
      ],
      fuwutaiConfirm: {
        title: 'MQ设置',
        icon: 'Check',
        status: 'success'
      }
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
              let status = 'warning'
              let icon = 'Clock'
              if (item['Status'] == 1) {
                status = 'success'
                icon = 'Check'
              }
              let prepare = {}
              prepare['id'] = item['Id']
              prepare['title'] = item['Title']
              prepare['status'] = status
              prepare['icon'] = icon
              if (item['Sequence'] == 0) {
                this.prepareGroup.push(prepare)
              } else {
                this.fuwutaiConfirm = prepare
              }
            })
          },
          error => {
            console.log(error)
          }
      )
    },
    totalConfirm() {
      let notFinished = ''
      this.prepareGroup.forEach((item) => {
        if (item['status'] == 'warning') {
          notFinished += ' ' + item['title']
        }
      })
      if (notFinished != '') {
        ElMessage({
          message: h('p', null, [
            h('span', null, '以下操作域还未准备完成：'),
            h('i', { style: 'color: red' }, notFinished),
          ]),
        })
        // ElMessage.error(item['title']+'准备项还未完成！')
        return
      }
      const url = '/#/switch/preparation/detail/'+this.fuwutaiConfirm.id
      window.open(url, '_self')
    },
    goTo(index) {
      let url = '/#/switch/preparation'
      if (index===3) {
        url = '/#/switch/drill'
      }
      window.open(url, '_self')
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
