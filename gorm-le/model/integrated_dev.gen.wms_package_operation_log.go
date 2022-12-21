package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsPackageOperationLogMgr struct {
	*_BaseMgr
}

// WmsPackageOperationLogMgr open func
func WmsPackageOperationLogMgr(db *gorm.DB) *_WmsPackageOperationLogMgr {
	if db == nil {
		panic(fmt.Errorf("WmsPackageOperationLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsPackageOperationLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_package_operation_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsPackageOperationLogMgr) GetTableName() string {
	return "wms_package_operation_log"
}

// Reset 重置gorm会话
func (obj *_WmsPackageOperationLogMgr) Reset() *_WmsPackageOperationLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsPackageOperationLogMgr) Get() (result WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsPackageOperationLogMgr) Gets() (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsPackageOperationLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsPackageOperationLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPackageID package_id获取 包裹id
func (obj *_WmsPackageOperationLogMgr) WithPackageID(packageID int) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// WithOperationContent operation_content获取 操作内容
func (obj *_WmsPackageOperationLogMgr) WithOperationContent(operationContent string) Option {
	return optionFunc(func(o *options) { o.query["operation_content"] = operationContent })
}

// WithRemark remark获取 备注
func (obj *_WmsPackageOperationLogMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsPackageOperationLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsPackageOperationLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsPackageOperationLogMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsPackageOperationLogMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsPackageOperationLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsPackageOperationLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithIP ip获取 ip地址
func (obj *_WmsPackageOperationLogMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// GetByOption 功能选项模式获取
func (obj *_WmsPackageOperationLogMgr) GetByOption(opts ...Option) (result WmsPackageOperationLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsPackageOperationLogMgr) GetByOptions(opts ...Option) (results []*WmsPackageOperationLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsPackageOperationLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsPackageOperationLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where(options.query)
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
func (obj *_WmsPackageOperationLogMgr) GetFromID(id int) (result WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsPackageOperationLogMgr) GetBatchFromID(ids []int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容 包裹id
func (obj *_WmsPackageOperationLogMgr) GetFromPackageID(packageID int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找 包裹id
func (obj *_WmsPackageOperationLogMgr) GetBatchFromPackageID(packageIDs []int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

// GetFromOperationContent 通过operation_content获取内容 操作内容
func (obj *_WmsPackageOperationLogMgr) GetFromOperationContent(operationContent string) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`operation_content` = ?", operationContent).Find(&results).Error

	return
}

// GetBatchFromOperationContent 批量查找 操作内容
func (obj *_WmsPackageOperationLogMgr) GetBatchFromOperationContent(operationContents []string) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`operation_content` IN (?)", operationContents).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WmsPackageOperationLogMgr) GetFromRemark(remark string) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WmsPackageOperationLogMgr) GetBatchFromRemark(remarks []string) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsPackageOperationLogMgr) GetFromCreateTime(createTime time.Time) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsPackageOperationLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsPackageOperationLogMgr) GetFromCreateUser(createUser int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsPackageOperationLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsPackageOperationLogMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsPackageOperationLogMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsPackageOperationLogMgr) GetFromUpdateUser(updateUser int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsPackageOperationLogMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsPackageOperationLogMgr) GetFromVersion(version int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsPackageOperationLogMgr) GetBatchFromVersion(versions []int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsPackageOperationLogMgr) GetFromDeleted(deleted int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsPackageOperationLogMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 ip地址
func (obj *_WmsPackageOperationLogMgr) GetFromIP(ip string) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找 ip地址
func (obj *_WmsPackageOperationLogMgr) GetBatchFromIP(ips []string) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsPackageOperationLogMgr) FetchByPrimaryKey(id int) (result WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByWmsPackageOperationLogPackageIDIndex  获取多个内容
func (obj *_WmsPackageOperationLogMgr) FetchIndexByWmsPackageOperationLogPackageIDIndex(packageID int) (results []*WmsPackageOperationLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLog{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}
