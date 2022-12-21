package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UUserPermissionMgr struct {
	*_BaseMgr
}

// UUserPermissionMgr open func
func UUserPermissionMgr(db *gorm.DB) *_UUserPermissionMgr {
	if db == nil {
		panic(fmt.Errorf("UUserPermissionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UUserPermissionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_user_permission"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UUserPermissionMgr) GetTableName() string {
	return "u_user_permission"
}

// Reset 重置gorm会话
func (obj *_UUserPermissionMgr) Reset() *_UUserPermissionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UUserPermissionMgr) Get() (result UUserPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_user").Where("id = ?", result.UserID).Find(&result.UUser).Error; err != nil { // 用户表
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
func (obj *_UUserPermissionMgr) Gets() (results []*UUserPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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
func (obj *_UUserPermissionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取 用户ID
func (obj *_UUserPermissionMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithPermissionID permission_id获取 权限ID
func (obj *_UUserPermissionMgr) WithPermissionID(permissionID int) Option {
	return optionFunc(func(o *options) { o.query["permission_id"] = permissionID })
}

// GetByOption 功能选项模式获取
func (obj *_UUserPermissionMgr) GetByOption(opts ...Option) (result UUserPermission, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_user").Where("id = ?", result.UserID).Find(&result.UUser).Error; err != nil { // 用户表
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
func (obj *_UUserPermissionMgr) GetByOptions(opts ...Option) (results []*UUserPermission, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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
func (obj *_UUserPermissionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UUserPermission, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_UUserPermissionMgr) GetFromUserID(userID int) (results []*UUserPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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

// GetBatchFromUserID 批量查找 用户ID
func (obj *_UUserPermissionMgr) GetBatchFromUserID(userIDs []int) (results []*UUserPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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

// GetFromPermissionID 通过permission_id获取内容 权限ID
func (obj *_UUserPermissionMgr) GetFromPermissionID(permissionID int) (results []*UUserPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where("`permission_id` = ?", permissionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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

// GetBatchFromPermissionID 批量查找 权限ID
func (obj *_UUserPermissionMgr) GetBatchFromPermissionID(permissionIDs []int) (results []*UUserPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where("`permission_id` IN (?)", permissionIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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

// FetchIndexByUUserPermissionUUserIDFk  获取多个内容
func (obj *_UUserPermissionMgr) FetchIndexByUUserPermissionUUserIDFk(userID int) (results []*UUserPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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

// FetchIndexByUUserPermissionUPermissionIDFk  获取多个内容
func (obj *_UUserPermissionMgr) FetchIndexByUUserPermissionUPermissionIDFk(permissionID int) (results []*UUserPermission, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUserPermission{}).Where("`permission_id` = ?", permissionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_user").Where("id = ?", results[i].UserID).Find(&results[i].UUser).Error; err != nil { // 用户表
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
