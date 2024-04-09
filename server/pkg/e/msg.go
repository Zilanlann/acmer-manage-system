package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	ERROR:                 "fail",
	INVALID_PARAMS:        "请求参数错误",
	ERROR_USER_CHECK_FAIL: "用户登录失败",
	ERROR_GEN_TOKEN:       "Token生成失败",
	ERROR_AUTH:            "Token错误",
	ERROR_NOT_VALID_USER:  "用户名或密码错误",
	ERROR_TOKEN_INVALID:   "Token无效",
	INVALID_REFRESH_TOKEN: "RefreshToken无效",
	USER_ALREADY_EXIST:    "用户已存在",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
