package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WWorkOrderDailyReportMgr struct {
	*_BaseMgr
}

// WWorkOrderDailyReportMgr open func
func WWorkOrderDailyReportMgr(db *gorm.DB) *_WWorkOrderDailyReportMgr {
	if db == nil {
		panic(fmt.Errorf("WWorkOrderDailyReportMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WWorkOrderDailyReportMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_work_order_daily_report"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WWorkOrderDailyReportMgr) GetTableName() string {
	return "w_work_order_daily_report"
}

// Reset 重置gorm会话
func (obj *_WWorkOrderDailyReportMgr) Reset() *_WWorkOrderDailyReportMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WWorkOrderDailyReportMgr) Get() (result WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WWorkOrderDailyReportMgr) Gets() (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WWorkOrderDailyReportMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_WWorkOrderDailyReportMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCompleteTask complete_task获取 完成的任务
func (obj *_WWorkOrderDailyReportMgr) WithCompleteTask(completeTask string) Option {
	return optionFunc(func(o *options) { o.query["complete_task"] = completeTask })
}

// WithTodoTask todo_task获取 待办任务
func (obj *_WWorkOrderDailyReportMgr) WithTodoTask(todoTask string) Option {
	return optionFunc(func(o *options) { o.query["todo_task"] = todoTask })
}

// WithAssistTask assist_task获取 需要协助任务
func (obj *_WWorkOrderDailyReportMgr) WithAssistTask(assistTask string) Option {
	return optionFunc(func(o *options) { o.query["assist_task"] = assistTask })
}

// WithState state获取 日报状态
func (obj *_WWorkOrderDailyReportMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithRemark remark获取 备注
func (obj *_WWorkOrderDailyReportMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithResources resources获取 附件
func (obj *_WWorkOrderDailyReportMgr) WithResources(resources []byte) Option {
	return optionFunc(func(o *options) { o.query["resources"] = resources })
}

// WithCopyToUserID copy_to_user_id获取 抄送人ID
func (obj *_WWorkOrderDailyReportMgr) WithCopyToUserID(copyToUserID []byte) Option {
	return optionFunc(func(o *options) { o.query["copy_to_user_id"] = copyToUserID })
}

// WithReadTime read_time获取 阅读时间
func (obj *_WWorkOrderDailyReportMgr) WithReadTime(readTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["read_time"] = readTime })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WWorkOrderDailyReportMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WWorkOrderDailyReportMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WWorkOrderDailyReportMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WWorkOrderDailyReportMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WWorkOrderDailyReportMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WWorkOrderDailyReportMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WWorkOrderDailyReportMgr) GetByOption(opts ...Option) (result WWorkOrderDailyReport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WWorkOrderDailyReportMgr) GetByOptions(opts ...Option) (results []*WWorkOrderDailyReport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WWorkOrderDailyReportMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WWorkOrderDailyReport, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_WWorkOrderDailyReportMgr) GetFromID(id int) (result WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromID(ids []int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCompleteTask 通过complete_task获取内容 完成的任务
func (obj *_WWorkOrderDailyReportMgr) GetFromCompleteTask(completeTask string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`complete_task` = ?", completeTask).Find(&results).Error

	return
}

// GetBatchFromCompleteTask 批量查找 完成的任务
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromCompleteTask(completeTasks []string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`complete_task` IN (?)", completeTasks).Find(&results).Error

	return
}

// GetFromTodoTask 通过todo_task获取内容 待办任务
func (obj *_WWorkOrderDailyReportMgr) GetFromTodoTask(todoTask string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`todo_task` = ?", todoTask).Find(&results).Error

	return
}

// GetBatchFromTodoTask 批量查找 待办任务
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromTodoTask(todoTasks []string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`todo_task` IN (?)", todoTasks).Find(&results).Error

	return
}

// GetFromAssistTask 通过assist_task获取内容 需要协助任务
func (obj *_WWorkOrderDailyReportMgr) GetFromAssistTask(assistTask string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`assist_task` = ?", assistTask).Find(&results).Error

	return
}

// GetBatchFromAssistTask 批量查找 需要协助任务
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromAssistTask(assistTasks []string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`assist_task` IN (?)", assistTasks).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 日报状态
func (obj *_WWorkOrderDailyReportMgr) GetFromState(state string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 日报状态
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromState(states []string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WWorkOrderDailyReportMgr) GetFromRemark(remark string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromRemark(remarks []string) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromResources 通过resources获取内容 附件
func (obj *_WWorkOrderDailyReportMgr) GetFromResources(resources []byte) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`resources` = ?", resources).Find(&results).Error

	return
}

// GetBatchFromResources 批量查找 附件
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromResources(resourcess [][]byte) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`resources` IN (?)", resourcess).Find(&results).Error

	return
}

// GetFromCopyToUserID 通过copy_to_user_id获取内容 抄送人ID
func (obj *_WWorkOrderDailyReportMgr) GetFromCopyToUserID(copyToUserID []byte) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`copy_to_user_id` = ?", copyToUserID).Find(&results).Error

	return
}

// GetBatchFromCopyToUserID 批量查找 抄送人ID
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromCopyToUserID(copyToUserIDs [][]byte) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`copy_to_user_id` IN (?)", copyToUserIDs).Find(&results).Error

	return
}

// GetFromReadTime 通过read_time获取内容 阅读时间
func (obj *_WWorkOrderDailyReportMgr) GetFromReadTime(readTime time.Time) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`read_time` = ?", readTime).Find(&results).Error

	return
}

// GetBatchFromReadTime 批量查找 阅读时间
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromReadTime(readTimes []time.Time) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`read_time` IN (?)", readTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WWorkOrderDailyReportMgr) GetFromCreateTime(createTime time.Time) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WWorkOrderDailyReportMgr) GetFromCreateUser(createUser int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromCreateUser(createUsers []int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WWorkOrderDailyReportMgr) GetFromUpdateTime(updateTime time.Time) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WWorkOrderDailyReportMgr) GetFromUpdateUser(updateUser int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WWorkOrderDailyReportMgr) GetFromVersion(version int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromVersion(versions []int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WWorkOrderDailyReportMgr) GetFromDeleted(deleted int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WWorkOrderDailyReportMgr) GetBatchFromDeleted(deleteds []int) (results []*WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WWorkOrderDailyReportMgr) FetchByPrimaryKey(id int) (result WWorkOrderDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderDailyReport{}).Where("`id` = ?", id).First(&result).Error

	return
}
