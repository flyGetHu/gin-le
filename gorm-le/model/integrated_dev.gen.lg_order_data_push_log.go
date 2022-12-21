package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderDataPushLogMgr struct {
	*_BaseMgr
}

// LgOrderDataPushLogMgr open func
func LgOrderDataPushLogMgr(db *gorm.DB) *_LgOrderDataPushLogMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderDataPushLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderDataPushLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_data_push_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderDataPushLogMgr) GetTableName() string {
	return "lg_order_data_push_log"
}

// Reset 重置gorm会话
func (obj *_LgOrderDataPushLogMgr) Reset() *_LgOrderDataPushLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderDataPushLogMgr) Get() (result LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderDataPushLogMgr) Gets() (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderDataPushLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgOrderDataPushLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithClientID client_id获取 客户id
func (obj *_LgOrderDataPushLogMgr) WithClientID(clientID int) Option {
	return optionFunc(func(o *options) { o.query["client_id"] = clientID })
}

// WithClientName client_name获取 客户名称(客户别名)
func (obj *_LgOrderDataPushLogMgr) WithClientName(clientName string) Option {
	return optionFunc(func(o *options) { o.query["client_name"] = clientName })
}

// WithOrderNumber order_number获取 订单号
func (obj *_LgOrderDataPushLogMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_LgOrderDataPushLogMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_LgOrderDataPushLogMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_LgOrderDataPushLogMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithLabelURL label_url获取 面单url
func (obj *_LgOrderDataPushLogMgr) WithLabelURL(labelURL string) Option {
	return optionFunc(func(o *options) { o.query["label_url"] = labelURL })
}

// WithLabelURLFinal label_url_final获取 最终面单地址
func (obj *_LgOrderDataPushLogMgr) WithLabelURLFinal(labelURLFinal string) Option {
	return optionFunc(func(o *options) { o.query["label_url_final"] = labelURLFinal })
}

// WithReceiptTime receipt_time获取 入库时间
func (obj *_LgOrderDataPushLogMgr) WithReceiptTime(receiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_time"] = receiptTime })
}

// WithSendTime send_time获取 出库时间
func (obj *_LgOrderDataPushLogMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithCustomerChannelName customer_channel_name获取 渠道名称
func (obj *_LgOrderDataPushLogMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithWeight weight获取 称重重量
func (obj *_LgOrderDataPushLogMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithNode node获取 推送节点(入库，复重，换单，出库)
func (obj *_LgOrderDataPushLogMgr) WithNode(node string) Option {
	return optionFunc(func(o *options) { o.query["node"] = node })
}

// WithSuccess success获取 是否推送成功：0-失败，1-成功
func (obj *_LgOrderDataPushLogMgr) WithSuccess(success []uint8) Option {
	return optionFunc(func(o *options) { o.query["success"] = success })
}

// WithFrequency frequency获取 推送次数
func (obj *_LgOrderDataPushLogMgr) WithFrequency(frequency int) Option {
	return optionFunc(func(o *options) { o.query["frequency"] = frequency })
}

// WithPushResult push_result获取 推送结果
func (obj *_LgOrderDataPushLogMgr) WithPushResult(pushResult string) Option {
	return optionFunc(func(o *options) { o.query["push_result"] = pushResult })
}

// WithRequestIP request_ip获取 请求地址
func (obj *_LgOrderDataPushLogMgr) WithRequestIP(requestIP string) Option {
	return optionFunc(func(o *options) { o.query["request_ip"] = requestIP })
}

// WithRequestParams request_params获取 请求报文
func (obj *_LgOrderDataPushLogMgr) WithRequestParams(requestParams string) Option {
	return optionFunc(func(o *options) { o.query["request_params"] = requestParams })
}

// WithResponseParams response_params获取 响应报文
func (obj *_LgOrderDataPushLogMgr) WithResponseParams(responseParams string) Option {
	return optionFunc(func(o *options) { o.query["response_params"] = responseParams })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderDataPushLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderDataPushLogMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderDataPushLogMgr) GetByOption(opts ...Option) (result LgOrderDataPushLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderDataPushLogMgr) GetByOptions(opts ...Option) (results []*LgOrderDataPushLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderDataPushLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderDataPushLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where(options.query)
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
func (obj *_LgOrderDataPushLogMgr) GetFromID(id int) (result LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgOrderDataPushLogMgr) GetBatchFromID(ids []int) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromClientID 通过client_id获取内容 客户id
func (obj *_LgOrderDataPushLogMgr) GetFromClientID(clientID int) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`client_id` = ?", clientID).Find(&results).Error

	return
}

// GetBatchFromClientID 批量查找 客户id
func (obj *_LgOrderDataPushLogMgr) GetBatchFromClientID(clientIDs []int) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`client_id` IN (?)", clientIDs).Find(&results).Error

	return
}

// GetFromClientName 通过client_name获取内容 客户名称(客户别名)
func (obj *_LgOrderDataPushLogMgr) GetFromClientName(clientName string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`client_name` = ?", clientName).Find(&results).Error

	return
}

// GetBatchFromClientName 批量查找 客户名称(客户别名)
func (obj *_LgOrderDataPushLogMgr) GetBatchFromClientName(clientNames []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`client_name` IN (?)", clientNames).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_LgOrderDataPushLogMgr) GetFromOrderNumber(orderNumber string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_LgOrderDataPushLogMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_LgOrderDataPushLogMgr) GetFromReferenceNumber(referenceNumber string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_LgOrderDataPushLogMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_LgOrderDataPushLogMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_LgOrderDataPushLogMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_LgOrderDataPushLogMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_LgOrderDataPushLogMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromLabelURL 通过label_url获取内容 面单url
func (obj *_LgOrderDataPushLogMgr) GetFromLabelURL(labelURL string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`label_url` = ?", labelURL).Find(&results).Error

	return
}

// GetBatchFromLabelURL 批量查找 面单url
func (obj *_LgOrderDataPushLogMgr) GetBatchFromLabelURL(labelURLs []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`label_url` IN (?)", labelURLs).Find(&results).Error

	return
}

// GetFromLabelURLFinal 通过label_url_final获取内容 最终面单地址
func (obj *_LgOrderDataPushLogMgr) GetFromLabelURLFinal(labelURLFinal string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`label_url_final` = ?", labelURLFinal).Find(&results).Error

	return
}

// GetBatchFromLabelURLFinal 批量查找 最终面单地址
func (obj *_LgOrderDataPushLogMgr) GetBatchFromLabelURLFinal(labelURLFinals []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`label_url_final` IN (?)", labelURLFinals).Find(&results).Error

	return
}

// GetFromReceiptTime 通过receipt_time获取内容 入库时间
func (obj *_LgOrderDataPushLogMgr) GetFromReceiptTime(receiptTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// GetBatchFromReceiptTime 批量查找 入库时间
func (obj *_LgOrderDataPushLogMgr) GetBatchFromReceiptTime(receiptTimes []time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`receipt_time` IN (?)", receiptTimes).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 出库时间
func (obj *_LgOrderDataPushLogMgr) GetFromSendTime(sendTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 出库时间
func (obj *_LgOrderDataPushLogMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 渠道名称
func (obj *_LgOrderDataPushLogMgr) GetFromCustomerChannelName(customerChannelName string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 渠道名称
func (obj *_LgOrderDataPushLogMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 称重重量
func (obj *_LgOrderDataPushLogMgr) GetFromWeight(weight float64) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 称重重量
func (obj *_LgOrderDataPushLogMgr) GetBatchFromWeight(weights []float64) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromNode 通过node获取内容 推送节点(入库，复重，换单，出库)
func (obj *_LgOrderDataPushLogMgr) GetFromNode(node string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`node` = ?", node).Find(&results).Error

	return
}

// GetBatchFromNode 批量查找 推送节点(入库，复重，换单，出库)
func (obj *_LgOrderDataPushLogMgr) GetBatchFromNode(nodes []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`node` IN (?)", nodes).Find(&results).Error

	return
}

// GetFromSuccess 通过success获取内容 是否推送成功：0-失败，1-成功
func (obj *_LgOrderDataPushLogMgr) GetFromSuccess(success []uint8) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`success` = ?", success).Find(&results).Error

	return
}

// GetBatchFromSuccess 批量查找 是否推送成功：0-失败，1-成功
func (obj *_LgOrderDataPushLogMgr) GetBatchFromSuccess(successs [][]uint8) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`success` IN (?)", successs).Find(&results).Error

	return
}

// GetFromFrequency 通过frequency获取内容 推送次数
func (obj *_LgOrderDataPushLogMgr) GetFromFrequency(frequency int) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`frequency` = ?", frequency).Find(&results).Error

	return
}

// GetBatchFromFrequency 批量查找 推送次数
func (obj *_LgOrderDataPushLogMgr) GetBatchFromFrequency(frequencys []int) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`frequency` IN (?)", frequencys).Find(&results).Error

	return
}

// GetFromPushResult 通过push_result获取内容 推送结果
func (obj *_LgOrderDataPushLogMgr) GetFromPushResult(pushResult string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`push_result` = ?", pushResult).Find(&results).Error

	return
}

// GetBatchFromPushResult 批量查找 推送结果
func (obj *_LgOrderDataPushLogMgr) GetBatchFromPushResult(pushResults []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`push_result` IN (?)", pushResults).Find(&results).Error

	return
}

// GetFromRequestIP 通过request_ip获取内容 请求地址
func (obj *_LgOrderDataPushLogMgr) GetFromRequestIP(requestIP string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`request_ip` = ?", requestIP).Find(&results).Error

	return
}

// GetBatchFromRequestIP 批量查找 请求地址
func (obj *_LgOrderDataPushLogMgr) GetBatchFromRequestIP(requestIPs []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`request_ip` IN (?)", requestIPs).Find(&results).Error

	return
}

// GetFromRequestParams 通过request_params获取内容 请求报文
func (obj *_LgOrderDataPushLogMgr) GetFromRequestParams(requestParams string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`request_params` = ?", requestParams).Find(&results).Error

	return
}

// GetBatchFromRequestParams 批量查找 请求报文
func (obj *_LgOrderDataPushLogMgr) GetBatchFromRequestParams(requestParamss []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`request_params` IN (?)", requestParamss).Find(&results).Error

	return
}

// GetFromResponseParams 通过response_params获取内容 响应报文
func (obj *_LgOrderDataPushLogMgr) GetFromResponseParams(responseParams string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`response_params` = ?", responseParams).Find(&results).Error

	return
}

// GetBatchFromResponseParams 批量查找 响应报文
func (obj *_LgOrderDataPushLogMgr) GetBatchFromResponseParams(responseParamss []string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`response_params` IN (?)", responseParamss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderDataPushLogMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderDataPushLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderDataPushLogMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderDataPushLogMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderDataPushLogMgr) FetchByPrimaryKey(id int) (result LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgOrderDataPushLogIDUIndex primary or index 获取唯一内容
func (obj *_LgOrderDataPushLogMgr) FetchUniqueByLgOrderDataPushLogIDUIndex(id int) (result LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByLgOrderDataPushLogNodeReferenceNumberUIndex primary or index 获取唯一内容
func (obj *_LgOrderDataPushLogMgr) FetchUniqueIndexByLgOrderDataPushLogNodeReferenceNumberUIndex(referenceNumber string, node string) (result LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`reference_number` = ? AND `node` = ?", referenceNumber, node).First(&result).Error

	return
}

// FetchIndexByLgOrderDataPushLogClientIDIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogClientIDIndex(clientID int) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`client_id` = ?", clientID).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogClientNameIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogClientNameIndex(clientName string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`client_name` = ?", clientName).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogOrderNumberIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogOrderNumberIndex(orderNumber string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// FetchIndexByLgDataPushLogReferenceNumberIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgDataPushLogReferenceNumberIndex(referenceNumber string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogLogisticsNumberIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogLogisticsNumberIndex(logisticsNumber string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogReceiptTimeIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogReceiptTimeIndex(receiptTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogSendTimeIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogSendTimeIndex(sendTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// FetchIndexByLgDataPushLogCustomerChannelNameIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgDataPushLogCustomerChannelNameIndex(customerChannelName string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogNodeIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogNodeIndex(node string) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`node` = ?", node).Find(&results).Error

	return
}

// FetchIndexByIDxNodeSuccessFrequencyCreatetime  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByIDxNodeSuccessFrequencyCreatetime(node string, success []uint8, frequency int, createTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`node` = ? AND `success` = ? AND `frequency` = ? AND `create_time` = ?", node, success, frequency, createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogSuccessIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogSuccessIndex(success []uint8) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`success` = ?", success).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogFrequencyIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogFrequencyIndex(frequency int) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`frequency` = ?", frequency).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogCreateTimeIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogCreateTimeIndex(createTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderDataPushLogUpdateTimeIndex  获取多个内容
func (obj *_LgOrderDataPushLogMgr) FetchIndexByLgOrderDataPushLogUpdateTimeIndex(updateTime time.Time) (results []*LgOrderDataPushLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderDataPushLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}
