package main

import (
	"fmt"
	"net"
)

var limitChan = make(chan bool, 1000)

// UDP goroutine 实现并发读取UDP数据
func udpProcess(conn *net.UDPConn) {

	// 最大读取数据大小
	data := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed read udp msg, error: " + err.Error())
	}
	str := string(data[:n])
	fmt.Println("receive from client, data:" + str)
	<-limitChan
}

func udp() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 8080})
	if err != nil {
		fmt.Println(err)
	}
	for {
		limitChan <- true
		go udpProcess(conn)
	}

}

func tcp() {
	l, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go func() {
			for {
				data := make([]byte, 1024)
				n, _ := conn.Read(data)
				str := string(data[:n])
				fmt.Println("receive from client, data:" + str)
			}
		}()

	}
}

func main() {
	tcp()

}
