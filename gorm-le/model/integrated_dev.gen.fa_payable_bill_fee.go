package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaPayableBillFeeMgr struct {
	*_BaseMgr
}

// FaPayableBillFeeMgr open func
func FaPayableBillFeeMgr(db *gorm.DB) *_FaPayableBillFeeMgr {
	if db == nil {
		panic(fmt.Errorf("FaPayableBillFeeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaPayableBillFeeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_payable_bill_fee"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaPayableBillFeeMgr) GetTableName() string {
	return "fa_payable_bill_fee"
}

// Reset 重置gorm会话
func (obj *_FaPayableBillFeeMgr) Reset() *_FaPayableBillFeeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaPayableBillFeeMgr) Get() (result FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaPayableBillFeeMgr) Gets() (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaPayableBillFeeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaPayableBillFeeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_FaPayableBillFeeMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_FaPayableBillFeeMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_FaPayableBillFeeMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_FaPayableBillFeeMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithBusinessEntityCode business_entity_code获取 业务主体代码
func (obj *_FaPayableBillFeeMgr) WithBusinessEntityCode(businessEntityCode string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_code"] = businessEntityCode })
}

// WithBusinessEntityName business_entity_name获取 业务主体名称
func (obj *_FaPayableBillFeeMgr) WithBusinessEntityName(businessEntityName string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_name"] = businessEntityName })
}

// WithBusinessEntityType business_entity_type获取 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillFeeMgr) WithBusinessEntityType(businessEntityType bool) Option {
	return optionFunc(func(o *options) { o.query["business_entity_type"] = businessEntityType })
}

// WithProviderChannelCode provider_channel_code获取 服务商渠道代码
func (obj *_FaPayableBillFeeMgr) WithProviderChannelCode(providerChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_code"] = providerChannelCode })
}

// WithReceiptTime receipt_time获取 入库时间
func (obj *_FaPayableBillFeeMgr) WithReceiptTime(receiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_time"] = receiptTime })
}

// WithSendTime send_time获取 出库时间
func (obj *_FaPayableBillFeeMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithIsSend is_send获取 是否出库
func (obj *_FaPayableBillFeeMgr) WithIsSend(isSend []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_send"] = isSend })
}

// WithWeight weight获取 预报重量
func (obj *_FaPayableBillFeeMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithWeighingWeight weighing_weight获取 称重重量
func (obj *_FaPayableBillFeeMgr) WithWeighingWeight(weighingWeight float64) Option {
	return optionFunc(func(o *options) { o.query["weighing_weight"] = weighingWeight })
}

// WithFoamWeight foam_weight获取 泡重
func (obj *_FaPayableBillFeeMgr) WithFoamWeight(foamWeight float64) Option {
	return optionFunc(func(o *options) { o.query["foam_weight"] = foamWeight })
}

// WithChargeWeight charge_weight获取 计费重量
func (obj *_FaPayableBillFeeMgr) WithChargeWeight(chargeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["charge_weight"] = chargeWeight })
}

// WithFeeQty fee_qty获取 费用数量
func (obj *_FaPayableBillFeeMgr) WithFeeQty(feeQty int) Option {
	return optionFunc(func(o *options) { o.query["fee_qty"] = feeQty })
}

// WithPayableRmb payable_rmb获取 应付费用(RMB),折前应付RMB
func (obj *_FaPayableBillFeeMgr) WithPayableRmb(payableRmb float64) Option {
	return optionFunc(func(o *options) { o.query["payable_rmb"] = payableRmb })
}

// WithBillBatchNumbers bill_batch_numbers获取 对账单批次号，多个以逗号隔开
func (obj *_FaPayableBillFeeMgr) WithBillBatchNumbers(billBatchNumbers string) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_numbers"] = billBatchNumbers })
}

// WithPaymentNumbers payment_numbers获取 付款单号，多个以逗号隔开
func (obj *_FaPayableBillFeeMgr) WithPaymentNumbers(paymentNumbers string) Option {
	return optionFunc(func(o *options) { o.query["payment_numbers"] = paymentNumbers })
}

// WithRemark remark获取 备注
func (obj *_FaPayableBillFeeMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithFinancialRemark financial_remark获取 财务备注
func (obj *_FaPayableBillFeeMgr) WithFinancialRemark(financialRemark string) Option {
	return optionFunc(func(o *options) { o.query["financial_remark"] = financialRemark })
}

// WithCreateUser create_user获取 创建人
func (obj *_FaPayableBillFeeMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaPayableBillFeeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaPayableBillFeeMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaPayableBillFeeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaPayableBillFeeMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_FaPayableBillFeeMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_FaPayableBillFeeMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_FaPayableBillFeeMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_FaPayableBillFeeMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// GetByOption 功能选项模式获取
func (obj *_FaPayableBillFeeMgr) GetByOption(opts ...Option) (result FaPayableBillFee, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaPayableBillFeeMgr) GetByOptions(opts ...Option) (results []*FaPayableBillFee, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaPayableBillFeeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaPayableBillFee, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where(options.query)
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
func (obj *_FaPayableBillFeeMgr) GetFromID(id int64) (result FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaPayableBillFeeMgr) GetBatchFromID(ids []int64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_FaPayableBillFeeMgr) GetFromOrderNumber(orderNumber string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_FaPayableBillFeeMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_FaPayableBillFeeMgr) GetFromReferenceNumber(referenceNumber string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_FaPayableBillFeeMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_FaPayableBillFeeMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_FaPayableBillFeeMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_FaPayableBillFeeMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_FaPayableBillFeeMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromBusinessEntityCode 通过business_entity_code获取内容 业务主体代码
func (obj *_FaPayableBillFeeMgr) GetFromBusinessEntityCode(businessEntityCode string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityCode 批量查找 业务主体代码
func (obj *_FaPayableBillFeeMgr) GetBatchFromBusinessEntityCode(businessEntityCodes []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_code` IN (?)", businessEntityCodes).Find(&results).Error

	return
}

// GetFromBusinessEntityName 通过business_entity_name获取内容 业务主体名称
func (obj *_FaPayableBillFeeMgr) GetFromBusinessEntityName(businessEntityName string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_name` = ?", businessEntityName).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityName 批量查找 业务主体名称
func (obj *_FaPayableBillFeeMgr) GetBatchFromBusinessEntityName(businessEntityNames []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_name` IN (?)", businessEntityNames).Find(&results).Error

	return
}

// GetFromBusinessEntityType 通过business_entity_type获取内容 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillFeeMgr) GetFromBusinessEntityType(businessEntityType bool) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_type` = ?", businessEntityType).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityType 批量查找 业务主体类型，0:服务商，1:航司
func (obj *_FaPayableBillFeeMgr) GetBatchFromBusinessEntityType(businessEntityTypes []bool) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_type` IN (?)", businessEntityTypes).Find(&results).Error

	return
}

// GetFromProviderChannelCode 通过provider_channel_code获取内容 服务商渠道代码
func (obj *_FaPayableBillFeeMgr) GetFromProviderChannelCode(providerChannelCode string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error

	return
}

// GetBatchFromProviderChannelCode 批量查找 服务商渠道代码
func (obj *_FaPayableBillFeeMgr) GetBatchFromProviderChannelCode(providerChannelCodes []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`provider_channel_code` IN (?)", providerChannelCodes).Find(&results).Error

	return
}

// GetFromReceiptTime 通过receipt_time获取内容 入库时间
func (obj *_FaPayableBillFeeMgr) GetFromReceiptTime(receiptTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// GetBatchFromReceiptTime 批量查找 入库时间
func (obj *_FaPayableBillFeeMgr) GetBatchFromReceiptTime(receiptTimes []time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`receipt_time` IN (?)", receiptTimes).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 出库时间
func (obj *_FaPayableBillFeeMgr) GetFromSendTime(sendTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 出库时间
func (obj *_FaPayableBillFeeMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromIsSend 通过is_send获取内容 是否出库
func (obj *_FaPayableBillFeeMgr) GetFromIsSend(isSend []uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`is_send` = ?", isSend).Find(&results).Error

	return
}

// GetBatchFromIsSend 批量查找 是否出库
func (obj *_FaPayableBillFeeMgr) GetBatchFromIsSend(isSends [][]uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`is_send` IN (?)", isSends).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 预报重量
func (obj *_FaPayableBillFeeMgr) GetFromWeight(weight float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 预报重量
func (obj *_FaPayableBillFeeMgr) GetBatchFromWeight(weights []float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromWeighingWeight 通过weighing_weight获取内容 称重重量
func (obj *_FaPayableBillFeeMgr) GetFromWeighingWeight(weighingWeight float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`weighing_weight` = ?", weighingWeight).Find(&results).Error

	return
}

// GetBatchFromWeighingWeight 批量查找 称重重量
func (obj *_FaPayableBillFeeMgr) GetBatchFromWeighingWeight(weighingWeights []float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`weighing_weight` IN (?)", weighingWeights).Find(&results).Error

	return
}

// GetFromFoamWeight 通过foam_weight获取内容 泡重
func (obj *_FaPayableBillFeeMgr) GetFromFoamWeight(foamWeight float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`foam_weight` = ?", foamWeight).Find(&results).Error

	return
}

// GetBatchFromFoamWeight 批量查找 泡重
func (obj *_FaPayableBillFeeMgr) GetBatchFromFoamWeight(foamWeights []float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`foam_weight` IN (?)", foamWeights).Find(&results).Error

	return
}

// GetFromChargeWeight 通过charge_weight获取内容 计费重量
func (obj *_FaPayableBillFeeMgr) GetFromChargeWeight(chargeWeight float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`charge_weight` = ?", chargeWeight).Find(&results).Error

	return
}

// GetBatchFromChargeWeight 批量查找 计费重量
func (obj *_FaPayableBillFeeMgr) GetBatchFromChargeWeight(chargeWeights []float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`charge_weight` IN (?)", chargeWeights).Find(&results).Error

	return
}

// GetFromFeeQty 通过fee_qty获取内容 费用数量
func (obj *_FaPayableBillFeeMgr) GetFromFeeQty(feeQty int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`fee_qty` = ?", feeQty).Find(&results).Error

	return
}

// GetBatchFromFeeQty 批量查找 费用数量
func (obj *_FaPayableBillFeeMgr) GetBatchFromFeeQty(feeQtys []int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`fee_qty` IN (?)", feeQtys).Find(&results).Error

	return
}

// GetFromPayableRmb 通过payable_rmb获取内容 应付费用(RMB),折前应付RMB
func (obj *_FaPayableBillFeeMgr) GetFromPayableRmb(payableRmb float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`payable_rmb` = ?", payableRmb).Find(&results).Error

	return
}

// GetBatchFromPayableRmb 批量查找 应付费用(RMB),折前应付RMB
func (obj *_FaPayableBillFeeMgr) GetBatchFromPayableRmb(payableRmbs []float64) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`payable_rmb` IN (?)", payableRmbs).Find(&results).Error

	return
}

// GetFromBillBatchNumbers 通过bill_batch_numbers获取内容 对账单批次号，多个以逗号隔开
func (obj *_FaPayableBillFeeMgr) GetFromBillBatchNumbers(billBatchNumbers string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`bill_batch_numbers` = ?", billBatchNumbers).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumbers 批量查找 对账单批次号，多个以逗号隔开
func (obj *_FaPayableBillFeeMgr) GetBatchFromBillBatchNumbers(billBatchNumberss []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`bill_batch_numbers` IN (?)", billBatchNumberss).Find(&results).Error

	return
}

// GetFromPaymentNumbers 通过payment_numbers获取内容 付款单号，多个以逗号隔开
func (obj *_FaPayableBillFeeMgr) GetFromPaymentNumbers(paymentNumbers string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`payment_numbers` = ?", paymentNumbers).Find(&results).Error

	return
}

// GetBatchFromPaymentNumbers 批量查找 付款单号，多个以逗号隔开
func (obj *_FaPayableBillFeeMgr) GetBatchFromPaymentNumbers(paymentNumberss []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`payment_numbers` IN (?)", paymentNumberss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaPayableBillFeeMgr) GetFromRemark(remark string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaPayableBillFeeMgr) GetBatchFromRemark(remarks []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromFinancialRemark 通过financial_remark获取内容 财务备注
func (obj *_FaPayableBillFeeMgr) GetFromFinancialRemark(financialRemark string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`financial_remark` = ?", financialRemark).Find(&results).Error

	return
}

// GetBatchFromFinancialRemark 批量查找 财务备注
func (obj *_FaPayableBillFeeMgr) GetBatchFromFinancialRemark(financialRemarks []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`financial_remark` IN (?)", financialRemarks).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_FaPayableBillFeeMgr) GetFromCreateUser(createUser int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_FaPayableBillFeeMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaPayableBillFeeMgr) GetFromCreateTime(createTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaPayableBillFeeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaPayableBillFeeMgr) GetFromUpdateUser(updateUser int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaPayableBillFeeMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaPayableBillFeeMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaPayableBillFeeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaPayableBillFeeMgr) GetFromVersion(version int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaPayableBillFeeMgr) GetBatchFromVersion(versions []int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_FaPayableBillFeeMgr) GetFromCustomerChannelID(customerChannelID int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_FaPayableBillFeeMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_FaPayableBillFeeMgr) GetFromCustomerChannelName(customerChannelName string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_FaPayableBillFeeMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_FaPayableBillFeeMgr) GetFromCustomerID(customerID int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_FaPayableBillFeeMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_FaPayableBillFeeMgr) GetFromCustomerAlias(customerAlias string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_FaPayableBillFeeMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaPayableBillFeeMgr) FetchByPrimaryKey(id int64) (result FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByUniqueReferenceNumberBusinessEntityCode primary or index 获取唯一内容
func (obj *_FaPayableBillFeeMgr) FetchUniqueIndexByUniqueReferenceNumberBusinessEntityCode(referenceNumber string, businessEntityCode string) (result FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`reference_number` = ? AND `business_entity_code` = ?", referenceNumber, businessEntityCode).First(&result).Error

	return
}

// FetchIndexByIndexOrderNumber  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexOrderNumber(orderNumber string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByIndexReferenceNumber  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexReferenceNumber(referenceNumber string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumber  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexLogisticsNumber(logisticsNumber string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumberBusinessEntityCode  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexLogisticsNumberBusinessEntityCode(logisticsNumber string, businessEntityCode string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`logistics_number` = ? AND `business_entity_code` = ?", logisticsNumber, businessEntityCode).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumberFinal  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexByIndexBusinessEntityCode  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexBusinessEntityCode(businessEntityCode string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// FetchIndexByIndexBusinessEntityCodeCreateTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexBusinessEntityCodeCreateTimeIsSend(businessEntityCode string, isSend []uint8, createTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_code` = ? AND `is_send` = ? AND `create_time` = ?", businessEntityCode, isSend, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexBusinessEntityCodeReceiptTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexBusinessEntityCodeReceiptTimeIsSend(businessEntityCode string, receiptTime time.Time, isSend []uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_code` = ? AND `receipt_time` = ? AND `is_send` = ?", businessEntityCode, receiptTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexBusinessEntityCodeSendTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexBusinessEntityCodeSendTimeIsSend(businessEntityCode string, sendTime time.Time, isSend []uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_code` = ? AND `send_time` = ? AND `is_send` = ?", businessEntityCode, sendTime, isSend).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeBusinessEntityCodeReceiptTimeIndex  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByFaPayableBillFeeBusinessEntityCodeReceiptTimeIndex(businessEntityCode string, receiptTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_code` = ? AND `receipt_time` = ?", businessEntityCode, receiptTime).Find(&results).Error

	return
}

// FetchIndexByFaPayableBillFeeBusinessEntityCodeSendTimeIndex  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByFaPayableBillFeeBusinessEntityCodeSendTimeIndex(businessEntityCode string, sendTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_code` = ? AND `send_time` = ?", businessEntityCode, sendTime).Find(&results).Error

	return
}

// FetchIndexByIndexBusinessEntityType  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexBusinessEntityType(businessEntityType bool) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`business_entity_type` = ?", businessEntityType).Find(&results).Error

	return
}

// FetchIndexByIndexProviderChannelCode  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexProviderChannelCode(providerChannelCode string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error

	return
}

// FetchIndexByIndexProviderChannelCodeCreateTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexProviderChannelCodeCreateTimeIsSend(providerChannelCode string, isSend []uint8, createTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`provider_channel_code` = ? AND `is_send` = ? AND `create_time` = ?", providerChannelCode, isSend, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexProviderChannelCodeReceiptTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexProviderChannelCodeReceiptTimeIsSend(providerChannelCode string, receiptTime time.Time, isSend []uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`provider_channel_code` = ? AND `receipt_time` = ? AND `is_send` = ?", providerChannelCode, receiptTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexProviderChannelCodeSendTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexProviderChannelCodeSendTimeIsSend(providerChannelCode string, sendTime time.Time, isSend []uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`provider_channel_code` = ? AND `send_time` = ? AND `is_send` = ?", providerChannelCode, sendTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexReceiptTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexReceiptTimeIsSend(receiptTime time.Time, isSend []uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`receipt_time` = ? AND `is_send` = ?", receiptTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexSendTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexSendTimeIsSend(sendTime time.Time, isSend []uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`send_time` = ? AND `is_send` = ?", sendTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexIsSend(isSend []uint8) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`is_send` = ?", isSend).Find(&results).Error

	return
}

// FetchIndexByIndexCreateTimeIsSend  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexCreateTimeIsSend(isSend []uint8, createTime time.Time) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`is_send` = ? AND `create_time` = ?", isSend, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexBillBatchNumbers  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexBillBatchNumbers(billBatchNumbers string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`bill_batch_numbers` = ?", billBatchNumbers).Find(&results).Error

	return
}

// FetchIndexByIndexPaymentNumbers  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexPaymentNumbers(paymentNumbers string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`payment_numbers` = ?", paymentNumbers).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelID  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexCustomerChannelID(customerChannelID int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelName  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexCustomerChannelName(customerChannelName string) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerID  获取多个内容
func (obj *_FaPayableBillFeeMgr) FetchIndexByIndexCustomerID(customerID int) (results []*FaPayableBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPayableBillFee{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}
