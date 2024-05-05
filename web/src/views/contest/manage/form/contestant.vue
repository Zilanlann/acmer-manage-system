<script setup lang="ts">
import { ref } from "vue";
import ReCol from "@/components/ReCol";
import { formRules } from "../utils/rule";
import { ContestantFormProps, FormProps } from "../utils/types";
import { usePublicHooks } from "../../hooks";

const props = withDefaults(defineProps<ContestantFormProps>(), {
  formInline: () => ({
    title: "新增",
    acmerOptions: [],
    userID: 0
  })
});

const ruleFormRef = ref();
const { switchStyle } = usePublicHooks();
const newFormInline = ref(props.formInline);

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-row :gutter="30">
      <re-col>
        <el-form-item label="选手" prop="id">
          <el-select
            v-model="newFormInline.userID"
            placeholder="请选择"
            class="w-full"
            clearable
          >
            <el-option
              v-for="(item, index) in newFormInline.acmerOptions"
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
  </el-form>
</template>
