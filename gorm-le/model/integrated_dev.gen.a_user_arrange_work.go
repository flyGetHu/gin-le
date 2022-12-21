package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AUserArrangeWorkMgr struct {
	*_BaseMgr
}

// AUserArrangeWorkMgr open func
func AUserArrangeWorkMgr(db *gorm.DB) *_AUserArrangeWorkMgr {
	if db == nil {
		panic(fmt.Errorf("AUserArrangeWorkMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AUserArrangeWorkMgr{_BaseMgr: &_BaseMgr{DB: db.Table("a_user_arrange_work"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AUserArrangeWorkMgr) GetTableName() string {
	return "a_user_arrange_work"
}

// Reset 重置gorm会话
func (obj *_AUserArrangeWorkMgr) Reset() *_AUserArrangeWorkMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AUserArrangeWorkMgr) Get() (result AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AUserArrangeWorkMgr) Gets() (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AUserArrangeWorkMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AUserArrangeWorkMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserID user_id获取 用户ID
func (obj *_AUserArrangeWorkMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithPlanMonth plan_month获取 排班月份
func (obj *_AUserArrangeWorkMgr) WithPlanMonth(planMonth string) Option {
	return optionFunc(func(o *options) { o.query["plan_month"] = planMonth })
}

// WithShiftNames shift_names获取 班次名称字符串集
func (obj *_AUserArrangeWorkMgr) WithShiftNames(shiftNames string) Option {
	return optionFunc(func(o *options) { o.query["shift_names"] = shiftNames })
}

// WithShiftNameIDs shift_name_ids获取 班次名称ID字符串集
func (obj *_AUserArrangeWorkMgr) WithShiftNameIDs(shiftNameIDs string) Option {
	return optionFunc(func(o *options) { o.query["shift_name_ids"] = shiftNameIDs })
}

// WithRemark remark获取 备注
func (obj *_AUserArrangeWorkMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_AUserArrangeWorkMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_AUserArrangeWorkMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_AUserArrangeWorkMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_AUserArrangeWorkMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_AUserArrangeWorkMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_AUserArrangeWorkMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_AUserArrangeWorkMgr) GetByOption(opts ...Option) (result AUserArrangeWork, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AUserArrangeWorkMgr) GetByOptions(opts ...Option) (results []*AUserArrangeWork, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AUserArrangeWorkMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AUserArrangeWork, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where(options.query)
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
func (obj *_AUserArrangeWorkMgr) GetFromID(id int) (result AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AUserArrangeWorkMgr) GetBatchFromID(ids []int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容 用户ID
func (obj *_AUserArrangeWorkMgr) GetFromUserID(userID int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找 用户ID
func (obj *_AUserArrangeWorkMgr) GetBatchFromUserID(userIDs []int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromPlanMonth 通过plan_month获取内容 排班月份
func (obj *_AUserArrangeWorkMgr) GetFromPlanMonth(planMonth string) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`plan_month` = ?", planMonth).Find(&results).Error

	return
}

// GetBatchFromPlanMonth 批量查找 排班月份
func (obj *_AUserArrangeWorkMgr) GetBatchFromPlanMonth(planMonths []string) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`plan_month` IN (?)", planMonths).Find(&results).Error

	return
}

// GetFromShiftNames 通过shift_names获取内容 班次名称字符串集
func (obj *_AUserArrangeWorkMgr) GetFromShiftNames(shiftNames string) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`shift_names` = ?", shiftNames).Find(&results).Error

	return
}

// GetBatchFromShiftNames 批量查找 班次名称字符串集
func (obj *_AUserArrangeWorkMgr) GetBatchFromShiftNames(shiftNamess []string) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`shift_names` IN (?)", shiftNamess).Find(&results).Error

	return
}

// GetFromShiftNameIDs 通过shift_name_ids获取内容 班次名称ID字符串集
func (obj *_AUserArrangeWorkMgr) GetFromShiftNameIDs(shiftNameIDs string) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`shift_name_ids` = ?", shiftNameIDs).Find(&results).Error

	return
}

// GetBatchFromShiftNameIDs 批量查找 班次名称ID字符串集
func (obj *_AUserArrangeWorkMgr) GetBatchFromShiftNameIDs(shiftNameIDss []string) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`shift_name_ids` IN (?)", shiftNameIDss).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_AUserArrangeWorkMgr) GetFromRemark(remark string) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_AUserArrangeWorkMgr) GetBatchFromRemark(remarks []string) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_AUserArrangeWorkMgr) GetFromCreateTime(createTime time.Time) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_AUserArrangeWorkMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_AUserArrangeWorkMgr) GetFromCreateUser(createUser int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_AUserArrangeWorkMgr) GetBatchFromCreateUser(createUsers []int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_AUserArrangeWorkMgr) GetFromUpdateTime(updateTime time.Time) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_AUserArrangeWorkMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_AUserArrangeWorkMgr) GetFromUpdateUser(updateUser int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_AUserArrangeWorkMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_AUserArrangeWorkMgr) GetFromVersion(version int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_AUserArrangeWorkMgr) GetBatchFromVersion(versions []int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_AUserArrangeWorkMgr) GetFromDeleted(deleted int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_AUserArrangeWorkMgr) GetBatchFromDeleted(deleteds []int) (results []*AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AUserArrangeWorkMgr) FetchByPrimaryKey(id int) (result AUserArrangeWork, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AUserArrangeWork{}).Where("`id` = ?", id).First(&result).Error

	return
}
