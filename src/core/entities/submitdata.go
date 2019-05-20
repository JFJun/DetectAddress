package entities

type ResqData struct {
	Hash string				`json:"hash"`
	BlockHeight int64		`json:"block_height"`
	BlockTime string		`json:"block_time"`
	Address string			`json:"address"`
	Amount float64			`json:"amount"`
	Symbol int				`json:"symbol"`
	Coin string				`json:"coin"`
}
