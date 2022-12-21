package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _WmsCn35DispatchNumRecordMgr struct {
	*_BaseMgr
}

// WmsCn35DispatchNumRecordMgr open func
func WmsCn35DispatchNumRecordMgr(db *gorm.DB) *_WmsCn35DispatchNumRecordMgr {
	if db == nil {
		panic(fmt.Errorf("WmsCn35DispatchNumRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsCn35DispatchNumRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_cn35_dispatch_num_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsCn35DispatchNumRecordMgr) GetTableName() string {
	return "wms_cn35_dispatch_num_record"
}

// Reset 重置gorm会话
func (obj *_WmsCn35DispatchNumRecordMgr) Reset() *_WmsCn35DispatchNumRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsCn35DispatchNumRecordMgr) Get() (result WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsCn35DispatchNumRecordMgr) Gets() (results []*WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsCn35DispatchNumRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsCn35DispatchNumRecordMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBigBagID big_bag_id获取 大包id
func (obj *_WmsCn35DispatchNumRecordMgr) WithBigBagID(bigBagID int) Option {
	return optionFunc(func(o *options) { o.query["big_bag_id"] = bigBagID })
}

// WithCn35DispatchNumber cn35_dispatch_number获取 cn35预报唯一标识(最大 999999)
func (obj *_WmsCn35DispatchNumRecordMgr) WithCn35DispatchNumber(cn35DispatchNumber float64) Option {
	return optionFunc(func(o *options) { o.query["cn35_dispatch_number"] = cn35DispatchNumber })
}

// GetByOption 功能选项模式获取
func (obj *_WmsCn35DispatchNumRecordMgr) GetByOption(opts ...Option) (result WmsCn35DispatchNumRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsCn35DispatchNumRecordMgr) GetByOptions(opts ...Option) (results []*WmsCn35DispatchNumRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsCn35DispatchNumRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsCn35DispatchNumRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where(options.query)
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
func (obj *_WmsCn35DispatchNumRecordMgr) GetFromID(id int) (result WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsCn35DispatchNumRecordMgr) GetBatchFromID(ids []int) (results []*WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBigBagID 通过big_bag_id获取内容 大包id
func (obj *_WmsCn35DispatchNumRecordMgr) GetFromBigBagID(bigBagID int) (results []*WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where("`big_bag_id` = ?", bigBagID).Find(&results).Error

	return
}

// GetBatchFromBigBagID 批量查找 大包id
func (obj *_WmsCn35DispatchNumRecordMgr) GetBatchFromBigBagID(bigBagIDs []int) (results []*WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where("`big_bag_id` IN (?)", bigBagIDs).Find(&results).Error

	return
}

// GetFromCn35DispatchNumber 通过cn35_dispatch_number获取内容 cn35预报唯一标识(最大 999999)
func (obj *_WmsCn35DispatchNumRecordMgr) GetFromCn35DispatchNumber(cn35DispatchNumber float64) (results []*WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where("`cn35_dispatch_number` = ?", cn35DispatchNumber).Find(&results).Error

	return
}

// GetBatchFromCn35DispatchNumber 批量查找 cn35预报唯一标识(最大 999999)
func (obj *_WmsCn35DispatchNumRecordMgr) GetBatchFromCn35DispatchNumber(cn35DispatchNumbers []float64) (results []*WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where("`cn35_dispatch_number` IN (?)", cn35DispatchNumbers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsCn35DispatchNumRecordMgr) FetchByPrimaryKey(id int) (result WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByWmsCn35NumRecordBigBagIDCn35DispatchNumberIndex primary or index 获取唯一内容
func (obj *_WmsCn35DispatchNumRecordMgr) FetchUniqueIndexByWmsCn35NumRecordBigBagIDCn35DispatchNumberIndex(bigBagID int, cn35DispatchNumber float64) (result WmsCn35DispatchNumRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsCn35DispatchNumRecord{}).Where("`big_bag_id` = ? AND `cn35_dispatch_number` = ?", bigBagID, cn35DispatchNumber).First(&result).Error

	return
}
