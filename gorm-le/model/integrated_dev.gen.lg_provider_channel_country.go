package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgProviderChannelCountryMgr struct {
	*_BaseMgr
}

// LgProviderChannelCountryMgr open func
func LgProviderChannelCountryMgr(db *gorm.DB) *_LgProviderChannelCountryMgr {
	if db == nil {
		panic(fmt.Errorf("LgProviderChannelCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgProviderChannelCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_provider_channel_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgProviderChannelCountryMgr) GetTableName() string {
	return "lg_provider_channel_country"
}

// Reset 重置gorm会话
func (obj *_LgProviderChannelCountryMgr) Reset() *_LgProviderChannelCountryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgProviderChannelCountryMgr) Get() (result LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgProviderChannelCountryMgr) Gets() (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgProviderChannelCountryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgProviderChannelCountryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithProviderChannelID provider_channel_id获取 服务商渠道ID
func (obj *_LgProviderChannelCountryMgr) WithProviderChannelID(providerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_id"] = providerChannelID })
}

// WithProviderChannelCode provider_channel_code获取 服务商渠道code
func (obj *_LgProviderChannelCountryMgr) WithProviderChannelCode(providerChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_code"] = providerChannelCode })
}

// WithCountryCode country_code获取 国家二字代码
func (obj *_LgProviderChannelCountryMgr) WithCountryCode(countryCode string) Option {
	return optionFunc(func(o *options) { o.query["country_code"] = countryCode })
}

// WithZipcodes zipcodes获取 黑名单邮编
func (obj *_LgProviderChannelCountryMgr) WithZipcodes(zipcodes string) Option {
	return optionFunc(func(o *options) { o.query["zipcodes"] = zipcodes })
}

// WithSenderAddressBookID sender_address_book_id获取 发件人地址簿id
func (obj *_LgProviderChannelCountryMgr) WithSenderAddressBookID(senderAddressBookID int) Option {
	return optionFunc(func(o *options) { o.query["sender_address_book_id"] = senderAddressBookID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgProviderChannelCountryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgProviderChannelCountryMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgProviderChannelCountryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgProviderChannelCountryMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgProviderChannelCountryMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgProviderChannelCountryMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgProviderChannelCountryMgr) GetByOption(opts ...Option) (result LgProviderChannelCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgProviderChannelCountryMgr) GetByOptions(opts ...Option) (results []*LgProviderChannelCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgProviderChannelCountryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgProviderChannelCountry, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where(options.query)
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
func (obj *_LgProviderChannelCountryMgr) GetFromID(id int) (result LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgProviderChannelCountryMgr) GetBatchFromID(ids []int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromProviderChannelID 通过provider_channel_id获取内容 服务商渠道ID
func (obj *_LgProviderChannelCountryMgr) GetFromProviderChannelID(providerChannelID int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`provider_channel_id` = ?", providerChannelID).Find(&results).Error

	return
}

// GetBatchFromProviderChannelID 批量查找 服务商渠道ID
func (obj *_LgProviderChannelCountryMgr) GetBatchFromProviderChannelID(providerChannelIDs []int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`provider_channel_id` IN (?)", providerChannelIDs).Find(&results).Error

	return
}

// GetFromProviderChannelCode 通过provider_channel_code获取内容 服务商渠道code
func (obj *_LgProviderChannelCountryMgr) GetFromProviderChannelCode(providerChannelCode string) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error

	return
}

// GetBatchFromProviderChannelCode 批量查找 服务商渠道code
func (obj *_LgProviderChannelCountryMgr) GetBatchFromProviderChannelCode(providerChannelCodes []string) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`provider_channel_code` IN (?)", providerChannelCodes).Find(&results).Error

	return
}

// GetFromCountryCode 通过country_code获取内容 国家二字代码
func (obj *_LgProviderChannelCountryMgr) GetFromCountryCode(countryCode string) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`country_code` = ?", countryCode).Find(&results).Error

	return
}

// GetBatchFromCountryCode 批量查找 国家二字代码
func (obj *_LgProviderChannelCountryMgr) GetBatchFromCountryCode(countryCodes []string) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`country_code` IN (?)", countryCodes).Find(&results).Error

	return
}

// GetFromZipcodes 通过zipcodes获取内容 黑名单邮编
func (obj *_LgProviderChannelCountryMgr) GetFromZipcodes(zipcodes string) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`zipcodes` = ?", zipcodes).Find(&results).Error

	return
}

// GetBatchFromZipcodes 批量查找 黑名单邮编
func (obj *_LgProviderChannelCountryMgr) GetBatchFromZipcodes(zipcodess []string) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`zipcodes` IN (?)", zipcodess).Find(&results).Error

	return
}

// GetFromSenderAddressBookID 通过sender_address_book_id获取内容 发件人地址簿id
func (obj *_LgProviderChannelCountryMgr) GetFromSenderAddressBookID(senderAddressBookID int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`sender_address_book_id` = ?", senderAddressBookID).Find(&results).Error

	return
}

// GetBatchFromSenderAddressBookID 批量查找 发件人地址簿id
func (obj *_LgProviderChannelCountryMgr) GetBatchFromSenderAddressBookID(senderAddressBookIDs []int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`sender_address_book_id` IN (?)", senderAddressBookIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgProviderChannelCountryMgr) GetFromCreateTime(createTime time.Time) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgProviderChannelCountryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgProviderChannelCountryMgr) GetFromCreateUser(createUser int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgProviderChannelCountryMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgProviderChannelCountryMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgProviderChannelCountryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgProviderChannelCountryMgr) GetFromUpdateUser(updateUser int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgProviderChannelCountryMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgProviderChannelCountryMgr) GetFromVersion(version int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgProviderChannelCountryMgr) GetBatchFromVersion(versions []int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgProviderChannelCountryMgr) GetFromDeleted(deleted int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgProviderChannelCountryMgr) GetBatchFromDeleted(deleteds []int) (results []*LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgProviderChannelCountryMgr) FetchByPrimaryKey(id int) (result LgProviderChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCountry{}).Where("`id` = ?", id).First(&result).Error

	return
}
