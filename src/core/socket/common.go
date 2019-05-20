package socket

import "net"

type SockServer struct {
	Port string
	IsStartSock int
}

func (sock *SockServer)SocketHandler(conn net.Conn){

}
