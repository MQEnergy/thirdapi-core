package drug

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/drug"
)

type Drug struct {
	Base *ele.Base
}

func New(source, secret string) *Drug {
	return &Drug{
		Base: ele.New(source, secret),
	}
}

// DrugTradePrescriptionDetailQuery 订单关联处方笺查询
func (c *Drug) DrugTradePrescriptionDetailQuery(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(drug.DRUG_TRADE_PRES_Q_URL, params).WithClient()
}

// DrugPrescriptionFiles 获取订单相关处方单
func (c *Drug) DrugPrescriptionFiles(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(drug.DRUG_PRES_FILES_URL, params).WithClient()
}

// MedicareSettlementQuery 医保订单结算信息查询
func (c *Drug) MedicareSettlementQuery(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(drug.MEDICARE_SETTLE_Q_URL, params).WithClient()
}

// DrugPresOrderAuditResult 药师审方结果回传
func (c *Drug) DrugPresOrderAuditResult(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(drug.DRUG_PRES_ORDER_AUDIT_RES_URL, params).WithClient()
}

// DrugSpuAuditApply 医药标品提报
func (c *Drug) DrugSpuAuditApply(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"barcode": params["barcode"],
		"title":   params["title"],
		"shopId":  params["shopId"],
		"images":  params["images"],
	}
	return c.Base.WithRequestParams(drug.DRUG_SPU_AUDIT_APPLY_URL, data).WithClient()
}

// DrugSpuAuditQuery 查询医药标品提报申请
func (c *Drug) DrugSpuAuditQuery(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(drug.DRUG_SPU_AUDIT_Q_URL, params).WithClient()
}
