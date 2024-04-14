<template>
  <v-chart class="w-full h-[100px]" :option="option" autoresize />
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
  series: [
    {
      name: "Skills",
      type: "pie",
      radius: "55%",
      center: ["50%", "60%"],
      data: [
        { value: 335, name: "brute force" },
        { value: 310, name: "data structures" },
        { value: 234, name: "dp" },
        { value: 135, name: "math" },
        { value: 1548, name: "number theory" }
      ],
      emphasis: {
        itemStyle: {
          shadowBlur: 10,
          shadowOffsetX: 0,
          shadowColor: "rgba(0, 0, 0, 0.5)"
        }
      }
    }
  ]
});
</script>
