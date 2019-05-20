package btc_api

//解析交易

type RespLatestTrans struct {
	Data RespData			`json:"data"`
}

type RespData struct {
	List []Transaction		`json:"list"`
}

type Transaction struct {
	BlockHeight int64		`json:"block_height"`
	Hash string				`json:"hash"`
	BlockTime	int64		`json:"block_time"`
	IsCoinbase bool			`json:"is_coinbase"`
	IsDoubleSpend bool		`json:"is_double_spend"`
	Inputs []RespInput		`json:"inputs"`
	Outputs []RespOutput	`json:"outputs"`
}

type RespInput struct {
	PrevAddresses []string	`json:"prev_addresses"`
	PrevValue float64		`json:"prev_value"`
}

type RespOutput struct {
	Addresses []string		`json:"addresses"`
	Value float64			`json:"value"`
}