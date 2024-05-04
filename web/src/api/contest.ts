import { http } from "@/utils/http";
import { baseUrlApi, v1BaseUrlApi } from "./utils";

type Result = {
  success: boolean;
  code: number;
  message: string;
  data?: Array<any>;
};

type ResultTable = {
  success: boolean;
  data?: {
    /** 列表数据 */
    list: Array<any>;
    /** 总条目数 */
    total?: number;
    /** 每页显示条目个数 */
    pageSize?: number;
    /** 当前页数 */
    currentPage?: number;
  };
};

/** 获取比赛管理列表 */
export const getContestList = () => {
  return http.request<ResultTable>("get", v1BaseUrlApi("contests"));
};
