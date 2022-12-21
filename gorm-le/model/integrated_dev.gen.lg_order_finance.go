package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _LgOrderFinanceMgr struct {
	*_BaseMgr
}

// LgOrderFinanceMgr open func
func LgOrderFinanceMgr(db *gorm.DB) *_LgOrderFinanceMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderFinanceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderFinanceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_finance"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderFinanceMgr) GetTableName() string {
	return "lg_order_finance"
}

// Reset 重置gorm会话
func (obj *_LgOrderFinanceMgr) Reset() *_LgOrderFinanceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderFinanceMgr) Get() (result LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderFinanceMgr) Gets() (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderFinanceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOrderID order_id获取 订单ID
func (obj *_LgOrderFinanceMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithCustomerBillBatchNumbers customer_bill_batch_numbers获取 客户对账单批次号,多个批次号逗号分割
func (obj *_LgOrderFinanceMgr) WithCustomerBillBatchNumbers(customerBillBatchNumbers string) Option {
	return optionFunc(func(o *options) { o.query["customer_bill_batch_numbers"] = customerBillBatchNumbers })
}

// WithCustomerReceiptNumbers customer_receipt_numbers获取 客户收款单批次号,多个批次号逗号分割
func (obj *_LgOrderFinanceMgr) WithCustomerReceiptNumbers(customerReceiptNumbers string) Option {
	return optionFunc(func(o *options) { o.query["customer_receipt_numbers"] = customerReceiptNumbers })
}

// WithPayableBillBatchNumbers payable_bill_batch_numbers获取 应付对账单批次号,多个批次逗号分割
func (obj *_LgOrderFinanceMgr) WithPayableBillBatchNumbers(payableBillBatchNumbers string) Option {
	return optionFunc(func(o *options) { o.query["payable_bill_batch_numbers"] = payableBillBatchNumbers })
}

// WithPayablePaymentNumber payable_payment_number获取 应付付款单批次号,多个批次逗号分割
func (obj *_LgOrderFinanceMgr) WithPayablePaymentNumber(payablePaymentNumber string) Option {
	return optionFunc(func(o *options) { o.query["payable_payment_number"] = payablePaymentNumber })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderFinanceMgr) GetByOption(opts ...Option) (result LgOrderFinance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderFinanceMgr) GetByOptions(opts ...Option) (results []*LgOrderFinance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderFinanceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderFinance, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where(options.query)
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

// GetFromOrderID 通过order_id获取内容 订单ID
func (obj *_LgOrderFinanceMgr) GetFromOrderID(orderID int) (result LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// GetBatchFromOrderID 批量查找 订单ID
func (obj *_LgOrderFinanceMgr) GetBatchFromOrderID(orderIDs []int) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromCustomerBillBatchNumbers 通过customer_bill_batch_numbers获取内容 客户对账单批次号,多个批次号逗号分割
func (obj *_LgOrderFinanceMgr) GetFromCustomerBillBatchNumbers(customerBillBatchNumbers string) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`customer_bill_batch_numbers` = ?", customerBillBatchNumbers).Find(&results).Error

	return
}

// GetBatchFromCustomerBillBatchNumbers 批量查找 客户对账单批次号,多个批次号逗号分割
func (obj *_LgOrderFinanceMgr) GetBatchFromCustomerBillBatchNumbers(customerBillBatchNumberss []string) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`customer_bill_batch_numbers` IN (?)", customerBillBatchNumberss).Find(&results).Error

	return
}

// GetFromCustomerReceiptNumbers 通过customer_receipt_numbers获取内容 客户收款单批次号,多个批次号逗号分割
func (obj *_LgOrderFinanceMgr) GetFromCustomerReceiptNumbers(customerReceiptNumbers string) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`customer_receipt_numbers` = ?", customerReceiptNumbers).Find(&results).Error

	return
}

// GetBatchFromCustomerReceiptNumbers 批量查找 客户收款单批次号,多个批次号逗号分割
func (obj *_LgOrderFinanceMgr) GetBatchFromCustomerReceiptNumbers(customerReceiptNumberss []string) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`customer_receipt_numbers` IN (?)", customerReceiptNumberss).Find(&results).Error

	return
}

// GetFromPayableBillBatchNumbers 通过payable_bill_batch_numbers获取内容 应付对账单批次号,多个批次逗号分割
func (obj *_LgOrderFinanceMgr) GetFromPayableBillBatchNumbers(payableBillBatchNumbers string) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`payable_bill_batch_numbers` = ?", payableBillBatchNumbers).Find(&results).Error

	return
}

// GetBatchFromPayableBillBatchNumbers 批量查找 应付对账单批次号,多个批次逗号分割
func (obj *_LgOrderFinanceMgr) GetBatchFromPayableBillBatchNumbers(payableBillBatchNumberss []string) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`payable_bill_batch_numbers` IN (?)", payableBillBatchNumberss).Find(&results).Error

	return
}

// GetFromPayablePaymentNumber 通过payable_payment_number获取内容 应付付款单批次号,多个批次逗号分割
func (obj *_LgOrderFinanceMgr) GetFromPayablePaymentNumber(payablePaymentNumber string) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`payable_payment_number` = ?", payablePaymentNumber).Find(&results).Error

	return
}

// GetBatchFromPayablePaymentNumber 批量查找 应付付款单批次号,多个批次逗号分割
func (obj *_LgOrderFinanceMgr) GetBatchFromPayablePaymentNumber(payablePaymentNumbers []string) (results []*LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`payable_payment_number` IN (?)", payablePaymentNumbers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderFinanceMgr) FetchByPrimaryKey(orderID int) (result LgOrderFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderFinance{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}
