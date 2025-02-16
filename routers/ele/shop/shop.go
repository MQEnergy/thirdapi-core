package shop

const (
	SHOP_GET_URL               = "shop.get"                     // 查看商户
	SHOP_LIST_URL              = "shop.list"                    // 商户列表
	SHOP_STATUS_GET_URL        = "shop.status.get"              // 查看商户状态
	SHOP_ID_BATCHUPDATE_URL    = "shop.id.batchupdate"          // 商户三方门店ID映射
	SHOP_UPDATE_URL            = "shop.update"                  // 修改商户
	SHOP_APTITUDE_URL          = "shop.aptitude.get"            // 获取商户资质
	SHOP_EXT_MOD_URL           = "shop.extendAttribute.modify"  // 修改扩展属性
	COMMON_SHOP_CATE_GET_URL   = "common.shop.category.get"     // 查询商户分类信息
	SHOP_ID_LIST_URL           = "shop.id.list"                 // 商户ID列表
	SHOP_BATCH_GET_URL         = "shop.batch.get"               // 根据门店ID批量查询门店基本信息
	SHOP_BUSSTATUS_GET_URL     = "shop.busstatus.get"           // 查看商户的营业状态
	SHOP_OPEN_URL              = "shop.open"                    // 商户开业
	SHOP_CLOSE_URL             = "shop.close"                   // 商户休息中
	SHOP_DYINFO_GET_URL        = "shop.deliveryinfo.get"        // 查询自配送物流信息
	SHOP_DYINFO_SYNC_URL       = "shop.deliveryinfo.sync"       // 同步自配送配送范围
	SHOP_BOUNDRULE_GET_URL     = "shop.boundrule.get"           // 查询自配送商家圈价费约束规则
	SHOP_ANNOUNCE_GET_URL      = "shop.announcement.get"        // 获取商户公告
	SHOP_ANNOUNCE_SET_URL      = "shop.announcement.set"        // 设置商户公告
	SHOP_CREATE_URL            = "shop.create"                  // 创建商户
	SHOP_APTITUDE_UPLOAD_URL   = "shop.aptitude.upload"         // 上传资质
	SHOP_APTITUDE_GETTYPES_URL = "shop.aptitude.gettypes"       // 获取全部资质类型
	SHOP_CREATE_ACCOUNT_URL    = "shop.create.account"          // 创建门店账号
	COMMON_BIZCATE_URL         = "common.businesscategories"    // 业态分类
	SHOP_CREATE_SUBMIT_URL     = "shop.create.submit"           // 门店开店信息提审
	MEMBER_ORD_REFUND_CB_URL   = "member.order.refund.callback" // 退款回调
	MEMBER_ORDER_CB_URL        = "member.order.callback"        // 付费结果回调
	MEMBER_UPDATE_URL          = "member.update"                // 会员身份变更
)
