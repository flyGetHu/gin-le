package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaReCalculatePriceRecordMgr struct {
	*_BaseMgr
}

// FaReCalculatePriceRecordMgr open func
func FaReCalculatePriceRecordMgr(db *gorm.DB) *_FaReCalculatePriceRecordMgr {
	if db == nil {
		panic(fmt.Errorf("FaReCalculatePriceRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaReCalculatePriceRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_re_calculate_price_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaReCalculatePriceRecordMgr) GetTableName() string {
	return "fa_re_calculate_price_record"
}

// Reset 重置gorm会话
func (obj *_FaReCalculatePriceRecordMgr) Reset() *_FaReCalculatePriceRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaReCalculatePriceRecordMgr) Get() (result FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaReCalculatePriceRecordMgr) Gets() (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaReCalculatePriceRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaReCalculatePriceRecordMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_FaReCalculatePriceRecordMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_FaReCalculatePriceRecordMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_FaReCalculatePriceRecordMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_FaReCalculatePriceRecordMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithUpdateField update_field获取 修改字段
func (obj *_FaReCalculatePriceRecordMgr) WithUpdateField(updateField string) Option {
	return optionFunc(func(o *options) { o.query["update_field"] = updateField })
}

// WithOldValue old_value获取 操作前值
func (obj *_FaReCalculatePriceRecordMgr) WithOldValue(oldValue string) Option {
	return optionFunc(func(o *options) { o.query["old_value"] = oldValue })
}

// WithNewValue new_value获取 操作后值
func (obj *_FaReCalculatePriceRecordMgr) WithNewValue(newValue string) Option {
	return optionFunc(func(o *options) { o.query["new_value"] = newValue })
}

// WithCreateTime create_time获取 操作时间
func (obj *_FaReCalculatePriceRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 操作人
func (obj *_FaReCalculatePriceRecordMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaReCalculatePriceRecordMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaReCalculatePriceRecordMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaReCalculatePriceRecordMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaReCalculatePriceRecordMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaReCalculatePriceRecordMgr) GetByOption(opts ...Option) (result FaReCalculatePriceRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaReCalculatePriceRecordMgr) GetByOptions(opts ...Option) (results []*FaReCalculatePriceRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaReCalculatePriceRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaReCalculatePriceRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_FaReCalculatePriceRecordMgr) GetFromID(id int) (result FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromID(ids []int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_FaReCalculatePriceRecordMgr) GetFromOrderNumber(orderNumber string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_FaReCalculatePriceRecordMgr) GetFromReferenceNumber(referenceNumber string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_FaReCalculatePriceRecordMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_FaReCalculatePriceRecordMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromUpdateField 通过update_field获取内容 修改字段
func (obj *_FaReCalculatePriceRecordMgr) GetFromUpdateField(updateField string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`update_field` = ?", updateField).Find(&results).Error

	return
}

// GetBatchFromUpdateField 批量查找 修改字段
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromUpdateField(updateFields []string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`update_field` IN (?)", updateFields).Find(&results).Error

	return
}

// GetFromOldValue 通过old_value获取内容 操作前值
func (obj *_FaReCalculatePriceRecordMgr) GetFromOldValue(oldValue string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`old_value` = ?", oldValue).Find(&results).Error

	return
}

// GetBatchFromOldValue 批量查找 操作前值
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromOldValue(oldValues []string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`old_value` IN (?)", oldValues).Find(&results).Error

	return
}

// GetFromNewValue 通过new_value获取内容 操作后值
func (obj *_FaReCalculatePriceRecordMgr) GetFromNewValue(newValue string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`new_value` = ?", newValue).Find(&results).Error

	return
}

// GetBatchFromNewValue 批量查找 操作后值
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromNewValue(newValues []string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`new_value` IN (?)", newValues).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 操作时间
func (obj *_FaReCalculatePriceRecordMgr) GetFromCreateTime(createTime time.Time) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 操作时间
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 操作人
func (obj *_FaReCalculatePriceRecordMgr) GetFromCreateUser(createUser int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 操作人
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaReCalculatePriceRecordMgr) GetFromUpdateUser(updateUser int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaReCalculatePriceRecordMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaReCalculatePriceRecordMgr) GetFromVersion(version int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromVersion(versions []int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaReCalculatePriceRecordMgr) GetFromDeleted(deleted int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaReCalculatePriceRecordMgr) GetBatchFromDeleted(deleteds []int) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaReCalculatePriceRecordMgr) FetchByPrimaryKey(id int) (result FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIndexReferenceNumber  获取多个内容
func (obj *_FaReCalculatePriceRecordMgr) FetchIndexByIndexReferenceNumber(referenceNumber string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumber  获取多个内容
func (obj *_FaReCalculatePriceRecordMgr) FetchIndexByIndexLogisticsNumber(logisticsNumber string) (results []*FaReCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaReCalculatePriceRecord{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}
