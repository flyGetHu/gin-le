package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderNumPoolSubMgr struct {
	*_BaseMgr
}

// LgOrderNumPoolSubMgr open func
func LgOrderNumPoolSubMgr(db *gorm.DB) *_LgOrderNumPoolSubMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderNumPoolSubMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderNumPoolSubMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_num_pool_sub"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderNumPoolSubMgr) GetTableName() string {
	return "lg_order_num_pool_sub"
}

// Reset 重置gorm会话
func (obj *_LgOrderNumPoolSubMgr) Reset() *_LgOrderNumPoolSubMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderNumPoolSubMgr) Get() (result LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderNumPoolSubMgr) Gets() (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderNumPoolSubMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 单号池副表主键ID
func (obj *_LgOrderNumPoolSubMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumPoolID order_num_pool_id获取 绑定的单号池主单ID
func (obj *_LgOrderNumPoolSubMgr) WithOrderNumPoolID(orderNumPoolID int64) Option {
	return optionFunc(func(o *options) { o.query["order_num_pool_id"] = orderNumPoolID })
}

// WithOrderNumPoolName order_num_pool_name获取 单号池名称关联主单名称
func (obj *_LgOrderNumPoolSubMgr) WithOrderNumPoolName(orderNumPoolName string) Option {
	return optionFunc(func(o *options) { o.query["order_num_pool_name"] = orderNumPoolName })
}

// WithTrackNo track_no获取 物流单号
func (obj *_LgOrderNumPoolSubMgr) WithTrackNo(trackNo string) Option {
	return optionFunc(func(o *options) { o.query["track_no"] = trackNo })
}

// WithExtractionOrderNum extraction_order_num获取 提取单号:取号后回填客户单号
func (obj *_LgOrderNumPoolSubMgr) WithExtractionOrderNum(extractionOrderNum string) Option {
	return optionFunc(func(o *options) { o.query["extraction_order_num"] = extractionOrderNum })
}

// WithExtractionCustomerID extraction_customer_id获取 提取客户ID
func (obj *_LgOrderNumPoolSubMgr) WithExtractionCustomerID(extractionCustomerID int64) Option {
	return optionFunc(func(o *options) { o.query["extraction_customer_id"] = extractionCustomerID })
}

// WithExtractionCustomerName extraction_customer_name获取 提取客户名称
func (obj *_LgOrderNumPoolSubMgr) WithExtractionCustomerName(extractionCustomerName string) Option {
	return optionFunc(func(o *options) { o.query["extraction_customer_name"] = extractionCustomerName })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderNumPoolSubMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderNumPoolSubMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderNumPoolSubMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderNumPoolSubMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderNumPoolSubMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderNumPoolSubMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderNumPoolSubMgr) GetByOption(opts ...Option) (result LgOrderNumPoolSub, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderNumPoolSubMgr) GetByOptions(opts ...Option) (results []*LgOrderNumPoolSub, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderNumPoolSubMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderNumPoolSub, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where(options.query)
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
func (obj *_LgOrderNumPoolSubMgr) GetFromID(id int64) (result LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 单号池副表主键ID
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromID(ids []int64) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumPoolID 通过order_num_pool_id获取内容 绑定的单号池主单ID
func (obj *_LgOrderNumPoolSubMgr) GetFromOrderNumPoolID(orderNumPoolID int64) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`order_num_pool_id` = ?", orderNumPoolID).Find(&results).Error

	return
}

// GetBatchFromOrderNumPoolID 批量查找 绑定的单号池主单ID
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromOrderNumPoolID(orderNumPoolIDs []int64) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`order_num_pool_id` IN (?)", orderNumPoolIDs).Find(&results).Error

	return
}

// GetFromOrderNumPoolName 通过order_num_pool_name获取内容 单号池名称关联主单名称
func (obj *_LgOrderNumPoolSubMgr) GetFromOrderNumPoolName(orderNumPoolName string) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`order_num_pool_name` = ?", orderNumPoolName).Find(&results).Error

	return
}

// GetBatchFromOrderNumPoolName 批量查找 单号池名称关联主单名称
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromOrderNumPoolName(orderNumPoolNames []string) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`order_num_pool_name` IN (?)", orderNumPoolNames).Find(&results).Error

	return
}

// GetFromTrackNo 通过track_no获取内容 物流单号
func (obj *_LgOrderNumPoolSubMgr) GetFromTrackNo(trackNo string) (result LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`track_no` = ?", trackNo).First(&result).Error

	return
}

// GetBatchFromTrackNo 批量查找 物流单号
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromTrackNo(trackNos []string) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`track_no` IN (?)", trackNos).Find(&results).Error

	return
}

// GetFromExtractionOrderNum 通过extraction_order_num获取内容 提取单号:取号后回填客户单号
func (obj *_LgOrderNumPoolSubMgr) GetFromExtractionOrderNum(extractionOrderNum string) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`extraction_order_num` = ?", extractionOrderNum).Find(&results).Error

	return
}

// GetBatchFromExtractionOrderNum 批量查找 提取单号:取号后回填客户单号
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromExtractionOrderNum(extractionOrderNums []string) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`extraction_order_num` IN (?)", extractionOrderNums).Find(&results).Error

	return
}

// GetFromExtractionCustomerID 通过extraction_customer_id获取内容 提取客户ID
func (obj *_LgOrderNumPoolSubMgr) GetFromExtractionCustomerID(extractionCustomerID int64) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`extraction_customer_id` = ?", extractionCustomerID).Find(&results).Error

	return
}

// GetBatchFromExtractionCustomerID 批量查找 提取客户ID
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromExtractionCustomerID(extractionCustomerIDs []int64) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`extraction_customer_id` IN (?)", extractionCustomerIDs).Find(&results).Error

	return
}

// GetFromExtractionCustomerName 通过extraction_customer_name获取内容 提取客户名称
func (obj *_LgOrderNumPoolSubMgr) GetFromExtractionCustomerName(extractionCustomerName string) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`extraction_customer_name` = ?", extractionCustomerName).Find(&results).Error

	return
}

// GetBatchFromExtractionCustomerName 批量查找 提取客户名称
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromExtractionCustomerName(extractionCustomerNames []string) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`extraction_customer_name` IN (?)", extractionCustomerNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderNumPoolSubMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderNumPoolSubMgr) GetFromCreateUser(createUser int) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderNumPoolSubMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderNumPoolSubMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderNumPoolSubMgr) GetFromVersion(version int) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromVersion(versions []int) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderNumPoolSubMgr) GetFromDeleted(deleted int) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderNumPoolSubMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderNumPoolSubMgr) FetchByPrimaryKey(id int64) (result LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgOrderNumPoolSubTrackNoUIndex primary or index 获取唯一内容
func (obj *_LgOrderNumPoolSubMgr) FetchUniqueByLgOrderNumPoolSubTrackNoUIndex(trackNo string) (result LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`track_no` = ?", trackNo).First(&result).Error

	return
}

// FetchIndexByLgOrderNumPoolSubLgOrderNumPoolIDFk  获取多个内容
func (obj *_LgOrderNumPoolSubMgr) FetchIndexByLgOrderNumPoolSubLgOrderNumPoolIDFk(orderNumPoolID int64) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`order_num_pool_id` = ?", orderNumPoolID).Find(&results).Error

	return
}

// FetchIndexByLgOrderNumPoolSubLgOrderNumPoolNameFk  获取多个内容
func (obj *_LgOrderNumPoolSubMgr) FetchIndexByLgOrderNumPoolSubLgOrderNumPoolNameFk(orderNumPoolName string) (results []*LgOrderNumPoolSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPoolSub{}).Where("`order_num_pool_name` = ?", orderNumPoolName).Find(&results).Error

	return
}
