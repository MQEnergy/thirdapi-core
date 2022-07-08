package poi

import (
	"github.com/chenxi/thirdapi-core/core/mt"
	"github.com/chenxi/thirdapi-core/routers/mt/poi"
)

type Poi struct {
	Base *mt.Base
}

// New 实例化
func New(appId, appSecret string) *Poi {
	return &Poi{
		Base: mt.New(appId, appSecret),
	}
}

// Save 创建或者更新门店信息
func (p *Poi) Save(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	if value, ok := params["name"]; ok {
		data["name"] = value
	}
	if value, ok := params["address"]; ok {
		data["address"] = value
	}
	if value, ok := params["latitude"]; ok {
		data["latitude"] = value
	}
	if value, ok := params["longitude"]; ok {
		data["longitude"] = value
	}
	if value, ok := params["phone"]; ok {
		data["phone"] = value
	}
	if value, ok := params["shipping_fee"]; ok {
		data["shipping_fee"] = value
	}
	if value, ok := params["shipping_time"]; ok {
		data["shipping_time"] = value
	}
	if value, ok := params["open_level"]; ok {
		data["open_level"] = value
	}
	if value, ok := params["is_online"]; ok {
		data["is_online"] = value
	}
	if value, ok := params["third_tag_name"]; ok {
		data["third_tag_name"] = value
	}
	if value, ok := params["pic_url"]; ok {
		data["pic_url"] = value
	}
	if value, ok := params["pic_url_large"]; ok {
		data["pic_url_large"] = value
	}
	if value, ok := params["standby_tel"]; ok {
		data["standby_tel"] = value
	}
	if value, ok := params["promotion_info"]; ok {
		data["promotion_info"] = value
	}
	if value, ok := params["invoice_support"]; ok {
		data["invoice_support"] = value
	}
	if value, ok := params["invoice_min_price"]; ok {
		data["invoice_min_price"] = value
	}
	if value, ok := params["invoice_description"]; ok {
		data["invoice_description"] = value
	}
	if value, ok := params["pre_book"]; ok {
		data["pre_book"] = value
	}
	if value, ok := params["time_select"]; ok {
		data["time_select"] = value
	}
	if value, ok := params["app_brand_code"]; ok {
		data["app_brand_code"] = value
	}
	if value, ok := params["pre_book_min_days"]; ok {
		data["pre_book_min_days"] = value
	}
	if value, ok := params["pre_book_max_days"]; ok {
		data["pre_book_max_days"] = value
	}
	return p.Base.WithRequestParams(poi.SAVE_URL, data).WithClient("POST")
}

// GetIds 获取门店ID
func (p *Poi) GetIds() (string, error) {
	return p.Base.WithRequestParams(poi.GET_IDS_URL, map[string]interface{}{}).WithClient("GET")
}

// Mget 批量获取门店详细信息
func (p *Poi) Mget(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	return p.Base.WithRequestParams(poi.MGET_URL, data).WithClient("GET")
}

// Open 批量获取门店详细信息
func (p *Poi) Open(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	return p.Base.WithRequestParams(poi.OPEN_URL, data).WithClient("POST")
}

// Close 门店设置为休息状态
func (p *Poi) Close(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	return p.Base.WithRequestParams(poi.CLOSE_URL, data).WithClient("POST")
}

// Offline 门店设置为下线状态
func (p *Poi) Offline(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	return p.Base.WithRequestParams(poi.OFFLINE_URL, data).WithClient("POST")
}

// Online 门店设置为上线状态
func (p *Poi) Online(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	return p.Base.WithRequestParams(poi.ONLINE_URL, data).WithClient("POST")
}

// AdditionalSave 更改门店附加信息（已下线）
func (p *Poi) AdditionalSave(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	if value, ok := params["app_poi_email"]; ok {
		data["app_poi_email"] = value
	}
	if value, ok := params["app_brand_code"]; ok {
		data["app_brand_code"] = value
	}
	if value, ok := params["app_org_id"]; ok {
		data["app_org_id"] = value
	}
	return p.Base.WithRequestParams(poi.ADDITIONAL_SAVE_URL, data).WithClient("POST")
}

// UpdatePromoteInfo 更改门店公告信息
func (p *Poi) UpdatePromoteInfo(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
	}
	if value, ok := params["promotion_info"]; ok {
		data["promotion_info"] = value
	}
	return p.Base.WithRequestParams(poi.UPDATE_PROMOTION_INFO_URL, data).WithClient("POST")
}

// TagList 获取门店品类列表
func (p *Poi) TagList() (string, error) {
	return p.Base.WithRequestParams(poi.TAG_LIST_URL, map[string]interface{}{}).WithClient("POST")
}

// ShippingtimeUpdate 更新门店营业时间
func (p *Poi) ShippingtimeUpdate(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
		"shipping_time": params["shipping_time"],
	}
	return p.Base.WithRequestParams(poi.SHIPPINGTIME_UPDATE_URL, data).WithClient("POST")
}

// LogisticsIsDelayPush 查询门店是否延迟发配送
func (p *Poi) LogisticsIsDelayPush(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
		"shipping_time": params["shipping_time"],
	}
	return p.Base.WithRequestParams(poi.LOGISTICS_ISDELAY_PUSH_URL, data).WithClient("POST")
}

// LogisticsSetDelayPush 设置门店延迟发配送时间
func (p *Poi) LogisticsSetDelayPush(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_codes": params["app_poi_codes"],
		"delay_seconds": params["delay_seconds"],
	}
	return p.Base.WithRequestParams(poi.LOGISTICS_SETDELAY_PUSH_URL, data).WithClient("POST")
}
