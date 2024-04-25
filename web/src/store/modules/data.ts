import { defineStore } from "pinia";
import { store } from "@/store";
import type { testType as Type, userStatusType, userType } from "./types";
import { getLogin, refreshTokenApi } from "@/api/user";
import { type DataInfo, setToken, removeToken, userKey } from "@/utils/auth";
import { getAllUserStatus, type Result, Status } from "@/api/data";

export const useDataStore = defineStore({
  id: "userStatus",
  state: (): userStatusType => ({
    status: []
  }),
  actions: {
    /** 存储status */
    SET_STATUS(status: Array<Status>) {
      this.status = status;
    },
    /** getUserStatus */
    async getUserStatus() {
      return new Promise<Result>((resolve, reject) => {
        getAllUserStatus()
          .then(data => {
            this.SET_STATUS(data.data.userdata);
            resolve(data);
          })
          .catch(error => {
            reject(error);
          });
      });
    }
  }
});

export function useDataStoreHook() {
  return useDataStore(store);
}
