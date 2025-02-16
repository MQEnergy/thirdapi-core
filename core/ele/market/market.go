package market

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/market"
)

type Market struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Market {
	return &Market{
		Base: ele.New(appId, appSecret),
	}
}

// OrderList 获取某一段时间内的服务市场订单
func (a *Market) OrderList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(market.MARKET_ORDER_LIST_URL, params).WithClient()
}

// OrderGet 查询服务市场订单
func (a *Market) OrderGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(market.MARKET_ORDER_GET_URL, params).WithClient()
}
