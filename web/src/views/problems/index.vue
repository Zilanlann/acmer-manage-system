<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useProblem } from "./utils/hook";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";

import Refresh from "@iconify-icons/ep/refresh";
import { getAllProblemTagList } from "@/api/oj";

defineOptions({
  name: "ProblemsList"
});

const treeRef = ref();
const formRef = ref();
const tableRef = ref();
const tags = ref([]); // Define tags as a ref

const {
  form,
  loading,
  columns,
  dataList,
  pagination,
  buttonClass,
  deviceDetection,
  onSearch,
  resetForm,
  handleSizeChange,
  handleCurrentChange,
  handleSelectionChange
} = useProblem(tableRef, treeRef);

onMounted(async () => {
  try {
    const response = await getAllProblemTagList();
    tags.value = response.data.list.map(tag => ({
      value: tag.ID,
      label: tag.name
    }));
  } catch (error) {
    console.error("Failed to fetch tags:", error);
  }
});
</script>

<template>
  <div :class="['flex', 'justify-between', deviceDetection() && 'flex-wrap']">
    <div :class="['w-full', 'mt-2']">
      <el-form
        ref="formRef"
        :inline="true"
        :model="form"
        class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px] overflow-auto"
      >
        <el-form-item label="题目名称：" prop="contestName">
          <el-input
            v-model="form.name"
            placeholder="请输入题目名称"
            clearable
            class="!w-[180px]"
          />
        </el-form-item>
        <el-form-item label="Rating范围：" prop="rating">
          <el-slider
            v-model="form.rating"
            :min="0"
            :max="3000"
            range
            class="!w-[360px]"
          />
        </el-form-item>
        <el-form-item label="题目类型：" prop="problemType">
          <el-select
            v-model="form.tags"
            placeholder="请选择题目类型"
            multiple
            clearable
            class="!w-[180px]"
          >
            <el-option
              v-for="item in tags"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="提交时间：" prop="time">
          <el-date-picker
            v-model="form.time"
            type="daterange"
            range-separator="到"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            :icon="useRenderIcon('ep:search')"
            :loading="loading"
            @click="onSearch"
          >
            搜索
          </el-button>
          <el-button :icon="useRenderIcon(Refresh)" @click="resetForm(formRef)">
            重置
          </el-button>
        </el-form-item>
      </el-form>

      <PureTableBar title="刷题统计" :columns="columns" @refresh="onSearch">
        <template v-slot="{ size, dynamicColumns }">
          <pure-table
            ref="tableRef"
            row-key="id"
            adaptive
            :adaptiveConfig="{ offsetBottom: 108 }"
            align-whole="center"
            table-layout="auto"
            :loading="loading"
            :size="size"
            :data="dataList"
            :columns="dynamicColumns"
            :pagination="pagination"
            :paginationSmall="size === 'small' ? true : false"
            :header-cell-style="{
              background: 'var(--el-fill-color-light)',
              color: 'var(--el-text-color-primary)'
            }"
            @selection-change="handleSelectionChange"
            @page-size-change="handleSizeChange"
            @page-current-change="handleCurrentChange"
          >
          </pure-table>
        </template>
      </PureTableBar>
    </div>
  </div>
</template>

<style scoped lang="scss">
:deep(.el-dropdown-menu__item i) {
  margin: 0;
}

:deep(.el-button:focus-visible) {
  outline: none;
}

.main-content {
  margin: 24px 24px 0 !important;
}

.search-form {
  :deep(.el-form-item) {
    margin-bottom: 12px;
  }
}
</style>
