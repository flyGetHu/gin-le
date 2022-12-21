package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _LgOrderStatisticalCopy1Mgr struct {
	*_BaseMgr
}

// LgOrderStatisticalCopy1Mgr open func
func LgOrderStatisticalCopy1Mgr(db *gorm.DB) *_LgOrderStatisticalCopy1Mgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderStatisticalCopy1Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderStatisticalCopy1Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_statistical_copy1"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderStatisticalCopy1Mgr) GetTableName() string {
	return "lg_order_statistical_copy1"
}

// Reset 重置gorm会话
func (obj *_LgOrderStatisticalCopy1Mgr) Reset() *_LgOrderStatisticalCopy1Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderStatisticalCopy1Mgr) Get() (result LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderStatisticalCopy1Mgr) Gets() (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderStatisticalCopy1Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithType type获取 统计类型
func (obj *_LgOrderStatisticalCopy1Mgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithDate date获取 日期
func (obj *_LgOrderStatisticalCopy1Mgr) WithDate(date datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// WithDateType date_type获取 时间类型
func (obj *_LgOrderStatisticalCopy1Mgr) WithDateType(dateType int) Option {
	return optionFunc(func(o *options) { o.query["date_type"] = dateType })
}

// WithData data获取 订单统计数据
func (obj *_LgOrderStatisticalCopy1Mgr) WithData(data datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["data"] = data })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderStatisticalCopy1Mgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCustomerChannelID customer_channel_id获取 渠道id
func (obj *_LgOrderStatisticalCopy1Mgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithPlatformCode platform_code获取 平台代码
func (obj *_LgOrderStatisticalCopy1Mgr) WithPlatformCode(platformCode string) Option {
	return optionFunc(func(o *options) { o.query["platform_code"] = platformCode })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderStatisticalCopy1Mgr) GetByOption(opts ...Option) (result LgOrderStatisticalCopy1, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderStatisticalCopy1Mgr) GetByOptions(opts ...Option) (results []*LgOrderStatisticalCopy1, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderStatisticalCopy1Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderStatisticalCopy1, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where(options.query)
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
func (obj *_LgOrderStatisticalCopy1Mgr) GetFromType(_type int) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 统计类型
func (obj *_LgOrderStatisticalCopy1Mgr) GetBatchFromType(_types []int) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容 日期
func (obj *_LgOrderStatisticalCopy1Mgr) GetFromDate(date datatypes.Date) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找 日期
func (obj *_LgOrderStatisticalCopy1Mgr) GetBatchFromDate(dates []datatypes.Date) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

// GetFromDateType 通过date_type获取内容 时间类型
func (obj *_LgOrderStatisticalCopy1Mgr) GetFromDateType(dateType int) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`date_type` = ?", dateType).Find(&results).Error

	return
}

// GetBatchFromDateType 批量查找 时间类型
func (obj *_LgOrderStatisticalCopy1Mgr) GetBatchFromDateType(dateTypes []int) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`date_type` IN (?)", dateTypes).Find(&results).Error

	return
}

// GetFromData 通过data获取内容 订单统计数据
func (obj *_LgOrderStatisticalCopy1Mgr) GetFromData(data datatypes.JSON) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`data` = ?", data).Find(&results).Error

	return
}

// GetBatchFromData 批量查找 订单统计数据
func (obj *_LgOrderStatisticalCopy1Mgr) GetBatchFromData(datas []datatypes.JSON) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`data` IN (?)", datas).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderStatisticalCopy1Mgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderStatisticalCopy1Mgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 渠道id
func (obj *_LgOrderStatisticalCopy1Mgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 渠道id
func (obj *_LgOrderStatisticalCopy1Mgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromPlatformCode 通过platform_code获取内容 平台代码
func (obj *_LgOrderStatisticalCopy1Mgr) GetFromPlatformCode(platformCode string) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`platform_code` = ?", platformCode).Find(&results).Error

	return
}

// GetBatchFromPlatformCode 批量查找 平台代码
func (obj *_LgOrderStatisticalCopy1Mgr) GetBatchFromPlatformCode(platformCodes []string) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`platform_code` IN (?)", platformCodes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByLgOrderStatisticalDateIndex  获取多个内容
func (obj *_LgOrderStatisticalCopy1Mgr) FetchIndexByLgOrderStatisticalDateIndex(date datatypes.Date) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// FetchIndexByLgOrderStatisticalCreateTimeIndex  获取多个内容
func (obj *_LgOrderStatisticalCopy1Mgr) FetchIndexByLgOrderStatisticalCreateTimeIndex(createTime time.Time) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderStatisticalCustomerChannelIDIndex  获取多个内容
func (obj *_LgOrderStatisticalCopy1Mgr) FetchIndexByLgOrderStatisticalCustomerChannelIDIndex(customerChannelID int) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByLgOrderStatisticalPlatformCodeIndex  获取多个内容
func (obj *_LgOrderStatisticalCopy1Mgr) FetchIndexByLgOrderStatisticalPlatformCodeIndex(platformCode string) (results []*LgOrderStatisticalCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderStatisticalCopy1{}).Where("`platform_code` = ?", platformCode).Find(&results).Error

	return
}
