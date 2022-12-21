package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOperateOrderLogMgr struct {
	*_BaseMgr
}

// LgOperateOrderLogMgr open func
func LgOperateOrderLogMgr(db *gorm.DB) *_LgOperateOrderLogMgr {
	if db == nil {
		panic(fmt.Errorf("LgOperateOrderLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOperateOrderLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_operate_order_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOperateOrderLogMgr) GetTableName() string {
	return "lg_operate_order_log"
}

// Reset 重置gorm会话
func (obj *_LgOperateOrderLogMgr) Reset() *_LgOperateOrderLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOperateOrderLogMgr) Get() (result LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOperateOrderLogMgr) Gets() (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOperateOrderLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgOperateOrderLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithReferenceNumber reference_number获取 参考号(安骏单号)
func (obj *_LgOperateOrderLogMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithOperateType operate_type获取 操作类型
func (obj *_LgOperateOrderLogMgr) WithOperateType(operateType string) Option {
	return optionFunc(func(o *options) { o.query["operate_type"] = operateType })
}

// WithOperateDescription operate_description获取 操作描述
func (obj *_LgOperateOrderLogMgr) WithOperateDescription(operateDescription string) Option {
	return optionFunc(func(o *options) { o.query["operate_description"] = operateDescription })
}

// WithOperateIP operate_ip获取 操作IP
func (obj *_LgOperateOrderLogMgr) WithOperateIP(operateIP string) Option {
	return optionFunc(func(o *options) { o.query["operate_ip"] = operateIP })
}

// WithCreateUser create_user获取 操作人id
func (obj *_LgOperateOrderLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 操作时间
func (obj *_LgOperateOrderLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithVersion version获取 乐观锁
func (obj *_LgOperateOrderLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOperateOrderLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOperateOrderLogMgr) GetByOption(opts ...Option) (result LgOperateOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOperateOrderLogMgr) GetByOptions(opts ...Option) (results []*LgOperateOrderLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOperateOrderLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOperateOrderLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where(options.query)
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
func (obj *_LgOperateOrderLogMgr) GetFromID(id int) (result LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgOperateOrderLogMgr) GetBatchFromID(ids []int) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号(安骏单号)
func (obj *_LgOperateOrderLogMgr) GetFromReferenceNumber(referenceNumber string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号(安骏单号)
func (obj *_LgOperateOrderLogMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromOperateType 通过operate_type获取内容 操作类型
func (obj *_LgOperateOrderLogMgr) GetFromOperateType(operateType string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`operate_type` = ?", operateType).Find(&results).Error

	return
}

// GetBatchFromOperateType 批量查找 操作类型
func (obj *_LgOperateOrderLogMgr) GetBatchFromOperateType(operateTypes []string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`operate_type` IN (?)", operateTypes).Find(&results).Error

	return
}

// GetFromOperateDescription 通过operate_description获取内容 操作描述
func (obj *_LgOperateOrderLogMgr) GetFromOperateDescription(operateDescription string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`operate_description` = ?", operateDescription).Find(&results).Error

	return
}

// GetBatchFromOperateDescription 批量查找 操作描述
func (obj *_LgOperateOrderLogMgr) GetBatchFromOperateDescription(operateDescriptions []string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`operate_description` IN (?)", operateDescriptions).Find(&results).Error

	return
}

// GetFromOperateIP 通过operate_ip获取内容 操作IP
func (obj *_LgOperateOrderLogMgr) GetFromOperateIP(operateIP string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`operate_ip` = ?", operateIP).Find(&results).Error

	return
}

// GetBatchFromOperateIP 批量查找 操作IP
func (obj *_LgOperateOrderLogMgr) GetBatchFromOperateIP(operateIPs []string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`operate_ip` IN (?)", operateIPs).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 操作人id
func (obj *_LgOperateOrderLogMgr) GetFromCreateUser(createUser int) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 操作人id
func (obj *_LgOperateOrderLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 操作时间
func (obj *_LgOperateOrderLogMgr) GetFromCreateTime(createTime time.Time) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 操作时间
func (obj *_LgOperateOrderLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOperateOrderLogMgr) GetFromVersion(version int) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOperateOrderLogMgr) GetBatchFromVersion(versions []int) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOperateOrderLogMgr) GetFromDeleted(deleted int) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOperateOrderLogMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOperateOrderLogMgr) FetchByPrimaryKey(id int) (result LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgOperateOrderLogIDUIndex primary or index 获取唯一内容
func (obj *_LgOperateOrderLogMgr) FetchUniqueByLgOperateOrderLogIDUIndex(id int) (result LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByLgOperateOrderLogReferenceNumberIndex  获取多个内容
func (obj *_LgOperateOrderLogMgr) FetchIndexByLgOperateOrderLogReferenceNumberIndex(referenceNumber string) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// FetchIndexByLgOperateOrderLogCreateTimeIndex  获取多个内容
func (obj *_LgOperateOrderLogMgr) FetchIndexByLgOperateOrderLogCreateTimeIndex(createTime time.Time) (results []*LgOperateOrderLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOperateOrderLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
