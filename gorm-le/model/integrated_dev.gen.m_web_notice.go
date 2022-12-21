package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MWebNoticeMgr struct {
	*_BaseMgr
}

// MWebNoticeMgr open func
func MWebNoticeMgr(db *gorm.DB) *_MWebNoticeMgr {
	if db == nil {
		panic(fmt.Errorf("MWebNoticeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MWebNoticeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("m_web_notice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MWebNoticeMgr) GetTableName() string {
	return "m_web_notice"
}

// Reset 重置gorm会话
func (obj *_MWebNoticeMgr) Reset() *_MWebNoticeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MWebNoticeMgr) Get() (result MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MWebNoticeMgr) Gets() (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MWebNoticeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 公告表ID
func (obj *_MWebNoticeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 标题
func (obj *_MWebNoticeMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_MWebNoticeMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithWatchVolume watch_volume获取 查看量
func (obj *_MWebNoticeMgr) WithWatchVolume(watchVolume int) Option {
	return optionFunc(func(o *options) { o.query["watch_volume"] = watchVolume })
}

// WithFileID file_id获取 文件表id
func (obj *_MWebNoticeMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_MWebNoticeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_MWebNoticeMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateUserName create_user_name获取 创建用户名称
func (obj *_MWebNoticeMgr) WithCreateUserName(createUserName string) Option {
	return optionFunc(func(o *options) { o.query["create_user_name"] = createUserName })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_MWebNoticeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_MWebNoticeMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateUserName update_user_name获取 更新用户名称
func (obj *_MWebNoticeMgr) WithUpdateUserName(updateUserName string) Option {
	return optionFunc(func(o *options) { o.query["update_user_name"] = updateUserName })
}

// WithVersion version获取 乐观锁
func (obj *_MWebNoticeMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_MWebNoticeMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_MWebNoticeMgr) GetByOption(opts ...Option) (result MWebNotice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MWebNoticeMgr) GetByOptions(opts ...Option) (results []*MWebNotice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MWebNoticeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MWebNotice, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where(options.query)
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

// GetFromID 通过id获取内容 公告表ID
func (obj *_MWebNoticeMgr) GetFromID(id int64) (result MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 公告表ID
func (obj *_MWebNoticeMgr) GetBatchFromID(ids []int64) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_MWebNoticeMgr) GetFromTitle(title string) (result MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`title` = ?", title).First(&result).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_MWebNoticeMgr) GetBatchFromTitle(titles []string) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_MWebNoticeMgr) GetFromContent(content string) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 内容
func (obj *_MWebNoticeMgr) GetBatchFromContent(contents []string) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromWatchVolume 通过watch_volume获取内容 查看量
func (obj *_MWebNoticeMgr) GetFromWatchVolume(watchVolume int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`watch_volume` = ?", watchVolume).Find(&results).Error

	return
}

// GetBatchFromWatchVolume 批量查找 查看量
func (obj *_MWebNoticeMgr) GetBatchFromWatchVolume(watchVolumes []int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`watch_volume` IN (?)", watchVolumes).Find(&results).Error

	return
}

// GetFromFileID 通过file_id获取内容 文件表id
func (obj *_MWebNoticeMgr) GetFromFileID(fileID int64) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`file_id` = ?", fileID).Find(&results).Error

	return
}

// GetBatchFromFileID 批量查找 文件表id
func (obj *_MWebNoticeMgr) GetBatchFromFileID(fileIDs []int64) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`file_id` IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_MWebNoticeMgr) GetFromCreateTime(createTime time.Time) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_MWebNoticeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_MWebNoticeMgr) GetFromCreateUser(createUser int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_MWebNoticeMgr) GetBatchFromCreateUser(createUsers []int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateUserName 通过create_user_name获取内容 创建用户名称
func (obj *_MWebNoticeMgr) GetFromCreateUserName(createUserName string) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`create_user_name` = ?", createUserName).Find(&results).Error

	return
}

// GetBatchFromCreateUserName 批量查找 创建用户名称
func (obj *_MWebNoticeMgr) GetBatchFromCreateUserName(createUserNames []string) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`create_user_name` IN (?)", createUserNames).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_MWebNoticeMgr) GetFromUpdateTime(updateTime time.Time) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_MWebNoticeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_MWebNoticeMgr) GetFromUpdateUser(updateUser int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_MWebNoticeMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateUserName 通过update_user_name获取内容 更新用户名称
func (obj *_MWebNoticeMgr) GetFromUpdateUserName(updateUserName string) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`update_user_name` = ?", updateUserName).Find(&results).Error

	return
}

// GetBatchFromUpdateUserName 批量查找 更新用户名称
func (obj *_MWebNoticeMgr) GetBatchFromUpdateUserName(updateUserNames []string) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`update_user_name` IN (?)", updateUserNames).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_MWebNoticeMgr) GetFromVersion(version int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_MWebNoticeMgr) GetBatchFromVersion(versions []int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_MWebNoticeMgr) GetFromDeleted(deleted int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_MWebNoticeMgr) GetBatchFromDeleted(deleteds []int) (results []*MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MWebNoticeMgr) FetchByPrimaryKey(id int64) (result MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByMNoticeTitleUIndex primary or index 获取唯一内容
func (obj *_MWebNoticeMgr) FetchUniqueByMNoticeTitleUIndex(title string) (result MWebNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MWebNotice{}).Where("`title` = ?", title).First(&result).Error

	return
}
