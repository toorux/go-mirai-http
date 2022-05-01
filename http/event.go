package http

import (
	"fmt"
	"github.com/toorux/go-mirai-http/internal/pkg"
	"github.com/toorux/go-mirai-http/model"
)

// NewFriendEvent 添加好友申请
//  使用此方法处理添加好友申请, 关于operate参数的取值：
//  - 0: 同意添加好友
//  - 1: 拒绝添加好友
//  - 2: 拒绝添加好友并添加黑名单，不再接收该用户的好友申请
//  其余用法同官方文档，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#添加好友申请
// ⚠️警告: 所有公告相关接口需要 mirai-api-http 2.5.0 及以上版本
func NewFriendEvent(sessionKey string, event model.Message, operate int, message string) (result any, err error) {
	const url = "/resp/newFriendRequestEvent"
	if event.Type != model.TypeNewFriendRequestEvent {
		err = fmt.Errorf("event参数非 TypeNewFriendRequestEvent 类型")
	}
	body := map[string]any{
		"sessionKey": sessionKey,
		"eventId":    event.EventId,
		"operate":    operate,
		"fromId":     event.FromId,
		"groupId":    event.GroupId,
		"message":    message,
	}
	ret, _err := pkg.HttpPost[any](url, body)
	if _err != nil {
		err = _err
	} else {
		result = ret
	}
	return
}

// MemberJoinRequestEvent 用户入群申请
//  使用此方法处理用户入群申请, 关于operate参数的取值：
//  - 0: 同意入群
//  - 1: 拒绝入群
//  - 2: 忽略请求
//  - 3: 拒绝入群并添加黑名单，不再接收该用户的入群申请
//  - 4: 忽略入群并添加黑名单，不再接收该用户的入群申请
//  其余用法同官方文档，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#用户入群申请
// ⚠️警告: 所有公告相关接口需要 mirai-api-http 2.5.0 及以上版本
func MemberJoinRequestEvent(sessionKey string, event model.Message, operate int, message string) (result any, err error) {
	const url = "/resp/memberJoinRequestEvent"
	if event.Type != model.TypeMemberJoinRequestEvent {
		err = fmt.Errorf("event参数非 TypeMemberJoinEvent 类型")
	}
	body := map[string]any{
		"sessionKey": sessionKey,
		"eventId":    event.EventId,
		"operate":    operate,
		"fromId":     event.FromId,
		"groupId":    event.GroupId,
		"message":    message,
	}
	ret, _err := pkg.HttpPost[any](url, body)
	if _err != nil {
		err = _err
	} else {
		result = ret
	}
	return
}

// BotInvJoinGroupEvent Bot被邀请入群申请
//  使用此方法处理Bot被邀请入群申请, 关于operate参数的取值：
//  - 0: 同意
//  - 1: 拒绝
//  其余用法同官方文档，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#Bot被邀请入群申请
// ⚠️警告: 所有公告相关接口需要 mirai-api-http 2.5.0 及以上版本
func BotInvJoinGroupEvent(sessionKey string, event model.Message, operate int, message string) (result any, err error) {
	const url = "/resp/botInvitedJoinGroupRequestEvent"
	if event.Type != model.TypeBotInvitedJoinGroupRequestEvent {
		err = fmt.Errorf("event参数非 TypeBotInvitedJoinGroupRequestEvent 类型")
	}
	body := map[string]any{
		"sessionKey": sessionKey,
		"eventId":    event.EventId,
		"operate":    operate,
		"fromId":     event.FromId,
		"groupId":    event.GroupId,
		"message":    message,
	}
	ret, _err := pkg.HttpPost[any](url, body)
	if _err != nil {
		err = _err
	} else {
		result = ret
	}
	return
}
