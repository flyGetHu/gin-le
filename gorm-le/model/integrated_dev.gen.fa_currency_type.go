package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaCurrencyTypeMgr struct {
	*_BaseMgr
}

// FaCurrencyTypeMgr open func
func FaCurrencyTypeMgr(db *gorm.DB) *_FaCurrencyTypeMgr {
	if db == nil {
		panic(fmt.Errorf("FaCurrencyTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaCurrencyTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_currency_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaCurrencyTypeMgr) GetTableName() string {
	return "fa_currency_type"
}

// Reset 重置gorm会话
func (obj *_FaCurrencyTypeMgr) Reset() *_FaCurrencyTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaCurrencyTypeMgr) Get() (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaCurrencyTypeMgr) Gets() (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaCurrencyTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_FaCurrencyTypeMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 货币编码
func (obj *_FaCurrencyTypeMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithChineseName chinese_name获取 中文名称
func (obj *_FaCurrencyTypeMgr) WithChineseName(chineseName string) Option {
	return optionFunc(func(o *options) { o.query["chinese_name"] = chineseName })
}

// WithEnglishName english_name获取 英文名称
func (obj *_FaCurrencyTypeMgr) WithEnglishName(englishName string) Option {
	return optionFunc(func(o *options) { o.query["english_name"] = englishName })
}

// WithExchangeRate exchange_rate获取 汇率
func (obj *_FaCurrencyTypeMgr) WithExchangeRate(exchangeRate float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_rate"] = exchangeRate })
}

// WithStatus status获取 状态： 1 启用 0 禁用
func (obj *_FaCurrencyTypeMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaCurrencyTypeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 修改时间
func (obj *_FaCurrencyTypeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaCurrencyTypeMgr) WithCreateUser(createUser string) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 修改用户
func (obj *_FaCurrencyTypeMgr) WithUpdateUser(updateUser string) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_FaCurrencyTypeMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithRemark remark获取 备注
func (obj *_FaCurrencyTypeMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// GetByOption 功能选项模式获取
func (obj *_FaCurrencyTypeMgr) GetByOption(opts ...Option) (result FaCurrencyType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaCurrencyTypeMgr) GetByOptions(opts ...Option) (results []*FaCurrencyType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaCurrencyTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaCurrencyType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where(options.query)
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
func (obj *_FaCurrencyTypeMgr) GetFromID(id int64) (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_FaCurrencyTypeMgr) GetBatchFromID(ids []int64) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 货币编码
func (obj *_FaCurrencyTypeMgr) GetFromCode(code string) (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 货币编码
func (obj *_FaCurrencyTypeMgr) GetBatchFromCode(codes []string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromChineseName 通过chinese_name获取内容 中文名称
func (obj *_FaCurrencyTypeMgr) GetFromChineseName(chineseName string) (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`chinese_name` = ?", chineseName).First(&result).Error

	return
}

// GetBatchFromChineseName 批量查找 中文名称
func (obj *_FaCurrencyTypeMgr) GetBatchFromChineseName(chineseNames []string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`chinese_name` IN (?)", chineseNames).Find(&results).Error

	return
}

// GetFromEnglishName 通过english_name获取内容 英文名称
func (obj *_FaCurrencyTypeMgr) GetFromEnglishName(englishName string) (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`english_name` = ?", englishName).First(&result).Error

	return
}

// GetBatchFromEnglishName 批量查找 英文名称
func (obj *_FaCurrencyTypeMgr) GetBatchFromEnglishName(englishNames []string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`english_name` IN (?)", englishNames).Find(&results).Error

	return
}

// GetFromExchangeRate 通过exchange_rate获取内容 汇率
func (obj *_FaCurrencyTypeMgr) GetFromExchangeRate(exchangeRate float64) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`exchange_rate` = ?", exchangeRate).Find(&results).Error

	return
}

// GetBatchFromExchangeRate 批量查找 汇率
func (obj *_FaCurrencyTypeMgr) GetBatchFromExchangeRate(exchangeRates []float64) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`exchange_rate` IN (?)", exchangeRates).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态： 1 启用 0 禁用
func (obj *_FaCurrencyTypeMgr) GetFromStatus(status int) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态： 1 启用 0 禁用
func (obj *_FaCurrencyTypeMgr) GetBatchFromStatus(statuss []int) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaCurrencyTypeMgr) GetFromCreateTime(createTime time.Time) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaCurrencyTypeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 修改时间
func (obj *_FaCurrencyTypeMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 修改时间
func (obj *_FaCurrencyTypeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaCurrencyTypeMgr) GetFromCreateUser(createUser string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaCurrencyTypeMgr) GetBatchFromCreateUser(createUsers []string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 修改用户
func (obj *_FaCurrencyTypeMgr) GetFromUpdateUser(updateUser string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 修改用户
func (obj *_FaCurrencyTypeMgr) GetBatchFromUpdateUser(updateUsers []string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaCurrencyTypeMgr) GetFromVersion(version int) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaCurrencyTypeMgr) GetBatchFromVersion(versions []int) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaCurrencyTypeMgr) GetFromRemark(remark string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaCurrencyTypeMgr) GetBatchFromRemark(remarks []string) (results []*FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaCurrencyTypeMgr) FetchByPrimaryKey(id int64) (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByCodeIndex primary or index 获取唯一内容
func (obj *_FaCurrencyTypeMgr) FetchUniqueByCodeIndex(code string) (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`code` = ?", code).First(&result).Error

	return
}

// FetchUniqueByZhIndex primary or index 获取唯一内容
func (obj *_FaCurrencyTypeMgr) FetchUniqueByZhIndex(chineseName string) (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`chinese_name` = ?", chineseName).First(&result).Error

	return
}

// FetchUniqueByEnIndex primary or index 获取唯一内容
func (obj *_FaCurrencyTypeMgr) FetchUniqueByEnIndex(englishName string) (result FaCurrencyType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCurrencyType{}).Where("`english_name` = ?", englishName).First(&result).Error

	return
}
