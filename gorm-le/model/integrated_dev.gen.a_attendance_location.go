package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AAttendanceLocationMgr struct {
	*_BaseMgr
}

// AAttendanceLocationMgr open func
func AAttendanceLocationMgr(db *gorm.DB) *_AAttendanceLocationMgr {
	if db == nil {
		panic(fmt.Errorf("AAttendanceLocationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AAttendanceLocationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("a_attendance_location"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AAttendanceLocationMgr) GetTableName() string {
	return "a_attendance_location"
}

// Reset 重置gorm会话
func (obj *_AAttendanceLocationMgr) Reset() *_AAttendanceLocationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AAttendanceLocationMgr) Get() (result AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AAttendanceLocationMgr) Gets() (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AAttendanceLocationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOrgID org_id获取
func (obj *_AAttendanceLocationMgr) WithOrgID(orgID int) Option {
	return optionFunc(func(o *options) { o.query["org_id"] = orgID })
}

// WithLongitude longitude获取 经度
func (obj *_AAttendanceLocationMgr) WithLongitude(longitude float64) Option {
	return optionFunc(func(o *options) { o.query["longitude"] = longitude })
}

// WithDimension dimension获取 纬度
func (obj *_AAttendanceLocationMgr) WithDimension(dimension float64) Option {
	return optionFunc(func(o *options) { o.query["dimension"] = dimension })
}

// WithRadius radius获取 打卡半径
func (obj *_AAttendanceLocationMgr) WithRadius(radius int) Option {
	return optionFunc(func(o *options) { o.query["radius"] = radius })
}

// WithLocation location获取 位置
func (obj *_AAttendanceLocationMgr) WithLocation(location string) Option {
	return optionFunc(func(o *options) { o.query["location"] = location })
}

// WithRemark remark获取 备注
func (obj *_AAttendanceLocationMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_AAttendanceLocationMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_AAttendanceLocationMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_AAttendanceLocationMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_AAttendanceLocationMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_AAttendanceLocationMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_AAttendanceLocationMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_AAttendanceLocationMgr) GetByOption(opts ...Option) (result AAttendanceLocation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AAttendanceLocationMgr) GetByOptions(opts ...Option) (results []*AAttendanceLocation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AAttendanceLocationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AAttendanceLocation, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where(options.query)
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

// GetFromOrgID 通过org_id获取内容
func (obj *_AAttendanceLocationMgr) GetFromOrgID(orgID int) (result AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`org_id` = ?", orgID).First(&result).Error

	return
}

// GetBatchFromOrgID 批量查找
func (obj *_AAttendanceLocationMgr) GetBatchFromOrgID(orgIDs []int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`org_id` IN (?)", orgIDs).Find(&results).Error

	return
}

// GetFromLongitude 通过longitude获取内容 经度
func (obj *_AAttendanceLocationMgr) GetFromLongitude(longitude float64) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`longitude` = ?", longitude).Find(&results).Error

	return
}

// GetBatchFromLongitude 批量查找 经度
func (obj *_AAttendanceLocationMgr) GetBatchFromLongitude(longitudes []float64) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`longitude` IN (?)", longitudes).Find(&results).Error

	return
}

// GetFromDimension 通过dimension获取内容 纬度
func (obj *_AAttendanceLocationMgr) GetFromDimension(dimension float64) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`dimension` = ?", dimension).Find(&results).Error

	return
}

// GetBatchFromDimension 批量查找 纬度
func (obj *_AAttendanceLocationMgr) GetBatchFromDimension(dimensions []float64) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`dimension` IN (?)", dimensions).Find(&results).Error

	return
}

// GetFromRadius 通过radius获取内容 打卡半径
func (obj *_AAttendanceLocationMgr) GetFromRadius(radius int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`radius` = ?", radius).Find(&results).Error

	return
}

// GetBatchFromRadius 批量查找 打卡半径
func (obj *_AAttendanceLocationMgr) GetBatchFromRadius(radiuss []int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`radius` IN (?)", radiuss).Find(&results).Error

	return
}

// GetFromLocation 通过location获取内容 位置
func (obj *_AAttendanceLocationMgr) GetFromLocation(location string) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`location` = ?", location).Find(&results).Error

	return
}

// GetBatchFromLocation 批量查找 位置
func (obj *_AAttendanceLocationMgr) GetBatchFromLocation(locations []string) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`location` IN (?)", locations).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_AAttendanceLocationMgr) GetFromRemark(remark string) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_AAttendanceLocationMgr) GetBatchFromRemark(remarks []string) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_AAttendanceLocationMgr) GetFromCreateTime(createTime time.Time) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_AAttendanceLocationMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_AAttendanceLocationMgr) GetFromCreateUser(createUser int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_AAttendanceLocationMgr) GetBatchFromCreateUser(createUsers []int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_AAttendanceLocationMgr) GetFromUpdateTime(updateTime time.Time) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_AAttendanceLocationMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_AAttendanceLocationMgr) GetFromUpdateUser(updateUser int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_AAttendanceLocationMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_AAttendanceLocationMgr) GetFromVersion(version int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_AAttendanceLocationMgr) GetBatchFromVersion(versions []int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_AAttendanceLocationMgr) GetFromDeleted(deleted int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_AAttendanceLocationMgr) GetBatchFromDeleted(deleteds []int) (results []*AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AAttendanceLocationMgr) FetchByPrimaryKey(orgID int) (result AAttendanceLocation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AAttendanceLocation{}).Where("`org_id` = ?", orgID).First(&result).Error

	return
}
