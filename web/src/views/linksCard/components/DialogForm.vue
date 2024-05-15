<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { message } from "@/utils/message";
import { FormInstance } from "element-plus";
import { createSite, getSiteTypeList } from "@/api/favSite";

const SELECT_OPTIONS = ref<Array<{ label: string; value: number }>>([]);

const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  },
  data: {
    type: Object,
    default: () => {
      return {
        name: "",
        host: "",
        siteTypeID: 0,
        desc: ""
      };
    }
  }
});

const ruleFormRef = ref<FormInstance>();

const formVisible = ref(false);
const formData = ref(props.data);
const textareaValue = ref("");

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async valid => {
    if (valid) {
      try {
        await createSite(formData.value);
        message("提交成功", { type: "success" });
        formVisible.value = false;
        resetForm(formEl);
      } catch (error) {
        message("提交失败", { type: "error" });
      }
    }
  });
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};

const closeDialog = () => {
  formVisible.value = false;
  resetForm(ruleFormRef.value);
};

const emit = defineEmits(["update:visible"]);
watch(
  () => formVisible.value,
  val => {
    emit("update:visible", val);
  }
);

watch(
  () => props.visible,
  val => {
    formVisible.value = val;
  }
);

watch(
  () => props.data,
  val => {
    formData.value = val;
  }
);

const rules = {
  name: [{ required: true, message: "请输入网站名称", trigger: "blur" }],
  host: [{ required: true, message: "请输入网站链接", trigger: "blur" }]
};

const fetchSiteTypeList = async () => {
  try {
    const response = await getSiteTypeList();
    const list = response.data.list; // Assuming 'list' contains the array of site types
    SELECT_OPTIONS.value = list.map((item: { name: string; ID: number }) => ({
      label: item.name,
      value: item.ID
    }));
  } catch (error) {
    message("获取网站类型列表失败", { type: "error" });
  }
};

onMounted(() => {
  fetchSiteTypeList();
});
</script>

<template>
  <el-dialog
    v-model="formVisible"
    title="添加网站"
    :width="680"
    draggable
    :before-close="closeDialog"
  >
    <!-- 表单内容 -->
    <el-form
      ref="ruleFormRef"
      :model="formData"
      :rules="rules"
      label-width="100px"
    >
      <el-form-item label="网站名称" prop="name">
        <el-input
          v-model="formData.name"
          :style="{ width: '480px' }"
          placeholder="请输入网站名称"
        />
      </el-form-item>
      <el-form-item label="网站链接" prop="host">
        <el-input
          v-model="formData.host"
          :style="{ width: '480px' }"
          placeholder="请输入网站链接"
        />
      </el-form-item>
      <el-form-item label="网站类型" prop="type">
        <el-select
          v-model="formData.siteTypeID"
          clearable
          :style="{ width: '480px' }"
        >
          <el-option
            v-for="(item, index) in SELECT_OPTIONS"
            :key="index"
            :value="item.value"
            :label="item.label"
          >
            {{ item.label }}
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="网站描述" prop="desc">
        <el-input
          v-model="formData.desc"
          type="textarea"
          :style="{ width: '480px' }"
          placeholder="请输入内容"
        />
      </el-form-item>
    </el-form>
    <template #footer>
      <el-button @click="closeDialog">取消</el-button>
      <el-button type="primary" @click="submitForm(ruleFormRef)">
        确定
      </el-button>
    </template>
  </el-dialog>
</template>
