package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgProviderChannelCheckMgr struct {
	*_BaseMgr
}

// LgProviderChannelCheckMgr open func
func LgProviderChannelCheckMgr(db *gorm.DB) *_LgProviderChannelCheckMgr {
	if db == nil {
		panic(fmt.Errorf("LgProviderChannelCheckMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgProviderChannelCheckMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_provider_channel_check"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgProviderChannelCheckMgr) GetTableName() string {
	return "lg_provider_channel_check"
}

// Reset 重置gorm会话
func (obj *_LgProviderChannelCheckMgr) Reset() *_LgProviderChannelCheckMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgProviderChannelCheckMgr) Get() (result LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgProviderChannelCheckMgr) Gets() (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgProviderChannelCheckMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgProviderChannelCheckMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithChannelFieldID channel_field_id获取 渠道字段id
func (obj *_LgProviderChannelCheckMgr) WithChannelFieldID(channelFieldID int) Option {
	return optionFunc(func(o *options) { o.query["channel_field_id"] = channelFieldID })
}

// WithIsRequired is_required获取 是否必填
func (obj *_LgProviderChannelCheckMgr) WithIsRequired(isRequired int) Option {
	return optionFunc(func(o *options) { o.query["is_required"] = isRequired })
}

// WithRule rule获取 限制规则
func (obj *_LgProviderChannelCheckMgr) WithRule(rule string) Option {
	return optionFunc(func(o *options) { o.query["rule"] = rule })
}

// WithErrorMessage error_message获取 错误消息提示
func (obj *_LgProviderChannelCheckMgr) WithErrorMessage(errorMessage string) Option {
	return optionFunc(func(o *options) { o.query["error_message"] = errorMessage })
}

// WithRemarks remarks获取 备注
func (obj *_LgProviderChannelCheckMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgProviderChannelCheckMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgProviderChannelCheckMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgProviderChannelCheckMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgProviderChannelCheckMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgProviderChannelCheckMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgProviderChannelCheckMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgProviderChannelCheckMgr) GetByOption(opts ...Option) (result LgProviderChannelCheck, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgProviderChannelCheckMgr) GetByOptions(opts ...Option) (results []*LgProviderChannelCheck, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgProviderChannelCheckMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgProviderChannelCheck, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where(options.query)
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

// GetFromID 通过id获取内容 ID
func (obj *_LgProviderChannelCheckMgr) GetFromID(id int) (result LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgProviderChannelCheckMgr) GetBatchFromID(ids []int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromChannelFieldID 通过channel_field_id获取内容 渠道字段id
func (obj *_LgProviderChannelCheckMgr) GetFromChannelFieldID(channelFieldID int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`channel_field_id` = ?", channelFieldID).Find(&results).Error

	return
}

// GetBatchFromChannelFieldID 批量查找 渠道字段id
func (obj *_LgProviderChannelCheckMgr) GetBatchFromChannelFieldID(channelFieldIDs []int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`channel_field_id` IN (?)", channelFieldIDs).Find(&results).Error

	return
}

// GetFromIsRequired 通过is_required获取内容 是否必填
func (obj *_LgProviderChannelCheckMgr) GetFromIsRequired(isRequired int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`is_required` = ?", isRequired).Find(&results).Error

	return
}

// GetBatchFromIsRequired 批量查找 是否必填
func (obj *_LgProviderChannelCheckMgr) GetBatchFromIsRequired(isRequireds []int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`is_required` IN (?)", isRequireds).Find(&results).Error

	return
}

// GetFromRule 通过rule获取内容 限制规则
func (obj *_LgProviderChannelCheckMgr) GetFromRule(rule string) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`rule` = ?", rule).Find(&results).Error

	return
}

// GetBatchFromRule 批量查找 限制规则
func (obj *_LgProviderChannelCheckMgr) GetBatchFromRule(rules []string) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`rule` IN (?)", rules).Find(&results).Error

	return
}

// GetFromErrorMessage 通过error_message获取内容 错误消息提示
func (obj *_LgProviderChannelCheckMgr) GetFromErrorMessage(errorMessage string) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`error_message` = ?", errorMessage).Find(&results).Error

	return
}

// GetBatchFromErrorMessage 批量查找 错误消息提示
func (obj *_LgProviderChannelCheckMgr) GetBatchFromErrorMessage(errorMessages []string) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`error_message` IN (?)", errorMessages).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注
func (obj *_LgProviderChannelCheckMgr) GetFromRemarks(remarks string) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注
func (obj *_LgProviderChannelCheckMgr) GetBatchFromRemarks(remarkss []string) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgProviderChannelCheckMgr) GetFromCreateTime(createTime time.Time) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgProviderChannelCheckMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgProviderChannelCheckMgr) GetFromCreateUser(createUser int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgProviderChannelCheckMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgProviderChannelCheckMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgProviderChannelCheckMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgProviderChannelCheckMgr) GetFromUpdateUser(updateUser int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgProviderChannelCheckMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgProviderChannelCheckMgr) GetFromVersion(version int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgProviderChannelCheckMgr) GetBatchFromVersion(versions []int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgProviderChannelCheckMgr) GetFromDeleted(deleted int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgProviderChannelCheckMgr) GetBatchFromDeleted(deleteds []int) (results []*LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgProviderChannelCheckMgr) FetchByPrimaryKey(id int) (result LgProviderChannelCheck, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelCheck{}).Where("`id` = ?", id).First(&result).Error

	return
}
