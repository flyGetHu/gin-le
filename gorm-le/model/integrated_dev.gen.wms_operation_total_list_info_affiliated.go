package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsOperationTotalListInfoAffiliatedMgr struct {
	*_BaseMgr
}

// WmsOperationTotalListInfoAffiliatedMgr open func
func WmsOperationTotalListInfoAffiliatedMgr(db *gorm.DB) *_WmsOperationTotalListInfoAffiliatedMgr {
	if db == nil {
		panic(fmt.Errorf("WmsOperationTotalListInfoAffiliatedMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsOperationTotalListInfoAffiliatedMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_operation_total_list_info_affiliated"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetTableName() string {
	return "wms_operation_total_list_info_affiliated"
}

// Reset 重置gorm会话
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) Reset() *_WmsOperationTotalListInfoAffiliatedMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) Get() (result WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) Gets() (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTotalListNo total_list_no获取 总单号
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithTotalListNo(totalListNo string) Option {
	return optionFunc(func(o *options) { o.query["total_list_no"] = totalListNo })
}

// WithSiteNumber site_number获取 站点编码
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithSiteNumber(siteNumber string) Option {
	return optionFunc(func(o *options) { o.query["site_number"] = siteNumber })
}

// WithCountryCode country_code获取 国家二字码
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithCountryCode(countryCode string) Option {
	return optionFunc(func(o *options) { o.query["country_code"] = countryCode })
}

// WithScannedCount scanned_count获取 已扫袋数
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithScannedCount(scannedCount int) Option {
	return optionFunc(func(o *options) { o.query["scanned_count"] = scannedCount })
}

// WithScannedWeight scanned_weight获取 已扫重量
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithScannedWeight(scannedWeight float64) Option {
	return optionFunc(func(o *options) { o.query["scanned_weight"] = scannedWeight })
}

// WithScannedUser scanned_user获取 扫描人
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithScannedUser(scannedUser int) Option {
	return optionFunc(func(o *options) { o.query["scanned_user"] = scannedUser })
}

// WithScannedTime scanned_time获取 扫描时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithScannedTime(scannedTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["scanned_time"] = scannedTime })
}

// WithRemarks remarks获取 备注信息
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithIsDelete is_delete获取 逻辑删除
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// GetByOption 功能选项模式获取
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetByOption(opts ...Option) (result WmsOperationTotalListInfoAffiliated, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetByOptions(opts ...Option) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsOperationTotalListInfoAffiliated, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where(options.query)
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
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromID(id int) (result WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromID(ids []int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTotalListNo 通过total_list_no获取内容 总单号
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromTotalListNo(totalListNo string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// GetBatchFromTotalListNo 批量查找 总单号
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromTotalListNo(totalListNos []string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`total_list_no` IN (?)", totalListNos).Find(&results).Error

	return
}

// GetFromSiteNumber 通过site_number获取内容 站点编码
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromSiteNumber(siteNumber string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`site_number` = ?", siteNumber).Find(&results).Error

	return
}

// GetBatchFromSiteNumber 批量查找 站点编码
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromSiteNumber(siteNumbers []string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`site_number` IN (?)", siteNumbers).Find(&results).Error

	return
}

// GetFromCountryCode 通过country_code获取内容 国家二字码
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromCountryCode(countryCode string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`country_code` = ?", countryCode).Find(&results).Error

	return
}

// GetBatchFromCountryCode 批量查找 国家二字码
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromCountryCode(countryCodes []string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`country_code` IN (?)", countryCodes).Find(&results).Error

	return
}

// GetFromScannedCount 通过scanned_count获取内容 已扫袋数
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromScannedCount(scannedCount int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`scanned_count` = ?", scannedCount).Find(&results).Error

	return
}

// GetBatchFromScannedCount 批量查找 已扫袋数
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromScannedCount(scannedCounts []int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`scanned_count` IN (?)", scannedCounts).Find(&results).Error

	return
}

// GetFromScannedWeight 通过scanned_weight获取内容 已扫重量
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromScannedWeight(scannedWeight float64) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`scanned_weight` = ?", scannedWeight).Find(&results).Error

	return
}

// GetBatchFromScannedWeight 批量查找 已扫重量
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromScannedWeight(scannedWeights []float64) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`scanned_weight` IN (?)", scannedWeights).Find(&results).Error

	return
}

// GetFromScannedUser 通过scanned_user获取内容 扫描人
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromScannedUser(scannedUser int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`scanned_user` = ?", scannedUser).Find(&results).Error

	return
}

// GetBatchFromScannedUser 批量查找 扫描人
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromScannedUser(scannedUsers []int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`scanned_user` IN (?)", scannedUsers).Find(&results).Error

	return
}

// GetFromScannedTime 通过scanned_time获取内容 扫描时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromScannedTime(scannedTime time.Time) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`scanned_time` = ?", scannedTime).Find(&results).Error

	return
}

// GetBatchFromScannedTime 批量查找 扫描时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromScannedTime(scannedTimes []time.Time) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`scanned_time` IN (?)", scannedTimes).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注信息
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromRemarks(remarks string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注信息
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromRemarks(remarkss []string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromCreateUser(createUser int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromCreateTime(createTime time.Time) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromUpdateUser(updateUser int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromVersion(version int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromVersion(versions []int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 逻辑删除
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetFromIsDelete(isDelete int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 逻辑删除
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) GetBatchFromIsDelete(isDeletes []int) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) FetchByPrimaryKey(id int) (result WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByTotalListNo  获取多个内容
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) FetchIndexByTotalListNo(totalListNo string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// FetchIndexBySiteNumber  获取多个内容
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) FetchIndexBySiteNumber(siteNumber string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`site_number` = ?", siteNumber).Find(&results).Error

	return
}

// FetchIndexByCountryCode  获取多个内容
func (obj *_WmsOperationTotalListInfoAffiliatedMgr) FetchIndexByCountryCode(countryCode string) (results []*WmsOperationTotalListInfoAffiliated, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsOperationTotalListInfoAffiliated{}).Where("`country_code` = ?", countryCode).Find(&results).Error

	return
}
