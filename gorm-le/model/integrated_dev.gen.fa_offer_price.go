package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaOfferPriceMgr struct {
	*_BaseMgr
}

// FaOfferPriceMgr open func
func FaOfferPriceMgr(db *gorm.DB) *_FaOfferPriceMgr {
	if db == nil {
		panic(fmt.Errorf("FaOfferPriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaOfferPriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_offer_price"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaOfferPriceMgr) GetTableName() string {
	return "fa_offer_price"
}

// Reset 重置gorm会话
func (obj *_FaOfferPriceMgr) Reset() *_FaOfferPriceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaOfferPriceMgr) Get() (result FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaOfferPriceMgr) Gets() (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaOfferPriceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaOfferPriceMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_FaOfferPriceMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOfferType offer_type获取 报价类型 客户报价 渠道报价 成本价等
func (obj *_FaOfferPriceMgr) WithOfferType(offerType string) Option {
	return optionFunc(func(o *options) { o.query["offer_type"] = offerType })
}

// WithPriceType price_type获取 价格类型 公开报价等
func (obj *_FaOfferPriceMgr) WithPriceType(priceType string) Option {
	return optionFunc(func(o *options) { o.query["price_type"] = priceType })
}

// WithOfferAttr offer_attr获取 报价属性 1:应收  2:应付
func (obj *_FaOfferPriceMgr) WithOfferAttr(offerAttr int) Option {
	return optionFunc(func(o *options) { o.query["offer_attr"] = offerAttr })
}

// WithStatus status获取 报价表状态:启用/禁用
func (obj *_FaOfferPriceMgr) WithStatus(status []uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCustomerChannelID customer_channel_id获取 关联的客户渠道ID
func (obj *_FaOfferPriceMgr) WithCustomerChannelID(customerChannelID int64) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_FaOfferPriceMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_FaOfferPriceMgr) WithCustomerID(customerID int64) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_FaOfferPriceMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithProviderChannelCode provider_channel_code获取 服务商渠道代码
func (obj *_FaOfferPriceMgr) WithProviderChannelCode(providerChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_code"] = providerChannelCode })
}

// WithProviderChannelName provider_channel_name获取 服务商渠道名称
func (obj *_FaOfferPriceMgr) WithProviderChannelName(providerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_name"] = providerChannelName })
}

// WithCurrency currency获取 币种
func (obj *_FaOfferPriceMgr) WithCurrency(currency string) Option {
	return optionFunc(func(o *options) { o.query["currency"] = currency })
}

// WithExtraCurrency extra_currency获取 额外加价币种
func (obj *_FaOfferPriceMgr) WithExtraCurrency(extraCurrency string) Option {
	return optionFunc(func(o *options) { o.query["extra_currency"] = extraCurrency })
}

// WithPartitionID partition_id获取 关联的分区ID
func (obj *_FaOfferPriceMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithPartitionName partition_name获取 关联的分区名称
func (obj *_FaOfferPriceMgr) WithPartitionName(partitionName string) Option {
	return optionFunc(func(o *options) { o.query["partition_name"] = partitionName })
}

// WithHaveSpecialPrice have_special_price获取 是否有特殊价格
func (obj *_FaOfferPriceMgr) WithHaveSpecialPrice(haveSpecialPrice []uint8) Option {
	return optionFunc(func(o *options) { o.query["have_special_price"] = haveSpecialPrice })
}

// WithBeginTime begin_time获取 开始时间
func (obj *_FaOfferPriceMgr) WithBeginTime(beginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["begin_time"] = beginTime })
}

// WithEndTime end_time获取 结束时间
func (obj *_FaOfferPriceMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithRemark remark获取 备注
func (obj *_FaOfferPriceMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaOfferPriceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaOfferPriceMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaOfferPriceMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_FaOfferPriceMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_FaOfferPriceMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaOfferPriceMgr) GetByOption(opts ...Option) (result FaOfferPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaOfferPriceMgr) GetByOptions(opts ...Option) (results []*FaOfferPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaOfferPriceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaOfferPrice, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_FaOfferPriceMgr) GetFromID(id int) (result FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaOfferPriceMgr) GetBatchFromID(ids []int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_FaOfferPriceMgr) GetFromName(name string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_FaOfferPriceMgr) GetBatchFromName(names []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromOfferType 通过offer_type获取内容 报价类型 客户报价 渠道报价 成本价等
func (obj *_FaOfferPriceMgr) GetFromOfferType(offerType string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`offer_type` = ?", offerType).Find(&results).Error

	return
}

// GetBatchFromOfferType 批量查找 报价类型 客户报价 渠道报价 成本价等
func (obj *_FaOfferPriceMgr) GetBatchFromOfferType(offerTypes []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`offer_type` IN (?)", offerTypes).Find(&results).Error

	return
}

// GetFromPriceType 通过price_type获取内容 价格类型 公开报价等
func (obj *_FaOfferPriceMgr) GetFromPriceType(priceType string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`price_type` = ?", priceType).Find(&results).Error

	return
}

// GetBatchFromPriceType 批量查找 价格类型 公开报价等
func (obj *_FaOfferPriceMgr) GetBatchFromPriceType(priceTypes []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`price_type` IN (?)", priceTypes).Find(&results).Error

	return
}

// GetFromOfferAttr 通过offer_attr获取内容 报价属性 1:应收  2:应付
func (obj *_FaOfferPriceMgr) GetFromOfferAttr(offerAttr int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`offer_attr` = ?", offerAttr).Find(&results).Error

	return
}

// GetBatchFromOfferAttr 批量查找 报价属性 1:应收  2:应付
func (obj *_FaOfferPriceMgr) GetBatchFromOfferAttr(offerAttrs []int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`offer_attr` IN (?)", offerAttrs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 报价表状态:启用/禁用
func (obj *_FaOfferPriceMgr) GetFromStatus(status []uint8) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 报价表状态:启用/禁用
func (obj *_FaOfferPriceMgr) GetBatchFromStatus(statuss [][]uint8) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 关联的客户渠道ID
func (obj *_FaOfferPriceMgr) GetFromCustomerChannelID(customerChannelID int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 关联的客户渠道ID
func (obj *_FaOfferPriceMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_FaOfferPriceMgr) GetFromCustomerChannelName(customerChannelName string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_FaOfferPriceMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_FaOfferPriceMgr) GetFromCustomerID(customerID int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_FaOfferPriceMgr) GetBatchFromCustomerID(customerIDs []int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_FaOfferPriceMgr) GetFromCustomerAlias(customerAlias string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_FaOfferPriceMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromProviderChannelCode 通过provider_channel_code获取内容 服务商渠道代码
func (obj *_FaOfferPriceMgr) GetFromProviderChannelCode(providerChannelCode string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error

	return
}

// GetBatchFromProviderChannelCode 批量查找 服务商渠道代码
func (obj *_FaOfferPriceMgr) GetBatchFromProviderChannelCode(providerChannelCodes []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`provider_channel_code` IN (?)", providerChannelCodes).Find(&results).Error

	return
}

// GetFromProviderChannelName 通过provider_channel_name获取内容 服务商渠道名称
func (obj *_FaOfferPriceMgr) GetFromProviderChannelName(providerChannelName string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`provider_channel_name` = ?", providerChannelName).Find(&results).Error

	return
}

// GetBatchFromProviderChannelName 批量查找 服务商渠道名称
func (obj *_FaOfferPriceMgr) GetBatchFromProviderChannelName(providerChannelNames []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`provider_channel_name` IN (?)", providerChannelNames).Find(&results).Error

	return
}

// GetFromCurrency 通过currency获取内容 币种
func (obj *_FaOfferPriceMgr) GetFromCurrency(currency string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`currency` = ?", currency).Find(&results).Error

	return
}

// GetBatchFromCurrency 批量查找 币种
func (obj *_FaOfferPriceMgr) GetBatchFromCurrency(currencys []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`currency` IN (?)", currencys).Find(&results).Error

	return
}

// GetFromExtraCurrency 通过extra_currency获取内容 额外加价币种
func (obj *_FaOfferPriceMgr) GetFromExtraCurrency(extraCurrency string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`extra_currency` = ?", extraCurrency).Find(&results).Error

	return
}

// GetBatchFromExtraCurrency 批量查找 额外加价币种
func (obj *_FaOfferPriceMgr) GetBatchFromExtraCurrency(extraCurrencys []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`extra_currency` IN (?)", extraCurrencys).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容 关联的分区ID
func (obj *_FaOfferPriceMgr) GetFromPartitionID(partitionID int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找 关联的分区ID
func (obj *_FaOfferPriceMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromPartitionName 通过partition_name获取内容 关联的分区名称
func (obj *_FaOfferPriceMgr) GetFromPartitionName(partitionName string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`partition_name` = ?", partitionName).Find(&results).Error

	return
}

// GetBatchFromPartitionName 批量查找 关联的分区名称
func (obj *_FaOfferPriceMgr) GetBatchFromPartitionName(partitionNames []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`partition_name` IN (?)", partitionNames).Find(&results).Error

	return
}

// GetFromHaveSpecialPrice 通过have_special_price获取内容 是否有特殊价格
func (obj *_FaOfferPriceMgr) GetFromHaveSpecialPrice(haveSpecialPrice []uint8) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`have_special_price` = ?", haveSpecialPrice).Find(&results).Error

	return
}

// GetBatchFromHaveSpecialPrice 批量查找 是否有特殊价格
func (obj *_FaOfferPriceMgr) GetBatchFromHaveSpecialPrice(haveSpecialPrices [][]uint8) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`have_special_price` IN (?)", haveSpecialPrices).Find(&results).Error

	return
}

// GetFromBeginTime 通过begin_time获取内容 开始时间
func (obj *_FaOfferPriceMgr) GetFromBeginTime(beginTime time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`begin_time` = ?", beginTime).Find(&results).Error

	return
}

// GetBatchFromBeginTime 批量查找 开始时间
func (obj *_FaOfferPriceMgr) GetBatchFromBeginTime(beginTimes []time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`begin_time` IN (?)", beginTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 结束时间
func (obj *_FaOfferPriceMgr) GetFromEndTime(endTime time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 结束时间
func (obj *_FaOfferPriceMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaOfferPriceMgr) GetFromRemark(remark string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaOfferPriceMgr) GetBatchFromRemark(remarks []string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaOfferPriceMgr) GetFromCreateTime(createTime time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaOfferPriceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaOfferPriceMgr) GetFromCreateUser(createUser int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaOfferPriceMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaOfferPriceMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaOfferPriceMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_FaOfferPriceMgr) GetFromUpdateUser(updateUser int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_FaOfferPriceMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaOfferPriceMgr) GetFromVersion(version int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaOfferPriceMgr) GetBatchFromVersion(versions []int) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaOfferPriceMgr) FetchByPrimaryKey(id int) (result FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByFaOfferPriceNameIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPriceNameIndex(name string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// FetchIndexByFaOfferPriceOfferTypeIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPriceOfferTypeIndex(offerType string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`offer_type` = ?", offerType).Find(&results).Error

	return
}

// FetchIndexByFaOfferPricePriceTypeIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPricePriceTypeIndex(priceType string) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`price_type` = ?", priceType).Find(&results).Error

	return
}

// FetchIndexByFaOfferPriceStatusIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPriceStatusIndex(status []uint8) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// FetchIndexByFaOfferPriceCustomerChannelIDIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPriceCustomerChannelIDIndex(customerChannelID int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByFaOfferPriceCustomerIDIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPriceCustomerIDIndex(customerID int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexByFaOfferPricePartitionIDIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPricePartitionIDIndex(partitionID int64) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// FetchIndexByFaOfferPriceHaveSpecialPriceIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPriceHaveSpecialPriceIndex(haveSpecialPrice []uint8) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`have_special_price` = ?", haveSpecialPrice).Find(&results).Error

	return
}

// FetchIndexByFaOfferPriceBeginTimeIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPriceBeginTimeIndex(beginTime time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`begin_time` = ?", beginTime).Find(&results).Error

	return
}

// FetchIndexByFaOfferPriceEndTimeIndex  获取多个内容
func (obj *_FaOfferPriceMgr) FetchIndexByFaOfferPriceEndTimeIndex(endTime time.Time) (results []*FaOfferPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPrice{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}
