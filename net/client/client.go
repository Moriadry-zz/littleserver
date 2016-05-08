package client

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn net.Conn
}

func (c *Client) Connect(addr string) {
	vals := strings.Split(addr, ":")
	if len(vals) != 2 {
		fmt.Println("err: wrong addr without :")
		return
	}
	var err error
	c.conn, err = net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("connect error,", err)
		return
	}
}

func (c *Client) Write(buf []byte) {
	_, err := c.conn.Write(buf)
	if err != nil {
		fmt.Println("write error,", err)
		return
	}
}

func (c *Client) Read(buf []byte) (n int) {
	var err error
	n, err := c.conn.Read(buf)
	if err != nil {
		fmt.Println("read error,", err)
		return
	}
}

func (c *Client) Close() {
	c.conn.Close()
}
