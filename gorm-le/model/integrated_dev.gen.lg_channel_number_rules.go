package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _LgChannelNumberRulesMgr struct {
	*_BaseMgr
}

// LgChannelNumberRulesMgr open func
func LgChannelNumberRulesMgr(db *gorm.DB) *_LgChannelNumberRulesMgr {
	if db == nil {
		panic(fmt.Errorf("LgChannelNumberRulesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgChannelNumberRulesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_channel_number_rules"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgChannelNumberRulesMgr) GetTableName() string {
	return "lg_channel_number_rules"
}

// Reset 重置gorm会话
func (obj *_LgChannelNumberRulesMgr) Reset() *_LgChannelNumberRulesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgChannelNumberRulesMgr) Get() (result LgChannelNumberRules, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgChannelNumberRulesMgr) Gets() (results []*LgChannelNumberRules, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgChannelNumberRulesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithChannelID channel_id获取 渠道ID
func (obj *_LgChannelNumberRulesMgr) WithChannelID(channelID int) Option {
	return optionFunc(func(o *options) { o.query["channel_id"] = channelID })
}

// WithNumberRules number_rules获取 单号规则 | 隔开
func (obj *_LgChannelNumberRulesMgr) WithNumberRules(numberRules string) Option {
	return optionFunc(func(o *options) { o.query["number_rules"] = numberRules })
}

// GetByOption 功能选项模式获取
func (obj *_LgChannelNumberRulesMgr) GetByOption(opts ...Option) (result LgChannelNumberRules, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgChannelNumberRulesMgr) GetByOptions(opts ...Option) (results []*LgChannelNumberRules, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgChannelNumberRulesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgChannelNumberRules, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Where(options.query)
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
func (obj *_LgChannelNumberRulesMgr) GetFromChannelID(channelID int) (result LgChannelNumberRules, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Where("`channel_id` = ?", channelID).First(&result).Error

	return
}

// GetBatchFromChannelID 批量查找 渠道ID
func (obj *_LgChannelNumberRulesMgr) GetBatchFromChannelID(channelIDs []int) (results []*LgChannelNumberRules, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Where("`channel_id` IN (?)", channelIDs).Find(&results).Error

	return
}

// GetFromNumberRules 通过number_rules获取内容 单号规则 | 隔开
func (obj *_LgChannelNumberRulesMgr) GetFromNumberRules(numberRules string) (results []*LgChannelNumberRules, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Where("`number_rules` = ?", numberRules).Find(&results).Error

	return
}

// GetBatchFromNumberRules 批量查找 单号规则 | 隔开
func (obj *_LgChannelNumberRulesMgr) GetBatchFromNumberRules(numberRuless []string) (results []*LgChannelNumberRules, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Where("`number_rules` IN (?)", numberRuless).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgChannelNumberRulesMgr) FetchByPrimaryKey(channelID int) (result LgChannelNumberRules, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelNumberRules{}).Where("`channel_id` = ?", channelID).First(&result).Error

	return
}
