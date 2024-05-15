import { http } from "@/utils/http";
import { v1BaseUrlApi } from "./utils";
import { number } from "echarts/types/src/echarts.all.js";

type ListResult = {
  success: boolean;
  data?: {
    /** 列表数据 */
    list: Array<any>;
  };
};

type Result = {
  success: boolean;
  code: number;
  msg: string;
  data?: Array<any>;
};

export const createSiteType = (data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("favorite/site-type"), {
    data
  });
};

export const createSite = (data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("favorite/site"), { data });
};

export const getSiteTypeList = (data?: object) => {
  return http.request<ListResult>("get", v1BaseUrlApi("favorite/site-types"), {
    data
  });
};

export const deleteSite = (id?: number) => {
  return http.request<Result>("delete", v1BaseUrlApi("favorite/site/" + id));
};

export const getCardList = (data?: object) => {
  return http.request<ListResult>("get", v1BaseUrlApi("favorite/sites"), {
    data
  });
};
