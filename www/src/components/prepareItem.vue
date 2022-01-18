<template>
  <div style="text-align: center">
    <h3>切换准备详情</h3>
  </div>
  <div style="margin: 20px 0 0 20px;">
    <el-button :type="prepare.status" :icon="prepare.icon" circle></el-button>
    {{prepare.title}}
  </div>
  <div style="margin: 20px 30px 0 20px;">
    <el-row :gutter="20" :model="form" >
      <span class="ml-3 w-35 text-gray-600 inline-flex items-center"
      >准备项</span>
      <el-input v-model="form.title" readonly></el-input>
    </el-row>
    <el-row :gutter="20" :model="form" >
      <span class="ml-3 w-35 text-gray-600 inline-flex items-center"
      >详细描述</span>
      <el-input v-model="form.content" type="textarea" :rows="15" placeholder="请输入相关准备信息"></el-input>
    </el-row>
    <el-row :gutter="20" justify="end">
      <el-button type="success" @click="onSubmit">准备完成</el-button>
      <el-button>取消</el-button>
    </el-row>
  </div>
</template>

<script setup>
import {
  Check,
  Close
} from '@element-plus/icons-vue'
import { reactive } from 'vue'

// do not use same name with ref
const form = reactive({
  title: 'mq脚本',
  content: ''
})

const onSubmit = () => {
  console.log('submit!')
}
</script>

<script>
import axios from "axios";

export default {
  data() {
    return {
      prepare: {
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
    },
    getAllPreparesByVersionId(versionId) {
      const url = 'http://127.0.0.1:80/switch/preparations/' + versionId
      axios.get(url).then(
          res => {
            this.prepareGroup = []
            res.data['prepares'].forEach((item) => {
                let prepare = {}
                this.prepareGroup.push(prepare)
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

.el-form-item__label {
  font-size: 20px!important;
  font-weight: bold!important;
}
</style>
