package trx_api

type RespTrxTransaction struct {
	Data []RespTrxData						`json:"data"`
}

type RespTrxData struct {
	OwnerAddress string						`json:"ownerAddress"`
	ToAddress string						`json:"toAddress"`
	ContractData RespTrxContractData		`json:"contractData"`
}

type RespTrxContractData struct {	
	Amount float64							`json:"amount"`
}