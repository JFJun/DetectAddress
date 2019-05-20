package btc_api

import (
	"config"
	"encoding/json"
	"fmt"
	"strconv"
	"util"
)

//获取数据

type ApiClinet struct {
	url string
}

func New(cfg *config.Config)*ApiClinet{
	return &ApiClinet{
		url:cfg.BtcUrl,
	}
}

//获取最新的区块
func (api *ApiClinet)GetLatestBlock(){
	apiPath:="block/latest"

	resp,err:=util.HttpGet(api.url+apiPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func (api *ApiClinet)GetBlockByHeight(height int64){
	apiPath:= "block/%s"
	h:=strconv.FormatInt(height,10)
	a:=fmt.Sprintf(apiPath,h)
	resp,err:=util.HttpGet(api.url+a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

func (api *ApiClinet)GetLatestTransaction()(*RespLatestTrans,error){
	apiPath:="block/latest/tx"
	resp,err:=util.HttpNewGet(api.url+apiPath)
	if err != nil {
		return nil,err
	}
	response:=RespLatestTrans{}
	err1:=json.Unmarshal(resp,&response)
	if err1 != nil {
		return nil,err1
	}

	return &response,nil
}