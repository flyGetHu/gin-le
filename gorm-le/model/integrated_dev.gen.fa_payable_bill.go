package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaPayableBillMgr struct {
	*_BaseMgr
}

// FaPayableBillMgr open func
func FaPayableBillMgr(db *gorm.DB) *_FaPayableBillMgr {
	if db == nil {
		panic(fmt.Errorf("FaPayableBillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaPayableBillMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_payable_bill"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaPayableBillMgr) GetTableName() string {
	return "fa_payable_bill"
}

// Reset 重置gorm会话
func (obj *_FaPayableBillMgr) Reset() *_FaPayableBillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaPayableBillMgr) Get() (result FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaPayableBillMgr) Gets() (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaPayableBillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaPayableBillMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBillBatchNumber bill_batch_number获取 账单批次号
func (obj *_FaPayableBillMgr) WithBillBatchNumber(billBatchNumber string) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number"] = billBatchNumber })
}

// WithBillBatchStatus bill_batch_status获取 账单状态
func (obj *_FaPayableBillMgr) WithBillBatchStatus(billBatchStatus int) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_status"] = billBatchStatus })
}

// WithPaymentNumber payment_number获取 付款单号
func (obj *_FaPayableBillMgr) WithPaymentNumber(paymentNumber string) Option {
	return optionFunc(func(o *options) { o.query["payment_number"] = paymentNumber })
}

// WithBusinessEntityCode business_entity_code获取 业务主体代码
func (obj *_FaPayableBillMgr) WithBusinessEntityCode(businessEntityCode string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_code"] = businessEntityCode })
}

// WithBusinessEntityName business_entity_name获取 业务实体名称
func (obj *_FaPayableBillMgr) WithBusinessEntityName(businessEntityName string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_name"] = businessEntityName })
}

// WithBusinessEntityType business_entity_type获取 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillMgr) WithBusinessEntityType(businessEntityType bool) Option {
	return optionFunc(func(o *options) { o.query["business_entity_type"] = businessEntityType })
}

// WithOrderQuantity order_quantity获取 订单数量
func (obj *_FaPayableBillMgr) WithOrderQuantity(orderQuantity int) Option {
	return optionFunc(func(o *options) { o.query["order_quantity"] = orderQuantity })
}

// WithFeeQuantity fee_quantity获取 费用条数
func (obj *_FaPayableBillMgr) WithFeeQuantity(feeQuantity int) Option {
	return optionFunc(func(o *options) { o.query["fee_quantity"] = feeQuantity })
}

// WithWeight weight获取 计费总重量
func (obj *_FaPayableBillMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithPayableRmb payable_rmb获取 应付总费用(RMB)
func (obj *_FaPayableBillMgr) WithPayableRmb(payableRmb float64) Option {
	return optionFunc(func(o *options) { o.query["payable_rmb"] = payableRmb })
}

// WithActualPaymentRmb actual_payment_rmb获取 实际付款金额(RMB)
func (obj *_FaPayableBillMgr) WithActualPaymentRmb(actualPaymentRmb float64) Option {
	return optionFunc(func(o *options) { o.query["actual_payment_rmb"] = actualPaymentRmb })
}

// WithCreateTime create_time获取 生成时间
func (obj *_FaPayableBillMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应付人员
func (obj *_FaPayableBillMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithIsAudit is_audit获取 审核状态(关联对账单批次号状态)
func (obj *_FaPayableBillMgr) WithIsAudit(isAudit []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_audit"] = isAudit })
}

// WithIsAuditRollback is_audit_rollback获取 是否反审标识 默认false ,被反审后true,重新审核后false
func (obj *_FaPayableBillMgr) WithIsAuditRollback(isAuditRollback []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_audit_rollback"] = isAuditRollback })
}

// WithAuditRollbackUser audit_rollback_user获取 反审核人员
func (obj *_FaPayableBillMgr) WithAuditRollbackUser(auditRollbackUser string) Option {
	return optionFunc(func(o *options) { o.query["audit_rollback_user"] = auditRollbackUser })
}

// WithAuditRollbackUserID audit_rollback_user_id获取 反审核人员ID
func (obj *_FaPayableBillMgr) WithAuditRollbackUserID(auditRollbackUserID int) Option {
	return optionFunc(func(o *options) { o.query["audit_rollback_user_id"] = auditRollbackUserID })
}

// WithCancelUser cancel_user获取 取消应付用户
func (obj *_FaPayableBillMgr) WithCancelUser(cancelUser string) Option {
	return optionFunc(func(o *options) { o.query["cancel_user"] = cancelUser })
}

// WithCancelUserID cancel_user_id获取 取消应付用户ID
func (obj *_FaPayableBillMgr) WithCancelUserID(cancelUserID int) Option {
	return optionFunc(func(o *options) { o.query["cancel_user_id"] = cancelUserID })
}

// WithAuditUser audit_user获取 审核用户
func (obj *_FaPayableBillMgr) WithAuditUser(auditUser string) Option {
	return optionFunc(func(o *options) { o.query["audit_user"] = auditUser })
}

// WithAuditUserID audit_user_id获取 审核用户
func (obj *_FaPayableBillMgr) WithAuditUserID(auditUserID int) Option {
	return optionFunc(func(o *options) { o.query["audit_user_id"] = auditUserID })
}

// WithAuditTime audit_time获取 审核时间
func (obj *_FaPayableBillMgr) WithAuditTime(auditTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["audit_time"] = auditTime })
}

// WithIsPayment is_payment获取 已生成付款单
func (obj *_FaPayableBillMgr) WithIsPayment(isPayment []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_payment"] = isPayment })
}

// WithPaymentNumberUser payment_number_user获取 付款单生成用户
func (obj *_FaPayableBillMgr) WithPaymentNumberUser(paymentNumberUser string) Option {
	return optionFunc(func(o *options) { o.query["payment_number_user"] = paymentNumberUser })
}

// WithPaymentNumberUserID payment_number_user_id获取 付款单生成用户ID
func (obj *_FaPayableBillMgr) WithPaymentNumberUserID(paymentNumberUserID int) Option {
	return optionFunc(func(o *options) { o.query["payment_number_user_id"] = paymentNumberUserID })
}

// WithPaymentNumberTime payment_number_time获取 付款单生成用户生成时间
func (obj *_FaPayableBillMgr) WithPaymentNumberTime(paymentNumberTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["payment_number_time"] = paymentNumberTime })
}

// WithIsAccept is_accept获取 已核收(关联付款单状态)
func (obj *_FaPayableBillMgr) WithIsAccept(isAccept []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_accept"] = isAccept })
}

// WithAcceptTime accept_time获取 核收时间
func (obj *_FaPayableBillMgr) WithAcceptTime(acceptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["accept_time"] = acceptTime })
}

// WithAcceptUser accept_user获取 核收用户
func (obj *_FaPayableBillMgr) WithAcceptUser(acceptUser string) Option {
	return optionFunc(func(o *options) { o.query["accept_user"] = acceptUser })
}

// WithAcceptUserID accept_user_id获取 核收用户ID
func (obj *_FaPayableBillMgr) WithAcceptUserID(acceptUserID int) Option {
	return optionFunc(func(o *options) { o.query["accept_user_id"] = acceptUserID })
}

// WithRemark remark获取 备注
func (obj *_FaPayableBillMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithFinancialRemark financial_remark获取 财务备注
func (obj *_FaPayableBillMgr) WithFinancialRemark(financialRemark string) Option {
	return optionFunc(func(o *options) { o.query["financial_remark"] = financialRemark })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaPayableBillMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaPayableBillMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaPayableBillMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaPayableBillMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithPayableRemark payable_remark获取 核付备注
func (obj *_FaPayableBillMgr) WithPayableRemark(payableRemark string) Option {
	return optionFunc(func(o *options) { o.query["payable_remark"] = payableRemark })
}

// WithChannelHaul channel_haul获取 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_FaPayableBillMgr) WithChannelHaul(channelHaul string) Option {
	return optionFunc(func(o *options) { o.query["channel_haul"] = channelHaul })
}

// WithCustomerChannelNames customer_channel_names获取 客户渠道名称，多个用,隔开
func (obj *_FaPayableBillMgr) WithCustomerChannelNames(customerChannelNames string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_names"] = customerChannelNames })
}

// WithKingdeeID kingdee_id获取 金蝶客户内码
func (obj *_FaPayableBillMgr) WithKingdeeID(kingdeeID string) Option {
	return optionFunc(func(o *options) { o.query["kingdee_id"] = kingdeeID })
}

// WithLastSyncTime last_sync_time获取 最后同步时间
func (obj *_FaPayableBillMgr) WithLastSyncTime(lastSyncTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_sync_time"] = lastSyncTime })
}

// WithLastSyncUser last_sync_user获取 最后同步人
func (obj *_FaPayableBillMgr) WithLastSyncUser(lastSyncUser int) Option {
	return optionFunc(func(o *options) { o.query["last_sync_user"] = lastSyncUser })
}

// GetByOption 功能选项模式获取
func (obj *_FaPayableBillMgr) GetByOption(opts ...Option) (result FaPayableBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaPayableBillMgr) GetByOptions(opts ...Option) (results []*FaPayableBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaPayableBillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaPayableBill, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where(options.query)
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
func (obj *_FaPayableBillMgr) GetFromID(id int) (result FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaPayableBillMgr) GetBatchFromID(ids []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBillBatchNumber 通过bill_batch_number获取内容 账单批次号
func (obj *_FaPayableBillMgr) GetFromBillBatchNumber(billBatchNumber string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumber 批量查找 账单批次号
func (obj *_FaPayableBillMgr) GetBatchFromBillBatchNumber(billBatchNumbers []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`bill_batch_number` IN (?)", billBatchNumbers).Find(&results).Error

	return
}

// GetFromBillBatchStatus 通过bill_batch_status获取内容 账单状态
func (obj *_FaPayableBillMgr) GetFromBillBatchStatus(billBatchStatus int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`bill_batch_status` = ?", billBatchStatus).Find(&results).Error

	return
}

// GetBatchFromBillBatchStatus 批量查找 账单状态
func (obj *_FaPayableBillMgr) GetBatchFromBillBatchStatus(billBatchStatuss []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`bill_batch_status` IN (?)", billBatchStatuss).Find(&results).Error

	return
}

// GetFromPaymentNumber 通过payment_number获取内容 付款单号
func (obj *_FaPayableBillMgr) GetFromPaymentNumber(paymentNumber string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payment_number` = ?", paymentNumber).Find(&results).Error

	return
}

// GetBatchFromPaymentNumber 批量查找 付款单号
func (obj *_FaPayableBillMgr) GetBatchFromPaymentNumber(paymentNumbers []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payment_number` IN (?)", paymentNumbers).Find(&results).Error

	return
}

// GetFromBusinessEntityCode 通过business_entity_code获取内容 业务主体代码
func (obj *_FaPayableBillMgr) GetFromBusinessEntityCode(businessEntityCode string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityCode 批量查找 业务主体代码
func (obj *_FaPayableBillMgr) GetBatchFromBusinessEntityCode(businessEntityCodes []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`business_entity_code` IN (?)", businessEntityCodes).Find(&results).Error

	return
}

// GetFromBusinessEntityName 通过business_entity_name获取内容 业务实体名称
func (obj *_FaPayableBillMgr) GetFromBusinessEntityName(businessEntityName string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`business_entity_name` = ?", businessEntityName).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityName 批量查找 业务实体名称
func (obj *_FaPayableBillMgr) GetBatchFromBusinessEntityName(businessEntityNames []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`business_entity_name` IN (?)", businessEntityNames).Find(&results).Error

	return
}

// GetFromBusinessEntityType 通过business_entity_type获取内容 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillMgr) GetFromBusinessEntityType(businessEntityType bool) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`business_entity_type` = ?", businessEntityType).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityType 批量查找 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillMgr) GetBatchFromBusinessEntityType(businessEntityTypes []bool) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`business_entity_type` IN (?)", businessEntityTypes).Find(&results).Error

	return
}

// GetFromOrderQuantity 通过order_quantity获取内容 订单数量
func (obj *_FaPayableBillMgr) GetFromOrderQuantity(orderQuantity int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`order_quantity` = ?", orderQuantity).Find(&results).Error

	return
}

// GetBatchFromOrderQuantity 批量查找 订单数量
func (obj *_FaPayableBillMgr) GetBatchFromOrderQuantity(orderQuantitys []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`order_quantity` IN (?)", orderQuantitys).Find(&results).Error

	return
}

// GetFromFeeQuantity 通过fee_quantity获取内容 费用条数
func (obj *_FaPayableBillMgr) GetFromFeeQuantity(feeQuantity int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`fee_quantity` = ?", feeQuantity).Find(&results).Error

	return
}

// GetBatchFromFeeQuantity 批量查找 费用条数
func (obj *_FaPayableBillMgr) GetBatchFromFeeQuantity(feeQuantitys []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`fee_quantity` IN (?)", feeQuantitys).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 计费总重量
func (obj *_FaPayableBillMgr) GetFromWeight(weight float64) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 计费总重量
func (obj *_FaPayableBillMgr) GetBatchFromWeight(weights []float64) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromPayableRmb 通过payable_rmb获取内容 应付总费用(RMB)
func (obj *_FaPayableBillMgr) GetFromPayableRmb(payableRmb float64) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payable_rmb` = ?", payableRmb).Find(&results).Error

	return
}

// GetBatchFromPayableRmb 批量查找 应付总费用(RMB)
func (obj *_FaPayableBillMgr) GetBatchFromPayableRmb(payableRmbs []float64) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payable_rmb` IN (?)", payableRmbs).Find(&results).Error

	return
}

// GetFromActualPaymentRmb 通过actual_payment_rmb获取内容 实际付款金额(RMB)
func (obj *_FaPayableBillMgr) GetFromActualPaymentRmb(actualPaymentRmb float64) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`actual_payment_rmb` = ?", actualPaymentRmb).Find(&results).Error

	return
}

// GetBatchFromActualPaymentRmb 批量查找 实际付款金额(RMB)
func (obj *_FaPayableBillMgr) GetBatchFromActualPaymentRmb(actualPaymentRmbs []float64) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`actual_payment_rmb` IN (?)", actualPaymentRmbs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 生成时间
func (obj *_FaPayableBillMgr) GetFromCreateTime(createTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 生成时间
func (obj *_FaPayableBillMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应付人员
func (obj *_FaPayableBillMgr) GetFromCreateUser(createUser int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应付人员
func (obj *_FaPayableBillMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromIsAudit 通过is_audit获取内容 审核状态(关联对账单批次号状态)
func (obj *_FaPayableBillMgr) GetFromIsAudit(isAudit []uint8) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`is_audit` = ?", isAudit).Find(&results).Error

	return
}

// GetBatchFromIsAudit 批量查找 审核状态(关联对账单批次号状态)
func (obj *_FaPayableBillMgr) GetBatchFromIsAudit(isAudits [][]uint8) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`is_audit` IN (?)", isAudits).Find(&results).Error

	return
}

// GetFromIsAuditRollback 通过is_audit_rollback获取内容 是否反审标识 默认false ,被反审后true,重新审核后false
func (obj *_FaPayableBillMgr) GetFromIsAuditRollback(isAuditRollback []uint8) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`is_audit_rollback` = ?", isAuditRollback).Find(&results).Error

	return
}

// GetBatchFromIsAuditRollback 批量查找 是否反审标识 默认false ,被反审后true,重新审核后false
func (obj *_FaPayableBillMgr) GetBatchFromIsAuditRollback(isAuditRollbacks [][]uint8) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`is_audit_rollback` IN (?)", isAuditRollbacks).Find(&results).Error

	return
}

// GetFromAuditRollbackUser 通过audit_rollback_user获取内容 反审核人员
func (obj *_FaPayableBillMgr) GetFromAuditRollbackUser(auditRollbackUser string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_rollback_user` = ?", auditRollbackUser).Find(&results).Error

	return
}

// GetBatchFromAuditRollbackUser 批量查找 反审核人员
func (obj *_FaPayableBillMgr) GetBatchFromAuditRollbackUser(auditRollbackUsers []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_rollback_user` IN (?)", auditRollbackUsers).Find(&results).Error

	return
}

// GetFromAuditRollbackUserID 通过audit_rollback_user_id获取内容 反审核人员ID
func (obj *_FaPayableBillMgr) GetFromAuditRollbackUserID(auditRollbackUserID int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_rollback_user_id` = ?", auditRollbackUserID).Find(&results).Error

	return
}

// GetBatchFromAuditRollbackUserID 批量查找 反审核人员ID
func (obj *_FaPayableBillMgr) GetBatchFromAuditRollbackUserID(auditRollbackUserIDs []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_rollback_user_id` IN (?)", auditRollbackUserIDs).Find(&results).Error

	return
}

// GetFromCancelUser 通过cancel_user获取内容 取消应付用户
func (obj *_FaPayableBillMgr) GetFromCancelUser(cancelUser string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`cancel_user` = ?", cancelUser).Find(&results).Error

	return
}

// GetBatchFromCancelUser 批量查找 取消应付用户
func (obj *_FaPayableBillMgr) GetBatchFromCancelUser(cancelUsers []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`cancel_user` IN (?)", cancelUsers).Find(&results).Error

	return
}

// GetFromCancelUserID 通过cancel_user_id获取内容 取消应付用户ID
func (obj *_FaPayableBillMgr) GetFromCancelUserID(cancelUserID int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`cancel_user_id` = ?", cancelUserID).Find(&results).Error

	return
}

// GetBatchFromCancelUserID 批量查找 取消应付用户ID
func (obj *_FaPayableBillMgr) GetBatchFromCancelUserID(cancelUserIDs []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`cancel_user_id` IN (?)", cancelUserIDs).Find(&results).Error

	return
}

// GetFromAuditUser 通过audit_user获取内容 审核用户
func (obj *_FaPayableBillMgr) GetFromAuditUser(auditUser string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_user` = ?", auditUser).Find(&results).Error

	return
}

// GetBatchFromAuditUser 批量查找 审核用户
func (obj *_FaPayableBillMgr) GetBatchFromAuditUser(auditUsers []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_user` IN (?)", auditUsers).Find(&results).Error

	return
}

// GetFromAuditUserID 通过audit_user_id获取内容 审核用户
func (obj *_FaPayableBillMgr) GetFromAuditUserID(auditUserID int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_user_id` = ?", auditUserID).Find(&results).Error

	return
}

// GetBatchFromAuditUserID 批量查找 审核用户
func (obj *_FaPayableBillMgr) GetBatchFromAuditUserID(auditUserIDs []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_user_id` IN (?)", auditUserIDs).Find(&results).Error

	return
}

// GetFromAuditTime 通过audit_time获取内容 审核时间
func (obj *_FaPayableBillMgr) GetFromAuditTime(auditTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_time` = ?", auditTime).Find(&results).Error

	return
}

// GetBatchFromAuditTime 批量查找 审核时间
func (obj *_FaPayableBillMgr) GetBatchFromAuditTime(auditTimes []time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_time` IN (?)", auditTimes).Find(&results).Error

	return
}

// GetFromIsPayment 通过is_payment获取内容 已生成付款单
func (obj *_FaPayableBillMgr) GetFromIsPayment(isPayment []uint8) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`is_payment` = ?", isPayment).Find(&results).Error

	return
}

// GetBatchFromIsPayment 批量查找 已生成付款单
func (obj *_FaPayableBillMgr) GetBatchFromIsPayment(isPayments [][]uint8) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`is_payment` IN (?)", isPayments).Find(&results).Error

	return
}

// GetFromPaymentNumberUser 通过payment_number_user获取内容 付款单生成用户
func (obj *_FaPayableBillMgr) GetFromPaymentNumberUser(paymentNumberUser string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payment_number_user` = ?", paymentNumberUser).Find(&results).Error

	return
}

// GetBatchFromPaymentNumberUser 批量查找 付款单生成用户
func (obj *_FaPayableBillMgr) GetBatchFromPaymentNumberUser(paymentNumberUsers []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payment_number_user` IN (?)", paymentNumberUsers).Find(&results).Error

	return
}

// GetFromPaymentNumberUserID 通过payment_number_user_id获取内容 付款单生成用户ID
func (obj *_FaPayableBillMgr) GetFromPaymentNumberUserID(paymentNumberUserID int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payment_number_user_id` = ?", paymentNumberUserID).Find(&results).Error

	return
}

// GetBatchFromPaymentNumberUserID 批量查找 付款单生成用户ID
func (obj *_FaPayableBillMgr) GetBatchFromPaymentNumberUserID(paymentNumberUserIDs []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payment_number_user_id` IN (?)", paymentNumberUserIDs).Find(&results).Error

	return
}

// GetFromPaymentNumberTime 通过payment_number_time获取内容 付款单生成用户生成时间
func (obj *_FaPayableBillMgr) GetFromPaymentNumberTime(paymentNumberTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payment_number_time` = ?", paymentNumberTime).Find(&results).Error

	return
}

// GetBatchFromPaymentNumberTime 批量查找 付款单生成用户生成时间
func (obj *_FaPayableBillMgr) GetBatchFromPaymentNumberTime(paymentNumberTimes []time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payment_number_time` IN (?)", paymentNumberTimes).Find(&results).Error

	return
}

// GetFromIsAccept 通过is_accept获取内容 已核收(关联付款单状态)
func (obj *_FaPayableBillMgr) GetFromIsAccept(isAccept []uint8) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`is_accept` = ?", isAccept).Find(&results).Error

	return
}

// GetBatchFromIsAccept 批量查找 已核收(关联付款单状态)
func (obj *_FaPayableBillMgr) GetBatchFromIsAccept(isAccepts [][]uint8) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`is_accept` IN (?)", isAccepts).Find(&results).Error

	return
}

// GetFromAcceptTime 通过accept_time获取内容 核收时间
func (obj *_FaPayableBillMgr) GetFromAcceptTime(acceptTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// GetBatchFromAcceptTime 批量查找 核收时间
func (obj *_FaPayableBillMgr) GetBatchFromAcceptTime(acceptTimes []time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`accept_time` IN (?)", acceptTimes).Find(&results).Error

	return
}

// GetFromAcceptUser 通过accept_user获取内容 核收用户
func (obj *_FaPayableBillMgr) GetFromAcceptUser(acceptUser string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`accept_user` = ?", acceptUser).Find(&results).Error

	return
}

// GetBatchFromAcceptUser 批量查找 核收用户
func (obj *_FaPayableBillMgr) GetBatchFromAcceptUser(acceptUsers []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`accept_user` IN (?)", acceptUsers).Find(&results).Error

	return
}

// GetFromAcceptUserID 通过accept_user_id获取内容 核收用户ID
func (obj *_FaPayableBillMgr) GetFromAcceptUserID(acceptUserID int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`accept_user_id` = ?", acceptUserID).Find(&results).Error

	return
}

// GetBatchFromAcceptUserID 批量查找 核收用户ID
func (obj *_FaPayableBillMgr) GetBatchFromAcceptUserID(acceptUserIDs []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`accept_user_id` IN (?)", acceptUserIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaPayableBillMgr) GetFromRemark(remark string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaPayableBillMgr) GetBatchFromRemark(remarks []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromFinancialRemark 通过financial_remark获取内容 财务备注
func (obj *_FaPayableBillMgr) GetFromFinancialRemark(financialRemark string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`financial_remark` = ?", financialRemark).Find(&results).Error

	return
}

// GetBatchFromFinancialRemark 批量查找 财务备注
func (obj *_FaPayableBillMgr) GetBatchFromFinancialRemark(financialRemarks []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`financial_remark` IN (?)", financialRemarks).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaPayableBillMgr) GetFromUpdateUser(updateUser int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaPayableBillMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaPayableBillMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaPayableBillMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaPayableBillMgr) GetFromVersion(version int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaPayableBillMgr) GetBatchFromVersion(versions []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaPayableBillMgr) GetFromDeleted(deleted int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaPayableBillMgr) GetBatchFromDeleted(deleteds []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromPayableRemark 通过payable_remark获取内容 核付备注
func (obj *_FaPayableBillMgr) GetFromPayableRemark(payableRemark string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payable_remark` = ?", payableRemark).Find(&results).Error

	return
}

// GetBatchFromPayableRemark 批量查找 核付备注
func (obj *_FaPayableBillMgr) GetBatchFromPayableRemark(payableRemarks []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`payable_remark` IN (?)", payableRemarks).Find(&results).Error

	return
}

// GetFromChannelHaul 通过channel_haul获取内容 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_FaPayableBillMgr) GetFromChannelHaul(channelHaul string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`channel_haul` = ?", channelHaul).Find(&results).Error

	return
}

// GetBatchFromChannelHaul 批量查找 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_FaPayableBillMgr) GetBatchFromChannelHaul(channelHauls []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`channel_haul` IN (?)", channelHauls).Find(&results).Error

	return
}

// GetFromCustomerChannelNames 通过customer_channel_names获取内容 客户渠道名称，多个用,隔开
func (obj *_FaPayableBillMgr) GetFromCustomerChannelNames(customerChannelNames string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`customer_channel_names` = ?", customerChannelNames).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelNames 批量查找 客户渠道名称，多个用,隔开
func (obj *_FaPayableBillMgr) GetBatchFromCustomerChannelNames(customerChannelNamess []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`customer_channel_names` IN (?)", customerChannelNamess).Find(&results).Error

	return
}

// GetFromKingdeeID 通过kingdee_id获取内容 金蝶客户内码
func (obj *_FaPayableBillMgr) GetFromKingdeeID(kingdeeID string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}

// GetBatchFromKingdeeID 批量查找 金蝶客户内码
func (obj *_FaPayableBillMgr) GetBatchFromKingdeeID(kingdeeIDs []string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`kingdee_id` IN (?)", kingdeeIDs).Find(&results).Error

	return
}

// GetFromLastSyncTime 通过last_sync_time获取内容 最后同步时间
func (obj *_FaPayableBillMgr) GetFromLastSyncTime(lastSyncTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`last_sync_time` = ?", lastSyncTime).Find(&results).Error

	return
}

// GetBatchFromLastSyncTime 批量查找 最后同步时间
func (obj *_FaPayableBillMgr) GetBatchFromLastSyncTime(lastSyncTimes []time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`last_sync_time` IN (?)", lastSyncTimes).Find(&results).Error

	return
}

// GetFromLastSyncUser 通过last_sync_user获取内容 最后同步人
func (obj *_FaPayableBillMgr) GetFromLastSyncUser(lastSyncUser int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`last_sync_user` = ?", lastSyncUser).Find(&results).Error

	return
}

// GetBatchFromLastSyncUser 批量查找 最后同步人
func (obj *_FaPayableBillMgr) GetBatchFromLastSyncUser(lastSyncUsers []int) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`last_sync_user` IN (?)", lastSyncUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaPayableBillMgr) FetchByPrimaryKey(id int) (result FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUniqueBillBatchNumber primary or index 获取唯一内容
func (obj *_FaPayableBillMgr) FetchUniqueByUniqueBillBatchNumber(billBatchNumber string) (result FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`bill_batch_number` = ?", billBatchNumber).First(&result).Error

	return
}

// FetchIndexByIndexBillBatchNumber  获取多个内容
func (obj *_FaPayableBillMgr) FetchIndexByIndexBillBatchNumber(billBatchNumber string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCreateTime  获取多个内容
func (obj *_FaPayableBillMgr) FetchIndexByIndexCreateTime(createTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByIndexAuditTime  获取多个内容
func (obj *_FaPayableBillMgr) FetchIndexByIndexAuditTime(auditTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`audit_time` = ?", auditTime).Find(&results).Error

	return
}

// FetchIndexByIndexAcceptTime  获取多个内容
func (obj *_FaPayableBillMgr) FetchIndexByIndexAcceptTime(acceptTime time.Time) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// FetchIndexByIndexChannelHaul  获取多个内容
func (obj *_FaPayableBillMgr) FetchIndexByIndexChannelHaul(channelHaul string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`channel_haul` = ?", channelHaul).Find(&results).Error

	return
}

// FetchIndexByIndexKingdeeID  获取多个内容
func (obj *_FaPayableBillMgr) FetchIndexByIndexKingdeeID(kingdeeID string) (results []*FaPayableBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBill{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}
