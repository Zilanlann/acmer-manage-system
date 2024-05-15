export default {
  path: "/",
  name: "Links",
  redirect: "/links",
  meta: {
    icon: "ep:collection",
    // showLink: false,
    title: "常用网站",
    rank: 20
  },
  children: [
    {
      path: "/links",
      name: "LinksPage",
      component: () => import("@/views/linksCard/index.vue"),
      meta: {
        title: "常用网站"
      }
    }
  ]
} satisfies RouteConfigsTable;
