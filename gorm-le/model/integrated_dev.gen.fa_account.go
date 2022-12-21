package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaAccountMgr struct {
	*_BaseMgr
}

// FaAccountMgr open func
func FaAccountMgr(db *gorm.DB) *_FaAccountMgr {
	if db == nil {
		panic(fmt.Errorf("FaAccountMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaAccountMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_account"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaAccountMgr) GetTableName() string {
	return "fa_account"
}

// Reset 重置gorm会话
func (obj *_FaAccountMgr) Reset() *_FaAccountMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaAccountMgr) Get() (result FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaAccountMgr) Gets() (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaAccountMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_FaAccountMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAccountAlias account_alias获取 账户别称
func (obj *_FaAccountMgr) WithAccountAlias(accountAlias string) Option {
	return optionFunc(func(o *options) { o.query["account_alias"] = accountAlias })
}

// WithStatus status获取 状态 启用/禁用
func (obj *_FaAccountMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBank bank获取 开户行
func (obj *_FaAccountMgr) WithBank(bank string) Option {
	return optionFunc(func(o *options) { o.query["bank"] = bank })
}

// WithAccountNumber account_number获取 账户号码
func (obj *_FaAccountMgr) WithAccountNumber(accountNumber string) Option {
	return optionFunc(func(o *options) { o.query["account_number"] = accountNumber })
}

// WithAccountName account_name获取 账户人名称
func (obj *_FaAccountMgr) WithAccountName(accountName string) Option {
	return optionFunc(func(o *options) { o.query["account_name"] = accountName })
}

// WithRemark remark获取 备注
func (obj *_FaAccountMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaAccountMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaAccountMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaAccountMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_FaAccountMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_FaAccountMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaAccountMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaAccountMgr) GetByOption(opts ...Option) (result FaAccount, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaAccountMgr) GetByOptions(opts ...Option) (results []*FaAccount, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaAccountMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaAccount, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where(options.query)
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
func (obj *_FaAccountMgr) GetFromID(id int) (result FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_FaAccountMgr) GetBatchFromID(ids []int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAccountAlias 通过account_alias获取内容 账户别称
func (obj *_FaAccountMgr) GetFromAccountAlias(accountAlias string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`account_alias` = ?", accountAlias).Find(&results).Error

	return
}

// GetBatchFromAccountAlias 批量查找 账户别称
func (obj *_FaAccountMgr) GetBatchFromAccountAlias(accountAliass []string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`account_alias` IN (?)", accountAliass).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态 启用/禁用
func (obj *_FaAccountMgr) GetFromStatus(status int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态 启用/禁用
func (obj *_FaAccountMgr) GetBatchFromStatus(statuss []int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBank 通过bank获取内容 开户行
func (obj *_FaAccountMgr) GetFromBank(bank string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`bank` = ?", bank).Find(&results).Error

	return
}

// GetBatchFromBank 批量查找 开户行
func (obj *_FaAccountMgr) GetBatchFromBank(banks []string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`bank` IN (?)", banks).Find(&results).Error

	return
}

// GetFromAccountNumber 通过account_number获取内容 账户号码
func (obj *_FaAccountMgr) GetFromAccountNumber(accountNumber string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`account_number` = ?", accountNumber).Find(&results).Error

	return
}

// GetBatchFromAccountNumber 批量查找 账户号码
func (obj *_FaAccountMgr) GetBatchFromAccountNumber(accountNumbers []string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`account_number` IN (?)", accountNumbers).Find(&results).Error

	return
}

// GetFromAccountName 通过account_name获取内容 账户人名称
func (obj *_FaAccountMgr) GetFromAccountName(accountName string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`account_name` = ?", accountName).Find(&results).Error

	return
}

// GetBatchFromAccountName 批量查找 账户人名称
func (obj *_FaAccountMgr) GetBatchFromAccountName(accountNames []string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`account_name` IN (?)", accountNames).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaAccountMgr) GetFromRemark(remark string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaAccountMgr) GetBatchFromRemark(remarks []string) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaAccountMgr) GetFromCreateTime(createTime time.Time) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaAccountMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaAccountMgr) GetFromCreateUser(createUser int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaAccountMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaAccountMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaAccountMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_FaAccountMgr) GetFromUpdateUser(updateUser int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_FaAccountMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaAccountMgr) GetFromVersion(version int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaAccountMgr) GetBatchFromVersion(versions []int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaAccountMgr) GetFromDeleted(deleted int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaAccountMgr) GetBatchFromDeleted(deleteds []int) (results []*FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaAccountMgr) FetchByPrimaryKey(id int) (result FaAccount, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaAccount{}).Where("`id` = ?", id).First(&result).Error

	return
}
