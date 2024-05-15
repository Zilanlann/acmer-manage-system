<script setup lang="ts">
import Card from "./components/Card.vue";
import { getCardList } from "@/api/favSite";
import { message } from "@/utils/message";
import { ElMessageBox } from "element-plus";
import { ref, onMounted } from "vue";
import dialogForm from "./components/DialogForm.vue";
import dialogAddType from "./components/DialogAddType.vue";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import AddFill from "@iconify-icons/ri/add-circle-line";
import { deleteSite, getSiteTypeList } from "@/api/favSite";
import { useUserStoreHook } from "@/store/modules/user";

defineOptions({
  name: "ListCard"
});

const svg = `
        <path class="path" d="
          M 30 15
          L 28 17
          M 25.61 25.61
          A 15 15, 0, 0, 1, 15 30
          A 15 15, 0, 1, 1, 27.99 7.5
          L 15 15
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)"/>
      `;

const INITIAL_DATA = {
  name: "",
  host: "",
  desc: "",
  siteTypeID: ""
};

const pagination = ref({ current: 1, pageSize: 12, total: 0 });

const siteList = ref([]);
const dataLoading = ref(true);

// 网站类型数据和选中值
const siteTypes = ref([{ id: "", name: "所有类型" }]);
const selectedType = ref("");

// 获取用户角色
const userStore = useUserStoreHook();
const userRoles = userStore.roles;

// 获取网站类型列表数据
const getSiteTypesData = async () => {
  try {
    const { data } = await getSiteTypeList();
    siteTypes.value = [
      { id: "", name: "所有类型" },
      ...data.list.map(type => ({ id: type.ID, name: type.name }))
    ];
  } catch (e) {
    console.log(e);
  }
};

// 获取网站列表数据
const getCardListData = async () => {
  try {
    const { data } = await getCardList();
    siteList.value = data.list;
    pagination.value = {
      ...pagination.value,
      total: data.list.length
    };
  } catch (e) {
    console.log(e);
  } finally {
    setTimeout(() => {
      dataLoading.value = false;
    }, 500);
  }
};

onMounted(() => {
  getSiteTypesData(); // 获取网站类型列表数据
  getCardListData(); // 获取网站列表数据
});

const formDialogVisible = ref(false);
const formAddTypeVisible = ref(false);
const formData = ref({ ...INITIAL_DATA });
const searchValue = ref("");

const onPageSizeChange = (size: number) => {
  pagination.value.pageSize = size;
  pagination.value.current = 1;
};
const onCurrentChange = (current: number) => {
  pagination.value.current = current;
};

const handleDeleteItem = site => {
  ElMessageBox.confirm(
    site ? `确认删除后${site.name}网站信息将被清空, 且无法恢复` : "",
    "提示",
    {
      type: "warning"
    }
  )
    .then(async () => {
      try {
        const response = await deleteSite(site.ID);
        if (response && response.success) {
          message("删除成功", { type: "success" });
          // 删除后重新获取列表数据
          await getCardListData();
        } else {
          message("删除失败", { type: "error" });
        }
      } catch (error) {
        message("删除失败", { type: "error" });
      }
    })
    .catch(() => {});
};
</script>

<template>
  <div>
    <div class="w-full flex justify-between mb-4">
      <div class="flex">
        <!-- Only show these buttons for admin and teacher roles -->
        <el-button
          v-if="userRoles.includes('admin') || userRoles.includes('teacher')"
          :icon="useRenderIcon(AddFill)"
          @click="formDialogVisible = true"
        >
          添加网站
        </el-button>
        <el-button
          v-if="userRoles.includes('admin') || userRoles.includes('teacher')"
          :icon="useRenderIcon(AddFill)"
          @click="formAddTypeVisible = true"
        >
          添加类型
        </el-button>
      </div>
      <div class="flex items-center">
        <el-select
          v-model="selectedType"
          width="300px"
          placeholder="请选择类型"
          clearable
          style="margin-right: 10px"
        >
          <el-option
            v-for="type in siteTypes"
            :key="type.id"
            :label="type.name"
            :value="type.id"
          />
        </el-select>
        <el-input
          v-model="searchValue"
          style="width: 500px"
          placeholder="请输入网站名称"
          clearable
        >
          <template #suffix>
            <el-icon class="el-input__icon">
              <IconifyIconOffline
                v-show="searchValue.length === 0"
                icon="ri:search-line"
              />
            </el-icon>
          </template>
        </el-input>
      </div>
    </div>
    <div
      v-loading="dataLoading"
      :element-loading-svg="svg"
      element-loading-svg-view-box="-10, -10, 50, 50"
    >
      <el-empty
        v-show="
          siteList
            .slice(
              pagination.pageSize * (pagination.current - 1),
              pagination.pageSize * pagination.current
            )
            .filter(
              v =>
                v.name.toLowerCase().includes(searchValue.toLowerCase()) &&
                (selectedType === '' || v.siteTypeID === selectedType)
            ).length === 0
        "
        :description="`${searchValue} 网站不存在`"
      />
      <template v-if="pagination.total > 0">
        <el-row :gutter="16">
          <el-col
            v-for="(site, index) in siteList
              .slice(
                pagination.pageSize * (pagination.current - 1),
                pagination.pageSize * pagination.current
              )
              .filter(
                v =>
                  v.name.toLowerCase().includes(searchValue.toLowerCase()) &&
                  (selectedType === '' || v.siteTypeID === selectedType)
              )"
            :key="index"
            :xs="24"
            :sm="12"
            :md="8"
            :lg="6"
            :xl="4"
          >
            <Card :site="site" @delete-item="handleDeleteItem" />
          </el-col>
        </el-row>
        <el-pagination
          v-model:currentPage="pagination.current"
          class="float-right"
          :page-size="pagination.pageSize"
          :total="pagination.total"
          :page-sizes="[12, 24, 36]"
          :background="true"
          layout="total, sizes, prev, pager, next, jumper"
          @size-change="onPageSizeChange"
          @current-change="onCurrentChange"
        />
      </template>
    </div>
    <dialogForm v-model:visible="formDialogVisible" :data="formData" />
    <dialogAddType v-model:visible="formAddTypeVisible" :data="formData" />
  </div>
</template>
