package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgProviderChannelRuleMgr struct {
	*_BaseMgr
}

// LgProviderChannelRuleMgr open func
func LgProviderChannelRuleMgr(db *gorm.DB) *_LgProviderChannelRuleMgr {
	if db == nil {
		panic(fmt.Errorf("LgProviderChannelRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgProviderChannelRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_provider_channel_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgProviderChannelRuleMgr) GetTableName() string {
	return "lg_provider_channel_rule"
}

// Reset 重置gorm会话
func (obj *_LgProviderChannelRuleMgr) Reset() *_LgProviderChannelRuleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgProviderChannelRuleMgr) Get() (result LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgProviderChannelRuleMgr) Gets() (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgProviderChannelRuleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgProviderChannelRuleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithChannelID channel_id获取 服务商渠道ID
func (obj *_LgProviderChannelRuleMgr) WithChannelID(channelID int) Option {
	return optionFunc(func(o *options) { o.query["channel_id"] = channelID })
}

// WithCountryTwoCode country_two_code获取 国家二字代码
func (obj *_LgProviderChannelRuleMgr) WithCountryTwoCode(countryTwoCode string) Option {
	return optionFunc(func(o *options) { o.query["country_two_code"] = countryTwoCode })
}

// WithChannelCode channel_code获取 服务商渠道Code
func (obj *_LgProviderChannelRuleMgr) WithChannelCode(channelCode string) Option {
	return optionFunc(func(o *options) { o.query["channel_code"] = channelCode })
}

// WithLabelCode label_code获取 面单代码
func (obj *_LgProviderChannelRuleMgr) WithLabelCode(labelCode string) Option {
	return optionFunc(func(o *options) { o.query["label_code"] = labelCode })
}

// WithStartZipcode start_zipcode获取 开始邮编
func (obj *_LgProviderChannelRuleMgr) WithStartZipcode(startZipcode int) Option {
	return optionFunc(func(o *options) { o.query["start_zipcode"] = startZipcode })
}

// WithEndZipcode end_zipcode获取 结束邮编
func (obj *_LgProviderChannelRuleMgr) WithEndZipcode(endZipcode int) Option {
	return optionFunc(func(o *options) { o.query["end_zipcode"] = endZipcode })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgProviderChannelRuleMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgProviderChannelRuleMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgProviderChannelRuleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgProviderChannelRuleMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgProviderChannelRuleMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgProviderChannelRuleMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgProviderChannelRuleMgr) GetByOption(opts ...Option) (result LgProviderChannelRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgProviderChannelRuleMgr) GetByOptions(opts ...Option) (results []*LgProviderChannelRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgProviderChannelRuleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgProviderChannelRule, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where(options.query)
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
func (obj *_LgProviderChannelRuleMgr) GetFromID(id int) (result LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgProviderChannelRuleMgr) GetBatchFromID(ids []int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromChannelID 通过channel_id获取内容 服务商渠道ID
func (obj *_LgProviderChannelRuleMgr) GetFromChannelID(channelID int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`channel_id` = ?", channelID).Find(&results).Error

	return
}

// GetBatchFromChannelID 批量查找 服务商渠道ID
func (obj *_LgProviderChannelRuleMgr) GetBatchFromChannelID(channelIDs []int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`channel_id` IN (?)", channelIDs).Find(&results).Error

	return
}

// GetFromCountryTwoCode 通过country_two_code获取内容 国家二字代码
func (obj *_LgProviderChannelRuleMgr) GetFromCountryTwoCode(countryTwoCode string) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`country_two_code` = ?", countryTwoCode).Find(&results).Error

	return
}

// GetBatchFromCountryTwoCode 批量查找 国家二字代码
func (obj *_LgProviderChannelRuleMgr) GetBatchFromCountryTwoCode(countryTwoCodes []string) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`country_two_code` IN (?)", countryTwoCodes).Find(&results).Error

	return
}

// GetFromChannelCode 通过channel_code获取内容 服务商渠道Code
func (obj *_LgProviderChannelRuleMgr) GetFromChannelCode(channelCode string) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`channel_code` = ?", channelCode).Find(&results).Error

	return
}

// GetBatchFromChannelCode 批量查找 服务商渠道Code
func (obj *_LgProviderChannelRuleMgr) GetBatchFromChannelCode(channelCodes []string) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`channel_code` IN (?)", channelCodes).Find(&results).Error

	return
}

// GetFromLabelCode 通过label_code获取内容 面单代码
func (obj *_LgProviderChannelRuleMgr) GetFromLabelCode(labelCode string) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`label_code` = ?", labelCode).Find(&results).Error

	return
}

// GetBatchFromLabelCode 批量查找 面单代码
func (obj *_LgProviderChannelRuleMgr) GetBatchFromLabelCode(labelCodes []string) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`label_code` IN (?)", labelCodes).Find(&results).Error

	return
}

// GetFromStartZipcode 通过start_zipcode获取内容 开始邮编
func (obj *_LgProviderChannelRuleMgr) GetFromStartZipcode(startZipcode int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`start_zipcode` = ?", startZipcode).Find(&results).Error

	return
}

// GetBatchFromStartZipcode 批量查找 开始邮编
func (obj *_LgProviderChannelRuleMgr) GetBatchFromStartZipcode(startZipcodes []int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`start_zipcode` IN (?)", startZipcodes).Find(&results).Error

	return
}

// GetFromEndZipcode 通过end_zipcode获取内容 结束邮编
func (obj *_LgProviderChannelRuleMgr) GetFromEndZipcode(endZipcode int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`end_zipcode` = ?", endZipcode).Find(&results).Error

	return
}

// GetBatchFromEndZipcode 批量查找 结束邮编
func (obj *_LgProviderChannelRuleMgr) GetBatchFromEndZipcode(endZipcodes []int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`end_zipcode` IN (?)", endZipcodes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgProviderChannelRuleMgr) GetFromCreateTime(createTime time.Time) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgProviderChannelRuleMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgProviderChannelRuleMgr) GetFromCreateUser(createUser int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgProviderChannelRuleMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgProviderChannelRuleMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgProviderChannelRuleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgProviderChannelRuleMgr) GetFromUpdateUser(updateUser int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgProviderChannelRuleMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgProviderChannelRuleMgr) GetFromVersion(version int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgProviderChannelRuleMgr) GetBatchFromVersion(versions []int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgProviderChannelRuleMgr) GetFromDeleted(deleted int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgProviderChannelRuleMgr) GetBatchFromDeleted(deleteds []int) (results []*LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgProviderChannelRuleMgr) FetchByPrimaryKey(id int) (result LgProviderChannelRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelRule{}).Where("`id` = ?", id).First(&result).Error

	return
}
