package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _WmsOperationTotalListInfoMgr struct {
	*_BaseMgr
}

// WmsOperationTotalListInfoMgr open func
func WmsOperationTotalListInfoMgr(db *gorm.DB) *_WmsOperationTotalListInfoMgr {
	if db == nil {
		panic(fmt.Errorf("WmsOperationTotalListInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsOperationTotalListInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_operation_total_list_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsOperationTotalListInfoMgr) GetTableName() string {
	return "wms_operation_total_list_info"
}

// Reset 重置gorm会话
func (obj *_WmsOperationTotalListInfoMgr) Reset() *_WmsOperationTotalListInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsOperationTotalListInfoMgr) Get() (result WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsOperationTotalListInfoMgr) Gets() (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsOperationTotalListInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsOperationTotalListInfoMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTotalListNo total_list_no获取 总单号
func (obj *_WmsOperationTotalListInfoMgr) WithTotalListNo(totalListNo string) Option {
	return optionFunc(func(o *options) { o.query["total_list_no"] = totalListNo })
}

// WithTotalListType total_list_type获取 总单类型: normal-普货，battery-纯电，站点扫描
func (obj *_WmsOperationTotalListInfoMgr) WithTotalListType(totalListType string) Option {
	return optionFunc(func(o *options) { o.query["total_list_type"] = totalListType })
}

// WithAPITotalListNo api_total_list_no获取 api总单号
func (obj *_WmsOperationTotalListInfoMgr) WithAPITotalListNo(apiTotalListNo string) Option {
	return optionFunc(func(o *options) { o.query["api_total_list_no"] = apiTotalListNo })
}

// WithCustomerTotalListNo customer_total_list_no获取 客户总单号
func (obj *_WmsOperationTotalListInfoMgr) WithCustomerTotalListNo(customerTotalListNo string) Option {
	return optionFunc(func(o *options) { o.query["customer_total_list_no"] = customerTotalListNo })
}

// WithExpressWaybillOssLink express_waybill_oss_link获取 总单面单OSS地址
func (obj *_WmsOperationTotalListInfoMgr) WithExpressWaybillOssLink(expressWaybillOssLink string) Option {
	return optionFunc(func(o *options) { o.query["express_waybill_oss_link"] = expressWaybillOssLink })
}

// WithFlightNo flight_no获取 航班号
func (obj *_WmsOperationTotalListInfoMgr) WithFlightNo(flightNo string) Option {
	return optionFunc(func(o *options) { o.query["flight_no"] = flightNo })
}

// WithDownstream downstream获取 下家
func (obj *_WmsOperationTotalListInfoMgr) WithDownstream(downstream string) Option {
	return optionFunc(func(o *options) { o.query["downstream"] = downstream })
}

// WithWeight weight获取 重量
func (obj *_WmsOperationTotalListInfoMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithBigBagCount big_bag_count获取 大包数量
func (obj *_WmsOperationTotalListInfoMgr) WithBigBagCount(bigBagCount int) Option {
	return optionFunc(func(o *options) { o.query["big_bag_count"] = bigBagCount })
}

// WithSmallBagCount small_bag_count获取 小包数量
func (obj *_WmsOperationTotalListInfoMgr) WithSmallBagCount(smallBagCount int) Option {
	return optionFunc(func(o *options) { o.query["small_bag_count"] = smallBagCount })
}

// WithFitterUserID fitter_user_id获取 装配人用户id
func (obj *_WmsOperationTotalListInfoMgr) WithFitterUserID(fitterUserID int) Option {
	return optionFunc(func(o *options) { o.query["fitter_user_id"] = fitterUserID })
}

// WithFitterUserName fitter_user_name获取 装配人用户名称
func (obj *_WmsOperationTotalListInfoMgr) WithFitterUserName(fitterUserName string) Option {
	return optionFunc(func(o *options) { o.query["fitter_user_name"] = fitterUserName })
}

// WithFitStatus fit_status获取 装配状态，initial:未装配，done:已装配
func (obj *_WmsOperationTotalListInfoMgr) WithFitStatus(fitStatus string) Option {
	return optionFunc(func(o *options) { o.query["fit_status"] = fitStatus })
}

// WithFitOverTime fit_over_time获取 装配完结时间
func (obj *_WmsOperationTotalListInfoMgr) WithFitOverTime(fitOverTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["fit_over_time"] = fitOverTime })
}

// WithWarehouseCode warehouse_code获取 仓库code
func (obj *_WmsOperationTotalListInfoMgr) WithWarehouseCode(warehouseCode string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_code"] = warehouseCode })
}

// WithWarehouseName warehouse_name获取 仓库名字
func (obj *_WmsOperationTotalListInfoMgr) WithWarehouseName(warehouseName string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_name"] = warehouseName })
}

// WithExtra extra获取 扩展字段
func (obj *_WmsOperationTotalListInfoMgr) WithExtra(extra datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithRemark remark获取 备注
func (obj *_WmsOperationTotalListInfoMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsOperationTotalListInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsOperationTotalListInfoMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsOperationTotalListInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsOperationTotalListInfoMgr) WithUpdateUser(updateUser uint32) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsOperationTotalListInfoMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsOperationTotalListInfoMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsOperationTotalListInfoMgr) GetByOption(opts ...Option) (result WmsOperationTotalListInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsOperationTotalListInfoMgr) GetByOptions(opts ...Option) (results []*WmsOperationTotalListInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsOperationTotalListInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsOperationTotalListInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where(options.query)
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
func (obj *_WmsOperationTotalListInfoMgr) GetFromID(id int) (result WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromID(ids []int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTotalListNo 通过total_list_no获取内容 总单号
func (obj *_WmsOperationTotalListInfoMgr) GetFromTotalListNo(totalListNo string) (result WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`total_list_no` = ?", totalListNo).First(&result).Error

	return
}

// GetBatchFromTotalListNo 批量查找 总单号
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromTotalListNo(totalListNos []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`total_list_no` IN (?)", totalListNos).Find(&results).Error

	return
}

// GetFromTotalListType 通过total_list_type获取内容 总单类型: normal-普货，battery-纯电，站点扫描
func (obj *_WmsOperationTotalListInfoMgr) GetFromTotalListType(totalListType string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`total_list_type` = ?", totalListType).Find(&results).Error

	return
}

// GetBatchFromTotalListType 批量查找 总单类型: normal-普货，battery-纯电，站点扫描
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromTotalListType(totalListTypes []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`total_list_type` IN (?)", totalListTypes).Find(&results).Error

	return
}

// GetFromAPITotalListNo 通过api_total_list_no获取内容 api总单号
func (obj *_WmsOperationTotalListInfoMgr) GetFromAPITotalListNo(apiTotalListNo string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`api_total_list_no` = ?", apiTotalListNo).Find(&results).Error

	return
}

// GetBatchFromAPITotalListNo 批量查找 api总单号
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromAPITotalListNo(apiTotalListNos []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`api_total_list_no` IN (?)", apiTotalListNos).Find(&results).Error

	return
}

// GetFromCustomerTotalListNo 通过customer_total_list_no获取内容 客户总单号
func (obj *_WmsOperationTotalListInfoMgr) GetFromCustomerTotalListNo(customerTotalListNo string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`customer_total_list_no` = ?", customerTotalListNo).Find(&results).Error

	return
}

// GetBatchFromCustomerTotalListNo 批量查找 客户总单号
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromCustomerTotalListNo(customerTotalListNos []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`customer_total_list_no` IN (?)", customerTotalListNos).Find(&results).Error

	return
}

// GetFromExpressWaybillOssLink 通过express_waybill_oss_link获取内容 总单面单OSS地址
func (obj *_WmsOperationTotalListInfoMgr) GetFromExpressWaybillOssLink(expressWaybillOssLink string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`express_waybill_oss_link` = ?", expressWaybillOssLink).Find(&results).Error

	return
}

// GetBatchFromExpressWaybillOssLink 批量查找 总单面单OSS地址
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromExpressWaybillOssLink(expressWaybillOssLinks []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`express_waybill_oss_link` IN (?)", expressWaybillOssLinks).Find(&results).Error

	return
}

// GetFromFlightNo 通过flight_no获取内容 航班号
func (obj *_WmsOperationTotalListInfoMgr) GetFromFlightNo(flightNo string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`flight_no` = ?", flightNo).Find(&results).Error

	return
}

// GetBatchFromFlightNo 批量查找 航班号
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromFlightNo(flightNos []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`flight_no` IN (?)", flightNos).Find(&results).Error

	return
}

// GetFromDownstream 通过downstream获取内容 下家
func (obj *_WmsOperationTotalListInfoMgr) GetFromDownstream(downstream string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`downstream` = ?", downstream).Find(&results).Error

	return
}

// GetBatchFromDownstream 批量查找 下家
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromDownstream(downstreams []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`downstream` IN (?)", downstreams).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *_WmsOperationTotalListInfoMgr) GetFromWeight(weight float64) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromWeight(weights []float64) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromBigBagCount 通过big_bag_count获取内容 大包数量
func (obj *_WmsOperationTotalListInfoMgr) GetFromBigBagCount(bigBagCount int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`big_bag_count` = ?", bigBagCount).Find(&results).Error

	return
}

// GetBatchFromBigBagCount 批量查找 大包数量
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromBigBagCount(bigBagCounts []int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`big_bag_count` IN (?)", bigBagCounts).Find(&results).Error

	return
}

// GetFromSmallBagCount 通过small_bag_count获取内容 小包数量
func (obj *_WmsOperationTotalListInfoMgr) GetFromSmallBagCount(smallBagCount int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`small_bag_count` = ?", smallBagCount).Find(&results).Error

	return
}

// GetBatchFromSmallBagCount 批量查找 小包数量
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromSmallBagCount(smallBagCounts []int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`small_bag_count` IN (?)", smallBagCounts).Find(&results).Error

	return
}

// GetFromFitterUserID 通过fitter_user_id获取内容 装配人用户id
func (obj *_WmsOperationTotalListInfoMgr) GetFromFitterUserID(fitterUserID int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fitter_user_id` = ?", fitterUserID).Find(&results).Error

	return
}

// GetBatchFromFitterUserID 批量查找 装配人用户id
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromFitterUserID(fitterUserIDs []int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fitter_user_id` IN (?)", fitterUserIDs).Find(&results).Error

	return
}

// GetFromFitterUserName 通过fitter_user_name获取内容 装配人用户名称
func (obj *_WmsOperationTotalListInfoMgr) GetFromFitterUserName(fitterUserName string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fitter_user_name` = ?", fitterUserName).Find(&results).Error

	return
}

// GetBatchFromFitterUserName 批量查找 装配人用户名称
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromFitterUserName(fitterUserNames []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fitter_user_name` IN (?)", fitterUserNames).Find(&results).Error

	return
}

// GetFromFitStatus 通过fit_status获取内容 装配状态，initial:未装配，done:已装配
func (obj *_WmsOperationTotalListInfoMgr) GetFromFitStatus(fitStatus string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fit_status` = ?", fitStatus).Find(&results).Error

	return
}

// GetBatchFromFitStatus 批量查找 装配状态，initial:未装配，done:已装配
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromFitStatus(fitStatuss []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fit_status` IN (?)", fitStatuss).Find(&results).Error

	return
}

// GetFromFitOverTime 通过fit_over_time获取内容 装配完结时间
func (obj *_WmsOperationTotalListInfoMgr) GetFromFitOverTime(fitOverTime time.Time) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fit_over_time` = ?", fitOverTime).Find(&results).Error

	return
}

// GetBatchFromFitOverTime 批量查找 装配完结时间
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromFitOverTime(fitOverTimes []time.Time) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fit_over_time` IN (?)", fitOverTimes).Find(&results).Error

	return
}

// GetFromWarehouseCode 通过warehouse_code获取内容 仓库code
func (obj *_WmsOperationTotalListInfoMgr) GetFromWarehouseCode(warehouseCode string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`warehouse_code` = ?", warehouseCode).Find(&results).Error

	return
}

// GetBatchFromWarehouseCode 批量查找 仓库code
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromWarehouseCode(warehouseCodes []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`warehouse_code` IN (?)", warehouseCodes).Find(&results).Error

	return
}

// GetFromWarehouseName 通过warehouse_name获取内容 仓库名字
func (obj *_WmsOperationTotalListInfoMgr) GetFromWarehouseName(warehouseName string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`warehouse_name` = ?", warehouseName).Find(&results).Error

	return
}

// GetBatchFromWarehouseName 批量查找 仓库名字
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromWarehouseName(warehouseNames []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`warehouse_name` IN (?)", warehouseNames).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 扩展字段
func (obj *_WmsOperationTotalListInfoMgr) GetFromExtra(extra datatypes.JSON) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 扩展字段
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromExtra(extras []datatypes.JSON) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WmsOperationTotalListInfoMgr) GetFromRemark(remark string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromRemark(remarks []string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsOperationTotalListInfoMgr) GetFromCreateTime(createTime time.Time) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsOperationTotalListInfoMgr) GetFromCreateUser(createUser int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsOperationTotalListInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsOperationTotalListInfoMgr) GetFromUpdateUser(updateUser uint32) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromUpdateUser(updateUsers []uint32) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsOperationTotalListInfoMgr) GetFromVersion(version int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromVersion(versions []int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsOperationTotalListInfoMgr) GetFromDeleted(deleted int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsOperationTotalListInfoMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsOperationTotalListInfoMgr) FetchByPrimaryKey(id int) (result WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexTotalListNo primary or index 获取唯一内容
func (obj *_WmsOperationTotalListInfoMgr) FetchUniqueByIndexTotalListNo(totalListNo string) (result WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`total_list_no` = ?", totalListNo).First(&result).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoTotalListTypeIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoTotalListTypeIndex(totalListType string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`total_list_type` = ?", totalListType).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoAPITotalListNoIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoAPITotalListNoIndex(apiTotalListNo string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`api_total_list_no` = ?", apiTotalListNo).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoCustomerTotalListNoIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoCustomerTotalListNoIndex(customerTotalListNo string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`customer_total_list_no` = ?", customerTotalListNo).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoFlightNoIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoFlightNoIndex(flightNo string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`flight_no` = ?", flightNo).Find(&results).Error

	return
}

// FetchIndexByIndexDownstream  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByIndexDownstream(downstream string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`downstream` = ?", downstream).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoFitterUserIDIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoFitterUserIDIndex(fitterUserID int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fitter_user_id` = ?", fitterUserID).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoFitStatusIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoFitStatusIndex(fitStatus string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fit_status` = ?", fitStatus).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoFitOverTimeIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoFitOverTimeIndex(fitOverTime time.Time) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`fit_over_time` = ?", fitOverTime).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoWarehouseCodeIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoWarehouseCodeIndex(warehouseCode string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`warehouse_code` = ?", warehouseCode).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoRemarkIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoRemarkIndex(remark string) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoCreateTimeIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoCreateTimeIndex(createTime time.Time) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByWmsOperationTotalListInfoDeletedIndex  获取多个内容
func (obj *_WmsOperationTotalListInfoMgr) FetchIndexByWmsOperationTotalListInfoDeletedIndex(deleted int) (results []*WmsOperationTotalListInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfo{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}
