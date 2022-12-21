package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _TUserMgr struct {
	*_BaseMgr
}

// TUserMgr open func
func TUserMgr(db *gorm.DB) *_TUserMgr {
	if db == nil {
		panic(fmt.Errorf("TUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("t_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TUserMgr) GetTableName() string {
	return "t_user"
}

// Reset 重置gorm会话
func (obj *_TUserMgr) Reset() *_TUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TUserMgr) Get() (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TUserMgr) Gets() (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_TUserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 账户名
func (obj *_TUserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码(必须加密)
func (obj *_TUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithRealName real_name获取 真实姓名
func (obj *_TUserMgr) WithRealName(realName string) Option {
	return optionFunc(func(o *options) { o.query["real_name"] = realName })
}

// WithPhone phone获取 手机
func (obj *_TUserMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithCreated created获取 创建时间
func (obj *_TUserMgr) WithCreated(created time.Time) Option {
	return optionFunc(func(o *options) { o.query["created"] = created })
}

// WithUpdated updated获取 更新时间
func (obj *_TUserMgr) WithUpdated(updated time.Time) Option {
	return optionFunc(func(o *options) { o.query["updated"] = updated })
}

// GetByOption 功能选项模式获取
func (obj *_TUserMgr) GetByOption(opts ...Option) (result TUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TUserMgr) GetByOptions(opts ...Option) (results []*TUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TUser{}).Where(options.query)
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
func (obj *_TUserMgr) GetFromID(id int) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_TUserMgr) GetBatchFromID(ids []int) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 账户名
func (obj *_TUserMgr) GetFromUsername(username string) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`username` = ?", username).First(&result).Error

	return
}

// GetBatchFromUsername 批量查找 账户名
func (obj *_TUserMgr) GetBatchFromUsername(usernames []string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码(必须加密)
func (obj *_TUserMgr) GetFromPassword(password string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码(必须加密)
func (obj *_TUserMgr) GetBatchFromPassword(passwords []string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromRealName 通过real_name获取内容 真实姓名
func (obj *_TUserMgr) GetFromRealName(realName string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`real_name` = ?", realName).Find(&results).Error

	return
}

// GetBatchFromRealName 批量查找 真实姓名
func (obj *_TUserMgr) GetBatchFromRealName(realNames []string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`real_name` IN (?)", realNames).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 手机
func (obj *_TUserMgr) GetFromPhone(phone string) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`phone` = ?", phone).First(&result).Error

	return
}

// GetBatchFromPhone 批量查找 手机
func (obj *_TUserMgr) GetBatchFromPhone(phones []string) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromCreated 通过created获取内容 创建时间
func (obj *_TUserMgr) GetFromCreated(created time.Time) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`created` = ?", created).Find(&results).Error

	return
}

// GetBatchFromCreated 批量查找 创建时间
func (obj *_TUserMgr) GetBatchFromCreated(createds []time.Time) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`created` IN (?)", createds).Find(&results).Error

	return
}

// GetFromUpdated 通过updated获取内容 更新时间
func (obj *_TUserMgr) GetFromUpdated(updated time.Time) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`updated` = ?", updated).Find(&results).Error

	return
}

// GetBatchFromUpdated 批量查找 更新时间
func (obj *_TUserMgr) GetBatchFromUpdated(updateds []time.Time) (results []*TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`updated` IN (?)", updateds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TUserMgr) FetchByPrimaryKey(id int) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUsername primary or index 获取唯一内容
func (obj *_TUserMgr) FetchUniqueByUsername(username string) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`username` = ?", username).First(&result).Error

	return
}

// FetchUniqueByUniquePhone primary or index 获取唯一内容
func (obj *_TUserMgr) FetchUniqueByUniquePhone(phone string) (result TUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TUser{}).Where("`phone` = ?", phone).First(&result).Error

	return
}
