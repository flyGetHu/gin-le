package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaPartitionAffiliateMgr struct {
	*_BaseMgr
}

// FaPartitionAffiliateMgr open func
func FaPartitionAffiliateMgr(db *gorm.DB) *_FaPartitionAffiliateMgr {
	if db == nil {
		panic(fmt.Errorf("FaPartitionAffiliateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaPartitionAffiliateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_partition_affiliate"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaPartitionAffiliateMgr) GetTableName() string {
	return "fa_partition_affiliate"
}

// Reset 重置gorm会话
func (obj *_FaPartitionAffiliateMgr) Reset() *_FaPartitionAffiliateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaPartitionAffiliateMgr) Get() (result FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaPartitionAffiliateMgr) Gets() (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaPartitionAffiliateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaPartitionAffiliateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPartitionID partition_id获取 主表分区ID
func (obj *_FaPartitionAffiliateMgr) WithPartitionID(partitionID int) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithPartitionIDentifier partition_identifier获取 分区标识:1区,2区
func (obj *_FaPartitionAffiliateMgr) WithPartitionIDentifier(partitionIDentifier string) Option {
	return optionFunc(func(o *options) { o.query["partition_identifier"] = partitionIDentifier })
}

// WithCountry country获取 国家
func (obj *_FaPartitionAffiliateMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithCountryExtra country_extra获取 国家扩展
func (obj *_FaPartitionAffiliateMgr) WithCountryExtra(countryExtra string) Option {
	return optionFunc(func(o *options) { o.query["country_extra"] = countryExtra })
}

// WithProvince province获取 省份
func (obj *_FaPartitionAffiliateMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithProvinceExtra province_extra获取 省份扩展
func (obj *_FaPartitionAffiliateMgr) WithProvinceExtra(provinceExtra string) Option {
	return optionFunc(func(o *options) { o.query["province_extra"] = provinceExtra })
}

// WithCity city获取 城市
func (obj *_FaPartitionAffiliateMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithCityExtra city_extra获取 城市扩展
func (obj *_FaPartitionAffiliateMgr) WithCityExtra(cityExtra string) Option {
	return optionFunc(func(o *options) { o.query["city_extra"] = cityExtra })
}

// WithPostalCode postal_code获取 邮编
func (obj *_FaPartitionAffiliateMgr) WithPostalCode(postalCode string) Option {
	return optionFunc(func(o *options) { o.query["postal_code"] = postalCode })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaPartitionAffiliateMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaPartitionAffiliateMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaPartitionAffiliateMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_FaPartitionAffiliateMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_FaPartitionAffiliateMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaPartitionAffiliateMgr) GetByOption(opts ...Option) (result FaPartitionAffiliate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaPartitionAffiliateMgr) GetByOptions(opts ...Option) (results []*FaPartitionAffiliate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaPartitionAffiliateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaPartitionAffiliate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where(options.query)
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
func (obj *_FaPartitionAffiliateMgr) GetFromID(id int) (result FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaPartitionAffiliateMgr) GetBatchFromID(ids []int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容 主表分区ID
func (obj *_FaPartitionAffiliateMgr) GetFromPartitionID(partitionID int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找 主表分区ID
func (obj *_FaPartitionAffiliateMgr) GetBatchFromPartitionID(partitionIDs []int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromPartitionIDentifier 通过partition_identifier获取内容 分区标识:1区,2区
func (obj *_FaPartitionAffiliateMgr) GetFromPartitionIDentifier(partitionIDentifier string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`partition_identifier` = ?", partitionIDentifier).Find(&results).Error

	return
}

// GetBatchFromPartitionIDentifier 批量查找 分区标识:1区,2区
func (obj *_FaPartitionAffiliateMgr) GetBatchFromPartitionIDentifier(partitionIDentifiers []string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`partition_identifier` IN (?)", partitionIDentifiers).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 国家
func (obj *_FaPartitionAffiliateMgr) GetFromCountry(country string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`country` = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量查找 国家
func (obj *_FaPartitionAffiliateMgr) GetBatchFromCountry(countrys []string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`country` IN (?)", countrys).Find(&results).Error

	return
}

// GetFromCountryExtra 通过country_extra获取内容 国家扩展
func (obj *_FaPartitionAffiliateMgr) GetFromCountryExtra(countryExtra string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`country_extra` = ?", countryExtra).Find(&results).Error

	return
}

// GetBatchFromCountryExtra 批量查找 国家扩展
func (obj *_FaPartitionAffiliateMgr) GetBatchFromCountryExtra(countryExtras []string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`country_extra` IN (?)", countryExtras).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 省份
func (obj *_FaPartitionAffiliateMgr) GetFromProvince(province string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 省份
func (obj *_FaPartitionAffiliateMgr) GetBatchFromProvince(provinces []string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromProvinceExtra 通过province_extra获取内容 省份扩展
func (obj *_FaPartitionAffiliateMgr) GetFromProvinceExtra(provinceExtra string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`province_extra` = ?", provinceExtra).Find(&results).Error

	return
}

// GetBatchFromProvinceExtra 批量查找 省份扩展
func (obj *_FaPartitionAffiliateMgr) GetBatchFromProvinceExtra(provinceExtras []string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`province_extra` IN (?)", provinceExtras).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 城市
func (obj *_FaPartitionAffiliateMgr) GetFromCity(city string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 城市
func (obj *_FaPartitionAffiliateMgr) GetBatchFromCity(citys []string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromCityExtra 通过city_extra获取内容 城市扩展
func (obj *_FaPartitionAffiliateMgr) GetFromCityExtra(cityExtra string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`city_extra` = ?", cityExtra).Find(&results).Error

	return
}

// GetBatchFromCityExtra 批量查找 城市扩展
func (obj *_FaPartitionAffiliateMgr) GetBatchFromCityExtra(cityExtras []string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`city_extra` IN (?)", cityExtras).Find(&results).Error

	return
}

// GetFromPostalCode 通过postal_code获取内容 邮编
func (obj *_FaPartitionAffiliateMgr) GetFromPostalCode(postalCode string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`postal_code` = ?", postalCode).Find(&results).Error

	return
}

// GetBatchFromPostalCode 批量查找 邮编
func (obj *_FaPartitionAffiliateMgr) GetBatchFromPostalCode(postalCodes []string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`postal_code` IN (?)", postalCodes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaPartitionAffiliateMgr) GetFromCreateTime(createTime time.Time) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaPartitionAffiliateMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaPartitionAffiliateMgr) GetFromCreateUser(createUser int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaPartitionAffiliateMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaPartitionAffiliateMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaPartitionAffiliateMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_FaPartitionAffiliateMgr) GetFromUpdateUser(updateUser int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_FaPartitionAffiliateMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaPartitionAffiliateMgr) GetFromVersion(version int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaPartitionAffiliateMgr) GetBatchFromVersion(versions []int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaPartitionAffiliateMgr) FetchByPrimaryKey(id int) (result FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByFaPartitionAffiliatePartitionIDIndex  获取多个内容
func (obj *_FaPartitionAffiliateMgr) FetchIndexByFaPartitionAffiliatePartitionIDIndex(partitionID int) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// FetchIndexByFaPartitionAffiliatePartitionIDentifierIndex  获取多个内容
func (obj *_FaPartitionAffiliateMgr) FetchIndexByFaPartitionAffiliatePartitionIDentifierIndex(partitionIDentifier string) (results []*FaPartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartitionAffiliate{}).Where("`partition_identifier` = ?", partitionIDentifier).Find(&results).Error

	return
}
