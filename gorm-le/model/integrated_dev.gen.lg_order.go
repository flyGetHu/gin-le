package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _LgOrderMgr struct {
	*_BaseMgr
}

// LgOrderMgr open func
func LgOrderMgr(db *gorm.DB) *_LgOrderMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderMgr) GetTableName() string {
	return "lg_order"
}

// Reset 重置gorm会话
func (obj *_LgOrderMgr) Reset() *_LgOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderMgr) Get() (result LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgOrderMgr) Gets() (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgOrderMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumber order_number获取 订单号
func (obj *_LgOrderMgr) WithOrderNumber(orderNumber string) Option {
	return optionFunc(func(o *options) { o.query["order_number"] = orderNumber })
}

// WithReferenceNumber reference_number获取 参考号
func (obj *_LgOrderMgr) WithReferenceNumber(referenceNumber string) Option {
	return optionFunc(func(o *options) { o.query["reference_number"] = referenceNumber })
}

// WithLogisticsNumber logistics_number获取 物流单号
func (obj *_LgOrderMgr) WithLogisticsNumber(logisticsNumber string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number"] = logisticsNumber })
}

// WithLogisticsNumberFinal logistics_number_final获取 最终物流单号
func (obj *_LgOrderMgr) WithLogisticsNumberFinal(logisticsNumberFinal string) Option {
	return optionFunc(func(o *options) { o.query["logistics_number_final"] = logisticsNumberFinal })
}

// WithExpressNum express_num获取 订单快递单号
func (obj *_LgOrderMgr) WithExpressNum(expressNum string) Option {
	return optionFunc(func(o *options) { o.query["express_num"] = expressNum })
}

// WithPlatformNumber platform_number获取 平台订单
func (obj *_LgOrderMgr) WithPlatformNumber(platformNumber string) Option {
	return optionFunc(func(o *options) { o.query["platform_number"] = platformNumber })
}

// WithProviderOrderID provider_order_id获取 服务商系统订单ID
func (obj *_LgOrderMgr) WithProviderOrderID(providerOrderID string) Option {
	return optionFunc(func(o *options) { o.query["provider_order_id"] = providerOrderID })
}

// WithLabelURL label_url获取 面单地址
func (obj *_LgOrderMgr) WithLabelURL(labelURL string) Option {
	return optionFunc(func(o *options) { o.query["label_url"] = labelURL })
}

// WithLabelURLFinal label_url_final获取 最终面单地址
func (obj *_LgOrderMgr) WithLabelURLFinal(labelURLFinal string) Option {
	return optionFunc(func(o *options) { o.query["label_url_final"] = labelURLFinal })
}

// WithOnlineNumber online_number获取 上网单号(17单号)
func (obj *_LgOrderMgr) WithOnlineNumber(onlineNumber string) Option {
	return optionFunc(func(o *options) { o.query["online_number"] = onlineNumber })
}

// WithBagNumber bag_number获取 袋号
func (obj *_LgOrderMgr) WithBagNumber(bagNumber string) Option {
	return optionFunc(func(o *options) { o.query["bag_number"] = bagNumber })
}

// WithTotalListNo total_list_no获取 总单号
func (obj *_LgOrderMgr) WithTotalListNo(totalListNo string) Option {
	return optionFunc(func(o *options) { o.query["total_list_no"] = totalListNo })
}

// WithAirBillNumber air_bill_number获取 航空提单号
func (obj *_LgOrderMgr) WithAirBillNumber(airBillNumber string) Option {
	return optionFunc(func(o *options) { o.query["air_bill_number"] = airBillNumber })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgOrderMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCustomerChannelName customer_channel_name获取 客户渠道名称
func (obj *_LgOrderMgr) WithCustomerChannelName(customerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_name"] = customerChannelName })
}

// WithProviderCode provider_code获取 服务商系统code
func (obj *_LgOrderMgr) WithProviderCode(providerCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_code"] = providerCode })
}

// WithProviderChannelCode provider_channel_code获取 服务商系统渠道Code
func (obj *_LgOrderMgr) WithProviderChannelCode(providerChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_code"] = providerChannelCode })
}

// WithProviderChannelName provider_Channel_Name获取 服务商渠道名称
func (obj *_LgOrderMgr) WithProviderChannelName(providerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["provider_Channel_Name"] = providerChannelName })
}

// WithReceiveCountry receive_country获取 收件人国家
func (obj *_LgOrderMgr) WithReceiveCountry(receiveCountry string) Option {
	return optionFunc(func(o *options) { o.query["receive_country"] = receiveCountry })
}

// WithReceiveName receive_name获取 收件人姓名
func (obj *_LgOrderMgr) WithReceiveName(receiveName string) Option {
	return optionFunc(func(o *options) { o.query["receive_name"] = receiveName })
}

// WithReceivePhone receive_phone获取 收件人电话
func (obj *_LgOrderMgr) WithReceivePhone(receivePhone string) Option {
	return optionFunc(func(o *options) { o.query["receive_phone"] = receivePhone })
}

// WithReceiveMobile receive_mobile获取 收件人手机
func (obj *_LgOrderMgr) WithReceiveMobile(receiveMobile string) Option {
	return optionFunc(func(o *options) { o.query["receive_mobile"] = receiveMobile })
}

// WithReceiveTax receive_tax获取 收件人税号
func (obj *_LgOrderMgr) WithReceiveTax(receiveTax string) Option {
	return optionFunc(func(o *options) { o.query["receive_tax"] = receiveTax })
}

// WithReceivePassport receive_passport获取 收件人护照号
func (obj *_LgOrderMgr) WithReceivePassport(receivePassport string) Option {
	return optionFunc(func(o *options) { o.query["receive_passport"] = receivePassport })
}

// WithReceiveMail receive_mail获取 收件人邮箱
func (obj *_LgOrderMgr) WithReceiveMail(receiveMail string) Option {
	return optionFunc(func(o *options) { o.query["receive_mail"] = receiveMail })
}

// WithReceiveCompany receive_company获取 收件人公司
func (obj *_LgOrderMgr) WithReceiveCompany(receiveCompany string) Option {
	return optionFunc(func(o *options) { o.query["receive_company"] = receiveCompany })
}

// WithReceiveZipcode receive_zipcode获取 收件人邮编
func (obj *_LgOrderMgr) WithReceiveZipcode(receiveZipcode string) Option {
	return optionFunc(func(o *options) { o.query["receive_zipcode"] = receiveZipcode })
}

// WithReceiveProvince receive_province获取 收件人省份
func (obj *_LgOrderMgr) WithReceiveProvince(receiveProvince string) Option {
	return optionFunc(func(o *options) { o.query["receive_province"] = receiveProvince })
}

// WithReceiveHouseNumber receive_house_number获取 收件人门牌号
func (obj *_LgOrderMgr) WithReceiveHouseNumber(receiveHouseNumber string) Option {
	return optionFunc(func(o *options) { o.query["receive_house_number"] = receiveHouseNumber })
}

// WithReceiveArea receive_area获取 收件人区
func (obj *_LgOrderMgr) WithReceiveArea(receiveArea string) Option {
	return optionFunc(func(o *options) { o.query["receive_area"] = receiveArea })
}

// WithReceiveStreet receive_street获取 收件人街道
func (obj *_LgOrderMgr) WithReceiveStreet(receiveStreet string) Option {
	return optionFunc(func(o *options) { o.query["receive_street"] = receiveStreet })
}

// WithReceiveCity receive_city获取 收件人城市
func (obj *_LgOrderMgr) WithReceiveCity(receiveCity string) Option {
	return optionFunc(func(o *options) { o.query["receive_city"] = receiveCity })
}

// WithReceiveAddress1 receive_address1获取 收件人地址1
func (obj *_LgOrderMgr) WithReceiveAddress1(receiveAddress1 string) Option {
	return optionFunc(func(o *options) { o.query["receive_address1"] = receiveAddress1 })
}

// WithReceiveAddress2 receive_address2获取 收件人地址2
func (obj *_LgOrderMgr) WithReceiveAddress2(receiveAddress2 string) Option {
	return optionFunc(func(o *options) { o.query["receive_address2"] = receiveAddress2 })
}

// WithSenderCountry sender_country获取 发件人国家
func (obj *_LgOrderMgr) WithSenderCountry(senderCountry string) Option {
	return optionFunc(func(o *options) { o.query["sender_country"] = senderCountry })
}

// WithSenderName sender_name获取 发件人姓名
func (obj *_LgOrderMgr) WithSenderName(senderName string) Option {
	return optionFunc(func(o *options) { o.query["sender_name"] = senderName })
}

// WithSenderPhone sender_phone获取 发件人电话
func (obj *_LgOrderMgr) WithSenderPhone(senderPhone string) Option {
	return optionFunc(func(o *options) { o.query["sender_phone"] = senderPhone })
}

// WithSenderMobile sender_mobile获取 发件人手机
func (obj *_LgOrderMgr) WithSenderMobile(senderMobile string) Option {
	return optionFunc(func(o *options) { o.query["sender_mobile"] = senderMobile })
}

// WithSenderTax sender_tax获取 发件人税号
func (obj *_LgOrderMgr) WithSenderTax(senderTax string) Option {
	return optionFunc(func(o *options) { o.query["sender_tax"] = senderTax })
}

// WithSenderPassport sender_passport获取 发件人护照号
func (obj *_LgOrderMgr) WithSenderPassport(senderPassport string) Option {
	return optionFunc(func(o *options) { o.query["sender_passport"] = senderPassport })
}

// WithSenderMail sender_mail获取 发件人邮箱
func (obj *_LgOrderMgr) WithSenderMail(senderMail string) Option {
	return optionFunc(func(o *options) { o.query["sender_mail"] = senderMail })
}

// WithSenderCompany sender_company获取 发件人公司
func (obj *_LgOrderMgr) WithSenderCompany(senderCompany string) Option {
	return optionFunc(func(o *options) { o.query["sender_company"] = senderCompany })
}

// WithSenderZipcode sender_zipcode获取 发件人邮编
func (obj *_LgOrderMgr) WithSenderZipcode(senderZipcode string) Option {
	return optionFunc(func(o *options) { o.query["sender_zipcode"] = senderZipcode })
}

// WithSenderProvince sender_province获取 发件人省份
func (obj *_LgOrderMgr) WithSenderProvince(senderProvince string) Option {
	return optionFunc(func(o *options) { o.query["sender_province"] = senderProvince })
}

// WithSenderCity sender_city获取 发件人城市
func (obj *_LgOrderMgr) WithSenderCity(senderCity string) Option {
	return optionFunc(func(o *options) { o.query["sender_city"] = senderCity })
}

// WithSenderArea sender_area获取 发件人区
func (obj *_LgOrderMgr) WithSenderArea(senderArea string) Option {
	return optionFunc(func(o *options) { o.query["sender_area"] = senderArea })
}

// WithSenderStreet sender_street获取 发件人街道
func (obj *_LgOrderMgr) WithSenderStreet(senderStreet string) Option {
	return optionFunc(func(o *options) { o.query["sender_street"] = senderStreet })
}

// WithSenderHouseNumber sender_house_number获取 发件人门牌号
func (obj *_LgOrderMgr) WithSenderHouseNumber(senderHouseNumber string) Option {
	return optionFunc(func(o *options) { o.query["sender_house_number"] = senderHouseNumber })
}

// WithSenderAddress sender_address获取 发件人地址
func (obj *_LgOrderMgr) WithSenderAddress(senderAddress string) Option {
	return optionFunc(func(o *options) { o.query["sender_address"] = senderAddress })
}

// WithPackageType package_type获取 包裹类型
func (obj *_LgOrderMgr) WithPackageType(packageType string) Option {
	return optionFunc(func(o *options) { o.query["package_type"] = packageType })
}

// WithHasBattery has_battery获取 是否带电
func (obj *_LgOrderMgr) WithHasBattery(hasBattery int) Option {
	return optionFunc(func(o *options) { o.query["has_battery"] = hasBattery })
}

// WithBatteryType battery_type获取 电池类型
func (obj *_LgOrderMgr) WithBatteryType(batteryType string) Option {
	return optionFunc(func(o *options) { o.query["battery_type"] = batteryType })
}

// WithHasBack has_back获取 是否退回
func (obj *_LgOrderMgr) WithHasBack(hasBack int) Option {
	return optionFunc(func(o *options) { o.query["has_back"] = hasBack })
}

// WithRemarks remarks获取 备注信息
func (obj *_LgOrderMgr) WithRemarks(remarks string) Option {
	return optionFunc(func(o *options) { o.query["remarks"] = remarks })
}

// WithDistributionInformation distribution_information获取 配货信息
func (obj *_LgOrderMgr) WithDistributionInformation(distributionInformation string) Option {
	return optionFunc(func(o *options) { o.query["distribution_information"] = distributionInformation })
}

// WithInsurance insurance获取 保险
func (obj *_LgOrderMgr) WithInsurance(insurance string) Option {
	return optionFunc(func(o *options) { o.query["insurance"] = insurance })
}

// WithSalesPlatform sales_platform获取 销售平台
func (obj *_LgOrderMgr) WithSalesPlatform(salesPlatform string) Option {
	return optionFunc(func(o *options) { o.query["sales_platform"] = salesPlatform })
}

// WithWeight weight获取 预报重量
func (obj *_LgOrderMgr) WithWeight(weight float64) Option {
	return optionFunc(func(o *options) { o.query["weight"] = weight })
}

// WithQuantity quantity获取 包裹数量
func (obj *_LgOrderMgr) WithQuantity(quantity int) Option {
	return optionFunc(func(o *options) { o.query["quantity"] = quantity })
}

// WithDeclaredValue declared_value获取 申报价值
func (obj *_LgOrderMgr) WithDeclaredValue(declaredValue float64) Option {
	return optionFunc(func(o *options) { o.query["declared_value"] = declaredValue })
}

// WithStatus status获取 订单状态
func (obj *_LgOrderMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTransportStatus transport_status获取 转运状态
func (obj *_LgOrderMgr) WithTransportStatus(transportStatus string) Option {
	return optionFunc(func(o *options) { o.query["transport_status"] = transportStatus })
}

// WithLastTrackStatus last_track_status获取 最后一条轨迹状态
func (obj *_LgOrderMgr) WithLastTrackStatus(lastTrackStatus string) Option {
	return optionFunc(func(o *options) { o.query["last_track_status"] = lastTrackStatus })
}

// WithLastTrackAddress last_track_address获取 最后一条轨迹
func (obj *_LgOrderMgr) WithLastTrackAddress(lastTrackAddress string) Option {
	return optionFunc(func(o *options) { o.query["last_track_address"] = lastTrackAddress })
}

// WithLastTrackTime last_track_time获取 最后一条轨迹的时间
func (obj *_LgOrderMgr) WithLastTrackTime(lastTrackTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_track_time"] = lastTrackTime })
}

// WithProblemType problem_type获取 问题件类型
func (obj *_LgOrderMgr) WithProblemType(problemType string) Option {
	return optionFunc(func(o *options) { o.query["problem_type"] = problemType })
}

// WithProblemReason problem_reason获取 问题件原因
func (obj *_LgOrderMgr) WithProblemReason(problemReason string) Option {
	return optionFunc(func(o *options) { o.query["problem_reason"] = problemReason })
}

// WithWeighingWeight weighing_weight获取 称重重量
func (obj *_LgOrderMgr) WithWeighingWeight(weighingWeight float64) Option {
	return optionFunc(func(o *options) { o.query["weighing_weight"] = weighingWeight })
}

// WithVolumeWeight volume_weight获取 材积重量
func (obj *_LgOrderMgr) WithVolumeWeight(volumeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["volume_weight"] = volumeWeight })
}

// WithChargeWeight charge_weight获取 收费重量
func (obj *_LgOrderMgr) WithChargeWeight(chargeWeight float64) Option {
	return optionFunc(func(o *options) { o.query["charge_weight"] = chargeWeight })
}

// WithStandardFee standard_fee获取 标准费用
func (obj *_LgOrderMgr) WithStandardFee(standardFee float64) Option {
	return optionFunc(func(o *options) { o.query["standard_fee"] = standardFee })
}

// WithCurrency currency获取 币种
func (obj *_LgOrderMgr) WithCurrency(currency string) Option {
	return optionFunc(func(o *options) { o.query["currency"] = currency })
}

// WithBuyersFee buyers_fee获取 下家费用
func (obj *_LgOrderMgr) WithBuyersFee(buyersFee float64) Option {
	return optionFunc(func(o *options) { o.query["buyers_fee"] = buyersFee })
}

// WithBuyersWeight buyers_weight获取 下家重量
func (obj *_LgOrderMgr) WithBuyersWeight(buyersWeight float64) Option {
	return optionFunc(func(o *options) { o.query["buyers_weight"] = buyersWeight })
}

// WithOtherFee other_fee获取 其他费用
func (obj *_LgOrderMgr) WithOtherFee(otherFee float64) Option {
	return optionFunc(func(o *options) { o.query["other_fee"] = otherFee })
}

// WithCustomerID customer_id获取 归属客户ID
func (obj *_LgOrderMgr) WithCustomerID(customerID int64) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithCustomerName customer_name获取 归属客户名称
func (obj *_LgOrderMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_LgOrderMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithPrepaymentVat prepayment_vat获取 "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
func (obj *_LgOrderMgr) WithPrepaymentVat(prepaymentVat string) Option {
	return optionFunc(func(o *options) { o.query["prepayment_vat"] = prepaymentVat })
}

// WithVat vat获取 vat号(英国方向必填)
func (obj *_LgOrderMgr) WithVat(vat string) Option {
	return optionFunc(func(o *options) { o.query["vat"] = vat })
}

// WithReceiveCertificateType receive_certificate_type获取 收件人证件类型
func (obj *_LgOrderMgr) WithReceiveCertificateType(receiveCertificateType string) Option {
	return optionFunc(func(o *options) { o.query["receive_certificate_type"] = receiveCertificateType })
}

// WithReceiveCertificateCode receive_certificate_code获取 收件人证件号码
func (obj *_LgOrderMgr) WithReceiveCertificateCode(receiveCertificateCode string) Option {
	return optionFunc(func(o *options) { o.query["receive_certificate_code"] = receiveCertificateCode })
}

// WithSenderCertificateType sender_certificate_type获取 发件人证件类型
func (obj *_LgOrderMgr) WithSenderCertificateType(senderCertificateType string) Option {
	return optionFunc(func(o *options) { o.query["sender_certificate_type"] = senderCertificateType })
}

// WithSenderCertificateCode sender_certificate_code获取 发件人证件号码
func (obj *_LgOrderMgr) WithSenderCertificateCode(senderCertificateCode string) Option {
	return optionFunc(func(o *options) { o.query["sender_certificate_code"] = senderCertificateCode })
}

// WithServiceChannelCode service_channel_code获取 服务商渠道代码
func (obj *_LgOrderMgr) WithServiceChannelCode(serviceChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["service_channel_code"] = serviceChannelCode })
}

// WithDeliveryTerms delivery_terms获取 贸易条款-(部分渠道巴西方向必填)- DDU DDP
func (obj *_LgOrderMgr) WithDeliveryTerms(deliveryTerms string) Option {
	return optionFunc(func(o *options) { o.query["delivery_terms"] = deliveryTerms })
}

// WithReceiveCountryName receive_country_name获取 收件人国家中文
func (obj *_LgOrderMgr) WithReceiveCountryName(receiveCountryName string) Option {
	return optionFunc(func(o *options) { o.query["receive_country_name"] = receiveCountryName })
}

// WithSenderCountryName sender_country_name获取 发件国家中文
func (obj *_LgOrderMgr) WithSenderCountryName(senderCountryName string) Option {
	return optionFunc(func(o *options) { o.query["sender_country_name"] = senderCountryName })
}

// WithBillNumber bill_number获取 账单编号
func (obj *_LgOrderMgr) WithBillNumber(billNumber string) Option {
	return optionFunc(func(o *options) { o.query["bill_number"] = billNumber })
}

// WithApproximateCost approximate_cost获取 预报费用
func (obj *_LgOrderMgr) WithApproximateCost(approximateCost float64) Option {
	return optionFunc(func(o *options) { o.query["approximate_cost"] = approximateCost })
}

// WithWeighingCost weighing_cost获取 称重费用
func (obj *_LgOrderMgr) WithWeighingCost(weighingCost float64) Option {
	return optionFunc(func(o *options) { o.query["weighing_cost"] = weighingCost })
}

// WithPrepaidAmount prepaid_amount获取 预扣款金额
func (obj *_LgOrderMgr) WithPrepaidAmount(prepaidAmount float64) Option {
	return optionFunc(func(o *options) { o.query["prepaid_amount"] = prepaidAmount })
}

// WithActualCost actual_cost获取 实收费用
func (obj *_LgOrderMgr) WithActualCost(actualCost float64) Option {
	return optionFunc(func(o *options) { o.query["actual_cost"] = actualCost })
}

// WithPrimeCost prime_cost获取 成本
func (obj *_LgOrderMgr) WithPrimeCost(primeCost float64) Option {
	return optionFunc(func(o *options) { o.query["prime_cost"] = primeCost })
}

// WithOrderType order_type获取 订单类型(0-预报订单，1-物流订单)
func (obj *_LgOrderMgr) WithOrderType(orderType int) Option {
	return optionFunc(func(o *options) { o.query["order_type"] = orderType })
}

// WithExchangeOrder exchange_order获取 是否需要换单 0 不需要1 需要
func (obj *_LgOrderMgr) WithExchangeOrder(exchangeOrder []uint8) Option {
	return optionFunc(func(o *options) { o.query["exchange_order"] = exchangeOrder })
}

// WithReceiptTime receipt_time获取 入库时间
func (obj *_LgOrderMgr) WithReceiptTime(receiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["receipt_time"] = receiptTime })
}

// WithReceiptUser receipt_user获取 入库人
func (obj *_LgOrderMgr) WithReceiptUser(receiptUser int) Option {
	return optionFunc(func(o *options) { o.query["receipt_user"] = receiptUser })
}

// WithSendTime send_time获取 出库时间
func (obj *_LgOrderMgr) WithSendTime(sendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithSendUser send_user获取 出库人
func (obj *_LgOrderMgr) WithSendUser(sendUser int) Option {
	return optionFunc(func(o *options) { o.query["send_user"] = sendUser })
}

// WithPicOssLink pic_oss_link获取 包裹图片oss地址
func (obj *_LgOrderMgr) WithPicOssLink(picOssLink string) Option {
	return optionFunc(func(o *options) { o.query["pic_oss_link"] = picOssLink })
}

// WithTest test获取 是否为测试订单
func (obj *_LgOrderMgr) WithTest(test []uint8) Option {
	return optionFunc(func(o *options) { o.query["test"] = test })
}

// WithOrderSource order_source获取 订单来源：1-默认，2-平台订单
func (obj *_LgOrderMgr) WithOrderSource(orderSource int) Option {
	return optionFunc(func(o *options) { o.query["order_source"] = orderSource })
}

// WithPlatformCode platform_code获取 平台代码
func (obj *_LgOrderMgr) WithPlatformCode(platformCode string) Option {
	return optionFunc(func(o *options) { o.query["platform_code"] = platformCode })
}

// WithPlatformName platform_name获取 平台名称
func (obj *_LgOrderMgr) WithPlatformName(platformName string) Option {
	return optionFunc(func(o *options) { o.query["platform_name"] = platformName })
}

// WithPlatformSellerName platform_seller_name获取 平台卖家名称
func (obj *_LgOrderMgr) WithPlatformSellerName(platformSellerName string) Option {
	return optionFunc(func(o *options) { o.query["platform_seller_name"] = platformSellerName })
}

// WithPlatformOrderTime platform_order_time获取 平台订单时间
func (obj *_LgOrderMgr) WithPlatformOrderTime(platformOrderTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["platform_order_time"] = platformOrderTime })
}

// WithPostOfficeBranchName post_office_branch_name获取 邮政局分行名称
func (obj *_LgOrderMgr) WithPostOfficeBranchName(postOfficeBranchName string) Option {
	return optionFunc(func(o *options) { o.query["post_office_branch_name"] = postOfficeBranchName })
}

// WithExpressSendTime express_send_time获取 快递寄送时间
func (obj *_LgOrderMgr) WithExpressSendTime(expressSendTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["express_send_time"] = expressSendTime })
}

// WithExpressReceiptTime express_receipt_time获取 快递签收时间
func (obj *_LgOrderMgr) WithExpressReceiptTime(expressReceiptTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["express_receipt_time"] = expressReceiptTime })
}

// WithFlightDeparturesTime flight_departures_time获取 航班起飞时间
func (obj *_LgOrderMgr) WithFlightDeparturesTime(flightDeparturesTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["flight_departures_time"] = flightDeparturesTime })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithDeliveredTime delivered_time获取 签收时间
func (obj *_LgOrderMgr) WithDeliveredTime(deliveredTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["delivered_time"] = deliveredTime })
}

// WithFlightArrivalTime flight_arrival_time获取 航班落地时间
func (obj *_LgOrderMgr) WithFlightArrivalTime(flightArrivalTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["flight_arrival_time"] = flightArrivalTime })
}

// WithExtra extra获取 额外信息
func (obj *_LgOrderMgr) WithExtra(extra datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithChannelRemarks channel_remarks获取 渠道备注
func (obj *_LgOrderMgr) WithChannelRemarks(channelRemarks string) Option {
	return optionFunc(func(o *options) { o.query["channel_remarks"] = channelRemarks })
}

// WithFinanceRemarks finance_remarks获取 财务备注
func (obj *_LgOrderMgr) WithFinanceRemarks(financeRemarks string) Option {
	return optionFunc(func(o *options) { o.query["finance_remarks"] = financeRemarks })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderMgr) GetByOption(opts ...Option) (result LgOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderMgr) GetByOptions(opts ...Option) (results []*LgOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where(options.query)
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
func (obj *_LgOrderMgr) GetFromID(id int) (result LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgOrderMgr) GetBatchFromID(ids []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderNumber 通过order_number获取内容 订单号
func (obj *_LgOrderMgr) GetFromOrderNumber(orderNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_number` = ?", orderNumber).Find(&results).Error

	return
}

// GetBatchFromOrderNumber 批量查找 订单号
func (obj *_LgOrderMgr) GetBatchFromOrderNumber(orderNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_number` IN (?)", orderNumbers).Find(&results).Error

	return
}

// GetFromReferenceNumber 通过reference_number获取内容 参考号
func (obj *_LgOrderMgr) GetFromReferenceNumber(referenceNumber string) (result LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`reference_number` = ?", referenceNumber).First(&result).Error

	return
}

// GetBatchFromReferenceNumber 批量查找 参考号
func (obj *_LgOrderMgr) GetBatchFromReferenceNumber(referenceNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`reference_number` IN (?)", referenceNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumber 通过logistics_number获取内容 物流单号
func (obj *_LgOrderMgr) GetFromLogisticsNumber(logisticsNumber string) (result LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`logistics_number` = ?", logisticsNumber).First(&result).Error

	return
}

// GetBatchFromLogisticsNumber 批量查找 物流单号
func (obj *_LgOrderMgr) GetBatchFromLogisticsNumber(logisticsNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`logistics_number` IN (?)", logisticsNumbers).Find(&results).Error

	return
}

// GetFromLogisticsNumberFinal 通过logistics_number_final获取内容 最终物流单号
func (obj *_LgOrderMgr) GetFromLogisticsNumberFinal(logisticsNumberFinal string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// GetBatchFromLogisticsNumberFinal 批量查找 最终物流单号
func (obj *_LgOrderMgr) GetBatchFromLogisticsNumberFinal(logisticsNumberFinals []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`logistics_number_final` IN (?)", logisticsNumberFinals).Find(&results).Error

	return
}

// GetFromExpressNum 通过express_num获取内容 订单快递单号
func (obj *_LgOrderMgr) GetFromExpressNum(expressNum string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`express_num` = ?", expressNum).Find(&results).Error

	return
}

// GetBatchFromExpressNum 批量查找 订单快递单号
func (obj *_LgOrderMgr) GetBatchFromExpressNum(expressNums []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`express_num` IN (?)", expressNums).Find(&results).Error

	return
}

// GetFromPlatformNumber 通过platform_number获取内容 平台订单
func (obj *_LgOrderMgr) GetFromPlatformNumber(platformNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_number` = ?", platformNumber).Find(&results).Error

	return
}

// GetBatchFromPlatformNumber 批量查找 平台订单
func (obj *_LgOrderMgr) GetBatchFromPlatformNumber(platformNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_number` IN (?)", platformNumbers).Find(&results).Error

	return
}

// GetFromProviderOrderID 通过provider_order_id获取内容 服务商系统订单ID
func (obj *_LgOrderMgr) GetFromProviderOrderID(providerOrderID string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_order_id` = ?", providerOrderID).Find(&results).Error

	return
}

// GetBatchFromProviderOrderID 批量查找 服务商系统订单ID
func (obj *_LgOrderMgr) GetBatchFromProviderOrderID(providerOrderIDs []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_order_id` IN (?)", providerOrderIDs).Find(&results).Error

	return
}

// GetFromLabelURL 通过label_url获取内容 面单地址
func (obj *_LgOrderMgr) GetFromLabelURL(labelURL string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`label_url` = ?", labelURL).Find(&results).Error

	return
}

// GetBatchFromLabelURL 批量查找 面单地址
func (obj *_LgOrderMgr) GetBatchFromLabelURL(labelURLs []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`label_url` IN (?)", labelURLs).Find(&results).Error

	return
}

// GetFromLabelURLFinal 通过label_url_final获取内容 最终面单地址
func (obj *_LgOrderMgr) GetFromLabelURLFinal(labelURLFinal string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`label_url_final` = ?", labelURLFinal).Find(&results).Error

	return
}

// GetBatchFromLabelURLFinal 批量查找 最终面单地址
func (obj *_LgOrderMgr) GetBatchFromLabelURLFinal(labelURLFinals []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`label_url_final` IN (?)", labelURLFinals).Find(&results).Error

	return
}

// GetFromOnlineNumber 通过online_number获取内容 上网单号(17单号)
func (obj *_LgOrderMgr) GetFromOnlineNumber(onlineNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`online_number` = ?", onlineNumber).Find(&results).Error

	return
}

// GetBatchFromOnlineNumber 批量查找 上网单号(17单号)
func (obj *_LgOrderMgr) GetBatchFromOnlineNumber(onlineNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`online_number` IN (?)", onlineNumbers).Find(&results).Error

	return
}

// GetFromBagNumber 通过bag_number获取内容 袋号
func (obj *_LgOrderMgr) GetFromBagNumber(bagNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`bag_number` = ?", bagNumber).Find(&results).Error

	return
}

// GetBatchFromBagNumber 批量查找 袋号
func (obj *_LgOrderMgr) GetBatchFromBagNumber(bagNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`bag_number` IN (?)", bagNumbers).Find(&results).Error

	return
}

// GetFromTotalListNo 通过total_list_no获取内容 总单号
func (obj *_LgOrderMgr) GetFromTotalListNo(totalListNo string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// GetBatchFromTotalListNo 批量查找 总单号
func (obj *_LgOrderMgr) GetBatchFromTotalListNo(totalListNos []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`total_list_no` IN (?)", totalListNos).Find(&results).Error

	return
}

// GetFromAirBillNumber 通过air_bill_number获取内容 航空提单号
func (obj *_LgOrderMgr) GetFromAirBillNumber(airBillNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`air_bill_number` = ?", airBillNumber).Find(&results).Error

	return
}

// GetBatchFromAirBillNumber 批量查找 航空提单号
func (obj *_LgOrderMgr) GetBatchFromAirBillNumber(airBillNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`air_bill_number` IN (?)", airBillNumbers).Find(&results).Error

	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgOrderMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgOrderMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error

	return
}

// GetFromCustomerChannelName 通过customer_channel_name获取内容 客户渠道名称
func (obj *_LgOrderMgr) GetFromCustomerChannelName(customerChannelName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelName 批量查找 客户渠道名称
func (obj *_LgOrderMgr) GetBatchFromCustomerChannelName(customerChannelNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_name` IN (?)", customerChannelNames).Find(&results).Error

	return
}

// GetFromProviderCode 通过provider_code获取内容 服务商系统code
func (obj *_LgOrderMgr) GetFromProviderCode(providerCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_code` = ?", providerCode).Find(&results).Error

	return
}

// GetBatchFromProviderCode 批量查找 服务商系统code
func (obj *_LgOrderMgr) GetBatchFromProviderCode(providerCodes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_code` IN (?)", providerCodes).Find(&results).Error

	return
}

// GetFromProviderChannelCode 通过provider_channel_code获取内容 服务商系统渠道Code
func (obj *_LgOrderMgr) GetFromProviderChannelCode(providerChannelCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error

	return
}

// GetBatchFromProviderChannelCode 批量查找 服务商系统渠道Code
func (obj *_LgOrderMgr) GetBatchFromProviderChannelCode(providerChannelCodes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_channel_code` IN (?)", providerChannelCodes).Find(&results).Error

	return
}

// GetFromProviderChannelName 通过provider_Channel_Name获取内容 服务商渠道名称
func (obj *_LgOrderMgr) GetFromProviderChannelName(providerChannelName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_Channel_Name` = ?", providerChannelName).Find(&results).Error

	return
}

// GetBatchFromProviderChannelName 批量查找 服务商渠道名称
func (obj *_LgOrderMgr) GetBatchFromProviderChannelName(providerChannelNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_Channel_Name` IN (?)", providerChannelNames).Find(&results).Error

	return
}

// GetFromReceiveCountry 通过receive_country获取内容 收件人国家
func (obj *_LgOrderMgr) GetFromReceiveCountry(receiveCountry string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// GetBatchFromReceiveCountry 批量查找 收件人国家
func (obj *_LgOrderMgr) GetBatchFromReceiveCountry(receiveCountrys []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_country` IN (?)", receiveCountrys).Find(&results).Error

	return
}

// GetFromReceiveName 通过receive_name获取内容 收件人姓名
func (obj *_LgOrderMgr) GetFromReceiveName(receiveName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_name` = ?", receiveName).Find(&results).Error

	return
}

// GetBatchFromReceiveName 批量查找 收件人姓名
func (obj *_LgOrderMgr) GetBatchFromReceiveName(receiveNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_name` IN (?)", receiveNames).Find(&results).Error

	return
}

// GetFromReceivePhone 通过receive_phone获取内容 收件人电话
func (obj *_LgOrderMgr) GetFromReceivePhone(receivePhone string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_phone` = ?", receivePhone).Find(&results).Error

	return
}

// GetBatchFromReceivePhone 批量查找 收件人电话
func (obj *_LgOrderMgr) GetBatchFromReceivePhone(receivePhones []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_phone` IN (?)", receivePhones).Find(&results).Error

	return
}

// GetFromReceiveMobile 通过receive_mobile获取内容 收件人手机
func (obj *_LgOrderMgr) GetFromReceiveMobile(receiveMobile string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_mobile` = ?", receiveMobile).Find(&results).Error

	return
}

// GetBatchFromReceiveMobile 批量查找 收件人手机
func (obj *_LgOrderMgr) GetBatchFromReceiveMobile(receiveMobiles []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_mobile` IN (?)", receiveMobiles).Find(&results).Error

	return
}

// GetFromReceiveTax 通过receive_tax获取内容 收件人税号
func (obj *_LgOrderMgr) GetFromReceiveTax(receiveTax string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_tax` = ?", receiveTax).Find(&results).Error

	return
}

// GetBatchFromReceiveTax 批量查找 收件人税号
func (obj *_LgOrderMgr) GetBatchFromReceiveTax(receiveTaxs []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_tax` IN (?)", receiveTaxs).Find(&results).Error

	return
}

// GetFromReceivePassport 通过receive_passport获取内容 收件人护照号
func (obj *_LgOrderMgr) GetFromReceivePassport(receivePassport string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_passport` = ?", receivePassport).Find(&results).Error

	return
}

// GetBatchFromReceivePassport 批量查找 收件人护照号
func (obj *_LgOrderMgr) GetBatchFromReceivePassport(receivePassports []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_passport` IN (?)", receivePassports).Find(&results).Error

	return
}

// GetFromReceiveMail 通过receive_mail获取内容 收件人邮箱
func (obj *_LgOrderMgr) GetFromReceiveMail(receiveMail string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_mail` = ?", receiveMail).Find(&results).Error

	return
}

// GetBatchFromReceiveMail 批量查找 收件人邮箱
func (obj *_LgOrderMgr) GetBatchFromReceiveMail(receiveMails []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_mail` IN (?)", receiveMails).Find(&results).Error

	return
}

// GetFromReceiveCompany 通过receive_company获取内容 收件人公司
func (obj *_LgOrderMgr) GetFromReceiveCompany(receiveCompany string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_company` = ?", receiveCompany).Find(&results).Error

	return
}

// GetBatchFromReceiveCompany 批量查找 收件人公司
func (obj *_LgOrderMgr) GetBatchFromReceiveCompany(receiveCompanys []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_company` IN (?)", receiveCompanys).Find(&results).Error

	return
}

// GetFromReceiveZipcode 通过receive_zipcode获取内容 收件人邮编
func (obj *_LgOrderMgr) GetFromReceiveZipcode(receiveZipcode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_zipcode` = ?", receiveZipcode).Find(&results).Error

	return
}

// GetBatchFromReceiveZipcode 批量查找 收件人邮编
func (obj *_LgOrderMgr) GetBatchFromReceiveZipcode(receiveZipcodes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_zipcode` IN (?)", receiveZipcodes).Find(&results).Error

	return
}

// GetFromReceiveProvince 通过receive_province获取内容 收件人省份
func (obj *_LgOrderMgr) GetFromReceiveProvince(receiveProvince string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_province` = ?", receiveProvince).Find(&results).Error

	return
}

// GetBatchFromReceiveProvince 批量查找 收件人省份
func (obj *_LgOrderMgr) GetBatchFromReceiveProvince(receiveProvinces []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_province` IN (?)", receiveProvinces).Find(&results).Error

	return
}

// GetFromReceiveHouseNumber 通过receive_house_number获取内容 收件人门牌号
func (obj *_LgOrderMgr) GetFromReceiveHouseNumber(receiveHouseNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_house_number` = ?", receiveHouseNumber).Find(&results).Error

	return
}

// GetBatchFromReceiveHouseNumber 批量查找 收件人门牌号
func (obj *_LgOrderMgr) GetBatchFromReceiveHouseNumber(receiveHouseNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_house_number` IN (?)", receiveHouseNumbers).Find(&results).Error

	return
}

// GetFromReceiveArea 通过receive_area获取内容 收件人区
func (obj *_LgOrderMgr) GetFromReceiveArea(receiveArea string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_area` = ?", receiveArea).Find(&results).Error

	return
}

// GetBatchFromReceiveArea 批量查找 收件人区
func (obj *_LgOrderMgr) GetBatchFromReceiveArea(receiveAreas []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_area` IN (?)", receiveAreas).Find(&results).Error

	return
}

// GetFromReceiveStreet 通过receive_street获取内容 收件人街道
func (obj *_LgOrderMgr) GetFromReceiveStreet(receiveStreet string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_street` = ?", receiveStreet).Find(&results).Error

	return
}

// GetBatchFromReceiveStreet 批量查找 收件人街道
func (obj *_LgOrderMgr) GetBatchFromReceiveStreet(receiveStreets []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_street` IN (?)", receiveStreets).Find(&results).Error

	return
}

// GetFromReceiveCity 通过receive_city获取内容 收件人城市
func (obj *_LgOrderMgr) GetFromReceiveCity(receiveCity string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_city` = ?", receiveCity).Find(&results).Error

	return
}

// GetBatchFromReceiveCity 批量查找 收件人城市
func (obj *_LgOrderMgr) GetBatchFromReceiveCity(receiveCitys []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_city` IN (?)", receiveCitys).Find(&results).Error

	return
}

// GetFromReceiveAddress1 通过receive_address1获取内容 收件人地址1
func (obj *_LgOrderMgr) GetFromReceiveAddress1(receiveAddress1 string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_address1` = ?", receiveAddress1).Find(&results).Error

	return
}

// GetBatchFromReceiveAddress1 批量查找 收件人地址1
func (obj *_LgOrderMgr) GetBatchFromReceiveAddress1(receiveAddress1s []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_address1` IN (?)", receiveAddress1s).Find(&results).Error

	return
}

// GetFromReceiveAddress2 通过receive_address2获取内容 收件人地址2
func (obj *_LgOrderMgr) GetFromReceiveAddress2(receiveAddress2 string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_address2` = ?", receiveAddress2).Find(&results).Error

	return
}

// GetBatchFromReceiveAddress2 批量查找 收件人地址2
func (obj *_LgOrderMgr) GetBatchFromReceiveAddress2(receiveAddress2s []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_address2` IN (?)", receiveAddress2s).Find(&results).Error

	return
}

// GetFromSenderCountry 通过sender_country获取内容 发件人国家
func (obj *_LgOrderMgr) GetFromSenderCountry(senderCountry string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_country` = ?", senderCountry).Find(&results).Error

	return
}

// GetBatchFromSenderCountry 批量查找 发件人国家
func (obj *_LgOrderMgr) GetBatchFromSenderCountry(senderCountrys []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_country` IN (?)", senderCountrys).Find(&results).Error

	return
}

// GetFromSenderName 通过sender_name获取内容 发件人姓名
func (obj *_LgOrderMgr) GetFromSenderName(senderName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_name` = ?", senderName).Find(&results).Error

	return
}

// GetBatchFromSenderName 批量查找 发件人姓名
func (obj *_LgOrderMgr) GetBatchFromSenderName(senderNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_name` IN (?)", senderNames).Find(&results).Error

	return
}

// GetFromSenderPhone 通过sender_phone获取内容 发件人电话
func (obj *_LgOrderMgr) GetFromSenderPhone(senderPhone string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_phone` = ?", senderPhone).Find(&results).Error

	return
}

// GetBatchFromSenderPhone 批量查找 发件人电话
func (obj *_LgOrderMgr) GetBatchFromSenderPhone(senderPhones []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_phone` IN (?)", senderPhones).Find(&results).Error

	return
}

// GetFromSenderMobile 通过sender_mobile获取内容 发件人手机
func (obj *_LgOrderMgr) GetFromSenderMobile(senderMobile string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_mobile` = ?", senderMobile).Find(&results).Error

	return
}

// GetBatchFromSenderMobile 批量查找 发件人手机
func (obj *_LgOrderMgr) GetBatchFromSenderMobile(senderMobiles []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_mobile` IN (?)", senderMobiles).Find(&results).Error

	return
}

// GetFromSenderTax 通过sender_tax获取内容 发件人税号
func (obj *_LgOrderMgr) GetFromSenderTax(senderTax string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_tax` = ?", senderTax).Find(&results).Error

	return
}

// GetBatchFromSenderTax 批量查找 发件人税号
func (obj *_LgOrderMgr) GetBatchFromSenderTax(senderTaxs []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_tax` IN (?)", senderTaxs).Find(&results).Error

	return
}

// GetFromSenderPassport 通过sender_passport获取内容 发件人护照号
func (obj *_LgOrderMgr) GetFromSenderPassport(senderPassport string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_passport` = ?", senderPassport).Find(&results).Error

	return
}

// GetBatchFromSenderPassport 批量查找 发件人护照号
func (obj *_LgOrderMgr) GetBatchFromSenderPassport(senderPassports []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_passport` IN (?)", senderPassports).Find(&results).Error

	return
}

// GetFromSenderMail 通过sender_mail获取内容 发件人邮箱
func (obj *_LgOrderMgr) GetFromSenderMail(senderMail string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_mail` = ?", senderMail).Find(&results).Error

	return
}

// GetBatchFromSenderMail 批量查找 发件人邮箱
func (obj *_LgOrderMgr) GetBatchFromSenderMail(senderMails []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_mail` IN (?)", senderMails).Find(&results).Error

	return
}

// GetFromSenderCompany 通过sender_company获取内容 发件人公司
func (obj *_LgOrderMgr) GetFromSenderCompany(senderCompany string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_company` = ?", senderCompany).Find(&results).Error

	return
}

// GetBatchFromSenderCompany 批量查找 发件人公司
func (obj *_LgOrderMgr) GetBatchFromSenderCompany(senderCompanys []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_company` IN (?)", senderCompanys).Find(&results).Error

	return
}

// GetFromSenderZipcode 通过sender_zipcode获取内容 发件人邮编
func (obj *_LgOrderMgr) GetFromSenderZipcode(senderZipcode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_zipcode` = ?", senderZipcode).Find(&results).Error

	return
}

// GetBatchFromSenderZipcode 批量查找 发件人邮编
func (obj *_LgOrderMgr) GetBatchFromSenderZipcode(senderZipcodes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_zipcode` IN (?)", senderZipcodes).Find(&results).Error

	return
}

// GetFromSenderProvince 通过sender_province获取内容 发件人省份
func (obj *_LgOrderMgr) GetFromSenderProvince(senderProvince string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_province` = ?", senderProvince).Find(&results).Error

	return
}

// GetBatchFromSenderProvince 批量查找 发件人省份
func (obj *_LgOrderMgr) GetBatchFromSenderProvince(senderProvinces []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_province` IN (?)", senderProvinces).Find(&results).Error

	return
}

// GetFromSenderCity 通过sender_city获取内容 发件人城市
func (obj *_LgOrderMgr) GetFromSenderCity(senderCity string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_city` = ?", senderCity).Find(&results).Error

	return
}

// GetBatchFromSenderCity 批量查找 发件人城市
func (obj *_LgOrderMgr) GetBatchFromSenderCity(senderCitys []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_city` IN (?)", senderCitys).Find(&results).Error

	return
}

// GetFromSenderArea 通过sender_area获取内容 发件人区
func (obj *_LgOrderMgr) GetFromSenderArea(senderArea string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_area` = ?", senderArea).Find(&results).Error

	return
}

// GetBatchFromSenderArea 批量查找 发件人区
func (obj *_LgOrderMgr) GetBatchFromSenderArea(senderAreas []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_area` IN (?)", senderAreas).Find(&results).Error

	return
}

// GetFromSenderStreet 通过sender_street获取内容 发件人街道
func (obj *_LgOrderMgr) GetFromSenderStreet(senderStreet string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_street` = ?", senderStreet).Find(&results).Error

	return
}

// GetBatchFromSenderStreet 批量查找 发件人街道
func (obj *_LgOrderMgr) GetBatchFromSenderStreet(senderStreets []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_street` IN (?)", senderStreets).Find(&results).Error

	return
}

// GetFromSenderHouseNumber 通过sender_house_number获取内容 发件人门牌号
func (obj *_LgOrderMgr) GetFromSenderHouseNumber(senderHouseNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_house_number` = ?", senderHouseNumber).Find(&results).Error

	return
}

// GetBatchFromSenderHouseNumber 批量查找 发件人门牌号
func (obj *_LgOrderMgr) GetBatchFromSenderHouseNumber(senderHouseNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_house_number` IN (?)", senderHouseNumbers).Find(&results).Error

	return
}

// GetFromSenderAddress 通过sender_address获取内容 发件人地址
func (obj *_LgOrderMgr) GetFromSenderAddress(senderAddress string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_address` = ?", senderAddress).Find(&results).Error

	return
}

// GetBatchFromSenderAddress 批量查找 发件人地址
func (obj *_LgOrderMgr) GetBatchFromSenderAddress(senderAddresss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_address` IN (?)", senderAddresss).Find(&results).Error

	return
}

// GetFromPackageType 通过package_type获取内容 包裹类型
func (obj *_LgOrderMgr) GetFromPackageType(packageType string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`package_type` = ?", packageType).Find(&results).Error

	return
}

// GetBatchFromPackageType 批量查找 包裹类型
func (obj *_LgOrderMgr) GetBatchFromPackageType(packageTypes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`package_type` IN (?)", packageTypes).Find(&results).Error

	return
}

// GetFromHasBattery 通过has_battery获取内容 是否带电
func (obj *_LgOrderMgr) GetFromHasBattery(hasBattery int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`has_battery` = ?", hasBattery).Find(&results).Error

	return
}

// GetBatchFromHasBattery 批量查找 是否带电
func (obj *_LgOrderMgr) GetBatchFromHasBattery(hasBatterys []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`has_battery` IN (?)", hasBatterys).Find(&results).Error

	return
}

// GetFromBatteryType 通过battery_type获取内容 电池类型
func (obj *_LgOrderMgr) GetFromBatteryType(batteryType string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`battery_type` = ?", batteryType).Find(&results).Error

	return
}

// GetBatchFromBatteryType 批量查找 电池类型
func (obj *_LgOrderMgr) GetBatchFromBatteryType(batteryTypes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`battery_type` IN (?)", batteryTypes).Find(&results).Error

	return
}

// GetFromHasBack 通过has_back获取内容 是否退回
func (obj *_LgOrderMgr) GetFromHasBack(hasBack int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`has_back` = ?", hasBack).Find(&results).Error

	return
}

// GetBatchFromHasBack 批量查找 是否退回
func (obj *_LgOrderMgr) GetBatchFromHasBack(hasBacks []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`has_back` IN (?)", hasBacks).Find(&results).Error

	return
}

// GetFromRemarks 通过remarks获取内容 备注信息
func (obj *_LgOrderMgr) GetFromRemarks(remarks string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`remarks` = ?", remarks).Find(&results).Error

	return
}

// GetBatchFromRemarks 批量查找 备注信息
func (obj *_LgOrderMgr) GetBatchFromRemarks(remarkss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`remarks` IN (?)", remarkss).Find(&results).Error

	return
}

// GetFromDistributionInformation 通过distribution_information获取内容 配货信息
func (obj *_LgOrderMgr) GetFromDistributionInformation(distributionInformation string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`distribution_information` = ?", distributionInformation).Find(&results).Error

	return
}

// GetBatchFromDistributionInformation 批量查找 配货信息
func (obj *_LgOrderMgr) GetBatchFromDistributionInformation(distributionInformations []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`distribution_information` IN (?)", distributionInformations).Find(&results).Error

	return
}

// GetFromInsurance 通过insurance获取内容 保险
func (obj *_LgOrderMgr) GetFromInsurance(insurance string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`insurance` = ?", insurance).Find(&results).Error

	return
}

// GetBatchFromInsurance 批量查找 保险
func (obj *_LgOrderMgr) GetBatchFromInsurance(insurances []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`insurance` IN (?)", insurances).Find(&results).Error

	return
}

// GetFromSalesPlatform 通过sales_platform获取内容 销售平台
func (obj *_LgOrderMgr) GetFromSalesPlatform(salesPlatform string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sales_platform` = ?", salesPlatform).Find(&results).Error

	return
}

// GetBatchFromSalesPlatform 批量查找 销售平台
func (obj *_LgOrderMgr) GetBatchFromSalesPlatform(salesPlatforms []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sales_platform` IN (?)", salesPlatforms).Find(&results).Error

	return
}

// GetFromWeight 通过weight获取内容 预报重量
func (obj *_LgOrderMgr) GetFromWeight(weight float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`weight` = ?", weight).Find(&results).Error

	return
}

// GetBatchFromWeight 批量查找 预报重量
func (obj *_LgOrderMgr) GetBatchFromWeight(weights []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`weight` IN (?)", weights).Find(&results).Error

	return
}

// GetFromQuantity 通过quantity获取内容 包裹数量
func (obj *_LgOrderMgr) GetFromQuantity(quantity int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`quantity` = ?", quantity).Find(&results).Error

	return
}

// GetBatchFromQuantity 批量查找 包裹数量
func (obj *_LgOrderMgr) GetBatchFromQuantity(quantitys []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`quantity` IN (?)", quantitys).Find(&results).Error

	return
}

// GetFromDeclaredValue 通过declared_value获取内容 申报价值
func (obj *_LgOrderMgr) GetFromDeclaredValue(declaredValue float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`declared_value` = ?", declaredValue).Find(&results).Error

	return
}

// GetBatchFromDeclaredValue 批量查找 申报价值
func (obj *_LgOrderMgr) GetBatchFromDeclaredValue(declaredValues []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`declared_value` IN (?)", declaredValues).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 订单状态
func (obj *_LgOrderMgr) GetFromStatus(status string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 订单状态
func (obj *_LgOrderMgr) GetBatchFromStatus(statuss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromTransportStatus 通过transport_status获取内容 转运状态
func (obj *_LgOrderMgr) GetFromTransportStatus(transportStatus string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`transport_status` = ?", transportStatus).Find(&results).Error

	return
}

// GetBatchFromTransportStatus 批量查找 转运状态
func (obj *_LgOrderMgr) GetBatchFromTransportStatus(transportStatuss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`transport_status` IN (?)", transportStatuss).Find(&results).Error

	return
}

// GetFromLastTrackStatus 通过last_track_status获取内容 最后一条轨迹状态
func (obj *_LgOrderMgr) GetFromLastTrackStatus(lastTrackStatus string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`last_track_status` = ?", lastTrackStatus).Find(&results).Error

	return
}

// GetBatchFromLastTrackStatus 批量查找 最后一条轨迹状态
func (obj *_LgOrderMgr) GetBatchFromLastTrackStatus(lastTrackStatuss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`last_track_status` IN (?)", lastTrackStatuss).Find(&results).Error

	return
}

// GetFromLastTrackAddress 通过last_track_address获取内容 最后一条轨迹
func (obj *_LgOrderMgr) GetFromLastTrackAddress(lastTrackAddress string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`last_track_address` = ?", lastTrackAddress).Find(&results).Error

	return
}

// GetBatchFromLastTrackAddress 批量查找 最后一条轨迹
func (obj *_LgOrderMgr) GetBatchFromLastTrackAddress(lastTrackAddresss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`last_track_address` IN (?)", lastTrackAddresss).Find(&results).Error

	return
}

// GetFromLastTrackTime 通过last_track_time获取内容 最后一条轨迹的时间
func (obj *_LgOrderMgr) GetFromLastTrackTime(lastTrackTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`last_track_time` = ?", lastTrackTime).Find(&results).Error

	return
}

// GetBatchFromLastTrackTime 批量查找 最后一条轨迹的时间
func (obj *_LgOrderMgr) GetBatchFromLastTrackTime(lastTrackTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`last_track_time` IN (?)", lastTrackTimes).Find(&results).Error

	return
}

// GetFromProblemType 通过problem_type获取内容 问题件类型
func (obj *_LgOrderMgr) GetFromProblemType(problemType string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`problem_type` = ?", problemType).Find(&results).Error

	return
}

// GetBatchFromProblemType 批量查找 问题件类型
func (obj *_LgOrderMgr) GetBatchFromProblemType(problemTypes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`problem_type` IN (?)", problemTypes).Find(&results).Error

	return
}

// GetFromProblemReason 通过problem_reason获取内容 问题件原因
func (obj *_LgOrderMgr) GetFromProblemReason(problemReason string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`problem_reason` = ?", problemReason).Find(&results).Error

	return
}

// GetBatchFromProblemReason 批量查找 问题件原因
func (obj *_LgOrderMgr) GetBatchFromProblemReason(problemReasons []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`problem_reason` IN (?)", problemReasons).Find(&results).Error

	return
}

// GetFromWeighingWeight 通过weighing_weight获取内容 称重重量
func (obj *_LgOrderMgr) GetFromWeighingWeight(weighingWeight float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`weighing_weight` = ?", weighingWeight).Find(&results).Error

	return
}

// GetBatchFromWeighingWeight 批量查找 称重重量
func (obj *_LgOrderMgr) GetBatchFromWeighingWeight(weighingWeights []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`weighing_weight` IN (?)", weighingWeights).Find(&results).Error

	return
}

// GetFromVolumeWeight 通过volume_weight获取内容 材积重量
func (obj *_LgOrderMgr) GetFromVolumeWeight(volumeWeight float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`volume_weight` = ?", volumeWeight).Find(&results).Error

	return
}

// GetBatchFromVolumeWeight 批量查找 材积重量
func (obj *_LgOrderMgr) GetBatchFromVolumeWeight(volumeWeights []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`volume_weight` IN (?)", volumeWeights).Find(&results).Error

	return
}

// GetFromChargeWeight 通过charge_weight获取内容 收费重量
func (obj *_LgOrderMgr) GetFromChargeWeight(chargeWeight float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`charge_weight` = ?", chargeWeight).Find(&results).Error

	return
}

// GetBatchFromChargeWeight 批量查找 收费重量
func (obj *_LgOrderMgr) GetBatchFromChargeWeight(chargeWeights []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`charge_weight` IN (?)", chargeWeights).Find(&results).Error

	return
}

// GetFromStandardFee 通过standard_fee获取内容 标准费用
func (obj *_LgOrderMgr) GetFromStandardFee(standardFee float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`standard_fee` = ?", standardFee).Find(&results).Error

	return
}

// GetBatchFromStandardFee 批量查找 标准费用
func (obj *_LgOrderMgr) GetBatchFromStandardFee(standardFees []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`standard_fee` IN (?)", standardFees).Find(&results).Error

	return
}

// GetFromCurrency 通过currency获取内容 币种
func (obj *_LgOrderMgr) GetFromCurrency(currency string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`currency` = ?", currency).Find(&results).Error

	return
}

// GetBatchFromCurrency 批量查找 币种
func (obj *_LgOrderMgr) GetBatchFromCurrency(currencys []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`currency` IN (?)", currencys).Find(&results).Error

	return
}

// GetFromBuyersFee 通过buyers_fee获取内容 下家费用
func (obj *_LgOrderMgr) GetFromBuyersFee(buyersFee float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`buyers_fee` = ?", buyersFee).Find(&results).Error

	return
}

// GetBatchFromBuyersFee 批量查找 下家费用
func (obj *_LgOrderMgr) GetBatchFromBuyersFee(buyersFees []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`buyers_fee` IN (?)", buyersFees).Find(&results).Error

	return
}

// GetFromBuyersWeight 通过buyers_weight获取内容 下家重量
func (obj *_LgOrderMgr) GetFromBuyersWeight(buyersWeight float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`buyers_weight` = ?", buyersWeight).Find(&results).Error

	return
}

// GetBatchFromBuyersWeight 批量查找 下家重量
func (obj *_LgOrderMgr) GetBatchFromBuyersWeight(buyersWeights []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`buyers_weight` IN (?)", buyersWeights).Find(&results).Error

	return
}

// GetFromOtherFee 通过other_fee获取内容 其他费用
func (obj *_LgOrderMgr) GetFromOtherFee(otherFee float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`other_fee` = ?", otherFee).Find(&results).Error

	return
}

// GetBatchFromOtherFee 批量查找 其他费用
func (obj *_LgOrderMgr) GetBatchFromOtherFee(otherFees []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`other_fee` IN (?)", otherFees).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 归属客户ID
func (obj *_LgOrderMgr) GetFromCustomerID(customerID int64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 归属客户ID
func (obj *_LgOrderMgr) GetBatchFromCustomerID(customerIDs []int64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 归属客户名称
func (obj *_LgOrderMgr) GetFromCustomerName(customerName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 归属客户名称
func (obj *_LgOrderMgr) GetBatchFromCustomerName(customerNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_LgOrderMgr) GetFromCustomerAlias(customerAlias string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_LgOrderMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromPrepaymentVat 通过prepayment_vat获取内容 "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
func (obj *_LgOrderMgr) GetFromPrepaymentVat(prepaymentVat string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`prepayment_vat` = ?", prepaymentVat).Find(&results).Error

	return
}

// GetBatchFromPrepaymentVat 批量查找 "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
func (obj *_LgOrderMgr) GetBatchFromPrepaymentVat(prepaymentVats []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`prepayment_vat` IN (?)", prepaymentVats).Find(&results).Error

	return
}

// GetFromVat 通过vat获取内容 vat号(英国方向必填)
func (obj *_LgOrderMgr) GetFromVat(vat string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`vat` = ?", vat).Find(&results).Error

	return
}

// GetBatchFromVat 批量查找 vat号(英国方向必填)
func (obj *_LgOrderMgr) GetBatchFromVat(vats []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`vat` IN (?)", vats).Find(&results).Error

	return
}

// GetFromReceiveCertificateType 通过receive_certificate_type获取内容 收件人证件类型
func (obj *_LgOrderMgr) GetFromReceiveCertificateType(receiveCertificateType string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_certificate_type` = ?", receiveCertificateType).Find(&results).Error

	return
}

// GetBatchFromReceiveCertificateType 批量查找 收件人证件类型
func (obj *_LgOrderMgr) GetBatchFromReceiveCertificateType(receiveCertificateTypes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_certificate_type` IN (?)", receiveCertificateTypes).Find(&results).Error

	return
}

// GetFromReceiveCertificateCode 通过receive_certificate_code获取内容 收件人证件号码
func (obj *_LgOrderMgr) GetFromReceiveCertificateCode(receiveCertificateCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_certificate_code` = ?", receiveCertificateCode).Find(&results).Error

	return
}

// GetBatchFromReceiveCertificateCode 批量查找 收件人证件号码
func (obj *_LgOrderMgr) GetBatchFromReceiveCertificateCode(receiveCertificateCodes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_certificate_code` IN (?)", receiveCertificateCodes).Find(&results).Error

	return
}

// GetFromSenderCertificateType 通过sender_certificate_type获取内容 发件人证件类型
func (obj *_LgOrderMgr) GetFromSenderCertificateType(senderCertificateType string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_certificate_type` = ?", senderCertificateType).Find(&results).Error

	return
}

// GetBatchFromSenderCertificateType 批量查找 发件人证件类型
func (obj *_LgOrderMgr) GetBatchFromSenderCertificateType(senderCertificateTypes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_certificate_type` IN (?)", senderCertificateTypes).Find(&results).Error

	return
}

// GetFromSenderCertificateCode 通过sender_certificate_code获取内容 发件人证件号码
func (obj *_LgOrderMgr) GetFromSenderCertificateCode(senderCertificateCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_certificate_code` = ?", senderCertificateCode).Find(&results).Error

	return
}

// GetBatchFromSenderCertificateCode 批量查找 发件人证件号码
func (obj *_LgOrderMgr) GetBatchFromSenderCertificateCode(senderCertificateCodes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_certificate_code` IN (?)", senderCertificateCodes).Find(&results).Error

	return
}

// GetFromServiceChannelCode 通过service_channel_code获取内容 服务商渠道代码
func (obj *_LgOrderMgr) GetFromServiceChannelCode(serviceChannelCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`service_channel_code` = ?", serviceChannelCode).Find(&results).Error

	return
}

// GetBatchFromServiceChannelCode 批量查找 服务商渠道代码
func (obj *_LgOrderMgr) GetBatchFromServiceChannelCode(serviceChannelCodes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`service_channel_code` IN (?)", serviceChannelCodes).Find(&results).Error

	return
}

// GetFromDeliveryTerms 通过delivery_terms获取内容 贸易条款-(部分渠道巴西方向必填)- DDU DDP
func (obj *_LgOrderMgr) GetFromDeliveryTerms(deliveryTerms string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`delivery_terms` = ?", deliveryTerms).Find(&results).Error

	return
}

// GetBatchFromDeliveryTerms 批量查找 贸易条款-(部分渠道巴西方向必填)- DDU DDP
func (obj *_LgOrderMgr) GetBatchFromDeliveryTerms(deliveryTermss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`delivery_terms` IN (?)", deliveryTermss).Find(&results).Error

	return
}

// GetFromReceiveCountryName 通过receive_country_name获取内容 收件人国家中文
func (obj *_LgOrderMgr) GetFromReceiveCountryName(receiveCountryName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_country_name` = ?", receiveCountryName).Find(&results).Error

	return
}

// GetBatchFromReceiveCountryName 批量查找 收件人国家中文
func (obj *_LgOrderMgr) GetBatchFromReceiveCountryName(receiveCountryNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_country_name` IN (?)", receiveCountryNames).Find(&results).Error

	return
}

// GetFromSenderCountryName 通过sender_country_name获取内容 发件国家中文
func (obj *_LgOrderMgr) GetFromSenderCountryName(senderCountryName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_country_name` = ?", senderCountryName).Find(&results).Error

	return
}

// GetBatchFromSenderCountryName 批量查找 发件国家中文
func (obj *_LgOrderMgr) GetBatchFromSenderCountryName(senderCountryNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_country_name` IN (?)", senderCountryNames).Find(&results).Error

	return
}

// GetFromBillNumber 通过bill_number获取内容 账单编号
func (obj *_LgOrderMgr) GetFromBillNumber(billNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`bill_number` = ?", billNumber).Find(&results).Error

	return
}

// GetBatchFromBillNumber 批量查找 账单编号
func (obj *_LgOrderMgr) GetBatchFromBillNumber(billNumbers []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`bill_number` IN (?)", billNumbers).Find(&results).Error

	return
}

// GetFromApproximateCost 通过approximate_cost获取内容 预报费用
func (obj *_LgOrderMgr) GetFromApproximateCost(approximateCost float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`approximate_cost` = ?", approximateCost).Find(&results).Error

	return
}

// GetBatchFromApproximateCost 批量查找 预报费用
func (obj *_LgOrderMgr) GetBatchFromApproximateCost(approximateCosts []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`approximate_cost` IN (?)", approximateCosts).Find(&results).Error

	return
}

// GetFromWeighingCost 通过weighing_cost获取内容 称重费用
func (obj *_LgOrderMgr) GetFromWeighingCost(weighingCost float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`weighing_cost` = ?", weighingCost).Find(&results).Error

	return
}

// GetBatchFromWeighingCost 批量查找 称重费用
func (obj *_LgOrderMgr) GetBatchFromWeighingCost(weighingCosts []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`weighing_cost` IN (?)", weighingCosts).Find(&results).Error

	return
}

// GetFromPrepaidAmount 通过prepaid_amount获取内容 预扣款金额
func (obj *_LgOrderMgr) GetFromPrepaidAmount(prepaidAmount float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`prepaid_amount` = ?", prepaidAmount).Find(&results).Error

	return
}

// GetBatchFromPrepaidAmount 批量查找 预扣款金额
func (obj *_LgOrderMgr) GetBatchFromPrepaidAmount(prepaidAmounts []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`prepaid_amount` IN (?)", prepaidAmounts).Find(&results).Error

	return
}

// GetFromActualCost 通过actual_cost获取内容 实收费用
func (obj *_LgOrderMgr) GetFromActualCost(actualCost float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`actual_cost` = ?", actualCost).Find(&results).Error

	return
}

// GetBatchFromActualCost 批量查找 实收费用
func (obj *_LgOrderMgr) GetBatchFromActualCost(actualCosts []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`actual_cost` IN (?)", actualCosts).Find(&results).Error

	return
}

// GetFromPrimeCost 通过prime_cost获取内容 成本
func (obj *_LgOrderMgr) GetFromPrimeCost(primeCost float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`prime_cost` = ?", primeCost).Find(&results).Error

	return
}

// GetBatchFromPrimeCost 批量查找 成本
func (obj *_LgOrderMgr) GetBatchFromPrimeCost(primeCosts []float64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`prime_cost` IN (?)", primeCosts).Find(&results).Error

	return
}

// GetFromOrderType 通过order_type获取内容 订单类型(0-预报订单，1-物流订单)
func (obj *_LgOrderMgr) GetFromOrderType(orderType int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_type` = ?", orderType).Find(&results).Error

	return
}

// GetBatchFromOrderType 批量查找 订单类型(0-预报订单，1-物流订单)
func (obj *_LgOrderMgr) GetBatchFromOrderType(orderTypes []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_type` IN (?)", orderTypes).Find(&results).Error

	return
}

// GetFromExchangeOrder 通过exchange_order获取内容 是否需要换单 0 不需要1 需要
func (obj *_LgOrderMgr) GetFromExchangeOrder(exchangeOrder []uint8) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`exchange_order` = ?", exchangeOrder).Find(&results).Error

	return
}

// GetBatchFromExchangeOrder 批量查找 是否需要换单 0 不需要1 需要
func (obj *_LgOrderMgr) GetBatchFromExchangeOrder(exchangeOrders [][]uint8) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`exchange_order` IN (?)", exchangeOrders).Find(&results).Error

	return
}

// GetFromReceiptTime 通过receipt_time获取内容 入库时间
func (obj *_LgOrderMgr) GetFromReceiptTime(receiptTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// GetBatchFromReceiptTime 批量查找 入库时间
func (obj *_LgOrderMgr) GetBatchFromReceiptTime(receiptTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receipt_time` IN (?)", receiptTimes).Find(&results).Error

	return
}

// GetFromReceiptUser 通过receipt_user获取内容 入库人
func (obj *_LgOrderMgr) GetFromReceiptUser(receiptUser int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receipt_user` = ?", receiptUser).Find(&results).Error

	return
}

// GetBatchFromReceiptUser 批量查找 入库人
func (obj *_LgOrderMgr) GetBatchFromReceiptUser(receiptUsers []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receipt_user` IN (?)", receiptUsers).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 出库时间
func (obj *_LgOrderMgr) GetFromSendTime(sendTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 出库时间
func (obj *_LgOrderMgr) GetBatchFromSendTime(sendTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromSendUser 通过send_user获取内容 出库人
func (obj *_LgOrderMgr) GetFromSendUser(sendUser int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`send_user` = ?", sendUser).Find(&results).Error

	return
}

// GetBatchFromSendUser 批量查找 出库人
func (obj *_LgOrderMgr) GetBatchFromSendUser(sendUsers []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`send_user` IN (?)", sendUsers).Find(&results).Error

	return
}

// GetFromPicOssLink 通过pic_oss_link获取内容 包裹图片oss地址
func (obj *_LgOrderMgr) GetFromPicOssLink(picOssLink string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`pic_oss_link` = ?", picOssLink).Find(&results).Error

	return
}

// GetBatchFromPicOssLink 批量查找 包裹图片oss地址
func (obj *_LgOrderMgr) GetBatchFromPicOssLink(picOssLinks []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`pic_oss_link` IN (?)", picOssLinks).Find(&results).Error

	return
}

// GetFromTest 通过test获取内容 是否为测试订单
func (obj *_LgOrderMgr) GetFromTest(test []uint8) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`test` = ?", test).Find(&results).Error

	return
}

// GetBatchFromTest 批量查找 是否为测试订单
func (obj *_LgOrderMgr) GetBatchFromTest(tests [][]uint8) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`test` IN (?)", tests).Find(&results).Error

	return
}

// GetFromOrderSource 通过order_source获取内容 订单来源：1-默认，2-平台订单
func (obj *_LgOrderMgr) GetFromOrderSource(orderSource int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_source` = ?", orderSource).Find(&results).Error

	return
}

// GetBatchFromOrderSource 批量查找 订单来源：1-默认，2-平台订单
func (obj *_LgOrderMgr) GetBatchFromOrderSource(orderSources []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_source` IN (?)", orderSources).Find(&results).Error

	return
}

// GetFromPlatformCode 通过platform_code获取内容 平台代码
func (obj *_LgOrderMgr) GetFromPlatformCode(platformCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_code` = ?", platformCode).Find(&results).Error

	return
}

// GetBatchFromPlatformCode 批量查找 平台代码
func (obj *_LgOrderMgr) GetBatchFromPlatformCode(platformCodes []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_code` IN (?)", platformCodes).Find(&results).Error

	return
}

// GetFromPlatformName 通过platform_name获取内容 平台名称
func (obj *_LgOrderMgr) GetFromPlatformName(platformName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_name` = ?", platformName).Find(&results).Error

	return
}

// GetBatchFromPlatformName 批量查找 平台名称
func (obj *_LgOrderMgr) GetBatchFromPlatformName(platformNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_name` IN (?)", platformNames).Find(&results).Error

	return
}

// GetFromPlatformSellerName 通过platform_seller_name获取内容 平台卖家名称
func (obj *_LgOrderMgr) GetFromPlatformSellerName(platformSellerName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_seller_name` = ?", platformSellerName).Find(&results).Error

	return
}

// GetBatchFromPlatformSellerName 批量查找 平台卖家名称
func (obj *_LgOrderMgr) GetBatchFromPlatformSellerName(platformSellerNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_seller_name` IN (?)", platformSellerNames).Find(&results).Error

	return
}

// GetFromPlatformOrderTime 通过platform_order_time获取内容 平台订单时间
func (obj *_LgOrderMgr) GetFromPlatformOrderTime(platformOrderTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_order_time` = ?", platformOrderTime).Find(&results).Error

	return
}

// GetBatchFromPlatformOrderTime 批量查找 平台订单时间
func (obj *_LgOrderMgr) GetBatchFromPlatformOrderTime(platformOrderTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_order_time` IN (?)", platformOrderTimes).Find(&results).Error

	return
}

// GetFromPostOfficeBranchName 通过post_office_branch_name获取内容 邮政局分行名称
func (obj *_LgOrderMgr) GetFromPostOfficeBranchName(postOfficeBranchName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`post_office_branch_name` = ?", postOfficeBranchName).Find(&results).Error

	return
}

// GetBatchFromPostOfficeBranchName 批量查找 邮政局分行名称
func (obj *_LgOrderMgr) GetBatchFromPostOfficeBranchName(postOfficeBranchNames []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`post_office_branch_name` IN (?)", postOfficeBranchNames).Find(&results).Error

	return
}

// GetFromExpressSendTime 通过express_send_time获取内容 快递寄送时间
func (obj *_LgOrderMgr) GetFromExpressSendTime(expressSendTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`express_send_time` = ?", expressSendTime).Find(&results).Error

	return
}

// GetBatchFromExpressSendTime 批量查找 快递寄送时间
func (obj *_LgOrderMgr) GetBatchFromExpressSendTime(expressSendTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`express_send_time` IN (?)", expressSendTimes).Find(&results).Error

	return
}

// GetFromExpressReceiptTime 通过express_receipt_time获取内容 快递签收时间
func (obj *_LgOrderMgr) GetFromExpressReceiptTime(expressReceiptTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`express_receipt_time` = ?", expressReceiptTime).Find(&results).Error

	return
}

// GetBatchFromExpressReceiptTime 批量查找 快递签收时间
func (obj *_LgOrderMgr) GetBatchFromExpressReceiptTime(expressReceiptTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`express_receipt_time` IN (?)", expressReceiptTimes).Find(&results).Error

	return
}

// GetFromFlightDeparturesTime 通过flight_departures_time获取内容 航班起飞时间
func (obj *_LgOrderMgr) GetFromFlightDeparturesTime(flightDeparturesTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`flight_departures_time` = ?", flightDeparturesTime).Find(&results).Error

	return
}

// GetBatchFromFlightDeparturesTime 批量查找 航班起飞时间
func (obj *_LgOrderMgr) GetBatchFromFlightDeparturesTime(flightDeparturesTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`flight_departures_time` IN (?)", flightDeparturesTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromDeliveredTime 通过delivered_time获取内容 签收时间
func (obj *_LgOrderMgr) GetFromDeliveredTime(deliveredTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`delivered_time` = ?", deliveredTime).Find(&results).Error

	return
}

// GetBatchFromDeliveredTime 批量查找 签收时间
func (obj *_LgOrderMgr) GetBatchFromDeliveredTime(deliveredTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`delivered_time` IN (?)", deliveredTimes).Find(&results).Error

	return
}

// GetFromFlightArrivalTime 通过flight_arrival_time获取内容 航班落地时间
func (obj *_LgOrderMgr) GetFromFlightArrivalTime(flightArrivalTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`flight_arrival_time` = ?", flightArrivalTime).Find(&results).Error

	return
}

// GetBatchFromFlightArrivalTime 批量查找 航班落地时间
func (obj *_LgOrderMgr) GetBatchFromFlightArrivalTime(flightArrivalTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`flight_arrival_time` IN (?)", flightArrivalTimes).Find(&results).Error

	return
}

// GetFromExtra 通过extra获取内容 额外信息
func (obj *_LgOrderMgr) GetFromExtra(extra datatypes.JSON) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`extra` = ?", extra).Find(&results).Error

	return
}

// GetBatchFromExtra 批量查找 额外信息
func (obj *_LgOrderMgr) GetBatchFromExtra(extras []datatypes.JSON) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`extra` IN (?)", extras).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderMgr) GetFromCreateUser(createUser int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderMgr) GetFromDeleted(deleted int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderMgr) GetFromVersion(version int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderMgr) GetBatchFromVersion(versions []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderMgr) GetFromUpdateUser(updateUser int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromChannelRemarks 通过channel_remarks获取内容 渠道备注
func (obj *_LgOrderMgr) GetFromChannelRemarks(channelRemarks string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`channel_remarks` = ?", channelRemarks).Find(&results).Error

	return
}

// GetBatchFromChannelRemarks 批量查找 渠道备注
func (obj *_LgOrderMgr) GetBatchFromChannelRemarks(channelRemarkss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`channel_remarks` IN (?)", channelRemarkss).Find(&results).Error

	return
}

// GetFromFinanceRemarks 通过finance_remarks获取内容 财务备注
func (obj *_LgOrderMgr) GetFromFinanceRemarks(financeRemarks string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`finance_remarks` = ?", financeRemarks).Find(&results).Error

	return
}

// GetBatchFromFinanceRemarks 批量查找 财务备注
func (obj *_LgOrderMgr) GetBatchFromFinanceRemarks(financeRemarkss []string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`finance_remarks` IN (?)", financeRemarkss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgOrderMgr) FetchByPrimaryKey(id int) (result LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexReferenceNumber primary or index 获取唯一内容
func (obj *_LgOrderMgr) FetchUniqueByIndexReferenceNumber(referenceNumber string) (result LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`reference_number` = ?", referenceNumber).First(&result).Error

	return
}

// FetchUniqueByIndexLogisticsNumber primary or index 获取唯一内容
func (obj *_LgOrderMgr) FetchUniqueByIndexLogisticsNumber(logisticsNumber string) (result LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`logistics_number` = ?", logisticsNumber).First(&result).Error

	return
}

// FetchUniqueIndexByLgOrderOrderNumberCustomerIDUIndex primary or index 获取唯一内容
func (obj *_LgOrderMgr) FetchUniqueIndexByLgOrderOrderNumberCustomerIDUIndex(orderNumber string, customerID int64) (result LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_number` = ? AND `customer_id` = ?", orderNumber, customerID).First(&result).Error

	return
}

// FetchIndexByLgOrderLogisticsNumberFinalIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderLogisticsNumberFinalIndex(logisticsNumberFinal string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`logistics_number_final` = ?", logisticsNumberFinal).Find(&results).Error

	return
}

// FetchIndexByLgOrderExpressNumIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderExpressNumIndex(expressNum string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`express_num` = ?", expressNum).Find(&results).Error

	return
}

// FetchIndexByLgOrderPlatformNumberIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderPlatformNumberIndex(platformNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_number` = ?", platformNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderBagNumberIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderBagNumberIndex(bagNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`bag_number` = ?", bagNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderTotalListNoIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderTotalListNoIndex(totalListNo string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`total_list_no` = ?", totalListNo).Find(&results).Error

	return
}

// FetchIndexByLgOrderAirBillNumberIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderAirBillNumberIndex(airBillNumber string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`air_bill_number` = ?", airBillNumber).Find(&results).Error

	return
}

// FetchIndexByLgOrderLgCustomerChannelIDFk  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderLgCustomerChannelIDFk(customerChannelID int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error

	return
}

// FetchIndexByIDxCustomerchannelidCreatetime  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByIDxCustomerchannelidCreatetime(customerChannelID int, createTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_id` = ? AND `create_time` = ?", customerChannelID, createTime).Find(&results).Error

	return
}

// FetchIndexByIDxCustomerchannelidReceipttime  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByIDxCustomerchannelidReceipttime(customerChannelID int, receiptTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_id` = ? AND `receipt_time` = ?", customerChannelID, receiptTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderCustomerChannelIDSendTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderCustomerChannelIDSendTimeIndex(customerChannelID int, sendTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_id` = ? AND `send_time` = ?", customerChannelID, sendTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderChannelNameIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderChannelNameIndex(customerChannelName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_channel_name` = ?", customerChannelName).Find(&results).Error

	return
}

// FetchIndexByLgOrderProviderCodeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderProviderCodeIndex(providerCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_code` = ?", providerCode).Find(&results).Error

	return
}

// FetchIndexByLgOrderProviderCodeCreateTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderProviderCodeCreateTimeIndex(providerCode string, createTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_code` = ? AND `create_time` = ?", providerCode, createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderProviderCodeReceiptTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderProviderCodeReceiptTimeIndex(providerCode string, receiptTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_code` = ? AND `receipt_time` = ?", providerCode, receiptTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderProviderCodeSendTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderProviderCodeSendTimeIndex(providerCode string, sendTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_code` = ? AND `send_time` = ?", providerCode, sendTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderProviderChannelCodeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderProviderChannelCodeIndex(providerChannelCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error

	return
}

// FetchIndexByLgOrderReceiveCountryIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderReceiveCountryIndex(receiveCountry string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_country` = ?", receiveCountry).Find(&results).Error

	return
}

// FetchIndexByLgOrderReceiveNameIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderReceiveNameIndex(receiveName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_name` = ?", receiveName).Find(&results).Error

	return
}

// FetchIndexByLgOrderReceivePhoneIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderReceivePhoneIndex(receivePhone string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_phone` = ?", receivePhone).Find(&results).Error

	return
}

// FetchIndexByLgOrderReceiveZipcodeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderReceiveZipcodeIndex(receiveZipcode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_zipcode` = ?", receiveZipcode).Find(&results).Error

	return
}

// FetchIndexByLgOrderReceiveProvinceIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderReceiveProvinceIndex(receiveProvince string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_province` = ?", receiveProvince).Find(&results).Error

	return
}

// FetchIndexByLgOrderReceiveCityIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderReceiveCityIndex(receiveCity string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receive_city` = ?", receiveCity).Find(&results).Error

	return
}

// FetchIndexByLgOrderSenderCountryIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderSenderCountryIndex(senderCountry string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_country` = ?", senderCountry).Find(&results).Error

	return
}

// FetchIndexByLgOrderSenderNameIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderSenderNameIndex(senderName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_name` = ?", senderName).Find(&results).Error

	return
}

// FetchIndexByLgOrderSenderPhoneIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderSenderPhoneIndex(senderPhone string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_phone` = ?", senderPhone).Find(&results).Error

	return
}

// FetchIndexByLgOrderSenderZipcodeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderSenderZipcodeIndex(senderZipcode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_zipcode` = ?", senderZipcode).Find(&results).Error

	return
}

// FetchIndexByLgOrderSenderProvinceIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderSenderProvinceIndex(senderProvince string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_province` = ?", senderProvince).Find(&results).Error

	return
}

// FetchIndexByLgOrderSenderCityIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderSenderCityIndex(senderCity string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`sender_city` = ?", senderCity).Find(&results).Error

	return
}

// FetchIndexByLgOrderStatusIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderStatusIndex(status string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// FetchIndexByLgOrderTransportStatusIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderTransportStatusIndex(transportStatus string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`transport_status` = ?", transportStatus).Find(&results).Error

	return
}

// FetchIndexByLgOrderProblemTypeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderProblemTypeIndex(problemType string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`problem_type` = ?", problemType).Find(&results).Error

	return
}

// FetchIndexByLgOrderCustomerIDIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderCustomerIDIndex(customerID int64) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// FetchIndexByIDxCustomeridCreatetime  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByIDxCustomeridCreatetime(customerID int64, createTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_id` = ? AND `create_time` = ?", customerID, createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderCustomerIDReceiptTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderCustomerIDReceiptTimeIndex(customerID int64, receiptTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_id` = ? AND `receipt_time` = ?", customerID, receiptTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderCustomerIDCreateTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderCustomerIDCreateTimeIndex(customerID int64, createTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_id` = ? AND `create_time` = ?", customerID, createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderCustomerIDSendTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderCustomerIDSendTimeIndex(customerID int64, sendTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`customer_id` = ? AND `send_time` = ?", customerID, sendTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderOrderTypeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderOrderTypeIndex(orderType int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_type` = ?", orderType).Find(&results).Error

	return
}

// FetchIndexByLgOrderOrderTypeReceiptTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderOrderTypeReceiptTimeIndex(orderType int, receiptTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_type` = ? AND `receipt_time` = ?", orderType, receiptTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderOrderTypeCreateTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderOrderTypeCreateTimeIndex(orderType int, createTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_type` = ? AND `create_time` = ?", orderType, createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderOrderTypeSendTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderOrderTypeSendTimeIndex(orderType int, sendTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`order_type` = ? AND `send_time` = ?", orderType, sendTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderExchangeOrderIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderExchangeOrderIndex(exchangeOrder []uint8) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`exchange_order` = ?", exchangeOrder).Find(&results).Error

	return
}

// FetchIndexByLgOrderReceiptTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderReceiptTimeIndex(receiptTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receipt_time` = ?", receiptTime).Find(&results).Error

	return
}

// FetchIndexByIDxPlatformcodeReceipttime  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByIDxPlatformcodeReceipttime(receiptTime time.Time, platformCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receipt_time` = ? AND `platform_code` = ?", receiptTime, platformCode).Find(&results).Error

	return
}

// FetchIndexByLgOrderReceiptUserIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderReceiptUserIndex(receiptUser int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`receipt_user` = ?", receiptUser).Find(&results).Error

	return
}

// FetchIndexByLgOrderSendTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderSendTimeIndex(sendTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderPlatformCodeSendTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderPlatformCodeSendTimeIndex(sendTime time.Time, platformCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`send_time` = ? AND `platform_code` = ?", sendTime, platformCode).Find(&results).Error

	return
}

// FetchIndexByLgOrderSendUserIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderSendUserIndex(sendUser int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`send_user` = ?", sendUser).Find(&results).Error

	return
}

// FetchIndexByLgOrderTestIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderTestIndex(test []uint8) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`test` = ?", test).Find(&results).Error

	return
}

// FetchIndexByLgOrderPlatformCodeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderPlatformCodeIndex(platformCode string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_code` = ?", platformCode).Find(&results).Error

	return
}

// FetchIndexByIDxPlatformcodeCreatetime  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByIDxPlatformcodeCreatetime(platformCode string, createTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_code` = ? AND `create_time` = ?", platformCode, createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderPlatformSellerNameIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderPlatformSellerNameIndex(platformSellerName string) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`platform_seller_name` = ?", platformSellerName).Find(&results).Error

	return
}

// FetchIndexByLgOrderCreateTimeIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderCreateTimeIndex(createTime time.Time) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByLgOrderCreateUserIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderCreateUserIndex(createUser int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// FetchIndexByLgOrderDeletedIndex  获取多个内容
func (obj *_LgOrderMgr) FetchIndexByLgOrderDeletedIndex(deleted int) (results []*LgOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrder{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}
