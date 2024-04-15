<template>
  <div>
    <v-chart class="w-[800px] h-[800px]" :option="option" autoresize />
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
import VChart, { THEME_KEY } from "vue-echarts";
import { ref, provide } from "vue";

defineOptions({
  name: "PieChartNormal"
});

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent
]);

// provide(THEME_KEY, "dark");

const option = ref({
  tooltip: {
    trigger: "item",
    formatter: "{b} : {c} ({d}%)",
    confine: true
  },
  legend: {
    top: "5%",
    left: "center"
  },
  series: [
    {
      name: "Skills",
      type: "pie",
      radius: ["40%", "70%"],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 10,
        borderColor: "#fff",
        borderWidth: 2
      },
      label: {
        show: false,
        position: "center"
      },
      emphasis: {
        label: {
          show: true,
          fontSize: 40,
          fontWeight: "bold"
        }
      },
      labelLine: {
        show: false
      },
      data: [
        { value: 335, name: "brute force" },
        { value: 310, name: "data structures" },
        { value: 234, name: "dp" },
        { value: 135, name: "math" },
        { value: 1548, name: "number theory" }
      ]
    }
  ]
});
</script>
