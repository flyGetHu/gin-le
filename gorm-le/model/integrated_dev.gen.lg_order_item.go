package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderItemMgr struct {
	*_BaseMgr
}

// LgOrderItemMgr open func
func LgOrderItemMgr(db *gorm.DB) *_LgOrderItemMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderItemMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderItemMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_item"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderItemMgr) GetTableName() string {
	return "lg_order_item"
}

// Reset 重置gorm会话
func (obj *_LgOrderItemMgr) Reset() *_LgOrderItemMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderItemMgr) Get() (result LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderItemMgr) Gets() (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderItemMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgOrderItemMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 订单ID
func (obj *_LgOrderItemMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithNameEn name_en获取 英文品名
func (obj *_LgOrderItemMgr) WithNameEn(nameEn string) Option {
	return optionFunc(func(o *options) { o.query["name_en"] = nameEn })
}

// WithNameCn name_cn获取 中文品名
func (obj *_LgOrderItemMgr) WithNameCn(nameCn string) Option {
	return optionFunc(func(o *options) { o.query["name_cn"] = nameCn })
}

// WithDeclaredWeight declared_weight获取 申报重量
func (obj *_LgOrderItemMgr) WithDeclaredWeight(declaredWeight float64) Option {
	return optionFunc(func(o *options) { o.query["declared_weight"] = declaredWeight })
}

// WithDeclaredValue declared_value获取 申报价值
func (obj *_LgOrderItemMgr) WithDeclaredValue(declaredValue float64) Option {
	return optionFunc(func(o *options) { o.query["declared_value"] = declaredValue })
}

// WithQuantity quantity获取 数量
func (obj *_LgOrderItemMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithHscode hscode获取 HSCODE
func (obj *_LgOrderItemMgr) WithHscode(hscode string) Option {
	return optionFunc(func(o *options) { o.query["hscode"] = hscode })
}

// WithGoodsURL goods_url获取 商品链接
func (obj *_LgOrderItemMgr) WithGoodsURL(goodsURL string) Option {
	return optionFunc(func(o *options) { o.query["goods_url"] = goodsURL })
}

// WithSku sku获取 SKU
func (obj *_LgOrderItemMgr) WithSku(sku string) Option {
	return optionFunc(func(o *options) { o.query["sku"] = sku })
}

// WithLength length获取 长
func (obj *_LgOrderItemMgr) WithLength(length string) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithWidth width获取 宽
func (obj *_LgOrderItemMgr) WithWidth(width string) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

// WithHeight height获取 高
func (obj *_LgOrderItemMgr) WithHeight(height string) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithMaterial material获取 材质
func (obj *_LgOrderItemMgr) WithMaterial(material int) Option {
	return optionFunc(func(o *options) { o.query["material"] = material })
}

// WithPurpose purpose获取 用途
func (obj *_LgOrderItemMgr) WithPurpose(purpose string) Option {
	return optionFunc(func(o *options) { o.query["purpose"] = purpose })
}

// WithCaseNumber case_number获取 箱号
func (obj *_LgOrderItemMgr) WithCaseNumber(caseNumber string) Option {
	return optionFunc(func(o *options) { o.query["case_number"] = caseNumber })
}

// WithPlatformProductID platform_product_id获取 平台产品id
func (obj *_LgOrderItemMgr) WithPlatformProductID(platformProductID string) Option {
	return optionFunc(func(o *options) { o.query["platform_product_id"] = platformProductID })
}

// WithSpecs specs获取 规格
func (obj *_LgOrderItemMgr) WithSpecs(specs string) Option {
	return optionFunc(func(o *options) { o.query["specs"] = specs })
}

// WithPlaceOrigin place_origin获取 产地
func (obj *_LgOrderItemMgr) WithPlaceOrigin(placeOrigin string) Option {
	return optionFunc(func(o *options) { o.query["place_origin"] = placeOrigin })
}

// WithCurrency currency获取 币种
func (obj *_LgOrderItemMgr) WithCurrency(currency string) Option {
	return optionFunc(func(o *options) { o.query["currency"] = currency })
}

// WithCompany company获取 单位
func (obj *_LgOrderItemMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithPicturesURL pictures_url获取 图片链接
func (obj *_LgOrderItemMgr) WithPicturesURL(picturesURL string) Option {
	return optionFunc(func(o *options) { o.query["pictures_url"] = picturesURL })
}

// WithBrand brand获取 品牌
func (obj *_LgOrderItemMgr) WithBrand(brand string) Option {
	return optionFunc(func(o *options) { o.query["brand"] = brand })
}

// WithHasBattery has_battery获取 是否带电
func (obj *_LgOrderItemMgr) WithHasBattery(hasBattery int) Option {
	return optionFunc(func(o *options) { o.query["has_battery"] = hasBattery })
}

// WithLiquid liquid获取 是否液体,1-液体，0-非液体
func (obj *_LgOrderItemMgr) WithLiquid(liquid string) Option {
	return optionFunc(func(o *options) { o.query["liquid"] = liquid })
}

// WithPowder powder获取 是否粉末,1-粉末，0-非粉末
func (obj *_LgOrderItemMgr) WithPowder(powder string) Option {
	return optionFunc(func(o *options) { o.query["powder"] = powder })
}

// WithMagnetic magnetic获取 是否带磁,1-带磁，0-不带磁
func (obj *_LgOrderItemMgr) WithMagnetic(magnetic string) Option {
	return optionFunc(func(o *options) { o.query["magnetic"] = magnetic })
}

// WithDistributionInfo distribution_info获取 配货信息
func (obj *_LgOrderItemMgr) WithDistributionInfo(distributionInfo string) Option {
	return optionFunc(func(o *options) { o.query["distribution_info"] = distributionInfo })
}

// WithBatteryType battery_type获取 电池类型
func (obj *_LgOrderItemMgr) WithBatteryType(batteryType string) Option {
	return optionFunc(func(o *options) { o.query["battery_type"] = batteryType })
}

// WithSalesPlatform sales_platform获取 销售平台
func (obj *_LgOrderItemMgr) WithSalesPlatform(salesPlatform string) Option {
	return optionFunc(func(o *options) { o.query["sales_platform"] = salesPlatform })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderItemMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderItemMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderItemMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderItemMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderItemMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderItemMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderItemMgr) GetByOption(opts ...Option) (result LgOrderItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderItemMgr) GetByOptions(opts ...Option) (results []*LgOrderItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderItemMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderItem, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 ID
func (obj *_LgOrderItemMgr) GetFromID(id int) (result LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgOrderItemMgr) GetBatchFromID(ids []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单ID
func (obj *_LgOrderItemMgr) GetFromOrderID(orderID int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 订单ID
func (obj *_LgOrderItemMgr) GetBatchFromOrderID(orderIDs []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromNameEn 通过name_en获取内容 英文品名
func (obj *_LgOrderItemMgr) GetFromNameEn(nameEn string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`name_en` = ?", nameEn).Find(&results).Error

	return
}

// GetBatchFromNameEn 批量查找 英文品名
func (obj *_LgOrderItemMgr) GetBatchFromNameEn(nameEns []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`name_en` IN (?)", nameEns).Find(&results).Error

	return
}

// GetFromNameCn 通过name_cn获取内容 中文品名
func (obj *_LgOrderItemMgr) GetFromNameCn(nameCn string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`name_cn` = ?", nameCn).Find(&results).Error

	return
}

// GetBatchFromNameCn 批量查找 中文品名
func (obj *_LgOrderItemMgr) GetBatchFromNameCn(nameCns []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`name_cn` IN (?)", nameCns).Find(&results).Error

	return
}

// GetFromDeclaredWeight 通过declared_weight获取内容 申报重量
func (obj *_LgOrderItemMgr) GetFromDeclaredWeight(declaredWeight float64) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`declared_weight` = ?", declaredWeight).Find(&results).Error

	return
}

// GetBatchFromDeclaredWeight 批量查找 申报重量
func (obj *_LgOrderItemMgr) GetBatchFromDeclaredWeight(declaredWeights []float64) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`declared_weight` IN (?)", declaredWeights).Find(&results).Error

	return
}

// GetFromDeclaredValue 通过declared_value获取内容 申报价值
func (obj *_LgOrderItemMgr) GetFromDeclaredValue(declaredValue float64) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`declared_value` = ?", declaredValue).Find(&results).Error

	return
}

// GetBatchFromDeclaredValue 批量查找 申报价值
func (obj *_LgOrderItemMgr) GetBatchFromDeclaredValue(declaredValues []float64) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`declared_value` IN (?)", declaredValues).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 数量
func (obj *_LgOrderItemMgr) GetFromQuantity(quantity int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 数量
func (obj *_LgOrderItemMgr) GetBatchFromQuantity(quantitys []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromHscode 通过hscode获取内容 HSCODE
func (obj *_LgOrderItemMgr) GetFromHscode(hscode string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`hscode` = ?", hscode).Find(&results).Error

	return
}

// GetBatchFromHscode 批量查找 HSCODE
func (obj *_LgOrderItemMgr) GetBatchFromHscode(hscodes []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`hscode` IN (?)", hscodes).Find(&results).Error

	return
}

// GetFromGoodsURL 通过goods_url获取内容 商品链接
func (obj *_LgOrderItemMgr) GetFromGoodsURL(goodsURL string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`goods_url` = ?", goodsURL).Find(&results).Error

	return
}

// GetBatchFromGoodsURL 批量查找 商品链接
func (obj *_LgOrderItemMgr) GetBatchFromGoodsURL(goodsURLs []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`goods_url` IN (?)", goodsURLs).Find(&results).Error

	return
}

// GetFromSku 通过sku获取内容 SKU
func (obj *_LgOrderItemMgr) GetFromSku(sku string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`sku` = ?", sku).Find(&results).Error

	return
}

// GetBatchFromSku 批量查找 SKU
func (obj *_LgOrderItemMgr) GetBatchFromSku(skus []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`sku` IN (?)", skus).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容 长
func (obj *_LgOrderItemMgr) GetFromLength(length string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找 长
func (obj *_LgOrderItemMgr) GetBatchFromLength(lengths []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromWidth 通过width获取内容 宽
func (obj *_LgOrderItemMgr) GetFromWidth(width string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`width` = ?", width).Find(&results).Error

	return
}

// GetBatchFromWidth 批量查找 宽
func (obj *_LgOrderItemMgr) GetBatchFromWidth(widths []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`width` IN (?)", widths).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 高
func (obj *_LgOrderItemMgr) GetFromHeight(height string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`height` = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 高
func (obj *_LgOrderItemMgr) GetBatchFromHeight(heights []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`height` IN (?)", heights).Find(&results).Error

	return
}

// GetFromMaterial 通过material获取内容 材质
func (obj *_LgOrderItemMgr) GetFromMaterial(material int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`material` = ?", material).Find(&results).Error

	return
}

// GetBatchFromMaterial 批量查找 材质
func (obj *_LgOrderItemMgr) GetBatchFromMaterial(materials []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`material` IN (?)", materials).Find(&results).Error

	return
}

// GetFromPurpose 通过purpose获取内容 用途
func (obj *_LgOrderItemMgr) GetFromPurpose(purpose string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`purpose` = ?", purpose).Find(&results).Error

	return
}

// GetBatchFromPurpose 批量查找 用途
func (obj *_LgOrderItemMgr) GetBatchFromPurpose(purposes []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`purpose` IN (?)", purposes).Find(&results).Error

	return
}

// GetFromCaseNumber 通过case_number获取内容 箱号
func (obj *_LgOrderItemMgr) GetFromCaseNumber(caseNumber string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`case_number` = ?", caseNumber).Find(&results).Error

	return
}

// GetBatchFromCaseNumber 批量查找 箱号
func (obj *_LgOrderItemMgr) GetBatchFromCaseNumber(caseNumbers []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`case_number` IN (?)", caseNumbers).Find(&results).Error

	return
}

// GetFromPlatformProductID 通过platform_product_id获取内容 平台产品id
func (obj *_LgOrderItemMgr) GetFromPlatformProductID(platformProductID string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`platform_product_id` = ?", platformProductID).Find(&results).Error

	return
}

// GetBatchFromPlatformProductID 批量查找 平台产品id
func (obj *_LgOrderItemMgr) GetBatchFromPlatformProductID(platformProductIDs []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`platform_product_id` IN (?)", platformProductIDs).Find(&results).Error

	return
}

// GetFromSpecs 通过specs获取内容 规格
func (obj *_LgOrderItemMgr) GetFromSpecs(specs string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`specs` = ?", specs).Find(&results).Error

	return
}

// GetBatchFromSpecs 批量查找 规格
func (obj *_LgOrderItemMgr) GetBatchFromSpecs(specss []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`specs` IN (?)", specss).Find(&results).Error

	return
}

// GetFromPlaceOrigin 通过place_origin获取内容 产地
func (obj *_LgOrderItemMgr) GetFromPlaceOrigin(placeOrigin string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`place_origin` = ?", placeOrigin).Find(&results).Error

	return
}

// GetBatchFromPlaceOrigin 批量查找 产地
func (obj *_LgOrderItemMgr) GetBatchFromPlaceOrigin(placeOrigins []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`place_origin` IN (?)", placeOrigins).Find(&results).Error

	return
}

// GetFromCurrency 通过currency获取内容 币种
func (obj *_LgOrderItemMgr) GetFromCurrency(currency string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`currency` = ?", currency).Find(&results).Error

	return
}

// GetBatchFromCurrency 批量查找 币种
func (obj *_LgOrderItemMgr) GetBatchFromCurrency(currencys []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`currency` IN (?)", currencys).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 单位
func (obj *_LgOrderItemMgr) GetFromCompany(company string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量查找 单位
func (obj *_LgOrderItemMgr) GetBatchFromCompany(companys []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`company` IN (?)", companys).Find(&results).Error

	return
}

// GetFromPicturesURL 通过pictures_url获取内容 图片链接
func (obj *_LgOrderItemMgr) GetFromPicturesURL(picturesURL string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`pictures_url` = ?", picturesURL).Find(&results).Error

	return
}

// GetBatchFromPicturesURL 批量查找 图片链接
func (obj *_LgOrderItemMgr) GetBatchFromPicturesURL(picturesURLs []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`pictures_url` IN (?)", picturesURLs).Find(&results).Error

	return
}

// GetFromBrand 通过brand获取内容 品牌
func (obj *_LgOrderItemMgr) GetFromBrand(brand string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`brand` = ?", brand).Find(&results).Error

	return
}

// GetBatchFromBrand 批量查找 品牌
func (obj *_LgOrderItemMgr) GetBatchFromBrand(brands []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`brand` IN (?)", brands).Find(&results).Error

	return
}

// GetFromHasBattery 通过has_battery获取内容 是否带电
func (obj *_LgOrderItemMgr) GetFromHasBattery(hasBattery int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`has_battery` = ?", hasBattery).Find(&results).Error

	return
}

// GetBatchFromHasBattery 批量查找 是否带电
func (obj *_LgOrderItemMgr) GetBatchFromHasBattery(hasBatterys []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`has_battery` IN (?)", hasBatterys).Find(&results).Error

	return
}

// GetFromLiquid 通过liquid获取内容 是否液体,1-液体，0-非液体
func (obj *_LgOrderItemMgr) GetFromLiquid(liquid string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`liquid` = ?", liquid).Find(&results).Error

	return
}

// GetBatchFromLiquid 批量查找 是否液体,1-液体，0-非液体
func (obj *_LgOrderItemMgr) GetBatchFromLiquid(liquids []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`liquid` IN (?)", liquids).Find(&results).Error

	return
}

// GetFromPowder 通过powder获取内容 是否粉末,1-粉末，0-非粉末
func (obj *_LgOrderItemMgr) GetFromPowder(powder string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`powder` = ?", powder).Find(&results).Error

	return
}

// GetBatchFromPowder 批量查找 是否粉末,1-粉末，0-非粉末
func (obj *_LgOrderItemMgr) GetBatchFromPowder(powders []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`powder` IN (?)", powders).Find(&results).Error

	return
}

// GetFromMagnetic 通过magnetic获取内容 是否带磁,1-带磁，0-不带磁
func (obj *_LgOrderItemMgr) GetFromMagnetic(magnetic string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`magnetic` = ?", magnetic).Find(&results).Error

	return
}

// GetBatchFromMagnetic 批量查找 是否带磁,1-带磁，0-不带磁
func (obj *_LgOrderItemMgr) GetBatchFromMagnetic(magnetics []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`magnetic` IN (?)", magnetics).Find(&results).Error

	return
}

// GetFromDistributionInfo 通过distribution_info获取内容 配货信息
func (obj *_LgOrderItemMgr) GetFromDistributionInfo(distributionInfo string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`distribution_info` = ?", distributionInfo).Find(&results).Error

	return
}

// GetBatchFromDistributionInfo 批量查找 配货信息
func (obj *_LgOrderItemMgr) GetBatchFromDistributionInfo(distributionInfos []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`distribution_info` IN (?)", distributionInfos).Find(&results).Error

	return
}

// GetFromBatteryType 通过battery_type获取内容 电池类型
func (obj *_LgOrderItemMgr) GetFromBatteryType(batteryType string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`battery_type` = ?", batteryType).Find(&results).Error

	return
}

// GetBatchFromBatteryType 批量查找 电池类型
func (obj *_LgOrderItemMgr) GetBatchFromBatteryType(batteryTypes []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`battery_type` IN (?)", batteryTypes).Find(&results).Error

	return
}

// GetFromSalesPlatform 通过sales_platform获取内容 销售平台
func (obj *_LgOrderItemMgr) GetFromSalesPlatform(salesPlatform string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`sales_platform` = ?", salesPlatform).Find(&results).Error

	return
}

// GetBatchFromSalesPlatform 批量查找 销售平台
func (obj *_LgOrderItemMgr) GetBatchFromSalesPlatform(salesPlatforms []string) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`sales_platform` IN (?)", salesPlatforms).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderItemMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderItemMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderItemMgr) GetFromCreateUser(createUser int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderItemMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderItemMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderItemMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderItemMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderItemMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderItemMgr) GetFromVersion(version int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderItemMgr) GetBatchFromVersion(versions []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderItemMgr) GetFromDeleted(deleted int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderItemMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderItemMgr) FetchByPrimaryKey(id int) (result LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByLgOrderItemOrderIDIndex  获取多个内容
func (obj *_LgOrderItemMgr) FetchIndexByLgOrderItemOrderIDIndex(orderID int) (results []*LgOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderItem{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}
