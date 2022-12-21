package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CCustomerFinanceMgr struct {
	*_BaseMgr
}

// CCustomerFinanceMgr open func
func CCustomerFinanceMgr(db *gorm.DB) *_CCustomerFinanceMgr {
	if db == nil {
		panic(fmt.Errorf("CCustomerFinanceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CCustomerFinanceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_customer_finance"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CCustomerFinanceMgr) GetTableName() string {
	return "c_customer_finance"
}

// Reset 重置gorm会话
func (obj *_CCustomerFinanceMgr) Reset() *_CCustomerFinanceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CCustomerFinanceMgr) Get() (result CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CCustomerFinanceMgr) Gets() (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CCustomerFinanceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_CCustomerFinanceMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerID customer_id获取 客户id
func (obj *_CCustomerFinanceMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerName customer_name获取 客户名称
func (obj *_CCustomerFinanceMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithBalance balance获取 账户余额
func (obj *_CCustomerFinanceMgr) WithBalance(balance float64) Option {
	return optionFunc(func(o *options) { o.query["balance"] = balance })
}

// WithCreateUser create_user获取 创建用户
func (obj *_CCustomerFinanceMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CCustomerFinanceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_CCustomerFinanceMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CCustomerFinanceMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_CCustomerFinanceMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_CCustomerFinanceMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_CCustomerFinanceMgr) GetByOption(opts ...Option) (result CCustomerFinance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CCustomerFinanceMgr) GetByOptions(opts ...Option) (results []*CCustomerFinance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CCustomerFinanceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CCustomerFinance, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where(options.query)
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
func (obj *_CCustomerFinanceMgr) GetFromID(id int) (result CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_CCustomerFinanceMgr) GetBatchFromID(ids []int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_CCustomerFinanceMgr) GetFromCustomerID(customerID int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_CCustomerFinanceMgr) GetBatchFromCustomerID(customerIDs []int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 客户名称
func (obj *_CCustomerFinanceMgr) GetFromCustomerName(customerName string) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 客户名称
func (obj *_CCustomerFinanceMgr) GetBatchFromCustomerName(customerNames []string) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

// GetFromBalance 通过balance获取内容 账户余额
func (obj *_CCustomerFinanceMgr) GetFromBalance(balance float64) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`balance` = ?", balance).Find(&results).Error

	return
}

// GetBatchFromBalance 批量查找 账户余额
func (obj *_CCustomerFinanceMgr) GetBatchFromBalance(balances []float64) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`balance` IN (?)", balances).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_CCustomerFinanceMgr) GetFromCreateUser(createUser int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_CCustomerFinanceMgr) GetBatchFromCreateUser(createUsers []int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CCustomerFinanceMgr) GetFromCreateTime(createTime time.Time) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CCustomerFinanceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_CCustomerFinanceMgr) GetFromUpdateUser(updateUser int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_CCustomerFinanceMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CCustomerFinanceMgr) GetFromUpdateTime(updateTime time.Time) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_CCustomerFinanceMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_CCustomerFinanceMgr) GetFromVersion(version int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_CCustomerFinanceMgr) GetBatchFromVersion(versions []int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_CCustomerFinanceMgr) GetFromDeleted(deleted int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_CCustomerFinanceMgr) GetBatchFromDeleted(deleteds []int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CCustomerFinanceMgr) FetchByPrimaryKey(id int) (result CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCCustomerFinanceIDUIndex primary or index 获取唯一内容
func (obj *_CCustomerFinanceMgr) FetchUniqueByCCustomerFinanceIDUIndex(id int) (result CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByCCustomerFinanceCustomerIDIndex  获取多个内容
func (obj *_CCustomerFinanceMgr) FetchIndexByCCustomerFinanceCustomerIDIndex(customerID int) (results []*CCustomerFinance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFinance{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}
