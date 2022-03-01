<template>
    <div id="pie" ref="pie_wrp" style="height: 250px; width: 31%;display: inline-block;text-align: center"></div>
</template>

<script>
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
      moduleIds: [],
    }
  },
  mounted() {
    this.init()
  },
  methods: {
    init() {
      this.getPie();
    },
    getPie() {
      echarts.init(this.$refs.pie_wrp).setOption({
        title: {
          text: '整体进度',
          // subtext: 'Fake Data',
          left: 'center'
        },
        tooltip: {
          trigger: 'item'
        },
        legend: {
          top: 'bottom'
        },
        series: [
          {
            name: 'Access From',
            type: 'pie',
            radius: '50%',
            data: [
              { value: 1048, name: 'Search Engine' },
              { value: 735, name: 'Direct' },
              { value: 580, name: 'Email' },
              { value: 484, name: 'Union Ads' },
              { value: 300, name: 'Video Ads' }
            ],
            emphasis: {
              itemStyle: {
                shadowBlur: 10,
                shadowOffsetX: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          }
        ]
      })

    }
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
