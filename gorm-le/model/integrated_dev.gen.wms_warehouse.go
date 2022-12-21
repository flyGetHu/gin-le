package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsWarehouseMgr struct {
	*_BaseMgr
}

// WmsWarehouseMgr open func
func WmsWarehouseMgr(db *gorm.DB) *_WmsWarehouseMgr {
	if db == nil {
		panic(fmt.Errorf("WmsWarehouseMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsWarehouseMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_warehouse"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsWarehouseMgr) GetTableName() string {
	return "wms_warehouse"
}

// Reset 重置gorm会话
func (obj *_WmsWarehouseMgr) Reset() *_WmsWarehouseMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsWarehouseMgr) Get() (result WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsWarehouseMgr) Gets() (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsWarehouseMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_WmsWarehouseMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWarehouseCode warehouse_code获取 仓库code
func (obj *_WmsWarehouseMgr) WithWarehouseCode(warehouseCode string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_code"] = warehouseCode })
}

// WithWarehouseName warehouse_name获取 仓库名
func (obj *_WmsWarehouseMgr) WithWarehouseName(warehouseName string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_name"] = warehouseName })
}

// WithWarehouseType warehouse_type获取 仓库类型(0客户仓库,1安骏仓库)
func (obj *_WmsWarehouseMgr) WithWarehouseType(warehouseType int) Option {
	return optionFunc(func(o *options) { o.query["warehouse_type"] = warehouseType })
}

// WithCustomerID customer_id获取 客户id(客户仓库必填)
func (obj *_WmsWarehouseMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsWarehouseMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsWarehouseMgr) WithCreateUser(createUser uint) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsWarehouseMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_WmsWarehouseMgr) WithUpdateUser(updateUser uint) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsWarehouseMgr) WithVersion(version uint) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsWarehouseMgr) WithDeleted(deleted uint32) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithRemark remark获取 备注
func (obj *_WmsWarehouseMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithStatus status获取 使用状态是否禁用0代表否1代表是
func (obj *_WmsWarehouseMgr) WithStatus(status uint32) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_WmsWarehouseMgr) GetByOption(opts ...Option) (result WmsWarehouse, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsWarehouseMgr) GetByOptions(opts ...Option) (results []*WmsWarehouse, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsWarehouseMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsWarehouse, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *_WmsWarehouseMgr) GetFromID(id uint) (result WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_WmsWarehouseMgr) GetBatchFromID(ids []uint) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWarehouseCode 通过warehouse_code获取内容 仓库code
func (obj *_WmsWarehouseMgr) GetFromWarehouseCode(warehouseCode string) (result WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`warehouse_code` = ?", warehouseCode).First(&result).Error

	return
}

// GetBatchFromWarehouseCode 批量查找 仓库code
func (obj *_WmsWarehouseMgr) GetBatchFromWarehouseCode(warehouseCodes []string) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`warehouse_code` IN (?)", warehouseCodes).Find(&results).Error

	return
}

// GetFromWarehouseName 通过warehouse_name获取内容 仓库名
func (obj *_WmsWarehouseMgr) GetFromWarehouseName(warehouseName string) (result WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`warehouse_name` = ?", warehouseName).First(&result).Error

	return
}

// GetBatchFromWarehouseName 批量查找 仓库名
func (obj *_WmsWarehouseMgr) GetBatchFromWarehouseName(warehouseNames []string) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`warehouse_name` IN (?)", warehouseNames).Find(&results).Error

	return
}

// GetFromWarehouseType 通过warehouse_type获取内容 仓库类型(0客户仓库,1安骏仓库)
func (obj *_WmsWarehouseMgr) GetFromWarehouseType(warehouseType int) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`warehouse_type` = ?", warehouseType).Find(&results).Error

	return
}

// GetBatchFromWarehouseType 批量查找 仓库类型(0客户仓库,1安骏仓库)
func (obj *_WmsWarehouseMgr) GetBatchFromWarehouseType(warehouseTypes []int) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`warehouse_type` IN (?)", warehouseTypes).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id(客户仓库必填)
func (obj *_WmsWarehouseMgr) GetFromCustomerID(customerID int) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id(客户仓库必填)
func (obj *_WmsWarehouseMgr) GetBatchFromCustomerID(customerIDs []int) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsWarehouseMgr) GetFromCreateTime(createTime time.Time) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsWarehouseMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsWarehouseMgr) GetFromCreateUser(createUser uint) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsWarehouseMgr) GetBatchFromCreateUser(createUsers []uint) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsWarehouseMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsWarehouseMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_WmsWarehouseMgr) GetFromUpdateUser(updateUser uint) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_WmsWarehouseMgr) GetBatchFromUpdateUser(updateUsers []uint) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsWarehouseMgr) GetFromVersion(version uint) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsWarehouseMgr) GetBatchFromVersion(versions []uint) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsWarehouseMgr) GetFromDeleted(deleted uint32) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsWarehouseMgr) GetBatchFromDeleted(deleteds []uint32) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WmsWarehouseMgr) GetFromRemark(remark string) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WmsWarehouseMgr) GetBatchFromRemark(remarks []string) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 使用状态是否禁用0代表否1代表是
func (obj *_WmsWarehouseMgr) GetFromStatus(status uint32) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 使用状态是否禁用0代表否1代表是
func (obj *_WmsWarehouseMgr) GetBatchFromStatus(statuss []uint32) (results []*WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsWarehouseMgr) FetchByPrimaryKey(id uint) (result WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUNIQUECODE primary or index 获取唯一内容
func (obj *_WmsWarehouseMgr) FetchUniqueByUNIQUECODE(warehouseCode string) (result WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`warehouse_code` = ?", warehouseCode).First(&result).Error

	return
}

// FetchUniqueByUNIQUENAME primary or index 获取唯一内容
func (obj *_WmsWarehouseMgr) FetchUniqueByUNIQUENAME(warehouseName string) (result WmsWarehouse, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsWarehouse{}).Where("`warehouse_name` = ?", warehouseName).First(&result).Error

	return
}
