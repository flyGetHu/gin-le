package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _WWorkOrderReviewingMgr struct {
	*_BaseMgr
}

// WWorkOrderReviewingMgr open func
func WWorkOrderReviewingMgr(db *gorm.DB) *_WWorkOrderReviewingMgr {
	if db == nil {
		panic(fmt.Errorf("WWorkOrderReviewingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WWorkOrderReviewingMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_work_order_reviewing"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WWorkOrderReviewingMgr) GetTableName() string {
	return "w_work_order_reviewing"
}

// Reset 重置gorm会话
func (obj *_WWorkOrderReviewingMgr) Reset() *_WWorkOrderReviewingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WWorkOrderReviewingMgr) Get() (result WWorkOrderReviewing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WWorkOrderReviewingMgr) Gets() (results []*WWorkOrderReviewing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WWorkOrderReviewingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOrderID order_id获取
func (obj *_WWorkOrderReviewingMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithReviewerID reviewer_id获取
func (obj *_WWorkOrderReviewingMgr) WithReviewerID(reviewerID int) Option {
	return optionFunc(func(o *options) { o.query["reviewer_id"] = reviewerID })
}

// GetByOption 功能选项模式获取
func (obj *_WWorkOrderReviewingMgr) GetByOption(opts ...Option) (result WWorkOrderReviewing, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WWorkOrderReviewingMgr) GetByOptions(opts ...Option) (results []*WWorkOrderReviewing, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WWorkOrderReviewingMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WWorkOrderReviewing, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Where(options.query)
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

// GetFromOrderID 通过order_id获取内容
func (obj *_WWorkOrderReviewingMgr) GetFromOrderID(orderID int) (results []*WWorkOrderReviewing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找
func (obj *_WWorkOrderReviewingMgr) GetBatchFromOrderID(orderIDs []int) (results []*WWorkOrderReviewing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromReviewerID 通过reviewer_id获取内容
func (obj *_WWorkOrderReviewingMgr) GetFromReviewerID(reviewerID int) (results []*WWorkOrderReviewing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Where("`reviewer_id` = ?", reviewerID).Find(&results).Error

	return
}

// GetBatchFromReviewerID 批量查找
func (obj *_WWorkOrderReviewingMgr) GetBatchFromReviewerID(reviewerIDs []int) (results []*WWorkOrderReviewing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderReviewing{}).Where("`reviewer_id` IN (?)", reviewerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
