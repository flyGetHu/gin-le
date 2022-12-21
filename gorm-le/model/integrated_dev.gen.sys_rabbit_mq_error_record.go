package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysRabbitMqErrorRecordMgr struct {
	*_BaseMgr
}

// SysRabbitMqErrorRecordMgr open func
func SysRabbitMqErrorRecordMgr(db *gorm.DB) *_SysRabbitMqErrorRecordMgr {
	if db == nil {
		panic(fmt.Errorf("SysRabbitMqErrorRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysRabbitMqErrorRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_rabbit_mq_error_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysRabbitMqErrorRecordMgr) GetTableName() string {
	return "sys_rabbit_mq_error_record"
}

// Reset 重置gorm会话
func (obj *_SysRabbitMqErrorRecordMgr) Reset() *_SysRabbitMqErrorRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SysRabbitMqErrorRecordMgr) Get() (result SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysRabbitMqErrorRecordMgr) Gets() (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SysRabbitMqErrorRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_SysRabbitMqErrorRecordMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithExchangeName exchange_name获取 交换机名称
func (obj *_SysRabbitMqErrorRecordMgr) WithExchangeName(exchangeName string) Option {
	return optionFunc(func(o *options) { o.query["exchange_name"] = exchangeName })
}

// WithExchangeType exchange_type获取 交换机类型，direct:精确路由、fanout:扇出,广播、topic:模糊路由
func (obj *_SysRabbitMqErrorRecordMgr) WithExchangeType(exchangeType string) Option {
	return optionFunc(func(o *options) { o.query["exchange_type"] = exchangeType })
}

// WithQueueName queue_name获取 队列名称
func (obj *_SysRabbitMqErrorRecordMgr) WithQueueName(queueName string) Option {
	return optionFunc(func(o *options) { o.query["queue_name"] = queueName })
}

// WithModelName model_name获取 模块名称
func (obj *_SysRabbitMqErrorRecordMgr) WithModelName(modelName string) Option {
	return optionFunc(func(o *options) { o.query["model_name"] = modelName })
}

// WithMsgContent msg_content获取 消息内容
func (obj *_SysRabbitMqErrorRecordMgr) WithMsgContent(msgContent string) Option {
	return optionFunc(func(o *options) { o.query["msg_content"] = msgContent })
}

// WithErrorReason error_reason获取 错误原因
func (obj *_SysRabbitMqErrorRecordMgr) WithErrorReason(errorReason string) Option {
	return optionFunc(func(o *options) { o.query["error_reason"] = errorReason })
}

// WithErrorType error_type获取 错误类型，1:生产失败、2:消费失败
func (obj *_SysRabbitMqErrorRecordMgr) WithErrorType(errorType bool) Option {
	return optionFunc(func(o *options) { o.query["error_type"] = errorType })
}

// WithIsResolve is_resolve获取 已处理，0:否，1:是
func (obj *_SysRabbitMqErrorRecordMgr) WithIsResolve(isResolve []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_resolve"] = isResolve })
}

// WithRemark remark获取 备注
func (obj *_SysRabbitMqErrorRecordMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysRabbitMqErrorRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 生成用户,生成应收人员
func (obj *_SysRabbitMqErrorRecordMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_SysRabbitMqErrorRecordMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysRabbitMqErrorRecordMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_SysRabbitMqErrorRecordMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_SysRabbitMqErrorRecordMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_SysRabbitMqErrorRecordMgr) GetByOption(opts ...Option) (result SysRabbitMqErrorRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysRabbitMqErrorRecordMgr) GetByOptions(opts ...Option) (results []*SysRabbitMqErrorRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SysRabbitMqErrorRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SysRabbitMqErrorRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where(options.query)
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
func (obj *_SysRabbitMqErrorRecordMgr) GetFromID(id int64) (result SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromID(ids []int64) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromExchangeName 通过exchange_name获取内容 交换机名称
func (obj *_SysRabbitMqErrorRecordMgr) GetFromExchangeName(exchangeName string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`exchange_name` = ?", exchangeName).Find(&results).Error

	return
}

// GetBatchFromExchangeName 批量查找 交换机名称
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromExchangeName(exchangeNames []string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`exchange_name` IN (?)", exchangeNames).Find(&results).Error

	return
}

// GetFromExchangeType 通过exchange_type获取内容 交换机类型，direct:精确路由、fanout:扇出,广播、topic:模糊路由
func (obj *_SysRabbitMqErrorRecordMgr) GetFromExchangeType(exchangeType string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`exchange_type` = ?", exchangeType).Find(&results).Error

	return
}

// GetBatchFromExchangeType 批量查找 交换机类型，direct:精确路由、fanout:扇出,广播、topic:模糊路由
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromExchangeType(exchangeTypes []string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`exchange_type` IN (?)", exchangeTypes).Find(&results).Error

	return
}

// GetFromQueueName 通过queue_name获取内容 队列名称
func (obj *_SysRabbitMqErrorRecordMgr) GetFromQueueName(queueName string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`queue_name` = ?", queueName).Find(&results).Error

	return
}

// GetBatchFromQueueName 批量查找 队列名称
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromQueueName(queueNames []string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`queue_name` IN (?)", queueNames).Find(&results).Error

	return
}

// GetFromModelName 通过model_name获取内容 模块名称
func (obj *_SysRabbitMqErrorRecordMgr) GetFromModelName(modelName string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`model_name` = ?", modelName).Find(&results).Error

	return
}

// GetBatchFromModelName 批量查找 模块名称
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromModelName(modelNames []string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`model_name` IN (?)", modelNames).Find(&results).Error

	return
}

// GetFromMsgContent 通过msg_content获取内容 消息内容
func (obj *_SysRabbitMqErrorRecordMgr) GetFromMsgContent(msgContent string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`msg_content` = ?", msgContent).Find(&results).Error

	return
}

// GetBatchFromMsgContent 批量查找 消息内容
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromMsgContent(msgContents []string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`msg_content` IN (?)", msgContents).Find(&results).Error

	return
}

// GetFromErrorReason 通过error_reason获取内容 错误原因
func (obj *_SysRabbitMqErrorRecordMgr) GetFromErrorReason(errorReason string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`error_reason` = ?", errorReason).Find(&results).Error

	return
}

// GetBatchFromErrorReason 批量查找 错误原因
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromErrorReason(errorReasons []string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`error_reason` IN (?)", errorReasons).Find(&results).Error

	return
}

// GetFromErrorType 通过error_type获取内容 错误类型，1:生产失败、2:消费失败
func (obj *_SysRabbitMqErrorRecordMgr) GetFromErrorType(errorType bool) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`error_type` = ?", errorType).Find(&results).Error

	return
}

// GetBatchFromErrorType 批量查找 错误类型，1:生产失败、2:消费失败
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromErrorType(errorTypes []bool) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`error_type` IN (?)", errorTypes).Find(&results).Error

	return
}

// GetFromIsResolve 通过is_resolve获取内容 已处理，0:否，1:是
func (obj *_SysRabbitMqErrorRecordMgr) GetFromIsResolve(isResolve []uint8) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`is_resolve` = ?", isResolve).Find(&results).Error

	return
}

// GetBatchFromIsResolve 批量查找 已处理，0:否，1:是
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromIsResolve(isResolves [][]uint8) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`is_resolve` IN (?)", isResolves).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysRabbitMqErrorRecordMgr) GetFromRemark(remark string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromRemark(remarks []string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysRabbitMqErrorRecordMgr) GetFromCreateTime(createTime time.Time) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 生成用户,生成应收人员
func (obj *_SysRabbitMqErrorRecordMgr) GetFromCreateUser(createUser int) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 生成用户,生成应收人员
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromCreateUser(createUsers []int) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_SysRabbitMqErrorRecordMgr) GetFromUpdateUser(updateUser int) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysRabbitMqErrorRecordMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_SysRabbitMqErrorRecordMgr) GetFromVersion(version int) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromVersion(versions []int) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_SysRabbitMqErrorRecordMgr) GetFromDeleted(deleted int) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_SysRabbitMqErrorRecordMgr) GetBatchFromDeleted(deleteds []int) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchByPrimaryKey(id int64) (result SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIndexExchangeName  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexExchangeName(exchangeName string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`exchange_name` = ?", exchangeName).Find(&results).Error

	return
}

// FetchIndexByIndexExchangeNameExchangeType  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexExchangeNameExchangeType(exchangeName string, exchangeType string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`exchange_name` = ? AND `exchange_type` = ?", exchangeName, exchangeType).Find(&results).Error

	return
}

// FetchIndexByIndexExchangeNameIsResolve  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexExchangeNameIsResolve(exchangeName string, isResolve []uint8) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`exchange_name` = ? AND `is_resolve` = ?", exchangeName, isResolve).Find(&results).Error

	return
}

// FetchIndexByIndexExchangeType  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexExchangeType(exchangeType string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`exchange_type` = ?", exchangeType).Find(&results).Error

	return
}

// FetchIndexByIndexQueueName  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexQueueName(queueName string) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`queue_name` = ?", queueName).Find(&results).Error

	return
}

// FetchIndexByIndexQueueNameIsResolve  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexQueueNameIsResolve(queueName string, isResolve []uint8) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`queue_name` = ? AND `is_resolve` = ?", queueName, isResolve).Find(&results).Error

	return
}

// FetchIndexByIndexModelNameIsResolve  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexModelNameIsResolve(modelName string, isResolve []uint8) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`model_name` = ? AND `is_resolve` = ?", modelName, isResolve).Find(&results).Error

	return
}

// FetchIndexByIndexErrorType  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexErrorType(errorType bool) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`error_type` = ?", errorType).Find(&results).Error

	return
}

// FetchIndexByIndexIsResolve  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexIsResolve(isResolve []uint8) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`is_resolve` = ?", isResolve).Find(&results).Error

	return
}

// FetchIndexByIndexCreateTime  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexCreateTime(createTime time.Time) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByIndexUpdateTime  获取多个内容
func (obj *_SysRabbitMqErrorRecordMgr) FetchIndexByIndexUpdateTime(updateTime time.Time) (results []*SysRabbitMqErrorRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysRabbitMqErrorRecord{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}
