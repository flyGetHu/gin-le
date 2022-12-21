package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderAddressErrorCorrectionMgr struct {
	*_BaseMgr
}

// LgOrderAddressErrorCorrectionMgr open func
func LgOrderAddressErrorCorrectionMgr(db *gorm.DB) *_LgOrderAddressErrorCorrectionMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderAddressErrorCorrectionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderAddressErrorCorrectionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_address_error_correction"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderAddressErrorCorrectionMgr) GetTableName() string {
	return "lg_order_address_error_correction"
}

// Reset 重置gorm会话
func (obj *_LgOrderAddressErrorCorrectionMgr) Reset() *_LgOrderAddressErrorCorrectionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderAddressErrorCorrectionMgr) Get() (result LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderAddressErrorCorrectionMgr) Gets() (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderAddressErrorCorrectionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgOrderAddressErrorCorrectionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithErrorZipCode error_zip_code获取 错误邮编
func (obj *_LgOrderAddressErrorCorrectionMgr) WithErrorZipCode(errorZipCode string) Option {
	return optionFunc(func(o *options) { o.query["error_zip_code"] = errorZipCode })
}

// WithCorrectZipCode correct_zip_code获取 正确邮编
func (obj *_LgOrderAddressErrorCorrectionMgr) WithCorrectZipCode(correctZipCode string) Option {
	return optionFunc(func(o *options) { o.query["correct_zip_code"] = correctZipCode })
}

// WithReceiveProvince receive_province获取 收件人省份
func (obj *_LgOrderAddressErrorCorrectionMgr) WithReceiveProvince(receiveProvince string) Option {
	return optionFunc(func(o *options) { o.query["receive_province"] = receiveProvince })
}

// WithReceiveCity receive_city获取 收件人城市
func (obj *_LgOrderAddressErrorCorrectionMgr) WithReceiveCity(receiveCity string) Option {
	return optionFunc(func(o *options) { o.query["receive_city"] = receiveCity })
}

// WithReceiveCountry receive_country获取 收件人国家
func (obj *_LgOrderAddressErrorCorrectionMgr) WithReceiveCountry(receiveCountry string) Option {
	return optionFunc(func(o *options) { o.query["receive_country"] = receiveCountry })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgOrderAddressErrorCorrectionMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_LgOrderAddressErrorCorrectionMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithState state获取 状态(0:禁用，1:启用)
func (obj *_LgOrderAddressErrorCorrectionMgr) WithState(state []uint8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderAddressErrorCorrectionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderAddressErrorCorrectionMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_LgOrderAddressErrorCorrectionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderAddressErrorCorrectionMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderAddressErrorCorrectionMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderAddressErrorCorrectionMgr) GetByOption(opts ...Option) (result LgOrderAddressErrorCorrection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderAddressErrorCorrectionMgr) GetByOptions(opts ...Option) (results []*LgOrderAddressErrorCorrection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderAddressErrorCorrectionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderAddressErrorCorrection, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where(options.query)
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
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromID(id int) (result LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromID(ids []int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromErrorZipCode 通过error_zip_code获取内容 错误邮编
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromErrorZipCode(errorZipCode string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`error_zip_code` = ?", errorZipCode).Find(&results).Error

	return
}

// GetBatchFromErrorZipCode 批量查找 错误邮编
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromErrorZipCode(errorZipCodes []string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`error_zip_code` IN (?)", errorZipCodes).Find(&results).Error

	return
}

// GetFromCorrectZipCode 通过correct_zip_code获取内容 正确邮编
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromCorrectZipCode(correctZipCode string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`correct_zip_code` = ?", correctZipCode).Find(&results).Error

	return
}

// GetBatchFromCorrectZipCode 批量查找 正确邮编
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromCorrectZipCode(correctZipCodes []string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`correct_zip_code` IN (?)", correctZipCodes).Find(&results).Error

	return
}

// GetFromReceiveProvince 通过receive_province获取内容 收件人省份
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromReceiveProvince(receiveProvince string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_province` = ?", receiveProvince).Find(&results).Error

	return
}

// GetBatchFromReceiveProvince 批量查找 收件人省份
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromReceiveProvince(receiveProvinces []string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_province` IN (?)", receiveProvinces).Find(&results).Error

	return
}

// GetFromReceiveCity 通过receive_city获取内容 收件人城市
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromReceiveCity(receiveCity string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_city` = ?", receiveCity).Find(&results).Error

	return
}

// GetBatchFromReceiveCity 批量查找 收件人城市
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromReceiveCity(receiveCitys []string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_city` IN (?)", receiveCitys).Find(&results).Error

	return
}

// GetFromReceiveCountry 通过receive_country获取内容 收件人国家
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromReceiveCountry(receiveCountry string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// GetBatchFromReceiveCountry 批量查找 收件人国家
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromReceiveCountry(receiveCountrys []string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_country` IN (?)", receiveCountrys).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromCustomerChannelName(customerChannelName string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态(0:禁用，1:启用)
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromState(state []uint8) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态(0:禁用，1:启用)
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromState(states [][]uint8) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromCreateUser(createUser int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 修改时间
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderAddressErrorCorrectionMgr) GetFromVersion(version int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderAddressErrorCorrectionMgr) GetBatchFromVersion(versions []int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderAddressErrorCorrectionMgr) FetchByPrimaryKey(id int) (result LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByErrorZipcodeReceiveCountryCustomerChannelIDIndex primary or index 获取唯一内容
func (obj *_LgOrderAddressErrorCorrectionMgr) FetchUniqueIndexByErrorZipcodeReceiveCountryCustomerChannelIDIndex(errorZipCode string, receiveCountry string, customerChannelID int) (result LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`error_zip_code` = ? AND `receive_country` = ? AND `customer_channel_id` = ?", errorZipCode, receiveCountry, customerChannelID).First(&result).Error

	return
}

// FetchIndexByLgOrderAddressErrorCorrectionErrorZipCodeIndex  获取多个内容
func (obj *_LgOrderAddressErrorCorrectionMgr) FetchIndexByLgOrderAddressErrorCorrectionErrorZipCodeIndex(errorZipCode string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`error_zip_code` = ?", errorZipCode).Find(&results).Error

	return
}

// FetchIndexByLgOrderAddressErrorCorrectionCorrectZipCodeIndex  获取多个内容
func (obj *_LgOrderAddressErrorCorrectionMgr) FetchIndexByLgOrderAddressErrorCorrectionCorrectZipCodeIndex(correctZipCode string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`correct_zip_code` = ?", correctZipCode).Find(&results).Error

	return
}

// FetchIndexByLgOrderAddressErrorCorrectionReceiveProvinceIndex  获取多个内容
func (obj *_LgOrderAddressErrorCorrectionMgr) FetchIndexByLgOrderAddressErrorCorrectionReceiveProvinceIndex(receiveProvince string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_province` = ?", receiveProvince).Find(&results).Error

	return
}

// FetchIndexByLgOrderAddressErrorCorrectionReceiveCityIndex  获取多个内容
func (obj *_LgOrderAddressErrorCorrectionMgr) FetchIndexByLgOrderAddressErrorCorrectionReceiveCityIndex(receiveCity string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_city` = ?", receiveCity).Find(&results).Error

	return
}

// FetchIndexByLgOrderAddressErrorCorrectionReceiveCountryIndex  获取多个内容
func (obj *_LgOrderAddressErrorCorrectionMgr) FetchIndexByLgOrderAddressErrorCorrectionReceiveCountryIndex(receiveCountry string) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// FetchIndexByLgOrderAddressErrorCorrectionCustomerChannelIDIndex  获取多个内容
func (obj *_LgOrderAddressErrorCorrectionMgr) FetchIndexByLgOrderAddressErrorCorrectionCustomerChannelIDIndex(customerChannelID int) (results []*LgOrderAddressErrorCorrection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderAddressErrorCorrection{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}
