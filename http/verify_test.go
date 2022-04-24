package http

import (
	"fmt"
	"github.com/kgysf/go-mirai-http/conf"
	"testing"
)

func TestVerify(t *testing.T) {
	ret := Verify(conf.VerifyKey)
	fmt.Printf("Verify Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
	session := ret.Session

	retBind := Bind(session, conf.QQ)
	fmt.Printf("Bind Code: %d, Msg: %s\n", retBind.Code, retBind.Msg)

	retRelease := Release(session, conf.QQ)
	fmt.Printf("Release Code: %d, Msg: %s\n", retRelease.Code, retRelease.Msg)
}
