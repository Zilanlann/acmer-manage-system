export default {
  path: "/",
  name: "Contest",
  redirect: "/contest",
  meta: {
    icon: "tdesign:activity",
    // showLink: false,
    title: "比赛管理",
    rank: 20
  },
  children: [
    {
      path: "/contest",
      name: "ContestManage",
      component: () => import("@/views/contest/index.vue"),
      meta: {
        title: "比赛管理"
      }
    }
  ]
} satisfies RouteConfigsTable;