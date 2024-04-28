export default {
  path: "/",
  name: "Test",
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
      name: "TestPage",
      component: () => import("@/views/testPage/index.vue"),
      meta: {
        title: "测试页面",
        roles: ["admin"]
      }
    }
  ]
} satisfies RouteConfigsTable;
