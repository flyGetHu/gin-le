package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgNumPoolFetchRecordMgr struct {
	*_BaseMgr
}

// LgNumPoolFetchRecordMgr open func
func LgNumPoolFetchRecordMgr(db *gorm.DB) *_LgNumPoolFetchRecordMgr {
	if db == nil {
		panic(fmt.Errorf("LgNumPoolFetchRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgNumPoolFetchRecordMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_num_pool_fetch_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgNumPoolFetchRecordMgr) GetTableName() string {
	return "lg_num_pool_fetch_record"
}

// Reset 重置gorm会话
func (obj *_LgNumPoolFetchRecordMgr) Reset() *_LgNumPoolFetchRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgNumPoolFetchRecordMgr) Get() (result LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgNumPoolFetchRecordMgr) Gets() (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgNumPoolFetchRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 单号池提取记录ID
func (obj *_LgNumPoolFetchRecordMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumID order_num_id获取 单号池关联ID
func (obj *_LgNumPoolFetchRecordMgr) WithOrderNumID(orderNumID int64) Option {
	return optionFunc(func(o *options) { o.query["order_num_id"] = orderNumID })
}

// WithName name获取 单号池名称
func (obj *_LgNumPoolFetchRecordMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrefix prefix获取 单号池前缀
func (obj *_LgNumPoolFetchRecordMgr) WithPrefix(prefix string) Option {
	return optionFunc(func(o *options) { o.query["prefix"] = prefix })
}

// WithSuffix suffix获取 单号池后缀
func (obj *_LgNumPoolFetchRecordMgr) WithSuffix(suffix string) Option {
	return optionFunc(func(o *options) { o.query["suffix"] = suffix })
}

// WithStartNum start_num获取 单号池提取开始位置,数字部分
func (obj *_LgNumPoolFetchRecordMgr) WithStartNum(startNum int64) Option {
	return optionFunc(func(o *options) { o.query["start_num"] = startNum })
}

// WithEndNum end_num获取 单号池提取结束位置,数字部分
func (obj *_LgNumPoolFetchRecordMgr) WithEndNum(endNum int64) Option {
	return optionFunc(func(o *options) { o.query["end_num"] = endNum })
}

// WithType type获取 提取类型:self:自用 customer:客户使用
func (obj *_LgNumPoolFetchRecordMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_LgNumPoolFetchRecordMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerName customer_name获取 客户名称
func (obj *_LgNumPoolFetchRecordMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithRemark remark获取 备注
func (obj *_LgNumPoolFetchRecordMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgNumPoolFetchRecordMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgNumPoolFetchRecordMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// GetByOption 功能选项模式获取
func (obj *_LgNumPoolFetchRecordMgr) GetByOption(opts ...Option) (result LgNumPoolFetchRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgNumPoolFetchRecordMgr) GetByOptions(opts ...Option) (results []*LgNumPoolFetchRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgNumPoolFetchRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgNumPoolFetchRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where(options.query)
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

// GetFromID 通过id获取内容 单号池提取记录ID
func (obj *_LgNumPoolFetchRecordMgr) GetFromID(id int64) (result LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 单号池提取记录ID
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromID(ids []int64) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumID 通过order_num_id获取内容 单号池关联ID
func (obj *_LgNumPoolFetchRecordMgr) GetFromOrderNumID(orderNumID int64) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`order_num_id` = ?", orderNumID).Find(&results).Error

	return
}

// GetBatchFromOrderNumID 批量查找 单号池关联ID
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromOrderNumID(orderNumIDs []int64) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`order_num_id` IN (?)", orderNumIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 单号池名称
func (obj *_LgNumPoolFetchRecordMgr) GetFromName(name string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 单号池名称
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromName(names []string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPrefix 通过prefix获取内容 单号池前缀
func (obj *_LgNumPoolFetchRecordMgr) GetFromPrefix(prefix string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`prefix` = ?", prefix).Find(&results).Error

	return
}

// GetBatchFromPrefix 批量查找 单号池前缀
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromPrefix(prefixs []string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`prefix` IN (?)", prefixs).Find(&results).Error

	return
}

// GetFromSuffix 通过suffix获取内容 单号池后缀
func (obj *_LgNumPoolFetchRecordMgr) GetFromSuffix(suffix string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`suffix` = ?", suffix).Find(&results).Error

	return
}

// GetBatchFromSuffix 批量查找 单号池后缀
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromSuffix(suffixs []string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`suffix` IN (?)", suffixs).Find(&results).Error

	return
}

// GetFromStartNum 通过start_num获取内容 单号池提取开始位置,数字部分
func (obj *_LgNumPoolFetchRecordMgr) GetFromStartNum(startNum int64) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`start_num` = ?", startNum).Find(&results).Error

	return
}

// GetBatchFromStartNum 批量查找 单号池提取开始位置,数字部分
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromStartNum(startNums []int64) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`start_num` IN (?)", startNums).Find(&results).Error

	return
}

// GetFromEndNum 通过end_num获取内容 单号池提取结束位置,数字部分
func (obj *_LgNumPoolFetchRecordMgr) GetFromEndNum(endNum int64) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`end_num` = ?", endNum).Find(&results).Error

	return
}

// GetBatchFromEndNum 批量查找 单号池提取结束位置,数字部分
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromEndNum(endNums []int64) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`end_num` IN (?)", endNums).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 提取类型:self:自用 customer:客户使用
func (obj *_LgNumPoolFetchRecordMgr) GetFromType(_type string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 提取类型:self:自用 customer:客户使用
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromType(_types []string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_LgNumPoolFetchRecordMgr) GetFromCustomerID(customerID int) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromCustomerID(customerIDs []int) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 客户名称
func (obj *_LgNumPoolFetchRecordMgr) GetFromCustomerName(customerName string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 客户名称
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromCustomerName(customerNames []string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_LgNumPoolFetchRecordMgr) GetFromRemark(remark string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromRemark(remarks []string) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgNumPoolFetchRecordMgr) GetFromCreateTime(createTime time.Time) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgNumPoolFetchRecordMgr) GetFromCreateUser(createUser int) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgNumPoolFetchRecordMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgNumPoolFetchRecordMgr) FetchByPrimaryKey(id int64) (result LgNumPoolFetchRecord, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolFetchRecord{}).Where("`id` = ?", id).First(&result).Error

	return
}
