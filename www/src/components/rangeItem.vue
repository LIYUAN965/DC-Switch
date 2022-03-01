<template>
  <div style="text-align: center">
    <h3>{{switchDomain}}切换范围</h3>
  </div>
  <div style="margin: 20px 30px 0 20px; width: 95%; display: inline-block;">
    <el-tabs
        v-model="activeModuleName"
        class="demo-tabs"
        @tab-click="handleClick"
    >
      <el-tab-pane v-for="(module,moduleId) in moduleList" :key="moduleId" :label=module :name=module>
        {{module}}
        <el-upload
            accept=".xls,.xlsx,.csv"
            :multiple="false"
            :show-file-list="false"
            :http-request="handleChange"
            action='string'
        >
          <el-button size="medium" type="primary" @click="updateModule(moduleId, module)">上传文件</el-button>
        </el-upload>
        <el-table :data="tableBody[module]" border  height="400" style="width: 100%">
          <el-table-column v-for="header in tableHeader[module]">
            <template #header>
              {{header}}
            </template>
            <template #default="scope">
              {{scope.row[header]}}
            </template>
          </el-table-column>
        </el-table>
      </el-tab-pane>
    </el-tabs>
  </div>


</template>

<script>
import axios from "axios";
import {ElMessageBox, ElMessage} from "element-plus";
import {UploadFile} from 'element-plus/es/components/upload/src/upload.type.mjs'

export default {
  data() {
    return {
      switchDomain: 'PE',
      activeModuleName: '专线',
      activeModuleId: 1,
      moduleList: {1:'专线', 2:'域名', 3:'调度', 4:'中间件'},
      fileList: [
        {
          name: 'food.jpeg',
          url: 'https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100',
        },
      ],
      myHeaders: {
        'Content-Type': 'multipart/form-data'
      },
      uploadApi: 'http://127.0.0.1:80/switch/range/upload/',
      tableHeader: {
        '专线': ['first', 'second', 'third'],
      },
      tableBody: {
        '专线': [
          {
            'first': '2016-05-03',
            'second': 'Tom',
            'third': 'No. 189, Grove St, Los Angeles',
          },
          {
            'first': '2016-05-02',
            'second': 'Tom',
            'third': 'No. 189, Grove St, Los Angeles',
          },
        ],
      }
    }
  },
  mounted() {
    this.init()
  },
  watch: {
    $route(to, from) {
      if (this.$route.params.id) {
        console.log(this.activeModuleName)
        this.getModulesByDomainId(this.$route.params.id)
        this.getSwitchDomainById(this.$route.params.id)
      }
    },
  },
  methods: {
    init() {
      this.getModulesByDomainId(this.$route.params.id)
      this.getSwitchDomainById(this.$route.params.id)
    },
    updateModule(id, module) {
      this.activeModuleId = id
      this.activeModuleName = module
    },
    getSwitchDomainById(domainId) {
      const url = 'http://127.0.0.1:80/switch/domain/' + domainId
      axios.get(url).then(
          res => {
            this.switchDomain = res.data['switchDomain']['AppDomain']
          },
          error => {
            console.log(error)
          }
      )
    },
    getModulesByDomainId(domainId) {
      const url = 'http://127.0.0.1:80/switch/modules/domain/' + domainId
      axios.get(url).then(
          res => {
            let header = {}
            let body = {}
            this.moduleList = {}
            res.data['modules'].forEach((item) => {
              this.moduleList[item['Id']] = item['Module']
              header[item['Module']] = []
              body[item['Module']] = []
              if (item['Details']!=null) {
                item['Details'].forEach((row) => {
                  if (header[item['Module']].length==0) {
                    header[item['Module']] = row['Detail']
                  } else {
                    let detail = {}
                    row['Detail'].forEach(function(cell, i) {
                      detail[header[item['Module']][i]] = cell
                    })
                    body[item['Module']].push(detail)
                  }
                })
              }
            })
            this.tableHeader = header
            this.tableBody = body
            this.activeModuleName = Object.values(this.moduleList)[0]
          },
          error => {
            console.log(error)
          }
      )
    },
    getTableByModuleId(moduleId) {
      const url = 'http://127.0.0.1:80/switch/module/domain/' + moduleId
      axios.get(url).then(
          res => {
            let header = []
            let body = []
            if (['Details']!=null) {
              res.data['module']['Details'].forEach((row) => {
                if (header.length==0) {
                  header = row['Detail']
                } else {
                  let detail = {}
                  row['Detail'].forEach(function(cell, i) {
                    detail[header[i]] = cell
                  })
                  body.push(detail)
                }
              })
            }
            this.tableHeader[res.data['module']['Module']] = header
            this.tableBody[res.data['module']['Module']] = body
          },
          error => {
            console.log(error)
          }
      )
    },
    axiosPost(url,params,callBack){
      axios.post(url, params)
          .then(res=>{
            callBack(res)
            this.getTableByModuleId(params.get('moduleId'))
            console.log(params)
            console.log(params.get('moduleName'))
            console.log('-------------------------------')
            this.activeModuleName = params.get('moduleName')
            console.log(this.activeModuleName)
          })
          .catch(err=>{
            console.log('error');
          });
    },
    //文件上传的函数
    handleChange(param) {
      let fd = new FormData();
      let self = this;
      fd.append('file',param.file);//传文件
      fd.append('moduleId',this.activeModuleId);//传
      fd.append('moduleName',this.activeModuleName);//传
      console.log(this.activeModuleName)
      this.axiosPost(
          this.uploadApi+this.activeModuleId, //接口地址
          fd,  //formdata对象参数
          res => {
            console.log(this.activeModuleName)
            console.log(res)
          }
      );
    },
    changeFile (file) {
      console.log('file', file)
      let fd = new FormData()
      fd.append('file', file.file)// 传文件
      fd.append('second_path', '/Users/liyuan/LY/projects/uploads/dcswitch')
      let self = this
      // this.loading = false
      fetch(this.uploadApi, {//fetces6语法，同axios/ajax作用
        method: 'POST',
        headers: {
          credentials: 'same-origin'
          // 'Content-Type': 'multipart/form-data',不用写这个，写了报错
        },
        body: fd//不为data为body否则报错
      })
          .then(res => res.json())
          .then(res => {
            console.log(res)
            //报错提示
            if (res.code === -100) {
              ElMessage({
                showClose: true,
                message: res.msg,
                type: 'error',
              })
              // this.loading = true
              return
            }
            self.loadData(0)
          }).catch(function (e) {
        // this.loading = true
        console.log( e)
      })
    },
    handleClick(tab, event) {
      console.log("handleClick")
      console.log(tab, event)
    },
    handleRemove(file, fileList) {
      console.log("handleRemove")
      console.log(file, fileList)
    },
    handlePreview(file) {
      console.log("handlePreview")
      console.log(file)
    },
    handleExceed(files, fileList) {
      console.log("handleExceed")
      ElMessage.warning(
          `The limit is 3, you selected ${files.length} files this time, add up to ${
              files.length + fileList.length
          } totally`
      )
    },
    beforeRemove(file, fileList) {
      console.log("beforeRemove")
      return ElMessageBox.confirm(`Cancel the transfer of ${file.name} ?`)
    }
  },
}
</script>

<style>
.demo-tabs > .el-tabs__content {
  padding: 32px;
  background-color: #f4f5f7;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}
</style>
