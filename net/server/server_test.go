package server

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println("test")
	server := &Server{}
	server.Listen(12222)
}
