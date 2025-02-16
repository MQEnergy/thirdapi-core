package invoice

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/invoice"
)

type Invoice struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Invoice {
	return &Invoice{
		Base: ele.New(appId, appSecret),
	}
}

// InvoiceSettingsBatchUpdate 批量更新店铺开票设置
func (a *Invoice) InvoiceSettingsBatchUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(invoice.INVOICE_SETTINGS_BATCH_UPDATE_URL, params).WithClient()
}

// InvoiceApplyList 商家获取开票申请列表
func (a *Invoice) InvoiceApplyList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(invoice.INVOICE_APPLY_LIST_URL, params).WithClient()
}

// InvoiceSettingsList 查询店铺开票设置
func (a *Invoice) InvoiceSettingsList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(invoice.INVOICE_SETTINGS_LIST_URL, params).WithClient()
}

// InvoiceApplyBatchUpdate 批量更新店铺开票申请
func (a *Invoice) InvoiceApplyBatchUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(invoice.INVOICE_APPLY_BATCH_UPDATE_URL, params).WithClient()
}
