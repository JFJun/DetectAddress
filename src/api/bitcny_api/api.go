package bitcny_api

import (
	"config"
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"strings"
	"util"
)

type BitcnyApiClient struct {
	url string
}

func New(cfg *config.Config)*BitcnyApiClient{
	return &BitcnyApiClient{
		url:cfg.BitcnyUrl,
	}
}

//获取最新区块高度
func (api *BitcnyApiClient)GetLatestBlockHeight()string{
	apiPath:="/last_block_number"

	resp,err:=util.HttpNewGet(api.url+apiPath)
	if err != nil {
		logrus.Error("获取bitcny最新区块高度错误")
		return ""
	}
	height:=string(resp)
	return strings.Split(height,"\n")[0]
}

//获取最新区块信息
func (api *BitcnyApiClient)GetBlock(height string)([]RespBitcnyTransactions,error){
	apiPath:="/block?"
	v:=util.Value{}
	v.Set("block_num",height)
	//fmt.Println(api.url+apiPath+v.Encode())
	resp,err:=util.HttpNewGet(api.url+apiPath+v.Encode())
	if err != nil {
		logrus.Error("获取bitcny最新区块错误")
		return nil,err
	}
	response:=&Resp_Bitcny_Block{}
	err=json.Unmarshal(resp,response)
	if err != nil {
		logrus.Errorf("解析获取最新的bitcny的区块数据错误【%v】",err)
		return nil ,err
	}

	return response.Transactions,nil
}

//这个方法的写在扫描区块里面
func (api *BitcnyApiClient)ParseTransactions(txs []RespBitcnyTransactions){
	if len(txs)==0{
		//TODO
		return
	}
	for _,tx:=range txs{
		b,err:=json.MarshalIndent(tx.Operations[0][1],""," ")
		if err != nil {
			logrus.Errorf("解析bitcny交易错误【%v】",err)
		}
		response:=RespBitcnyData{}
		err = json.Unmarshal(b,&response)
		if err != nil {
			logrus.Errorf("解析bitcny交易错误【%v】",err)
		}
		fmt.Println(response.From)
	}
}