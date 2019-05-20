package trx_api

import (
	"config"
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"util"
)

type TrxApiClient struct {
	url string
}


func New(cfg *config.Config)*TrxApiClient{
	return &TrxApiClient{
		url:cfg.TrxUrl,
	}
}


func (api *TrxApiClient)GetLatestBlockHeight()(float64,float64){
	apiPath:="block/latest"
	resp,err:=util.HttpGet(api.url+apiPath)
	if err != nil {
		logrus.Errorf("获取Trx最新的区块高度错误【%v】",err)
		return -1,-1
	}
	return resp["number"].(float64),resp["timestamp"].(float64)
}

func (api *TrxApiClient)GetBlock(height float64){
	apiPath:="block?"
	v:=util.Value{}
	v.Set("number",height)
	resp,err:=util.HttpNewGet(api.url+apiPath+v.Encode())
	if err != nil {
		logrus.Errorf("获取Trx最新的区块错误【%v】",err)
	}
	fmt.Println(string(resp))
}

func (api *TrxApiClient)GetTransactionInfo(timestamp int64)([]RespTrxData,error){
	apiPath:="transaction?"
	v:=util.Value{}
	v.Set("sort","-timestamp")
	v.Set("count",true)
	v.Set("limit",10000)
	v.Set("start",0)
	v.Set("start_timestamp",int64(timestamp))
	v.Set("end_timestamp",int64(timestamp))
	resp,err:=util.HttpNewGet(api.url+apiPath+v.Encode())
	if err != nil {
		logrus.Errorf("获取Trx最新的区块错误【%v】",err)
		return nil,err
	}
	response:= RespTrxTransaction{}
	err = json.Unmarshal(resp,&response)
	if err != nil {
		logrus.Errorf("解析Trx交易json数据错误【%v】",err)
		return nil,err
	}
	return response.Data,nil
}