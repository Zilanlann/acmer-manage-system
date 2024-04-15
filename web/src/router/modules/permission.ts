export default {
  path: "/permission",
  redirect: "/permission/page/index",
  meta: {
    title: "权限管理",
    icon: "ep:lollipop",
    rank: 10
  },
  children: [
    {
      path: "/permission/page/index",
      name: "PermissionPage",
      component: () => import("@/views/permission/page/index.vue"),
      meta: {
        title: "页面权限",
        roles: ["admin"]
      }
    },
    {
      path: "/permission/button/index",
      name: "PermissionButton",
      component: () => import("@/views/permission/button/index.vue"),
      meta: {
        title: "按钮权限",
        roles: ["admin"],
        auths: [
          "permission:btn:add",
          "permission:btn:edit"
          // "permission:btn:delete"
        ]
      }
    }
  ]
} satisfies RouteConfigsTable;
