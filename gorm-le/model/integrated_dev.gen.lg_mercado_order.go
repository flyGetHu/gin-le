package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgMercadoOrderMgr struct {
	*_BaseMgr
}

// LgMercadoOrderMgr open func
func LgMercadoOrderMgr(db *gorm.DB) *_LgMercadoOrderMgr {
	if db == nil {
		panic(fmt.Errorf("LgMercadoOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgMercadoOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_mercado_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgMercadoOrderMgr) GetTableName() string {
	return "lg_mercado_order"
}

// Reset 重置gorm会话
func (obj *_LgMercadoOrderMgr) Reset() *_LgMercadoOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgMercadoOrderMgr) Get() (result LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgMercadoOrderMgr) Gets() (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgMercadoOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgMercadoOrderMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithChannelCode channel_code获取 渠道代码
func (obj *_LgMercadoOrderMgr) WithChannelCode(channelCode string) Option {
	return optionFunc(func(o *options) { o.query["channel_code"] = channelCode })
}

// WithMercadoID mercado_id获取 美卡多id
func (obj *_LgMercadoOrderMgr) WithMercadoID(mercadoID string) Option {
	return optionFunc(func(o *options) { o.query["mercado_id"] = mercadoID })
}

// WithOrderNum order_num获取 订单号
func (obj *_LgMercadoOrderMgr) WithOrderNum(orderNum string) Option {
	return optionFunc(func(o *options) { o.query["order_num"] = orderNum })
}

// WithAccount account获取 下单账户
func (obj *_LgMercadoOrderMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// WithAuthorizationCode authorization_code获取 账户授权码
func (obj *_LgMercadoOrderMgr) WithAuthorizationCode(authorizationCode string) Option {
	return optionFunc(func(o *options) { o.query["authorization_code"] = authorizationCode })
}

// WithOrderStatus order_status获取 订单状态
func (obj *_LgMercadoOrderMgr) WithOrderStatus(orderStatus int) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithPostOfficeStatus post_office_status获取 取号状态
func (obj *_LgMercadoOrderMgr) WithPostOfficeStatus(postOfficeStatus int) Option {
	return optionFunc(func(o *options) { o.query["post_office_status"] = postOfficeStatus })
}

// WithTrackStatus track_status获取 轨迹状态
func (obj *_LgMercadoOrderMgr) WithTrackStatus(trackStatus int) Option {
	return optionFunc(func(o *options) { o.query["track_status"] = trackStatus })
}

// WithTrackNumber track_number获取 跟踪号
func (obj *_LgMercadoOrderMgr) WithTrackNumber(trackNumber string) Option {
	return optionFunc(func(o *options) { o.query["track_number"] = trackNumber })
}

// WithReceiveName receive_name获取 收件人名称
func (obj *_LgMercadoOrderMgr) WithReceiveName(receiveName string) Option {
	return optionFunc(func(o *options) { o.query["receive_name"] = receiveName })
}

// WithReceiveTax receive_tax获取 收件人税号
func (obj *_LgMercadoOrderMgr) WithReceiveTax(receiveTax string) Option {
	return optionFunc(func(o *options) { o.query["receive_tax"] = receiveTax })
}

// WithReceivePhone receive_phone获取 收件人电话
func (obj *_LgMercadoOrderMgr) WithReceivePhone(receivePhone string) Option {
	return optionFunc(func(o *options) { o.query["receive_phone"] = receivePhone })
}

// WithReceiveZipCode receive_zip_code获取 收件人邮编
func (obj *_LgMercadoOrderMgr) WithReceiveZipCode(receiveZipCode string) Option {
	return optionFunc(func(o *options) { o.query["receive_zip_code"] = receiveZipCode })
}

// WithReceiveCountry receive_country获取 收件人国家
func (obj *_LgMercadoOrderMgr) WithReceiveCountry(receiveCountry string) Option {
	return optionFunc(func(o *options) { o.query["receive_country"] = receiveCountry })
}

// WithReceiveProvince receive_province获取 收件人省份
func (obj *_LgMercadoOrderMgr) WithReceiveProvince(receiveProvince string) Option {
	return optionFunc(func(o *options) { o.query["receive_province"] = receiveProvince })
}

// WithReceiveCity receive_city获取 收件人城市
func (obj *_LgMercadoOrderMgr) WithReceiveCity(receiveCity string) Option {
	return optionFunc(func(o *options) { o.query["receive_city"] = receiveCity })
}

// WithReceiveAddress receive_address获取 收件人地址
func (obj *_LgMercadoOrderMgr) WithReceiveAddress(receiveAddress string) Option {
	return optionFunc(func(o *options) { o.query["receive_address"] = receiveAddress })
}

// WithDeliveryPreference delivery_preference获取 地址类型
func (obj *_LgMercadoOrderMgr) WithDeliveryPreference(deliveryPreference string) Option {
	return optionFunc(func(o *options) { o.query["delivery_preference"] = deliveryPreference })
}

// WithSenderID sender_id获取 发件人id
func (obj *_LgMercadoOrderMgr) WithSenderID(senderID int) Option {
	return optionFunc(func(o *options) { o.query["sender_id"] = senderID })
}

// WithSenderName sender_name获取 发件人名称
func (obj *_LgMercadoOrderMgr) WithSenderName(senderName string) Option {
	return optionFunc(func(o *options) { o.query["sender_name"] = senderName })
}

// WithSenderPhone sender_phone获取 发件人电话
func (obj *_LgMercadoOrderMgr) WithSenderPhone(senderPhone string) Option {
	return optionFunc(func(o *options) { o.query["sender_phone"] = senderPhone })
}

// WithSenderTax sender_tax获取 发件人税号
func (obj *_LgMercadoOrderMgr) WithSenderTax(senderTax string) Option {
	return optionFunc(func(o *options) { o.query["sender_tax"] = senderTax })
}

// WithSenderAddress sender_address获取 发件人地址
func (obj *_LgMercadoOrderMgr) WithSenderAddress(senderAddress string) Option {
	return optionFunc(func(o *options) { o.query["sender_address"] = senderAddress })
}

// WithItemInfo item_info获取 商品信息
func (obj *_LgMercadoOrderMgr) WithItemInfo(itemInfo string) Option {
	return optionFunc(func(o *options) { o.query["item_info"] = itemInfo })
}

// WithQuantity quantity获取 申报数量
func (obj *_LgMercadoOrderMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithDeclaredValue declared_value获取 申报金额
func (obj *_LgMercadoOrderMgr) WithDeclaredValue(declaredValue float64) Option {
	return optionFunc(func(o *options) { o.query["declared_value"] = declaredValue })
}

// WithDeclaredWeight declared_weight获取 申报重量
func (obj *_LgMercadoOrderMgr) WithDeclaredWeight(declaredWeight float64) Option {
	return optionFunc(func(o *options) { o.query["declared_weight"] = declaredWeight })
}

// WithRequestParams request_params获取 请求参数
func (obj *_LgMercadoOrderMgr) WithRequestParams(requestParams string) Option {
	return optionFunc(func(o *options) { o.query["request_params"] = requestParams })
}

// WithResponseParams response_params获取 响应参数
func (obj *_LgMercadoOrderMgr) WithResponseParams(responseParams string) Option {
	return optionFunc(func(o *options) { o.query["response_params"] = responseParams })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgMercadoOrderMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgMercadoOrderMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgMercadoOrderMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithVersion version获取 乐观锁
func (obj *_LgMercadoOrderMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_LgMercadoOrderMgr) GetByOption(opts ...Option) (result LgMercadoOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgMercadoOrderMgr) GetByOptions(opts ...Option) (results []*LgMercadoOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgMercadoOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgMercadoOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where(options.query)
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
func (obj *_LgMercadoOrderMgr) GetFromID(id int) (result LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgMercadoOrderMgr) GetBatchFromID(ids []int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromChannelCode 通过channel_code获取内容 渠道代码
func (obj *_LgMercadoOrderMgr) GetFromChannelCode(channelCode string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`channel_code` = ?", channelCode).Find(&results).Error

	return
}

// GetBatchFromChannelCode 批量查找 渠道代码
func (obj *_LgMercadoOrderMgr) GetBatchFromChannelCode(channelCodes []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`channel_code` IN (?)", channelCodes).Find(&results).Error

	return
}

// GetFromMercadoID 通过mercado_id获取内容 美卡多id
func (obj *_LgMercadoOrderMgr) GetFromMercadoID(mercadoID string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`mercado_id` = ?", mercadoID).Find(&results).Error

	return
}

// GetBatchFromMercadoID 批量查找 美卡多id
func (obj *_LgMercadoOrderMgr) GetBatchFromMercadoID(mercadoIDs []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`mercado_id` IN (?)", mercadoIDs).Find(&results).Error

	return
}

// GetFromOrderNum 通过order_num获取内容 订单号
func (obj *_LgMercadoOrderMgr) GetFromOrderNum(orderNum string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`order_num` = ?", orderNum).Find(&results).Error

	return
}

// GetBatchFromOrderNum 批量查找 订单号
func (obj *_LgMercadoOrderMgr) GetBatchFromOrderNum(orderNums []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`order_num` IN (?)", orderNums).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 下单账户
func (obj *_LgMercadoOrderMgr) GetFromAccount(account string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找 下单账户
func (obj *_LgMercadoOrderMgr) GetBatchFromAccount(accounts []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromAuthorizationCode 通过authorization_code获取内容 账户授权码
func (obj *_LgMercadoOrderMgr) GetFromAuthorizationCode(authorizationCode string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`authorization_code` = ?", authorizationCode).Find(&results).Error

	return
}

// GetBatchFromAuthorizationCode 批量查找 账户授权码
func (obj *_LgMercadoOrderMgr) GetBatchFromAuthorizationCode(authorizationCodes []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`authorization_code` IN (?)", authorizationCodes).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容 订单状态
func (obj *_LgMercadoOrderMgr) GetFromOrderStatus(orderStatus int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`order_status` = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量查找 订单状态
func (obj *_LgMercadoOrderMgr) GetBatchFromOrderStatus(orderStatuss []int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`order_status` IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromPostOfficeStatus 通过post_office_status获取内容 取号状态
func (obj *_LgMercadoOrderMgr) GetFromPostOfficeStatus(postOfficeStatus int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`post_office_status` = ?", postOfficeStatus).Find(&results).Error

	return
}

// GetBatchFromPostOfficeStatus 批量查找 取号状态
func (obj *_LgMercadoOrderMgr) GetBatchFromPostOfficeStatus(postOfficeStatuss []int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`post_office_status` IN (?)", postOfficeStatuss).Find(&results).Error

	return
}

// GetFromTrackStatus 通过track_status获取内容 轨迹状态
func (obj *_LgMercadoOrderMgr) GetFromTrackStatus(trackStatus int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`track_status` = ?", trackStatus).Find(&results).Error

	return
}

// GetBatchFromTrackStatus 批量查找 轨迹状态
func (obj *_LgMercadoOrderMgr) GetBatchFromTrackStatus(trackStatuss []int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`track_status` IN (?)", trackStatuss).Find(&results).Error

	return
}

// GetFromTrackNumber 通过track_number获取内容 跟踪号
func (obj *_LgMercadoOrderMgr) GetFromTrackNumber(trackNumber string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`track_number` = ?", trackNumber).Find(&results).Error

	return
}

// GetBatchFromTrackNumber 批量查找 跟踪号
func (obj *_LgMercadoOrderMgr) GetBatchFromTrackNumber(trackNumbers []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`track_number` IN (?)", trackNumbers).Find(&results).Error

	return
}

// GetFromReceiveName 通过receive_name获取内容 收件人名称
func (obj *_LgMercadoOrderMgr) GetFromReceiveName(receiveName string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_name` = ?", receiveName).Find(&results).Error

	return
}

// GetBatchFromReceiveName 批量查找 收件人名称
func (obj *_LgMercadoOrderMgr) GetBatchFromReceiveName(receiveNames []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_name` IN (?)", receiveNames).Find(&results).Error

	return
}

// GetFromReceiveTax 通过receive_tax获取内容 收件人税号
func (obj *_LgMercadoOrderMgr) GetFromReceiveTax(receiveTax string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_tax` = ?", receiveTax).Find(&results).Error

	return
}

// GetBatchFromReceiveTax 批量查找 收件人税号
func (obj *_LgMercadoOrderMgr) GetBatchFromReceiveTax(receiveTaxs []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_tax` IN (?)", receiveTaxs).Find(&results).Error

	return
}

// GetFromReceivePhone 通过receive_phone获取内容 收件人电话
func (obj *_LgMercadoOrderMgr) GetFromReceivePhone(receivePhone string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_phone` = ?", receivePhone).Find(&results).Error

	return
}

// GetBatchFromReceivePhone 批量查找 收件人电话
func (obj *_LgMercadoOrderMgr) GetBatchFromReceivePhone(receivePhones []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_phone` IN (?)", receivePhones).Find(&results).Error

	return
}

// GetFromReceiveZipCode 通过receive_zip_code获取内容 收件人邮编
func (obj *_LgMercadoOrderMgr) GetFromReceiveZipCode(receiveZipCode string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_zip_code` = ?", receiveZipCode).Find(&results).Error

	return
}

// GetBatchFromReceiveZipCode 批量查找 收件人邮编
func (obj *_LgMercadoOrderMgr) GetBatchFromReceiveZipCode(receiveZipCodes []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_zip_code` IN (?)", receiveZipCodes).Find(&results).Error

	return
}

// GetFromReceiveCountry 通过receive_country获取内容 收件人国家
func (obj *_LgMercadoOrderMgr) GetFromReceiveCountry(receiveCountry string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// GetBatchFromReceiveCountry 批量查找 收件人国家
func (obj *_LgMercadoOrderMgr) GetBatchFromReceiveCountry(receiveCountrys []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_country` IN (?)", receiveCountrys).Find(&results).Error

	return
}

// GetFromReceiveProvince 通过receive_province获取内容 收件人省份
func (obj *_LgMercadoOrderMgr) GetFromReceiveProvince(receiveProvince string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_province` = ?", receiveProvince).Find(&results).Error

	return
}

// GetBatchFromReceiveProvince 批量查找 收件人省份
func (obj *_LgMercadoOrderMgr) GetBatchFromReceiveProvince(receiveProvinces []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_province` IN (?)", receiveProvinces).Find(&results).Error

	return
}

// GetFromReceiveCity 通过receive_city获取内容 收件人城市
func (obj *_LgMercadoOrderMgr) GetFromReceiveCity(receiveCity string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_city` = ?", receiveCity).Find(&results).Error

	return
}

// GetBatchFromReceiveCity 批量查找 收件人城市
func (obj *_LgMercadoOrderMgr) GetBatchFromReceiveCity(receiveCitys []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_city` IN (?)", receiveCitys).Find(&results).Error

	return
}

// GetFromReceiveAddress 通过receive_address获取内容 收件人地址
func (obj *_LgMercadoOrderMgr) GetFromReceiveAddress(receiveAddress string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_address` = ?", receiveAddress).Find(&results).Error

	return
}

// GetBatchFromReceiveAddress 批量查找 收件人地址
func (obj *_LgMercadoOrderMgr) GetBatchFromReceiveAddress(receiveAddresss []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`receive_address` IN (?)", receiveAddresss).Find(&results).Error

	return
}

// GetFromDeliveryPreference 通过delivery_preference获取内容 地址类型
func (obj *_LgMercadoOrderMgr) GetFromDeliveryPreference(deliveryPreference string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`delivery_preference` = ?", deliveryPreference).Find(&results).Error

	return
}

// GetBatchFromDeliveryPreference 批量查找 地址类型
func (obj *_LgMercadoOrderMgr) GetBatchFromDeliveryPreference(deliveryPreferences []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`delivery_preference` IN (?)", deliveryPreferences).Find(&results).Error

	return
}

// GetFromSenderID 通过sender_id获取内容 发件人id
func (obj *_LgMercadoOrderMgr) GetFromSenderID(senderID int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_id` = ?", senderID).Find(&results).Error

	return
}

// GetBatchFromSenderID 批量查找 发件人id
func (obj *_LgMercadoOrderMgr) GetBatchFromSenderID(senderIDs []int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_id` IN (?)", senderIDs).Find(&results).Error

	return
}

// GetFromSenderName 通过sender_name获取内容 发件人名称
func (obj *_LgMercadoOrderMgr) GetFromSenderName(senderName string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_name` = ?", senderName).Find(&results).Error

	return
}

// GetBatchFromSenderName 批量查找 发件人名称
func (obj *_LgMercadoOrderMgr) GetBatchFromSenderName(senderNames []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_name` IN (?)", senderNames).Find(&results).Error

	return
}

// GetFromSenderPhone 通过sender_phone获取内容 发件人电话
func (obj *_LgMercadoOrderMgr) GetFromSenderPhone(senderPhone string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_phone` = ?", senderPhone).Find(&results).Error

	return
}

// GetBatchFromSenderPhone 批量查找 发件人电话
func (obj *_LgMercadoOrderMgr) GetBatchFromSenderPhone(senderPhones []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_phone` IN (?)", senderPhones).Find(&results).Error

	return
}

// GetFromSenderTax 通过sender_tax获取内容 发件人税号
func (obj *_LgMercadoOrderMgr) GetFromSenderTax(senderTax string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_tax` = ?", senderTax).Find(&results).Error

	return
}

// GetBatchFromSenderTax 批量查找 发件人税号
func (obj *_LgMercadoOrderMgr) GetBatchFromSenderTax(senderTaxs []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_tax` IN (?)", senderTaxs).Find(&results).Error

	return
}

// GetFromSenderAddress 通过sender_address获取内容 发件人地址
func (obj *_LgMercadoOrderMgr) GetFromSenderAddress(senderAddress string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_address` = ?", senderAddress).Find(&results).Error

	return
}

// GetBatchFromSenderAddress 批量查找 发件人地址
func (obj *_LgMercadoOrderMgr) GetBatchFromSenderAddress(senderAddresss []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`sender_address` IN (?)", senderAddresss).Find(&results).Error

	return
}

// GetFromItemInfo 通过item_info获取内容 商品信息
func (obj *_LgMercadoOrderMgr) GetFromItemInfo(itemInfo string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`item_info` = ?", itemInfo).Find(&results).Error

	return
}

// GetBatchFromItemInfo 批量查找 商品信息
func (obj *_LgMercadoOrderMgr) GetBatchFromItemInfo(itemInfos []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`item_info` IN (?)", itemInfos).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 申报数量
func (obj *_LgMercadoOrderMgr) GetFromQuantity(quantity int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 申报数量
func (obj *_LgMercadoOrderMgr) GetBatchFromQuantity(quantitys []int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromDeclaredValue 通过declared_value获取内容 申报金额
func (obj *_LgMercadoOrderMgr) GetFromDeclaredValue(declaredValue float64) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`declared_value` = ?", declaredValue).Find(&results).Error

	return
}

// GetBatchFromDeclaredValue 批量查找 申报金额
func (obj *_LgMercadoOrderMgr) GetBatchFromDeclaredValue(declaredValues []float64) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`declared_value` IN (?)", declaredValues).Find(&results).Error

	return
}

// GetFromDeclaredWeight 通过declared_weight获取内容 申报重量
func (obj *_LgMercadoOrderMgr) GetFromDeclaredWeight(declaredWeight float64) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`declared_weight` = ?", declaredWeight).Find(&results).Error

	return
}

// GetBatchFromDeclaredWeight 批量查找 申报重量
func (obj *_LgMercadoOrderMgr) GetBatchFromDeclaredWeight(declaredWeights []float64) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`declared_weight` IN (?)", declaredWeights).Find(&results).Error

	return
}

// GetFromRequestParams 通过request_params获取内容 请求参数
func (obj *_LgMercadoOrderMgr) GetFromRequestParams(requestParams string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`request_params` = ?", requestParams).Find(&results).Error

	return
}

// GetBatchFromRequestParams 批量查找 请求参数
func (obj *_LgMercadoOrderMgr) GetBatchFromRequestParams(requestParamss []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`request_params` IN (?)", requestParamss).Find(&results).Error

	return
}

// GetFromResponseParams 通过response_params获取内容 响应参数
func (obj *_LgMercadoOrderMgr) GetFromResponseParams(responseParams string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`response_params` = ?", responseParams).Find(&results).Error

	return
}

// GetBatchFromResponseParams 批量查找 响应参数
func (obj *_LgMercadoOrderMgr) GetBatchFromResponseParams(responseParamss []string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`response_params` IN (?)", responseParamss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgMercadoOrderMgr) GetFromCreateTime(createTime time.Time) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgMercadoOrderMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgMercadoOrderMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgMercadoOrderMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgMercadoOrderMgr) GetFromDeleted(deleted int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgMercadoOrderMgr) GetBatchFromDeleted(deleteds []int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgMercadoOrderMgr) GetFromVersion(version int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgMercadoOrderMgr) GetBatchFromVersion(versions []int) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgMercadoOrderMgr) FetchByPrimaryKey(id int) (result LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByTableNameIDUIndex primary or index 获取唯一内容
func (obj *_LgMercadoOrderMgr) FetchUniqueByTableNameIDUIndex(id int) (result LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByTableNameMercadoOrderIDIndex  获取多个内容
func (obj *_LgMercadoOrderMgr) FetchIndexByTableNameMercadoOrderIDIndex(mercadoID string) (results []*LgMercadoOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgMercadoOrder{}).Where("`mercado_id` = ?", mercadoID).Find(&results).Error

	return
}
