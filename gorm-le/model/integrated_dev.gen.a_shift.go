package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AShiftMgr struct {
	*_BaseMgr
}

// AShiftMgr open func
func AShiftMgr(db *gorm.DB) *_AShiftMgr {
	if db == nil {
		panic(fmt.Errorf("AShiftMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AShiftMgr{_BaseMgr: &_BaseMgr{DB: db.Table("a_shift"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AShiftMgr) GetTableName() string {
	return "a_shift"
}

// Reset 重置gorm会话
func (obj *_AShiftMgr) Reset() *_AShiftMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AShiftMgr) Get() (result AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AShiftMgr) Gets() (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AShiftMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AShift{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AShiftMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithShiftName shift_name获取 班次名称
func (obj *_AShiftMgr) WithShiftName(shiftName string) Option {
	return optionFunc(func(o *options) { o.query["shift_name"] = shiftName })
}

// WithUpWorkTime up_work_time获取 上班打卡
func (obj *_AShiftMgr) WithUpWorkTime(upWorkTime string) Option {
	return optionFunc(func(o *options) { o.query["up_work_time"] = upWorkTime })
}

// WithDownWorkTime down_work_time获取 下班时间
func (obj *_AShiftMgr) WithDownWorkTime(downWorkTime string) Option {
	return optionFunc(func(o *options) { o.query["down_work_time"] = downWorkTime })
}

// WithUpWorkEarliest up_work_earliest获取 上班最早打卡时间
func (obj *_AShiftMgr) WithUpWorkEarliest(upWorkEarliest int) Option {
	return optionFunc(func(o *options) { o.query["up_work_earliest"] = upWorkEarliest })
}

// WithDownWorkEarliest down_work_earliest获取 下班最早打卡时间
func (obj *_AShiftMgr) WithDownWorkEarliest(downWorkEarliest int) Option {
	return optionFunc(func(o *options) { o.query["down_work_earliest"] = downWorkEarliest })
}

// WithUpWorkLatest up_work_latest获取 上班最晚打卡时间
func (obj *_AShiftMgr) WithUpWorkLatest(upWorkLatest int) Option {
	return optionFunc(func(o *options) { o.query["up_work_latest"] = upWorkLatest })
}

// WithDownWorkLatest down_work_latest获取 下班最晚打卡时间
func (obj *_AShiftMgr) WithDownWorkLatest(downWorkLatest int) Option {
	return optionFunc(func(o *options) { o.query["down_work_latest"] = downWorkLatest })
}

// WithIsPhoto is_photo获取 是否拍照
func (obj *_AShiftMgr) WithIsPhoto(isPhoto int) Option {
	return optionFunc(func(o *options) { o.query["is_photo"] = isPhoto })
}

// WithIsLockRange is_lock_range获取 是否锁定打卡范围
func (obj *_AShiftMgr) WithIsLockRange(isLockRange int) Option {
	return optionFunc(func(o *options) { o.query["is_lock_range"] = isLockRange })
}

// WithIsRest is_rest获取 是否休息/考勤开关
func (obj *_AShiftMgr) WithIsRest(isRest int) Option {
	return optionFunc(func(o *options) { o.query["is_rest"] = isRest })
}

// WithRemark remark获取 备注
func (obj *_AShiftMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_AShiftMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_AShiftMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_AShiftMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_AShiftMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_AShiftMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_AShiftMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_AShiftMgr) GetByOption(opts ...Option) (result AShift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AShiftMgr) GetByOptions(opts ...Option) (results []*AShift, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AShiftMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AShift, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AShift{}).Where(options.query)
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
func (obj *_AShiftMgr) GetFromID(id int) (result AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AShiftMgr) GetBatchFromID(ids []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromShiftName 通过shift_name获取内容 班次名称
func (obj *_AShiftMgr) GetFromShiftName(shiftName string) (result AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`shift_name` = ?", shiftName).First(&result).Error

	return
}

// GetBatchFromShiftName 批量查找 班次名称
func (obj *_AShiftMgr) GetBatchFromShiftName(shiftNames []string) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`shift_name` IN (?)", shiftNames).Find(&results).Error

	return
}

// GetFromUpWorkTime 通过up_work_time获取内容 上班打卡
func (obj *_AShiftMgr) GetFromUpWorkTime(upWorkTime string) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`up_work_time` = ?", upWorkTime).Find(&results).Error

	return
}

// GetBatchFromUpWorkTime 批量查找 上班打卡
func (obj *_AShiftMgr) GetBatchFromUpWorkTime(upWorkTimes []string) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`up_work_time` IN (?)", upWorkTimes).Find(&results).Error

	return
}

// GetFromDownWorkTime 通过down_work_time获取内容 下班时间
func (obj *_AShiftMgr) GetFromDownWorkTime(downWorkTime string) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`down_work_time` = ?", downWorkTime).Find(&results).Error

	return
}

// GetBatchFromDownWorkTime 批量查找 下班时间
func (obj *_AShiftMgr) GetBatchFromDownWorkTime(downWorkTimes []string) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`down_work_time` IN (?)", downWorkTimes).Find(&results).Error

	return
}

// GetFromUpWorkEarliest 通过up_work_earliest获取内容 上班最早打卡时间
func (obj *_AShiftMgr) GetFromUpWorkEarliest(upWorkEarliest int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`up_work_earliest` = ?", upWorkEarliest).Find(&results).Error

	return
}

// GetBatchFromUpWorkEarliest 批量查找 上班最早打卡时间
func (obj *_AShiftMgr) GetBatchFromUpWorkEarliest(upWorkEarliests []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`up_work_earliest` IN (?)", upWorkEarliests).Find(&results).Error

	return
}

// GetFromDownWorkEarliest 通过down_work_earliest获取内容 下班最早打卡时间
func (obj *_AShiftMgr) GetFromDownWorkEarliest(downWorkEarliest int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`down_work_earliest` = ?", downWorkEarliest).Find(&results).Error

	return
}

// GetBatchFromDownWorkEarliest 批量查找 下班最早打卡时间
func (obj *_AShiftMgr) GetBatchFromDownWorkEarliest(downWorkEarliests []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`down_work_earliest` IN (?)", downWorkEarliests).Find(&results).Error

	return
}

// GetFromUpWorkLatest 通过up_work_latest获取内容 上班最晚打卡时间
func (obj *_AShiftMgr) GetFromUpWorkLatest(upWorkLatest int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`up_work_latest` = ?", upWorkLatest).Find(&results).Error

	return
}

// GetBatchFromUpWorkLatest 批量查找 上班最晚打卡时间
func (obj *_AShiftMgr) GetBatchFromUpWorkLatest(upWorkLatests []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`up_work_latest` IN (?)", upWorkLatests).Find(&results).Error

	return
}

// GetFromDownWorkLatest 通过down_work_latest获取内容 下班最晚打卡时间
func (obj *_AShiftMgr) GetFromDownWorkLatest(downWorkLatest int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`down_work_latest` = ?", downWorkLatest).Find(&results).Error

	return
}

// GetBatchFromDownWorkLatest 批量查找 下班最晚打卡时间
func (obj *_AShiftMgr) GetBatchFromDownWorkLatest(downWorkLatests []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`down_work_latest` IN (?)", downWorkLatests).Find(&results).Error

	return
}

// GetFromIsPhoto 通过is_photo获取内容 是否拍照
func (obj *_AShiftMgr) GetFromIsPhoto(isPhoto int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`is_photo` = ?", isPhoto).Find(&results).Error

	return
}

// GetBatchFromIsPhoto 批量查找 是否拍照
func (obj *_AShiftMgr) GetBatchFromIsPhoto(isPhotos []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`is_photo` IN (?)", isPhotos).Find(&results).Error

	return
}

// GetFromIsLockRange 通过is_lock_range获取内容 是否锁定打卡范围
func (obj *_AShiftMgr) GetFromIsLockRange(isLockRange int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`is_lock_range` = ?", isLockRange).Find(&results).Error

	return
}

// GetBatchFromIsLockRange 批量查找 是否锁定打卡范围
func (obj *_AShiftMgr) GetBatchFromIsLockRange(isLockRanges []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`is_lock_range` IN (?)", isLockRanges).Find(&results).Error

	return
}

// GetFromIsRest 通过is_rest获取内容 是否休息/考勤开关
func (obj *_AShiftMgr) GetFromIsRest(isRest int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`is_rest` = ?", isRest).Find(&results).Error

	return
}

// GetBatchFromIsRest 批量查找 是否休息/考勤开关
func (obj *_AShiftMgr) GetBatchFromIsRest(isRests []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`is_rest` IN (?)", isRests).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_AShiftMgr) GetFromRemark(remark string) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_AShiftMgr) GetBatchFromRemark(remarks []string) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_AShiftMgr) GetFromCreateTime(createTime time.Time) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_AShiftMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_AShiftMgr) GetFromCreateUser(createUser int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_AShiftMgr) GetBatchFromCreateUser(createUsers []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_AShiftMgr) GetFromUpdateTime(updateTime time.Time) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_AShiftMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_AShiftMgr) GetFromUpdateUser(updateUser int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_AShiftMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_AShiftMgr) GetFromVersion(version int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_AShiftMgr) GetBatchFromVersion(versions []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_AShiftMgr) GetFromDeleted(deleted int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_AShiftMgr) GetBatchFromDeleted(deleteds []int) (results []*AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AShiftMgr) FetchByPrimaryKey(id int) (result AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUniqueShiftName primary or index 获取唯一内容
func (obj *_AShiftMgr) FetchUniqueByUniqueShiftName(shiftName string) (result AShift, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AShift{}).Where("`shift_name` = ?", shiftName).First(&result).Error

	return
}
