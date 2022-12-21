package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgChannelMgr struct {
	*_BaseMgr
}

// LgChannelMgr open func
func LgChannelMgr(db *gorm.DB) *_LgChannelMgr {
	if db == nil {
		panic(fmt.Errorf("LgChannelMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgChannelMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_channel"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgChannelMgr) GetTableName() string {
	return "lg_channel"
}

// Reset 重置gorm会话
func (obj *_LgChannelMgr) Reset() *_LgChannelMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgChannelMgr) Get() (result LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgChannelMgr) Gets() (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgChannelMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgChannelMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithChannelID channel_id获取 渠道ID
func (obj *_LgChannelMgr) WithChannelID(channelID int) Option {
	return optionFunc(func(o *options) { o.query["channel_id"] = channelID })
}

// WithChannelNameA channel_name_a获取 渠道名称
func (obj *_LgChannelMgr) WithChannelNameA(channelNameA string) Option {
	return optionFunc(func(o *options) { o.query["channel_name_a"] = channelNameA })
}

// WithChannelNameB channel_name_b获取 客户渠道名称
func (obj *_LgChannelMgr) WithChannelNameB(channelNameB string) Option {
	return optionFunc(func(o *options) { o.query["channel_name_b"] = channelNameB })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgChannelMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgChannelMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgChannelMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgChannelMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgChannelMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgChannelMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgChannelMgr) GetByOption(opts ...Option) (result LgChannel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgChannelMgr) GetByOptions(opts ...Option) (results []*LgChannel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgChannelMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgChannel, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where(options.query)
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

// GetFromID 通过id获取内容 ID
func (obj *_LgChannelMgr) GetFromID(id int) (result LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgChannelMgr) GetBatchFromID(ids []int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromChannelID 通过channel_id获取内容 渠道ID
func (obj *_LgChannelMgr) GetFromChannelID(channelID int) (result LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`channel_id` = ?", channelID).First(&result).Error

	return
}

// GetBatchFromChannelID 批量查找 渠道ID
func (obj *_LgChannelMgr) GetBatchFromChannelID(channelIDs []int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`channel_id` IN (?)", channelIDs).Find(&results).Error

	return
}

// GetFromChannelNameA 通过channel_name_a获取内容 渠道名称
func (obj *_LgChannelMgr) GetFromChannelNameA(channelNameA string) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`channel_name_a` = ?", channelNameA).Find(&results).Error

	return
}

// GetBatchFromChannelNameA 批量查找 渠道名称
func (obj *_LgChannelMgr) GetBatchFromChannelNameA(channelNameAs []string) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`channel_name_a` IN (?)", channelNameAs).Find(&results).Error

	return
}

// GetFromChannelNameB 通过channel_name_b获取内容 客户渠道名称
func (obj *_LgChannelMgr) GetFromChannelNameB(channelNameB string) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`channel_name_b` = ?", channelNameB).Find(&results).Error

	return
}

// GetBatchFromChannelNameB 批量查找 客户渠道名称
func (obj *_LgChannelMgr) GetBatchFromChannelNameB(channelNameBs []string) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`channel_name_b` IN (?)", channelNameBs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgChannelMgr) GetFromCreateTime(createTime time.Time) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgChannelMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgChannelMgr) GetFromCreateUser(createUser int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgChannelMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgChannelMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgChannelMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgChannelMgr) GetFromUpdateUser(updateUser int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgChannelMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgChannelMgr) GetFromVersion(version int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgChannelMgr) GetBatchFromVersion(versions []int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgChannelMgr) GetFromDeleted(deleted int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgChannelMgr) GetBatchFromDeleted(deleteds []int) (results []*LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgChannelMgr) FetchByPrimaryKey(id int) (result LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUniqueChannelID primary or index 获取唯一内容
func (obj *_LgChannelMgr) FetchUniqueByUniqueChannelID(channelID int) (result LgChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannel{}).Where("`channel_id` = ?", channelID).First(&result).Error

	return
}
