package go_mirai_http

import (
	"fmt"
	"testing"
)

var host = "localhost"
var port = 8080

func TestInit(t *testing.T) {
	Init(host, port)
	fmt.Println("测试初始化结果", Host, Port)
	if Host != host || Port != port {
		t.Errorf("初始化出错！")
	}
}
