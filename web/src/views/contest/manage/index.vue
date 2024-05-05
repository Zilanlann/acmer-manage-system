<script setup lang="ts">
import { ref } from "vue";
import { useContest } from "./utils/hook";
import { PureTableBar } from "@/components/RePureTableBar";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";

import Password from "@iconify-icons/ri/lock-password-line";
import More from "@iconify-icons/ep/more-filled";
import Delete from "@iconify-icons/ep/delete";
import EditPen from "@iconify-icons/ep/edit-pen";
import Refresh from "@iconify-icons/ep/refresh";
import AddFill from "@iconify-icons/ri/add-circle-line";

defineOptions({
  name: "ContestManage"
});

const treeRef = ref();
const formRef = ref();
const tableRef = ref();

const {
  form,
  loading,
  columns,
  teamColumns,
  contestantColumns,
  dataList,
  selectedNum,
  pagination,
  buttonClass,
  deviceDetection,
  onSearch,
  resetForm,
  onbatchDel,
  openDialog,
  openTeamDialog,
  openContestantDialog,
  handleDeleteContest,
  handleDeleteTeam,
  handleDeleteContestant,
  handleSizeChange,
  onSelectionCancel,
  handleCurrentChange,
  handleSelectionChange
} = useContest(tableRef, treeRef);
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
        <el-form-item label="比赛名称：" prop="contestName">
          <el-input
            v-model="form.contestName"
            placeholder="请输入比赛名称"
            clearable
            class="!w-[180px]"
          />
        </el-form-item>
        <el-form-item label="时间：" prop="time">
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

      <PureTableBar title="比赛管理" :columns="columns" @refresh="onSearch">
        <template #buttons>
          <el-button
            type="primary"
            :icon="useRenderIcon(AddFill)"
            @click="openDialog('新增')"
          >
            新增比赛
          </el-button>
        </template>
        <template v-slot="{ size, dynamicColumns }">
          <div
            v-if="selectedNum > 0"
            v-motion-fade
            class="bg-[var(--el-fill-color-light)] w-full h-[46px] mb-2 pl-4 flex items-center"
          >
            <div class="flex-auto">
              <span
                style="font-size: var(--el-font-size-base)"
                class="text-[rgba(42,46,54,0.5)] dark:text-[rgba(220,220,242,0.5)]"
              >
                已选 {{ selectedNum }} 项
              </span>
              <el-button type="primary" text @click="onSelectionCancel">
                取消选择
              </el-button>
            </div>
            <el-popconfirm title="是否确认删除?" @confirm="onbatchDel">
              <template #reference>
                <el-button type="danger" text class="mr-1">
                  批量删除
                </el-button>
              </template>
            </el-popconfirm>
          </div>
          <!-- 比赛表格 -->
          <pure-table
            ref="tableRef"
            row-key="ID"
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
            <template #expand="{ row }">
              <div>
                <div style="margin-left: 20px">
                  <p>比赛详情以及备忘信息：</p>
                  <p>{{ row.desc }}</p>
                </div>
                <!-- 队伍表格 -->
                <PureTable
                  row-key="ID"
                  align-whole="center"
                  table-layout="auto"
                  :columns="teamColumns"
                  :data="row.teams"
                >
                  <template #expand="{ row }">
                    <div>
                      <div style="margin-left: 20px">
                        <p>队伍详情以及备忘信息：</p>
                        <p>{{ row.desc }}</p>
                      </div>
                      <!-- 队员表格 -->
                      <PureTable
                        row-key="ID"
                        align-whole="center"
                        table-layout="auto"
                        :columns="contestantColumns"
                        :data="row.contestants"
                      >
                        <!-- 自定义操作 -->
                        <template #operation="{ row }">
                          <el-button
                            class="reset-margin"
                            link
                            type="primary"
                            :size="size"
                            :icon="useRenderIcon(EditPen)"
                            @click="openContestantDialog('修改', row)"
                          >
                            修改
                          </el-button>
                          <el-popconfirm
                            :title="`是否确认删除队员编号为${row.ID}的这名队员`"
                            @confirm="handleDeleteContestant(row)"
                          >
                            <template #reference>
                              <el-button
                                class="reset-margin"
                                link
                                type="primary"
                                :size="size"
                                :icon="useRenderIcon(Delete)"
                              >
                                删除
                              </el-button>
                            </template>
                          </el-popconfirm>
                        </template>
                      </PureTable>
                    </div>
                  </template>
                  <template #operation="{ row }">
                    <el-button
                      class="reset-margin"
                      link
                      type="primary"
                      :size="size"
                      :icon="useRenderIcon(EditPen)"
                      @click="openTeamDialog('修改', row)"
                    >
                      修改
                    </el-button>
                    <el-button
                      class="reset-margin"
                      link
                      type="primary"
                      :size="size"
                      :icon="useRenderIcon(AddFill)"
                      @click="openContestantDialog('新增', row)"
                    >
                      添加队员
                    </el-button>
                    <el-popconfirm
                      :title="`是否确认删除队伍编号为${row.ID}的这条数据`"
                      @confirm="handleDeleteTeam(row)"
                    >
                      <template #reference>
                        <el-button
                          class="reset-margin"
                          link
                          type="primary"
                          :size="size"
                          :icon="useRenderIcon(Delete)"
                        >
                          删除
                        </el-button>
                      </template>
                    </el-popconfirm>
                  </template>
                </PureTable>
              </div>
            </template>
            <template #operation="{ row }">
              <el-button
                class="reset-margin"
                link
                type="primary"
                :size="size"
                :icon="useRenderIcon(EditPen)"
                @click="openDialog('修改', row)"
              >
                修改
              </el-button>
              <el-button
                class="reset-margin"
                link
                type="primary"
                :size="size"
                :icon="useRenderIcon(AddFill)"
                @click="openTeamDialog('新增', row)"
              >
                添加队伍
              </el-button>
              <el-popconfirm
                :title="`是否确认删除比赛编号为${row.ID}的这条数据`"
                @confirm="handleDeleteContest(row)"
              >
                <template #reference>
                  <el-button
                    class="reset-margin"
                    link
                    type="primary"
                    :size="size"
                    :icon="useRenderIcon(Delete)"
                  >
                    删除
                  </el-button>
                </template>
              </el-popconfirm>
              <el-dropdown>
                <el-button
                  class="ml-3 mt-[2px]"
                  link
                  type="primary"
                  :size="size"
                  :icon="useRenderIcon(More)"
                />
              </el-dropdown>
            </template>
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
