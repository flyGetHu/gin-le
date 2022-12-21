package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _FaPaymentOrderInfoMgr struct {
	*_BaseMgr
}

// FaPaymentOrderInfoMgr open func
func FaPaymentOrderInfoMgr(db *gorm.DB) *_FaPaymentOrderInfoMgr {
	if db == nil {
		panic(fmt.Errorf("FaPaymentOrderInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaPaymentOrderInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_payment_order_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaPaymentOrderInfoMgr) GetTableName() string {
	return "fa_payment_order_info"
}

// Reset 重置gorm会话
func (obj *_FaPaymentOrderInfoMgr) Reset() *_FaPaymentOrderInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaPaymentOrderInfoMgr) Get() (result FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaPaymentOrderInfoMgr) Gets() (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaPaymentOrderInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_FaPaymentOrderInfoMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSystemFlowingNumber system_flowing_number获取 系统流水号
func (obj *_FaPaymentOrderInfoMgr) WithSystemFlowingNumber(systemFlowingNumber string) Option {
	return optionFunc(func(o *options) { o.query["system_flowing_number"] = systemFlowingNumber })
}

// WithOrderFlowingNumber order_flowing_number获取 订单流水号
func (obj *_FaPaymentOrderInfoMgr) WithOrderFlowingNumber(orderFlowingNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_flowing_number"] = orderFlowingNumber })
}

// WithPayAmount pay_amount获取 支付金额
func (obj *_FaPaymentOrderInfoMgr) WithPayAmount(payAmount float64) Option {
	return optionFunc(func(o *options) { o.query["pay_amount"] = payAmount })
}

// WithPayType pay_type获取 支付方式（0-微信支付，1-支付宝支付，2-线下支付）
func (obj *_FaPaymentOrderInfoMgr) WithPayType(payType int) Option {
	return optionFunc(func(o *options) { o.query["pay_type"] = payType })
}

// WithPayStatus pay_status获取 支付状态（0-支付中，1-支付完成，2-转入退款，3-未支付，4-已关闭，5-已撤销（付款码支付），6-用户支付中（付款码支付），7-支付失败(其他原因，如银行返回失败)）
func (obj *_FaPaymentOrderInfoMgr) WithPayStatus(payStatus int) Option {
	return optionFunc(func(o *options) { o.query["pay_status"] = payStatus })
}

// WithPayTime pay_time获取 支付完成时间
func (obj *_FaPaymentOrderInfoMgr) WithPayTime(payTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["pay_time"] = payTime })
}

// WithThirdPartyNumber third_party_number获取 第三方单号
func (obj *_FaPaymentOrderInfoMgr) WithThirdPartyNumber(thirdPartyNumber string) Option {
	return optionFunc(func(o *options) { o.query["third_party_number"] = thirdPartyNumber })
}

// WithMerchantNumber merchant_number获取 商户订单号
func (obj *_FaPaymentOrderInfoMgr) WithMerchantNumber(merchantNumber string) Option {
	return optionFunc(func(o *options) { o.query["merchant_number"] = merchantNumber })
}

// WithCustomerID customer_id获取 客户id
func (obj *_FaPaymentOrderInfoMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerName customer_name获取 客户名称
func (obj *_FaPaymentOrderInfoMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithRequestSource request_source获取 请求来源（0：内部，1：外部）
func (obj *_FaPaymentOrderInfoMgr) WithRequestSource(requestSource string) Option {
	return optionFunc(func(o *options) { o.query["request_source"] = requestSource })
}

// WithCallbackStatus callback_status获取 回调状态（0：未回调，1：已回调）
func (obj *_FaPaymentOrderInfoMgr) WithCallbackStatus(callbackStatus int) Option {
	return optionFunc(func(o *options) { o.query["callback_status"] = callbackStatus })
}

// WithExtra extra获取 扩展字段
func (obj *_FaPaymentOrderInfoMgr) WithExtra(extra datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithRemarks remarks获取 备注信息
func (obj *_FaPaymentOrderInfoMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaPaymentOrderInfoMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaPaymentOrderInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_FaPaymentOrderInfoMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaPaymentOrderInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaPaymentOrderInfoMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithIsDelete is_delete获取 逻辑删除
func (obj *_FaPaymentOrderInfoMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// GetByOption 功能选项模式获取
func (obj *_FaPaymentOrderInfoMgr) GetByOption(opts ...Option) (result FaPaymentOrderInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaPaymentOrderInfoMgr) GetByOptions(opts ...Option) (results []*FaPaymentOrderInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaPaymentOrderInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaPaymentOrderInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where(options.query)
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

// GetFromID 通过id获取内容 ID
func (obj *_FaPaymentOrderInfoMgr) GetFromID(id int) (result FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromID(ids []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSystemFlowingNumber 通过system_flowing_number获取内容 系统流水号
func (obj *_FaPaymentOrderInfoMgr) GetFromSystemFlowingNumber(systemFlowingNumber string) (result FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`system_flowing_number` = ?", systemFlowingNumber).First(&result).Error

	return
}

// GetBatchFromSystemFlowingNumber 批量查找 系统流水号
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromSystemFlowingNumber(systemFlowingNumbers []string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`system_flowing_number` IN (?)", systemFlowingNumbers).Find(&results).Error

	return
}

// GetFromOrderFlowingNumber 通过order_flowing_number获取内容 订单流水号
func (obj *_FaPaymentOrderInfoMgr) GetFromOrderFlowingNumber(orderFlowingNumber string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`order_flowing_number` = ?", orderFlowingNumber).Find(&results).Error

	return
}

// GetBatchFromOrderFlowingNumber 批量查找 订单流水号
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromOrderFlowingNumber(orderFlowingNumbers []string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`order_flowing_number` IN (?)", orderFlowingNumbers).Find(&results).Error

	return
}

// GetFromPayAmount 通过pay_amount获取内容 支付金额
func (obj *_FaPaymentOrderInfoMgr) GetFromPayAmount(payAmount float64) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_amount` = ?", payAmount).Find(&results).Error

	return
}

// GetBatchFromPayAmount 批量查找 支付金额
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromPayAmount(payAmounts []float64) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_amount` IN (?)", payAmounts).Find(&results).Error

	return
}

// GetFromPayType 通过pay_type获取内容 支付方式（0-微信支付，1-支付宝支付，2-线下支付）
func (obj *_FaPaymentOrderInfoMgr) GetFromPayType(payType int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_type` = ?", payType).Find(&results).Error

	return
}

// GetBatchFromPayType 批量查找 支付方式（0-微信支付，1-支付宝支付，2-线下支付）
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromPayType(payTypes []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_type` IN (?)", payTypes).Find(&results).Error

	return
}

// GetFromPayStatus 通过pay_status获取内容 支付状态（0-支付中，1-支付完成，2-转入退款，3-未支付，4-已关闭，5-已撤销（付款码支付），6-用户支付中（付款码支付），7-支付失败(其他原因，如银行返回失败)）
func (obj *_FaPaymentOrderInfoMgr) GetFromPayStatus(payStatus int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_status` = ?", payStatus).Find(&results).Error

	return
}

// GetBatchFromPayStatus 批量查找 支付状态（0-支付中，1-支付完成，2-转入退款，3-未支付，4-已关闭，5-已撤销（付款码支付），6-用户支付中（付款码支付），7-支付失败(其他原因，如银行返回失败)）
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromPayStatus(payStatuss []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_status` IN (?)", payStatuss).Find(&results).Error

	return
}

// GetFromPayTime 通过pay_time获取内容 支付完成时间
func (obj *_FaPaymentOrderInfoMgr) GetFromPayTime(payTime time.Time) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_time` = ?", payTime).Find(&results).Error

	return
}

// GetBatchFromPayTime 批量查找 支付完成时间
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromPayTime(payTimes []time.Time) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_time` IN (?)", payTimes).Find(&results).Error

	return
}

// GetFromThirdPartyNumber 通过third_party_number获取内容 第三方单号
func (obj *_FaPaymentOrderInfoMgr) GetFromThirdPartyNumber(thirdPartyNumber string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`third_party_number` = ?", thirdPartyNumber).Find(&results).Error

	return
}

// GetBatchFromThirdPartyNumber 批量查找 第三方单号
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromThirdPartyNumber(thirdPartyNumbers []string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`third_party_number` IN (?)", thirdPartyNumbers).Find(&results).Error

	return
}

// GetFromMerchantNumber 通过merchant_number获取内容 商户订单号
func (obj *_FaPaymentOrderInfoMgr) GetFromMerchantNumber(merchantNumber string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`merchant_number` = ?", merchantNumber).Find(&results).Error

	return
}

// GetBatchFromMerchantNumber 批量查找 商户订单号
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromMerchantNumber(merchantNumbers []string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`merchant_number` IN (?)", merchantNumbers).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_FaPaymentOrderInfoMgr) GetFromCustomerID(customerID int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 客户名称
func (obj *_FaPaymentOrderInfoMgr) GetFromCustomerName(customerName string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 客户名称
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromCustomerName(customerNames []string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

// GetFromRequestSource 通过request_source获取内容 请求来源（0：内部，1：外部）
func (obj *_FaPaymentOrderInfoMgr) GetFromRequestSource(requestSource string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`request_source` = ?", requestSource).Find(&results).Error

	return
}

// GetBatchFromRequestSource 批量查找 请求来源（0：内部，1：外部）
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromRequestSource(requestSources []string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`request_source` IN (?)", requestSources).Find(&results).Error

	return
}

// GetFromCallbackStatus 通过callback_status获取内容 回调状态（0：未回调，1：已回调）
func (obj *_FaPaymentOrderInfoMgr) GetFromCallbackStatus(callbackStatus int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`callback_status` = ?", callbackStatus).Find(&results).Error

	return
}

// GetBatchFromCallbackStatus 批量查找 回调状态（0：未回调，1：已回调）
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromCallbackStatus(callbackStatuss []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`callback_status` IN (?)", callbackStatuss).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 扩展字段
func (obj *_FaPaymentOrderInfoMgr) GetFromExtra(extra datatypes.JSON) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 扩展字段
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromExtra(extras []datatypes.JSON) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注信息
func (obj *_FaPaymentOrderInfoMgr) GetFromRemarks(remarks string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注信息
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromRemarks(remarkss []string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaPaymentOrderInfoMgr) GetFromCreateUser(createUser int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaPaymentOrderInfoMgr) GetFromCreateTime(createTime time.Time) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_FaPaymentOrderInfoMgr) GetFromUpdateUser(updateUser int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaPaymentOrderInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaPaymentOrderInfoMgr) GetFromVersion(version int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromVersion(versions []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 逻辑删除
func (obj *_FaPaymentOrderInfoMgr) GetFromIsDelete(isDelete int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 逻辑删除
func (obj *_FaPaymentOrderInfoMgr) GetBatchFromIsDelete(isDeletes []int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaPaymentOrderInfoMgr) FetchByPrimaryKey(id int) (result FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueBySystemFlowingNumber primary or index 获取唯一内容
func (obj *_FaPaymentOrderInfoMgr) FetchUniqueBySystemFlowingNumber(systemFlowingNumber string) (result FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`system_flowing_number` = ?", systemFlowingNumber).First(&result).Error

	return
}

// FetchIndexByFaPaymentOrderInfoOrderFlowingNumberIndex  获取多个内容
func (obj *_FaPaymentOrderInfoMgr) FetchIndexByFaPaymentOrderInfoOrderFlowingNumberIndex(orderFlowingNumber string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`order_flowing_number` = ?", orderFlowingNumber).Find(&results).Error

	return
}

// FetchIndexByFaPaymentOrderInfoPayTypeIndex  获取多个内容
func (obj *_FaPaymentOrderInfoMgr) FetchIndexByFaPaymentOrderInfoPayTypeIndex(payType int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_type` = ?", payType).Find(&results).Error

	return
}

// FetchIndexByFaPaymentOrderInfoPayStatusIndex  获取多个内容
func (obj *_FaPaymentOrderInfoMgr) FetchIndexByFaPaymentOrderInfoPayStatusIndex(payStatus int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`pay_status` = ?", payStatus).Find(&results).Error

	return
}

// FetchIndexByFaPaymentOrderInfoMerchantNumberIndex  获取多个内容
func (obj *_FaPaymentOrderInfoMgr) FetchIndexByFaPaymentOrderInfoMerchantNumberIndex(merchantNumber string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`merchant_number` = ?", merchantNumber).Find(&results).Error

	return
}

// FetchIndexByFaPaymentOrderInfoCustomerIDIndex  获取多个内容
func (obj *_FaPaymentOrderInfoMgr) FetchIndexByFaPaymentOrderInfoCustomerIDIndex(customerID int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexByFaPaymentOrderInfoRequestSourceIndex  获取多个内容
func (obj *_FaPaymentOrderInfoMgr) FetchIndexByFaPaymentOrderInfoRequestSourceIndex(requestSource string) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`request_source` = ?", requestSource).Find(&results).Error

	return
}

// FetchIndexByFaPaymentOrderInfoCallbackStatusIndex  获取多个内容
func (obj *_FaPaymentOrderInfoMgr) FetchIndexByFaPaymentOrderInfoCallbackStatusIndex(callbackStatus int) (results []*FaPaymentOrderInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPaymentOrderInfo{}).Where("`callback_status` = ?", callbackStatus).Find(&results).Error

	return
}
