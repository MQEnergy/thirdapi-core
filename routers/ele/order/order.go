package order

const (
	CONFIRM_URL                    = "order.confirm"                  // 确认订单
	CANCEL_URL                     = "order.cancel"                   // 取消订单
	COMPLETE_URL                   = "order.complete"                 // 订单送达
	STATUS_GET_URL                 = "order.status.get"               // 查看订单状态
	GET_URL                        = "order.get"                      // 查看订单详情
	LIST_URL                       = "order.list"                     // 查看订单列表
	STATUS_PUSH_URL                = "order.status.push"              // 订单状态推送
	CREATE_URL                     = "order.create"                   // 创建订单
	PARTREFUND_URL                 = "order.partrefund"               // 商户发起部分退款申请
	PARTREFUND_PUSH_URL            = "order.partrefund.push"          // 部分退款订单信息推送
	ID_CONVERT_URL                 = "order.id.convert"               // 订单ID转换
	DISAGREEREFUND_URL             = "order.disagreerefund"           // 拒绝用户取消单/退单
	AGREEREFUND_URL                = "order.agreerefund"              // 同意用户取消单/退单
	PRIVATEINFO_URL                = "order.privateinfo"              // 订单隐私信息
	SEND_COUPON_URL                = "order.send.coupon"              // 订单发送红包
	COMPENSATE_COUPON_URL          = "order.compensate.coupon.list"   // 订单补偿红包列表
	PARTREFUND_GET_URL             = "order.partrefund.get"           // 查看部分退款订单详情
	CANCELLIST_URL                 = "order.cancellist"               // 订单申请取消/退单列表
	PARTREFUND_UNTRELIST_URL       = "order.partrefund.untrelist"     // 订单申请取消/退单列表用户侧部分退款未处理订单列表
	GET_DELIVERY_FEE_FOR_CROWD_URL = "order.getDeliveryFeeForCrowd"   // 获取众包订单配送费
	CALL_DELIVERY_URL              = "order.callDelivery"             // 呼叫配送
	CANCEL_DELIVERY_URL            = "order.cancelDelivery"           // 取消呼叫配送
	DELIVERY_GET_URL               = "order.delivery.get"             // 获取订单配送信息
	REFUND_DELIVERY_GET_URL        = "order.refundDelivery.get"       // 逆向单配送信息获取
	STOP_DELIVERY_URL              = "order.stopdelivery"             // 配送异常不再配送
	SWITCH_SELF_DELIVERY_URL       = "order.switchselfdelivery"       // 配送异常切自配送
	SENDOUT_URL                    = "order.sendout"                  // 订单送出
	REMIND_GET_URL                 = "order.remind.get"               // 获取商户未处理催单
	REMIND_REPLY_URL               = "order.remind.reply"             // 商户回复催单
	LIMITED_TAKING_URL             = "order.limited.taking"           // 设置门店限单值
	PICK_COMPLETE_URL              = "order.pickcomplete"             // 订单拣货完成
	USER_CANCEL_URL                = "order.user.cancel"              // 用户申请订单取消/退款
	CHECKOUT_URL                   = "order.selfpick.checkout"        // 到店自提订单核销
	EPIDEMIC_GET_URL               = "order.epidemic.get"             // 获取医药订单疫情登记数据
	SELFDELIVERYSTATESYNC          = "order.selfDeliveryStateSync"    // 自配送接入骑手状态
	SELFDELIVERYLOCATIONSYNC       = "order.selfDeliveryLocationSync" // 自配送接入骑手轨迹
	REVERSE_PROCESS                = "order.reverse.process"          // 商家审核逆向单
	REVERSE_QUERY                  = "order.reverse.query"            // 商家查询逆向单
	REVERSE_APPLY                  = "order.reverse.apply"            // 商家发起逆向退款
	UNPROCESSED_LIST               = "order.reverse.unprocessedlist"  // 查询待处理的逆向单列表
	REVERSE_CONSULT                = "order.reverse.consult"          // 查询订单逆向可退信息及退款预览
	REVERSE_LIST                   = "order.reverse.list"             // 查询全量逆向单列表
	EXPRESS_GET_URL                = "order.express.get"              //【快递发货】获取订单运单信息
	BIND_EXPRESS_URL               = "order.bind.express"             //【快递发货】订单绑定运单
	UPDATE_EXPRESS_URL             = "order.update.express"           //【快递发货】更新订单运单信息
	CALL_PLT_DELIVERY_URL          = "order.callPlatformDelivery"     // 订单呼叫平台配送
	DELIVERY_REMARK_URL            = "order.delivery.remark"          // 商家给骑手备注
	CANCEL_UNION_DELIVERY_URL      = "order.cancelUnionDelivery"      //【蜂鸟跑腿】取消呼叫蜂鸟跑腿配送
	OPEN_UNION_DELIVERY_URL        = "order.query.openUnionDelivery"  //【蜂鸟跑腿】查询商家是否开通蜂鸟跑腿
	DELIVERY_UNION_URL             = "order.getDeliveryInfoForUnion"  //【蜂鸟跑腿】查询商家是否开通蜂鸟跑腿
	CALL_UNION_DELIVERY_URL        = "order.callUnionDelivery"        //【蜂鸟跑腿】呼叫蜂鸟跑腿配送
	ADD_UNION_TIP_URL              = "order.addUnionTip"              //【蜂鸟跑腿】蜂鸟跑腿加小费
	BILL_GET_URL                   = "bill.get"                       // 获取账单
	ORDER_DETAIL_URL               = "bill.orderdetail"               // 获取账单订单明细信息
	APPLY_COMPENSATION_URL         = "order.apply.compensation"       // 订单申请索赔
	COMPENSATION_RESULT_URL        = "order.compensation.result"      // 订单索赔结果
	COMPENSATION_LIST_URL          = "order.compensation.list"        // 获取可索赔订单列表
)
