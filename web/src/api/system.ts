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

/** 获取系统管理-用户管理列表 */
export const getUserList = () => {
  return http.request<ResultTable>("get", v1BaseUrlApi("users"));
};

/** 系统管理-用户管理-获取所有角色列表 */
export const getAllRoleList = () => {
  return http.request<Result>("get", "/list-all-role");
};

/** 系统管理-用户管理-添加用户 */
export const addUser = (data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("user"), { data });
};

/** 系统管理-用户管理-修改用户信息 */
export const updateUser = (id?: number, data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("user/") + id, { data });
};

/** 系统管理-用户管理-修改用户角色 */
export const updateUserRole = (id?: number, data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("user/") + id + "/role", {
    data
  });
};

/** 系统管理-用户管理-修改用户密码 */
export const updateUserPassword = (id?: number, data?: object) => {
  return http.request<Result>("put", v1BaseUrlApi("user/") + id + "/password", {
    data
  });
};

/** 系统管理-用户管理-删除用户 */
export const deleteUser = (data?: object) => {
  return http.request<Result>("delete", v1BaseUrlApi("user/") + data, { data });
};

/** 系统管理-用户管理-删除多个用户 */
export const deleteUsers = (data?: object) => {
  return http.request<Result>("delete", v1BaseUrlApi("users"), { data });
};
