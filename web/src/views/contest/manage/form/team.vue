<script setup lang="ts">
import { ref } from "vue";
import ReCol from "@/components/ReCol";
import { teamFormRules } from "../utils/rule";
import { TeamFormProps } from "../utils/types";
import { usePublicHooks } from "../../hooks";

const props = withDefaults(defineProps<TeamFormProps>(), {
  formInline: () => ({
    title: "新增",
    zhName: "",
    enName: "",
    desc: "",
    teacherOptions: [],
    coachID: 0
  })
});

const ruleFormRef = ref();
const { switchStyle } = usePublicHooks();
const newFormInline = ref(props.formInline);

function getRef() {
  console.log(ruleFormRef.value);
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="teamFormRules"
    label-width="82px"
  >
    <el-row :gutter="30">
      <re-col :value="12" :xs="24" :sm="24">
        <el-form-item label="队伍中文名" prop="name">
          <el-input
            v-model="newFormInline.zhName"
            clearable
            placeholder="请输入队伍中文名"
          />
        </el-form-item>
      </re-col>
      <re-col :value="12" :xs="24" :sm="24">
        <el-form-item label="队伍英文名" prop="name">
          <el-input
            v-model="newFormInline.enName"
            clearable
            placeholder="请输入队伍英文名"
          />
        </el-form-item>
      </re-col>
      <re-col>
        <el-form-item label="教练" prop="id">
          <el-select
            v-model="newFormInline.coachID"
            placeholder="请选择"
            class="w-full"
            clearable
          >
            <el-option
              v-for="(item, index) in newFormInline.teacherOptions"
              :key="index"
              :value="item.ID"
              :label="item.realname"
            >
              {{ item.realname }}
            </el-option>
          </el-select>
        </el-form-item>
      </re-col>
    </el-row>
    <el-row :gutter="30">
      <re-col :value="24" :xs="24" :sm="24">
        <el-form-item label="队伍备注" prop="desc">
          <el-input
            v-model="newFormInline.desc"
            :autosize="{ minRows: 2, maxRows: 4 }"
            type="textarea"
            placeholder="请输入队伍备注"
          />
        </el-form-item>
      </re-col>
    </el-row>
  </el-form>
</template>
