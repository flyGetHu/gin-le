package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaCustomerBillFeeMgr struct {
	*_BaseMgr
}

// FaCustomerBillFeeMgr open func
func FaCustomerBillFeeMgr(db *gorm.DB) *_FaCustomerBillFeeMgr {
	if db == nil {
		panic(fmt.Errorf("FaCustomerBillFeeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaCustomerBillFeeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_customer_bill_fee"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaCustomerBillFeeMgr) GetTableName() string {
	return "fa_customer_bill_fee"
}

// Reset 重置gorm会话
func (obj *_FaCustomerBillFeeMgr) Reset() *_FaCustomerBillFeeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaCustomerBillFeeMgr) Get() (result FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaCustomerBillFeeMgr) Gets() (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaCustomerBillFeeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaCustomerBillFeeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 订单ID
func (obj *_FaCustomerBillFeeMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithOrderNumber order_number获取 订单号
func (obj *_FaCustomerBillFeeMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_FaCustomerBillFeeMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_FaCustomerBillFeeMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_FaCustomerBillFeeMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_FaCustomerBillFeeMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_FaCustomerBillFeeMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_FaCustomerBillFeeMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_FaCustomerBillFeeMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithReceiptTime receipt_time获取 入库时间
func (obj *_FaCustomerBillFeeMgr) WithReceiptTime(receiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_time"] = receiptTime })
}

// WithSendTime send_time获取 出库时间
func (obj *_FaCustomerBillFeeMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithIsSend is_send获取 是否出库
func (obj *_FaCustomerBillFeeMgr) WithIsSend(isSend []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_send"] = isSend })
}

// WithWeight weight获取 预报重量
func (obj *_FaCustomerBillFeeMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithWeighingWeight weighing_weight获取 称重重量
func (obj *_FaCustomerBillFeeMgr) WithWeighingWeight(weighingWeight float64) Option {
	return optionFunc(func(o *options) { o.query["weighing_weight"] = weighingWeight })
}

// WithFoamWeight foam_weight获取 泡重
func (obj *_FaCustomerBillFeeMgr) WithFoamWeight(foamWeight float64) Option {
	return optionFunc(func(o *options) { o.query["foam_weight"] = foamWeight })
}

// WithChargeWeight charge_weight获取 计费重量
func (obj *_FaCustomerBillFeeMgr) WithChargeWeight(chargeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["charge_weight"] = chargeWeight })
}

// WithFeeQty fee_qty获取 费用数量
func (obj *_FaCustomerBillFeeMgr) WithFeeQty(feeQty int) Option {
	return optionFunc(func(o *options) { o.query["fee_qty"] = feeQty })
}

// WithReceivablesRmb receivables_rmb获取 应收费用(RMB)
func (obj *_FaCustomerBillFeeMgr) WithReceivablesRmb(receivablesRmb float64) Option {
	return optionFunc(func(o *options) { o.query["receivables_rmb"] = receivablesRmb })
}

// WithBillBatchNumbers bill_batch_numbers获取 客户对账单批次号
func (obj *_FaCustomerBillFeeMgr) WithBillBatchNumbers(billBatchNumbers string) Option {
	return optionFunc(func(o *options) { o.query["bill_batch_numbers"] = billBatchNumbers })
}

// WithReceiptNumbers receipt_numbers获取 客户收款单批次号
func (obj *_FaCustomerBillFeeMgr) WithReceiptNumbers(receiptNumbers string) Option {
	return optionFunc(func(o *options) { o.query["receipt_numbers"] = receiptNumbers })
}

// WithRemark remark获取 备注
func (obj *_FaCustomerBillFeeMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithFinancialRemark financial_remark获取 财务备注
func (obj *_FaCustomerBillFeeMgr) WithFinancialRemark(financialRemark string) Option {
	return optionFunc(func(o *options) { o.query["financial_remark"] = financialRemark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaCustomerBillFeeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应收人员
func (obj *_FaCustomerBillFeeMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaCustomerBillFeeMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaCustomerBillFeeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaCustomerBillFeeMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaCustomerBillFeeMgr) GetByOption(opts ...Option) (result FaCustomerBillFee, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaCustomerBillFeeMgr) GetByOptions(opts ...Option) (results []*FaCustomerBillFee, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaCustomerBillFeeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaCustomerBillFee, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where(options.query)
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
func (obj *_FaCustomerBillFeeMgr) GetFromID(id int) (result FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaCustomerBillFeeMgr) GetBatchFromID(ids []int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单ID
func (obj *_FaCustomerBillFeeMgr) GetFromOrderID(orderID int) (result FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// GetBatchFromOrderID 批量查找 订单ID
func (obj *_FaCustomerBillFeeMgr) GetBatchFromOrderID(orderIDs []int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_FaCustomerBillFeeMgr) GetFromOrderNumber(orderNumber string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_FaCustomerBillFeeMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_FaCustomerBillFeeMgr) GetFromReferenceNumber(referenceNumber string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_FaCustomerBillFeeMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_FaCustomerBillFeeMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_FaCustomerBillFeeMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_FaCustomerBillFeeMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_FaCustomerBillFeeMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_FaCustomerBillFeeMgr) GetFromCustomerID(customerID int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_FaCustomerBillFeeMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_FaCustomerBillFeeMgr) GetFromCustomerAlias(customerAlias string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_FaCustomerBillFeeMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_FaCustomerBillFeeMgr) GetFromCustomerChannelID(customerChannelID int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_FaCustomerBillFeeMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_FaCustomerBillFeeMgr) GetFromCustomerChannelName(customerChannelName string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_FaCustomerBillFeeMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromReceiptTime 通过receipt_time获取内容 入库时间
func (obj *_FaCustomerBillFeeMgr) GetFromReceiptTime(receiptTime time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// GetBatchFromReceiptTime 批量查找 入库时间
func (obj *_FaCustomerBillFeeMgr) GetBatchFromReceiptTime(receiptTimes []time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`receipt_time` IN (?)", receiptTimes).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 出库时间
func (obj *_FaCustomerBillFeeMgr) GetFromSendTime(sendTime time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 出库时间
func (obj *_FaCustomerBillFeeMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromIsSend 通过is_send获取内容 是否出库
func (obj *_FaCustomerBillFeeMgr) GetFromIsSend(isSend []uint8) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`is_send` = ?", isSend).Find(&results).Error

	return
}

// GetBatchFromIsSend 批量查找 是否出库
func (obj *_FaCustomerBillFeeMgr) GetBatchFromIsSend(isSends [][]uint8) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`is_send` IN (?)", isSends).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 预报重量
func (obj *_FaCustomerBillFeeMgr) GetFromWeight(weight float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 预报重量
func (obj *_FaCustomerBillFeeMgr) GetBatchFromWeight(weights []float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromWeighingWeight 通过weighing_weight获取内容 称重重量
func (obj *_FaCustomerBillFeeMgr) GetFromWeighingWeight(weighingWeight float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`weighing_weight` = ?", weighingWeight).Find(&results).Error

	return
}

// GetBatchFromWeighingWeight 批量查找 称重重量
func (obj *_FaCustomerBillFeeMgr) GetBatchFromWeighingWeight(weighingWeights []float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`weighing_weight` IN (?)", weighingWeights).Find(&results).Error

	return
}

// GetFromFoamWeight 通过foam_weight获取内容 泡重
func (obj *_FaCustomerBillFeeMgr) GetFromFoamWeight(foamWeight float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`foam_weight` = ?", foamWeight).Find(&results).Error

	return
}

// GetBatchFromFoamWeight 批量查找 泡重
func (obj *_FaCustomerBillFeeMgr) GetBatchFromFoamWeight(foamWeights []float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`foam_weight` IN (?)", foamWeights).Find(&results).Error

	return
}

// GetFromChargeWeight 通过charge_weight获取内容 计费重量
func (obj *_FaCustomerBillFeeMgr) GetFromChargeWeight(chargeWeight float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`charge_weight` = ?", chargeWeight).Find(&results).Error

	return
}

// GetBatchFromChargeWeight 批量查找 计费重量
func (obj *_FaCustomerBillFeeMgr) GetBatchFromChargeWeight(chargeWeights []float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`charge_weight` IN (?)", chargeWeights).Find(&results).Error

	return
}

// GetFromFeeQty 通过fee_qty获取内容 费用数量
func (obj *_FaCustomerBillFeeMgr) GetFromFeeQty(feeQty int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`fee_qty` = ?", feeQty).Find(&results).Error

	return
}

// GetBatchFromFeeQty 批量查找 费用数量
func (obj *_FaCustomerBillFeeMgr) GetBatchFromFeeQty(feeQtys []int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`fee_qty` IN (?)", feeQtys).Find(&results).Error

	return
}

// GetFromReceivablesRmb 通过receivables_rmb获取内容 应收费用(RMB)
func (obj *_FaCustomerBillFeeMgr) GetFromReceivablesRmb(receivablesRmb float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`receivables_rmb` = ?", receivablesRmb).Find(&results).Error

	return
}

// GetBatchFromReceivablesRmb 批量查找 应收费用(RMB)
func (obj *_FaCustomerBillFeeMgr) GetBatchFromReceivablesRmb(receivablesRmbs []float64) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`receivables_rmb` IN (?)", receivablesRmbs).Find(&results).Error

	return
}

// GetFromBillBatchNumbers 通过bill_batch_numbers获取内容 客户对账单批次号
func (obj *_FaCustomerBillFeeMgr) GetFromBillBatchNumbers(billBatchNumbers string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`bill_batch_numbers` = ?", billBatchNumbers).Find(&results).Error

	return
}

// GetBatchFromBillBatchNumbers 批量查找 客户对账单批次号
func (obj *_FaCustomerBillFeeMgr) GetBatchFromBillBatchNumbers(billBatchNumberss []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`bill_batch_numbers` IN (?)", billBatchNumberss).Find(&results).Error

	return
}

// GetFromReceiptNumbers 通过receipt_numbers获取内容 客户收款单批次号
func (obj *_FaCustomerBillFeeMgr) GetFromReceiptNumbers(receiptNumbers string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`receipt_numbers` = ?", receiptNumbers).Find(&results).Error

	return
}

// GetBatchFromReceiptNumbers 批量查找 客户收款单批次号
func (obj *_FaCustomerBillFeeMgr) GetBatchFromReceiptNumbers(receiptNumberss []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`receipt_numbers` IN (?)", receiptNumberss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaCustomerBillFeeMgr) GetFromRemark(remark string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaCustomerBillFeeMgr) GetBatchFromRemark(remarks []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromFinancialRemark 通过financial_remark获取内容 财务备注
func (obj *_FaCustomerBillFeeMgr) GetFromFinancialRemark(financialRemark string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`financial_remark` = ?", financialRemark).Find(&results).Error

	return
}

// GetBatchFromFinancialRemark 批量查找 财务备注
func (obj *_FaCustomerBillFeeMgr) GetBatchFromFinancialRemark(financialRemarks []string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`financial_remark` IN (?)", financialRemarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaCustomerBillFeeMgr) GetFromCreateTime(createTime time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaCustomerBillFeeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应收人员
func (obj *_FaCustomerBillFeeMgr) GetFromCreateUser(createUser int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应收人员
func (obj *_FaCustomerBillFeeMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaCustomerBillFeeMgr) GetFromUpdateUser(updateUser int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaCustomerBillFeeMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaCustomerBillFeeMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaCustomerBillFeeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaCustomerBillFeeMgr) GetFromVersion(version int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaCustomerBillFeeMgr) GetBatchFromVersion(versions []int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaCustomerBillFeeMgr) FetchByPrimaryKey(id int) (result FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByFaCustomerBillFeeOrderIDUIndex primary or index 获取唯一内容
func (obj *_FaCustomerBillFeeMgr) FetchUniqueByFaCustomerBillFeeOrderIDUIndex(orderID int) (result FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// FetchUniqueByUniqueReferenceNumber primary or index 获取唯一内容
func (obj *_FaCustomerBillFeeMgr) FetchUniqueByUniqueReferenceNumber(referenceNumber string) (result FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`reference_number` = ?", referenceNumber).First(&result).Error

	return
}

// FetchIndexByIndexOrderNumber  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexOrderNumber(orderNumber string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByIndexReferenceNumber  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexReferenceNumber(referenceNumber string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumber  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexLogisticsNumber(logisticsNumber string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByIndexLogisticsNumberFinal  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerID  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCustomerID(customerID int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerIDReceiptTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCustomerIDReceiptTimeIsSend(customerID int, receiptTime time.Time, isSend []uint8) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_id` = ? AND `receipt_time` = ? AND `is_send` = ?", customerID, receiptTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerIDSendTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCustomerIDSendTimeIsSend(customerID int, sendTime time.Time, isSend []uint8) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_id` = ? AND `send_time` = ? AND `is_send` = ?", customerID, sendTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerIDCreateTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCustomerIDCreateTimeIsSend(customerID int, isSend []uint8, createTime time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_id` = ? AND `is_send` = ? AND `create_time` = ?", customerID, isSend, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelID  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCustomerChannelID(customerChannelID int) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelIDReceiptTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCustomerChannelIDReceiptTimeIsSend(customerChannelID int, receiptTime time.Time, isSend []uint8) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_channel_id` = ? AND `receipt_time` = ? AND `is_send` = ?", customerChannelID, receiptTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelIDSendTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCustomerChannelIDSendTimeIsSend(customerChannelID int, sendTime time.Time, isSend []uint8) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_channel_id` = ? AND `send_time` = ? AND `is_send` = ?", customerChannelID, sendTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexCustomerChannelIDCreateTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCustomerChannelIDCreateTimeIsSend(customerChannelID int, isSend []uint8, createTime time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`customer_channel_id` = ? AND `is_send` = ? AND `create_time` = ?", customerChannelID, isSend, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexReceiptTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexReceiptTimeIsSend(receiptTime time.Time, isSend []uint8) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`receipt_time` = ? AND `is_send` = ?", receiptTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexSendTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexSendTimeIsSend(sendTime time.Time, isSend []uint8) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`send_time` = ? AND `is_send` = ?", sendTime, isSend).Find(&results).Error

	return
}

// FetchIndexByIndexCreateTimeIsSend  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexCreateTimeIsSend(isSend []uint8, createTime time.Time) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`is_send` = ? AND `create_time` = ?", isSend, createTime).Find(&results).Error

	return
}

// FetchIndexByIndexBillBatchNumbers  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexBillBatchNumbers(billBatchNumbers string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`bill_batch_numbers` = ?", billBatchNumbers).Find(&results).Error

	return
}

// FetchIndexByIndexReceiptNumbers  获取多个内容
func (obj *_FaCustomerBillFeeMgr) FetchIndexByIndexReceiptNumbers(receiptNumbers string) (results []*FaCustomerBillFee, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerBillFee{}).Where("`receipt_numbers` = ?", receiptNumbers).Find(&results).Error

	return
}
