<template>
  <div>
    <v-chart class="w-full h-[100px]" :option="option" autoresize />
  </div>
</template>

<script setup>
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart } from "echarts/charts";
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent
} from "echarts/components";
import VChart from "vue-echarts";
import { ref, watch, defineProps } from "vue";

defineOptions({
  name: "PieChartSmall"
});

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent
]);

const props = defineProps({
  userData: Array
});

const option = ref({});

watch(
  () => props.userData,
  newData => {
    option.value = {
      tooltip: {
        trigger: "item",
        formatter: "{b} : {c} ({d}%)",
        confine: true
      },
      series: [
        {
          name: "Skills",
          type: "pie",
          radius: "55%",
          center: ["50%", "60%"],
          data: newData,
          emphasis: {
            itemStyle: {
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowColor: "rgba(0, 0, 0, 0.5)"
            }
          }
        }
      ]
    };
  },
  { immediate: true }
);
</script>
