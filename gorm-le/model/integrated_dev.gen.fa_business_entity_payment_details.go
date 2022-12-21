package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaBusinessEntityPaymentDetailsMgr struct {
	*_BaseMgr
}

// FaBusinessEntityPaymentDetailsMgr open func
func FaBusinessEntityPaymentDetailsMgr(db *gorm.DB) *_FaBusinessEntityPaymentDetailsMgr {
	if db == nil {
		panic(fmt.Errorf("FaBusinessEntityPaymentDetailsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaBusinessEntityPaymentDetailsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_business_entity_payment_details"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetTableName() string {
	return "fa_business_entity_payment_details"
}

// Reset 重置gorm会话
func (obj *_FaBusinessEntityPaymentDetailsMgr) Reset() *_FaBusinessEntityPaymentDetailsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaBusinessEntityPaymentDetailsMgr) Get() (result FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaBusinessEntityPaymentDetailsMgr) Gets() (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaBusinessEntityPaymentDetailsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPaymentNumber payment_number获取 付款单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithPaymentNumber(paymentNumber string) Option {
	return optionFunc(func(o *options) { o.query["payment_number"] = paymentNumber })
}

// WithOrderNumber order_number获取 订单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithBusinessEntityCode business_entity_code获取 业务主体代码
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithBusinessEntityCode(businessEntityCode string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_code"] = businessEntityCode })
}

// WithBusinessEntityName business_entity_name获取 业务实体名称
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithBusinessEntityName(businessEntityName string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_name"] = businessEntityName })
}

// WithBusinessEntityType business_entity_type获取 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithBusinessEntityType(businessEntityType bool) Option {
	return optionFunc(func(o *options) { o.query["business_entity_type"] = businessEntityType })
}

// WithIsAccept is_accept获取 已核付(关联付款单状态)
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithIsAccept(isAccept []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_accept"] = isAccept })
}

// WithAcceptTime accept_time获取 核付时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithAcceptTime(acceptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["accept_time"] = acceptTime })
}

// WithAcceptUser accept_user获取 核付用户
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithAcceptUser(acceptUser string) Option {
	return optionFunc(func(o *options) { o.query["accept_user"] = acceptUser })
}

// WithAcceptUserID accept_user_id获取 核付用户ID
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithAcceptUserID(acceptUserID int) Option {
	return optionFunc(func(o *options) { o.query["accept_user_id"] = acceptUserID })
}

// WithWeight weight获取 计费总重量
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithPayableRmb payable_rmb获取 应付总费用(RMB)
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithPayableRmb(payableRmb float64) Option {
	return optionFunc(func(o *options) { o.query["payable_rmb"] = payableRmb })
}

// WithRemark remark获取 备注
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应收人员
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaBusinessEntityPaymentDetailsMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetByOption(opts ...Option) (result FaBusinessEntityPaymentDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetByOptions(opts ...Option) (results []*FaBusinessEntityPaymentDetails, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaBusinessEntityPaymentDetailsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaBusinessEntityPaymentDetails, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where(options.query)
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
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromID(id int) (result FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromID(ids []int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPaymentNumber 通过payment_number获取内容 付款单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromPaymentNumber(paymentNumber string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`payment_number` = ?", paymentNumber).Find(&results).Error

	return
}

// GetBatchFromPaymentNumber 批量查找 付款单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromPaymentNumber(paymentNumbers []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`payment_number` IN (?)", paymentNumbers).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromOrderNumber(orderNumber string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromReferenceNumber(referenceNumber string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromBusinessEntityCode 通过business_entity_code获取内容 业务主体代码
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromBusinessEntityCode(businessEntityCode string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityCode 批量查找 业务主体代码
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromBusinessEntityCode(businessEntityCodes []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`business_entity_code` IN (?)", businessEntityCodes).Find(&results).Error

	return
}

// GetFromBusinessEntityName 通过business_entity_name获取内容 业务实体名称
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromBusinessEntityName(businessEntityName string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`business_entity_name` = ?", businessEntityName).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityName 批量查找 业务实体名称
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromBusinessEntityName(businessEntityNames []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`business_entity_name` IN (?)", businessEntityNames).Find(&results).Error

	return
}

// GetFromBusinessEntityType 通过business_entity_type获取内容 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromBusinessEntityType(businessEntityType bool) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`business_entity_type` = ?", businessEntityType).Find(&results).Error

	return
}

// GetBatchFromBusinessEntityType 批量查找 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromBusinessEntityType(businessEntityTypes []bool) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`business_entity_type` IN (?)", businessEntityTypes).Find(&results).Error

	return
}

// GetFromIsAccept 通过is_accept获取内容 已核付(关联付款单状态)
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromIsAccept(isAccept []uint8) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`is_accept` = ?", isAccept).Find(&results).Error

	return
}

// GetBatchFromIsAccept 批量查找 已核付(关联付款单状态)
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromIsAccept(isAccepts [][]uint8) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`is_accept` IN (?)", isAccepts).Find(&results).Error

	return
}

// GetFromAcceptTime 通过accept_time获取内容 核付时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromAcceptTime(acceptTime time.Time) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`accept_time` = ?", acceptTime).Find(&results).Error

	return
}

// GetBatchFromAcceptTime 批量查找 核付时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromAcceptTime(acceptTimes []time.Time) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`accept_time` IN (?)", acceptTimes).Find(&results).Error

	return
}

// GetFromAcceptUser 通过accept_user获取内容 核付用户
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromAcceptUser(acceptUser string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`accept_user` = ?", acceptUser).Find(&results).Error

	return
}

// GetBatchFromAcceptUser 批量查找 核付用户
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromAcceptUser(acceptUsers []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`accept_user` IN (?)", acceptUsers).Find(&results).Error

	return
}

// GetFromAcceptUserID 通过accept_user_id获取内容 核付用户ID
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromAcceptUserID(acceptUserID int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`accept_user_id` = ?", acceptUserID).Find(&results).Error

	return
}

// GetBatchFromAcceptUserID 批量查找 核付用户ID
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromAcceptUserID(acceptUserIDs []int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`accept_user_id` IN (?)", acceptUserIDs).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 计费总重量
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromWeight(weight float64) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 计费总重量
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromWeight(weights []float64) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromPayableRmb 通过payable_rmb获取内容 应付总费用(RMB)
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromPayableRmb(payableRmb float64) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`payable_rmb` = ?", payableRmb).Find(&results).Error

	return
}

// GetBatchFromPayableRmb 批量查找 应付总费用(RMB)
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromPayableRmb(payableRmbs []float64) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`payable_rmb` IN (?)", payableRmbs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromRemark(remark string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromRemark(remarks []string) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromCreateTime(createTime time.Time) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应收人员
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromCreateUser(createUser int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应收人员
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromUpdateUser(updateUser int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetFromVersion(version int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaBusinessEntityPaymentDetailsMgr) GetBatchFromVersion(versions []int) (results []*FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaBusinessEntityPaymentDetailsMgr) FetchByPrimaryKey(id int) (result FaBusinessEntityPaymentDetails, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntityPaymentDetails{}).Where("`id` = ?", id).First(&result).Error

	return
}
