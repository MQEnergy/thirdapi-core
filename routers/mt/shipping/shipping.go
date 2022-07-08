package shipping

/*
配送类
*/
const (
	SAVE_URL       = "shipping/save"      // 创建/更新门店配送范围
	LIST_URL       = "shipping/list"      // 获取门店配送范围（自配送）
	BATCH_SAVE_URL = "shipping/batchsave" // 批量创建/更新配送范围
	FETCH_URL      = "shipping/fetch"     // 获取门店配送范围（混合送）
	DELETE_URL     = "shipping/delete"    // 删除配送范围
	SPEC_SAVE_URL  = "shipping/spec/save" // 新增/更新特殊时段配送范围
)
