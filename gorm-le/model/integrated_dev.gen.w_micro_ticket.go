package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _WMicroTicketMgr struct {
	*_BaseMgr
}

// WMicroTicketMgr open func
func WMicroTicketMgr(db *gorm.DB) *_WMicroTicketMgr {
	if db == nil {
		panic(fmt.Errorf("WMicroTicketMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WMicroTicketMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_micro_ticket"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WMicroTicketMgr) GetTableName() string {
	return "w_micro_ticket"
}

// Reset 重置gorm会话
func (obj *_WMicroTicketMgr) Reset() *_WMicroTicketMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WMicroTicketMgr) Get() (result WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WMicroTicketMgr) Gets() (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WMicroTicketMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 微工单
func (obj *_WMicroTicketMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTicketName ticket_name获取 工单名称
func (obj *_WMicroTicketMgr) WithTicketName(ticketName string) Option {
	return optionFunc(func(o *options) { o.query["ticket_name"] = ticketName })
}

// WithTicketType ticket_type获取 工单类型
func (obj *_WMicroTicketMgr) WithTicketType(ticketType string) Option {
	return optionFunc(func(o *options) { o.query["ticket_type"] = ticketType })
}

// WithTicketStatus ticket_status获取 工单状态
func (obj *_WMicroTicketMgr) WithTicketStatus(ticketStatus string) Option {
	return optionFunc(func(o *options) { o.query["ticket_status"] = ticketStatus })
}

// WithSort sort获取 优先级
func (obj *_WMicroTicketMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithTicketDescribe ticket_describe获取 工单描述
func (obj *_WMicroTicketMgr) WithTicketDescribe(ticketDescribe string) Option {
	return optionFunc(func(o *options) { o.query["ticket_describe"] = ticketDescribe })
}

// WithAppendix appendix获取 工单附件
func (obj *_WMicroTicketMgr) WithAppendix(appendix string) Option {
	return optionFunc(func(o *options) { o.query["appendix"] = appendix })
}

// WithLastFollower last_follower获取 最后跟进人
func (obj *_WMicroTicketMgr) WithLastFollower(lastFollower int) Option {
	return optionFunc(func(o *options) { o.query["last_follower"] = lastFollower })
}

// WithLastFollowTime last_follow_time获取 最后跟进时间
func (obj *_WMicroTicketMgr) WithLastFollowTime(lastFollowTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_follow_time"] = lastFollowTime })
}

// WithExecuteUser execute_user获取 执行人
func (obj *_WMicroTicketMgr) WithExecuteUser(executeUser datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["execute_user"] = executeUser })
}

// WithPlanFinishTime plan_finish_time获取 计划完成时间
func (obj *_WMicroTicketMgr) WithPlanFinishTime(planFinishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["plan_finish_time"] = planFinishTime })
}

// WithActualFinishTime actual_finish_time获取 实际完成时间
func (obj *_WMicroTicketMgr) WithActualFinishTime(actualFinishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["actual_finish_time"] = actualFinishTime })
}

// WithTimeConsuming time_consuming获取 工时
func (obj *_WMicroTicketMgr) WithTimeConsuming(timeConsuming int) Option {
	return optionFunc(func(o *options) { o.query["time_consuming"] = timeConsuming })
}

// WithCreateUser create_user获取 创建人
func (obj *_WMicroTicketMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WMicroTicketMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_WMicroTicketMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WMicroTicketMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_WMicroTicketMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WMicroTicketMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WMicroTicketMgr) GetByOption(opts ...Option) (result WMicroTicket, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WMicroTicketMgr) GetByOptions(opts ...Option) (results []*WMicroTicket, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WMicroTicketMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WMicroTicket, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where(options.query)
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

// GetFromID 通过id获取内容 微工单
func (obj *_WMicroTicketMgr) GetFromID(id int) (result WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 微工单
func (obj *_WMicroTicketMgr) GetBatchFromID(ids []int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTicketName 通过ticket_name获取内容 工单名称
func (obj *_WMicroTicketMgr) GetFromTicketName(ticketName string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`ticket_name` = ?", ticketName).Find(&results).Error

	return
}

// GetBatchFromTicketName 批量查找 工单名称
func (obj *_WMicroTicketMgr) GetBatchFromTicketName(ticketNames []string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`ticket_name` IN (?)", ticketNames).Find(&results).Error

	return
}

// GetFromTicketType 通过ticket_type获取内容 工单类型
func (obj *_WMicroTicketMgr) GetFromTicketType(ticketType string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`ticket_type` = ?", ticketType).Find(&results).Error

	return
}

// GetBatchFromTicketType 批量查找 工单类型
func (obj *_WMicroTicketMgr) GetBatchFromTicketType(ticketTypes []string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`ticket_type` IN (?)", ticketTypes).Find(&results).Error

	return
}

// GetFromTicketStatus 通过ticket_status获取内容 工单状态
func (obj *_WMicroTicketMgr) GetFromTicketStatus(ticketStatus string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`ticket_status` = ?", ticketStatus).Find(&results).Error

	return
}

// GetBatchFromTicketStatus 批量查找 工单状态
func (obj *_WMicroTicketMgr) GetBatchFromTicketStatus(ticketStatuss []string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`ticket_status` IN (?)", ticketStatuss).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 优先级
func (obj *_WMicroTicketMgr) GetFromSort(sort int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 优先级
func (obj *_WMicroTicketMgr) GetBatchFromSort(sorts []int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromTicketDescribe 通过ticket_describe获取内容 工单描述
func (obj *_WMicroTicketMgr) GetFromTicketDescribe(ticketDescribe string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`ticket_describe` = ?", ticketDescribe).Find(&results).Error

	return
}

// GetBatchFromTicketDescribe 批量查找 工单描述
func (obj *_WMicroTicketMgr) GetBatchFromTicketDescribe(ticketDescribes []string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`ticket_describe` IN (?)", ticketDescribes).Find(&results).Error

	return
}

// GetFromAppendix 通过appendix获取内容 工单附件
func (obj *_WMicroTicketMgr) GetFromAppendix(appendix string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`appendix` = ?", appendix).Find(&results).Error

	return
}

// GetBatchFromAppendix 批量查找 工单附件
func (obj *_WMicroTicketMgr) GetBatchFromAppendix(appendixs []string) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`appendix` IN (?)", appendixs).Find(&results).Error

	return
}

// GetFromLastFollower 通过last_follower获取内容 最后跟进人
func (obj *_WMicroTicketMgr) GetFromLastFollower(lastFollower int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`last_follower` = ?", lastFollower).Find(&results).Error

	return
}

// GetBatchFromLastFollower 批量查找 最后跟进人
func (obj *_WMicroTicketMgr) GetBatchFromLastFollower(lastFollowers []int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`last_follower` IN (?)", lastFollowers).Find(&results).Error

	return
}

// GetFromLastFollowTime 通过last_follow_time获取内容 最后跟进时间
func (obj *_WMicroTicketMgr) GetFromLastFollowTime(lastFollowTime time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`last_follow_time` = ?", lastFollowTime).Find(&results).Error

	return
}

// GetBatchFromLastFollowTime 批量查找 最后跟进时间
func (obj *_WMicroTicketMgr) GetBatchFromLastFollowTime(lastFollowTimes []time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`last_follow_time` IN (?)", lastFollowTimes).Find(&results).Error

	return
}

// GetFromExecuteUser 通过execute_user获取内容 执行人
func (obj *_WMicroTicketMgr) GetFromExecuteUser(executeUser datatypes.JSON) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`execute_user` = ?", executeUser).Find(&results).Error

	return
}

// GetBatchFromExecuteUser 批量查找 执行人
func (obj *_WMicroTicketMgr) GetBatchFromExecuteUser(executeUsers []datatypes.JSON) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`execute_user` IN (?)", executeUsers).Find(&results).Error

	return
}

// GetFromPlanFinishTime 通过plan_finish_time获取内容 计划完成时间
func (obj *_WMicroTicketMgr) GetFromPlanFinishTime(planFinishTime time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`plan_finish_time` = ?", planFinishTime).Find(&results).Error

	return
}

// GetBatchFromPlanFinishTime 批量查找 计划完成时间
func (obj *_WMicroTicketMgr) GetBatchFromPlanFinishTime(planFinishTimes []time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`plan_finish_time` IN (?)", planFinishTimes).Find(&results).Error

	return
}

// GetFromActualFinishTime 通过actual_finish_time获取内容 实际完成时间
func (obj *_WMicroTicketMgr) GetFromActualFinishTime(actualFinishTime time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`actual_finish_time` = ?", actualFinishTime).Find(&results).Error

	return
}

// GetBatchFromActualFinishTime 批量查找 实际完成时间
func (obj *_WMicroTicketMgr) GetBatchFromActualFinishTime(actualFinishTimes []time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`actual_finish_time` IN (?)", actualFinishTimes).Find(&results).Error

	return
}

// GetFromTimeConsuming 通过time_consuming获取内容 工时
func (obj *_WMicroTicketMgr) GetFromTimeConsuming(timeConsuming int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`time_consuming` = ?", timeConsuming).Find(&results).Error

	return
}

// GetBatchFromTimeConsuming 批量查找 工时
func (obj *_WMicroTicketMgr) GetBatchFromTimeConsuming(timeConsumings []int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`time_consuming` IN (?)", timeConsumings).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_WMicroTicketMgr) GetFromCreateUser(createUser int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_WMicroTicketMgr) GetBatchFromCreateUser(createUsers []int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WMicroTicketMgr) GetFromCreateTime(createTime time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WMicroTicketMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_WMicroTicketMgr) GetFromUpdateUser(updateUser int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_WMicroTicketMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WMicroTicketMgr) GetFromUpdateTime(updateTime time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WMicroTicketMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WMicroTicketMgr) GetFromVersion(version int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WMicroTicketMgr) GetBatchFromVersion(versions []int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WMicroTicketMgr) GetFromDeleted(deleted int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WMicroTicketMgr) GetBatchFromDeleted(deleteds []int) (results []*WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WMicroTicketMgr) FetchByPrimaryKey(id int) (result WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByWMicroTicketIDUIndex primary or index 获取唯一内容
func (obj *_WMicroTicketMgr) FetchUniqueByWMicroTicketIDUIndex(id int) (result WMicroTicket, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicket{}).Where("`id` = ?", id).First(&result).Error

	return
}
