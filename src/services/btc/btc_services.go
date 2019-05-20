package btc

import (
	"api/btc_api"
	"config"
	"github.com/Sirupsen/logrus"
	"math/rand"

	"services/util"

	"time"
)

//比特币服务

type BTC struct {
	Run bool
	Btc_Api *btc_api.ApiClinet
	Txs chan map[string]interface{}
	Addresses []string  //地址集合，需要检测的地址
	CurrentBlockHeight int64

}
var address = []string{}

func init(){
	//读取比特币的配置文件
	address =util.ReadCoinConfig("btc")
	logrus.Println("检测的比特币地址为： ",address)
}


func New(cfg *config.Config)*BTC{
	return &BTC{
		Run:true,
		Btc_Api:btc_api.New(cfg),
		Txs:make(chan map[string]interface{}),
		Addresses:address,
		CurrentBlockHeight:0,
	}
}

//开始服务
func (btc *BTC)SubscribeBlock(){
	logrus.Println("开启扫描BTC最新的交易信息")
	for btc.Run{
		resp,err:=btc.Btc_Api.GetLatestTransaction()
		if err!=nil{
			//发生错误，休眠一段时间后继续扫描
			n:=rand.Intn(15)
			time.Sleep(time.Duration(n)*time.Second)
			continue
		}

		if resp!=nil{
			if resp.Data.List[0].BlockHeight !=btc.CurrentBlockHeight{
				logrus.Printf("处理区块高度为【%d】的交易",resp.Data.List[0].BlockHeight)
				if len(resp.Data.List)!=0{
					for _,tx:=range resp.Data.List  {
						if !tx.IsCoinbase{

							for _,detectAddress:=range btc.Addresses{
								//输入集合
								if len(tx.Inputs)!=0{
									for _,input:=range tx.Inputs {
										if len(input.PrevAddresses)!=0{
											address:=input.PrevAddresses[0]
											//当地址等于我想要查看的地址时
											if address==detectAddress{
												inputInfo:=make( map[string]interface{})
												inputInfo["hash"] = tx.Hash
												inputInfo["block_height"] = tx.BlockHeight
												inputInfo["block_time"] = util.Timestamp(tx.BlockTime)
												inputInfo["address"] = address
												inputInfo["amount"] = input.PrevValue
												inputInfo["symbol"]= 0
												logrus.Println("发送的数据为：",inputInfo)
												btc.Txs<-inputInfo
												//dao.InsertBtc(inputInfo)
											}
										}
									}
								}
								//输出集合
								if len(tx.Outputs)!=0{
									for _,output:=range tx.Outputs{
										if len(output.Addresses)!=0 {
											address:=output.Addresses[0]
											if address==detectAddress{
												inputInfo:=make( map[string]interface{})
												inputInfo["hash"] = tx.Hash
												inputInfo["block_height"] = tx.BlockHeight
												inputInfo["block_time"] = util.Timestamp(tx.BlockTime)
												inputInfo["address"] = address
												inputInfo["amount"] = output.Value
												inputInfo["symbol"]= 1
												logrus.Println("发送的数据为：",inputInfo)
												btc.Txs<-inputInfo
												//dao.InsertBtc(inputInfo)
											}
										}
									}
								}
							}
						}
					}
				}else {
					logrus.Println("没有数据")
				}
			}
			btc.CurrentBlockHeight=resp.Data.List[0].BlockHeight
		}else {
			logrus.Error("解析最新的交易数据错误")
		}
		//处理玩一笔交易，休眠5分钟
		logrus.Println("处理完一笔交易，歇息5分钟")
		time.Sleep(5*time.Minute)
	}
}

//处理区块数据
func (btc *BTC)DealWithBlock(){
	for{
		//主线程，用来处理返回的数据
		tx:=<-btc.Txs
		if tx["symbol"].(int)==0{
			//这是一笔比特币的转出交易

		}
		if tx["symbol"].(int)==1 {
			//这是一笔比特币的转入交易
		}
	}
}
