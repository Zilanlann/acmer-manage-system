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
    async testApi() {
      return new Promise<TestResult>((resolve, reject) => {
        console.log("Start testApi - requesting getTest");
        getTest()
          .then(data => {
            console.log("Data received from getTest:", data);
            if (data) {
              console.log(
                "Data is valid - setting test and resolving promise."
              );
              this.SET_TEST(data.data.expires);
              resolve(data);
            } else {
              console.log("Received data is null or undefined.");
              reject(data);
            }
          })
          .catch(error => {
            console.log("Error occurred in getTest:", error);
            reject(error);
          });
      });
    }
  }
});

export function useTestStoreHook() {
  return useTestStore(store);
}
