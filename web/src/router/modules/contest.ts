export default {
  path: "/",
  name: "Contest",
  redirect: "/contest",
  meta: {
    icon: "ep:medal",
    // showLink: false,
    title: "比赛管理",
    rank: 20
  },
  children: [
    {
      path: "/contest",
      name: "ContestManage",
      component: () => import("@/views/contest/manage/index.vue"),
      meta: {
        title: "比赛管理",
        roles: ["admin", "teacher"]
      }
    }
  ]
} satisfies RouteConfigsTable;
