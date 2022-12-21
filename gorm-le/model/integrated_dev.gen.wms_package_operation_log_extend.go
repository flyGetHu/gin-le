package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsPackageOperationLogExtendMgr struct {
	*_BaseMgr
}

// WmsPackageOperationLogExtendMgr open func
func WmsPackageOperationLogExtendMgr(db *gorm.DB) *_WmsPackageOperationLogExtendMgr {
	if db == nil {
		panic(fmt.Errorf("WmsPackageOperationLogExtendMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsPackageOperationLogExtendMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_package_operation_log_extend"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsPackageOperationLogExtendMgr) GetTableName() string {
	return "wms_package_operation_log_extend"
}

// Reset 重置gorm会话
func (obj *_WmsPackageOperationLogExtendMgr) Reset() *_WmsPackageOperationLogExtendMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsPackageOperationLogExtendMgr) Get() (result WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsPackageOperationLogExtendMgr) Gets() (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsPackageOperationLogExtendMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsPackageOperationLogExtendMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 订单id
func (obj *_WmsPackageOperationLogExtendMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithOperationContent operation_content获取 操作内容
func (obj *_WmsPackageOperationLogExtendMgr) WithOperationContent(operationContent string) Option {
	return optionFunc(func(o *options) { o.query["operation_content"] = operationContent })
}

// WithRemark remark获取 备注
func (obj *_WmsPackageOperationLogExtendMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsPackageOperationLogExtendMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsPackageOperationLogExtendMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsPackageOperationLogExtendMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WmsPackageOperationLogExtendMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsPackageOperationLogExtendMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsPackageOperationLogExtendMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithIP ip获取 ip地址
func (obj *_WmsPackageOperationLogExtendMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// GetByOption 功能选项模式获取
func (obj *_WmsPackageOperationLogExtendMgr) GetByOption(opts ...Option) (result WmsPackageOperationLogExtend, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsPackageOperationLogExtendMgr) GetByOptions(opts ...Option) (results []*WmsPackageOperationLogExtend, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsPackageOperationLogExtendMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsPackageOperationLogExtend, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where(options.query)
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
func (obj *_WmsPackageOperationLogExtendMgr) GetFromID(id int) (result WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromID(ids []int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单id
func (obj *_WmsPackageOperationLogExtendMgr) GetFromOrderID(orderID int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 订单id
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromOrderID(orderIDs []int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromOperationContent 通过operation_content获取内容 操作内容
func (obj *_WmsPackageOperationLogExtendMgr) GetFromOperationContent(operationContent string) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`operation_content` = ?", operationContent).Find(&results).Error

	return
}

// GetBatchFromOperationContent 批量查找 操作内容
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromOperationContent(operationContents []string) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`operation_content` IN (?)", operationContents).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WmsPackageOperationLogExtendMgr) GetFromRemark(remark string) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromRemark(remarks []string) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsPackageOperationLogExtendMgr) GetFromCreateTime(createTime time.Time) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsPackageOperationLogExtendMgr) GetFromCreateUser(createUser int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsPackageOperationLogExtendMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WmsPackageOperationLogExtendMgr) GetFromUpdateUser(updateUser int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsPackageOperationLogExtendMgr) GetFromVersion(version int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromVersion(versions []int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsPackageOperationLogExtendMgr) GetFromDeleted(deleted int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 ip地址
func (obj *_WmsPackageOperationLogExtendMgr) GetFromIP(ip string) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找 ip地址
func (obj *_WmsPackageOperationLogExtendMgr) GetBatchFromIP(ips []string) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsPackageOperationLogExtendMgr) FetchByPrimaryKey(id int) (result WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByWmsPackageOperationLogOrderIDIndex  获取多个内容
func (obj *_WmsPackageOperationLogExtendMgr) FetchIndexByWmsPackageOperationLogOrderIDIndex(orderID int) (results []*WmsPackageOperationLogExtend, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsPackageOperationLogExtend{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}
