package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysCheckTheInformationMgr struct {
	*_BaseMgr
}

// SysCheckTheInformationMgr open func
func SysCheckTheInformationMgr(db *gorm.DB) *_SysCheckTheInformationMgr {
	if db == nil {
		panic(fmt.Errorf("SysCheckTheInformationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysCheckTheInformationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_check_the_information"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysCheckTheInformationMgr) GetTableName() string {
	return "sys_check_the_information"
}

// Reset 重置gorm会话
func (obj *_SysCheckTheInformationMgr) Reset() *_SysCheckTheInformationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SysCheckTheInformationMgr) Get() (result SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysCheckTheInformationMgr) Gets() (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SysCheckTheInformationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_SysCheckTheInformationMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFilterKey filter_key获取 校验代码
func (obj *_SysCheckTheInformationMgr) WithFilterKey(filterKey string) Option {
	return optionFunc(func(o *options) { o.query["filter_key"] = filterKey })
}

// WithFilterValue filter_value获取 校验过滤值
func (obj *_SysCheckTheInformationMgr) WithFilterValue(filterValue string) Option {
	return optionFunc(func(o *options) { o.query["filter_value"] = filterValue })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysCheckTheInformationMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_SysCheckTheInformationMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysCheckTheInformationMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_SysCheckTheInformationMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_SysCheckTheInformationMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithRemarks remarks获取 备注
func (obj *_SysCheckTheInformationMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// GetByOption 功能选项模式获取
func (obj *_SysCheckTheInformationMgr) GetByOption(opts ...Option) (result SysCheckTheInformation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysCheckTheInformationMgr) GetByOptions(opts ...Option) (results []*SysCheckTheInformation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SysCheckTheInformationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SysCheckTheInformation, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where(options.query)
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
func (obj *_SysCheckTheInformationMgr) GetFromID(id int) (result SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_SysCheckTheInformationMgr) GetBatchFromID(ids []int) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromFilterKey 通过filter_key获取内容 校验代码
func (obj *_SysCheckTheInformationMgr) GetFromFilterKey(filterKey string) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`filter_key` = ?", filterKey).Find(&results).Error

	return
}

// GetBatchFromFilterKey 批量查找 校验代码
func (obj *_SysCheckTheInformationMgr) GetBatchFromFilterKey(filterKeys []string) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`filter_key` IN (?)", filterKeys).Find(&results).Error

	return
}

// GetFromFilterValue 通过filter_value获取内容 校验过滤值
func (obj *_SysCheckTheInformationMgr) GetFromFilterValue(filterValue string) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`filter_value` = ?", filterValue).Find(&results).Error

	return
}

// GetBatchFromFilterValue 批量查找 校验过滤值
func (obj *_SysCheckTheInformationMgr) GetBatchFromFilterValue(filterValues []string) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`filter_value` IN (?)", filterValues).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysCheckTheInformationMgr) GetFromCreateTime(createTime time.Time) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SysCheckTheInformationMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_SysCheckTheInformationMgr) GetFromCreateUser(createUser int) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_SysCheckTheInformationMgr) GetBatchFromCreateUser(createUsers []int) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysCheckTheInformationMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_SysCheckTheInformationMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_SysCheckTheInformationMgr) GetFromUpdateUser(updateUser int) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_SysCheckTheInformationMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_SysCheckTheInformationMgr) GetFromVersion(version int) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_SysCheckTheInformationMgr) GetBatchFromVersion(versions []int) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注
func (obj *_SysCheckTheInformationMgr) GetFromRemarks(remarks string) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注
func (obj *_SysCheckTheInformationMgr) GetBatchFromRemarks(remarkss []string) (results []*SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SysCheckTheInformationMgr) FetchByPrimaryKey(id int) (result SysCheckTheInformation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysCheckTheInformation{}).Where("`id` = ?", id).First(&result).Error

	return
}
