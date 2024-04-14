import { tableData } from "./data";
import { clone, delay, useDark, useECharts } from "@pureadmin/utils";
import { ref, computed, onMounted, reactive, watchEffect } from "vue";
import type { PaginationProps, LoadingConfig, Align } from "@pureadmin/table";
import { templateRef } from "@vueuse/core";

export function useColumns() {
  const dataList = ref([]);
  const loading = ref(true);
  const select = ref("no");
  const hidePie = ref(false);
  const tableSize = ref("default");
  const paginationSmall = ref(false);
  const paginationAlign = ref("center");
  const columns: TableColumnList = [
    {
      type: "selection",
      align: "left",
      reserveSelection: true,
      hide: () => (select.value === "no" ? true : false)
    },
    {
      label: "姓名",
      prop: "name",
      minWidth: 80
    },
    {
      label: "Codeforces竞赛分",
      prop: "cfRating",
      minWidth: 170,
      sortable: true
    },
    {
      label: "Atcoder竞赛分",
      prop: "atcRating",
      minWidth: 145,
      sortable: true
    },
    {
      label: "周统计",
      children: [
        {
          label: "周Rating变化",
          prop: "weeklyRating",
          sortable: true,
          minWidth: 140
        },
        {
          label: "周活跃度",
          prop: "weeklyActive",
          sortable: true,
          minWidth: 110
        }
      ]
    },
    {
      label: "月统计",
      children: [
        {
          label: "月Rating变化",
          prop: "monthlyRating",
          sortable: true,
          minWidth: 137
        },
        {
          label: "月活跃度",
          prop: "monthlyActive",
          sortable: true,
          minWidth: 110
        }
      ]
    },
    {
      label: "能力分布",
      slot: "echart",
      minWidth: 350,
      hide: () => hidePie.value
    }
  ];

  /** 分页配置 */
  const pagination = reactive<PaginationProps>({
    pageSize: 15,
    currentPage: 1,
    pageSizes: [10, 15, 20, 25],
    total: 0,
    align: "center",
    background: true,
    small: false
  });

  /** 加载动画配置 */
  const loadingConfig = reactive<LoadingConfig>({
    text: "正在加载第一页...",
    viewBox: "-10, -10, 50, 50",
    spinner: `
        <path class="path" d="
          M 30 15
          L 28 17
          M 25.61 25.61
          A 15 15, 0, 0, 1, 15 30
          A 15 15, 0, 1, 1, 27.99 7.5
          L 15 15
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)"/>
      `
    // svg: "",
    // background: rgba()
  });

  const { isDark } = useDark();
  const theme = computed(() => (isDark.value ? "dark" : "light"));

  dataList.value.forEach((_, i) => {
    const { setOptions } = useECharts(templateRef(`PieChartRef${i}`), {
      theme
    });

    setOptions({
      tooltip: {
        trigger: "item",
        // 将 tooltip 控制在图表区域里
        confine: true
      },
      series: [
        {
          name: "Github信息",
          type: "pie",
          // center: ["30%", "50%"],
          data: [
            { value: 1067, name: "watchers" },
            { value: 4037, name: "star" },
            { value: 859, name: "forks" }
          ],
          emphasis: {
            itemStyle: {
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowColor: "rgba(0, 0, 0, 0.5)"
            }
          }
        }
      ]
    });
  });

  function onChange(val) {
    pagination.small = val;
  }

  function onSizeChange(val) {
    console.log("onSizeChange", val);
  }

  function onCurrentChange(val) {
    loadingConfig.text = `正在加载第${val}页...`;
    loading.value = true;
    delay(600).then(() => {
      loading.value = false;
    });
  }

  watchEffect(() => {
    pagination.align = paginationAlign.value as Align;
  });

  onMounted(() => {
    delay(600).then(() => {
      const newList = [];
      Array.from({ length: 6 }).forEach(() => {
        newList.push(clone(tableData, true));
      });
      newList.flat(Infinity).forEach((item, index) => {
        dataList.value.push({ id: index, ...item });
      });
      pagination.total = dataList.value.length;
      loading.value = false;
    });
  });

  return {
    loading,
    columns,
    dataList,
    select,
    hidePie,
    tableSize,
    pagination,
    loadingConfig,
    paginationAlign,
    paginationSmall,
    onChange,
    onSizeChange,
    onCurrentChange
  };
}
