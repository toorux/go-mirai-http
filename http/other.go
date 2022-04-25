package http

import (
	"github.com/kgysf/go-mirai-http/internal/pkg"
	. "github.com/kgysf/go-mirai-http/model"
)

// About 获取插件信息
//  - 使用此方法获取插件的信息，如版本号
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取插件信息
func About() (result PluginInfoResult) {
	const url = "/about"
	ret, err := pkg.HttpGet[PluginInfoResult](url)
	if err != nil {
		result = PluginInfoResult{Code: -10000, Msg: err.Error(), Err: err}
	} else {
		result = ret
	}
	return
}
