package bitcny_api


type Resp_Bitcny_Block struct {
	Transactions []RespBitcnyTransactions		`json:"transactions"`
}

type RespBitcnyTransactions struct {
	Operations [][]interface{}			`json:"operations"`
	Expiration string						`json:"expiration"`
}

type RespBitcnyData struct {
	Amount RespBitcnyAmount					`json:"amount"`
	From string								`json:"from"`
	To string								`json:"to"`
}
type RespBitcnyAmount struct {
	BitcnyAmount float64					`json:"amount"`
	AssetId string							`json:"asset_id"`
}

