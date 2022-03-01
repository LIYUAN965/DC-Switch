<template>
  <div style="text-align: center">
    <h3>切换演练详情</h3>
  </div>
  <div style="margin: 20px 30px 0 20px; width: 95%; display: inline-block;">
    <el-button type="success" size="small" circle>&nbsp;{{moduleProgress[2]}}&nbsp;</el-button>
    切换成功
    <el-button type="danger" size="small" circle>&nbsp;{{moduleProgress[3]}}&nbsp;</el-button>
    切换失败
    <el-button type="warning" size="small" circle>&nbsp;{{moduleProgress[1]}}&nbsp;</el-button>
    进行中
    <el-button type="info" size="small" circle>&nbsp;{{moduleProgress[0]}}&nbsp;</el-button>
    未开始
    <div style="float: right;">
      总计：<span style="font-weight: bold;font-size: 24px;margin: 0 10px 0 5px;">{{moduleProgress[0]+moduleProgress[1]+moduleProgress[2]+moduleProgress[3]}}</span>
      其中：
      切换完成<span style="font-weight: bold;font-size: 24px;margin: 0 10px 0 5px;">{{moduleProgress[2]+moduleProgress[3]}}</span>
      进行中<span style="font-weight: bold;font-size: 24px;margin: 0 10px 0 5px;">{{moduleProgress[1]}}</span>
      未开始<span style="font-weight: bold;font-size: 24px;margin: 0 0 0 5px;">{{moduleProgress[0]}}</span>
    </div>
  </div>
  <div style="margin: 20px 30px 0 20px;width: 95%;">
    <el-table :data="moduleDetails" style="width: 80%;display: inline-block;" height="330"
              :row-class-name="tableRowClassName"
    >
<!--      <el-table-column fixed label="模块" width="70">-->
<!--        <template #default="scope">-->
<!--          <router-link :to="'/switch/drill/detail/'+scope.row.id" target="_blank">-->
<!--            {{ scope.row.module }}-->
<!--          </router-link>-->
<!--        </template>-->
<!--      </el-table-column>-->
      <el-table-column fixed prop="detail" label="内容" width="100" />
      <el-table-column prop="duration" label="切换耗时" width="100" />
      <el-table-column prop="status" label="切换进展" width="100" />
      <el-table-column prop="comment" label="失败原因" width="280" />
      <el-table-column prop="beginTime" label="开始时间" width="100" />
      <el-table-column prop="endTime" label="结束时间" width="100" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="small" @click="editComment(scope.$index, scope.row)">修改失败原因</el-button><br>
          <el-button size="small" @click="editEndTime(scope.$index, scope.row)">修改结束时间</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div style="background: transparent;display: inline-block;text-align: center;margin: 0 20px 0 20px;">
      <div v-if="module.count!==0">
        <el-badge :value=module.count class="item">
          <el-card class="box-card module-progress">
            <template #header>
              <div class="card-header module-header">
                {{module.name}}
              </div>
            </template>
            <div class="text item module-title">开始时间</div>
            <div class="text item module-content">{{module.beginTime}}</div>
            <div class="text item module-title">结束时间</div>
            <div class="text item module-content">{{module.endTime}}</div>
            <div class="text item module-title">整体耗时</div>
            <div class="text item module-content">{{module.duration}}</div>
          </el-card>
        </el-badge>
      </div>
      <div v-else>
        <el-card class="box-card module-progress">
          <template #header>
            <div class="card-header module-header">
              {{module.name}}
            </div>
          </template>
          <div class="text item module-title">开始时间</div>
          <div class="text item module-content">{{module.beginTime}}</div>
          <div class="text item module-title">结束时间</div>
          <div class="text item module-content">{{module.endTime}}</div>
          <div class="text item module-title">整体耗时</div>
          <div class="text item module-content">{{module.duration}}</div>
        </el-card>
      </div>
      <div style="margin: 10px 0 0 20px;text-align: center;">
        <span class="module-title">进度</span>
        <el-progress text-inside="true" stroke-width="24" :percentage=module.percent :status=module.status />
        <span>{{module.content}}</span>
      </div>
    </div>
  </div>

  <div style="margin: 20px 30px 0 20px;width: 95%;">
    <div style="display: inline-block;" v-for="(detail, i) in moduleDetails" :key="i">
      <el-button :type="detail.type" icon="Clock" size="small" circle style="margin: 20px 5px 0 30px;"></el-button>
      {{detail.detail}}
    </div>
  </div>

  <el-dialog v-model="dialogVisible" title="修改结束时间" width="30%" draggable>
    <span>请选择 <b>{{detail}}</b> 的新结束时间</span><br><br>
    <span>原结束时间：</span>{{endTime}}<br>
    <span>新结束时间：</span><el-time-picker v-model="value1" placeholder="请选择" value-format="YYYY-MM-DD HH:mm:ss"></el-time-picker>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="submitEndTime(false)">取消</el-button>
        <el-button type="primary" @click="submitEndTime(true)">确认</el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup>
import {
  Check,
  Clock
} from '@element-plus/icons-vue'
</script>

<script>
import axios from "axios";
import {ElMessageBox, ElMessage} from "element-plus";

export default {
  data() {
    return {
      moduleDetails: [
        // {
        //   id: 1,
        //   detail: 'oracle1',
        //   status: '失败',
        //   beginTime: '00:00:00',
        //   endTime: '00:03:02',
        //   duration: '00:30:50',
        //   comment: 'No. 189, Grove St, Los Angeles',
        //   type: 'success'
        // },
      ],
      module: {
        // id: 1,
        // name: 'oracle',
        // beginTime: '00:00:00',
        // endTime: '01:00:32',
        // duration: '01:00:32',
        // count: 2,
        // status: 'exception',
        // percent: 67,
        // content: '2/3'
      },
      moduleProgress: {
        0: 0,
        1: 0,
        2: 0,
        3: 0
      },
      beginTime: '',
      endTime: '',
      detail: '',
      detailId: 0,
      dialogVisible: false,
      value1: '',
      row: null
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    editEndTime(index, row) {
      this.row = row
      this.beginTime = row.beginTime
      this.endTime = row.endTime
      this.detail = row.detail
      this.detailId = row.id
      this.dialogVisible = true
    },
    editComment(index, row) {
      ElMessageBox.prompt('请输入 '+row.detail+' 切换失败的原因', '修改', {
        confirmButtonText: '确认',
        cancelButtonText: '取消',
        inputPattern: /^.{5,}$/,
        inputErrorMessage: '不得少于五个字',
        inputValue: row.comment
      })
          .then(({ value }) => {
            if (row.comment != value) {
              this.editCommentByDetailId(row.id, value);
              ElMessage({
                type: 'success',
                message: `修改成功`,
              });
              row.comment = value
            } else {
              ElMessage({
                type: 'info',
                message: `无改动`,
              });
            }
          })
          .catch(() => {
            ElMessage({
              type: 'info',
              message: '取消修改',
            })
          })
    },
    tableRowClassName({row, rowIndex}) {
      if (row.status === '失败') {
        return 'warning-row'
      }
    },
    init() {
      this.getDetailsByModuleId(this.$route.params.id)
      this.getProgressByModuleId(this.$route.params.id)
    },
    getDetailsByModuleId(moduleId) {
      const url = 'http://127.0.0.1:80/switch/module/details/' + moduleId
      axios.get(url).then(
          res => {
            this.moduleDetails = []
            const STATUS = {0:'未开始', 1:'进行中', 2:'成功', 3:'失败'}
            const TYPE = {0:'info', 1:'warning', 2:'success', 3:'danger'}
            res.data['details'].forEach((item) => {
              let detail = {}
              detail['id'] = item['Id']
              detail['detail'] = item['Detail']
              detail['status'] = STATUS[item['Status']]
              detail['type'] = TYPE[item['Status']]
              detail['beginTime'] = item['BeginTime'].substring(11)
              detail['endTime'] = item['EndTime'].substring(11)
              detail['comment'] = item['Comment']
              detail['duration'] = item['Duration'].substring(0,8)
              this.moduleDetails.push(detail)
            })
          },
          error => {
            console.log(error)
          }
      )
    },
    getProgressByModuleId(moduleId) {
      const url = 'http://127.0.0.1:80/switch/module/progress/' + moduleId
      axios.get(url).then(
          res => {
            this.module['name'] = res.data['progress']['Module']
            this.module['beginTime'] = '00:00:00'
            this.module['endTime'] = '00:00:00'
            this.module['duration'] = '00:00:00'
            if (res.data['progress']['BeginTime']) {
              this.module['beginTime'] = res.data['progress']['BeginTime'].substring(11)
            }
            if (res.data['progress']['EndTime']) {
              this.module['endTime'] = res.data['progress']['EndTime'].substring(11)
            }
            if (res.data['progress']['Duration']) {
              this.module['duration'] = res.data['progress']['Duration'].substring(0,8)
            }
            let finished = 0
            let total = 0
            res.data['progress']['ModuleStatus'].forEach((item) => {
              if (item['Status'] === 3) {
                this.module['count'] = item['Count']
                this.module['status'] = 'exception'
                finished += item['Count']
              } else if (item['Status'] === 2) {
                finished += item['Count']
              }
              total += item['Count']
              this.moduleProgress[item['Status']] = item['Count']
            })
            this.module['content'] = finished+'/'+total
            this.module['percent'] = Math.round(finished/total*100)
          },
          error => {
            console.log(error)
          }
      )
    },
    editCommentByDetailId(id, comment) {
      const url = 'http://127.0.0.1:80/switch/detail/editComment/'+id
      let data = {
        "comment": comment
      }
      axios.get(url, {params:data}).then(
          res => {
            this.$router.go(-1);
          },
          error => {
            console.log(error)
          }
      )
    },
    editEndTimeByDetailId(id, endTime) {
      const url = 'http://127.0.0.1:80/switch/detail/editEndTime/'+id
      let data = {
        "endTime": endTime
      }
      axios.get(url, {params:data}).then(
          res => {
            this.$router.go(-1);
          },
          error => {
            console.log(error)
          }
      )
    },
    submitEndTime(dialogVisible) {
      if (dialogVisible == false) {
        this.dialogVisible = false
        this.value1 = ''
      } else {
        let newEndTime = this.value1.split(' ')[1]
        if (newEndTime <= this.endTime) {
          ElMessage({
            showClose: true,
            message: '新结束时间不得小于（等于）原结束时间',
            type: 'error',
          })
        } else {
          this.editEndTimeByDetailId(this.detailId, this.value1)
          ElMessage({
            type: 'success',
            message: `修改成功`,
          });
          this.dialogVisible = false
          this.value1 = ''
          this.row.endTime = newEndTime
        }
      }
    }
  },
}
</script>

<style>
.demo-shadow {
  height: 100px;
  width: 18%;
  display: inline-block;
  margin: 20px 0 0 20px;
  line-height: 100px;
  padding-left: 5%;
}

.el-table .warning-row {
  --el-table-tr-bg-color: var(--el-color-warning-lighter);
  color: red;
}

.el-table .success-row {
  --el-table-tr-bg-color: var(--el-color-success-lighter);
  color: green;
}

.module-content {
  font-weight: bold;
  font-size: 20px;
}

.module-title {
  color: #0d5aa7;
}

.module-header {
  font-size: 20px;
  font-weight: bold;
}

.module-progress {
  width: 160px;
  min-height: 245px;
  margin-left: 20px;
  cursor: pointer;
}
</style>
