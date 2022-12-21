package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UContractUseTypeMgr struct {
	*_BaseMgr
}

// UContractUseTypeMgr open func
func UContractUseTypeMgr(db *gorm.DB) *_UContractUseTypeMgr {
	if db == nil {
		panic(fmt.Errorf("UContractUseTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UContractUseTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_contract_use_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UContractUseTypeMgr) GetTableName() string {
	return "u_contract_use_type"
}

// Reset 重置gorm会话
func (obj *_UContractUseTypeMgr) Reset() *_UContractUseTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UContractUseTypeMgr) Get() (result UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UContractUseTypeMgr) Gets() (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UContractUseTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UContractUseTypeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTypeName type_name获取 类型名称
func (obj *_UContractUseTypeMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithCode code获取 类型code
func (obj *_UContractUseTypeMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithUserIDs user_ids获取 部分可见的用户id（用","逗号隔开）
func (obj *_UContractUseTypeMgr) WithUserIDs(userIDs string) Option {
	return optionFunc(func(o *options) { o.query["user_ids"] = userIDs })
}

// WithIsDelete is_delete获取 逻辑删除 0 未删除 1 已删除
func (obj *_UContractUseTypeMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UContractUseTypeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_UContractUseTypeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_UContractUseTypeMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 修改人
func (obj *_UContractUseTypeMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UContractUseTypeMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithRemark remark获取 合同用途备注
func (obj *_UContractUseTypeMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// GetByOption 功能选项模式获取
func (obj *_UContractUseTypeMgr) GetByOption(opts ...Option) (result UContractUseType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UContractUseTypeMgr) GetByOptions(opts ...Option) (results []*UContractUseType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UContractUseTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UContractUseType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where(options.query)
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
func (obj *_UContractUseTypeMgr) GetFromID(id int64) (result UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UContractUseTypeMgr) GetBatchFromID(ids []int64) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容 类型名称
func (obj *_UContractUseTypeMgr) GetFromTypeName(typeName string) (result UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`type_name` = ?", typeName).First(&result).Error

	return
}

// GetBatchFromTypeName 批量查找 类型名称
func (obj *_UContractUseTypeMgr) GetBatchFromTypeName(typeNames []string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 类型code
func (obj *_UContractUseTypeMgr) GetFromCode(code string) (result UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 类型code
func (obj *_UContractUseTypeMgr) GetBatchFromCode(codes []string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromUserIDs 通过user_ids获取内容 部分可见的用户id（用","逗号隔开）
func (obj *_UContractUseTypeMgr) GetFromUserIDs(userIDs string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`user_ids` = ?", userIDs).Find(&results).Error

	return
}

// GetBatchFromUserIDs 批量查找 部分可见的用户id（用","逗号隔开）
func (obj *_UContractUseTypeMgr) GetBatchFromUserIDs(userIDss []string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`user_ids` IN (?)", userIDss).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 逻辑删除 0 未删除 1 已删除
func (obj *_UContractUseTypeMgr) GetFromIsDelete(isDelete int) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 逻辑删除 0 未删除 1 已删除
func (obj *_UContractUseTypeMgr) GetBatchFromIsDelete(isDeletes []int) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UContractUseTypeMgr) GetFromCreateTime(createTime time.Time) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UContractUseTypeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_UContractUseTypeMgr) GetFromUpdateTime(updateTime time.Time) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 修改时间
func (obj *_UContractUseTypeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_UContractUseTypeMgr) GetFromCreateUser(createUser string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_UContractUseTypeMgr) GetBatchFromCreateUser(createUsers []string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 修改人
func (obj *_UContractUseTypeMgr) GetFromUpdateUser(updateUser string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 修改人
func (obj *_UContractUseTypeMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UContractUseTypeMgr) GetFromVersion(version int) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UContractUseTypeMgr) GetBatchFromVersion(versions []int) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 合同用途备注
func (obj *_UContractUseTypeMgr) GetFromRemark(remark string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 合同用途备注
func (obj *_UContractUseTypeMgr) GetBatchFromRemark(remarks []string) (results []*UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UContractUseTypeMgr) FetchByPrimaryKey(id int64) (result UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByTypeNameIndex primary or index 获取唯一内容
func (obj *_UContractUseTypeMgr) FetchUniqueByTypeNameIndex(typeName string) (result UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`type_name` = ?", typeName).First(&result).Error

	return
}

// FetchUniqueByCodeIndex primary or index 获取唯一内容
func (obj *_UContractUseTypeMgr) FetchUniqueByCodeIndex(code string) (result UContractUseType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContractUseType{}).Where("`code` = ?", code).First(&result).Error

	return
}
