package core_services

import (
	"core/util"
	"fmt"
	"github.com/Sirupsen/logrus"
	"sync"
)

//提交数据的服务

type SubmitData struct {
	BaseServices
	sync.Once
}

var sd  *SubmitData

func GetSubmitData()*SubmitData{
	if sd ==nil{
		sd = new(SubmitData)
		sd.Once  = sync.Once{}
		sd.Once.Do(func() {

			sd.create()

		})
	}
	return sd
}

func (sd *SubmitData)create(){

	sd.BaseServices.create()

	defer sd.status.RegObs(sd) //注册为观察者
}

func (sd *SubmitData)BeforeTurn(status *util.Status, tgtStt int){
	switch tgtStt {
	case INIT:
		//TODO进行相关的初始化
		fmt.Println("init")
	case START:
		logrus.Println("【启动提交数据服务】")
	}
}
func (sd *SubmitData)AfterTurn(status *util.Status, tgtStt int){
	switch status.Current() {
	case INIT:

		fmt.Println("Init提交数据服务完成")
	case START:

		go sd.DealWithData()  //处理数据

	}

}

func (sd *SubmitData)DealWithData(){

	for sd.status.Current() == START{
		//从缓冲中拿取数据
		resp:=<-SubmitDataChan
		logrus.Printf("处理数据为：【%v】 ",resp)
		switch resp.Coin {
		case "btc":
			//TODO 处理币种的数据
			logrus.Println("开始对btc数据处理")
		case "bitcny":
			logrus.Println("开始对btc数据处理")
		case "trx":
			logrus.Println("开始对btc数据处理")
		}

	}
}