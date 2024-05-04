<script setup lang="ts">
import { useDataStoreHook } from "@/store/modules/data";
import { useColumns } from "./columns";
import smallPie from "@/views/charts/pieCharts/small/index.vue";
import { ref, onMounted, onUnmounted } from "vue";

defineOptions({
  name: "ScoreTable"
});

const {
  loading,
  columns,
  dataList,
  select,
  hidePie,
  tableSize,
  pagination,
  loadingConfig,
  onCurrentChange
} = useColumns();

const height = ref(document.body.clientHeight - 255);

const updateHeight = () => {
  height.value = document.body.clientHeight - 255;
};

onMounted(() => {
  window.addEventListener("resize", updateHeight);
});

onUnmounted(() => {
  window.removeEventListener("resize", updateHeight);
});
</script>

<template>
  <div style="width: 100%; height: 100%">
    <el-space class="float-right mb-4">
      <!-- <p class="text-sm">多选：</p>
      <el-radio-group v-model="select" size="small">
        <el-radio-button value="yes">是</el-radio-button>
        <el-radio-button value="no">否</el-radio-button>
      </el-radio-group> -->
      <p class="text-sm">是否隐藏饼图：</p>
      <el-switch v-model="hidePie" />
      <!-- <el-divider direction="vertical" /> -->
    </el-space>

    <pure-table
      border
      adaptive
      :adaptiveConfig="{ offsetBottom: 108 }"
      align-whole="center"
      table-layout="auto"
      row-key="id"
      showOverflowTooltip
      :size="tableSize as any"
      :loading="loading"
      :loading-config="loadingConfig"
      :height="height"
      :data="
        dataList.slice(
          (pagination.currentPage - 1) * pagination.pageSize,
          pagination.currentPage * pagination.pageSize
        )
      "
      :columns="columns"
      :pagination="pagination"
      @page-current-change="onCurrentChange"
    >
      <template #echart="{ index }">
        <smallPie />
      </template>
    </pure-table>
  </div>
</template>
