package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CCustomerChannelConfigMgr struct {
	*_BaseMgr
}

// CCustomerChannelConfigMgr open func
func CCustomerChannelConfigMgr(db *gorm.DB) *_CCustomerChannelConfigMgr {
	if db == nil {
		panic(fmt.Errorf("CCustomerChannelConfigMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CCustomerChannelConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_customer_channel_config"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CCustomerChannelConfigMgr) GetTableName() string {
	return "c_customer_channel_config"
}

// Reset 重置gorm会话
func (obj *_CCustomerChannelConfigMgr) Reset() *_CCustomerChannelConfigMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CCustomerChannelConfigMgr) Get() (result CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CCustomerChannelConfigMgr) Gets() (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CCustomerChannelConfigMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CCustomerChannelConfigMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerID customer_id获取 客户id
func (obj *_CCustomerChannelConfigMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道id
func (obj *_CCustomerChannelConfigMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithOrderCountMax order_count_max获取 下单量/天
func (obj *_CCustomerChannelConfigMgr) WithOrderCountMax(orderCountMax int) Option {
	return optionFunc(func(o *options) { o.query["order_count_max"] = orderCountMax })
}

// WithAmountMax amount_max获取 限额/元
func (obj *_CCustomerChannelConfigMgr) WithAmountMax(amountMax float64) Option {
	return optionFunc(func(o *options) { o.query["amount_max"] = amountMax })
}

// WithLabelCode label_code获取 客户配置此渠道使用的面单模板code
func (obj *_CCustomerChannelConfigMgr) WithLabelCode(labelCode string) Option {
	return optionFunc(func(o *options) { o.query["label_code"] = labelCode })
}

// WithSenderAddressBookID sender_address_book_id获取 发件人地址簿id
func (obj *_CCustomerChannelConfigMgr) WithSenderAddressBookID(senderAddressBookID int) Option {
	return optionFunc(func(o *options) { o.query["sender_address_book_id"] = senderAddressBookID })
}

// WithWeightingWeightDiff weighting_weight_diff获取 称重重量差kg
func (obj *_CCustomerChannelConfigMgr) WithWeightingWeightDiff(weightingWeightDiff float64) Option {
	return optionFunc(func(o *options) { o.query["weighting_weight_diff"] = weightingWeightDiff })
}

// WithWeightingWeightNegativeDiff weighting_weight_negative_diff获取 称重重量差kg(负)
func (obj *_CCustomerChannelConfigMgr) WithWeightingWeightNegativeDiff(weightingWeightNegativeDiff float64) Option {
	return optionFunc(func(o *options) { o.query["weighting_weight_negative_diff"] = weightingWeightNegativeDiff })
}

// WithProportion proportion获取 抛比
func (obj *_CCustomerChannelConfigMgr) WithProportion(proportion int) Option {
	return optionFunc(func(o *options) { o.query["proportion"] = proportion })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CCustomerChannelConfigMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_CCustomerChannelConfigMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CCustomerChannelConfigMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_CCustomerChannelConfigMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_CCustomerChannelConfigMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_CCustomerChannelConfigMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithSalesLink sales_link获取 销售链接
func (obj *_CCustomerChannelConfigMgr) WithSalesLink(salesLink string) Option {
	return optionFunc(func(o *options) { o.query["sales_link"] = salesLink })
}

// GetByOption 功能选项模式获取
func (obj *_CCustomerChannelConfigMgr) GetByOption(opts ...Option) (result CCustomerChannelConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CCustomerChannelConfigMgr) GetByOptions(opts ...Option) (results []*CCustomerChannelConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CCustomerChannelConfigMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CCustomerChannelConfig, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where(options.query)
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
func (obj *_CCustomerChannelConfigMgr) GetFromID(id int) (result CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_CCustomerChannelConfigMgr) GetBatchFromID(ids []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_CCustomerChannelConfigMgr) GetFromCustomerID(customerID int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_CCustomerChannelConfigMgr) GetBatchFromCustomerID(customerIDs []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道id
func (obj *_CCustomerChannelConfigMgr) GetFromCustomerChannelID(customerChannelID int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道id
func (obj *_CCustomerChannelConfigMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromOrderCountMax 通过order_count_max获取内容 下单量/天
func (obj *_CCustomerChannelConfigMgr) GetFromOrderCountMax(orderCountMax int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`order_count_max` = ?", orderCountMax).Find(&results).Error

	return
}

// GetBatchFromOrderCountMax 批量查找 下单量/天
func (obj *_CCustomerChannelConfigMgr) GetBatchFromOrderCountMax(orderCountMaxs []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`order_count_max` IN (?)", orderCountMaxs).Find(&results).Error

	return
}

// GetFromAmountMax 通过amount_max获取内容 限额/元
func (obj *_CCustomerChannelConfigMgr) GetFromAmountMax(amountMax float64) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`amount_max` = ?", amountMax).Find(&results).Error

	return
}

// GetBatchFromAmountMax 批量查找 限额/元
func (obj *_CCustomerChannelConfigMgr) GetBatchFromAmountMax(amountMaxs []float64) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`amount_max` IN (?)", amountMaxs).Find(&results).Error

	return
}

// GetFromLabelCode 通过label_code获取内容 客户配置此渠道使用的面单模板code
func (obj *_CCustomerChannelConfigMgr) GetFromLabelCode(labelCode string) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`label_code` = ?", labelCode).Find(&results).Error

	return
}

// GetBatchFromLabelCode 批量查找 客户配置此渠道使用的面单模板code
func (obj *_CCustomerChannelConfigMgr) GetBatchFromLabelCode(labelCodes []string) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`label_code` IN (?)", labelCodes).Find(&results).Error

	return
}

// GetFromSenderAddressBookID 通过sender_address_book_id获取内容 发件人地址簿id
func (obj *_CCustomerChannelConfigMgr) GetFromSenderAddressBookID(senderAddressBookID int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`sender_address_book_id` = ?", senderAddressBookID).Find(&results).Error

	return
}

// GetBatchFromSenderAddressBookID 批量查找 发件人地址簿id
func (obj *_CCustomerChannelConfigMgr) GetBatchFromSenderAddressBookID(senderAddressBookIDs []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`sender_address_book_id` IN (?)", senderAddressBookIDs).Find(&results).Error

	return
}

// GetFromWeightingWeightDiff 通过weighting_weight_diff获取内容 称重重量差kg
func (obj *_CCustomerChannelConfigMgr) GetFromWeightingWeightDiff(weightingWeightDiff float64) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`weighting_weight_diff` = ?", weightingWeightDiff).Find(&results).Error

	return
}

// GetBatchFromWeightingWeightDiff 批量查找 称重重量差kg
func (obj *_CCustomerChannelConfigMgr) GetBatchFromWeightingWeightDiff(weightingWeightDiffs []float64) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`weighting_weight_diff` IN (?)", weightingWeightDiffs).Find(&results).Error

	return
}

// GetFromWeightingWeightNegativeDiff 通过weighting_weight_negative_diff获取内容 称重重量差kg(负)
func (obj *_CCustomerChannelConfigMgr) GetFromWeightingWeightNegativeDiff(weightingWeightNegativeDiff float64) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`weighting_weight_negative_diff` = ?", weightingWeightNegativeDiff).Find(&results).Error

	return
}

// GetBatchFromWeightingWeightNegativeDiff 批量查找 称重重量差kg(负)
func (obj *_CCustomerChannelConfigMgr) GetBatchFromWeightingWeightNegativeDiff(weightingWeightNegativeDiffs []float64) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`weighting_weight_negative_diff` IN (?)", weightingWeightNegativeDiffs).Find(&results).Error

	return
}

// GetFromProportion 通过proportion获取内容 抛比
func (obj *_CCustomerChannelConfigMgr) GetFromProportion(proportion int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`proportion` = ?", proportion).Find(&results).Error

	return
}

// GetBatchFromProportion 批量查找 抛比
func (obj *_CCustomerChannelConfigMgr) GetBatchFromProportion(proportions []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`proportion` IN (?)", proportions).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CCustomerChannelConfigMgr) GetFromCreateTime(createTime time.Time) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CCustomerChannelConfigMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_CCustomerChannelConfigMgr) GetFromCreateUser(createUser int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_CCustomerChannelConfigMgr) GetBatchFromCreateUser(createUsers []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CCustomerChannelConfigMgr) GetFromUpdateTime(updateTime time.Time) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_CCustomerChannelConfigMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_CCustomerChannelConfigMgr) GetFromUpdateUser(updateUser int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_CCustomerChannelConfigMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_CCustomerChannelConfigMgr) GetFromVersion(version int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_CCustomerChannelConfigMgr) GetBatchFromVersion(versions []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_CCustomerChannelConfigMgr) GetFromDeleted(deleted int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_CCustomerChannelConfigMgr) GetBatchFromDeleted(deleteds []int) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromSalesLink 通过sales_link获取内容 销售链接
func (obj *_CCustomerChannelConfigMgr) GetFromSalesLink(salesLink string) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`sales_link` = ?", salesLink).Find(&results).Error

	return
}

// GetBatchFromSalesLink 批量查找 销售链接
func (obj *_CCustomerChannelConfigMgr) GetBatchFromSalesLink(salesLinks []string) (results []*CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`sales_link` IN (?)", salesLinks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CCustomerChannelConfigMgr) FetchByPrimaryKey(id int) (result CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByCustomerIDCustomerChannelIDIndex primary or index 获取唯一内容
func (obj *_CCustomerChannelConfigMgr) FetchUniqueIndexByCustomerIDCustomerChannelIDIndex(customerID int, customerChannelID int) (result CCustomerChannelConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerChannelConfig{}).Where("`customer_id` = ? AND `customer_channel_id` = ?", customerID, customerChannelID).First(&result).Error

	return
}
