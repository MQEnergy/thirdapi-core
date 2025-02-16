package sms

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/sms"
)

type Sms struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Sms {
	return &Sms{
		Base: ele.New(appId, appSecret),
	}
}

// TemplateCreate 短信模板创建
func (c *Sms) TemplateCreate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(sms.SMS_TEMP_CREATE_URL, params).WithClient()
}

// TemplateQuery 短信模板创建
func (c *Sms) TemplateQuery(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(sms.SMS_QUERY_URL, params).WithClient()
}

// Send 发送短信
func (c *Sms) Send(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(sms.SMS_SEND_URL, params).WithClient()
}

// TemplateList 短信模板列表
func (c *Sms) TemplateList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(sms.SMS_TEMP_LIST_URL, params).WithClient()
}
