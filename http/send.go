package http

import (
	"github.com/kgysf/go-mirai-http/internal/pkg"
	. "github.com/kgysf/go-mirai-http/model"
)

// SendFriendMessage 发送好友消息
//  - 使用此方法向指定好友发送消息, quoteId为引用消息id， 无需引用传0
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#发送好友消息
func SendFriendMessage(sessionKey string, target int, chain []MessageChain, quoteId int) (result SendMsgResult) {
	const url = "/sendFriendMessage"
	body := map[string]any{
		"sessionKey":   sessionKey,
		"target":       target,
		"messageChain": chain,
	}
	if quoteId != 0 {
		body["quote"] = quoteId
	}
	ret, _err := pkg.HttpPost[SendMsgResult](url, body)
	if _err != nil {
		result = SendMsgResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// SendGroupMessage 发送群消息
//  - 使用此方法向指定群发送消息, quoteId为引用消息id， 无需引用传0
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#发送群消息
func SendGroupMessage(sessionKey string, target int, chain []MessageChain, quoteId int) (result SendMsgResult) {
	const url = "/sendGroupMessage"
	body := map[string]any{
		"sessionKey":   sessionKey,
		"target":       target,
		"messageChain": chain,
	}
	if quoteId != 0 {
		body["quote"] = quoteId
	}
	ret, _err := pkg.HttpPost[SendMsgResult](url, body)
	if _err != nil {
		result = SendMsgResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// SendTempMessage 发送临时会话消息
//  - 使用此方法向指定群成员发送临时消息, quoteId为引用消息id， 无需引用传0
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#发送临时会话消息
func SendTempMessage(sessionKey string, group int, qq int, chain []MessageChain, quoteId int) (result SendMsgResult) {
	const url = "/sendTempMessage"
	body := map[string]any{
		"sessionKey":   sessionKey,
		"group":        group,
		"qq":           qq,
		"messageChain": chain,
	}
	if quoteId != 0 {
		body["quote"] = quoteId
	}
	ret, _err := pkg.HttpPost[SendMsgResult](url, body)
	if _err != nil {
		result = SendMsgResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// SendNudge 发送头像戳一戳消息
//  - 使用此方法发送头像戳一戳消息
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#发送头像戳一戳消息
// ⚠️警告: 不建议直接使用此函数，建议使用 SendNudgeFriend / SendNudgeMember
func SendNudge(sessionKey string, target int, subject int, kind string) (result HttpResult[any]) {
	const url = "/sendNudge"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     target,
		"subject":    subject,
		"kind":       kind,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// SendNudgeFriend 发送好友头像戳一戳消息
//  - 使用此方法向指定好友发送头像戳一戳消息
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#发送头像戳一戳消息
func SendNudgeFriend(sessionKey string, qq int) (result HttpResult[any]) {
	return SendNudge(sessionKey, qq, qq, "Friend")
}

// SendNudgeMember 发送群成员头像戳一戳消息
//  - 使用此方法向指定群成员发送头像戳一戳消息
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#发送头像戳一戳消息
func SendNudgeMember(sessionKey string, group int, qq int) (result HttpResult[any]) {
	return SendNudge(sessionKey, qq, group, "Group")
}

// Recall 撤回消息
//  - 使用此方法根据messageId撤回指定消息
//  - 用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  - NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#撤回消息
// ⚠️警告: 不建议直接使用此函数，建议使用 SendNudgeFriend / SendNudgeMember
func Recall(sessionKey string, messageId int) (result HttpResult[any]) {
	const url = "/recall"
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
