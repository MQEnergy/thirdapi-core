package product

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/product"
)

type Product struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Product {
	return &Product{
		Base: ele.New(appId, appSecret),
	}
}

// SkuStdupcExist 根据门店id批量查询评价信息（新版）
func (c *Product) SkuStdupcExist(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"shop_id": params["shop_id"],
		"upc":     params["upc"],
	}
	return c.Base.WithRequestParams(product.SKU_STDUPC_EXIST_URL, data).WithClient()
}