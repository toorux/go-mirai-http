package go_mirai_http

import "sync"

var Host string = "localhost"
var Port int = 8080

// Init 初始化host与int
func Init(host string, port int) {
	lock := sync.Mutex{}
	lock.Lock()
	Host, Port = host, port
	lock.Unlock()
}
