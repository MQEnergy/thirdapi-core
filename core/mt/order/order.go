package order

import (
	"github.com/MQEnergy/thirdapi-core/core/mt"
	"github.com/MQEnergy/thirdapi-core/routers/mt/order"
)

type Order struct {
	Base *mt.Base
}

// New 实例化
func New(appId, appSecret string) *Order {
	return &Order{
		Base: mt.New(appId, appSecret),
	}
}

// PoiReceived 设订单为商家已收到
func (c *Order) PoiReceived(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.POI_RECEIVED_URL, data).WithClient("GET")
}

// Confirm 商家确认订单（必接）
func (c *Order) Confirm(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.CONFIRM_URL, data).WithClient("GET")
}

// Cancel 商家取消订单（必接）
func (c *Order) Cancel(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	if value, ok := params["reason"]; ok {
		data["reason"] = value
	}
	if value, ok := params["reason_code"]; ok {
		data["reason_code"] = value
	}
	return c.Base.WithRequestParams(order.CANCEL_URL, data).WithClient("GET")
}

// Delivering 自配订单配送中
func (c *Order) Delivering(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.DELIVERING_URL, data).WithClient("GET")
}

// Arrived 自配订单已到达
func (c *Order) Arrived(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.ARRIVED_URL, data).WithClient("GET")
}

// RefundAgree 订单确认退款请求（必接）
func (c *Order) RefundAgree(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
		"reason":   params["reason"],
	}
	return c.Base.WithRequestParams(order.REFUND_AGREE_URL, data).WithClient("GET")
}

// RefundReject 驳回订单退款申请（必接）
func (c *Order) RefundReject(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
		"reason":   params["reason"],
	}
	return c.Base.WithRequestParams(order.REFUND_REJECT_URL, data).WithClient("GET")
}

// ViewStatus 查询订单状态
func (c *Order) ViewStatus(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.VIEWSTATUS_URL, data).WithClient("GET")
}

// GetOrderDetail 获取订单详细信息
func (c *Order) GetOrderDetail(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	if value, ok := params["is_mt_logistics"]; ok {
		data["is_mt_logistics"] = value
	}
	return c.Base.WithRequestParams(order.GET_ORDER_DETAIL_URL, data).WithClient("GET")
}

// LogisticsPush 下发美团配送订单
func (c *Order) LogisticsPush(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.LOGISTICS_PUSH_URL, data).WithClient("GET")
}

// LogisticsCancel 取消美团配送订单
func (c *Order) LogisticsCancel(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.LOGISTICS_CANCEL_URL, data).WithClient("GET")
}

// LogisticsStatus 获取订单配送状态
func (c *Order) LogisticsStatus(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.LOGISTICS_STATUS_URL, data).WithClient("GET")
}

// GetOrderDaySeq 获取当日最新订单流水号
func (c *Order) GetOrderDaySeq(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_code": params["app_poi_code"],
	}
	return c.Base.WithRequestParams(order.GET_ORDER_DAY_SEQ_URL, data).WithClient("GET")
}

// GetOrderIdByDaySeq 根据流水号获取订单ID
func (c *Order) GetOrderIdByDaySeq(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":  params["order_id"],
		"date_time": params["date_time"],
		"day_seq":   params["day_seq"],
	}
	return c.Base.WithRequestParams(order.GET_ORDER_ID_BY_DAY_SEQ_URL, data).WithClient("GET")
}

// ZhongbaoShippingFee 批量查询众包配送费
func (c *Order) ZhongbaoShippingFee(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_ids": params["order_ids"],
	}
	return c.Base.WithRequestParams(order.ZHONGBAO_SHIPPING_FEE_URL, data).WithClient("GET")
}

// ZhongbaoDispatch 众包发配送
func (c *Order) ZhongbaoDispatch(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":     params["order_id"],
		"shipping_fee": params["shipping_fee"],
		"tip_amount":   params["tip_amount"],
	}
	return c.Base.WithRequestParams(order.ZHONGBAO_DISPATCH_URL, data).WithClient("GET")
}

// ZhongbaoUpdateTip 众包配送单追加小费
func (c *Order) ZhongbaoUpdateTip(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":   params["order_id"],
		"tip_amount": params["tip_amount"],
	}
	return c.Base.WithRequestParams(order.ZHONGBAO_UPDATE_TIP_URL, data).WithClient("GET")
}

// GetPartRefundFoods 查询可被部分退款的商品详情
func (c *Order) GetPartRefundFoods(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.GET_PART_REFUND_FOODS_URL, data).WithClient("GET")
}

// ApplyPartRefund 发起部分退款
func (c *Order) ApplyPartRefund(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":  params["order_id"],
		"reason":    params["reason"],
		"food_data": params["food_data"],
	}
	return c.Base.WithRequestParams(order.APPLY_PART_REFUND_URL, data).WithClient("POST")
}

// RemindReply 催单回复接口
func (c *Order) RemindReply(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":      params["order_id"],
		"remind_id":     params["remind_id"],
		"reply_id":      params["reply_id"],
		"reply_content": params["reply_content"],
	}
	return c.Base.WithRequestParams(order.REMIND_REPLY_URL, data).WithClient("POST")
}

// PreparationMealComplete 商家确认已完成拣货
func (c *Order) PreparationMealComplete(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.PREPARATION_MEAL_COMPLETE_URL, data).WithClient("GET")
}

// GetPreparationMealTime 商家获取备货时长
func (c *Order) GetPreparationMealTime(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.GET_PREPARATION_MEAL_TIME_URL, data).WithClient("GET")
}

// BatchPullPhoneNumber 拉取用户真实手机号（必接）
func (c *Order) BatchPullPhoneNumber(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"offset": params["offset"],
		"limit":  params["limit"],
	}
	if value, ok := params["app_poi_code"]; ok {
		data["app_poi_code"] = value
	}
	return c.Base.WithRequestParams(order.BATCH_PULL_PHONE_NUMBER_URL, data).WithClient("POST")
}

// GetSupportedCompensation 查询可申请货损赔付的订单
func (c *Order) GetSupportedCompensation(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_code": params["app_poi_code"],
		"offset":       params["offset"],
		"limit":        params["limit"],
	}
	return c.Base.WithRequestParams(order.GET_SUPPORTED_COMPENSATION_URL, data).WithClient("GET")
}

// ApplyCompensation 申请货损赔付
func (c *Order) ApplyCompensation(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":     params["order_id"],
		"reason":       params["reason"],
		"apply_status": params["apply_status"],
		"amount":       params["amount"],
	}
	return c.Base.WithRequestParams(order.APPLY_COMPENSATION_URL, data).WithClient("GET")
}

// GetCompensationResult 查询货损赔付结果
func (c *Order) GetCompensationResult(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.GET_COMPENSATION_RESULT_URL, data).WithClient("GET")
}

// LogisticsChangePoiSelf 专快混配送转为商家自配送
func (c *Order) LogisticsChangePoiSelf(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id": params["order_id"],
	}
	return c.Base.WithRequestParams(order.LOGISTICS_CHANGE_POI_SELF_URL, data).WithClient("POST")
}

// BatchFetchAbnormalOrder 批量拉取异常订单
func (c *Order) BatchFetchAbnormalOrder(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"start_time": params["start_time"],
		"end_time":   params["end_time"],
		"type":       params["type"],
		"offset":     params["offset"],
		"limit":      params["limit"],
	}
	return c.Base.WithRequestParams(order.BATCH_FETCH_ABNORMAL_ORDER_URL, data).WithClient("POST")
}

// GetEcoOrderRefundDetail 获取订单退款记录
func (c *Order) GetEcoOrderRefundDetail(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"wm_order_id_view": params["wm_order_id_view"],
	}
	if value, ok := params["refund_type"]; ok {
		data["refund_type"] = value
	}
	return c.Base.WithRequestParams(order.ECO_GET_ORDER_REFUND_DETAIL_URL, data).WithClient("GET")
}

// GetEcoOrderActDetail 查询订单中的活动信息
func (c *Order) GetEcoOrderActDetail(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"query_data": params["query_data"],
	}
	return c.Base.WithRequestParams(order.ECO_GET_ORDER_ACT_DETAIL_URL, data).WithClient("GET")
}

// GetEcoOrderIdByDaySeq 根据流水号获取订单ID（分段）
func (c *Order) GetEcoOrderIdByDaySeq(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_code":  params["app_poi_code"],
		"date_time":     params["date_time"],
		"day_seq_start": params["day_seq_start"],
		"day_seq_end":   params["day_seq_end"],
	}
	return c.Base.WithRequestParams(order.ECO_GET_ORDER_ID_BY_DAY_SEQ_URL, data).WithClient("GET")
}

// EcoOrderLogisticsSync 自配送商家同步配送信息
func (c *Order) EcoOrderLogisticsSync(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":               params["order_id"],
		"courier_name":           params["courier_name"],
		"courier_phone":          params["courier_phone"],
		"third_carrier_order_id": params["third_carrier_order_id"],
	}
	if value, ok := params["logistics_provider_code"]; ok {
		data["logistics_provider_code"] = value
	}
	if value, ok := params["logistics_status"]; ok {
		data["logistics_status"] = value
	}
	if value, ok := params["latitude"]; ok {
		data["latitude"] = value
	}
	if value, ok := params["longitude"]; ok {
		data["longitude"] = value
	}
	return c.Base.WithRequestParams(order.ECO_LOGISTICS_SYNC_URL, data).WithClient("POST")
}

// EcoOrderLogisticsBtocSync 根据流水号获取订单ID（分段）
func (c *Order) EcoOrderLogisticsBtocSync(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":                params["order_id"],
		"logistics_provider_code": params["logistics_provider_code"],
		"logistics_code":          params["logistics_code"],
	}
	return c.Base.WithRequestParams(order.ECO_LOGISTICS_BTOC_SYNC_URL, data).WithClient("POST")
}

// BatchCompensationOrder 快递配送商家同步配送信息接口
func (c *Order) BatchCompensationOrder(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"orderViewIds": params["orderViewIds"],
		"app_poi_code": params["app_poi_code"],
	}
	return c.Base.WithRequestParams(order.BATCH_COMPENSATION_ORDER_URL, data).WithClient("POST")
}

// GetRiderInfoPhoneNumber 拉取骑手真实手机号（必接）
func (c *Order) GetRiderInfoPhoneNumber(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"offset": params["offset"],
		"limit":  params["limit"],
	}
	if value, ok := params["app_poi_code"]; ok {
		data["app_poi_code"] = value
	}
	return c.Base.WithRequestParams(order.GET_RIDER_INFO_PHONE_NUMBER_URL, data).WithClient("POST")
}

// EcoOrderReviewAfterSales 拉取骑手真实手机号（必接）
func (c *Order) EcoOrderReviewAfterSales(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"wm_order_id_view": params["wm_order_id_view"],
		"review_type":      params["review_type"],
	}
	if value, ok := params["reject_reason_code"]; ok {
		data["reject_reason_code"] = value
	}
	if value, ok := params["reject_other_reason"]; ok {
		data["reject_other_reason"] = value
	}
	return c.Base.WithRequestParams(order.ECO_REVIEW_AFTER_SALES_URL, data).WithClient("GET")
}

// GetGwRpPictureList 根据订单id集合查询对应处方图片列表
func (c *Order) GetGwRpPictureList(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"app_poi_code": params["app_poi_code"],
		"order_ids":    params["order_ids"],
	}
	return c.Base.WithRequestParams(order.GW_RP_PICTURE_LIST_URL, data).WithClient("GET")
}

// SyncEstimateArrivalTime 自配送商家同步预计送达时间接口
func (c *Order) SyncEstimateArrivalTime(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"order_id":              params["order_id"],
		"estimate_arrival_time": params["estimate_arrival_time"],
	}
	return c.Base.WithRequestParams(order.ECO_SYNC_ARRIVAL_TIME, data).WithClient("GET")
}
