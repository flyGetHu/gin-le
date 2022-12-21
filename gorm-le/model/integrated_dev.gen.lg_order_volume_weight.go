package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderVolumeWeightMgr struct {
	*_BaseMgr
}

// LgOrderVolumeWeightMgr open func
func LgOrderVolumeWeightMgr(db *gorm.DB) *_LgOrderVolumeWeightMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderVolumeWeightMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderVolumeWeightMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_volume_weight"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderVolumeWeightMgr) GetTableName() string {
	return "lg_order_volume_weight"
}

// Reset 重置gorm会话
func (obj *_LgOrderVolumeWeightMgr) Reset() *_LgOrderVolumeWeightMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderVolumeWeightMgr) Get() (result LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderVolumeWeightMgr) Gets() (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderVolumeWeightMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgOrderVolumeWeightMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_LgOrderVolumeWeightMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLength length获取 长
func (obj *_LgOrderVolumeWeightMgr) WithLength(length float64) Option {
	return optionFunc(func(o *options) { o.query["length"] = length })
}

// WithWidth width获取 宽
func (obj *_LgOrderVolumeWeightMgr) WithWidth(width float64) Option {
	return optionFunc(func(o *options) { o.query["width"] = width })
}

// WithHeight height获取 高
func (obj *_LgOrderVolumeWeightMgr) WithHeight(height float64) Option {
	return optionFunc(func(o *options) { o.query["height"] = height })
}

// WithWeight weight获取 称重重量
func (obj *_LgOrderVolumeWeightMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithVolumeWeight volume_weight获取 材积重
func (obj *_LgOrderVolumeWeightMgr) WithVolumeWeight(volumeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["volume_weight"] = volumeWeight })
}

// WithProportion proportion获取 抛比
func (obj *_LgOrderVolumeWeightMgr) WithProportion(proportion int) Option {
	return optionFunc(func(o *options) { o.query["proportion"] = proportion })
}

// WithWeightDifference weight_difference获取 重量差
func (obj *_LgOrderVolumeWeightMgr) WithWeightDifference(weightDifference float64) Option {
	return optionFunc(func(o *options) { o.query["weight_difference"] = weightDifference })
}

// WithProcessingStatus processing_status获取 处理状态，initial:待处理，done:已处理
func (obj *_LgOrderVolumeWeightMgr) WithProcessingStatus(processingStatus string) Option {
	return optionFunc(func(o *options) { o.query["processing_status"] = processingStatus })
}

// WithProcessingUser processing_user获取 处理状态人
func (obj *_LgOrderVolumeWeightMgr) WithProcessingUser(processingUser int) Option {
	return optionFunc(func(o *options) { o.query["processing_user"] = processingUser })
}

// WithProcessingTime processing_time获取 处理时间
func (obj *_LgOrderVolumeWeightMgr) WithProcessingTime(processingTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["processing_time"] = processingTime })
}

// WithRemark remark获取 说明
func (obj *_LgOrderVolumeWeightMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateUser create_user获取 创建人
func (obj *_LgOrderVolumeWeightMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderVolumeWeightMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_LgOrderVolumeWeightMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderVolumeWeightMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderVolumeWeightMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderVolumeWeightMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderVolumeWeightMgr) GetByOption(opts ...Option) (result LgOrderVolumeWeight, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderVolumeWeightMgr) GetByOptions(opts ...Option) (results []*LgOrderVolumeWeight, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderVolumeWeightMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderVolumeWeight, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where(options.query)
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
func (obj *_LgOrderVolumeWeightMgr) GetFromID(id int) (result LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromID(ids []int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_LgOrderVolumeWeightMgr) GetFromLogisticsNumber(logisticsNumber string) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLength 通过length获取内容 长
func (obj *_LgOrderVolumeWeightMgr) GetFromLength(length float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`length` = ?", length).Find(&results).Error

	return
}

// GetBatchFromLength 批量查找 长
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromLength(lengths []float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`length` IN (?)", lengths).Find(&results).Error

	return
}

// GetFromWidth 通过width获取内容 宽
func (obj *_LgOrderVolumeWeightMgr) GetFromWidth(width float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`width` = ?", width).Find(&results).Error

	return
}

// GetBatchFromWidth 批量查找 宽
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromWidth(widths []float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`width` IN (?)", widths).Find(&results).Error

	return
}

// GetFromHeight 通过height获取内容 高
func (obj *_LgOrderVolumeWeightMgr) GetFromHeight(height float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`height` = ?", height).Find(&results).Error

	return
}

// GetBatchFromHeight 批量查找 高
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromHeight(heights []float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`height` IN (?)", heights).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 称重重量
func (obj *_LgOrderVolumeWeightMgr) GetFromWeight(weight float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 称重重量
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromWeight(weights []float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromVolumeWeight 通过volume_weight获取内容 材积重
func (obj *_LgOrderVolumeWeightMgr) GetFromVolumeWeight(volumeWeight float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`volume_weight` = ?", volumeWeight).Find(&results).Error

	return
}

// GetBatchFromVolumeWeight 批量查找 材积重
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromVolumeWeight(volumeWeights []float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`volume_weight` IN (?)", volumeWeights).Find(&results).Error

	return
}

// GetFromProportion 通过proportion获取内容 抛比
func (obj *_LgOrderVolumeWeightMgr) GetFromProportion(proportion int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`proportion` = ?", proportion).Find(&results).Error

	return
}

// GetBatchFromProportion 批量查找 抛比
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromProportion(proportions []int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`proportion` IN (?)", proportions).Find(&results).Error

	return
}

// GetFromWeightDifference 通过weight_difference获取内容 重量差
func (obj *_LgOrderVolumeWeightMgr) GetFromWeightDifference(weightDifference float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`weight_difference` = ?", weightDifference).Find(&results).Error

	return
}

// GetBatchFromWeightDifference 批量查找 重量差
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromWeightDifference(weightDifferences []float64) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`weight_difference` IN (?)", weightDifferences).Find(&results).Error

	return
}

// GetFromProcessingStatus 通过processing_status获取内容 处理状态，initial:待处理，done:已处理
func (obj *_LgOrderVolumeWeightMgr) GetFromProcessingStatus(processingStatus string) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`processing_status` = ?", processingStatus).Find(&results).Error

	return
}

// GetBatchFromProcessingStatus 批量查找 处理状态，initial:待处理，done:已处理
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromProcessingStatus(processingStatuss []string) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`processing_status` IN (?)", processingStatuss).Find(&results).Error

	return
}

// GetFromProcessingUser 通过processing_user获取内容 处理状态人
func (obj *_LgOrderVolumeWeightMgr) GetFromProcessingUser(processingUser int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`processing_user` = ?", processingUser).Find(&results).Error

	return
}

// GetBatchFromProcessingUser 批量查找 处理状态人
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromProcessingUser(processingUsers []int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`processing_user` IN (?)", processingUsers).Find(&results).Error

	return
}

// GetFromProcessingTime 通过processing_time获取内容 处理时间
func (obj *_LgOrderVolumeWeightMgr) GetFromProcessingTime(processingTime time.Time) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`processing_time` = ?", processingTime).Find(&results).Error

	return
}

// GetBatchFromProcessingTime 批量查找 处理时间
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromProcessingTime(processingTimes []time.Time) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`processing_time` IN (?)", processingTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 说明
func (obj *_LgOrderVolumeWeightMgr) GetFromRemark(remark string) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 说明
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromRemark(remarks []string) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_LgOrderVolumeWeightMgr) GetFromCreateUser(createUser int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderVolumeWeightMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_LgOrderVolumeWeightMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderVolumeWeightMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderVolumeWeightMgr) GetFromVersion(version int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromVersion(versions []int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderVolumeWeightMgr) GetFromDeleted(deleted int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderVolumeWeightMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderVolumeWeightMgr) FetchByPrimaryKey(id int) (result LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgOrderVolumeWeightIDUIndex primary or index 获取唯一内容
func (obj *_LgOrderVolumeWeightMgr) FetchUniqueByLgOrderVolumeWeightIDUIndex(id int) (result LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByTableNameLogisticsNumberIndex  获取多个内容
func (obj *_LgOrderVolumeWeightMgr) FetchIndexByTableNameLogisticsNumberIndex(logisticsNumber string) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`logistics_number` = ?", logisticsNumber).Find(&results).Error

	return
}

// FetchIndexByTableNameCreateTimeIndex  获取多个内容
func (obj *_LgOrderVolumeWeightMgr) FetchIndexByTableNameCreateTimeIndex(createTime time.Time) (results []*LgOrderVolumeWeight, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderVolumeWeight{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}
