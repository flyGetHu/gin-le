package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsSorterGridMgr struct {
	*_BaseMgr
}

// WmsSorterGridMgr open func
func WmsSorterGridMgr(db *gorm.DB) *_WmsSorterGridMgr {
	if db == nil {
		panic(fmt.Errorf("WmsSorterGridMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsSorterGridMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_sorter_grid"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsSorterGridMgr) GetTableName() string {
	return "wms_sorter_grid"
}

// Reset 重置gorm会话
func (obj *_WmsSorterGridMgr) Reset() *_WmsSorterGridMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsSorterGridMgr) Get() (result WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsSorterGridMgr) Gets() (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsSorterGridMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsSorterGridMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGridNum grid_num获取 格口
func (obj *_WmsSorterGridMgr) WithGridNum(gridNum int) Option {
	return optionFunc(func(o *options) { o.query["grid_num"] = gridNum })
}

// WithRemark remark获取 备注
func (obj *_WmsSorterGridMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsSorterGridMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsSorterGridMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsSorterGridMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsSorterGridMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsSorterGridMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsSorterGridMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsSorterGridMgr) GetByOption(opts ...Option) (result WmsSorterGrid, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsSorterGridMgr) GetByOptions(opts ...Option) (results []*WmsSorterGrid, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsSorterGridMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsSorterGrid, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where(options.query)
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
func (obj *_WmsSorterGridMgr) GetFromID(id int) (result WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsSorterGridMgr) GetBatchFromID(ids []int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGridNum 通过grid_num获取内容 格口
func (obj *_WmsSorterGridMgr) GetFromGridNum(gridNum int) (result WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`grid_num` = ?", gridNum).First(&result).Error

	return
}

// GetBatchFromGridNum 批量查找 格口
func (obj *_WmsSorterGridMgr) GetBatchFromGridNum(gridNums []int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`grid_num` IN (?)", gridNums).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WmsSorterGridMgr) GetFromRemark(remark string) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WmsSorterGridMgr) GetBatchFromRemark(remarks []string) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsSorterGridMgr) GetFromCreateTime(createTime time.Time) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsSorterGridMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsSorterGridMgr) GetFromCreateUser(createUser int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsSorterGridMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsSorterGridMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsSorterGridMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsSorterGridMgr) GetFromUpdateUser(updateUser int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsSorterGridMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsSorterGridMgr) GetFromVersion(version int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsSorterGridMgr) GetBatchFromVersion(versions []int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsSorterGridMgr) GetFromDeleted(deleted int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsSorterGridMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsSorterGridMgr) FetchByPrimaryKey(id int) (result WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByGridNum primary or index 获取唯一内容
func (obj *_WmsSorterGridMgr) FetchUniqueByGridNum(gridNum int) (result WmsSorterGrid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGrid{}).Where("`grid_num` = ?", gridNum).First(&result).Error

	return
}
