package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaPartitionMgr struct {
	*_BaseMgr
}

// FaPartitionMgr open func
func FaPartitionMgr(db *gorm.DB) *_FaPartitionMgr {
	if db == nil {
		panic(fmt.Errorf("FaPartitionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaPartitionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_partition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaPartitionMgr) GetTableName() string {
	return "fa_partition"
}

// Reset 重置gorm会话
func (obj *_FaPartitionMgr) Reset() *_FaPartitionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaPartitionMgr) Get() (result FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaPartitionMgr) Gets() (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaPartitionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaPartitionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *_FaPartitionMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithState state获取 是否启用 0:未启用 1:启用
func (obj *_FaPartitionMgr) WithState(state int) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithRemark remark获取 备注
func (obj *_FaPartitionMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaPartitionMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaPartitionMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaPartitionMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_FaPartitionMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_FaPartitionMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaPartitionMgr) GetByOption(opts ...Option) (result FaPartition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaPartitionMgr) GetByOptions(opts ...Option) (results []*FaPartition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaPartitionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaPartition, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where(options.query)
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
func (obj *_FaPartitionMgr) GetFromID(id int) (result FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaPartitionMgr) GetBatchFromID(ids []int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *_FaPartitionMgr) GetFromName(name string) (result FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *_FaPartitionMgr) GetBatchFromName(names []string) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 是否启用 0:未启用 1:启用
func (obj *_FaPartitionMgr) GetFromState(state int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 是否启用 0:未启用 1:启用
func (obj *_FaPartitionMgr) GetBatchFromState(states []int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaPartitionMgr) GetFromRemark(remark string) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaPartitionMgr) GetBatchFromRemark(remarks []string) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaPartitionMgr) GetFromCreateTime(createTime time.Time) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaPartitionMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaPartitionMgr) GetFromCreateUser(createUser int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaPartitionMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaPartitionMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaPartitionMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_FaPartitionMgr) GetFromUpdateUser(updateUser int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_FaPartitionMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaPartitionMgr) GetFromVersion(version int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaPartitionMgr) GetBatchFromVersion(versions []int) (results []*FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaPartitionMgr) FetchByPrimaryKey(id int) (result FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexName primary or index 获取唯一内容
func (obj *_FaPartitionMgr) FetchUniqueByIndexName(name string) (result FaPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaPartition{}).Where("`name` = ?", name).First(&result).Error

	return
}
