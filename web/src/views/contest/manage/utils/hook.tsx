import "./reset.css";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import duration from "dayjs/plugin/duration";
import "dayjs/locale/zh-cn";
import editForm from "../form/index.vue";
import teamForm from "../form/team.vue";
import contestantForm from "../form/contestant.vue";
import { zxcvbn } from "@zxcvbn-ts/core";
import { message } from "@/utils/message";
import userAvatar from "@/assets/user.jpg";
import { usePublicHooks } from "../../hooks";
import { addDialog } from "@/components/ReDialog";
import type { PaginationProps } from "@pureadmin/table";
import type {
  ContestantFormItemProps,
  FormItemProps,
  TeamFormItemProps
} from "../utils/types";
import { getKeyList, isAllEmpty, deviceDetection } from "@pureadmin/utils";
import {
  getACMerList,
  getTeachersList,
  updateUser,
  updateUserPassword
} from "@/api/system";
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
import {
  createContest,
  createContestant,
  createTeam,
  deleteContest,
  deleteContestant,
  deleteTeam,
  getContestList,
  updateContest,
  updateContestant,
  updateTeam
} from "@/api/contest";

export function useContest(tableRef: Ref, treeRef: Ref) {
  const form = reactive({
    username: "",
    contestName: "",
    time: ""
  });
  const formRef = ref();
  const ruleFormRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  // 上传头像信息
  const avatarInfo = ref();
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
      label: "勾选列",
      type: "selection",
      fixed: "left",
      reserveSelection: true // 数据刷新后保留选项
    },
    {
      type: "expand",
      slot: "expand"
    },
    {
      label: "比赛编号",
      prop: "ID",
      width: 90
    },
    {
      label: "比赛名称",
      prop: "name",
      minWidth: 130
    },
    {
      label: "正式赛开始时间",
      prop: "startTime",
      minWidth: 90,
      sortable: true,
      formatter: ({ startTime }) =>
        dayjs(startTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "正式赛结束时间",
      prop: "endTime",
      minWidth: 90,
      sortable: true,
      formatter: ({ endTime }) => dayjs(endTime).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "比赛持续时长",
      prop: "length",
      minWidth: 90,
      formatter: ({ startTime, endTime }) =>
        dayjs.duration(dayjs(endTime).diff(dayjs(startTime))).format("HH:mm:ss")
    },
    {
      label: "创建时间",
      prop: "CreatedAt",
      minWidth: 90,
      formatter: ({ CreatedAt }) =>
        dayjs(CreatedAt).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      fixed: "right",
      width: 260,
      slot: "operation"
    }
  ];
  const teamColumns: TableColumnList = [
    {
      label: "勾选列",
      type: "selection",
      fixed: "left",
      reserveSelection: true // 数据刷新后保留选项
    },
    {
      type: "expand",
      slot: "expand"
    },
    {
      label: "队伍编号",
      prop: "ID",
      width: 90
    },
    {
      label: "队伍中文名",
      prop: "zhName",
      minWidth: 130
    },
    {
      label: "队伍英文名",
      prop: "enName",
      minWidth: 130
    },
    {
      label: "操作",
      fixed: "right",
      width: 230,
      slot: "operation"
    }
  ];
  const contestantColumns: TableColumnList = [
    {
      label: "勾选列",
      type: "selection",
      fixed: "left",
      reserveSelection: true // 数据刷新后保留选项
    },
    {
      label: "队员编号",
      prop: "ID",
      width: 90
    },
    {
      label: "姓名",
      prop: "user.realname",
      minWidth: 130
    },
    {
      label: "班级",
      prop: "user.class",
      minWidth: 130
    },
    {
      label: "学号",
      prop: "user.studentID",
      minWidth: 130
    },
    {
      label: "手机号码",
      prop: "user.phone",
      minWidth: 90
      // formatter: ({ phone }) => hideTextAtIndex(phone, { start: 3, end: 6 })
    },
    {
      label: "用户邮箱",
      prop: "user.email",
      minWidth: 180
    },
    {
      label: "操作",
      fixed: "right",
      width: 180,
      slot: "operation"
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
  // 重置的新密码
  const pwdForm = reactive({
    password: ""
  });
  const pwdProgress = [
    { color: "#e74242", text: "非常弱" },
    { color: "#EFBD47", text: "弱" },
    { color: "#ffa500", text: "一般" },
    { color: "#1bbf1b", text: "强" },
    { color: "#008000", text: "非常强" }
  ];
  // 当前密码强度（0-4）
  const curScore = ref();
  const roleOptions = ref([]);
  const teacherOptions = ref([]);
  const acmerOptions = ref([]);

  function onChange({ row, index }) {
    ElMessageBox.confirm(
      `确认要<strong>${
        row.status === 0 ? "停用" : "启用"
      }</strong><strong style='color:var(--el-color-primary)'>${
        row.username
      }</strong>用户吗?`,
      "系统提示",
      {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
        dangerouslyUseHTMLString: true,
        draggable: true
      }
    )
      .then(() => {
        switchLoadMap.value[index] = Object.assign(
          {},
          switchLoadMap.value[index],
          {
            loading: true
          }
        );
        setTimeout(() => {
          switchLoadMap.value[index] = Object.assign(
            {},
            switchLoadMap.value[index],
            {
              loading: false
            }
          );
          message("已成功修改用户状态", {
            type: "success"
          });
        }, 300);
      })
      .catch(() => {
        row.status === 0 ? (row.status = 1) : (row.status = 0);
      });
  }

  async function handleDeleteContest(row) {
    await deleteContest(row.ID).then(() => {
      message(`您删除了比赛编号为${row.ID}的这条数据`, { type: "success" });
    });
    onSearch();
  }

  async function handleDeleteTeam(row) {
    await deleteTeam(row.ID).then(() => {
      message(`您删除了队伍编号为${row.ID}的这条数据`, { type: "success" });
    });
    onSearch();
  }

  async function handleDeleteContestant(row) {
    await deleteContestant(row.ID).then(() => {
      message(`您删除了选手编号为${row.ID}的这条数据`, { type: "success" });
    });
    onSearch();
  }

  function handleSizeChange(val: number) {
    console.log(`${val} items per page`);
  }

  function handleCurrentChange(val: number) {
    console.log(`current page: ${val}`);
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

  /** 批量删除 */
  function onbatchDelContest() {
    // 返回当前选中的行
    const curSelected = tableRef.value.getTableRef().getSelectionRows();
    console.log(curSelected);
    // 接下来根据实际业务，通过选中行的某项数据，比如下面的id，调用接口进行批量删除
    message(`已删除用户编号为 ${getKeyList(curSelected, "ID")} 的数据`, {
      type: "success"
    });
    tableRef.value.getTableRef().clearSelection();
    onSearch();
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getContestList();
    const body = toRaw(form);
    let list = data.list;
    list = list.filter(item => item.name.includes(body?.contestName));
    // if (body.phone) list = list.filter(item => item.phone === body.phone);

    dataList.value = list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;

    loading.value = false;
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    console.log(formEl);
    onSearch();
  };

  function openDialog(title?: string, row?: FormItemProps) {
    addDialog({
      title: `${title}比赛`,
      props: {
        formInline: {
          title,
          name: row?.name ?? "",
          time: row?.time ?? "",
          startTime: row?.startTime ?? "",
          endTime: row?.endTime ?? "",
          desc: row?.desc ?? ""
        }
      },
      width: "46%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { ref: formRef }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as FormItemProps;
        /** 添加比赛 */
        async function handleCreateContest(data) {
          await createContest(data)
            .then(res => {
              chores();
            })
            .catch(err => {
              message(`比赛名称为${curData.name}的这条数据${title}失败`, {
                type: "error"
              });
            });
        }
        /** 更新比赛 */
        async function handleUpdateContest(data) {
          const id = row?.ID;
          await updateContest(id, data)
            .then(() => {
              chores();
            })
            .catch(err => {
              message(`比赛名称为${curData.name}的这条数据${title}失败`, {
                type: "error"
              });
            });
        }
        function chores() {
          message(`您${title}了比赛名称为${curData.name}的这条数据`, {
            type: "success"
          });
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(valid => {
          if (valid) {
            // 表单规则校验通过
            curData.startTime = dayjs(curData.time[0]).format(
              "YYYY-MM-DDTHH:mm:ssZ"
            );
            curData.endTime = dayjs(curData.time[1]).format(
              "YYYY-MM-DDTHH:mm:ssZ"
            );
            const { time, ...cleanedData } = curData;
            if (title === "新增") {
              handleCreateContest(cleanedData);
            } else {
              handleUpdateContest(cleanedData);
            }
          }
        });
      }
    });
  }

  /** 打开队伍添加/修改弹框 */
  function openTeamDialog(title?: string, row?: TeamFormItemProps) {
    addDialog({
      title: `${title}队伍`,
      props: {
        formInline: {
          title,
          contestID: row?.ID ?? 0,
          zhName: row?.zhName ?? "",
          enName: row?.enName ?? "",
          teacherOptions: teacherOptions.value,
          coachID: teacherOptions.value[0]?.ID ?? 0,
          desc: row?.desc ?? ""
        }
      },
      width: "46%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(teamForm, { ref: formRef }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as TeamFormItemProps;
        const { teacherOptions, ...cleanedData } = curData;
        /** 添加队伍 */
        async function handleCreateTeam(data) {
          await createTeam(data)
            .then(res => {
              chores();
            })
            .catch(err => {
              message(`队伍名称为${curData.zhName}的这条数据${title}失败`, {
                type: "error"
              });
            });
        }
        /** 更新队伍 */
        async function handleUpdateTeam(data) {
          const id = row?.ID;
          await updateTeam(id, data)
            .then(() => {
              chores();
            })
            .catch(err => {
              message(`队伍名称为${curData.zhName}的这条数据${title}失败`, {
                type: "error"
              });
            });
        }
        function chores() {
          message(`您${title}了队伍名称为${curData.zhName}的这条数据`, {
            type: "success"
          });
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(valid => {
          if (valid) {
            // 表单规则校验通过
            if (title === "新增") {
              handleCreateTeam(cleanedData);
            } else {
              handleUpdateTeam(cleanedData);
            }
          }
        });
      }
    });
  }

  /** 打开选手添加/修改弹框 */
  function openContestantDialog(title?: string, row?: ContestantFormItemProps) {
    addDialog({
      title: `${title}选手`,
      props: {
        formInline: {
          title,
          teamID: row?.ID ?? 0,
          acmerOptions: acmerOptions.value,
          userID: acmerOptions.value[0]?.ID ?? 0
        }
      },
      width: "46%",
      draggable: true,
      fullscreen: deviceDetection(),
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(contestantForm, { ref: formRef }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as ContestantFormItemProps;
        /** 添加选手 */
        async function handleCreateContestant(data) {
          await createContestant(data)
            .then(res => {
              chores();
            })
            .catch(err => {
              message(`选手ID为${curData.userID}的这条数据${title}失败`, {
                type: "error"
              });
            });
        }
        /** 更新选手 */
        async function handleUpdateContestant(data) {
          const id = row?.ID;
          await updateContestant(id, data)
            .then(() => {
              chores();
            })
            .catch(err => {
              message(`队伍名称为${curData.userID}的这条数据${title}失败`, {
                type: "error"
              });
            });
        }
        function chores() {
          message(`您${title}了队伍名称为${curData.userID}的这条数据`, {
            type: "success"
          });
          done(); // 关闭弹框
          onSearch(); // 刷新表格数据
        }
        FormRef.validate(valid => {
          if (valid) {
            // 表单规则校验通过
            const { acmerOptions, ...cleanedData } = curData;
            if (title === "新增") {
              handleCreateContestant(cleanedData);
            } else {
              handleUpdateContestant(cleanedData);
            }
          }
        });
      }
    });
  }

  watch(
    pwdForm,
    ({ password: newPwd }) =>
      (curScore.value = isAllEmpty(newPwd) ? -1 : zxcvbn(newPwd).score)
  );

  onMounted(async () => {
    treeLoading.value = true;
    dayjs.extend(relativeTime);
    dayjs.extend(duration);
    onSearch();
    teacherOptions.value = (await getTeachersList()).data.list;
    acmerOptions.value = (await getACMerList()).data.list;
    console.log(teacherOptions.value);
  });

  return {
    form,
    loading,
    columns,
    teamColumns,
    contestantColumns,
    dataList,
    treeLoading,
    selectedNum,
    pagination,
    buttonClass,
    deviceDetection,
    onSearch,
    resetForm,
    onbatchDel: onbatchDelContest,
    openDialog,
    openTeamDialog,
    openContestantDialog,
    handleDeleteContest,
    handleDeleteTeam,
    handleDeleteContestant,
    handleSizeChange,
    onSelectionCancel,
    handleCurrentChange,
    handleSelectionChange
  };
}
