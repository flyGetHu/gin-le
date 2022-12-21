package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _URolePermissionMgr struct {
	*_BaseMgr
}

// URolePermissionMgr open func
func URolePermissionMgr(db *gorm.DB) *_URolePermissionMgr {
	if db == nil {
		panic(fmt.Errorf("URolePermissionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_URolePermissionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_role_permission"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_URolePermissionMgr) GetTableName() string {
	return "u_role_permission"
}

// Reset 重置gorm会话
func (obj *_URolePermissionMgr) Reset() *_URolePermissionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_URolePermissionMgr) Get() (result URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_role").Where("id = ?", result.RoleID).Find(&result.URole).Error; err != nil { // 角色表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("u_permission").Where("id = ?", result.PermissionID).Find(&result.UPermission).Error; err != nil { // 权限表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_URolePermissionMgr) Gets() (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_URolePermissionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_URolePermissionMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRoleID role_id获取 角色id
func (obj *_URolePermissionMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithPermissionID permission_id获取 权限id
func (obj *_URolePermissionMgr) WithPermissionID(permissionID int) Option {
	return optionFunc(func(o *options) { o.query["permission_id"] = permissionID })
}

// GetByOption 功能选项模式获取
func (obj *_URolePermissionMgr) GetByOption(opts ...Option) (result URolePermission, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_role").Where("id = ?", result.RoleID).Find(&result.URole).Error; err != nil { // 角色表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("u_permission").Where("id = ?", result.PermissionID).Find(&result.UPermission).Error; err != nil { // 权限表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_URolePermissionMgr) GetByOptions(opts ...Option) (results []*URolePermission, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_URolePermissionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]URolePermission, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_URolePermissionMgr) GetFromID(id int64) (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where("`id` = ?", id).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromID 批量查找
func (obj *_URolePermissionMgr) GetBatchFromID(ids []int64) (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRoleID 通过role_id获取内容 角色id
func (obj *_URolePermissionMgr) GetFromRoleID(roleID int) (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where("`role_id` = ?", roleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRoleID 批量查找 角色id
func (obj *_URolePermissionMgr) GetBatchFromRoleID(roleIDs []int) (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPermissionID 通过permission_id获取内容 权限id
func (obj *_URolePermissionMgr) GetFromPermissionID(permissionID int) (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where("`permission_id` = ?", permissionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPermissionID 批量查找 权限id
func (obj *_URolePermissionMgr) GetBatchFromPermissionID(permissionIDs []int) (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where("`permission_id` IN (?)", permissionIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByURolePermissionURoleIDFk  获取多个内容
func (obj *_URolePermissionMgr) FetchIndexByURolePermissionURoleIDFk(roleID int) (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where("`role_id` = ?", roleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByURolePermissionUPermissionIDFk  获取多个内容
func (obj *_URolePermissionMgr) FetchIndexByURolePermissionUPermissionIDFk(permissionID int) (results []*URolePermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URolePermission{}).Where("`permission_id` = ?", permissionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_role").Where("id = ?", results[i].RoleID).Find(&results[i].URole).Error; err != nil { // 角色表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("u_permission").Where("id = ?", results[i].PermissionID).Find(&results[i].UPermission).Error; err != nil { // 权限表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
