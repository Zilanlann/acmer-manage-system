<script setup lang="ts">
defineOptions({
  name: "Calendar"
});
import { reactive, ref } from "vue";
import type { CalendarDateType, CalendarInstance } from "element-plus";

const calendar = ref<CalendarInstance>();
const selectDate = (val: CalendarDateType) => {
  if (!calendar.value) return;
  calendar.value.selectDate(val);
};

const contests = reactive([
  {
    name: "Codeforces Round (Div. 4)",
    oj: "codeforces",
    link: "https://codeforces.com/",
    date: "2024-04-23"
  },
  {
    name: "Codeforces Round (Div. 3)",
    link: "https://codeforces.com/",
    date: "2024-04-23"
  },
  {
    name: "Codeforces Round (Div. 3)",
    link: "https://codeforces.com/",
    date: "2024-04-23"
  }
]);

const openUrl = (url: string) => {
  window.open(url, "_blank");
};
</script>

<template>
  <el-calendar ref="calendar">
    <template #header="{ date }">
      <span>{{ date }}</span>
      <b><span>各大OJ竞赛日历</span></b>
      <el-button-group>
        <el-button size="small" @click="selectDate('prev-month')">
          Previous Month
        </el-button>
        <el-button size="small" @click="selectDate('today')">Today</el-button>
        <el-button size="small" @click="selectDate('next-month')">
          Next Month
        </el-button>
      </el-button-group>
    </template>
    <template #date-cell="{ data }">
      <div class="date-cell">
        <span class="day">{{ data.day.split("-").slice(2).join("") }}</span>
        <div
          v-for="contest in contests.filter(item => item.date === data.day)"
          :key="contest.name"
        >
          <span class="contest-tag" @click="openUrl(contest.link)">
            {{ contest.name }}
          </span>
          <a :href="contest.link"> </a>
        </div>
      </div>
    </template>
  </el-calendar>
</template>

<style lang="scss" scoped>
:deep(.el-calendar-table .el-calendar-day) {
  height: auto;
}
.date-cell {
  display: flex;
  flex-direction: column;
  box-sizing: border-box;
  overflow: hidden;
  position: relative; // 添加相对定位
  min-height: 80px; // 设置最小高度
}

.day {
  font-size: 14px;
  font-weight: bold;
}
.contest-tag {
  display: inline-block;
  background-color: #84b9ef;
  color: #ffffff;
  padding: 2px 4px;
  border-radius: 4px;
  font-size: 12px;
  cursor: pointer;
  word-break: break-all;
  margin-bottom: 2px; // 添加底部间距
  position: relative; // 添加相对定位
  z-index: 1; // 设置层级
}
</style>
