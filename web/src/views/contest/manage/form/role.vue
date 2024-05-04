<script setup lang="ts">
import { ref } from "vue";
import ReCol from "@/components/ReCol";
import { RoleFormProps } from "../utils/types";
import { string } from "vue-types";

const props = withDefaults(defineProps<RoleFormProps>(), {
  formInline: () => ({
    realname: "",
    roleOptions: [],
    role: ""
  })
});

const newFormInline = ref(props.formInline);
</script>

<template>
  <el-form :model="newFormInline">
    <el-row :gutter="30">
      <re-col>
        <el-form-item label="用户名称" prop="realname">
          <el-input disabled v-model="newFormInline.realname" />
        </el-form-item>
      </re-col>
      <re-col>
        <el-form-item label="角色列表" prop="id">
          <el-select
            v-model="newFormInline.role"
            placeholder="请选择"
            class="w-full"
            clearable
          >
            <el-option
              v-for="(item, index) in newFormInline.roleOptions"
              :key="index"
              :value="item.id"
              :label="item.name"
            >
              {{ item.name }}
            </el-option>
          </el-select>
        </el-form-item>
      </re-col>
    </el-row>
  </el-form>
</template>
