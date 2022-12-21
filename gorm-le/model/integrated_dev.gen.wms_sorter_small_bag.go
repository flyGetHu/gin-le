package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsSorterSmallBagMgr struct {
	*_BaseMgr
}

// WmsSorterSmallBagMgr open func
func WmsSorterSmallBagMgr(db *gorm.DB) *_WmsSorterSmallBagMgr {
	if db == nil {
		panic(fmt.Errorf("WmsSorterSmallBagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsSorterSmallBagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_sorter_small_bag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsSorterSmallBagMgr) GetTableName() string {
	return "wms_sorter_small_bag"
}

// Reset 重置gorm会话
func (obj *_WmsSorterSmallBagMgr) Reset() *_WmsSorterSmallBagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsSorterSmallBagMgr) Get() (result WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsSorterSmallBagMgr) Gets() (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsSorterSmallBagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsSorterSmallBagMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGridNum grid_num获取 格口
func (obj *_WmsSorterSmallBagMgr) WithGridNum(gridNum int) Option {
	return optionFunc(func(o *options) { o.query["grid_num"] = gridNum })
}

// WithOrderNumber order_number获取 订单号
func (obj *_WmsSorterSmallBagMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_WmsSorterSmallBagMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_WmsSorterSmallBagMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithWeight weight获取 重量
func (obj *_WmsSorterSmallBagMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithSorterBigBagNo sorter_big_bag_no获取 分拣机大包号
func (obj *_WmsSorterSmallBagMgr) WithSorterBigBagNo(sorterBigBagNo string) Option {
	return optionFunc(func(o *options) { o.query["sorter_big_bag_no"] = sorterBigBagNo })
}

// WithDownstreamCode downstream_code获取 下家code
func (obj *_WmsSorterSmallBagMgr) WithDownstreamCode(downstreamCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_code"] = downstreamCode })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道id
func (obj *_WmsSorterSmallBagMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsSorterSmallBagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsSorterSmallBagMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsSorterSmallBagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsSorterSmallBagMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsSorterSmallBagMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsSorterSmallBagMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsSorterSmallBagMgr) GetByOption(opts ...Option) (result WmsSorterSmallBag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsSorterSmallBagMgr) GetByOptions(opts ...Option) (results []*WmsSorterSmallBag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsSorterSmallBagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsSorterSmallBag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where(options.query)
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
func (obj *_WmsSorterSmallBagMgr) GetFromID(id int) (result WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsSorterSmallBagMgr) GetBatchFromID(ids []int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGridNum 通过grid_num获取内容 格口
func (obj *_WmsSorterSmallBagMgr) GetFromGridNum(gridNum int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`grid_num` = ?", gridNum).Find(&results).Error

	return
}

// GetBatchFromGridNum 批量查找 格口
func (obj *_WmsSorterSmallBagMgr) GetBatchFromGridNum(gridNums []int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`grid_num` IN (?)", gridNums).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_WmsSorterSmallBagMgr) GetFromOrderNumber(orderNumber string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_WmsSorterSmallBagMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_WmsSorterSmallBagMgr) GetFromReferenceNumber(referenceNumber string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_WmsSorterSmallBagMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_WmsSorterSmallBagMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_WmsSorterSmallBagMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 重量
func (obj *_WmsSorterSmallBagMgr) GetFromWeight(weight float64) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 重量
func (obj *_WmsSorterSmallBagMgr) GetBatchFromWeight(weights []float64) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromSorterBigBagNo 通过sorter_big_bag_no获取内容 分拣机大包号
func (obj *_WmsSorterSmallBagMgr) GetFromSorterBigBagNo(sorterBigBagNo string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`sorter_big_bag_no` = ?", sorterBigBagNo).Find(&results).Error

	return
}

// GetBatchFromSorterBigBagNo 批量查找 分拣机大包号
func (obj *_WmsSorterSmallBagMgr) GetBatchFromSorterBigBagNo(sorterBigBagNos []string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`sorter_big_bag_no` IN (?)", sorterBigBagNos).Find(&results).Error

	return
}

// GetFromDownstreamCode 通过downstream_code获取内容 下家code
func (obj *_WmsSorterSmallBagMgr) GetFromDownstreamCode(downstreamCode string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`downstream_code` = ?", downstreamCode).Find(&results).Error

	return
}

// GetBatchFromDownstreamCode 批量查找 下家code
func (obj *_WmsSorterSmallBagMgr) GetBatchFromDownstreamCode(downstreamCodes []string) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`downstream_code` IN (?)", downstreamCodes).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道id
func (obj *_WmsSorterSmallBagMgr) GetFromCustomerChannelID(customerChannelID int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道id
func (obj *_WmsSorterSmallBagMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsSorterSmallBagMgr) GetFromCreateTime(createTime time.Time) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsSorterSmallBagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsSorterSmallBagMgr) GetFromCreateUser(createUser int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsSorterSmallBagMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsSorterSmallBagMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsSorterSmallBagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsSorterSmallBagMgr) GetFromUpdateUser(updateUser int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsSorterSmallBagMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsSorterSmallBagMgr) GetFromVersion(version int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsSorterSmallBagMgr) GetBatchFromVersion(versions []int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsSorterSmallBagMgr) GetFromDeleted(deleted int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsSorterSmallBagMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsSorterSmallBagMgr) FetchByPrimaryKey(id int) (result WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByReferenceNumber primary or index 获取唯一内容
func (obj *_WmsSorterSmallBagMgr) FetchUniqueByReferenceNumber(referenceNumber string) (result WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`reference_number` = ?", referenceNumber).First(&result).Error

	return
}

// FetchUniqueByLogisticsNumber primary or index 获取唯一内容
func (obj *_WmsSorterSmallBagMgr) FetchUniqueByLogisticsNumber(logisticsNumber string) (result WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`logistics_number` = ?", logisticsNumber).First(&result).Error

	return
}

// FetchUniqueIndexByIndexSorterBigBagNoReferenceNumber primary or index 获取唯一内容
func (obj *_WmsSorterSmallBagMgr) FetchUniqueIndexByIndexSorterBigBagNoReferenceNumber(referenceNumber string, sorterBigBagNo string) (result WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`reference_number` = ? AND `sorter_big_bag_no` = ?", referenceNumber, sorterBigBagNo).First(&result).Error

	return
}

// FetchUniqueIndexByIndexSorterBigBagNoLogisticsNumber primary or index 获取唯一内容
func (obj *_WmsSorterSmallBagMgr) FetchUniqueIndexByIndexSorterBigBagNoLogisticsNumber(logisticsNumber string, sorterBigBagNo string) (result WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`logistics_number` = ? AND `sorter_big_bag_no` = ?", logisticsNumber, sorterBigBagNo).First(&result).Error

	return
}

// FetchIndexByIndexCreateTime  获取多个内容
func (obj *_WmsSorterSmallBagMgr) FetchIndexByIndexCreateTime(createTime time.Time) (results []*WmsSorterSmallBag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterSmallBag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
