import { defineStore } from "pinia";
import { store } from "@/store";
import type { testType, userType } from "./types";
import { getLogin, refreshTokenApi } from "@/api/user";
import { type DataInfo, setToken, removeToken, userKey } from "@/utils/auth";
import { getTest, type TestResult } from "@/api/routes";

export const useTestStore = defineStore({
  id: "test",
  state: (): testType => ({
    test: ""
  }),
  actions: {
    /** 存储test */
    SET_TEST(test: string) {
      this.test = test;
    },
    /** test */
    async aaa() {
      // 调用getTest获取测试结果
      const result = await getTest();
      // 将获取的test设置到Pinia的state中
      this.SET_TEST(result.data.expires);
    },
    async testApi(data) {
      return new Promise<TestResult>((resolve, reject) => {
        getTest()
          .then(data => {
            if (data) {
              this.SET_TEST(data.data.expires);
              resolve(data);
            }
          })
          .catch(error => {
            reject(error);
          });
      });
    }
  }
});

export function useTestStoreHook() {
  return useTestStore(store);
}
