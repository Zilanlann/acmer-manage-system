export default {
  path: "/test",
  redirect: "/test/index",
  meta: {
    icon: "twemoji:curling-stone",
    // showLink: false,
    title: "测试页面目录",
    rank: 99
  },
  children: [
    {
      path: "/test/index",
      name: "Test",
      component: () => import("@/views/test.vue"),
      meta: {
        title: "测试页面"
      }
    }
  ]
} satisfies RouteConfigsTable;
