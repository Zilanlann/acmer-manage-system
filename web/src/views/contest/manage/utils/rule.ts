import { reactive } from "vue";
import type { FormRules } from "element-plus";
import { isPhone, isEmail } from "@pureadmin/utils";

/** 自定义表单规则校验 */
export const formRules = reactive(<FormRules>{
  name: [{ required: true, message: "比赛名称为必填项", trigger: "blur" }]
});

export const teamFormRules = reactive(<FormRules>{
  zhName: [{ required: true, message: "队伍中文名为必填项", trigger: "blur" }]
});
