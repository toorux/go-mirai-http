package http

import (
	"github.com/toorux/go-mirai-http/internal/pkg"
	. "github.com/toorux/go-mirai-http/model"
	"strconv"
)

// Mute 禁言群成员
//  使用此方法禁言群成员, 需要对应权限
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#禁言群成员
func Mute(sessionKey string, group int, qq int, time int) (result HttpResult[any]) {
	const url = "/mute"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
		"memberId":   qq,
		"time":       time,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// Unmute 解除群成员禁言
//  使用此方法解除群成员禁言, 需要对应权限
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#解除群成员禁言
func Unmute(sessionKey string, group int, qq int) (result HttpResult[any]) {
	const url = "/unmute"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
		"memberId":   qq,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// Kick 移除群成员
//  使用此方法移除群成员, 需要对应权限
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#移除群成员
func Kick(sessionKey string, group int, qq int, msg string) (result HttpResult[any]) {
	const url = "/kick"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
		"memberId":   qq,
		"msg":        msg,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// Quit 退出群聊
//  使用此方法退出群聊
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#退出群聊
func Quit(sessionKey string, group int) (result HttpResult[any]) {
	const url = "/quit"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// MuteAll 全体禁言
//  使用此方法全体禁言, 需要对应权限
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#全体禁言
func MuteAll(sessionKey string, group int) (result HttpResult[any]) {
	const url = "/muteAll"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// UnmuteAll 解除全体禁言
//  使用此方法解除全体禁言, 需要对应权限
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#解除全体禁言
func UnmuteAll(sessionKey string, group int) (result HttpResult[any]) {
	const url = "/unmuteAll"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// SetEssence 解除全体禁言
//  使用此方法解除全体禁言, 需要对应权限
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#解除全体禁言
func SetEssence(sessionKey string, messageId int) (result HttpResult[any]) {
	const url = "/setEssence"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     messageId,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// GetGroupConfig 获取群设置
//  使用此方法获取群设置
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取群设置
func GetGroupConfig(sessionKey string, group int) (result GroupConfig, e error) {
	const url = "/groupConfig"
	params := pkg.HttpParams{
		"sessionKey": sessionKey,
		"target":     strconv.Itoa(group),
	}
	ret, _err := pkg.HttpGet[GroupConfig](url, params)
	if _err != nil {
		e = _err
	} else {
		result = ret
	}
	return
}

// SetGroupConfig 修改群设置
//  使用此方法修改群设置（需要有相关限权）
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#修改群设置
// ⚠️警告: 根据官方源码来看，目前只支持修改Name和AllowMemberInvite，后续官方支持后即可支持
func SetGroupConfig(sessionKey string, group int, config GroupConfig) (result HttpResult[any]) {
	const url = "/groupConfig"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
		"config":     config,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// GetMemberInfo 获取群员设置
//  使用此方法获取群员设置(资料)
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取群员设置
func GetMemberInfo(sessionKey string, group int, qq int) (result Member, e error) {
	const url = "/memberInfo"
	params := pkg.HttpParams{
		"sessionKey": sessionKey,
		"target":     strconv.Itoa(group),
		"memberId":   strconv.Itoa(qq),
	}
	ret, _err := pkg.HttpGet[Member](url, params)
	if _err != nil {
		e = _err
	} else {
		result = ret
	}
	return
}

// SetMemberInfo 修改群员设置(只支持改群名片)
//  使用此方法修改群员资料（需要有相关限权）
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#修改群员设置
func SetMemberInfo(sessionKey string, group int, qq int, name string) (result HttpResult[any]) {
	const url = "/memberInfo"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
		"memberId":   qq,
		"info": map[string]string{
			"name": name,
		},
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// SetMemberAdmin 修改群员管理员
//  使用此方法修改群员的管理员权限（需要有群主限权）
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#修改群员管理员
// ⚠️警告: 未测试
func SetMemberAdmin(sessionKey string, group int, qq int, assign bool) (result HttpResult[any]) {
	const url = "/memberAdmin"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     group,
		"memberId":   qq,
		"assign":     assign,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}
