export default {
  path: "/",
  name: "Problems",
  redirect: "/problems",
  meta: {
    icon: "ep:list",
    // showLink: false,
    title: "刷题统计",
    rank: 20
  },
  children: [
    {
      path: "/problems",
      name: "ProblemsPage",
      component: () => import("@/views/problems/index.vue"),
      meta: {
        title: "刷题统计"
      }
    }
  ]
} satisfies RouteConfigsTable;
