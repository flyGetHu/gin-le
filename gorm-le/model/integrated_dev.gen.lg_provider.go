package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgProviderMgr struct {
	*_BaseMgr
}

// LgProviderMgr open func
func LgProviderMgr(db *gorm.DB) *_LgProviderMgr {
	if db == nil {
		panic(fmt.Errorf("LgProviderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgProviderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_provider"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgProviderMgr) GetTableName() string {
	return "lg_provider"
}

// Reset 重置gorm会话
func (obj *_LgProviderMgr) Reset() *_LgProviderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgProviderMgr) Get() (result LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgProviderMgr) Gets() (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgProviderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_LgProviderMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithProviderID provider_id获取 服务商code
func (obj *_LgProviderMgr) WithProviderID(providerID int) Option {
	return optionFunc(func(o *options) { o.query["provider_id"] = providerID })
}

// WithProviderName provider_name获取 服务商名称
func (obj *_LgProviderMgr) WithProviderName(providerName string) Option {
	return optionFunc(func(o *options) { o.query["provider_name"] = providerName })
}

// WithAddress address获取 地址
func (obj *_LgProviderMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithWebsite website获取 主页
func (obj *_LgProviderMgr) WithWebsite(website string) Option {
	return optionFunc(func(o *options) { o.query["website"] = website })
}

// WithContact contact获取 联系人
func (obj *_LgProviderMgr) WithContact(contact string) Option {
	return optionFunc(func(o *options) { o.query["contact"] = contact })
}

// WithPhone phone获取 电话
func (obj *_LgProviderMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithEmail email获取 邮箱
func (obj *_LgProviderMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithClearingBank clearing_bank获取 结算银行
func (obj *_LgProviderMgr) WithClearingBank(clearingBank string) Option {
	return optionFunc(func(o *options) { o.query["clearing_bank"] = clearingBank })
}

// WithBankAccount bank_account获取 银行账户
func (obj *_LgProviderMgr) WithBankAccount(bankAccount string) Option {
	return optionFunc(func(o *options) { o.query["bank_account"] = bankAccount })
}

// WithState state获取 状态
func (obj *_LgProviderMgr) WithState(state int) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgProviderMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgProviderMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgProviderMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgProviderMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgProviderMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgProviderMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgProviderMgr) GetByOption(opts ...Option) (result LgProvider, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgProviderMgr) GetByOptions(opts ...Option) (results []*LgProvider, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgProviderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgProvider, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where(options.query)
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
func (obj *_LgProviderMgr) GetFromID(id int) (result LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_LgProviderMgr) GetBatchFromID(ids []int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromProviderID 通过provider_id获取内容 服务商code
func (obj *_LgProviderMgr) GetFromProviderID(providerID int) (result LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`provider_id` = ?", providerID).First(&result).Error

	return
}

// GetBatchFromProviderID 批量查找 服务商code
func (obj *_LgProviderMgr) GetBatchFromProviderID(providerIDs []int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`provider_id` IN (?)", providerIDs).Find(&results).Error

	return
}

// GetFromProviderName 通过provider_name获取内容 服务商名称
func (obj *_LgProviderMgr) GetFromProviderName(providerName string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`provider_name` = ?", providerName).Find(&results).Error

	return
}

// GetBatchFromProviderName 批量查找 服务商名称
func (obj *_LgProviderMgr) GetBatchFromProviderName(providerNames []string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`provider_name` IN (?)", providerNames).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 地址
func (obj *_LgProviderMgr) GetFromAddress(address string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 地址
func (obj *_LgProviderMgr) GetBatchFromAddress(addresss []string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromWebsite 通过website获取内容 主页
func (obj *_LgProviderMgr) GetFromWebsite(website string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`website` = ?", website).Find(&results).Error

	return
}

// GetBatchFromWebsite 批量查找 主页
func (obj *_LgProviderMgr) GetBatchFromWebsite(websites []string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`website` IN (?)", websites).Find(&results).Error

	return
}

// GetFromContact 通过contact获取内容 联系人
func (obj *_LgProviderMgr) GetFromContact(contact string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`contact` = ?", contact).Find(&results).Error

	return
}

// GetBatchFromContact 批量查找 联系人
func (obj *_LgProviderMgr) GetBatchFromContact(contacts []string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`contact` IN (?)", contacts).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 电话
func (obj *_LgProviderMgr) GetFromPhone(phone string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 电话
func (obj *_LgProviderMgr) GetBatchFromPhone(phones []string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_LgProviderMgr) GetFromEmail(email string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 邮箱
func (obj *_LgProviderMgr) GetBatchFromEmail(emails []string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromClearingBank 通过clearing_bank获取内容 结算银行
func (obj *_LgProviderMgr) GetFromClearingBank(clearingBank string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`clearing_bank` = ?", clearingBank).Find(&results).Error

	return
}

// GetBatchFromClearingBank 批量查找 结算银行
func (obj *_LgProviderMgr) GetBatchFromClearingBank(clearingBanks []string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`clearing_bank` IN (?)", clearingBanks).Find(&results).Error

	return
}

// GetFromBankAccount 通过bank_account获取内容 银行账户
func (obj *_LgProviderMgr) GetFromBankAccount(bankAccount string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`bank_account` = ?", bankAccount).Find(&results).Error

	return
}

// GetBatchFromBankAccount 批量查找 银行账户
func (obj *_LgProviderMgr) GetBatchFromBankAccount(bankAccounts []string) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`bank_account` IN (?)", bankAccounts).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态
func (obj *_LgProviderMgr) GetFromState(state int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态
func (obj *_LgProviderMgr) GetBatchFromState(states []int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgProviderMgr) GetFromCreateTime(createTime time.Time) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgProviderMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgProviderMgr) GetFromCreateUser(createUser int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgProviderMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgProviderMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgProviderMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgProviderMgr) GetFromUpdateUser(updateUser int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgProviderMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgProviderMgr) GetFromVersion(version int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgProviderMgr) GetBatchFromVersion(versions []int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgProviderMgr) GetFromDeleted(deleted int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgProviderMgr) GetBatchFromDeleted(deleteds []int) (results []*LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgProviderMgr) FetchByPrimaryKey(id int) (result LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUniqueProviderID primary or index 获取唯一内容
func (obj *_LgProviderMgr) FetchUniqueByUniqueProviderID(providerID int) (result LgProvider, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProvider{}).Where("`provider_id` = ?", providerID).First(&result).Error

	return
}
