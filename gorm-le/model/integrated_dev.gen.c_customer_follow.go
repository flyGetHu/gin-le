package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CCustomerFollowMgr struct {
	*_BaseMgr
}

// CCustomerFollowMgr open func
func CCustomerFollowMgr(db *gorm.DB) *_CCustomerFollowMgr {
	if db == nil {
		panic(fmt.Errorf("CCustomerFollowMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CCustomerFollowMgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_customer_follow"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CCustomerFollowMgr) GetTableName() string {
	return "c_customer_follow"
}

// Reset 重置gorm会话
func (obj *_CCustomerFollowMgr) Reset() *_CCustomerFollowMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CCustomerFollowMgr) Get() (result CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CCustomerFollowMgr) Gets() (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CCustomerFollowMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CCustomerFollowMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerID customer_id获取 客户id
func (obj *_CCustomerFollowMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithContent content获取 跟进内容
func (obj *_CCustomerFollowMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithStatus status获取 状态，0:无效，1:有效
func (obj *_CCustomerFollowMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CCustomerFollowMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_CCustomerFollowMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CCustomerFollowMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_CCustomerFollowMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_CCustomerFollowMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_CCustomerFollowMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_CCustomerFollowMgr) GetByOption(opts ...Option) (result CCustomerFollow, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CCustomerFollowMgr) GetByOptions(opts ...Option) (results []*CCustomerFollow, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CCustomerFollowMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CCustomerFollow, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where(options.query)
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
func (obj *_CCustomerFollowMgr) GetFromID(id int) (result CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_CCustomerFollowMgr) GetBatchFromID(ids []int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_CCustomerFollowMgr) GetFromCustomerID(customerID int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_CCustomerFollowMgr) GetBatchFromCustomerID(customerIDs []int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 跟进内容
func (obj *_CCustomerFollowMgr) GetFromContent(content string) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 跟进内容
func (obj *_CCustomerFollowMgr) GetBatchFromContent(contents []string) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态，0:无效，1:有效
func (obj *_CCustomerFollowMgr) GetFromStatus(status int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态，0:无效，1:有效
func (obj *_CCustomerFollowMgr) GetBatchFromStatus(statuss []int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CCustomerFollowMgr) GetFromCreateTime(createTime time.Time) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CCustomerFollowMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_CCustomerFollowMgr) GetFromCreateUser(createUser int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_CCustomerFollowMgr) GetBatchFromCreateUser(createUsers []int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CCustomerFollowMgr) GetFromUpdateTime(updateTime time.Time) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_CCustomerFollowMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_CCustomerFollowMgr) GetFromUpdateUser(updateUser int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_CCustomerFollowMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_CCustomerFollowMgr) GetFromVersion(version int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_CCustomerFollowMgr) GetBatchFromVersion(versions []int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_CCustomerFollowMgr) GetFromDeleted(deleted int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_CCustomerFollowMgr) GetBatchFromDeleted(deleteds []int) (results []*CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CCustomerFollowMgr) FetchByPrimaryKey(id int) (result CCustomerFollow, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerFollow{}).Where("`id` = ?", id).First(&result).Error

	return
}
