package core_services

import (
	"core/entities"
	"core/util"
)

var SubmitDataChan = make(chan entities.ResqData)


//定义枚举：用于观察者分发消息
const (
	DESTORY = iota
	STOP
	NONE
	CREATE
	INIT
	START
)

type BaseServices struct {
	status *util.Status
}
func(service *BaseServices) create(){
	service.status = util.NewStatus([]int{DESTORY,STOP,NONE,CREATE,INIT,START})
	//service.status.Init([]int{DESTORY,STOP,NONE,CREATE,INIT,START})
}
func (service *BaseServices) Init() {
	service.status.TurnTo(INIT)
}
func (service *BaseServices) Start() {
	service.status.TurnTo(START)
}
func (service *BaseServices) Stop() {
	if service.status.Current() == START {
		service.status.TurnTo(STOP)
	} else {
		service.status.TurnTo(DESTORY)
	}
}

func (service *BaseServices) IsInit() bool {
	return service.status.Current() >= INIT
}

func (service *BaseServices) IsDestroy() bool {
	return service.status.Current() == DESTORY
}

func (service *BaseServices) CurrentStatus() int {
	return service.status.Current()
}