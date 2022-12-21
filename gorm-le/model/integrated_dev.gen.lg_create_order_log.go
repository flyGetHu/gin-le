package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgCreateOrderLogMgr struct {
	*_BaseMgr
}

// LgCreateOrderLogMgr open func
func LgCreateOrderLogMgr(db *gorm.DB) *_LgCreateOrderLogMgr {
	if db == nil {
		panic(fmt.Errorf("LgCreateOrderLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgCreateOrderLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_create_order_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgCreateOrderLogMgr) GetTableName() string {
	return "lg_create_order_log"
}

// Reset 重置gorm会话
func (obj *_LgCreateOrderLogMgr) Reset() *_LgCreateOrderLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgCreateOrderLogMgr) Get() (result LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgCreateOrderLogMgr) Gets() (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgCreateOrderLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgCreateOrderLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithClientID client_id获取 客户id
func (obj *_LgCreateOrderLogMgr) WithClientID(clientID int) Option {
	return optionFunc(func(o *options) { o.query["client_id"] = clientID })
}

// WithClientName client_name获取 客户名
func (obj *_LgCreateOrderLogMgr) WithClientName(clientName string) Option {
	return optionFunc(func(o *options) { o.query["client_name"] = clientName })
}

// WithOrderNum order_num获取 订单号
func (obj *_LgCreateOrderLogMgr) WithOrderNum(orderNum string) Option {
	return optionFunc(func(o *options) { o.query["order_num"] = orderNum })
}

// WithSuccess success获取 是否下单成功 0 失败，1 成功
func (obj *_LgCreateOrderLogMgr) WithSuccess(success []uint8) Option {
	return optionFunc(func(o *options) { o.query["success"] = success })
}

// WithRequestParams request_params获取 请求报文
func (obj *_LgCreateOrderLogMgr) WithRequestParams(requestParams string) Option {
	return optionFunc(func(o *options) { o.query["request_params"] = requestParams })
}

// WithOriginalRequestParams original_request_params获取 原始请求报文
func (obj *_LgCreateOrderLogMgr) WithOriginalRequestParams(originalRequestParams string) Option {
	return optionFunc(func(o *options) { o.query["original_request_params"] = originalRequestParams })
}

// WithResponseParams response_params获取 响应报文
func (obj *_LgCreateOrderLogMgr) WithResponseParams(responseParams string) Option {
	return optionFunc(func(o *options) { o.query["response_params"] = responseParams })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgCreateOrderLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 操作人
func (obj *_LgCreateOrderLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgCreateOrderLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgCreateOrderLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_LgCreateOrderLogMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgCreateOrderLogMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithFailureMessage failure_message获取 失败原因
func (obj *_LgCreateOrderLogMgr) WithFailureMessage(failureMessage string) Option {
	return optionFunc(func(o *options) { o.query["failure_message"] = failureMessage })
}

// GetByOption 功能选项模式获取
func (obj *_LgCreateOrderLogMgr) GetByOption(opts ...Option) (result LgCreateOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgCreateOrderLogMgr) GetByOptions(opts ...Option) (results []*LgCreateOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgCreateOrderLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgCreateOrderLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *_LgCreateOrderLogMgr) GetFromID(id int) (result LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgCreateOrderLogMgr) GetBatchFromID(ids []int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromClientID 通过client_id获取内容 客户id
func (obj *_LgCreateOrderLogMgr) GetFromClientID(clientID int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`client_id` = ?", clientID).Find(&results).Error

	return
}

// GetBatchFromClientID 批量查找 客户id
func (obj *_LgCreateOrderLogMgr) GetBatchFromClientID(clientIDs []int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`client_id` IN (?)", clientIDs).Find(&results).Error

	return
}

// GetFromClientName 通过client_name获取内容 客户名
func (obj *_LgCreateOrderLogMgr) GetFromClientName(clientName string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`client_name` = ?", clientName).Find(&results).Error

	return
}

// GetBatchFromClientName 批量查找 客户名
func (obj *_LgCreateOrderLogMgr) GetBatchFromClientName(clientNames []string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`client_name` IN (?)", clientNames).Find(&results).Error

	return
}

// GetFromOrderNum 通过order_num获取内容 订单号
func (obj *_LgCreateOrderLogMgr) GetFromOrderNum(orderNum string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`order_num` = ?", orderNum).Find(&results).Error

	return
}

// GetBatchFromOrderNum 批量查找 订单号
func (obj *_LgCreateOrderLogMgr) GetBatchFromOrderNum(orderNums []string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`order_num` IN (?)", orderNums).Find(&results).Error

	return
}

// GetFromSuccess 通过success获取内容 是否下单成功 0 失败，1 成功
func (obj *_LgCreateOrderLogMgr) GetFromSuccess(success []uint8) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`success` = ?", success).Find(&results).Error

	return
}

// GetBatchFromSuccess 批量查找 是否下单成功 0 失败，1 成功
func (obj *_LgCreateOrderLogMgr) GetBatchFromSuccess(successs [][]uint8) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`success` IN (?)", successs).Find(&results).Error

	return
}

// GetFromRequestParams 通过request_params获取内容 请求报文
func (obj *_LgCreateOrderLogMgr) GetFromRequestParams(requestParams string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`request_params` = ?", requestParams).Find(&results).Error

	return
}

// GetBatchFromRequestParams 批量查找 请求报文
func (obj *_LgCreateOrderLogMgr) GetBatchFromRequestParams(requestParamss []string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`request_params` IN (?)", requestParamss).Find(&results).Error

	return
}

// GetFromOriginalRequestParams 通过original_request_params获取内容 原始请求报文
func (obj *_LgCreateOrderLogMgr) GetFromOriginalRequestParams(originalRequestParams string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`original_request_params` = ?", originalRequestParams).Find(&results).Error

	return
}

// GetBatchFromOriginalRequestParams 批量查找 原始请求报文
func (obj *_LgCreateOrderLogMgr) GetBatchFromOriginalRequestParams(originalRequestParamss []string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`original_request_params` IN (?)", originalRequestParamss).Find(&results).Error

	return
}

// GetFromResponseParams 通过response_params获取内容 响应报文
func (obj *_LgCreateOrderLogMgr) GetFromResponseParams(responseParams string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`response_params` = ?", responseParams).Find(&results).Error

	return
}

// GetBatchFromResponseParams 批量查找 响应报文
func (obj *_LgCreateOrderLogMgr) GetBatchFromResponseParams(responseParamss []string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`response_params` IN (?)", responseParamss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgCreateOrderLogMgr) GetFromCreateTime(createTime time.Time) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgCreateOrderLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 操作人
func (obj *_LgCreateOrderLogMgr) GetFromCreateUser(createUser int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 操作人
func (obj *_LgCreateOrderLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgCreateOrderLogMgr) GetFromVersion(version int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgCreateOrderLogMgr) GetBatchFromVersion(versions []int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgCreateOrderLogMgr) GetFromDeleted(deleted int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgCreateOrderLogMgr) GetBatchFromDeleted(deleteds []int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_LgCreateOrderLogMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_LgCreateOrderLogMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgCreateOrderLogMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgCreateOrderLogMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromFailureMessage 通过failure_message获取内容 失败原因
func (obj *_LgCreateOrderLogMgr) GetFromFailureMessage(failureMessage string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`failure_message` = ?", failureMessage).Find(&results).Error

	return
}

// GetBatchFromFailureMessage 批量查找 失败原因
func (obj *_LgCreateOrderLogMgr) GetBatchFromFailureMessage(failureMessages []string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`failure_message` IN (?)", failureMessages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgCreateOrderLogMgr) FetchByPrimaryKey(id int) (result LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByOrderOperateLogIDUIndex primary or index 获取唯一内容
func (obj *_LgCreateOrderLogMgr) FetchUniqueByOrderOperateLogIDUIndex(id int) (result LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByLgCreateOrderLogClientIDIndex  获取多个内容
func (obj *_LgCreateOrderLogMgr) FetchIndexByLgCreateOrderLogClientIDIndex(clientID int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`client_id` = ?", clientID).Find(&results).Error

	return
}

// FetchIndexByLgCreateOrderLogOrderNumIndex  获取多个内容
func (obj *_LgCreateOrderLogMgr) FetchIndexByLgCreateOrderLogOrderNumIndex(orderNum string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`order_num` = ?", orderNum).Find(&results).Error

	return
}

// FetchIndexByLgCreateOrderLogSuccessIndex  获取多个内容
func (obj *_LgCreateOrderLogMgr) FetchIndexByLgCreateOrderLogSuccessIndex(success []uint8) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`success` = ?", success).Find(&results).Error

	return
}

// FetchIndexByLgCreateOrderLogCreateTimeIndex  获取多个内容
func (obj *_LgCreateOrderLogMgr) FetchIndexByLgCreateOrderLogCreateTimeIndex(createTime time.Time) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByLgCreateOrderLogDeletedIndex  获取多个内容
func (obj *_LgCreateOrderLogMgr) FetchIndexByLgCreateOrderLogDeletedIndex(deleted int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// FetchIndexByLgCreateOrderLogLogisticsNumberIndex  获取多个内容
func (obj *_LgCreateOrderLogMgr) FetchIndexByLgCreateOrderLogLogisticsNumberIndex(logisticsNumber string) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByLgCreateOrderLogCustomerChannelIDIndex  获取多个内容
func (obj *_LgCreateOrderLogMgr) FetchIndexByLgCreateOrderLogCustomerChannelIDIndex(customerChannelID int) (results []*LgCreateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCreateOrderLog{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}
