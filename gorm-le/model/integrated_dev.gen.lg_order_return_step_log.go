package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderReturnStepLogMgr struct {
	*_BaseMgr
}

// LgOrderReturnStepLogMgr open func
func LgOrderReturnStepLogMgr(db *gorm.DB) *_LgOrderReturnStepLogMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderReturnStepLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderReturnStepLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_return_step_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderReturnStepLogMgr) GetTableName() string {
	return "lg_order_return_step_log"
}

// Reset 重置gorm会话
func (obj *_LgOrderReturnStepLogMgr) Reset() *_LgOrderReturnStepLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderReturnStepLogMgr) Get() (result LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderReturnStepLogMgr) Gets() (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderReturnStepLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgOrderReturnStepLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_LgOrderReturnStepLogMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_LgOrderReturnStepLogMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_LgOrderReturnStepLogMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_LgOrderReturnStepLogMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithCustomerID customer_id获取 归属客户ID
func (obj *_LgOrderReturnStepLogMgr) WithCustomerID(customerID int64) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_LgOrderReturnStepLogMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgOrderReturnStepLogMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 渠道名称
func (obj *_LgOrderReturnStepLogMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithReceiveCountry receive_country获取 收件人国家
func (obj *_LgOrderReturnStepLogMgr) WithReceiveCountry(receiveCountry string) Option {
	return optionFunc(func(o *options) { o.query["receive_country"] = receiveCountry })
}

// WithReturnStepLocation return_step_location获取 退件地点
func (obj *_LgOrderReturnStepLogMgr) WithReturnStepLocation(returnStepLocation string) Option {
	return optionFunc(func(o *options) { o.query["return_step_location"] = returnStepLocation })
}

// WithReturnStepReason return_step_reason获取 退件原因
func (obj *_LgOrderReturnStepLogMgr) WithReturnStepReason(returnStepReason string) Option {
	return optionFunc(func(o *options) { o.query["return_step_reason"] = returnStepReason })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderReturnStepLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderReturnStepLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateUserName create_user_name获取 创建用户名称
func (obj *_LgOrderReturnStepLogMgr) WithCreateUserName(createUserName string) Option {
	return optionFunc(func(o *options) { o.query["create_user_name"] = createUserName })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderReturnStepLogMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_LgOrderReturnStepLogMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateUserName update_user_name获取 更新用户名称
func (obj *_LgOrderReturnStepLogMgr) WithUpdateUserName(updateUserName string) Option {
	return optionFunc(func(o *options) { o.query["update_user_name"] = updateUserName })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderReturnStepLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderReturnStepLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderReturnStepLogMgr) GetByOption(opts ...Option) (result LgOrderReturnStepLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderReturnStepLogMgr) GetByOptions(opts ...Option) (results []*LgOrderReturnStepLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderReturnStepLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderReturnStepLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where(options.query)
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
func (obj *_LgOrderReturnStepLogMgr) GetFromID(id int64) (result LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromID(ids []int64) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_LgOrderReturnStepLogMgr) GetFromOrderNumber(orderNumber string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_LgOrderReturnStepLogMgr) GetFromReferenceNumber(referenceNumber string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_LgOrderReturnStepLogMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_LgOrderReturnStepLogMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 归属客户ID
func (obj *_LgOrderReturnStepLogMgr) GetFromCustomerID(customerID int64) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 归属客户ID
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromCustomerID(customerIDs []int64) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_LgOrderReturnStepLogMgr) GetFromCustomerAlias(customerAlias string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgOrderReturnStepLogMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 渠道名称
func (obj *_LgOrderReturnStepLogMgr) GetFromCustomerChannelName(customerChannelName string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 渠道名称
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromReceiveCountry 通过receive_country获取内容 收件人国家
func (obj *_LgOrderReturnStepLogMgr) GetFromReceiveCountry(receiveCountry string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// GetBatchFromReceiveCountry 批量查找 收件人国家
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromReceiveCountry(receiveCountrys []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`receive_country` IN (?)", receiveCountrys).Find(&results).Error

	return
}

// GetFromReturnStepLocation 通过return_step_location获取内容 退件地点
func (obj *_LgOrderReturnStepLogMgr) GetFromReturnStepLocation(returnStepLocation string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`return_step_location` = ?", returnStepLocation).Find(&results).Error

	return
}

// GetBatchFromReturnStepLocation 批量查找 退件地点
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromReturnStepLocation(returnStepLocations []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`return_step_location` IN (?)", returnStepLocations).Find(&results).Error

	return
}

// GetFromReturnStepReason 通过return_step_reason获取内容 退件原因
func (obj *_LgOrderReturnStepLogMgr) GetFromReturnStepReason(returnStepReason string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`return_step_reason` = ?", returnStepReason).Find(&results).Error

	return
}

// GetBatchFromReturnStepReason 批量查找 退件原因
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromReturnStepReason(returnStepReasons []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`return_step_reason` IN (?)", returnStepReasons).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderReturnStepLogMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderReturnStepLogMgr) GetFromCreateUser(createUser int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateUserName 通过create_user_name获取内容 创建用户名称
func (obj *_LgOrderReturnStepLogMgr) GetFromCreateUserName(createUserName string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`create_user_name` = ?", createUserName).Find(&results).Error

	return
}

// GetBatchFromCreateUserName 批量查找 创建用户名称
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromCreateUserName(createUserNames []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`create_user_name` IN (?)", createUserNames).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderReturnStepLogMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_LgOrderReturnStepLogMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateUserName 通过update_user_name获取内容 更新用户名称
func (obj *_LgOrderReturnStepLogMgr) GetFromUpdateUserName(updateUserName string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`update_user_name` = ?", updateUserName).Find(&results).Error

	return
}

// GetBatchFromUpdateUserName 批量查找 更新用户名称
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromUpdateUserName(updateUserNames []string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`update_user_name` IN (?)", updateUserNames).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderReturnStepLogMgr) GetFromVersion(version int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromVersion(versions []int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderReturnStepLogMgr) GetFromDeleted(deleted int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderReturnStepLogMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderReturnStepLogMgr) FetchByPrimaryKey(id int64) (result LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByLgOrderReturnStepLogNumberCombinationIndex primary or index 获取唯一内容
func (obj *_LgOrderReturnStepLogMgr) FetchUniqueIndexByLgOrderReturnStepLogNumberCombinationIndex(orderNumber string, referenceNumber string, logisticsNumber string, logisticsNumberFinal string) (result LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`order_number` = ? AND `reference_number` = ? AND `logistics_number` = ? AND `logistics_number_final` = ?", orderNumber, referenceNumber, logisticsNumber, logisticsNumberFinal).First(&result).Error

	return
}

// FetchIndexByLgOrderReturnStepLogOrderNumberIndex  获取多个内容
func (obj *_LgOrderReturnStepLogMgr) FetchIndexByLgOrderReturnStepLogOrderNumberIndex(orderNumber string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderReturnStepLogReferenceNumberIndex  获取多个内容
func (obj *_LgOrderReturnStepLogMgr) FetchIndexByLgOrderReturnStepLogReferenceNumberIndex(referenceNumber string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderReturnStepLogLogisticsNumberIndex  获取多个内容
func (obj *_LgOrderReturnStepLogMgr) FetchIndexByLgOrderReturnStepLogLogisticsNumberIndex(logisticsNumber string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderReturnStepLogLogisticsNumberFinalIndex  获取多个内容
func (obj *_LgOrderReturnStepLogMgr) FetchIndexByLgOrderReturnStepLogLogisticsNumberFinalIndex(logisticsNumberFinal string) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexByLgOrderReturnStepLogCreateTimeIndex  获取多个内容
func (obj *_LgOrderReturnStepLogMgr) FetchIndexByLgOrderReturnStepLogCreateTimeIndex(createTime time.Time) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderReturnStepLogUpdateTimeIndex  获取多个内容
func (obj *_LgOrderReturnStepLogMgr) FetchIndexByLgOrderReturnStepLogUpdateTimeIndex(updateTime time.Time) (results []*LgOrderReturnStepLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderReturnStepLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}
