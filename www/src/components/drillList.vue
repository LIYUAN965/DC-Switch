<template>
  <div>
    <el-steps :active="2" finish-status="success" simple style="margin: 0 20px 0 20px;">
      <el-step @click="goTo(1)" title="事件升级与预判" style="cursor: pointer;"></el-step>
      <el-step @click="goTo(2)" title="切换准备" style="cursor: pointer;"></el-step>
      <el-step @click="goTo(3)" title="切换演练" style="cursor: pointer;"></el-step>
    </el-steps>
  </div>

  <div style="margin: 20px 0 0 20px;">
    <el-table :data="tableFailedDetails" style="width: 34%;display: inline-block;" height="250">
      <el-table-column fixed label="模块" width="70">
        <template #default="scope">
          <router-link :to="'/switch/drill/detail/'+scope.row.id" target="_blank">
            {{ scope.row.module }}
          </router-link>
        </template>
      </el-table-column>
      <el-table-column prop="detail" label="内容" width="100" />
      <el-table-column prop="comment" label="失败原因" width="200" />
      <el-table-column prop="beginTime" label="开始时间" width="110" />
      <el-table-column prop="endTime" label="结束时间" width="110" />
    </el-table>
    <div id="pie" ref="pie_wrp" style="height: 250px; width: 31%;display: inline-block;text-align: center"></div>
    <iframe :src="url" class="pc" style="height: 250px;width: 34%;float: right;"/>
  </div>

  <div style="margin: 20px 0 0 0px;">
    <div v-for="(module, i) in moduleGroup" :key="i" style="background: transparent;display: inline-block;text-align: center;">
      <div v-if="module.count!==0" style="margin-top: 15px;">
        <el-badge :value=module.count class="item">
          <el-card class="box-card" style="width: 140px; min-height: 245px;margin-left: 20px;cursor: pointer;" @click=openDrillItem(module.id)>
            <template #header>
              <div class="card-header">
                <span style="font-size: 20px;"><router-link :to="'/switch/drill/detail/'+module.id" target="_blank">{{module.name}}</router-link></span>
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
      <div v-else style="margin-top: 15px;">
        <el-card class="box-card" style="width: 140px; min-height: 245px;margin-left: 20px;cursor: pointer;" @click=openDrillItem(module.id)>
          <template #header>
            <div class="card-header">
              <span style="font-size: 20px;"><router-link :to="'/switch/drill/detail/'+module.id" target="_blank">{{module.name}}</router-link></span>
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
</template>

<script>
import axios from "axios";
// 引入 echarts 核心模块，核心模块提供了 echarts 使用必须要的接口。
import * as echarts from "echarts/core";
// 引入圆形图
import { PieChart } from "echarts/charts";
// 引入提示框，标题，直角坐标系组件，组件后缀都为 Component
import {
  TitleComponent,
  TooltipComponent,
  GridComponent,
  LegendComponent
} from "echarts/components";
import { LabelLayout } from 'echarts/features';
// 引入 Canvas 渲染器，注意引入 CanvasRenderer 或者 SVGRenderer 是必须的一步
import { CanvasRenderer } from "echarts/renderers";
echarts.use([TitleComponent, TooltipComponent, LabelLayout, LegendComponent, GridComponent, PieChart, CanvasRenderer]);

export default {
  data() {
    return {
      url: 'https://element-plus.gitee.io/zh-CN/component/skeleton.html#%E6%B8%B2%E6%9F%93%E5%A4%9A%E6%9D%A1%E6%95%B0%E6%8D%AE',
      tableFailedDetails: [
        {
          id: 1,
          module: 'oracle',
          detail: 'oracle1',
          beginTime: '2021-12-26 00：00：00',
          endTime: '2021-12-26 00：03：02',
          comment: 'No. 189, Grove St, Los Angeles',
        },
      ],
      moduleGroup: {
        1: {
          id: 1,
          name: 'oracle',
          beginTime: '00:00:00',
          endTime: '00:03:02',
          duration: '00:30:50',
          count: 2,
          status: 'exception',
          percent: 66,
          content: '2/3'
        },
        2: {
          id: 2,
          name: 'mysql',
          beginTime: '00:00:00',
          endTime: '00:03:02',
          duration: '00:30:50',
          count: 0,
          status: 'success',
          percent: 20,
          content: '1/5'
        },
      },
      moduleIds: [],
      pieProgress: [
        { value: 1048, name: 'Search Engine' },
        { value: 735, name: 'Direct' }
      ]
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.getAllSwitchModules(1);
      this.getAllFailedDetails(1);
      this.getVersionProgressPie(1);
    },
    getAllFailedDetails(versionId) {
      const url = 'http://127.0.0.1:80/switch/version/failed/details/' + versionId
      axios.get(url).then(
          res => {
            this.tableFailedDetails = []
            if (res.data['details']) {
              res.data['details'].forEach((item) => {
                let detail = {}
                detail['id'] = item['Module']['Id']
                detail['module'] = item['Module']['Module']
                detail['detail'] = item['Detail']
                detail['beginTime'] = item['BeginTime']
                detail['endTime'] = item['EndTime']
                detail['comment'] = item['Comment']
                this.tableFailedDetails.push(detail)
              })
            }
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
            this.moduleGroup[moduleId]['beginTime'] = '00:00:00'
            this.moduleGroup[moduleId]['endTime'] = '00:00:00'
            this.moduleGroup[moduleId]['duration'] = '00:00:00'
            if (res.data['progress']['BeginTime']) {
              this.moduleGroup[moduleId]['beginTime'] = res.data['progress']['BeginTime'].substring(11)
            }
            if (res.data['progress']['EndTime']) {
              this.moduleGroup[moduleId]['endTime'] = res.data['progress']['EndTime'].substring(11)
            }
            if (res.data['progress']['Duration']) {
              this.moduleGroup[moduleId]['duration'] = res.data['progress']['Duration'].substring(0,8)
            }
            let finished = 0
            let total = 0
            if (res.data['progress']['ModuleStatus']) {
              res.data['progress']['ModuleStatus'].forEach((item) => {
                if (item['Status'] === 3) {
                  this.moduleGroup[moduleId]['count'] = item['Count']
                  this.moduleGroup[moduleId]['status'] = 'exception'
                  finished += item['Count']
                } else if (item['Status'] === 2) {
                  finished += item['Count']
                }
                total += item['Count']
              })
            }
            this.moduleGroup[moduleId]['content'] = finished+'/'+total
            this.moduleGroup[moduleId]['percent'] = Math.round(finished/total*100)
          },
          error => {
            console.log(error)
          }
      )
    },
    getAllSwitchModules(versionId) {
      const url = 'http://127.0.0.1:80/switch/modules/' + versionId
      axios.get(url).then(
          res => {
            res.data['modules'].forEach((item) => {
              this.moduleIds.push(item['Id'])
              let module = {}
              module['id'] = item['Id']
              module['name'] = item['Module']
              module['beginTime'] = '00:00:00'
              module['endTime'] = '00:00:00'
              module['duration'] = '00:00:00'
              module['count'] = 0
              module['status'] = 'success'
              module['percent'] = 0
              module['content'] = '0/0'
              this.moduleGroup[item['Id']] = module
            })
            this.getProgress();
          },
          error => {
            console.log(error)
          }
      )
    },
    getProgress() {
      this.moduleIds.forEach((moduleId) => {
        this.getProgressByModuleId(moduleId);
      })
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
    openDrillItem(moduleId) {
      const url = '/#/switch/drill/detail/'+moduleId
      window.open(url, '_blank')
    },
    goTo(index) {
      let url = '/#/switch/preparation'
      if (index===3) {
        url = '/#/switch/drill'
      }
      window.open(url, '_self')
    },
    getVersionProgressPie(versionId) {
      const url = 'http://127.0.0.1:80/switch/version/progress/' + versionId
      axios.get(url).then(
          res => {
            const STATUS = {0:'未开始', 1:'进行中', 2:'成功', 3:'失败'}
            this.pieProgress = []
            res.data['progress']['ModuleStatus'].forEach((item) => {
              let progress = {}
              progress['value'] = item['Count']
              progress['name'] = STATUS[item['Status']]
              this.pieProgress.push(progress)
            })
            this.getPie(this.pieProgress)
          },
          error => {
            console.log(error)
          }
      )
    },
    getPie(data) {
      echarts.init(this.$refs.pie_wrp).setOption({
        color: ['#f56c6c', '#909399', '#e6a23c', '#67c23a'],
        title: {
          text: '整体进度',
          // subtext: 'Fake Data',
          left: 'center'
        },
        tooltip: {
          trigger: 'item',
          formatter: '{b}: {c} ({d}%)'
        },
        legend: {
          top: 'bottom'
        },
        series: [
          {
            // name: 'Access From',
            type: 'pie',
            radius: '50%',
            data: data,
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            },
            label: {
              formatter: '{per|{d}%} ({count|{c}})',
              rich: {
                per: {
                  fontSize: 18,
                },
                count: {
                  color: '#999'
                },
              }
            },
          }
        ]
      })
    },
  },
}
</script>


<style scoped>
.module-content {
  font-weight: bold;
  font-size: 20px;
}

.module-title {
  color: #0d5aa7;
}
</style>
