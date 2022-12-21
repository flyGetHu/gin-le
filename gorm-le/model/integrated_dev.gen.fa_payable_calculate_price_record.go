package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaPayableCalculatePriceRecordMgr struct {
	*_BaseMgr
}

// FaPayableCalculatePriceRecordMgr open func
func FaPayableCalculatePriceRecordMgr(db *gorm.DB) *_FaPayableCalculatePriceRecordMgr {
	if db == nil {
		panic(fmt.Errorf("FaPayableCalculatePriceRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaPayableCalculatePriceRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_payable_calculate_price_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaPayableCalculatePriceRecordMgr) GetTableName() string {
	return "fa_payable_calculate_price_record"
}

// Reset 重置gorm会话
func (obj *_FaPayableCalculatePriceRecordMgr) Reset() *_FaPayableCalculatePriceRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaPayableCalculatePriceRecordMgr) Get() (result FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaPayableCalculatePriceRecordMgr) Gets() (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaPayableCalculatePriceRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaPayableCalculatePriceRecordMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_FaPayableCalculatePriceRecordMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_FaPayableCalculatePriceRecordMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_FaPayableCalculatePriceRecordMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_FaPayableCalculatePriceRecordMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithUpdateField update_field获取 修改字段
func (obj *_FaPayableCalculatePriceRecordMgr) WithUpdateField(updateField string) Option {
	return optionFunc(func(o *options) { o.query["update_field"] = updateField })
}

// WithOldValue old_value获取 操作前值
func (obj *_FaPayableCalculatePriceRecordMgr) WithOldValue(oldValue string) Option {
	return optionFunc(func(o *options) { o.query["old_value"] = oldValue })
}

// WithNewValue new_value获取 操作后值
func (obj *_FaPayableCalculatePriceRecordMgr) WithNewValue(newValue string) Option {
	return optionFunc(func(o *options) { o.query["new_value"] = newValue })
}

// WithCreateTime create_time获取 操作时间
func (obj *_FaPayableCalculatePriceRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 操作人
func (obj *_FaPayableCalculatePriceRecordMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaPayableCalculatePriceRecordMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaPayableCalculatePriceRecordMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaPayableCalculatePriceRecordMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaPayableCalculatePriceRecordMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaPayableCalculatePriceRecordMgr) GetByOption(opts ...Option) (result FaPayableCalculatePriceRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaPayableCalculatePriceRecordMgr) GetByOptions(opts ...Option) (results []*FaPayableCalculatePriceRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaPayableCalculatePriceRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaPayableCalculatePriceRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where(options.query)
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
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromID(id int64) (result FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromID(ids []int64) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromOrderNumber(orderNumber string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromReferenceNumber(referenceNumber string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromUpdateField 通过update_field获取内容 修改字段
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromUpdateField(updateField string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`update_field` = ?", updateField).Find(&results).Error

	return
}

// GetBatchFromUpdateField 批量查找 修改字段
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromUpdateField(updateFields []string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`update_field` IN (?)", updateFields).Find(&results).Error

	return
}

// GetFromOldValue 通过old_value获取内容 操作前值
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromOldValue(oldValue string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`old_value` = ?", oldValue).Find(&results).Error

	return
}

// GetBatchFromOldValue 批量查找 操作前值
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromOldValue(oldValues []string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`old_value` IN (?)", oldValues).Find(&results).Error

	return
}

// GetFromNewValue 通过new_value获取内容 操作后值
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromNewValue(newValue string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`new_value` = ?", newValue).Find(&results).Error

	return
}

// GetBatchFromNewValue 批量查找 操作后值
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromNewValue(newValues []string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`new_value` IN (?)", newValues).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 操作时间
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromCreateTime(createTime time.Time) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 操作时间
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 操作人
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromCreateUser(createUser int) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 操作人
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromUpdateUser(updateUser int) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromVersion(version int) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromVersion(versions []int) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaPayableCalculatePriceRecordMgr) GetFromDeleted(deleted int) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaPayableCalculatePriceRecordMgr) GetBatchFromDeleted(deleteds []int) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaPayableCalculatePriceRecordMgr) FetchByPrimaryKey(id int64) (result FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIndexOrderNumber  获取多个内容
func (obj *_FaPayableCalculatePriceRecordMgr) FetchIndexByIndexOrderNumber(orderNumber string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByIndexReferenceNumber  获取多个内容
func (obj *_FaPayableCalculatePriceRecordMgr) FetchIndexByIndexReferenceNumber(referenceNumber string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumber  获取多个内容
func (obj *_FaPayableCalculatePriceRecordMgr) FetchIndexByIndexLogisticsNumber(logisticsNumber string) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCreateTime  获取多个内容
func (obj *_FaPayableCalculatePriceRecordMgr) FetchIndexByIndexCreateTime(createTime time.Time) (results []*FaPayableCalculatePriceRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableCalculatePriceRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
