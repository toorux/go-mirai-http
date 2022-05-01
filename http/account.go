package http

import (
	"github.com/toorux/go-mirai-http/internal/pkg"
	. "github.com/toorux/go-mirai-http/model"
	"strconv"
)

type FriendsResult HttpResult[[]Friend]
type GroupsResult HttpResult[[]Group]
type MembersResult HttpResult[[]Member]

// FriendList 获取好友列表
//  使用此方法获取bot的好友列表
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取好友列表
func FriendList(sessionKey string) (result FriendsResult) {
	const url = "/friendList"
	params := pkg.HttpParams{"sessionKey": sessionKey}
	ret, err := pkg.HttpGet[FriendsResult](url, params)
	if err != nil {
		result = FriendsResult{Code: -10000, Msg: err.Error()}
	} else {
		result = ret
	}
	return
}

// GroupList 获取群列表
//  使用此方法获取bot的群列表
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取群列表
func GroupList(sessionKey string) (result GroupsResult) {
	const url = "/groupList"
	params := pkg.HttpParams{"sessionKey": sessionKey}
	ret, err := pkg.HttpGet[GroupsResult](url, params)
	if err != nil {
		result = GroupsResult{Code: -10000, Msg: err.Error()}
	} else {
		result = ret
	}
	return
}

// MemberList 获取群成员列表
//  使用此方法获取bot指定群中的成员列表
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取群成员列表
func MemberList(sessionKey string, groupId int) (result MembersResult) {
	const url = "/memberList"
	params := pkg.HttpParams{"sessionKey": sessionKey, "target": strconv.Itoa(groupId)}
	ret, err := pkg.HttpGet[MembersResult](url, params)
	if err != nil {
		result = MembersResult{Code: -10000, Msg: err.Error()}
	} else {
		result = ret
	}
	return
}

// BotProfile 获取Bot资料
//  此接口获取 session 绑定 bot 的详细资料
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取Bot资料
func BotProfile(sessionKey string) (result Profile, err error) {
	const url = "/botProfile"
	params := pkg.HttpParams{"sessionKey": sessionKey}
	ret, _err := pkg.HttpGet[Profile](url, params)
	if _err != nil {
		err = _err
	} else {
		result = ret
	}
	return
}

// FriendProfile 获取好友资料
//  此接口获取好友的详细资料
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取好友资料
func FriendProfile(sessionKey string, friendId int) (result Profile, err error) {
	const url = "/friendProfile"
	params := pkg.HttpParams{"sessionKey": sessionKey, "target": strconv.Itoa(friendId)}
	ret, _err := pkg.HttpGet[Profile](url, params)
	if _err != nil {
		err = _err
	} else {
		result = ret
	}
	return
}

// MemberProfile 获取群成员资料
//  此接口获取群成员的消息资料
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取群成员资料
func MemberProfile(sessionKey string, groupId int, memberId int) (result Profile, err error) {
	const url = "/memberProfile"
	params := pkg.HttpParams{"sessionKey": sessionKey, "target": strconv.Itoa(groupId), "memberId": strconv.Itoa(memberId)}
	ret, _err := pkg.HttpGet[Profile](url, params)
	if _err != nil {
		err = _err
	} else {
		result = ret
	}
	return
}

// UserProfile 获取QQ用户资料
//  此接口获取任意QQ用户的资料
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#获取QQ用户资料
// ⚠️警告: 此条API需要 mirai-api-http 2.5.0及以上版本
func UserProfile(sessionKey string, id int) (result Profile, err error) {
	const url = "/userProfile"
	params := pkg.HttpParams{"sessionKey": sessionKey, "target": strconv.Itoa(id)}
	ret, _err := pkg.HttpGet[Profile](url, params)
	if _err != nil {
		err = _err
	} else {
		result = ret
	}
	return
}

// DeleteFriend 删除好友
//  使用此方法删除指定好友
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#删除好友
// ⚠️警告: 此接口未测试
func DeleteFriend(sessionKey string, target int) (result HttpResult[any]) {
	const url = "/deleteFriend"
	body := map[string]any{
		"sessionKey": sessionKey,
		"target":     target,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}
