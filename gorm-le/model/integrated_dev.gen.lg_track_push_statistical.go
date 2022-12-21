package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgTrackPushStatisticalMgr struct {
	*_BaseMgr
}

// LgTrackPushStatisticalMgr open func
func LgTrackPushStatisticalMgr(db *gorm.DB) *_LgTrackPushStatisticalMgr {
	if db == nil {
		panic(fmt.Errorf("LgTrackPushStatisticalMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgTrackPushStatisticalMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_track_push_statistical"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgTrackPushStatisticalMgr) GetTableName() string {
	return "lg_track_push_statistical"
}

// Reset 重置gorm会话
func (obj *_LgTrackPushStatisticalMgr) Reset() *_LgTrackPushStatisticalMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgTrackPushStatisticalMgr) Get() (result LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgTrackPushStatisticalMgr) Gets() (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgTrackPushStatisticalMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgTrackPushStatisticalMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCount count获取 本次推送数量
func (obj *_LgTrackPushStatisticalMgr) WithCount(count int64) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithType type获取 业务类型:比如巴西邮政菜鸟推送
func (obj *_LgTrackPushStatisticalMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCountry country获取 业务国家
func (obj *_LgTrackPushStatisticalMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithTrackState track_state获取 推送的节点状态码
func (obj *_LgTrackPushStatisticalMgr) WithTrackState(trackState string) Option {
	return optionFunc(func(o *options) { o.query["track_state"] = trackState })
}

// WithPushDate push_date获取 推送日期 yyyy-MM-dd
func (obj *_LgTrackPushStatisticalMgr) WithPushDate(pushDate string) Option {
	return optionFunc(func(o *options) { o.query["push_date"] = pushDate })
}

// WithCreateTime create_time获取 推送时间
func (obj *_LgTrackPushStatisticalMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_LgTrackPushStatisticalMgr) GetByOption(opts ...Option) (result LgTrackPushStatistical, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgTrackPushStatisticalMgr) GetByOptions(opts ...Option) (results []*LgTrackPushStatistical, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgTrackPushStatisticalMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgTrackPushStatistical, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where(options.query)
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
func (obj *_LgTrackPushStatisticalMgr) GetFromID(id int) (result LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LgTrackPushStatisticalMgr) GetBatchFromID(ids []int) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容 本次推送数量
func (obj *_LgTrackPushStatisticalMgr) GetFromCount(count int64) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找 本次推送数量
func (obj *_LgTrackPushStatisticalMgr) GetBatchFromCount(counts []int64) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 业务类型:比如巴西邮政菜鸟推送
func (obj *_LgTrackPushStatisticalMgr) GetFromType(_type string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 业务类型:比如巴西邮政菜鸟推送
func (obj *_LgTrackPushStatisticalMgr) GetBatchFromType(_types []string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 业务国家
func (obj *_LgTrackPushStatisticalMgr) GetFromCountry(country string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`country` = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量查找 业务国家
func (obj *_LgTrackPushStatisticalMgr) GetBatchFromCountry(countrys []string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`country` IN (?)", countrys).Find(&results).Error

	return
}

// GetFromTrackState 通过track_state获取内容 推送的节点状态码
func (obj *_LgTrackPushStatisticalMgr) GetFromTrackState(trackState string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`track_state` = ?", trackState).Find(&results).Error

	return
}

// GetBatchFromTrackState 批量查找 推送的节点状态码
func (obj *_LgTrackPushStatisticalMgr) GetBatchFromTrackState(trackStates []string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`track_state` IN (?)", trackStates).Find(&results).Error

	return
}

// GetFromPushDate 通过push_date获取内容 推送日期 yyyy-MM-dd
func (obj *_LgTrackPushStatisticalMgr) GetFromPushDate(pushDate string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`push_date` = ?", pushDate).Find(&results).Error

	return
}

// GetBatchFromPushDate 批量查找 推送日期 yyyy-MM-dd
func (obj *_LgTrackPushStatisticalMgr) GetBatchFromPushDate(pushDates []string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`push_date` IN (?)", pushDates).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 推送时间
func (obj *_LgTrackPushStatisticalMgr) GetFromCreateTime(createTime time.Time) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 推送时间
func (obj *_LgTrackPushStatisticalMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgTrackPushStatisticalMgr) FetchByPrimaryKey(id int) (result LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByLgTrajectoryTrackingTypeIndex  获取多个内容
func (obj *_LgTrackPushStatisticalMgr) FetchIndexByLgTrajectoryTrackingTypeIndex(_type string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// FetchIndexByLgTrajectoryTrackingTrackStateIndex  获取多个内容
func (obj *_LgTrackPushStatisticalMgr) FetchIndexByLgTrajectoryTrackingTrackStateIndex(trackState string) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`track_state` = ?", trackState).Find(&results).Error

	return
}

// FetchIndexByLgTrajectoryTrackingDatetimeIndex  获取多个内容
func (obj *_LgTrackPushStatisticalMgr) FetchIndexByLgTrajectoryTrackingDatetimeIndex(createTime time.Time) (results []*LgTrackPushStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackPushStatistical{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
