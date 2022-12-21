package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaCustomerBillFeeDetailsMgr struct {
	*_BaseMgr
}

// FaCustomerBillFeeDetailsMgr open func
func FaCustomerBillFeeDetailsMgr(db *gorm.DB) *_FaCustomerBillFeeDetailsMgr {
	if db == nil {
		panic(fmt.Errorf("FaCustomerBillFeeDetailsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaCustomerBillFeeDetailsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_customer_bill_fee_details"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaCustomerBillFeeDetailsMgr) GetTableName() string {
	return "fa_customer_bill_fee_details"
}

// Reset 重置gorm会话
func (obj *_FaCustomerBillFeeDetailsMgr) Reset() *_FaCustomerBillFeeDetailsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaCustomerBillFeeDetailsMgr) Get() (result FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaCustomerBillFeeDetailsMgr) Gets() (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaCustomerBillFeeDetailsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaCustomerBillFeeDetailsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_FaCustomerBillFeeDetailsMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_FaCustomerBillFeeDetailsMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_FaCustomerBillFeeDetailsMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_FaCustomerBillFeeDetailsMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_FaCustomerBillFeeDetailsMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_FaCustomerBillFeeDetailsMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_FaCustomerBillFeeDetailsMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_FaCustomerBillFeeDetailsMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithReceiptTime receipt_time获取 入库时间
func (obj *_FaCustomerBillFeeDetailsMgr) WithReceiptTime(receiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_time"] = receiptTime })
}

// WithSendTime send_time获取 出库时间
func (obj *_FaCustomerBillFeeDetailsMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithIsSend is_send获取 是否出库
func (obj *_FaCustomerBillFeeDetailsMgr) WithIsSend(isSend []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_send"] = isSend })
}

// WithWeight weight获取 预报重量
func (obj *_FaCustomerBillFeeDetailsMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithWeighingWeight weighing_weight获取 称重重量
func (obj *_FaCustomerBillFeeDetailsMgr) WithWeighingWeight(weighingWeight float64) Option {
	return optionFunc(func(o *options) { o.query["weighing_weight"] = weighingWeight })
}

// WithFoamWeight foam_weight获取 泡重
func (obj *_FaCustomerBillFeeDetailsMgr) WithFoamWeight(foamWeight float64) Option {
	return optionFunc(func(o *options) { o.query["foam_weight"] = foamWeight })
}

// WithChargeWeight charge_weight获取 计费重量
func (obj *_FaCustomerBillFeeDetailsMgr) WithChargeWeight(chargeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["charge_weight"] = chargeWeight })
}

// WithFeeTypeName fee_type_name获取 费用类型名称
func (obj *_FaCustomerBillFeeDetailsMgr) WithFeeTypeName(feeTypeName string) Option {
	return optionFunc(func(o *options) { o.query["fee_type_name"] = feeTypeName })
}

// WithReceivables receivables获取 应收费用
func (obj *_FaCustomerBillFeeDetailsMgr) WithReceivables(receivables float64) Option {
	return optionFunc(func(o *options) { o.query["receivables"] = receivables })
}

// WithCurrencyCode currency_code获取 币种
func (obj *_FaCustomerBillFeeDetailsMgr) WithCurrencyCode(currencyCode string) Option {
	return optionFunc(func(o *options) { o.query["currency_code"] = currencyCode })
}

// WithCurrencyName currency_name获取 币种中文名称
func (obj *_FaCustomerBillFeeDetailsMgr) WithCurrencyName(currencyName string) Option {
	return optionFunc(func(o *options) { o.query["currency_name"] = currencyName })
}

// WithExchangeRate exchange_rate获取 汇率
func (obj *_FaCustomerBillFeeDetailsMgr) WithExchangeRate(exchangeRate float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_rate"] = exchangeRate })
}

// WithReceivablesRmb receivables_rmb获取 应收费用(RMB)
func (obj *_FaCustomerBillFeeDetailsMgr) WithReceivablesRmb(receivablesRmb float64) Option {
	return optionFunc(func(o *options) { o.query["receivables_rmb"] = receivablesRmb })
}

// WithFeeSource fee_source获取 费用来源  0:系统生成  1:手工添加
func (obj *_FaCustomerBillFeeDetailsMgr) WithFeeSource(feeSource int) Option {
	return optionFunc(func(o *options) { o.query["fee_source"] = feeSource })
}

// WithBillBatchNumber bill_batch_number获取 对账单批次号
func (obj *_FaCustomerBillFeeDetailsMgr) WithBillBatchNumber(billBatchNumber string) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number"] = billBatchNumber })
}

// WithBillBatchNumberUser bill_batch_number_user获取 生成对账用户
func (obj *_FaCustomerBillFeeDetailsMgr) WithBillBatchNumberUser(billBatchNumberUser string) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number_user"] = billBatchNumberUser })
}

// WithBillBatchNumberUserID bill_batch_number_user_id获取 生成对账用户
func (obj *_FaCustomerBillFeeDetailsMgr) WithBillBatchNumberUserID(billBatchNumberUserID int) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number_user_id"] = billBatchNumberUserID })
}

// WithBillBatchNumberTime bill_batch_number_time获取 生成对账时间
func (obj *_FaCustomerBillFeeDetailsMgr) WithBillBatchNumberTime(billBatchNumberTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number_time"] = billBatchNumberTime })
}

// WithIsGenerateBillBatchNumber is_generate_bill_batch_number获取 已生成对账单
func (obj *_FaCustomerBillFeeDetailsMgr) WithIsGenerateBillBatchNumber(isGenerateBillBatchNumber []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_generate_bill_batch_number"] = isGenerateBillBatchNumber })
}

// WithIsInvalid is_invalid获取 作废 0:正常  1:作废
func (obj *_FaCustomerBillFeeDetailsMgr) WithIsInvalid(isInvalid []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_invalid"] = isInvalid })
}

// WithIsAudit is_audit获取 审核状态(关联对账单批次号状态)
func (obj *_FaCustomerBillFeeDetailsMgr) WithIsAudit(isAudit []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_audit"] = isAudit })
}

// WithAuditUser audit_user获取 审核用户
func (obj *_FaCustomerBillFeeDetailsMgr) WithAuditUser(auditUser string) Option {
	return optionFunc(func(o *options) { o.query["audit_user"] = auditUser })
}

// WithAuditUserID audit_user_id获取 审核用户ID
func (obj *_FaCustomerBillFeeDetailsMgr) WithAuditUserID(auditUserID int) Option {
	return optionFunc(func(o *options) { o.query["audit_user_id"] = auditUserID })
}

// WithAuditTime audit_time获取 审核时间
func (obj *_FaCustomerBillFeeDetailsMgr) WithAuditTime(auditTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["audit_time"] = auditTime })
}

// WithIsAccept is_accept获取 已核收(关联收款单状态)
func (obj *_FaCustomerBillFeeDetailsMgr) WithIsAccept(isAccept []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_accept"] = isAccept })
}

// WithReceiptNumber receipt_number获取 收款单号
func (obj *_FaCustomerBillFeeDetailsMgr) WithReceiptNumber(receiptNumber string) Option {
	return optionFunc(func(o *options) { o.query["receipt_number"] = receiptNumber })
}

// WithAcceptUser accept_user获取 核收用户
func (obj *_FaCustomerBillFeeDetailsMgr) WithAcceptUser(acceptUser string) Option {
	return optionFunc(func(o *options) { o.query["accept_user"] = acceptUser })
}

// WithAcceptUserID accept_user_id获取 核收人员id
func (obj *_FaCustomerBillFeeDetailsMgr) WithAcceptUserID(acceptUserID int) Option {
	return optionFunc(func(o *options) { o.query["accept_user_id"] = acceptUserID })
}

// WithAcceptTime accept_time获取 核收时间
func (obj *_FaCustomerBillFeeDetailsMgr) WithAcceptTime(acceptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["accept_time"] = acceptTime })
}

// WithIsSyncedOutCumulativeFee is_synced_out_cumulative_fee获取 是否同步了出库累加金额
func (obj *_FaCustomerBillFeeDetailsMgr) WithIsSyncedOutCumulativeFee(isSyncedOutCumulativeFee []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_synced_out_cumulative_fee"] = isSyncedOutCumulativeFee })
}

// WithRemark remark获取 备注
func (obj *_FaCustomerBillFeeDetailsMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithFinancialRemark financial_remark获取 财务备注
func (obj *_FaCustomerBillFeeDetailsMgr) WithFinancialRemark(financialRemark string) Option {
	return optionFunc(func(o *options) { o.query["financial_remark"] = financialRemark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaCustomerBillFeeDetailsMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应收人员
func (obj *_FaCustomerBillFeeDetailsMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaCustomerBillFeeDetailsMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaCustomerBillFeeDetailsMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaCustomerBillFeeDetailsMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithIsDiffAmount is_diff_amount获取 是否差额
func (obj *_FaCustomerBillFeeDetailsMgr) WithIsDiffAmount(isDiffAmount []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_diff_amount"] = isDiffAmount })
}

// GetByOption 功能选项模式获取
func (obj *_FaCustomerBillFeeDetailsMgr) GetByOption(opts ...Option) (result FaCustomerBillFeeDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaCustomerBillFeeDetailsMgr) GetByOptions(opts ...Option) (results []*FaCustomerBillFeeDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaCustomerBillFeeDetailsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaCustomerBillFeeDetails, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where(options.query)
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
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromID(id int64) (result FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromID(ids []int64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromOrderNumber(orderNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromReferenceNumber(referenceNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromCustomerID(customerID int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromCustomerAlias(customerAlias string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromCustomerChannelID(customerChannelID int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromCustomerChannelName(customerChannelName string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromReceiptTime 通过receipt_time获取内容 入库时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromReceiptTime(receiptTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// GetBatchFromReceiptTime 批量查找 入库时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromReceiptTime(receiptTimes []time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receipt_time` IN (?)", receiptTimes).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 出库时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromSendTime(sendTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 出库时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromIsSend 通过is_send获取内容 是否出库
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromIsSend(isSend []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_send` = ?", isSend).Find(&results).Error

	return
}

// GetBatchFromIsSend 批量查找 是否出库
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromIsSend(isSends [][]uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_send` IN (?)", isSends).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 预报重量
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromWeight(weight float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 预报重量
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromWeight(weights []float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromWeighingWeight 通过weighing_weight获取内容 称重重量
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromWeighingWeight(weighingWeight float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`weighing_weight` = ?", weighingWeight).Find(&results).Error

	return
}

// GetBatchFromWeighingWeight 批量查找 称重重量
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromWeighingWeight(weighingWeights []float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`weighing_weight` IN (?)", weighingWeights).Find(&results).Error

	return
}

// GetFromFoamWeight 通过foam_weight获取内容 泡重
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromFoamWeight(foamWeight float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`foam_weight` = ?", foamWeight).Find(&results).Error

	return
}

// GetBatchFromFoamWeight 批量查找 泡重
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromFoamWeight(foamWeights []float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`foam_weight` IN (?)", foamWeights).Find(&results).Error

	return
}

// GetFromChargeWeight 通过charge_weight获取内容 计费重量
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromChargeWeight(chargeWeight float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`charge_weight` = ?", chargeWeight).Find(&results).Error

	return
}

// GetBatchFromChargeWeight 批量查找 计费重量
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromChargeWeight(chargeWeights []float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`charge_weight` IN (?)", chargeWeights).Find(&results).Error

	return
}

// GetFromFeeTypeName 通过fee_type_name获取内容 费用类型名称
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromFeeTypeName(feeTypeName string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`fee_type_name` = ?", feeTypeName).Find(&results).Error

	return
}

// GetBatchFromFeeTypeName 批量查找 费用类型名称
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromFeeTypeName(feeTypeNames []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`fee_type_name` IN (?)", feeTypeNames).Find(&results).Error

	return
}

// GetFromReceivables 通过receivables获取内容 应收费用
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromReceivables(receivables float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receivables` = ?", receivables).Find(&results).Error

	return
}

// GetBatchFromReceivables 批量查找 应收费用
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromReceivables(receivabless []float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receivables` IN (?)", receivabless).Find(&results).Error

	return
}

// GetFromCurrencyCode 通过currency_code获取内容 币种
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromCurrencyCode(currencyCode string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`currency_code` = ?", currencyCode).Find(&results).Error

	return
}

// GetBatchFromCurrencyCode 批量查找 币种
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromCurrencyCode(currencyCodes []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`currency_code` IN (?)", currencyCodes).Find(&results).Error

	return
}

// GetFromCurrencyName 通过currency_name获取内容 币种中文名称
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromCurrencyName(currencyName string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`currency_name` = ?", currencyName).Find(&results).Error

	return
}

// GetBatchFromCurrencyName 批量查找 币种中文名称
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromCurrencyName(currencyNames []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`currency_name` IN (?)", currencyNames).Find(&results).Error

	return
}

// GetFromExchangeRate 通过exchange_rate获取内容 汇率
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromExchangeRate(exchangeRate float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`exchange_rate` = ?", exchangeRate).Find(&results).Error

	return
}

// GetBatchFromExchangeRate 批量查找 汇率
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromExchangeRate(exchangeRates []float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`exchange_rate` IN (?)", exchangeRates).Find(&results).Error

	return
}

// GetFromReceivablesRmb 通过receivables_rmb获取内容 应收费用(RMB)
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromReceivablesRmb(receivablesRmb float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receivables_rmb` = ?", receivablesRmb).Find(&results).Error

	return
}

// GetBatchFromReceivablesRmb 批量查找 应收费用(RMB)
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromReceivablesRmb(receivablesRmbs []float64) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receivables_rmb` IN (?)", receivablesRmbs).Find(&results).Error

	return
}

// GetFromFeeSource 通过fee_source获取内容 费用来源  0:系统生成  1:手工添加
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromFeeSource(feeSource int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`fee_source` = ?", feeSource).Find(&results).Error

	return
}

// GetBatchFromFeeSource 批量查找 费用来源  0:系统生成  1:手工添加
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromFeeSource(feeSources []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`fee_source` IN (?)", feeSources).Find(&results).Error

	return
}

// GetFromBillBatchNumber 通过bill_batch_number获取内容 对账单批次号
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromBillBatchNumber(billBatchNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumber 批量查找 对账单批次号
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromBillBatchNumber(billBatchNumbers []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number` IN (?)", billBatchNumbers).Find(&results).Error

	return
}

// GetFromBillBatchNumberUser 通过bill_batch_number_user获取内容 生成对账用户
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromBillBatchNumberUser(billBatchNumberUser string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number_user` = ?", billBatchNumberUser).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumberUser 批量查找 生成对账用户
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromBillBatchNumberUser(billBatchNumberUsers []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number_user` IN (?)", billBatchNumberUsers).Find(&results).Error

	return
}

// GetFromBillBatchNumberUserID 通过bill_batch_number_user_id获取内容 生成对账用户
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromBillBatchNumberUserID(billBatchNumberUserID int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number_user_id` = ?", billBatchNumberUserID).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumberUserID 批量查找 生成对账用户
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromBillBatchNumberUserID(billBatchNumberUserIDs []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number_user_id` IN (?)", billBatchNumberUserIDs).Find(&results).Error

	return
}

// GetFromBillBatchNumberTime 通过bill_batch_number_time获取内容 生成对账时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromBillBatchNumberTime(billBatchNumberTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number_time` = ?", billBatchNumberTime).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumberTime 批量查找 生成对账时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromBillBatchNumberTime(billBatchNumberTimes []time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number_time` IN (?)", billBatchNumberTimes).Find(&results).Error

	return
}

// GetFromIsGenerateBillBatchNumber 通过is_generate_bill_batch_number获取内容 已生成对账单
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromIsGenerateBillBatchNumber(isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_generate_bill_batch_number` = ?", isGenerateBillBatchNumber).Find(&results).Error

	return
}

// GetBatchFromIsGenerateBillBatchNumber 批量查找 已生成对账单
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromIsGenerateBillBatchNumber(isGenerateBillBatchNumbers [][]uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_generate_bill_batch_number` IN (?)", isGenerateBillBatchNumbers).Find(&results).Error

	return
}

// GetFromIsInvalid 通过is_invalid获取内容 作废 0:正常  1:作废
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromIsInvalid(isInvalid []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_invalid` = ?", isInvalid).Find(&results).Error

	return
}

// GetBatchFromIsInvalid 批量查找 作废 0:正常  1:作废
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromIsInvalid(isInvalids [][]uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_invalid` IN (?)", isInvalids).Find(&results).Error

	return
}

// GetFromIsAudit 通过is_audit获取内容 审核状态(关联对账单批次号状态)
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromIsAudit(isAudit []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_audit` = ?", isAudit).Find(&results).Error

	return
}

// GetBatchFromIsAudit 批量查找 审核状态(关联对账单批次号状态)
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromIsAudit(isAudits [][]uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_audit` IN (?)", isAudits).Find(&results).Error

	return
}

// GetFromAuditUser 通过audit_user获取内容 审核用户
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromAuditUser(auditUser string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`audit_user` = ?", auditUser).Find(&results).Error

	return
}

// GetBatchFromAuditUser 批量查找 审核用户
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromAuditUser(auditUsers []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`audit_user` IN (?)", auditUsers).Find(&results).Error

	return
}

// GetFromAuditUserID 通过audit_user_id获取内容 审核用户ID
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromAuditUserID(auditUserID int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`audit_user_id` = ?", auditUserID).Find(&results).Error

	return
}

// GetBatchFromAuditUserID 批量查找 审核用户ID
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromAuditUserID(auditUserIDs []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`audit_user_id` IN (?)", auditUserIDs).Find(&results).Error

	return
}

// GetFromAuditTime 通过audit_time获取内容 审核时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromAuditTime(auditTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`audit_time` = ?", auditTime).Find(&results).Error

	return
}

// GetBatchFromAuditTime 批量查找 审核时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromAuditTime(auditTimes []time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`audit_time` IN (?)", auditTimes).Find(&results).Error

	return
}

// GetFromIsAccept 通过is_accept获取内容 已核收(关联收款单状态)
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromIsAccept(isAccept []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_accept` = ?", isAccept).Find(&results).Error

	return
}

// GetBatchFromIsAccept 批量查找 已核收(关联收款单状态)
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromIsAccept(isAccepts [][]uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_accept` IN (?)", isAccepts).Find(&results).Error

	return
}

// GetFromReceiptNumber 通过receipt_number获取内容 收款单号
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromReceiptNumber(receiptNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receipt_number` = ?", receiptNumber).Find(&results).Error

	return
}

// GetBatchFromReceiptNumber 批量查找 收款单号
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromReceiptNumber(receiptNumbers []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receipt_number` IN (?)", receiptNumbers).Find(&results).Error

	return
}

// GetFromAcceptUser 通过accept_user获取内容 核收用户
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromAcceptUser(acceptUser string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`accept_user` = ?", acceptUser).Find(&results).Error

	return
}

// GetBatchFromAcceptUser 批量查找 核收用户
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromAcceptUser(acceptUsers []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`accept_user` IN (?)", acceptUsers).Find(&results).Error

	return
}

// GetFromAcceptUserID 通过accept_user_id获取内容 核收人员id
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromAcceptUserID(acceptUserID int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`accept_user_id` = ?", acceptUserID).Find(&results).Error

	return
}

// GetBatchFromAcceptUserID 批量查找 核收人员id
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromAcceptUserID(acceptUserIDs []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`accept_user_id` IN (?)", acceptUserIDs).Find(&results).Error

	return
}

// GetFromAcceptTime 通过accept_time获取内容 核收时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromAcceptTime(acceptTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// GetBatchFromAcceptTime 批量查找 核收时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromAcceptTime(acceptTimes []time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`accept_time` IN (?)", acceptTimes).Find(&results).Error

	return
}

// GetFromIsSyncedOutCumulativeFee 通过is_synced_out_cumulative_fee获取内容 是否同步了出库累加金额
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromIsSyncedOutCumulativeFee(isSyncedOutCumulativeFee []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_synced_out_cumulative_fee` = ?", isSyncedOutCumulativeFee).Find(&results).Error

	return
}

// GetBatchFromIsSyncedOutCumulativeFee 批量查找 是否同步了出库累加金额
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromIsSyncedOutCumulativeFee(isSyncedOutCumulativeFees [][]uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_synced_out_cumulative_fee` IN (?)", isSyncedOutCumulativeFees).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromRemark(remark string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromRemark(remarks []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromFinancialRemark 通过financial_remark获取内容 财务备注
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromFinancialRemark(financialRemark string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`financial_remark` = ?", financialRemark).Find(&results).Error

	return
}

// GetBatchFromFinancialRemark 批量查找 财务备注
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromFinancialRemark(financialRemarks []string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`financial_remark` IN (?)", financialRemarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromCreateTime(createTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应收人员
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromCreateUser(createUser int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应收人员
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromUpdateUser(updateUser int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromVersion(version int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromVersion(versions []int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromIsDiffAmount 通过is_diff_amount获取内容 是否差额
func (obj *_FaCustomerBillFeeDetailsMgr) GetFromIsDiffAmount(isDiffAmount []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_diff_amount` = ?", isDiffAmount).Find(&results).Error

	return
}

// GetBatchFromIsDiffAmount 批量查找 是否差额
func (obj *_FaCustomerBillFeeDetailsMgr) GetBatchFromIsDiffAmount(isDiffAmounts [][]uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_diff_amount` IN (?)", isDiffAmounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchByPrimaryKey(id int64) (result FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIndexOrderNumber  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexOrderNumber(orderNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByIndexReferenceNumber  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexReferenceNumber(referenceNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByIndexReferenceNumberIsGenerateBillBatchNumber  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexReferenceNumberIsGenerateBillBatchNumber(referenceNumber string, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`reference_number` = ? AND `is_generate_bill_batch_number` = ?", referenceNumber, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumber  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexLogisticsNumber(logisticsNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumberFinal  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerID  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerID(customerID int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerIDReceiptTimeIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerIDReceiptTimeIsGen(customerID int, receiptTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_id` = ? AND `receipt_time` = ? AND `is_generate_bill_batch_number` = ?", customerID, receiptTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerIDSendTimeIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerIDSendTimeIsGen(customerID int, sendTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_id` = ? AND `send_time` = ? AND `is_generate_bill_batch_number` = ?", customerID, sendTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerIDIsSendIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerIDIsSendIsGen(customerID int, isSend []uint8, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_id` = ? AND `is_send` = ? AND `is_generate_bill_batch_number` = ?", customerID, isSend, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerIDIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerIDIsGen(customerID int, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_id` = ? AND `is_generate_bill_batch_number` = ?", customerID, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerIDCreatedTimeIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerIDCreatedTimeIsGen(customerID int, isGenerateBillBatchNumber []uint8, createTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_id` = ? AND `is_generate_bill_batch_number` = ? AND `create_time` = ?", customerID, isGenerateBillBatchNumber, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelID  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerChannelID(customerChannelID int) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelIDReceiptTimeIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerChannelIDReceiptTimeIsGen(customerChannelID int, receiptTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_id` = ? AND `receipt_time` = ? AND `is_generate_bill_batch_number` = ?", customerChannelID, receiptTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelIDSendTimeIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerChannelIDSendTimeIsGen(customerChannelID int, sendTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_id` = ? AND `send_time` = ? AND `is_generate_bill_batch_number` = ?", customerChannelID, sendTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelIDIsSendIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerChannelIDIsSendIsGen(customerChannelID int, isSend []uint8, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_id` = ? AND `is_send` = ? AND `is_generate_bill_batch_number` = ?", customerChannelID, isSend, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelIDIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerChannelIDIsGen(customerChannelID int, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_id` = ? AND `is_generate_bill_batch_number` = ?", customerChannelID, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelIDCreatedTimeIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCustomerChannelIDCreatedTimeIsGen(customerChannelID int, isGenerateBillBatchNumber []uint8, createTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`customer_channel_id` = ? AND `is_generate_bill_batch_number` = ? AND `create_time` = ?", customerChannelID, isGenerateBillBatchNumber, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexReceiptTime  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexReceiptTime(receiptTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// FetchIndexByIndexReceiptTimeIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexReceiptTimeIsGen(receiptTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receipt_time` = ? AND `is_generate_bill_batch_number` = ?", receiptTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIsGenerateBillBatchNumberReceiptTimeIndex  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIsGenerateBillBatchNumberReceiptTimeIndex(receiptTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receipt_time` = ? AND `is_generate_bill_batch_number` = ?", receiptTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexSendTime  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexSendTime(sendTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// FetchIndexByIndexSendTimeIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexSendTimeIsGen(sendTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`send_time` = ? AND `is_generate_bill_batch_number` = ?", sendTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexIsSendIsIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexIsSendIsIsGen(isSend []uint8, isGenerateBillBatchNumber []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_send` = ? AND `is_generate_bill_batch_number` = ?", isSend, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexBillBatchNumber  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexBillBatchNumber(billBatchNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexCreatedIsGen  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexCreatedIsGen(isGenerateBillBatchNumber []uint8, createTime time.Time) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_generate_bill_batch_number` = ? AND `create_time` = ?", isGenerateBillBatchNumber, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexAcceptNumber  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexAcceptNumber(receiptNumber string) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`receipt_number` = ?", receiptNumber).Find(&results).Error

	return
}

// FetchIndexByIndexIsSyncedOutCumulativeFee  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexIsSyncedOutCumulativeFee(isSyncedOutCumulativeFee []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_synced_out_cumulative_fee` = ?", isSyncedOutCumulativeFee).Find(&results).Error

	return
}

// FetchIndexByIndexIsDiffAmount  获取多个内容
func (obj *_FaCustomerBillFeeDetailsMgr) FetchIndexByIndexIsDiffAmount(isDiffAmount []uint8) (results []*FaCustomerBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFeeDetails{}).Where("`is_diff_amount` = ?", isDiffAmount).Find(&results).Error

	return
}
