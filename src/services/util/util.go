package util

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

//获取配置文件

func ReadCoinConfig(coin string)[]string{
	viper.SetConfigName(coin)
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		logrus.Errorf("读取%s配置文件失败：Error=【%v】",coin,err)
	}
	address:=getCoinAddress()
	return address
}

func getCoinAddress()[]string{
	num:=viper.GetInt("num")
	addresses:=[]string{}
	for i:=1;i<=num;i++{
		address:=viper.GetString(fmt.Sprintf("address.%s",strconv.FormatInt(int64(i),10)))
		addresses = append(addresses,address)
	}
	return addresses
}

func Timestamp(seconds int64) string {
	var timelayout = "2006-01-02 T 15:04:05.000"  //时间格式

	datatimestr := time.Unix(seconds,0).Format(timelayout)

	return datatimestr
}