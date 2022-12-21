package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaCustomerBillMgr struct {
	*_BaseMgr
}

// FaCustomerBillMgr open func
func FaCustomerBillMgr(db *gorm.DB) *_FaCustomerBillMgr {
	if db == nil {
		panic(fmt.Errorf("FaCustomerBillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaCustomerBillMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_customer_bill"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaCustomerBillMgr) GetTableName() string {
	return "fa_customer_bill"
}

// Reset 重置gorm会话
func (obj *_FaCustomerBillMgr) Reset() *_FaCustomerBillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaCustomerBillMgr) Get() (result FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaCustomerBillMgr) Gets() (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaCustomerBillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaCustomerBillMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBillBatchNumber bill_batch_number获取 账单批次号
func (obj *_FaCustomerBillMgr) WithBillBatchNumber(billBatchNumber string) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number"] = billBatchNumber })
}

// WithBillBatchStatus bill_batch_status获取 账单状态
func (obj *_FaCustomerBillMgr) WithBillBatchStatus(billBatchStatus int) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_status"] = billBatchStatus })
}

// WithReceiptNumber receipt_number获取 收款单号
func (obj *_FaCustomerBillMgr) WithReceiptNumber(receiptNumber string) Option {
	return optionFunc(func(o *options) { o.query["receipt_number"] = receiptNumber })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_FaCustomerBillMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_FaCustomerBillMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithCustomerChannelNames customer_channel_names获取 客户渠道名称
func (obj *_FaCustomerBillMgr) WithCustomerChannelNames(customerChannelNames string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_names"] = customerChannelNames })
}

// WithOrderQuantity order_quantity获取 订单数量
func (obj *_FaCustomerBillMgr) WithOrderQuantity(orderQuantity int) Option {
	return optionFunc(func(o *options) { o.query["order_quantity"] = orderQuantity })
}

// WithFeeQuantity fee_quantity获取 费用条数
func (obj *_FaCustomerBillMgr) WithFeeQuantity(feeQuantity int) Option {
	return optionFunc(func(o *options) { o.query["fee_quantity"] = feeQuantity })
}

// WithWeight weight获取 计费总重量
func (obj *_FaCustomerBillMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithReceivablesRmb receivables_rmb获取 应收总费用(RMB)
func (obj *_FaCustomerBillMgr) WithReceivablesRmb(receivablesRmb float64) Option {
	return optionFunc(func(o *options) { o.query["receivables_rmb"] = receivablesRmb })
}

// WithActualFeeRmb actual_fee_rmb获取 实际收款金额rmb
func (obj *_FaCustomerBillMgr) WithActualFeeRmb(actualFeeRmb float64) Option {
	return optionFunc(func(o *options) { o.query["actual_fee_rmb"] = actualFeeRmb })
}

// WithCreateTime create_time获取 生成时间
func (obj *_FaCustomerBillMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应收人员
func (obj *_FaCustomerBillMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithIsAudit is_audit获取 审核状态(关联对账单批次号状态)
func (obj *_FaCustomerBillMgr) WithIsAudit(isAudit []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_audit"] = isAudit })
}

// WithIsAuditRollback is_audit_rollback获取 是否反审标识 默认false ,被反审后true,重新审核后false
func (obj *_FaCustomerBillMgr) WithIsAuditRollback(isAuditRollback []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_audit_rollback"] = isAuditRollback })
}

// WithAuditRollbackUser audit_rollback_user获取 反审核人员
func (obj *_FaCustomerBillMgr) WithAuditRollbackUser(auditRollbackUser string) Option {
	return optionFunc(func(o *options) { o.query["audit_rollback_user"] = auditRollbackUser })
}

// WithAuditRollbackUserID audit_rollback_user_id获取 反审核人员ID
func (obj *_FaCustomerBillMgr) WithAuditRollbackUserID(auditRollbackUserID int) Option {
	return optionFunc(func(o *options) { o.query["audit_rollback_user_id"] = auditRollbackUserID })
}

// WithCancelUser cancel_user获取 取消应收用户
func (obj *_FaCustomerBillMgr) WithCancelUser(cancelUser string) Option {
	return optionFunc(func(o *options) { o.query["cancel_user"] = cancelUser })
}

// WithCancelUserID cancel_user_id获取 取消应收用户ID
func (obj *_FaCustomerBillMgr) WithCancelUserID(cancelUserID int) Option {
	return optionFunc(func(o *options) { o.query["cancel_user_id"] = cancelUserID })
}

// WithAuditUser audit_user获取 审核用户
func (obj *_FaCustomerBillMgr) WithAuditUser(auditUser string) Option {
	return optionFunc(func(o *options) { o.query["audit_user"] = auditUser })
}

// WithAuditUserID audit_user_id获取 审核用户
func (obj *_FaCustomerBillMgr) WithAuditUserID(auditUserID int) Option {
	return optionFunc(func(o *options) { o.query["audit_user_id"] = auditUserID })
}

// WithAuditTime audit_time获取 审核时间
func (obj *_FaCustomerBillMgr) WithAuditTime(auditTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["audit_time"] = auditTime })
}

// WithIsReceipt is_receipt获取 已生成收款单
func (obj *_FaCustomerBillMgr) WithIsReceipt(isReceipt []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_receipt"] = isReceipt })
}

// WithReceiptNumberUser receipt_number_user获取 收款单生成用户
func (obj *_FaCustomerBillMgr) WithReceiptNumberUser(receiptNumberUser string) Option {
	return optionFunc(func(o *options) { o.query["receipt_number_user"] = receiptNumberUser })
}

// WithReceiptNumberUserID receipt_number_user_id获取 收款单生成用户ID
func (obj *_FaCustomerBillMgr) WithReceiptNumberUserID(receiptNumberUserID int) Option {
	return optionFunc(func(o *options) { o.query["receipt_number_user_id"] = receiptNumberUserID })
}

// WithReceiptNumberTime receipt_number_time获取 收款单生成用户生成时间
func (obj *_FaCustomerBillMgr) WithReceiptNumberTime(receiptNumberTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_number_time"] = receiptNumberTime })
}

// WithIsAccept is_accept获取 已核收(关联收款单状态)
func (obj *_FaCustomerBillMgr) WithIsAccept(isAccept []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_accept"] = isAccept })
}

// WithAcceptTime accept_time获取 核收时间
func (obj *_FaCustomerBillMgr) WithAcceptTime(acceptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["accept_time"] = acceptTime })
}

// WithAcceptUser accept_user获取 核收用户
func (obj *_FaCustomerBillMgr) WithAcceptUser(acceptUser string) Option {
	return optionFunc(func(o *options) { o.query["accept_user"] = acceptUser })
}

// WithAcceptUserID accept_user_id获取 核收用户ID
func (obj *_FaCustomerBillMgr) WithAcceptUserID(acceptUserID int) Option {
	return optionFunc(func(o *options) { o.query["accept_user_id"] = acceptUserID })
}

// WithRemark remark获取 备注
func (obj *_FaCustomerBillMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithFinancialRemark financial_remark获取 财务备注
func (obj *_FaCustomerBillMgr) WithFinancialRemark(financialRemark string) Option {
	return optionFunc(func(o *options) { o.query["financial_remark"] = financialRemark })
}

// WithAcceptRemark accept_remark获取 核收备注
func (obj *_FaCustomerBillMgr) WithAcceptRemark(acceptRemark string) Option {
	return optionFunc(func(o *options) { o.query["accept_remark"] = acceptRemark })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaCustomerBillMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaCustomerBillMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaCustomerBillMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaCustomerBillMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithChannelHaul channel_haul获取 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_FaCustomerBillMgr) WithChannelHaul(channelHaul string) Option {
	return optionFunc(func(o *options) { o.query["channel_haul"] = channelHaul })
}

// WithKingdeeID kingdee_id获取 金蝶客户内码
func (obj *_FaCustomerBillMgr) WithKingdeeID(kingdeeID string) Option {
	return optionFunc(func(o *options) { o.query["kingdee_id"] = kingdeeID })
}

// WithLastSyncTime last_sync_time获取 最后同步时间
func (obj *_FaCustomerBillMgr) WithLastSyncTime(lastSyncTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_sync_time"] = lastSyncTime })
}

// WithLastSyncUser last_sync_user获取 最后同步人
func (obj *_FaCustomerBillMgr) WithLastSyncUser(lastSyncUser int) Option {
	return optionFunc(func(o *options) { o.query["last_sync_user"] = lastSyncUser })
}

// GetByOption 功能选项模式获取
func (obj *_FaCustomerBillMgr) GetByOption(opts ...Option) (result FaCustomerBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaCustomerBillMgr) GetByOptions(opts ...Option) (results []*FaCustomerBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaCustomerBillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaCustomerBill, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where(options.query)
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
func (obj *_FaCustomerBillMgr) GetFromID(id int) (result FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaCustomerBillMgr) GetBatchFromID(ids []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBillBatchNumber 通过bill_batch_number获取内容 账单批次号
func (obj *_FaCustomerBillMgr) GetFromBillBatchNumber(billBatchNumber string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumber 批量查找 账单批次号
func (obj *_FaCustomerBillMgr) GetBatchFromBillBatchNumber(billBatchNumbers []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`bill_batch_number` IN (?)", billBatchNumbers).Find(&results).Error

	return
}

// GetFromBillBatchStatus 通过bill_batch_status获取内容 账单状态
func (obj *_FaCustomerBillMgr) GetFromBillBatchStatus(billBatchStatus int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`bill_batch_status` = ?", billBatchStatus).Find(&results).Error

	return
}

// GetBatchFromBillBatchStatus 批量查找 账单状态
func (obj *_FaCustomerBillMgr) GetBatchFromBillBatchStatus(billBatchStatuss []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`bill_batch_status` IN (?)", billBatchStatuss).Find(&results).Error

	return
}

// GetFromReceiptNumber 通过receipt_number获取内容 收款单号
func (obj *_FaCustomerBillMgr) GetFromReceiptNumber(receiptNumber string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receipt_number` = ?", receiptNumber).Find(&results).Error

	return
}

// GetBatchFromReceiptNumber 批量查找 收款单号
func (obj *_FaCustomerBillMgr) GetBatchFromReceiptNumber(receiptNumbers []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receipt_number` IN (?)", receiptNumbers).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_FaCustomerBillMgr) GetFromCustomerID(customerID int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_FaCustomerBillMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_FaCustomerBillMgr) GetFromCustomerAlias(customerAlias string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_FaCustomerBillMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromCustomerChannelNames 通过customer_channel_names获取内容 客户渠道名称
func (obj *_FaCustomerBillMgr) GetFromCustomerChannelNames(customerChannelNames string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`customer_channel_names` = ?", customerChannelNames).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelNames 批量查找 客户渠道名称
func (obj *_FaCustomerBillMgr) GetBatchFromCustomerChannelNames(customerChannelNamess []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`customer_channel_names` IN (?)", customerChannelNamess).Find(&results).Error

	return
}

// GetFromOrderQuantity 通过order_quantity获取内容 订单数量
func (obj *_FaCustomerBillMgr) GetFromOrderQuantity(orderQuantity int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`order_quantity` = ?", orderQuantity).Find(&results).Error

	return
}

// GetBatchFromOrderQuantity 批量查找 订单数量
func (obj *_FaCustomerBillMgr) GetBatchFromOrderQuantity(orderQuantitys []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`order_quantity` IN (?)", orderQuantitys).Find(&results).Error

	return
}

// GetFromFeeQuantity 通过fee_quantity获取内容 费用条数
func (obj *_FaCustomerBillMgr) GetFromFeeQuantity(feeQuantity int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`fee_quantity` = ?", feeQuantity).Find(&results).Error

	return
}

// GetBatchFromFeeQuantity 批量查找 费用条数
func (obj *_FaCustomerBillMgr) GetBatchFromFeeQuantity(feeQuantitys []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`fee_quantity` IN (?)", feeQuantitys).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 计费总重量
func (obj *_FaCustomerBillMgr) GetFromWeight(weight float64) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 计费总重量
func (obj *_FaCustomerBillMgr) GetBatchFromWeight(weights []float64) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromReceivablesRmb 通过receivables_rmb获取内容 应收总费用(RMB)
func (obj *_FaCustomerBillMgr) GetFromReceivablesRmb(receivablesRmb float64) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receivables_rmb` = ?", receivablesRmb).Find(&results).Error

	return
}

// GetBatchFromReceivablesRmb 批量查找 应收总费用(RMB)
func (obj *_FaCustomerBillMgr) GetBatchFromReceivablesRmb(receivablesRmbs []float64) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receivables_rmb` IN (?)", receivablesRmbs).Find(&results).Error

	return
}

// GetFromActualFeeRmb 通过actual_fee_rmb获取内容 实际收款金额rmb
func (obj *_FaCustomerBillMgr) GetFromActualFeeRmb(actualFeeRmb float64) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`actual_fee_rmb` = ?", actualFeeRmb).Find(&results).Error

	return
}

// GetBatchFromActualFeeRmb 批量查找 实际收款金额rmb
func (obj *_FaCustomerBillMgr) GetBatchFromActualFeeRmb(actualFeeRmbs []float64) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`actual_fee_rmb` IN (?)", actualFeeRmbs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 生成时间
func (obj *_FaCustomerBillMgr) GetFromCreateTime(createTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 生成时间
func (obj *_FaCustomerBillMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应收人员
func (obj *_FaCustomerBillMgr) GetFromCreateUser(createUser int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应收人员
func (obj *_FaCustomerBillMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromIsAudit 通过is_audit获取内容 审核状态(关联对账单批次号状态)
func (obj *_FaCustomerBillMgr) GetFromIsAudit(isAudit []uint8) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`is_audit` = ?", isAudit).Find(&results).Error

	return
}

// GetBatchFromIsAudit 批量查找 审核状态(关联对账单批次号状态)
func (obj *_FaCustomerBillMgr) GetBatchFromIsAudit(isAudits [][]uint8) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`is_audit` IN (?)", isAudits).Find(&results).Error

	return
}

// GetFromIsAuditRollback 通过is_audit_rollback获取内容 是否反审标识 默认false ,被反审后true,重新审核后false
func (obj *_FaCustomerBillMgr) GetFromIsAuditRollback(isAuditRollback []uint8) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`is_audit_rollback` = ?", isAuditRollback).Find(&results).Error

	return
}

// GetBatchFromIsAuditRollback 批量查找 是否反审标识 默认false ,被反审后true,重新审核后false
func (obj *_FaCustomerBillMgr) GetBatchFromIsAuditRollback(isAuditRollbacks [][]uint8) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`is_audit_rollback` IN (?)", isAuditRollbacks).Find(&results).Error

	return
}

// GetFromAuditRollbackUser 通过audit_rollback_user获取内容 反审核人员
func (obj *_FaCustomerBillMgr) GetFromAuditRollbackUser(auditRollbackUser string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_rollback_user` = ?", auditRollbackUser).Find(&results).Error

	return
}

// GetBatchFromAuditRollbackUser 批量查找 反审核人员
func (obj *_FaCustomerBillMgr) GetBatchFromAuditRollbackUser(auditRollbackUsers []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_rollback_user` IN (?)", auditRollbackUsers).Find(&results).Error

	return
}

// GetFromAuditRollbackUserID 通过audit_rollback_user_id获取内容 反审核人员ID
func (obj *_FaCustomerBillMgr) GetFromAuditRollbackUserID(auditRollbackUserID int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_rollback_user_id` = ?", auditRollbackUserID).Find(&results).Error

	return
}

// GetBatchFromAuditRollbackUserID 批量查找 反审核人员ID
func (obj *_FaCustomerBillMgr) GetBatchFromAuditRollbackUserID(auditRollbackUserIDs []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_rollback_user_id` IN (?)", auditRollbackUserIDs).Find(&results).Error

	return
}

// GetFromCancelUser 通过cancel_user获取内容 取消应收用户
func (obj *_FaCustomerBillMgr) GetFromCancelUser(cancelUser string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`cancel_user` = ?", cancelUser).Find(&results).Error

	return
}

// GetBatchFromCancelUser 批量查找 取消应收用户
func (obj *_FaCustomerBillMgr) GetBatchFromCancelUser(cancelUsers []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`cancel_user` IN (?)", cancelUsers).Find(&results).Error

	return
}

// GetFromCancelUserID 通过cancel_user_id获取内容 取消应收用户ID
func (obj *_FaCustomerBillMgr) GetFromCancelUserID(cancelUserID int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`cancel_user_id` = ?", cancelUserID).Find(&results).Error

	return
}

// GetBatchFromCancelUserID 批量查找 取消应收用户ID
func (obj *_FaCustomerBillMgr) GetBatchFromCancelUserID(cancelUserIDs []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`cancel_user_id` IN (?)", cancelUserIDs).Find(&results).Error

	return
}

// GetFromAuditUser 通过audit_user获取内容 审核用户
func (obj *_FaCustomerBillMgr) GetFromAuditUser(auditUser string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_user` = ?", auditUser).Find(&results).Error

	return
}

// GetBatchFromAuditUser 批量查找 审核用户
func (obj *_FaCustomerBillMgr) GetBatchFromAuditUser(auditUsers []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_user` IN (?)", auditUsers).Find(&results).Error

	return
}

// GetFromAuditUserID 通过audit_user_id获取内容 审核用户
func (obj *_FaCustomerBillMgr) GetFromAuditUserID(auditUserID int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_user_id` = ?", auditUserID).Find(&results).Error

	return
}

// GetBatchFromAuditUserID 批量查找 审核用户
func (obj *_FaCustomerBillMgr) GetBatchFromAuditUserID(auditUserIDs []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_user_id` IN (?)", auditUserIDs).Find(&results).Error

	return
}

// GetFromAuditTime 通过audit_time获取内容 审核时间
func (obj *_FaCustomerBillMgr) GetFromAuditTime(auditTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_time` = ?", auditTime).Find(&results).Error

	return
}

// GetBatchFromAuditTime 批量查找 审核时间
func (obj *_FaCustomerBillMgr) GetBatchFromAuditTime(auditTimes []time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_time` IN (?)", auditTimes).Find(&results).Error

	return
}

// GetFromIsReceipt 通过is_receipt获取内容 已生成收款单
func (obj *_FaCustomerBillMgr) GetFromIsReceipt(isReceipt []uint8) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`is_receipt` = ?", isReceipt).Find(&results).Error

	return
}

// GetBatchFromIsReceipt 批量查找 已生成收款单
func (obj *_FaCustomerBillMgr) GetBatchFromIsReceipt(isReceipts [][]uint8) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`is_receipt` IN (?)", isReceipts).Find(&results).Error

	return
}

// GetFromReceiptNumberUser 通过receipt_number_user获取内容 收款单生成用户
func (obj *_FaCustomerBillMgr) GetFromReceiptNumberUser(receiptNumberUser string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receipt_number_user` = ?", receiptNumberUser).Find(&results).Error

	return
}

// GetBatchFromReceiptNumberUser 批量查找 收款单生成用户
func (obj *_FaCustomerBillMgr) GetBatchFromReceiptNumberUser(receiptNumberUsers []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receipt_number_user` IN (?)", receiptNumberUsers).Find(&results).Error

	return
}

// GetFromReceiptNumberUserID 通过receipt_number_user_id获取内容 收款单生成用户ID
func (obj *_FaCustomerBillMgr) GetFromReceiptNumberUserID(receiptNumberUserID int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receipt_number_user_id` = ?", receiptNumberUserID).Find(&results).Error

	return
}

// GetBatchFromReceiptNumberUserID 批量查找 收款单生成用户ID
func (obj *_FaCustomerBillMgr) GetBatchFromReceiptNumberUserID(receiptNumberUserIDs []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receipt_number_user_id` IN (?)", receiptNumberUserIDs).Find(&results).Error

	return
}

// GetFromReceiptNumberTime 通过receipt_number_time获取内容 收款单生成用户生成时间
func (obj *_FaCustomerBillMgr) GetFromReceiptNumberTime(receiptNumberTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receipt_number_time` = ?", receiptNumberTime).Find(&results).Error

	return
}

// GetBatchFromReceiptNumberTime 批量查找 收款单生成用户生成时间
func (obj *_FaCustomerBillMgr) GetBatchFromReceiptNumberTime(receiptNumberTimes []time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`receipt_number_time` IN (?)", receiptNumberTimes).Find(&results).Error

	return
}

// GetFromIsAccept 通过is_accept获取内容 已核收(关联收款单状态)
func (obj *_FaCustomerBillMgr) GetFromIsAccept(isAccept []uint8) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`is_accept` = ?", isAccept).Find(&results).Error

	return
}

// GetBatchFromIsAccept 批量查找 已核收(关联收款单状态)
func (obj *_FaCustomerBillMgr) GetBatchFromIsAccept(isAccepts [][]uint8) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`is_accept` IN (?)", isAccepts).Find(&results).Error

	return
}

// GetFromAcceptTime 通过accept_time获取内容 核收时间
func (obj *_FaCustomerBillMgr) GetFromAcceptTime(acceptTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// GetBatchFromAcceptTime 批量查找 核收时间
func (obj *_FaCustomerBillMgr) GetBatchFromAcceptTime(acceptTimes []time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_time` IN (?)", acceptTimes).Find(&results).Error

	return
}

// GetFromAcceptUser 通过accept_user获取内容 核收用户
func (obj *_FaCustomerBillMgr) GetFromAcceptUser(acceptUser string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_user` = ?", acceptUser).Find(&results).Error

	return
}

// GetBatchFromAcceptUser 批量查找 核收用户
func (obj *_FaCustomerBillMgr) GetBatchFromAcceptUser(acceptUsers []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_user` IN (?)", acceptUsers).Find(&results).Error

	return
}

// GetFromAcceptUserID 通过accept_user_id获取内容 核收用户ID
func (obj *_FaCustomerBillMgr) GetFromAcceptUserID(acceptUserID int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_user_id` = ?", acceptUserID).Find(&results).Error

	return
}

// GetBatchFromAcceptUserID 批量查找 核收用户ID
func (obj *_FaCustomerBillMgr) GetBatchFromAcceptUserID(acceptUserIDs []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_user_id` IN (?)", acceptUserIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaCustomerBillMgr) GetFromRemark(remark string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaCustomerBillMgr) GetBatchFromRemark(remarks []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromFinancialRemark 通过financial_remark获取内容 财务备注
func (obj *_FaCustomerBillMgr) GetFromFinancialRemark(financialRemark string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`financial_remark` = ?", financialRemark).Find(&results).Error

	return
}

// GetBatchFromFinancialRemark 批量查找 财务备注
func (obj *_FaCustomerBillMgr) GetBatchFromFinancialRemark(financialRemarks []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`financial_remark` IN (?)", financialRemarks).Find(&results).Error

	return
}

// GetFromAcceptRemark 通过accept_remark获取内容 核收备注
func (obj *_FaCustomerBillMgr) GetFromAcceptRemark(acceptRemark string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_remark` = ?", acceptRemark).Find(&results).Error

	return
}

// GetBatchFromAcceptRemark 批量查找 核收备注
func (obj *_FaCustomerBillMgr) GetBatchFromAcceptRemark(acceptRemarks []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_remark` IN (?)", acceptRemarks).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaCustomerBillMgr) GetFromUpdateUser(updateUser int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaCustomerBillMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaCustomerBillMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaCustomerBillMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaCustomerBillMgr) GetFromVersion(version int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaCustomerBillMgr) GetBatchFromVersion(versions []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaCustomerBillMgr) GetFromDeleted(deleted int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaCustomerBillMgr) GetBatchFromDeleted(deleteds []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromChannelHaul 通过channel_haul获取内容 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_FaCustomerBillMgr) GetFromChannelHaul(channelHaul string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`channel_haul` = ?", channelHaul).Find(&results).Error

	return
}

// GetBatchFromChannelHaul 批量查找 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_FaCustomerBillMgr) GetBatchFromChannelHaul(channelHauls []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`channel_haul` IN (?)", channelHauls).Find(&results).Error

	return
}

// GetFromKingdeeID 通过kingdee_id获取内容 金蝶客户内码
func (obj *_FaCustomerBillMgr) GetFromKingdeeID(kingdeeID string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}

// GetBatchFromKingdeeID 批量查找 金蝶客户内码
func (obj *_FaCustomerBillMgr) GetBatchFromKingdeeID(kingdeeIDs []string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`kingdee_id` IN (?)", kingdeeIDs).Find(&results).Error

	return
}

// GetFromLastSyncTime 通过last_sync_time获取内容 最后同步时间
func (obj *_FaCustomerBillMgr) GetFromLastSyncTime(lastSyncTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`last_sync_time` = ?", lastSyncTime).Find(&results).Error

	return
}

// GetBatchFromLastSyncTime 批量查找 最后同步时间
func (obj *_FaCustomerBillMgr) GetBatchFromLastSyncTime(lastSyncTimes []time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`last_sync_time` IN (?)", lastSyncTimes).Find(&results).Error

	return
}

// GetFromLastSyncUser 通过last_sync_user获取内容 最后同步人
func (obj *_FaCustomerBillMgr) GetFromLastSyncUser(lastSyncUser int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`last_sync_user` = ?", lastSyncUser).Find(&results).Error

	return
}

// GetBatchFromLastSyncUser 批量查找 最后同步人
func (obj *_FaCustomerBillMgr) GetBatchFromLastSyncUser(lastSyncUsers []int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`last_sync_user` IN (?)", lastSyncUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaCustomerBillMgr) FetchByPrimaryKey(id int) (result FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUniqueBillBatchNumber primary or index 获取唯一内容
func (obj *_FaCustomerBillMgr) FetchUniqueByUniqueBillBatchNumber(billBatchNumber string) (result FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`bill_batch_number` = ?", billBatchNumber).First(&result).Error

	return
}

// FetchIndexByIndexBillBatchNumber  获取多个内容
func (obj *_FaCustomerBillMgr) FetchIndexByIndexBillBatchNumber(billBatchNumber string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// FetchIndexByFaCustomerBillCustomerIDIndex  获取多个内容
func (obj *_FaCustomerBillMgr) FetchIndexByFaCustomerBillCustomerIDIndex(customerID int) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexByIndexCreateTime  获取多个内容
func (obj *_FaCustomerBillMgr) FetchIndexByIndexCreateTime(createTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByIndexAuditTime  获取多个内容
func (obj *_FaCustomerBillMgr) FetchIndexByIndexAuditTime(auditTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`audit_time` = ?", auditTime).Find(&results).Error

	return
}

// FetchIndexByIndexAcceptTime  获取多个内容
func (obj *_FaCustomerBillMgr) FetchIndexByIndexAcceptTime(acceptTime time.Time) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// FetchIndexByIndexChannelHaul  获取多个内容
func (obj *_FaCustomerBillMgr) FetchIndexByIndexChannelHaul(channelHaul string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`channel_haul` = ?", channelHaul).Find(&results).Error

	return
}

// FetchIndexByIndexKingdeeID  获取多个内容
func (obj *_FaCustomerBillMgr) FetchIndexByIndexKingdeeID(kingdeeID string) (results []*FaCustomerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBill{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}
