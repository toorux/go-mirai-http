package http

import (
	"fmt"
	"testing"
)

func TestVerify(t *testing.T) {
	ret := Verify("H1OF6SHU41D3F4G")
	fmt.Printf("Code: %d, Msg: %s, Session: %s\n", ret.Code, ret.Msg, ret.Session)
}
