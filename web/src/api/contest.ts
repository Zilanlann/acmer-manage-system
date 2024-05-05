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

/** 添加比赛 */
export const createContest = (data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("contest"), { data });
};

/** 更新比赛 */
export const updateContest = (id?: number, data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("contest/" + id), { data });
};

/** 删除比赛 */
export const deleteContest = (id?: number) => {
  return http.request<Result>("delete", v1BaseUrlApi("contest/") + id);
};

/** 添加队伍 */
export const createTeam = (data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("team"), { data });
};

/** 修改队伍 */
export const updateTeam = (id?: number, data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("team/" + id), { data });
};

/** 删除队伍 */
export const deleteTeam = (id?: number, data?: object) => {
  return http.request<Result>("delete", v1BaseUrlApi("team/" + id), { data });
};

/** 添加选手 */
export const createContestant = (data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("contestant"), { data });
};

/** 修改选手 */
export const updateContestant = (id?: number, data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("contestant/" + id), {
    data
  });
};

/** 删除选手 */
export const deleteContestant = (id?: number, data?: object) => {
  return http.request<Result>("delete", v1BaseUrlApi("contestant/" + id), {
    data
  });
};
