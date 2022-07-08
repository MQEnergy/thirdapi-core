package order

/*
订单类
*/
const (
	POI_RECEIVED_URL                = "order/poi_received"                      // 设订单为商家已收到
	CONFIRM_URL                     = "order/confirm"                           // 商家确认订单（必接）
	CANCEL_URL                      = "order/cancel"                            // 商家取消订单（必接）
	DELIVERING_URL                  = "order/delivering"                        // 自配订单配送中
	ARRIVED_URL                     = "order/arrived"                           // 自配订单已送达
	REFUND_AGREE_URL                = "order/refund/agree"                      // 订单确认退款请求（必接）
	REFUND_REJECT_URL               = "order/refund/reject"                     // 驳回订单退款申请（必接）
	VIEWSTATUS_URL                  = "order/viewstatus"                        // 查询订单状态
	GET_ORDER_DETAIL_URL            = "order/getOrderDetail"                    // 获取订单详细信息
	LOGISTICS_PUSH_URL              = "order/logistics/push"                    // 下发美团配送订单
	LOGISTICS_CANCEL_URL            = "order/logistics/cancel"                  // 取消美团配送订单
	LOGISTICS_STATUS_URL            = "order/logistics/status"                  // 获取订单配送状态
	GET_ORDER_DAY_SEQ_URL           = "order/getOrderDaySeq"                    // 获取当日最新订单流水号
	GET_ORDER_ID_BY_DAY_SEQ_URL     = "order/getOrderIdByDaySeq"                // 根据流水号获取订单ID
	ZHONGBAO_SHIPPING_FEE_URL       = "order/zhongbao/shippingFee"              // 批量查询众包配送费
	ZHONGBAO_DISPATCH_URL           = "order/zhongbao/dispatch"                 // 众包发配送
	ZHONGBAO_UPDATE_TIP_URL         = "order/zhongbao/update/tip"               // 众包配送单追加小费
	GET_PART_REFUND_FOODS_URL       = "order/getPartRefundFoods"                // 查询可被部分退款的商品详情
	APPLY_PART_REFUND_URL           = "order/applyPartRefund"                   // 发起部分退款
	REMIND_REPLY_URL                = "order/remindReply"                       // 催单回复接口
	PREPARATION_MEAL_COMPLETE_URL   = "order/preparationMealComplete"           // 商家确认已完成拣货
	GET_PREPARATION_MEAL_TIME_URL   = "order/getPreparationMealTime"            // 商家获取备货时长
	BATCH_PULL_PHONE_NUMBER_URL     = "order/batchPullPhoneNumber"              // 拉取用户真实手机号（必接）
	GET_SUPPORTED_COMPENSATION_URL  = "order/getSupportedCompensation"          // 查询可申请货损赔付的订单
	APPLY_COMPENSATION_URL          = "order/applyCompensation"                 // 申请货损赔付
	GET_COMPENSATION_RESULT_URL     = "order/getCompensationResult"             // 查询货损赔付结果
	LOGISTICS_CHANGE_POI_SELF_URL   = "order/logistics/change/poi_self"         // 专快混配送转为商家自配送
	BATCH_FETCH_ABNORMAL_ORDER_URL  = "order/batchFetchAbnormalOrder"           // 批量拉取异常订单
	ECO_GET_ORDER_REFUND_DETAIL_URL = "ecommerce/order/getOrderRefundDetail"    // 获取订单退款记录
	ECO_GET_ORDER_ACT_DETAIL_URL    = "ecommerce/order/getOrderActDetail"       // 查询订单中的活动信息
	ECO_GET_ORDER_ID_BY_DAY_SEQ_URL = "ecommerce/order/getOrderIdByDaySeq"      // 查询订单中的活动信息
	ECO_LOGISTICS_SYNC_URL          = "ecommerce/order/logistics/sync"          // 自配送商家同步配送信息
	ECO_LOGISTICS_BTOC_SYNC_URL     = "ecommerce/order/logistics/btoc/sync"     // 快递配送商家同步配送信息接口
	BATCH_COMPENSATION_ORDER_URL    = "order/batchCompensationOrder"            // 批量查询客服赔付商家责任订单信息
	GET_RIDER_INFO_PHONE_NUMBER_URL = "order/getRiderInfoPhoneNumber"           // 拉取骑手真实手机号（必接）
	ECO_REVIEW_AFTER_SALES_URL      = "ecommerce/order/reviewAfterSales"        // 售后审查接口
	GW_RP_PICTURE_LIST_URL          = "gw/rp/picture/list"                      // 根据订单id查询对应处方图片列表
	ECO_SYNC_ARRIVAL_TIME           = "ecommerce/order/syncEstimateArrivalTime" // 自配送商家同步预计送达时间接口
)
