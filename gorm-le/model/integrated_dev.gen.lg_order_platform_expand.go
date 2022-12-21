package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderPlatformExpandMgr struct {
	*_BaseMgr
}

// LgOrderPlatformExpandMgr open func
func LgOrderPlatformExpandMgr(db *gorm.DB) *_LgOrderPlatformExpandMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderPlatformExpandMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderPlatformExpandMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_platform_expand"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderPlatformExpandMgr) GetTableName() string {
	return "lg_order_platform_expand"
}

// Reset 重置gorm会话
func (obj *_LgOrderPlatformExpandMgr) Reset() *_LgOrderPlatformExpandMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderPlatformExpandMgr) Get() (result LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderPlatformExpandMgr) Gets() (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderPlatformExpandMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgOrderPlatformExpandMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取  订单id
func (obj *_LgOrderPlatformExpandMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithPlatformNumber platform_number获取 平台订单号
func (obj *_LgOrderPlatformExpandMgr) WithPlatformNumber(platformNumber string) Option {
	return optionFunc(func(o *options) { o.query["platform_number"] = platformNumber })
}

// WithPlatformCode platform_code获取 平台代码
func (obj *_LgOrderPlatformExpandMgr) WithPlatformCode(platformCode string) Option {
	return optionFunc(func(o *options) { o.query["platform_code"] = platformCode })
}

// WithPlatformName platform_name获取 平台名称
func (obj *_LgOrderPlatformExpandMgr) WithPlatformName(platformName string) Option {
	return optionFunc(func(o *options) { o.query["platform_name"] = platformName })
}

// WithPlatformSellerName platform_seller_name获取 平台卖家名称
func (obj *_LgOrderPlatformExpandMgr) WithPlatformSellerName(platformSellerName string) Option {
	return optionFunc(func(o *options) { o.query["platform_seller_name"] = platformSellerName })
}

// WithPlatformOrderTime platform_order_time获取 平台订单时间
func (obj *_LgOrderPlatformExpandMgr) WithPlatformOrderTime(platformOrderTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["platform_order_time"] = platformOrderTime })
}

// WithCreateTime create_time获取  创建时间
func (obj *_LgOrderPlatformExpandMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取  创建用户
func (obj *_LgOrderPlatformExpandMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取  更新时间
func (obj *_LgOrderPlatformExpandMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取  更新用户
func (obj *_LgOrderPlatformExpandMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取  乐观锁
func (obj *_LgOrderPlatformExpandMgr) WithVersion(version uint) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取  逻辑删除
func (obj *_LgOrderPlatformExpandMgr) WithDeleted(deleted uint) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderPlatformExpandMgr) GetByOption(opts ...Option) (result LgOrderPlatformExpand, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderPlatformExpandMgr) GetByOptions(opts ...Option) (results []*LgOrderPlatformExpand, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderPlatformExpandMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderPlatformExpand, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where(options.query)
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
func (obj *_LgOrderPlatformExpandMgr) GetFromID(id int) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromID(ids []int) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容  订单id
func (obj *_LgOrderPlatformExpandMgr) GetFromOrderID(orderID int) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找  订单id
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromOrderID(orderIDs []int) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromPlatformNumber 通过platform_number获取内容 平台订单号
func (obj *_LgOrderPlatformExpandMgr) GetFromPlatformNumber(platformNumber string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_number` = ?", platformNumber).Find(&results).Error

	return
}

// GetBatchFromPlatformNumber 批量查找 平台订单号
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromPlatformNumber(platformNumbers []string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_number` IN (?)", platformNumbers).Find(&results).Error

	return
}

// GetFromPlatformCode 通过platform_code获取内容 平台代码
func (obj *_LgOrderPlatformExpandMgr) GetFromPlatformCode(platformCode string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_code` = ?", platformCode).Find(&results).Error

	return
}

// GetBatchFromPlatformCode 批量查找 平台代码
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromPlatformCode(platformCodes []string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_code` IN (?)", platformCodes).Find(&results).Error

	return
}

// GetFromPlatformName 通过platform_name获取内容 平台名称
func (obj *_LgOrderPlatformExpandMgr) GetFromPlatformName(platformName string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_name` = ?", platformName).Find(&results).Error

	return
}

// GetBatchFromPlatformName 批量查找 平台名称
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromPlatformName(platformNames []string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_name` IN (?)", platformNames).Find(&results).Error

	return
}

// GetFromPlatformSellerName 通过platform_seller_name获取内容 平台卖家名称
func (obj *_LgOrderPlatformExpandMgr) GetFromPlatformSellerName(platformSellerName string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_seller_name` = ?", platformSellerName).Find(&results).Error

	return
}

// GetBatchFromPlatformSellerName 批量查找 平台卖家名称
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromPlatformSellerName(platformSellerNames []string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_seller_name` IN (?)", platformSellerNames).Find(&results).Error

	return
}

// GetFromPlatformOrderTime 通过platform_order_time获取内容 平台订单时间
func (obj *_LgOrderPlatformExpandMgr) GetFromPlatformOrderTime(platformOrderTime time.Time) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_order_time` = ?", platformOrderTime).Find(&results).Error

	return
}

// GetBatchFromPlatformOrderTime 批量查找 平台订单时间
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromPlatformOrderTime(platformOrderTimes []time.Time) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_order_time` IN (?)", platformOrderTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容  创建时间
func (obj *_LgOrderPlatformExpandMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找  创建时间
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容  创建用户
func (obj *_LgOrderPlatformExpandMgr) GetFromCreateUser(createUser int) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找  创建用户
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容  更新时间
func (obj *_LgOrderPlatformExpandMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找  更新时间
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容  更新用户
func (obj *_LgOrderPlatformExpandMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找  更新用户
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容  乐观锁
func (obj *_LgOrderPlatformExpandMgr) GetFromVersion(version uint) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找  乐观锁
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromVersion(versions []uint) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容  逻辑删除
func (obj *_LgOrderPlatformExpandMgr) GetFromDeleted(deleted uint) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找  逻辑删除
func (obj *_LgOrderPlatformExpandMgr) GetBatchFromDeleted(deleteds []uint) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderPlatformExpandMgr) FetchByPrimaryKey(id int, orderID int) (result LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`id` = ? AND `order_id` = ?", id, orderID).First(&result).Error

	return
}

// FetchUniqueByUNIQUEORDERID primary or index 获取唯一内容
func (obj *_LgOrderPlatformExpandMgr) FetchUniqueByUNIQUEORDERID(orderID int) (result LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// FetchUniqueIndexByUNIQUEPLATFORMNUMBERPLATFORMCODE primary or index 获取唯一内容
func (obj *_LgOrderPlatformExpandMgr) FetchUniqueIndexByUNIQUEPLATFORMNUMBERPLATFORMCODE(platformNumber string, platformCode string) (result LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_number` = ? AND `platform_code` = ?", platformNumber, platformCode).First(&result).Error

	return
}

// FetchIndexByUNIQUEPLATFORMNUMBER  获取多个内容
func (obj *_LgOrderPlatformExpandMgr) FetchIndexByUNIQUEPLATFORMNUMBER(platformNumber string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_number` = ?", platformNumber).Find(&results).Error

	return
}

// FetchIndexByNORMALPLATFORMCODE  获取多个内容
func (obj *_LgOrderPlatformExpandMgr) FetchIndexByNORMALPLATFORMCODE(platformCode string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_code` = ?", platformCode).Find(&results).Error

	return
}

// FetchIndexByNORMALPLATFORMNAME  获取多个内容
func (obj *_LgOrderPlatformExpandMgr) FetchIndexByNORMALPLATFORMNAME(platformName string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_name` = ?", platformName).Find(&results).Error

	return
}

// FetchIndexByNORMALPLATFORMSELLERNAME  获取多个内容
func (obj *_LgOrderPlatformExpandMgr) FetchIndexByNORMALPLATFORMSELLERNAME(platformSellerName string) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_seller_name` = ?", platformSellerName).Find(&results).Error

	return
}

// FetchIndexByNORMALPLATFORMORDERTIME  获取多个内容
func (obj *_LgOrderPlatformExpandMgr) FetchIndexByNORMALPLATFORMORDERTIME(platformOrderTime time.Time) (results []*LgOrderPlatformExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderPlatformExpand{}).Where("`platform_order_time` = ?", platformOrderTime).Find(&results).Error

	return
}
