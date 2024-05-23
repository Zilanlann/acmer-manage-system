import { http } from "@/utils/http";
import { baseUrlApi, v1BaseUrlApi } from "./utils";

type ListResult = {
  success: boolean;
  code: number;
  msg: string;
  data?: {
    /** 列表数据 */
    list: Array<any>;
    total: number;
  };
};

export const getAllUserStatus = () => {
  return http.request<ListResult>("get", v1BaseUrlApi("users/status"));
};
