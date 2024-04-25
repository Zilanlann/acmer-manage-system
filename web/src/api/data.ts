import { http } from "@/utils/http";
import { baseUrlApi, v1BaseUrlApi } from "./utils";

export type Result = {
  success: boolean;
  code: number;
  msg: string;
  data: {
    userdata: Array<Status>;
  };
};

export type Status = {
  userName: string;
  realName: string;
  cfHandle: string;
  cfRating: number;
  weeklyRating: number;
  monthlyRating: number;
  weeklyActive: number;
  monthlyActive: number;
};

export const getAllUserStatus = () => {
  return http.request<Result>("get", v1BaseUrlApi("all-user-status"));
};
