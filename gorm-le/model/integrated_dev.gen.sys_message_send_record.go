package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysMessageSendRecordMgr struct {
	*_BaseMgr
}

// SysMessageSendRecordMgr open func
func SysMessageSendRecordMgr(db *gorm.DB) *_SysMessageSendRecordMgr {
	if db == nil {
		panic(fmt.Errorf("SysMessageSendRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysMessageSendRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_message_send_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysMessageSendRecordMgr) GetTableName() string {
	return "sys_message_send_record"
}

// Reset 重置gorm会话
func (obj *_SysMessageSendRecordMgr) Reset() *_SysMessageSendRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SysMessageSendRecordMgr) Get() (result SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysMessageSendRecordMgr) Gets() (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SysMessageSendRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SysMessageSendRecordMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_SysMessageSendRecordMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_SysMessageSendRecordMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_SysMessageSendRecordMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_SysMessageSendRecordMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithSendContent send_content获取 发送内容
func (obj *_SysMessageSendRecordMgr) WithSendContent(sendContent string) Option {
	return optionFunc(func(o *options) { o.query["send_content"] = sendContent })
}

// WithMessageType message_type获取 消息类型，sms:短信，email:邮件
func (obj *_SysMessageSendRecordMgr) WithMessageType(messageType string) Option {
	return optionFunc(func(o *options) { o.query["message_type"] = messageType })
}

// WithSendNode send_node获取 发送节点
func (obj *_SysMessageSendRecordMgr) WithSendNode(sendNode string) Option {
	return optionFunc(func(o *options) { o.query["send_node"] = sendNode })
}

// WithSendStatus send_status获取 发送状态，0:发送中，1:发送成功，2:发送失败
func (obj *_SysMessageSendRecordMgr) WithSendStatus(sendStatus int8) Option {
	return optionFunc(func(o *options) { o.query["send_status"] = sendStatus })
}

// WithSendCount send_count获取 发送次数
func (obj *_SysMessageSendRecordMgr) WithSendCount(sendCount int) Option {
	return optionFunc(func(o *options) { o.query["send_count"] = sendCount })
}

// WithNumber number获取 发送号码
func (obj *_SysMessageSendRecordMgr) WithNumber(number string) Option {
	return optionFunc(func(o *options) { o.query["number"] = number })
}

// WithSendTime send_time获取 发送时间
func (obj *_SysMessageSendRecordMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithServiceProviderCode service_provider_code获取 发送消息服务商code
func (obj *_SysMessageSendRecordMgr) WithServiceProviderCode(serviceProviderCode string) Option {
	return optionFunc(func(o *options) { o.query["service_provider_code"] = serviceProviderCode })
}

// WithOperatorsID operators_id获取 运营商ID
func (obj *_SysMessageSendRecordMgr) WithOperatorsID(operatorsID string) Option {
	return optionFunc(func(o *options) { o.query["operators_id"] = operatorsID })
}

// WithRemark remark获取 备注
func (obj *_SysMessageSendRecordMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysMessageSendRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_SysMessageSendRecordMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysMessageSendRecordMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_SysMessageSendRecordMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_SysMessageSendRecordMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_SysMessageSendRecordMgr) GetByOption(opts ...Option) (result SysMessageSendRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysMessageSendRecordMgr) GetByOptions(opts ...Option) (results []*SysMessageSendRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SysMessageSendRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SysMessageSendRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where(options.query)
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
func (obj *_SysMessageSendRecordMgr) GetFromID(id int) (result SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SysMessageSendRecordMgr) GetBatchFromID(ids []int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_SysMessageSendRecordMgr) GetFromOrderNumber(orderNumber string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_SysMessageSendRecordMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_SysMessageSendRecordMgr) GetFromReferenceNumber(referenceNumber string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_SysMessageSendRecordMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_SysMessageSendRecordMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_SysMessageSendRecordMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_SysMessageSendRecordMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_SysMessageSendRecordMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromSendContent 通过send_content获取内容 发送内容
func (obj *_SysMessageSendRecordMgr) GetFromSendContent(sendContent string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_content` = ?", sendContent).Find(&results).Error

	return
}

// GetBatchFromSendContent 批量查找 发送内容
func (obj *_SysMessageSendRecordMgr) GetBatchFromSendContent(sendContents []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_content` IN (?)", sendContents).Find(&results).Error

	return
}

// GetFromMessageType 通过message_type获取内容 消息类型，sms:短信，email:邮件
func (obj *_SysMessageSendRecordMgr) GetFromMessageType(messageType string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`message_type` = ?", messageType).Find(&results).Error

	return
}

// GetBatchFromMessageType 批量查找 消息类型，sms:短信，email:邮件
func (obj *_SysMessageSendRecordMgr) GetBatchFromMessageType(messageTypes []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`message_type` IN (?)", messageTypes).Find(&results).Error

	return
}

// GetFromSendNode 通过send_node获取内容 发送节点
func (obj *_SysMessageSendRecordMgr) GetFromSendNode(sendNode string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_node` = ?", sendNode).Find(&results).Error

	return
}

// GetBatchFromSendNode 批量查找 发送节点
func (obj *_SysMessageSendRecordMgr) GetBatchFromSendNode(sendNodes []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_node` IN (?)", sendNodes).Find(&results).Error

	return
}

// GetFromSendStatus 通过send_status获取内容 发送状态，0:发送中，1:发送成功，2:发送失败
func (obj *_SysMessageSendRecordMgr) GetFromSendStatus(sendStatus int8) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_status` = ?", sendStatus).Find(&results).Error

	return
}

// GetBatchFromSendStatus 批量查找 发送状态，0:发送中，1:发送成功，2:发送失败
func (obj *_SysMessageSendRecordMgr) GetBatchFromSendStatus(sendStatuss []int8) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_status` IN (?)", sendStatuss).Find(&results).Error

	return
}

// GetFromSendCount 通过send_count获取内容 发送次数
func (obj *_SysMessageSendRecordMgr) GetFromSendCount(sendCount int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_count` = ?", sendCount).Find(&results).Error

	return
}

// GetBatchFromSendCount 批量查找 发送次数
func (obj *_SysMessageSendRecordMgr) GetBatchFromSendCount(sendCounts []int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_count` IN (?)", sendCounts).Find(&results).Error

	return
}

// GetFromNumber 通过number获取内容 发送号码
func (obj *_SysMessageSendRecordMgr) GetFromNumber(number string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`number` = ?", number).Find(&results).Error

	return
}

// GetBatchFromNumber 批量查找 发送号码
func (obj *_SysMessageSendRecordMgr) GetBatchFromNumber(numbers []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`number` IN (?)", numbers).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 发送时间
func (obj *_SysMessageSendRecordMgr) GetFromSendTime(sendTime time.Time) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 发送时间
func (obj *_SysMessageSendRecordMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromServiceProviderCode 通过service_provider_code获取内容 发送消息服务商code
func (obj *_SysMessageSendRecordMgr) GetFromServiceProviderCode(serviceProviderCode string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`service_provider_code` = ?", serviceProviderCode).Find(&results).Error

	return
}

// GetBatchFromServiceProviderCode 批量查找 发送消息服务商code
func (obj *_SysMessageSendRecordMgr) GetBatchFromServiceProviderCode(serviceProviderCodes []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`service_provider_code` IN (?)", serviceProviderCodes).Find(&results).Error

	return
}

// GetFromOperatorsID 通过operators_id获取内容 运营商ID
func (obj *_SysMessageSendRecordMgr) GetFromOperatorsID(operatorsID string) (result SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`operators_id` = ?", operatorsID).First(&result).Error

	return
}

// GetBatchFromOperatorsID 批量查找 运营商ID
func (obj *_SysMessageSendRecordMgr) GetBatchFromOperatorsID(operatorsIDs []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`operators_id` IN (?)", operatorsIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysMessageSendRecordMgr) GetFromRemark(remark string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SysMessageSendRecordMgr) GetBatchFromRemark(remarks []string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysMessageSendRecordMgr) GetFromCreateTime(createTime time.Time) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SysMessageSendRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_SysMessageSendRecordMgr) GetFromCreateUser(createUser int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_SysMessageSendRecordMgr) GetBatchFromCreateUser(createUsers []int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysMessageSendRecordMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_SysMessageSendRecordMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_SysMessageSendRecordMgr) GetFromUpdateUser(updateUser int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_SysMessageSendRecordMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_SysMessageSendRecordMgr) GetFromVersion(version int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_SysMessageSendRecordMgr) GetBatchFromVersion(versions []int) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SysMessageSendRecordMgr) FetchByPrimaryKey(id int) (result SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueBySysMessageSendRecordOperatorsIDUIndex primary or index 获取唯一内容
func (obj *_SysMessageSendRecordMgr) FetchUniqueBySysMessageSendRecordOperatorsIDUIndex(operatorsID string) (result SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`operators_id` = ?", operatorsID).First(&result).Error

	return
}

// FetchIndexByOrderNumberIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexByOrderNumberIndex(orderNumber string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByReferenceNumberIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexByReferenceNumberIndex(referenceNumber string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByLogisticsNumberIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexByLogisticsNumberIndex(logisticsNumber string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByLogisticsNumberFinalIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexByLogisticsNumberFinalIndex(logisticsNumberFinal string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexBySendNodeIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexBySendNodeIndex(sendNode string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_node` = ?", sendNode).Find(&results).Error

	return
}

// FetchIndexBySendStatusIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexBySendStatusIndex(sendStatus int8) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`send_status` = ?", sendStatus).Find(&results).Error

	return
}

// FetchIndexByNumberIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexByNumberIndex(number string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`number` = ?", number).Find(&results).Error

	return
}

// FetchIndexBySysMessageSendRecordServiceProviderCodeIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexBySysMessageSendRecordServiceProviderCodeIndex(serviceProviderCode string) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`service_provider_code` = ?", serviceProviderCode).Find(&results).Error

	return
}

// FetchIndexBySysMessageSendRecordCreateTimeIndex  获取多个内容
func (obj *_SysMessageSendRecordMgr) FetchIndexBySysMessageSendRecordCreateTimeIndex(createTime time.Time) (results []*SysMessageSendRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysMessageSendRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
