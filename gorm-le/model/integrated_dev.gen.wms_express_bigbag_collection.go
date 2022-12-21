package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsExpressBigbagCollectionMgr struct {
	*_BaseMgr
}

// WmsExpressBigbagCollectionMgr open func
func WmsExpressBigbagCollectionMgr(db *gorm.DB) *_WmsExpressBigbagCollectionMgr {
	if db == nil {
		panic(fmt.Errorf("WmsExpressBigbagCollectionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsExpressBigbagCollectionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_express_bigbag_collection"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsExpressBigbagCollectionMgr) GetTableName() string {
	return "wms_express_bigbag_collection"
}

// Reset 重置gorm会话
func (obj *_WmsExpressBigbagCollectionMgr) Reset() *_WmsExpressBigbagCollectionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsExpressBigbagCollectionMgr) Get() (result WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsExpressBigbagCollectionMgr) Gets() (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsExpressBigbagCollectionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_WmsExpressBigbagCollectionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBigBagNum bigBag_num获取 大包号
func (obj *_WmsExpressBigbagCollectionMgr) WithBigBagNum(bigBagNum string) Option {
	return optionFunc(func(o *options) { o.query["bigBag_num"] = bigBagNum })
}

// WithBigBagWeight bigBag_weight获取 大包重量
func (obj *_WmsExpressBigbagCollectionMgr) WithBigBagWeight(bigBagWeight float64) Option {
	return optionFunc(func(o *options) { o.query["bigBag_weight"] = bigBagWeight })
}

// WithBigBagQuantity bigBag_quantity获取 包裹数量
func (obj *_WmsExpressBigbagCollectionMgr) WithBigBagQuantity(bigBagQuantity int) Option {
	return optionFunc(func(o *options) { o.query["bigBag_quantity"] = bigBagQuantity })
}

// WithCreateUser create_user获取 创建人
func (obj *_WmsExpressBigbagCollectionMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsExpressBigbagCollectionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_WmsExpressBigbagCollectionMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsExpressBigbagCollectionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_WmsExpressBigbagCollectionMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsExpressBigbagCollectionMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsExpressBigbagCollectionMgr) GetByOption(opts ...Option) (result WmsExpressBigbagCollection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsExpressBigbagCollectionMgr) GetByOptions(opts ...Option) (results []*WmsExpressBigbagCollection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsExpressBigbagCollectionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsExpressBigbagCollection, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *_WmsExpressBigbagCollectionMgr) GetFromID(id int) (result WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromID(ids []int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromBigBagNum 通过bigBag_num获取内容 大包号
func (obj *_WmsExpressBigbagCollectionMgr) GetFromBigBagNum(bigBagNum string) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`bigBag_num` = ?", bigBagNum).Find(&results).Error

	return
}

// GetBatchFromBigBagNum 批量查找 大包号
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromBigBagNum(bigBagNums []string) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`bigBag_num` IN (?)", bigBagNums).Find(&results).Error

	return
}

// GetFromBigBagWeight 通过bigBag_weight获取内容 大包重量
func (obj *_WmsExpressBigbagCollectionMgr) GetFromBigBagWeight(bigBagWeight float64) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`bigBag_weight` = ?", bigBagWeight).Find(&results).Error

	return
}

// GetBatchFromBigBagWeight 批量查找 大包重量
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromBigBagWeight(bigBagWeights []float64) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`bigBag_weight` IN (?)", bigBagWeights).Find(&results).Error

	return
}

// GetFromBigBagQuantity 通过bigBag_quantity获取内容 包裹数量
func (obj *_WmsExpressBigbagCollectionMgr) GetFromBigBagQuantity(bigBagQuantity int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`bigBag_quantity` = ?", bigBagQuantity).Find(&results).Error

	return
}

// GetBatchFromBigBagQuantity 批量查找 包裹数量
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromBigBagQuantity(bigBagQuantitys []int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`bigBag_quantity` IN (?)", bigBagQuantitys).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_WmsExpressBigbagCollectionMgr) GetFromCreateUser(createUser int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsExpressBigbagCollectionMgr) GetFromCreateTime(createTime time.Time) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_WmsExpressBigbagCollectionMgr) GetFromUpdateUser(updateUser int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsExpressBigbagCollectionMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsExpressBigbagCollectionMgr) GetFromVersion(version int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromVersion(versions []int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsExpressBigbagCollectionMgr) GetFromDeleted(deleted int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsExpressBigbagCollectionMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsExpressBigbagCollectionMgr) FetchByPrimaryKey(id int) (result WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByWmsExpressBigBagCollectionBigBagNumIndex  获取多个内容
func (obj *_WmsExpressBigbagCollectionMgr) FetchIndexByWmsExpressBigBagCollectionBigBagNumIndex(bigBagNum string) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`bigBag_num` = ?", bigBagNum).Find(&results).Error

	return
}

// FetchIndexByWmsExpressBigBagCollectionCreateTimeIndex  获取多个内容
func (obj *_WmsExpressBigbagCollectionMgr) FetchIndexByWmsExpressBigBagCollectionCreateTimeIndex(createTime time.Time) (results []*WmsExpressBigbagCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressBigbagCollection{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
