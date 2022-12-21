package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderProblemLogMgr struct {
	*_BaseMgr
}

// LgOrderProblemLogMgr open func
func LgOrderProblemLogMgr(db *gorm.DB) *_LgOrderProblemLogMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderProblemLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderProblemLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_problem_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderProblemLogMgr) GetTableName() string {
	return "lg_order_problem_log"
}

// Reset 重置gorm会话
func (obj *_LgOrderProblemLogMgr) Reset() *_LgOrderProblemLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderProblemLogMgr) Get() (result LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderProblemLogMgr) Gets() (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderProblemLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgOrderProblemLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_LgOrderProblemLogMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_LgOrderProblemLogMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_LgOrderProblemLogMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_LgOrderProblemLogMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithCustomerID customer_id获取 归属客户ID
func (obj *_LgOrderProblemLogMgr) WithCustomerID(customerID int64) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_LgOrderProblemLogMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgOrderProblemLogMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 渠道名称
func (obj *_LgOrderProblemLogMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithProblemType problem_type获取 问题件类型
func (obj *_LgOrderProblemLogMgr) WithProblemType(problemType string) Option {
	return optionFunc(func(o *options) { o.query["problem_type"] = problemType })
}

// WithProblemReason problem_reason获取 问题件原因
func (obj *_LgOrderProblemLogMgr) WithProblemReason(problemReason string) Option {
	return optionFunc(func(o *options) { o.query["problem_reason"] = problemReason })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderProblemLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderProblemLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateUserName create_user_name获取 创建用户名称
func (obj *_LgOrderProblemLogMgr) WithCreateUserName(createUserName string) Option {
	return optionFunc(func(o *options) { o.query["create_user_name"] = createUserName })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderProblemLogMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_LgOrderProblemLogMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateUserName update_user_name获取 更新用户名称
func (obj *_LgOrderProblemLogMgr) WithUpdateUserName(updateUserName string) Option {
	return optionFunc(func(o *options) { o.query["update_user_name"] = updateUserName })
}

// WithHandlerID handler_id获取 经办人id
func (obj *_LgOrderProblemLogMgr) WithHandlerID(handlerID int) Option {
	return optionFunc(func(o *options) { o.query["handler_id"] = handlerID })
}

// WithHandlerName handler_name获取 经办人名称
func (obj *_LgOrderProblemLogMgr) WithHandlerName(handlerName string) Option {
	return optionFunc(func(o *options) { o.query["handler_name"] = handlerName })
}

// WithHandleStatus handle_status获取 处理状态 1:待处理 2:已处理
func (obj *_LgOrderProblemLogMgr) WithHandleStatus(handleStatus int) Option {
	return optionFunc(func(o *options) { o.query["handle_status"] = handleStatus })
}

// WithOperationType operation_type获取 操作类型
func (obj *_LgOrderProblemLogMgr) WithOperationType(operationType string) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithRemark remark获取 说明
func (obj *_LgOrderProblemLogMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderProblemLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderProblemLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithHandleTime handle_time获取 处理时间
func (obj *_LgOrderProblemLogMgr) WithHandleTime(handleTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["handle_time"] = handleTime })
}

// WithHandleUser handle_user获取 处理用户
func (obj *_LgOrderProblemLogMgr) WithHandleUser(handleUser int) Option {
	return optionFunc(func(o *options) { o.query["handle_user"] = handleUser })
}

// WithHandleUserName handle_user_name获取 处理人名称
func (obj *_LgOrderProblemLogMgr) WithHandleUserName(handleUserName string) Option {
	return optionFunc(func(o *options) { o.query["handle_user_name"] = handleUserName })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderProblemLogMgr) GetByOption(opts ...Option) (result LgOrderProblemLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderProblemLogMgr) GetByOptions(opts ...Option) (results []*LgOrderProblemLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderProblemLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderProblemLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_LgOrderProblemLogMgr) GetFromID(id int64) (result LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LgOrderProblemLogMgr) GetBatchFromID(ids []int64) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_LgOrderProblemLogMgr) GetFromOrderNumber(orderNumber string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_LgOrderProblemLogMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_LgOrderProblemLogMgr) GetFromReferenceNumber(referenceNumber string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_LgOrderProblemLogMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_LgOrderProblemLogMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_LgOrderProblemLogMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_LgOrderProblemLogMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_LgOrderProblemLogMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 归属客户ID
func (obj *_LgOrderProblemLogMgr) GetFromCustomerID(customerID int64) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 归属客户ID
func (obj *_LgOrderProblemLogMgr) GetBatchFromCustomerID(customerIDs []int64) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_LgOrderProblemLogMgr) GetFromCustomerAlias(customerAlias string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_LgOrderProblemLogMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgOrderProblemLogMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgOrderProblemLogMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 渠道名称
func (obj *_LgOrderProblemLogMgr) GetFromCustomerChannelName(customerChannelName string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 渠道名称
func (obj *_LgOrderProblemLogMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromProblemType 通过problem_type获取内容 问题件类型
func (obj *_LgOrderProblemLogMgr) GetFromProblemType(problemType string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`problem_type` = ?", problemType).Find(&results).Error

	return
}

// GetBatchFromProblemType 批量查找 问题件类型
func (obj *_LgOrderProblemLogMgr) GetBatchFromProblemType(problemTypes []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`problem_type` IN (?)", problemTypes).Find(&results).Error

	return
}

// GetFromProblemReason 通过problem_reason获取内容 问题件原因
func (obj *_LgOrderProblemLogMgr) GetFromProblemReason(problemReason string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`problem_reason` = ?", problemReason).Find(&results).Error

	return
}

// GetBatchFromProblemReason 批量查找 问题件原因
func (obj *_LgOrderProblemLogMgr) GetBatchFromProblemReason(problemReasons []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`problem_reason` IN (?)", problemReasons).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderProblemLogMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderProblemLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderProblemLogMgr) GetFromCreateUser(createUser int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderProblemLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateUserName 通过create_user_name获取内容 创建用户名称
func (obj *_LgOrderProblemLogMgr) GetFromCreateUserName(createUserName string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`create_user_name` = ?", createUserName).Find(&results).Error

	return
}

// GetBatchFromCreateUserName 批量查找 创建用户名称
func (obj *_LgOrderProblemLogMgr) GetBatchFromCreateUserName(createUserNames []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`create_user_name` IN (?)", createUserNames).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderProblemLogMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderProblemLogMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_LgOrderProblemLogMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_LgOrderProblemLogMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateUserName 通过update_user_name获取内容 更新用户名称
func (obj *_LgOrderProblemLogMgr) GetFromUpdateUserName(updateUserName string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`update_user_name` = ?", updateUserName).Find(&results).Error

	return
}

// GetBatchFromUpdateUserName 批量查找 更新用户名称
func (obj *_LgOrderProblemLogMgr) GetBatchFromUpdateUserName(updateUserNames []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`update_user_name` IN (?)", updateUserNames).Find(&results).Error

	return
}

// GetFromHandlerID 通过handler_id获取内容 经办人id
func (obj *_LgOrderProblemLogMgr) GetFromHandlerID(handlerID int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handler_id` = ?", handlerID).Find(&results).Error

	return
}

// GetBatchFromHandlerID 批量查找 经办人id
func (obj *_LgOrderProblemLogMgr) GetBatchFromHandlerID(handlerIDs []int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handler_id` IN (?)", handlerIDs).Find(&results).Error

	return
}

// GetFromHandlerName 通过handler_name获取内容 经办人名称
func (obj *_LgOrderProblemLogMgr) GetFromHandlerName(handlerName string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handler_name` = ?", handlerName).Find(&results).Error

	return
}

// GetBatchFromHandlerName 批量查找 经办人名称
func (obj *_LgOrderProblemLogMgr) GetBatchFromHandlerName(handlerNames []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handler_name` IN (?)", handlerNames).Find(&results).Error

	return
}

// GetFromHandleStatus 通过handle_status获取内容 处理状态 1:待处理 2:已处理
func (obj *_LgOrderProblemLogMgr) GetFromHandleStatus(handleStatus int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handle_status` = ?", handleStatus).Find(&results).Error

	return
}

// GetBatchFromHandleStatus 批量查找 处理状态 1:待处理 2:已处理
func (obj *_LgOrderProblemLogMgr) GetBatchFromHandleStatus(handleStatuss []int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handle_status` IN (?)", handleStatuss).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容 操作类型
func (obj *_LgOrderProblemLogMgr) GetFromOperationType(operationType string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找 操作类型
func (obj *_LgOrderProblemLogMgr) GetBatchFromOperationType(operationTypes []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 说明
func (obj *_LgOrderProblemLogMgr) GetFromRemark(remark string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 说明
func (obj *_LgOrderProblemLogMgr) GetBatchFromRemark(remarks []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderProblemLogMgr) GetFromVersion(version int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderProblemLogMgr) GetBatchFromVersion(versions []int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderProblemLogMgr) GetFromDeleted(deleted int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderProblemLogMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromHandleTime 通过handle_time获取内容 处理时间
func (obj *_LgOrderProblemLogMgr) GetFromHandleTime(handleTime time.Time) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handle_time` = ?", handleTime).Find(&results).Error

	return
}

// GetBatchFromHandleTime 批量查找 处理时间
func (obj *_LgOrderProblemLogMgr) GetBatchFromHandleTime(handleTimes []time.Time) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handle_time` IN (?)", handleTimes).Find(&results).Error

	return
}

// GetFromHandleUser 通过handle_user获取内容 处理用户
func (obj *_LgOrderProblemLogMgr) GetFromHandleUser(handleUser int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handle_user` = ?", handleUser).Find(&results).Error

	return
}

// GetBatchFromHandleUser 批量查找 处理用户
func (obj *_LgOrderProblemLogMgr) GetBatchFromHandleUser(handleUsers []int) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handle_user` IN (?)", handleUsers).Find(&results).Error

	return
}

// GetFromHandleUserName 通过handle_user_name获取内容 处理人名称
func (obj *_LgOrderProblemLogMgr) GetFromHandleUserName(handleUserName string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handle_user_name` = ?", handleUserName).Find(&results).Error

	return
}

// GetBatchFromHandleUserName 批量查找 处理人名称
func (obj *_LgOrderProblemLogMgr) GetBatchFromHandleUserName(handleUserNames []string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`handle_user_name` IN (?)", handleUserNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderProblemLogMgr) FetchByPrimaryKey(id int64) (result LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByLgOrderProblemLogNumberCombinationProblemTypeIndex primary or index 获取唯一内容
func (obj *_LgOrderProblemLogMgr) FetchUniqueIndexByLgOrderProblemLogNumberCombinationProblemTypeIndex(orderNumber string, referenceNumber string, logisticsNumber string, logisticsNumberFinal string, problemType string) (result LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`order_number` = ? AND `reference_number` = ? AND `logistics_number` = ? AND `logistics_number_final` = ? AND `problem_type` = ?", orderNumber, referenceNumber, logisticsNumber, logisticsNumberFinal, problemType).First(&result).Error

	return
}

// FetchIndexByLgOrderProblemLogOrderNumberProblemTypeIndex  获取多个内容
func (obj *_LgOrderProblemLogMgr) FetchIndexByLgOrderProblemLogOrderNumberProblemTypeIndex(orderNumber string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderProblemLogReferenceNumberProblemTypeIndex  获取多个内容
func (obj *_LgOrderProblemLogMgr) FetchIndexByLgOrderProblemLogReferenceNumberProblemTypeIndex(referenceNumber string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderProblemLogLogisticsNumberProblemTypeIndex  获取多个内容
func (obj *_LgOrderProblemLogMgr) FetchIndexByLgOrderProblemLogLogisticsNumberProblemTypeIndex(logisticsNumber string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderProblemLogLogisticsNumberFinalProblemTypeIndex  获取多个内容
func (obj *_LgOrderProblemLogMgr) FetchIndexByLgOrderProblemLogLogisticsNumberFinalProblemTypeIndex(logisticsNumberFinal string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexByLgOrderProblemProblemReasonIndex  获取多个内容
func (obj *_LgOrderProblemLogMgr) FetchIndexByLgOrderProblemProblemReasonIndex(problemReason string) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`problem_reason` = ?", problemReason).Find(&results).Error

	return
}

// FetchIndexByLgOrderProblemCreateTimeIndex  获取多个内容
func (obj *_LgOrderProblemLogMgr) FetchIndexByLgOrderProblemCreateTimeIndex(createTime time.Time) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderProblemUpdateTimeIndex  获取多个内容
func (obj *_LgOrderProblemLogMgr) FetchIndexByLgOrderProblemUpdateTimeIndex(updateTime time.Time) (results []*LgOrderProblemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderProblemLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}
