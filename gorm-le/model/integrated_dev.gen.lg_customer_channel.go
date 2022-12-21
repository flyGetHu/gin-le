package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgCustomerChannelMgr struct {
	*_BaseMgr
}

// LgCustomerChannelMgr open func
func LgCustomerChannelMgr(db *gorm.DB) *_LgCustomerChannelMgr {
	if db == nil {
		panic(fmt.Errorf("LgCustomerChannelMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgCustomerChannelMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_customer_channel"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgCustomerChannelMgr) GetTableName() string {
	return "lg_customer_channel"
}

// Reset 重置gorm会话
func (obj *_LgCustomerChannelMgr) Reset() *_LgCustomerChannelMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgCustomerChannelMgr) Get() (result LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgCustomerChannelMgr) Gets() (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgCustomerChannelMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 渠道ID
func (obj *_LgCustomerChannelMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerChannelType customer_channel_type获取 渠道类型
func (obj *_LgCustomerChannelMgr) WithCustomerChannelType(customerChannelType int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_type"] = customerChannelType })
}

// WithCustomerChannelName customer_channel_name获取 渠道名称
func (obj *_LgCustomerChannelMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithMinDeclare min_declare获取 最低申报价值
func (obj *_LgCustomerChannelMgr) WithMinDeclare(minDeclare float64) Option {
	return optionFunc(func(o *options) { o.query["min_declare"] = minDeclare })
}

// WithMaxDeclare max_declare获取 最高申报价值
func (obj *_LgCustomerChannelMgr) WithMaxDeclare(maxDeclare float64) Option {
	return optionFunc(func(o *options) { o.query["max_declare"] = maxDeclare })
}

// WithMinWeight min_weight获取 最低下单重量
func (obj *_LgCustomerChannelMgr) WithMinWeight(minWeight float64) Option {
	return optionFunc(func(o *options) { o.query["min_weight"] = minWeight })
}

// WithMaxWeight max_weight获取 最高下单重量
func (obj *_LgCustomerChannelMgr) WithMaxWeight(maxWeight float64) Option {
	return optionFunc(func(o *options) { o.query["max_weight"] = maxWeight })
}

// WithMinSendWeight min_send_weight获取 最小出库重量
func (obj *_LgCustomerChannelMgr) WithMinSendWeight(minSendWeight float64) Option {
	return optionFunc(func(o *options) { o.query["min_send_weight"] = minSendWeight })
}

// WithMaxSendWeight max_send_weight获取 最大出库重量
func (obj *_LgCustomerChannelMgr) WithMaxSendWeight(maxSendWeight float64) Option {
	return optionFunc(func(o *options) { o.query["max_send_weight"] = maxSendWeight })
}

// WithInStorageUseOrderWeightFlag in_storage_use_order_weight_flag获取 入库是否使用订单预报重量，0:false，1:true，默认0
func (obj *_LgCustomerChannelMgr) WithInStorageUseOrderWeightFlag(inStorageUseOrderWeightFlag bool) Option {
	return optionFunc(func(o *options) { o.query["in_storage_use_order_weight_flag"] = inStorageUseOrderWeightFlag })
}

// WithInStorageOrderWeightMin in_storage_order_weight_min获取 入库最小预报重量，单位:KG
func (obj *_LgCustomerChannelMgr) WithInStorageOrderWeightMin(inStorageOrderWeightMin float64) Option {
	return optionFunc(func(o *options) { o.query["in_storage_order_weight_min"] = inStorageOrderWeightMin })
}

// WithInStorageOrderWeightMax in_storage_order_weight_max获取 入库最大预报重量，单位:KG
func (obj *_LgCustomerChannelMgr) WithInStorageOrderWeightMax(inStorageOrderWeightMax float64) Option {
	return optionFunc(func(o *options) { o.query["in_storage_order_weight_max"] = inStorageOrderWeightMax })
}

// WithSenderBlackList sender_black_list获取 寄件人黑名单
func (obj *_LgCustomerChannelMgr) WithSenderBlackList(senderBlackList string) Option {
	return optionFunc(func(o *options) { o.query["sender_black_list"] = senderBlackList })
}

// WithBanGoodsList ban_goods_list获取 禁止商品
func (obj *_LgCustomerChannelMgr) WithBanGoodsList(banGoodsList string) Option {
	return optionFunc(func(o *options) { o.query["ban_goods_list"] = banGoodsList })
}

// WithNotePublic note_public获取 客户渠道说明_公开
func (obj *_LgCustomerChannelMgr) WithNotePublic(notePublic string) Option {
	return optionFunc(func(o *options) { o.query["note_public"] = notePublic })
}

// WithNotePrivate note_private获取 客户渠道说明_私有
func (obj *_LgCustomerChannelMgr) WithNotePrivate(notePrivate string) Option {
	return optionFunc(func(o *options) { o.query["note_private"] = notePrivate })
}

// WithOrderNumPoolID order_num_pool_id获取 绑定的单号池ID
func (obj *_LgCustomerChannelMgr) WithOrderNumPoolID(orderNumPoolID int64) Option {
	return optionFunc(func(o *options) { o.query["order_num_pool_id"] = orderNumPoolID })
}

// WithOrderNumPoolName order_num_pool_name获取 绑定的单号池名称
func (obj *_LgCustomerChannelMgr) WithOrderNumPoolName(orderNumPoolName string) Option {
	return optionFunc(func(o *options) { o.query["order_num_pool_name"] = orderNumPoolName })
}

// WithExchangeOrder exchange_order获取 是否需要换单0 不需要1 需要
func (obj *_LgCustomerChannelMgr) WithExchangeOrder(exchangeOrder []uint8) Option {
	return optionFunc(func(o *options) { o.query["exchange_order"] = exchangeOrder })
}

// WithExchangeType exchange_type获取 换单类型 1:只换面单不换号 2:即换单又换号
func (obj *_LgCustomerChannelMgr) WithExchangeType(exchangeType int) Option {
	return optionFunc(func(o *options) { o.query["exchange_type"] = exchangeType })
}

// WithLabelCode label_code获取 渠道为换单渠道则选取制定的面单code
func (obj *_LgCustomerChannelMgr) WithLabelCode(labelCode string) Option {
	return optionFunc(func(o *options) { o.query["label_code"] = labelCode })
}

// WithDataPushAPI data_push_api获取 订单数据推送地址
func (obj *_LgCustomerChannelMgr) WithDataPushAPI(dataPushAPI string) Option {
	return optionFunc(func(o *options) { o.query["data_push_api"] = dataPushAPI })
}

// WithProportion proportion获取 抛比
func (obj *_LgCustomerChannelMgr) WithProportion(proportion int) Option {
	return optionFunc(func(o *options) { o.query["proportion"] = proportion })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgCustomerChannelMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgCustomerChannelMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgCustomerChannelMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgCustomerChannelMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgCustomerChannelMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgCustomerChannelMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithStatus status获取 状态，0:无效，1:有效
func (obj *_LgCustomerChannelMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithRequiredPlatformNo required_platform_no获取 是否必填平台订单号，0-不必填，1-必填
func (obj *_LgCustomerChannelMgr) WithRequiredPlatformNo(requiredPlatformNo []uint8) Option {
	return optionFunc(func(o *options) { o.query["required_platform_no"] = requiredPlatformNo })
}

// WithChannelHaul channel_haul获取 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_LgCustomerChannelMgr) WithChannelHaul(channelHaul string) Option {
	return optionFunc(func(o *options) { o.query["channel_haul"] = channelHaul })
}

// WithKingdeeID kingdee_id获取 金蝶客户内码
func (obj *_LgCustomerChannelMgr) WithKingdeeID(kingdeeID string) Option {
	return optionFunc(func(o *options) { o.query["kingdee_id"] = kingdeeID })
}

// WithCustomerChannelTypeID customer_channel_type_id获取 客户渠道类型id
func (obj *_LgCustomerChannelMgr) WithCustomerChannelTypeID(customerChannelTypeID string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_type_id"] = customerChannelTypeID })
}

// WithChannelCountryAffiliationID channel_country_affiliation_id获取 渠道所属国家id,lg_area，对应lg_area表的id（注意：只有level 为0 的才是国家）
func (obj *_LgCustomerChannelMgr) WithChannelCountryAffiliationID(channelCountryAffiliationID int) Option {
	return optionFunc(func(o *options) { o.query["channel_country_affiliation_id"] = channelCountryAffiliationID })
}

// GetByOption 功能选项模式获取
func (obj *_LgCustomerChannelMgr) GetByOption(opts ...Option) (result LgCustomerChannel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgCustomerChannelMgr) GetByOptions(opts ...Option) (results []*LgCustomerChannel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgCustomerChannelMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgCustomerChannel, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where(options.query)
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

// GetFromID 通过id获取内容 渠道ID
func (obj *_LgCustomerChannelMgr) GetFromID(id int) (result LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 渠道ID
func (obj *_LgCustomerChannelMgr) GetBatchFromID(ids []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCustomerChannelType 通过customer_channel_type获取内容 渠道类型
func (obj *_LgCustomerChannelMgr) GetFromCustomerChannelType(customerChannelType int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`customer_channel_type` = ?", customerChannelType).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelType 批量查找 渠道类型
func (obj *_LgCustomerChannelMgr) GetBatchFromCustomerChannelType(customerChannelTypes []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`customer_channel_type` IN (?)", customerChannelTypes).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 渠道名称
func (obj *_LgCustomerChannelMgr) GetFromCustomerChannelName(customerChannelName string) (result LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`customer_channel_name` = ?", customerChannelName).First(&result).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 渠道名称
func (obj *_LgCustomerChannelMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromMinDeclare 通过min_declare获取内容 最低申报价值
func (obj *_LgCustomerChannelMgr) GetFromMinDeclare(minDeclare float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`min_declare` = ?", minDeclare).Find(&results).Error

	return
}

// GetBatchFromMinDeclare 批量查找 最低申报价值
func (obj *_LgCustomerChannelMgr) GetBatchFromMinDeclare(minDeclares []float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`min_declare` IN (?)", minDeclares).Find(&results).Error

	return
}

// GetFromMaxDeclare 通过max_declare获取内容 最高申报价值
func (obj *_LgCustomerChannelMgr) GetFromMaxDeclare(maxDeclare float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`max_declare` = ?", maxDeclare).Find(&results).Error

	return
}

// GetBatchFromMaxDeclare 批量查找 最高申报价值
func (obj *_LgCustomerChannelMgr) GetBatchFromMaxDeclare(maxDeclares []float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`max_declare` IN (?)", maxDeclares).Find(&results).Error

	return
}

// GetFromMinWeight 通过min_weight获取内容 最低下单重量
func (obj *_LgCustomerChannelMgr) GetFromMinWeight(minWeight float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`min_weight` = ?", minWeight).Find(&results).Error

	return
}

// GetBatchFromMinWeight 批量查找 最低下单重量
func (obj *_LgCustomerChannelMgr) GetBatchFromMinWeight(minWeights []float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`min_weight` IN (?)", minWeights).Find(&results).Error

	return
}

// GetFromMaxWeight 通过max_weight获取内容 最高下单重量
func (obj *_LgCustomerChannelMgr) GetFromMaxWeight(maxWeight float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`max_weight` = ?", maxWeight).Find(&results).Error

	return
}

// GetBatchFromMaxWeight 批量查找 最高下单重量
func (obj *_LgCustomerChannelMgr) GetBatchFromMaxWeight(maxWeights []float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`max_weight` IN (?)", maxWeights).Find(&results).Error

	return
}

// GetFromMinSendWeight 通过min_send_weight获取内容 最小出库重量
func (obj *_LgCustomerChannelMgr) GetFromMinSendWeight(minSendWeight float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`min_send_weight` = ?", minSendWeight).Find(&results).Error

	return
}

// GetBatchFromMinSendWeight 批量查找 最小出库重量
func (obj *_LgCustomerChannelMgr) GetBatchFromMinSendWeight(minSendWeights []float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`min_send_weight` IN (?)", minSendWeights).Find(&results).Error

	return
}

// GetFromMaxSendWeight 通过max_send_weight获取内容 最大出库重量
func (obj *_LgCustomerChannelMgr) GetFromMaxSendWeight(maxSendWeight float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`max_send_weight` = ?", maxSendWeight).Find(&results).Error

	return
}

// GetBatchFromMaxSendWeight 批量查找 最大出库重量
func (obj *_LgCustomerChannelMgr) GetBatchFromMaxSendWeight(maxSendWeights []float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`max_send_weight` IN (?)", maxSendWeights).Find(&results).Error

	return
}

// GetFromInStorageUseOrderWeightFlag 通过in_storage_use_order_weight_flag获取内容 入库是否使用订单预报重量，0:false，1:true，默认0
func (obj *_LgCustomerChannelMgr) GetFromInStorageUseOrderWeightFlag(inStorageUseOrderWeightFlag bool) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`in_storage_use_order_weight_flag` = ?", inStorageUseOrderWeightFlag).Find(&results).Error

	return
}

// GetBatchFromInStorageUseOrderWeightFlag 批量查找 入库是否使用订单预报重量，0:false，1:true，默认0
func (obj *_LgCustomerChannelMgr) GetBatchFromInStorageUseOrderWeightFlag(inStorageUseOrderWeightFlags []bool) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`in_storage_use_order_weight_flag` IN (?)", inStorageUseOrderWeightFlags).Find(&results).Error

	return
}

// GetFromInStorageOrderWeightMin 通过in_storage_order_weight_min获取内容 入库最小预报重量，单位:KG
func (obj *_LgCustomerChannelMgr) GetFromInStorageOrderWeightMin(inStorageOrderWeightMin float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`in_storage_order_weight_min` = ?", inStorageOrderWeightMin).Find(&results).Error

	return
}

// GetBatchFromInStorageOrderWeightMin 批量查找 入库最小预报重量，单位:KG
func (obj *_LgCustomerChannelMgr) GetBatchFromInStorageOrderWeightMin(inStorageOrderWeightMins []float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`in_storage_order_weight_min` IN (?)", inStorageOrderWeightMins).Find(&results).Error

	return
}

// GetFromInStorageOrderWeightMax 通过in_storage_order_weight_max获取内容 入库最大预报重量，单位:KG
func (obj *_LgCustomerChannelMgr) GetFromInStorageOrderWeightMax(inStorageOrderWeightMax float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`in_storage_order_weight_max` = ?", inStorageOrderWeightMax).Find(&results).Error

	return
}

// GetBatchFromInStorageOrderWeightMax 批量查找 入库最大预报重量，单位:KG
func (obj *_LgCustomerChannelMgr) GetBatchFromInStorageOrderWeightMax(inStorageOrderWeightMaxs []float64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`in_storage_order_weight_max` IN (?)", inStorageOrderWeightMaxs).Find(&results).Error

	return
}

// GetFromSenderBlackList 通过sender_black_list获取内容 寄件人黑名单
func (obj *_LgCustomerChannelMgr) GetFromSenderBlackList(senderBlackList string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`sender_black_list` = ?", senderBlackList).Find(&results).Error

	return
}

// GetBatchFromSenderBlackList 批量查找 寄件人黑名单
func (obj *_LgCustomerChannelMgr) GetBatchFromSenderBlackList(senderBlackLists []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`sender_black_list` IN (?)", senderBlackLists).Find(&results).Error

	return
}

// GetFromBanGoodsList 通过ban_goods_list获取内容 禁止商品
func (obj *_LgCustomerChannelMgr) GetFromBanGoodsList(banGoodsList string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`ban_goods_list` = ?", banGoodsList).Find(&results).Error

	return
}

// GetBatchFromBanGoodsList 批量查找 禁止商品
func (obj *_LgCustomerChannelMgr) GetBatchFromBanGoodsList(banGoodsLists []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`ban_goods_list` IN (?)", banGoodsLists).Find(&results).Error

	return
}

// GetFromNotePublic 通过note_public获取内容 客户渠道说明_公开
func (obj *_LgCustomerChannelMgr) GetFromNotePublic(notePublic string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`note_public` = ?", notePublic).Find(&results).Error

	return
}

// GetBatchFromNotePublic 批量查找 客户渠道说明_公开
func (obj *_LgCustomerChannelMgr) GetBatchFromNotePublic(notePublics []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`note_public` IN (?)", notePublics).Find(&results).Error

	return
}

// GetFromNotePrivate 通过note_private获取内容 客户渠道说明_私有
func (obj *_LgCustomerChannelMgr) GetFromNotePrivate(notePrivate string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`note_private` = ?", notePrivate).Find(&results).Error

	return
}

// GetBatchFromNotePrivate 批量查找 客户渠道说明_私有
func (obj *_LgCustomerChannelMgr) GetBatchFromNotePrivate(notePrivates []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`note_private` IN (?)", notePrivates).Find(&results).Error

	return
}

// GetFromOrderNumPoolID 通过order_num_pool_id获取内容 绑定的单号池ID
func (obj *_LgCustomerChannelMgr) GetFromOrderNumPoolID(orderNumPoolID int64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`order_num_pool_id` = ?", orderNumPoolID).Find(&results).Error

	return
}

// GetBatchFromOrderNumPoolID 批量查找 绑定的单号池ID
func (obj *_LgCustomerChannelMgr) GetBatchFromOrderNumPoolID(orderNumPoolIDs []int64) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`order_num_pool_id` IN (?)", orderNumPoolIDs).Find(&results).Error

	return
}

// GetFromOrderNumPoolName 通过order_num_pool_name获取内容 绑定的单号池名称
func (obj *_LgCustomerChannelMgr) GetFromOrderNumPoolName(orderNumPoolName string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`order_num_pool_name` = ?", orderNumPoolName).Find(&results).Error

	return
}

// GetBatchFromOrderNumPoolName 批量查找 绑定的单号池名称
func (obj *_LgCustomerChannelMgr) GetBatchFromOrderNumPoolName(orderNumPoolNames []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`order_num_pool_name` IN (?)", orderNumPoolNames).Find(&results).Error

	return
}

// GetFromExchangeOrder 通过exchange_order获取内容 是否需要换单0 不需要1 需要
func (obj *_LgCustomerChannelMgr) GetFromExchangeOrder(exchangeOrder []uint8) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`exchange_order` = ?", exchangeOrder).Find(&results).Error

	return
}

// GetBatchFromExchangeOrder 批量查找 是否需要换单0 不需要1 需要
func (obj *_LgCustomerChannelMgr) GetBatchFromExchangeOrder(exchangeOrders [][]uint8) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`exchange_order` IN (?)", exchangeOrders).Find(&results).Error

	return
}

// GetFromExchangeType 通过exchange_type获取内容 换单类型 1:只换面单不换号 2:即换单又换号
func (obj *_LgCustomerChannelMgr) GetFromExchangeType(exchangeType int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`exchange_type` = ?", exchangeType).Find(&results).Error

	return
}

// GetBatchFromExchangeType 批量查找 换单类型 1:只换面单不换号 2:即换单又换号
func (obj *_LgCustomerChannelMgr) GetBatchFromExchangeType(exchangeTypes []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`exchange_type` IN (?)", exchangeTypes).Find(&results).Error

	return
}

// GetFromLabelCode 通过label_code获取内容 渠道为换单渠道则选取制定的面单code
func (obj *_LgCustomerChannelMgr) GetFromLabelCode(labelCode string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`label_code` = ?", labelCode).Find(&results).Error

	return
}

// GetBatchFromLabelCode 批量查找 渠道为换单渠道则选取制定的面单code
func (obj *_LgCustomerChannelMgr) GetBatchFromLabelCode(labelCodes []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`label_code` IN (?)", labelCodes).Find(&results).Error

	return
}

// GetFromDataPushAPI 通过data_push_api获取内容 订单数据推送地址
func (obj *_LgCustomerChannelMgr) GetFromDataPushAPI(dataPushAPI string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`data_push_api` = ?", dataPushAPI).Find(&results).Error

	return
}

// GetBatchFromDataPushAPI 批量查找 订单数据推送地址
func (obj *_LgCustomerChannelMgr) GetBatchFromDataPushAPI(dataPushAPIs []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`data_push_api` IN (?)", dataPushAPIs).Find(&results).Error

	return
}

// GetFromProportion 通过proportion获取内容 抛比
func (obj *_LgCustomerChannelMgr) GetFromProportion(proportion int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`proportion` = ?", proportion).Find(&results).Error

	return
}

// GetBatchFromProportion 批量查找 抛比
func (obj *_LgCustomerChannelMgr) GetBatchFromProportion(proportions []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`proportion` IN (?)", proportions).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgCustomerChannelMgr) GetFromCreateTime(createTime time.Time) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgCustomerChannelMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgCustomerChannelMgr) GetFromCreateUser(createUser int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgCustomerChannelMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgCustomerChannelMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgCustomerChannelMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgCustomerChannelMgr) GetFromUpdateUser(updateUser int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgCustomerChannelMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgCustomerChannelMgr) GetFromVersion(version int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgCustomerChannelMgr) GetBatchFromVersion(versions []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgCustomerChannelMgr) GetFromDeleted(deleted int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgCustomerChannelMgr) GetBatchFromDeleted(deleteds []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态，0:无效，1:有效
func (obj *_LgCustomerChannelMgr) GetFromStatus(status int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态，0:无效，1:有效
func (obj *_LgCustomerChannelMgr) GetBatchFromStatus(statuss []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromRequiredPlatformNo 通过required_platform_no获取内容 是否必填平台订单号，0-不必填，1-必填
func (obj *_LgCustomerChannelMgr) GetFromRequiredPlatformNo(requiredPlatformNo []uint8) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`required_platform_no` = ?", requiredPlatformNo).Find(&results).Error

	return
}

// GetBatchFromRequiredPlatformNo 批量查找 是否必填平台订单号，0-不必填，1-必填
func (obj *_LgCustomerChannelMgr) GetBatchFromRequiredPlatformNo(requiredPlatformNos [][]uint8) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`required_platform_no` IN (?)", requiredPlatformNos).Find(&results).Error

	return
}

// GetFromChannelHaul 通过channel_haul获取内容 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_LgCustomerChannelMgr) GetFromChannelHaul(channelHaul string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`channel_haul` = ?", channelHaul).Find(&results).Error

	return
}

// GetBatchFromChannelHaul 批量查找 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
func (obj *_LgCustomerChannelMgr) GetBatchFromChannelHaul(channelHauls []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`channel_haul` IN (?)", channelHauls).Find(&results).Error

	return
}

// GetFromKingdeeID 通过kingdee_id获取内容 金蝶客户内码
func (obj *_LgCustomerChannelMgr) GetFromKingdeeID(kingdeeID string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}

// GetBatchFromKingdeeID 批量查找 金蝶客户内码
func (obj *_LgCustomerChannelMgr) GetBatchFromKingdeeID(kingdeeIDs []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`kingdee_id` IN (?)", kingdeeIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelTypeID 通过customer_channel_type_id获取内容 客户渠道类型id
func (obj *_LgCustomerChannelMgr) GetFromCustomerChannelTypeID(customerChannelTypeID string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`customer_channel_type_id` = ?", customerChannelTypeID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelTypeID 批量查找 客户渠道类型id
func (obj *_LgCustomerChannelMgr) GetBatchFromCustomerChannelTypeID(customerChannelTypeIDs []string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`customer_channel_type_id` IN (?)", customerChannelTypeIDs).Find(&results).Error

	return
}

// GetFromChannelCountryAffiliationID 通过channel_country_affiliation_id获取内容 渠道所属国家id,lg_area，对应lg_area表的id（注意：只有level 为0 的才是国家）
func (obj *_LgCustomerChannelMgr) GetFromChannelCountryAffiliationID(channelCountryAffiliationID int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`channel_country_affiliation_id` = ?", channelCountryAffiliationID).Find(&results).Error

	return
}

// GetBatchFromChannelCountryAffiliationID 批量查找 渠道所属国家id,lg_area，对应lg_area表的id（注意：只有level 为0 的才是国家）
func (obj *_LgCustomerChannelMgr) GetBatchFromChannelCountryAffiliationID(channelCountryAffiliationIDs []int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`channel_country_affiliation_id` IN (?)", channelCountryAffiliationIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgCustomerChannelMgr) FetchByPrimaryKey(id int) (result LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgCustomerChannelChannelNameUIndex primary or index 获取唯一内容
func (obj *_LgCustomerChannelMgr) FetchUniqueByLgCustomerChannelChannelNameUIndex(customerChannelName string) (result LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`customer_channel_name` = ?", customerChannelName).First(&result).Error

	return
}

// FetchIndexByIndexChannelHaul  获取多个内容
func (obj *_LgCustomerChannelMgr) FetchIndexByIndexChannelHaul(channelHaul string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`channel_haul` = ?", channelHaul).Find(&results).Error

	return
}

// FetchIndexByIndexKingdeeID  获取多个内容
func (obj *_LgCustomerChannelMgr) FetchIndexByIndexKingdeeID(kingdeeID string) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}

// FetchIndexByLgCustomerChannelChannelCountryAffiliationIDIndex  获取多个内容
func (obj *_LgCustomerChannelMgr) FetchIndexByLgCustomerChannelChannelCountryAffiliationIDIndex(channelCountryAffiliationID int) (results []*LgCustomerChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannel{}).Where("`channel_country_affiliation_id` = ?", channelCountryAffiliationID).Find(&results).Error

	return
}
