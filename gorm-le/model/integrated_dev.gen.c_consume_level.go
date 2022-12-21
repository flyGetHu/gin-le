package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _CConsumeLevelMgr struct {
	*_BaseMgr
}

// CConsumeLevelMgr open func
func CConsumeLevelMgr(db *gorm.DB) *_CConsumeLevelMgr {
	if db == nil {
		panic(fmt.Errorf("CConsumeLevelMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CConsumeLevelMgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_consume_level"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CConsumeLevelMgr) GetTableName() string {
	return "c_consume_level"
}

// Reset 重置gorm会话
func (obj *_CConsumeLevelMgr) Reset() *_CConsumeLevelMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CConsumeLevelMgr) Get() (result CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CConsumeLevelMgr) Gets() (results []*CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CConsumeLevelMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CConsumeLevelMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithLevel level获取 消费等级
func (obj *_CConsumeLevelMgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithLevelDesc level_desc获取 等级描述
func (obj *_CConsumeLevelMgr) WithLevelDesc(levelDesc string) Option {
	return optionFunc(func(o *options) { o.query["level_desc"] = levelDesc })
}

// GetByOption 功能选项模式获取
func (obj *_CConsumeLevelMgr) GetByOption(opts ...Option) (result CConsumeLevel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CConsumeLevelMgr) GetByOptions(opts ...Option) (results []*CConsumeLevel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CConsumeLevelMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CConsumeLevel, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where(options.query)
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
func (obj *_CConsumeLevelMgr) GetFromID(id int) (result CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_CConsumeLevelMgr) GetBatchFromID(ids []int) (results []*CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 消费等级
func (obj *_CConsumeLevelMgr) GetFromLevel(level int) (result CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`level` = ?", level).First(&result).Error

	return
}

// GetBatchFromLevel 批量查找 消费等级
func (obj *_CConsumeLevelMgr) GetBatchFromLevel(levels []int) (results []*CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromLevelDesc 通过level_desc获取内容 等级描述
func (obj *_CConsumeLevelMgr) GetFromLevelDesc(levelDesc string) (result CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`level_desc` = ?", levelDesc).First(&result).Error

	return
}

// GetBatchFromLevelDesc 批量查找 等级描述
func (obj *_CConsumeLevelMgr) GetBatchFromLevelDesc(levelDescs []string) (results []*CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`level_desc` IN (?)", levelDescs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CConsumeLevelMgr) FetchByPrimaryKey(id int) (result CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexLevel primary or index 获取唯一内容
func (obj *_CConsumeLevelMgr) FetchUniqueByIndexLevel(level int) (result CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`level` = ?", level).First(&result).Error

	return
}

// FetchUniqueByIndexLevelDesc primary or index 获取唯一内容
func (obj *_CConsumeLevelMgr) FetchUniqueByIndexLevelDesc(levelDesc string) (result CConsumeLevel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeLevel{}).Where("`level_desc` = ?", levelDesc).First(&result).Error

	return
}
