package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaPayableBillFeeDetailsMgr struct {
	*_BaseMgr
}

// FaPayableBillFeeDetailsMgr open func
func FaPayableBillFeeDetailsMgr(db *gorm.DB) *_FaPayableBillFeeDetailsMgr {
	if db == nil {
		panic(fmt.Errorf("FaPayableBillFeeDetailsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaPayableBillFeeDetailsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_payable_bill_fee_details"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaPayableBillFeeDetailsMgr) GetTableName() string {
	return "fa_payable_bill_fee_details"
}

// Reset 重置gorm会话
func (obj *_FaPayableBillFeeDetailsMgr) Reset() *_FaPayableBillFeeDetailsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaPayableBillFeeDetailsMgr) Get() (result FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaPayableBillFeeDetailsMgr) Gets() (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaPayableBillFeeDetailsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaPayableBillFeeDetailsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_FaPayableBillFeeDetailsMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_FaPayableBillFeeDetailsMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_FaPayableBillFeeDetailsMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_FaPayableBillFeeDetailsMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithBusinessEntityCode business_entity_code获取 业务主体代码
func (obj *_FaPayableBillFeeDetailsMgr) WithBusinessEntityCode(businessEntityCode string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_code"] = businessEntityCode })
}

// WithBusinessEntityName business_entity_name获取 业务主体名称
func (obj *_FaPayableBillFeeDetailsMgr) WithBusinessEntityName(businessEntityName string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_name"] = businessEntityName })
}

// WithProviderChannelCode provider_channel_code获取 服务商渠道代码
func (obj *_FaPayableBillFeeDetailsMgr) WithProviderChannelCode(providerChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_code"] = providerChannelCode })
}

// WithBusinessEntityType business_entity_type获取 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillFeeDetailsMgr) WithBusinessEntityType(businessEntityType bool) Option {
	return optionFunc(func(o *options) { o.query["business_entity_type"] = businessEntityType })
}

// WithReceiptTime receipt_time获取 入库时间
func (obj *_FaPayableBillFeeDetailsMgr) WithReceiptTime(receiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_time"] = receiptTime })
}

// WithSendTime send_time获取 出库时间
func (obj *_FaPayableBillFeeDetailsMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithIsSend is_send获取 是否出库
func (obj *_FaPayableBillFeeDetailsMgr) WithIsSend(isSend []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_send"] = isSend })
}

// WithWeight weight获取 预报重量
func (obj *_FaPayableBillFeeDetailsMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithWeighingWeight weighing_weight获取 称重重量
func (obj *_FaPayableBillFeeDetailsMgr) WithWeighingWeight(weighingWeight float64) Option {
	return optionFunc(func(o *options) { o.query["weighing_weight"] = weighingWeight })
}

// WithFoamWeight foam_weight获取 泡重
func (obj *_FaPayableBillFeeDetailsMgr) WithFoamWeight(foamWeight float64) Option {
	return optionFunc(func(o *options) { o.query["foam_weight"] = foamWeight })
}

// WithChargeWeight charge_weight获取 计费重量
func (obj *_FaPayableBillFeeDetailsMgr) WithChargeWeight(chargeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["charge_weight"] = chargeWeight })
}

// WithFeeTypeName fee_type_name获取 费用类型名称
func (obj *_FaPayableBillFeeDetailsMgr) WithFeeTypeName(feeTypeName string) Option {
	return optionFunc(func(o *options) { o.query["fee_type_name"] = feeTypeName })
}

// WithPayable payable获取 应付费用
func (obj *_FaPayableBillFeeDetailsMgr) WithPayable(payable float64) Option {
	return optionFunc(func(o *options) { o.query["payable"] = payable })
}

// WithCurrencyCode currency_code获取 币种
func (obj *_FaPayableBillFeeDetailsMgr) WithCurrencyCode(currencyCode string) Option {
	return optionFunc(func(o *options) { o.query["currency_code"] = currencyCode })
}

// WithCurrencyName currency_name获取 币种中文名称
func (obj *_FaPayableBillFeeDetailsMgr) WithCurrencyName(currencyName string) Option {
	return optionFunc(func(o *options) { o.query["currency_name"] = currencyName })
}

// WithExchangeRate exchange_rate获取 汇率
func (obj *_FaPayableBillFeeDetailsMgr) WithExchangeRate(exchangeRate float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_rate"] = exchangeRate })
}

// WithPayableRmb payable_rmb获取 应付费用(RMB),折扣前
func (obj *_FaPayableBillFeeDetailsMgr) WithPayableRmb(payableRmb float64) Option {
	return optionFunc(func(o *options) { o.query["payable_rmb"] = payableRmb })
}

// WithFeeSource fee_source获取 费用来源  0:系统生成  1:手工添加
func (obj *_FaPayableBillFeeDetailsMgr) WithFeeSource(feeSource int) Option {
	return optionFunc(func(o *options) { o.query["fee_source"] = feeSource })
}

// WithBillBatchNumber bill_batch_number获取 对账单批次号
func (obj *_FaPayableBillFeeDetailsMgr) WithBillBatchNumber(billBatchNumber string) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number"] = billBatchNumber })
}

// WithBillBatchNumberUser bill_batch_number_user获取 生成对账用户
func (obj *_FaPayableBillFeeDetailsMgr) WithBillBatchNumberUser(billBatchNumberUser string) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number_user"] = billBatchNumberUser })
}

// WithBillBatchNumberUserID bill_batch_number_user_id获取 生成对账用户
func (obj *_FaPayableBillFeeDetailsMgr) WithBillBatchNumberUserID(billBatchNumberUserID int) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number_user_id"] = billBatchNumberUserID })
}

// WithBillBatchNumberTime bill_batch_number_time获取 生成对账时间
func (obj *_FaPayableBillFeeDetailsMgr) WithBillBatchNumberTime(billBatchNumberTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_number_time"] = billBatchNumberTime })
}

// WithIsGenerateBillBatchNumber is_generate_bill_batch_number获取 已生成对账单
func (obj *_FaPayableBillFeeDetailsMgr) WithIsGenerateBillBatchNumber(isGenerateBillBatchNumber []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_generate_bill_batch_number"] = isGenerateBillBatchNumber })
}

// WithIsInvalid is_invalid获取 作废 0:正常  1:作废
func (obj *_FaPayableBillFeeDetailsMgr) WithIsInvalid(isInvalid []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_invalid"] = isInvalid })
}

// WithIsAudit is_audit获取 审核状态(关联对账单批次号状态)
func (obj *_FaPayableBillFeeDetailsMgr) WithIsAudit(isAudit []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_audit"] = isAudit })
}

// WithAuditUser audit_user获取 审核用户
func (obj *_FaPayableBillFeeDetailsMgr) WithAuditUser(auditUser string) Option {
	return optionFunc(func(o *options) { o.query["audit_user"] = auditUser })
}

// WithAuditUserID audit_user_id获取 审核用户ID
func (obj *_FaPayableBillFeeDetailsMgr) WithAuditUserID(auditUserID int) Option {
	return optionFunc(func(o *options) { o.query["audit_user_id"] = auditUserID })
}

// WithAuditTime audit_time获取 审核时间
func (obj *_FaPayableBillFeeDetailsMgr) WithAuditTime(auditTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["audit_time"] = auditTime })
}

// WithIsAccept is_accept获取 已核付(关联付款单状态)
func (obj *_FaPayableBillFeeDetailsMgr) WithIsAccept(isAccept []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_accept"] = isAccept })
}

// WithPaymentNumber payment_number获取 付款单号
func (obj *_FaPayableBillFeeDetailsMgr) WithPaymentNumber(paymentNumber string) Option {
	return optionFunc(func(o *options) { o.query["payment_number"] = paymentNumber })
}

// WithAcceptUser accept_user获取 核付用户
func (obj *_FaPayableBillFeeDetailsMgr) WithAcceptUser(acceptUser string) Option {
	return optionFunc(func(o *options) { o.query["accept_user"] = acceptUser })
}

// WithAcceptUserID accept_user_id获取 核付人员id
func (obj *_FaPayableBillFeeDetailsMgr) WithAcceptUserID(acceptUserID int) Option {
	return optionFunc(func(o *options) { o.query["accept_user_id"] = acceptUserID })
}

// WithAcceptTime accept_time获取 核付时间
func (obj *_FaPayableBillFeeDetailsMgr) WithAcceptTime(acceptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["accept_time"] = acceptTime })
}

// WithRemark remark获取 备注
func (obj *_FaPayableBillFeeDetailsMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithFinancialRemark financial_remark获取 财务备注
func (obj *_FaPayableBillFeeDetailsMgr) WithFinancialRemark(financialRemark string) Option {
	return optionFunc(func(o *options) { o.query["financial_remark"] = financialRemark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaPayableBillFeeDetailsMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应付人员
func (obj *_FaPayableBillFeeDetailsMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaPayableBillFeeDetailsMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaPayableBillFeeDetailsMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaPayableBillFeeDetailsMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_FaPayableBillFeeDetailsMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_FaPayableBillFeeDetailsMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithDiscountRate discount_rate获取 折扣率(小数),默认1 原币种
func (obj *_FaPayableBillFeeDetailsMgr) WithDiscountRate(discountRate float64) Option {
	return optionFunc(func(o *options) { o.query["discount_rate"] = discountRate })
}

// WithActualPayable actual_payable获取 实际原币种金额
func (obj *_FaPayableBillFeeDetailsMgr) WithActualPayable(actualPayable float64) Option {
	return optionFunc(func(o *options) { o.query["actual_payable"] = actualPayable })
}

// WithIsDiffAmount is_diff_amount获取 是否差额
func (obj *_FaPayableBillFeeDetailsMgr) WithIsDiffAmount(isDiffAmount []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_diff_amount"] = isDiffAmount })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_FaPayableBillFeeDetailsMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_FaPayableBillFeeDetailsMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// GetByOption 功能选项模式获取
func (obj *_FaPayableBillFeeDetailsMgr) GetByOption(opts ...Option) (result FaPayableBillFeeDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaPayableBillFeeDetailsMgr) GetByOptions(opts ...Option) (results []*FaPayableBillFeeDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaPayableBillFeeDetailsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaPayableBillFeeDetails, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where(options.query)
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
func (obj *_FaPayableBillFeeDetailsMgr) GetFromID(id int64) (result FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromID(ids []int64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_FaPayableBillFeeDetailsMgr) GetFromOrderNumber(orderNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_FaPayableBillFeeDetailsMgr) GetFromReferenceNumber(referenceNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_FaPayableBillFeeDetailsMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_FaPayableBillFeeDetailsMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromBusinessEntityCode 通过business_entity_code获取内容 业务主体代码
func (obj *_FaPayableBillFeeDetailsMgr) GetFromBusinessEntityCode(businessEntityCode string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityCode 批量查找 业务主体代码
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromBusinessEntityCode(businessEntityCodes []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`business_entity_code` IN (?)", businessEntityCodes).Find(&results).Error

	return
}

// GetFromBusinessEntityName 通过business_entity_name获取内容 业务主体名称
func (obj *_FaPayableBillFeeDetailsMgr) GetFromBusinessEntityName(businessEntityName string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`business_entity_name` = ?", businessEntityName).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityName 批量查找 业务主体名称
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromBusinessEntityName(businessEntityNames []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`business_entity_name` IN (?)", businessEntityNames).Find(&results).Error

	return
}

// GetFromProviderChannelCode 通过provider_channel_code获取内容 服务商渠道代码
func (obj *_FaPayableBillFeeDetailsMgr) GetFromProviderChannelCode(providerChannelCode string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error

	return
}

// GetBatchFromProviderChannelCode 批量查找 服务商渠道代码
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromProviderChannelCode(providerChannelCodes []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`provider_channel_code` IN (?)", providerChannelCodes).Find(&results).Error

	return
}

// GetFromBusinessEntityType 通过business_entity_type获取内容 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillFeeDetailsMgr) GetFromBusinessEntityType(businessEntityType bool) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`business_entity_type` = ?", businessEntityType).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityType 批量查找 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromBusinessEntityType(businessEntityTypes []bool) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`business_entity_type` IN (?)", businessEntityTypes).Find(&results).Error

	return
}

// GetFromReceiptTime 通过receipt_time获取内容 入库时间
func (obj *_FaPayableBillFeeDetailsMgr) GetFromReceiptTime(receiptTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// GetBatchFromReceiptTime 批量查找 入库时间
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromReceiptTime(receiptTimes []time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`receipt_time` IN (?)", receiptTimes).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 出库时间
func (obj *_FaPayableBillFeeDetailsMgr) GetFromSendTime(sendTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 出库时间
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromIsSend 通过is_send获取内容 是否出库
func (obj *_FaPayableBillFeeDetailsMgr) GetFromIsSend(isSend []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_send` = ?", isSend).Find(&results).Error

	return
}

// GetBatchFromIsSend 批量查找 是否出库
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromIsSend(isSends [][]uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_send` IN (?)", isSends).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 预报重量
func (obj *_FaPayableBillFeeDetailsMgr) GetFromWeight(weight float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 预报重量
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromWeight(weights []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromWeighingWeight 通过weighing_weight获取内容 称重重量
func (obj *_FaPayableBillFeeDetailsMgr) GetFromWeighingWeight(weighingWeight float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`weighing_weight` = ?", weighingWeight).Find(&results).Error

	return
}

// GetBatchFromWeighingWeight 批量查找 称重重量
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromWeighingWeight(weighingWeights []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`weighing_weight` IN (?)", weighingWeights).Find(&results).Error

	return
}

// GetFromFoamWeight 通过foam_weight获取内容 泡重
func (obj *_FaPayableBillFeeDetailsMgr) GetFromFoamWeight(foamWeight float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`foam_weight` = ?", foamWeight).Find(&results).Error

	return
}

// GetBatchFromFoamWeight 批量查找 泡重
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromFoamWeight(foamWeights []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`foam_weight` IN (?)", foamWeights).Find(&results).Error

	return
}

// GetFromChargeWeight 通过charge_weight获取内容 计费重量
func (obj *_FaPayableBillFeeDetailsMgr) GetFromChargeWeight(chargeWeight float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`charge_weight` = ?", chargeWeight).Find(&results).Error

	return
}

// GetBatchFromChargeWeight 批量查找 计费重量
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromChargeWeight(chargeWeights []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`charge_weight` IN (?)", chargeWeights).Find(&results).Error

	return
}

// GetFromFeeTypeName 通过fee_type_name获取内容 费用类型名称
func (obj *_FaPayableBillFeeDetailsMgr) GetFromFeeTypeName(feeTypeName string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`fee_type_name` = ?", feeTypeName).Find(&results).Error

	return
}

// GetBatchFromFeeTypeName 批量查找 费用类型名称
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromFeeTypeName(feeTypeNames []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`fee_type_name` IN (?)", feeTypeNames).Find(&results).Error

	return
}

// GetFromPayable 通过payable获取内容 应付费用
func (obj *_FaPayableBillFeeDetailsMgr) GetFromPayable(payable float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`payable` = ?", payable).Find(&results).Error

	return
}

// GetBatchFromPayable 批量查找 应付费用
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromPayable(payables []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`payable` IN (?)", payables).Find(&results).Error

	return
}

// GetFromCurrencyCode 通过currency_code获取内容 币种
func (obj *_FaPayableBillFeeDetailsMgr) GetFromCurrencyCode(currencyCode string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`currency_code` = ?", currencyCode).Find(&results).Error

	return
}

// GetBatchFromCurrencyCode 批量查找 币种
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromCurrencyCode(currencyCodes []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`currency_code` IN (?)", currencyCodes).Find(&results).Error

	return
}

// GetFromCurrencyName 通过currency_name获取内容 币种中文名称
func (obj *_FaPayableBillFeeDetailsMgr) GetFromCurrencyName(currencyName string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`currency_name` = ?", currencyName).Find(&results).Error

	return
}

// GetBatchFromCurrencyName 批量查找 币种中文名称
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromCurrencyName(currencyNames []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`currency_name` IN (?)", currencyNames).Find(&results).Error

	return
}

// GetFromExchangeRate 通过exchange_rate获取内容 汇率
func (obj *_FaPayableBillFeeDetailsMgr) GetFromExchangeRate(exchangeRate float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`exchange_rate` = ?", exchangeRate).Find(&results).Error

	return
}

// GetBatchFromExchangeRate 批量查找 汇率
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromExchangeRate(exchangeRates []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`exchange_rate` IN (?)", exchangeRates).Find(&results).Error

	return
}

// GetFromPayableRmb 通过payable_rmb获取内容 应付费用(RMB),折扣前
func (obj *_FaPayableBillFeeDetailsMgr) GetFromPayableRmb(payableRmb float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`payable_rmb` = ?", payableRmb).Find(&results).Error

	return
}

// GetBatchFromPayableRmb 批量查找 应付费用(RMB),折扣前
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromPayableRmb(payableRmbs []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`payable_rmb` IN (?)", payableRmbs).Find(&results).Error

	return
}

// GetFromFeeSource 通过fee_source获取内容 费用来源  0:系统生成  1:手工添加
func (obj *_FaPayableBillFeeDetailsMgr) GetFromFeeSource(feeSource int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`fee_source` = ?", feeSource).Find(&results).Error

	return
}

// GetBatchFromFeeSource 批量查找 费用来源  0:系统生成  1:手工添加
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromFeeSource(feeSources []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`fee_source` IN (?)", feeSources).Find(&results).Error

	return
}

// GetFromBillBatchNumber 通过bill_batch_number获取内容 对账单批次号
func (obj *_FaPayableBillFeeDetailsMgr) GetFromBillBatchNumber(billBatchNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumber 批量查找 对账单批次号
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromBillBatchNumber(billBatchNumbers []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number` IN (?)", billBatchNumbers).Find(&results).Error

	return
}

// GetFromBillBatchNumberUser 通过bill_batch_number_user获取内容 生成对账用户
func (obj *_FaPayableBillFeeDetailsMgr) GetFromBillBatchNumberUser(billBatchNumberUser string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number_user` = ?", billBatchNumberUser).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumberUser 批量查找 生成对账用户
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromBillBatchNumberUser(billBatchNumberUsers []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number_user` IN (?)", billBatchNumberUsers).Find(&results).Error

	return
}

// GetFromBillBatchNumberUserID 通过bill_batch_number_user_id获取内容 生成对账用户
func (obj *_FaPayableBillFeeDetailsMgr) GetFromBillBatchNumberUserID(billBatchNumberUserID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number_user_id` = ?", billBatchNumberUserID).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumberUserID 批量查找 生成对账用户
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromBillBatchNumberUserID(billBatchNumberUserIDs []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number_user_id` IN (?)", billBatchNumberUserIDs).Find(&results).Error

	return
}

// GetFromBillBatchNumberTime 通过bill_batch_number_time获取内容 生成对账时间
func (obj *_FaPayableBillFeeDetailsMgr) GetFromBillBatchNumberTime(billBatchNumberTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number_time` = ?", billBatchNumberTime).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumberTime 批量查找 生成对账时间
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromBillBatchNumberTime(billBatchNumberTimes []time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number_time` IN (?)", billBatchNumberTimes).Find(&results).Error

	return
}

// GetFromIsGenerateBillBatchNumber 通过is_generate_bill_batch_number获取内容 已生成对账单
func (obj *_FaPayableBillFeeDetailsMgr) GetFromIsGenerateBillBatchNumber(isGenerateBillBatchNumber []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_generate_bill_batch_number` = ?", isGenerateBillBatchNumber).Find(&results).Error

	return
}

// GetBatchFromIsGenerateBillBatchNumber 批量查找 已生成对账单
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromIsGenerateBillBatchNumber(isGenerateBillBatchNumbers [][]uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_generate_bill_batch_number` IN (?)", isGenerateBillBatchNumbers).Find(&results).Error

	return
}

// GetFromIsInvalid 通过is_invalid获取内容 作废 0:正常  1:作废
func (obj *_FaPayableBillFeeDetailsMgr) GetFromIsInvalid(isInvalid []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_invalid` = ?", isInvalid).Find(&results).Error

	return
}

// GetBatchFromIsInvalid 批量查找 作废 0:正常  1:作废
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromIsInvalid(isInvalids [][]uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_invalid` IN (?)", isInvalids).Find(&results).Error

	return
}

// GetFromIsAudit 通过is_audit获取内容 审核状态(关联对账单批次号状态)
func (obj *_FaPayableBillFeeDetailsMgr) GetFromIsAudit(isAudit []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_audit` = ?", isAudit).Find(&results).Error

	return
}

// GetBatchFromIsAudit 批量查找 审核状态(关联对账单批次号状态)
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromIsAudit(isAudits [][]uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_audit` IN (?)", isAudits).Find(&results).Error

	return
}

// GetFromAuditUser 通过audit_user获取内容 审核用户
func (obj *_FaPayableBillFeeDetailsMgr) GetFromAuditUser(auditUser string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`audit_user` = ?", auditUser).Find(&results).Error

	return
}

// GetBatchFromAuditUser 批量查找 审核用户
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromAuditUser(auditUsers []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`audit_user` IN (?)", auditUsers).Find(&results).Error

	return
}

// GetFromAuditUserID 通过audit_user_id获取内容 审核用户ID
func (obj *_FaPayableBillFeeDetailsMgr) GetFromAuditUserID(auditUserID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`audit_user_id` = ?", auditUserID).Find(&results).Error

	return
}

// GetBatchFromAuditUserID 批量查找 审核用户ID
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromAuditUserID(auditUserIDs []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`audit_user_id` IN (?)", auditUserIDs).Find(&results).Error

	return
}

// GetFromAuditTime 通过audit_time获取内容 审核时间
func (obj *_FaPayableBillFeeDetailsMgr) GetFromAuditTime(auditTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`audit_time` = ?", auditTime).Find(&results).Error

	return
}

// GetBatchFromAuditTime 批量查找 审核时间
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromAuditTime(auditTimes []time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`audit_time` IN (?)", auditTimes).Find(&results).Error

	return
}

// GetFromIsAccept 通过is_accept获取内容 已核付(关联付款单状态)
func (obj *_FaPayableBillFeeDetailsMgr) GetFromIsAccept(isAccept []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_accept` = ?", isAccept).Find(&results).Error

	return
}

// GetBatchFromIsAccept 批量查找 已核付(关联付款单状态)
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromIsAccept(isAccepts [][]uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_accept` IN (?)", isAccepts).Find(&results).Error

	return
}

// GetFromPaymentNumber 通过payment_number获取内容 付款单号
func (obj *_FaPayableBillFeeDetailsMgr) GetFromPaymentNumber(paymentNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`payment_number` = ?", paymentNumber).Find(&results).Error

	return
}

// GetBatchFromPaymentNumber 批量查找 付款单号
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromPaymentNumber(paymentNumbers []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`payment_number` IN (?)", paymentNumbers).Find(&results).Error

	return
}

// GetFromAcceptUser 通过accept_user获取内容 核付用户
func (obj *_FaPayableBillFeeDetailsMgr) GetFromAcceptUser(acceptUser string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_user` = ?", acceptUser).Find(&results).Error

	return
}

// GetBatchFromAcceptUser 批量查找 核付用户
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromAcceptUser(acceptUsers []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_user` IN (?)", acceptUsers).Find(&results).Error

	return
}

// GetFromAcceptUserID 通过accept_user_id获取内容 核付人员id
func (obj *_FaPayableBillFeeDetailsMgr) GetFromAcceptUserID(acceptUserID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_user_id` = ?", acceptUserID).Find(&results).Error

	return
}

// GetBatchFromAcceptUserID 批量查找 核付人员id
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromAcceptUserID(acceptUserIDs []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_user_id` IN (?)", acceptUserIDs).Find(&results).Error

	return
}

// GetFromAcceptTime 通过accept_time获取内容 核付时间
func (obj *_FaPayableBillFeeDetailsMgr) GetFromAcceptTime(acceptTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// GetBatchFromAcceptTime 批量查找 核付时间
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromAcceptTime(acceptTimes []time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_time` IN (?)", acceptTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaPayableBillFeeDetailsMgr) GetFromRemark(remark string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromRemark(remarks []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromFinancialRemark 通过financial_remark获取内容 财务备注
func (obj *_FaPayableBillFeeDetailsMgr) GetFromFinancialRemark(financialRemark string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`financial_remark` = ?", financialRemark).Find(&results).Error

	return
}

// GetBatchFromFinancialRemark 批量查找 财务备注
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromFinancialRemark(financialRemarks []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`financial_remark` IN (?)", financialRemarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaPayableBillFeeDetailsMgr) GetFromCreateTime(createTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应付人员
func (obj *_FaPayableBillFeeDetailsMgr) GetFromCreateUser(createUser int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应付人员
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaPayableBillFeeDetailsMgr) GetFromUpdateUser(updateUser int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaPayableBillFeeDetailsMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaPayableBillFeeDetailsMgr) GetFromVersion(version int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromVersion(versions []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_FaPayableBillFeeDetailsMgr) GetFromCustomerChannelID(customerChannelID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_FaPayableBillFeeDetailsMgr) GetFromCustomerChannelName(customerChannelName string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromDiscountRate 通过discount_rate获取内容 折扣率(小数),默认1 原币种
func (obj *_FaPayableBillFeeDetailsMgr) GetFromDiscountRate(discountRate float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`discount_rate` = ?", discountRate).Find(&results).Error

	return
}

// GetBatchFromDiscountRate 批量查找 折扣率(小数),默认1 原币种
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromDiscountRate(discountRates []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`discount_rate` IN (?)", discountRates).Find(&results).Error

	return
}

// GetFromActualPayable 通过actual_payable获取内容 实际原币种金额
func (obj *_FaPayableBillFeeDetailsMgr) GetFromActualPayable(actualPayable float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`actual_payable` = ?", actualPayable).Find(&results).Error

	return
}

// GetBatchFromActualPayable 批量查找 实际原币种金额
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromActualPayable(actualPayables []float64) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`actual_payable` IN (?)", actualPayables).Find(&results).Error

	return
}

// GetFromIsDiffAmount 通过is_diff_amount获取内容 是否差额
func (obj *_FaPayableBillFeeDetailsMgr) GetFromIsDiffAmount(isDiffAmount []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_diff_amount` = ?", isDiffAmount).Find(&results).Error

	return
}

// GetBatchFromIsDiffAmount 批量查找 是否差额
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromIsDiffAmount(isDiffAmounts [][]uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_diff_amount` IN (?)", isDiffAmounts).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_FaPayableBillFeeDetailsMgr) GetFromCustomerID(customerID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_FaPayableBillFeeDetailsMgr) GetFromCustomerAlias(customerAlias string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_FaPayableBillFeeDetailsMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchByPrimaryKey(id int64) (result FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIndexOrderNumber  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexOrderNumber(orderNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByIndexReferenceNumber  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexReferenceNumber(referenceNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByIndexReferenceNumberBusinessEntityCode  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexReferenceNumberBusinessEntityCode(referenceNumber string, businessEntityCode string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`reference_number` = ? AND `business_entity_code` = ?", referenceNumber, businessEntityCode).Find(&results).Error

	return
}

// FetchIndexByIndexReferenceNumberIsGenerateBillBatchNumber  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexReferenceNumberIsGenerateBillBatchNumber(referenceNumber string, isGenerateBillBatchNumber []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`reference_number` = ? AND `is_generate_bill_batch_number` = ?", referenceNumber, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumber  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexLogisticsNumber(logisticsNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumberFinal  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsBusinessEntityCodeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsBusinessEntityCodeIndex(businessEntityCode string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsProviderChannelCodeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsProviderChannelCodeIndex(providerChannelCode string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsBusinessEntityTypeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsBusinessEntityTypeIndex(businessEntityType bool) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`business_entity_type` = ?", businessEntityType).Find(&results).Error

	return
}

// FetchIndexByIndexReceiptTimeIsGen  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexReceiptTimeIsGen(receiptTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`receipt_time` = ? AND `is_generate_bill_batch_number` = ?", receiptTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexReceiptTimeIsSendIsGen  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexReceiptTimeIsSendIsGen(receiptTime time.Time, isSend []uint8, isGenerateBillBatchNumber []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`receipt_time` = ? AND `is_send` = ? AND `is_generate_bill_batch_number` = ?", receiptTime, isSend, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsReceiptTimeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsReceiptTimeIndex(receiptTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// FetchIndexByIndexSendTimeIsGen  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexSendTimeIsGen(sendTime time.Time, isGenerateBillBatchNumber []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`send_time` = ? AND `is_generate_bill_batch_number` = ?", sendTime, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByIndexSendTimeIsSendIsGen  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexSendTimeIsSendIsGen(sendTime time.Time, isSend []uint8, isGenerateBillBatchNumber []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`send_time` = ? AND `is_send` = ? AND `is_generate_bill_batch_number` = ?", sendTime, isSend, isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsSendTimeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsSendTimeIndex(sendTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsIsSendIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsIsSendIndex(isSend []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_send` = ?", isSend).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsFeeTypeNameIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsFeeTypeNameIndex(feeTypeName string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`fee_type_name` = ?", feeTypeName).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsCurrencyCodeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsCurrencyCodeIndex(currencyCode string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`currency_code` = ?", currencyCode).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsFeeSourceIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsFeeSourceIndex(feeSource int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`fee_source` = ?", feeSource).Find(&results).Error

	return
}

// FetchIndexByIndexBillBatchNumber  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexBillBatchNumber(billBatchNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsBillBatchNumberIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsBillBatchNumberIndex(billBatchNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number` = ?", billBatchNumber).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsBillBatchNumberUserIDIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsBillBatchNumberUserIDIndex(billBatchNumberUserID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`bill_batch_number_user_id` = ?", billBatchNumberUserID).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsIsGenerateBillBatchNumberIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsIsGenerateBillBatchNumberIndex(isGenerateBillBatchNumber []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_generate_bill_batch_number` = ?", isGenerateBillBatchNumber).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsIsInvalidIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsIsInvalidIndex(isInvalid []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_invalid` = ?", isInvalid).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsIsAuditIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsIsAuditIndex(isAudit []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_audit` = ?", isAudit).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsAuditUserIDIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsAuditUserIDIndex(auditUserID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`audit_user_id` = ?", auditUserID).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsAuditTimeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsAuditTimeIndex(auditTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`audit_time` = ?", auditTime).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsIsAcceptIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsIsAcceptIndex(isAccept []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_accept` = ?", isAccept).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsPaymentNumberIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsPaymentNumberIndex(paymentNumber string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`payment_number` = ?", paymentNumber).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsAcceptUserIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsAcceptUserIndex(acceptUser string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_user` = ?", acceptUser).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsAcceptUserIDIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsAcceptUserIDIndex(acceptUserID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_user_id` = ?", acceptUserID).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsAcceptTimeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsAcceptTimeIndex(acceptTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsCreateTimeIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsCreateTimeIndex(createTime time.Time) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeDetailsCreateUserIndex  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByFaPayableBillFeeDetailsCreateUserIndex(createUser int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelID  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexCustomerChannelID(customerChannelID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelName  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexCustomerChannelName(customerChannelName string) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// FetchIndexByIndexIsDiffAmount  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexIsDiffAmount(isDiffAmount []uint8) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`is_diff_amount` = ?", isDiffAmount).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerID  获取多个内容
func (obj *_FaPayableBillFeeDetailsMgr) FetchIndexByIndexCustomerID(customerID int) (results []*FaPayableBillFeeDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFeeDetails{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}
