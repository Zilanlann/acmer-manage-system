<script setup lang="ts">
import { message } from "@/utils/message";
import { useTestStoreHook } from "@/store/modules/test";
import PieChartNormal from "@/views/charts/pieCharts/normal/index.vue";
defineOptions({
  name: "TestPage"
});
async function onTest() {
  try {
    const result = await useTestStoreHook().testApi();
    message(useTestStoreHook().test, { type: "success" });
  } catch (error) {
    console.error("Test API Failed:", error);
  }
}
</script>

<template>
  <div>
    <el-button type="primary" @click="onTest">测试</el-button>
    <PieChartNormal />
  </div>
</template>

<style lang="scss" scoped></style>
