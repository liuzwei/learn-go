package socket

import (
	log "github.com/sirupsen/logrus"
	"net"
	"testing"
)

func Test_process(t *testing.T) {
	// 创建一个监听服务
	listen, err := net.Listen("tcp", "10.56.180.181:20230")
	if err != nil {
		log.Error("Open socket failed!")
		return
	}
	// 创建连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error("Accept and create Connection failed!")
		}
		// 处理连接
		go process(conn)
	}
}
