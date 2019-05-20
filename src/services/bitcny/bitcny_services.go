package bitcny

import (
	"api/bitcny_api"
	"config"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"services/util"
	"strconv"
	"time"
)

type Bitcny struct {
	run bool
	api *bitcny_api.BitcnyApiClient
	Txs chan map[string]interface{}
	CurrentBlockHeight int64
	//检测的地址
	Addresses []string
}

var address = []string{}

func init(){
	//读取比特币的配置文件
	address =util.ReadCoinConfig("bitcny")
}


func New(cfg *config.Config)*Bitcny{
	return &Bitcny{
		run:true,
		api:bitcny_api.New(cfg),
		Txs:make(chan map[string]interface{}),
		Addresses:address,
	}
}

func (bitcny *Bitcny)SubscribeBlock(){
	logrus.Println("开启扫描Bitcny最新的交易信息")
	for bitcny.run{
		//获取最新高度
		heightStr:=bitcny.api.GetLatestBlockHeight()
		if heightStr==""{
			time.Sleep(1*time.Second)
			continue
		}
		height,_:=strconv.ParseInt(heightStr,10,64)
		if height!=bitcny.CurrentBlockHeight{
			txs,err:=bitcny.api.GetBlock(heightStr)
			if err != nil {
				logrus.Error(err)
			}
			if txs!=nil{
				if len(txs)!=0{

					for _,tx:=range txs{
						createTime:=tx.Expiration  //交易创造时间
						b,err:=json.MarshalIndent(tx.Operations[0][1],""," ")
						if err != nil {
							logrus.Errorf("解析bitcny交易错误【%v】",err)
						}
						response:=bitcny_api.RespBitcnyData{}
						err = json.Unmarshal(b,&response)
						if err != nil {
							logrus.Errorf("解析bitcny交易错误【%v】",err)
						}
						//简单的判断这只一笔bitcny交易
						if response.From!=""&&response.To!=""{
							for _,detectAddress:=range bitcny.Addresses{
								if detectAddress==response.From{
									//这是一笔转出交易
									inputInfo:=make( map[string]interface{})
									inputInfo["block_height"] = height
									inputInfo["block_time"] = createTime
									inputInfo["address"] = response.From
									inputInfo["amount"] = response.Amount.BitcnyAmount
									inputInfo["symbol"] = 1
									bitcny.Txs<-inputInfo
								}
								if detectAddress==response.To{
									//这是一笔转入交易
									inputInfo:=make( map[string]interface{})
									inputInfo["block_height"] = height
									inputInfo["block_time"] = createTime
									inputInfo["address"] = response.To
									inputInfo["amount"] = response.Amount.BitcnyAmount
									inputInfo["symbol"] = 0
									bitcny.Txs<-inputInfo
								}
							}
						}
					}
				}
			}
			bitcny.CurrentBlockHeight = height
		}
		logrus.Println("处理完一笔交易，歇息10秒")
		//休眠10秒
		time.Sleep(10*time.Second)
	}
}
func (bitcny *Bitcny)DealWithBlock(){

}