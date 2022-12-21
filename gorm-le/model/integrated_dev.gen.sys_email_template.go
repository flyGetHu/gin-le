package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysEmailTemplateMgr struct {
	*_BaseMgr
}

// SysEmailTemplateMgr open func
func SysEmailTemplateMgr(db *gorm.DB) *_SysEmailTemplateMgr {
	if db == nil {
		panic(fmt.Errorf("SysEmailTemplateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysEmailTemplateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_email_template"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysEmailTemplateMgr) GetTableName() string {
	return "sys_email_template"
}

// Reset 重置gorm会话
func (obj *_SysEmailTemplateMgr) Reset() *_SysEmailTemplateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SysEmailTemplateMgr) Get() (result SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysEmailTemplateMgr) Gets() (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SysEmailTemplateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_SysEmailTemplateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTemplateID template_id获取 模板id
func (obj *_SysEmailTemplateMgr) WithTemplateID(templateID int) Option {
	return optionFunc(func(o *options) { o.query["template_id"] = templateID })
}

// WithTemplateName template_name获取 模板名称
func (obj *_SysEmailTemplateMgr) WithTemplateName(templateName string) Option {
	return optionFunc(func(o *options) { o.query["template_name"] = templateName })
}

// WithSendNode send_node获取 发送节点
func (obj *_SysEmailTemplateMgr) WithSendNode(sendNode string) Option {
	return optionFunc(func(o *options) { o.query["send_node"] = sendNode })
}

// WithTemplateContent template_content获取 模板内容
func (obj *_SysEmailTemplateMgr) WithTemplateContent(templateContent string) Option {
	return optionFunc(func(o *options) { o.query["template_content"] = templateContent })
}

// WithStatus status获取 是否启用(0-否，1-启用)
func (obj *_SysEmailTemplateMgr) WithStatus(status []uint8) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysEmailTemplateMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_SysEmailTemplateMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysEmailTemplateMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_SysEmailTemplateMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_SysEmailTemplateMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_SysEmailTemplateMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithReviewStatus review_status获取 审核状态 1-审核中|0-已通过|2-拒绝|其它-不可用
func (obj *_SysEmailTemplateMgr) WithReviewStatus(reviewStatus int) Option {
	return optionFunc(func(o *options) { o.query["review_status"] = reviewStatus })
}

// WithReviewReason review_reason获取 审核原因 审核的失败原因
func (obj *_SysEmailTemplateMgr) WithReviewReason(reviewReason string) Option {
	return optionFunc(func(o *options) { o.query["review_reason"] = reviewReason })
}

// GetByOption 功能选项模式获取
func (obj *_SysEmailTemplateMgr) GetByOption(opts ...Option) (result SysEmailTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysEmailTemplateMgr) GetByOptions(opts ...Option) (results []*SysEmailTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SysEmailTemplateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SysEmailTemplate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where(options.query)
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
func (obj *_SysEmailTemplateMgr) GetFromID(id int) (result SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_SysEmailTemplateMgr) GetBatchFromID(ids []int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTemplateID 通过template_id获取内容 模板id
func (obj *_SysEmailTemplateMgr) GetFromTemplateID(templateID int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`template_id` = ?", templateID).Find(&results).Error

	return
}

// GetBatchFromTemplateID 批量查找 模板id
func (obj *_SysEmailTemplateMgr) GetBatchFromTemplateID(templateIDs []int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`template_id` IN (?)", templateIDs).Find(&results).Error

	return
}

// GetFromTemplateName 通过template_name获取内容 模板名称
func (obj *_SysEmailTemplateMgr) GetFromTemplateName(templateName string) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`template_name` = ?", templateName).Find(&results).Error

	return
}

// GetBatchFromTemplateName 批量查找 模板名称
func (obj *_SysEmailTemplateMgr) GetBatchFromTemplateName(templateNames []string) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`template_name` IN (?)", templateNames).Find(&results).Error

	return
}

// GetFromSendNode 通过send_node获取内容 发送节点
func (obj *_SysEmailTemplateMgr) GetFromSendNode(sendNode string) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`send_node` = ?", sendNode).Find(&results).Error

	return
}

// GetBatchFromSendNode 批量查找 发送节点
func (obj *_SysEmailTemplateMgr) GetBatchFromSendNode(sendNodes []string) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`send_node` IN (?)", sendNodes).Find(&results).Error

	return
}

// GetFromTemplateContent 通过template_content获取内容 模板内容
func (obj *_SysEmailTemplateMgr) GetFromTemplateContent(templateContent string) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`template_content` = ?", templateContent).Find(&results).Error

	return
}

// GetBatchFromTemplateContent 批量查找 模板内容
func (obj *_SysEmailTemplateMgr) GetBatchFromTemplateContent(templateContents []string) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`template_content` IN (?)", templateContents).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 是否启用(0-否，1-启用)
func (obj *_SysEmailTemplateMgr) GetFromStatus(status []uint8) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 是否启用(0-否，1-启用)
func (obj *_SysEmailTemplateMgr) GetBatchFromStatus(statuss [][]uint8) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysEmailTemplateMgr) GetFromCreateTime(createTime time.Time) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SysEmailTemplateMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_SysEmailTemplateMgr) GetFromCreateUser(createUser int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_SysEmailTemplateMgr) GetBatchFromCreateUser(createUsers []int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysEmailTemplateMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_SysEmailTemplateMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_SysEmailTemplateMgr) GetFromUpdateUser(updateUser int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_SysEmailTemplateMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_SysEmailTemplateMgr) GetFromVersion(version int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_SysEmailTemplateMgr) GetBatchFromVersion(versions []int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_SysEmailTemplateMgr) GetFromDeleted(deleted int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_SysEmailTemplateMgr) GetBatchFromDeleted(deleteds []int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromReviewStatus 通过review_status获取内容 审核状态 1-审核中|0-已通过|2-拒绝|其它-不可用
func (obj *_SysEmailTemplateMgr) GetFromReviewStatus(reviewStatus int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`review_status` = ?", reviewStatus).Find(&results).Error

	return
}

// GetBatchFromReviewStatus 批量查找 审核状态 1-审核中|0-已通过|2-拒绝|其它-不可用
func (obj *_SysEmailTemplateMgr) GetBatchFromReviewStatus(reviewStatuss []int) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`review_status` IN (?)", reviewStatuss).Find(&results).Error

	return
}

// GetFromReviewReason 通过review_reason获取内容 审核原因 审核的失败原因
func (obj *_SysEmailTemplateMgr) GetFromReviewReason(reviewReason string) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`review_reason` = ?", reviewReason).Find(&results).Error

	return
}

// GetBatchFromReviewReason 批量查找 审核原因 审核的失败原因
func (obj *_SysEmailTemplateMgr) GetBatchFromReviewReason(reviewReasons []string) (results []*SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`review_reason` IN (?)", reviewReasons).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SysEmailTemplateMgr) FetchByPrimaryKey(id int) (result SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueBySysEmailTemplateIDUIndex primary or index 获取唯一内容
func (obj *_SysEmailTemplateMgr) FetchUniqueBySysEmailTemplateIDUIndex(id int) (result SysEmailTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysEmailTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}
