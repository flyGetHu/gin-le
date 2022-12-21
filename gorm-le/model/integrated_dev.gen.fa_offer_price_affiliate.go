package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaOfferPriceAffiliateMgr struct {
	*_BaseMgr
}

// FaOfferPriceAffiliateMgr open func
func FaOfferPriceAffiliateMgr(db *gorm.DB) *_FaOfferPriceAffiliateMgr {
	if db == nil {
		panic(fmt.Errorf("FaOfferPriceAffiliateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaOfferPriceAffiliateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_offer_price_affiliate"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaOfferPriceAffiliateMgr) GetTableName() string {
	return "fa_offer_price_affiliate"
}

// Reset 重置gorm会话
func (obj *_FaOfferPriceAffiliateMgr) Reset() *_FaOfferPriceAffiliateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaOfferPriceAffiliateMgr) Get() (result FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaOfferPriceAffiliateMgr) Gets() (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaOfferPriceAffiliateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FaOfferPriceAffiliateMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOfferPriceID offer_price_id获取 关联主表ID
func (obj *_FaOfferPriceAffiliateMgr) WithOfferPriceID(offerPriceID int64) Option {
	return optionFunc(func(o *options) { o.query["offer_price_id"] = offerPriceID })
}

// WithPackageType package_type获取 包裹类型
func (obj *_FaOfferPriceAffiliateMgr) WithPackageType(packageType string) Option {
	return optionFunc(func(o *options) { o.query["package_type"] = packageType })
}

// WithPartitionIDentifier partition_identifier获取 分区标识
func (obj *_FaOfferPriceAffiliateMgr) WithPartitionIDentifier(partitionIDentifier int) Option {
	return optionFunc(func(o *options) { o.query["partition_identifier"] = partitionIDentifier })
}

// WithBeginWeight begin_weight获取 重量开始不包含
func (obj *_FaOfferPriceAffiliateMgr) WithBeginWeight(beginWeight float64) Option {
	return optionFunc(func(o *options) { o.query["begin_weight"] = beginWeight })
}

// WithEndWeight end_weight获取 重量结束:包含
func (obj *_FaOfferPriceAffiliateMgr) WithEndWeight(endWeight float64) Option {
	return optionFunc(func(o *options) { o.query["end_weight"] = endWeight })
}

// WithPriceType price_type获取 价格类型:
// 1.速查：只要重量在某1个重量段范围内。价格固定不变。比如“首重”就是其中一种例子。
// 2.总乘：重量在某1个重量段范围，没有起重，直接把公斤数字分割成多少段。这个段数与单价相乘得出的结果。
// 比如，快递行业常见的“大货价格算法”。按照1公斤为1段。
// 举例：20公斤到50公斤之间，按照1公斤10元来计算。这里的就是1公斤为1段。
// 20公斤的价格就是 20x10=200元
// 21公斤的价格就是 21x10=210元
// 20.5公斤的其中0.5也要单独分割为1段。不管是20.1公斤，还是20.9公斤，都被进位视为21公斤。价格就是21x10=210元。只要重量在就是20公斤到50公斤之间这个算法。比如，现在FEDEX、DHL等国际大件都是这个算法。
// 3.增量：以某个公斤数为起重，每增加多少公斤，对应的增加运费。比如“续重”其中1个例子。
// 比如国际EMS和国内EMS有这个算法。
// 4.直乘：没有首重。在某个公斤段范围内。把公斤数字和价格单价，直接进行数字相乘得出价格，就是直乘算法。
// 也没有小数点进位。算出价格是多少就是多少，小数点也算。比如现在的航空小包就是这个算法。
func (obj *_FaOfferPriceAffiliateMgr) WithPriceType(priceType int) Option {
	return optionFunc(func(o *options) { o.query["price_type"] = priceType })
}

// WithWeightUnit weight_unit获取 重量单位
func (obj *_FaOfferPriceAffiliateMgr) WithWeightUnit(weightUnit float64) Option {
	return optionFunc(func(o *options) { o.query["weight_unit"] = weightUnit })
}

// WithUnitPrice unit_price获取 单价
func (obj *_FaOfferPriceAffiliateMgr) WithUnitPrice(unitPrice float64) Option {
	return optionFunc(func(o *options) { o.query["unit_price"] = unitPrice })
}

// WithExtraPrice extra_price获取 额外加价
func (obj *_FaOfferPriceAffiliateMgr) WithExtraPrice(extraPrice float64) Option {
	return optionFunc(func(o *options) { o.query["extra_price"] = extraPrice })
}

// WithSumPrice sum_price获取 总和价格
func (obj *_FaOfferPriceAffiliateMgr) WithSumPrice(sumPrice float64) Option {
	return optionFunc(func(o *options) { o.query["sum_price"] = sumPrice })
}

// WithFirstWeight first_weight获取 首重重量
func (obj *_FaOfferPriceAffiliateMgr) WithFirstWeight(firstWeight float64) Option {
	return optionFunc(func(o *options) { o.query["first_weight"] = firstWeight })
}

// WithFirstWeightPrice first_weight_price获取 首重价格
func (obj *_FaOfferPriceAffiliateMgr) WithFirstWeightPrice(firstWeightPrice float64) Option {
	return optionFunc(func(o *options) { o.query["first_weight_price"] = firstWeightPrice })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaOfferPriceAffiliateMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaOfferPriceAffiliateMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaOfferPriceAffiliateMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_FaOfferPriceAffiliateMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_FaOfferPriceAffiliateMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaOfferPriceAffiliateMgr) GetByOption(opts ...Option) (result FaOfferPriceAffiliate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaOfferPriceAffiliateMgr) GetByOptions(opts ...Option) (results []*FaOfferPriceAffiliate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaOfferPriceAffiliateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaOfferPriceAffiliate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_FaOfferPriceAffiliateMgr) GetFromID(id int64) (result FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromID(ids []int64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOfferPriceID 通过offer_price_id获取内容 关联主表ID
func (obj *_FaOfferPriceAffiliateMgr) GetFromOfferPriceID(offerPriceID int64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`offer_price_id` = ?", offerPriceID).Find(&results).Error

	return
}

// GetBatchFromOfferPriceID 批量查找 关联主表ID
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromOfferPriceID(offerPriceIDs []int64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`offer_price_id` IN (?)", offerPriceIDs).Find(&results).Error

	return
}

// GetFromPackageType 通过package_type获取内容 包裹类型
func (obj *_FaOfferPriceAffiliateMgr) GetFromPackageType(packageType string) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`package_type` = ?", packageType).Find(&results).Error

	return
}

// GetBatchFromPackageType 批量查找 包裹类型
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromPackageType(packageTypes []string) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`package_type` IN (?)", packageTypes).Find(&results).Error

	return
}

// GetFromPartitionIDentifier 通过partition_identifier获取内容 分区标识
func (obj *_FaOfferPriceAffiliateMgr) GetFromPartitionIDentifier(partitionIDentifier int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`partition_identifier` = ?", partitionIDentifier).Find(&results).Error

	return
}

// GetBatchFromPartitionIDentifier 批量查找 分区标识
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromPartitionIDentifier(partitionIDentifiers []int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`partition_identifier` IN (?)", partitionIDentifiers).Find(&results).Error

	return
}

// GetFromBeginWeight 通过begin_weight获取内容 重量开始不包含
func (obj *_FaOfferPriceAffiliateMgr) GetFromBeginWeight(beginWeight float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`begin_weight` = ?", beginWeight).Find(&results).Error

	return
}

// GetBatchFromBeginWeight 批量查找 重量开始不包含
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromBeginWeight(beginWeights []float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`begin_weight` IN (?)", beginWeights).Find(&results).Error

	return
}

// GetFromEndWeight 通过end_weight获取内容 重量结束:包含
func (obj *_FaOfferPriceAffiliateMgr) GetFromEndWeight(endWeight float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`end_weight` = ?", endWeight).Find(&results).Error

	return
}

// GetBatchFromEndWeight 批量查找 重量结束:包含
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromEndWeight(endWeights []float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`end_weight` IN (?)", endWeights).Find(&results).Error

	return
}

// GetFromPriceType 通过price_type获取内容 价格类型:
// 1.速查：只要重量在某1个重量段范围内。价格固定不变。比如“首重”就是其中一种例子。
// 2.总乘：重量在某1个重量段范围，没有起重，直接把公斤数字分割成多少段。这个段数与单价相乘得出的结果。
// 比如，快递行业常见的“大货价格算法”。按照1公斤为1段。
// 举例：20公斤到50公斤之间，按照1公斤10元来计算。这里的就是1公斤为1段。
// 20公斤的价格就是 20x10=200元
// 21公斤的价格就是 21x10=210元
// 20.5公斤的其中0.5也要单独分割为1段。不管是20.1公斤，还是20.9公斤，都被进位视为21公斤。价格就是21x10=210元。只要重量在就是20公斤到50公斤之间这个算法。比如，现在FEDEX、DHL等国际大件都是这个算法。
// 3.增量：以某个公斤数为起重，每增加多少公斤，对应的增加运费。比如“续重”其中1个例子。
// 比如国际EMS和国内EMS有这个算法。
// 4.直乘：没有首重。在某个公斤段范围内。把公斤数字和价格单价，直接进行数字相乘得出价格，就是直乘算法。
// 也没有小数点进位。算出价格是多少就是多少，小数点也算。比如现在的航空小包就是这个算法。
func (obj *_FaOfferPriceAffiliateMgr) GetFromPriceType(priceType int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`price_type` = ?", priceType).Find(&results).Error

	return
}

// GetBatchFromPriceType 批量查找 价格类型:
// 1.速查：只要重量在某1个重量段范围内。价格固定不变。比如“首重”就是其中一种例子。
// 2.总乘：重量在某1个重量段范围，没有起重，直接把公斤数字分割成多少段。这个段数与单价相乘得出的结果。
// 比如，快递行业常见的“大货价格算法”。按照1公斤为1段。
// 举例：20公斤到50公斤之间，按照1公斤10元来计算。这里的就是1公斤为1段。
// 20公斤的价格就是 20x10=200元
// 21公斤的价格就是 21x10=210元
// 20.5公斤的其中0.5也要单独分割为1段。不管是20.1公斤，还是20.9公斤，都被进位视为21公斤。价格就是21x10=210元。只要重量在就是20公斤到50公斤之间这个算法。比如，现在FEDEX、DHL等国际大件都是这个算法。
// 3.增量：以某个公斤数为起重，每增加多少公斤，对应的增加运费。比如“续重”其中1个例子。
// 比如国际EMS和国内EMS有这个算法。
// 4.直乘：没有首重。在某个公斤段范围内。把公斤数字和价格单价，直接进行数字相乘得出价格，就是直乘算法。
// 也没有小数点进位。算出价格是多少就是多少，小数点也算。比如现在的航空小包就是这个算法。
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromPriceType(priceTypes []int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`price_type` IN (?)", priceTypes).Find(&results).Error

	return
}

// GetFromWeightUnit 通过weight_unit获取内容 重量单位
func (obj *_FaOfferPriceAffiliateMgr) GetFromWeightUnit(weightUnit float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`weight_unit` = ?", weightUnit).Find(&results).Error

	return
}

// GetBatchFromWeightUnit 批量查找 重量单位
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromWeightUnit(weightUnits []float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`weight_unit` IN (?)", weightUnits).Find(&results).Error

	return
}

// GetFromUnitPrice 通过unit_price获取内容 单价
func (obj *_FaOfferPriceAffiliateMgr) GetFromUnitPrice(unitPrice float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`unit_price` = ?", unitPrice).Find(&results).Error

	return
}

// GetBatchFromUnitPrice 批量查找 单价
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromUnitPrice(unitPrices []float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`unit_price` IN (?)", unitPrices).Find(&results).Error

	return
}

// GetFromExtraPrice 通过extra_price获取内容 额外加价
func (obj *_FaOfferPriceAffiliateMgr) GetFromExtraPrice(extraPrice float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`extra_price` = ?", extraPrice).Find(&results).Error

	return
}

// GetBatchFromExtraPrice 批量查找 额外加价
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromExtraPrice(extraPrices []float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`extra_price` IN (?)", extraPrices).Find(&results).Error

	return
}

// GetFromSumPrice 通过sum_price获取内容 总和价格
func (obj *_FaOfferPriceAffiliateMgr) GetFromSumPrice(sumPrice float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`sum_price` = ?", sumPrice).Find(&results).Error

	return
}

// GetBatchFromSumPrice 批量查找 总和价格
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromSumPrice(sumPrices []float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`sum_price` IN (?)", sumPrices).Find(&results).Error

	return
}

// GetFromFirstWeight 通过first_weight获取内容 首重重量
func (obj *_FaOfferPriceAffiliateMgr) GetFromFirstWeight(firstWeight float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`first_weight` = ?", firstWeight).Find(&results).Error

	return
}

// GetBatchFromFirstWeight 批量查找 首重重量
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromFirstWeight(firstWeights []float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`first_weight` IN (?)", firstWeights).Find(&results).Error

	return
}

// GetFromFirstWeightPrice 通过first_weight_price获取内容 首重价格
func (obj *_FaOfferPriceAffiliateMgr) GetFromFirstWeightPrice(firstWeightPrice float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`first_weight_price` = ?", firstWeightPrice).Find(&results).Error

	return
}

// GetBatchFromFirstWeightPrice 批量查找 首重价格
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromFirstWeightPrice(firstWeightPrices []float64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`first_weight_price` IN (?)", firstWeightPrices).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaOfferPriceAffiliateMgr) GetFromCreateTime(createTime time.Time) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaOfferPriceAffiliateMgr) GetFromCreateUser(createUser int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaOfferPriceAffiliateMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_FaOfferPriceAffiliateMgr) GetFromUpdateUser(updateUser int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaOfferPriceAffiliateMgr) GetFromVersion(version int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaOfferPriceAffiliateMgr) GetBatchFromVersion(versions []int) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaOfferPriceAffiliateMgr) FetchByPrimaryKey(id int64) (result FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByFaOfferPriceAffiliateOfferPriceIDIndex  获取多个内容
func (obj *_FaOfferPriceAffiliateMgr) FetchIndexByFaOfferPriceAffiliateOfferPriceIDIndex(offerPriceID int64) (results []*FaOfferPriceAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceAffiliate{}).Where("`offer_price_id` = ?", offerPriceID).Find(&results).Error

	return
}
