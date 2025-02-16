package activity

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/activity"
)

type Activity struct {
	Base *ele.Base
}

func New(source, appSecret string) *Activity {
	return &Activity{Base: ele.New(source, appSecret)}
}

// ActivityDisable 【通用】下线商品营销活动（所有活动类型）
func (a *Activity) ActivityDisable(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_DISABLE_URL, params).WithClient()
}

// ActMultiChannelQueryInfos 【多渠道查询】多渠道查询活动详情
func (a *Activity) ActMultiChannelQueryInfos(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_MULTI_CHAN_QINFOS_URL, params).WithClient()
}

// ActMultiChannelQueryIds 多渠道查询获取参与的活动Id列表
func (a *Activity) ActMultiChannelQueryIds(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_MULTI_CHAN_QIDS_URL, params).WithClient()
}

// QuerySkuActivityList 【通用】查询商品参与的活动（非券活动）
func (a *Activity) QuerySkuActivityList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.Q_SKU_ACTIVITY_LIST_URL, params).WithClient()
}

// ActivityCreate 【单品直降、品类满减】创建商品营销活动
func (a *Activity) ActivityCreate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_CREATE_URL, params).WithClient()
}

// ActivityGet 查看活动(直降、品类满减、N选一)
func (a *Activity) ActivityGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_GET_URL, params).WithClient()
}

// ActivityUpdate 【单品直降、品类满减】更新活动信息
func (a *Activity) ActivityUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_UPDATE_URL, params).WithClient()
}

// ActivitySkuAdd 【单品直降、品类满减】添加活动商品
func (a *Activity) ActivitySkuAdd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_SKU_ADD_URL, params).WithClient()
}

// ActivitySkuUpdate 【单品直降、品类满减】更新活动商品
func (a *Activity) ActivitySkuUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_SKU_UPDATE_URL, params).WithClient()
}

// ActivitySkuList 查看活动商品信息（直降、N选1、品类满减、第N件特价）
func (a *Activity) ActivitySkuList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_SKU_LIST_URL, params).WithClient()
}

// ActivitySkuDelete 删除活动商品（除N选1外，所有活动类型）
func (a *Activity) ActivitySkuDelete(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_SKU_DELETE_URL, params).WithClient()
}

// ActivitySkuUpdateBatch 批量更新活动商品（直降）
func (a *Activity) ActivitySkuUpdateBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_SKU_UPDATE_BATCH_URL, params).WithClient()
}

// ActivitySkuAddBatch 【单品直降】批量添加活动商品
func (a *Activity) ActivitySkuAddBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_SKU_ADD_BATCH_URL, params).WithClient()
}

// ActivitySkuDeleteBatch 批量删除活动商品（除N选1外，所有活动类型）
func (a *Activity) ActivitySkuDeleteBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_SKU_DELETE_BATCH_URL, params).WithClient()
}

// ActItemSkuUpdateBatch 商品真折扣批量修改商品明细
func (a *Activity) ActItemSkuUpdateBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_ITEM_SKU_UPDATE_BATCH_URL, params).WithClient()
}

// ActItemCreate 商品真折扣创建活动
func (a *Activity) ActItemCreate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_ITEM_CREATE_URL, params).WithClient()
}

// ActItemModify 商品真折扣修改活动信息
func (a *Activity) ActItemModify(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_ITEM_MODIFY_URL, params).WithClient()
}

// ActItemSkuUpdate 商品真折扣修改商品明细
func (a *Activity) ActItemSkuUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_ITEM_SKU_UPDATE_URL, params).WithClient()
}

// ActItemSkuAddBatch 商品真折扣批量添加商品明细
func (a *Activity) ActItemSkuAddBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_ITEM_SKU_ADD_BATCH_URL, params).WithClient()
}

// ActItemSkuAdd 商品真折扣添加商品明细
func (a *Activity) ActItemSkuAdd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_ITEM_SKU_ADD_URL, params).WithClient()
}

// ActivityNgiftmUpdate 更新买N赠M活动信息
func (a *Activity) ActivityNgiftmUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NGIFTM_UPDATE_URL, params).WithClient()
}

// ActivityNgiftmSkuAdd 添加买N赠M活动商品和相应赠品
func (a *Activity) ActivityNgiftmSkuAdd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NGIFTM_SKU_ADD_URL, params).WithClient()
}

// ActivityNgiftmSkuAddBatch 买N赠M活动批量添加商品
func (a *Activity) ActivityNgiftmSkuAddBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NGIFTM_SKU_ADD_BATCH_URL, params).WithClient()
}

// ActivityNgiftmGet 查看买N赠M活动信息
func (a *Activity) ActivityNgiftmGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NGIFTM_GET_URL, params).WithClient()
}

// ActivityNgiftmSkuList 买N赠M查询活动商品信息列表
func (a *Activity) ActivityNgiftmSkuList(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NGIFTM_SKU_LIST_URL, params).WithClient()
}

// ActivityNgiftmSkuUpdateBatch 买N赠M批量更新商品库存、日库存
func (a *Activity) ActivityNgiftmSkuUpdateBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NGIFTM_SKU_UPDATE_BATCH_URL, params).WithClient()
}

// ActivityNgiftmSkuUpdate 买N赠M更新商品库存、日库存
func (a *Activity) ActivityNgiftmSkuUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NGIFTM_SKU_UPDATE_URL, params).WithClient()
}

// ActivityNgiftmCreate 创建买N赠M营销活动
func (a *Activity) ActivityNgiftmCreate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NGIFTM_CREATE_URL, params).WithClient()
}

// ActivityNitesSkuAddBatch 第N件特价批量添加商品
func (a *Activity) ActivityNitesSkuAddBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NITES_SKU_ADD_BATCH_URL, params).WithClient()
}

// ActivityNitesCreate 创建第N件特价活动
func (a *Activity) ActivityNitesCreate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NITES_CREATE_URL, params).WithClient()
}

// ActivityNitesSkuUpdate 第N件特价更新商品库存、日库存
func (a *Activity) ActivityNitesSkuUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NITES_SKU_UPDATE_URL, params).WithClient()
}

// ActivityNitesGet 查看第N件特价活动信息
func (a *Activity) ActivityNitesGet(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NITES_GET_URL, params).WithClient()
}

// ActivityNitesSkuAdd 添加第N件特价活动商品
func (a *Activity) ActivityNitesSkuAdd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NITES_SKU_ADD_URL, params).WithClient()
}

// ActivityNitesUpdate 更新第N件特价活动信息
func (a *Activity) ActivityNitesUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NITES_UPDATE_URL, params).WithClient()
}

// ActivityNitesSkuUpdateBatch 第N件特价批量更新商品库存、日库存
func (a *Activity) ActivityNitesSkuUpdateBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACTIVITY_NITES_SKU_UPDATE_BATCH_URL, params).WithClient()
}

// ActNitedSkuAddBatch 商品第N件特价真折扣批量添加商品明细
func (a *Activity) ActNitedSkuAddBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NITED_SKU_ADD_BATCH_URL, params).WithClient()
}

// ActNitedSkuAdd 商品第N件特价真折扣添加商品明细
func (a *Activity) ActNitedSkuAdd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NITED_SKU_ADD_URL, params).WithClient()
}

// ActNitedModify 商品第N件特价真折扣修改活动信息
func (a *Activity) ActNitedModify(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NITED_MODIFY_URL, params).WithClient()
}

// ActNitedSkuUpdate 商品第N件特价真折扣修改商品明细
func (a *Activity) ActNitedSkuUpdate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NITED_SKU_UPDATE_URL, params).WithClient()
}

// ActNitedSkuUpdateBatch 商品第N件特价真折扣批量修改商品明细
func (a *Activity) ActNitedSkuUpdateBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NITED_SKU_UPDATE_BATCH_URL, params).WithClient()
}

// ActNitedCreate 商品第N件特价真折扣创建活动
func (a *Activity) ActNitedCreate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NITED_CREATE_URL, params).WithClient()
}

// ActXdiscSkuAddBatch 每满X件优惠批量添加商品明细
func (a *Activity) ActXdiscSkuAddBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_XDISC_SKU_ADD_BATCH_URL, params).WithClient()
}

// ActXdiscCreate 每满X件优惠创建活动
func (a *Activity) ActXdiscCreate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_XDISC_CREATE_URL, params).WithClient()
}

// ActXdiscModify 每满X件优惠修改活动信息
func (a *Activity) ActXdiscModify(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_XDISC_MODIFY_URL, params).WithClient()
}

// ActXdiscSkuAdd 添加每满X件优惠商品明细
func (a *Activity) ActXdiscSkuAdd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_XDISC_SKU_ADD_URL, params).WithClient()
}

// ActNchSkuAdd N选M活动添加商品明细
func (a *Activity) ActNchSkuAdd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NCH_SKU_ADD_URL, params).WithClient()
}

// ActNchCreate N选M活动创建活动
func (a *Activity) ActNchCreate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NCH_CREATE_URL, params).WithClient()
}

// ActNchModify N选M活动修改活动信息
func (a *Activity) ActNchModify(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NCH_MODIFY_URL, params).WithClient()
}

// ActNchSkuAddBatch N选M活动批量添加商品明细
func (a *Activity) ActNchSkuAddBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NCH_SKU_ADD_BATCH_URL, params).WithClient()
}

// ActNchSkuUpdateBatch N选M活动批量修改商品明细
func (a *Activity) ActNchSkuUpdateBatch(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_NCH_SKU_UPDATE_BATCH_URL, params).WithClient()
}

// ActShopactQ 店铺类活动查询活动信息
func (a *Activity) ActShopactQ(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_SHOPACT_Q_URL, params).WithClient()
}

// ActShopactDetailAddSingle 店铺类活动添加单个店铺
func (a *Activity) ActShopactDetailAddSingle(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_SHOPACT_DETAIL_ADD_SINGLE_URL, params).WithClient()
}

// ActShopactCreate 店铺类活动创建活动
func (a *Activity) ActShopactCreate(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_SHOPACT_CREATE_URL, params).WithClient()
}

// ActShopactDetailQ 店铺类活动查询活动明细信息
func (a *Activity) ActShopactDetailQ(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_SHOPACT_DETAIL_Q_URL, params).WithClient()
}

// ActShopactModify 店铺类活动信息修改
func (a *Activity) ActShopactModify(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_SHOPACT_MODIFY_URL, params).WithClient()
}

// ActShopactDetailDel 店铺类活动删除店铺
func (a *Activity) ActShopactDetailDel(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_SHOPACT_DETAIL_DEL_URL, params).WithClient()
}

// ActShopactDetailDelSingle 店铺类活动删除单个店铺
func (a *Activity) ActShopactDetailDelSingle(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_SHOPACT_DETAIL_DEL_SINGLE_URL, params).WithClient()
}

// ActShopactDetailAdd 店铺类活动添加店铺
func (a *Activity) ActShopactDetailAdd(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(activity.ACT_SHOPACT_DETAIL_ADD_URL, params).WithClient()
}
