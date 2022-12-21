package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WWorkOrderCommentMgr struct {
	*_BaseMgr
}

// WWorkOrderCommentMgr open func
func WWorkOrderCommentMgr(db *gorm.DB) *_WWorkOrderCommentMgr {
	if db == nil {
		panic(fmt.Errorf("WWorkOrderCommentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WWorkOrderCommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_work_order_comment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WWorkOrderCommentMgr) GetTableName() string {
	return "w_work_order_comment"
}

// Reset 重置gorm会话
func (obj *_WWorkOrderCommentMgr) Reset() *_WWorkOrderCommentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WWorkOrderCommentMgr) Get() (result WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WWorkOrderCommentMgr) Gets() (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WWorkOrderCommentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WWorkOrderCommentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 工单ID
func (obj *_WWorkOrderCommentMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithType type获取 类型 1:工单 2:任务
func (obj *_WWorkOrderCommentMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithContent content获取 留言内容
func (obj *_WWorkOrderCommentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithRemark remark获取 备注
func (obj *_WWorkOrderCommentMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WWorkOrderCommentMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WWorkOrderCommentMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WWorkOrderCommentMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WWorkOrderCommentMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WWorkOrderCommentMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WWorkOrderCommentMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WWorkOrderCommentMgr) GetByOption(opts ...Option) (result WWorkOrderComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WWorkOrderCommentMgr) GetByOptions(opts ...Option) (results []*WWorkOrderComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WWorkOrderCommentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WWorkOrderComment, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where(options.query)
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
func (obj *_WWorkOrderCommentMgr) GetFromID(id int) (result WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WWorkOrderCommentMgr) GetBatchFromID(ids []int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 工单ID
func (obj *_WWorkOrderCommentMgr) GetFromParentID(parentID int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 工单ID
func (obj *_WWorkOrderCommentMgr) GetBatchFromParentID(parentIDs []int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型 1:工单 2:任务
func (obj *_WWorkOrderCommentMgr) GetFromType(_type int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型 1:工单 2:任务
func (obj *_WWorkOrderCommentMgr) GetBatchFromType(_types []int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 留言内容
func (obj *_WWorkOrderCommentMgr) GetFromContent(content string) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 留言内容
func (obj *_WWorkOrderCommentMgr) GetBatchFromContent(contents []string) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WWorkOrderCommentMgr) GetFromRemark(remark string) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WWorkOrderCommentMgr) GetBatchFromRemark(remarks []string) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WWorkOrderCommentMgr) GetFromCreateTime(createTime time.Time) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WWorkOrderCommentMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WWorkOrderCommentMgr) GetFromCreateUser(createUser int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WWorkOrderCommentMgr) GetBatchFromCreateUser(createUsers []int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WWorkOrderCommentMgr) GetFromUpdateTime(updateTime time.Time) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WWorkOrderCommentMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WWorkOrderCommentMgr) GetFromUpdateUser(updateUser int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WWorkOrderCommentMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WWorkOrderCommentMgr) GetFromVersion(version int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WWorkOrderCommentMgr) GetBatchFromVersion(versions []int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WWorkOrderCommentMgr) GetFromDeleted(deleted int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WWorkOrderCommentMgr) GetBatchFromDeleted(deleteds []int) (results []*WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WWorkOrderCommentMgr) FetchByPrimaryKey(id int) (result WWorkOrderComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderComment{}).Where("`id` = ?", id).First(&result).Error

	return
}
