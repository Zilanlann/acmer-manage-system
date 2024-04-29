<template>
  <!-- 操作按钮组 -->
  <div class="table-header">
    <!-- <slot name="refreshPrepend"></slot> -->
    <el-tooltip content="刷新" placement="top">
      <el-button
        v-blur
        color="#40485b"
        class="table-header-operate"
        type="info"
      >
        <component :is="useRenderIcon('ep:refresh')" />
      </el-button>
    </el-tooltip>
    <slot name="refreshAppend"></slot>
    <el-tooltip content="添加比赛" placement="top">
      <el-button v-blur class="table-header-operate" type="primary">
        <component :is="useRenderIcon('ep:circle-plus')" />
        <span class="table-header-operate-text">添加比赛</span>
      </el-button>
    </el-tooltip>
    <el-popconfirm
      confirm-button-text="删除"
      cancel-button-text="取消"
      confirmButtonType="danger"
      title="确定删除选中比赛？"
    >
      <template #reference>
        <div class="mlr-12">
          <el-tooltip content="删除选中比赛" placement="top">
            <el-button v-blur class="table-header-operate" type="danger">
              <component :is="useRenderIcon('ep:delete')" />
              <span class="table-header-operate-text">删除</span>
            </el-button>
          </el-tooltip>
        </div>
      </template>
    </el-popconfirm>

    <!-- slot -->
    <slot></slot>

    <!-- 右侧搜索框和工具按钮 -->
    <div class="table-search">
      <slot name="quickSearchPrepend"></slot>
      <el-input
        class="xs-hidden quick-search"
        :placeholder="'请输入比赛名称'"
        clearable
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { computed } from "vue";
import { ElTable } from "element-plus";

// const enableBatchOpt = computed(() =>
//   ElTable.table.selection!.length > 0 ? true : false
// );

// const onAction = (event: string, data: any = {}) => {
//   ElTable.onTableHeaderAction(event, data);
// };

// const onSearchInput = () => {
//   ElTable.onTableHeaderAction("quick-search", {
//     keyword: ElTable.table.filter!.quickSearch
//   });
// };

// const onChangeShowColumn = (
//   value: string | number | boolean,
//   field: string
// ) => {
//   ElTable.onTableHeaderAction("change-show-column", {
//     field: field,
//     value: value
//   });
// };
</script>

<style scoped lang="scss">
.table-header {
  position: relative;
  overflow-x: auto;
  box-sizing: border-box;
  display: flex;
  align-items: center;
  width: 100%;
  max-width: 100%;
  background-color: var(--ba-bg-color-overlay);
  border: 1px solid var(--ba-border-color);
  border-bottom: none;
  padding: 13px 15px;
  font-size: 14px;
  .table-header-operate-text {
    margin-left: 6px;
  }
}

.mlr-12 {
  margin-left: 12px;
}
.mlr-12 + .el-button {
  margin-left: 12px;
}
.table-search {
  display: flex;
  margin-left: auto;
  .quick-search {
    width: auto;
  }
}
.table-search-button-group {
  display: flex;
  margin-left: 12px;
  border: 1px solid var(--el-border-color);
  border-radius: var(--el-border-radius-base);
  overflow: hidden;
  button:focus,
  button:active {
    background-color: var(--ba-bg-color-overlay);
  }
  button:hover {
    background-color: var(--el-color-info-light-7);
  }
  .table-search-button-item {
    height: 30px;
    border: none;
    border-radius: 0;
  }
  .el-button + .el-button {
    margin: 0;
  }
  .right-border {
    border-right: 1px solid var(--el-border-color);
  }
}

html.dark {
  .table-search-button-group {
    button:focus,
    button:active {
      background-color: var(--el-color-info-dark-2);
    }
    button:hover {
      background-color: var(--el-color-info-light-7);
    }
    button {
      background-color: var(--ba-bg-color-overlay);
      el-icon {
        color: white !important;
      }
    }
  }
}
</style>
