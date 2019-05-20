package bitcny_api

import (
	"config"
	"testing"
)

var (
	cfg=config.NewDefaultConfig()
	api=New(cfg)
)

func TestBitCnyApiClient_GetLatestBlockHeight(t *testing.T) {
	api.GetLatestBlockHeight()
}

func TestBitCnyApiClient_GetBlock(t *testing.T) {
	height:="37423812"
	api.GetBlock(height)

}