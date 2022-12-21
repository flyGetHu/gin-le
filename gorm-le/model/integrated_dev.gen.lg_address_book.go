package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgAddressBookMgr struct {
	*_BaseMgr
}

// LgAddressBookMgr open func
func LgAddressBookMgr(db *gorm.DB) *_LgAddressBookMgr {
	if db == nil {
		panic(fmt.Errorf("LgAddressBookMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgAddressBookMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_address_book"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgAddressBookMgr) GetTableName() string {
	return "lg_address_book"
}

// Reset 重置gorm会话
func (obj *_LgAddressBookMgr) Reset() *_LgAddressBookMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgAddressBookMgr) Get() (result LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgAddressBookMgr) Gets() (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgAddressBookMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgAddressBookMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerID customer_id获取 客户id
func (obj *_LgAddressBookMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithAddressName address_name获取 地址簿名称
func (obj *_LgAddressBookMgr) WithAddressName(addressName string) Option {
	return optionFunc(func(o *options) { o.query["address_name"] = addressName })
}

// WithApplicationType application_type获取 应用类型，customer:客户，provider_channel:服务商渠道
func (obj *_LgAddressBookMgr) WithApplicationType(applicationType string) Option {
	return optionFunc(func(o *options) { o.query["application_type"] = applicationType })
}

// WithType type获取 类型 收件人/发件人
func (obj *_LgAddressBookMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithCountry country获取 国家二字码
func (obj *_LgAddressBookMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithAreaID area_id获取 地区id
func (obj *_LgAddressBookMgr) WithAreaID(areaID int) Option {
	return optionFunc(func(o *options) { o.query["area_id"] = areaID })
}

// WithProvince province获取 省份
func (obj *_LgAddressBookMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 城市
func (obj *_LgAddressBookMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithPhone phone获取 电话
func (obj *_LgAddressBookMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithName name获取 姓名
func (obj *_LgAddressBookMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithZipcode zipcode获取 邮编
func (obj *_LgAddressBookMgr) WithZipcode(zipcode string) Option {
	return optionFunc(func(o *options) { o.query["zipcode"] = zipcode })
}

// WithArea area获取 区/县
func (obj *_LgAddressBookMgr) WithArea(area string) Option {
	return optionFunc(func(o *options) { o.query["area"] = area })
}

// WithStreet street获取 街道
func (obj *_LgAddressBookMgr) WithStreet(street string) Option {
	return optionFunc(func(o *options) { o.query["street"] = street })
}

// WithHouseNumber house_number获取 门牌号
func (obj *_LgAddressBookMgr) WithHouseNumber(houseNumber string) Option {
	return optionFunc(func(o *options) { o.query["house_number"] = houseNumber })
}

// WithCertificateType certificate_type获取 证件类型
func (obj *_LgAddressBookMgr) WithCertificateType(certificateType string) Option {
	return optionFunc(func(o *options) { o.query["certificate_type"] = certificateType })
}

// WithCertificateCode certificate_code获取 证件号
func (obj *_LgAddressBookMgr) WithCertificateCode(certificateCode string) Option {
	return optionFunc(func(o *options) { o.query["certificate_code"] = certificateCode })
}

// WithMail mail获取 邮箱
func (obj *_LgAddressBookMgr) WithMail(mail string) Option {
	return optionFunc(func(o *options) { o.query["mail"] = mail })
}

// WithRemark remark获取 备注
func (obj *_LgAddressBookMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithMobile mobile获取 手机
func (obj *_LgAddressBookMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithTax tax获取 税号
func (obj *_LgAddressBookMgr) WithTax(tax string) Option {
	return optionFunc(func(o *options) { o.query["tax"] = tax })
}

// WithAddress address获取 地址
func (obj *_LgAddressBookMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithAddress2 address_2获取 地址2
func (obj *_LgAddressBookMgr) WithAddress2(address2 string) Option {
	return optionFunc(func(o *options) { o.query["address_2"] = address2 })
}

// WithCompany company获取 公司
func (obj *_LgAddressBookMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgAddressBookMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgAddressBookMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgAddressBookMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgAddressBookMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgAddressBookMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgAddressBookMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgAddressBookMgr) GetByOption(opts ...Option) (result LgAddressBook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgAddressBookMgr) GetByOptions(opts ...Option) (results []*LgAddressBook, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgAddressBookMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgAddressBook, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where(options.query)
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
func (obj *_LgAddressBookMgr) GetFromID(id int) (result LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgAddressBookMgr) GetBatchFromID(ids []int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_LgAddressBookMgr) GetFromCustomerID(customerID int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_LgAddressBookMgr) GetBatchFromCustomerID(customerIDs []int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromAddressName 通过address_name获取内容 地址簿名称
func (obj *_LgAddressBookMgr) GetFromAddressName(addressName string) (result LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`address_name` = ?", addressName).First(&result).Error

	return
}

// GetBatchFromAddressName 批量查找 地址簿名称
func (obj *_LgAddressBookMgr) GetBatchFromAddressName(addressNames []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`address_name` IN (?)", addressNames).Find(&results).Error

	return
}

// GetFromApplicationType 通过application_type获取内容 应用类型，customer:客户，provider_channel:服务商渠道
func (obj *_LgAddressBookMgr) GetFromApplicationType(applicationType string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`application_type` = ?", applicationType).Find(&results).Error

	return
}

// GetBatchFromApplicationType 批量查找 应用类型，customer:客户，provider_channel:服务商渠道
func (obj *_LgAddressBookMgr) GetBatchFromApplicationType(applicationTypes []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`application_type` IN (?)", applicationTypes).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型 收件人/发件人
func (obj *_LgAddressBookMgr) GetFromType(_type string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型 收件人/发件人
func (obj *_LgAddressBookMgr) GetBatchFromType(_types []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 国家二字码
func (obj *_LgAddressBookMgr) GetFromCountry(country string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`country` = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量查找 国家二字码
func (obj *_LgAddressBookMgr) GetBatchFromCountry(countrys []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`country` IN (?)", countrys).Find(&results).Error

	return
}

// GetFromAreaID 通过area_id获取内容 地区id
func (obj *_LgAddressBookMgr) GetFromAreaID(areaID int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`area_id` = ?", areaID).Find(&results).Error

	return
}

// GetBatchFromAreaID 批量查找 地区id
func (obj *_LgAddressBookMgr) GetBatchFromAreaID(areaIDs []int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`area_id` IN (?)", areaIDs).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 省份
func (obj *_LgAddressBookMgr) GetFromProvince(province string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 省份
func (obj *_LgAddressBookMgr) GetBatchFromProvince(provinces []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 城市
func (obj *_LgAddressBookMgr) GetFromCity(city string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 城市
func (obj *_LgAddressBookMgr) GetBatchFromCity(citys []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromPhone 通过phone获取内容 电话
func (obj *_LgAddressBookMgr) GetFromPhone(phone string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`phone` = ?", phone).Find(&results).Error

	return
}

// GetBatchFromPhone 批量查找 电话
func (obj *_LgAddressBookMgr) GetBatchFromPhone(phones []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`phone` IN (?)", phones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_LgAddressBookMgr) GetFromName(name string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_LgAddressBookMgr) GetBatchFromName(names []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromZipcode 通过zipcode获取内容 邮编
func (obj *_LgAddressBookMgr) GetFromZipcode(zipcode string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`zipcode` = ?", zipcode).Find(&results).Error

	return
}

// GetBatchFromZipcode 批量查找 邮编
func (obj *_LgAddressBookMgr) GetBatchFromZipcode(zipcodes []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`zipcode` IN (?)", zipcodes).Find(&results).Error

	return
}

// GetFromArea 通过area获取内容 区/县
func (obj *_LgAddressBookMgr) GetFromArea(area string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`area` = ?", area).Find(&results).Error

	return
}

// GetBatchFromArea 批量查找 区/县
func (obj *_LgAddressBookMgr) GetBatchFromArea(areas []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`area` IN (?)", areas).Find(&results).Error

	return
}

// GetFromStreet 通过street获取内容 街道
func (obj *_LgAddressBookMgr) GetFromStreet(street string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`street` = ?", street).Find(&results).Error

	return
}

// GetBatchFromStreet 批量查找 街道
func (obj *_LgAddressBookMgr) GetBatchFromStreet(streets []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`street` IN (?)", streets).Find(&results).Error

	return
}

// GetFromHouseNumber 通过house_number获取内容 门牌号
func (obj *_LgAddressBookMgr) GetFromHouseNumber(houseNumber string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`house_number` = ?", houseNumber).Find(&results).Error

	return
}

// GetBatchFromHouseNumber 批量查找 门牌号
func (obj *_LgAddressBookMgr) GetBatchFromHouseNumber(houseNumbers []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`house_number` IN (?)", houseNumbers).Find(&results).Error

	return
}

// GetFromCertificateType 通过certificate_type获取内容 证件类型
func (obj *_LgAddressBookMgr) GetFromCertificateType(certificateType string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`certificate_type` = ?", certificateType).Find(&results).Error

	return
}

// GetBatchFromCertificateType 批量查找 证件类型
func (obj *_LgAddressBookMgr) GetBatchFromCertificateType(certificateTypes []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`certificate_type` IN (?)", certificateTypes).Find(&results).Error

	return
}

// GetFromCertificateCode 通过certificate_code获取内容 证件号
func (obj *_LgAddressBookMgr) GetFromCertificateCode(certificateCode string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`certificate_code` = ?", certificateCode).Find(&results).Error

	return
}

// GetBatchFromCertificateCode 批量查找 证件号
func (obj *_LgAddressBookMgr) GetBatchFromCertificateCode(certificateCodes []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`certificate_code` IN (?)", certificateCodes).Find(&results).Error

	return
}

// GetFromMail 通过mail获取内容 邮箱
func (obj *_LgAddressBookMgr) GetFromMail(mail string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`mail` = ?", mail).Find(&results).Error

	return
}

// GetBatchFromMail 批量查找 邮箱
func (obj *_LgAddressBookMgr) GetBatchFromMail(mails []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`mail` IN (?)", mails).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_LgAddressBookMgr) GetFromRemark(remark string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_LgAddressBookMgr) GetBatchFromRemark(remarks []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机
func (obj *_LgAddressBookMgr) GetFromMobile(mobile string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机
func (obj *_LgAddressBookMgr) GetBatchFromMobile(mobiles []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromTax 通过tax获取内容 税号
func (obj *_LgAddressBookMgr) GetFromTax(tax string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`tax` = ?", tax).Find(&results).Error

	return
}

// GetBatchFromTax 批量查找 税号
func (obj *_LgAddressBookMgr) GetBatchFromTax(taxs []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`tax` IN (?)", taxs).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 地址
func (obj *_LgAddressBookMgr) GetFromAddress(address string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 地址
func (obj *_LgAddressBookMgr) GetBatchFromAddress(addresss []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromAddress2 通过address_2获取内容 地址2
func (obj *_LgAddressBookMgr) GetFromAddress2(address2 string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`address_2` = ?", address2).Find(&results).Error

	return
}

// GetBatchFromAddress2 批量查找 地址2
func (obj *_LgAddressBookMgr) GetBatchFromAddress2(address2s []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`address_2` IN (?)", address2s).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 公司
func (obj *_LgAddressBookMgr) GetFromCompany(company string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量查找 公司
func (obj *_LgAddressBookMgr) GetBatchFromCompany(companys []string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`company` IN (?)", companys).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgAddressBookMgr) GetFromCreateTime(createTime time.Time) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgAddressBookMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgAddressBookMgr) GetFromCreateUser(createUser int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgAddressBookMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgAddressBookMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgAddressBookMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgAddressBookMgr) GetFromUpdateUser(updateUser int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgAddressBookMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgAddressBookMgr) GetFromVersion(version int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgAddressBookMgr) GetBatchFromVersion(versions []int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgAddressBookMgr) GetFromDeleted(deleted int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgAddressBookMgr) GetBatchFromDeleted(deleteds []int) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgAddressBookMgr) FetchByPrimaryKey(id int) (result LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgAddressBookAddressNameUIndex primary or index 获取唯一内容
func (obj *_LgAddressBookMgr) FetchUniqueByLgAddressBookAddressNameUIndex(addressName string) (result LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`address_name` = ?", addressName).First(&result).Error

	return
}

// FetchIndexByLgAddressBookCustomerIDTypeCountryPhoneIndex  获取多个内容
func (obj *_LgAddressBookMgr) FetchIndexByLgAddressBookCustomerIDTypeCountryPhoneIndex(customerID int, _type string, country string, phone string) (results []*LgAddressBook, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgAddressBook{}).Where("`customer_id` = ? AND `type` = ? AND `country` = ? AND `phone` = ?", customerID, _type, country, phone).Find(&results).Error

	return
}
