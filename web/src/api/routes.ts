import { http } from "@/utils/http";
import { baseUrlApi, v1BaseUrlApi } from "./utils";

type Result = {
  success: boolean;
  data: Array<any>;
};

export type TestResult = {
  success: boolean;
  code: number;
  msg: string;
  data: {
    expires: string;
  };
};

export const getAsyncRoutes = () => {
  return http.request<Result>("get", "/get-async-routes");
};

export const getTest = () => {
  return http.request<TestResult>("get", v1BaseUrlApi("test"));
};
