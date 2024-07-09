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

// SkuStdupcExist 根据条形码查询是否平台药品SPU
func (c *Product) SkuStdupcExist(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"shop_id": params["shop_id"],
		"upc":     params["upc"],
	}
	return c.Base.WithRequestParams(product.SKU_STDUPC_EXIST_URL, data).WithClient()
}

// DrugSpuAuditApply 医药标品提报
func (c *Product) DrugSpuAuditApply(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"barcode": params["barcode"],
		"title":   params["title"],
		"shopId":  params["shopId"],
		"images":  params["images"],
	}
	return c.Base.WithRequestParams(product.DRUG_SPU_AUDIT_APPLY_URL, data).WithClient()
}
