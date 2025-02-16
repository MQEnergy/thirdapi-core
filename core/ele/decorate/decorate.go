package decorate

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/decorate"
)

type Decorate struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Decorate {
	return &Decorate{
		Base: ele.New(appId, appSecret),
	}
}

// DecorateWindowEdit 创建编辑老板推荐橱窗
func (a *Decorate) DecorateWindowEdit(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(decorate.DECORATE_WIN_EDIT_URL, params).WithClient()
}

// DecorateWindowGetList 查询门店老板推荐列表
func (a *Decorate) DecorateWindowGetList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(decorate.DECORATE_WIN_GETLIST_URL, params).WithClient()
}

// DecoratePosterEdit 创建编辑电子海报
func (a *Decorate) DecoratePosterEdit(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(decorate.DECORATE_POSTER_EDIT_URL, params).WithClient()
}

// DecorateWindowGetDetail 查询指定老板推荐内容
func (a *Decorate) DecorateWindowGetDetail(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(decorate.DECORATE_WIN_GETDETAIL_URL, params).WithClient()
}

// DecoratePosterOffline 下线电子海报
func (a *Decorate) DecoratePosterOffline(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(decorate.DECORATE_POSTER_OFFLINE_URL, params).WithClient()
}

// DecoratePosterGetDetail 查询指定海报内容
func (a *Decorate) DecoratePosterGetDetail(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(decorate.DECORATE_POSTER_GETDETAIL_URL, params).WithClient()
}
