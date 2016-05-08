package server

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	conn    net.Conn
	Handler EventHandler
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println(conn.(*net.Conn).RemoteAddr())
	buf := make([]byte, 1024)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error", err)
			break
		}
		conn.write([]byte("back info"))

	}
}

func (s *Server) Listen(port int) error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		fmt.Println("listen error", err)
		return err
	}
	go func() {
		for {
			c, err := listener.Accept()
			if err != nil {
				fmt.Println("accept error", err)
				break
			}
			go handleConn(c)
		}
	}()
	return nil
}
