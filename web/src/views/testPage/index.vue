<script setup lang="ts">
import { message } from "@/utils/message";
import { useTestStoreHook } from "@/store/modules/test";
import { ref, reactive, toRaw, onMounted, onBeforeUnmount } from "vue";
import { useDataThemeChange } from "@/layout/hooks/useDataThemeChange";
defineOptions({
  name: "TestPage"
});
async function onTest() {
  try {
    const result = await useTestStoreHook().testApi();
    message(useTestStoreHook().test, { type: "success" });
  } catch (error) {
    console.error("Test API Failed:", error);
  }
}
</script>

<template>
  <div>
    <el-button type="primary" @click="onTest">测试</el-button>
  </div>
</template>

<style lang="scss" scoped></style>
