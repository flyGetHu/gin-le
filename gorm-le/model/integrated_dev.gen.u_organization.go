package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UOrganizationMgr struct {
	*_BaseMgr
}

// UOrganizationMgr open func
func UOrganizationMgr(db *gorm.DB) *_UOrganizationMgr {
	if db == nil {
		panic(fmt.Errorf("UOrganizationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UOrganizationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_organization"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UOrganizationMgr) GetTableName() string {
	return "u_organization"
}

// Reset 重置gorm会话
func (obj *_UOrganizationMgr) Reset() *_UOrganizationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UOrganizationMgr) Get() (result UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UOrganizationMgr) Gets() (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UOrganizationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_UOrganizationMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 父ID
func (obj *_UOrganizationMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithName name获取 名称
func (obj *_UOrganizationMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOrgCode org_code获取 组织结构代码
func (obj *_UOrganizationMgr) WithOrgCode(orgCode string) Option {
	return optionFunc(func(o *options) { o.query["org_code"] = orgCode })
}

// WithSort sort获取 排序 值越小越靠前
func (obj *_UOrganizationMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithIsParent is_parent获取 该类目是否为父类目，1为true，0为false
func (obj *_UOrganizationMgr) WithIsParent(isParent int8) Option {
	return optionFunc(func(o *options) { o.query["is_parent"] = isParent })
}

// WithLevel level获取 级别
func (obj *_UOrganizationMgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithOrgType org_type获取 组织类型
func (obj *_UOrganizationMgr) WithOrgType(orgType int) Option {
	return optionFunc(func(o *options) { o.query["org_type"] = orgType })
}

// WithRemark remark获取 备注
func (obj *_UOrganizationMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UOrganizationMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UOrganizationMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UOrganizationMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_UOrganizationMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UOrganizationMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_UOrganizationMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_UOrganizationMgr) GetByOption(opts ...Option) (result UOrganization, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UOrganizationMgr) GetByOptions(opts ...Option) (results []*UOrganization, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UOrganizationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UOrganization, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_UOrganizationMgr) GetFromID(id int) (result UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_UOrganizationMgr) GetBatchFromID(ids []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父ID
func (obj *_UOrganizationMgr) GetFromParentID(parentID int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父ID
func (obj *_UOrganizationMgr) GetBatchFromParentID(parentIDs []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_UOrganizationMgr) GetFromName(name string) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_UOrganizationMgr) GetBatchFromName(names []string) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromOrgCode 通过org_code获取内容 组织结构代码
func (obj *_UOrganizationMgr) GetFromOrgCode(orgCode string) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`org_code` = ?", orgCode).Find(&results).Error

	return
}

// GetBatchFromOrgCode 批量查找 组织结构代码
func (obj *_UOrganizationMgr) GetBatchFromOrgCode(orgCodes []string) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`org_code` IN (?)", orgCodes).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序 值越小越靠前
func (obj *_UOrganizationMgr) GetFromSort(sort int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序 值越小越靠前
func (obj *_UOrganizationMgr) GetBatchFromSort(sorts []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromIsParent 通过is_parent获取内容 该类目是否为父类目，1为true，0为false
func (obj *_UOrganizationMgr) GetFromIsParent(isParent int8) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`is_parent` = ?", isParent).Find(&results).Error

	return
}

// GetBatchFromIsParent 批量查找 该类目是否为父类目，1为true，0为false
func (obj *_UOrganizationMgr) GetBatchFromIsParent(isParents []int8) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`is_parent` IN (?)", isParents).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 级别
func (obj *_UOrganizationMgr) GetFromLevel(level int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找 级别
func (obj *_UOrganizationMgr) GetBatchFromLevel(levels []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromOrgType 通过org_type获取内容 组织类型
func (obj *_UOrganizationMgr) GetFromOrgType(orgType int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`org_type` = ?", orgType).Find(&results).Error

	return
}

// GetBatchFromOrgType 批量查找 组织类型
func (obj *_UOrganizationMgr) GetBatchFromOrgType(orgTypes []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`org_type` IN (?)", orgTypes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_UOrganizationMgr) GetFromRemark(remark string) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_UOrganizationMgr) GetBatchFromRemark(remarks []string) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UOrganizationMgr) GetFromCreateTime(createTime time.Time) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UOrganizationMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UOrganizationMgr) GetFromCreateUser(createUser int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UOrganizationMgr) GetBatchFromCreateUser(createUsers []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UOrganizationMgr) GetFromUpdateTime(updateTime time.Time) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UOrganizationMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_UOrganizationMgr) GetFromUpdateUser(updateUser int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_UOrganizationMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UOrganizationMgr) GetFromVersion(version int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UOrganizationMgr) GetBatchFromVersion(versions []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_UOrganizationMgr) GetFromDeleted(deleted int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_UOrganizationMgr) GetBatchFromDeleted(deleteds []int) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UOrganizationMgr) FetchByPrimaryKey(id int) (result UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByUOrganizationNameIndex  获取多个内容
func (obj *_UOrganizationMgr) FetchIndexByUOrganizationNameIndex(name string) (results []*UOrganization, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UOrganization{}).Where("`name` = ?", name).Find(&results).Error

	return
}
