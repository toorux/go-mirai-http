package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMessageChain(t *testing.T) {
	//a := SourceMsgChain{
	//	Type: "Source",
	//	Id:   45665656,
	//	Time: 6546546,
	//}
	//b, ok := a.(MessageChain)

	a := MessageChain{Type: "Plain", Text: "asdasdasdasdasd"}
	v, err := json.Marshal(a)
	fmt.Println(string(v), err)

}
