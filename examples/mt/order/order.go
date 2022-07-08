package main

import (
	"fmt"
	"github.com/chenxi/thirdapi-core/core/mt/order"
)

func main() {
	result, err := order.
		New("", "").
		GetOrderDetail(map[string]interface{}{
			"order_id": "4612593754970170004",
		})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
