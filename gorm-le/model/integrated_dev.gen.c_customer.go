package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CCustomerMgr struct {
	*_BaseMgr
}

// CCustomerMgr open func
func CCustomerMgr(db *gorm.DB) *_CCustomerMgr {
	if db == nil {
		panic(fmt.Errorf("CCustomerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CCustomerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_customer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CCustomerMgr) GetTableName() string {
	return "c_customer"
}

// Reset 重置gorm会话
func (obj *_CCustomerMgr) Reset() *_CCustomerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CCustomerMgr) Get() (result CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CCustomerMgr) Gets() (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CCustomerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CCustomerMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 客户编号
func (obj *_CCustomerMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithUsername username获取 账户名称
func (obj *_CCustomerMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码/不得明文
func (obj *_CCustomerMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_CCustomerMgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithLevel level获取 等级，2:一级账户，3:一级账户下面的级别
func (obj *_CCustomerMgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithParentID parent_id获取 上级id(客户公司的上下级关系)
func (obj *_CCustomerMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithIntegral integral获取 积分
func (obj *_CCustomerMgr) WithIntegral(integral int) Option {
	return optionFunc(func(o *options) { o.query["integral"] = integral })
}

// WithSource source获取 客户来源，corporation:企业，person:个人
func (obj *_CCustomerMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithCompany company获取 公司名称
func (obj *_CCustomerMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithAddress address获取 公司(个人)地址
func (obj *_CCustomerMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithIDentityCard identity_card获取 营业执照或身份证号
func (obj *_CCustomerMgr) WithIDentityCard(identityCard string) Option {
	return optionFunc(func(o *options) { o.query["identity_card"] = identityCard })
}

// WithIDentityCardAttachment identity_card_attachment获取 证件照附件
func (obj *_CCustomerMgr) WithIDentityCardAttachment(identityCardAttachment string) Option {
	return optionFunc(func(o *options) { o.query["identity_card_attachment"] = identityCardAttachment })
}

// WithCoPrincipalName co_principal_name获取 企业负责人名字
func (obj *_CCustomerMgr) WithCoPrincipalName(coPrincipalName string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_name"] = coPrincipalName })
}

// WithCoPrincipalTel co_principal_tel获取 企业负责人电话
func (obj *_CCustomerMgr) WithCoPrincipalTel(coPrincipalTel string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_tel"] = coPrincipalTel })
}

// WithCoPrincipalPhone co_principal_phone获取 企业负责人手机
func (obj *_CCustomerMgr) WithCoPrincipalPhone(coPrincipalPhone string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_phone"] = coPrincipalPhone })
}

// WithCoPrincipalQq co_principal_qq获取 企业负责人qq
func (obj *_CCustomerMgr) WithCoPrincipalQq(coPrincipalQq string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_qq"] = coPrincipalQq })
}

// WithCoPrincipalEmail co_principal_email获取 企业负责人邮箱
func (obj *_CCustomerMgr) WithCoPrincipalEmail(coPrincipalEmail string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_email"] = coPrincipalEmail })
}

// WithLogisticsContactsName logistics_contacts_name获取 物流联系人名字
func (obj *_CCustomerMgr) WithLogisticsContactsName(logisticsContactsName string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_name"] = logisticsContactsName })
}

// WithLogisticsContactsTel logistics_contacts_tel获取 物流联系人电话
func (obj *_CCustomerMgr) WithLogisticsContactsTel(logisticsContactsTel string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_tel"] = logisticsContactsTel })
}

// WithLogisticsContactsPhone logistics_contacts_phone获取 物流联系人手机
func (obj *_CCustomerMgr) WithLogisticsContactsPhone(logisticsContactsPhone string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_phone"] = logisticsContactsPhone })
}

// WithLogisticsContactsQq logistics_contacts_qq获取 物流联系人qq
func (obj *_CCustomerMgr) WithLogisticsContactsQq(logisticsContactsQq string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_qq"] = logisticsContactsQq })
}

// WithLogisticsContactsEmail logistics_contacts_email获取 物流联系人邮箱
func (obj *_CCustomerMgr) WithLogisticsContactsEmail(logisticsContactsEmail string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_email"] = logisticsContactsEmail })
}

// WithFinanceContactsName finance_contacts_name获取 财务联系人名字
func (obj *_CCustomerMgr) WithFinanceContactsName(financeContactsName string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_name"] = financeContactsName })
}

// WithFinanceContactsTel finance_contacts_tel获取 财务联系人电话
func (obj *_CCustomerMgr) WithFinanceContactsTel(financeContactsTel string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_tel"] = financeContactsTel })
}

// WithFinanceContactsPhone finance_contacts_phone获取 财务联系人手机
func (obj *_CCustomerMgr) WithFinanceContactsPhone(financeContactsPhone string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_phone"] = financeContactsPhone })
}

// WithFinanceContactsQq finance_contacts_qq获取 财务联系人qq
func (obj *_CCustomerMgr) WithFinanceContactsQq(financeContactsQq string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_qq"] = financeContactsQq })
}

// WithFinanceContactsEmail finance_contacts_email获取 财务联系人邮箱
func (obj *_CCustomerMgr) WithFinanceContactsEmail(financeContactsEmail string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_email"] = financeContactsEmail })
}

// WithCustomServiceUserID custom_service_user_id获取 客服id(oa系统中的用户id)
func (obj *_CCustomerMgr) WithCustomServiceUserID(customServiceUserID int) Option {
	return optionFunc(func(o *options) { o.query["custom_service_user_id"] = customServiceUserID })
}

// WithLockState lock_state获取 账号锁定，0：正常，1：锁定
func (obj *_CCustomerMgr) WithLockState(lockState int) Option {
	return optionFunc(func(o *options) { o.query["lock_state"] = lockState })
}

// WithAuditState audit_state获取 账户审核，0：未审核，1：已审核
func (obj *_CCustomerMgr) WithAuditState(auditState int) Option {
	return optionFunc(func(o *options) { o.query["audit_state"] = auditState })
}

// WithAccountStatus account_status获取 账户状态 0：正常，1：欠费，2：禁用
func (obj *_CCustomerMgr) WithAccountStatus(accountStatus int) Option {
	return optionFunc(func(o *options) { o.query["account_status"] = accountStatus })
}

// WithBalance balance获取 账户余额(账面余额)
func (obj *_CCustomerMgr) WithBalance(balance float64) Option {
	return optionFunc(func(o *options) { o.query["balance"] = balance })
}

// WithInitialCreditLimit initial_credit_limit获取 信用额度
func (obj *_CCustomerMgr) WithInitialCreditLimit(initialCreditLimit float64) Option {
	return optionFunc(func(o *options) { o.query["initial_credit_limit"] = initialCreditLimit })
}

// WithTempCreditLimit temp_credit_limit获取 临时信用额度
func (obj *_CCustomerMgr) WithTempCreditLimit(tempCreditLimit float64) Option {
	return optionFunc(func(o *options) { o.query["temp_credit_limit"] = tempCreditLimit })
}

// WithTotalBalance total_balance获取 总计余额
// 总计余额是把未出货的那些一起预扣了
func (obj *_CCustomerMgr) WithTotalBalance(totalBalance float64) Option {
	return optionFunc(func(o *options) { o.query["total_balance"] = totalBalance })
}

// WithCreditBalance credit_balance获取 信用余额
func (obj *_CCustomerMgr) WithCreditBalance(creditBalance float64) Option {
	return optionFunc(func(o *options) { o.query["credit_balance"] = creditBalance })
}

// WithTempCreditExpireTime temp_credit_expire_time获取 临时信用额度过期时间
func (obj *_CCustomerMgr) WithTempCreditExpireTime(tempCreditExpireTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["temp_credit_expire_time"] = tempCreditExpireTime })
}

// WithOutCumulativeFee out_cumulative_fee获取 出库累加费用
// Note:距离 上次结算 到 当前时间 累计出库余额 变动
func (obj *_CCustomerMgr) WithOutCumulativeFee(outCumulativeFee float64) Option {
	return optionFunc(func(o *options) { o.query["out_cumulative_fee"] = outCumulativeFee })
}

// WithInputCumulativeFee input_cumulative_fee获取 入库累计费用
func (obj *_CCustomerMgr) WithInputCumulativeFee(inputCumulativeFee float64) Option {
	return optionFunc(func(o *options) { o.query["input_cumulative_fee"] = inputCumulativeFee })
}

// WithLastCollectionDate last_collection_date获取 最后收款时间
func (obj *_CCustomerMgr) WithLastCollectionDate(lastCollectionDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_collection_date"] = lastCollectionDate })
}

// WithPaymentCycle payment_cycle获取 收款周期
func (obj *_CCustomerMgr) WithPaymentCycle(paymentCycle int) Option {
	return optionFunc(func(o *options) { o.query["payment_cycle"] = paymentCycle })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CCustomerMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_CCustomerMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CCustomerMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_CCustomerMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_CCustomerMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_CCustomerMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithAuthorizationCode authorization_code获取 授权码
func (obj *_CCustomerMgr) WithAuthorizationCode(authorizationCode string) Option {
	return optionFunc(func(o *options) { o.query["authorization_code"] = authorizationCode })
}

// WithHandlerName handler_name获取 经办人名称
func (obj *_CCustomerMgr) WithHandlerName(handlerName string) Option {
	return optionFunc(func(o *options) { o.query["handler_name"] = handlerName })
}

// WithHandlerID handler_id获取 经办人id
func (obj *_CCustomerMgr) WithHandlerID(handlerID int) Option {
	return optionFunc(func(o *options) { o.query["handler_id"] = handlerID })
}

// WithSalesLink sales_link获取 销售链接
func (obj *_CCustomerMgr) WithSalesLink(salesLink string) Option {
	return optionFunc(func(o *options) { o.query["sales_link"] = salesLink })
}

// WithAttributeState attribute_state获取 长宽高是否必填
func (obj *_CCustomerMgr) WithAttributeState(attributeState int) Option {
	return optionFunc(func(o *options) { o.query["attribute_state"] = attributeState })
}

// WithPlatformCode platform_code获取 平台代码
func (obj *_CCustomerMgr) WithPlatformCode(platformCode string) Option {
	return optionFunc(func(o *options) { o.query["platform_code"] = platformCode })
}

// WithOrganizationCode organization_code获取 客户组织编码
func (obj *_CCustomerMgr) WithOrganizationCode(organizationCode string) Option {
	return optionFunc(func(o *options) { o.query["organization_code"] = organizationCode })
}

// WithOrganizationName organization_name获取 客户组织名称
func (obj *_CCustomerMgr) WithOrganizationName(organizationName string) Option {
	return optionFunc(func(o *options) { o.query["organization_name"] = organizationName })
}

// WithKingdeeID kingdee_id获取 金蝶客户内码
func (obj *_CCustomerMgr) WithKingdeeID(kingdeeID string) Option {
	return optionFunc(func(o *options) { o.query["kingdee_id"] = kingdeeID })
}

// GetByOption 功能选项模式获取
func (obj *_CCustomerMgr) GetByOption(opts ...Option) (result CCustomer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CCustomerMgr) GetByOptions(opts ...Option) (results []*CCustomer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CCustomerMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CCustomer, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_CCustomerMgr) GetFromID(id int) (result CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_CCustomerMgr) GetBatchFromID(ids []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 客户编号
func (obj *_CCustomerMgr) GetFromCode(code string) (result CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 客户编号
func (obj *_CCustomerMgr) GetBatchFromCode(codes []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 账户名称
func (obj *_CCustomerMgr) GetFromUsername(username string) (result CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`username` = ?", username).First(&result).Error

	return
}

// GetBatchFromUsername 批量查找 账户名称
func (obj *_CCustomerMgr) GetBatchFromUsername(usernames []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码/不得明文
func (obj *_CCustomerMgr) GetFromPassword(password string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码/不得明文
func (obj *_CCustomerMgr) GetBatchFromPassword(passwords []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_CCustomerMgr) GetFromCustomerAlias(customerAlias string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_CCustomerMgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 等级，2:一级账户，3:一级账户下面的级别
func (obj *_CCustomerMgr) GetFromLevel(level int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找 等级，2:一级账户，3:一级账户下面的级别
func (obj *_CCustomerMgr) GetBatchFromLevel(levels []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 上级id(客户公司的上下级关系)
func (obj *_CCustomerMgr) GetFromParentID(parentID int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 上级id(客户公司的上下级关系)
func (obj *_CCustomerMgr) GetBatchFromParentID(parentIDs []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromIntegral 通过integral获取内容 积分
func (obj *_CCustomerMgr) GetFromIntegral(integral int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`integral` = ?", integral).Find(&results).Error

	return
}

// GetBatchFromIntegral 批量查找 积分
func (obj *_CCustomerMgr) GetBatchFromIntegral(integrals []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`integral` IN (?)", integrals).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容 客户来源，corporation:企业，person:个人
func (obj *_CCustomerMgr) GetFromSource(source string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找 客户来源，corporation:企业，person:个人
func (obj *_CCustomerMgr) GetBatchFromSource(sources []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 公司名称
func (obj *_CCustomerMgr) GetFromCompany(company string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量查找 公司名称
func (obj *_CCustomerMgr) GetBatchFromCompany(companys []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`company` IN (?)", companys).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 公司(个人)地址
func (obj *_CCustomerMgr) GetFromAddress(address string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 公司(个人)地址
func (obj *_CCustomerMgr) GetBatchFromAddress(addresss []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromIDentityCard 通过identity_card获取内容 营业执照或身份证号
func (obj *_CCustomerMgr) GetFromIDentityCard(identityCard string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`identity_card` = ?", identityCard).Find(&results).Error

	return
}

// GetBatchFromIDentityCard 批量查找 营业执照或身份证号
func (obj *_CCustomerMgr) GetBatchFromIDentityCard(identityCards []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`identity_card` IN (?)", identityCards).Find(&results).Error

	return
}

// GetFromIDentityCardAttachment 通过identity_card_attachment获取内容 证件照附件
func (obj *_CCustomerMgr) GetFromIDentityCardAttachment(identityCardAttachment string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`identity_card_attachment` = ?", identityCardAttachment).Find(&results).Error

	return
}

// GetBatchFromIDentityCardAttachment 批量查找 证件照附件
func (obj *_CCustomerMgr) GetBatchFromIDentityCardAttachment(identityCardAttachments []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`identity_card_attachment` IN (?)", identityCardAttachments).Find(&results).Error

	return
}

// GetFromCoPrincipalName 通过co_principal_name获取内容 企业负责人名字
func (obj *_CCustomerMgr) GetFromCoPrincipalName(coPrincipalName string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_name` = ?", coPrincipalName).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalName 批量查找 企业负责人名字
func (obj *_CCustomerMgr) GetBatchFromCoPrincipalName(coPrincipalNames []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_name` IN (?)", coPrincipalNames).Find(&results).Error

	return
}

// GetFromCoPrincipalTel 通过co_principal_tel获取内容 企业负责人电话
func (obj *_CCustomerMgr) GetFromCoPrincipalTel(coPrincipalTel string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_tel` = ?", coPrincipalTel).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalTel 批量查找 企业负责人电话
func (obj *_CCustomerMgr) GetBatchFromCoPrincipalTel(coPrincipalTels []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_tel` IN (?)", coPrincipalTels).Find(&results).Error

	return
}

// GetFromCoPrincipalPhone 通过co_principal_phone获取内容 企业负责人手机
func (obj *_CCustomerMgr) GetFromCoPrincipalPhone(coPrincipalPhone string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_phone` = ?", coPrincipalPhone).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalPhone 批量查找 企业负责人手机
func (obj *_CCustomerMgr) GetBatchFromCoPrincipalPhone(coPrincipalPhones []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_phone` IN (?)", coPrincipalPhones).Find(&results).Error

	return
}

// GetFromCoPrincipalQq 通过co_principal_qq获取内容 企业负责人qq
func (obj *_CCustomerMgr) GetFromCoPrincipalQq(coPrincipalQq string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_qq` = ?", coPrincipalQq).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalQq 批量查找 企业负责人qq
func (obj *_CCustomerMgr) GetBatchFromCoPrincipalQq(coPrincipalQqs []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_qq` IN (?)", coPrincipalQqs).Find(&results).Error

	return
}

// GetFromCoPrincipalEmail 通过co_principal_email获取内容 企业负责人邮箱
func (obj *_CCustomerMgr) GetFromCoPrincipalEmail(coPrincipalEmail string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_email` = ?", coPrincipalEmail).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalEmail 批量查找 企业负责人邮箱
func (obj *_CCustomerMgr) GetBatchFromCoPrincipalEmail(coPrincipalEmails []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`co_principal_email` IN (?)", coPrincipalEmails).Find(&results).Error

	return
}

// GetFromLogisticsContactsName 通过logistics_contacts_name获取内容 物流联系人名字
func (obj *_CCustomerMgr) GetFromLogisticsContactsName(logisticsContactsName string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_name` = ?", logisticsContactsName).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsName 批量查找 物流联系人名字
func (obj *_CCustomerMgr) GetBatchFromLogisticsContactsName(logisticsContactsNames []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_name` IN (?)", logisticsContactsNames).Find(&results).Error

	return
}

// GetFromLogisticsContactsTel 通过logistics_contacts_tel获取内容 物流联系人电话
func (obj *_CCustomerMgr) GetFromLogisticsContactsTel(logisticsContactsTel string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_tel` = ?", logisticsContactsTel).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsTel 批量查找 物流联系人电话
func (obj *_CCustomerMgr) GetBatchFromLogisticsContactsTel(logisticsContactsTels []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_tel` IN (?)", logisticsContactsTels).Find(&results).Error

	return
}

// GetFromLogisticsContactsPhone 通过logistics_contacts_phone获取内容 物流联系人手机
func (obj *_CCustomerMgr) GetFromLogisticsContactsPhone(logisticsContactsPhone string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_phone` = ?", logisticsContactsPhone).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsPhone 批量查找 物流联系人手机
func (obj *_CCustomerMgr) GetBatchFromLogisticsContactsPhone(logisticsContactsPhones []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_phone` IN (?)", logisticsContactsPhones).Find(&results).Error

	return
}

// GetFromLogisticsContactsQq 通过logistics_contacts_qq获取内容 物流联系人qq
func (obj *_CCustomerMgr) GetFromLogisticsContactsQq(logisticsContactsQq string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_qq` = ?", logisticsContactsQq).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsQq 批量查找 物流联系人qq
func (obj *_CCustomerMgr) GetBatchFromLogisticsContactsQq(logisticsContactsQqs []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_qq` IN (?)", logisticsContactsQqs).Find(&results).Error

	return
}

// GetFromLogisticsContactsEmail 通过logistics_contacts_email获取内容 物流联系人邮箱
func (obj *_CCustomerMgr) GetFromLogisticsContactsEmail(logisticsContactsEmail string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_email` = ?", logisticsContactsEmail).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsEmail 批量查找 物流联系人邮箱
func (obj *_CCustomerMgr) GetBatchFromLogisticsContactsEmail(logisticsContactsEmails []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`logistics_contacts_email` IN (?)", logisticsContactsEmails).Find(&results).Error

	return
}

// GetFromFinanceContactsName 通过finance_contacts_name获取内容 财务联系人名字
func (obj *_CCustomerMgr) GetFromFinanceContactsName(financeContactsName string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_name` = ?", financeContactsName).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsName 批量查找 财务联系人名字
func (obj *_CCustomerMgr) GetBatchFromFinanceContactsName(financeContactsNames []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_name` IN (?)", financeContactsNames).Find(&results).Error

	return
}

// GetFromFinanceContactsTel 通过finance_contacts_tel获取内容 财务联系人电话
func (obj *_CCustomerMgr) GetFromFinanceContactsTel(financeContactsTel string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_tel` = ?", financeContactsTel).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsTel 批量查找 财务联系人电话
func (obj *_CCustomerMgr) GetBatchFromFinanceContactsTel(financeContactsTels []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_tel` IN (?)", financeContactsTels).Find(&results).Error

	return
}

// GetFromFinanceContactsPhone 通过finance_contacts_phone获取内容 财务联系人手机
func (obj *_CCustomerMgr) GetFromFinanceContactsPhone(financeContactsPhone string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_phone` = ?", financeContactsPhone).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsPhone 批量查找 财务联系人手机
func (obj *_CCustomerMgr) GetBatchFromFinanceContactsPhone(financeContactsPhones []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_phone` IN (?)", financeContactsPhones).Find(&results).Error

	return
}

// GetFromFinanceContactsQq 通过finance_contacts_qq获取内容 财务联系人qq
func (obj *_CCustomerMgr) GetFromFinanceContactsQq(financeContactsQq string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_qq` = ?", financeContactsQq).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsQq 批量查找 财务联系人qq
func (obj *_CCustomerMgr) GetBatchFromFinanceContactsQq(financeContactsQqs []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_qq` IN (?)", financeContactsQqs).Find(&results).Error

	return
}

// GetFromFinanceContactsEmail 通过finance_contacts_email获取内容 财务联系人邮箱
func (obj *_CCustomerMgr) GetFromFinanceContactsEmail(financeContactsEmail string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_email` = ?", financeContactsEmail).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsEmail 批量查找 财务联系人邮箱
func (obj *_CCustomerMgr) GetBatchFromFinanceContactsEmail(financeContactsEmails []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`finance_contacts_email` IN (?)", financeContactsEmails).Find(&results).Error

	return
}

// GetFromCustomServiceUserID 通过custom_service_user_id获取内容 客服id(oa系统中的用户id)
func (obj *_CCustomerMgr) GetFromCustomServiceUserID(customServiceUserID int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`custom_service_user_id` = ?", customServiceUserID).Find(&results).Error

	return
}

// GetBatchFromCustomServiceUserID 批量查找 客服id(oa系统中的用户id)
func (obj *_CCustomerMgr) GetBatchFromCustomServiceUserID(customServiceUserIDs []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`custom_service_user_id` IN (?)", customServiceUserIDs).Find(&results).Error

	return
}

// GetFromLockState 通过lock_state获取内容 账号锁定，0：正常，1：锁定
func (obj *_CCustomerMgr) GetFromLockState(lockState int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`lock_state` = ?", lockState).Find(&results).Error

	return
}

// GetBatchFromLockState 批量查找 账号锁定，0：正常，1：锁定
func (obj *_CCustomerMgr) GetBatchFromLockState(lockStates []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`lock_state` IN (?)", lockStates).Find(&results).Error

	return
}

// GetFromAuditState 通过audit_state获取内容 账户审核，0：未审核，1：已审核
func (obj *_CCustomerMgr) GetFromAuditState(auditState int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`audit_state` = ?", auditState).Find(&results).Error

	return
}

// GetBatchFromAuditState 批量查找 账户审核，0：未审核，1：已审核
func (obj *_CCustomerMgr) GetBatchFromAuditState(auditStates []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`audit_state` IN (?)", auditStates).Find(&results).Error

	return
}

// GetFromAccountStatus 通过account_status获取内容 账户状态 0：正常，1：欠费，2：禁用
func (obj *_CCustomerMgr) GetFromAccountStatus(accountStatus int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`account_status` = ?", accountStatus).Find(&results).Error

	return
}

// GetBatchFromAccountStatus 批量查找 账户状态 0：正常，1：欠费，2：禁用
func (obj *_CCustomerMgr) GetBatchFromAccountStatus(accountStatuss []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`account_status` IN (?)", accountStatuss).Find(&results).Error

	return
}

// GetFromBalance 通过balance获取内容 账户余额(账面余额)
func (obj *_CCustomerMgr) GetFromBalance(balance float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`balance` = ?", balance).Find(&results).Error

	return
}

// GetBatchFromBalance 批量查找 账户余额(账面余额)
func (obj *_CCustomerMgr) GetBatchFromBalance(balances []float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`balance` IN (?)", balances).Find(&results).Error

	return
}

// GetFromInitialCreditLimit 通过initial_credit_limit获取内容 信用额度
func (obj *_CCustomerMgr) GetFromInitialCreditLimit(initialCreditLimit float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`initial_credit_limit` = ?", initialCreditLimit).Find(&results).Error

	return
}

// GetBatchFromInitialCreditLimit 批量查找 信用额度
func (obj *_CCustomerMgr) GetBatchFromInitialCreditLimit(initialCreditLimits []float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`initial_credit_limit` IN (?)", initialCreditLimits).Find(&results).Error

	return
}

// GetFromTempCreditLimit 通过temp_credit_limit获取内容 临时信用额度
func (obj *_CCustomerMgr) GetFromTempCreditLimit(tempCreditLimit float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`temp_credit_limit` = ?", tempCreditLimit).Find(&results).Error

	return
}

// GetBatchFromTempCreditLimit 批量查找 临时信用额度
func (obj *_CCustomerMgr) GetBatchFromTempCreditLimit(tempCreditLimits []float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`temp_credit_limit` IN (?)", tempCreditLimits).Find(&results).Error

	return
}

// GetFromTotalBalance 通过total_balance获取内容 总计余额
// 总计余额是把未出货的那些一起预扣了
func (obj *_CCustomerMgr) GetFromTotalBalance(totalBalance float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`total_balance` = ?", totalBalance).Find(&results).Error

	return
}

// GetBatchFromTotalBalance 批量查找 总计余额
// 总计余额是把未出货的那些一起预扣了
func (obj *_CCustomerMgr) GetBatchFromTotalBalance(totalBalances []float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`total_balance` IN (?)", totalBalances).Find(&results).Error

	return
}

// GetFromCreditBalance 通过credit_balance获取内容 信用余额
func (obj *_CCustomerMgr) GetFromCreditBalance(creditBalance float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`credit_balance` = ?", creditBalance).Find(&results).Error

	return
}

// GetBatchFromCreditBalance 批量查找 信用余额
func (obj *_CCustomerMgr) GetBatchFromCreditBalance(creditBalances []float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`credit_balance` IN (?)", creditBalances).Find(&results).Error

	return
}

// GetFromTempCreditExpireTime 通过temp_credit_expire_time获取内容 临时信用额度过期时间
func (obj *_CCustomerMgr) GetFromTempCreditExpireTime(tempCreditExpireTime time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`temp_credit_expire_time` = ?", tempCreditExpireTime).Find(&results).Error

	return
}

// GetBatchFromTempCreditExpireTime 批量查找 临时信用额度过期时间
func (obj *_CCustomerMgr) GetBatchFromTempCreditExpireTime(tempCreditExpireTimes []time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`temp_credit_expire_time` IN (?)", tempCreditExpireTimes).Find(&results).Error

	return
}

// GetFromOutCumulativeFee 通过out_cumulative_fee获取内容 出库累加费用
// Note:距离 上次结算 到 当前时间 累计出库余额 变动
func (obj *_CCustomerMgr) GetFromOutCumulativeFee(outCumulativeFee float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`out_cumulative_fee` = ?", outCumulativeFee).Find(&results).Error

	return
}

// GetBatchFromOutCumulativeFee 批量查找 出库累加费用
// Note:距离 上次结算 到 当前时间 累计出库余额 变动
func (obj *_CCustomerMgr) GetBatchFromOutCumulativeFee(outCumulativeFees []float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`out_cumulative_fee` IN (?)", outCumulativeFees).Find(&results).Error

	return
}

// GetFromInputCumulativeFee 通过input_cumulative_fee获取内容 入库累计费用
func (obj *_CCustomerMgr) GetFromInputCumulativeFee(inputCumulativeFee float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`input_cumulative_fee` = ?", inputCumulativeFee).Find(&results).Error

	return
}

// GetBatchFromInputCumulativeFee 批量查找 入库累计费用
func (obj *_CCustomerMgr) GetBatchFromInputCumulativeFee(inputCumulativeFees []float64) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`input_cumulative_fee` IN (?)", inputCumulativeFees).Find(&results).Error

	return
}

// GetFromLastCollectionDate 通过last_collection_date获取内容 最后收款时间
func (obj *_CCustomerMgr) GetFromLastCollectionDate(lastCollectionDate time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`last_collection_date` = ?", lastCollectionDate).Find(&results).Error

	return
}

// GetBatchFromLastCollectionDate 批量查找 最后收款时间
func (obj *_CCustomerMgr) GetBatchFromLastCollectionDate(lastCollectionDates []time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`last_collection_date` IN (?)", lastCollectionDates).Find(&results).Error

	return
}

// GetFromPaymentCycle 通过payment_cycle获取内容 收款周期
func (obj *_CCustomerMgr) GetFromPaymentCycle(paymentCycle int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`payment_cycle` = ?", paymentCycle).Find(&results).Error

	return
}

// GetBatchFromPaymentCycle 批量查找 收款周期
func (obj *_CCustomerMgr) GetBatchFromPaymentCycle(paymentCycles []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`payment_cycle` IN (?)", paymentCycles).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CCustomerMgr) GetFromCreateTime(createTime time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CCustomerMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_CCustomerMgr) GetFromCreateUser(createUser int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_CCustomerMgr) GetBatchFromCreateUser(createUsers []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CCustomerMgr) GetFromUpdateTime(updateTime time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_CCustomerMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_CCustomerMgr) GetFromUpdateUser(updateUser int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_CCustomerMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_CCustomerMgr) GetFromVersion(version int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_CCustomerMgr) GetBatchFromVersion(versions []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_CCustomerMgr) GetFromDeleted(deleted int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_CCustomerMgr) GetBatchFromDeleted(deleteds []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromAuthorizationCode 通过authorization_code获取内容 授权码
func (obj *_CCustomerMgr) GetFromAuthorizationCode(authorizationCode string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`authorization_code` = ?", authorizationCode).Find(&results).Error

	return
}

// GetBatchFromAuthorizationCode 批量查找 授权码
func (obj *_CCustomerMgr) GetBatchFromAuthorizationCode(authorizationCodes []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`authorization_code` IN (?)", authorizationCodes).Find(&results).Error

	return
}

// GetFromHandlerName 通过handler_name获取内容 经办人名称
func (obj *_CCustomerMgr) GetFromHandlerName(handlerName string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`handler_name` = ?", handlerName).Find(&results).Error

	return
}

// GetBatchFromHandlerName 批量查找 经办人名称
func (obj *_CCustomerMgr) GetBatchFromHandlerName(handlerNames []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`handler_name` IN (?)", handlerNames).Find(&results).Error

	return
}

// GetFromHandlerID 通过handler_id获取内容 经办人id
func (obj *_CCustomerMgr) GetFromHandlerID(handlerID int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`handler_id` = ?", handlerID).Find(&results).Error

	return
}

// GetBatchFromHandlerID 批量查找 经办人id
func (obj *_CCustomerMgr) GetBatchFromHandlerID(handlerIDs []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`handler_id` IN (?)", handlerIDs).Find(&results).Error

	return
}

// GetFromSalesLink 通过sales_link获取内容 销售链接
func (obj *_CCustomerMgr) GetFromSalesLink(salesLink string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`sales_link` = ?", salesLink).Find(&results).Error

	return
}

// GetBatchFromSalesLink 批量查找 销售链接
func (obj *_CCustomerMgr) GetBatchFromSalesLink(salesLinks []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`sales_link` IN (?)", salesLinks).Find(&results).Error

	return
}

// GetFromAttributeState 通过attribute_state获取内容 长宽高是否必填
func (obj *_CCustomerMgr) GetFromAttributeState(attributeState int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`attribute_state` = ?", attributeState).Find(&results).Error

	return
}

// GetBatchFromAttributeState 批量查找 长宽高是否必填
func (obj *_CCustomerMgr) GetBatchFromAttributeState(attributeStates []int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`attribute_state` IN (?)", attributeStates).Find(&results).Error

	return
}

// GetFromPlatformCode 通过platform_code获取内容 平台代码
func (obj *_CCustomerMgr) GetFromPlatformCode(platformCode string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`platform_code` = ?", platformCode).Find(&results).Error

	return
}

// GetBatchFromPlatformCode 批量查找 平台代码
func (obj *_CCustomerMgr) GetBatchFromPlatformCode(platformCodes []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`platform_code` IN (?)", platformCodes).Find(&results).Error

	return
}

// GetFromOrganizationCode 通过organization_code获取内容 客户组织编码
func (obj *_CCustomerMgr) GetFromOrganizationCode(organizationCode string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`organization_code` = ?", organizationCode).Find(&results).Error

	return
}

// GetBatchFromOrganizationCode 批量查找 客户组织编码
func (obj *_CCustomerMgr) GetBatchFromOrganizationCode(organizationCodes []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`organization_code` IN (?)", organizationCodes).Find(&results).Error

	return
}

// GetFromOrganizationName 通过organization_name获取内容 客户组织名称
func (obj *_CCustomerMgr) GetFromOrganizationName(organizationName string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`organization_name` = ?", organizationName).Find(&results).Error

	return
}

// GetBatchFromOrganizationName 批量查找 客户组织名称
func (obj *_CCustomerMgr) GetBatchFromOrganizationName(organizationNames []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`organization_name` IN (?)", organizationNames).Find(&results).Error

	return
}

// GetFromKingdeeID 通过kingdee_id获取内容 金蝶客户内码
func (obj *_CCustomerMgr) GetFromKingdeeID(kingdeeID string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}

// GetBatchFromKingdeeID 批量查找 金蝶客户内码
func (obj *_CCustomerMgr) GetBatchFromKingdeeID(kingdeeIDs []string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`kingdee_id` IN (?)", kingdeeIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CCustomerMgr) FetchByPrimaryKey(id int) (result CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexCode primary or index 获取唯一内容
func (obj *_CCustomerMgr) FetchUniqueByIndexCode(code string) (result CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`code` = ?", code).First(&result).Error

	return
}

// FetchUniqueByIndexUsername primary or index 获取唯一内容
func (obj *_CCustomerMgr) FetchUniqueByIndexUsername(username string) (result CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`username` = ?", username).First(&result).Error

	return
}

// FetchIndexByCCustomerCustomerAliasIndex  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByCCustomerCustomerAliasIndex(customerAlias string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// FetchIndexByCCustomerCompanyIndex  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByCCustomerCompanyIndex(company string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// FetchIndexByCCustomerLockStateIndex  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByCCustomerLockStateIndex(lockState int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`lock_state` = ?", lockState).Find(&results).Error

	return
}

// FetchIndexByCCustomerAuditStateIndex  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByCCustomerAuditStateIndex(auditState int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`audit_state` = ?", auditState).Find(&results).Error

	return
}

// FetchIndexByCCustomerCreateTimeIndex  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByCCustomerCreateTimeIndex(createTime time.Time) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByCCustomerDeletedIndex  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByCCustomerDeletedIndex(deleted int) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// FetchIndexByCCustomerPlatformCodeIndex  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByCCustomerPlatformCodeIndex(platformCode string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`platform_code` = ?", platformCode).Find(&results).Error

	return
}

// FetchIndexByIndexOrganizationCode  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByIndexOrganizationCode(organizationCode string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`organization_code` = ?", organizationCode).Find(&results).Error

	return
}

// FetchIndexByIndexOrganizationName  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByIndexOrganizationName(organizationName string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`organization_name` = ?", organizationName).Find(&results).Error

	return
}

// FetchIndexByIndexKingdeeID  获取多个内容
func (obj *_CCustomerMgr) FetchIndexByIndexKingdeeID(kingdeeID string) (results []*CCustomer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomer{}).Where("`kingdee_id` = ?", kingdeeID).Find(&results).Error

	return
}
