package p2p

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// This File is only for test

// https://blog.csdn.net/Jailman/article/details/122702807
type PeerTunnel struct {
	Host  net.UDPAddr // host 主机
	Token string      // 链接码 （ token 唯一，不会重复添加）
	// peers []net.UDPAddr // 连接客户端
}

// TODO 立即移除已连接的主机 或定时销毁， 如何处理已注册的主机
var hosts []PeerTunnel // 所有注册的主机

// p2p服务的server
// ip: net.IPv4zero
func StartServerTest(ip net.IP, port int) {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: port})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("本地地址: <%s> \n", listener.LocalAddr().String())
	peers := make([]net.UDPAddr, 0, 2)
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("<%s> %s\n", remoteAddr.String(), data[:n])
		peers = append(peers, *remoteAddr)
		if len(peers) == 2 {
			log.Printf("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
			listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
			listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
			time.Sleep(time.Second * 8)
			log.Println("中转服务器退出,仍不影响peers间通信")
			return
		}
	}
}

// 查找主机端点信息
func findByToken(token string) *PeerTunnel {
	var findH PeerTunnel
	// var err string
	for _, host := range hosts {
		if host.Token == token {
			findH = host
			// err = "no such token"
			break
		}
	}
	return &findH
}

// 添加新的主机端点信息
func addToken(token string, host net.UDPAddr) PeerTunnel {
	// var founded PeerTunnel
	// var existed bool = false
	// for _, host := range hosts {
	// 	if host.Token == token {
	// 		founded = host
	// 		existed = true
	// 		break
	// 	}
	// }
	founded := findByToken(token)

	// if existed {
	if founded != nil { // token 已经存在（覆盖ip）
		founded.Host = host
	} else {
		newPeer := PeerTunnel{host, token}
		hosts = append(hosts, newPeer)
		founded = &newPeer
	}

	return *founded
}

/** 启动中间服务 */
// ip: net.IPv4zero
func StartServer(ip net.IP, port int) {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: ip, Port: port})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Printf("本地地址: <%s> \n", listener.LocalAddr().String())
	// peers := make([]net.UDPAddr, 0, 2)
	data := make([]byte, 1024)
	for {
		// 监听注册服务
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("Receive <%s> %s\n", remoteAddr.String(), data[:n])

		dataStr := string(data[:])
		if strings.HasPrefix(dataStr, "host:") {
			token := dataStr[5:]
			// 这里应该先检查是否存在
			host := findByToken(token)
			if host != nil { // 返回错误信息
				listener.WriteToUDP([]byte("error:Token已存在！Token existed!"), remoteAddr) // 回复失败消息
			} else { // 添加新的主机
				addToken(token, *remoteAddr)
			}
		} else if strings.HasPrefix(dataStr, "client:") {
			token := dataStr[7:]
			host := findByToken(token)

			// if err == "" {
			if host == nil {
				listener.WriteToUDP([]byte("error:Token不存在！Token not existed!"), remoteAddr) // 回复主机不存在
				return
			}

			// 通知双方连接成功
			log.Printf("进行UDP打洞,建立 %s <--> %s 的连接\n", &host.Host, remoteAddr)
			listener.WriteToUDP([]byte(remoteAddr.String()), &host.Host) // 先给host发送client的信息，让其连接
			listener.WriteToUDP([]byte(host.Host.String()), remoteAddr)
			time.Sleep(time.Second * 2)
			log.Println("中转服务器退出,仍不影响peers间通信")
			return
		} else {
			return
		}

		// peers = append(peers, *remoteAddr)
		// if len(peers) == 2 {
		// 	log.Printf("进行UDP打洞,建立 %s <--> %s 的连接\n", peers[0].String(), peers[1].String())
		// 	listener.WriteToUDP([]byte(peers[1].String()), &peers[0])
		// 	listener.WriteToUDP([]byte(peers[0].String()), &peers[1])
		// 	time.Sleep(time.Second * 8)
		// 	log.Println("中转服务器退出,仍不影响peers间通信")
		// 	return
		// }
	}
}
