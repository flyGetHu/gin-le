package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderNumSegmentMgr struct {
	*_BaseMgr
}

// LgOrderNumSegmentMgr open func
func LgOrderNumSegmentMgr(db *gorm.DB) *_LgOrderNumSegmentMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderNumSegmentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderNumSegmentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_num_segment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderNumSegmentMgr) GetTableName() string {
	return "lg_order_num_segment"
}

// Reset 重置gorm会话
func (obj *_LgOrderNumSegmentMgr) Reset() *_LgOrderNumSegmentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderNumSegmentMgr) Get() (result LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderNumSegmentMgr) Gets() (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderNumSegmentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 号段主表ID
func (obj *_LgOrderNumSegmentMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 号段名称
func (obj *_LgOrderNumSegmentMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrefix prefix获取 号段前缀
func (obj *_LgOrderNumSegmentMgr) WithPrefix(prefix string) Option {
	return optionFunc(func(o *options) { o.query["prefix"] = prefix })
}

// WithSuffix suffix获取 号段后缀
func (obj *_LgOrderNumSegmentMgr) WithSuffix(suffix string) Option {
	return optionFunc(func(o *options) { o.query["suffix"] = suffix })
}

// WithStartNum start_num获取 开始号码
func (obj *_LgOrderNumSegmentMgr) WithStartNum(startNum int64) Option {
	return optionFunc(func(o *options) { o.query["start_num"] = startNum })
}

// WithEndNum end_num获取 结束号码
func (obj *_LgOrderNumSegmentMgr) WithEndNum(endNum int64) Option {
	return optionFunc(func(o *options) { o.query["end_num"] = endNum })
}

// WithVerificationMode verification_mode获取 验证号码合法方式:保存验证方式code
func (obj *_LgOrderNumSegmentMgr) WithVerificationMode(verificationMode string) Option {
	return optionFunc(func(o *options) { o.query["verification_mode"] = verificationMode })
}

// WithUsedNum used_num获取 已用到号码段位置:默认喂号码开始位置
func (obj *_LgOrderNumSegmentMgr) WithUsedNum(usedNum int64) Option {
	return optionFunc(func(o *options) { o.query["used_num"] = usedNum })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderNumSegmentMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderNumSegmentMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderNumSegmentMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderNumSegmentMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderNumSegmentMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderNumSegmentMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderNumSegmentMgr) GetByOption(opts ...Option) (result LgOrderNumSegment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderNumSegmentMgr) GetByOptions(opts ...Option) (results []*LgOrderNumSegment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderNumSegmentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderNumSegment, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where(options.query)
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

// GetFromID 通过id获取内容 号段主表ID
func (obj *_LgOrderNumSegmentMgr) GetFromID(id int64) (result LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 号段主表ID
func (obj *_LgOrderNumSegmentMgr) GetBatchFromID(ids []int64) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 号段名称
func (obj *_LgOrderNumSegmentMgr) GetFromName(name string) (result LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 号段名称
func (obj *_LgOrderNumSegmentMgr) GetBatchFromName(names []string) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPrefix 通过prefix获取内容 号段前缀
func (obj *_LgOrderNumSegmentMgr) GetFromPrefix(prefix string) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`prefix` = ?", prefix).Find(&results).Error

	return
}

// GetBatchFromPrefix 批量查找 号段前缀
func (obj *_LgOrderNumSegmentMgr) GetBatchFromPrefix(prefixs []string) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`prefix` IN (?)", prefixs).Find(&results).Error

	return
}

// GetFromSuffix 通过suffix获取内容 号段后缀
func (obj *_LgOrderNumSegmentMgr) GetFromSuffix(suffix string) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`suffix` = ?", suffix).Find(&results).Error

	return
}

// GetBatchFromSuffix 批量查找 号段后缀
func (obj *_LgOrderNumSegmentMgr) GetBatchFromSuffix(suffixs []string) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`suffix` IN (?)", suffixs).Find(&results).Error

	return
}

// GetFromStartNum 通过start_num获取内容 开始号码
func (obj *_LgOrderNumSegmentMgr) GetFromStartNum(startNum int64) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`start_num` = ?", startNum).Find(&results).Error

	return
}

// GetBatchFromStartNum 批量查找 开始号码
func (obj *_LgOrderNumSegmentMgr) GetBatchFromStartNum(startNums []int64) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`start_num` IN (?)", startNums).Find(&results).Error

	return
}

// GetFromEndNum 通过end_num获取内容 结束号码
func (obj *_LgOrderNumSegmentMgr) GetFromEndNum(endNum int64) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`end_num` = ?", endNum).Find(&results).Error

	return
}

// GetBatchFromEndNum 批量查找 结束号码
func (obj *_LgOrderNumSegmentMgr) GetBatchFromEndNum(endNums []int64) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`end_num` IN (?)", endNums).Find(&results).Error

	return
}

// GetFromVerificationMode 通过verification_mode获取内容 验证号码合法方式:保存验证方式code
func (obj *_LgOrderNumSegmentMgr) GetFromVerificationMode(verificationMode string) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`verification_mode` = ?", verificationMode).Find(&results).Error

	return
}

// GetBatchFromVerificationMode 批量查找 验证号码合法方式:保存验证方式code
func (obj *_LgOrderNumSegmentMgr) GetBatchFromVerificationMode(verificationModes []string) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`verification_mode` IN (?)", verificationModes).Find(&results).Error

	return
}

// GetFromUsedNum 通过used_num获取内容 已用到号码段位置:默认喂号码开始位置
func (obj *_LgOrderNumSegmentMgr) GetFromUsedNum(usedNum int64) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`used_num` = ?", usedNum).Find(&results).Error

	return
}

// GetBatchFromUsedNum 批量查找 已用到号码段位置:默认喂号码开始位置
func (obj *_LgOrderNumSegmentMgr) GetBatchFromUsedNum(usedNums []int64) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`used_num` IN (?)", usedNums).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderNumSegmentMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderNumSegmentMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderNumSegmentMgr) GetFromCreateUser(createUser int) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderNumSegmentMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderNumSegmentMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderNumSegmentMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderNumSegmentMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderNumSegmentMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderNumSegmentMgr) GetFromVersion(version int) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderNumSegmentMgr) GetBatchFromVersion(versions []int) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderNumSegmentMgr) GetFromDeleted(deleted int) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderNumSegmentMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderNumSegmentMgr) FetchByPrimaryKey(id int64) (result LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgOrderNumSegmentNameUIndex primary or index 获取唯一内容
func (obj *_LgOrderNumSegmentMgr) FetchUniqueByLgOrderNumSegmentNameUIndex(name string) (result LgOrderNumSegment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegment{}).Where("`name` = ?", name).First(&result).Error

	return
}
