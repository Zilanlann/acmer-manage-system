export default {
  path: "/calendar",
  redirect: "/calendar/index",
  meta: {
    icon: "material-symbols:calendar-month",
    // showLink: false,
    title: "竞赛日历",
    rank: 99
  },
  children: [
    {
      path: "/calendar/index",
      name: "Calendar",
      component: () => import("@/views/calender/index.vue"),
      meta: {
        title: "竞赛日历"
      }
    }
  ]
} satisfies RouteConfigsTable;
