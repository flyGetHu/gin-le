package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysTemplateFieldMgr struct {
	*_BaseMgr
}

// SysTemplateFieldMgr open func
func SysTemplateFieldMgr(db *gorm.DB) *_SysTemplateFieldMgr {
	if db == nil {
		panic(fmt.Errorf("SysTemplateFieldMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysTemplateFieldMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_template_field"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysTemplateFieldMgr) GetTableName() string {
	return "sys_template_field"
}

// Reset 重置gorm会话
func (obj *_SysTemplateFieldMgr) Reset() *_SysTemplateFieldMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SysTemplateFieldMgr) Get() (result SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysTemplateFieldMgr) Gets() (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SysTemplateFieldMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SysTemplateFieldMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithModuleName module_name获取 模块
func (obj *_SysTemplateFieldMgr) WithModuleName(moduleName string) Option {
	return optionFunc(func(o *options) { o.query["module_name"] = moduleName })
}

// WithFieldName field_name获取 字段名
func (obj *_SysTemplateFieldMgr) WithFieldName(fieldName string) Option {
	return optionFunc(func(o *options) { o.query["field_name"] = fieldName })
}

// WithFieldChinese field_chinese获取 中文
func (obj *_SysTemplateFieldMgr) WithFieldChinese(fieldChinese string) Option {
	return optionFunc(func(o *options) { o.query["field_chinese"] = fieldChinese })
}

// WithCreateTime create_time获取
func (obj *_SysTemplateFieldMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_SysTemplateFieldMgr) GetByOption(opts ...Option) (result SysTemplateField, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysTemplateFieldMgr) GetByOptions(opts ...Option) (results []*SysTemplateField, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SysTemplateFieldMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SysTemplateField, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where(options.query)
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
func (obj *_SysTemplateFieldMgr) GetFromID(id int) (result SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SysTemplateFieldMgr) GetBatchFromID(ids []int) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromModuleName 通过module_name获取内容 模块
func (obj *_SysTemplateFieldMgr) GetFromModuleName(moduleName string) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`module_name` = ?", moduleName).Find(&results).Error

	return
}

// GetBatchFromModuleName 批量查找 模块
func (obj *_SysTemplateFieldMgr) GetBatchFromModuleName(moduleNames []string) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`module_name` IN (?)", moduleNames).Find(&results).Error

	return
}

// GetFromFieldName 通过field_name获取内容 字段名
func (obj *_SysTemplateFieldMgr) GetFromFieldName(fieldName string) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`field_name` = ?", fieldName).Find(&results).Error

	return
}

// GetBatchFromFieldName 批量查找 字段名
func (obj *_SysTemplateFieldMgr) GetBatchFromFieldName(fieldNames []string) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`field_name` IN (?)", fieldNames).Find(&results).Error

	return
}

// GetFromFieldChinese 通过field_chinese获取内容 中文
func (obj *_SysTemplateFieldMgr) GetFromFieldChinese(fieldChinese string) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`field_chinese` = ?", fieldChinese).Find(&results).Error

	return
}

// GetBatchFromFieldChinese 批量查找 中文
func (obj *_SysTemplateFieldMgr) GetBatchFromFieldChinese(fieldChineses []string) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`field_chinese` IN (?)", fieldChineses).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *_SysTemplateFieldMgr) GetFromCreateTime(createTime time.Time) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *_SysTemplateFieldMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SysTemplateFieldMgr) FetchByPrimaryKey(id int) (result SysTemplateField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysTemplateField{}).Where("`id` = ?", id).First(&result).Error

	return
}
