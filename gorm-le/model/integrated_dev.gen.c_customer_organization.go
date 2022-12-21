package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CCustomerOrganizationMgr struct {
	*_BaseMgr
}

// CCustomerOrganizationMgr open func
func CCustomerOrganizationMgr(db *gorm.DB) *_CCustomerOrganizationMgr {
	if db == nil {
		panic(fmt.Errorf("CCustomerOrganizationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CCustomerOrganizationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_customer_organization"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CCustomerOrganizationMgr) GetTableName() string {
	return "c_customer_organization"
}

// Reset 重置gorm会话
func (obj *_CCustomerOrganizationMgr) Reset() *_CCustomerOrganizationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CCustomerOrganizationMgr) Get() (result CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CCustomerOrganizationMgr) Gets() (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CCustomerOrganizationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_CCustomerOrganizationMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrganizationCode organization_code获取 编码
func (obj *_CCustomerOrganizationMgr) WithOrganizationCode(organizationCode string) Option {
	return optionFunc(func(o *options) { o.query["organization_code"] = organizationCode })
}

// WithOrganizationName organization_name获取 名称
func (obj *_CCustomerOrganizationMgr) WithOrganizationName(organizationName string) Option {
	return optionFunc(func(o *options) { o.query["organization_name"] = organizationName })
}

// WithOrganizationStatus organization_status获取 是否启用 0：启用，1：禁用
func (obj *_CCustomerOrganizationMgr) WithOrganizationStatus(organizationStatus int) Option {
	return optionFunc(func(o *options) { o.query["organization_status"] = organizationStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CCustomerOrganizationMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_CCustomerOrganizationMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CCustomerOrganizationMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_CCustomerOrganizationMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_CCustomerOrganizationMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_CCustomerOrganizationMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_CCustomerOrganizationMgr) GetByOption(opts ...Option) (result CCustomerOrganization, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CCustomerOrganizationMgr) GetByOptions(opts ...Option) (results []*CCustomerOrganization, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CCustomerOrganizationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CCustomerOrganization, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where(options.query)
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
func (obj *_CCustomerOrganizationMgr) GetFromID(id int) (result CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_CCustomerOrganizationMgr) GetBatchFromID(ids []int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrganizationCode 通过organization_code获取内容 编码
func (obj *_CCustomerOrganizationMgr) GetFromOrganizationCode(organizationCode string) (result CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`organization_code` = ?", organizationCode).First(&result).Error

	return
}

// GetBatchFromOrganizationCode 批量查找 编码
func (obj *_CCustomerOrganizationMgr) GetBatchFromOrganizationCode(organizationCodes []string) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`organization_code` IN (?)", organizationCodes).Find(&results).Error

	return
}

// GetFromOrganizationName 通过organization_name获取内容 名称
func (obj *_CCustomerOrganizationMgr) GetFromOrganizationName(organizationName string) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`organization_name` = ?", organizationName).Find(&results).Error

	return
}

// GetBatchFromOrganizationName 批量查找 名称
func (obj *_CCustomerOrganizationMgr) GetBatchFromOrganizationName(organizationNames []string) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`organization_name` IN (?)", organizationNames).Find(&results).Error

	return
}

// GetFromOrganizationStatus 通过organization_status获取内容 是否启用 0：启用，1：禁用
func (obj *_CCustomerOrganizationMgr) GetFromOrganizationStatus(organizationStatus int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`organization_status` = ?", organizationStatus).Find(&results).Error

	return
}

// GetBatchFromOrganizationStatus 批量查找 是否启用 0：启用，1：禁用
func (obj *_CCustomerOrganizationMgr) GetBatchFromOrganizationStatus(organizationStatuss []int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`organization_status` IN (?)", organizationStatuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CCustomerOrganizationMgr) GetFromCreateTime(createTime time.Time) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CCustomerOrganizationMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_CCustomerOrganizationMgr) GetFromCreateUser(createUser int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_CCustomerOrganizationMgr) GetBatchFromCreateUser(createUsers []int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CCustomerOrganizationMgr) GetFromUpdateTime(updateTime time.Time) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_CCustomerOrganizationMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_CCustomerOrganizationMgr) GetFromUpdateUser(updateUser int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_CCustomerOrganizationMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_CCustomerOrganizationMgr) GetFromVersion(version int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_CCustomerOrganizationMgr) GetBatchFromVersion(versions []int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_CCustomerOrganizationMgr) GetFromDeleted(deleted int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_CCustomerOrganizationMgr) GetBatchFromDeleted(deleteds []int) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CCustomerOrganizationMgr) FetchByPrimaryKey(id int) (result CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexOrganizationCode primary or index 获取唯一内容
func (obj *_CCustomerOrganizationMgr) FetchUniqueByIndexOrganizationCode(organizationCode string) (result CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`organization_code` = ?", organizationCode).First(&result).Error

	return
}

// FetchIndexByIndexOrganizationName  获取多个内容
func (obj *_CCustomerOrganizationMgr) FetchIndexByIndexOrganizationName(organizationName string) (results []*CCustomerOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerOrganization{}).Where("`organization_name` = ?", organizationName).Find(&results).Error

	return
}
