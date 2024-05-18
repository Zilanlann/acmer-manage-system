import "./reset.css";
import dayjs from "dayjs";
import "dayjs/locale/zh-cn";
import { message } from "@/utils/message";
import { usePublicHooks } from "../hooks";
import type { PaginationProps } from "@pureadmin/table";
import { getKeyList, isAllEmpty, deviceDetection } from "@pureadmin/utils";
import {
  ElForm,
  ElInput,
  ElFormItem,
  ElProgress,
  ElMessageBox
} from "element-plus";
import {
  type Ref,
  h,
  ref,
  toRaw,
  watch,
  computed,
  reactive,
  onMounted
} from "vue";
import { getSubmissionsByUser } from "@/api/oj";
import { useUserStoreHook } from "@/store/modules/user";

export function useProblem(tableRef: Ref, treeRef: Ref) {
  const form = reactive({
    name: "",
    rating: [0, 3000],
    tags: [],
    time: "",
    verdict: ""
  });
  const formRef = ref();
  const ruleFormRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const allData = ref([]); // 存储所有数据
  const switchLoadMap = ref({});
  const { switchStyle } = usePublicHooks();
  const treeLoading = ref(true);
  const selectedNum = ref(0);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "题目编号",
      prop: "ID",
      width: 90
    },
    {
      label: "题目名称",
      prop: "name",
      minWidth: 130
    },
    {
      label: "题目难度",
      prop: "rating",
      sortable: true,
      width: 90
    },
    {
      label: "提交状态",
      prop: "verdict",
      sortable: true,
      width: 90
    },
    {
      label: "题目类型",
      prop: "tags",
      width: 130,
      cellRenderer: ({ row, props }) => (
        <span>
          {row.tags
            ? row.tags.map(tag => (
                <el-tag key={tag.name} size={props.size} effect="plain">
                  {tag.name}
                </el-tag>
              ))
            : null}
        </span>
      )
    },
    {
      label: "提交时间",
      prop: "time",
      minWidth: 90,
      sortable: true,
      formatter: ({ time }) => dayjs(time).format("YYYY-MM-DD HH:mm:ss")
    }
  ];

  const buttonClass = computed(() => {
    return [
      "!h-[20px]",
      "reset-margin",
      "!text-gray-500",
      "dark:!text-white",
      "dark:hover:!text-primary"
    ];
  });

  function updatePaginationData() {
    const start = (pagination.currentPage - 1) * pagination.pageSize;
    const end = start + pagination.pageSize;
    dataList.value = allData.value.slice(start, end); // 根据分页配置截取数据
    pagination.total = allData.value.length; // 更新总数
  }

  function handleSizeChange(val: number) {
    pagination.pageSize = val;
    onSearch();
  }

  function handleCurrentChange(val: number) {
    pagination.currentPage = val;
    onSearch();
  }

  /** 当CheckBox选择项发生变化时会触发该事件 */
  function handleSelectionChange(val) {
    selectedNum.value = val.length;
    // 重置表格高度
    tableRef.value.setAdaptive();
  }

  /** 取消选择 */
  function onSelectionCancel() {
    selectedNum.value = 0;
    // 用于多选表格，清空用户的选择
    tableRef.value.getTableRef().clearSelection();
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getSubmissionsByUser(useUserStoreHook().id);
    const body = toRaw(form);

    allData.value = data.list;

    if (body?.name) {
      allData.value = allData.value.filter(item =>
        item.name.includes(body.name)
      );
    }

    allData.value = allData.value.filter(
      item => item.rating >= body.rating[0] && item.rating <= body.rating[1]
    );

    if (body?.verdict) {
      allData.value = allData.value.filter(
        item => item.verdict === body.verdict
      );
    }

    if (body?.time && body.time.length === 2) {
      const startTime = dayjs(body.time[0]);
      const endTime = dayjs(body.time[1]);

      allData.value = allData.value.filter(item => {
        const itemTime = dayjs(item.time);
        return (
          (itemTime.isAfter(startTime) && itemTime.isBefore(endTime)) ||
          itemTime.isSame(startTime) ||
          itemTime.isSame(endTime)
        );
      });
    }

    // 筛选Tag
    if (body?.tags && body.tags.length > 0) {
      allData.value = allData.value.filter(item => {
        const itemTagIds = item.tags.map(tag => tag.ID);
        return body.tags.every(tagId => itemTagIds.includes(tagId));
      });
    }

    updatePaginationData(); // 更新分页数据
    loading.value = false;
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  onMounted(async () => {
    treeLoading.value = true;
    onSearch();
  });

  return {
    form,
    loading,
    columns,
    dataList,
    treeLoading,
    selectedNum,
    pagination,
    buttonClass,
    deviceDetection,
    onSearch,
    resetForm,
    handleSizeChange,
    onSelectionCancel,
    handleCurrentChange,
    handleSelectionChange
  };
}
