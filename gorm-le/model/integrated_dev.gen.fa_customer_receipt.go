package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaCustomerReceiptMgr struct {
	*_BaseMgr
}

// FaCustomerReceiptMgr open func
func FaCustomerReceiptMgr(db *gorm.DB) *_FaCustomerReceiptMgr {
	if db == nil {
		panic(fmt.Errorf("FaCustomerReceiptMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaCustomerReceiptMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_customer_receipt"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaCustomerReceiptMgr) GetTableName() string {
	return "fa_customer_receipt"
}

// Reset 重置gorm会话
func (obj *_FaCustomerReceiptMgr) Reset() *_FaCustomerReceiptMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaCustomerReceiptMgr) Get() (result FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaCustomerReceiptMgr) Gets() (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaCustomerReceiptMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaCustomerReceiptMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithReceiptNumber receipt_number获取 收款单号
func (obj *_FaCustomerReceiptMgr) WithReceiptNumber(receiptNumber string) Option {
	return optionFunc(func(o *options) { o.query["receipt_number"] = receiptNumber })
}

// WithParentID parent_id获取 父级iD
func (obj *_FaCustomerReceiptMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithIsParent is_parent获取 是否是父级节点
func (obj *_FaCustomerReceiptMgr) WithIsParent(isParent []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_parent"] = isParent })
}

// WithInvoiceNumber invoice_number获取 发票号
func (obj *_FaCustomerReceiptMgr) WithInvoiceNumber(invoiceNumber string) Option {
	return optionFunc(func(o *options) { o.query["invoice_number"] = invoiceNumber })
}

// WithVoucherNumber voucher_number获取 凭证号
func (obj *_FaCustomerReceiptMgr) WithVoucherNumber(voucherNumber string) Option {
	return optionFunc(func(o *options) { o.query["voucher_number"] = voucherNumber })
}

// WithTicketNumber ticket_number获取 票证号
func (obj *_FaCustomerReceiptMgr) WithTicketNumber(ticketNumber string) Option {
	return optionFunc(func(o *options) { o.query["ticket_number"] = ticketNumber })
}

// WithReceiptStatus receipt_status获取 收款状态
func (obj *_FaCustomerReceiptMgr) WithReceiptStatus(receiptStatus int) Option {
	return optionFunc(func(o *options) { o.query["receipt_status"] = receiptStatus })
}

// WithIsAccept is_accept获取 是否已收款
func (obj *_FaCustomerReceiptMgr) WithIsAccept(isAccept []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_accept"] = isAccept })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_FaCustomerReceiptMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_FaCustomerReceiptMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithBillFee bill_fee获取 票据金额
func (obj *_FaCustomerReceiptMgr) WithBillFee(billFee float64) Option {
	return optionFunc(func(o *options) { o.query["bill_fee"] = billFee })
}

// WithActualFee actual_fee获取 实际收款金额
func (obj *_FaCustomerReceiptMgr) WithActualFee(actualFee float64) Option {
	return optionFunc(func(o *options) { o.query["actual_fee"] = actualFee })
}

// WithCurrencyCode currency_code获取 币种
func (obj *_FaCustomerReceiptMgr) WithCurrencyCode(currencyCode string) Option {
	return optionFunc(func(o *options) { o.query["currency_code"] = currencyCode })
}

// WithCurrencyName currency_name获取 币种中文名称
func (obj *_FaCustomerReceiptMgr) WithCurrencyName(currencyName string) Option {
	return optionFunc(func(o *options) { o.query["currency_name"] = currencyName })
}

// WithExchangeRate exchange_rate获取 汇率
func (obj *_FaCustomerReceiptMgr) WithExchangeRate(exchangeRate float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_rate"] = exchangeRate })
}

// WithActualFeeRmb actual_fee_rmb获取 实际收款金额rmb
func (obj *_FaCustomerReceiptMgr) WithActualFeeRmb(actualFeeRmb float64) Option {
	return optionFunc(func(o *options) { o.query["actual_fee_rmb"] = actualFeeRmb })
}

// WithDealWithUser deal_with_user获取 经办人
func (obj *_FaCustomerReceiptMgr) WithDealWithUser(dealWithUser string) Option {
	return optionFunc(func(o *options) { o.query["deal_with_user"] = dealWithUser })
}

// WithBillInputUser bill_input_user获取 票据录入人员
func (obj *_FaCustomerReceiptMgr) WithBillInputUser(billInputUser string) Option {
	return optionFunc(func(o *options) { o.query["bill_input_user"] = billInputUser })
}

// WithBillInputUserID bill_input_user_id获取 票据录入人员ID
func (obj *_FaCustomerReceiptMgr) WithBillInputUserID(billInputUserID int) Option {
	return optionFunc(func(o *options) { o.query["bill_input_user_id"] = billInputUserID })
}

// WithReceiptType receipt_type获取 收款类别 0:核收 1:清缴
func (obj *_FaCustomerReceiptMgr) WithReceiptType(receiptType int) Option {
	return optionFunc(func(o *options) { o.query["receipt_type"] = receiptType })
}

// WithReceiptMethod receipt_method获取 收款方式
func (obj *_FaCustomerReceiptMgr) WithReceiptMethod(receiptMethod string) Option {
	return optionFunc(func(o *options) { o.query["receipt_method"] = receiptMethod })
}

// WithPaymentTime payment_time获取 收款时间/支付时间
func (obj *_FaCustomerReceiptMgr) WithPaymentTime(paymentTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["payment_time"] = paymentTime })
}

// WithMark mark获取 标签/标记
func (obj *_FaCustomerReceiptMgr) WithMark(mark string) Option {
	return optionFunc(func(o *options) { o.query["mark"] = mark })
}

// WithAccountAlias account_alias获取 账户别称/走账
func (obj *_FaCustomerReceiptMgr) WithAccountAlias(accountAlias string) Option {
	return optionFunc(func(o *options) { o.query["account_alias"] = accountAlias })
}

// WithAttachmentURL attachment_url获取 附件
func (obj *_FaCustomerReceiptMgr) WithAttachmentURL(attachmentURL string) Option {
	return optionFunc(func(o *options) { o.query["attachment_url"] = attachmentURL })
}

// WithRemark remark获取 备注
func (obj *_FaCustomerReceiptMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaCustomerReceiptMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应收人员
func (obj *_FaCustomerReceiptMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaCustomerReceiptMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaCustomerReceiptMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaCustomerReceiptMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaCustomerReceiptMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaCustomerReceiptMgr) GetByOption(opts ...Option) (result FaCustomerReceipt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaCustomerReceiptMgr) GetByOptions(opts ...Option) (results []*FaCustomerReceipt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaCustomerReceiptMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaCustomerReceipt, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where(options.query)
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
func (obj *_FaCustomerReceiptMgr) GetFromID(id int) (result FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaCustomerReceiptMgr) GetBatchFromID(ids []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromReceiptNumber 通过receipt_number获取内容 收款单号
func (obj *_FaCustomerReceiptMgr) GetFromReceiptNumber(receiptNumber string) (result FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_number` = ?", receiptNumber).First(&result).Error

	return
}

// GetBatchFromReceiptNumber 批量查找 收款单号
func (obj *_FaCustomerReceiptMgr) GetBatchFromReceiptNumber(receiptNumbers []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_number` IN (?)", receiptNumbers).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父级iD
func (obj *_FaCustomerReceiptMgr) GetFromParentID(parentID int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父级iD
func (obj *_FaCustomerReceiptMgr) GetBatchFromParentID(parentIDs []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromIsParent 通过is_parent获取内容 是否是父级节点
func (obj *_FaCustomerReceiptMgr) GetFromIsParent(isParent []uint8) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`is_parent` = ?", isParent).Find(&results).Error

	return
}

// GetBatchFromIsParent 批量查找 是否是父级节点
func (obj *_FaCustomerReceiptMgr) GetBatchFromIsParent(isParents [][]uint8) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`is_parent` IN (?)", isParents).Find(&results).Error

	return
}

// GetFromInvoiceNumber 通过invoice_number获取内容 发票号
func (obj *_FaCustomerReceiptMgr) GetFromInvoiceNumber(invoiceNumber string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`invoice_number` = ?", invoiceNumber).Find(&results).Error

	return
}

// GetBatchFromInvoiceNumber 批量查找 发票号
func (obj *_FaCustomerReceiptMgr) GetBatchFromInvoiceNumber(invoiceNumbers []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`invoice_number` IN (?)", invoiceNumbers).Find(&results).Error

	return
}

// GetFromVoucherNumber 通过voucher_number获取内容 凭证号
func (obj *_FaCustomerReceiptMgr) GetFromVoucherNumber(voucherNumber string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`voucher_number` = ?", voucherNumber).Find(&results).Error

	return
}

// GetBatchFromVoucherNumber 批量查找 凭证号
func (obj *_FaCustomerReceiptMgr) GetBatchFromVoucherNumber(voucherNumbers []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`voucher_number` IN (?)", voucherNumbers).Find(&results).Error

	return
}

// GetFromTicketNumber 通过ticket_number获取内容 票证号
func (obj *_FaCustomerReceiptMgr) GetFromTicketNumber(ticketNumber string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`ticket_number` = ?", ticketNumber).Find(&results).Error

	return
}

// GetBatchFromTicketNumber 批量查找 票证号
func (obj *_FaCustomerReceiptMgr) GetBatchFromTicketNumber(ticketNumbers []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`ticket_number` IN (?)", ticketNumbers).Find(&results).Error

	return
}

// GetFromReceiptStatus 通过receipt_status获取内容 收款状态
func (obj *_FaCustomerReceiptMgr) GetFromReceiptStatus(receiptStatus int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_status` = ?", receiptStatus).Find(&results).Error

	return
}

// GetBatchFromReceiptStatus 批量查找 收款状态
func (obj *_FaCustomerReceiptMgr) GetBatchFromReceiptStatus(receiptStatuss []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_status` IN (?)", receiptStatuss).Find(&results).Error

	return
}

// GetFromIsAccept 通过is_accept获取内容 是否已收款
func (obj *_FaCustomerReceiptMgr) GetFromIsAccept(isAccept []uint8) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`is_accept` = ?", isAccept).Find(&results).Error

	return
}

// GetBatchFromIsAccept 批量查找 是否已收款
func (obj *_FaCustomerReceiptMgr) GetBatchFromIsAccept(isAccepts [][]uint8) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`is_accept` IN (?)", isAccepts).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_FaCustomerReceiptMgr) GetFromCustomerID(customerID int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_FaCustomerReceiptMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_FaCustomerReceiptMgr) GetFromCustomerAlias(customerAlias string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_FaCustomerReceiptMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromBillFee 通过bill_fee获取内容 票据金额
func (obj *_FaCustomerReceiptMgr) GetFromBillFee(billFee float64) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`bill_fee` = ?", billFee).Find(&results).Error

	return
}

// GetBatchFromBillFee 批量查找 票据金额
func (obj *_FaCustomerReceiptMgr) GetBatchFromBillFee(billFees []float64) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`bill_fee` IN (?)", billFees).Find(&results).Error

	return
}

// GetFromActualFee 通过actual_fee获取内容 实际收款金额
func (obj *_FaCustomerReceiptMgr) GetFromActualFee(actualFee float64) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`actual_fee` = ?", actualFee).Find(&results).Error

	return
}

// GetBatchFromActualFee 批量查找 实际收款金额
func (obj *_FaCustomerReceiptMgr) GetBatchFromActualFee(actualFees []float64) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`actual_fee` IN (?)", actualFees).Find(&results).Error

	return
}

// GetFromCurrencyCode 通过currency_code获取内容 币种
func (obj *_FaCustomerReceiptMgr) GetFromCurrencyCode(currencyCode string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`currency_code` = ?", currencyCode).Find(&results).Error

	return
}

// GetBatchFromCurrencyCode 批量查找 币种
func (obj *_FaCustomerReceiptMgr) GetBatchFromCurrencyCode(currencyCodes []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`currency_code` IN (?)", currencyCodes).Find(&results).Error

	return
}

// GetFromCurrencyName 通过currency_name获取内容 币种中文名称
func (obj *_FaCustomerReceiptMgr) GetFromCurrencyName(currencyName string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`currency_name` = ?", currencyName).Find(&results).Error

	return
}

// GetBatchFromCurrencyName 批量查找 币种中文名称
func (obj *_FaCustomerReceiptMgr) GetBatchFromCurrencyName(currencyNames []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`currency_name` IN (?)", currencyNames).Find(&results).Error

	return
}

// GetFromExchangeRate 通过exchange_rate获取内容 汇率
func (obj *_FaCustomerReceiptMgr) GetFromExchangeRate(exchangeRate float64) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`exchange_rate` = ?", exchangeRate).Find(&results).Error

	return
}

// GetBatchFromExchangeRate 批量查找 汇率
func (obj *_FaCustomerReceiptMgr) GetBatchFromExchangeRate(exchangeRates []float64) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`exchange_rate` IN (?)", exchangeRates).Find(&results).Error

	return
}

// GetFromActualFeeRmb 通过actual_fee_rmb获取内容 实际收款金额rmb
func (obj *_FaCustomerReceiptMgr) GetFromActualFeeRmb(actualFeeRmb float64) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`actual_fee_rmb` = ?", actualFeeRmb).Find(&results).Error

	return
}

// GetBatchFromActualFeeRmb 批量查找 实际收款金额rmb
func (obj *_FaCustomerReceiptMgr) GetBatchFromActualFeeRmb(actualFeeRmbs []float64) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`actual_fee_rmb` IN (?)", actualFeeRmbs).Find(&results).Error

	return
}

// GetFromDealWithUser 通过deal_with_user获取内容 经办人
func (obj *_FaCustomerReceiptMgr) GetFromDealWithUser(dealWithUser string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`deal_with_user` = ?", dealWithUser).Find(&results).Error

	return
}

// GetBatchFromDealWithUser 批量查找 经办人
func (obj *_FaCustomerReceiptMgr) GetBatchFromDealWithUser(dealWithUsers []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`deal_with_user` IN (?)", dealWithUsers).Find(&results).Error

	return
}

// GetFromBillInputUser 通过bill_input_user获取内容 票据录入人员
func (obj *_FaCustomerReceiptMgr) GetFromBillInputUser(billInputUser string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`bill_input_user` = ?", billInputUser).Find(&results).Error

	return
}

// GetBatchFromBillInputUser 批量查找 票据录入人员
func (obj *_FaCustomerReceiptMgr) GetBatchFromBillInputUser(billInputUsers []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`bill_input_user` IN (?)", billInputUsers).Find(&results).Error

	return
}

// GetFromBillInputUserID 通过bill_input_user_id获取内容 票据录入人员ID
func (obj *_FaCustomerReceiptMgr) GetFromBillInputUserID(billInputUserID int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`bill_input_user_id` = ?", billInputUserID).Find(&results).Error

	return
}

// GetBatchFromBillInputUserID 批量查找 票据录入人员ID
func (obj *_FaCustomerReceiptMgr) GetBatchFromBillInputUserID(billInputUserIDs []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`bill_input_user_id` IN (?)", billInputUserIDs).Find(&results).Error

	return
}

// GetFromReceiptType 通过receipt_type获取内容 收款类别 0:核收 1:清缴
func (obj *_FaCustomerReceiptMgr) GetFromReceiptType(receiptType int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_type` = ?", receiptType).Find(&results).Error

	return
}

// GetBatchFromReceiptType 批量查找 收款类别 0:核收 1:清缴
func (obj *_FaCustomerReceiptMgr) GetBatchFromReceiptType(receiptTypes []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_type` IN (?)", receiptTypes).Find(&results).Error

	return
}

// GetFromReceiptMethod 通过receipt_method获取内容 收款方式
func (obj *_FaCustomerReceiptMgr) GetFromReceiptMethod(receiptMethod string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_method` = ?", receiptMethod).Find(&results).Error

	return
}

// GetBatchFromReceiptMethod 批量查找 收款方式
func (obj *_FaCustomerReceiptMgr) GetBatchFromReceiptMethod(receiptMethods []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_method` IN (?)", receiptMethods).Find(&results).Error

	return
}

// GetFromPaymentTime 通过payment_time获取内容 收款时间/支付时间
func (obj *_FaCustomerReceiptMgr) GetFromPaymentTime(paymentTime time.Time) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`payment_time` = ?", paymentTime).Find(&results).Error

	return
}

// GetBatchFromPaymentTime 批量查找 收款时间/支付时间
func (obj *_FaCustomerReceiptMgr) GetBatchFromPaymentTime(paymentTimes []time.Time) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`payment_time` IN (?)", paymentTimes).Find(&results).Error

	return
}

// GetFromMark 通过mark获取内容 标签/标记
func (obj *_FaCustomerReceiptMgr) GetFromMark(mark string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`mark` = ?", mark).Find(&results).Error

	return
}

// GetBatchFromMark 批量查找 标签/标记
func (obj *_FaCustomerReceiptMgr) GetBatchFromMark(marks []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`mark` IN (?)", marks).Find(&results).Error

	return
}

// GetFromAccountAlias 通过account_alias获取内容 账户别称/走账
func (obj *_FaCustomerReceiptMgr) GetFromAccountAlias(accountAlias string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`account_alias` = ?", accountAlias).Find(&results).Error

	return
}

// GetBatchFromAccountAlias 批量查找 账户别称/走账
func (obj *_FaCustomerReceiptMgr) GetBatchFromAccountAlias(accountAliass []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`account_alias` IN (?)", accountAliass).Find(&results).Error

	return
}

// GetFromAttachmentURL 通过attachment_url获取内容 附件
func (obj *_FaCustomerReceiptMgr) GetFromAttachmentURL(attachmentURL string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`attachment_url` = ?", attachmentURL).Find(&results).Error

	return
}

// GetBatchFromAttachmentURL 批量查找 附件
func (obj *_FaCustomerReceiptMgr) GetBatchFromAttachmentURL(attachmentURLs []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`attachment_url` IN (?)", attachmentURLs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaCustomerReceiptMgr) GetFromRemark(remark string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaCustomerReceiptMgr) GetBatchFromRemark(remarks []string) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaCustomerReceiptMgr) GetFromCreateTime(createTime time.Time) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaCustomerReceiptMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应收人员
func (obj *_FaCustomerReceiptMgr) GetFromCreateUser(createUser int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应收人员
func (obj *_FaCustomerReceiptMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaCustomerReceiptMgr) GetFromUpdateUser(updateUser int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaCustomerReceiptMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaCustomerReceiptMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaCustomerReceiptMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaCustomerReceiptMgr) GetFromVersion(version int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaCustomerReceiptMgr) GetBatchFromVersion(versions []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaCustomerReceiptMgr) GetFromDeleted(deleted int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaCustomerReceiptMgr) GetBatchFromDeleted(deleteds []int) (results []*FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaCustomerReceiptMgr) FetchByPrimaryKey(id int) (result FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUniqueReceiptNumber primary or index 获取唯一内容
func (obj *_FaCustomerReceiptMgr) FetchUniqueByUniqueReceiptNumber(receiptNumber string) (result FaCustomerReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerReceipt{}).Where("`receipt_number` = ?", receiptNumber).First(&result).Error

	return
}
