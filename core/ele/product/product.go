package product

import (
	"github.com/MQEnergy/thirdapi-core/core/ele"
	"github.com/MQEnergy/thirdapi-core/routers/ele/product"
)

type Product struct {
	Base *ele.Base
}

// New 实例化
func New(appId, appSecret string) *Product {
	return &Product{
		Base: ele.New(appId, appSecret),
	}
}

// SkuCreate 单个创建商品
func (c *Product) SkuCreate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_CREATE_URL, params).WithClient()
}

// SkuUpdate 单个商品修改
func (c *Product) SkuUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_UPDATE_URL, params).WithClient()
}

// SkuBatchCreate 批量创建渠道商品
func (c *Product) SkuBatchCreate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.BATCH_SKU_CREATE_URL, params).WithClient()
}

// SkuBatchUpdate 批量更新渠道商品
func (c *Product) SkuBatchUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.BATCH_SKU_UPDATE_URL, params).WithClient()
}

// SkuList 商品列表
func (c *Product) SkuList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_LIST_URL, params).WithClient()
}

// SkuDelete 批量删除商品
func (c *Product) SkuDelete(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_DELETE_URL, params).WithClient()
}

// SkuShopCustomskuMap 绑定商品与商品自定义ID
func (c *Product) SkuShopCustomskuMap(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_CUSTOMSKU_MAP_URL, params).WithClient()
}

// SkuBrandList 获取品牌列表
func (c *Product) SkuBrandList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_BRAND_LIST_URL, params).WithClient()
}

// SkuStdupcExist 根据条形码查询是否平台药品SPU
func (c *Product) SkuStdupcExist(params map[string]interface{}) (string, error) {
	data := map[string]interface{}{
		"shop_id": params["shop_id"],
		"upc":     params["upc"],
	}
	return c.Base.WithRequestParams(product.SKU_STDUPC_EXIST_URL, data).WithClient()
}

// SkuSpecUpdateAppend 渠道商品追加多规格sku
func (c *Product) SkuSpecUpdateAppend(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SPEC_UPDATE_APPEND_URL, params).WithClient()
}

// SkuStdupcBarcode 查询平台商品条码
func (c *Product) SkuStdupcBarcode(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_STDUPC_BARCODE_URL, params).WithClient()
}

// SkuShopAuthControl 商品店铺权限控制
func (c *Product) SkuShopAuthControl(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_AUTH_CONTROL_URL, params).WithClient()
}

// SkuShopAuthGet 商品店铺权限查询
func (c *Product) SkuShopAuthGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_AUTH_GET_URL, params).WithClient()
}

// SkuCopy 商品复制
func (c *Product) SkuCopy(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_COPY_URL, params).WithClient()
}

// SkuJobList 查询异步任务列表
func (c *Product) SkuJobList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_JOB_LIST_URL, params).WithClient()
}

// SkuStockList 获取门店商品缺货数据
func (c *Product) SkuStockList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_STOCK_LIST_URL, params).WithClient()
}

// SkuCategoryList 获取商品类目列表
func (c *Product) SkuCategoryList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_CATEGORY_LIST_URL, params).WithClient()
}

// SkuCategoryPropertyList 查询类目属性值列表
func (c *Product) SkuCategoryPropertyList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_CATEGORY_PROPERTY_LIST_URL, params).WithClient()
}

// SkuCategoryPropertyQuery 查询渠道商品的类目属性值
func (c *Product) SkuCategoryPropertyQuery(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_CATEGORY_PROPERTY_QUERY_URL, params).WithClient()
}

// SkuCategoryPropertyPredict 根据商品条码或商品名获取商品类目属性值
func (c *Product) SkuCategoryPropertyPredict(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_CATEGORY_PROPERTY_PREDICT_URL, params).WithClient()
}

// SkuStockUpdateOne 单个修改商品库存
func (c *Product) SkuStockUpdateOne(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_STOCK_UPDATE_ONE_URL, params).WithClient()
}

// SkuStockUpdateBatch 批量修改商品库存
func (c *Product) SkuStockUpdateBatch(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_STOCK_UPDATE_BATCH_URL, params).WithClient()
}

// SkuStockIncreaseOne 增量更新单个商品库存
func (c *Product) SkuStockIncreaseOne(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_STOCK_INCREASE_ONE_URL, params).WithClient()
}

// SkuSpecStockUpdate 商品多规格修改库存
func (c *Product) SkuSpecStockUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SPEC_STOCK_UPDATE_URL, params).WithClient()
}

// SkuWarehouseStockUpdate 修改仓货品库存
func (c *Product) SkuWarehouseStockUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_WAREHOUSE_STOCK_UPDATE_URL, params).WithClient()
}

// SkuUpdateStockNodeId 更新商品库存同步节点识别
func (c *Product) SkuUpdateStockNodeId(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_UPDATE_STOCK_NODE_ID_URL, params).WithClient()
}

// SkuSpecStockIncrease 增量更新商品多规格库存
func (c *Product) SkuSpecStockIncrease(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SPEC_STOCK_INCREASE_URL, params).WithClient()
}

// SkuShopCategoryMap 绑定商品与店铺内分类
func (c *Product) SkuShopCategoryMap(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_CATEGORY_MAP_URL, params).WithClient()
}

// SkuShopCategoryGet 获取店铺内分类
func (c *Product) SkuShopCategoryGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_CATEGORY_GET_URL, params).WithClient()
}

// SkuShopCategoryCreate 新增店铺内分类
func (c *Product) SkuShopCategoryCreate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_CATEGORY_CREATE_URL, params).WithClient()
}

// SkuShopCategoryUpdate 修改店铺内分类
func (c *Product) SkuShopCategoryUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_CATEGORY_UPDATE_URL, params).WithClient()
}

// SkuShopCategoryDelete 删除店铺内分类
func (c *Product) SkuShopCategoryDelete(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_CATEGORY_DELETE_URL, params).WithClient()
}

// SkuShopCustomskuList 获取店铺内分类下商品
func (c *Product) SkuShopCustomskuList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SHOP_CUSTOMSKU_LIST_URL, params).WithClient()
}

// SkuOnlineOne 单个商品上架
func (c *Product) SkuOnlineOne(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_ONLINE_ONE_URL, params).WithClient()
}

// SkuOfflineOne 单个商品下架
func (c *Product) SkuOfflineOne(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_OFFLINE_ONE_URL, params).WithClient()
}

// SkuOnline 批量商品上架
func (c *Product) SkuOnline(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_ONLINE_URL, params).WithClient()
}

// SkuOffline 批量商品下架
func (c *Product) SkuOffline(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_OFFLINE_URL, params).WithClient()
}

// SkuPriceUpdateBatch 批量修改商品价格
func (c *Product) SkuPriceUpdateBatch(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_PRICE_UPDATE_BATCH_URL, params).WithClient()
}

// SkuPriceUpdateOne 单个修改商品价格
func (c *Product) SkuPriceUpdateOne(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_PRICE_UPDATE_ONE_URL, params).WithClient()
}

// SkuSpecUpdatePrice 更新渠道商品多规格价格
func (c *Product) SkuSpecUpdatePrice(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SPEC_UPDATE_PRICE_URL, params).WithClient()
}

// ProductList 查询连锁商品
func (c *Product) ProductList(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.PRODUCT_LIST_URL, params).WithClient()
}

// ProductRemSecBcTag 解除商家产品和关联的所有渠道商品与特定类目的绑定关系
func (c *Product) ProductRemSecBcTag(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.PRODUCT_REM_SEC_BC_TAG_URL, params).WithClient()
}

// ProductUpdate 修改连锁商品
func (c *Product) ProductUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.PRODUCT_UPDATE_URL, params).WithClient()
}

// SkuSellerAuthGet 查询连锁权限
func (c *Product) SkuSellerAuthGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SELLER_AUTH_GET_URL, params).WithClient()
}

// ProductCustomMap 绑定商家商品与自定义商品ID
func (c *Product) ProductCustomMap(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.PRODUCT_CUSTOM_MAP_URL, params).WithClient()
}

// ProductDelete 删除连锁商品
func (c *Product) ProductDelete(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.PRODUCT_DELETE_URL, params).WithClient()
}

// ProductSyncItem 连锁商品同步至连锁下多门店
func (c *Product) ProductSyncItem(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.PRODUCT_SYNC_ITEM_URL, params).WithClient()
}

// BatchProductUpdate 跨连锁批量修改商家产品
func (c *Product) BatchProductUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.BATCH_PRODUCT_UPDATE_URL, params).WithClient()
}

// BatchProductDelete 跨连锁批量删除商家产品
func (c *Product) BatchProductDelete(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.BATCH_PRODUCT_DELETE_URL, params).WithClient()
}

// ProductCreate 创建连锁商品
func (c *Product) ProductCreate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.PRODUCT_CREATE_URL, params).WithClient()
}

// SkuSellerCateCreate 创建连锁商品前台分类
func (c *Product) SkuSellerCateCreate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SELLER_CATE_CREATE_URL, params).WithClient()
}

// SkuSellerCateDelete 删除连锁商品前台分类
func (c *Product) SkuSellerCateDelete(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SELLER_CATE_DELETE_URL, params).WithClient()
}

// SkuSellerCateMap 绑定或更新连锁商品到前台分类
func (c *Product) SkuSellerCateMap(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SELLER_CATE_MAP_URL, params).WithClient()
}

// SkuSellerCateUpdate 修改连锁商品前台分类
func (c *Product) SkuSellerCateUpdate(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SELLER_CATE_UPDATE_URL, params).WithClient()
}

// SkuSellerCateProducts 查询分类下连锁商品详情列表
func (c *Product) SkuSellerCateProducts(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SELLER_CATE_PRODUCTS_URL, params).WithClient()
}

// SkuSellerCateSync 同步连锁前台分类到子门店
func (c *Product) SkuSellerCateSync(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SELLER_CATE_SYNC_URL, params).WithClient()
}

// SkuSellerCateGet 连锁商品前台分类列表
func (c *Product) SkuSellerCateGet(params map[string]interface{}) (string, error) {
	return c.Base.WithRequestParams(product.SKU_SELLER_CATE_GET_URL, params).WithClient()
}
