package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WMicroTicketTypeMgr struct {
	*_BaseMgr
}

// WMicroTicketTypeMgr open func
func WMicroTicketTypeMgr(db *gorm.DB) *_WMicroTicketTypeMgr {
	if db == nil {
		panic(fmt.Errorf("WMicroTicketTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WMicroTicketTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_micro_ticket_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WMicroTicketTypeMgr) GetTableName() string {
	return "w_micro_ticket_type"
}

// Reset 重置gorm会话
func (obj *_WMicroTicketTypeMgr) Reset() *_WMicroTicketTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WMicroTicketTypeMgr) Get() (result WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WMicroTicketTypeMgr) Gets() (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WMicroTicketTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_WMicroTicketTypeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTypeName type_name获取 类型名称
func (obj *_WMicroTicketTypeMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithEnglishName english_name获取 英文名称
func (obj *_WMicroTicketTypeMgr) WithEnglishName(englishName string) Option {
	return optionFunc(func(o *options) { o.query["english_name"] = englishName })
}

// WithTypeStatus type_status获取 状态（0:禁用1:启用）
func (obj *_WMicroTicketTypeMgr) WithTypeStatus(typeStatus []uint8) Option {
	return optionFunc(func(o *options) { o.query["type_status"] = typeStatus })
}

// WithTypeDescribe type_describe获取 类型描述
func (obj *_WMicroTicketTypeMgr) WithTypeDescribe(typeDescribe string) Option {
	return optionFunc(func(o *options) { o.query["type_describe"] = typeDescribe })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WMicroTicketTypeMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WMicroTicketTypeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_WMicroTicketTypeMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WMicroTicketTypeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_WMicroTicketTypeMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WMicroTicketTypeMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WMicroTicketTypeMgr) GetByOption(opts ...Option) (result WMicroTicketType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WMicroTicketTypeMgr) GetByOptions(opts ...Option) (results []*WMicroTicketType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WMicroTicketTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WMicroTicketType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *_WMicroTicketTypeMgr) GetFromID(id int) (result WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_WMicroTicketTypeMgr) GetBatchFromID(ids []int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容 类型名称
func (obj *_WMicroTicketTypeMgr) GetFromTypeName(typeName string) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量查找 类型名称
func (obj *_WMicroTicketTypeMgr) GetBatchFromTypeName(typeNames []string) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromEnglishName 通过english_name获取内容 英文名称
func (obj *_WMicroTicketTypeMgr) GetFromEnglishName(englishName string) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`english_name` = ?", englishName).Find(&results).Error

	return
}

// GetBatchFromEnglishName 批量查找 英文名称
func (obj *_WMicroTicketTypeMgr) GetBatchFromEnglishName(englishNames []string) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`english_name` IN (?)", englishNames).Find(&results).Error

	return
}

// GetFromTypeStatus 通过type_status获取内容 状态（0:禁用1:启用）
func (obj *_WMicroTicketTypeMgr) GetFromTypeStatus(typeStatus []uint8) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`type_status` = ?", typeStatus).Find(&results).Error

	return
}

// GetBatchFromTypeStatus 批量查找 状态（0:禁用1:启用）
func (obj *_WMicroTicketTypeMgr) GetBatchFromTypeStatus(typeStatuss [][]uint8) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`type_status` IN (?)", typeStatuss).Find(&results).Error

	return
}

// GetFromTypeDescribe 通过type_describe获取内容 类型描述
func (obj *_WMicroTicketTypeMgr) GetFromTypeDescribe(typeDescribe string) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`type_describe` = ?", typeDescribe).Find(&results).Error

	return
}

// GetBatchFromTypeDescribe 批量查找 类型描述
func (obj *_WMicroTicketTypeMgr) GetBatchFromTypeDescribe(typeDescribes []string) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`type_describe` IN (?)", typeDescribes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WMicroTicketTypeMgr) GetFromCreateUser(createUser int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WMicroTicketTypeMgr) GetBatchFromCreateUser(createUsers []int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WMicroTicketTypeMgr) GetFromCreateTime(createTime time.Time) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WMicroTicketTypeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_WMicroTicketTypeMgr) GetFromUpdateUser(updateUser int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_WMicroTicketTypeMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WMicroTicketTypeMgr) GetFromUpdateTime(updateTime time.Time) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WMicroTicketTypeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WMicroTicketTypeMgr) GetFromVersion(version int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WMicroTicketTypeMgr) GetBatchFromVersion(versions []int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WMicroTicketTypeMgr) GetFromDeleted(deleted int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WMicroTicketTypeMgr) GetBatchFromDeleted(deleteds []int) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchUniqueByWMicroTicketTypeIDUIndex primary or index 获取唯一内容
func (obj *_WMicroTicketTypeMgr) FetchUniqueByWMicroTicketTypeIDUIndex(id int) (result WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByWMicroTicketTypeTypeNameIndex  获取多个内容
func (obj *_WMicroTicketTypeMgr) FetchIndexByWMicroTicketTypeTypeNameIndex(typeName string) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// FetchIndexByWMicroTicketTypeTypeStatusIndex  获取多个内容
func (obj *_WMicroTicketTypeMgr) FetchIndexByWMicroTicketTypeTypeStatusIndex(typeStatus []uint8) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`type_status` = ?", typeStatus).Find(&results).Error

	return
}

// FetchIndexByWMicroTicketTypeCreateTimeIndex  获取多个内容
func (obj *_WMicroTicketTypeMgr) FetchIndexByWMicroTicketTypeCreateTimeIndex(createTime time.Time) (results []*WMicroTicketType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WMicroTicketType{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
