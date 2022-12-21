package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderCustomTemplateExportMgr struct {
	*_BaseMgr
}

// LgOrderCustomTemplateExportMgr open func
func LgOrderCustomTemplateExportMgr(db *gorm.DB) *_LgOrderCustomTemplateExportMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderCustomTemplateExportMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderCustomTemplateExportMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_custom_template_export"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderCustomTemplateExportMgr) GetTableName() string {
	return "lg_order_custom_template_export"
}

// Reset 重置gorm会话
func (obj *_LgOrderCustomTemplateExportMgr) Reset() *_LgOrderCustomTemplateExportMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderCustomTemplateExportMgr) Get() (result LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderCustomTemplateExportMgr) Gets() (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderCustomTemplateExportMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgOrderCustomTemplateExportMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_LgOrderCustomTemplateExportMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithCustomerID customer_id获取 客户id
func (obj *_LgOrderCustomTemplateExportMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithTemplateName template_name获取 模板名称
func (obj *_LgOrderCustomTemplateExportMgr) WithTemplateName(templateName string) Option {
	return optionFunc(func(o *options) { o.query["template_name"] = templateName })
}

// WithTemplateField template_field获取 自定义字段模板
func (obj *_LgOrderCustomTemplateExportMgr) WithTemplateField(templateField string) Option {
	return optionFunc(func(o *options) { o.query["template_field"] = templateField })
}

// WithType type获取 模板类型（0：页面模板，1：导出模板，3：客户顶带你自定义模板）
func (obj *_LgOrderCustomTemplateExportMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithRemarks remarks获取 备注信息
func (obj *_LgOrderCustomTemplateExportMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderCustomTemplateExportMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderCustomTemplateExportMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderCustomTemplateExportMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderCustomTemplateExportMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderCustomTemplateExportMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithIsDelete is_delete获取 逻辑删除
func (obj *_LgOrderCustomTemplateExportMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderCustomTemplateExportMgr) GetByOption(opts ...Option) (result LgOrderCustomTemplateExport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderCustomTemplateExportMgr) GetByOptions(opts ...Option) (results []*LgOrderCustomTemplateExport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderCustomTemplateExportMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderCustomTemplateExport, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where(options.query)
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
func (obj *_LgOrderCustomTemplateExportMgr) GetFromID(id int) (result LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromID(ids []int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_LgOrderCustomTemplateExportMgr) GetFromUserID(userID int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户id
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromUserID(userIDs []int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_LgOrderCustomTemplateExportMgr) GetFromCustomerID(customerID int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromCustomerID(customerIDs []int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromTemplateName 通过template_name获取内容 模板名称
func (obj *_LgOrderCustomTemplateExportMgr) GetFromTemplateName(templateName string) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`template_name` = ?", templateName).Find(&results).Error

	return
}

// GetBatchFromTemplateName 批量查找 模板名称
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromTemplateName(templateNames []string) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`template_name` IN (?)", templateNames).Find(&results).Error

	return
}

// GetFromTemplateField 通过template_field获取内容 自定义字段模板
func (obj *_LgOrderCustomTemplateExportMgr) GetFromTemplateField(templateField string) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`template_field` = ?", templateField).Find(&results).Error

	return
}

// GetBatchFromTemplateField 批量查找 自定义字段模板
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromTemplateField(templateFields []string) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`template_field` IN (?)", templateFields).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 模板类型（0：页面模板，1：导出模板，3：客户顶带你自定义模板）
func (obj *_LgOrderCustomTemplateExportMgr) GetFromType(_type int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 模板类型（0：页面模板，1：导出模板，3：客户顶带你自定义模板）
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromType(_types []int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注信息
func (obj *_LgOrderCustomTemplateExportMgr) GetFromRemarks(remarks string) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注信息
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromRemarks(remarkss []string) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderCustomTemplateExportMgr) GetFromCreateUser(createUser int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderCustomTemplateExportMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderCustomTemplateExportMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderCustomTemplateExportMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderCustomTemplateExportMgr) GetFromVersion(version int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromVersion(versions []int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 逻辑删除
func (obj *_LgOrderCustomTemplateExportMgr) GetFromIsDelete(isDelete int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 逻辑删除
func (obj *_LgOrderCustomTemplateExportMgr) GetBatchFromIsDelete(isDeletes []int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderCustomTemplateExportMgr) FetchByPrimaryKey(id int) (result LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByUserID  获取多个内容
func (obj *_LgOrderCustomTemplateExportMgr) FetchIndexByUserID(userID int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// FetchIndexByLgOrderCustomTemplateExportCustomerIDIndex  获取多个内容
func (obj *_LgOrderCustomTemplateExportMgr) FetchIndexByLgOrderCustomTemplateExportCustomerIDIndex(customerID int) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexByTemplateName  获取多个内容
func (obj *_LgOrderCustomTemplateExportMgr) FetchIndexByTemplateName(templateName string) (results []*LgOrderCustomTemplateExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderCustomTemplateExport{}).Where("`template_name` = ?", templateName).Find(&results).Error

	return
}
