package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _SysServerHealthMgr struct {
	*_BaseMgr
}

// SysServerHealthMgr open func
func SysServerHealthMgr(db *gorm.DB) *_SysServerHealthMgr {
	if db == nil {
		panic(fmt.Errorf("SysServerHealthMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SysServerHealthMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sys_server_health"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SysServerHealthMgr) GetTableName() string {
	return "sys_server_health"
}

// Reset 重置gorm会话
func (obj *_SysServerHealthMgr) Reset() *_SysServerHealthMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SysServerHealthMgr) Get() (result SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_SysServerHealthMgr) Gets() (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SysServerHealthMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_SysServerHealthMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 服务名称
func (obj *_SysServerHealthMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithURL url获取 服务健康检查请求地址
func (obj *_SysServerHealthMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithState state获取 正常:1异常:0
func (obj *_SysServerHealthMgr) WithState(state []uint8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithRemark remark获取 备注
func (obj *_SysServerHealthMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_SysServerHealthMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_SysServerHealthMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_SysServerHealthMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_SysServerHealthMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_SysServerHealthMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_SysServerHealthMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_SysServerHealthMgr) GetByOption(opts ...Option) (result SysServerHealth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SysServerHealthMgr) GetByOptions(opts ...Option) (results []*SysServerHealth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_SysServerHealthMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]SysServerHealth, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where(options.query)
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
func (obj *_SysServerHealthMgr) GetFromID(id int64) (result SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_SysServerHealthMgr) GetBatchFromID(ids []int64) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 服务名称
func (obj *_SysServerHealthMgr) GetFromName(name string) (result SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 服务名称
func (obj *_SysServerHealthMgr) GetBatchFromName(names []string) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 服务健康检查请求地址
func (obj *_SysServerHealthMgr) GetFromURL(url string) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找 服务健康检查请求地址
func (obj *_SysServerHealthMgr) GetBatchFromURL(urls []string) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 正常:1异常:0
func (obj *_SysServerHealthMgr) GetFromState(state []uint8) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 正常:1异常:0
func (obj *_SysServerHealthMgr) GetBatchFromState(states [][]uint8) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_SysServerHealthMgr) GetFromRemark(remark string) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_SysServerHealthMgr) GetBatchFromRemark(remarks []string) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_SysServerHealthMgr) GetFromCreateTime(createTime time.Time) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_SysServerHealthMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_SysServerHealthMgr) GetFromCreateUser(createUser int) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_SysServerHealthMgr) GetBatchFromCreateUser(createUsers []int) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_SysServerHealthMgr) GetFromUpdateTime(updateTime time.Time) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_SysServerHealthMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_SysServerHealthMgr) GetFromUpdateUser(updateUser int) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_SysServerHealthMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_SysServerHealthMgr) GetFromVersion(version int) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_SysServerHealthMgr) GetBatchFromVersion(versions []int) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_SysServerHealthMgr) GetFromDeleted(deleted int) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_SysServerHealthMgr) GetBatchFromDeleted(deleteds []int) (results []*SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SysServerHealthMgr) FetchByPrimaryKey(id int64) (result SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByTServerHealthNameUIndex primary or index 获取唯一内容
func (obj *_SysServerHealthMgr) FetchUniqueByTServerHealthNameUIndex(name string) (result SysServerHealth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SysServerHealth{}).Where("`name` = ?", name).First(&result).Error

	return
}
