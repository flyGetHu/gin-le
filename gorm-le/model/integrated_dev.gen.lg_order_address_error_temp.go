package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderAddressErrorTempMgr struct {
	*_BaseMgr
}

// LgOrderAddressErrorTempMgr open func
func LgOrderAddressErrorTempMgr(db *gorm.DB) *_LgOrderAddressErrorTempMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderAddressErrorTempMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderAddressErrorTempMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_address_error_temp"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderAddressErrorTempMgr) GetTableName() string {
	return "lg_order_address_error_temp"
}

// Reset 重置gorm会话
func (obj *_LgOrderAddressErrorTempMgr) Reset() *_LgOrderAddressErrorTempMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderAddressErrorTempMgr) Get() (result LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderAddressErrorTempMgr) Gets() (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderAddressErrorTempMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgOrderAddressErrorTempMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_LgOrderAddressErrorTempMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReceiveCity receive_city获取 收件人城市
func (obj *_LgOrderAddressErrorTempMgr) WithReceiveCity(receiveCity string) Option {
	return optionFunc(func(o *options) { o.query["receive_city"] = receiveCity })
}

// WithReceiveProvince receive_province获取 收件人省份
func (obj *_LgOrderAddressErrorTempMgr) WithReceiveProvince(receiveProvince string) Option {
	return optionFunc(func(o *options) { o.query["receive_province"] = receiveProvince })
}

// WithReceiveCountry receive_country获取 收件人国家
func (obj *_LgOrderAddressErrorTempMgr) WithReceiveCountry(receiveCountry string) Option {
	return optionFunc(func(o *options) { o.query["receive_country"] = receiveCountry })
}

// WithReceiveZipCode receive_zip_code获取 收件人邮编
func (obj *_LgOrderAddressErrorTempMgr) WithReceiveZipCode(receiveZipCode string) Option {
	return optionFunc(func(o *options) { o.query["receive_zip_code"] = receiveZipCode })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgOrderAddressErrorTempMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_LgOrderAddressErrorTempMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithCustomerID customer_id获取 归属客户ID
func (obj *_LgOrderAddressErrorTempMgr) WithCustomerID(customerID int64) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_LgOrderAddressErrorTempMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithErrorInfo error_info获取 错误信息
func (obj *_LgOrderAddressErrorTempMgr) WithErrorInfo(errorInfo string) Option {
	return optionFunc(func(o *options) { o.query["error_info"] = errorInfo })
}

// WithStatus status获取 状态 1 已处理 0 未处理
func (obj *_LgOrderAddressErrorTempMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderAddressErrorTempMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderAddressErrorTempMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_LgOrderAddressErrorTempMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderAddressErrorTempMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderAddressErrorTempMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderAddressErrorTempMgr) GetByOption(opts ...Option) (result LgOrderAddressErrorTemp, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderAddressErrorTempMgr) GetByOptions(opts ...Option) (results []*LgOrderAddressErrorTemp, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderAddressErrorTempMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderAddressErrorTemp, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where(options.query)
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
func (obj *_LgOrderAddressErrorTempMgr) GetFromID(id int) (result LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromID(ids []int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_LgOrderAddressErrorTempMgr) GetFromOrderNumber(orderNumber string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReceiveCity 通过receive_city获取内容 收件人城市
func (obj *_LgOrderAddressErrorTempMgr) GetFromReceiveCity(receiveCity string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_city` = ?", receiveCity).Find(&results).Error

	return
}

// GetBatchFromReceiveCity 批量查找 收件人城市
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromReceiveCity(receiveCitys []string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_city` IN (?)", receiveCitys).Find(&results).Error

	return
}

// GetFromReceiveProvince 通过receive_province获取内容 收件人省份
func (obj *_LgOrderAddressErrorTempMgr) GetFromReceiveProvince(receiveProvince string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_province` = ?", receiveProvince).Find(&results).Error

	return
}

// GetBatchFromReceiveProvince 批量查找 收件人省份
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromReceiveProvince(receiveProvinces []string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_province` IN (?)", receiveProvinces).Find(&results).Error

	return
}

// GetFromReceiveCountry 通过receive_country获取内容 收件人国家
func (obj *_LgOrderAddressErrorTempMgr) GetFromReceiveCountry(receiveCountry string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// GetBatchFromReceiveCountry 批量查找 收件人国家
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromReceiveCountry(receiveCountrys []string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_country` IN (?)", receiveCountrys).Find(&results).Error

	return
}

// GetFromReceiveZipCode 通过receive_zip_code获取内容 收件人邮编
func (obj *_LgOrderAddressErrorTempMgr) GetFromReceiveZipCode(receiveZipCode string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_zip_code` = ?", receiveZipCode).Find(&results).Error

	return
}

// GetBatchFromReceiveZipCode 批量查找 收件人邮编
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromReceiveZipCode(receiveZipCodes []string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_zip_code` IN (?)", receiveZipCodes).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgOrderAddressErrorTempMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_LgOrderAddressErrorTempMgr) GetFromCustomerChannelName(customerChannelName string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 归属客户ID
func (obj *_LgOrderAddressErrorTempMgr) GetFromCustomerID(customerID int64) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 归属客户ID
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromCustomerID(customerIDs []int64) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_LgOrderAddressErrorTempMgr) GetFromCustomerAlias(customerAlias string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromErrorInfo 通过error_info获取内容 错误信息
func (obj *_LgOrderAddressErrorTempMgr) GetFromErrorInfo(errorInfo string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`error_info` = ?", errorInfo).Find(&results).Error

	return
}

// GetBatchFromErrorInfo 批量查找 错误信息
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromErrorInfo(errorInfos []string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`error_info` IN (?)", errorInfos).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态 1 已处理 0 未处理
func (obj *_LgOrderAddressErrorTempMgr) GetFromStatus(status int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态 1 已处理 0 未处理
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromStatus(statuss []int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderAddressErrorTempMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderAddressErrorTempMgr) GetFromCreateUser(createUser int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_LgOrderAddressErrorTempMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 修改时间
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderAddressErrorTempMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderAddressErrorTempMgr) GetFromVersion(version int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderAddressErrorTempMgr) GetBatchFromVersion(versions []int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderAddressErrorTempMgr) FetchByPrimaryKey(id int) (result LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByLgOrderReceiveZipcodeReceiveCountryCustomerChannelID primary or index 获取唯一内容
func (obj *_LgOrderAddressErrorTempMgr) FetchUniqueIndexByLgOrderReceiveZipcodeReceiveCountryCustomerChannelID(receiveCountry string, receiveZipCode string, customerChannelID int) (result LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`receive_country` = ? AND `receive_zip_code` = ? AND `customer_channel_id` = ?", receiveCountry, receiveZipCode, customerChannelID).First(&result).Error

	return
}

// FetchIndexByLgOrderAddressErrorTempOrderNumberIndex  获取多个内容
func (obj *_LgOrderAddressErrorTempMgr) FetchIndexByLgOrderAddressErrorTempOrderNumberIndex(orderNumber string) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderAddressErrorTempCustomerChannelIDIndex  获取多个内容
func (obj *_LgOrderAddressErrorTempMgr) FetchIndexByLgOrderAddressErrorTempCustomerChannelIDIndex(customerChannelID int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByLgOrderAddressErrorTempCustomerIDIndex  获取多个内容
func (obj *_LgOrderAddressErrorTempMgr) FetchIndexByLgOrderAddressErrorTempCustomerIDIndex(customerID int64) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexByIDxStatusCreatetime  获取多个内容
func (obj *_LgOrderAddressErrorTempMgr) FetchIndexByIDxStatusCreatetime(status int, createTime time.Time) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`status` = ? AND `create_time` = ?", status, createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderAddressErrorTempStatusIndex  获取多个内容
func (obj *_LgOrderAddressErrorTempMgr) FetchIndexByLgOrderAddressErrorTempStatusIndex(status int) (results []*LgOrderAddressErrorTemp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorTemp{}).Where("`status` = ?", status).Find(&results).Error

	return
}
