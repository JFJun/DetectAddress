package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

var(
	Header = []string{
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.63 Safari/537.36",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; .NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729; InfoPath.3; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; rv:38.0) Gecko/20100101 Firefox/38.0",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",

	}
)

func _httpRequest(reqType string,reqUrl string,postData string,requestHeaders map[string]string)([]byte,error){
	req,_:= http.NewRequest(reqType,reqUrl,strings.NewReader(postData))
	//设置随机的请求头
	n:=rand.Intn(7)
	req.Header.Set("User-Agent", Header[n])
	if requestHeaders!=nil{
		for k,v:=range requestHeaders{
			req.Header.Add(k,v)
		}
	}else{
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	client:=http.DefaultClient
	resp,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if resp.StatusCode !=200{
		return  nil,errors.New(fmt.Sprintf("HttpStatusCode:%d ,Desc:%s", resp.StatusCode, resp.Status))
	}
	bodyData,err:= ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return bodyData,nil
}

func HttpNewGet(reqUrl string)([]byte,error){
	return _httpRequest(http.MethodGet,reqUrl,"",nil)
}

func HttpGet(reqUrl string)(map[string]interface{},error){
	respData,err:= _httpRequest(http.MethodGet,reqUrl,"",nil)
	if err != nil {
		return nil,err
	}
	var bodyDataMap map[string]interface{}
	err = json.Unmarshal(respData,&bodyDataMap)
	if err != nil {
		return nil,err
	}
	return bodyDataMap,nil
}
func HttpGetNeedParams( reqUrl string,params Value)([]byte,error){
	respData,err:= _httpRequest(http.MethodGet,reqUrl+"?"+params.Encode(),"",nil)
	fmt.Println("请求地址：" ,reqUrl+"?"+params.Encode())
	if err != nil {
		return nil,err
	}
	return respData,nil
}

func HttpPost(reqUrl ,postData string)([]byte,error){
	headers:= map[string]string{
		"Content-Type": "application/json;charset=UTF-8"}
	return _httpRequest(http.MethodPost,reqUrl,postData,headers)
}


func HttpPostForm(reqUrl string,postData Value)([]byte,error){
	headers := map[string]string{
		"Content-Type": "application/json;charset=UTF-8"}
	params,err:=json.Marshal(postData)
	if err != nil {
		return nil,err
	}
	log.Printf("reqUrl:%s",reqUrl)
	return _httpRequest(http.MethodPost,reqUrl,string(params),headers)
}



func HttpDelete(reqUrl,postData string)([]byte,error){
	headers := map[string]string{
		"Content-Type": "application/json;charset=UTF-8"}
	return _httpRequest(http.MethodDelete,reqUrl,postData,headers)
}


type Value map[string]interface{}

func (v Value)Get(key string)interface{}{
	if v==nil{
		return ""
	}
	return v[key]
}
func (v Value)Set(key string,value interface{}){
	v[key] = value
}

func(v Value) Del(key string){
	delete(v,key)
}
// Encode params to url like "bar=baz&foo=quux"
func (v Value)Encode() string{
	if v == nil{
		return ""
	}

	var buf bytes.Buffer
	keys := make([]string,0,len(v))
	for k:=range v{
		keys = append(keys,k)
	}
	sort.Strings(keys)
	for _,k:=range keys{
		vs := v[k]

		prefix:= url.QueryEscape(k)+"="
		if buf.Len()>0{
			buf.WriteByte('&')
		}
		buf.WriteString(prefix)
		buf.WriteString(url.QueryEscape(fmt.Sprintf("%v",vs)))
	}
	return buf.String()
}


