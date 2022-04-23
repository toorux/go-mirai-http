package go_mirai_http

import "sync"

var SessionKey string
var Host string = "localhost"
var Port int = 8080

// Init 初始化host与int
func Init(host string, port int) {
	lock := sync.Mutex{}
	lock.Lock()
	Host, Port = host, port
	lock.Unlock()
}

// SetSessionKey 设置全局SessionKey
func SetSessionKey(sessionKey string) {
	lock := sync.Mutex{}
	lock.Lock()
	SessionKey = sessionKey
	lock.Unlock()
}
