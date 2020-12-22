package main

import (
	"./easycc"
	"encoding/json"
	"fmt"
)

type Person struct {
	//序列化,后面跟tag可以将其转化为你需要的名称
}

func main() {
	var concurrency uint = 40 // 并发量

	data := Person{
		//填写序列化后的参数
	}
	jsonData, _ := json.Marshal(data)
	fmt.Println("json data is " + string(jsonData)) // body使用json数据

	BASE_URL := "TestURL"
	req := &easycc.CCRequest{
		Method: "POST",
		URL:    BASE_URL,
		Body:   jsonData,
		Headers: map[string]string{
			"Content-Type": "application/json;charset=utf-8",
			"User-Agent":   "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1 wechatdevtools/1.03.2011120 MicroMessenger/7.0.4 Language/zh_CN webview/",
		},
	}
	respArr := easycc.CCTest(req, concurrency)

	fmt.Println("All requests finished")

	for i := 0; i < len(respArr); i++ {
		resp := respArr[i]
		if resp.Err != nil {
			fmt.Println("This request encouter error " + resp.Err.Error())
		}else{
			// 此处依据返回状态码进行判断,可自行改造反馈信息
			if int(resp.StatusCode) != 200 {
				fmt.Println("Requests fail " + string(resp.Status))
			}else{
				fmt.Println("Requests success " + string(resp.Status))
			}
		}
	}
}
