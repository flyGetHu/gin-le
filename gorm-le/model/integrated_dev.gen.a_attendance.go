package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AAttendanceMgr struct {
	*_BaseMgr
}

// AAttendanceMgr open func
func AAttendanceMgr(db *gorm.DB) *_AAttendanceMgr {
	if db == nil {
		panic(fmt.Errorf("AAttendanceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AAttendanceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("a_attendance"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AAttendanceMgr) GetTableName() string {
	return "a_attendance"
}

// Reset 重置gorm会话
func (obj *_AAttendanceMgr) Reset() *_AAttendanceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AAttendanceMgr) Get() (result AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AAttendanceMgr) Gets() (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AAttendanceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AAttendanceMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户ID
func (obj *_AAttendanceMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithOrgID org_id获取 部门ID
func (obj *_AAttendanceMgr) WithOrgID(orgID int) Option {
	return optionFunc(func(o *options) { o.query["org_id"] = orgID })
}

// WithWorkDate work_date获取 日期
func (obj *_AAttendanceMgr) WithWorkDate(workDate string) Option {
	return optionFunc(func(o *options) { o.query["work_date"] = workDate })
}

// WithShiftName shift_name获取 考勤班次名称
func (obj *_AAttendanceMgr) WithShiftName(shiftName string) Option {
	return optionFunc(func(o *options) { o.query["shift_name"] = shiftName })
}

// WithUpWorkDatetime up_work_datetime获取 上班时间
func (obj *_AAttendanceMgr) WithUpWorkDatetime(upWorkDatetime time.Time) Option {
	return optionFunc(func(o *options) { o.query["up_work_datetime"] = upWorkDatetime })
}

// WithDownWorkDatatime down_work_datatime获取 下班时间
func (obj *_AAttendanceMgr) WithDownWorkDatatime(downWorkDatatime time.Time) Option {
	return optionFunc(func(o *options) { o.query["down_work_datatime"] = downWorkDatatime })
}

// WithUpWorkLate up_work_late获取 迟到时间
func (obj *_AAttendanceMgr) WithUpWorkLate(upWorkLate int) Option {
	return optionFunc(func(o *options) { o.query["up_work_late"] = upWorkLate })
}

// WithDownWorkEarly down_work_early获取 早退时间
func (obj *_AAttendanceMgr) WithDownWorkEarly(downWorkEarly int) Option {
	return optionFunc(func(o *options) { o.query["down_work_early"] = downWorkEarly })
}

// WithRemark remark获取 备注
func (obj *_AAttendanceMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_AAttendanceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_AAttendanceMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_AAttendanceMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_AAttendanceMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_AAttendanceMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_AAttendanceMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_AAttendanceMgr) GetByOption(opts ...Option) (result AAttendance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AAttendanceMgr) GetByOptions(opts ...Option) (results []*AAttendance, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AAttendanceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AAttendance, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where(options.query)
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
func (obj *_AAttendanceMgr) GetFromID(id int) (result AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AAttendanceMgr) GetBatchFromID(ids []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_AAttendanceMgr) GetFromUserID(userID int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_AAttendanceMgr) GetBatchFromUserID(userIDs []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromOrgID 通过org_id获取内容 部门ID
func (obj *_AAttendanceMgr) GetFromOrgID(orgID int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`org_id` = ?", orgID).Find(&results).Error

	return
}

// GetBatchFromOrgID 批量查找 部门ID
func (obj *_AAttendanceMgr) GetBatchFromOrgID(orgIDs []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`org_id` IN (?)", orgIDs).Find(&results).Error

	return
}

// GetFromWorkDate 通过work_date获取内容 日期
func (obj *_AAttendanceMgr) GetFromWorkDate(workDate string) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`work_date` = ?", workDate).Find(&results).Error

	return
}

// GetBatchFromWorkDate 批量查找 日期
func (obj *_AAttendanceMgr) GetBatchFromWorkDate(workDates []string) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`work_date` IN (?)", workDates).Find(&results).Error

	return
}

// GetFromShiftName 通过shift_name获取内容 考勤班次名称
func (obj *_AAttendanceMgr) GetFromShiftName(shiftName string) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`shift_name` = ?", shiftName).Find(&results).Error

	return
}

// GetBatchFromShiftName 批量查找 考勤班次名称
func (obj *_AAttendanceMgr) GetBatchFromShiftName(shiftNames []string) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`shift_name` IN (?)", shiftNames).Find(&results).Error

	return
}

// GetFromUpWorkDatetime 通过up_work_datetime获取内容 上班时间
func (obj *_AAttendanceMgr) GetFromUpWorkDatetime(upWorkDatetime time.Time) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`up_work_datetime` = ?", upWorkDatetime).Find(&results).Error

	return
}

// GetBatchFromUpWorkDatetime 批量查找 上班时间
func (obj *_AAttendanceMgr) GetBatchFromUpWorkDatetime(upWorkDatetimes []time.Time) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`up_work_datetime` IN (?)", upWorkDatetimes).Find(&results).Error

	return
}

// GetFromDownWorkDatatime 通过down_work_datatime获取内容 下班时间
func (obj *_AAttendanceMgr) GetFromDownWorkDatatime(downWorkDatatime time.Time) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`down_work_datatime` = ?", downWorkDatatime).Find(&results).Error

	return
}

// GetBatchFromDownWorkDatatime 批量查找 下班时间
func (obj *_AAttendanceMgr) GetBatchFromDownWorkDatatime(downWorkDatatimes []time.Time) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`down_work_datatime` IN (?)", downWorkDatatimes).Find(&results).Error

	return
}

// GetFromUpWorkLate 通过up_work_late获取内容 迟到时间
func (obj *_AAttendanceMgr) GetFromUpWorkLate(upWorkLate int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`up_work_late` = ?", upWorkLate).Find(&results).Error

	return
}

// GetBatchFromUpWorkLate 批量查找 迟到时间
func (obj *_AAttendanceMgr) GetBatchFromUpWorkLate(upWorkLates []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`up_work_late` IN (?)", upWorkLates).Find(&results).Error

	return
}

// GetFromDownWorkEarly 通过down_work_early获取内容 早退时间
func (obj *_AAttendanceMgr) GetFromDownWorkEarly(downWorkEarly int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`down_work_early` = ?", downWorkEarly).Find(&results).Error

	return
}

// GetBatchFromDownWorkEarly 批量查找 早退时间
func (obj *_AAttendanceMgr) GetBatchFromDownWorkEarly(downWorkEarlys []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`down_work_early` IN (?)", downWorkEarlys).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_AAttendanceMgr) GetFromRemark(remark string) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_AAttendanceMgr) GetBatchFromRemark(remarks []string) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_AAttendanceMgr) GetFromCreateTime(createTime time.Time) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_AAttendanceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_AAttendanceMgr) GetFromCreateUser(createUser int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_AAttendanceMgr) GetBatchFromCreateUser(createUsers []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_AAttendanceMgr) GetFromUpdateTime(updateTime time.Time) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_AAttendanceMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_AAttendanceMgr) GetFromUpdateUser(updateUser int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_AAttendanceMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_AAttendanceMgr) GetFromVersion(version int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_AAttendanceMgr) GetBatchFromVersion(versions []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_AAttendanceMgr) GetFromDeleted(deleted int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_AAttendanceMgr) GetBatchFromDeleted(deleteds []int) (results []*AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AAttendanceMgr) FetchByPrimaryKey(id int) (result AAttendance, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendance{}).Where("`id` = ?", id).First(&result).Error

	return
}
