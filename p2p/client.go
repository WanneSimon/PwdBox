package p2p

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// This File is only for test

var tag string

const HAND_SHAKE_MSG = "NAT MESSAGE"

/** 连接中间服务，并尝试 p2p。 可以封装自动重试功能 */
func TryConnect(localAddr net.UDPAddr, serverAddr net.UDPAddr, isHost bool, token string) {
	// 当前进程标记字符串,便于显示
	tag = os.Args[1]
	// srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 9982} // 注意端口必须固定
	// dstAddr := &net.UDPAddr{IP: net.ParseIP("207.148.70.129"), Port: 9981}
	srcAddr := &localAddr
	dstAddr := &serverAddr

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		// fmt.Println(err)
		log.Panic(err)
	}

	// dataStr := "hello, I'm new peer:" + tag
	dataStr := "server:" + token
	if isHost {
		dataStr = "client:" + token
	}

	// 向中间服务发送 token
	if _, err = conn.Write([]byte(dataStr)); err != nil {
		log.Panic(err)
	}

	// 等待中间服务返回另一个端点的信息
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("error during read: %s", err)
		log.Panic(err)
	}
	conn.Close()

	// 检查中间服务是否返回错误
	reData := string(data[:n])
	if strings.HasPrefix(reData, "error:") {
		log.Panic(reData[6:])
	} else {
		anotherPeer := parseAddr(reData)
		fmt.Printf("local:%s server:%s another:%s\n", srcAddr, remoteAddr, anotherPeer.String())

		// 开始打洞
		bidirectionHole(srcAddr, &anotherPeer)
	}
}

/** 字符串IP转对象 */
func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{
		IP:   net.ParseIP(t[0]),
		Port: port,
	}
}

/** 打洞 */
func bidirectionHole(srcAddr *net.UDPAddr, anotherAddr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", srcAddr, anotherAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	// 向另一个peer发送一条udp消息(对方peer的nat设备会丢弃该消息,非法来源),用意是在自身的nat设备打开一条可进入的通道,这样对方peer就可以发过来udp消息
	if _, err = conn.Write([]byte(HAND_SHAKE_MSG)); err != nil {
		log.Println("send handshake:", err)
	}

	go func() {
		// 发送
		for {
			time.Sleep(3 * time.Second)
			if _, err = conn.Write([]byte("from [" + tag + "]")); err != nil {
				log.Println("send msg fail", err)
			}
		}
	}()

	// 接收
	for {
		data := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			log.Printf("error during read: %s\n", err)
		} else {
			log.Printf("收到数据:%s\n", data[:n])
			log.Println("对方地址：" + addr.IP.String() + ":" + strconv.Itoa(addr.Port))
		}
	}
}
