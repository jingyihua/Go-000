package api

import (
	"fmt"
	"log"
	"week02/service"
)

type Handler struct {
	Service *service.Service
}

func NewHandler() *Handler {
	return &Handler{
		Service: service.NewService(),
	}
}

func (handler *Handler) GetUserAge(id int) {
	age, err := handler.Service.GetUserAge(id)
	if err != nil {
		//log 记录日志
		log.Printf("ERROR: %+v\n", err)
		backMsg := ErrTranslate(err)
		//返回失败
		fmt.Printf("failed get age: %s\n", backMsg)
		return
	}

	//查询成功 回复信息
	fmt.Printf("success get age: %d\n", age)
}
