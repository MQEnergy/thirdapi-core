package sunny

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/sunny"
)

type Sunny struct {
	Base *ele.Base
}

// New 实例化
func New(source, appSecret string) *Sunny {
	return &Sunny{
		Base: ele.New(source, appSecret),
	}
}

// RelationPriceUpdate 供货价修改
func (s *Sunny) RelationPriceUpdate(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_RELATION_PRICE_UPDATE_URL, params).WithClient()
}

// CartsAdd 阳光菜场加购
func (s *Sunny) CartsAdd(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_CARTS_ADD_URL, params).WithClient()
}

// CouponActAttend 零售发券
func (s *Sunny) CouponActAttend(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.COUPON_ACT_ATTEND_URL, params).WithClient()
}

// RelationDelete 供货关系删除
func (s *Sunny) RelationDelete(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_RELATION_DELETE_URL, params).WithClient()
}

// ActivityIdsGetact 营销活动补全(多维度)
func (s *Sunny) ActivityIdsGetact(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.ACTIVITY_IDS_GETACT_URL, params).WithClient()
}

// SunnyItemRecallByAct 菜场商品召回
func (s *Sunny) SunnyItemRecallByAct(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_ITEM_RECALLBYACT_URL, params).WithClient()
}

// SunnyStallUpdate 摊位信息更新
func (s *Sunny) SunnyStallUpdate(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_STALL_UPDATE_URL, params).WithClient()
}

// SunnyRelationCreate 创建供货关系
func (s *Sunny) SunnyRelationCreate(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_RELATION_CREATE_URL, params).WithClient()
}

// SunnyRelationQuery 供货关系查询
func (s *Sunny) SunnyRelationQuery(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_RELATION_QUERY_URL, params).WithClient()
}

// SunnyRelationQuerySingle 单条供货关系查询
func (s *Sunny) SunnyRelationQuerySingle(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_RELATION_QUERY_SINGLE_URL, params).WithClient()
}

// SunnyStallCreate 摊位信息增加
func (s *Sunny) SunnyStallCreate(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_STALL_CREATE_URL, params).WithClient()
}

// SunnyRelationStatusUpdate 供货关系状态更新
func (s *Sunny) SunnyRelationStatusUpdate(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_RELATION_STATUS_UPDATE_URL, params).WithClient()
}

// SunnyStallQuery 摊位信息查询
func (s *Sunny) SunnyStallQuery(params map[string]interface{}) (string, error) {
	return s.Base.WithRequestParams(sunny.SUNNY_STALL_QUERY_URL, params).WithClient()
}
