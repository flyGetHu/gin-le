package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _FaStatisticalMgr struct {
	*_BaseMgr
}

// FaStatisticalMgr open func
func FaStatisticalMgr(db *gorm.DB) *_FaStatisticalMgr {
	if db == nil {
		panic(fmt.Errorf("FaStatisticalMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaStatisticalMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_statistical"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaStatisticalMgr) GetTableName() string {
	return "fa_statistical"
}

// Reset 重置gorm会话
func (obj *_FaStatisticalMgr) Reset() *_FaStatisticalMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaStatisticalMgr) Get() (result FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaStatisticalMgr) Gets() (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaStatisticalMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithType type获取 报表类型，1:客户渠道毛利计算报表
func (obj *_FaStatisticalMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithDateType date_type获取 日期类型，1:订单创建时间、2:应收费用明细创建时间、3:应付费用明细创建时间、4:入库时间、5:出库时间
func (obj *_FaStatisticalMgr) WithDateType(dateType int) Option {
	return optionFunc(func(o *options) { o.query["date_type"] = dateType })
}

// WithDate date获取 日期
func (obj *_FaStatisticalMgr) WithDate(date datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// WithData data获取 统计数据
func (obj *_FaStatisticalMgr) WithData(data datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["data"] = data })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaStatisticalMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_FaStatisticalMgr) GetByOption(opts ...Option) (result FaStatistical, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaStatisticalMgr) GetByOptions(opts ...Option) (results []*FaStatistical, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaStatisticalMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaStatistical, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where(options.query)
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

// GetFromType 通过type获取内容 报表类型，1:客户渠道毛利计算报表
func (obj *_FaStatisticalMgr) GetFromType(_type int) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 报表类型，1:客户渠道毛利计算报表
func (obj *_FaStatisticalMgr) GetBatchFromType(_types []int) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromDateType 通过date_type获取内容 日期类型，1:订单创建时间、2:应收费用明细创建时间、3:应付费用明细创建时间、4:入库时间、5:出库时间
func (obj *_FaStatisticalMgr) GetFromDateType(dateType int) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`date_type` = ?", dateType).Find(&results).Error

	return
}

// GetBatchFromDateType 批量查找 日期类型，1:订单创建时间、2:应收费用明细创建时间、3:应付费用明细创建时间、4:入库时间、5:出库时间
func (obj *_FaStatisticalMgr) GetBatchFromDateType(dateTypes []int) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`date_type` IN (?)", dateTypes).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容 日期
func (obj *_FaStatisticalMgr) GetFromDate(date datatypes.Date) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找 日期
func (obj *_FaStatisticalMgr) GetBatchFromDate(dates []datatypes.Date) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

// GetFromData 通过data获取内容 统计数据
func (obj *_FaStatisticalMgr) GetFromData(data datatypes.JSON) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`data` = ?", data).Find(&results).Error

	return
}

// GetBatchFromData 批量查找 统计数据
func (obj *_FaStatisticalMgr) GetBatchFromData(datas []datatypes.JSON) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`data` IN (?)", datas).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaStatisticalMgr) GetFromCreateTime(createTime time.Time) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaStatisticalMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByIndexType  获取多个内容
func (obj *_FaStatisticalMgr) FetchIndexByIndexType(_type int) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// FetchIndexByIndexTypeDateDateType  获取多个内容
func (obj *_FaStatisticalMgr) FetchIndexByIndexTypeDateDateType(_type int, dateType int, date datatypes.Date) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`type` = ? AND `date_type` = ? AND `date` = ?", _type, dateType, date).Find(&results).Error

	return
}

// FetchIndexByIndexDateType  获取多个内容
func (obj *_FaStatisticalMgr) FetchIndexByIndexDateType(dateType int) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`date_type` = ?", dateType).Find(&results).Error

	return
}

// FetchIndexByIndexDate  获取多个内容
func (obj *_FaStatisticalMgr) FetchIndexByIndexDate(date datatypes.Date) (results []*FaStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaStatistical{}).Where("`date` = ?", date).Find(&results).Error

	return
}
