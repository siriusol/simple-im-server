package common

type JsonResult struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(code int32, message string, data interface{}) JsonResult {
	e := HTTPError(code)
	return JsonResult{
		Code: e.Code(),
		Msg:  e.Msg(),
		Data: data,
	}
}
