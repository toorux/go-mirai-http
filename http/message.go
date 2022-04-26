package http

import (
	"github.com/kgysf/go-mirai-http/internal/pkg"
	"github.com/kgysf/go-mirai-http/model"
	"strconv"
)

type CountMsgResult model.HttpResult[int]
type MessagesResult model.HttpResult[[]model.Message]
type MessageResult model.HttpResult[model.Message]

// CountMessage 查看队列大小
//  使用此方法获取 session 未读缓存消息的数量
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#%E6%9F%A5%E7%9C%8B%E9%98%9F%E5%88%97%E5%A4%A7%E5%B0%8F
func CountMessage(sessionKey string) (result CountMsgResult) {
	const url = "/countMessage"
	params := pkg.HttpParams{"sessionKey": sessionKey}
	ret, err := pkg.HttpGet[CountMsgResult](url, params)
	if err != nil {
		result = CountMsgResult{Code: -10000, Msg: err.Error(), Err: err}
	} else {
		result = ret
	}
	return
}

// GetMessage 获取消息队列
//  使用此方法整合了官网获取消息队列相关4个接口, flag字段取值范围1-4，分别表示：
//  1. 按时间顺序获取消息，获取消息后从队列中移除
//  2. 获取最新的消息，获取消息后从队列中移除
//  3. 按时间顺序查看消息，查看消息后不从队列中移除
//  4. 查看最新的消息，查看消息后不从队列中移除
//  其余用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#%E6%8E%A5%E6%94%B6%E6%B6%88%E6%81%AF%E4%B8%8E%E4%BA%8B%E4%BB%B6
func GetMessage(flag int, sessionKey string, count int) (result MessagesResult) {
	urls := [4]string{
		"/fetchMessage",
		"/fetchLatestMessage",
		"/peekMessage",
		"/peekLatestMessage",
	}
	params := pkg.HttpParams{
		"sessionKey": sessionKey,
		"count":      strconv.Itoa(count),
	}
	//fmt.Println(urls[flag-1])
	ret, err := pkg.HttpGet[MessagesResult](urls[flag-1], params, true)
	if err != nil {
		result = MessagesResult{Code: -10000, Msg: err.Error(), Err: err}
	} else {
		result = ret
	}
	return
}

// GetMessageForId 通过messageId获取信息
//  当该messageId没有被缓存或缓存失效时，返回code 5(指定对象不存在)
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/api/API.html#通过messageid获取消息
func GetMessageForId(sessionKey string, id int) (result MessageResult) {
	const url = "/messageFromId"
	params := pkg.HttpParams{"sessionKey": sessionKey, "id": strconv.Itoa(id)}
	ret, err := pkg.HttpGet[MessageResult](url, params)
	if err != nil {
		result = MessageResult{Code: -10000, Msg: err.Error(), Err: err}
	} else {
		result = ret
	}
	return
}
