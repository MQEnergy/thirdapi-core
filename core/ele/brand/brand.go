package brand

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/brand"
)

type Brand struct {
	Base *ele.Base
}

func New(appId, appSecret string) *Brand {
	return &Brand{
		Base: ele.New(appId, appSecret),
	}
}

// CommunityQuerySkuExt 根据店铺查询店铺下的商品信息-扩展补充
func (a *Brand) CommunityQuerySkuExt(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(brand.BRAND_COMM_Q_SKU_EXT_URL, params).WithClient()
}

// CrowdImpact 人群撞库
func (a *Brand) CrowdImpact(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(brand.BRAND_CROWD_IMPACT_URL, params).WithClient()
}

// SupplyItemIncrementQuery 品牌供给商品增量信息查询
func (a *Brand) SupplyItemIncrementQuery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(brand.BRAND_SUPPLY_ITEM_I_Q_URL, params).WithClient()
}

// CommunityQueryStore 社群导购品牌店铺列表
func (a *Brand) CommunityQueryStore(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(brand.BRAND_COMM_Q_STORE_URL, params).WithClient()
}

// CrowdUpload 品牌人群上传
func (a *Brand) CrowdUpload(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(brand.BRAND_CROWD_UPLOAD_URL, params).WithClient()
}

// CommunityQueryOrderList 社群导购品牌交易订单列表
func (a *Brand) CommunityQueryOrderList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(brand.BRAND_COMM_Q_ORDER_LIST_URL, params).WithClient()
}

// SupplyItemQuery 品牌供给商品查询
func (a *Brand) SupplyItemQuery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(brand.BRAND_SUPPLY_ITEM_Q_URL, params).WithClient()
}

// CommunityQuerySku 社群导购品牌商品列表
func (a *Brand) CommunityQuerySku(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(brand.BRAND_COMM_Q_SKU_URL, params).WithClient()
}
