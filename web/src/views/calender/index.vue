<script setup lang="ts">
import { onMounted, reactive, ref } from "vue";
import type { CalendarDateType, CalendarInstance } from "element-plus";
import { getAllOJContestList } from "@/api/oj";

defineOptions({
  name: "Calendar"
});

const calendar = ref<CalendarInstance>();
const selectDate = (val: CalendarDateType) => {
  if (!calendar.value) return;
  calendar.value.selectDate(val);
};

const contests = reactive<
  Array<{ name: string; oj?: string; link: string; date: string }>
>([]);

const openUrl = (url: string) => {
  window.open(url, "_blank");
};

const getOJLink = (oj: string, contestID: number) => {
  switch (oj.toLowerCase()) {
    case "codeforces":
      return `https://codeforces.com/contest/${contestID}`;
    case "leetcode":
      return `https://leetcode.com/contest/${contestID}`;
    case "atcoder":
      return `https://atcoder.jp/contests/${contestID}`;
    // Add more OJs as needed
    default:
      return "#"; // Default link if OJ is not recognized
  }
};

const transformData = (data: Array<any>) => {
  return data.map(item => ({
    name: item.name,
    oj: item.oj,
    link: getOJLink(item.oj, item.contestID), // Set the link based on the OJ and contestID
    date: item.startTime.split("T")[0] // Extracting the date part from the startTime
  }));
};

onMounted(async () => {
  try {
    const response = await getAllOJContestList();
    if (response.success && response.data?.list) {
      const transformedData = transformData(response.data.list);
      contests.push(...transformedData);
    } else {
      console.error("Failed to fetch contest data:", response);
    }
  } catch (error) {
    console.error("Error fetching contest data:", error);
  }
});
</script>

<template>
  <el-calendar ref="calendar">
    <template #header="{ date }">
      <span>{{ date }}</span>
      <b><span>各大OJ竞赛日历</span></b>
      <el-button-group>
        <el-button size="small" @click="selectDate('prev-month')"
          >上个月</el-button
        >
        <el-button size="small" @click="selectDate('today')">今天</el-button>
        <el-button size="small" @click="selectDate('next-month')"
          >下个月</el-button
        >
      </el-button-group>
    </template>
    <template #date-cell="{ data }">
      <div class="date-cell">
        <span class="day">{{ data.day.split("-").slice(2).join("") }}</span>
        <div
          v-for="contest in contests.filter(item => item.date === data.day)"
          :key="contest.name"
        >
          <span class="contest-tag" @click="openUrl(contest.link)">{{
            contest.name
          }}</span>
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
