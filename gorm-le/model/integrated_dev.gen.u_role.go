package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _URoleMgr struct {
	*_BaseMgr
}

// URoleMgr open func
func URoleMgr(db *gorm.DB) *_URoleMgr {
	if db == nil {
		panic(fmt.Errorf("URoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_URoleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_URoleMgr) GetTableName() string {
	return "u_role"
}

// Reset 重置gorm会话
func (obj *_URoleMgr) Reset() *_URoleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_URoleMgr) Get() (result URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_URoleMgr) Gets() (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_URoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(URole{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_URoleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRoleName role_name获取 角色名称
func (obj *_URoleMgr) WithRoleName(roleName string) Option {
	return optionFunc(func(o *options) { o.query["role_name"] = roleName })
}

// WithRemark remark获取 备注
func (obj *_URoleMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_URoleMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_URoleMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_URoleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_URoleMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_URoleMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_URoleMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_URoleMgr) GetByOption(opts ...Option) (result URole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_URoleMgr) GetByOptions(opts ...Option) (results []*URole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_URoleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]URole, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(URole{}).Where(options.query)
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
func (obj *_URoleMgr) GetFromID(id int) (result URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_URoleMgr) GetBatchFromID(ids []int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromRoleName 通过role_name获取内容 角色名称
func (obj *_URoleMgr) GetFromRoleName(roleName string) (result URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`role_name` = ?", roleName).First(&result).Error

	return
}

// GetBatchFromRoleName 批量查找 角色名称
func (obj *_URoleMgr) GetBatchFromRoleName(roleNames []string) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`role_name` IN (?)", roleNames).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_URoleMgr) GetFromRemark(remark string) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_URoleMgr) GetBatchFromRemark(remarks []string) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_URoleMgr) GetFromCreateTime(createTime time.Time) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_URoleMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_URoleMgr) GetFromCreateUser(createUser int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_URoleMgr) GetBatchFromCreateUser(createUsers []int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_URoleMgr) GetFromUpdateTime(updateTime time.Time) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_URoleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_URoleMgr) GetFromUpdateUser(updateUser int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_URoleMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_URoleMgr) GetFromVersion(version int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_URoleMgr) GetBatchFromVersion(versions []int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_URoleMgr) GetFromDeleted(deleted int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_URoleMgr) GetBatchFromDeleted(deleteds []int) (results []*URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_URoleMgr) FetchByPrimaryKey(id int) (result URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexRoleName primary or index 获取唯一内容
func (obj *_URoleMgr) FetchUniqueByIndexRoleName(roleName string) (result URole, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(URole{}).Where("`role_name` = ?", roleName).First(&result).Error

	return
}
