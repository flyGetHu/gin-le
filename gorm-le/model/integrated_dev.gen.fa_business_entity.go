package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaBusinessEntityMgr struct {
	*_BaseMgr
}

// FaBusinessEntityMgr open func
func FaBusinessEntityMgr(db *gorm.DB) *_FaBusinessEntityMgr {
	if db == nil {
		panic(fmt.Errorf("FaBusinessEntityMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaBusinessEntityMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_business_entity"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaBusinessEntityMgr) GetTableName() string {
	return "fa_business_entity"
}

// Reset 重置gorm会话
func (obj *_FaBusinessEntityMgr) Reset() *_FaBusinessEntityMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaBusinessEntityMgr) Get() (result FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaBusinessEntityMgr) Gets() (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaBusinessEntityMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaBusinessEntityMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 业务主体代码
func (obj *_FaBusinessEntityMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithName name获取 业务主体名称
func (obj *_FaBusinessEntityMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithType type获取 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityMgr) WithType(_type bool) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithRemark remark获取 备注
func (obj *_FaBusinessEntityMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaBusinessEntityMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_FaBusinessEntityMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaBusinessEntityMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaBusinessEntityMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaBusinessEntityMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaBusinessEntityMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithKingdeeID kingdee_id获取 金蝶客户内码
func (obj *_FaBusinessEntityMgr) WithKingdeeID(kingdeeID string) Option {
	return optionFunc(func(o *options) { o.query["kingdee_id"] = kingdeeID })
}

// WithInputCumulativeFee input_cumulative_fee获取 入库预扣款
func (obj *_FaBusinessEntityMgr) WithInputCumulativeFee(inputCumulativeFee float64) Option {
	return optionFunc(func(o *options) { o.query["input_cumulative_fee"] = inputCumulativeFee })
}

// WithOutCumulativeFee out_cumulative_fee获取 出库未结算
func (obj *_FaBusinessEntityMgr) WithOutCumulativeFee(outCumulativeFee float64) Option {
	return optionFunc(func(o *options) { o.query["out_cumulative_fee"] = outCumulativeFee })
}

// GetByOption 功能选项模式获取
func (obj *_FaBusinessEntityMgr) GetByOption(opts ...Option) (result FaBusinessEntity, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaBusinessEntityMgr) GetByOptions(opts ...Option) (results []*FaBusinessEntity, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaBusinessEntityMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaBusinessEntity, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where(options.query)
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
func (obj *_FaBusinessEntityMgr) GetFromID(id int64) (result FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaBusinessEntityMgr) GetBatchFromID(ids []int64) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 业务主体代码
func (obj *_FaBusinessEntityMgr) GetFromCode(code string) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 业务主体代码
func (obj *_FaBusinessEntityMgr) GetBatchFromCode(codes []string) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 业务主体名称
func (obj *_FaBusinessEntityMgr) GetFromName(name string) (result FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`name` = ?", name).First(&result).Error

	return
}

// GetBatchFromName 批量查找 业务主体名称
func (obj *_FaBusinessEntityMgr) GetBatchFromName(names []string) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityMgr) GetFromType(_type bool) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 业务主体类型，0:服务商，1:航司
func (obj *_FaBusinessEntityMgr) GetBatchFromType(_types []bool) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaBusinessEntityMgr) GetFromRemark(remark string) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaBusinessEntityMgr) GetBatchFromRemark(remarks []string) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaBusinessEntityMgr) GetFromCreateTime(createTime time.Time) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaBusinessEntityMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_FaBusinessEntityMgr) GetFromCreateUser(createUser int) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_FaBusinessEntityMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaBusinessEntityMgr) GetFromUpdateUser(updateUser int) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaBusinessEntityMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaBusinessEntityMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaBusinessEntityMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaBusinessEntityMgr) GetFromVersion(version int) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaBusinessEntityMgr) GetBatchFromVersion(versions []int) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaBusinessEntityMgr) GetFromDeleted(deleted int) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaBusinessEntityMgr) GetBatchFromDeleted(deleteds []int) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromKingdeeID 通过kingdee_id获取内容 金蝶客户内码
func (obj *_FaBusinessEntityMgr) GetFromKingdeeID(kingdeeID string) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}

// GetBatchFromKingdeeID 批量查找 金蝶客户内码
func (obj *_FaBusinessEntityMgr) GetBatchFromKingdeeID(kingdeeIDs []string) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`kingdee_id` IN (?)", kingdeeIDs).Find(&results).Error

	return
}

// GetFromInputCumulativeFee 通过input_cumulative_fee获取内容 入库预扣款
func (obj *_FaBusinessEntityMgr) GetFromInputCumulativeFee(inputCumulativeFee float64) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`input_cumulative_fee` = ?", inputCumulativeFee).Find(&results).Error

	return
}

// GetBatchFromInputCumulativeFee 批量查找 入库预扣款
func (obj *_FaBusinessEntityMgr) GetBatchFromInputCumulativeFee(inputCumulativeFees []float64) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`input_cumulative_fee` IN (?)", inputCumulativeFees).Find(&results).Error

	return
}

// GetFromOutCumulativeFee 通过out_cumulative_fee获取内容 出库未结算
func (obj *_FaBusinessEntityMgr) GetFromOutCumulativeFee(outCumulativeFee float64) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`out_cumulative_fee` = ?", outCumulativeFee).Find(&results).Error

	return
}

// GetBatchFromOutCumulativeFee 批量查找 出库未结算
func (obj *_FaBusinessEntityMgr) GetBatchFromOutCumulativeFee(outCumulativeFees []float64) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`out_cumulative_fee` IN (?)", outCumulativeFees).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaBusinessEntityMgr) FetchByPrimaryKey(id int64) (result FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUnidexCode primary or index 获取唯一内容
func (obj *_FaBusinessEntityMgr) FetchUniqueByUnidexCode(code string) (result FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`code` = ?", code).First(&result).Error

	return
}

// FetchUniqueByUnidexName primary or index 获取唯一内容
func (obj *_FaBusinessEntityMgr) FetchUniqueByUnidexName(name string) (result FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`name` = ?", name).First(&result).Error

	return
}

// FetchIndexByIndexCodeType  获取多个内容
func (obj *_FaBusinessEntityMgr) FetchIndexByIndexCodeType(code string, _type bool) (results []*FaBusinessEntity, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBusinessEntity{}).Where("`code` = ? AND `type` = ?", code, _type).Find(&results).Error

	return
}
