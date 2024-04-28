<script setup lang="ts">
import { message } from "@/utils/message";
import { useTestStoreHook } from "@/store/modules/test";
import PieChartNormal from "@/views/charts/pieCharts/normal/index.vue";
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
    <el-form
      ref="formRef"
      :inline="true"
      :model="form"
      class="search-form bg-bg_color w-[99/100] pl-8 pt-[12px] overflow-auto"
    >
      <el-form-item label="用户名称：" prop="username">
        <el-input
          v-model="form.username"
          placeholder="请输入用户名称"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="手机号码：" prop="phone">
        <el-input
          v-model="form.phone"
          placeholder="请输入手机号码"
          clearable
          class="!w-[180px]"
        />
      </el-form-item>
      <el-form-item label="状态：" prop="status">
        <el-select
          v-model="form.status"
          placeholder="请选择"
          clearable
          class="!w-[180px]"
        >
          <el-option label="已开启" value="1" />
          <el-option label="已关闭" value="0" />
        </el-select>
      </el-form-item>
      <el-button type="primary"> 添加比赛 </el-button>
    </el-form>
  </div>
</template>

<style scoped lang="scss"></style>
