package http

import (
	"fmt"
	"github.com/toorux/go-mirai-http/conf"
	"testing"
)

func TestFriendList(t *testing.T) {

	ret := Verify(conf.VerifyKey)
	fmt.Printf("Verify Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
	session := ret.Session

	retBind := Bind(session, conf.QQ)
	fmt.Printf("Bind Code: %d, Msg: %s\n", retBind.Code, retBind.Msg)

	defer func() {
		retRelease := Release(session, conf.QQ)
		fmt.Printf("Release Code: %d, Msg: %s\n", retRelease.Code, retRelease.Msg)
	}()

	retFriend := FriendList(session)
	fmt.Println("friends: ", retFriend)

	retGroup := GroupList(session)
	fmt.Println("groups: ", retGroup)

	retMember := MemberList(session, retGroup.Data[0].Id)
	fmt.Println("members: ", retMember)

	var profile, err = BotProfile(session)
	fmt.Println("profile: ", profile, err)

	profile, err = FriendProfile(session, conf.FriendQQ)
	fmt.Println("friend profile: ", profile, err)

	profile, err = MemberProfile(session, retGroup.Data[0].Id, conf.FriendQQ)
	fmt.Println("member profile: ", profile, err)

	profile, err = UserProfile(session, conf.FriendQQ)
	fmt.Println("user profile: ", profile, err)

}
