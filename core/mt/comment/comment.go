package comment

import (
	"github.com/chenxi/thirdapi-core/core/mt"
	"github.com/chenxi/thirdapi-core/routers/mt/comment"
)

type Comment struct {
	Base *mt.Base
}

// New 实例化
func New(appId, appSecret string) *Comment {
	return &Comment{
		Base: mt.New(appId, appSecret),
	}
}

// CommentQuery 根据门店id批量查询评价信息（新版）
func (c *Comment) CommentQuery(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_code": params["app_poi_code"],
		"start_time":   params["start_time"],
		"end_time":     params["end_time"],
		"pageoffset":   params["pageoffset"],
		"pagesize":     params["pagesize"],
		"replyStatus":  params["replyStatus"],
	}
	return c.Base.WithRequestParams(comment.COMMENT_QUERY_URL, data).WithClient("GET")
}

// AddReply 根据评价id添加商家回复
func (c *Comment) AddReply(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_code": params["app_poi_code"],
		"start_time":   params["start_time"],
		"end_time":     params["end_time"],
		"pageoffset":   params["pageoffset"],
		"pagesize":     params["pagesize"],
		"replyStatus":  params["replyStatus"],
	}
	return c.Base.WithRequestParams(comment.COMMENT_ADD_REPLY_URL, data).WithClient("POST")
}

// CommentScore 通过门店ID获取当前门店评分
func (c *Comment) CommentScore(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_code": params["app_poi_code"],
	}
	return c.Base.WithRequestParams(comment.COMMENT_SCORE_URL, data).WithClient("GET")
}
