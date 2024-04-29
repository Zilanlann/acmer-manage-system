<script setup lang="ts">
import { message } from "@/utils/message";
import { useTestStoreHook } from "@/store/modules/test";
import PieChartNormal from "@/views/charts/pieCharts/normal/index.vue";
import Header from "@/views/user/manage/header/index.vue";
import ContestTable from "@/views/user/manage/table/index.vue";
import { reactive } from "vue";
defineOptions({
  name: "ContestPage"
});

const form = reactive({
  username: "",
  status: "",
  phone: ""
});
</script>

<template>
  <div>
    <Header />
    <ContestTable />
  </div>
</template>

<style scoped lang="scss"></style>
