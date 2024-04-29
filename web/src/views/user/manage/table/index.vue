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
    <pure-table
      border
      row-key="id"
      alignWhole="center"
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
