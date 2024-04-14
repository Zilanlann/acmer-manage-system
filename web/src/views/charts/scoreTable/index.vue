<script setup lang="ts">
import { useColumns } from "./columns";
import smallPie from "@/views/charts/pieCharts/small/index.vue";

const {
  loading,
  columns,
  dataList,
  select,
  hidePie,
  tableSize,
  pagination,
  loadingConfig,
  onChange,
  onSizeChange,
  onCurrentChange
} = useColumns();
</script>

<template>
  <div style="width: 100%">
    <el-space class="float-right mb-4">
      <!-- <p class="text-sm">多选：</p>
      <el-radio-group v-model="select" size="small">
        <el-radio-button value="yes">是</el-radio-button>
        <el-radio-button value="no">否</el-radio-button>
      </el-radio-group> -->
      <p class="text-sm">是否隐藏饼图：</p>
      <!-- <el-radio-group v-model="hideVal" size="small">
        <el-radio-button value="nohide">不隐藏</el-radio-button>
        <el-radio-button value="hidePie">隐藏饼图</el-radio-button>
      </el-radio-group> -->
      <el-switch v-model="hidePie" />
      <!-- <el-divider direction="vertical" /> -->
    </el-space>

    <pure-table
      border
      row-key="id"
      alignWhole="center"
      showOverflowTooltip
      :size="tableSize as any"
      :loading="loading"
      :loading-config="loadingConfig"
      :height="tableSize === 'small' ? 352 : 700"
      :data="
        dataList.slice(
          (pagination.currentPage - 1) * pagination.pageSize,
          pagination.currentPage * pagination.pageSize
        )
      "
      :columns="columns"
      :pagination="pagination"
      @page-size-change="onSizeChange"
      @page-current-change="onCurrentChange"
    >
      <template #echart="{ index }">
        <smallPie />
      </template>
    </pure-table>
  </div>
</template>
