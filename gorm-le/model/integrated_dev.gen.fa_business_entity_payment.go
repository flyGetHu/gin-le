package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaBusinessEntityPaymentMgr struct {
	*_BaseMgr
}

// FaBusinessEntityPaymentMgr open func
func FaBusinessEntityPaymentMgr(db *gorm.DB) *_FaBusinessEntityPaymentMgr {
	if db == nil {
		panic(fmt.Errorf("FaBusinessEntityPaymentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaBusinessEntityPaymentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_business_entity_payment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaBusinessEntityPaymentMgr) GetTableName() string {
	return "fa_business_entity_payment"
}

// Reset 重置gorm会话
func (obj *_FaBusinessEntityPaymentMgr) Reset() *_FaBusinessEntityPaymentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaBusinessEntityPaymentMgr) Get() (result FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaBusinessEntityPaymentMgr) Gets() (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaBusinessEntityPaymentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaBusinessEntityPaymentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPaymentNumber payment_number获取 付款单号
func (obj *_FaBusinessEntityPaymentMgr) WithPaymentNumber(paymentNumber string) Option {
	return optionFunc(func(o *options) { o.query["payment_number"] = paymentNumber })
}

// WithParentID parent_id获取 父级id
func (obj *_FaBusinessEntityPaymentMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithIsParent is_parent获取 是否为父级；0-否，1-是
func (obj *_FaBusinessEntityPaymentMgr) WithIsParent(isParent []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_parent"] = isParent })
}

// WithInvoiceNumber invoice_number获取 发票号
func (obj *_FaBusinessEntityPaymentMgr) WithInvoiceNumber(invoiceNumber string) Option {
	return optionFunc(func(o *options) { o.query["invoice_number"] = invoiceNumber })
}

// WithVoucherNumber voucher_number获取 凭证号
func (obj *_FaBusinessEntityPaymentMgr) WithVoucherNumber(voucherNumber string) Option {
	return optionFunc(func(o *options) { o.query["voucher_number"] = voucherNumber })
}

// WithTicketNumber ticket_number获取 票证号
func (obj *_FaBusinessEntityPaymentMgr) WithTicketNumber(ticketNumber string) Option {
	return optionFunc(func(o *options) { o.query["ticket_number"] = ticketNumber })
}

// WithPaymentStatus payment_status获取 付款状态:0未付,  1已付清, 2未付清
func (obj *_FaBusinessEntityPaymentMgr) WithPaymentStatus(paymentStatus int) Option {
	return optionFunc(func(o *options) { o.query["payment_status"] = paymentStatus })
}

// WithIsPayment is_payment获取 是否已付款
func (obj *_FaBusinessEntityPaymentMgr) WithIsPayment(isPayment []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_payment"] = isPayment })
}

// WithBusinessEntityCode business_entity_code获取 业务主体代码
func (obj *_FaBusinessEntityPaymentMgr) WithBusinessEntityCode(businessEntityCode string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_code"] = businessEntityCode })
}

// WithBusinessEntityName business_entity_name获取 业务实体名称
func (obj *_FaBusinessEntityPaymentMgr) WithBusinessEntityName(businessEntityName string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_name"] = businessEntityName })
}

// WithBusinessEntityType business_entity_type获取 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentMgr) WithBusinessEntityType(businessEntityType bool) Option {
	return optionFunc(func(o *options) { o.query["business_entity_type"] = businessEntityType })
}

// WithBillFee bill_fee获取 票据金额
func (obj *_FaBusinessEntityPaymentMgr) WithBillFee(billFee float64) Option {
	return optionFunc(func(o *options) { o.query["bill_fee"] = billFee })
}

// WithActualFee actual_fee获取 实际付款金额
func (obj *_FaBusinessEntityPaymentMgr) WithActualFee(actualFee float64) Option {
	return optionFunc(func(o *options) { o.query["actual_fee"] = actualFee })
}

// WithCurrencyCode currency_code获取 币种
func (obj *_FaBusinessEntityPaymentMgr) WithCurrencyCode(currencyCode string) Option {
	return optionFunc(func(o *options) { o.query["currency_code"] = currencyCode })
}

// WithCurrencyName currency_name获取 币种中文名称
func (obj *_FaBusinessEntityPaymentMgr) WithCurrencyName(currencyName string) Option {
	return optionFunc(func(o *options) { o.query["currency_name"] = currencyName })
}

// WithExchangeRate exchange_rate获取 汇率
func (obj *_FaBusinessEntityPaymentMgr) WithExchangeRate(exchangeRate float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_rate"] = exchangeRate })
}

// WithActualFeeRmb actual_fee_rmb获取 实际收款金额rmb
func (obj *_FaBusinessEntityPaymentMgr) WithActualFeeRmb(actualFeeRmb float64) Option {
	return optionFunc(func(o *options) { o.query["actual_fee_rmb"] = actualFeeRmb })
}

// WithDealWithUser deal_with_user获取 经办人
func (obj *_FaBusinessEntityPaymentMgr) WithDealWithUser(dealWithUser string) Option {
	return optionFunc(func(o *options) { o.query["deal_with_user"] = dealWithUser })
}

// WithBillInputUser bill_input_user获取 票据录入人员
func (obj *_FaBusinessEntityPaymentMgr) WithBillInputUser(billInputUser string) Option {
	return optionFunc(func(o *options) { o.query["bill_input_user"] = billInputUser })
}

// WithBillInputUserID bill_input_user_id获取 票据录入人员ID
func (obj *_FaBusinessEntityPaymentMgr) WithBillInputUserID(billInputUserID int) Option {
	return optionFunc(func(o *options) { o.query["bill_input_user_id"] = billInputUserID })
}

// WithPaymentType payment_type获取 付款类别 0:冲销，1:核付
func (obj *_FaBusinessEntityPaymentMgr) WithPaymentType(paymentType int) Option {
	return optionFunc(func(o *options) { o.query["payment_type"] = paymentType })
}

// WithPaymentMethod payment_method获取 付款方式
func (obj *_FaBusinessEntityPaymentMgr) WithPaymentMethod(paymentMethod string) Option {
	return optionFunc(func(o *options) { o.query["payment_method"] = paymentMethod })
}

// WithPaymentTime payment_time获取 支付时间
func (obj *_FaBusinessEntityPaymentMgr) WithPaymentTime(paymentTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["payment_time"] = paymentTime })
}

// WithMark mark获取 标签/标记
func (obj *_FaBusinessEntityPaymentMgr) WithMark(mark string) Option {
	return optionFunc(func(o *options) { o.query["mark"] = mark })
}

// WithAccountAlias account_alias获取 账户别称/走账
func (obj *_FaBusinessEntityPaymentMgr) WithAccountAlias(accountAlias string) Option {
	return optionFunc(func(o *options) { o.query["account_alias"] = accountAlias })
}

// WithAttachmentURL attachment_url获取 附件
func (obj *_FaBusinessEntityPaymentMgr) WithAttachmentURL(attachmentURL string) Option {
	return optionFunc(func(o *options) { o.query["attachment_url"] = attachmentURL })
}

// WithRemark remark获取 备注
func (obj *_FaBusinessEntityPaymentMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaBusinessEntityPaymentMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应付人员
func (obj *_FaBusinessEntityPaymentMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaBusinessEntityPaymentMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaBusinessEntityPaymentMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaBusinessEntityPaymentMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaBusinessEntityPaymentMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaBusinessEntityPaymentMgr) GetByOption(opts ...Option) (result FaBusinessEntityPayment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaBusinessEntityPaymentMgr) GetByOptions(opts ...Option) (results []*FaBusinessEntityPayment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaBusinessEntityPaymentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaBusinessEntityPayment, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where(options.query)
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
func (obj *_FaBusinessEntityPaymentMgr) GetFromID(id int) (result FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromID(ids []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPaymentNumber 通过payment_number获取内容 付款单号
func (obj *_FaBusinessEntityPaymentMgr) GetFromPaymentNumber(paymentNumber string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_number` = ?", paymentNumber).Find(&results).Error

	return
}

// GetBatchFromPaymentNumber 批量查找 付款单号
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromPaymentNumber(paymentNumbers []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_number` IN (?)", paymentNumbers).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父级id
func (obj *_FaBusinessEntityPaymentMgr) GetFromParentID(parentID int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父级id
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromParentID(parentIDs []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromIsParent 通过is_parent获取内容 是否为父级；0-否，1-是
func (obj *_FaBusinessEntityPaymentMgr) GetFromIsParent(isParent []uint8) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`is_parent` = ?", isParent).Find(&results).Error

	return
}

// GetBatchFromIsParent 批量查找 是否为父级；0-否，1-是
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromIsParent(isParents [][]uint8) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`is_parent` IN (?)", isParents).Find(&results).Error

	return
}

// GetFromInvoiceNumber 通过invoice_number获取内容 发票号
func (obj *_FaBusinessEntityPaymentMgr) GetFromInvoiceNumber(invoiceNumber string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`invoice_number` = ?", invoiceNumber).Find(&results).Error

	return
}

// GetBatchFromInvoiceNumber 批量查找 发票号
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromInvoiceNumber(invoiceNumbers []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`invoice_number` IN (?)", invoiceNumbers).Find(&results).Error

	return
}

// GetFromVoucherNumber 通过voucher_number获取内容 凭证号
func (obj *_FaBusinessEntityPaymentMgr) GetFromVoucherNumber(voucherNumber string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`voucher_number` = ?", voucherNumber).Find(&results).Error

	return
}

// GetBatchFromVoucherNumber 批量查找 凭证号
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromVoucherNumber(voucherNumbers []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`voucher_number` IN (?)", voucherNumbers).Find(&results).Error

	return
}

// GetFromTicketNumber 通过ticket_number获取内容 票证号
func (obj *_FaBusinessEntityPaymentMgr) GetFromTicketNumber(ticketNumber string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`ticket_number` = ?", ticketNumber).Find(&results).Error

	return
}

// GetBatchFromTicketNumber 批量查找 票证号
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromTicketNumber(ticketNumbers []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`ticket_number` IN (?)", ticketNumbers).Find(&results).Error

	return
}

// GetFromPaymentStatus 通过payment_status获取内容 付款状态:0未付,  1已付清, 2未付清
func (obj *_FaBusinessEntityPaymentMgr) GetFromPaymentStatus(paymentStatus int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_status` = ?", paymentStatus).Find(&results).Error

	return
}

// GetBatchFromPaymentStatus 批量查找 付款状态:0未付,  1已付清, 2未付清
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromPaymentStatus(paymentStatuss []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_status` IN (?)", paymentStatuss).Find(&results).Error

	return
}

// GetFromIsPayment 通过is_payment获取内容 是否已付款
func (obj *_FaBusinessEntityPaymentMgr) GetFromIsPayment(isPayment []uint8) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`is_payment` = ?", isPayment).Find(&results).Error

	return
}

// GetBatchFromIsPayment 批量查找 是否已付款
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromIsPayment(isPayments [][]uint8) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`is_payment` IN (?)", isPayments).Find(&results).Error

	return
}

// GetFromBusinessEntityCode 通过business_entity_code获取内容 业务主体代码
func (obj *_FaBusinessEntityPaymentMgr) GetFromBusinessEntityCode(businessEntityCode string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityCode 批量查找 业务主体代码
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromBusinessEntityCode(businessEntityCodes []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`business_entity_code` IN (?)", businessEntityCodes).Find(&results).Error

	return
}

// GetFromBusinessEntityName 通过business_entity_name获取内容 业务实体名称
func (obj *_FaBusinessEntityPaymentMgr) GetFromBusinessEntityName(businessEntityName string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`business_entity_name` = ?", businessEntityName).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityName 批量查找 业务实体名称
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromBusinessEntityName(businessEntityNames []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`business_entity_name` IN (?)", businessEntityNames).Find(&results).Error

	return
}

// GetFromBusinessEntityType 通过business_entity_type获取内容 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentMgr) GetFromBusinessEntityType(businessEntityType bool) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`business_entity_type` = ?", businessEntityType).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityType 批量查找 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromBusinessEntityType(businessEntityTypes []bool) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`business_entity_type` IN (?)", businessEntityTypes).Find(&results).Error

	return
}

// GetFromBillFee 通过bill_fee获取内容 票据金额
func (obj *_FaBusinessEntityPaymentMgr) GetFromBillFee(billFee float64) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`bill_fee` = ?", billFee).Find(&results).Error

	return
}

// GetBatchFromBillFee 批量查找 票据金额
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromBillFee(billFees []float64) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`bill_fee` IN (?)", billFees).Find(&results).Error

	return
}

// GetFromActualFee 通过actual_fee获取内容 实际付款金额
func (obj *_FaBusinessEntityPaymentMgr) GetFromActualFee(actualFee float64) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`actual_fee` = ?", actualFee).Find(&results).Error

	return
}

// GetBatchFromActualFee 批量查找 实际付款金额
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromActualFee(actualFees []float64) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`actual_fee` IN (?)", actualFees).Find(&results).Error

	return
}

// GetFromCurrencyCode 通过currency_code获取内容 币种
func (obj *_FaBusinessEntityPaymentMgr) GetFromCurrencyCode(currencyCode string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`currency_code` = ?", currencyCode).Find(&results).Error

	return
}

// GetBatchFromCurrencyCode 批量查找 币种
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromCurrencyCode(currencyCodes []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`currency_code` IN (?)", currencyCodes).Find(&results).Error

	return
}

// GetFromCurrencyName 通过currency_name获取内容 币种中文名称
func (obj *_FaBusinessEntityPaymentMgr) GetFromCurrencyName(currencyName string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`currency_name` = ?", currencyName).Find(&results).Error

	return
}

// GetBatchFromCurrencyName 批量查找 币种中文名称
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromCurrencyName(currencyNames []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`currency_name` IN (?)", currencyNames).Find(&results).Error

	return
}

// GetFromExchangeRate 通过exchange_rate获取内容 汇率
func (obj *_FaBusinessEntityPaymentMgr) GetFromExchangeRate(exchangeRate float64) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`exchange_rate` = ?", exchangeRate).Find(&results).Error

	return
}

// GetBatchFromExchangeRate 批量查找 汇率
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromExchangeRate(exchangeRates []float64) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`exchange_rate` IN (?)", exchangeRates).Find(&results).Error

	return
}

// GetFromActualFeeRmb 通过actual_fee_rmb获取内容 实际收款金额rmb
func (obj *_FaBusinessEntityPaymentMgr) GetFromActualFeeRmb(actualFeeRmb float64) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`actual_fee_rmb` = ?", actualFeeRmb).Find(&results).Error

	return
}

// GetBatchFromActualFeeRmb 批量查找 实际收款金额rmb
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromActualFeeRmb(actualFeeRmbs []float64) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`actual_fee_rmb` IN (?)", actualFeeRmbs).Find(&results).Error

	return
}

// GetFromDealWithUser 通过deal_with_user获取内容 经办人
func (obj *_FaBusinessEntityPaymentMgr) GetFromDealWithUser(dealWithUser string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`deal_with_user` = ?", dealWithUser).Find(&results).Error

	return
}

// GetBatchFromDealWithUser 批量查找 经办人
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromDealWithUser(dealWithUsers []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`deal_with_user` IN (?)", dealWithUsers).Find(&results).Error

	return
}

// GetFromBillInputUser 通过bill_input_user获取内容 票据录入人员
func (obj *_FaBusinessEntityPaymentMgr) GetFromBillInputUser(billInputUser string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`bill_input_user` = ?", billInputUser).Find(&results).Error

	return
}

// GetBatchFromBillInputUser 批量查找 票据录入人员
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromBillInputUser(billInputUsers []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`bill_input_user` IN (?)", billInputUsers).Find(&results).Error

	return
}

// GetFromBillInputUserID 通过bill_input_user_id获取内容 票据录入人员ID
func (obj *_FaBusinessEntityPaymentMgr) GetFromBillInputUserID(billInputUserID int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`bill_input_user_id` = ?", billInputUserID).Find(&results).Error

	return
}

// GetBatchFromBillInputUserID 批量查找 票据录入人员ID
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromBillInputUserID(billInputUserIDs []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`bill_input_user_id` IN (?)", billInputUserIDs).Find(&results).Error

	return
}

// GetFromPaymentType 通过payment_type获取内容 付款类别 0:冲销，1:核付
func (obj *_FaBusinessEntityPaymentMgr) GetFromPaymentType(paymentType int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_type` = ?", paymentType).Find(&results).Error

	return
}

// GetBatchFromPaymentType 批量查找 付款类别 0:冲销，1:核付
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromPaymentType(paymentTypes []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_type` IN (?)", paymentTypes).Find(&results).Error

	return
}

// GetFromPaymentMethod 通过payment_method获取内容 付款方式
func (obj *_FaBusinessEntityPaymentMgr) GetFromPaymentMethod(paymentMethod string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_method` = ?", paymentMethod).Find(&results).Error

	return
}

// GetBatchFromPaymentMethod 批量查找 付款方式
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromPaymentMethod(paymentMethods []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_method` IN (?)", paymentMethods).Find(&results).Error

	return
}

// GetFromPaymentTime 通过payment_time获取内容 支付时间
func (obj *_FaBusinessEntityPaymentMgr) GetFromPaymentTime(paymentTime time.Time) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_time` = ?", paymentTime).Find(&results).Error

	return
}

// GetBatchFromPaymentTime 批量查找 支付时间
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromPaymentTime(paymentTimes []time.Time) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`payment_time` IN (?)", paymentTimes).Find(&results).Error

	return
}

// GetFromMark 通过mark获取内容 标签/标记
func (obj *_FaBusinessEntityPaymentMgr) GetFromMark(mark string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`mark` = ?", mark).Find(&results).Error

	return
}

// GetBatchFromMark 批量查找 标签/标记
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromMark(marks []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`mark` IN (?)", marks).Find(&results).Error

	return
}

// GetFromAccountAlias 通过account_alias获取内容 账户别称/走账
func (obj *_FaBusinessEntityPaymentMgr) GetFromAccountAlias(accountAlias string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`account_alias` = ?", accountAlias).Find(&results).Error

	return
}

// GetBatchFromAccountAlias 批量查找 账户别称/走账
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromAccountAlias(accountAliass []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`account_alias` IN (?)", accountAliass).Find(&results).Error

	return
}

// GetFromAttachmentURL 通过attachment_url获取内容 附件
func (obj *_FaBusinessEntityPaymentMgr) GetFromAttachmentURL(attachmentURL string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`attachment_url` = ?", attachmentURL).Find(&results).Error

	return
}

// GetBatchFromAttachmentURL 批量查找 附件
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromAttachmentURL(attachmentURLs []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`attachment_url` IN (?)", attachmentURLs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaBusinessEntityPaymentMgr) GetFromRemark(remark string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromRemark(remarks []string) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaBusinessEntityPaymentMgr) GetFromCreateTime(createTime time.Time) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应付人员
func (obj *_FaBusinessEntityPaymentMgr) GetFromCreateUser(createUser int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应付人员
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaBusinessEntityPaymentMgr) GetFromUpdateUser(updateUser int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaBusinessEntityPaymentMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaBusinessEntityPaymentMgr) GetFromVersion(version int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromVersion(versions []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaBusinessEntityPaymentMgr) GetFromDeleted(deleted int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaBusinessEntityPaymentMgr) GetBatchFromDeleted(deleteds []int) (results []*FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaBusinessEntityPaymentMgr) FetchByPrimaryKey(id int) (result FaBusinessEntityPayment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPayment{}).Where("`id` = ?", id).First(&result).Error

	return
}
