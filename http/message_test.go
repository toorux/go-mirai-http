package http

import (
	"fmt"
	"github.com/toorux/go-mirai-http/conf"
	"github.com/toorux/go-mirai-http/model"
	"testing"
	"time"
)

func TestCountMessage(t *testing.T) {
	// 休眠多久
	var wait time.Duration = 5

	ret := Verify(conf.VerifyKey)
	fmt.Printf("Verify Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
	session := ret.Session

	retBind := Bind(session, conf.QQ)
	fmt.Printf("Bind Code: %d, Msg: %s\n", retBind.Code, retBind.Msg)

	fmt.Printf("Sleep %ds......\n", wait)
	time.Sleep(time.Millisecond * 1000 * wait)

	defer func() {
		retRelease := Release(session, conf.QQ)
		fmt.Printf("Release Code: %d, Msg: %s\n", retRelease.Code, retRelease.Msg)
	}()

	for {
		retCM := CountMessage(session)
		//fmt.Printf("CountMessage Code: %d, Msg: %s, Data: %d\n", retCM.Code, retCM.Msg, retCM.Data)

		if retCM.Data > 0 {
			retMsg := GetMessage(2, session, retCM.Data)
			for _, el := range retMsg.Data {
				fmt.Println(el)
				// 判断如果是私聊并且信息内容为 exit 则退出程序并释放session
				if el.Type == model.TypeFriendMessage && len(el.MessageChain) == 2 {
					if el.MessageChain[1].Type == model.MsgChainTypePlain && el.MessageChain[1].Text == "exit" {
						fmt.Println("exit.")
						return
					}
				}
				if el.Type == model.TypeGroupMessage {
					chain := el.MessageChain
					var msg = ""
					for _, item := range chain {
						if item.Type == model.MsgChainTypeImage {
							msg += "[图片]"
						} else if item.Type == model.MsgChainTypeAt {
							msg += item.Display
						} else if item.Type == model.MsgChainTypePlain {
							msg += item.Text + " "
						}
					}
					fmt.Printf("[%s] %s: %s\n", el.Sender.Group.Name, el.Sender.MemberName, msg)
					if msg == "精华测试 " {
						mid := el.MessageChain[0].Id
						retX := SetEssence(session, mid)
						fmt.Println("set essence: ", retX, mid)
					}
				}
				if el.Type == model.TypeNewFriendRequestEvent {
					friend, err := NewFriendEvent(session, el, 1, "肯定要先拒绝在同意啊")
					fmt.Println("[event] new friend: ", friend, err)
				}
				if el.Type == model.TypeMemberJoinRequestEvent {
					group, err := MemberJoinRequestEvent(session, el, 1, "肯定要先拒绝在同意啊")
					fmt.Println("[event] new member: ", group, err)
				}
				if el.Type == model.TypeBotInvitedJoinGroupRequestEvent {
					group, err := BotInvJoinGroupEvent(session, el, 1, "肯定要先拒绝在同意啊")
					fmt.Println("[event] BotInvitedJoinGroupRequestEvent: ", group, err)
				}
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}

}
