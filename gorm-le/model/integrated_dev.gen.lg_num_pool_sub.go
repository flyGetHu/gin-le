package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgNumPoolSubMgr struct {
	*_BaseMgr
}

// LgNumPoolSubMgr open func
func LgNumPoolSubMgr(db *gorm.DB) *_LgNumPoolSubMgr {
	if db == nil {
		panic(fmt.Errorf("LgNumPoolSubMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgNumPoolSubMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_num_pool_sub"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgNumPoolSubMgr) GetTableName() string {
	return "lg_num_pool_sub"
}

// Reset 重置gorm会话
func (obj *_LgNumPoolSubMgr) Reset() *_LgNumPoolSubMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgNumPoolSubMgr) Get() (result LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgNumPoolSubMgr) Gets() (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgNumPoolSubMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 单号池副表主键ID
func (obj *_LgNumPoolSubMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumPoolID order_num_pool_id获取 绑定的单号池主单ID
func (obj *_LgNumPoolSubMgr) WithOrderNumPoolID(orderNumPoolID int64) Option {
	return optionFunc(func(o *options) { o.query["order_num_pool_id"] = orderNumPoolID })
}

// WithTrackNo track_no获取 物流单号
func (obj *_LgNumPoolSubMgr) WithTrackNo(trackNo string) Option {
	return optionFunc(func(o *options) { o.query["track_no"] = trackNo })
}

// WithExtractionOrderNum extraction_order_num获取 提取单号:取号后回填客户单号
func (obj *_LgNumPoolSubMgr) WithExtractionOrderNum(extractionOrderNum string) Option {
	return optionFunc(func(o *options) { o.query["extraction_order_num"] = extractionOrderNum })
}

// WithExtractionCustomerID extraction_customer_id获取 提取客户ID
func (obj *_LgNumPoolSubMgr) WithExtractionCustomerID(extractionCustomerID int64) Option {
	return optionFunc(func(o *options) { o.query["extraction_customer_id"] = extractionCustomerID })
}

// WithExtractionCustomerName extraction_customer_name获取 提取客户名称
func (obj *_LgNumPoolSubMgr) WithExtractionCustomerName(extractionCustomerName string) Option {
	return optionFunc(func(o *options) { o.query["extraction_customer_name"] = extractionCustomerName })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgNumPoolSubMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgNumPoolSubMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// GetByOption 功能选项模式获取
func (obj *_LgNumPoolSubMgr) GetByOption(opts ...Option) (result LgNumPoolSub, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgNumPoolSubMgr) GetByOptions(opts ...Option) (results []*LgNumPoolSub, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgNumPoolSubMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgNumPoolSub, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where(options.query)
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

// GetFromID 通过id获取内容 单号池副表主键ID
func (obj *_LgNumPoolSubMgr) GetFromID(id int64) (result LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 单号池副表主键ID
func (obj *_LgNumPoolSubMgr) GetBatchFromID(ids []int64) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumPoolID 通过order_num_pool_id获取内容 绑定的单号池主单ID
func (obj *_LgNumPoolSubMgr) GetFromOrderNumPoolID(orderNumPoolID int64) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`order_num_pool_id` = ?", orderNumPoolID).Find(&results).Error

	return
}

// GetBatchFromOrderNumPoolID 批量查找 绑定的单号池主单ID
func (obj *_LgNumPoolSubMgr) GetBatchFromOrderNumPoolID(orderNumPoolIDs []int64) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`order_num_pool_id` IN (?)", orderNumPoolIDs).Find(&results).Error

	return
}

// GetFromTrackNo 通过track_no获取内容 物流单号
func (obj *_LgNumPoolSubMgr) GetFromTrackNo(trackNo string) (result LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`track_no` = ?", trackNo).First(&result).Error

	return
}

// GetBatchFromTrackNo 批量查找 物流单号
func (obj *_LgNumPoolSubMgr) GetBatchFromTrackNo(trackNos []string) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`track_no` IN (?)", trackNos).Find(&results).Error

	return
}

// GetFromExtractionOrderNum 通过extraction_order_num获取内容 提取单号:取号后回填客户单号
func (obj *_LgNumPoolSubMgr) GetFromExtractionOrderNum(extractionOrderNum string) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`extraction_order_num` = ?", extractionOrderNum).Find(&results).Error

	return
}

// GetBatchFromExtractionOrderNum 批量查找 提取单号:取号后回填客户单号
func (obj *_LgNumPoolSubMgr) GetBatchFromExtractionOrderNum(extractionOrderNums []string) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`extraction_order_num` IN (?)", extractionOrderNums).Find(&results).Error

	return
}

// GetFromExtractionCustomerID 通过extraction_customer_id获取内容 提取客户ID
func (obj *_LgNumPoolSubMgr) GetFromExtractionCustomerID(extractionCustomerID int64) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`extraction_customer_id` = ?", extractionCustomerID).Find(&results).Error

	return
}

// GetBatchFromExtractionCustomerID 批量查找 提取客户ID
func (obj *_LgNumPoolSubMgr) GetBatchFromExtractionCustomerID(extractionCustomerIDs []int64) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`extraction_customer_id` IN (?)", extractionCustomerIDs).Find(&results).Error

	return
}

// GetFromExtractionCustomerName 通过extraction_customer_name获取内容 提取客户名称
func (obj *_LgNumPoolSubMgr) GetFromExtractionCustomerName(extractionCustomerName string) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`extraction_customer_name` = ?", extractionCustomerName).Find(&results).Error

	return
}

// GetBatchFromExtractionCustomerName 批量查找 提取客户名称
func (obj *_LgNumPoolSubMgr) GetBatchFromExtractionCustomerName(extractionCustomerNames []string) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`extraction_customer_name` IN (?)", extractionCustomerNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgNumPoolSubMgr) GetFromCreateTime(createTime time.Time) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgNumPoolSubMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgNumPoolSubMgr) GetFromCreateUser(createUser int) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgNumPoolSubMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgNumPoolSubMgr) FetchByPrimaryKey(id int64) (result LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgOrderNumPoolSubTrackNoUIndex primary or index 获取唯一内容
func (obj *_LgNumPoolSubMgr) FetchUniqueByLgOrderNumPoolSubTrackNoUIndex(trackNo string) (result LgNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPoolSub{}).Where("`track_no` = ?", trackNo).First(&result).Error

	return
}
