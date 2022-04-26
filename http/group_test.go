package http

import (
	"fmt"
	"github.com/kgysf/go-mirai-http/conf"
	"testing"
)

func TestMute(t *testing.T) {
	ret := Verify(conf.VerifyKey)
	fmt.Printf("Verify Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
	session := ret.Session

	retBind := Bind(session, conf.QQ)
	fmt.Printf("Bind Code: %d, Msg: %s\n", retBind.Code, retBind.Msg)

	defer func() {
		retRelease := Release(session, conf.QQ)
		fmt.Printf("Release Code: %d, Msg: %s\n", retRelease.Code, retRelease.Msg)
	}()

	//retMute := Mute(session, conf.Group, conf.MuteQQ, 60)
	//fmt.Println("mute: ", retMute)
	//
	//time.Sleep(time.Millisecond * 1000)
	//
	//retUnmute := Unmute(session, conf.Group, conf.MuteQQ)
	//fmt.Println("unmute: ", retUnmute)

	//retX := Kick(session, conf.Group, conf.MuteQQ, "滚")
	//fmt.Println("kick: ", retX)

	//retX := Quit(session, conf.Group)
	//fmt.Println("quit: ", retX)

	//retMute := MuteAll(session, conf.Group)
	//fmt.Println("mute: ", retMute)
	//
	//time.Sleep(time.Millisecond * 1000)
	//
	//retUnmute := UnmuteAll(session, conf.Group)
	//fmt.Println("unmute: ", retUnmute)

	//retX, err := GetGroupConfig(session, conf.Group)
	//fmt.Println("get config 1", retX, " |err: ", err)
	//
	//retX.ConfessTalk = !retX.ConfessTalk
	//
	//retY := SetGroupConfig(session, conf.Group, retX)
	//fmt.Println("set config ", retY)

	retX, err := GetMemberInfo(session, conf.Group, conf.FriendQQ)
	fmt.Println("get member ", retX, " |err: ", err)

	retY := SetMemberInfo(session, conf.Group, conf.MuteQQ, "")
	fmt.Println("set member ", retY)
}
