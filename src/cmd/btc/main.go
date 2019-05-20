package main

import (
	"config"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"services"
	"syscall"
)

var(
	cfg = config.NewDefaultConfig()
	coin  string
)

func initConfig(){
	//加载配置文件
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		logrus.Errorf("读取配置文件错误，使用默认配置文件:Error=【%v】",err)
	}
	cfg = config.NewConfig()
}

var rootCmd = &cobra.Command{
	Use:"CoinData",
	Short:"Transaction Model",
	Long:"A new transaction model by some important address",
	Run: func(cmd *cobra.Command, args []string) {
		if coin==""{
			fmt.Println("解析指令出错")
			os.Exit(1)
		}
		fmt.Println(coin)
		service:=services.New(cfg)
		service.Start(coin)

	},
}
func init(){
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&coin, "cointype", "c", "btc", "select cointype by ['btc','eth','bitcny','eos','trx']")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//做一个信号处理
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		logrus.Printf("app-admin get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			logrus.Printf("app-admin exit")
			return
		case syscall.SIGHUP:
		// TODO reload
		default:
			return
		}
	}
}
