package http

import (
	"fmt"
	"github.com/toorux/go-mirai-http/internal/pkg"
	. "github.com/toorux/go-mirai-http/model"
	"strconv"
)

type AnnoListResult HttpResult[[]Anno]
type AnnoResult HttpResult[Anno]

// AnnoList 获取群公告
//  此方法获取指定群公告列表
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#获取群公告
// ⚠️警告: 所有公告相关接口需要 mirai-api-http 2.5.0 及以上版本
func AnnoList(sessionKey string, group int, offset int, size int) (result AnnoListResult) {
	const url = "/anno/list"
	params := pkg.HttpParams{
		"sessionKey": sessionKey,
		"id":         strconv.Itoa(group),
		"offset":     strconv.Itoa(offset),
		"size":       strconv.Itoa(size),
	}
	ret, _err := pkg.HttpGet[AnnoListResult](url, params)
	if _err != nil {
		result = AnnoListResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// AnnoPublish 发布群公告
//  此方法获取指定群公告列表, imageUrl参数设置为""表示无图片，args参数解释如下
//  - args[0]: 是否发送给新成员
//  - args[1]: 是否引导用户修改名片
//  - args[2]: 是否自动弹窗显示
//  - args[3]: 是否需要成员确认
//  以上参数默认皆为false, 多余的参数将会被忽略
//  其余用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#发布群公告
// ⚠️警告: 所有公告相关接口需要 mirai-api-http 2.5.0 及以上版本
func AnnoPublish(sessionKey string, group int, content string, imageUrl string, pinned bool, args ...bool) (result AnnoResult) {
	const url = "/anno/publish"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
		"content":    content,
		"pinned":     pinned,
	}
	keys := []string{"sendToNewMember", "showEditCard", "showPopup", "requireConfirmation"}
	for i, b := range args {
		if i > 3 {
			break
		}
		key := keys[i]
		body[key] = b
	}
	if imageUrl != "" {
		body["imageUrl"] = imageUrl
	}
	fmt.Println(body)
	ret, _err := pkg.HttpPost[AnnoResult](url, body)
	if _err != nil {
		result = AnnoResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// AnnoDelete 删除群公告
//  此方法删除指定群中一条公告
//  其余用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#删除群公告
// ⚠️警告: 所有公告相关接口需要 mirai-api-http 2.5.0 及以上版本
func AnnoDelete(sessionKey string, group int, id string) (result AnnoResult) {
	const url = "/anno/delete"
	body := map[string]any{
		"sessionKey": sessionKey,
		"id":         group,
		"fid":        id,
	}
	ret, _err := pkg.HttpPost[AnnoResult](url, body)
	if _err != nil {
		result = AnnoResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}
