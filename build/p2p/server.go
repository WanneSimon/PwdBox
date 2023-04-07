package p2p

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// https://blog.csdn.net/Jailman/article/details/122702807
type PeerTunnel struct {
	Host  net.UDPAddr // host 主机
	Token string      // 链接码
	// peers []net.UDPAddr // 连接客户端
}

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

func findByToken(token string) (PeerTunnel, string) {
	var findH PeerTunnel
	var err string
	for _, host := range hosts {
		if host.Token == token {
			findH = host
			err = "no such token"
			break
		}
	}
	return findH, err
}
func addToken(token string, host net.UDPAddr) PeerTunnel {
	var founded PeerTunnel
	var existed bool = false
	for _, host := range hosts {
		if host.Token == token {
			founded = host
			existed = true
			break
		}
	}

	if existed {
		founded.Host = host
	} else {
		newPeer := PeerTunnel{host, token}
		hosts = append(hosts, newPeer)
		founded = newPeer
	}

	return founded
}

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
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		log.Printf("Receive <%s> %s\n", remoteAddr.String(), data[:n])

		dataStr := string(data[:])
		if strings.HasPrefix(dataStr, "host:") {
			token := dataStr[5:]
			addToken(token, *remoteAddr)
		} else if strings.HasPrefix(dataStr, "client:") {
			token := dataStr[7:]
			host, err := findByToken(token)

			if err == "" {
				return
			}

			//TODO host 和 任意 client 建立连接后，其公网 IP 已经暴露出来，后续 client 直接即可
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
