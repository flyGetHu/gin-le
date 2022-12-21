package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgProviderChannelFieldMgr struct {
	*_BaseMgr
}

// LgProviderChannelFieldMgr open func
func LgProviderChannelFieldMgr(db *gorm.DB) *_LgProviderChannelFieldMgr {
	if db == nil {
		panic(fmt.Errorf("LgProviderChannelFieldMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgProviderChannelFieldMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_provider_channel_field"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgProviderChannelFieldMgr) GetTableName() string {
	return "lg_provider_channel_field"
}

// Reset 重置gorm会话
func (obj *_LgProviderChannelFieldMgr) Reset() *_LgProviderChannelFieldMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgProviderChannelFieldMgr) Get() (result LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgProviderChannelFieldMgr) Gets() (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgProviderChannelFieldMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgProviderChannelFieldMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 所属类型
func (obj *_LgProviderChannelFieldMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithName name获取 字段名
func (obj *_LgProviderChannelFieldMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithNameCn name_cn获取 字段描述
func (obj *_LgProviderChannelFieldMgr) WithNameCn(nameCn string) Option {
	return optionFunc(func(o *options) { o.query["name_cn"] = nameCn })
}

// WithRemarks remarks获取 备注
func (obj *_LgProviderChannelFieldMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgProviderChannelFieldMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgProviderChannelFieldMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgProviderChannelFieldMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgProviderChannelFieldMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgProviderChannelFieldMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgProviderChannelFieldMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgProviderChannelFieldMgr) GetByOption(opts ...Option) (result LgProviderChannelField, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgProviderChannelFieldMgr) GetByOptions(opts ...Option) (results []*LgProviderChannelField, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgProviderChannelFieldMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgProviderChannelField, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where(options.query)
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
func (obj *_LgProviderChannelFieldMgr) GetFromID(id int) (result LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgProviderChannelFieldMgr) GetBatchFromID(ids []int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 所属类型
func (obj *_LgProviderChannelFieldMgr) GetFromType(_type string) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 所属类型
func (obj *_LgProviderChannelFieldMgr) GetBatchFromType(_types []string) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 字段名
func (obj *_LgProviderChannelFieldMgr) GetFromName(name string) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 字段名
func (obj *_LgProviderChannelFieldMgr) GetBatchFromName(names []string) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromNameCn 通过name_cn获取内容 字段描述
func (obj *_LgProviderChannelFieldMgr) GetFromNameCn(nameCn string) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`name_cn` = ?", nameCn).Find(&results).Error

	return
}

// GetBatchFromNameCn 批量查找 字段描述
func (obj *_LgProviderChannelFieldMgr) GetBatchFromNameCn(nameCns []string) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`name_cn` IN (?)", nameCns).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注
func (obj *_LgProviderChannelFieldMgr) GetFromRemarks(remarks string) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注
func (obj *_LgProviderChannelFieldMgr) GetBatchFromRemarks(remarkss []string) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgProviderChannelFieldMgr) GetFromCreateTime(createTime time.Time) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgProviderChannelFieldMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgProviderChannelFieldMgr) GetFromCreateUser(createUser int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgProviderChannelFieldMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgProviderChannelFieldMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgProviderChannelFieldMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgProviderChannelFieldMgr) GetFromUpdateUser(updateUser int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgProviderChannelFieldMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgProviderChannelFieldMgr) GetFromVersion(version int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgProviderChannelFieldMgr) GetBatchFromVersion(versions []int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgProviderChannelFieldMgr) GetFromDeleted(deleted int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgProviderChannelFieldMgr) GetBatchFromDeleted(deleteds []int) (results []*LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgProviderChannelFieldMgr) FetchByPrimaryKey(id int) (result LgProviderChannelField, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannelField{}).Where("`id` = ?", id).First(&result).Error

	return
}
