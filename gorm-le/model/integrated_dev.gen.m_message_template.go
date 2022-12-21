package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MMessageTemplateMgr struct {
	*_BaseMgr
}

// MMessageTemplateMgr open func
func MMessageTemplateMgr(db *gorm.DB) *_MMessageTemplateMgr {
	if db == nil {
		panic(fmt.Errorf("MMessageTemplateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MMessageTemplateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("m_message_template"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MMessageTemplateMgr) GetTableName() string {
	return "m_message_template"
}

// Reset 重置gorm会话
func (obj *_MMessageTemplateMgr) Reset() *_MMessageTemplateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MMessageTemplateMgr) Get() (result MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MMessageTemplateMgr) Gets() (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MMessageTemplateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MMessageTemplateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 模板名称
func (obj *_MMessageTemplateMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithModuleName module_name获取 所属模块
func (obj *_MMessageTemplateMgr) WithModuleName(moduleName string) Option {
	return optionFunc(func(o *options) { o.query["module_name"] = moduleName })
}

// WithTemplate template获取 模板
func (obj *_MMessageTemplateMgr) WithTemplate(template string) Option {
	return optionFunc(func(o *options) { o.query["template"] = template })
}

// WithIsPrivate is_private获取 是否私有
func (obj *_MMessageTemplateMgr) WithIsPrivate(isPrivate int) Option {
	return optionFunc(func(o *options) { o.query["is_private"] = isPrivate })
}

// WithRemark remark获取 备注
func (obj *_MMessageTemplateMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_MMessageTemplateMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_MMessageTemplateMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_MMessageTemplateMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_MMessageTemplateMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_MMessageTemplateMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_MMessageTemplateMgr) GetByOption(opts ...Option) (result MMessageTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MMessageTemplateMgr) GetByOptions(opts ...Option) (results []*MMessageTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MMessageTemplateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MMessageTemplate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where(options.query)
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
func (obj *_MMessageTemplateMgr) GetFromID(id int) (result MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MMessageTemplateMgr) GetBatchFromID(ids []int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 模板名称
func (obj *_MMessageTemplateMgr) GetFromName(name string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 模板名称
func (obj *_MMessageTemplateMgr) GetBatchFromName(names []string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromModuleName 通过module_name获取内容 所属模块
func (obj *_MMessageTemplateMgr) GetFromModuleName(moduleName string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`module_name` = ?", moduleName).Find(&results).Error

	return
}

// GetBatchFromModuleName 批量查找 所属模块
func (obj *_MMessageTemplateMgr) GetBatchFromModuleName(moduleNames []string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`module_name` IN (?)", moduleNames).Find(&results).Error

	return
}

// GetFromTemplate 通过template获取内容 模板
func (obj *_MMessageTemplateMgr) GetFromTemplate(template string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`template` = ?", template).Find(&results).Error

	return
}

// GetBatchFromTemplate 批量查找 模板
func (obj *_MMessageTemplateMgr) GetBatchFromTemplate(templates []string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`template` IN (?)", templates).Find(&results).Error

	return
}

// GetFromIsPrivate 通过is_private获取内容 是否私有
func (obj *_MMessageTemplateMgr) GetFromIsPrivate(isPrivate int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`is_private` = ?", isPrivate).Find(&results).Error

	return
}

// GetBatchFromIsPrivate 批量查找 是否私有
func (obj *_MMessageTemplateMgr) GetBatchFromIsPrivate(isPrivates []int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`is_private` IN (?)", isPrivates).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_MMessageTemplateMgr) GetFromRemark(remark string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_MMessageTemplateMgr) GetBatchFromRemark(remarks []string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_MMessageTemplateMgr) GetFromCreateTime(createTime time.Time) (result MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`create_time` = ?", createTime).First(&result).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_MMessageTemplateMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_MMessageTemplateMgr) GetFromCreateUser(createUser int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_MMessageTemplateMgr) GetBatchFromCreateUser(createUsers []int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_MMessageTemplateMgr) GetFromUpdateTime(updateTime time.Time) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_MMessageTemplateMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_MMessageTemplateMgr) GetFromUpdateUser(updateUser int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_MMessageTemplateMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_MMessageTemplateMgr) GetFromVersion(version int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_MMessageTemplateMgr) GetBatchFromVersion(versions []int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MMessageTemplateMgr) FetchByPrimaryKey(id int) (result MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByMMessageTemplatePk primary or index 获取唯一内容
func (obj *_MMessageTemplateMgr) FetchUniqueByMMessageTemplatePk(createTime time.Time) (result MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`create_time` = ?", createTime).First(&result).Error

	return
}

// FetchIndexByMMessageTemplateNameIndex  获取多个内容
func (obj *_MMessageTemplateMgr) FetchIndexByMMessageTemplateNameIndex(name string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// FetchIndexByMMessageTemplateModuleNameIndex  获取多个内容
func (obj *_MMessageTemplateMgr) FetchIndexByMMessageTemplateModuleNameIndex(moduleName string) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`module_name` = ?", moduleName).Find(&results).Error

	return
}

// FetchIndexByMMessageTemplateIsPrivateIndex  获取多个内容
func (obj *_MMessageTemplateMgr) FetchIndexByMMessageTemplateIsPrivateIndex(isPrivate int) (results []*MMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageTemplate{}).Where("`is_private` = ?", isPrivate).Find(&results).Error

	return
}
