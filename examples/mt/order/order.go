package main

import (
	"fmt"
	"github.com/MQEnergy/thirdapi-core/core/mt/order"
)

func main() {
	// 获取订单详情
	order := order.New("", "")
	params := map[string]interface{}{
		"order_id": "62980670405312532",
	}
	result, err := order.GetOrderDetail(params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", result)

	// 验证sig签名是否正确
	order.Base.SystemData["order_id"] = params["order_id"]
	res := order.Base.VerifySign(order.Base.SystemData)
	fmt.Println("res:", res)
}
