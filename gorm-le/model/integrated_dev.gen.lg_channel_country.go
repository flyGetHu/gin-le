package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _LgChannelCountryMgr struct {
	*_BaseMgr
}

// LgChannelCountryMgr open func
func LgChannelCountryMgr(db *gorm.DB) *_LgChannelCountryMgr {
	if db == nil {
		panic(fmt.Errorf("LgChannelCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgChannelCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_channel_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgChannelCountryMgr) GetTableName() string {
	return "lg_channel_country"
}

// Reset 重置gorm会话
func (obj *_LgChannelCountryMgr) Reset() *_LgChannelCountryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgChannelCountryMgr) Get() (result LgChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgChannelCountryMgr) Gets() (results []*LgChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgChannelCountryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithChannelID channel_id获取 渠道ID
func (obj *_LgChannelCountryMgr) WithChannelID(channelID int) Option {
	return optionFunc(func(o *options) { o.query["channel_id"] = channelID })
}

// WithCountryID country_id获取 国家表ID
func (obj *_LgChannelCountryMgr) WithCountryID(countryID int) Option {
	return optionFunc(func(o *options) { o.query["country_id"] = countryID })
}

// GetByOption 功能选项模式获取
func (obj *_LgChannelCountryMgr) GetByOption(opts ...Option) (result LgChannelCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgChannelCountryMgr) GetByOptions(opts ...Option) (results []*LgChannelCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgChannelCountryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgChannelCountry, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Where(options.query)
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

// GetFromChannelID 通过channel_id获取内容 渠道ID
func (obj *_LgChannelCountryMgr) GetFromChannelID(channelID int) (results []*LgChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Where("`channel_id` = ?", channelID).Find(&results).Error

	return
}

// GetBatchFromChannelID 批量查找 渠道ID
func (obj *_LgChannelCountryMgr) GetBatchFromChannelID(channelIDs []int) (results []*LgChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Where("`channel_id` IN (?)", channelIDs).Find(&results).Error

	return
}

// GetFromCountryID 通过country_id获取内容 国家表ID
func (obj *_LgChannelCountryMgr) GetFromCountryID(countryID int) (results []*LgChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Where("`country_id` = ?", countryID).Find(&results).Error

	return
}

// GetBatchFromCountryID 批量查找 国家表ID
func (obj *_LgChannelCountryMgr) GetBatchFromCountryID(countryIDs []int) (results []*LgChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelCountry{}).Where("`country_id` IN (?)", countryIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
