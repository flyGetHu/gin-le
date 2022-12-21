package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaBusinessEntityPaymentRecordMgr struct {
	*_BaseMgr
}

// FaBusinessEntityPaymentRecordMgr open func
func FaBusinessEntityPaymentRecordMgr(db *gorm.DB) *_FaBusinessEntityPaymentRecordMgr {
	if db == nil {
		panic(fmt.Errorf("FaBusinessEntityPaymentRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaBusinessEntityPaymentRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_business_entity_payment_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaBusinessEntityPaymentRecordMgr) GetTableName() string {
	return "fa_business_entity_payment_record"
}

// Reset 重置gorm会话
func (obj *_FaBusinessEntityPaymentRecordMgr) Reset() *_FaBusinessEntityPaymentRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaBusinessEntityPaymentRecordMgr) Get() (result FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaBusinessEntityPaymentRecordMgr) Gets() (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaBusinessEntityPaymentRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaBusinessEntityPaymentRecordMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPaymentNumber payment_number获取 付款单号
func (obj *_FaBusinessEntityPaymentRecordMgr) WithPaymentNumber(paymentNumber string) Option {
	return optionFunc(func(o *options) { o.query["payment_number"] = paymentNumber })
}

// WithInvoiceNumber invoice_number获取 发票号
func (obj *_FaBusinessEntityPaymentRecordMgr) WithInvoiceNumber(invoiceNumber string) Option {
	return optionFunc(func(o *options) { o.query["invoice_number"] = invoiceNumber })
}

// WithVoucherNumber voucher_number获取 凭证号
func (obj *_FaBusinessEntityPaymentRecordMgr) WithVoucherNumber(voucherNumber string) Option {
	return optionFunc(func(o *options) { o.query["voucher_number"] = voucherNumber })
}

// WithTicketNumber ticket_number获取 票证号
func (obj *_FaBusinessEntityPaymentRecordMgr) WithTicketNumber(ticketNumber string) Option {
	return optionFunc(func(o *options) { o.query["ticket_number"] = ticketNumber })
}

// WithBusinessEntityCode business_entity_code获取 业务主体代码
func (obj *_FaBusinessEntityPaymentRecordMgr) WithBusinessEntityCode(businessEntityCode string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_code"] = businessEntityCode })
}

// WithBusinessEntityName business_entity_name获取 业务实体名称
func (obj *_FaBusinessEntityPaymentRecordMgr) WithBusinessEntityName(businessEntityName string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_name"] = businessEntityName })
}

// WithBusinessEntityType business_entity_type获取 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentRecordMgr) WithBusinessEntityType(businessEntityType bool) Option {
	return optionFunc(func(o *options) { o.query["business_entity_type"] = businessEntityType })
}

// WithPayableFeeRmb payable_fee_rmb获取 应付金额
func (obj *_FaBusinessEntityPaymentRecordMgr) WithPayableFeeRmb(payableFeeRmb float64) Option {
	return optionFunc(func(o *options) { o.query["payable_fee_rmb"] = payableFeeRmb })
}

// WithActualPaymentFee actual_payment_fee获取 实付金额
func (obj *_FaBusinessEntityPaymentRecordMgr) WithActualPaymentFee(actualPaymentFee float64) Option {
	return optionFunc(func(o *options) { o.query["actual_payment_fee"] = actualPaymentFee })
}

// WithCurrencyCode currency_code获取 币种
func (obj *_FaBusinessEntityPaymentRecordMgr) WithCurrencyCode(currencyCode string) Option {
	return optionFunc(func(o *options) { o.query["currency_code"] = currencyCode })
}

// WithCurrencyName currency_name获取 币种中文名称
func (obj *_FaBusinessEntityPaymentRecordMgr) WithCurrencyName(currencyName string) Option {
	return optionFunc(func(o *options) { o.query["currency_name"] = currencyName })
}

// WithExchangeRate exchange_rate获取 汇率
func (obj *_FaBusinessEntityPaymentRecordMgr) WithExchangeRate(exchangeRate float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_rate"] = exchangeRate })
}

// WithActualPaymentFeeRmb actual_payment_fee_rmb获取 实付金额(人民币)
func (obj *_FaBusinessEntityPaymentRecordMgr) WithActualPaymentFeeRmb(actualPaymentFeeRmb float64) Option {
	return optionFunc(func(o *options) { o.query["actual_payment_fee_rmb"] = actualPaymentFeeRmb })
}

// WithIsPaid is_paid获取 已付清
func (obj *_FaBusinessEntityPaymentRecordMgr) WithIsPaid(isPaid []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_paid"] = isPaid })
}

// WithPaymentMethod payment_method获取 付款方式
func (obj *_FaBusinessEntityPaymentRecordMgr) WithPaymentMethod(paymentMethod string) Option {
	return optionFunc(func(o *options) { o.query["payment_method"] = paymentMethod })
}

// WithPaymentAccount payment_account获取 付款账户
func (obj *_FaBusinessEntityPaymentRecordMgr) WithPaymentAccount(paymentAccount string) Option {
	return optionFunc(func(o *options) { o.query["payment_account"] = paymentAccount })
}

// WithPaymentAccountAlias payment_account_alias获取 付款账户别称
func (obj *_FaBusinessEntityPaymentRecordMgr) WithPaymentAccountAlias(paymentAccountAlias string) Option {
	return optionFunc(func(o *options) { o.query["payment_account_alias"] = paymentAccountAlias })
}

// WithPaymentTime payment_time获取 支付时间
func (obj *_FaBusinessEntityPaymentRecordMgr) WithPaymentTime(paymentTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["payment_time"] = paymentTime })
}

// WithDealWithUser deal_with_user获取 经办人
func (obj *_FaBusinessEntityPaymentRecordMgr) WithDealWithUser(dealWithUser string) Option {
	return optionFunc(func(o *options) { o.query["deal_with_user"] = dealWithUser })
}

// WithAttachmentURL attachment_url获取 附件
func (obj *_FaBusinessEntityPaymentRecordMgr) WithAttachmentURL(attachmentURL string) Option {
	return optionFunc(func(o *options) { o.query["attachment_url"] = attachmentURL })
}

// WithRemark remark获取 备注
func (obj *_FaBusinessEntityPaymentRecordMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaBusinessEntityPaymentRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 流水生成用户
func (obj *_FaBusinessEntityPaymentRecordMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaBusinessEntityPaymentRecordMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaBusinessEntityPaymentRecordMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaBusinessEntityPaymentRecordMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaBusinessEntityPaymentRecordMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaBusinessEntityPaymentRecordMgr) GetByOption(opts ...Option) (result FaBusinessEntityPaymentRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaBusinessEntityPaymentRecordMgr) GetByOptions(opts ...Option) (results []*FaBusinessEntityPaymentRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaBusinessEntityPaymentRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaBusinessEntityPaymentRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where(options.query)
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
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromID(id int) (result FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromID(ids []int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPaymentNumber 通过payment_number获取内容 付款单号
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromPaymentNumber(paymentNumber string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_number` = ?", paymentNumber).Find(&results).Error

	return
}

// GetBatchFromPaymentNumber 批量查找 付款单号
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromPaymentNumber(paymentNumbers []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_number` IN (?)", paymentNumbers).Find(&results).Error

	return
}

// GetFromInvoiceNumber 通过invoice_number获取内容 发票号
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromInvoiceNumber(invoiceNumber string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`invoice_number` = ?", invoiceNumber).Find(&results).Error

	return
}

// GetBatchFromInvoiceNumber 批量查找 发票号
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromInvoiceNumber(invoiceNumbers []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`invoice_number` IN (?)", invoiceNumbers).Find(&results).Error

	return
}

// GetFromVoucherNumber 通过voucher_number获取内容 凭证号
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromVoucherNumber(voucherNumber string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`voucher_number` = ?", voucherNumber).Find(&results).Error

	return
}

// GetBatchFromVoucherNumber 批量查找 凭证号
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromVoucherNumber(voucherNumbers []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`voucher_number` IN (?)", voucherNumbers).Find(&results).Error

	return
}

// GetFromTicketNumber 通过ticket_number获取内容 票证号
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromTicketNumber(ticketNumber string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`ticket_number` = ?", ticketNumber).Find(&results).Error

	return
}

// GetBatchFromTicketNumber 批量查找 票证号
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromTicketNumber(ticketNumbers []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`ticket_number` IN (?)", ticketNumbers).Find(&results).Error

	return
}

// GetFromBusinessEntityCode 通过business_entity_code获取内容 业务主体代码
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromBusinessEntityCode(businessEntityCode string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityCode 批量查找 业务主体代码
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromBusinessEntityCode(businessEntityCodes []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`business_entity_code` IN (?)", businessEntityCodes).Find(&results).Error

	return
}

// GetFromBusinessEntityName 通过business_entity_name获取内容 业务实体名称
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromBusinessEntityName(businessEntityName string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`business_entity_name` = ?", businessEntityName).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityName 批量查找 业务实体名称
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromBusinessEntityName(businessEntityNames []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`business_entity_name` IN (?)", businessEntityNames).Find(&results).Error

	return
}

// GetFromBusinessEntityType 通过business_entity_type获取内容 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromBusinessEntityType(businessEntityType bool) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`business_entity_type` = ?", businessEntityType).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityType 批量查找 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromBusinessEntityType(businessEntityTypes []bool) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`business_entity_type` IN (?)", businessEntityTypes).Find(&results).Error

	return
}

// GetFromPayableFeeRmb 通过payable_fee_rmb获取内容 应付金额
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromPayableFeeRmb(payableFeeRmb float64) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payable_fee_rmb` = ?", payableFeeRmb).Find(&results).Error

	return
}

// GetBatchFromPayableFeeRmb 批量查找 应付金额
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromPayableFeeRmb(payableFeeRmbs []float64) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payable_fee_rmb` IN (?)", payableFeeRmbs).Find(&results).Error

	return
}

// GetFromActualPaymentFee 通过actual_payment_fee获取内容 实付金额
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromActualPaymentFee(actualPaymentFee float64) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`actual_payment_fee` = ?", actualPaymentFee).Find(&results).Error

	return
}

// GetBatchFromActualPaymentFee 批量查找 实付金额
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromActualPaymentFee(actualPaymentFees []float64) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`actual_payment_fee` IN (?)", actualPaymentFees).Find(&results).Error

	return
}

// GetFromCurrencyCode 通过currency_code获取内容 币种
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromCurrencyCode(currencyCode string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`currency_code` = ?", currencyCode).Find(&results).Error

	return
}

// GetBatchFromCurrencyCode 批量查找 币种
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromCurrencyCode(currencyCodes []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`currency_code` IN (?)", currencyCodes).Find(&results).Error

	return
}

// GetFromCurrencyName 通过currency_name获取内容 币种中文名称
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromCurrencyName(currencyName string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`currency_name` = ?", currencyName).Find(&results).Error

	return
}

// GetBatchFromCurrencyName 批量查找 币种中文名称
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromCurrencyName(currencyNames []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`currency_name` IN (?)", currencyNames).Find(&results).Error

	return
}

// GetFromExchangeRate 通过exchange_rate获取内容 汇率
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromExchangeRate(exchangeRate float64) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`exchange_rate` = ?", exchangeRate).Find(&results).Error

	return
}

// GetBatchFromExchangeRate 批量查找 汇率
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromExchangeRate(exchangeRates []float64) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`exchange_rate` IN (?)", exchangeRates).Find(&results).Error

	return
}

// GetFromActualPaymentFeeRmb 通过actual_payment_fee_rmb获取内容 实付金额(人民币)
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromActualPaymentFeeRmb(actualPaymentFeeRmb float64) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`actual_payment_fee_rmb` = ?", actualPaymentFeeRmb).Find(&results).Error

	return
}

// GetBatchFromActualPaymentFeeRmb 批量查找 实付金额(人民币)
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromActualPaymentFeeRmb(actualPaymentFeeRmbs []float64) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`actual_payment_fee_rmb` IN (?)", actualPaymentFeeRmbs).Find(&results).Error

	return
}

// GetFromIsPaid 通过is_paid获取内容 已付清
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromIsPaid(isPaid []uint8) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`is_paid` = ?", isPaid).Find(&results).Error

	return
}

// GetBatchFromIsPaid 批量查找 已付清
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromIsPaid(isPaids [][]uint8) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`is_paid` IN (?)", isPaids).Find(&results).Error

	return
}

// GetFromPaymentMethod 通过payment_method获取内容 付款方式
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromPaymentMethod(paymentMethod string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_method` = ?", paymentMethod).Find(&results).Error

	return
}

// GetBatchFromPaymentMethod 批量查找 付款方式
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromPaymentMethod(paymentMethods []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_method` IN (?)", paymentMethods).Find(&results).Error

	return
}

// GetFromPaymentAccount 通过payment_account获取内容 付款账户
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromPaymentAccount(paymentAccount string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_account` = ?", paymentAccount).Find(&results).Error

	return
}

// GetBatchFromPaymentAccount 批量查找 付款账户
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromPaymentAccount(paymentAccounts []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_account` IN (?)", paymentAccounts).Find(&results).Error

	return
}

// GetFromPaymentAccountAlias 通过payment_account_alias获取内容 付款账户别称
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromPaymentAccountAlias(paymentAccountAlias string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_account_alias` = ?", paymentAccountAlias).Find(&results).Error

	return
}

// GetBatchFromPaymentAccountAlias 批量查找 付款账户别称
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromPaymentAccountAlias(paymentAccountAliass []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_account_alias` IN (?)", paymentAccountAliass).Find(&results).Error

	return
}

// GetFromPaymentTime 通过payment_time获取内容 支付时间
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromPaymentTime(paymentTime time.Time) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_time` = ?", paymentTime).Find(&results).Error

	return
}

// GetBatchFromPaymentTime 批量查找 支付时间
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromPaymentTime(paymentTimes []time.Time) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_time` IN (?)", paymentTimes).Find(&results).Error

	return
}

// GetFromDealWithUser 通过deal_with_user获取内容 经办人
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromDealWithUser(dealWithUser string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`deal_with_user` = ?", dealWithUser).Find(&results).Error

	return
}

// GetBatchFromDealWithUser 批量查找 经办人
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromDealWithUser(dealWithUsers []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`deal_with_user` IN (?)", dealWithUsers).Find(&results).Error

	return
}

// GetFromAttachmentURL 通过attachment_url获取内容 附件
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromAttachmentURL(attachmentURL string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`attachment_url` = ?", attachmentURL).Find(&results).Error

	return
}

// GetBatchFromAttachmentURL 批量查找 附件
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromAttachmentURL(attachmentURLs []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`attachment_url` IN (?)", attachmentURLs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromRemark(remark string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromRemark(remarks []string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromCreateTime(createTime time.Time) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 流水生成用户
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromCreateUser(createUser int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 流水生成用户
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromUpdateUser(updateUser int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromVersion(version int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromVersion(versions []int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaBusinessEntityPaymentRecordMgr) GetFromDeleted(deleted int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaBusinessEntityPaymentRecordMgr) GetBatchFromDeleted(deleteds []int) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaBusinessEntityPaymentRecordMgr) FetchByPrimaryKey(id int) (result FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIndexPaymentNumber  获取多个内容
func (obj *_FaBusinessEntityPaymentRecordMgr) FetchIndexByIndexPaymentNumber(paymentNumber string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_number` = ?", paymentNumber).Find(&results).Error

	return
}

// FetchIndexByIndexBusinessEntityCode  获取多个内容
func (obj *_FaBusinessEntityPaymentRecordMgr) FetchIndexByIndexBusinessEntityCode(businessEntityCode string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// FetchIndexByIndexPaymentAccount  获取多个内容
func (obj *_FaBusinessEntityPaymentRecordMgr) FetchIndexByIndexPaymentAccount(paymentAccount string) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_account` = ?", paymentAccount).Find(&results).Error

	return
}

// FetchIndexByIndexPaymentTime  获取多个内容
func (obj *_FaBusinessEntityPaymentRecordMgr) FetchIndexByIndexPaymentTime(paymentTime time.Time) (results []*FaBusinessEntityPaymentRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentRecord{}).Where("`payment_time` = ?", paymentTime).Find(&results).Error

	return
}
