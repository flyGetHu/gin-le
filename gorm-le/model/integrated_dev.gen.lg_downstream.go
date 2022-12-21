package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgDownstreamMgr struct {
	*_BaseMgr
}

// LgDownstreamMgr open func
func LgDownstreamMgr(db *gorm.DB) *_LgDownstreamMgr {
	if db == nil {
		panic(fmt.Errorf("LgDownstreamMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgDownstreamMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_downstream"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgDownstreamMgr) GetTableName() string {
	return "lg_downstream"
}

// Reset 重置gorm会话
func (obj *_LgDownstreamMgr) Reset() *_LgDownstreamMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgDownstreamMgr) Get() (result LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgDownstreamMgr) Gets() (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgDownstreamMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgDownstreamMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDownstreamCode downstream_code获取 下家code
func (obj *_LgDownstreamMgr) WithDownstreamCode(downstreamCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_code"] = downstreamCode })
}

// WithDownstreamName downstream_name获取 下家名称
func (obj *_LgDownstreamMgr) WithDownstreamName(downstreamName string) Option {
	return optionFunc(func(o *options) { o.query["downstream_name"] = downstreamName })
}

// WithDownstreamProviderCode downstream_provider_code获取 下家服务商代码
func (obj *_LgDownstreamMgr) WithDownstreamProviderCode(downstreamProviderCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_provider_code"] = downstreamProviderCode })
}

// WithForecastTemplateCode forecast_template_code获取 预报模板code
func (obj *_LgDownstreamMgr) WithForecastTemplateCode(forecastTemplateCode string) Option {
	return optionFunc(func(o *options) { o.query["forecast_template_code"] = forecastTemplateCode })
}

// WithVariableList variable_list获取 变量列表
func (obj *_LgDownstreamMgr) WithVariableList(variableList string) Option {
	return optionFunc(func(o *options) { o.query["variable_list"] = variableList })
}

// WithVariableKeyValueList variable_key_value_list获取 变量key-value列表
func (obj *_LgDownstreamMgr) WithVariableKeyValueList(variableKeyValueList string) Option {
	return optionFunc(func(o *options) { o.query["variable_key_value_list"] = variableKeyValueList })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgDownstreamMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgDownstreamMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgDownstreamMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgDownstreamMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgDownstreamMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgDownstreamMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgDownstreamMgr) GetByOption(opts ...Option) (result LgDownstream, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgDownstreamMgr) GetByOptions(opts ...Option) (results []*LgDownstream, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgDownstreamMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgDownstream, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where(options.query)
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
func (obj *_LgDownstreamMgr) GetFromID(id int) (result LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LgDownstreamMgr) GetBatchFromID(ids []int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDownstreamCode 通过downstream_code获取内容 下家code
func (obj *_LgDownstreamMgr) GetFromDownstreamCode(downstreamCode string) (result LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_code` = ?", downstreamCode).First(&result).Error

	return
}

// GetBatchFromDownstreamCode 批量查找 下家code
func (obj *_LgDownstreamMgr) GetBatchFromDownstreamCode(downstreamCodes []string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_code` IN (?)", downstreamCodes).Find(&results).Error

	return
}

// GetFromDownstreamName 通过downstream_name获取内容 下家名称
func (obj *_LgDownstreamMgr) GetFromDownstreamName(downstreamName string) (result LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_name` = ?", downstreamName).First(&result).Error

	return
}

// GetBatchFromDownstreamName 批量查找 下家名称
func (obj *_LgDownstreamMgr) GetBatchFromDownstreamName(downstreamNames []string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_name` IN (?)", downstreamNames).Find(&results).Error

	return
}

// GetFromDownstreamProviderCode 通过downstream_provider_code获取内容 下家服务商代码
func (obj *_LgDownstreamMgr) GetFromDownstreamProviderCode(downstreamProviderCode string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_provider_code` = ?", downstreamProviderCode).Find(&results).Error

	return
}

// GetBatchFromDownstreamProviderCode 批量查找 下家服务商代码
func (obj *_LgDownstreamMgr) GetBatchFromDownstreamProviderCode(downstreamProviderCodes []string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_provider_code` IN (?)", downstreamProviderCodes).Find(&results).Error

	return
}

// GetFromForecastTemplateCode 通过forecast_template_code获取内容 预报模板code
func (obj *_LgDownstreamMgr) GetFromForecastTemplateCode(forecastTemplateCode string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`forecast_template_code` = ?", forecastTemplateCode).Find(&results).Error

	return
}

// GetBatchFromForecastTemplateCode 批量查找 预报模板code
func (obj *_LgDownstreamMgr) GetBatchFromForecastTemplateCode(forecastTemplateCodes []string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`forecast_template_code` IN (?)", forecastTemplateCodes).Find(&results).Error

	return
}

// GetFromVariableList 通过variable_list获取内容 变量列表
func (obj *_LgDownstreamMgr) GetFromVariableList(variableList string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`variable_list` = ?", variableList).Find(&results).Error

	return
}

// GetBatchFromVariableList 批量查找 变量列表
func (obj *_LgDownstreamMgr) GetBatchFromVariableList(variableLists []string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`variable_list` IN (?)", variableLists).Find(&results).Error

	return
}

// GetFromVariableKeyValueList 通过variable_key_value_list获取内容 变量key-value列表
func (obj *_LgDownstreamMgr) GetFromVariableKeyValueList(variableKeyValueList string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`variable_key_value_list` = ?", variableKeyValueList).Find(&results).Error

	return
}

// GetBatchFromVariableKeyValueList 批量查找 变量key-value列表
func (obj *_LgDownstreamMgr) GetBatchFromVariableKeyValueList(variableKeyValueLists []string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`variable_key_value_list` IN (?)", variableKeyValueLists).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgDownstreamMgr) GetFromCreateTime(createTime time.Time) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgDownstreamMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgDownstreamMgr) GetFromCreateUser(createUser int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgDownstreamMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgDownstreamMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgDownstreamMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgDownstreamMgr) GetFromUpdateUser(updateUser int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgDownstreamMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgDownstreamMgr) GetFromVersion(version int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgDownstreamMgr) GetBatchFromVersion(versions []int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgDownstreamMgr) GetFromDeleted(deleted int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgDownstreamMgr) GetBatchFromDeleted(deleteds []int) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgDownstreamMgr) FetchByPrimaryKey(id int) (result LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexDownstreamCode primary or index 获取唯一内容
func (obj *_LgDownstreamMgr) FetchUniqueByIndexDownstreamCode(downstreamCode string) (result LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_code` = ?", downstreamCode).First(&result).Error

	return
}

// FetchUniqueByIndexDownstreamName primary or index 获取唯一内容
func (obj *_LgDownstreamMgr) FetchUniqueByIndexDownstreamName(downstreamName string) (result LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_name` = ?", downstreamName).First(&result).Error

	return
}

// FetchIndexByDownstreamProviderCodeIndex  获取多个内容
func (obj *_LgDownstreamMgr) FetchIndexByDownstreamProviderCodeIndex(downstreamProviderCode string) (results []*LgDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDownstream{}).Where("`downstream_provider_code` = ?", downstreamProviderCode).Find(&results).Error

	return
}
