package ads

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/ads"
)

type Ads struct {
	Base *ele.Base
}

func New(source, secret string) *Ads {
	return &Ads{
		Base: ele.New(source, secret),
	}
}

// AdFindAccountBalance 查询斗金最大可用余额
func (a *Ads) AdFindAccountBalance(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_FIND_ACCOUNT_BALANCE_URL, params).WithClient()
}

// AdFindRoiEffectReport 查询斗金单店分时ROI效果数据
func (a *Ads) AdFindRoiEffectReport(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_FIND_ROI_EFFECT_REPORT_URL, params).WithClient()
}

// AdFindDistReport 查询斗金推广点击分布信息
func (a *Ads) AdFindDistReport(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_FIND_DIST_REPORT_URL, params).WithClient()
}

// AdUpdateTime 设置斗金时段
func (a *Ads) AdUpdateTime(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_UPDATE_TIME_URL, params).WithClient()
}

// AdFindEffectReport 查询斗金推广效果数据
func (a *Ads) AdFindEffectReport(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_FIND_EFFECT_REPORT_URL, params).WithClient()
}

// AdUpdateState 设置斗金状态
func (a *Ads) AdUpdateState(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_UPDATE_STATE_URL, params).WithClient()
}

// AdUpdateBudget 设置斗金预算
func (a *Ads) AdUpdateBudget(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_UPDATE_BUDGET_URL, params).WithClient()
}

// AdFindCpcSolution 查询斗金推广设置
func (a *Ads) AdFindCpcSolution(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_FIND_CPC_SOLUTION_URL, params).WithClient()
}

// AdUpdateBid 设置斗金出价
func (a *Ads) AdUpdateBid(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_UPDATE_BID_URL, params).WithClient()
}

// AdCreateSolution 创建斗金计划
func (a *Ads) AdCreateSolution(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_CREATE_SOLUTION_URL, params).WithClient()
}

// AdUpdateTarget 设置斗金定向
func (a *Ads) AdUpdateTarget(params map[string]interface{}) (string, error) {
	return a.Base.WithRequestParams(ads.AD_DJCPC_UPDATE_TARGET_URL, params).WithClient()
}
