package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgNumPoolMgr struct {
	*_BaseMgr
}

// LgNumPoolMgr open func
func LgNumPoolMgr(db *gorm.DB) *_LgNumPoolMgr {
	if db == nil {
		panic(fmt.Errorf("LgNumPoolMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgNumPoolMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_num_pool"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgNumPoolMgr) GetTableName() string {
	return "lg_num_pool"
}

// Reset 重置gorm会话
func (obj *_LgNumPoolMgr) Reset() *_LgNumPoolMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgNumPoolMgr) Get() (result LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgNumPoolMgr) Gets() (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgNumPoolMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 单号池ID
func (obj *_LgNumPoolMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 单号池名称
func (obj *_LgNumPoolMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPrefix prefix获取 单号池前缀
func (obj *_LgNumPoolMgr) WithPrefix(prefix string) Option {
	return optionFunc(func(o *options) { o.query["prefix"] = prefix })
}

// WithSuffix suffix获取 单号池后缀
func (obj *_LgNumPoolMgr) WithSuffix(suffix string) Option {
	return optionFunc(func(o *options) { o.query["suffix"] = suffix })
}

// WithStartNum start_num获取 开始号码
func (obj *_LgNumPoolMgr) WithStartNum(startNum int64) Option {
	return optionFunc(func(o *options) { o.query["start_num"] = startNum })
}

// WithEndNum end_num获取 结束号码
func (obj *_LgNumPoolMgr) WithEndNum(endNum int64) Option {
	return optionFunc(func(o *options) { o.query["end_num"] = endNum })
}

// WithVerificationMode verification_mode获取 验证号码合法方式:保存验证方式code
func (obj *_LgNumPoolMgr) WithVerificationMode(verificationMode string) Option {
	return optionFunc(func(o *options) { o.query["verification_mode"] = verificationMode })
}

// WithUsedNum used_num获取 已用到单号池位置:默认喂号码开始位置
func (obj *_LgNumPoolMgr) WithUsedNum(usedNum int64) Option {
	return optionFunc(func(o *options) { o.query["used_num"] = usedNum })
}

// WithLowestThreshold lowest_threshold获取 最低阈值
func (obj *_LgNumPoolMgr) WithLowestThreshold(lowestThreshold int64) Option {
	return optionFunc(func(o *options) { o.query["lowest_threshold"] = lowestThreshold })
}

// WithNoticeUserID notice_user_id获取 低于阈值通知人ID
func (obj *_LgNumPoolMgr) WithNoticeUserID(noticeUserID int64) Option {
	return optionFunc(func(o *options) { o.query["notice_user_id"] = noticeUserID })
}

// WithRemark remark获取 单号池备注
func (obj *_LgNumPoolMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithType type获取 单号池类型
// 1:默认类型,号段
// 2:倒入单号类型,所有单号需要导入
func (obj *_LgNumPoolMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgNumPoolMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgNumPoolMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgNumPoolMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_LgNumPoolMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgNumPoolMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgNumPoolMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgNumPoolMgr) GetByOption(opts ...Option) (result LgNumPool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgNumPoolMgr) GetByOptions(opts ...Option) (results []*LgNumPool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgNumPoolMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgNumPool, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where(options.query)
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

// GetFromID 通过id获取内容 单号池ID
func (obj *_LgNumPoolMgr) GetFromID(id int64) (result LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 单号池ID
func (obj *_LgNumPoolMgr) GetBatchFromID(ids []int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 单号池名称
func (obj *_LgNumPoolMgr) GetFromName(name string) (result LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 单号池名称
func (obj *_LgNumPoolMgr) GetBatchFromName(names []string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPrefix 通过prefix获取内容 单号池前缀
func (obj *_LgNumPoolMgr) GetFromPrefix(prefix string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`prefix` = ?", prefix).Find(&results).Error

	return
}

// GetBatchFromPrefix 批量查找 单号池前缀
func (obj *_LgNumPoolMgr) GetBatchFromPrefix(prefixs []string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`prefix` IN (?)", prefixs).Find(&results).Error

	return
}

// GetFromSuffix 通过suffix获取内容 单号池后缀
func (obj *_LgNumPoolMgr) GetFromSuffix(suffix string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`suffix` = ?", suffix).Find(&results).Error

	return
}

// GetBatchFromSuffix 批量查找 单号池后缀
func (obj *_LgNumPoolMgr) GetBatchFromSuffix(suffixs []string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`suffix` IN (?)", suffixs).Find(&results).Error

	return
}

// GetFromStartNum 通过start_num获取内容 开始号码
func (obj *_LgNumPoolMgr) GetFromStartNum(startNum int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`start_num` = ?", startNum).Find(&results).Error

	return
}

// GetBatchFromStartNum 批量查找 开始号码
func (obj *_LgNumPoolMgr) GetBatchFromStartNum(startNums []int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`start_num` IN (?)", startNums).Find(&results).Error

	return
}

// GetFromEndNum 通过end_num获取内容 结束号码
func (obj *_LgNumPoolMgr) GetFromEndNum(endNum int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`end_num` = ?", endNum).Find(&results).Error

	return
}

// GetBatchFromEndNum 批量查找 结束号码
func (obj *_LgNumPoolMgr) GetBatchFromEndNum(endNums []int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`end_num` IN (?)", endNums).Find(&results).Error

	return
}

// GetFromVerificationMode 通过verification_mode获取内容 验证号码合法方式:保存验证方式code
func (obj *_LgNumPoolMgr) GetFromVerificationMode(verificationMode string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`verification_mode` = ?", verificationMode).Find(&results).Error

	return
}

// GetBatchFromVerificationMode 批量查找 验证号码合法方式:保存验证方式code
func (obj *_LgNumPoolMgr) GetBatchFromVerificationMode(verificationModes []string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`verification_mode` IN (?)", verificationModes).Find(&results).Error

	return
}

// GetFromUsedNum 通过used_num获取内容 已用到单号池位置:默认喂号码开始位置
func (obj *_LgNumPoolMgr) GetFromUsedNum(usedNum int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`used_num` = ?", usedNum).Find(&results).Error

	return
}

// GetBatchFromUsedNum 批量查找 已用到单号池位置:默认喂号码开始位置
func (obj *_LgNumPoolMgr) GetBatchFromUsedNum(usedNums []int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`used_num` IN (?)", usedNums).Find(&results).Error

	return
}

// GetFromLowestThreshold 通过lowest_threshold获取内容 最低阈值
func (obj *_LgNumPoolMgr) GetFromLowestThreshold(lowestThreshold int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`lowest_threshold` = ?", lowestThreshold).Find(&results).Error

	return
}

// GetBatchFromLowestThreshold 批量查找 最低阈值
func (obj *_LgNumPoolMgr) GetBatchFromLowestThreshold(lowestThresholds []int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`lowest_threshold` IN (?)", lowestThresholds).Find(&results).Error

	return
}

// GetFromNoticeUserID 通过notice_user_id获取内容 低于阈值通知人ID
func (obj *_LgNumPoolMgr) GetFromNoticeUserID(noticeUserID int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`notice_user_id` = ?", noticeUserID).Find(&results).Error

	return
}

// GetBatchFromNoticeUserID 批量查找 低于阈值通知人ID
func (obj *_LgNumPoolMgr) GetBatchFromNoticeUserID(noticeUserIDs []int64) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`notice_user_id` IN (?)", noticeUserIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 单号池备注
func (obj *_LgNumPoolMgr) GetFromRemark(remark string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 单号池备注
func (obj *_LgNumPoolMgr) GetBatchFromRemark(remarks []string) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 单号池类型
// 1:默认类型,号段
// 2:倒入单号类型,所有单号需要导入
func (obj *_LgNumPoolMgr) GetFromType(_type int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 单号池类型
// 1:默认类型,号段
// 2:倒入单号类型,所有单号需要导入
func (obj *_LgNumPoolMgr) GetBatchFromType(_types []int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgNumPoolMgr) GetFromCreateTime(createTime time.Time) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgNumPoolMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgNumPoolMgr) GetFromCreateUser(createUser int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgNumPoolMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgNumPoolMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgNumPoolMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_LgNumPoolMgr) GetFromUpdateUser(updateUser int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_LgNumPoolMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgNumPoolMgr) GetFromVersion(version int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgNumPoolMgr) GetBatchFromVersion(versions []int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgNumPoolMgr) GetFromDeleted(deleted int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgNumPoolMgr) GetBatchFromDeleted(deleteds []int) (results []*LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgNumPoolMgr) FetchByPrimaryKey(id int64) (result LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgOrderNumPoolNameUIndex primary or index 获取唯一内容
func (obj *_LgNumPoolMgr) FetchUniqueByLgOrderNumPoolNameUIndex(name string) (result LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`name` = ?", name).First(&result).Error

	return
}

// FetchUniqueIndexByLgOrderNumPool1PrefixSuffixUIndex primary or index 获取唯一内容
func (obj *_LgNumPoolMgr) FetchUniqueIndexByLgOrderNumPool1PrefixSuffixUIndex(prefix string, suffix string) (result LgNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgNumPool{}).Where("`prefix` = ? AND `suffix` = ?", prefix, suffix).First(&result).Error

	return
}
