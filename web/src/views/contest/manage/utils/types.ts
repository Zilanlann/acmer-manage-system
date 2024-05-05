interface FormItemProps {
  ID?: number;
  /** 用于判断是`新增`还是`修改` */
  title: string;
  name: string;
  time: string;
  startTime: string;
  endTime: string;
  desc: string;
}
interface FormProps {
  formInline: FormItemProps;
}

interface RoleFormItemProps {
  ID?: number;
  realname: string;
  /** 角色列表 */
  roleOptions: any[];
  /** 选中的角色 */
  role: string;
}
interface RoleFormProps {
  formInline: RoleFormItemProps;
}

interface TeamFormItemProps {
  ID?: number;
  /** 用于判断是`新增`还是`修改` */
  title: string;
  zhName: string;
  enName: string;
  desc: string;
  /** 教师列表 */
  teacherOptions: any[];
  /** 选中的教师 */
  coachID: number;
}
interface TeamFormProps {
  formInline: TeamFormItemProps;
}

interface ContestantFormItemProps {
  ID?: number;
  /** 用于判断是`新增`还是`修改` */
  title: string;
  /** ACMer列表 */
  acmerOptions: any[];
  /** 选中的ACMer */
  userID: number;
}
interface ContestantFormProps {
  formInline: ContestantFormItemProps;
}

export type {
  FormItemProps,
  FormProps,
  RoleFormItemProps,
  RoleFormProps,
  TeamFormItemProps,
  TeamFormProps,
  ContestantFormItemProps,
  ContestantFormProps
};
