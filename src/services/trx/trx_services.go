package trx

import (
	"api/trx_api"
	"config"
	"github.com/Sirupsen/logrus"
	"math/rand"
	"services/util"
	"time"
)

type Trx struct {
	Run bool
	api *trx_api.TrxApiClient
	Txs chan map[string]interface{}
	Addresses []string  //地址集合，需要检测的地址
	CurrentBlockHeight int64

}
var address = []string{}

func init(){
	//读取比特币的配置文件
	address =util.ReadCoinConfig("trx")
}


func New(cfg *config.Config)*Trx{
	return &Trx{
		Run:true,
		api:trx_api.New(cfg),
		Txs:make(chan map[string]interface{}),
		Addresses:address,
		CurrentBlockHeight:0,
	}
}


func (trx *Trx)SubscribeBlock() {
	logrus.Println("开启扫描Trx最新的交易信息")
	for trx.Run{
		height,timestamp:=trx.api.GetLatestBlockHeight()
		if height==-1 ||timestamp==-1{
			logrus.Error("获取Trx最新区块高度错误，再次请求数据")
			n:=rand.Intn(5)
			time.Sleep(time.Duration(n)*time.Second)
			continue
		}
		datas,err:=trx.api.GetTransactionInfo(int64(timestamp))
		if err != nil {
			logrus.Error("获取Trx最新交易数据错误")
		}
		if len(datas)>0{
			for _,data :=range datas{
				if data.ContractData.Amount >0{
					for _,detectAddress:=range trx.Addresses{
						if detectAddress == data.OwnerAddress{
							//转出
							inputInfo:=make( map[string]interface{})
							inputInfo["block_height"] = height
							inputInfo["block_time"] = util.Timestamp(int64(timestamp))
							inputInfo["address"] = data.OwnerAddress
							inputInfo["amount"] = data.ContractData.Amount
							inputInfo["symbol"]= 1
							trx.Txs<-inputInfo
							//TODO 存入数据库

						}
						if detectAddress == data.ToAddress{
							//转入
							inputInfo:=make( map[string]interface{})
							inputInfo["block_height"] = height
							inputInfo["block_time"] = util.Timestamp(int64(timestamp))
							inputInfo["address"] = data.ToAddress
							inputInfo["amount"] = data.ContractData.Amount
							inputInfo["symbol"]= 0
							trx.Txs<-inputInfo
							//TODO 存入数据库

						}
					}
				}
			}
		}else {
			logrus.Errorf("没有区块信息，请检查高度 【%s】",height)
		}
		time.Sleep(10*time.Second)
	}
}