package http

import (
	"fmt"
	"github.com/kgysf/go-mirai-http/conf"
	"github.com/kgysf/go-mirai-http/model"
	"testing"
)

func TestFileList(t *testing.T) {

	ret := Verify(conf.VerifyKey)
	fmt.Printf("Verify Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
	session := ret.Session

	retBind := Bind(session, conf.QQ)
	fmt.Printf("Bind Code: %d, Msg: %s\n", retBind.Code, retBind.Msg)

	defer func() {
		retRelease := Release(session, conf.QQ)
		fmt.Printf("Release Code: %d, Msg: %s\n", retRelease.Code, retRelease.Msg)
	}()

	retFileList := FileList(session, model.FileParams{
		Id:               "",
		Path:             "",
		Target:           conf.Group,
		WithDownloadInfo: false,
	}, 0, 10)

	if retFileList.Code == 0 {
		for i, file := range retFileList.Data {
			fmt.Printf("file list [%d]: %s;\n", i, file.Name)
			//fmt.Println(file.Contact)
			if file.IsFile {
				ret := FileInfo(session, model.FileParams{Id: file.Id, Target: conf.Group})
				fmt.Println("file info: ", ret)
			}
			if file.Name == "delete.txt" {
				ret := FileDeleteForId(session, file.Id, conf.Group)
				fmt.Println("delete: ", ret)
			}
			if file.Name == "delete.JPG" {
				ret := FileRenameForId(session, file.Id, "ddd.jpg", conf.Group)
				fmt.Println("rename: ", ret)
			}
		}
	}

	//retMkdir := FileMkdir(session, model.FileParams{
	//	Id:               "",
	//	Path:             "",
	//	Target:           conf.Group,
	//	Group:            conf.Group,
	//	Qq:               0,
	//	WithDownloadInfo: false,
	//}, "测试文件夹")
	//fmt.Println("mkdir: ", retMkdir)

}
