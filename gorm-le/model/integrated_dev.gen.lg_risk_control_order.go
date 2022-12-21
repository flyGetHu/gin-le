package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgRiskControlOrderMgr struct {
	*_BaseMgr
}

// LgRiskControlOrderMgr open func
func LgRiskControlOrderMgr(db *gorm.DB) *_LgRiskControlOrderMgr {
	if db == nil {
		panic(fmt.Errorf("LgRiskControlOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgRiskControlOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_risk_control_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgRiskControlOrderMgr) GetTableName() string {
	return "lg_risk_control_order"
}

// Reset 重置gorm会话
func (obj *_LgRiskControlOrderMgr) Reset() *_LgRiskControlOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgRiskControlOrderMgr) Get() (result LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgRiskControlOrderMgr) Gets() (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgRiskControlOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgRiskControlOrderMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_LgRiskControlOrderMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_LgRiskControlOrderMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_LgRiskControlOrderMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLabelURL label_url获取 面单地址
func (obj *_LgRiskControlOrderMgr) WithLabelURL(labelURL string) Option {
	return optionFunc(func(o *options) { o.query["label_url"] = labelURL })
}

// WithOnlineNumber online_number获取 上网单号(17单号)
func (obj *_LgRiskControlOrderMgr) WithOnlineNumber(onlineNumber string) Option {
	return optionFunc(func(o *options) { o.query["online_number"] = onlineNumber })
}

// WithBagNumber bag_number获取 袋号
func (obj *_LgRiskControlOrderMgr) WithBagNumber(bagNumber string) Option {
	return optionFunc(func(o *options) { o.query["bag_number"] = bagNumber })
}

// WithAirBillNumber air_bill_number获取 航空提单号
func (obj *_LgRiskControlOrderMgr) WithAirBillNumber(airBillNumber string) Option {
	return optionFunc(func(o *options) { o.query["air_bill_number"] = airBillNumber })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgRiskControlOrderMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_LgRiskControlOrderMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithReceiveCountry receive_country获取 收件人国家
func (obj *_LgRiskControlOrderMgr) WithReceiveCountry(receiveCountry string) Option {
	return optionFunc(func(o *options) { o.query["receive_country"] = receiveCountry })
}

// WithReceiveName receive_name获取 收件人姓名
func (obj *_LgRiskControlOrderMgr) WithReceiveName(receiveName string) Option {
	return optionFunc(func(o *options) { o.query["receive_name"] = receiveName })
}

// WithReceivePhone receive_phone获取 收件人电话
func (obj *_LgRiskControlOrderMgr) WithReceivePhone(receivePhone string) Option {
	return optionFunc(func(o *options) { o.query["receive_phone"] = receivePhone })
}

// WithReceiveMobile receive_mobile获取 收件人手机
func (obj *_LgRiskControlOrderMgr) WithReceiveMobile(receiveMobile string) Option {
	return optionFunc(func(o *options) { o.query["receive_mobile"] = receiveMobile })
}

// WithReceiveTax receive_tax获取 收件人税号
func (obj *_LgRiskControlOrderMgr) WithReceiveTax(receiveTax string) Option {
	return optionFunc(func(o *options) { o.query["receive_tax"] = receiveTax })
}

// WithReceivePassport receive_passport获取 收件人护照号
func (obj *_LgRiskControlOrderMgr) WithReceivePassport(receivePassport string) Option {
	return optionFunc(func(o *options) { o.query["receive_passport"] = receivePassport })
}

// WithReceiveMail receive_mail获取 收件人邮箱
func (obj *_LgRiskControlOrderMgr) WithReceiveMail(receiveMail string) Option {
	return optionFunc(func(o *options) { o.query["receive_mail"] = receiveMail })
}

// WithReceiveCompany receive_company获取 收件人公司
func (obj *_LgRiskControlOrderMgr) WithReceiveCompany(receiveCompany string) Option {
	return optionFunc(func(o *options) { o.query["receive_company"] = receiveCompany })
}

// WithReceiveZipcode receive_zipcode获取 收件人邮编
func (obj *_LgRiskControlOrderMgr) WithReceiveZipcode(receiveZipcode string) Option {
	return optionFunc(func(o *options) { o.query["receive_zipcode"] = receiveZipcode })
}

// WithReceiveProvince receive_province获取 收件人省份
func (obj *_LgRiskControlOrderMgr) WithReceiveProvince(receiveProvince string) Option {
	return optionFunc(func(o *options) { o.query["receive_province"] = receiveProvince })
}

// WithReceiveCity receive_city获取 收件人城市
func (obj *_LgRiskControlOrderMgr) WithReceiveCity(receiveCity string) Option {
	return optionFunc(func(o *options) { o.query["receive_city"] = receiveCity })
}

// WithReceiveArea receive_area获取 收件人区
func (obj *_LgRiskControlOrderMgr) WithReceiveArea(receiveArea string) Option {
	return optionFunc(func(o *options) { o.query["receive_area"] = receiveArea })
}

// WithReceiveStreet receive_street获取 收件人街道
func (obj *_LgRiskControlOrderMgr) WithReceiveStreet(receiveStreet string) Option {
	return optionFunc(func(o *options) { o.query["receive_street"] = receiveStreet })
}

// WithReceiveHouseNumber receive_house_number获取 收件人门牌号
func (obj *_LgRiskControlOrderMgr) WithReceiveHouseNumber(receiveHouseNumber string) Option {
	return optionFunc(func(o *options) { o.query["receive_house_number"] = receiveHouseNumber })
}

// WithReceiveAddress1 receive_address1获取 收件人地址1
func (obj *_LgRiskControlOrderMgr) WithReceiveAddress1(receiveAddress1 string) Option {
	return optionFunc(func(o *options) { o.query["receive_address1"] = receiveAddress1 })
}

// WithReceiveAddress2 receive_address2获取 收件人地址2
func (obj *_LgRiskControlOrderMgr) WithReceiveAddress2(receiveAddress2 string) Option {
	return optionFunc(func(o *options) { o.query["receive_address2"] = receiveAddress2 })
}

// WithSenderCountry sender_country获取 发件人国家
func (obj *_LgRiskControlOrderMgr) WithSenderCountry(senderCountry string) Option {
	return optionFunc(func(o *options) { o.query["sender_country"] = senderCountry })
}

// WithSenderName sender_name获取 发件人姓名
func (obj *_LgRiskControlOrderMgr) WithSenderName(senderName string) Option {
	return optionFunc(func(o *options) { o.query["sender_name"] = senderName })
}

// WithSenderPhone sender_phone获取 发件人电话
func (obj *_LgRiskControlOrderMgr) WithSenderPhone(senderPhone string) Option {
	return optionFunc(func(o *options) { o.query["sender_phone"] = senderPhone })
}

// WithSenderMobile sender_mobile获取 发件人手机
func (obj *_LgRiskControlOrderMgr) WithSenderMobile(senderMobile string) Option {
	return optionFunc(func(o *options) { o.query["sender_mobile"] = senderMobile })
}

// WithSenderTax sender_tax获取 发件人税号
func (obj *_LgRiskControlOrderMgr) WithSenderTax(senderTax string) Option {
	return optionFunc(func(o *options) { o.query["sender_tax"] = senderTax })
}

// WithSenderPassport sender_passport获取 发件人护照号
func (obj *_LgRiskControlOrderMgr) WithSenderPassport(senderPassport string) Option {
	return optionFunc(func(o *options) { o.query["sender_passport"] = senderPassport })
}

// WithSenderMail sender_mail获取 发件人邮箱
func (obj *_LgRiskControlOrderMgr) WithSenderMail(senderMail string) Option {
	return optionFunc(func(o *options) { o.query["sender_mail"] = senderMail })
}

// WithSenderCompany sender_company获取 发件人公司
func (obj *_LgRiskControlOrderMgr) WithSenderCompany(senderCompany string) Option {
	return optionFunc(func(o *options) { o.query["sender_company"] = senderCompany })
}

// WithSenderZipcode sender_zipcode获取 发件人邮编
func (obj *_LgRiskControlOrderMgr) WithSenderZipcode(senderZipcode string) Option {
	return optionFunc(func(o *options) { o.query["sender_zipcode"] = senderZipcode })
}

// WithSenderProvince sender_province获取 发件人省份
func (obj *_LgRiskControlOrderMgr) WithSenderProvince(senderProvince string) Option {
	return optionFunc(func(o *options) { o.query["sender_province"] = senderProvince })
}

// WithSenderCity sender_city获取 发件人城市
func (obj *_LgRiskControlOrderMgr) WithSenderCity(senderCity string) Option {
	return optionFunc(func(o *options) { o.query["sender_city"] = senderCity })
}

// WithSenderArea sender_area获取 发件人区
func (obj *_LgRiskControlOrderMgr) WithSenderArea(senderArea string) Option {
	return optionFunc(func(o *options) { o.query["sender_area"] = senderArea })
}

// WithSenderStreet sender_street获取 发件人街道
func (obj *_LgRiskControlOrderMgr) WithSenderStreet(senderStreet string) Option {
	return optionFunc(func(o *options) { o.query["sender_street"] = senderStreet })
}

// WithSenderHouseNumber sender_house_number获取 发件人门牌号
func (obj *_LgRiskControlOrderMgr) WithSenderHouseNumber(senderHouseNumber string) Option {
	return optionFunc(func(o *options) { o.query["sender_house_number"] = senderHouseNumber })
}

// WithSenderAddress sender_address获取 发件人地址
func (obj *_LgRiskControlOrderMgr) WithSenderAddress(senderAddress string) Option {
	return optionFunc(func(o *options) { o.query["sender_address"] = senderAddress })
}

// WithPackageType package_type获取 包裹类型
func (obj *_LgRiskControlOrderMgr) WithPackageType(packageType string) Option {
	return optionFunc(func(o *options) { o.query["package_type"] = packageType })
}

// WithHasBattery has_battery获取 是否带电
func (obj *_LgRiskControlOrderMgr) WithHasBattery(hasBattery int) Option {
	return optionFunc(func(o *options) { o.query["has_battery"] = hasBattery })
}

// WithBatteryType battery_type获取 电池类型
func (obj *_LgRiskControlOrderMgr) WithBatteryType(batteryType string) Option {
	return optionFunc(func(o *options) { o.query["battery_type"] = batteryType })
}

// WithHasBack has_back获取 是否退回
func (obj *_LgRiskControlOrderMgr) WithHasBack(hasBack int) Option {
	return optionFunc(func(o *options) { o.query["has_back"] = hasBack })
}

// WithRemarks remarks获取 备注信息
func (obj *_LgRiskControlOrderMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithDistributionInformation distribution_information获取 配货信息
func (obj *_LgRiskControlOrderMgr) WithDistributionInformation(distributionInformation string) Option {
	return optionFunc(func(o *options) { o.query["distribution_information"] = distributionInformation })
}

// WithInsurance insurance获取 保险
func (obj *_LgRiskControlOrderMgr) WithInsurance(insurance string) Option {
	return optionFunc(func(o *options) { o.query["insurance"] = insurance })
}

// WithSalesPlatform sales_platform获取 销售平台
func (obj *_LgRiskControlOrderMgr) WithSalesPlatform(salesPlatform string) Option {
	return optionFunc(func(o *options) { o.query["sales_platform"] = salesPlatform })
}

// WithWeight weight获取 预报重量
func (obj *_LgRiskControlOrderMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithQuantity quantity获取 包裹数量
func (obj *_LgRiskControlOrderMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithDeclaredValue declared_value获取 申报价值
func (obj *_LgRiskControlOrderMgr) WithDeclaredValue(declaredValue float64) Option {
	return optionFunc(func(o *options) { o.query["declared_value"] = declaredValue })
}

// WithStatus status获取 订单状态
func (obj *_LgRiskControlOrderMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgRiskControlOrderMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithWeighingWeight weighing_weight获取 称重重量
func (obj *_LgRiskControlOrderMgr) WithWeighingWeight(weighingWeight float64) Option {
	return optionFunc(func(o *options) { o.query["weighing_weight"] = weighingWeight })
}

// WithVolumeWeight volume_weight获取 材积重量
func (obj *_LgRiskControlOrderMgr) WithVolumeWeight(volumeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["volume_weight"] = volumeWeight })
}

// WithChargeWeight charge_weight获取 收费重量
func (obj *_LgRiskControlOrderMgr) WithChargeWeight(chargeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["charge_weight"] = chargeWeight })
}

// WithStandardFee standard_fee获取 标准费用
func (obj *_LgRiskControlOrderMgr) WithStandardFee(standardFee float64) Option {
	return optionFunc(func(o *options) { o.query["standard_fee"] = standardFee })
}

// WithCurrency currency获取 币种
func (obj *_LgRiskControlOrderMgr) WithCurrency(currency string) Option {
	return optionFunc(func(o *options) { o.query["currency"] = currency })
}

// WithBuyersFee buyers_fee获取 下家费用
func (obj *_LgRiskControlOrderMgr) WithBuyersFee(buyersFee float64) Option {
	return optionFunc(func(o *options) { o.query["buyers_fee"] = buyersFee })
}

// WithBuyersWeight buyers_weight获取 下家重量
func (obj *_LgRiskControlOrderMgr) WithBuyersWeight(buyersWeight float64) Option {
	return optionFunc(func(o *options) { o.query["buyers_weight"] = buyersWeight })
}

// WithOtherFee other_fee获取 其他费用
func (obj *_LgRiskControlOrderMgr) WithOtherFee(otherFee float64) Option {
	return optionFunc(func(o *options) { o.query["other_fee"] = otherFee })
}

// WithCustomerID customer_id获取 归属客户ID
func (obj *_LgRiskControlOrderMgr) WithCustomerID(customerID int64) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerName customer_name获取 归属客户名称
func (obj *_LgRiskControlOrderMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithPrepaymentVat prepayment_vat获取 "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
func (obj *_LgRiskControlOrderMgr) WithPrepaymentVat(prepaymentVat string) Option {
	return optionFunc(func(o *options) { o.query["prepayment_vat"] = prepaymentVat })
}

// WithVat vat获取 vat号(英国方向必填)
func (obj *_LgRiskControlOrderMgr) WithVat(vat string) Option {
	return optionFunc(func(o *options) { o.query["vat"] = vat })
}

// WithReceiveCertificateType receive_certificate_type获取 收件人证件类型
func (obj *_LgRiskControlOrderMgr) WithReceiveCertificateType(receiveCertificateType string) Option {
	return optionFunc(func(o *options) { o.query["receive_certificate_type"] = receiveCertificateType })
}

// WithReceiveCertificateCode receive_certificate_code获取 收件人证件号码
func (obj *_LgRiskControlOrderMgr) WithReceiveCertificateCode(receiveCertificateCode string) Option {
	return optionFunc(func(o *options) { o.query["receive_certificate_code"] = receiveCertificateCode })
}

// WithSenderCertificateType sender_certificate_type获取 发件人证件类型
func (obj *_LgRiskControlOrderMgr) WithSenderCertificateType(senderCertificateType string) Option {
	return optionFunc(func(o *options) { o.query["sender_certificate_type"] = senderCertificateType })
}

// WithSenderCertificateCode sender_certificate_code获取 发件人证件号码
func (obj *_LgRiskControlOrderMgr) WithSenderCertificateCode(senderCertificateCode string) Option {
	return optionFunc(func(o *options) { o.query["sender_certificate_code"] = senderCertificateCode })
}

// WithServiceChannelCode service_channel_code获取 api渠道代码
func (obj *_LgRiskControlOrderMgr) WithServiceChannelCode(serviceChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["service_channel_code"] = serviceChannelCode })
}

// WithDeliveryTerms delivery_terms获取 贸易条款- DDU DDP
func (obj *_LgRiskControlOrderMgr) WithDeliveryTerms(deliveryTerms string) Option {
	return optionFunc(func(o *options) { o.query["delivery_terms"] = deliveryTerms })
}

// WithReceiveCountryName receive_country_name获取 收件人国家中文
func (obj *_LgRiskControlOrderMgr) WithReceiveCountryName(receiveCountryName string) Option {
	return optionFunc(func(o *options) { o.query["receive_country_name"] = receiveCountryName })
}

// WithSenderCountryName sender_country_name获取 发件国家中文
func (obj *_LgRiskControlOrderMgr) WithSenderCountryName(senderCountryName string) Option {
	return optionFunc(func(o *options) { o.query["sender_country_name"] = senderCountryName })
}

// WithOrderType order_type获取 订单类型(0-预报订单，1-物流订单)
func (obj *_LgRiskControlOrderMgr) WithOrderType(orderType int) Option {
	return optionFunc(func(o *options) { o.query["order_type"] = orderType })
}

// WithReceiptTime receipt_time获取 入库时间
func (obj *_LgRiskControlOrderMgr) WithReceiptTime(receiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_time"] = receiptTime })
}

// WithReceiptUser receipt_user获取 入库人
func (obj *_LgRiskControlOrderMgr) WithReceiptUser(receiptUser int) Option {
	return optionFunc(func(o *options) { o.query["receipt_user"] = receiptUser })
}

// WithSendTime send_time获取 出库时间
func (obj *_LgRiskControlOrderMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithSendUser send_user获取 出库人
func (obj *_LgRiskControlOrderMgr) WithSendUser(sendUser int) Option {
	return optionFunc(func(o *options) { o.query["send_user"] = sendUser })
}

// WithPicOssLink pic_oss_link获取 包裹图片oss地址
func (obj *_LgRiskControlOrderMgr) WithPicOssLink(picOssLink string) Option {
	return optionFunc(func(o *options) { o.query["pic_oss_link"] = picOssLink })
}

// WithRiskMessage risk_message获取 风控信息
func (obj *_LgRiskControlOrderMgr) WithRiskMessage(riskMessage string) Option {
	return optionFunc(func(o *options) { o.query["risk_message"] = riskMessage })
}

// WithStatusRisk status_risk获取 风控订单状态:normal:正常,abnormal:异常 warn:警告
func (obj *_LgRiskControlOrderMgr) WithStatusRisk(statusRisk string) Option {
	return optionFunc(func(o *options) { o.query["status_risk"] = statusRisk })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgRiskControlOrderMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgRiskControlOrderMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgRiskControlOrderMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgRiskControlOrderMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgRiskControlOrderMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgRiskControlOrderMgr) GetByOption(opts ...Option) (result LgRiskControlOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgRiskControlOrderMgr) GetByOptions(opts ...Option) (results []*LgRiskControlOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgRiskControlOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgRiskControlOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where(options.query)
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

// GetFromID 通过id获取内容 ID
func (obj *_LgRiskControlOrderMgr) GetFromID(id int) (result LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgRiskControlOrderMgr) GetBatchFromID(ids []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_LgRiskControlOrderMgr) GetFromOrderNumber(orderNumber string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_LgRiskControlOrderMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_LgRiskControlOrderMgr) GetFromReferenceNumber(referenceNumber string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`reference_number` = ?", referenceNumber).Find(&results).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_LgRiskControlOrderMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_LgRiskControlOrderMgr) GetFromLogisticsNumber(logisticsNumber string) (result LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`logistics_number` = ?", logisticsNumber).First(&result).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_LgRiskControlOrderMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLabelURL 通过label_url获取内容 面单地址
func (obj *_LgRiskControlOrderMgr) GetFromLabelURL(labelURL string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`label_url` = ?", labelURL).Find(&results).Error

	return
}

// GetBatchFromLabelURL 批量查找 面单地址
func (obj *_LgRiskControlOrderMgr) GetBatchFromLabelURL(labelURLs []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`label_url` IN (?)", labelURLs).Find(&results).Error

	return
}

// GetFromOnlineNumber 通过online_number获取内容 上网单号(17单号)
func (obj *_LgRiskControlOrderMgr) GetFromOnlineNumber(onlineNumber string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`online_number` = ?", onlineNumber).Find(&results).Error

	return
}

// GetBatchFromOnlineNumber 批量查找 上网单号(17单号)
func (obj *_LgRiskControlOrderMgr) GetBatchFromOnlineNumber(onlineNumbers []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`online_number` IN (?)", onlineNumbers).Find(&results).Error

	return
}

// GetFromBagNumber 通过bag_number获取内容 袋号
func (obj *_LgRiskControlOrderMgr) GetFromBagNumber(bagNumber string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`bag_number` = ?", bagNumber).Find(&results).Error

	return
}

// GetBatchFromBagNumber 批量查找 袋号
func (obj *_LgRiskControlOrderMgr) GetBatchFromBagNumber(bagNumbers []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`bag_number` IN (?)", bagNumbers).Find(&results).Error

	return
}

// GetFromAirBillNumber 通过air_bill_number获取内容 航空提单号
func (obj *_LgRiskControlOrderMgr) GetFromAirBillNumber(airBillNumber string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`air_bill_number` = ?", airBillNumber).Find(&results).Error

	return
}

// GetBatchFromAirBillNumber 批量查找 航空提单号
func (obj *_LgRiskControlOrderMgr) GetBatchFromAirBillNumber(airBillNumbers []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`air_bill_number` IN (?)", airBillNumbers).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgRiskControlOrderMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgRiskControlOrderMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_LgRiskControlOrderMgr) GetFromCustomerChannelName(customerChannelName string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_LgRiskControlOrderMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromReceiveCountry 通过receive_country获取内容 收件人国家
func (obj *_LgRiskControlOrderMgr) GetFromReceiveCountry(receiveCountry string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// GetBatchFromReceiveCountry 批量查找 收件人国家
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveCountry(receiveCountrys []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_country` IN (?)", receiveCountrys).Find(&results).Error

	return
}

// GetFromReceiveName 通过receive_name获取内容 收件人姓名
func (obj *_LgRiskControlOrderMgr) GetFromReceiveName(receiveName string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_name` = ?", receiveName).Find(&results).Error

	return
}

// GetBatchFromReceiveName 批量查找 收件人姓名
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveName(receiveNames []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_name` IN (?)", receiveNames).Find(&results).Error

	return
}

// GetFromReceivePhone 通过receive_phone获取内容 收件人电话
func (obj *_LgRiskControlOrderMgr) GetFromReceivePhone(receivePhone string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_phone` = ?", receivePhone).Find(&results).Error

	return
}

// GetBatchFromReceivePhone 批量查找 收件人电话
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceivePhone(receivePhones []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_phone` IN (?)", receivePhones).Find(&results).Error

	return
}

// GetFromReceiveMobile 通过receive_mobile获取内容 收件人手机
func (obj *_LgRiskControlOrderMgr) GetFromReceiveMobile(receiveMobile string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_mobile` = ?", receiveMobile).Find(&results).Error

	return
}

// GetBatchFromReceiveMobile 批量查找 收件人手机
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveMobile(receiveMobiles []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_mobile` IN (?)", receiveMobiles).Find(&results).Error

	return
}

// GetFromReceiveTax 通过receive_tax获取内容 收件人税号
func (obj *_LgRiskControlOrderMgr) GetFromReceiveTax(receiveTax string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_tax` = ?", receiveTax).Find(&results).Error

	return
}

// GetBatchFromReceiveTax 批量查找 收件人税号
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveTax(receiveTaxs []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_tax` IN (?)", receiveTaxs).Find(&results).Error

	return
}

// GetFromReceivePassport 通过receive_passport获取内容 收件人护照号
func (obj *_LgRiskControlOrderMgr) GetFromReceivePassport(receivePassport string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_passport` = ?", receivePassport).Find(&results).Error

	return
}

// GetBatchFromReceivePassport 批量查找 收件人护照号
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceivePassport(receivePassports []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_passport` IN (?)", receivePassports).Find(&results).Error

	return
}

// GetFromReceiveMail 通过receive_mail获取内容 收件人邮箱
func (obj *_LgRiskControlOrderMgr) GetFromReceiveMail(receiveMail string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_mail` = ?", receiveMail).Find(&results).Error

	return
}

// GetBatchFromReceiveMail 批量查找 收件人邮箱
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveMail(receiveMails []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_mail` IN (?)", receiveMails).Find(&results).Error

	return
}

// GetFromReceiveCompany 通过receive_company获取内容 收件人公司
func (obj *_LgRiskControlOrderMgr) GetFromReceiveCompany(receiveCompany string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_company` = ?", receiveCompany).Find(&results).Error

	return
}

// GetBatchFromReceiveCompany 批量查找 收件人公司
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveCompany(receiveCompanys []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_company` IN (?)", receiveCompanys).Find(&results).Error

	return
}

// GetFromReceiveZipcode 通过receive_zipcode获取内容 收件人邮编
func (obj *_LgRiskControlOrderMgr) GetFromReceiveZipcode(receiveZipcode string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_zipcode` = ?", receiveZipcode).Find(&results).Error

	return
}

// GetBatchFromReceiveZipcode 批量查找 收件人邮编
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveZipcode(receiveZipcodes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_zipcode` IN (?)", receiveZipcodes).Find(&results).Error

	return
}

// GetFromReceiveProvince 通过receive_province获取内容 收件人省份
func (obj *_LgRiskControlOrderMgr) GetFromReceiveProvince(receiveProvince string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_province` = ?", receiveProvince).Find(&results).Error

	return
}

// GetBatchFromReceiveProvince 批量查找 收件人省份
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveProvince(receiveProvinces []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_province` IN (?)", receiveProvinces).Find(&results).Error

	return
}

// GetFromReceiveCity 通过receive_city获取内容 收件人城市
func (obj *_LgRiskControlOrderMgr) GetFromReceiveCity(receiveCity string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_city` = ?", receiveCity).Find(&results).Error

	return
}

// GetBatchFromReceiveCity 批量查找 收件人城市
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveCity(receiveCitys []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_city` IN (?)", receiveCitys).Find(&results).Error

	return
}

// GetFromReceiveArea 通过receive_area获取内容 收件人区
func (obj *_LgRiskControlOrderMgr) GetFromReceiveArea(receiveArea string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_area` = ?", receiveArea).Find(&results).Error

	return
}

// GetBatchFromReceiveArea 批量查找 收件人区
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveArea(receiveAreas []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_area` IN (?)", receiveAreas).Find(&results).Error

	return
}

// GetFromReceiveStreet 通过receive_street获取内容 收件人街道
func (obj *_LgRiskControlOrderMgr) GetFromReceiveStreet(receiveStreet string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_street` = ?", receiveStreet).Find(&results).Error

	return
}

// GetBatchFromReceiveStreet 批量查找 收件人街道
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveStreet(receiveStreets []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_street` IN (?)", receiveStreets).Find(&results).Error

	return
}

// GetFromReceiveHouseNumber 通过receive_house_number获取内容 收件人门牌号
func (obj *_LgRiskControlOrderMgr) GetFromReceiveHouseNumber(receiveHouseNumber string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_house_number` = ?", receiveHouseNumber).Find(&results).Error

	return
}

// GetBatchFromReceiveHouseNumber 批量查找 收件人门牌号
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveHouseNumber(receiveHouseNumbers []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_house_number` IN (?)", receiveHouseNumbers).Find(&results).Error

	return
}

// GetFromReceiveAddress1 通过receive_address1获取内容 收件人地址1
func (obj *_LgRiskControlOrderMgr) GetFromReceiveAddress1(receiveAddress1 string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_address1` = ?", receiveAddress1).Find(&results).Error

	return
}

// GetBatchFromReceiveAddress1 批量查找 收件人地址1
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveAddress1(receiveAddress1s []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_address1` IN (?)", receiveAddress1s).Find(&results).Error

	return
}

// GetFromReceiveAddress2 通过receive_address2获取内容 收件人地址2
func (obj *_LgRiskControlOrderMgr) GetFromReceiveAddress2(receiveAddress2 string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_address2` = ?", receiveAddress2).Find(&results).Error

	return
}

// GetBatchFromReceiveAddress2 批量查找 收件人地址2
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveAddress2(receiveAddress2s []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_address2` IN (?)", receiveAddress2s).Find(&results).Error

	return
}

// GetFromSenderCountry 通过sender_country获取内容 发件人国家
func (obj *_LgRiskControlOrderMgr) GetFromSenderCountry(senderCountry string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_country` = ?", senderCountry).Find(&results).Error

	return
}

// GetBatchFromSenderCountry 批量查找 发件人国家
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderCountry(senderCountrys []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_country` IN (?)", senderCountrys).Find(&results).Error

	return
}

// GetFromSenderName 通过sender_name获取内容 发件人姓名
func (obj *_LgRiskControlOrderMgr) GetFromSenderName(senderName string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_name` = ?", senderName).Find(&results).Error

	return
}

// GetBatchFromSenderName 批量查找 发件人姓名
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderName(senderNames []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_name` IN (?)", senderNames).Find(&results).Error

	return
}

// GetFromSenderPhone 通过sender_phone获取内容 发件人电话
func (obj *_LgRiskControlOrderMgr) GetFromSenderPhone(senderPhone string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_phone` = ?", senderPhone).Find(&results).Error

	return
}

// GetBatchFromSenderPhone 批量查找 发件人电话
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderPhone(senderPhones []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_phone` IN (?)", senderPhones).Find(&results).Error

	return
}

// GetFromSenderMobile 通过sender_mobile获取内容 发件人手机
func (obj *_LgRiskControlOrderMgr) GetFromSenderMobile(senderMobile string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_mobile` = ?", senderMobile).Find(&results).Error

	return
}

// GetBatchFromSenderMobile 批量查找 发件人手机
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderMobile(senderMobiles []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_mobile` IN (?)", senderMobiles).Find(&results).Error

	return
}

// GetFromSenderTax 通过sender_tax获取内容 发件人税号
func (obj *_LgRiskControlOrderMgr) GetFromSenderTax(senderTax string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_tax` = ?", senderTax).Find(&results).Error

	return
}

// GetBatchFromSenderTax 批量查找 发件人税号
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderTax(senderTaxs []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_tax` IN (?)", senderTaxs).Find(&results).Error

	return
}

// GetFromSenderPassport 通过sender_passport获取内容 发件人护照号
func (obj *_LgRiskControlOrderMgr) GetFromSenderPassport(senderPassport string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_passport` = ?", senderPassport).Find(&results).Error

	return
}

// GetBatchFromSenderPassport 批量查找 发件人护照号
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderPassport(senderPassports []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_passport` IN (?)", senderPassports).Find(&results).Error

	return
}

// GetFromSenderMail 通过sender_mail获取内容 发件人邮箱
func (obj *_LgRiskControlOrderMgr) GetFromSenderMail(senderMail string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_mail` = ?", senderMail).Find(&results).Error

	return
}

// GetBatchFromSenderMail 批量查找 发件人邮箱
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderMail(senderMails []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_mail` IN (?)", senderMails).Find(&results).Error

	return
}

// GetFromSenderCompany 通过sender_company获取内容 发件人公司
func (obj *_LgRiskControlOrderMgr) GetFromSenderCompany(senderCompany string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_company` = ?", senderCompany).Find(&results).Error

	return
}

// GetBatchFromSenderCompany 批量查找 发件人公司
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderCompany(senderCompanys []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_company` IN (?)", senderCompanys).Find(&results).Error

	return
}

// GetFromSenderZipcode 通过sender_zipcode获取内容 发件人邮编
func (obj *_LgRiskControlOrderMgr) GetFromSenderZipcode(senderZipcode string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_zipcode` = ?", senderZipcode).Find(&results).Error

	return
}

// GetBatchFromSenderZipcode 批量查找 发件人邮编
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderZipcode(senderZipcodes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_zipcode` IN (?)", senderZipcodes).Find(&results).Error

	return
}

// GetFromSenderProvince 通过sender_province获取内容 发件人省份
func (obj *_LgRiskControlOrderMgr) GetFromSenderProvince(senderProvince string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_province` = ?", senderProvince).Find(&results).Error

	return
}

// GetBatchFromSenderProvince 批量查找 发件人省份
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderProvince(senderProvinces []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_province` IN (?)", senderProvinces).Find(&results).Error

	return
}

// GetFromSenderCity 通过sender_city获取内容 发件人城市
func (obj *_LgRiskControlOrderMgr) GetFromSenderCity(senderCity string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_city` = ?", senderCity).Find(&results).Error

	return
}

// GetBatchFromSenderCity 批量查找 发件人城市
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderCity(senderCitys []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_city` IN (?)", senderCitys).Find(&results).Error

	return
}

// GetFromSenderArea 通过sender_area获取内容 发件人区
func (obj *_LgRiskControlOrderMgr) GetFromSenderArea(senderArea string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_area` = ?", senderArea).Find(&results).Error

	return
}

// GetBatchFromSenderArea 批量查找 发件人区
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderArea(senderAreas []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_area` IN (?)", senderAreas).Find(&results).Error

	return
}

// GetFromSenderStreet 通过sender_street获取内容 发件人街道
func (obj *_LgRiskControlOrderMgr) GetFromSenderStreet(senderStreet string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_street` = ?", senderStreet).Find(&results).Error

	return
}

// GetBatchFromSenderStreet 批量查找 发件人街道
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderStreet(senderStreets []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_street` IN (?)", senderStreets).Find(&results).Error

	return
}

// GetFromSenderHouseNumber 通过sender_house_number获取内容 发件人门牌号
func (obj *_LgRiskControlOrderMgr) GetFromSenderHouseNumber(senderHouseNumber string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_house_number` = ?", senderHouseNumber).Find(&results).Error

	return
}

// GetBatchFromSenderHouseNumber 批量查找 发件人门牌号
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderHouseNumber(senderHouseNumbers []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_house_number` IN (?)", senderHouseNumbers).Find(&results).Error

	return
}

// GetFromSenderAddress 通过sender_address获取内容 发件人地址
func (obj *_LgRiskControlOrderMgr) GetFromSenderAddress(senderAddress string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_address` = ?", senderAddress).Find(&results).Error

	return
}

// GetBatchFromSenderAddress 批量查找 发件人地址
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderAddress(senderAddresss []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_address` IN (?)", senderAddresss).Find(&results).Error

	return
}

// GetFromPackageType 通过package_type获取内容 包裹类型
func (obj *_LgRiskControlOrderMgr) GetFromPackageType(packageType string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`package_type` = ?", packageType).Find(&results).Error

	return
}

// GetBatchFromPackageType 批量查找 包裹类型
func (obj *_LgRiskControlOrderMgr) GetBatchFromPackageType(packageTypes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`package_type` IN (?)", packageTypes).Find(&results).Error

	return
}

// GetFromHasBattery 通过has_battery获取内容 是否带电
func (obj *_LgRiskControlOrderMgr) GetFromHasBattery(hasBattery int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`has_battery` = ?", hasBattery).Find(&results).Error

	return
}

// GetBatchFromHasBattery 批量查找 是否带电
func (obj *_LgRiskControlOrderMgr) GetBatchFromHasBattery(hasBatterys []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`has_battery` IN (?)", hasBatterys).Find(&results).Error

	return
}

// GetFromBatteryType 通过battery_type获取内容 电池类型
func (obj *_LgRiskControlOrderMgr) GetFromBatteryType(batteryType string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`battery_type` = ?", batteryType).Find(&results).Error

	return
}

// GetBatchFromBatteryType 批量查找 电池类型
func (obj *_LgRiskControlOrderMgr) GetBatchFromBatteryType(batteryTypes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`battery_type` IN (?)", batteryTypes).Find(&results).Error

	return
}

// GetFromHasBack 通过has_back获取内容 是否退回
func (obj *_LgRiskControlOrderMgr) GetFromHasBack(hasBack int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`has_back` = ?", hasBack).Find(&results).Error

	return
}

// GetBatchFromHasBack 批量查找 是否退回
func (obj *_LgRiskControlOrderMgr) GetBatchFromHasBack(hasBacks []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`has_back` IN (?)", hasBacks).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注信息
func (obj *_LgRiskControlOrderMgr) GetFromRemarks(remarks string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注信息
func (obj *_LgRiskControlOrderMgr) GetBatchFromRemarks(remarkss []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromDistributionInformation 通过distribution_information获取内容 配货信息
func (obj *_LgRiskControlOrderMgr) GetFromDistributionInformation(distributionInformation string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`distribution_information` = ?", distributionInformation).Find(&results).Error

	return
}

// GetBatchFromDistributionInformation 批量查找 配货信息
func (obj *_LgRiskControlOrderMgr) GetBatchFromDistributionInformation(distributionInformations []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`distribution_information` IN (?)", distributionInformations).Find(&results).Error

	return
}

// GetFromInsurance 通过insurance获取内容 保险
func (obj *_LgRiskControlOrderMgr) GetFromInsurance(insurance string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`insurance` = ?", insurance).Find(&results).Error

	return
}

// GetBatchFromInsurance 批量查找 保险
func (obj *_LgRiskControlOrderMgr) GetBatchFromInsurance(insurances []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`insurance` IN (?)", insurances).Find(&results).Error

	return
}

// GetFromSalesPlatform 通过sales_platform获取内容 销售平台
func (obj *_LgRiskControlOrderMgr) GetFromSalesPlatform(salesPlatform string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sales_platform` = ?", salesPlatform).Find(&results).Error

	return
}

// GetBatchFromSalesPlatform 批量查找 销售平台
func (obj *_LgRiskControlOrderMgr) GetBatchFromSalesPlatform(salesPlatforms []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sales_platform` IN (?)", salesPlatforms).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 预报重量
func (obj *_LgRiskControlOrderMgr) GetFromWeight(weight float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 预报重量
func (obj *_LgRiskControlOrderMgr) GetBatchFromWeight(weights []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 包裹数量
func (obj *_LgRiskControlOrderMgr) GetFromQuantity(quantity int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 包裹数量
func (obj *_LgRiskControlOrderMgr) GetBatchFromQuantity(quantitys []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromDeclaredValue 通过declared_value获取内容 申报价值
func (obj *_LgRiskControlOrderMgr) GetFromDeclaredValue(declaredValue float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`declared_value` = ?", declaredValue).Find(&results).Error

	return
}

// GetBatchFromDeclaredValue 批量查找 申报价值
func (obj *_LgRiskControlOrderMgr) GetBatchFromDeclaredValue(declaredValues []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`declared_value` IN (?)", declaredValues).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 订单状态
func (obj *_LgRiskControlOrderMgr) GetFromStatus(status string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 订单状态
func (obj *_LgRiskControlOrderMgr) GetBatchFromStatus(statuss []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgRiskControlOrderMgr) GetFromCreateTime(createTime time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgRiskControlOrderMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromWeighingWeight 通过weighing_weight获取内容 称重重量
func (obj *_LgRiskControlOrderMgr) GetFromWeighingWeight(weighingWeight float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`weighing_weight` = ?", weighingWeight).Find(&results).Error

	return
}

// GetBatchFromWeighingWeight 批量查找 称重重量
func (obj *_LgRiskControlOrderMgr) GetBatchFromWeighingWeight(weighingWeights []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`weighing_weight` IN (?)", weighingWeights).Find(&results).Error

	return
}

// GetFromVolumeWeight 通过volume_weight获取内容 材积重量
func (obj *_LgRiskControlOrderMgr) GetFromVolumeWeight(volumeWeight float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`volume_weight` = ?", volumeWeight).Find(&results).Error

	return
}

// GetBatchFromVolumeWeight 批量查找 材积重量
func (obj *_LgRiskControlOrderMgr) GetBatchFromVolumeWeight(volumeWeights []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`volume_weight` IN (?)", volumeWeights).Find(&results).Error

	return
}

// GetFromChargeWeight 通过charge_weight获取内容 收费重量
func (obj *_LgRiskControlOrderMgr) GetFromChargeWeight(chargeWeight float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`charge_weight` = ?", chargeWeight).Find(&results).Error

	return
}

// GetBatchFromChargeWeight 批量查找 收费重量
func (obj *_LgRiskControlOrderMgr) GetBatchFromChargeWeight(chargeWeights []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`charge_weight` IN (?)", chargeWeights).Find(&results).Error

	return
}

// GetFromStandardFee 通过standard_fee获取内容 标准费用
func (obj *_LgRiskControlOrderMgr) GetFromStandardFee(standardFee float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`standard_fee` = ?", standardFee).Find(&results).Error

	return
}

// GetBatchFromStandardFee 批量查找 标准费用
func (obj *_LgRiskControlOrderMgr) GetBatchFromStandardFee(standardFees []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`standard_fee` IN (?)", standardFees).Find(&results).Error

	return
}

// GetFromCurrency 通过currency获取内容 币种
func (obj *_LgRiskControlOrderMgr) GetFromCurrency(currency string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`currency` = ?", currency).Find(&results).Error

	return
}

// GetBatchFromCurrency 批量查找 币种
func (obj *_LgRiskControlOrderMgr) GetBatchFromCurrency(currencys []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`currency` IN (?)", currencys).Find(&results).Error

	return
}

// GetFromBuyersFee 通过buyers_fee获取内容 下家费用
func (obj *_LgRiskControlOrderMgr) GetFromBuyersFee(buyersFee float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`buyers_fee` = ?", buyersFee).Find(&results).Error

	return
}

// GetBatchFromBuyersFee 批量查找 下家费用
func (obj *_LgRiskControlOrderMgr) GetBatchFromBuyersFee(buyersFees []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`buyers_fee` IN (?)", buyersFees).Find(&results).Error

	return
}

// GetFromBuyersWeight 通过buyers_weight获取内容 下家重量
func (obj *_LgRiskControlOrderMgr) GetFromBuyersWeight(buyersWeight float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`buyers_weight` = ?", buyersWeight).Find(&results).Error

	return
}

// GetBatchFromBuyersWeight 批量查找 下家重量
func (obj *_LgRiskControlOrderMgr) GetBatchFromBuyersWeight(buyersWeights []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`buyers_weight` IN (?)", buyersWeights).Find(&results).Error

	return
}

// GetFromOtherFee 通过other_fee获取内容 其他费用
func (obj *_LgRiskControlOrderMgr) GetFromOtherFee(otherFee float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`other_fee` = ?", otherFee).Find(&results).Error

	return
}

// GetBatchFromOtherFee 批量查找 其他费用
func (obj *_LgRiskControlOrderMgr) GetBatchFromOtherFee(otherFees []float64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`other_fee` IN (?)", otherFees).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 归属客户ID
func (obj *_LgRiskControlOrderMgr) GetFromCustomerID(customerID int64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 归属客户ID
func (obj *_LgRiskControlOrderMgr) GetBatchFromCustomerID(customerIDs []int64) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 归属客户名称
func (obj *_LgRiskControlOrderMgr) GetFromCustomerName(customerName string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 归属客户名称
func (obj *_LgRiskControlOrderMgr) GetBatchFromCustomerName(customerNames []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

// GetFromPrepaymentVat 通过prepayment_vat获取内容 "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
func (obj *_LgRiskControlOrderMgr) GetFromPrepaymentVat(prepaymentVat string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`prepayment_vat` = ?", prepaymentVat).Find(&results).Error

	return
}

// GetBatchFromPrepaymentVat 批量查找 "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
func (obj *_LgRiskControlOrderMgr) GetBatchFromPrepaymentVat(prepaymentVats []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`prepayment_vat` IN (?)", prepaymentVats).Find(&results).Error

	return
}

// GetFromVat 通过vat获取内容 vat号(英国方向必填)
func (obj *_LgRiskControlOrderMgr) GetFromVat(vat string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`vat` = ?", vat).Find(&results).Error

	return
}

// GetBatchFromVat 批量查找 vat号(英国方向必填)
func (obj *_LgRiskControlOrderMgr) GetBatchFromVat(vats []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`vat` IN (?)", vats).Find(&results).Error

	return
}

// GetFromReceiveCertificateType 通过receive_certificate_type获取内容 收件人证件类型
func (obj *_LgRiskControlOrderMgr) GetFromReceiveCertificateType(receiveCertificateType string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_certificate_type` = ?", receiveCertificateType).Find(&results).Error

	return
}

// GetBatchFromReceiveCertificateType 批量查找 收件人证件类型
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveCertificateType(receiveCertificateTypes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_certificate_type` IN (?)", receiveCertificateTypes).Find(&results).Error

	return
}

// GetFromReceiveCertificateCode 通过receive_certificate_code获取内容 收件人证件号码
func (obj *_LgRiskControlOrderMgr) GetFromReceiveCertificateCode(receiveCertificateCode string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_certificate_code` = ?", receiveCertificateCode).Find(&results).Error

	return
}

// GetBatchFromReceiveCertificateCode 批量查找 收件人证件号码
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveCertificateCode(receiveCertificateCodes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_certificate_code` IN (?)", receiveCertificateCodes).Find(&results).Error

	return
}

// GetFromSenderCertificateType 通过sender_certificate_type获取内容 发件人证件类型
func (obj *_LgRiskControlOrderMgr) GetFromSenderCertificateType(senderCertificateType string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_certificate_type` = ?", senderCertificateType).Find(&results).Error

	return
}

// GetBatchFromSenderCertificateType 批量查找 发件人证件类型
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderCertificateType(senderCertificateTypes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_certificate_type` IN (?)", senderCertificateTypes).Find(&results).Error

	return
}

// GetFromSenderCertificateCode 通过sender_certificate_code获取内容 发件人证件号码
func (obj *_LgRiskControlOrderMgr) GetFromSenderCertificateCode(senderCertificateCode string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_certificate_code` = ?", senderCertificateCode).Find(&results).Error

	return
}

// GetBatchFromSenderCertificateCode 批量查找 发件人证件号码
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderCertificateCode(senderCertificateCodes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_certificate_code` IN (?)", senderCertificateCodes).Find(&results).Error

	return
}

// GetFromServiceChannelCode 通过service_channel_code获取内容 api渠道代码
func (obj *_LgRiskControlOrderMgr) GetFromServiceChannelCode(serviceChannelCode string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`service_channel_code` = ?", serviceChannelCode).Find(&results).Error

	return
}

// GetBatchFromServiceChannelCode 批量查找 api渠道代码
func (obj *_LgRiskControlOrderMgr) GetBatchFromServiceChannelCode(serviceChannelCodes []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`service_channel_code` IN (?)", serviceChannelCodes).Find(&results).Error

	return
}

// GetFromDeliveryTerms 通过delivery_terms获取内容 贸易条款- DDU DDP
func (obj *_LgRiskControlOrderMgr) GetFromDeliveryTerms(deliveryTerms string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`delivery_terms` = ?", deliveryTerms).Find(&results).Error

	return
}

// GetBatchFromDeliveryTerms 批量查找 贸易条款- DDU DDP
func (obj *_LgRiskControlOrderMgr) GetBatchFromDeliveryTerms(deliveryTermss []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`delivery_terms` IN (?)", deliveryTermss).Find(&results).Error

	return
}

// GetFromReceiveCountryName 通过receive_country_name获取内容 收件人国家中文
func (obj *_LgRiskControlOrderMgr) GetFromReceiveCountryName(receiveCountryName string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_country_name` = ?", receiveCountryName).Find(&results).Error

	return
}

// GetBatchFromReceiveCountryName 批量查找 收件人国家中文
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiveCountryName(receiveCountryNames []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receive_country_name` IN (?)", receiveCountryNames).Find(&results).Error

	return
}

// GetFromSenderCountryName 通过sender_country_name获取内容 发件国家中文
func (obj *_LgRiskControlOrderMgr) GetFromSenderCountryName(senderCountryName string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_country_name` = ?", senderCountryName).Find(&results).Error

	return
}

// GetBatchFromSenderCountryName 批量查找 发件国家中文
func (obj *_LgRiskControlOrderMgr) GetBatchFromSenderCountryName(senderCountryNames []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`sender_country_name` IN (?)", senderCountryNames).Find(&results).Error

	return
}

// GetFromOrderType 通过order_type获取内容 订单类型(0-预报订单，1-物流订单)
func (obj *_LgRiskControlOrderMgr) GetFromOrderType(orderType int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`order_type` = ?", orderType).Find(&results).Error

	return
}

// GetBatchFromOrderType 批量查找 订单类型(0-预报订单，1-物流订单)
func (obj *_LgRiskControlOrderMgr) GetBatchFromOrderType(orderTypes []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`order_type` IN (?)", orderTypes).Find(&results).Error

	return
}

// GetFromReceiptTime 通过receipt_time获取内容 入库时间
func (obj *_LgRiskControlOrderMgr) GetFromReceiptTime(receiptTime time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// GetBatchFromReceiptTime 批量查找 入库时间
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiptTime(receiptTimes []time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receipt_time` IN (?)", receiptTimes).Find(&results).Error

	return
}

// GetFromReceiptUser 通过receipt_user获取内容 入库人
func (obj *_LgRiskControlOrderMgr) GetFromReceiptUser(receiptUser int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receipt_user` = ?", receiptUser).Find(&results).Error

	return
}

// GetBatchFromReceiptUser 批量查找 入库人
func (obj *_LgRiskControlOrderMgr) GetBatchFromReceiptUser(receiptUsers []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`receipt_user` IN (?)", receiptUsers).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 出库时间
func (obj *_LgRiskControlOrderMgr) GetFromSendTime(sendTime time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 出库时间
func (obj *_LgRiskControlOrderMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromSendUser 通过send_user获取内容 出库人
func (obj *_LgRiskControlOrderMgr) GetFromSendUser(sendUser int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`send_user` = ?", sendUser).Find(&results).Error

	return
}

// GetBatchFromSendUser 批量查找 出库人
func (obj *_LgRiskControlOrderMgr) GetBatchFromSendUser(sendUsers []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`send_user` IN (?)", sendUsers).Find(&results).Error

	return
}

// GetFromPicOssLink 通过pic_oss_link获取内容 包裹图片oss地址
func (obj *_LgRiskControlOrderMgr) GetFromPicOssLink(picOssLink string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`pic_oss_link` = ?", picOssLink).Find(&results).Error

	return
}

// GetBatchFromPicOssLink 批量查找 包裹图片oss地址
func (obj *_LgRiskControlOrderMgr) GetBatchFromPicOssLink(picOssLinks []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`pic_oss_link` IN (?)", picOssLinks).Find(&results).Error

	return
}

// GetFromRiskMessage 通过risk_message获取内容 风控信息
func (obj *_LgRiskControlOrderMgr) GetFromRiskMessage(riskMessage string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`risk_message` = ?", riskMessage).Find(&results).Error

	return
}

// GetBatchFromRiskMessage 批量查找 风控信息
func (obj *_LgRiskControlOrderMgr) GetBatchFromRiskMessage(riskMessages []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`risk_message` IN (?)", riskMessages).Find(&results).Error

	return
}

// GetFromStatusRisk 通过status_risk获取内容 风控订单状态:normal:正常,abnormal:异常 warn:警告
func (obj *_LgRiskControlOrderMgr) GetFromStatusRisk(statusRisk string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`status_risk` = ?", statusRisk).Find(&results).Error

	return
}

// GetBatchFromStatusRisk 批量查找 风控订单状态:normal:正常,abnormal:异常 warn:警告
func (obj *_LgRiskControlOrderMgr) GetBatchFromStatusRisk(statusRisks []string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`status_risk` IN (?)", statusRisks).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgRiskControlOrderMgr) GetFromCreateUser(createUser int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgRiskControlOrderMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgRiskControlOrderMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgRiskControlOrderMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgRiskControlOrderMgr) GetFromUpdateUser(updateUser int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgRiskControlOrderMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgRiskControlOrderMgr) GetFromVersion(version int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgRiskControlOrderMgr) GetBatchFromVersion(versions []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgRiskControlOrderMgr) GetFromDeleted(deleted int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgRiskControlOrderMgr) GetBatchFromDeleted(deleteds []int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgRiskControlOrderMgr) FetchByPrimaryKey(id int) (result LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgRiskControlOrderLogisticsNumber primary or index 获取唯一内容
func (obj *_LgRiskControlOrderMgr) FetchUniqueByLgRiskControlOrderLogisticsNumber(logisticsNumber string) (result LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`logistics_number` = ?", logisticsNumber).First(&result).Error

	return
}

// FetchIndexByLgRiskControlOrderLgCustomerChannelIDFk  获取多个内容
func (obj *_LgRiskControlOrderMgr) FetchIndexByLgRiskControlOrderLgCustomerChannelIDFk(customerChannelID int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByLgOrderChannelNameIndex  获取多个内容
func (obj *_LgRiskControlOrderMgr) FetchIndexByLgOrderChannelNameIndex(customerChannelName string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// FetchIndexByLgRiskControlOrderCreateTimeIndex  获取多个内容
func (obj *_LgRiskControlOrderMgr) FetchIndexByLgRiskControlOrderCreateTimeIndex(createTime time.Time) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByLgRiskControlOrderStatusRiskIndex  获取多个内容
func (obj *_LgRiskControlOrderMgr) FetchIndexByLgRiskControlOrderStatusRiskIndex(statusRisk string) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`status_risk` = ?", statusRisk).Find(&results).Error

	return
}

// FetchIndexByLgRiskControlOrderCreateUserIndex  获取多个内容
func (obj *_LgRiskControlOrderMgr) FetchIndexByLgRiskControlOrderCreateUserIndex(createUser int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// FetchIndexByLgRiskControlOrderDeletedIndex  获取多个内容
func (obj *_LgRiskControlOrderMgr) FetchIndexByLgRiskControlOrderDeletedIndex(deleted int) (results []*LgRiskControlOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgRiskControlOrder{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}
