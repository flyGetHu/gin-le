package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _WMicroTicketFollowUpInfoMgr struct {
	*_BaseMgr
}

// WMicroTicketFollowUpInfoMgr open func
func WMicroTicketFollowUpInfoMgr(db *gorm.DB) *_WMicroTicketFollowUpInfoMgr {
	if db == nil {
		panic(fmt.Errorf("WMicroTicketFollowUpInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WMicroTicketFollowUpInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_micro_ticket_follow_up_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WMicroTicketFollowUpInfoMgr) GetTableName() string {
	return "w_micro_ticket_follow_up_info"
}

// Reset 重置gorm会话
func (obj *_WMicroTicketFollowUpInfoMgr) Reset() *_WMicroTicketFollowUpInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WMicroTicketFollowUpInfoMgr) Get() (result WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WMicroTicketFollowUpInfoMgr) Gets() (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WMicroTicketFollowUpInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_WMicroTicketFollowUpInfoMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTicketID ticket_id获取 微工单id
func (obj *_WMicroTicketFollowUpInfoMgr) WithTicketID(ticketID int) Option {
	return optionFunc(func(o *options) { o.query["ticket_id"] = ticketID })
}

// WithFollowUpDescribe follow_up_describe获取 跟进描述
func (obj *_WMicroTicketFollowUpInfoMgr) WithFollowUpDescribe(followUpDescribe string) Option {
	return optionFunc(func(o *options) { o.query["follow_up_describe"] = followUpDescribe })
}

// WithFollower follower获取 跟进人
func (obj *_WMicroTicketFollowUpInfoMgr) WithFollower(follower datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["follower"] = follower })
}

// WithFollowUpStatus follow_up_status获取 跟进状态(0:正常跟进1:完结任务2:作废任务)
func (obj *_WMicroTicketFollowUpInfoMgr) WithFollowUpStatus(followUpStatus int) Option {
	return optionFunc(func(o *options) { o.query["follow_up_status"] = followUpStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WMicroTicketFollowUpInfoMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_WMicroTicketFollowUpInfoMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WMicroTicketFollowUpInfoMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_WMicroTicketFollowUpInfoMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WMicroTicketFollowUpInfoMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WMicroTicketFollowUpInfoMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WMicroTicketFollowUpInfoMgr) GetByOption(opts ...Option) (result WMicroTicketFollowUpInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WMicroTicketFollowUpInfoMgr) GetByOptions(opts ...Option) (results []*WMicroTicketFollowUpInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WMicroTicketFollowUpInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WMicroTicketFollowUpInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where(options.query)
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
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromID(id int) (result WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromID(ids []int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTicketID 通过ticket_id获取内容 微工单id
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromTicketID(ticketID int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`ticket_id` = ?", ticketID).Find(&results).Error

	return
}

// GetBatchFromTicketID 批量查找 微工单id
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromTicketID(ticketIDs []int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`ticket_id` IN (?)", ticketIDs).Find(&results).Error

	return
}

// GetFromFollowUpDescribe 通过follow_up_describe获取内容 跟进描述
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromFollowUpDescribe(followUpDescribe string) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`follow_up_describe` = ?", followUpDescribe).Find(&results).Error

	return
}

// GetBatchFromFollowUpDescribe 批量查找 跟进描述
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromFollowUpDescribe(followUpDescribes []string) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`follow_up_describe` IN (?)", followUpDescribes).Find(&results).Error

	return
}

// GetFromFollower 通过follower获取内容 跟进人
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromFollower(follower datatypes.JSON) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`follower` = ?", follower).Find(&results).Error

	return
}

// GetBatchFromFollower 批量查找 跟进人
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromFollower(followers []datatypes.JSON) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`follower` IN (?)", followers).Find(&results).Error

	return
}

// GetFromFollowUpStatus 通过follow_up_status获取内容 跟进状态(0:正常跟进1:完结任务2:作废任务)
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromFollowUpStatus(followUpStatus int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`follow_up_status` = ?", followUpStatus).Find(&results).Error

	return
}

// GetBatchFromFollowUpStatus 批量查找 跟进状态(0:正常跟进1:完结任务2:作废任务)
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromFollowUpStatus(followUpStatuss []int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`follow_up_status` IN (?)", followUpStatuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromCreateTime(createTime time.Time) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromCreateUser(createUser int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromCreateUser(createUsers []int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromUpdateTime(updateTime time.Time) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromUpdateUser(updateUser int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromVersion(version int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromVersion(versions []int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WMicroTicketFollowUpInfoMgr) GetFromDeleted(deleted int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WMicroTicketFollowUpInfoMgr) GetBatchFromDeleted(deleteds []int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WMicroTicketFollowUpInfoMgr) FetchByPrimaryKey(id int) (result WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByWMicroTicketFollowUpInfoIDUIndex primary or index 获取唯一内容
func (obj *_WMicroTicketFollowUpInfoMgr) FetchUniqueByWMicroTicketFollowUpInfoIDUIndex(id int) (result WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByWMicroTicketFollowUpLogTicketIDIndex  获取多个内容
func (obj *_WMicroTicketFollowUpInfoMgr) FetchIndexByWMicroTicketFollowUpLogTicketIDIndex(ticketID int) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`ticket_id` = ?", ticketID).Find(&results).Error

	return
}

// FetchIndexByWMicroTicketFollowUpLogCreateTimeIndex  获取多个内容
func (obj *_WMicroTicketFollowUpInfoMgr) FetchIndexByWMicroTicketFollowUpLogCreateTimeIndex(createTime time.Time) (results []*WMicroTicketFollowUpInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketFollowUpInfo{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
