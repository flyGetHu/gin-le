package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WWorkOrderOperationLogMgr struct {
	*_BaseMgr
}

// WWorkOrderOperationLogMgr open func
func WWorkOrderOperationLogMgr(db *gorm.DB) *_WWorkOrderOperationLogMgr {
	if db == nil {
		panic(fmt.Errorf("WWorkOrderOperationLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WWorkOrderOperationLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_work_order_operation_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WWorkOrderOperationLogMgr) GetTableName() string {
	return "w_work_order_operation_log"
}

// Reset 重置gorm会话
func (obj *_WWorkOrderOperationLogMgr) Reset() *_WWorkOrderOperationLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WWorkOrderOperationLogMgr) Get() (result WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WWorkOrderOperationLogMgr) Gets() (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WWorkOrderOperationLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WWorkOrderOperationLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 工单id
func (obj *_WWorkOrderOperationLogMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithTaskID task_id获取 任务id
func (obj *_WWorkOrderOperationLogMgr) WithTaskID(taskID int) Option {
	return optionFunc(func(o *options) { o.query["task_id"] = taskID })
}

// WithOperationContent operation_content获取 操作内容
func (obj *_WWorkOrderOperationLogMgr) WithOperationContent(operationContent string) Option {
	return optionFunc(func(o *options) { o.query["operation_content"] = operationContent })
}

// WithRemark remark获取 备注
func (obj *_WWorkOrderOperationLogMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WWorkOrderOperationLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WWorkOrderOperationLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WWorkOrderOperationLogMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WWorkOrderOperationLogMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WWorkOrderOperationLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WWorkOrderOperationLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WWorkOrderOperationLogMgr) GetByOption(opts ...Option) (result WWorkOrderOperationLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WWorkOrderOperationLogMgr) GetByOptions(opts ...Option) (results []*WWorkOrderOperationLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WWorkOrderOperationLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WWorkOrderOperationLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where(options.query)
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
func (obj *_WWorkOrderOperationLogMgr) GetFromID(id int) (result WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromID(ids []int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 工单id
func (obj *_WWorkOrderOperationLogMgr) GetFromOrderID(orderID int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 工单id
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromOrderID(orderIDs []int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromTaskID 通过task_id获取内容 任务id
func (obj *_WWorkOrderOperationLogMgr) GetFromTaskID(taskID int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`task_id` = ?", taskID).Find(&results).Error

	return
}

// GetBatchFromTaskID 批量查找 任务id
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromTaskID(taskIDs []int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`task_id` IN (?)", taskIDs).Find(&results).Error

	return
}

// GetFromOperationContent 通过operation_content获取内容 操作内容
func (obj *_WWorkOrderOperationLogMgr) GetFromOperationContent(operationContent string) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`operation_content` = ?", operationContent).Find(&results).Error

	return
}

// GetBatchFromOperationContent 批量查找 操作内容
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromOperationContent(operationContents []string) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`operation_content` IN (?)", operationContents).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WWorkOrderOperationLogMgr) GetFromRemark(remark string) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromRemark(remarks []string) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WWorkOrderOperationLogMgr) GetFromCreateTime(createTime time.Time) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WWorkOrderOperationLogMgr) GetFromCreateUser(createUser int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WWorkOrderOperationLogMgr) GetFromUpdateTime(updateTime time.Time) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WWorkOrderOperationLogMgr) GetFromUpdateUser(updateUser int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WWorkOrderOperationLogMgr) GetFromVersion(version int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromVersion(versions []int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WWorkOrderOperationLogMgr) GetFromDeleted(deleted int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WWorkOrderOperationLogMgr) GetBatchFromDeleted(deleteds []int) (results []*WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WWorkOrderOperationLogMgr) FetchByPrimaryKey(id int) (result WWorkOrderOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderOperationLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
