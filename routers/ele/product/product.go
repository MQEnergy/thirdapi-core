package product

/*
商品管理
*/
const (
	SKU_CREATE_URL                    = "sku.create"                        // 单个创建商品
	SKU_UPDATE_URL                    = "sku.update"                        // 单个商品修改
	BATCH_SKU_CREATE_URL              = "batch.sku.create"                  // 批量创建渠道商品
	BATCH_SKU_UPDATE_URL              = "batch.sku.update"                  // 批量更新渠道商品
	SKU_LIST_URL                      = "sku.list"                          // 商品列表
	SKU_DELETE_URL                    = "sku.delete"                        // 批量删除商品
	SKU_SHOP_CUSTOMSKU_MAP_URL        = "sku.shop.customsku.map"            // 绑定商品与商品自定义ID
	SKU_BRAND_LIST_URL                = "sku.brand.list"                    // 获取品牌列表
	SKU_STDUPC_EXIST_URL              = "sku.stdupc.exist"                  // 根据条形码查询是否平台药品SPU
	SKU_SPEC_UPDATE_APPEND_URL        = "sku.spec.update.append"            // 渠道商品追加多规格sku
	SKU_STDUPC_BARCODE_URL            = "sku.stdupc.barcode"                // 查询平台商品条码
	SKU_SHOP_AUTH_CONTROL_URL         = "sku.shop.auth.control"             // 商品店铺权限控制
	SKU_SHOP_AUTH_GET_URL             = "sku.shop.auth.get"                 // 商品店铺权限查询
	SKU_COPY_URL                      = "sku.copy"                          // 商品复制
	SKU_JOB_LIST_URL                  = "sku.job.list"                      // 查询异步任务列表
	SKU_STOCK_LIST_URL                = "sku.out.of.stock.list"             // 获取门店商品缺货数据
	SKU_CATEGORY_LIST_URL             = "sku.category.list"                 // 获取商品类目列表
	SKU_CATEGORY_PROPERTY_LIST_URL    = "sku.category.property.list"        // 查询类目属性值列表
	SKU_CATEGORY_PROPERTY_QUERY_URL   = "sku.category.property.query"       // 查询渠道商品的类目属性值
	SKU_CATEGORY_PROPERTY_PREDICT_URL = "sku.category.property.predict"     // 根据商品条码或商品名获取商品类目属性值
	SKU_STOCK_UPDATE_ONE_URL          = "sku.stock.update.one"              // 单个修改商品库存
	SKU_STOCK_UPDATE_BATCH_URL        = "sku.stock.update.batch"            // 批量修改商品库存
	SKU_STOCK_INCREASE_ONE_URL        = "sku.stock.increase.one"            // 增量更新单个商品库存
	SKU_SPEC_STOCK_UPDATE_URL         = "sku.spec.stock.update"             // 商品多规格修改库存
	SKU_WAREHOUSE_STOCK_UPDATE_URL    = "sku.warehouse.stock.update"        // 修改仓货品库存
	SKU_UPDATE_STOCK_NODE_ID_URL      = "sku.update.stock.node.identify"    // 商品库存同步节点识别
	SKU_SPEC_STOCK_INCREASE_URL       = "sku.spec.stock.increase"           // 增量更新商品多规格库存
	SKU_SHOP_CATEGORY_MAP_URL         = "sku.shop.category.map"             // 绑定商品与店铺内分类
	SKU_SHOP_CATEGORY_GET_URL         = "sku.shop.category.get"             // 获取店铺内分类
	SKU_SHOP_CATEGORY_CREATE_URL      = "sku.shop.category.create"          // 新增店铺内分类
	SKU_SHOP_CATEGORY_UPDATE_URL      = "sku.shop.category.update"          // 修改店铺内分类
	SKU_SHOP_CATEGORY_DELETE_URL      = "sku.shop.category.delete"          // 删除店铺内分类
	SKU_SHOP_CUSTOMSKU_LIST_URL       = "sku.shop.customsku.list"           // 获取店铺内分类下商品
	SKU_ONLINE_ONE_URL                = "sku.online.one"                    // 单个商品上架
	SKU_OFFLINE_ONE_URL               = "sku.offline.one"                   // 单个商品下架
	SKU_ONLINE_URL                    = "sku.online"                        // 批量商品上架
	SKU_OFFLINE_URL                   = "sku.offline"                       // 批量商品下架
	SKU_PRICE_UPDATE_BATCH_URL        = "sku.price.update.batch"            // 批量修改商品价格
	SKU_PRICE_UPDATE_ONE_URL          = "sku.price.update.one"              // 单个修改商品价格
	SKU_SPEC_UPDATE_PRICE_URL         = "sku.spec.update.price"             // 更新渠道商品多规格价格
	PRODUCT_LIST_URL                  = "product.list"                      // 查询连锁商品
	PRODUCT_REM_SEC_BC_TAG_URL        = "product.remove.second.barcode.tag" // 解除商家产品和关联的所有渠道商品与特定类目的绑定关系
	PRODUCT_UPDATE_URL                = "product.update"                    // 修改连锁商品
	SKU_SELLER_AUTH_GET_URL           = "sku.seller.auth.get"               // 查询连锁权限
	PRODUCT_CUSTOM_MAP_URL            = "product.custom.map"                // 绑定商家商品与自定义商品ID
	PRODUCT_DELETE_URL                = "product.delete"                    // 删除连锁商品
	PRODUCT_SYNC_ITEM_URL             = "product.sync.item"                 // 连锁商品同步至连锁下多门店
	BATCH_PRODUCT_UPDATE_URL          = "batch.product.update"              // 跨连锁批量修改商家产品
	BATCH_PRODUCT_DELETE_URL          = "batch.product.delete"              // 跨连锁批量删除商家产品
	PRODUCT_CREATE_URL                = "product.create"                    // 创建连锁商品
	SKU_SELLER_CATE_CREATE_URL        = "sku.seller.category.create"        // 创建连锁商品前台分类
	SKU_SELLER_CATE_DELETE_URL        = "sku.seller.category.delete"        // 删除连锁商品前台分类
	SKU_SELLER_CATE_MAP_URL           = "sku.seller.category.map"           // 绑定或更新连锁商品到前台分类
	SKU_SELLER_CATE_UPDATE_URL        = "sku.seller.category.update"        // 修改连锁商品前台分类
	SKU_SELLER_CATE_PRODUCTS_URL      = "sku.seller.category.products"      // 查询分类下连锁商品详情列表
	SKU_SELLER_CATE_SYNC_URL          = "sku.seller.category.sync"          // 同步连锁前台分类到子门店
	SKU_SELLER_CATE_GET_URL           = "sku.seller.category.get"           // 连锁商品前台分类列表
)
