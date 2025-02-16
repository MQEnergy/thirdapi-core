package shop

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/shop"
)

type Shop struct {
	Base *ele.Base
}

func New(source, appSecret string) *Shop {
	return &Shop{
		Base: ele.New(source, appSecret),
	}
}

// ShopGet 查看商户
func (c *Shop) ShopGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_GET_URL, params).WithClient()
}

// ShopList 商户列表
func (c *Shop) ShopList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_LIST_URL, params).WithClient()
}

// ShopStatusGet 查看商户状态
func (c *Shop) ShopStatusGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_STATUS_GET_URL, params).WithClient()
}

// ShopIdBatchupdate 商户三方门店ID映射
func (c *Shop) ShopIdBatchupdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_ID_BATCHUPDATE_URL, params).WithClient()
}

// ShopUpdate 修改商户
func (c *Shop) ShopUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_UPDATE_URL, params).WithClient()
}

// ShopAptitudeGet 获取商户资质
func (c *Shop) ShopAptitudeGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_APTITUDE_URL, params).WithClient()
}

// ShopExtMod 修改扩展属性
func (c *Shop) ShopExtMod(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_EXT_MOD_URL, params).WithClient()
}

// CommonShopCateGet 查询商户分类信息
func (c *Shop) CommonShopCateGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.COMMON_SHOP_CATE_GET_URL, params).WithClient()
}

// ShopIdList 商户ID列表
func (c *Shop) ShopIdList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_ID_LIST_URL, params).WithClient()
}

// ShopBatchGet 根据门店ID批量查询门店基本信息
func (c *Shop) ShopBatchGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_BATCH_GET_URL, params).WithClient()
}

// ShopBusstatusGet 查看商户的营业状态
func (c *Shop) ShopBusstatusGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_BUSSTATUS_GET_URL, params).WithClient()
}

// ShopOpen 商户开业
func (c *Shop) ShopOpen(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_OPEN_URL, params).WithClient()
}

// ShopClose 商户休息中
func (c *Shop) ShopClose(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_CLOSE_URL, params).WithClient()
}

// ShopDyinfoGet 查询自配送物流信息
func (c *Shop) ShopDyinfoGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_DYINFO_GET_URL, params).WithClient()
}

// ShopDyinfoSync 同步自配送配送范围
func (c *Shop) ShopDyinfoSync(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_DYINFO_SYNC_URL, params).WithClient()
}

// ShopBoundruleGet 查询自配送商家圈价费约束规则
func (c *Shop) ShopBoundruleGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_BOUNDRULE_GET_URL, params).WithClient()
}

// ShopAnnounceGet 获取商户公告
func (c *Shop) ShopAnnounceGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_ANNOUNCE_GET_URL, params).WithClient()
}

// ShopAnnounceSet 设置商户公告
func (c *Shop) ShopAnnounceSet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_ANNOUNCE_SET_URL, params).WithClient()
}

// ShopCreate 创建商户
func (c *Shop) ShopCreate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_CREATE_URL, params).WithClient()
}

// ShopAptitudeUpload 上传资质
func (c *Shop) ShopAptitudeUpload(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_APTITUDE_UPLOAD_URL, params).WithClient()
}

// ShopAptitudeGettypes 获取全部资质类型
func (c *Shop) ShopAptitudeGettypes(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_APTITUDE_GETTYPES_URL, params).WithClient()
}

// ShopCreateAccount 创建门店账号
func (c *Shop) ShopCreateAccount(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_CREATE_ACCOUNT_URL, params).WithClient()
}

// CommonBizcate 业态分类
func (c *Shop) CommonBizcate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.COMMON_BIZCATE_URL, params).WithClient()
}

// ShopCreateSubmit 门店开店信息提审
func (c *Shop) ShopCreateSubmit(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.SHOP_CREATE_SUBMIT_URL, params).WithClient()
}

// MemberOrderRefundCb 退款回调
func (c *Shop) MemberOrderRefundCb(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.MEMBER_ORD_REFUND_CB_URL, params).WithClient()
}

// MemberOrderCb 付费结果回调
func (c *Shop) MemberOrderCb(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.MEMBER_ORDER_CB_URL, params).WithClient()
}

// MemberUpdate 会员身份变更
func (c *Shop) MemberUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(shop.MEMBER_UPDATE_URL, params).WithClient()
}
