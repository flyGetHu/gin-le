package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _LgOrderStatisticalMgr struct {
	*_BaseMgr
}

// LgOrderStatisticalMgr open func
func LgOrderStatisticalMgr(db *gorm.DB) *_LgOrderStatisticalMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderStatisticalMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderStatisticalMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_statistical"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderStatisticalMgr) GetTableName() string {
	return "lg_order_statistical"
}

// Reset 重置gorm会话
func (obj *_LgOrderStatisticalMgr) Reset() *_LgOrderStatisticalMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderStatisticalMgr) Get() (result LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderStatisticalMgr) Gets() (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderStatisticalMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithType type获取 统计类型
func (obj *_LgOrderStatisticalMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithDate date获取 日期
func (obj *_LgOrderStatisticalMgr) WithDate(date datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// WithDateType date_type获取 时间类型
func (obj *_LgOrderStatisticalMgr) WithDateType(dateType int) Option {
	return optionFunc(func(o *options) { o.query["date_type"] = dateType })
}

// WithData data获取 订单统计数据
func (obj *_LgOrderStatisticalMgr) WithData(data datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["data"] = data })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderStatisticalMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderStatisticalMgr) GetByOption(opts ...Option) (result LgOrderStatistical, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderStatisticalMgr) GetByOptions(opts ...Option) (results []*LgOrderStatistical, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderStatisticalMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderStatistical, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where(options.query)
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

// GetFromType 通过type获取内容 统计类型
func (obj *_LgOrderStatisticalMgr) GetFromType(_type int) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 统计类型
func (obj *_LgOrderStatisticalMgr) GetBatchFromType(_types []int) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容 日期
func (obj *_LgOrderStatisticalMgr) GetFromDate(date datatypes.Date) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找 日期
func (obj *_LgOrderStatisticalMgr) GetBatchFromDate(dates []datatypes.Date) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

// GetFromDateType 通过date_type获取内容 时间类型
func (obj *_LgOrderStatisticalMgr) GetFromDateType(dateType int) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`date_type` = ?", dateType).Find(&results).Error

	return
}

// GetBatchFromDateType 批量查找 时间类型
func (obj *_LgOrderStatisticalMgr) GetBatchFromDateType(dateTypes []int) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`date_type` IN (?)", dateTypes).Find(&results).Error

	return
}

// GetFromData 通过data获取内容 订单统计数据
func (obj *_LgOrderStatisticalMgr) GetFromData(data datatypes.JSON) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`data` = ?", data).Find(&results).Error

	return
}

// GetBatchFromData 批量查找 订单统计数据
func (obj *_LgOrderStatisticalMgr) GetBatchFromData(datas []datatypes.JSON) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`data` IN (?)", datas).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderStatisticalMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderStatisticalMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByLgOrderStatisticalCreateTimeIndex  获取多个内容
func (obj *_LgOrderStatisticalMgr) FetchIndexByLgOrderStatisticalCreateTimeIndex(createTime time.Time) (results []*LgOrderStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatistical{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
