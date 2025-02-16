package bill

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/bill"
)

type Bill struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Bill {
	return &Bill{
		Base: ele.New(appId, appSecret),
	}
}

// GetBillDetail 半托管订单账单查询接口
func (a *Bill) GetBillDetail(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(bill.BILL_GETBILLDETAIL_URL, params).WithClient()
}

// GetBillDaily 半托管日汇总账单查询接口
func (a *Bill) GetBillDaily(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(bill.BILL_GETBILLDAILY_URL, params).WithClient()
}

// GetBillsGoods 半托管商品账单查询接口
func (a *Bill) GetBillsGoods(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(bill.BILL_GETBILLSGOODS_URL, params).WithClient()
}
