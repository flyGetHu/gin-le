package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MInternalMessageNoticeMgr struct {
	*_BaseMgr
}

// MInternalMessageNoticeMgr open func
func MInternalMessageNoticeMgr(db *gorm.DB) *_MInternalMessageNoticeMgr {
	if db == nil {
		panic(fmt.Errorf("MInternalMessageNoticeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MInternalMessageNoticeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("m_internal_message_notice"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MInternalMessageNoticeMgr) GetTableName() string {
	return "m_internal_message_notice"
}

// Reset 重置gorm会话
func (obj *_MInternalMessageNoticeMgr) Reset() *_MInternalMessageNoticeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MInternalMessageNoticeMgr) Get() (result MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MInternalMessageNoticeMgr) Gets() (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MInternalMessageNoticeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MInternalMessageNoticeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNoticeName notice_name获取 通知名称
func (obj *_MInternalMessageNoticeMgr) WithNoticeName(noticeName string) Option {
	return optionFunc(func(o *options) { o.query["notice_name"] = noticeName })
}

// WithNoticeType notice_type获取 通知类型 0 超过24小时未打大包 1 超过24小时未生成总单
func (obj *_MInternalMessageNoticeMgr) WithNoticeType(noticeType int) Option {
	return optionFunc(func(o *options) { o.query["notice_type"] = noticeType })
}

// WithTemplateContent template_content获取 通知模板（消息内容)
func (obj *_MInternalMessageNoticeMgr) WithTemplateContent(templateContent string) Option {
	return optionFunc(func(o *options) { o.query["template_content"] = templateContent })
}

// WithIsEnabled is_enabled获取 是否启用 0 未启用 1 已启用
func (obj *_MInternalMessageNoticeMgr) WithIsEnabled(isEnabled int) Option {
	return optionFunc(func(o *options) { o.query["is_enabled"] = isEnabled })
}

// WithRecipient recipient获取 消息接收人id 用逗号分割
func (obj *_MInternalMessageNoticeMgr) WithRecipient(recipient string) Option {
	return optionFunc(func(o *options) { o.query["recipient"] = recipient })
}

// WithCreateTime create_time获取 创建时间
func (obj *_MInternalMessageNoticeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_MInternalMessageNoticeMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_MInternalMessageNoticeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_MInternalMessageNoticeMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_MInternalMessageNoticeMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_MInternalMessageNoticeMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_MInternalMessageNoticeMgr) GetByOption(opts ...Option) (result MInternalMessageNotice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MInternalMessageNoticeMgr) GetByOptions(opts ...Option) (results []*MInternalMessageNotice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MInternalMessageNoticeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MInternalMessageNotice, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where(options.query)
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
func (obj *_MInternalMessageNoticeMgr) GetFromID(id int) (result MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MInternalMessageNoticeMgr) GetBatchFromID(ids []int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromNoticeName 通过notice_name获取内容 通知名称
func (obj *_MInternalMessageNoticeMgr) GetFromNoticeName(noticeName string) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`notice_name` = ?", noticeName).Find(&results).Error

	return
}

// GetBatchFromNoticeName 批量查找 通知名称
func (obj *_MInternalMessageNoticeMgr) GetBatchFromNoticeName(noticeNames []string) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`notice_name` IN (?)", noticeNames).Find(&results).Error

	return
}

// GetFromNoticeType 通过notice_type获取内容 通知类型 0 超过24小时未打大包 1 超过24小时未生成总单
func (obj *_MInternalMessageNoticeMgr) GetFromNoticeType(noticeType int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`notice_type` = ?", noticeType).Find(&results).Error

	return
}

// GetBatchFromNoticeType 批量查找 通知类型 0 超过24小时未打大包 1 超过24小时未生成总单
func (obj *_MInternalMessageNoticeMgr) GetBatchFromNoticeType(noticeTypes []int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`notice_type` IN (?)", noticeTypes).Find(&results).Error

	return
}

// GetFromTemplateContent 通过template_content获取内容 通知模板（消息内容)
func (obj *_MInternalMessageNoticeMgr) GetFromTemplateContent(templateContent string) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`template_content` = ?", templateContent).Find(&results).Error

	return
}

// GetBatchFromTemplateContent 批量查找 通知模板（消息内容)
func (obj *_MInternalMessageNoticeMgr) GetBatchFromTemplateContent(templateContents []string) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`template_content` IN (?)", templateContents).Find(&results).Error

	return
}

// GetFromIsEnabled 通过is_enabled获取内容 是否启用 0 未启用 1 已启用
func (obj *_MInternalMessageNoticeMgr) GetFromIsEnabled(isEnabled int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`is_enabled` = ?", isEnabled).Find(&results).Error

	return
}

// GetBatchFromIsEnabled 批量查找 是否启用 0 未启用 1 已启用
func (obj *_MInternalMessageNoticeMgr) GetBatchFromIsEnabled(isEnableds []int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`is_enabled` IN (?)", isEnableds).Find(&results).Error

	return
}

// GetFromRecipient 通过recipient获取内容 消息接收人id 用逗号分割
func (obj *_MInternalMessageNoticeMgr) GetFromRecipient(recipient string) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`recipient` = ?", recipient).Find(&results).Error

	return
}

// GetBatchFromRecipient 批量查找 消息接收人id 用逗号分割
func (obj *_MInternalMessageNoticeMgr) GetBatchFromRecipient(recipients []string) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`recipient` IN (?)", recipients).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_MInternalMessageNoticeMgr) GetFromCreateTime(createTime time.Time) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_MInternalMessageNoticeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_MInternalMessageNoticeMgr) GetFromCreateUser(createUser int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_MInternalMessageNoticeMgr) GetBatchFromCreateUser(createUsers []int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_MInternalMessageNoticeMgr) GetFromUpdateTime(updateTime time.Time) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_MInternalMessageNoticeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_MInternalMessageNoticeMgr) GetFromUpdateUser(updateUser int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_MInternalMessageNoticeMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_MInternalMessageNoticeMgr) GetFromVersion(version int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_MInternalMessageNoticeMgr) GetBatchFromVersion(versions []int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_MInternalMessageNoticeMgr) GetFromDeleted(deleted int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_MInternalMessageNoticeMgr) GetBatchFromDeleted(deleteds []int) (results []*MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MInternalMessageNoticeMgr) FetchByPrimaryKey(id int) (result MInternalMessageNotice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MInternalMessageNotice{}).Where("`id` = ?", id).First(&result).Error

	return
}
