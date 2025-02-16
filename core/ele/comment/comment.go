package comment

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/comment"
)

type Comment struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Comment {
	return &Comment{
		Base: ele.New(appId, appSecret),
	}
}

// UgcReply 回复评论
func (a *Comment) UgcReply(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(comment.UGC_REPLY_URL, params).WithClient()
}

// UgcOrderCommentGet 获取商户订单评论
func (a *Comment) UgcOrderCommentGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(comment.UGC_ORDERCOMMENT_GET_URL, params).WithClient()
}
