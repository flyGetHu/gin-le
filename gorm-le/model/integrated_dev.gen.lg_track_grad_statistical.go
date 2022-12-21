package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgTrackGradStatisticalMgr struct {
	*_BaseMgr
}

// LgTrackGradStatisticalMgr open func
func LgTrackGradStatisticalMgr(db *gorm.DB) *_LgTrackGradStatisticalMgr {
	if db == nil {
		panic(fmt.Errorf("LgTrackGradStatisticalMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgTrackGradStatisticalMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_track_grad_statistical"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgTrackGradStatisticalMgr) GetTableName() string {
	return "lg_track_grad_statistical"
}

// Reset 重置gorm会话
func (obj *_LgTrackGradStatisticalMgr) Reset() *_LgTrackGradStatisticalMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgTrackGradStatisticalMgr) Get() (result LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgTrackGradStatisticalMgr) Gets() (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgTrackGradStatisticalMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgTrackGradStatisticalMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithShipmentCode shipment_code获取 物流渠道代码
func (obj *_LgTrackGradStatisticalMgr) WithShipmentCode(shipmentCode string) Option {
	return optionFunc(func(o *options) { o.query["shipment_code"] = shipmentCode })
}

// WithStateCode state_code获取 轨迹状态
func (obj *_LgTrackGradStatisticalMgr) WithStateCode(stateCode string) Option {
	return optionFunc(func(o *options) { o.query["state_code"] = stateCode })
}

// WithCount count获取 抓取数量
func (obj *_LgTrackGradStatisticalMgr) WithCount(count int) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgTrackGradStatisticalMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_LgTrackGradStatisticalMgr) GetByOption(opts ...Option) (result LgTrackGradStatistical, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgTrackGradStatisticalMgr) GetByOptions(opts ...Option) (results []*LgTrackGradStatistical, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgTrackGradStatisticalMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgTrackGradStatistical, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where(options.query)
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
func (obj *_LgTrackGradStatisticalMgr) GetFromID(id int64) (result LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LgTrackGradStatisticalMgr) GetBatchFromID(ids []int64) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromShipmentCode 通过shipment_code获取内容 物流渠道代码
func (obj *_LgTrackGradStatisticalMgr) GetFromShipmentCode(shipmentCode string) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`shipment_code` = ?", shipmentCode).Find(&results).Error

	return
}

// GetBatchFromShipmentCode 批量查找 物流渠道代码
func (obj *_LgTrackGradStatisticalMgr) GetBatchFromShipmentCode(shipmentCodes []string) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`shipment_code` IN (?)", shipmentCodes).Find(&results).Error

	return
}

// GetFromStateCode 通过state_code获取内容 轨迹状态
func (obj *_LgTrackGradStatisticalMgr) GetFromStateCode(stateCode string) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`state_code` = ?", stateCode).Find(&results).Error

	return
}

// GetBatchFromStateCode 批量查找 轨迹状态
func (obj *_LgTrackGradStatisticalMgr) GetBatchFromStateCode(stateCodes []string) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`state_code` IN (?)", stateCodes).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容 抓取数量
func (obj *_LgTrackGradStatisticalMgr) GetFromCount(count int) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找 抓取数量
func (obj *_LgTrackGradStatisticalMgr) GetBatchFromCount(counts []int) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgTrackGradStatisticalMgr) GetFromCreateTime(createTime time.Time) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgTrackGradStatisticalMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgTrackGradStatisticalMgr) FetchByPrimaryKey(id int64) (result LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByLgTrackGradStatisticalShipmentCodeIndex  获取多个内容
func (obj *_LgTrackGradStatisticalMgr) FetchIndexByLgTrackGradStatisticalShipmentCodeIndex(shipmentCode string) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`shipment_code` = ?", shipmentCode).Find(&results).Error

	return
}

// FetchIndexByLgTrackGradStatisticalStateCodeIndex  获取多个内容
func (obj *_LgTrackGradStatisticalMgr) FetchIndexByLgTrackGradStatisticalStateCodeIndex(stateCode string) (results []*LgTrackGradStatistical, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgTrackGradStatistical{}).Where("`state_code` = ?", stateCode).Find(&results).Error

	return
}
