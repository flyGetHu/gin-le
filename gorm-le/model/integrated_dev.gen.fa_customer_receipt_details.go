package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaCustomerReceiptDetailsMgr struct {
	*_BaseMgr
}

// FaCustomerReceiptDetailsMgr open func
func FaCustomerReceiptDetailsMgr(db *gorm.DB) *_FaCustomerReceiptDetailsMgr {
	if db == nil {
		panic(fmt.Errorf("FaCustomerReceiptDetailsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaCustomerReceiptDetailsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_customer_receipt_details"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaCustomerReceiptDetailsMgr) GetTableName() string {
	return "fa_customer_receipt_details"
}

// Reset 重置gorm会话
func (obj *_FaCustomerReceiptDetailsMgr) Reset() *_FaCustomerReceiptDetailsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaCustomerReceiptDetailsMgr) Get() (result FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaCustomerReceiptDetailsMgr) Gets() (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaCustomerReceiptDetailsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaCustomerReceiptDetailsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithReceiptNumber receipt_number获取 收款单号
func (obj *_FaCustomerReceiptDetailsMgr) WithReceiptNumber(receiptNumber string) Option {
	return optionFunc(func(o *options) { o.query["receipt_number"] = receiptNumber })
}

// WithOrderNumber order_number获取 订单号
func (obj *_FaCustomerReceiptDetailsMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_FaCustomerReceiptDetailsMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_FaCustomerReceiptDetailsMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_FaCustomerReceiptDetailsMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_FaCustomerReceiptDetailsMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_FaCustomerReceiptDetailsMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_FaCustomerReceiptDetailsMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_FaCustomerReceiptDetailsMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithIsAccept is_accept获取 已核收(关联收款单状态)
func (obj *_FaCustomerReceiptDetailsMgr) WithIsAccept(isAccept []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_accept"] = isAccept })
}

// WithAcceptTime accept_time获取 核收时间
func (obj *_FaCustomerReceiptDetailsMgr) WithAcceptTime(acceptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["accept_time"] = acceptTime })
}

// WithAcceptUser accept_user获取 核收用户
func (obj *_FaCustomerReceiptDetailsMgr) WithAcceptUser(acceptUser string) Option {
	return optionFunc(func(o *options) { o.query["accept_user"] = acceptUser })
}

// WithAcceptUserID accept_user_id获取 核收用户ID
func (obj *_FaCustomerReceiptDetailsMgr) WithAcceptUserID(acceptUserID int) Option {
	return optionFunc(func(o *options) { o.query["accept_user_id"] = acceptUserID })
}

// WithWeight weight获取 计费总重量
func (obj *_FaCustomerReceiptDetailsMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithReceivablesRmb receivables_rmb获取 应收总费用(RMB)
func (obj *_FaCustomerReceiptDetailsMgr) WithReceivablesRmb(receivablesRmb float64) Option {
	return optionFunc(func(o *options) { o.query["receivables_rmb"] = receivablesRmb })
}

// WithBeforeBalance before_balance获取 扣款前余额
func (obj *_FaCustomerReceiptDetailsMgr) WithBeforeBalance(beforeBalance float64) Option {
	return optionFunc(func(o *options) { o.query["before_balance"] = beforeBalance })
}

// WithAfterBalance after_balance获取 扣款后余额
func (obj *_FaCustomerReceiptDetailsMgr) WithAfterBalance(afterBalance float64) Option {
	return optionFunc(func(o *options) { o.query["after_balance"] = afterBalance })
}

// WithRemark remark获取 备注
func (obj *_FaCustomerReceiptDetailsMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaCustomerReceiptDetailsMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应收人员
func (obj *_FaCustomerReceiptDetailsMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaCustomerReceiptDetailsMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaCustomerReceiptDetailsMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaCustomerReceiptDetailsMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaCustomerReceiptDetailsMgr) GetByOption(opts ...Option) (result FaCustomerReceiptDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaCustomerReceiptDetailsMgr) GetByOptions(opts ...Option) (results []*FaCustomerReceiptDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaCustomerReceiptDetailsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaCustomerReceiptDetails, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where(options.query)
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
func (obj *_FaCustomerReceiptDetailsMgr) GetFromID(id int) (result FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromID(ids []int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromReceiptNumber 通过receipt_number获取内容 收款单号
func (obj *_FaCustomerReceiptDetailsMgr) GetFromReceiptNumber(receiptNumber string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`receipt_number` = ?", receiptNumber).Find(&results).Error

	return
}

// GetBatchFromReceiptNumber 批量查找 收款单号
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromReceiptNumber(receiptNumbers []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`receipt_number` IN (?)", receiptNumbers).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_FaCustomerReceiptDetailsMgr) GetFromOrderNumber(orderNumber string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_FaCustomerReceiptDetailsMgr) GetFromReferenceNumber(referenceNumber string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_FaCustomerReceiptDetailsMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_FaCustomerReceiptDetailsMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_FaCustomerReceiptDetailsMgr) GetFromCustomerID(customerID int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_FaCustomerReceiptDetailsMgr) GetFromCustomerAlias(customerAlias string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_FaCustomerReceiptDetailsMgr) GetFromCustomerChannelID(customerChannelID int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_FaCustomerReceiptDetailsMgr) GetFromCustomerChannelName(customerChannelName string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromIsAccept 通过is_accept获取内容 已核收(关联收款单状态)
func (obj *_FaCustomerReceiptDetailsMgr) GetFromIsAccept(isAccept []uint8) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`is_accept` = ?", isAccept).Find(&results).Error

	return
}

// GetBatchFromIsAccept 批量查找 已核收(关联收款单状态)
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromIsAccept(isAccepts [][]uint8) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`is_accept` IN (?)", isAccepts).Find(&results).Error

	return
}

// GetFromAcceptTime 通过accept_time获取内容 核收时间
func (obj *_FaCustomerReceiptDetailsMgr) GetFromAcceptTime(acceptTime time.Time) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// GetBatchFromAcceptTime 批量查找 核收时间
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromAcceptTime(acceptTimes []time.Time) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`accept_time` IN (?)", acceptTimes).Find(&results).Error

	return
}

// GetFromAcceptUser 通过accept_user获取内容 核收用户
func (obj *_FaCustomerReceiptDetailsMgr) GetFromAcceptUser(acceptUser string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`accept_user` = ?", acceptUser).Find(&results).Error

	return
}

// GetBatchFromAcceptUser 批量查找 核收用户
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromAcceptUser(acceptUsers []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`accept_user` IN (?)", acceptUsers).Find(&results).Error

	return
}

// GetFromAcceptUserID 通过accept_user_id获取内容 核收用户ID
func (obj *_FaCustomerReceiptDetailsMgr) GetFromAcceptUserID(acceptUserID int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`accept_user_id` = ?", acceptUserID).Find(&results).Error

	return
}

// GetBatchFromAcceptUserID 批量查找 核收用户ID
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromAcceptUserID(acceptUserIDs []int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`accept_user_id` IN (?)", acceptUserIDs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 计费总重量
func (obj *_FaCustomerReceiptDetailsMgr) GetFromWeight(weight float64) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 计费总重量
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromWeight(weights []float64) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromReceivablesRmb 通过receivables_rmb获取内容 应收总费用(RMB)
func (obj *_FaCustomerReceiptDetailsMgr) GetFromReceivablesRmb(receivablesRmb float64) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`receivables_rmb` = ?", receivablesRmb).Find(&results).Error

	return
}

// GetBatchFromReceivablesRmb 批量查找 应收总费用(RMB)
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromReceivablesRmb(receivablesRmbs []float64) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`receivables_rmb` IN (?)", receivablesRmbs).Find(&results).Error

	return
}

// GetFromBeforeBalance 通过before_balance获取内容 扣款前余额
func (obj *_FaCustomerReceiptDetailsMgr) GetFromBeforeBalance(beforeBalance float64) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`before_balance` = ?", beforeBalance).Find(&results).Error

	return
}

// GetBatchFromBeforeBalance 批量查找 扣款前余额
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromBeforeBalance(beforeBalances []float64) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`before_balance` IN (?)", beforeBalances).Find(&results).Error

	return
}

// GetFromAfterBalance 通过after_balance获取内容 扣款后余额
func (obj *_FaCustomerReceiptDetailsMgr) GetFromAfterBalance(afterBalance float64) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`after_balance` = ?", afterBalance).Find(&results).Error

	return
}

// GetBatchFromAfterBalance 批量查找 扣款后余额
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromAfterBalance(afterBalances []float64) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`after_balance` IN (?)", afterBalances).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaCustomerReceiptDetailsMgr) GetFromRemark(remark string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromRemark(remarks []string) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaCustomerReceiptDetailsMgr) GetFromCreateTime(createTime time.Time) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应收人员
func (obj *_FaCustomerReceiptDetailsMgr) GetFromCreateUser(createUser int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应收人员
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaCustomerReceiptDetailsMgr) GetFromUpdateUser(updateUser int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaCustomerReceiptDetailsMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaCustomerReceiptDetailsMgr) GetFromVersion(version int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaCustomerReceiptDetailsMgr) GetBatchFromVersion(versions []int) (results []*FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaCustomerReceiptDetailsMgr) FetchByPrimaryKey(id int) (result FaCustomerReceiptDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceiptDetails{}).Where("`id` = ?", id).First(&result).Error

	return
}
