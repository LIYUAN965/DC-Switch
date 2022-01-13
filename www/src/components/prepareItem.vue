<template>
  <div style="text-align: center">
    <h3>切换准备详情</h3>
  </div>
  <div style="margin: 20px 0 0 20px;">
    <el-button :type="prepare.status" :icon="prepare.icon" circle></el-button>
    {{prepare.title}}
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
</style>
