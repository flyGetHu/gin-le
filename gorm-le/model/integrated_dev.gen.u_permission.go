package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UPermissionMgr struct {
	*_BaseMgr
}

// UPermissionMgr open func
func UPermissionMgr(db *gorm.DB) *_UPermissionMgr {
	if db == nil {
		panic(fmt.Errorf("UPermissionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UPermissionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_permission"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UPermissionMgr) GetTableName() string {
	return "u_permission"
}

// Reset 重置gorm会话
func (obj *_UPermissionMgr) Reset() *_UPermissionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UPermissionMgr) Get() (result UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UPermissionMgr) Gets() (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UPermissionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UPermission{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_UPermissionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 父节点ID
func (obj *_UPermissionMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithPermissionName permission_name获取 权限名称
func (obj *_UPermissionMgr) WithPermissionName(permissionName string) Option {
	return optionFunc(func(o *options) { o.query["permission_name"] = permissionName })
}

// WithPermissionCode permission_code获取 权限标识符
func (obj *_UPermissionMgr) WithPermissionCode(permissionCode string) Option {
	return optionFunc(func(o *options) { o.query["permission_code"] = permissionCode })
}

// WithSort sort获取 排序 值越小越靠前
func (obj *_UPermissionMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithIsParent is_parent获取 该类目是否为父类目，1为true，0为false
func (obj *_UPermissionMgr) WithIsParent(isParent int8) Option {
	return optionFunc(func(o *options) { o.query["is_parent"] = isParent })
}

// WithRemark remark获取 备注
func (obj *_UPermissionMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UPermissionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UPermissionMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UPermissionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_UPermissionMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UPermissionMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_UPermissionMgr) GetByOption(opts ...Option) (result UPermission, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UPermissionMgr) GetByOptions(opts ...Option) (results []*UPermission, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UPermissionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UPermission, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where(options.query)
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
func (obj *_UPermissionMgr) GetFromID(id int) (result UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_UPermissionMgr) GetBatchFromID(ids []int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父节点ID
func (obj *_UPermissionMgr) GetFromParentID(parentID int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父节点ID
func (obj *_UPermissionMgr) GetBatchFromParentID(parentIDs []int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromPermissionName 通过permission_name获取内容 权限名称
func (obj *_UPermissionMgr) GetFromPermissionName(permissionName string) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`permission_name` = ?", permissionName).Find(&results).Error

	return
}

// GetBatchFromPermissionName 批量查找 权限名称
func (obj *_UPermissionMgr) GetBatchFromPermissionName(permissionNames []string) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`permission_name` IN (?)", permissionNames).Find(&results).Error

	return
}

// GetFromPermissionCode 通过permission_code获取内容 权限标识符
func (obj *_UPermissionMgr) GetFromPermissionCode(permissionCode string) (result UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`permission_code` = ?", permissionCode).First(&result).Error

	return
}

// GetBatchFromPermissionCode 批量查找 权限标识符
func (obj *_UPermissionMgr) GetBatchFromPermissionCode(permissionCodes []string) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`permission_code` IN (?)", permissionCodes).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序 值越小越靠前
func (obj *_UPermissionMgr) GetFromSort(sort int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序 值越小越靠前
func (obj *_UPermissionMgr) GetBatchFromSort(sorts []int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromIsParent 通过is_parent获取内容 该类目是否为父类目，1为true，0为false
func (obj *_UPermissionMgr) GetFromIsParent(isParent int8) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`is_parent` = ?", isParent).Find(&results).Error

	return
}

// GetBatchFromIsParent 批量查找 该类目是否为父类目，1为true，0为false
func (obj *_UPermissionMgr) GetBatchFromIsParent(isParents []int8) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`is_parent` IN (?)", isParents).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_UPermissionMgr) GetFromRemark(remark string) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_UPermissionMgr) GetBatchFromRemark(remarks []string) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UPermissionMgr) GetFromCreateTime(createTime time.Time) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UPermissionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UPermissionMgr) GetFromCreateUser(createUser int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UPermissionMgr) GetBatchFromCreateUser(createUsers []int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UPermissionMgr) GetFromUpdateTime(updateTime time.Time) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UPermissionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_UPermissionMgr) GetFromUpdateUser(updateUser int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_UPermissionMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UPermissionMgr) GetFromVersion(version int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UPermissionMgr) GetBatchFromVersion(versions []int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UPermissionMgr) FetchByPrimaryKey(id int) (result UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexPermissionCode primary or index 获取唯一内容
func (obj *_UPermissionMgr) FetchUniqueByIndexPermissionCode(permissionCode string) (result UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`permission_code` = ?", permissionCode).First(&result).Error

	return
}

// FetchIndexByUPermissionParentIDIndex  获取多个内容
func (obj *_UPermissionMgr) FetchIndexByUPermissionParentIDIndex(parentID int) (results []*UPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UPermission{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}
