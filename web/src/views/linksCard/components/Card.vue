<script setup lang="ts">
import { PropType } from "vue";
import More2Fill from "@iconify-icons/ri/more-2-fill";
import { deleteSite } from "@/api/favSite";
import { message } from "@/utils/message";

defineOptions({
  name: "ReCard"
});

interface CardType {
  ID: number;
  siteTypeID: number;
  desc: string;
  name: string;
  host: string;
}

const props = defineProps({
  site: {
    type: Object as PropType<CardType>
  }
});

const emit = defineEmits(["modify-site", "delete-item"]);

const handleClickManage = (site: CardType) => {
  emit("modify-site", site);
};

const handleClickDelete = (site: CardType) => {
  emit("delete-item", site);
};

// 获取favicon的URL
const getFaviconUrl = (host: string) => {
  return `https://www.google.com/s2/favicons?sz=64&domain_url=${host}`;
};
</script>

<template>
  <div class="list-card-item">
    <div class="list-card-item_detail bg-bg_color">
      <el-row justify="space-between">
        <div class="list-card-item_detail--logo">
          <img :src="getFaviconUrl(site.host)" alt="favicon" />
        </div>
        <div class="list-card-item_detail--operation">
          <el-dropdown trigger="click">
            <IconifyIconOffline :icon="More2Fill" class="text-[24px]" />
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item @click="handleClickDelete(site)">
                  删除
                </el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-row>
      <p class="list-card-item_detail--name text-text_color_primary">
        <a :href="site.host" target="_blank" rel="noopener noreferrer">
          {{ site.name }}
        </a>
      </p>
      <p class="list-card-item_detail--desc text-text_color_regular">
        {{ site.desc }}
      </p>
    </div>
  </div>
</template>

<style scoped lang="scss">
.list-card-item {
  display: flex;
  flex-direction: column;
  margin-bottom: 12px;
  overflow: hidden;
  cursor: pointer;
  border-radius: 3px;

  &_detail {
    flex: 1;
    min-height: 140px;
    padding: 24px 32px;

    &--logo {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 46px;
      height: 46px;
      font-size: 26px;
      // Removed the background property
      // background: #e0ebff;
      border-radius: 50%;

      img {
        width: 100%;
        height: 100%;
        border-radius: 50%;
      }
    }

    &--operation {
      display: flex;
      height: 100%;
    }

    &--name {
      margin: 24px 0 8px;
      font-size: 20px;
      font-weight: 600;

      a {
        color: inherit;
        text-decoration: none;

        &:hover {
          text-decoration: underline;
        }
      }
    }

    &--desc {
      display: -webkit-box;
      height: 40px;
      margin-bottom: 24px;
      overflow: hidden;
      font-size: 12px;
      line-height: 20px;
      text-overflow: ellipsis;
      -webkit-line-clamp: 2;
      -webkit-box-orient: vertical;
    }
  }
}
</style>
