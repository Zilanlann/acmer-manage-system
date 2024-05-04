interface FormItemProps {
  ID?: number;
  /** 用于判断是`新增`还是`修改` */
  title: string;
  realname: string;
  username: string;
  password: string;
  phone: string | number;
  email: string;
  class: string;
  studentID: string;
  cfHandle: string;
  atcHandle: string;
  sex: boolean;
}
interface FormProps {
  formInline: FormItemProps;
}

interface RoleFormItemProps {
  realname: string;
  ID?: number;
  /** 角色列表 */
  roleOptions: any[];
  /** 选中的角色 */
  role: string;
}
interface RoleFormProps {
  formInline: RoleFormItemProps;
}

export type { FormItemProps, FormProps, RoleFormItemProps, RoleFormProps };
