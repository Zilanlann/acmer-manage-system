<script setup lang="ts">
import { ref, reactive } from "vue";
import Motion from "../utils/motion";
import { message } from "@/utils/message";
import { registerRules } from "../utils/rule";
import type { FormInstance } from "element-plus";
// import { useVerifyCode } from "../utils/verifyCode";
import { useUserStoreHook } from "@/store/modules/user";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import Lock from "@iconify-icons/ri/lock-fill";
import Iphone from "@iconify-icons/ep/iphone";
import User from "@iconify-icons/ri/user-3-fill";
import { useVerifyCode } from "../utils/verifyCode";

const checked = ref(true);
const loading = ref(false);
const ruleForm = reactive({
  username: "",
  email: "",
  verifyCode: "",
  password: "",
  repeatPassword: "",
  realname: "",
  cfHandle: "",
  atcHandle: ""
});
const ruleFormRef = ref<FormInstance>();
const { isDisabled, text } = useVerifyCode();
const repeatPasswordRule = [
  {
    validator: (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入确认密码"));
      } else if (ruleForm.password !== value) {
        callback(new Error("两次密码不一致!"));
      } else {
        callback();
      }
    },
    trigger: "blur"
  }
];

const onUpdate = async (formEl: FormInstance | undefined) => {
  loading.value = true;
  if (!formEl) return;
  await formEl.validate((valid, fields) => {
    if (valid) {
      if (checked.value) {
        // 模拟请求，需根据实际开发进行修改
        setTimeout(() => {
          message("注册成功", {
            type: "success"
          });
          loading.value = false;
        }, 2000);
      } else {
        loading.value = false;
        message("请勾选隐私政策", {
          type: "warning"
        });
      }
    } else {
      loading.value = false;
      return fields;
    }
  });
};

function onBack() {
  // useVerifyCode().end();
  useUserStoreHook().SET_CURRENTPAGE(0);
}
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="ruleForm"
    :rules="registerRules"
    size="large"
  >
    <Motion>
      <el-form-item
        :rules="[
          {
            required: true,
            message: '请输入用户名',
            trigger: 'blur'
          }
        ]"
        prop="username"
      >
        <el-input
          v-model="ruleForm.username"
          clearable
          placeholder="用户名"
          :prefix-icon="useRenderIcon(User)"
        />
      </el-form-item>
    </Motion>

    <Motion :delay="100">
      <el-form-item
        :rules="[
          {
            required: true,
            message: '请输入你的真实姓名',
            trigger: 'blur'
          }
        ]"
        prop="realname"
      >
        <el-input
          v-model="ruleForm.realname"
          clearable
          placeholder="真实姓名"
          :prefix-icon="useRenderIcon(User)"
        />
      </el-form-item>
    </Motion>

    <Motion :delay="100">
      <el-form-item
        :rules="[
          {
            required: true,
            message: '请输入你的邮箱',
            trigger: 'blur'
          }
        ]"
        prop="email"
      >
        <el-input
          v-model="ruleForm.email"
          clearable
          placeholder="邮箱"
          :prefix-icon="useRenderIcon('tdesign:mail')"
        />
      </el-form-item>
    </Motion>

    <Motion :delay="150">
      <el-form-item prop="verifyCode">
        <div class="w-full flex justify-between">
          <el-input
            v-model="ruleForm.verifyCode"
            clearable
            placeholder="邮箱验证码"
            :prefix-icon="useRenderIcon('ri:shield-keyhole-line')"
          />
          <el-button
            :disabled="isDisabled"
            class="ml-2"
            @click="useVerifyCode().start(ruleFormRef, 'phone')"
          >
            {{ text.length > 0 ? text + "秒后重新获取" : "获取验证码" }}
          </el-button>
        </div>
      </el-form-item>
    </Motion>

    <Motion :delay="200">
      <el-row :gutter="20">
        <el-col :span="12">
          <el-form-item prop="cf">
            <el-input
              v-model="ruleForm.cfHandle"
              clearable
              placeholder="Codeforces 用户名"
              :prefix-icon="useRenderIcon('tdesign:code')"
            />
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item prop="atc">
            <el-input
              v-model="ruleForm.atcHandle"
              clearable
              placeholder="Atcoder 用户名"
              :prefix-icon="useRenderIcon('tdesign:code')"
            />
          </el-form-item>
        </el-col>
      </el-row>
    </Motion>

    <Motion :delay="200">
      <el-form-item prop="password">
        <el-input
          v-model="ruleForm.password"
          clearable
          show-password
          placeholder="密码"
          :prefix-icon="useRenderIcon(Lock)"
        />
      </el-form-item>
    </Motion>

    <Motion :delay="250">
      <el-form-item :rules="repeatPasswordRule" prop="repeatPassword">
        <el-input
          v-model="ruleForm.repeatPassword"
          clearable
          show-password
          placeholder="确认密码"
          :prefix-icon="useRenderIcon(Lock)"
        />
      </el-form-item>
    </Motion>

    <Motion :delay="350">
      <el-form-item>
        <el-button
          class="w-full"
          size="default"
          type="primary"
          :loading="loading"
          @click="onUpdate(ruleFormRef)"
        >
          确定
        </el-button>
      </el-form-item>
    </Motion>

    <Motion :delay="400">
      <el-form-item>
        <el-button class="w-full" size="default" @click="onBack">
          返回
        </el-button>
      </el-form-item>
    </Motion>
  </el-form>
</template>