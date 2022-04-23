package http

import (
	"github.com/kgysf/go-mirai-http/internal/pkg"
	"github.com/kgysf/go-mirai-http/model"
)

type BindResult model.HttpResult[string]

// Verify 认证
// 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
// NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#%E8%AE%A4%E8%AF%81
func Verify(verifyKey string) (result model.VerifyResult) {
	const url = "/verify"
	body := map[string]string{"verifyKey": verifyKey}
	ret, err := pkg.HttpPost[model.VerifyResult](url, body, true)
	if err != nil {
		result = model.VerifyResult{Code: -10000, Msg: err.Error()}
	} else {
		result = ret
	}
	return
}

func Bind(sessionKey string, qq uint64) (result BindResult) {
	const url = "/verify"
	body := map[string]any{
		"sessionKey": sessionKey,
		"qq":         qq,
	}
	ret, err := pkg.HttpPost[BindResult](url, body, true)
	if err != nil {
		result = BindResult{Code: -10000, Msg: err.Error()}
	} else {
		result = ret
	}
	return
}
