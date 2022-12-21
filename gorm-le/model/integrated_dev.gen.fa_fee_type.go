package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaFeeTypeMgr struct {
	*_BaseMgr
}

// FaFeeTypeMgr open func
func FaFeeTypeMgr(db *gorm.DB) *_FaFeeTypeMgr {
	if db == nil {
		panic(fmt.Errorf("FaFeeTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaFeeTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_fee_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaFeeTypeMgr) GetTableName() string {
	return "fa_fee_type"
}

// Reset 重置gorm会话
func (obj *_FaFeeTypeMgr) Reset() *_FaFeeTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaFeeTypeMgr) Get() (result FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaFeeTypeMgr) Gets() (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaFeeTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaFeeTypeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithFeeName fee_name获取 费用名称
func (obj *_FaFeeTypeMgr) WithFeeName(feeName string) Option {
	return optionFunc(func(o *options) { o.query["fee_name"] = feeName })
}

// WithRemark remark获取 备注
func (obj *_FaFeeTypeMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithState state获取 状态(0:禁用1:启用)
func (obj *_FaFeeTypeMgr) WithState(state []uint8) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithFeeType fee_type获取 费用种类(1:应收费用,2:应付费用,3:即是应收也是应付)
func (obj *_FaFeeTypeMgr) WithFeeType(feeType int) Option {
	return optionFunc(func(o *options) { o.query["fee_type"] = feeType })
}

// WithFeeState fee_state获取 0:收款 1:退件退款 2:部分退款
func (obj *_FaFeeTypeMgr) WithFeeState(feeState int) Option {
	return optionFunc(func(o *options) { o.query["fee_state"] = feeState })
}

// WithCreateUser create_user获取 创建人
func (obj *_FaFeeTypeMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaFeeTypeMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaFeeTypeMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaFeeTypeMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaFeeTypeMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_FaFeeTypeMgr) GetByOption(opts ...Option) (result FaFeeType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaFeeTypeMgr) GetByOptions(opts ...Option) (results []*FaFeeType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaFeeTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaFeeType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_FaFeeTypeMgr) GetFromID(id int) (result FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaFeeTypeMgr) GetBatchFromID(ids []int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromFeeName 通过fee_name获取内容 费用名称
func (obj *_FaFeeTypeMgr) GetFromFeeName(feeName string) (result FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`fee_name` = ?", feeName).First(&result).Error

	return
}

// GetBatchFromFeeName 批量查找 费用名称
func (obj *_FaFeeTypeMgr) GetBatchFromFeeName(feeNames []string) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`fee_name` IN (?)", feeNames).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaFeeTypeMgr) GetFromRemark(remark string) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaFeeTypeMgr) GetBatchFromRemark(remarks []string) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromState 通过state获取内容 状态(0:禁用1:启用)
func (obj *_FaFeeTypeMgr) GetFromState(state []uint8) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// GetBatchFromState 批量查找 状态(0:禁用1:启用)
func (obj *_FaFeeTypeMgr) GetBatchFromState(states [][]uint8) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`state` IN (?)", states).Find(&results).Error

	return
}

// GetFromFeeType 通过fee_type获取内容 费用种类(1:应收费用,2:应付费用,3:即是应收也是应付)
func (obj *_FaFeeTypeMgr) GetFromFeeType(feeType int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`fee_type` = ?", feeType).Find(&results).Error

	return
}

// GetBatchFromFeeType 批量查找 费用种类(1:应收费用,2:应付费用,3:即是应收也是应付)
func (obj *_FaFeeTypeMgr) GetBatchFromFeeType(feeTypes []int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`fee_type` IN (?)", feeTypes).Find(&results).Error

	return
}

// GetFromFeeState 通过fee_state获取内容 0:收款 1:退件退款 2:部分退款
func (obj *_FaFeeTypeMgr) GetFromFeeState(feeState int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`fee_state` = ?", feeState).Find(&results).Error

	return
}

// GetBatchFromFeeState 批量查找 0:收款 1:退件退款 2:部分退款
func (obj *_FaFeeTypeMgr) GetBatchFromFeeState(feeStates []int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`fee_state` IN (?)", feeStates).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_FaFeeTypeMgr) GetFromCreateUser(createUser int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_FaFeeTypeMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaFeeTypeMgr) GetFromCreateTime(createTime time.Time) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaFeeTypeMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaFeeTypeMgr) GetFromUpdateUser(updateUser int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaFeeTypeMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaFeeTypeMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaFeeTypeMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaFeeTypeMgr) GetFromVersion(version int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaFeeTypeMgr) GetBatchFromVersion(versions []int) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaFeeTypeMgr) FetchByPrimaryKey(id int) (result FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByFaFeeTypeIDUIndex primary or index 获取唯一内容
func (obj *_FaFeeTypeMgr) FetchUniqueByFaFeeTypeIDUIndex(id int) (result FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByFaFeeTypeFeeNameUIndex primary or index 获取唯一内容
func (obj *_FaFeeTypeMgr) FetchUniqueByFaFeeTypeFeeNameUIndex(feeName string) (result FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`fee_name` = ?", feeName).First(&result).Error

	return
}

// FetchIndexByFaFeeTypeStateIndex  获取多个内容
func (obj *_FaFeeTypeMgr) FetchIndexByFaFeeTypeStateIndex(state []uint8) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`state` = ?", state).Find(&results).Error

	return
}

// FetchIndexByFaFeeTypeCreateTimeIndex  获取多个内容
func (obj *_FaFeeTypeMgr) FetchIndexByFaFeeTypeCreateTimeIndex(createTime time.Time) (results []*FaFeeType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaFeeType{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
