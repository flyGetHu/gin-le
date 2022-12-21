package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderWmsOperationExpandMgr struct {
	*_BaseMgr
}

// LgOrderWmsOperationExpandMgr open func
func LgOrderWmsOperationExpandMgr(db *gorm.DB) *_LgOrderWmsOperationExpandMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderWmsOperationExpandMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderWmsOperationExpandMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_wms_operation_expand"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderWmsOperationExpandMgr) GetTableName() string {
	return "lg_order_wms_operation_expand"
}

// Reset 重置gorm会话
func (obj *_LgOrderWmsOperationExpandMgr) Reset() *_LgOrderWmsOperationExpandMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderWmsOperationExpandMgr) Get() (result LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderWmsOperationExpandMgr) Gets() (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderWmsOperationExpandMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgOrderWmsOperationExpandMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取  订单id
func (obj *_LgOrderWmsOperationExpandMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithWarehouseID warehouse_id获取  仓库id
func (obj *_LgOrderWmsOperationExpandMgr) WithWarehouseID(warehouseID int) Option {
	return optionFunc(func(o *options) { o.query["warehouse_id"] = warehouseID })
}

// WithWarehouseCode warehouse_code获取  仓库code 默认code是没有
func (obj *_LgOrderWmsOperationExpandMgr) WithWarehouseCode(warehouseCode string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_code"] = warehouseCode })
}

// WithWarehouseName warehouse_name获取  仓库名称 （默认仓库是没有仓库）
func (obj *_LgOrderWmsOperationExpandMgr) WithWarehouseName(warehouseName string) Option {
	return optionFunc(func(o *options) { o.query["warehouse_name"] = warehouseName })
}

// WithReceiptTime receipt_time获取  入库时间
func (obj *_LgOrderWmsOperationExpandMgr) WithReceiptTime(receiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_time"] = receiptTime })
}

// WithSendTime send_time获取  出库时间
func (obj *_LgOrderWmsOperationExpandMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithCreateTime create_time获取  创建时间
func (obj *_LgOrderWmsOperationExpandMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取  创建用户
func (obj *_LgOrderWmsOperationExpandMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取  更新时间
func (obj *_LgOrderWmsOperationExpandMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取  更新用户
func (obj *_LgOrderWmsOperationExpandMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取  乐观锁
func (obj *_LgOrderWmsOperationExpandMgr) WithVersion(version uint) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取  逻辑删除
func (obj *_LgOrderWmsOperationExpandMgr) WithDeleted(deleted uint) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderWmsOperationExpandMgr) GetByOption(opts ...Option) (result LgOrderWmsOperationExpand, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderWmsOperationExpandMgr) GetByOptions(opts ...Option) (results []*LgOrderWmsOperationExpand, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderWmsOperationExpandMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderWmsOperationExpand, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where(options.query)
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
func (obj *_LgOrderWmsOperationExpandMgr) GetFromID(id int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromID(ids []int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容  订单id
func (obj *_LgOrderWmsOperationExpandMgr) GetFromOrderID(orderID int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找  订单id
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromOrderID(orderIDs []int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromWarehouseID 通过warehouse_id获取内容  仓库id
func (obj *_LgOrderWmsOperationExpandMgr) GetFromWarehouseID(warehouseID int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_id` = ?", warehouseID).Find(&results).Error

	return
}

// GetBatchFromWarehouseID 批量查找  仓库id
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromWarehouseID(warehouseIDs []int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_id` IN (?)", warehouseIDs).Find(&results).Error

	return
}

// GetFromWarehouseCode 通过warehouse_code获取内容  仓库code 默认code是没有
func (obj *_LgOrderWmsOperationExpandMgr) GetFromWarehouseCode(warehouseCode string) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_code` = ?", warehouseCode).Find(&results).Error

	return
}

// GetBatchFromWarehouseCode 批量查找  仓库code 默认code是没有
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromWarehouseCode(warehouseCodes []string) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_code` IN (?)", warehouseCodes).Find(&results).Error

	return
}

// GetFromWarehouseName 通过warehouse_name获取内容  仓库名称 （默认仓库是没有仓库）
func (obj *_LgOrderWmsOperationExpandMgr) GetFromWarehouseName(warehouseName string) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_name` = ?", warehouseName).Find(&results).Error

	return
}

// GetBatchFromWarehouseName 批量查找  仓库名称 （默认仓库是没有仓库）
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromWarehouseName(warehouseNames []string) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_name` IN (?)", warehouseNames).Find(&results).Error

	return
}

// GetFromReceiptTime 通过receipt_time获取内容  入库时间
func (obj *_LgOrderWmsOperationExpandMgr) GetFromReceiptTime(receiptTime time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// GetBatchFromReceiptTime 批量查找  入库时间
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromReceiptTime(receiptTimes []time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`receipt_time` IN (?)", receiptTimes).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容  出库时间
func (obj *_LgOrderWmsOperationExpandMgr) GetFromSendTime(sendTime time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找  出库时间
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容  创建时间
func (obj *_LgOrderWmsOperationExpandMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找  创建时间
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容  创建用户
func (obj *_LgOrderWmsOperationExpandMgr) GetFromCreateUser(createUser int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找  创建用户
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容  更新时间
func (obj *_LgOrderWmsOperationExpandMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找  更新时间
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容  更新用户
func (obj *_LgOrderWmsOperationExpandMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找  更新用户
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容  乐观锁
func (obj *_LgOrderWmsOperationExpandMgr) GetFromVersion(version uint) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找  乐观锁
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromVersion(versions []uint) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容  逻辑删除
func (obj *_LgOrderWmsOperationExpandMgr) GetFromDeleted(deleted uint) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找  逻辑删除
func (obj *_LgOrderWmsOperationExpandMgr) GetBatchFromDeleted(deleteds []uint) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchByPrimaryKey(id int, orderID int) (result LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`id` = ? AND `order_id` = ?", id, orderID).First(&result).Error

	return
}

// FetchUniqueByUNIQUEORDERID primary or index 获取唯一内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchUniqueByUNIQUEORDERID(orderID int) (result LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// FetchIndexByNORMALORDERIDANDRECEIPTTIME  获取多个内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchIndexByNORMALORDERIDANDRECEIPTTIME(orderID int, receiptTime time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`order_id` = ? AND `receipt_time` = ?", orderID, receiptTime).Find(&results).Error

	return
}

// FetchIndexByNORMALORDERIDANDSENDTIME  获取多个内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchIndexByNORMALORDERIDANDSENDTIME(orderID int, sendTime time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`order_id` = ? AND `send_time` = ?", orderID, sendTime).Find(&results).Error

	return
}

// FetchIndexByNORMALWAREHOUSEID  获取多个内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchIndexByNORMALWAREHOUSEID(warehouseID int) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_id` = ?", warehouseID).Find(&results).Error

	return
}

// FetchIndexByNORMALCODE  获取多个内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchIndexByNORMALCODE(warehouseCode string) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_code` = ?", warehouseCode).Find(&results).Error

	return
}

// FetchIndexByNORMALNAME  获取多个内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchIndexByNORMALNAME(warehouseName string) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`warehouse_name` = ?", warehouseName).Find(&results).Error

	return
}

// FetchIndexByNORMALRECEIPTTIME  获取多个内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchIndexByNORMALRECEIPTTIME(receiptTime time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// FetchIndexByNORMALSENDTIME  获取多个内容
func (obj *_LgOrderWmsOperationExpandMgr) FetchIndexByNORMALSENDTIME(sendTime time.Time) (results []*LgOrderWmsOperationExpand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderWmsOperationExpand{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}
