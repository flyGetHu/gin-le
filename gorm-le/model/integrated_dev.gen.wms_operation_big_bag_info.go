package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _WmsOperationBigBagInfoMgr struct {
	*_BaseMgr
}

// WmsOperationBigBagInfoMgr open func
func WmsOperationBigBagInfoMgr(db *gorm.DB) *_WmsOperationBigBagInfoMgr {
	if db == nil {
		panic(fmt.Errorf("WmsOperationBigBagInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsOperationBigBagInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_operation_big_bag_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsOperationBigBagInfoMgr) GetTableName() string {
	return "wms_operation_big_bag_info"
}

// Reset 重置gorm会话
func (obj *_WmsOperationBigBagInfoMgr) Reset() *_WmsOperationBigBagInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsOperationBigBagInfoMgr) Get() (result WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsOperationBigBagInfoMgr) Gets() (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsOperationBigBagInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsOperationBigBagInfoMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBigBagNo big_bag_no获取 大包号
func (obj *_WmsOperationBigBagInfoMgr) WithBigBagNo(bigBagNo string) Option {
	return optionFunc(func(o *options) { o.query["big_bag_no"] = bigBagNo })
}

// WithCustomerBigBagNo customer_big_bag_no获取 客户大包号，来源为内部时，默认为大包号
func (obj *_WmsOperationBigBagInfoMgr) WithCustomerBigBagNo(customerBigBagNo string) Option {
	return optionFunc(func(o *options) { o.query["customer_big_bag_no"] = customerBigBagNo })
}

// WithCustomerID customer_id获取 客户id，来源为内部时，默认为0
func (obj *_WmsOperationBigBagInfoMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_WmsOperationBigBagInfoMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithSource source获取 来源，0:内部，1:外部 2自飞
func (obj *_WmsOperationBigBagInfoMgr) WithSource(source int8) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithDownstreamCode downstream_code获取 下家code
func (obj *_WmsOperationBigBagInfoMgr) WithDownstreamCode(downstreamCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_code"] = downstreamCode })
}

// WithCustomerChannelIDs customer_channel_ids获取 客户渠道id，多个以逗号隔开
func (obj *_WmsOperationBigBagInfoMgr) WithCustomerChannelIDs(customerChannelIDs string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_ids"] = customerChannelIDs })
}

// WithAPIBigBagNo api_big_bag_no获取 api大包号
func (obj *_WmsOperationBigBagInfoMgr) WithAPIBigBagNo(apiBigBagNo string) Option {
	return optionFunc(func(o *options) { o.query["api_big_bag_no"] = apiBigBagNo })
}

// WithExpressWaybillOssLink express_waybill_oss_link获取 大包面单OSS地址
func (obj *_WmsOperationBigBagInfoMgr) WithExpressWaybillOssLink(expressWaybillOssLink string) Option {
	return optionFunc(func(o *options) { o.query["express_waybill_oss_link"] = expressWaybillOssLink })
}

// WithTotalListNo total_list_no获取 总单号
func (obj *_WmsOperationBigBagInfoMgr) WithTotalListNo(totalListNo string) Option {
	return optionFunc(func(o *options) { o.query["total_list_no"] = totalListNo })
}

// WithAirBillNumber air_bill_number获取 航空提单号
func (obj *_WmsOperationBigBagInfoMgr) WithAirBillNumber(airBillNumber string) Option {
	return optionFunc(func(o *options) { o.query["air_bill_number"] = airBillNumber })
}

// WithLength length获取 长(cm)
func (obj *_WmsOperationBigBagInfoMgr) WithLength(length float64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithWidth width获取 宽(cm)
func (obj *_WmsOperationBigBagInfoMgr) WithWidth(width float64) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

// WithHeight height获取 高(cm)
func (obj *_WmsOperationBigBagInfoMgr) WithHeight(height float64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithWeight weight获取 重量
func (obj *_WmsOperationBigBagInfoMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithSmallBagCount small_bag_count获取 小包数量
func (obj *_WmsOperationBigBagInfoMgr) WithSmallBagCount(smallBagCount int) Option {
	return optionFunc(func(o *options) { o.query["small_bag_count"] = smallBagCount })
}

// WithBigBagStatus big_bag_status获取 大包状态，packaged:已打包、unpacked:未打包
func (obj *_WmsOperationBigBagInfoMgr) WithBigBagStatus(bigBagStatus string) Option {
	return optionFunc(func(o *options) { o.query["big_bag_status"] = bigBagStatus })
}

// WithPackageUserID package_user_id获取 打包人用户id
func (obj *_WmsOperationBigBagInfoMgr) WithPackageUserID(packageUserID int) Option {
	return optionFunc(func(o *options) { o.query["package_user_id"] = packageUserID })
}

// WithPackageUserName package_user_name获取 打包人名称
func (obj *_WmsOperationBigBagInfoMgr) WithPackageUserName(packageUserName string) Option {
	return optionFunc(func(o *options) { o.query["package_user_name"] = packageUserName })
}

// WithBackPackageUserName back_package_user_name获取 上一个打包人名称
func (obj *_WmsOperationBigBagInfoMgr) WithBackPackageUserName(backPackageUserName string) Option {
	return optionFunc(func(o *options) { o.query["back_package_user_name"] = backPackageUserName })
}

// WithBackPackageUserID back_package_user_id获取 上一个打包人id
func (obj *_WmsOperationBigBagInfoMgr) WithBackPackageUserID(backPackageUserID int) Option {
	return optionFunc(func(o *options) { o.query["back_package_user_id"] = backPackageUserID })
}

// WithFitterUserID fitter_user_id获取 装配人用户id
func (obj *_WmsOperationBigBagInfoMgr) WithFitterUserID(fitterUserID int) Option {
	return optionFunc(func(o *options) { o.query["fitter_user_id"] = fitterUserID })
}

// WithFitterUserName fitter_user_name获取 装配人用户名称
func (obj *_WmsOperationBigBagInfoMgr) WithFitterUserName(fitterUserName string) Option {
	return optionFunc(func(o *options) { o.query["fitter_user_name"] = fitterUserName })
}

// WithWarehouseCode warehouse_code获取 仓库code
func (obj *_WmsOperationBigBagInfoMgr) WithWarehouseCode(warehouseCode string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_code"] = warehouseCode })
}

// WithWarehouseName warehouse_name获取 仓库name
func (obj *_WmsOperationBigBagInfoMgr) WithWarehouseName(warehouseName string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_name"] = warehouseName })
}

// WithExtra extra获取 扩展
func (obj *_WmsOperationBigBagInfoMgr) WithExtra(extra datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsOperationBigBagInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsOperationBigBagInfoMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsOperationBigBagInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsOperationBigBagInfoMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsOperationBigBagInfoMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsOperationBigBagInfoMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsOperationBigBagInfoMgr) GetByOption(opts ...Option) (result WmsOperationBigBagInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsOperationBigBagInfoMgr) GetByOptions(opts ...Option) (results []*WmsOperationBigBagInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsOperationBigBagInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsOperationBigBagInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where(options.query)
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
func (obj *_WmsOperationBigBagInfoMgr) GetFromID(id int) (result WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromID(ids []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBigBagNo 通过big_bag_no获取内容 大包号
func (obj *_WmsOperationBigBagInfoMgr) GetFromBigBagNo(bigBagNo string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`big_bag_no` = ?", bigBagNo).Find(&results).Error

	return
}

// GetBatchFromBigBagNo 批量查找 大包号
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromBigBagNo(bigBagNos []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`big_bag_no` IN (?)", bigBagNos).Find(&results).Error

	return
}

// GetFromCustomerBigBagNo 通过customer_big_bag_no获取内容 客户大包号，来源为内部时，默认为大包号
func (obj *_WmsOperationBigBagInfoMgr) GetFromCustomerBigBagNo(customerBigBagNo string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_big_bag_no` = ?", customerBigBagNo).Find(&results).Error

	return
}

// GetBatchFromCustomerBigBagNo 批量查找 客户大包号，来源为内部时，默认为大包号
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromCustomerBigBagNo(customerBigBagNos []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_big_bag_no` IN (?)", customerBigBagNos).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id，来源为内部时，默认为0
func (obj *_WmsOperationBigBagInfoMgr) GetFromCustomerID(customerID int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id，来源为内部时，默认为0
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromCustomerID(customerIDs []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_WmsOperationBigBagInfoMgr) GetFromCustomerAlias(customerAlias string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容 来源，0:内部，1:外部 2自飞
func (obj *_WmsOperationBigBagInfoMgr) GetFromSource(source int8) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找 来源，0:内部，1:外部 2自飞
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromSource(sources []int8) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromDownstreamCode 通过downstream_code获取内容 下家code
func (obj *_WmsOperationBigBagInfoMgr) GetFromDownstreamCode(downstreamCode string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`downstream_code` = ?", downstreamCode).Find(&results).Error

	return
}

// GetBatchFromDownstreamCode 批量查找 下家code
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromDownstreamCode(downstreamCodes []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`downstream_code` IN (?)", downstreamCodes).Find(&results).Error

	return
}

// GetFromCustomerChannelIDs 通过customer_channel_ids获取内容 客户渠道id，多个以逗号隔开
func (obj *_WmsOperationBigBagInfoMgr) GetFromCustomerChannelIDs(customerChannelIDs string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_channel_ids` = ?", customerChannelIDs).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelIDs 批量查找 客户渠道id，多个以逗号隔开
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromCustomerChannelIDs(customerChannelIDss []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_channel_ids` IN (?)", customerChannelIDss).Find(&results).Error

	return
}

// GetFromAPIBigBagNo 通过api_big_bag_no获取内容 api大包号
func (obj *_WmsOperationBigBagInfoMgr) GetFromAPIBigBagNo(apiBigBagNo string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`api_big_bag_no` = ?", apiBigBagNo).Find(&results).Error

	return
}

// GetBatchFromAPIBigBagNo 批量查找 api大包号
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromAPIBigBagNo(apiBigBagNos []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`api_big_bag_no` IN (?)", apiBigBagNos).Find(&results).Error

	return
}

// GetFromExpressWaybillOssLink 通过express_waybill_oss_link获取内容 大包面单OSS地址
func (obj *_WmsOperationBigBagInfoMgr) GetFromExpressWaybillOssLink(expressWaybillOssLink string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`express_waybill_oss_link` = ?", expressWaybillOssLink).Find(&results).Error

	return
}

// GetBatchFromExpressWaybillOssLink 批量查找 大包面单OSS地址
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromExpressWaybillOssLink(expressWaybillOssLinks []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`express_waybill_oss_link` IN (?)", expressWaybillOssLinks).Find(&results).Error

	return
}

// GetFromTotalListNo 通过total_list_no获取内容 总单号
func (obj *_WmsOperationBigBagInfoMgr) GetFromTotalListNo(totalListNo string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// GetBatchFromTotalListNo 批量查找 总单号
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromTotalListNo(totalListNos []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`total_list_no` IN (?)", totalListNos).Find(&results).Error

	return
}

// GetFromAirBillNumber 通过air_bill_number获取内容 航空提单号
func (obj *_WmsOperationBigBagInfoMgr) GetFromAirBillNumber(airBillNumber string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`air_bill_number` = ?", airBillNumber).Find(&results).Error

	return
}

// GetBatchFromAirBillNumber 批量查找 航空提单号
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromAirBillNumber(airBillNumbers []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`air_bill_number` IN (?)", airBillNumbers).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容 长(cm)
func (obj *_WmsOperationBigBagInfoMgr) GetFromLength(length float64) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找 长(cm)
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromLength(lengths []float64) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromWidth 通过width获取内容 宽(cm)
func (obj *_WmsOperationBigBagInfoMgr) GetFromWidth(width float64) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`width` = ?", width).Find(&results).Error

	return
}

// GetBatchFromWidth 批量查找 宽(cm)
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromWidth(widths []float64) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`width` IN (?)", widths).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 高(cm)
func (obj *_WmsOperationBigBagInfoMgr) GetFromHeight(height float64) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`height` = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 高(cm)
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromHeight(heights []float64) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`height` IN (?)", heights).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *_WmsOperationBigBagInfoMgr) GetFromWeight(weight float64) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromWeight(weights []float64) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromSmallBagCount 通过small_bag_count获取内容 小包数量
func (obj *_WmsOperationBigBagInfoMgr) GetFromSmallBagCount(smallBagCount int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`small_bag_count` = ?", smallBagCount).Find(&results).Error

	return
}

// GetBatchFromSmallBagCount 批量查找 小包数量
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromSmallBagCount(smallBagCounts []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`small_bag_count` IN (?)", smallBagCounts).Find(&results).Error

	return
}

// GetFromBigBagStatus 通过big_bag_status获取内容 大包状态，packaged:已打包、unpacked:未打包
func (obj *_WmsOperationBigBagInfoMgr) GetFromBigBagStatus(bigBagStatus string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`big_bag_status` = ?", bigBagStatus).Find(&results).Error

	return
}

// GetBatchFromBigBagStatus 批量查找 大包状态，packaged:已打包、unpacked:未打包
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromBigBagStatus(bigBagStatuss []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`big_bag_status` IN (?)", bigBagStatuss).Find(&results).Error

	return
}

// GetFromPackageUserID 通过package_user_id获取内容 打包人用户id
func (obj *_WmsOperationBigBagInfoMgr) GetFromPackageUserID(packageUserID int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`package_user_id` = ?", packageUserID).Find(&results).Error

	return
}

// GetBatchFromPackageUserID 批量查找 打包人用户id
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromPackageUserID(packageUserIDs []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`package_user_id` IN (?)", packageUserIDs).Find(&results).Error

	return
}

// GetFromPackageUserName 通过package_user_name获取内容 打包人名称
func (obj *_WmsOperationBigBagInfoMgr) GetFromPackageUserName(packageUserName string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`package_user_name` = ?", packageUserName).Find(&results).Error

	return
}

// GetBatchFromPackageUserName 批量查找 打包人名称
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromPackageUserName(packageUserNames []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`package_user_name` IN (?)", packageUserNames).Find(&results).Error

	return
}

// GetFromBackPackageUserName 通过back_package_user_name获取内容 上一个打包人名称
func (obj *_WmsOperationBigBagInfoMgr) GetFromBackPackageUserName(backPackageUserName string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`back_package_user_name` = ?", backPackageUserName).Find(&results).Error

	return
}

// GetBatchFromBackPackageUserName 批量查找 上一个打包人名称
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromBackPackageUserName(backPackageUserNames []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`back_package_user_name` IN (?)", backPackageUserNames).Find(&results).Error

	return
}

// GetFromBackPackageUserID 通过back_package_user_id获取内容 上一个打包人id
func (obj *_WmsOperationBigBagInfoMgr) GetFromBackPackageUserID(backPackageUserID int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`back_package_user_id` = ?", backPackageUserID).Find(&results).Error

	return
}

// GetBatchFromBackPackageUserID 批量查找 上一个打包人id
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromBackPackageUserID(backPackageUserIDs []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`back_package_user_id` IN (?)", backPackageUserIDs).Find(&results).Error

	return
}

// GetFromFitterUserID 通过fitter_user_id获取内容 装配人用户id
func (obj *_WmsOperationBigBagInfoMgr) GetFromFitterUserID(fitterUserID int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`fitter_user_id` = ?", fitterUserID).Find(&results).Error

	return
}

// GetBatchFromFitterUserID 批量查找 装配人用户id
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromFitterUserID(fitterUserIDs []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`fitter_user_id` IN (?)", fitterUserIDs).Find(&results).Error

	return
}

// GetFromFitterUserName 通过fitter_user_name获取内容 装配人用户名称
func (obj *_WmsOperationBigBagInfoMgr) GetFromFitterUserName(fitterUserName string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`fitter_user_name` = ?", fitterUserName).Find(&results).Error

	return
}

// GetBatchFromFitterUserName 批量查找 装配人用户名称
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromFitterUserName(fitterUserNames []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`fitter_user_name` IN (?)", fitterUserNames).Find(&results).Error

	return
}

// GetFromWarehouseCode 通过warehouse_code获取内容 仓库code
func (obj *_WmsOperationBigBagInfoMgr) GetFromWarehouseCode(warehouseCode string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`warehouse_code` = ?", warehouseCode).Find(&results).Error

	return
}

// GetBatchFromWarehouseCode 批量查找 仓库code
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromWarehouseCode(warehouseCodes []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`warehouse_code` IN (?)", warehouseCodes).Find(&results).Error

	return
}

// GetFromWarehouseName 通过warehouse_name获取内容 仓库name
func (obj *_WmsOperationBigBagInfoMgr) GetFromWarehouseName(warehouseName string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`warehouse_name` = ?", warehouseName).Find(&results).Error

	return
}

// GetBatchFromWarehouseName 批量查找 仓库name
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromWarehouseName(warehouseNames []string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`warehouse_name` IN (?)", warehouseNames).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 扩展
func (obj *_WmsOperationBigBagInfoMgr) GetFromExtra(extra datatypes.JSON) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 扩展
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromExtra(extras []datatypes.JSON) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsOperationBigBagInfoMgr) GetFromCreateTime(createTime time.Time) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsOperationBigBagInfoMgr) GetFromCreateUser(createUser int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsOperationBigBagInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsOperationBigBagInfoMgr) GetFromUpdateUser(updateUser int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsOperationBigBagInfoMgr) GetFromVersion(version int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromVersion(versions []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsOperationBigBagInfoMgr) GetFromDeleted(deleted int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsOperationBigBagInfoMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsOperationBigBagInfoMgr) FetchByPrimaryKey(id int) (result WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByWmsOperationBigBagInfoBigBagNoUIndex primary or index 获取唯一内容
func (obj *_WmsOperationBigBagInfoMgr) FetchUniqueByWmsOperationBigBagInfoBigBagNoUIndex(bigBagNo string) (result WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`big_bag_no` = ?", bigBagNo).First(&result).Error

	return
}

// FetchUniqueIndexByIndexBigBagNoAPIBigBagNo primary or index 获取唯一内容
func (obj *_WmsOperationBigBagInfoMgr) FetchUniqueIndexByIndexBigBagNoAPIBigBagNo(bigBagNo string, apiBigBagNo string) (result WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`big_bag_no` = ? AND `api_big_bag_no` = ?", bigBagNo, apiBigBagNo).First(&result).Error

	return
}

// FetchUniqueIndexByCustomerIDCustomerBigBagNoUnidex primary or index 获取唯一内容
func (obj *_WmsOperationBigBagInfoMgr) FetchUniqueIndexByCustomerIDCustomerBigBagNoUnidex(customerBigBagNo string, customerID int) (result WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_big_bag_no` = ? AND `customer_id` = ?", customerBigBagNo, customerID).First(&result).Error

	return
}

// FetchIndexByCustomerBigBagNoIndex  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexByCustomerBigBagNoIndex(customerBigBagNo string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_big_bag_no` = ?", customerBigBagNo).Find(&results).Error

	return
}

// FetchIndexByCustomerIDIndex  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexByCustomerIDIndex(customerID int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexBySourceIndex  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexBySourceIndex(source int8) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// FetchIndexByWmsOperationBigBagInfoTotalListNoIndex  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexByWmsOperationBigBagInfoTotalListNoIndex(totalListNo string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// FetchIndexByIndexTotalListNo  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexByIndexTotalListNo(totalListNo string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// FetchIndexByWmsOperationBigBagInfoBigBagStatusIndex  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexByWmsOperationBigBagInfoBigBagStatusIndex(bigBagStatus string) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`big_bag_status` = ?", bigBagStatus).Find(&results).Error

	return
}

// FetchIndexByIDxBigbagstatusDeletedCreatetime  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexByIDxBigbagstatusDeletedCreatetime(bigBagStatus string, createTime time.Time, deleted int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`big_bag_status` = ? AND `create_time` = ? AND `deleted` = ?", bigBagStatus, createTime, deleted).Find(&results).Error

	return
}

// FetchIndexByWmsOperationBigBagInfoFitterUserIDIndex  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexByWmsOperationBigBagInfoFitterUserIDIndex(fitterUserID int) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`fitter_user_id` = ?", fitterUserID).Find(&results).Error

	return
}

// FetchIndexByWmsOperationBigBagInfoCreateTimeIndex  获取多个内容
func (obj *_WmsOperationBigBagInfoMgr) FetchIndexByWmsOperationBigBagInfoCreateTimeIndex(createTime time.Time) (results []*WmsOperationBigBagInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
