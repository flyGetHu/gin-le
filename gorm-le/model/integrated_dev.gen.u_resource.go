package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UResourceMgr struct {
	*_BaseMgr
}

// UResourceMgr open func
func UResourceMgr(db *gorm.DB) *_UResourceMgr {
	if db == nil {
		panic(fmt.Errorf("UResourceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UResourceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_resource"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UResourceMgr) GetTableName() string {
	return "u_resource"
}

// Reset 重置gorm会话
func (obj *_UResourceMgr) Reset() *_UResourceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UResourceMgr) Get() (result UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", result.Tag).Find(&result.UResourceTag).Error; err != nil { // 资源标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UResourceMgr) Gets() (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UResourceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UResource{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_UResourceMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 资源编号
func (obj *_UResourceMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithUserID user_id获取 用户iD
func (obj *_UResourceMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithType type获取 文件标签类型(文件,图片等)
func (obj *_UResourceMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithTag tag获取 自定义标签(例如IT部文件)
func (obj *_UResourceMgr) WithTag(tag string) Option {
	return optionFunc(func(o *options) { o.query["tag"] = tag })
}

// WithName name获取 文件名
func (obj *_UResourceMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithResources resources获取 资源列表(下载链接)
func (obj *_UResourceMgr) WithResources(resources string) Option {
	return optionFunc(func(o *options) { o.query["resources"] = resources })
}

// WithIsOpen is_open获取 是否公开 0:不公开(私人) 1:公开(公共)
func (obj *_UResourceMgr) WithIsOpen(isOpen int) Option {
	return optionFunc(func(o *options) { o.query["is_open"] = isOpen })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UResourceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UResourceMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UResourceMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_UResourceMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UResourceMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_UResourceMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_UResourceMgr) GetByOption(opts ...Option) (result UResource, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", result.Tag).Find(&result.UResourceTag).Error; err != nil { // 资源标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UResourceMgr) GetByOptions(opts ...Option) (results []*UResource, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_UResourceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UResource, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UResource{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
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

// GetFromID 通过id获取内容 ID
func (obj *_UResourceMgr) GetFromID(id int) (result UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", result.Tag).Find(&result.UResourceTag).Error; err != nil { // 资源标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_UResourceMgr) GetBatchFromID(ids []int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCode 通过code获取内容 资源编号
func (obj *_UResourceMgr) GetFromCode(code string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`code` = ?", code).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCode 批量查找 资源编号
func (obj *_UResourceMgr) GetBatchFromCode(codes []string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`code` IN (?)", codes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容 用户iD
func (obj *_UResourceMgr) GetFromUserID(userID int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量查找 用户iD
func (obj *_UResourceMgr) GetBatchFromUserID(userIDs []int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromType 通过type获取内容 文件标签类型(文件,图片等)
func (obj *_UResourceMgr) GetFromType(_type string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`type` = ?", _type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromType 批量查找 文件标签类型(文件,图片等)
func (obj *_UResourceMgr) GetBatchFromType(_types []string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`type` IN (?)", _types).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTag 通过tag获取内容 自定义标签(例如IT部文件)
func (obj *_UResourceMgr) GetFromTag(tag string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`tag` = ?", tag).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTag 批量查找 自定义标签(例如IT部文件)
func (obj *_UResourceMgr) GetBatchFromTag(tags []string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`tag` IN (?)", tags).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容 文件名
func (obj *_UResourceMgr) GetFromName(name string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量查找 文件名
func (obj *_UResourceMgr) GetBatchFromName(names []string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromResources 通过resources获取内容 资源列表(下载链接)
func (obj *_UResourceMgr) GetFromResources(resources string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`resources` = ?", resources).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromResources 批量查找 资源列表(下载链接)
func (obj *_UResourceMgr) GetBatchFromResources(resourcess []string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`resources` IN (?)", resourcess).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsOpen 通过is_open获取内容 是否公开 0:不公开(私人) 1:公开(公共)
func (obj *_UResourceMgr) GetFromIsOpen(isOpen int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`is_open` = ?", isOpen).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsOpen 批量查找 是否公开 0:不公开(私人) 1:公开(公共)
func (obj *_UResourceMgr) GetBatchFromIsOpen(isOpens []int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`is_open` IN (?)", isOpens).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UResourceMgr) GetFromCreateTime(createTime time.Time) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UResourceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UResourceMgr) GetFromCreateUser(createUser int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UResourceMgr) GetBatchFromCreateUser(createUsers []int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UResourceMgr) GetFromUpdateTime(updateTime time.Time) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UResourceMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_UResourceMgr) GetFromUpdateUser(updateUser int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_UResourceMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UResourceMgr) GetFromVersion(version int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UResourceMgr) GetBatchFromVersion(versions []int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_UResourceMgr) GetFromDeleted(deleted int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_UResourceMgr) GetBatchFromDeleted(deleteds []int) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UResourceMgr) FetchByPrimaryKey(id int) (result UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", result.Tag).Find(&result.UResourceTag).Error; err != nil { // 资源标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByUResourceTagIndex  获取多个内容
func (obj *_UResourceMgr) FetchIndexByUResourceTagIndex(tag string) (results []*UResource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResource{}).Where("`tag` = ?", tag).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_resource_tag").Where("name = ?", results[i].Tag).Find(&results[i].UResourceTag).Error; err != nil { // 资源标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
