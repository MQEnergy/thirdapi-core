package shipping

import (
	"github.com/chenxi/thirdapi-core/core/mt"
	"github.com/chenxi/thirdapi-core/routers/mt/shipping"
)

type Shipping struct {
	Base *mt.Base
}

// New 实例化
func New(appId, appSecret string) *Shipping {
	return &Shipping{
		Base: mt.New(appId, appSecret),
	}
}

// Save 创建/更新门店配送范围
func (s *Shipping) Save(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes":     params["app_poi_codes"],
		"app_shipping_code": params["app_shipping_code"],
		"type":              params["type"],
		"area":              params["area"],
	}
	if value, ok := params["min_price"]; ok {
		data["min_price"] = value
	}
	if value, ok := params["shipping_fee"]; ok {
		data["shipping_fee"] = value
	}
	if value, ok := params["distance_markup_execute_type"]; ok {
		data["distance_markup_execute_type"] = value
	}
	if value, ok := params["distance_markup_factors"]; ok {
		data["distance_markup_factors"] = value
	}
	if value, ok := params["weight_markup_factors"]; ok {
		data["weight_markup_factors"] = value
	}
	if value, ok := params["time_markup_execute_type"]; ok {
		data["time_markup_execute_type"] = value
	}
	if value, ok := params["time_markup_factors"]; ok {
		data["time_markup_factors"] = value
	}
	return s.Base.WithRequestParams(shipping.SAVE_URL, data).WithClient("POST")
}

// List 获取门店配送范围（自配送）
func (s *Shipping) List(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	return s.Base.WithRequestParams(shipping.LIST_URL, data).WithClient("GET")
}

// BatchSave 批量创建/更新配送范围
func (s *Shipping) BatchSave(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
		"shipping_data": params["shipping_data"],
	}
	return s.Base.WithRequestParams(shipping.BATCH_SAVE_URL, data).WithClient("POST")
}

// Fetch 获取门店配送范围（混合送）
func (s *Shipping) Fetch(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	return s.Base.WithRequestParams(shipping.FETCH_URL, data).WithClient("GET")
}

// Delete 删除配送范围
func (s *Shipping) Delete(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes":     params["app_poi_codes"],
		"app_shipping_code": params["app_shipping_code"],
	}
	return s.Base.WithRequestParams(shipping.DELETE_URL, data).WithClient("POST")
}

// SpecSave 新增/更新特殊时段配送范围
func (s *Shipping) SpecSave(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes":     params["app_poi_codes"],
		"app_shipping_code": params["app_shipping_code"],
		"type":              params["type"],
		"area":              params["area"],
		"shipping_fee":      params["shipping_fee"],
		"time_range":        params["time_range"],
	}
	if value, ok := params["min_price"]; ok {
		data["min_price"] = value
	}
	if value, ok := params["distance_markup_execute_type"]; ok {
		data["distance_markup_execute_type"] = value
	}
	if value, ok := params["distance_markup_factors"]; ok {
		data["distance_markup_factors"] = value
	}
	if value, ok := params["weight_markup_execute_type"]; ok {
		data["weight_markup_execute_type"] = value
	}
	if value, ok := params["weight_markup_factors"]; ok {
		data["weight_markup_factors"] = value
	}
	if value, ok := params["time_markup_execute_type"]; ok {
		data["time_markup_execute_type"] = value
	}
	if value, ok := params["time_markup_factors"]; ok {
		data["time_markup_factors"] = value
	}
	return s.Base.WithRequestParams(shipping.SPEC_SAVE_URL, data).WithClient("POST")
}
