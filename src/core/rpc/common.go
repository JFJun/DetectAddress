package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type RpcServer struct {
	IsStartRpc int
	Port string
}

const SubmitData = "/api/submit"
const SD  =len(SubmitData)

type RespVO struct {
	Code int			`json:"code"`
	Msg string			`json:"message"`
}

func (rpc *RpcServer)HttpHandler(w http.ResponseWriter,r *http.Request){
	switch  {
	case len(r.RequestURI)>=SD&&r.RequestURI[:SD]==SubmitData:
		//处理该请求
		subHandler(w,r,sdRouteMap)
	default:
		reqError(w,r)

	}
}

//反射获取路由方法
func subHandler(w http.ResponseWriter, req *http.Request, routeMap map[string]interface {}){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT ,DELETE")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	uri := strings.Split(req.RequestURI, "?")
	if len(uri) == 2 {
		req.RequestURI = uri[0]
		req.ParseForm()
	}

	for route,handle :=range routeMap{
		reqGrp := strings.Split(route, ":")
		if len(reqGrp) != 2 {
			continue
		}
		method := reqGrp[0]  //请求的方法
		path := reqGrp[1]    //请求的路径
		//正则匹配，如果请求的路径不在path中，返回
		re := regexp.MustCompile(path)
		if !re.MatchString(req.RequestURI) {
			continue
		}

		if strings.ToUpper(req.Method) ==method{ //请求方法匹配
			//利用反射调用方法
			a:= reflect.ValueOf(handle).Call([]reflect.Value{
				reflect.ValueOf(w),reflect.ValueOf(req),
			})
			w.Write(a[0].Bytes())
			return
		}
	}
	reqError(w,req)
}

func reqError(w http.ResponseWriter,r *http.Request){
	var resp RespVO
	resp.Code = 404
	resp.Msg = fmt.Sprintf("获取接口失败 【%s】",r.RequestURI)
	respJSON , _:= json.Marshal(resp)
	w.Write(respJSON)
}