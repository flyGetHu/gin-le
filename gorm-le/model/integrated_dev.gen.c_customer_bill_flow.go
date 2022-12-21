package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CCustomerBillFlowMgr struct {
	*_BaseMgr
}

// CCustomerBillFlowMgr open func
func CCustomerBillFlowMgr(db *gorm.DB) *_CCustomerBillFlowMgr {
	if db == nil {
		panic(fmt.Errorf("CCustomerBillFlowMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CCustomerBillFlowMgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_customer_bill_flow"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CCustomerBillFlowMgr) GetTableName() string {
	return "c_customer_bill_flow"
}

// Reset 重置gorm会话
func (obj *_CCustomerBillFlowMgr) Reset() *_CCustomerBillFlowMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CCustomerBillFlowMgr) Get() (result CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CCustomerBillFlowMgr) Gets() (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CCustomerBillFlowMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_CCustomerBillFlowMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerID customer_id获取 客户id
func (obj *_CCustomerBillFlowMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerName customer_name获取 客户名称
func (obj *_CCustomerBillFlowMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithBillType bill_type获取 账单类型
func (obj *_CCustomerBillFlowMgr) WithBillType(billType string) Option {
	return optionFunc(func(o *options) { o.query["bill_type"] = billType })
}

// WithDeductionAmount deduction_amount获取 操作金额
func (obj *_CCustomerBillFlowMgr) WithDeductionAmount(deductionAmount float64) Option {
	return optionFunc(func(o *options) { o.query["deduction_amount"] = deductionAmount })
}

// WithBeforeBalance before_balance获取 操作前余额
func (obj *_CCustomerBillFlowMgr) WithBeforeBalance(beforeBalance float64) Option {
	return optionFunc(func(o *options) { o.query["before_balance"] = beforeBalance })
}

// WithAfterBalance after_balance获取 操作后余额
func (obj *_CCustomerBillFlowMgr) WithAfterBalance(afterBalance float64) Option {
	return optionFunc(func(o *options) { o.query["after_balance"] = afterBalance })
}

// WithCreateUser create_user获取 创建用户
func (obj *_CCustomerBillFlowMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CCustomerBillFlowMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithVersion version获取 乐观锁
func (obj *_CCustomerBillFlowMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_CCustomerBillFlowMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_CCustomerBillFlowMgr) GetByOption(opts ...Option) (result CCustomerBillFlow, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CCustomerBillFlowMgr) GetByOptions(opts ...Option) (results []*CCustomerBillFlow, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CCustomerBillFlowMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CCustomerBillFlow, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where(options.query)
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
func (obj *_CCustomerBillFlowMgr) GetFromID(id int) (result CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_CCustomerBillFlowMgr) GetBatchFromID(ids []int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_CCustomerBillFlowMgr) GetFromCustomerID(customerID int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_CCustomerBillFlowMgr) GetBatchFromCustomerID(customerIDs []int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 客户名称
func (obj *_CCustomerBillFlowMgr) GetFromCustomerName(customerName string) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 客户名称
func (obj *_CCustomerBillFlowMgr) GetBatchFromCustomerName(customerNames []string) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

// GetFromBillType 通过bill_type获取内容 账单类型
func (obj *_CCustomerBillFlowMgr) GetFromBillType(billType string) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`bill_type` = ?", billType).Find(&results).Error

	return
}

// GetBatchFromBillType 批量查找 账单类型
func (obj *_CCustomerBillFlowMgr) GetBatchFromBillType(billTypes []string) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`bill_type` IN (?)", billTypes).Find(&results).Error

	return
}

// GetFromDeductionAmount 通过deduction_amount获取内容 操作金额
func (obj *_CCustomerBillFlowMgr) GetFromDeductionAmount(deductionAmount float64) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`deduction_amount` = ?", deductionAmount).Find(&results).Error

	return
}

// GetBatchFromDeductionAmount 批量查找 操作金额
func (obj *_CCustomerBillFlowMgr) GetBatchFromDeductionAmount(deductionAmounts []float64) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`deduction_amount` IN (?)", deductionAmounts).Find(&results).Error

	return
}

// GetFromBeforeBalance 通过before_balance获取内容 操作前余额
func (obj *_CCustomerBillFlowMgr) GetFromBeforeBalance(beforeBalance float64) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`before_balance` = ?", beforeBalance).Find(&results).Error

	return
}

// GetBatchFromBeforeBalance 批量查找 操作前余额
func (obj *_CCustomerBillFlowMgr) GetBatchFromBeforeBalance(beforeBalances []float64) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`before_balance` IN (?)", beforeBalances).Find(&results).Error

	return
}

// GetFromAfterBalance 通过after_balance获取内容 操作后余额
func (obj *_CCustomerBillFlowMgr) GetFromAfterBalance(afterBalance float64) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`after_balance` = ?", afterBalance).Find(&results).Error

	return
}

// GetBatchFromAfterBalance 批量查找 操作后余额
func (obj *_CCustomerBillFlowMgr) GetBatchFromAfterBalance(afterBalances []float64) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`after_balance` IN (?)", afterBalances).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_CCustomerBillFlowMgr) GetFromCreateUser(createUser int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_CCustomerBillFlowMgr) GetBatchFromCreateUser(createUsers []int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CCustomerBillFlowMgr) GetFromCreateTime(createTime time.Time) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CCustomerBillFlowMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_CCustomerBillFlowMgr) GetFromVersion(version int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_CCustomerBillFlowMgr) GetBatchFromVersion(versions []int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_CCustomerBillFlowMgr) GetFromDeleted(deleted int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_CCustomerBillFlowMgr) GetBatchFromDeleted(deleteds []int) (results []*CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CCustomerBillFlowMgr) FetchByPrimaryKey(id int) (result CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCCustomerBillFlowIDUIndex primary or index 获取唯一内容
func (obj *_CCustomerBillFlowMgr) FetchUniqueByCCustomerBillFlowIDUIndex(id int) (result CCustomerBillFlow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerBillFlow{}).Where("`id` = ?", id).First(&result).Error

	return
}
