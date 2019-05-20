package rpc

import (
	"core/core_services"
	"core/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const submitDataPath = "^/api/submit$"

var sdRouteMap = map[string]interface{}{
	fmt.Sprintf("%s:%s",http.MethodPost,submitDataPath):submitData,
}



func submitData(w http.ResponseWriter,req*http.Request)[]byte{
	var resp RespVO
	//获取数据
	var body []byte
	var err error

	if body, err = ioutil.ReadAll(req.Body); err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		ret, _ := json.Marshal(resp)
		return ret
	}
	defer req.Body.Close()
	//解析数据
	var response entities.ResqData
	if err = json.Unmarshal(body, &response); err != nil {
		fmt.Println(222)
		resp.Code = 500
		resp.Msg = err.Error()
		ret, _ := json.Marshal(resp)
		return ret
	}
	if  &response==nil{
		resp.Code = 500
		resp.Msg = "解析提交的数据错误"
		ret, _ := json.Marshal(resp)
		return ret
	}
	//把数据写入缓冲中，交给服务去处理
	core_services.SubmitDataChan<-response
	resp.Code = 200
	resp.Msg = "Success"
	ret, _ := json.Marshal(resp)
	return ret



	//参数判断
	//switch response.Coin {
	//case "btc":
	//	//TODO 进行比特币数据处理
	//	fmt.Println("btc")
	//case "bitcny":
	//	fmt.Println("bitcny")
	//case "trx":
	//	fmt.Println("trx")
	//default:
	//	resp.Code = 500
	//	resp.Msg = "暂未开放该币种或币种名字错误"
	//	ret, _ := json.Marshal(resp)
	//	return ret
	//
	//}

}