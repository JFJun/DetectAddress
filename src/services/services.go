package services

import (
	"config"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"services/bitcny"
	"services/btc"
	"services/trx"
	"util"
)

type Serviceser interface {
	SubscribeBlock()
}

type Services struct {
	Cfg *config.Config
	transDataChan chan map[string]interface{}
	Btc *btc.BTC
	Bitcny *bitcny.Bitcny
	Trx *trx.Trx
}

func New(config *config.Config)*Services{
	services:=&Services{
		Cfg :config,
		transDataChan:make(chan map[string]interface{}),
	}
	//注册需要扫描的币
	services.Btc = btc.New(config)
	services.Bitcny = bitcny.New(config)
	services.Trx = trx.New(config)

	return services
}

//开始监听
func (service *Services)Start(coin string){
	switch coin {
	case "btc":
		logrus.Println("开启比特币服务")
		go service.Btc.SubscribeBlock()
		go service.sendBtc()

	case "bitcny":
		logrus.Println("开启Bitcny服务")
		go service.Bitcny.SubscribeBlock()
	case "trx":
		logrus.Println("开启Trx服务")
		go service.Trx.SubscribeBlock()
	default:
		break
	}
	//发送数据到服务器
	go service.SendDataToServer()
}

type ReqVo struct {
	Code int			`json:"code"`
	Msg string			`json:"message"`
}
func (service *Services)SendDataToServer(){
	for{
		data:=<-service.transDataChan
		logrus.Printf("发送给服务器的数据为： 【%v】",data)
		b,err:=json.MarshalIndent(data,""," ")
		if err != nil {
			logrus.Errorf("转化为json 数据错误 Error=【%v】",err)
		}
		resp,err:=util.HttpPost(service.Cfg.ServerUrl,string(b))
		if err != nil {
			logrus.Errorf("请求服务器接口错误，提交数据 Error = 【%v】",err)
		}
		var req ReqVo
		err =json.Unmarshal(resp,&req)
		if err != nil {
			logrus.Errorf("解压json 数据错误 Error=【%v】",err)
		}
		if req.Code!=200 || req.Msg!="Success"{
			logrus.Errorf("************提交数据出错了，请尽快查看,原因====【%s】",req.Msg)
		}
	}
}

func (service *Services)sendBtc(){
	for{
		tx:=<-service.Btc.Txs
		tx["coin"] = "btc"
		println("services接受到了数据，准备发送给服务器")
		service.transDataChan<-tx
	}
}
