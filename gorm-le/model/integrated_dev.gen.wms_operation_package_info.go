package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsOperationPackageInfoMgr struct {
	*_BaseMgr
}

// WmsOperationPackageInfoMgr open func
func WmsOperationPackageInfoMgr(db *gorm.DB) *_WmsOperationPackageInfoMgr {
	if db == nil {
		panic(fmt.Errorf("WmsOperationPackageInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsOperationPackageInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_operation_package_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsOperationPackageInfoMgr) GetTableName() string {
	return "wms_operation_package_info"
}

// Reset 重置gorm会话
func (obj *_WmsOperationPackageInfoMgr) Reset() *_WmsOperationPackageInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsOperationPackageInfoMgr) Get() (result WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsOperationPackageInfoMgr) Gets() (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsOperationPackageInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsOperationPackageInfoMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTrackingNo tracking_no获取 快递单号
func (obj *_WmsOperationPackageInfoMgr) WithTrackingNo(trackingNo string) Option {
	return optionFunc(func(o *options) { o.query["tracking_no"] = trackingNo })
}

// WithTrackingNoType tracking_no_type获取 快递单号类型，reference:参考号，transfer:转单号
func (obj *_WmsOperationPackageInfoMgr) WithTrackingNoType(trackingNoType string) Option {
	return optionFunc(func(o *options) { o.query["tracking_no_type"] = trackingNoType })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_WmsOperationPackageInfoMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_WmsOperationPackageInfoMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_WmsOperationPackageInfoMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithReceiveCountry receive_country获取 收件人国家
func (obj *_WmsOperationPackageInfoMgr) WithReceiveCountry(receiveCountry string) Option {
	return optionFunc(func(o *options) { o.query["receive_country"] = receiveCountry })
}

// WithOrderNo order_no获取 订单号
func (obj *_WmsOperationPackageInfoMgr) WithOrderNo(orderNo string) Option {
	return optionFunc(func(o *options) { o.query["order_no"] = orderNo })
}

// WithCustomerID customer_id获取 客户id
func (obj *_WmsOperationPackageInfoMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithSource source获取 来源，0:内部，1:外部
func (obj *_WmsOperationPackageInfoMgr) WithSource(source int8) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道id
func (obj *_WmsOperationPackageInfoMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithDownstreamCode downstream_code获取 下家code
func (obj *_WmsOperationPackageInfoMgr) WithDownstreamCode(downstreamCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_code"] = downstreamCode })
}

// WithAPIChannelCode api_channel_code获取 api渠道代码
func (obj *_WmsOperationPackageInfoMgr) WithAPIChannelCode(apiChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["api_channel_code"] = apiChannelCode })
}

// WithBigBagNo big_bag_no获取 大包号
func (obj *_WmsOperationPackageInfoMgr) WithBigBagNo(bigBagNo string) Option {
	return optionFunc(func(o *options) { o.query["big_bag_no"] = bigBagNo })
}

// WithTotalListNo total_list_no获取 总单号
func (obj *_WmsOperationPackageInfoMgr) WithTotalListNo(totalListNo string) Option {
	return optionFunc(func(o *options) { o.query["total_list_no"] = totalListNo })
}

// WithAirBillNumber air_bill_number获取 航空提单号
func (obj *_WmsOperationPackageInfoMgr) WithAirBillNumber(airBillNumber string) Option {
	return optionFunc(func(o *options) { o.query["air_bill_number"] = airBillNumber })
}

// WithWeight weight获取 包裹重量(KG),保留4位小数
func (obj *_WmsOperationPackageInfoMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithOrderWeight order_weight获取 订单预报重量
func (obj *_WmsOperationPackageInfoMgr) WithOrderWeight(orderWeight float64) Option {
	return optionFunc(func(o *options) { o.query["order_weight"] = orderWeight })
}

// WithWorkConsole work_console获取 操作台号
func (obj *_WmsOperationPackageInfoMgr) WithWorkConsole(workConsole string) Option {
	return optionFunc(func(o *options) { o.query["work_console"] = workConsole })
}

// WithPicOssLink pic_oss_link获取 包裹图片oss地址
func (obj *_WmsOperationPackageInfoMgr) WithPicOssLink(picOssLink string) Option {
	return optionFunc(func(o *options) { o.query["pic_oss_link"] = picOssLink })
}

// WithScanTime scan_time获取 扫描时间
func (obj *_WmsOperationPackageInfoMgr) WithScanTime(scanTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["scan_time"] = scanTime })
}

// WithScanUserID scan_user_id获取 扫描人用户id
func (obj *_WmsOperationPackageInfoMgr) WithScanUserID(scanUserID int) Option {
	return optionFunc(func(o *options) { o.query["scan_user_id"] = scanUserID })
}

// WithScanUserName scan_user_name获取 扫描人用户名字
func (obj *_WmsOperationPackageInfoMgr) WithScanUserName(scanUserName string) Option {
	return optionFunc(func(o *options) { o.query["scan_user_name"] = scanUserName })
}

// WithPackageUserID package_user_id获取 打包人用户id
func (obj *_WmsOperationPackageInfoMgr) WithPackageUserID(packageUserID int) Option {
	return optionFunc(func(o *options) { o.query["package_user_id"] = packageUserID })
}

// WithPackageUserName package_user_name获取 打包人名称
func (obj *_WmsOperationPackageInfoMgr) WithPackageUserName(packageUserName string) Option {
	return optionFunc(func(o *options) { o.query["package_user_name"] = packageUserName })
}

// WithPackageTime package_time获取 打包时间
func (obj *_WmsOperationPackageInfoMgr) WithPackageTime(packageTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["package_time"] = packageTime })
}

// WithFitterUserID fitter_user_id获取 装配人用户id
func (obj *_WmsOperationPackageInfoMgr) WithFitterUserID(fitterUserID int) Option {
	return optionFunc(func(o *options) { o.query["fitter_user_id"] = fitterUserID })
}

// WithFitterUserName fitter_user_name获取 装配人用户名称
func (obj *_WmsOperationPackageInfoMgr) WithFitterUserName(fitterUserName string) Option {
	return optionFunc(func(o *options) { o.query["fitter_user_name"] = fitterUserName })
}

// WithWarehouseCode warehouse_code获取 仓库code
func (obj *_WmsOperationPackageInfoMgr) WithWarehouseCode(warehouseCode string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_code"] = warehouseCode })
}

// WithWarehouseName warehouse_name获取 仓库名字
func (obj *_WmsOperationPackageInfoMgr) WithWarehouseName(warehouseName string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_name"] = warehouseName })
}

// WithScanStatus scan_status获取 扫描状态 done:已完成，unconfirmed:待确认(包裹异常，如超重、欠费等)，initial：初始化(包裹信息填写错误时，重新录入)
func (obj *_WmsOperationPackageInfoMgr) WithScanStatus(scanStatus string) Option {
	return optionFunc(func(o *options) { o.query["scan_status"] = scanStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsOperationPackageInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsOperationPackageInfoMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsOperationPackageInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsOperationPackageInfoMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsOperationPackageInfoMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsOperationPackageInfoMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithLength length获取 长 单位:cm
func (obj *_WmsOperationPackageInfoMgr) WithLength(length float64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithWidth width获取 宽 单位:cm
func (obj *_WmsOperationPackageInfoMgr) WithWidth(width float64) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

// WithHeight height获取 高 单位:cm
func (obj *_WmsOperationPackageInfoMgr) WithHeight(height float64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// GetByOption 功能选项模式获取
func (obj *_WmsOperationPackageInfoMgr) GetByOption(opts ...Option) (result WmsOperationPackageInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsOperationPackageInfoMgr) GetByOptions(opts ...Option) (results []*WmsOperationPackageInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsOperationPackageInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsOperationPackageInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where(options.query)
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
func (obj *_WmsOperationPackageInfoMgr) GetFromID(id int) (result WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromID(ids []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTrackingNo 通过tracking_no获取内容 快递单号
func (obj *_WmsOperationPackageInfoMgr) GetFromTrackingNo(trackingNo string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`tracking_no` = ?", trackingNo).Find(&results).Error

	return
}

// GetBatchFromTrackingNo 批量查找 快递单号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromTrackingNo(trackingNos []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`tracking_no` IN (?)", trackingNos).Find(&results).Error

	return
}

// GetFromTrackingNoType 通过tracking_no_type获取内容 快递单号类型，reference:参考号，transfer:转单号
func (obj *_WmsOperationPackageInfoMgr) GetFromTrackingNoType(trackingNoType string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`tracking_no_type` = ?", trackingNoType).Find(&results).Error

	return
}

// GetBatchFromTrackingNoType 批量查找 快递单号类型，reference:参考号，transfer:转单号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromTrackingNoType(trackingNoTypes []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`tracking_no_type` IN (?)", trackingNoTypes).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_WmsOperationPackageInfoMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_WmsOperationPackageInfoMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_WmsOperationPackageInfoMgr) GetFromReferenceNumber(referenceNumber string) (result WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`reference_number` = ?", referenceNumber).First(&result).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromReceiveCountry 通过receive_country获取内容 收件人国家
func (obj *_WmsOperationPackageInfoMgr) GetFromReceiveCountry(receiveCountry string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// GetBatchFromReceiveCountry 批量查找 收件人国家
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromReceiveCountry(receiveCountrys []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`receive_country` IN (?)", receiveCountrys).Find(&results).Error

	return
}

// GetFromOrderNo 通过order_no获取内容 订单号
func (obj *_WmsOperationPackageInfoMgr) GetFromOrderNo(orderNo string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`order_no` = ?", orderNo).Find(&results).Error

	return
}

// GetBatchFromOrderNo 批量查找 订单号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromOrderNo(orderNos []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`order_no` IN (?)", orderNos).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_WmsOperationPackageInfoMgr) GetFromCustomerID(customerID int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromCustomerID(customerIDs []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容 来源，0:内部，1:外部
func (obj *_WmsOperationPackageInfoMgr) GetFromSource(source int8) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找 来源，0:内部，1:外部
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromSource(sources []int8) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道id
func (obj *_WmsOperationPackageInfoMgr) GetFromCustomerChannelID(customerChannelID int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道id
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromDownstreamCode 通过downstream_code获取内容 下家code
func (obj *_WmsOperationPackageInfoMgr) GetFromDownstreamCode(downstreamCode string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`downstream_code` = ?", downstreamCode).Find(&results).Error

	return
}

// GetBatchFromDownstreamCode 批量查找 下家code
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromDownstreamCode(downstreamCodes []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`downstream_code` IN (?)", downstreamCodes).Find(&results).Error

	return
}

// GetFromAPIChannelCode 通过api_channel_code获取内容 api渠道代码
func (obj *_WmsOperationPackageInfoMgr) GetFromAPIChannelCode(apiChannelCode string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`api_channel_code` = ?", apiChannelCode).Find(&results).Error

	return
}

// GetBatchFromAPIChannelCode 批量查找 api渠道代码
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromAPIChannelCode(apiChannelCodes []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`api_channel_code` IN (?)", apiChannelCodes).Find(&results).Error

	return
}

// GetFromBigBagNo 通过big_bag_no获取内容 大包号
func (obj *_WmsOperationPackageInfoMgr) GetFromBigBagNo(bigBagNo string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`big_bag_no` = ?", bigBagNo).Find(&results).Error

	return
}

// GetBatchFromBigBagNo 批量查找 大包号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromBigBagNo(bigBagNos []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`big_bag_no` IN (?)", bigBagNos).Find(&results).Error

	return
}

// GetFromTotalListNo 通过total_list_no获取内容 总单号
func (obj *_WmsOperationPackageInfoMgr) GetFromTotalListNo(totalListNo string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// GetBatchFromTotalListNo 批量查找 总单号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromTotalListNo(totalListNos []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`total_list_no` IN (?)", totalListNos).Find(&results).Error

	return
}

// GetFromAirBillNumber 通过air_bill_number获取内容 航空提单号
func (obj *_WmsOperationPackageInfoMgr) GetFromAirBillNumber(airBillNumber string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`air_bill_number` = ?", airBillNumber).Find(&results).Error

	return
}

// GetBatchFromAirBillNumber 批量查找 航空提单号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromAirBillNumber(airBillNumbers []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`air_bill_number` IN (?)", airBillNumbers).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 包裹重量(KG),保留4位小数
func (obj *_WmsOperationPackageInfoMgr) GetFromWeight(weight float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 包裹重量(KG),保留4位小数
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromWeight(weights []float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromOrderWeight 通过order_weight获取内容 订单预报重量
func (obj *_WmsOperationPackageInfoMgr) GetFromOrderWeight(orderWeight float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`order_weight` = ?", orderWeight).Find(&results).Error

	return
}

// GetBatchFromOrderWeight 批量查找 订单预报重量
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromOrderWeight(orderWeights []float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`order_weight` IN (?)", orderWeights).Find(&results).Error

	return
}

// GetFromWorkConsole 通过work_console获取内容 操作台号
func (obj *_WmsOperationPackageInfoMgr) GetFromWorkConsole(workConsole string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`work_console` = ?", workConsole).Find(&results).Error

	return
}

// GetBatchFromWorkConsole 批量查找 操作台号
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromWorkConsole(workConsoles []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`work_console` IN (?)", workConsoles).Find(&results).Error

	return
}

// GetFromPicOssLink 通过pic_oss_link获取内容 包裹图片oss地址
func (obj *_WmsOperationPackageInfoMgr) GetFromPicOssLink(picOssLink string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`pic_oss_link` = ?", picOssLink).Find(&results).Error

	return
}

// GetBatchFromPicOssLink 批量查找 包裹图片oss地址
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromPicOssLink(picOssLinks []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`pic_oss_link` IN (?)", picOssLinks).Find(&results).Error

	return
}

// GetFromScanTime 通过scan_time获取内容 扫描时间
func (obj *_WmsOperationPackageInfoMgr) GetFromScanTime(scanTime time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_time` = ?", scanTime).Find(&results).Error

	return
}

// GetBatchFromScanTime 批量查找 扫描时间
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromScanTime(scanTimes []time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_time` IN (?)", scanTimes).Find(&results).Error

	return
}

// GetFromScanUserID 通过scan_user_id获取内容 扫描人用户id
func (obj *_WmsOperationPackageInfoMgr) GetFromScanUserID(scanUserID int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_user_id` = ?", scanUserID).Find(&results).Error

	return
}

// GetBatchFromScanUserID 批量查找 扫描人用户id
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromScanUserID(scanUserIDs []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_user_id` IN (?)", scanUserIDs).Find(&results).Error

	return
}

// GetFromScanUserName 通过scan_user_name获取内容 扫描人用户名字
func (obj *_WmsOperationPackageInfoMgr) GetFromScanUserName(scanUserName string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_user_name` = ?", scanUserName).Find(&results).Error

	return
}

// GetBatchFromScanUserName 批量查找 扫描人用户名字
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromScanUserName(scanUserNames []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_user_name` IN (?)", scanUserNames).Find(&results).Error

	return
}

// GetFromPackageUserID 通过package_user_id获取内容 打包人用户id
func (obj *_WmsOperationPackageInfoMgr) GetFromPackageUserID(packageUserID int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`package_user_id` = ?", packageUserID).Find(&results).Error

	return
}

// GetBatchFromPackageUserID 批量查找 打包人用户id
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromPackageUserID(packageUserIDs []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`package_user_id` IN (?)", packageUserIDs).Find(&results).Error

	return
}

// GetFromPackageUserName 通过package_user_name获取内容 打包人名称
func (obj *_WmsOperationPackageInfoMgr) GetFromPackageUserName(packageUserName string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`package_user_name` = ?", packageUserName).Find(&results).Error

	return
}

// GetBatchFromPackageUserName 批量查找 打包人名称
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromPackageUserName(packageUserNames []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`package_user_name` IN (?)", packageUserNames).Find(&results).Error

	return
}

// GetFromPackageTime 通过package_time获取内容 打包时间
func (obj *_WmsOperationPackageInfoMgr) GetFromPackageTime(packageTime time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`package_time` = ?", packageTime).Find(&results).Error

	return
}

// GetBatchFromPackageTime 批量查找 打包时间
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromPackageTime(packageTimes []time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`package_time` IN (?)", packageTimes).Find(&results).Error

	return
}

// GetFromFitterUserID 通过fitter_user_id获取内容 装配人用户id
func (obj *_WmsOperationPackageInfoMgr) GetFromFitterUserID(fitterUserID int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`fitter_user_id` = ?", fitterUserID).Find(&results).Error

	return
}

// GetBatchFromFitterUserID 批量查找 装配人用户id
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromFitterUserID(fitterUserIDs []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`fitter_user_id` IN (?)", fitterUserIDs).Find(&results).Error

	return
}

// GetFromFitterUserName 通过fitter_user_name获取内容 装配人用户名称
func (obj *_WmsOperationPackageInfoMgr) GetFromFitterUserName(fitterUserName string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`fitter_user_name` = ?", fitterUserName).Find(&results).Error

	return
}

// GetBatchFromFitterUserName 批量查找 装配人用户名称
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromFitterUserName(fitterUserNames []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`fitter_user_name` IN (?)", fitterUserNames).Find(&results).Error

	return
}

// GetFromWarehouseCode 通过warehouse_code获取内容 仓库code
func (obj *_WmsOperationPackageInfoMgr) GetFromWarehouseCode(warehouseCode string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`warehouse_code` = ?", warehouseCode).Find(&results).Error

	return
}

// GetBatchFromWarehouseCode 批量查找 仓库code
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromWarehouseCode(warehouseCodes []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`warehouse_code` IN (?)", warehouseCodes).Find(&results).Error

	return
}

// GetFromWarehouseName 通过warehouse_name获取内容 仓库名字
func (obj *_WmsOperationPackageInfoMgr) GetFromWarehouseName(warehouseName string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`warehouse_name` = ?", warehouseName).Find(&results).Error

	return
}

// GetBatchFromWarehouseName 批量查找 仓库名字
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromWarehouseName(warehouseNames []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`warehouse_name` IN (?)", warehouseNames).Find(&results).Error

	return
}

// GetFromScanStatus 通过scan_status获取内容 扫描状态 done:已完成，unconfirmed:待确认(包裹异常，如超重、欠费等)，initial：初始化(包裹信息填写错误时，重新录入)
func (obj *_WmsOperationPackageInfoMgr) GetFromScanStatus(scanStatus string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_status` = ?", scanStatus).Find(&results).Error

	return
}

// GetBatchFromScanStatus 批量查找 扫描状态 done:已完成，unconfirmed:待确认(包裹异常，如超重、欠费等)，initial：初始化(包裹信息填写错误时，重新录入)
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromScanStatus(scanStatuss []string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_status` IN (?)", scanStatuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsOperationPackageInfoMgr) GetFromCreateTime(createTime time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsOperationPackageInfoMgr) GetFromCreateUser(createUser int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsOperationPackageInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsOperationPackageInfoMgr) GetFromUpdateUser(updateUser int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsOperationPackageInfoMgr) GetFromVersion(version int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromVersion(versions []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsOperationPackageInfoMgr) GetFromDeleted(deleted int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容 长 单位:cm
func (obj *_WmsOperationPackageInfoMgr) GetFromLength(length float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找 长 单位:cm
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromLength(lengths []float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromWidth 通过width获取内容 宽 单位:cm
func (obj *_WmsOperationPackageInfoMgr) GetFromWidth(width float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`width` = ?", width).Find(&results).Error

	return
}

// GetBatchFromWidth 批量查找 宽 单位:cm
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromWidth(widths []float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`width` IN (?)", widths).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 高 单位:cm
func (obj *_WmsOperationPackageInfoMgr) GetFromHeight(height float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`height` = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 高 单位:cm
func (obj *_WmsOperationPackageInfoMgr) GetBatchFromHeight(heights []float64) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`height` IN (?)", heights).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsOperationPackageInfoMgr) FetchByPrimaryKey(id int) (result WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexReferenceNumber primary or index 获取唯一内容
func (obj *_WmsOperationPackageInfoMgr) FetchUniqueByIndexReferenceNumber(referenceNumber string) (result WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`reference_number` = ?", referenceNumber).First(&result).Error

	return
}

// FetchIndexByWmsOperationPackageInfoTrackingNoIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoTrackingNoIndex(trackingNo string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`tracking_no` = ?", trackingNo).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoTrackingNoTypeIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoTrackingNoTypeIndex(trackingNoType string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`tracking_no_type` = ?", trackingNoType).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoLogisticsNumberIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoLogisticsNumberIndex(logisticsNumber string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoLogisticsNumberFinalIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoLogisticsNumberFinalIndex(logisticsNumberFinal string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexByIDxOrdernoDeletedScantime  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByIDxOrdernoDeletedScantime(orderNo string, scanTime time.Time, deleted int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`order_no` = ? AND `scan_time` = ? AND `deleted` = ?", orderNo, scanTime, deleted).Find(&results).Error

	return
}

// FetchIndexBySourceIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexBySourceIndex(source int8) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoBigBagNoIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoBigBagNoIndex(bigBagNo string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`big_bag_no` = ?", bigBagNo).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoTotalListNoIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoTotalListNoIndex(totalListNo string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoAirBillNumberIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoAirBillNumberIndex(airBillNumber string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`air_bill_number` = ?", airBillNumber).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoWorkConsoleIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoWorkConsoleIndex(workConsole string) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`work_console` = ?", workConsole).Find(&results).Error

	return
}

// FetchIndexByIDxDeletedCreatetimeScantime  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByIDxDeletedCreatetimeScantime(scanTime time.Time, createTime time.Time, deleted int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_time` = ? AND `create_time` = ? AND `deleted` = ?", scanTime, createTime, deleted).Find(&results).Error

	return
}

// FetchIndexByIDxDeletedScantime  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByIDxDeletedScantime(scanTime time.Time, deleted int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_time` = ? AND `deleted` = ?", scanTime, deleted).Find(&results).Error

	return
}

// FetchIndexByIDxScanuseridDeletedScantime  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByIDxScanuseridDeletedScantime(scanTime time.Time, scanUserID int, deleted int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_time` = ? AND `scan_user_id` = ? AND `deleted` = ?", scanTime, scanUserID, deleted).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoScanTimeIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoScanTimeIndex(scanTime time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_time` = ?", scanTime).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoScanUserIDIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoScanUserIDIndex(scanUserID int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`scan_user_id` = ?", scanUserID).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoPackageTimeIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoPackageTimeIndex(packageTime time.Time) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`package_time` = ?", packageTime).Find(&results).Error

	return
}

// FetchIndexByWmsOperationPackageInfoDeletedIndex  获取多个内容
func (obj *_WmsOperationPackageInfoMgr) FetchIndexByWmsOperationPackageInfoDeletedIndex(deleted int) (results []*WmsOperationPackageInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationPackageInfo{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}
