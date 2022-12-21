package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _USiteMgr struct {
	*_BaseMgr
}

// USiteMgr open func
func USiteMgr(db *gorm.DB) *_USiteMgr {
	if db == nil {
		panic(fmt.Errorf("USiteMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_USiteMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_site"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_USiteMgr) GetTableName() string {
	return "u_site"
}

// Reset 重置gorm会话
func (obj *_USiteMgr) Reset() *_USiteMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_USiteMgr) Get() (result USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_USiteMgr) Gets() (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_USiteMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(USite{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_USiteMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSiteNumber site_number获取 站点编码
func (obj *_USiteMgr) WithSiteNumber(siteNumber string) Option {
	return optionFunc(func(o *options) { o.query["site_number"] = siteNumber })
}

// WithSiteName site_name获取 站点名称
func (obj *_USiteMgr) WithSiteName(siteName string) Option {
	return optionFunc(func(o *options) { o.query["site_name"] = siteName })
}

// WithCountryCode country_code获取 国家二字码
func (obj *_USiteMgr) WithCountryCode(countryCode string) Option {
	return optionFunc(func(o *options) { o.query["country_code"] = countryCode })
}

// WithTrackAddress track_address获取 轨迹地址
func (obj *_USiteMgr) WithTrackAddress(trackAddress string) Option {
	return optionFunc(func(o *options) { o.query["track_address"] = trackAddress })
}

// WithRemarks remarks获取 备注信息
func (obj *_USiteMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateUser create_user获取 创建用户
func (obj *_USiteMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_USiteMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_USiteMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_USiteMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_USiteMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithIsDelete is_delete获取 逻辑删除
func (obj *_USiteMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// GetByOption 功能选项模式获取
func (obj *_USiteMgr) GetByOption(opts ...Option) (result USite, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_USiteMgr) GetByOptions(opts ...Option) (results []*USite, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_USiteMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]USite, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(USite{}).Where(options.query)
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
func (obj *_USiteMgr) GetFromID(id int) (result USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_USiteMgr) GetBatchFromID(ids []int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSiteNumber 通过site_number获取内容 站点编码
func (obj *_USiteMgr) GetFromSiteNumber(siteNumber string) (result USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`site_number` = ?", siteNumber).First(&result).Error

	return
}

// GetBatchFromSiteNumber 批量查找 站点编码
func (obj *_USiteMgr) GetBatchFromSiteNumber(siteNumbers []string) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`site_number` IN (?)", siteNumbers).Find(&results).Error

	return
}

// GetFromSiteName 通过site_name获取内容 站点名称
func (obj *_USiteMgr) GetFromSiteName(siteName string) (result USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`site_name` = ?", siteName).First(&result).Error

	return
}

// GetBatchFromSiteName 批量查找 站点名称
func (obj *_USiteMgr) GetBatchFromSiteName(siteNames []string) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`site_name` IN (?)", siteNames).Find(&results).Error

	return
}

// GetFromCountryCode 通过country_code获取内容 国家二字码
func (obj *_USiteMgr) GetFromCountryCode(countryCode string) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`country_code` = ?", countryCode).Find(&results).Error

	return
}

// GetBatchFromCountryCode 批量查找 国家二字码
func (obj *_USiteMgr) GetBatchFromCountryCode(countryCodes []string) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`country_code` IN (?)", countryCodes).Find(&results).Error

	return
}

// GetFromTrackAddress 通过track_address获取内容 轨迹地址
func (obj *_USiteMgr) GetFromTrackAddress(trackAddress string) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`track_address` = ?", trackAddress).Find(&results).Error

	return
}

// GetBatchFromTrackAddress 批量查找 轨迹地址
func (obj *_USiteMgr) GetBatchFromTrackAddress(trackAddresss []string) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`track_address` IN (?)", trackAddresss).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注信息
func (obj *_USiteMgr) GetFromRemarks(remarks string) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注信息
func (obj *_USiteMgr) GetBatchFromRemarks(remarkss []string) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_USiteMgr) GetFromCreateUser(createUser int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_USiteMgr) GetBatchFromCreateUser(createUsers []int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_USiteMgr) GetFromCreateTime(createTime time.Time) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_USiteMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_USiteMgr) GetFromUpdateUser(updateUser int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_USiteMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_USiteMgr) GetFromUpdateTime(updateTime time.Time) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_USiteMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_USiteMgr) GetFromVersion(version int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_USiteMgr) GetBatchFromVersion(versions []int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 逻辑删除
func (obj *_USiteMgr) GetFromIsDelete(isDelete int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 逻辑删除
func (obj *_USiteMgr) GetBatchFromIsDelete(isDeletes []int) (results []*USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_USiteMgr) FetchByPrimaryKey(id int) (result USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueBySiteNumber primary or index 获取唯一内容
func (obj *_USiteMgr) FetchUniqueBySiteNumber(siteNumber string) (result USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`site_number` = ?", siteNumber).First(&result).Error

	return
}

// FetchUniqueBySiteName primary or index 获取唯一内容
func (obj *_USiteMgr) FetchUniqueBySiteName(siteName string) (result USite, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(USite{}).Where("`site_name` = ?", siteName).First(&result).Error

	return
}
