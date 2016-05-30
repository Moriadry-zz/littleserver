package client

import (
	"fmt"
	"testing"
)

func TestErrorAddr(t *testing.T) {
	client := &Client{}
	client.Connect("127.0.0.1")
}

func TestReadWrite(t *testing.T) {
	client := &Client{}
	defer client.Close()
	client.Connect("127.0.0.1:12222")
	fmt.Println([]byte("asd"))
	client.Write([]byte("asd"))
	buf := make([]byte, 1024)
	n := client.Read(buf)
	fmt.Printf("get %d bytes\n", n)
	fmt.Println(string(buf[:n]))
}
