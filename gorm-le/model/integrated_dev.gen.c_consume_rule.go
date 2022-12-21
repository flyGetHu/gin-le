package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CConsumeRuleMgr struct {
	*_BaseMgr
}

// CConsumeRuleMgr open func
func CConsumeRuleMgr(db *gorm.DB) *_CConsumeRuleMgr {
	if db == nil {
		panic(fmt.Errorf("CConsumeRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CConsumeRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_consume_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CConsumeRuleMgr) GetTableName() string {
	return "c_consume_rule"
}

// Reset 重置gorm会话
func (obj *_CConsumeRuleMgr) Reset() *_CConsumeRuleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CConsumeRuleMgr) Get() (result CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CConsumeRuleMgr) Gets() (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CConsumeRuleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CConsumeRuleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRuleType rule_type获取 规则类型，0:累计金额达标，1:周期内消费总金额达标
func (obj *_CConsumeRuleMgr) WithRuleType(ruleType int) Option {
	return optionFunc(func(o *options) { o.query["rule_type"] = ruleType })
}

// WithRuleDesc rule_desc获取 规则描述
func (obj *_CConsumeRuleMgr) WithRuleDesc(ruleDesc string) Option {
	return optionFunc(func(o *options) { o.query["rule_desc"] = ruleDesc })
}

// WithCalculatePeriod calculate_period获取 计费周期(单位:天)
func (obj *_CConsumeRuleMgr) WithCalculatePeriod(calculatePeriod int) Option {
	return optionFunc(func(o *options) { o.query["calculate_period"] = calculatePeriod })
}

// WithRuleAmount rule_amount获取 规则达标所需金额
func (obj *_CConsumeRuleMgr) WithRuleAmount(ruleAmount float64) Option {
	return optionFunc(func(o *options) { o.query["rule_amount"] = ruleAmount })
}

// WithStartTime start_time获取 规则开始时间
func (obj *_CConsumeRuleMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 规则结束时间
func (obj *_CConsumeRuleMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// GetByOption 功能选项模式获取
func (obj *_CConsumeRuleMgr) GetByOption(opts ...Option) (result CConsumeRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CConsumeRuleMgr) GetByOptions(opts ...Option) (results []*CConsumeRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CConsumeRuleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CConsumeRule, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where(options.query)
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
func (obj *_CConsumeRuleMgr) GetFromID(id int) (result CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_CConsumeRuleMgr) GetBatchFromID(ids []int) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromRuleType 通过rule_type获取内容 规则类型，0:累计金额达标，1:周期内消费总金额达标
func (obj *_CConsumeRuleMgr) GetFromRuleType(ruleType int) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`rule_type` = ?", ruleType).Find(&results).Error

	return
}

// GetBatchFromRuleType 批量查找 规则类型，0:累计金额达标，1:周期内消费总金额达标
func (obj *_CConsumeRuleMgr) GetBatchFromRuleType(ruleTypes []int) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`rule_type` IN (?)", ruleTypes).Find(&results).Error

	return
}

// GetFromRuleDesc 通过rule_desc获取内容 规则描述
func (obj *_CConsumeRuleMgr) GetFromRuleDesc(ruleDesc string) (result CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`rule_desc` = ?", ruleDesc).First(&result).Error

	return
}

// GetBatchFromRuleDesc 批量查找 规则描述
func (obj *_CConsumeRuleMgr) GetBatchFromRuleDesc(ruleDescs []string) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`rule_desc` IN (?)", ruleDescs).Find(&results).Error

	return
}

// GetFromCalculatePeriod 通过calculate_period获取内容 计费周期(单位:天)
func (obj *_CConsumeRuleMgr) GetFromCalculatePeriod(calculatePeriod int) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`calculate_period` = ?", calculatePeriod).Find(&results).Error

	return
}

// GetBatchFromCalculatePeriod 批量查找 计费周期(单位:天)
func (obj *_CConsumeRuleMgr) GetBatchFromCalculatePeriod(calculatePeriods []int) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`calculate_period` IN (?)", calculatePeriods).Find(&results).Error

	return
}

// GetFromRuleAmount 通过rule_amount获取内容 规则达标所需金额
func (obj *_CConsumeRuleMgr) GetFromRuleAmount(ruleAmount float64) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`rule_amount` = ?", ruleAmount).Find(&results).Error

	return
}

// GetBatchFromRuleAmount 批量查找 规则达标所需金额
func (obj *_CConsumeRuleMgr) GetBatchFromRuleAmount(ruleAmounts []float64) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`rule_amount` IN (?)", ruleAmounts).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 规则开始时间
func (obj *_CConsumeRuleMgr) GetFromStartTime(startTime time.Time) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 规则开始时间
func (obj *_CConsumeRuleMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 规则结束时间
func (obj *_CConsumeRuleMgr) GetFromEndTime(endTime time.Time) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 规则结束时间
func (obj *_CConsumeRuleMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CConsumeRuleMgr) FetchByPrimaryKey(id int) (result CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexRuleDesc primary or index 获取唯一内容
func (obj *_CConsumeRuleMgr) FetchUniqueByIndexRuleDesc(ruleDesc string) (result CConsumeRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CConsumeRule{}).Where("`rule_desc` = ?", ruleDesc).First(&result).Error

	return
}
