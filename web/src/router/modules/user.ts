export default {
  path: "/",
  name: "User",
  redirect: "/user-manage",
  meta: {
    icon: "ep:user",
    // showLink: false,
    title: "用户管理",
    rank: 20
  },
  children: [
    {
      path: "/user-manage",
      name: "UserManage",
      component: () => import("@/views/user/manage/index.vue"),
      meta: {
        title: "用户管理",
        roles: ["admin", "teacher"]
      }
    }
  ]
} satisfies RouteConfigsTable;
