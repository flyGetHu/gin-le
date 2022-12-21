package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysConfigMgr struct {
	*_BaseMgr
}

// SysConfigMgr open func
func SysConfigMgr(db *gorm.DB) *_SysConfigMgr {
	if db == nil {
		panic(fmt.Errorf("SysConfigMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_config"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysConfigMgr) GetTableName() string {
	return "sys_config"
}

// Reset 重置gorm会话
func (obj *_SysConfigMgr) Reset() *_SysConfigMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SysConfigMgr) Get() (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysConfigMgr) Gets() (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SysConfigMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_SysConfigMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithConfigKey config_key获取 配置key
func (obj *_SysConfigMgr) WithConfigKey(configKey string) Option {
	return optionFunc(func(o *options) { o.query["config_key"] = configKey })
}

// WithConfigValue config_value获取 配置值
func (obj *_SysConfigMgr) WithConfigValue(configValue string) Option {
	return optionFunc(func(o *options) { o.query["config_value"] = configValue })
}

// WithDefaultValue default_value获取 默认值
func (obj *_SysConfigMgr) WithDefaultValue(defaultValue string) Option {
	return optionFunc(func(o *options) { o.query["default_value"] = defaultValue })
}

// WithRemarks remarks获取 备注说明
func (obj *_SysConfigMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysConfigMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_SysConfigMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysConfigMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_SysConfigMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_SysConfigMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_SysConfigMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_SysConfigMgr) GetByOption(opts ...Option) (result SysConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysConfigMgr) GetByOptions(opts ...Option) (results []*SysConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SysConfigMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SysConfig, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where(options.query)
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
func (obj *_SysConfigMgr) GetFromID(id int) (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_SysConfigMgr) GetBatchFromID(ids []int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromConfigKey 通过config_key获取内容 配置key
func (obj *_SysConfigMgr) GetFromConfigKey(configKey string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`config_key` = ?", configKey).Find(&results).Error

	return
}

// GetBatchFromConfigKey 批量查找 配置key
func (obj *_SysConfigMgr) GetBatchFromConfigKey(configKeys []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`config_key` IN (?)", configKeys).Find(&results).Error

	return
}

// GetFromConfigValue 通过config_value获取内容 配置值
func (obj *_SysConfigMgr) GetFromConfigValue(configValue string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`config_value` = ?", configValue).Find(&results).Error

	return
}

// GetBatchFromConfigValue 批量查找 配置值
func (obj *_SysConfigMgr) GetBatchFromConfigValue(configValues []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`config_value` IN (?)", configValues).Find(&results).Error

	return
}

// GetFromDefaultValue 通过default_value获取内容 默认值
func (obj *_SysConfigMgr) GetFromDefaultValue(defaultValue string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`default_value` = ?", defaultValue).Find(&results).Error

	return
}

// GetBatchFromDefaultValue 批量查找 默认值
func (obj *_SysConfigMgr) GetBatchFromDefaultValue(defaultValues []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`default_value` IN (?)", defaultValues).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注说明
func (obj *_SysConfigMgr) GetFromRemarks(remarks string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注说明
func (obj *_SysConfigMgr) GetBatchFromRemarks(remarkss []string) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysConfigMgr) GetFromCreateTime(createTime time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SysConfigMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_SysConfigMgr) GetFromCreateUser(createUser int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_SysConfigMgr) GetBatchFromCreateUser(createUsers []int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysConfigMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_SysConfigMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_SysConfigMgr) GetFromUpdateUser(updateUser int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_SysConfigMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_SysConfigMgr) GetFromVersion(version int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_SysConfigMgr) GetBatchFromVersion(versions []int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_SysConfigMgr) GetFromDeleted(deleted int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_SysConfigMgr) GetBatchFromDeleted(deleteds []int) (results []*SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SysConfigMgr) FetchByPrimaryKey(id int) (result SysConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysConfig{}).Where("`id` = ?", id).First(&result).Error

	return
}
