package http

import (
	"fmt"
	"github.com/kgysf/go-mirai-http/conf"
	. "github.com/kgysf/go-mirai-http/model"
	"testing"
	"time"
)

func TestSendFriendMessage(t *testing.T) {

	ret := Verify(conf.VerifyKey)
	fmt.Printf("Verify Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
	session := ret.Session

	retBind := Bind(session, conf.QQ)
	fmt.Printf("Bind Code: %d, Msg: %s\n", retBind.Code, retBind.Msg)

	defer func() {
		retRelease := Release(session, conf.QQ)
		fmt.Printf("Release Code: %d, Msg: %s\n", retRelease.Code, retRelease.Msg)
	}()

	var msg = []MessageChain{
		MessageChain{Type: MsgChainTypePlain, Text: "哪个还可以"},
	}
	var retSend SendMsgResult

	retSend = SendFriendMessage(session, conf.FriendQQ, msg, 0)
	fmt.Println("send friend msg: ", retSend)

	//retSend = SendGroupMessage(session, conf.Group, msg, 0)
	//fmt.Println("send group msg: ", retSend)

	time.Sleep(time.Millisecond * 1000)

	//retSend = SendGroupMessage(session, conf.Group, msg, retSend.MessageId)
	//fmt.Println("send group quote msg: ", retSend)

	//retSend = SendTempMessage(session, conf.Group, conf.FriendQQ, msg, 0)
	//fmt.Println("send temp msg: ", retSend)

	var isOk HttpResult[any]
	//isOk = SendNudgeFriend(session, conf.FriendQQ)
	//fmt.Println("send nudge friend: ", isOk)

	//isOk = SendNudgeMember(session, conf.Group, conf.FriendQQ)
	//fmt.Println("send nudge friend: ", isOk)

	isOk = Recall(session, retSend.MessageId)
	fmt.Println("send recall: ", isOk)

}
