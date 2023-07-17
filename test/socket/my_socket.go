package socket

import (
	log "github.com/sirupsen/logrus"
	"net"
)

// process handle connection
func process(conn net.Conn) {
	// 最终关闭连接
	defer conn.Close()
	// 构造一个bytes
	bs := make([]byte, 1024)
	var readStr string
	// 读取数据
	for {
		n, err := conn.Read(bs)
		if err != nil {
			log.Error("read failed")
			break
		}
		content := string(bs[:n])

		if content == "exit\r\n" {
			log.Info("read over!")
			break
		}
		readStr = readStr + content
	}
	log.WithFields(log.Fields{
		"readObject": readStr,
	}).Info("数据读取完毕")

	// 写出数据
	conn.Write([]byte("Good Bye\n"))
}
