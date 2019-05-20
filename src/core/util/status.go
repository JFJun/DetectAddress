package util

import "errors"

//使用观察者模式
type Observer interface {
	BeforeTurn(status *Status, tgtStt int)
	AfterTurn(status *Status, srcStt int)
}


type Status struct {
	currentStatus int
	AllStatus []int
	AllObserver []Observer
}
func  NewStatus(stts []int)*Status {
	stt:=&Status{}
	if stts != nil {
		stt.AllStatus = stts
	}
	return stt
}

//注册为观察者
func (stt *Status)RegObs(obs Observer){
	stt.AllObserver = append(stt.AllObserver, obs)

}
//获取当前状态
func (stt *Status)Current()int{
	return stt.currentStatus
}

func(stt *Status)TurnTo(status int)(int,error){
	//获取当前状态
	orgStt:=stt.currentStatus
	if !IntArrayContains(stt.AllStatus,status){
		return orgStt,errors.New("Cloud not find identified status")
	}
	//如果该状态执行了，就不再执行
	if orgStt ==status{
		return orgStt,nil
	}

	for _,obs :=range stt.AllObserver{
		obs.BeforeTurn(stt,status)
	}
	stt.currentStatus = status
	for _,obs :=range stt.AllObserver{
		obs.AfterTurn(stt,status)
	}
	return status,nil
}

func IntArrayContains(array []int, target int) bool {
	for _, i := range array {
		if i == target { return true }
	}
	return false
}