package http

import (
	"github.com/kgysf/go-mirai-http/internal/pkg"
	"github.com/kgysf/go-mirai-http/model"
)

type BindResult model.HttpResult[string]
type ReleaseResult model.HttpResult[string]

// Verify 认证
//  - 使用此方法验证你的身份，并返回一个会话
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#%E8%AE%A4%E8%AF%81
func Verify(verifyKey string) (result model.VerifyResult) {
	const url = "/verify"
	body := map[string]string{"verifyKey": verifyKey}
	ret, err := pkg.HttpPost[model.VerifyResult](url, body)
	if err != nil {
		result = model.VerifyResult{Code: -10000, Msg: err.Error(), Err: err}
	} else {
		result = ret
	}
	return
}

// Bind 绑定
//  - 使用此方法校验并激活你的Session，同时将Session与一个已登录的Bot绑定
//  - 用法同官方文档，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#%E7%BB%91%E5%AE%9A
func Bind(sessionKey string, qq int) (result BindResult) {
	const url = "/bind"
	body := map[string]any{
		"sessionKey": sessionKey,
		"qq":         qq,
	}
	ret, err := pkg.HttpPost[BindResult](url, body)
	if err != nil {
		result = BindResult{Code: -10000, Msg: err.Error(), Err: err}
	} else {
		result = ret
	}
	return
}

// Release 释放
//  - 使用此方式释放session及其相关资源（Bot不会被释放） 不使用的Session应当被释放，长时间（30分钟）未使用的Session将自动释放，否则Session持续保存Bot收到的消息，将会导致内存泄露
//  - 用法同官方文档，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#%E9%87%8A%E6%94%BE
func Release(sessionKey string, qq int) (result ReleaseResult) {
	const url = "/release"
	body := map[string]any{
		"sessionKey": sessionKey,
		"qq":         qq,
	}
	ret, err := pkg.HttpPost[ReleaseResult](url, body)
	if err != nil {
		result = ReleaseResult{Code: -10000, Msg: err.Error(), Err: err}
	} else {
		result = ret
	}
	return
}
