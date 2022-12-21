package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsSorterBigBagMgr struct {
	*_BaseMgr
}

// WmsSorterBigBagMgr open func
func WmsSorterBigBagMgr(db *gorm.DB) *_WmsSorterBigBagMgr {
	if db == nil {
		panic(fmt.Errorf("WmsSorterBigBagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsSorterBigBagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_sorter_big_bag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsSorterBigBagMgr) GetTableName() string {
	return "wms_sorter_big_bag"
}

// Reset 重置gorm会话
func (obj *_WmsSorterBigBagMgr) Reset() *_WmsSorterBigBagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsSorterBigBagMgr) Get() (result WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsSorterBigBagMgr) Gets() (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsSorterBigBagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsSorterBigBagMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGridNum grid_num获取 格口
func (obj *_WmsSorterBigBagMgr) WithGridNum(gridNum int) Option {
	return optionFunc(func(o *options) { o.query["grid_num"] = gridNum })
}

// WithDownstreamCode downstream_code获取 下家code
func (obj *_WmsSorterBigBagMgr) WithDownstreamCode(downstreamCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_code"] = downstreamCode })
}

// WithSmallBagCount small_bag_count获取 小包数量
func (obj *_WmsSorterBigBagMgr) WithSmallBagCount(smallBagCount int) Option {
	return optionFunc(func(o *options) { o.query["small_bag_count"] = smallBagCount })
}

// WithWeight weight获取 重量
func (obj *_WmsSorterBigBagMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithBigBagNo big_bag_no获取 大包号
func (obj *_WmsSorterBigBagMgr) WithBigBagNo(bigBagNo string) Option {
	return optionFunc(func(o *options) { o.query["big_bag_no"] = bigBagNo })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsSorterBigBagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsSorterBigBagMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsSorterBigBagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsSorterBigBagMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsSorterBigBagMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsSorterBigBagMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsSorterBigBagMgr) GetByOption(opts ...Option) (result WmsSorterBigBag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsSorterBigBagMgr) GetByOptions(opts ...Option) (results []*WmsSorterBigBag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsSorterBigBagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsSorterBigBag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where(options.query)
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
func (obj *_WmsSorterBigBagMgr) GetFromID(id int) (result WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsSorterBigBagMgr) GetBatchFromID(ids []int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGridNum 通过grid_num获取内容 格口
func (obj *_WmsSorterBigBagMgr) GetFromGridNum(gridNum int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`grid_num` = ?", gridNum).Find(&results).Error

	return
}

// GetBatchFromGridNum 批量查找 格口
func (obj *_WmsSorterBigBagMgr) GetBatchFromGridNum(gridNums []int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`grid_num` IN (?)", gridNums).Find(&results).Error

	return
}

// GetFromDownstreamCode 通过downstream_code获取内容 下家code
func (obj *_WmsSorterBigBagMgr) GetFromDownstreamCode(downstreamCode string) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`downstream_code` = ?", downstreamCode).Find(&results).Error

	return
}

// GetBatchFromDownstreamCode 批量查找 下家code
func (obj *_WmsSorterBigBagMgr) GetBatchFromDownstreamCode(downstreamCodes []string) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`downstream_code` IN (?)", downstreamCodes).Find(&results).Error

	return
}

// GetFromSmallBagCount 通过small_bag_count获取内容 小包数量
func (obj *_WmsSorterBigBagMgr) GetFromSmallBagCount(smallBagCount int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`small_bag_count` = ?", smallBagCount).Find(&results).Error

	return
}

// GetBatchFromSmallBagCount 批量查找 小包数量
func (obj *_WmsSorterBigBagMgr) GetBatchFromSmallBagCount(smallBagCounts []int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`small_bag_count` IN (?)", smallBagCounts).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *_WmsSorterBigBagMgr) GetFromWeight(weight float64) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *_WmsSorterBigBagMgr) GetBatchFromWeight(weights []float64) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromBigBagNo 通过big_bag_no获取内容 大包号
func (obj *_WmsSorterBigBagMgr) GetFromBigBagNo(bigBagNo string) (result WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`big_bag_no` = ?", bigBagNo).First(&result).Error

	return
}

// GetBatchFromBigBagNo 批量查找 大包号
func (obj *_WmsSorterBigBagMgr) GetBatchFromBigBagNo(bigBagNos []string) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`big_bag_no` IN (?)", bigBagNos).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsSorterBigBagMgr) GetFromCreateTime(createTime time.Time) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsSorterBigBagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsSorterBigBagMgr) GetFromCreateUser(createUser int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsSorterBigBagMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsSorterBigBagMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsSorterBigBagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsSorterBigBagMgr) GetFromUpdateUser(updateUser int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsSorterBigBagMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsSorterBigBagMgr) GetFromVersion(version int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsSorterBigBagMgr) GetBatchFromVersion(versions []int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsSorterBigBagMgr) GetFromDeleted(deleted int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsSorterBigBagMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsSorterBigBagMgr) FetchByPrimaryKey(id int) (result WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByBigBagNo primary or index 获取唯一内容
func (obj *_WmsSorterBigBagMgr) FetchUniqueByBigBagNo(bigBagNo string) (result WmsSorterBigBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterBigBag{}).Where("`big_bag_no` = ?", bigBagNo).First(&result).Error

	return
}
