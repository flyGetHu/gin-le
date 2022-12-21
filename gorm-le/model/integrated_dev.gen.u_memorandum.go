package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UMemorandumMgr struct {
	*_BaseMgr
}

// UMemorandumMgr open func
func UMemorandumMgr(db *gorm.DB) *_UMemorandumMgr {
	if db == nil {
		panic(fmt.Errorf("UMemorandumMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UMemorandumMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_memorandum"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UMemorandumMgr) GetTableName() string {
	return "u_memorandum"
}

// Reset 重置gorm会话
func (obj *_UMemorandumMgr) Reset() *_UMemorandumMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UMemorandumMgr) Get() (result UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UMemorandumMgr) Gets() (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UMemorandumMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UMemorandumMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithContent content获取 随手记内容
func (obj *_UMemorandumMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithIsComplete is_complete获取
func (obj *_UMemorandumMgr) WithIsComplete(isComplete []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_complete"] = isComplete })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UMemorandumMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UMemorandumMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UMemorandumMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_UMemorandumMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UMemorandumMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_UMemorandumMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_UMemorandumMgr) GetByOption(opts ...Option) (result UMemorandum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UMemorandumMgr) GetByOptions(opts ...Option) (results []*UMemorandum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UMemorandumMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UMemorandum, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where(options.query)
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
func (obj *_UMemorandumMgr) GetFromID(id int) (result UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UMemorandumMgr) GetBatchFromID(ids []int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 随手记内容
func (obj *_UMemorandumMgr) GetFromContent(content string) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 随手记内容
func (obj *_UMemorandumMgr) GetBatchFromContent(contents []string) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromIsComplete 通过is_complete获取内容
func (obj *_UMemorandumMgr) GetFromIsComplete(isComplete []uint8) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`is_complete` = ?", isComplete).Find(&results).Error

	return
}

// GetBatchFromIsComplete 批量查找
func (obj *_UMemorandumMgr) GetBatchFromIsComplete(isCompletes [][]uint8) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`is_complete` IN (?)", isCompletes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UMemorandumMgr) GetFromCreateTime(createTime time.Time) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UMemorandumMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UMemorandumMgr) GetFromCreateUser(createUser int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UMemorandumMgr) GetBatchFromCreateUser(createUsers []int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UMemorandumMgr) GetFromUpdateTime(updateTime time.Time) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UMemorandumMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_UMemorandumMgr) GetFromUpdateUser(updateUser int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_UMemorandumMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UMemorandumMgr) GetFromVersion(version int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UMemorandumMgr) GetBatchFromVersion(versions []int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_UMemorandumMgr) GetFromDeleted(deleted int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_UMemorandumMgr) GetBatchFromDeleted(deleteds []int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UMemorandumMgr) FetchByPrimaryKey(id int) (result UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByUMemorandumCreateUserIndex  获取多个内容
func (obj *_UMemorandumMgr) FetchIndexByUMemorandumCreateUserIndex(createUser int) (results []*UMemorandum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UMemorandum{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}
