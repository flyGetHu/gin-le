package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UResourceTagMgr struct {
	*_BaseMgr
}

// UResourceTagMgr open func
func UResourceTagMgr(db *gorm.DB) *_UResourceTagMgr {
	if db == nil {
		panic(fmt.Errorf("UResourceTagMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UResourceTagMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_resource_tag"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UResourceTagMgr) GetTableName() string {
	return "u_resource_tag"
}

// Reset 重置gorm会话
func (obj *_UResourceTagMgr) Reset() *_UResourceTagMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UResourceTagMgr) Get() (result UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UResourceTagMgr) Gets() (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UResourceTagMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UResourceTagMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 资源标签名称
func (obj *_UResourceTagMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithType type获取 类型:保留字段
func (obj *_UResourceTagMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithRemark remark获取 备注
func (obj *_UResourceTagMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UResourceTagMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UResourceTagMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UResourceTagMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_UResourceTagMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UResourceTagMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_UResourceTagMgr) GetByOption(opts ...Option) (result UResourceTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UResourceTagMgr) GetByOptions(opts ...Option) (results []*UResourceTag, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UResourceTagMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UResourceTag, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where(options.query)
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
func (obj *_UResourceTagMgr) GetFromID(id int) (result UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UResourceTagMgr) GetBatchFromID(ids []int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 资源标签名称
func (obj *_UResourceTagMgr) GetFromName(name string) (result UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 资源标签名称
func (obj *_UResourceTagMgr) GetBatchFromName(names []string) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型:保留字段
func (obj *_UResourceTagMgr) GetFromType(_type int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型:保留字段
func (obj *_UResourceTagMgr) GetBatchFromType(_types []int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_UResourceTagMgr) GetFromRemark(remark string) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_UResourceTagMgr) GetBatchFromRemark(remarks []string) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UResourceTagMgr) GetFromCreateTime(createTime time.Time) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UResourceTagMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UResourceTagMgr) GetFromCreateUser(createUser int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UResourceTagMgr) GetBatchFromCreateUser(createUsers []int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UResourceTagMgr) GetFromUpdateTime(updateTime time.Time) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UResourceTagMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_UResourceTagMgr) GetFromUpdateUser(updateUser int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_UResourceTagMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UResourceTagMgr) GetFromVersion(version int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UResourceTagMgr) GetBatchFromVersion(versions []int) (results []*UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UResourceTagMgr) FetchByPrimaryKey(id int) (result UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUResourceTagNameUIndex primary or index 获取唯一内容
func (obj *_UResourceTagMgr) FetchUniqueByUResourceTagNameUIndex(name string) (result UResourceTag, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UResourceTag{}).Where("`name` = ?", name).First(&result).Error

	return
}
