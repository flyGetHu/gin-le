package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsAirBillMgr struct {
	*_BaseMgr
}

// WmsAirBillMgr open func
func WmsAirBillMgr(db *gorm.DB) *_WmsAirBillMgr {
	if db == nil {
		panic(fmt.Errorf("WmsAirBillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsAirBillMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_air_bill"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsAirBillMgr) GetTableName() string {
	return "wms_air_bill"
}

// Reset 重置gorm会话
func (obj *_WmsAirBillMgr) Reset() *_WmsAirBillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsAirBillMgr) Get() (result WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsAirBillMgr) Gets() (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsAirBillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsAirBillMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAirBillNumber air_bill_number获取 航空提单号
func (obj *_WmsAirBillMgr) WithAirBillNumber(airBillNumber string) Option {
	return optionFunc(func(o *options) { o.query["air_bill_number"] = airBillNumber })
}

// WithBigBagCount big_bag_count获取 大包数量
func (obj *_WmsAirBillMgr) WithBigBagCount(bigBagCount int) Option {
	return optionFunc(func(o *options) { o.query["big_bag_count"] = bigBagCount })
}

// WithBigBagWeight big_bag_weight获取 大包重量
func (obj *_WmsAirBillMgr) WithBigBagWeight(bigBagWeight float64) Option {
	return optionFunc(func(o *options) { o.query["big_bag_weight"] = bigBagWeight })
}

// WithSmallBagCount small_bag_count获取 小包数量
func (obj *_WmsAirBillMgr) WithSmallBagCount(smallBagCount int) Option {
	return optionFunc(func(o *options) { o.query["small_bag_count"] = smallBagCount })
}

// WithDownstreamCodes downstream_codes获取 下家三字码，可能有多个
func (obj *_WmsAirBillMgr) WithDownstreamCodes(downstreamCodes string) Option {
	return optionFunc(func(o *options) { o.query["downstream_codes"] = downstreamCodes })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsAirBillMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsAirBillMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsAirBillMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsAirBillMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsAirBillMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsAirBillMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsAirBillMgr) GetByOption(opts ...Option) (result WmsAirBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsAirBillMgr) GetByOptions(opts ...Option) (results []*WmsAirBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsAirBillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsAirBill, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where(options.query)
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
func (obj *_WmsAirBillMgr) GetFromID(id int) (result WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsAirBillMgr) GetBatchFromID(ids []int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAirBillNumber 通过air_bill_number获取内容 航空提单号
func (obj *_WmsAirBillMgr) GetFromAirBillNumber(airBillNumber string) (result WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`air_bill_number` = ?", airBillNumber).First(&result).Error

	return
}

// GetBatchFromAirBillNumber 批量查找 航空提单号
func (obj *_WmsAirBillMgr) GetBatchFromAirBillNumber(airBillNumbers []string) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`air_bill_number` IN (?)", airBillNumbers).Find(&results).Error

	return
}

// GetFromBigBagCount 通过big_bag_count获取内容 大包数量
func (obj *_WmsAirBillMgr) GetFromBigBagCount(bigBagCount int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`big_bag_count` = ?", bigBagCount).Find(&results).Error

	return
}

// GetBatchFromBigBagCount 批量查找 大包数量
func (obj *_WmsAirBillMgr) GetBatchFromBigBagCount(bigBagCounts []int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`big_bag_count` IN (?)", bigBagCounts).Find(&results).Error

	return
}

// GetFromBigBagWeight 通过big_bag_weight获取内容 大包重量
func (obj *_WmsAirBillMgr) GetFromBigBagWeight(bigBagWeight float64) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`big_bag_weight` = ?", bigBagWeight).Find(&results).Error

	return
}

// GetBatchFromBigBagWeight 批量查找 大包重量
func (obj *_WmsAirBillMgr) GetBatchFromBigBagWeight(bigBagWeights []float64) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`big_bag_weight` IN (?)", bigBagWeights).Find(&results).Error

	return
}

// GetFromSmallBagCount 通过small_bag_count获取内容 小包数量
func (obj *_WmsAirBillMgr) GetFromSmallBagCount(smallBagCount int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`small_bag_count` = ?", smallBagCount).Find(&results).Error

	return
}

// GetBatchFromSmallBagCount 批量查找 小包数量
func (obj *_WmsAirBillMgr) GetBatchFromSmallBagCount(smallBagCounts []int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`small_bag_count` IN (?)", smallBagCounts).Find(&results).Error

	return
}

// GetFromDownstreamCodes 通过downstream_codes获取内容 下家三字码，可能有多个
func (obj *_WmsAirBillMgr) GetFromDownstreamCodes(downstreamCodes string) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`downstream_codes` = ?", downstreamCodes).Find(&results).Error

	return
}

// GetBatchFromDownstreamCodes 批量查找 下家三字码，可能有多个
func (obj *_WmsAirBillMgr) GetBatchFromDownstreamCodes(downstreamCodess []string) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`downstream_codes` IN (?)", downstreamCodess).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsAirBillMgr) GetFromCreateTime(createTime time.Time) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsAirBillMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsAirBillMgr) GetFromCreateUser(createUser int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsAirBillMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsAirBillMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsAirBillMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsAirBillMgr) GetFromUpdateUser(updateUser int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsAirBillMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsAirBillMgr) GetFromVersion(version int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsAirBillMgr) GetBatchFromVersion(versions []int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsAirBillMgr) GetFromDeleted(deleted int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsAirBillMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsAirBillMgr) FetchByPrimaryKey(id int) (result WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexAirBillNumber primary or index 获取唯一内容
func (obj *_WmsAirBillMgr) FetchUniqueByIndexAirBillNumber(airBillNumber string) (result WmsAirBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsAirBill{}).Where("`air_bill_number` = ?", airBillNumber).First(&result).Error

	return
}
