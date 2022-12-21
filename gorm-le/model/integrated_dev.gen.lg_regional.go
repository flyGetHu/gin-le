package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgRegionalMgr struct {
	*_BaseMgr
}

// LgRegionalMgr open func
func LgRegionalMgr(db *gorm.DB) *_LgRegionalMgr {
	if db == nil {
		panic(fmt.Errorf("LgRegionalMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgRegionalMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_regional"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgRegionalMgr) GetTableName() string {
	return "lg_regional"
}

// Reset 重置gorm会话
func (obj *_LgRegionalMgr) Reset() *_LgRegionalMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgRegionalMgr) Get() (result LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgRegionalMgr) Gets() (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgRegionalMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgRegionalMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAreaID area_id获取 地区id
func (obj *_LgRegionalMgr) WithAreaID(areaID int) Option {
	return optionFunc(func(o *options) { o.query["area_id"] = areaID })
}

// WithAreaNameCn area_name_cn获取 地区中文名
func (obj *_LgRegionalMgr) WithAreaNameCn(areaNameCn string) Option {
	return optionFunc(func(o *options) { o.query["area_name_cn"] = areaNameCn })
}

// WithAreaTwoCode area_two_code获取 地区二字码
func (obj *_LgRegionalMgr) WithAreaTwoCode(areaTwoCode string) Option {
	return optionFunc(func(o *options) { o.query["area_two_code"] = areaTwoCode })
}

// WithNameCn name_cn获取 中文名
func (obj *_LgRegionalMgr) WithNameCn(nameCn string) Option {
	return optionFunc(func(o *options) { o.query["name_cn"] = nameCn })
}

// WithNameEn name_en获取 英文名
func (obj *_LgRegionalMgr) WithNameEn(nameEn string) Option {
	return optionFunc(func(o *options) { o.query["name_en"] = nameEn })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgRegionalMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgRegionalMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgRegionalMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_LgRegionalMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgRegionalMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgRegionalMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgRegionalMgr) GetByOption(opts ...Option) (result LgRegional, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgRegionalMgr) GetByOptions(opts ...Option) (results []*LgRegional, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgRegionalMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgRegional, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where(options.query)
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
func (obj *_LgRegionalMgr) GetFromID(id int) (result LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_LgRegionalMgr) GetBatchFromID(ids []int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAreaID 通过area_id获取内容 地区id
func (obj *_LgRegionalMgr) GetFromAreaID(areaID int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`area_id` = ?", areaID).Find(&results).Error

	return
}

// GetBatchFromAreaID 批量查找 地区id
func (obj *_LgRegionalMgr) GetBatchFromAreaID(areaIDs []int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`area_id` IN (?)", areaIDs).Find(&results).Error

	return
}

// GetFromAreaNameCn 通过area_name_cn获取内容 地区中文名
func (obj *_LgRegionalMgr) GetFromAreaNameCn(areaNameCn string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`area_name_cn` = ?", areaNameCn).Find(&results).Error

	return
}

// GetBatchFromAreaNameCn 批量查找 地区中文名
func (obj *_LgRegionalMgr) GetBatchFromAreaNameCn(areaNameCns []string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`area_name_cn` IN (?)", areaNameCns).Find(&results).Error

	return
}

// GetFromAreaTwoCode 通过area_two_code获取内容 地区二字码
func (obj *_LgRegionalMgr) GetFromAreaTwoCode(areaTwoCode string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`area_two_code` = ?", areaTwoCode).Find(&results).Error

	return
}

// GetBatchFromAreaTwoCode 批量查找 地区二字码
func (obj *_LgRegionalMgr) GetBatchFromAreaTwoCode(areaTwoCodes []string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`area_two_code` IN (?)", areaTwoCodes).Find(&results).Error

	return
}

// GetFromNameCn 通过name_cn获取内容 中文名
func (obj *_LgRegionalMgr) GetFromNameCn(nameCn string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`name_cn` = ?", nameCn).Find(&results).Error

	return
}

// GetBatchFromNameCn 批量查找 中文名
func (obj *_LgRegionalMgr) GetBatchFromNameCn(nameCns []string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`name_cn` IN (?)", nameCns).Find(&results).Error

	return
}

// GetFromNameEn 通过name_en获取内容 英文名
func (obj *_LgRegionalMgr) GetFromNameEn(nameEn string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`name_en` = ?", nameEn).Find(&results).Error

	return
}

// GetBatchFromNameEn 批量查找 英文名
func (obj *_LgRegionalMgr) GetBatchFromNameEn(nameEns []string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`name_en` IN (?)", nameEns).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgRegionalMgr) GetFromCreateTime(createTime time.Time) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgRegionalMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgRegionalMgr) GetFromCreateUser(createUser int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgRegionalMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgRegionalMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgRegionalMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_LgRegionalMgr) GetFromUpdateUser(updateUser int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_LgRegionalMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgRegionalMgr) GetFromVersion(version int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgRegionalMgr) GetBatchFromVersion(versions []int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgRegionalMgr) GetFromDeleted(deleted int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgRegionalMgr) GetBatchFromDeleted(deleteds []int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgRegionalMgr) FetchByPrimaryKey(id int) (result LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByLgAreaIDIndex  获取多个内容
func (obj *_LgRegionalMgr) FetchIndexByLgAreaIDIndex(areaID int) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`area_id` = ?", areaID).Find(&results).Error

	return
}

// FetchIndexByLgAreaNameCnIndex  获取多个内容
func (obj *_LgRegionalMgr) FetchIndexByLgAreaNameCnIndex(areaNameCn string) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`area_name_cn` = ?", areaNameCn).Find(&results).Error

	return
}

// FetchIndexByLgRegionalCreateTime  获取多个内容
func (obj *_LgRegionalMgr) FetchIndexByLgRegionalCreateTime(createTime time.Time) (results []*LgRegional, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRegional{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
