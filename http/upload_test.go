package http

import (
	"fmt"
	"github.com/toorux/go-mirai-http/conf"
	"testing"
)

func TestUploadImage(t *testing.T) {

	ret := Verify(conf.VerifyKey)
	fmt.Printf("Verify Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
	session := ret.Session

	retBind := Bind(session, conf.QQ)
	fmt.Printf("Bind Code: %d, Msg: %s\n", retBind.Code, retBind.Msg)

	defer func() {
		retRelease := Release(session, conf.QQ)
		fmt.Printf("Release Code: %d, Msg: %s\n", retRelease.Code, retRelease.Msg)
	}()

	retImg, err := UploadImage(session, UploadTypeFriend, conf.ImgPath)
	fmt.Println("image: ", retImg, err)

	//retFile, err := UploadFile(session, conf.ImgPath, conf.Group, "")
	//fmt.Println("image: ", retFile, err)
}
