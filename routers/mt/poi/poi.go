package poi

/*
门店类
*/
const (
	SAVE_URL                    = "poi/save"                   // 创建或者更新门店信息
	GET_IDS_URL                 = "poi/getids"                 // 获取到家门店编码列表接口
	MGET_URL                    = "poi/mget"                   // 批量获取门店详细信息
	OPEN_URL                    = "poi/open"                   // 门店设置为营业状态
	CLOSE_URL                   = "poi/close"                  // 门店设置为休息状态
	OFFLINE_URL                 = "poi/offline"                // 门店设置为下线状态
	ONLINE_URL                  = "poi/online"                 // 门店设置为上线状态
	ADDITIONAL_SAVE_URL         = "poi/additional/save"        // 更改门店附加信息（已下线）
	UPDATE_PROMOTION_INFO_URL   = "poi/updatepromoteinfo"      // 更改门店公告信息
	TAG_LIST_URL                = "poiTag/list"                // 获取门店品类列表
	SHIPPINGTIME_UPDATE_URL     = "shippingtime/update"        // 更新门店营业时间
	LOGISTICS_ISDELAY_PUSH_URL  = "poi/logistics/isDelayPush"  // 查询门店是否延迟发配送
	LOGISTICS_SETDELAY_PUSH_URL = "poi/logistics/setDelayPush" // 设置门店延迟发配送时间
)
