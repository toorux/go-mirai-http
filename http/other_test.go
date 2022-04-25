package http

import (
	"fmt"
	"testing"
)

func TestAbout(t *testing.T) {
	info := About()
	fmt.Println(info)
}
