package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderNumPoolMgr struct {
	*_BaseMgr
}

// LgOrderNumPoolMgr open func
func LgOrderNumPoolMgr(db *gorm.DB) *_LgOrderNumPoolMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderNumPoolMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderNumPoolMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_num_pool"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderNumPoolMgr) GetTableName() string {
	return "lg_order_num_pool"
}

// Reset 重置gorm会话
func (obj *_LgOrderNumPoolMgr) Reset() *_LgOrderNumPoolMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderNumPoolMgr) Get() (result LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderNumPoolMgr) Gets() (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderNumPoolMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 单号池ID
func (obj *_LgOrderNumPoolMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 单号池名称
func (obj *_LgOrderNumPoolMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithLowestThreshold lowest_threshold获取 最低阈值
func (obj *_LgOrderNumPoolMgr) WithLowestThreshold(lowestThreshold int64) Option {
	return optionFunc(func(o *options) { o.query["lowest_threshold"] = lowestThreshold })
}

// WithNoticeUserID notice_user_id获取 低于阈值通知人ID
func (obj *_LgOrderNumPoolMgr) WithNoticeUserID(noticeUserID int64) Option {
	return optionFunc(func(o *options) { o.query["notice_user_id"] = noticeUserID })
}

// WithRemark remark获取 单号池备注
func (obj *_LgOrderNumPoolMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderNumPoolMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderNumPoolMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderNumPoolMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderNumPoolMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderNumPoolMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderNumPoolMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithType type获取 类型(系统,外部)
func (obj *_LgOrderNumPoolMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithPrefix prefix获取 号段前缀
func (obj *_LgOrderNumPoolMgr) WithPrefix(prefix string) Option {
	return optionFunc(func(o *options) { o.query["prefix"] = prefix })
}

// WithSuffix suffix获取 号段后缀
func (obj *_LgOrderNumPoolMgr) WithSuffix(suffix string) Option {
	return optionFunc(func(o *options) { o.query["suffix"] = suffix })
}

// WithInitNum init_num获取 初始化值
func (obj *_LgOrderNumPoolMgr) WithInitNum(initNum int64) Option {
	return optionFunc(func(o *options) { o.query["init_num"] = initNum })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderNumPoolMgr) GetByOption(opts ...Option) (result LgOrderNumPool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderNumPoolMgr) GetByOptions(opts ...Option) (results []*LgOrderNumPool, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderNumPoolMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderNumPool, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where(options.query)
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
func (obj *_LgOrderNumPoolMgr) GetFromID(id int64) (result LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 单号池ID
func (obj *_LgOrderNumPoolMgr) GetBatchFromID(ids []int64) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 单号池名称
func (obj *_LgOrderNumPoolMgr) GetFromName(name string) (result LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 单号池名称
func (obj *_LgOrderNumPoolMgr) GetBatchFromName(names []string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromLowestThreshold 通过lowest_threshold获取内容 最低阈值
func (obj *_LgOrderNumPoolMgr) GetFromLowestThreshold(lowestThreshold int64) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`lowest_threshold` = ?", lowestThreshold).Find(&results).Error

	return
}

// GetBatchFromLowestThreshold 批量查找 最低阈值
func (obj *_LgOrderNumPoolMgr) GetBatchFromLowestThreshold(lowestThresholds []int64) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`lowest_threshold` IN (?)", lowestThresholds).Find(&results).Error

	return
}

// GetFromNoticeUserID 通过notice_user_id获取内容 低于阈值通知人ID
func (obj *_LgOrderNumPoolMgr) GetFromNoticeUserID(noticeUserID int64) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`notice_user_id` = ?", noticeUserID).Find(&results).Error

	return
}

// GetBatchFromNoticeUserID 批量查找 低于阈值通知人ID
func (obj *_LgOrderNumPoolMgr) GetBatchFromNoticeUserID(noticeUserIDs []int64) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`notice_user_id` IN (?)", noticeUserIDs).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 单号池备注
func (obj *_LgOrderNumPoolMgr) GetFromRemark(remark string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 单号池备注
func (obj *_LgOrderNumPoolMgr) GetBatchFromRemark(remarks []string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderNumPoolMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderNumPoolMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderNumPoolMgr) GetFromCreateUser(createUser int) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderNumPoolMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderNumPoolMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderNumPoolMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderNumPoolMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderNumPoolMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderNumPoolMgr) GetFromVersion(version int) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderNumPoolMgr) GetBatchFromVersion(versions []int) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderNumPoolMgr) GetFromDeleted(deleted int) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderNumPoolMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型(系统,外部)
func (obj *_LgOrderNumPoolMgr) GetFromType(_type string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型(系统,外部)
func (obj *_LgOrderNumPoolMgr) GetBatchFromType(_types []string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromPrefix 通过prefix获取内容 号段前缀
func (obj *_LgOrderNumPoolMgr) GetFromPrefix(prefix string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`prefix` = ?", prefix).Find(&results).Error

	return
}

// GetBatchFromPrefix 批量查找 号段前缀
func (obj *_LgOrderNumPoolMgr) GetBatchFromPrefix(prefixs []string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`prefix` IN (?)", prefixs).Find(&results).Error

	return
}

// GetFromSuffix 通过suffix获取内容 号段后缀
func (obj *_LgOrderNumPoolMgr) GetFromSuffix(suffix string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`suffix` = ?", suffix).Find(&results).Error

	return
}

// GetBatchFromSuffix 批量查找 号段后缀
func (obj *_LgOrderNumPoolMgr) GetBatchFromSuffix(suffixs []string) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`suffix` IN (?)", suffixs).Find(&results).Error

	return
}

// GetFromInitNum 通过init_num获取内容 初始化值
func (obj *_LgOrderNumPoolMgr) GetFromInitNum(initNum int64) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`init_num` = ?", initNum).Find(&results).Error

	return
}

// GetBatchFromInitNum 批量查找 初始化值
func (obj *_LgOrderNumPoolMgr) GetBatchFromInitNum(initNums []int64) (results []*LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`init_num` IN (?)", initNums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderNumPoolMgr) FetchByPrimaryKey(id int64) (result LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgOrderNumPoolNameUIndex primary or index 获取唯一内容
func (obj *_LgOrderNumPoolMgr) FetchUniqueByLgOrderNumPoolNameUIndex(name string) (result LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`name` = ?", name).First(&result).Error

	return
}

// FetchUniqueIndexByLgOrderNumPoolPrefixSuffixUIndex primary or index 获取唯一内容
func (obj *_LgOrderNumPoolMgr) FetchUniqueIndexByLgOrderNumPoolPrefixSuffixUIndex(prefix string, suffix string) (result LgOrderNumPool, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumPool{}).Where("`prefix` = ? AND `suffix` = ?", prefix, suffix).First(&result).Error

	return
}
