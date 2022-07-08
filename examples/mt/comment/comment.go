package main

import (
	"fmt"
	"github.com/MQEnergy/thirdapi-core/core/mt/comment"
)

func main() {
	result, err := comment.
		New("", "").
		CommentQuery(map[string]interface{}{
			"app_poi_code": "t_7sYOj3Ks9v",
			"start_time":   "",
			"end_time":     "",
			"pageoffset":   1,
			"pagesize":     10,
			"replyStatus":  "",
		})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
