package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgLabelTemplateMgr struct {
	*_BaseMgr
}

// LgLabelTemplateMgr open func
func LgLabelTemplateMgr(db *gorm.DB) *_LgLabelTemplateMgr {
	if db == nil {
		panic(fmt.Errorf("LgLabelTemplateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgLabelTemplateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_label_template"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgLabelTemplateMgr) GetTableName() string {
	return "lg_label_template"
}

// Reset 重置gorm会话
func (obj *_LgLabelTemplateMgr) Reset() *_LgLabelTemplateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgLabelTemplateMgr) Get() (result LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", result.ProviderSystemID).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", result.ProviderSystemName).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", result.ProviderSystemCode).Find(&result.LgProviderSystem).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_LgLabelTemplateMgr) Gets() (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgLabelTemplateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgLabelTemplateMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 面单名称
func (obj *_LgLabelTemplateMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取 面单代码
func (obj *_LgLabelTemplateMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithType type获取 面单类型1：自制 2：服务商
func (obj *_LgLabelTemplateMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithProviderSystemID provider_system_id获取 绑定的服务商系统id
func (obj *_LgLabelTemplateMgr) WithProviderSystemID(providerSystemID int) Option {
	return optionFunc(func(o *options) { o.query["provider_system_id"] = providerSystemID })
}

// WithProviderSystemName provider_system_name获取 绑定的服务商系统名称
func (obj *_LgLabelTemplateMgr) WithProviderSystemName(providerSystemName string) Option {
	return optionFunc(func(o *options) { o.query["provider_system_name"] = providerSystemName })
}

// WithProviderSystemCode provider_system_code获取 绑定服务商系统code
func (obj *_LgLabelTemplateMgr) WithProviderSystemCode(providerSystemCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_system_code"] = providerSystemCode })
}

// WithRemark remark获取 备注
func (obj *_LgLabelTemplateMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgLabelTemplateMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgLabelTemplateMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgLabelTemplateMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgLabelTemplateMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgLabelTemplateMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgLabelTemplateMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgLabelTemplateMgr) GetByOption(opts ...Option) (result LgLabelTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", result.ProviderSystemID).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", result.ProviderSystemName).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", result.ProviderSystemCode).Find(&result.LgProviderSystem).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgLabelTemplateMgr) GetByOptions(opts ...Option) (results []*LgLabelTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_LgLabelTemplateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgLabelTemplate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
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
func (obj *_LgLabelTemplateMgr) GetFromID(id int64) (result LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", result.ProviderSystemID).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", result.ProviderSystemName).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", result.ProviderSystemCode).Find(&result.LgProviderSystem).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_LgLabelTemplateMgr) GetBatchFromID(ids []int64) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromName 通过name获取内容 面单名称
func (obj *_LgLabelTemplateMgr) GetFromName(name string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`name` = ?", name).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromName 批量查找 面单名称
func (obj *_LgLabelTemplateMgr) GetBatchFromName(names []string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`name` IN (?)", names).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCode 通过code获取内容 面单代码
func (obj *_LgLabelTemplateMgr) GetFromCode(code string) (result LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`code` = ?", code).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", result.ProviderSystemID).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", result.ProviderSystemName).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", result.ProviderSystemCode).Find(&result.LgProviderSystem).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromCode 批量查找 面单代码
func (obj *_LgLabelTemplateMgr) GetBatchFromCode(codes []string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`code` IN (?)", codes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromType 通过type获取内容 面单类型1：自制 2：服务商
func (obj *_LgLabelTemplateMgr) GetFromType(_type int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`type` = ?", _type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromType 批量查找 面单类型1：自制 2：服务商
func (obj *_LgLabelTemplateMgr) GetBatchFromType(_types []int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`type` IN (?)", _types).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProviderSystemID 通过provider_system_id获取内容 绑定的服务商系统id
func (obj *_LgLabelTemplateMgr) GetFromProviderSystemID(providerSystemID int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_id` = ?", providerSystemID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProviderSystemID 批量查找 绑定的服务商系统id
func (obj *_LgLabelTemplateMgr) GetBatchFromProviderSystemID(providerSystemIDs []int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_id` IN (?)", providerSystemIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProviderSystemName 通过provider_system_name获取内容 绑定的服务商系统名称
func (obj *_LgLabelTemplateMgr) GetFromProviderSystemName(providerSystemName string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_name` = ?", providerSystemName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProviderSystemName 批量查找 绑定的服务商系统名称
func (obj *_LgLabelTemplateMgr) GetBatchFromProviderSystemName(providerSystemNames []string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_name` IN (?)", providerSystemNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProviderSystemCode 通过provider_system_code获取内容 绑定服务商系统code
func (obj *_LgLabelTemplateMgr) GetFromProviderSystemCode(providerSystemCode string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_code` = ?", providerSystemCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProviderSystemCode 批量查找 绑定服务商系统code
func (obj *_LgLabelTemplateMgr) GetBatchFromProviderSystemCode(providerSystemCodes []string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_code` IN (?)", providerSystemCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_LgLabelTemplateMgr) GetFromRemark(remark string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`remark` = ?", remark).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_LgLabelTemplateMgr) GetBatchFromRemark(remarks []string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`remark` IN (?)", remarks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgLabelTemplateMgr) GetFromCreateTime(createTime time.Time) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgLabelTemplateMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgLabelTemplateMgr) GetFromCreateUser(createUser int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgLabelTemplateMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgLabelTemplateMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgLabelTemplateMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgLabelTemplateMgr) GetFromUpdateUser(updateUser int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgLabelTemplateMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgLabelTemplateMgr) GetFromVersion(version int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgLabelTemplateMgr) GetBatchFromVersion(versions []int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgLabelTemplateMgr) GetFromDeleted(deleted int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgLabelTemplateMgr) GetBatchFromDeleted(deleteds []int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
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
func (obj *_LgLabelTemplateMgr) FetchByPrimaryKey(id int64) (result LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", result.ProviderSystemID).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", result.ProviderSystemName).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", result.ProviderSystemCode).Find(&result.LgProviderSystem).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByLgLabelTemplateCodeUIndex primary or index 获取唯一内容
func (obj *_LgLabelTemplateMgr) FetchUniqueByLgLabelTemplateCodeUIndex(code string) (result LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`code` = ?", code).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", result.ProviderSystemID).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", result.ProviderSystemName).Find(&result.LgProviderSystemList).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", result.ProviderSystemCode).Find(&result.LgProviderSystem).Error; err != nil { // 系统服务商
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByLgLabelTemplateLgProviderSystemIDFk  获取多个内容
func (obj *_LgLabelTemplateMgr) FetchIndexByLgLabelTemplateLgProviderSystemIDFk(providerSystemID int) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_id` = ?", providerSystemID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByLgLabelTemplateLgProviderSystemNameFk  获取多个内容
func (obj *_LgLabelTemplateMgr) FetchIndexByLgLabelTemplateLgProviderSystemNameFk(providerSystemName string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_name` = ?", providerSystemName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByLgLabelTemplateLgProviderSystemCodeFk  获取多个内容
func (obj *_LgLabelTemplateMgr) FetchIndexByLgLabelTemplateLgProviderSystemCodeFk(providerSystemCode string) (results []*LgLabelTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgLabelTemplate{}).Where("`provider_system_code` = ?", providerSystemCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_provider_system").Where("id = ?", results[i].ProviderSystemID).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("name = ?", results[i].ProviderSystemName).Find(&results[i].LgProviderSystemList).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_provider_system").Where("code = ?", results[i].ProviderSystemCode).Find(&results[i].LgProviderSystem).Error; err != nil { // 系统服务商
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
