package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsExpressCollectionMgr struct {
	*_BaseMgr
}

// WmsExpressCollectionMgr open func
func WmsExpressCollectionMgr(db *gorm.DB) *_WmsExpressCollectionMgr {
	if db == nil {
		panic(fmt.Errorf("WmsExpressCollectionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsExpressCollectionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_express_collection"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsExpressCollectionMgr) GetTableName() string {
	return "wms_express_collection"
}

// Reset 重置gorm会话
func (obj *_WmsExpressCollectionMgr) Reset() *_WmsExpressCollectionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsExpressCollectionMgr) Get() (result WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsExpressCollectionMgr) Gets() (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsExpressCollectionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_WmsExpressCollectionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithExpressNum express_num获取 快递单号
func (obj *_WmsExpressCollectionMgr) WithExpressNum(expressNum string) Option {
	return optionFunc(func(o *options) { o.query["express_num"] = expressNum })
}

// WithBigBagNum bigBag_num获取 大包号
func (obj *_WmsExpressCollectionMgr) WithBigBagNum(bigBagNum string) Option {
	return optionFunc(func(o *options) { o.query["bigBag_num"] = bigBagNum })
}

// WithWeight weight获取 称重重量
func (obj *_WmsExpressCollectionMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithExpressResource express_resource获取 快递来源
func (obj *_WmsExpressCollectionMgr) WithExpressResource(expressResource string) Option {
	return optionFunc(func(o *options) { o.query["express_resource"] = expressResource })
}

// WithImageURL image_url获取 快递图片路径
func (obj *_WmsExpressCollectionMgr) WithImageURL(imageURL string) Option {
	return optionFunc(func(o *options) { o.query["image_url"] = imageURL })
}

// WithCreateUser create_user获取 操作人
func (obj *_WmsExpressCollectionMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsExpressCollectionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_WmsExpressCollectionMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsExpressCollectionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_WmsExpressCollectionMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsExpressCollectionMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithStatus status获取 状态 0 作废 1 正常 2 预录状态 默认为1
func (obj *_WmsExpressCollectionMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithIsLogisticsOrder is_logistics_order获取 是否是物流订单 1 是物流订单 0 不是物流订单
func (obj *_WmsExpressCollectionMgr) WithIsLogisticsOrder(isLogisticsOrder int) Option {
	return optionFunc(func(o *options) { o.query["is_logistics_order"] = isLogisticsOrder })
}

// WithPreEntryUserID pre_entry_user_id获取 预录人
func (obj *_WmsExpressCollectionMgr) WithPreEntryUserID(preEntryUserID int) Option {
	return optionFunc(func(o *options) { o.query["pre_entry_user_id"] = preEntryUserID })
}

// WithPreEntryTime pre_entry_time获取 预录时间
func (obj *_WmsExpressCollectionMgr) WithPreEntryTime(preEntryTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["pre_entry_time"] = preEntryTime })
}

// GetByOption 功能选项模式获取
func (obj *_WmsExpressCollectionMgr) GetByOption(opts ...Option) (result WmsExpressCollection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsExpressCollectionMgr) GetByOptions(opts ...Option) (results []*WmsExpressCollection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsExpressCollectionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsExpressCollection, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where(options.query)
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
func (obj *_WmsExpressCollectionMgr) GetFromID(id int) (result WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_WmsExpressCollectionMgr) GetBatchFromID(ids []int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromExpressNum 通过express_num获取内容 快递单号
func (obj *_WmsExpressCollectionMgr) GetFromExpressNum(expressNum string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`express_num` = ?", expressNum).Find(&results).Error

	return
}

// GetBatchFromExpressNum 批量查找 快递单号
func (obj *_WmsExpressCollectionMgr) GetBatchFromExpressNum(expressNums []string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`express_num` IN (?)", expressNums).Find(&results).Error

	return
}

// GetFromBigBagNum 通过bigBag_num获取内容 大包号
func (obj *_WmsExpressCollectionMgr) GetFromBigBagNum(bigBagNum string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`bigBag_num` = ?", bigBagNum).Find(&results).Error

	return
}

// GetBatchFromBigBagNum 批量查找 大包号
func (obj *_WmsExpressCollectionMgr) GetBatchFromBigBagNum(bigBagNums []string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`bigBag_num` IN (?)", bigBagNums).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 称重重量
func (obj *_WmsExpressCollectionMgr) GetFromWeight(weight float64) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 称重重量
func (obj *_WmsExpressCollectionMgr) GetBatchFromWeight(weights []float64) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromExpressResource 通过express_resource获取内容 快递来源
func (obj *_WmsExpressCollectionMgr) GetFromExpressResource(expressResource string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`express_resource` = ?", expressResource).Find(&results).Error

	return
}

// GetBatchFromExpressResource 批量查找 快递来源
func (obj *_WmsExpressCollectionMgr) GetBatchFromExpressResource(expressResources []string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`express_resource` IN (?)", expressResources).Find(&results).Error

	return
}

// GetFromImageURL 通过image_url获取内容 快递图片路径
func (obj *_WmsExpressCollectionMgr) GetFromImageURL(imageURL string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`image_url` = ?", imageURL).Find(&results).Error

	return
}

// GetBatchFromImageURL 批量查找 快递图片路径
func (obj *_WmsExpressCollectionMgr) GetBatchFromImageURL(imageURLs []string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`image_url` IN (?)", imageURLs).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 操作人
func (obj *_WmsExpressCollectionMgr) GetFromCreateUser(createUser int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 操作人
func (obj *_WmsExpressCollectionMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsExpressCollectionMgr) GetFromCreateTime(createTime time.Time) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsExpressCollectionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_WmsExpressCollectionMgr) GetFromUpdateUser(updateUser int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_WmsExpressCollectionMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsExpressCollectionMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsExpressCollectionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsExpressCollectionMgr) GetFromVersion(version int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsExpressCollectionMgr) GetBatchFromVersion(versions []int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsExpressCollectionMgr) GetFromDeleted(deleted int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsExpressCollectionMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态 0 作废 1 正常 2 预录状态 默认为1
func (obj *_WmsExpressCollectionMgr) GetFromStatus(status int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态 0 作废 1 正常 2 预录状态 默认为1
func (obj *_WmsExpressCollectionMgr) GetBatchFromStatus(statuss []int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromIsLogisticsOrder 通过is_logistics_order获取内容 是否是物流订单 1 是物流订单 0 不是物流订单
func (obj *_WmsExpressCollectionMgr) GetFromIsLogisticsOrder(isLogisticsOrder int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`is_logistics_order` = ?", isLogisticsOrder).Find(&results).Error

	return
}

// GetBatchFromIsLogisticsOrder 批量查找 是否是物流订单 1 是物流订单 0 不是物流订单
func (obj *_WmsExpressCollectionMgr) GetBatchFromIsLogisticsOrder(isLogisticsOrders []int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`is_logistics_order` IN (?)", isLogisticsOrders).Find(&results).Error

	return
}

// GetFromPreEntryUserID 通过pre_entry_user_id获取内容 预录人
func (obj *_WmsExpressCollectionMgr) GetFromPreEntryUserID(preEntryUserID int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`pre_entry_user_id` = ?", preEntryUserID).Find(&results).Error

	return
}

// GetBatchFromPreEntryUserID 批量查找 预录人
func (obj *_WmsExpressCollectionMgr) GetBatchFromPreEntryUserID(preEntryUserIDs []int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`pre_entry_user_id` IN (?)", preEntryUserIDs).Find(&results).Error

	return
}

// GetFromPreEntryTime 通过pre_entry_time获取内容 预录时间
func (obj *_WmsExpressCollectionMgr) GetFromPreEntryTime(preEntryTime time.Time) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`pre_entry_time` = ?", preEntryTime).Find(&results).Error

	return
}

// GetBatchFromPreEntryTime 批量查找 预录时间
func (obj *_WmsExpressCollectionMgr) GetBatchFromPreEntryTime(preEntryTimes []time.Time) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`pre_entry_time` IN (?)", preEntryTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsExpressCollectionMgr) FetchByPrimaryKey(id int) (result WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByExpressCollectionIDUIndex primary or index 获取唯一内容
func (obj *_WmsExpressCollectionMgr) FetchUniqueByExpressCollectionIDUIndex(id int) (result WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByExpressCollectionExpressNumIndex  获取多个内容
func (obj *_WmsExpressCollectionMgr) FetchIndexByExpressCollectionExpressNumIndex(expressNum string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`express_num` = ?", expressNum).Find(&results).Error

	return
}

// FetchIndexByExpressCollectionBigBagNumIndex  获取多个内容
func (obj *_WmsExpressCollectionMgr) FetchIndexByExpressCollectionBigBagNumIndex(bigBagNum string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`bigBag_num` = ?", bigBagNum).Find(&results).Error

	return
}

// FetchIndexByExpressCollectionExpressResourceIndex  获取多个内容
func (obj *_WmsExpressCollectionMgr) FetchIndexByExpressCollectionExpressResourceIndex(expressResource string) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`express_resource` = ?", expressResource).Find(&results).Error

	return
}

// FetchIndexByExpressCollectionPreEntryUserID  获取多个内容
func (obj *_WmsExpressCollectionMgr) FetchIndexByExpressCollectionPreEntryUserID(preEntryUserID int) (results []*WmsExpressCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressCollection{}).Where("`pre_entry_user_id` = ?", preEntryUserID).Find(&results).Error

	return
}
