package order

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/order"
)

type Order struct {
	Base *ele.Base
}

// New 实例化
func New(source, appSecret string) *Order {
	return &Order{
		Base: ele.New(source, appSecret),
	}
}

// Confirm 确认订单
func (a *Order) Confirm(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CONFIRM_URL, params).WithClient()
}

// Cancel 取消订单
func (a *Order) Cancel(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CANCEL_URL, params).WithClient()
}

// Complete 订单送达
func (a *Order) Complete(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.COMPLETE_URL, params).WithClient()
}

// StatusGet 查看订单状态
func (a *Order) StatusGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.STATUS_GET_URL, params).WithClient()
}

// Get 查看订单详情
func (a *Order) Get(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.GET_URL, params).WithClient()
}

// List 查看订单列表
func (a *Order) List(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.LIST_URL, params).WithClient()
}

// StatusPush 订单状态推送
func (a *Order) StatusPush(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.STATUS_PUSH_URL, params).WithClient()
}

// Create 创建订单
func (a *Order) Create(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CREATE_URL, params).WithClient()
}

// Partrefund 商户发起部分退款申请
func (a *Order) Partrefund(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.PARTREFUND_URL, params).WithClient()
}

// PartrefundPush 部分退款订单信息推送
func (a *Order) PartrefundPush(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.PARTREFUND_PUSH_URL, params).WithClient()
}

// IdConvert 订单ID转换
func (a *Order) IdConvert(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.ID_CONVERT_URL, params).WithClient()
}

// Disagreerefund 拒绝用户取消单/退单
func (a *Order) Disagreerefund(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.DISAGREEREFUND_URL, params).WithClient()
}

// Agreerefund 同意用户取消单/退单
func (a *Order) Agreerefund(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.AGREEREFUND_URL, params).WithClient()
}

// Privateinfo 订单隐私信息
func (a *Order) Privateinfo(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.PRIVATEINFO_URL, params).WithClient()
}

// SendCoupon 订单发送红包
func (a *Order) SendCoupon(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.SEND_COUPON_URL, params).WithClient()
}

// CompensateCouponList 订单补偿红包列表
func (a *Order) CompensateCouponList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.COMPENSATE_COUPON_URL, params).WithClient()
}

// PartrefundGet 查看部分退款订单详情
func (a *Order) PartrefundGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.PARTREFUND_GET_URL, params).WithClient()
}

// Cancellist 订单申请取消/退单列表
func (a *Order) Cancellist(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CANCELLIST_URL, params).WithClient()
}

// PartrefundUntrelist 订单申请取消/退单列表用户侧部分退款未处理订单列表
func (a *Order) PartrefundUntrelist(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.PARTREFUND_UNTRELIST_URL, params).WithClient()
}

// GetDeliveryFeeForCrowd 获取众包订单配送费
func (a *Order) GetDeliveryFeeForCrowd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.GET_DELIVERY_FEE_FOR_CROWD_URL, params).WithClient()
}

// CallDelivery 呼叫配送
func (a *Order) CallDelivery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CALL_DELIVERY_URL, params).WithClient()
}

// CancelDelivery 取消呼叫配送
func (a *Order) CancelDelivery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CANCEL_DELIVERY_URL, params).WithClient()
}

// DeliveryGet 获取订单配送信息
func (a *Order) DeliveryGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.DELIVERY_GET_URL, params).WithClient()
}

// RefundDeliveryGet 逆向单配送信息获取
func (a *Order) RefundDeliveryGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.REFUND_DELIVERY_GET_URL, params).WithClient()
}

// StopDelivery 配送异常不再配送
func (a *Order) StopDelivery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.STOP_DELIVERY_URL, params).WithClient()
}

// SwitchSelfDelivery 配送异常切自配送
func (a *Order) SwitchSelfDelivery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.SWITCH_SELF_DELIVERY_URL, params).WithClient()
}

// Sendout 订单送出
func (a *Order) Sendout(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.SENDOUT_URL, params).WithClient()
}

// RemindGet 获取商户未处理催单
func (a *Order) RemindGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.REMIND_GET_URL, params).WithClient()
}

// RemindReply 商户回复催单
func (a *Order) RemindReply(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.REMIND_REPLY_URL, params).WithClient()
}

// LimitedTaking 设置门店限单值
func (a *Order) LimitedTaking(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.LIMITED_TAKING_URL, params).WithClient()
}

// PickComplete 订单拣货完成
func (a *Order) PickComplete(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.PICK_COMPLETE_URL, params).WithClient()
}

// UserCancel 用户申请订单取消/退款
func (a *Order) UserCancel(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.USER_CANCEL_URL, params).WithClient()
}

// Checkout 到店自提订单核销
func (a *Order) Checkout(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CHECKOUT_URL, params).WithClient()
}

// EpidemicGet 获取医药订单疫情登记数据
func (a *Order) EpidemicGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.EPIDEMIC_GET_URL, params).WithClient()
}

// SelfDeliveryStateSync 自配送接入骑手状态
func (a *Order) SelfDeliveryStateSync(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.SELFDELIVERYSTATESYNC, params).WithClient()
}

// SelfDeliveryLocationSync 自配送接入骑手轨迹
func (a *Order) SelfDeliveryLocationSync(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.SELFDELIVERYLOCATIONSYNC, params).WithClient()
}

// ReverseProcess 商家审核逆向单
func (a *Order) ReverseProcess(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.REVERSE_PROCESS, params).WithClient()
}

// ReverseQuery 商家查询逆向单
func (a *Order) ReverseQuery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.REVERSE_QUERY, params).WithClient()
}

// ReverseApply 商家发起逆向退款
func (a *Order) ReverseApply(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.REVERSE_APPLY, params).WithClient()
}

// UnprocessedList 查询待处理的逆向单列表
func (a *Order) UnprocessedList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.UNPROCESSED_LIST, params).WithClient()
}

// ReverseConsult 查询订单逆向可退信息及退款预览
func (a *Order) ReverseConsult(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.REVERSE_CONSULT, params).WithClient()
}

// ReverseList 查询全量逆向单列表
func (a *Order) ReverseList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.REVERSE_LIST, params).WithClient()
}

// ExpressGet 获取订单运单信息
func (a *Order) ExpressGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.EXPRESS_GET_URL, params).WithClient()
}

// BindExpress 订单绑定运单
func (a *Order) BindExpress(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.BIND_EXPRESS_URL, params).WithClient()
}

// UpdateExpress 更新订单运单信息
func (a *Order) UpdateExpress(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.UPDATE_EXPRESS_URL, params).WithClient()
}

// CallPltDelivery 订单呼叫平台配送
func (a *Order) CallPltDelivery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CALL_PLT_DELIVERY_URL, params).WithClient()
}

// DeliveryRemark 商家给骑手备注
func (a *Order) DeliveryRemark(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.DELIVERY_REMARK_URL, params).WithClient()
}

// CancelUnionDelivery 【蜂鸟跑腿】取消呼叫蜂鸟跑腿配送
func (a *Order) CancelUnionDelivery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CANCEL_UNION_DELIVERY_URL, params).WithClient()
}

// OpenUnionDelivery 【蜂鸟跑腿】查询商家是否开通蜂鸟跑腿
func (a *Order) OpenUnionDelivery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.OPEN_UNION_DELIVERY_URL, params).WithClient()
}

// DeliveryUnion 【蜂鸟跑腿】查询商家是否开通蜂鸟跑腿
func (a *Order) DeliveryUnion(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.DELIVERY_UNION_URL, params).WithClient()
}

// CallUnionDelivery 【蜂鸟跑腿】呼叫蜂鸟跑腿配送
func (a *Order) CallUnionDelivery(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.CALL_UNION_DELIVERY_URL, params).WithClient()
}

// AddUnionTip 【蜂鸟跑腿】蜂鸟跑腿加小费
func (a *Order) AddUnionTip(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.ADD_UNION_TIP_URL, params).WithClient()
}

// BillGet 获取账单
func (a *Order) BillGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.BILL_GET_URL, params).WithClient()
}

// OrderDetail 获取账单订单明细信息
func (a *Order) OrderDetail(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.ORDER_DETAIL_URL, params).WithClient()
}

// ApplyCompensation 订单申请索赔
func (a *Order) ApplyCompensation(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.APPLY_COMPENSATION_URL, params).WithClient()
}

// CompensationResult 订单索赔结果
func (a *Order) CompensationResult(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.COMPENSATION_RESULT_URL, params).WithClient()
}

// CompensationList 获取可索赔订单列表
func (a *Order) CompensationList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(order.COMPENSATION_LIST_URL, params).WithClient()
}
