package http

import (
	"fmt"
	"github.com/kgysf/go-mirai-http/conf"
	"testing"
)

func TestAnnoList(t *testing.T) {
	ret := Verify(conf.VerifyKey)
	fmt.Printf("Verify Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
	session := ret.Session

	retBind := Bind(session, conf.QQ)
	fmt.Printf("Bind Code: %d, Msg: %s\n", retBind.Code, retBind.Msg)

	defer func() {
		retRelease := Release(session, conf.QQ)
		fmt.Printf("Release Code: %d, Msg: %s\n", retRelease.Code, retRelease.Msg)
	}()

	retX := AnnoList(session, conf.Group, 0, 10)
	fmt.Println(retX)
	for _, anno := range retX.Data {
		fmt.Println("content: ", anno.Content)
		retDelete := AnnoDelete(session, conf.Group, anno.Fid)
		fmt.Println("delete: ", retDelete)
	}

	//retX := AnnoPublish(session, conf.Group, "我就是说， 首先你们不能骂人， 然后不能打架", "", false, false, false, true)
	//fmt.Println(retX)

}
