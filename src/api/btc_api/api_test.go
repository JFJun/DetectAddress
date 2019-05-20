package btc_api

import (
	config2 "config"
	"testing"
)

var  (
	cfg= config2.NewDefaultConfig()
	api = New(cfg)
)
func TestApiClinet_GetLatestBlock(t *testing.T) {
	api.GetLatestBlock()
}

func TestApiClinet_GetBlockByHeight(t *testing.T) {
	api.GetBlockByHeight(576881)
}

func TestApiClinet_GetLatestTransaction(t *testing.T) {
	api.GetLatestTransaction()
}

