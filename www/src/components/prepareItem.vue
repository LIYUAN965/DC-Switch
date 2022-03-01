<template>
  <div style="text-align: center">
    <h3>切换准备详情</h3>
  </div>
  <div style="margin: 20px 0 0 20px;">
    <el-button :type="prepare.status" :icon="prepare.icon" circle></el-button>
    {{prepare.title}}
  </div>
  <div style="margin: 20px 30px 0 20px;">
    <el-row :gutter="20" :model="prepare" >
      <span class="ml-3 w-35 text-gray-600 inline-flex items-center"
      >准备项</span>
      <el-input v-model="prepare.title" readonly></el-input>
    </el-row>
    <el-row :gutter="20" :model="prepare" >
      <span class="ml-3 w-35 text-gray-600 inline-flex items-center"
      >详细描述</span>
      <el-input v-model="prepare.content" type="textarea" :rows="15" placeholder="请输入相关准备信息"></el-input>
      <el-alert title="详细描述不得为空！" type="error" show-icon v-show="isShow"/>
    </el-row>
    <el-row :gutter="20" justify="end">
      <el-button type="primary" @click="editPrepareById(prepare.id)">准备完成</el-button>
      <el-button @click="$router.go(-1);">返回</el-button>
    </el-row>
  </div>
</template>

<script setup>
import {
  Check,
  Clock
} from '@element-plus/icons-vue'
</script>

<script>
import axios from "axios";
import {ElMessageBox} from "element-plus";

export default {
  data() {
    return {
      prepare: {
          title: '',
          icon: 'Check',
          status: 'success',
          content: ''
      },
      isShow: false
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.getPrepareById(this.$route.params.id)
    },
    getPrepareById(id) {
      const url = 'http://127.0.0.1:80/switch/preparation/' + id
      axios.get(url).then(
          res => {
            this.prepare['id'] = res.data['prepare']['Id']
            this.prepare['title'] = res.data['prepare']['Title']
            this.prepare['content'] = res.data['prepare']['Content']
            if (res.data['prepare']['Status'] == 0) {
              this.prepare['status'] = 'warning'
              this.prepare['icon'] = 'Clock'
            } else {
              this.prepare['status'] = 'success'
              this.prepare['icon'] = 'Check'
            }
          },
          error => {
            console.log(error)
          }
      )
    },
    editPrepareById(id) {
      console.log(id)
      if (!this.prepare.content) {
        this.isShow = true
        return
      }
      this.isShow = false
      const url = 'http://127.0.0.1:80/switch/preparation/edit/'+id
      let data = {
        "title": this.prepare.title,
        "content": this.prepare.content
      }
      ElMessageBox.confirm('确认准备信息填写完成?')
        .then(() => {
          axios.get(url, {params:data}).then(
              res => {
                this.$router.go(-1);
              },
              error => {
                console.log(error)
              }
          )
        })
        .catch(() => {
          // catch error
        })
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
