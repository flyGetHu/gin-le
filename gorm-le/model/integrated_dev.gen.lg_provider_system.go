package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgProviderSystemMgr struct {
	*_BaseMgr
}

// LgProviderSystemMgr open func
func LgProviderSystemMgr(db *gorm.DB) *_LgProviderSystemMgr {
	if db == nil {
		panic(fmt.Errorf("LgProviderSystemMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgProviderSystemMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_provider_system"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgProviderSystemMgr) GetTableName() string {
	return "lg_provider_system"
}

// Reset 重置gorm会话
func (obj *_LgProviderSystemMgr) Reset() *_LgProviderSystemMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgProviderSystemMgr) Get() (result LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgProviderSystemMgr) Gets() (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgProviderSystemMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgProviderSystemMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 系统名称
func (obj *_LgProviderSystemMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取 系统CODE
func (obj *_LgProviderSystemMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithCheckRule check_rule获取 字段校验规则(正则表达式)
func (obj *_LgProviderSystemMgr) WithCheckRule(checkRule string) Option {
	return optionFunc(func(o *options) { o.query["check_rule"] = checkRule })
}

// WithVariableList variable_list获取 下单字段列表
func (obj *_LgProviderSystemMgr) WithVariableList(variableList string) Option {
	return optionFunc(func(o *options) { o.query["variable_list"] = variableList })
}

// WithContent content获取 附件及留言内容
func (obj *_LgProviderSystemMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgProviderSystemMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgProviderSystemMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgProviderSystemMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgProviderSystemMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgProviderSystemMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgProviderSystemMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgProviderSystemMgr) GetByOption(opts ...Option) (result LgProviderSystem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgProviderSystemMgr) GetByOptions(opts ...Option) (results []*LgProviderSystem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgProviderSystemMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgProviderSystem, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where(options.query)
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
func (obj *_LgProviderSystemMgr) GetFromID(id int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgProviderSystemMgr) GetBatchFromID(ids []int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 系统名称
func (obj *_LgProviderSystemMgr) GetFromName(name string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 系统名称
func (obj *_LgProviderSystemMgr) GetBatchFromName(names []string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 系统CODE
func (obj *_LgProviderSystemMgr) GetFromCode(code string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 系统CODE
func (obj *_LgProviderSystemMgr) GetBatchFromCode(codes []string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromCheckRule 通过check_rule获取内容 字段校验规则(正则表达式)
func (obj *_LgProviderSystemMgr) GetFromCheckRule(checkRule string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`check_rule` = ?", checkRule).Find(&results).Error

	return
}

// GetBatchFromCheckRule 批量查找 字段校验规则(正则表达式)
func (obj *_LgProviderSystemMgr) GetBatchFromCheckRule(checkRules []string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`check_rule` IN (?)", checkRules).Find(&results).Error

	return
}

// GetFromVariableList 通过variable_list获取内容 下单字段列表
func (obj *_LgProviderSystemMgr) GetFromVariableList(variableList string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`variable_list` = ?", variableList).Find(&results).Error

	return
}

// GetBatchFromVariableList 批量查找 下单字段列表
func (obj *_LgProviderSystemMgr) GetBatchFromVariableList(variableLists []string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`variable_list` IN (?)", variableLists).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 附件及留言内容
func (obj *_LgProviderSystemMgr) GetFromContent(content string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 附件及留言内容
func (obj *_LgProviderSystemMgr) GetBatchFromContent(contents []string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgProviderSystemMgr) GetFromCreateTime(createTime time.Time) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgProviderSystemMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgProviderSystemMgr) GetFromCreateUser(createUser int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgProviderSystemMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgProviderSystemMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgProviderSystemMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgProviderSystemMgr) GetFromUpdateUser(updateUser int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgProviderSystemMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgProviderSystemMgr) GetFromVersion(version int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgProviderSystemMgr) GetBatchFromVersion(versions []int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgProviderSystemMgr) GetFromDeleted(deleted int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgProviderSystemMgr) GetBatchFromDeleted(deleteds []int) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgProviderSystemMgr) FetchByPrimaryKey(id int, code string) (result LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`id` = ? AND `code` = ?", id, code).First(&result).Error

	return
}

// FetchUniqueByIndexCode primary or index 获取唯一内容
func (obj *_LgProviderSystemMgr) FetchUniqueByIndexCode(code string) (result LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`code` = ?", code).First(&result).Error

	return
}

// FetchIndexByLgProviderSystemNameIndex  获取多个内容
func (obj *_LgProviderSystemMgr) FetchIndexByLgProviderSystemNameIndex(name string) (results []*LgProviderSystem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderSystem{}).Where("`name` = ?", name).Find(&results).Error

	return
}
