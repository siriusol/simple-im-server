package common

type Error interface {
	Code() int32
	Msg() string
}

type HTTPError int32

const (
	OK           HTTPError = 200
	NotLogin     HTTPError = 4001
	ParamIllegal HTTPError = 4002
)

func (e HTTPError) Code() int32 {
	return int32(e)
}

func (e HTTPError) Msg() string {
	switch e {
	case OK:
		return "成功"
	case NotLogin:
		return "未登录"
	case ParamIllegal:
		return "参数错误"
	default:
		return "未知错误"
	}
}
