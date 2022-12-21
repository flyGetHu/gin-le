package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgChannelFollowMgr struct {
	*_BaseMgr
}

// LgChannelFollowMgr open func
func LgChannelFollowMgr(db *gorm.DB) *_LgChannelFollowMgr {
	if db == nil {
		panic(fmt.Errorf("LgChannelFollowMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgChannelFollowMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_channel_follow"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgChannelFollowMgr) GetTableName() string {
	return "lg_channel_follow"
}

// Reset 重置gorm会话
func (obj *_LgChannelFollowMgr) Reset() *_LgChannelFollowMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgChannelFollowMgr) Get() (result LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgChannelFollowMgr) Gets() (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgChannelFollowMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgChannelFollowMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerID customer_id获取 客户ID
func (obj *_LgChannelFollowMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithChannelType channel_type获取 渠道类型
func (obj *_LgChannelFollowMgr) WithChannelType(channelType string) Option {
	return optionFunc(func(o *options) { o.query["channel_type"] = channelType })
}

// WithChannelID channel_id获取 渠道id
func (obj *_LgChannelFollowMgr) WithChannelID(channelID int) Option {
	return optionFunc(func(o *options) { o.query["channel_id"] = channelID })
}

// WithContent content获取 留言内容
func (obj *_LgChannelFollowMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithStatus status获取 状态，0:无效，1:有效
func (obj *_LgChannelFollowMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgChannelFollowMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgChannelFollowMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgChannelFollowMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgChannelFollowMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgChannelFollowMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgChannelFollowMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgChannelFollowMgr) GetByOption(opts ...Option) (result LgChannelFollow, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgChannelFollowMgr) GetByOptions(opts ...Option) (results []*LgChannelFollow, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgChannelFollowMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgChannelFollow, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where(options.query)
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
func (obj *_LgChannelFollowMgr) GetFromID(id int) (result LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgChannelFollowMgr) GetBatchFromID(ids []int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户ID
func (obj *_LgChannelFollowMgr) GetFromCustomerID(customerID int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户ID
func (obj *_LgChannelFollowMgr) GetBatchFromCustomerID(customerIDs []int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromChannelType 通过channel_type获取内容 渠道类型
func (obj *_LgChannelFollowMgr) GetFromChannelType(channelType string) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`channel_type` = ?", channelType).Find(&results).Error

	return
}

// GetBatchFromChannelType 批量查找 渠道类型
func (obj *_LgChannelFollowMgr) GetBatchFromChannelType(channelTypes []string) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`channel_type` IN (?)", channelTypes).Find(&results).Error

	return
}

// GetFromChannelID 通过channel_id获取内容 渠道id
func (obj *_LgChannelFollowMgr) GetFromChannelID(channelID int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`channel_id` = ?", channelID).Find(&results).Error

	return
}

// GetBatchFromChannelID 批量查找 渠道id
func (obj *_LgChannelFollowMgr) GetBatchFromChannelID(channelIDs []int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`channel_id` IN (?)", channelIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 留言内容
func (obj *_LgChannelFollowMgr) GetFromContent(content string) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 留言内容
func (obj *_LgChannelFollowMgr) GetBatchFromContent(contents []string) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态，0:无效，1:有效
func (obj *_LgChannelFollowMgr) GetFromStatus(status int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态，0:无效，1:有效
func (obj *_LgChannelFollowMgr) GetBatchFromStatus(statuss []int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgChannelFollowMgr) GetFromCreateTime(createTime time.Time) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgChannelFollowMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgChannelFollowMgr) GetFromCreateUser(createUser int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgChannelFollowMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgChannelFollowMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgChannelFollowMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgChannelFollowMgr) GetFromUpdateUser(updateUser int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgChannelFollowMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgChannelFollowMgr) GetFromVersion(version int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgChannelFollowMgr) GetBatchFromVersion(versions []int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgChannelFollowMgr) GetFromDeleted(deleted int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgChannelFollowMgr) GetBatchFromDeleted(deleteds []int) (results []*LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgChannelFollowMgr) FetchByPrimaryKey(id int) (result LgChannelFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelFollow{}).Where("`id` = ?", id).First(&result).Error

	return
}
