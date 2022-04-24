package http

import (
	"fmt"
	"github.com/kgysf/go-mirai-http/conf"
	"github.com/kgysf/go-mirai-http/model"
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
		fmt.Printf("CountMessage Code: %d, Msg: %s, Data: %d\n", retCM.Code, retCM.Msg, retCM.Data)

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
			}
		}
		time.Sleep(time.Millisecond * 1000)
	}

}
