package im

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/im"
)

type IM struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *IM {
	return &IM{
		Base: ele.New(appId, appSecret),
	}
}

// IMGetIMStatus 查询店铺IM状态
func (a *IM) IMGetIMStatus(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(im.IM_GETIMSTATUS_URL, params).WithClient()
}

// IMGetMediaUrl 获取多媒体文件url
func (a *IM) IMGetMediaUrl(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(im.IM_GET_MEDIA_URL, params).WithClient()
}

// IMUpdateOnlineStatus 设置im线上状态
func (a *IM) IMUpdateOnlineStatus(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(im.IM_UPDATE_ONLINE_STATUS_URL, params).WithClient()
}

// IMGetOnlineStatus 获取门店IM线上状态
func (a *IM) IMGetOnlineStatus(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(im.IM_GET_ONLINE_STATUS_URL, params).WithClient()
}

// IMMsgSend 商家IM消息发送
func (a *IM) IMMsgSend(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(im.IM_MESSAGE_SEND_URL, params).WithClient()
}

// IMUpdateStatus 更新店铺IM开关状态
func (a *IM) IMUpdateStatus(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(im.IM_UPDATE_STATUS_URL, params).WithClient()
}

// IMMsgRead 消息已读
func (a *IM) IMMsgRead(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(im.IM_MESSAGE_READ_URL, params).WithClient()
}
