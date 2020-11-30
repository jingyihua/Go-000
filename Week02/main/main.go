package main

import (
	"week02/api"
)

func main() {
	handler := api.NewHandler()
	//模拟进入http请求： 通过id获取年龄
	//测试用例 负数返回输入错误 0返回成功 正数返回sqlNowRows
	id := 1
	handler.GetUserAge(id)
}
