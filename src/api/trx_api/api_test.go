package trx_api

import (
	"config"
	"fmt"
	"github.com/Sirupsen/logrus"
	"testing"
)
var (
	cfg=config.NewDefaultConfig()
	api=New(cfg)
)

func TestTrxApiClient_GetLatestBlockHeight(t *testing.T) {
	api.GetLatestBlockHeight()
}

func TestTrxApiClient_GetBlock(t *testing.T) {
	height,_:=api.GetLatestBlockHeight()

	api.GetBlock(height)
}

func TestTrxApiClient_GetTransactionInfo(t *testing.T) {
	_,time:=api.GetLatestBlockHeight()
	data,err:=api.GetTransactionInfo(int64(time))
	if err != nil {
		logrus.Error(err)
	}
	for _,da:=range data{
		fmt.Println(da)
	}
}