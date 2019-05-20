package config

type Config struct {
	BtcUrl string
	BtcDecimal float64
	BitcnyUrl string
	BitcnyDecimal float64
	EosUrl string
	Eosdecimal float64
	EthUrl string
	EthDecimal float64
	TrxUrl string
	TrxDecimal float64
	ServerUrl string
}

func NewDefaultConfig()*Config{
	config:=&Config{
		BitcnyUrl:"http://185.208.208.184:5000",
		BtcDecimal:0,
		BtcUrl:"https://chain.api.btc.com/v3/",
		BitcnyDecimal:0,
		EthUrl:"",
		EthDecimal:0,
		Eosdecimal:0,
		EosUrl:"",
		TrxDecimal:0,
		TrxUrl:"https://apilist.tronscan.org/api/",
		ServerUrl:"http://127.0.0.1:6567/api/submit",
	}

	return config
}

func NewConfig()*Config{
	cfg:=NewDefaultConfig()
	cfg.ServerUrl = getString("server.url",cfg.ServerUrl)
	cfg.BtcUrl = getString("btc.url",cfg.BtcUrl)
	cfg.BitcnyUrl = getString("bitcny.url",cfg.BitcnyUrl)
	cfg.EosUrl = getString("eos.url",cfg.EosUrl)
	cfg.EthUrl = getString("eth.url",cfg.EthUrl)
	cfg.TrxUrl = getString("trx.url",cfg.TrxUrl)

	cfg.BtcDecimal = getFloat64("btc.coinDecimal",cfg.BtcDecimal)
	cfg.BitcnyDecimal = getFloat64("bitcny.coinDecimal",cfg.BitcnyDecimal)
	cfg.Eosdecimal = getFloat64("eos.coinDecimal",cfg.Eosdecimal)
	cfg.EthDecimal = getFloat64("eth.coinDecimal",cfg.EthDecimal)
	cfg.TrxDecimal = getFloat64("trx.coinDecimal",cfg.TrxDecimal)


	return cfg
}