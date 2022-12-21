package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgCountryMgr struct {
	*_BaseMgr
}

// LgCountryMgr open func
func LgCountryMgr(db *gorm.DB) *_LgCountryMgr {
	if db == nil {
		panic(fmt.Errorf("LgCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgCountryMgr) GetTableName() string {
	return "lg_country"
}

// Reset 重置gorm会话
func (obj *_LgCountryMgr) Reset() *_LgCountryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgCountryMgr) Get() (result LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgCountryMgr) Gets() (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgCountryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgCountryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNameEn name_en获取 国家中文
func (obj *_LgCountryMgr) WithNameEn(nameEn string) Option {
	return optionFunc(func(o *options) { o.query["name_en"] = nameEn })
}

// WithNameCn name_cn获取 国家英文
func (obj *_LgCountryMgr) WithNameCn(nameCn string) Option {
	return optionFunc(func(o *options) { o.query["name_cn"] = nameCn })
}

// WithTwoCode two_code获取 二字代码
func (obj *_LgCountryMgr) WithTwoCode(twoCode string) Option {
	return optionFunc(func(o *options) { o.query["two_code"] = twoCode })
}

// WithThreeCode three_code获取 三字代码
func (obj *_LgCountryMgr) WithThreeCode(threeCode string) Option {
	return optionFunc(func(o *options) { o.query["three_code"] = threeCode })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgCountryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgCountryMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgCountryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgCountryMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgCountryMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgCountryMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgCountryMgr) GetByOption(opts ...Option) (result LgCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgCountryMgr) GetByOptions(opts ...Option) (results []*LgCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgCountryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgCountry, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where(options.query)
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
func (obj *_LgCountryMgr) GetFromID(id int) (result LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgCountryMgr) GetBatchFromID(ids []int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromNameEn 通过name_en获取内容 国家中文
func (obj *_LgCountryMgr) GetFromNameEn(nameEn string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`name_en` = ?", nameEn).Find(&results).Error

	return
}

// GetBatchFromNameEn 批量查找 国家中文
func (obj *_LgCountryMgr) GetBatchFromNameEn(nameEns []string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`name_en` IN (?)", nameEns).Find(&results).Error

	return
}

// GetFromNameCn 通过name_cn获取内容 国家英文
func (obj *_LgCountryMgr) GetFromNameCn(nameCn string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`name_cn` = ?", nameCn).Find(&results).Error

	return
}

// GetBatchFromNameCn 批量查找 国家英文
func (obj *_LgCountryMgr) GetBatchFromNameCn(nameCns []string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`name_cn` IN (?)", nameCns).Find(&results).Error

	return
}

// GetFromTwoCode 通过two_code获取内容 二字代码
func (obj *_LgCountryMgr) GetFromTwoCode(twoCode string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`two_code` = ?", twoCode).Find(&results).Error

	return
}

// GetBatchFromTwoCode 批量查找 二字代码
func (obj *_LgCountryMgr) GetBatchFromTwoCode(twoCodes []string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`two_code` IN (?)", twoCodes).Find(&results).Error

	return
}

// GetFromThreeCode 通过three_code获取内容 三字代码
func (obj *_LgCountryMgr) GetFromThreeCode(threeCode string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`three_code` = ?", threeCode).Find(&results).Error

	return
}

// GetBatchFromThreeCode 批量查找 三字代码
func (obj *_LgCountryMgr) GetBatchFromThreeCode(threeCodes []string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`three_code` IN (?)", threeCodes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgCountryMgr) GetFromCreateTime(createTime time.Time) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgCountryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgCountryMgr) GetFromCreateUser(createUser int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgCountryMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgCountryMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgCountryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgCountryMgr) GetFromUpdateUser(updateUser int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgCountryMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgCountryMgr) GetFromVersion(version int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgCountryMgr) GetBatchFromVersion(versions []int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgCountryMgr) GetFromDeleted(deleted int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgCountryMgr) GetBatchFromDeleted(deleteds []int) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgCountryMgr) FetchByPrimaryKey(id int) (result LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByTwoCode  获取多个内容
func (obj *_LgCountryMgr) FetchIndexByTwoCode(twoCode string) (results []*LgCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCountry{}).Where("`two_code` = ?", twoCode).Find(&results).Error

	return
}
