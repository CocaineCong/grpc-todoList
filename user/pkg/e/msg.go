package e

var MsgFlags = map[uint]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	HaveSignUp:                 "已经报名了",
	ErrorActivityTimeout:       "活动过期了",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",
	ErrorNotCompare:            "不匹配",
	ErrorDatabase:              "数据库操作出错,请重试",
	ErrorAuthNotFound:          "Token不能为空",
}

// GetMsg 获取状态码对应信息
func GetMsg(code uint) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
