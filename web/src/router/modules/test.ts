export default {
  path: "/",
  redirect: "/test",
  meta: {
    icon: "tdesign:pearl-of-the-orient",
    // showLink: false,
    title: "测试页面",
    rank: 20
  },
  children: [
    {
      path: "/test",
      name: "test",
      component: () => import("@/views/testPage/index.vue"),
      meta: {
        title: "测试页面"
      }
    }
  ]
} satisfies RouteConfigsTable;
