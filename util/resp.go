package util

// Code 返回码
type Code int

// 返回码定义
const (
	CodeOK       = 0
	CodeParamErr = 1
	CodeSrvErr   = 2
)

// Resp Http请求返回Body
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewRespWithData(data interface{}) *Resp {
	return &Resp{
		Code: CodeOK,
		Data: data,
	}
}

func NewRespWithMsg(code Code, msg string) *Resp {
	return &Resp{
		Code: CodeOK,
		Msg:  msg,
	}
}
