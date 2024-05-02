package e

var MsgFlags = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "fail",
	INVALID_PARAMS:         "请求参数错误",
	ERROR_USER_CHECK_FAIL:  "用户登录失败",
	ERROR_GEN_TOKEN:        "Token生成失败",
	ERROR_AUTH:             "Token错误",
	ERROR_NOT_VALID_USER:   "用户名或密码错误",
	ERROR_TOKEN_INVALID:    "Token无效",
	AUTH_TOKEN_REQUIRED:    "缺少Token",
	INVALID_REFRESH_TOKEN:  "RefreshToken无效",
	USER_ALREADY_EXIST:     "用户已存在",
	INVALID_PERMISSION:     "您没有操作权限",
	INVALID_VERIFY_CODE:    "验证码错误",
	ERROR_USER_UPDATE_FAIL: "用户信息更新失败",
	ERROR_USER_DELETE_FAIL: "用户删除失败",
	ERROR_UPDATE_ROLE_FAIL: "用户角色更新失败",
	ERROR_UPDATE_PASSWORD_FAIL: "用户密码更新失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
