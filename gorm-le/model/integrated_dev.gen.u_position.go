package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UPositionMgr struct {
	*_BaseMgr
}

// UPositionMgr open func
func UPositionMgr(db *gorm.DB) *_UPositionMgr {
	if db == nil {
		panic(fmt.Errorf("UPositionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UPositionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_position"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UPositionMgr) GetTableName() string {
	return "u_position"
}

// Reset 重置gorm会话
func (obj *_UPositionMgr) Reset() *_UPositionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UPositionMgr) Get() (result UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UPositionMgr) Gets() (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UPositionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UPosition{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UPositionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 职务名称
func (obj *_UPositionMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOrgID org_id获取 职务绑定的部门ID
func (obj *_UPositionMgr) WithOrgID(orgID int) Option {
	return optionFunc(func(o *options) { o.query["org_id"] = orgID })
}

// WithOrgIDs org_ids获取 职务关联组织架构ids
func (obj *_UPositionMgr) WithOrgIDs(orgIDs string) Option {
	return optionFunc(func(o *options) { o.query["org_ids"] = orgIDs })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UPositionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UPositionMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UPositionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_UPositionMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UPositionMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_UPositionMgr) GetByOption(opts ...Option) (result UPosition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UPositionMgr) GetByOptions(opts ...Option) (results []*UPosition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UPositionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UPosition, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where(options.query)
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
func (obj *_UPositionMgr) GetFromID(id int) (result UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UPositionMgr) GetBatchFromID(ids []int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 职务名称
func (obj *_UPositionMgr) GetFromName(name string) (result UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 职务名称
func (obj *_UPositionMgr) GetBatchFromName(names []string) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromOrgID 通过org_id获取内容 职务绑定的部门ID
func (obj *_UPositionMgr) GetFromOrgID(orgID int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`org_id` = ?", orgID).Find(&results).Error

	return
}

// GetBatchFromOrgID 批量查找 职务绑定的部门ID
func (obj *_UPositionMgr) GetBatchFromOrgID(orgIDs []int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`org_id` IN (?)", orgIDs).Find(&results).Error

	return
}

// GetFromOrgIDs 通过org_ids获取内容 职务关联组织架构ids
func (obj *_UPositionMgr) GetFromOrgIDs(orgIDs string) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`org_ids` = ?", orgIDs).Find(&results).Error

	return
}

// GetBatchFromOrgIDs 批量查找 职务关联组织架构ids
func (obj *_UPositionMgr) GetBatchFromOrgIDs(orgIDss []string) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`org_ids` IN (?)", orgIDss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UPositionMgr) GetFromCreateTime(createTime time.Time) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UPositionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UPositionMgr) GetFromCreateUser(createUser int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UPositionMgr) GetBatchFromCreateUser(createUsers []int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UPositionMgr) GetFromUpdateTime(updateTime time.Time) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UPositionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_UPositionMgr) GetFromUpdateUser(updateUser int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_UPositionMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UPositionMgr) GetFromVersion(version int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UPositionMgr) GetBatchFromVersion(versions []int) (results []*UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UPositionMgr) FetchByPrimaryKey(id int) (result UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUPositionNameUIndex primary or index 获取唯一内容
func (obj *_UPositionMgr) FetchUniqueByUPositionNameUIndex(name string) (result UPosition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPosition{}).Where("`name` = ?", name).First(&result).Error

	return
}
