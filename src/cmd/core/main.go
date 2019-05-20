package main

import (
	"core/core_services"
	rpc2 "core/rpc"
	"core/socket"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

var (
	rpc  = new(rpc2.RpcServer)
	rpcPort string
	sockPort string
	sock = new(socket.SockServer)
)
//把所有服务对象放入切片中,使用不安全指针强转
var allServices = []*core_services.BaseServices{
	(*core_services.BaseServices)(unsafe.Pointer(core_services.GetSubmitData())),
}

func initServices(){
	for _,service:=range allServices {
		service.Init()
	}
	for _,service:=range allServices {
		service.Start()
	}
}
func init(){
	//读取配置文件
	viper.SetConfigName("core")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		logrus.Errorf("读取core配置文件失败：Error=【%v】",err)
	}
	readConfig()
}

func readConfig(){
	rpcPort=viper.GetString("rpc.port")

	if rpcPort ==""{
		rpc.Port = "6767"
	}else {
		rpc.Port = rpcPort
	}
	rpc.IsStartRpc =viper.GetInt("rpc.isStartRpc")
	sockPort=viper.GetString("socket.port")
	if sockPort ==""{
		sock.Port = "6768"
	}else {
		sock.Port = sockPort
	}
	sock.IsStartSock =viper.GetInt("socket.isStartSocket")
}

func main(){
	//启动所有的服务
	initServices()
	if rpc.IsStartRpc==1{
		//开启rpc服务
		go func() {
			http.HandleFunc("/",rpc.HttpHandler)
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", rpc.Port), nil))
			logrus.Println("正在安全退出rpc服务器")
		}()
	}

	if sock.IsStartSock==1{
		go func() {
			url := fmt.Sprintf("localhost:%s", sock.Port)
			var socket net.Listener
			var err error
			if socket, err = net.Listen("tcp", url); err != nil {
				log.Fatal(err)
			}
			defer socket.Close()
			for{
				var conn net.Conn
				if conn,err = socket.Accept();err!=nil{
					logrus.Errorf("接收客户端连接错误【%v】",err)
					continue
				}
				//异步处理连接消息
				go sock.SocketHandler(conn)
			}
		}()
	}



	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		logrus.Printf("app-admin get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logrus.Printf("app-admin exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
