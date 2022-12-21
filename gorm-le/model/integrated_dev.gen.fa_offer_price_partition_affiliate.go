package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaOfferPricePartitionAffiliateMgr struct {
	*_BaseMgr
}

// FaOfferPricePartitionAffiliateMgr open func
func FaOfferPricePartitionAffiliateMgr(db *gorm.DB) *_FaOfferPricePartitionAffiliateMgr {
	if db == nil {
		panic(fmt.Errorf("FaOfferPricePartitionAffiliateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaOfferPricePartitionAffiliateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_offer_price_partition_affiliate"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaOfferPricePartitionAffiliateMgr) GetTableName() string {
	return "fa_offer_price_partition_affiliate"
}

// Reset 重置gorm会话
func (obj *_FaOfferPricePartitionAffiliateMgr) Reset() *_FaOfferPricePartitionAffiliateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaOfferPricePartitionAffiliateMgr) Get() (result FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaOfferPricePartitionAffiliateMgr) Gets() (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaOfferPricePartitionAffiliateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaOfferPricePartitionAffiliateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOfferPriceID offer_price_id获取 关联报价表ID
func (obj *_FaOfferPricePartitionAffiliateMgr) WithOfferPriceID(offerPriceID int) Option {
	return optionFunc(func(o *options) { o.query["offer_price_id"] = offerPriceID })
}

// WithPartitionID partition_id获取 分区id
func (obj *_FaOfferPricePartitionAffiliateMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithPartitionName partition_name获取 分区名称
func (obj *_FaOfferPricePartitionAffiliateMgr) WithPartitionName(partitionName string) Option {
	return optionFunc(func(o *options) { o.query["partition_name"] = partitionName })
}

// WithPartitionIDentifier partition_identifier获取 分区标识:1区,2区
func (obj *_FaOfferPricePartitionAffiliateMgr) WithPartitionIDentifier(partitionIDentifier string) Option {
	return optionFunc(func(o *options) { o.query["partition_identifier"] = partitionIDentifier })
}

// WithCountry country获取 国家
func (obj *_FaOfferPricePartitionAffiliateMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithProvince province获取 省份
func (obj *_FaOfferPricePartitionAffiliateMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithProvinceExtra province_extra获取 省份扩展
func (obj *_FaOfferPricePartitionAffiliateMgr) WithProvinceExtra(provinceExtra string) Option {
	return optionFunc(func(o *options) { o.query["province_extra"] = provinceExtra })
}

// WithCity city获取 城市
func (obj *_FaOfferPricePartitionAffiliateMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithCityExtra city_extra获取 城市扩展
func (obj *_FaOfferPricePartitionAffiliateMgr) WithCityExtra(cityExtra string) Option {
	return optionFunc(func(o *options) { o.query["city_extra"] = cityExtra })
}

// WithPostalCode postal_code获取 邮编
func (obj *_FaOfferPricePartitionAffiliateMgr) WithPostalCode(postalCode string) Option {
	return optionFunc(func(o *options) { o.query["postal_code"] = postalCode })
}

// WithRemarks remarks获取 备注
func (obj *_FaOfferPricePartitionAffiliateMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaOfferPricePartitionAffiliateMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaOfferPricePartitionAffiliateMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaOfferPricePartitionAffiliateMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_FaOfferPricePartitionAffiliateMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_FaOfferPricePartitionAffiliateMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaOfferPricePartitionAffiliateMgr) GetByOption(opts ...Option) (result FaOfferPricePartitionAffiliate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaOfferPricePartitionAffiliateMgr) GetByOptions(opts ...Option) (results []*FaOfferPricePartitionAffiliate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaOfferPricePartitionAffiliateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaOfferPricePartitionAffiliate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where(options.query)
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
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromID(id int) (result FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromID(ids []int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOfferPriceID 通过offer_price_id获取内容 关联报价表ID
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromOfferPriceID(offerPriceID int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`offer_price_id` = ?", offerPriceID).Find(&results).Error

	return
}

// GetBatchFromOfferPriceID 批量查找 关联报价表ID
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromOfferPriceID(offerPriceIDs []int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`offer_price_id` IN (?)", offerPriceIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容 分区id
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromPartitionID(partitionID int64) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找 分区id
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromPartitionName 通过partition_name获取内容 分区名称
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromPartitionName(partitionName string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`partition_name` = ?", partitionName).Find(&results).Error

	return
}

// GetBatchFromPartitionName 批量查找 分区名称
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromPartitionName(partitionNames []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`partition_name` IN (?)", partitionNames).Find(&results).Error

	return
}

// GetFromPartitionIDentifier 通过partition_identifier获取内容 分区标识:1区,2区
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromPartitionIDentifier(partitionIDentifier string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`partition_identifier` = ?", partitionIDentifier).Find(&results).Error

	return
}

// GetBatchFromPartitionIDentifier 批量查找 分区标识:1区,2区
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromPartitionIDentifier(partitionIDentifiers []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`partition_identifier` IN (?)", partitionIDentifiers).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 国家
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromCountry(country string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`country` = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量查找 国家
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromCountry(countrys []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`country` IN (?)", countrys).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 省份
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromProvince(province string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 省份
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromProvince(provinces []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromProvinceExtra 通过province_extra获取内容 省份扩展
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromProvinceExtra(provinceExtra string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`province_extra` = ?", provinceExtra).Find(&results).Error

	return
}

// GetBatchFromProvinceExtra 批量查找 省份扩展
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromProvinceExtra(provinceExtras []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`province_extra` IN (?)", provinceExtras).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 城市
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromCity(city string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 城市
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromCity(citys []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromCityExtra 通过city_extra获取内容 城市扩展
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromCityExtra(cityExtra string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`city_extra` = ?", cityExtra).Find(&results).Error

	return
}

// GetBatchFromCityExtra 批量查找 城市扩展
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromCityExtra(cityExtras []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`city_extra` IN (?)", cityExtras).Find(&results).Error

	return
}

// GetFromPostalCode 通过postal_code获取内容 邮编
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromPostalCode(postalCode string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`postal_code` = ?", postalCode).Find(&results).Error

	return
}

// GetBatchFromPostalCode 批量查找 邮编
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromPostalCode(postalCodes []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`postal_code` IN (?)", postalCodes).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromRemarks(remarks string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromRemarks(remarkss []string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromCreateTime(createTime time.Time) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromCreateUser(createUser int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromUpdateUser(updateUser int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaOfferPricePartitionAffiliateMgr) GetFromVersion(version int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaOfferPricePartitionAffiliateMgr) GetBatchFromVersion(versions []int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaOfferPricePartitionAffiliateMgr) FetchByPrimaryKey(id int) (result FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByFaOfferPricePartitionOfferPriceIDIndex  获取多个内容
func (obj *_FaOfferPricePartitionAffiliateMgr) FetchIndexByFaOfferPricePartitionOfferPriceIDIndex(offerPriceID int) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`offer_price_id` = ?", offerPriceID).Find(&results).Error

	return
}

// FetchIndexByFaOfferPricePartitionPartitionIDentifierIndex  获取多个内容
func (obj *_FaOfferPricePartitionAffiliateMgr) FetchIndexByFaOfferPricePartitionPartitionIDentifierIndex(partitionIDentifier string) (results []*FaOfferPricePartitionAffiliate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPricePartitionAffiliate{}).Where("`partition_identifier` = ?", partitionIDentifier).Find(&results).Error

	return
}
