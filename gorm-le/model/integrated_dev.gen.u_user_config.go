package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UUserConfigMgr struct {
	*_BaseMgr
}

// UUserConfigMgr open func
func UUserConfigMgr(db *gorm.DB) *_UUserConfigMgr {
	if db == nil {
		panic(fmt.Errorf("UUserConfigMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UUserConfigMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_user_config"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UUserConfigMgr) GetTableName() string {
	return "u_user_config"
}

// Reset 重置gorm会话
func (obj *_UUserConfigMgr) Reset() *_UUserConfigMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UUserConfigMgr) Get() (result UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UUserConfigMgr) Gets() (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UUserConfigMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_UUserConfigMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户id
func (obj *_UUserConfigMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithConfigKey config_key获取 配置key
func (obj *_UUserConfigMgr) WithConfigKey(configKey string) Option {
	return optionFunc(func(o *options) { o.query["config_key"] = configKey })
}

// WithConfigValue config_value获取 配置值
func (obj *_UUserConfigMgr) WithConfigValue(configValue string) Option {
	return optionFunc(func(o *options) { o.query["config_value"] = configValue })
}

// WithDefaultValue default_value获取 默认值
func (obj *_UUserConfigMgr) WithDefaultValue(defaultValue string) Option {
	return optionFunc(func(o *options) { o.query["default_value"] = defaultValue })
}

// WithRemarks remarks获取 备注说明
func (obj *_UUserConfigMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UUserConfigMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UUserConfigMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UUserConfigMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_UUserConfigMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UUserConfigMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_UUserConfigMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_UUserConfigMgr) GetByOption(opts ...Option) (result UUserConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UUserConfigMgr) GetByOptions(opts ...Option) (results []*UUserConfig, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UUserConfigMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UUserConfig, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where(options.query)
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
func (obj *_UUserConfigMgr) GetFromID(id int) (result UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_UUserConfigMgr) GetBatchFromID(ids []int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户id
func (obj *_UUserConfigMgr) GetFromUserID(userID int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户id
func (obj *_UUserConfigMgr) GetBatchFromUserID(userIDs []int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromConfigKey 通过config_key获取内容 配置key
func (obj *_UUserConfigMgr) GetFromConfigKey(configKey string) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`config_key` = ?", configKey).Find(&results).Error

	return
}

// GetBatchFromConfigKey 批量查找 配置key
func (obj *_UUserConfigMgr) GetBatchFromConfigKey(configKeys []string) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`config_key` IN (?)", configKeys).Find(&results).Error

	return
}

// GetFromConfigValue 通过config_value获取内容 配置值
func (obj *_UUserConfigMgr) GetFromConfigValue(configValue string) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`config_value` = ?", configValue).Find(&results).Error

	return
}

// GetBatchFromConfigValue 批量查找 配置值
func (obj *_UUserConfigMgr) GetBatchFromConfigValue(configValues []string) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`config_value` IN (?)", configValues).Find(&results).Error

	return
}

// GetFromDefaultValue 通过default_value获取内容 默认值
func (obj *_UUserConfigMgr) GetFromDefaultValue(defaultValue string) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`default_value` = ?", defaultValue).Find(&results).Error

	return
}

// GetBatchFromDefaultValue 批量查找 默认值
func (obj *_UUserConfigMgr) GetBatchFromDefaultValue(defaultValues []string) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`default_value` IN (?)", defaultValues).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注说明
func (obj *_UUserConfigMgr) GetFromRemarks(remarks string) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注说明
func (obj *_UUserConfigMgr) GetBatchFromRemarks(remarkss []string) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UUserConfigMgr) GetFromCreateTime(createTime time.Time) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UUserConfigMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UUserConfigMgr) GetFromCreateUser(createUser int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UUserConfigMgr) GetBatchFromCreateUser(createUsers []int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UUserConfigMgr) GetFromUpdateTime(updateTime time.Time) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UUserConfigMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_UUserConfigMgr) GetFromUpdateUser(updateUser int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_UUserConfigMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UUserConfigMgr) GetFromVersion(version int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UUserConfigMgr) GetBatchFromVersion(versions []int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_UUserConfigMgr) GetFromDeleted(deleted int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_UUserConfigMgr) GetBatchFromDeleted(deleteds []int) (results []*UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UUserConfigMgr) FetchByPrimaryKey(id int) (result UUserConfig, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserConfig{}).Where("`id` = ?", id).First(&result).Error

	return
}
