package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsOperationBigBagInfoAffiliatedMgr struct {
	*_BaseMgr
}

// WmsOperationBigBagInfoAffiliatedMgr open func
func WmsOperationBigBagInfoAffiliatedMgr(db *gorm.DB) *_WmsOperationBigBagInfoAffiliatedMgr {
	if db == nil {
		panic(fmt.Errorf("WmsOperationBigBagInfoAffiliatedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsOperationBigBagInfoAffiliatedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_operation_big_bag_info_affiliated"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetTableName() string {
	return "wms_operation_big_bag_info_affiliated"
}

// Reset 重置gorm会话
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) Reset() *_WmsOperationBigBagInfoAffiliatedMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) Get() (result WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) Gets() (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBigBagNo big_bag_no获取 大包号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithBigBagNo(bigBagNo string) Option {
	return optionFunc(func(o *options) { o.query["big_bag_no"] = bigBagNo })
}

// WithTotalListNo total_list_no获取 总单号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithTotalListNo(totalListNo string) Option {
	return optionFunc(func(o *options) { o.query["total_list_no"] = totalListNo })
}

// WithSiteNumber site_number获取 站点编码
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithSiteNumber(siteNumber string) Option {
	return optionFunc(func(o *options) { o.query["site_number"] = siteNumber })
}

// WithFlightNo flight_no获取 航班号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithFlightNo(flightNo string) Option {
	return optionFunc(func(o *options) { o.query["flight_no"] = flightNo })
}

// WithWeight weight获取 重量
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithScannedStatus scanned_status获取 扫描状态 done:已完成，unconfirmed:待确认
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithScannedStatus(scannedStatus string) Option {
	return optionFunc(func(o *options) { o.query["scanned_status"] = scannedStatus })
}

// WithScannedUser scanned_user获取 扫描人
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithScannedUser(scannedUser int) Option {
	return optionFunc(func(o *options) { o.query["scanned_user"] = scannedUser })
}

// WithScannedTime scanned_time获取 扫描时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithScannedTime(scannedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["scanned_time"] = scannedTime })
}

// WithRemarks remarks获取 备注信息
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithIsDelete is_delete获取 逻辑删除
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// GetByOption 功能选项模式获取
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetByOption(opts ...Option) (result WmsOperationBigBagInfoAffiliated, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetByOptions(opts ...Option) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsOperationBigBagInfoAffiliated, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where(options.query)
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
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromID(id int) (result WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromID(ids []int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBigBagNo 通过big_bag_no获取内容 大包号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromBigBagNo(bigBagNo string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`big_bag_no` = ?", bigBagNo).Find(&results).Error

	return
}

// GetBatchFromBigBagNo 批量查找 大包号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromBigBagNo(bigBagNos []string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`big_bag_no` IN (?)", bigBagNos).Find(&results).Error

	return
}

// GetFromTotalListNo 通过total_list_no获取内容 总单号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromTotalListNo(totalListNo string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// GetBatchFromTotalListNo 批量查找 总单号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromTotalListNo(totalListNos []string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`total_list_no` IN (?)", totalListNos).Find(&results).Error

	return
}

// GetFromSiteNumber 通过site_number获取内容 站点编码
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromSiteNumber(siteNumber string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`site_number` = ?", siteNumber).Find(&results).Error

	return
}

// GetBatchFromSiteNumber 批量查找 站点编码
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromSiteNumber(siteNumbers []string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`site_number` IN (?)", siteNumbers).Find(&results).Error

	return
}

// GetFromFlightNo 通过flight_no获取内容 航班号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromFlightNo(flightNo string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`flight_no` = ?", flightNo).Find(&results).Error

	return
}

// GetBatchFromFlightNo 批量查找 航班号
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromFlightNo(flightNos []string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`flight_no` IN (?)", flightNos).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromWeight(weight float64) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromWeight(weights []float64) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromScannedStatus 通过scanned_status获取内容 扫描状态 done:已完成，unconfirmed:待确认
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromScannedStatus(scannedStatus string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`scanned_status` = ?", scannedStatus).Find(&results).Error

	return
}

// GetBatchFromScannedStatus 批量查找 扫描状态 done:已完成，unconfirmed:待确认
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromScannedStatus(scannedStatuss []string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`scanned_status` IN (?)", scannedStatuss).Find(&results).Error

	return
}

// GetFromScannedUser 通过scanned_user获取内容 扫描人
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromScannedUser(scannedUser int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`scanned_user` = ?", scannedUser).Find(&results).Error

	return
}

// GetBatchFromScannedUser 批量查找 扫描人
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromScannedUser(scannedUsers []int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`scanned_user` IN (?)", scannedUsers).Find(&results).Error

	return
}

// GetFromScannedTime 通过scanned_time获取内容 扫描时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromScannedTime(scannedTime time.Time) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`scanned_time` = ?", scannedTime).Find(&results).Error

	return
}

// GetBatchFromScannedTime 批量查找 扫描时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromScannedTime(scannedTimes []time.Time) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`scanned_time` IN (?)", scannedTimes).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注信息
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromRemarks(remarks string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注信息
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromRemarks(remarkss []string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromCreateUser(createUser int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromCreateTime(createTime time.Time) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromUpdateUser(updateUser int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromVersion(version int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromVersion(versions []int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 逻辑删除
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetFromIsDelete(isDelete int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 逻辑删除
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) GetBatchFromIsDelete(isDeletes []int) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) FetchByPrimaryKey(id int) (result WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByBigBagNo  获取多个内容
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) FetchIndexByBigBagNo(bigBagNo string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`big_bag_no` = ?", bigBagNo).Find(&results).Error

	return
}

// FetchIndexByTotalListNo  获取多个内容
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) FetchIndexByTotalListNo(totalListNo string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// FetchIndexBySiteNumber  获取多个内容
func (obj *_WmsOperationBigBagInfoAffiliatedMgr) FetchIndexBySiteNumber(siteNumber string) (results []*WmsOperationBigBagInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationBigBagInfoAffiliated{}).Where("`site_number` = ?", siteNumber).Find(&results).Error

	return
}
