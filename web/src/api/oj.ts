import { http } from "@/utils/http";
import { v1BaseUrlApi } from "./utils";

type ListResult = {
  success: boolean;
  data?: {
    /** 列表数据 */
    list: Array<any>;
    total: number;
    pageSize: number;
    currentPage: number;
  };
};

type Result = {
  success: boolean;
  code: number;
  msg: string;
  data?: Array<any>;
};

export const getAllOJContestList = (data?: object) => {
  return http.request<ListResult>("get", v1BaseUrlApi("oj/contests"));
};

export const getSubmissionsByUser = (id?: number) => {
  return http.request<ListResult>(
    "get",
    v1BaseUrlApi("oj/submissions/user/" + id)
  );
};

export const getAllProblemTagList = () => {
  return http.request<ListResult>("get", v1BaseUrlApi("oj/submissions/tags"));
};
