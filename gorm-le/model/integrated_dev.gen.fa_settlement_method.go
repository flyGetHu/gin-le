package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaSettlementMethodMgr struct {
	*_BaseMgr
}

// FaSettlementMethodMgr open func
func FaSettlementMethodMgr(db *gorm.DB) *_FaSettlementMethodMgr {
	if db == nil {
		panic(fmt.Errorf("FaSettlementMethodMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaSettlementMethodMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_settlement_method"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaSettlementMethodMgr) GetTableName() string {
	return "fa_settlement_method"
}

// Reset 重置gorm会话
func (obj *_FaSettlementMethodMgr) Reset() *_FaSettlementMethodMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaSettlementMethodMgr) Get() (result FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaSettlementMethodMgr) Gets() (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaSettlementMethodMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_FaSettlementMethodMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMethod method获取 结算方式
func (obj *_FaSettlementMethodMgr) WithMethod(method string) Option {
	return optionFunc(func(o *options) { o.query["method"] = method })
}

// WithState state获取 状态（0:禁用1:启用）
func (obj *_FaSettlementMethodMgr) WithState(state []uint8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithRemark remark获取 备注
func (obj *_FaSettlementMethodMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateUser create_user获取 创建人
func (obj *_FaSettlementMethodMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaSettlementMethodMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaSettlementMethodMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaSettlementMethodMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaSettlementMethodMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaSettlementMethodMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaSettlementMethodMgr) GetByOption(opts ...Option) (result FaSettlementMethod, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaSettlementMethodMgr) GetByOptions(opts ...Option) (results []*FaSettlementMethod, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaSettlementMethodMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaSettlementMethod, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *_FaSettlementMethodMgr) GetFromID(id int) (result FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_FaSettlementMethodMgr) GetBatchFromID(ids []int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMethod 通过method获取内容 结算方式
func (obj *_FaSettlementMethodMgr) GetFromMethod(method string) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`method` = ?", method).Find(&results).Error

	return
}

// GetBatchFromMethod 批量查找 结算方式
func (obj *_FaSettlementMethodMgr) GetBatchFromMethod(methods []string) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`method` IN (?)", methods).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态（0:禁用1:启用）
func (obj *_FaSettlementMethodMgr) GetFromState(state []uint8) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态（0:禁用1:启用）
func (obj *_FaSettlementMethodMgr) GetBatchFromState(states [][]uint8) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaSettlementMethodMgr) GetFromRemark(remark string) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaSettlementMethodMgr) GetBatchFromRemark(remarks []string) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_FaSettlementMethodMgr) GetFromCreateUser(createUser int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_FaSettlementMethodMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaSettlementMethodMgr) GetFromCreateTime(createTime time.Time) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaSettlementMethodMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaSettlementMethodMgr) GetFromUpdateUser(updateUser int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaSettlementMethodMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaSettlementMethodMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaSettlementMethodMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaSettlementMethodMgr) GetFromVersion(version int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaSettlementMethodMgr) GetBatchFromVersion(versions []int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaSettlementMethodMgr) GetFromDeleted(deleted int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaSettlementMethodMgr) GetBatchFromDeleted(deleteds []int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaSettlementMethodMgr) FetchByPrimaryKey(id int) (result FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByFaSettlementMethodIDUIndex primary or index 获取唯一内容
func (obj *_FaSettlementMethodMgr) FetchUniqueByFaSettlementMethodIDUIndex(id int) (result FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByFaSettlementMethodMethodIndex  获取多个内容
func (obj *_FaSettlementMethodMgr) FetchIndexByFaSettlementMethodMethodIndex(method string) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`method` = ?", method).Find(&results).Error

	return
}

// FetchIndexByFaSettlementMethodStateIndex  获取多个内容
func (obj *_FaSettlementMethodMgr) FetchIndexByFaSettlementMethodStateIndex(state []uint8) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// FetchIndexByFaSettlementMethodCreateUserIndex  获取多个内容
func (obj *_FaSettlementMethodMgr) FetchIndexByFaSettlementMethodCreateUserIndex(createUser int) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// FetchIndexByFaSettlementMethodCreateTimeIndex  获取多个内容
func (obj *_FaSettlementMethodMgr) FetchIndexByFaSettlementMethodCreateTimeIndex(createTime time.Time) (results []*FaSettlementMethod, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaSettlementMethod{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
