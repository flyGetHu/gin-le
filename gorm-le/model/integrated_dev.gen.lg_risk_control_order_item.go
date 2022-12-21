package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgRiskControlOrderItemMgr struct {
	*_BaseMgr
}

// LgRiskControlOrderItemMgr open func
func LgRiskControlOrderItemMgr(db *gorm.DB) *_LgRiskControlOrderItemMgr {
	if db == nil {
		panic(fmt.Errorf("LgRiskControlOrderItemMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgRiskControlOrderItemMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_risk_control_order_item"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgRiskControlOrderItemMgr) GetTableName() string {
	return "lg_risk_control_order_item"
}

// Reset 重置gorm会话
func (obj *_LgRiskControlOrderItemMgr) Reset() *_LgRiskControlOrderItemMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgRiskControlOrderItemMgr) Get() (result LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgRiskControlOrderItemMgr) Gets() (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgRiskControlOrderItemMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgRiskControlOrderItemMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRiskControlOrderID risk_control_order_id获取 订单编号
func (obj *_LgRiskControlOrderItemMgr) WithRiskControlOrderID(riskControlOrderID int64) Option {
	return optionFunc(func(o *options) { o.query["risk_control_order_id"] = riskControlOrderID })
}

// WithNameEn name_en获取 英文品名
func (obj *_LgRiskControlOrderItemMgr) WithNameEn(nameEn string) Option {
	return optionFunc(func(o *options) { o.query["name_en"] = nameEn })
}

// WithNameCn name_cn获取 中文品名
func (obj *_LgRiskControlOrderItemMgr) WithNameCn(nameCn string) Option {
	return optionFunc(func(o *options) { o.query["name_cn"] = nameCn })
}

// WithDeclaredWeight declared_weight获取 申报重量
func (obj *_LgRiskControlOrderItemMgr) WithDeclaredWeight(declaredWeight float64) Option {
	return optionFunc(func(o *options) { o.query["declared_weight"] = declaredWeight })
}

// WithDeclaredValue declared_value获取 申报价值
func (obj *_LgRiskControlOrderItemMgr) WithDeclaredValue(declaredValue float64) Option {
	return optionFunc(func(o *options) { o.query["declared_value"] = declaredValue })
}

// WithQuantity quantity获取 数量
func (obj *_LgRiskControlOrderItemMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithHscode hscode获取 HSCODE
func (obj *_LgRiskControlOrderItemMgr) WithHscode(hscode string) Option {
	return optionFunc(func(o *options) { o.query["hscode"] = hscode })
}

// WithGoodsURL goods_url获取 商品链接
func (obj *_LgRiskControlOrderItemMgr) WithGoodsURL(goodsURL string) Option {
	return optionFunc(func(o *options) { o.query["goods_url"] = goodsURL })
}

// WithSku sku获取 SKU
func (obj *_LgRiskControlOrderItemMgr) WithSku(sku string) Option {
	return optionFunc(func(o *options) { o.query["sku"] = sku })
}

// WithLength length获取 长
func (obj *_LgRiskControlOrderItemMgr) WithLength(length string) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithWidth width获取 宽
func (obj *_LgRiskControlOrderItemMgr) WithWidth(width string) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

// WithHeight height获取 高
func (obj *_LgRiskControlOrderItemMgr) WithHeight(height string) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithMaterial material获取 材质
func (obj *_LgRiskControlOrderItemMgr) WithMaterial(material int) Option {
	return optionFunc(func(o *options) { o.query["material"] = material })
}

// WithPurpose purpose获取 用途
func (obj *_LgRiskControlOrderItemMgr) WithPurpose(purpose string) Option {
	return optionFunc(func(o *options) { o.query["purpose"] = purpose })
}

// WithCaseNumber case_number获取 箱号
func (obj *_LgRiskControlOrderItemMgr) WithCaseNumber(caseNumber string) Option {
	return optionFunc(func(o *options) { o.query["case_number"] = caseNumber })
}

// WithPlatformProductID platform_product_id获取 平台产品id
func (obj *_LgRiskControlOrderItemMgr) WithPlatformProductID(platformProductID string) Option {
	return optionFunc(func(o *options) { o.query["platform_product_id"] = platformProductID })
}

// WithSpecs specs获取 规格
func (obj *_LgRiskControlOrderItemMgr) WithSpecs(specs string) Option {
	return optionFunc(func(o *options) { o.query["specs"] = specs })
}

// WithPlaceOrigin place_origin获取 产地
func (obj *_LgRiskControlOrderItemMgr) WithPlaceOrigin(placeOrigin string) Option {
	return optionFunc(func(o *options) { o.query["place_origin"] = placeOrigin })
}

// WithCurrency currency获取 币种
func (obj *_LgRiskControlOrderItemMgr) WithCurrency(currency string) Option {
	return optionFunc(func(o *options) { o.query["currency"] = currency })
}

// WithCompany company获取 单位
func (obj *_LgRiskControlOrderItemMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithPicturesURL pictures_url获取 图片链接
func (obj *_LgRiskControlOrderItemMgr) WithPicturesURL(picturesURL string) Option {
	return optionFunc(func(o *options) { o.query["pictures_url"] = picturesURL })
}

// WithBrand brand获取 品牌
func (obj *_LgRiskControlOrderItemMgr) WithBrand(brand string) Option {
	return optionFunc(func(o *options) { o.query["brand"] = brand })
}

// WithHasBattery has_battery获取 是否带电
func (obj *_LgRiskControlOrderItemMgr) WithHasBattery(hasBattery int) Option {
	return optionFunc(func(o *options) { o.query["has_battery"] = hasBattery })
}

// WithLiquid liquid获取 是否液体,1-液体，0-非液体
func (obj *_LgRiskControlOrderItemMgr) WithLiquid(liquid string) Option {
	return optionFunc(func(o *options) { o.query["liquid"] = liquid })
}

// WithPowder powder获取 是否粉末,1-粉末，0-非粉末
func (obj *_LgRiskControlOrderItemMgr) WithPowder(powder string) Option {
	return optionFunc(func(o *options) { o.query["powder"] = powder })
}

// WithMagnetic magnetic获取 是否带磁,1-带磁，0-不带磁
func (obj *_LgRiskControlOrderItemMgr) WithMagnetic(magnetic string) Option {
	return optionFunc(func(o *options) { o.query["magnetic"] = magnetic })
}

// WithDistributionInfo distribution_info获取 配货信息
func (obj *_LgRiskControlOrderItemMgr) WithDistributionInfo(distributionInfo string) Option {
	return optionFunc(func(o *options) { o.query["distribution_info"] = distributionInfo })
}

// WithBatteryType battery_type获取 电池类型
func (obj *_LgRiskControlOrderItemMgr) WithBatteryType(batteryType string) Option {
	return optionFunc(func(o *options) { o.query["battery_type"] = batteryType })
}

// WithSalesPlatform sales_platform获取 销售平台
func (obj *_LgRiskControlOrderItemMgr) WithSalesPlatform(salesPlatform string) Option {
	return optionFunc(func(o *options) { o.query["sales_platform"] = salesPlatform })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgRiskControlOrderItemMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgRiskControlOrderItemMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgRiskControlOrderItemMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgRiskControlOrderItemMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgRiskControlOrderItemMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgRiskControlOrderItemMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgRiskControlOrderItemMgr) GetByOption(opts ...Option) (result LgRiskControlOrderItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgRiskControlOrderItemMgr) GetByOptions(opts ...Option) (results []*LgRiskControlOrderItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgRiskControlOrderItemMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgRiskControlOrderItem, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where(options.query)
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
func (obj *_LgRiskControlOrderItemMgr) GetFromID(id int) (result LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromID(ids []int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromRiskControlOrderID 通过risk_control_order_id获取内容 订单编号
func (obj *_LgRiskControlOrderItemMgr) GetFromRiskControlOrderID(riskControlOrderID int64) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`risk_control_order_id` = ?", riskControlOrderID).Find(&results).Error

	return
}

// GetBatchFromRiskControlOrderID 批量查找 订单编号
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromRiskControlOrderID(riskControlOrderIDs []int64) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`risk_control_order_id` IN (?)", riskControlOrderIDs).Find(&results).Error

	return
}

// GetFromNameEn 通过name_en获取内容 英文品名
func (obj *_LgRiskControlOrderItemMgr) GetFromNameEn(nameEn string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`name_en` = ?", nameEn).Find(&results).Error

	return
}

// GetBatchFromNameEn 批量查找 英文品名
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromNameEn(nameEns []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`name_en` IN (?)", nameEns).Find(&results).Error

	return
}

// GetFromNameCn 通过name_cn获取内容 中文品名
func (obj *_LgRiskControlOrderItemMgr) GetFromNameCn(nameCn string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`name_cn` = ?", nameCn).Find(&results).Error

	return
}

// GetBatchFromNameCn 批量查找 中文品名
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromNameCn(nameCns []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`name_cn` IN (?)", nameCns).Find(&results).Error

	return
}

// GetFromDeclaredWeight 通过declared_weight获取内容 申报重量
func (obj *_LgRiskControlOrderItemMgr) GetFromDeclaredWeight(declaredWeight float64) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`declared_weight` = ?", declaredWeight).Find(&results).Error

	return
}

// GetBatchFromDeclaredWeight 批量查找 申报重量
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromDeclaredWeight(declaredWeights []float64) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`declared_weight` IN (?)", declaredWeights).Find(&results).Error

	return
}

// GetFromDeclaredValue 通过declared_value获取内容 申报价值
func (obj *_LgRiskControlOrderItemMgr) GetFromDeclaredValue(declaredValue float64) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`declared_value` = ?", declaredValue).Find(&results).Error

	return
}

// GetBatchFromDeclaredValue 批量查找 申报价值
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromDeclaredValue(declaredValues []float64) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`declared_value` IN (?)", declaredValues).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 数量
func (obj *_LgRiskControlOrderItemMgr) GetFromQuantity(quantity int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 数量
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromQuantity(quantitys []int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromHscode 通过hscode获取内容 HSCODE
func (obj *_LgRiskControlOrderItemMgr) GetFromHscode(hscode string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`hscode` = ?", hscode).Find(&results).Error

	return
}

// GetBatchFromHscode 批量查找 HSCODE
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromHscode(hscodes []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`hscode` IN (?)", hscodes).Find(&results).Error

	return
}

// GetFromGoodsURL 通过goods_url获取内容 商品链接
func (obj *_LgRiskControlOrderItemMgr) GetFromGoodsURL(goodsURL string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`goods_url` = ?", goodsURL).Find(&results).Error

	return
}

// GetBatchFromGoodsURL 批量查找 商品链接
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromGoodsURL(goodsURLs []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`goods_url` IN (?)", goodsURLs).Find(&results).Error

	return
}

// GetFromSku 通过sku获取内容 SKU
func (obj *_LgRiskControlOrderItemMgr) GetFromSku(sku string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`sku` = ?", sku).Find(&results).Error

	return
}

// GetBatchFromSku 批量查找 SKU
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromSku(skus []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`sku` IN (?)", skus).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容 长
func (obj *_LgRiskControlOrderItemMgr) GetFromLength(length string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找 长
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromLength(lengths []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromWidth 通过width获取内容 宽
func (obj *_LgRiskControlOrderItemMgr) GetFromWidth(width string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`width` = ?", width).Find(&results).Error

	return
}

// GetBatchFromWidth 批量查找 宽
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromWidth(widths []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`width` IN (?)", widths).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 高
func (obj *_LgRiskControlOrderItemMgr) GetFromHeight(height string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`height` = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 高
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromHeight(heights []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`height` IN (?)", heights).Find(&results).Error

	return
}

// GetFromMaterial 通过material获取内容 材质
func (obj *_LgRiskControlOrderItemMgr) GetFromMaterial(material int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`material` = ?", material).Find(&results).Error

	return
}

// GetBatchFromMaterial 批量查找 材质
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromMaterial(materials []int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`material` IN (?)", materials).Find(&results).Error

	return
}

// GetFromPurpose 通过purpose获取内容 用途
func (obj *_LgRiskControlOrderItemMgr) GetFromPurpose(purpose string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`purpose` = ?", purpose).Find(&results).Error

	return
}

// GetBatchFromPurpose 批量查找 用途
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromPurpose(purposes []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`purpose` IN (?)", purposes).Find(&results).Error

	return
}

// GetFromCaseNumber 通过case_number获取内容 箱号
func (obj *_LgRiskControlOrderItemMgr) GetFromCaseNumber(caseNumber string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`case_number` = ?", caseNumber).Find(&results).Error

	return
}

// GetBatchFromCaseNumber 批量查找 箱号
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromCaseNumber(caseNumbers []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`case_number` IN (?)", caseNumbers).Find(&results).Error

	return
}

// GetFromPlatformProductID 通过platform_product_id获取内容 平台产品id
func (obj *_LgRiskControlOrderItemMgr) GetFromPlatformProductID(platformProductID string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`platform_product_id` = ?", platformProductID).Find(&results).Error

	return
}

// GetBatchFromPlatformProductID 批量查找 平台产品id
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromPlatformProductID(platformProductIDs []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`platform_product_id` IN (?)", platformProductIDs).Find(&results).Error

	return
}

// GetFromSpecs 通过specs获取内容 规格
func (obj *_LgRiskControlOrderItemMgr) GetFromSpecs(specs string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`specs` = ?", specs).Find(&results).Error

	return
}

// GetBatchFromSpecs 批量查找 规格
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromSpecs(specss []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`specs` IN (?)", specss).Find(&results).Error

	return
}

// GetFromPlaceOrigin 通过place_origin获取内容 产地
func (obj *_LgRiskControlOrderItemMgr) GetFromPlaceOrigin(placeOrigin string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`place_origin` = ?", placeOrigin).Find(&results).Error

	return
}

// GetBatchFromPlaceOrigin 批量查找 产地
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromPlaceOrigin(placeOrigins []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`place_origin` IN (?)", placeOrigins).Find(&results).Error

	return
}

// GetFromCurrency 通过currency获取内容 币种
func (obj *_LgRiskControlOrderItemMgr) GetFromCurrency(currency string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`currency` = ?", currency).Find(&results).Error

	return
}

// GetBatchFromCurrency 批量查找 币种
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromCurrency(currencys []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`currency` IN (?)", currencys).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 单位
func (obj *_LgRiskControlOrderItemMgr) GetFromCompany(company string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量查找 单位
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromCompany(companys []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`company` IN (?)", companys).Find(&results).Error

	return
}

// GetFromPicturesURL 通过pictures_url获取内容 图片链接
func (obj *_LgRiskControlOrderItemMgr) GetFromPicturesURL(picturesURL string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`pictures_url` = ?", picturesURL).Find(&results).Error

	return
}

// GetBatchFromPicturesURL 批量查找 图片链接
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromPicturesURL(picturesURLs []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`pictures_url` IN (?)", picturesURLs).Find(&results).Error

	return
}

// GetFromBrand 通过brand获取内容 品牌
func (obj *_LgRiskControlOrderItemMgr) GetFromBrand(brand string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`brand` = ?", brand).Find(&results).Error

	return
}

// GetBatchFromBrand 批量查找 品牌
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromBrand(brands []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`brand` IN (?)", brands).Find(&results).Error

	return
}

// GetFromHasBattery 通过has_battery获取内容 是否带电
func (obj *_LgRiskControlOrderItemMgr) GetFromHasBattery(hasBattery int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`has_battery` = ?", hasBattery).Find(&results).Error

	return
}

// GetBatchFromHasBattery 批量查找 是否带电
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromHasBattery(hasBatterys []int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`has_battery` IN (?)", hasBatterys).Find(&results).Error

	return
}

// GetFromLiquid 通过liquid获取内容 是否液体,1-液体，0-非液体
func (obj *_LgRiskControlOrderItemMgr) GetFromLiquid(liquid string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`liquid` = ?", liquid).Find(&results).Error

	return
}

// GetBatchFromLiquid 批量查找 是否液体,1-液体，0-非液体
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromLiquid(liquids []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`liquid` IN (?)", liquids).Find(&results).Error

	return
}

// GetFromPowder 通过powder获取内容 是否粉末,1-粉末，0-非粉末
func (obj *_LgRiskControlOrderItemMgr) GetFromPowder(powder string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`powder` = ?", powder).Find(&results).Error

	return
}

// GetBatchFromPowder 批量查找 是否粉末,1-粉末，0-非粉末
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromPowder(powders []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`powder` IN (?)", powders).Find(&results).Error

	return
}

// GetFromMagnetic 通过magnetic获取内容 是否带磁,1-带磁，0-不带磁
func (obj *_LgRiskControlOrderItemMgr) GetFromMagnetic(magnetic string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`magnetic` = ?", magnetic).Find(&results).Error

	return
}

// GetBatchFromMagnetic 批量查找 是否带磁,1-带磁，0-不带磁
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromMagnetic(magnetics []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`magnetic` IN (?)", magnetics).Find(&results).Error

	return
}

// GetFromDistributionInfo 通过distribution_info获取内容 配货信息
func (obj *_LgRiskControlOrderItemMgr) GetFromDistributionInfo(distributionInfo string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`distribution_info` = ?", distributionInfo).Find(&results).Error

	return
}

// GetBatchFromDistributionInfo 批量查找 配货信息
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromDistributionInfo(distributionInfos []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`distribution_info` IN (?)", distributionInfos).Find(&results).Error

	return
}

// GetFromBatteryType 通过battery_type获取内容 电池类型
func (obj *_LgRiskControlOrderItemMgr) GetFromBatteryType(batteryType string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`battery_type` = ?", batteryType).Find(&results).Error

	return
}

// GetBatchFromBatteryType 批量查找 电池类型
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromBatteryType(batteryTypes []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`battery_type` IN (?)", batteryTypes).Find(&results).Error

	return
}

// GetFromSalesPlatform 通过sales_platform获取内容 销售平台
func (obj *_LgRiskControlOrderItemMgr) GetFromSalesPlatform(salesPlatform string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`sales_platform` = ?", salesPlatform).Find(&results).Error

	return
}

// GetBatchFromSalesPlatform 批量查找 销售平台
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromSalesPlatform(salesPlatforms []string) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`sales_platform` IN (?)", salesPlatforms).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgRiskControlOrderItemMgr) GetFromCreateTime(createTime time.Time) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgRiskControlOrderItemMgr) GetFromCreateUser(createUser int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgRiskControlOrderItemMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgRiskControlOrderItemMgr) GetFromUpdateUser(updateUser int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgRiskControlOrderItemMgr) GetFromVersion(version int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromVersion(versions []int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgRiskControlOrderItemMgr) GetFromDeleted(deleted int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgRiskControlOrderItemMgr) GetBatchFromDeleted(deleteds []int) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgRiskControlOrderItemMgr) FetchByPrimaryKey(id int) (result LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByLgRcOrderItemLgRcOrderRcNumberFk  获取多个内容
func (obj *_LgRiskControlOrderItemMgr) FetchIndexByLgRcOrderItemLgRcOrderRcNumberFk(riskControlOrderID int64) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`risk_control_order_id` = ?", riskControlOrderID).Find(&results).Error

	return
}

// FetchIndexByLgRiskControlOrderItemCreateTimeIndex  获取多个内容
func (obj *_LgRiskControlOrderItemMgr) FetchIndexByLgRiskControlOrderItemCreateTimeIndex(createTime time.Time) (results []*LgRiskControlOrderItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrderItem{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
