package model

// VerifyResult
// 认证接口直接返回值(无序嵌套HttpResult)
type VerifyResult struct {
	Code    int
	Msg     string
	Session string
	Err     error
}
