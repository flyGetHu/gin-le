package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _CCustomerCopy1Mgr struct {
	*_BaseMgr
}

// CCustomerCopy1Mgr open func
func CCustomerCopy1Mgr(db *gorm.DB) *_CCustomerCopy1Mgr {
	if db == nil {
		panic(fmt.Errorf("CCustomerCopy1Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_CCustomerCopy1Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("c_customer_copy1"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_CCustomerCopy1Mgr) GetTableName() string {
	return "c_customer_copy1"
}

// Reset 重置gorm会话
func (obj *_CCustomerCopy1Mgr) Reset() *_CCustomerCopy1Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_CCustomerCopy1Mgr) Get() (result CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_CCustomerCopy1Mgr) Gets() (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_CCustomerCopy1Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_CCustomerCopy1Mgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 客户编号
func (obj *_CCustomerCopy1Mgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithUsername username获取 账户名称
func (obj *_CCustomerCopy1Mgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码/不得明文
func (obj *_CCustomerCopy1Mgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithCustomerAlias customer_alias获取 客户别名
func (obj *_CCustomerCopy1Mgr) WithCustomerAlias(customerAlias string) Option {
	return optionFunc(func(o *options) { o.query["customer_alias"] = customerAlias })
}

// WithLevel level获取 等级，2:一级账户，3:一级账户下面的级别
func (obj *_CCustomerCopy1Mgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithParentID parent_id获取 上级id(客户公司的上下级关系)
func (obj *_CCustomerCopy1Mgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithIntegral integral获取 积分
func (obj *_CCustomerCopy1Mgr) WithIntegral(integral int) Option {
	return optionFunc(func(o *options) { o.query["integral"] = integral })
}

// WithSource source获取 客户来源，corporation:企业，person:个人
func (obj *_CCustomerCopy1Mgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithCompany company获取 公司名称
func (obj *_CCustomerCopy1Mgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithAddress address获取 公司(个人)地址
func (obj *_CCustomerCopy1Mgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithIDentityCard identity_card获取 营业执照或身份证号
func (obj *_CCustomerCopy1Mgr) WithIDentityCard(identityCard string) Option {
	return optionFunc(func(o *options) { o.query["identity_card"] = identityCard })
}

// WithIDentityCardAttachment identity_card_attachment获取 证件照附件
func (obj *_CCustomerCopy1Mgr) WithIDentityCardAttachment(identityCardAttachment string) Option {
	return optionFunc(func(o *options) { o.query["identity_card_attachment"] = identityCardAttachment })
}

// WithCoPrincipalName co_principal_name获取 企业负责人名字
func (obj *_CCustomerCopy1Mgr) WithCoPrincipalName(coPrincipalName string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_name"] = coPrincipalName })
}

// WithCoPrincipalTel co_principal_tel获取 企业负责人电话
func (obj *_CCustomerCopy1Mgr) WithCoPrincipalTel(coPrincipalTel string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_tel"] = coPrincipalTel })
}

// WithCoPrincipalPhone co_principal_phone获取 企业负责人手机
func (obj *_CCustomerCopy1Mgr) WithCoPrincipalPhone(coPrincipalPhone string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_phone"] = coPrincipalPhone })
}

// WithCoPrincipalQq co_principal_qq获取 企业负责人qq
func (obj *_CCustomerCopy1Mgr) WithCoPrincipalQq(coPrincipalQq string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_qq"] = coPrincipalQq })
}

// WithCoPrincipalEmail co_principal_email获取 企业负责人邮箱
func (obj *_CCustomerCopy1Mgr) WithCoPrincipalEmail(coPrincipalEmail string) Option {
	return optionFunc(func(o *options) { o.query["co_principal_email"] = coPrincipalEmail })
}

// WithLogisticsContactsName logistics_contacts_name获取 物流联系人名字
func (obj *_CCustomerCopy1Mgr) WithLogisticsContactsName(logisticsContactsName string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_name"] = logisticsContactsName })
}

// WithLogisticsContactsTel logistics_contacts_tel获取 物流联系人电话
func (obj *_CCustomerCopy1Mgr) WithLogisticsContactsTel(logisticsContactsTel string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_tel"] = logisticsContactsTel })
}

// WithLogisticsContactsPhone logistics_contacts_phone获取 物流联系人手机
func (obj *_CCustomerCopy1Mgr) WithLogisticsContactsPhone(logisticsContactsPhone string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_phone"] = logisticsContactsPhone })
}

// WithLogisticsContactsQq logistics_contacts_qq获取 物流联系人qq
func (obj *_CCustomerCopy1Mgr) WithLogisticsContactsQq(logisticsContactsQq string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_qq"] = logisticsContactsQq })
}

// WithLogisticsContactsEmail logistics_contacts_email获取 物流联系人邮箱
func (obj *_CCustomerCopy1Mgr) WithLogisticsContactsEmail(logisticsContactsEmail string) Option {
	return optionFunc(func(o *options) { o.query["logistics_contacts_email"] = logisticsContactsEmail })
}

// WithFinanceContactsName finance_contacts_name获取 财务联系人名字
func (obj *_CCustomerCopy1Mgr) WithFinanceContactsName(financeContactsName string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_name"] = financeContactsName })
}

// WithFinanceContactsTel finance_contacts_tel获取 财务联系人电话
func (obj *_CCustomerCopy1Mgr) WithFinanceContactsTel(financeContactsTel string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_tel"] = financeContactsTel })
}

// WithFinanceContactsPhone finance_contacts_phone获取 财务联系人手机
func (obj *_CCustomerCopy1Mgr) WithFinanceContactsPhone(financeContactsPhone string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_phone"] = financeContactsPhone })
}

// WithFinanceContactsQq finance_contacts_qq获取 财务联系人qq
func (obj *_CCustomerCopy1Mgr) WithFinanceContactsQq(financeContactsQq string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_qq"] = financeContactsQq })
}

// WithFinanceContactsEmail finance_contacts_email获取 财务联系人邮箱
func (obj *_CCustomerCopy1Mgr) WithFinanceContactsEmail(financeContactsEmail string) Option {
	return optionFunc(func(o *options) { o.query["finance_contacts_email"] = financeContactsEmail })
}

// WithCustomServiceUserID custom_service_user_id获取 客服id(oa系统中的用户id)
func (obj *_CCustomerCopy1Mgr) WithCustomServiceUserID(customServiceUserID int) Option {
	return optionFunc(func(o *options) { o.query["custom_service_user_id"] = customServiceUserID })
}

// WithLockState lock_state获取 账号锁定，0：正常，1：锁定
func (obj *_CCustomerCopy1Mgr) WithLockState(lockState int) Option {
	return optionFunc(func(o *options) { o.query["lock_state"] = lockState })
}

// WithAuditState audit_state获取 账户审核，0：未审核，1：已审核
func (obj *_CCustomerCopy1Mgr) WithAuditState(auditState int) Option {
	return optionFunc(func(o *options) { o.query["audit_state"] = auditState })
}

// WithAccountStatus account_status获取 账户状态 0：正常，1：欠费，2：禁用
func (obj *_CCustomerCopy1Mgr) WithAccountStatus(accountStatus int) Option {
	return optionFunc(func(o *options) { o.query["account_status"] = accountStatus })
}

// WithBalance balance获取 账户余额
func (obj *_CCustomerCopy1Mgr) WithBalance(balance float64) Option {
	return optionFunc(func(o *options) { o.query["balance"] = balance })
}

// WithInitialCreditLimit initial_credit_limit获取 信用额度
func (obj *_CCustomerCopy1Mgr) WithInitialCreditLimit(initialCreditLimit float64) Option {
	return optionFunc(func(o *options) { o.query["initial_credit_limit"] = initialCreditLimit })
}

// WithOutCumulativeFee out_cumulative_fee获取 出库累加费用
// Note:距离 上次结算 到 当前时间 累计出库余额 变动
func (obj *_CCustomerCopy1Mgr) WithOutCumulativeFee(outCumulativeFee float64) Option {
	return optionFunc(func(o *options) { o.query["out_cumulative_fee"] = outCumulativeFee })
}

// WithLastCollectionDate last_collection_date获取 最后收款时间
func (obj *_CCustomerCopy1Mgr) WithLastCollectionDate(lastCollectionDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_collection_date"] = lastCollectionDate })
}

// WithPaymentCycle payment_cycle获取 收款周期
func (obj *_CCustomerCopy1Mgr) WithPaymentCycle(paymentCycle int) Option {
	return optionFunc(func(o *options) { o.query["payment_cycle"] = paymentCycle })
}

// WithCreateTime create_time获取 创建时间
func (obj *_CCustomerCopy1Mgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_CCustomerCopy1Mgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_CCustomerCopy1Mgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_CCustomerCopy1Mgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_CCustomerCopy1Mgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_CCustomerCopy1Mgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithAuthorizationCode authorization_code获取 授权码
func (obj *_CCustomerCopy1Mgr) WithAuthorizationCode(authorizationCode string) Option {
	return optionFunc(func(o *options) { o.query["authorization_code"] = authorizationCode })
}

// WithHandlerName handler_name获取 经办人名称
func (obj *_CCustomerCopy1Mgr) WithHandlerName(handlerName string) Option {
	return optionFunc(func(o *options) { o.query["handler_name"] = handlerName })
}

// WithHandlerID handler_id获取 经办人id
func (obj *_CCustomerCopy1Mgr) WithHandlerID(handlerID int) Option {
	return optionFunc(func(o *options) { o.query["handler_id"] = handlerID })
}

// GetByOption 功能选项模式获取
func (obj *_CCustomerCopy1Mgr) GetByOption(opts ...Option) (result CCustomerCopy1, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_CCustomerCopy1Mgr) GetByOptions(opts ...Option) (results []*CCustomerCopy1, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_CCustomerCopy1Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]CCustomerCopy1, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where(options.query)
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
func (obj *_CCustomerCopy1Mgr) GetFromID(id int) (result CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_CCustomerCopy1Mgr) GetBatchFromID(ids []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 客户编号
func (obj *_CCustomerCopy1Mgr) GetFromCode(code string) (result CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 客户编号
func (obj *_CCustomerCopy1Mgr) GetBatchFromCode(codes []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 账户名称
func (obj *_CCustomerCopy1Mgr) GetFromUsername(username string) (result CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`username` = ?", username).First(&result).Error

	return
}

// GetBatchFromUsername 批量查找 账户名称
func (obj *_CCustomerCopy1Mgr) GetBatchFromUsername(usernames []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码/不得明文
func (obj *_CCustomerCopy1Mgr) GetFromPassword(password string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码/不得明文
func (obj *_CCustomerCopy1Mgr) GetBatchFromPassword(passwords []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromCustomerAlias 通过customer_alias获取内容 客户别名
func (obj *_CCustomerCopy1Mgr) GetFromCustomerAlias(customerAlias string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// GetBatchFromCustomerAlias 批量查找 客户别名
func (obj *_CCustomerCopy1Mgr) GetBatchFromCustomerAlias(customerAliass []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`customer_alias` IN (?)", customerAliass).Find(&results).Error

	return
}

// GetFromLevel 通过level获取内容 等级，2:一级账户，3:一级账户下面的级别
func (obj *_CCustomerCopy1Mgr) GetFromLevel(level int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`level` = ?", level).Find(&results).Error

	return
}

// GetBatchFromLevel 批量查找 等级，2:一级账户，3:一级账户下面的级别
func (obj *_CCustomerCopy1Mgr) GetBatchFromLevel(levels []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`level` IN (?)", levels).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 上级id(客户公司的上下级关系)
func (obj *_CCustomerCopy1Mgr) GetFromParentID(parentID int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 上级id(客户公司的上下级关系)
func (obj *_CCustomerCopy1Mgr) GetBatchFromParentID(parentIDs []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromIntegral 通过integral获取内容 积分
func (obj *_CCustomerCopy1Mgr) GetFromIntegral(integral int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`integral` = ?", integral).Find(&results).Error

	return
}

// GetBatchFromIntegral 批量查找 积分
func (obj *_CCustomerCopy1Mgr) GetBatchFromIntegral(integrals []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`integral` IN (?)", integrals).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容 客户来源，corporation:企业，person:个人
func (obj *_CCustomerCopy1Mgr) GetFromSource(source string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找 客户来源，corporation:企业，person:个人
func (obj *_CCustomerCopy1Mgr) GetBatchFromSource(sources []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 公司名称
func (obj *_CCustomerCopy1Mgr) GetFromCompany(company string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量查找 公司名称
func (obj *_CCustomerCopy1Mgr) GetBatchFromCompany(companys []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`company` IN (?)", companys).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 公司(个人)地址
func (obj *_CCustomerCopy1Mgr) GetFromAddress(address string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 公司(个人)地址
func (obj *_CCustomerCopy1Mgr) GetBatchFromAddress(addresss []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromIDentityCard 通过identity_card获取内容 营业执照或身份证号
func (obj *_CCustomerCopy1Mgr) GetFromIDentityCard(identityCard string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`identity_card` = ?", identityCard).Find(&results).Error

	return
}

// GetBatchFromIDentityCard 批量查找 营业执照或身份证号
func (obj *_CCustomerCopy1Mgr) GetBatchFromIDentityCard(identityCards []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`identity_card` IN (?)", identityCards).Find(&results).Error

	return
}

// GetFromIDentityCardAttachment 通过identity_card_attachment获取内容 证件照附件
func (obj *_CCustomerCopy1Mgr) GetFromIDentityCardAttachment(identityCardAttachment string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`identity_card_attachment` = ?", identityCardAttachment).Find(&results).Error

	return
}

// GetBatchFromIDentityCardAttachment 批量查找 证件照附件
func (obj *_CCustomerCopy1Mgr) GetBatchFromIDentityCardAttachment(identityCardAttachments []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`identity_card_attachment` IN (?)", identityCardAttachments).Find(&results).Error

	return
}

// GetFromCoPrincipalName 通过co_principal_name获取内容 企业负责人名字
func (obj *_CCustomerCopy1Mgr) GetFromCoPrincipalName(coPrincipalName string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_name` = ?", coPrincipalName).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalName 批量查找 企业负责人名字
func (obj *_CCustomerCopy1Mgr) GetBatchFromCoPrincipalName(coPrincipalNames []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_name` IN (?)", coPrincipalNames).Find(&results).Error

	return
}

// GetFromCoPrincipalTel 通过co_principal_tel获取内容 企业负责人电话
func (obj *_CCustomerCopy1Mgr) GetFromCoPrincipalTel(coPrincipalTel string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_tel` = ?", coPrincipalTel).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalTel 批量查找 企业负责人电话
func (obj *_CCustomerCopy1Mgr) GetBatchFromCoPrincipalTel(coPrincipalTels []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_tel` IN (?)", coPrincipalTels).Find(&results).Error

	return
}

// GetFromCoPrincipalPhone 通过co_principal_phone获取内容 企业负责人手机
func (obj *_CCustomerCopy1Mgr) GetFromCoPrincipalPhone(coPrincipalPhone string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_phone` = ?", coPrincipalPhone).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalPhone 批量查找 企业负责人手机
func (obj *_CCustomerCopy1Mgr) GetBatchFromCoPrincipalPhone(coPrincipalPhones []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_phone` IN (?)", coPrincipalPhones).Find(&results).Error

	return
}

// GetFromCoPrincipalQq 通过co_principal_qq获取内容 企业负责人qq
func (obj *_CCustomerCopy1Mgr) GetFromCoPrincipalQq(coPrincipalQq string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_qq` = ?", coPrincipalQq).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalQq 批量查找 企业负责人qq
func (obj *_CCustomerCopy1Mgr) GetBatchFromCoPrincipalQq(coPrincipalQqs []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_qq` IN (?)", coPrincipalQqs).Find(&results).Error

	return
}

// GetFromCoPrincipalEmail 通过co_principal_email获取内容 企业负责人邮箱
func (obj *_CCustomerCopy1Mgr) GetFromCoPrincipalEmail(coPrincipalEmail string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_email` = ?", coPrincipalEmail).Find(&results).Error

	return
}

// GetBatchFromCoPrincipalEmail 批量查找 企业负责人邮箱
func (obj *_CCustomerCopy1Mgr) GetBatchFromCoPrincipalEmail(coPrincipalEmails []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`co_principal_email` IN (?)", coPrincipalEmails).Find(&results).Error

	return
}

// GetFromLogisticsContactsName 通过logistics_contacts_name获取内容 物流联系人名字
func (obj *_CCustomerCopy1Mgr) GetFromLogisticsContactsName(logisticsContactsName string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_name` = ?", logisticsContactsName).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsName 批量查找 物流联系人名字
func (obj *_CCustomerCopy1Mgr) GetBatchFromLogisticsContactsName(logisticsContactsNames []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_name` IN (?)", logisticsContactsNames).Find(&results).Error

	return
}

// GetFromLogisticsContactsTel 通过logistics_contacts_tel获取内容 物流联系人电话
func (obj *_CCustomerCopy1Mgr) GetFromLogisticsContactsTel(logisticsContactsTel string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_tel` = ?", logisticsContactsTel).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsTel 批量查找 物流联系人电话
func (obj *_CCustomerCopy1Mgr) GetBatchFromLogisticsContactsTel(logisticsContactsTels []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_tel` IN (?)", logisticsContactsTels).Find(&results).Error

	return
}

// GetFromLogisticsContactsPhone 通过logistics_contacts_phone获取内容 物流联系人手机
func (obj *_CCustomerCopy1Mgr) GetFromLogisticsContactsPhone(logisticsContactsPhone string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_phone` = ?", logisticsContactsPhone).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsPhone 批量查找 物流联系人手机
func (obj *_CCustomerCopy1Mgr) GetBatchFromLogisticsContactsPhone(logisticsContactsPhones []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_phone` IN (?)", logisticsContactsPhones).Find(&results).Error

	return
}

// GetFromLogisticsContactsQq 通过logistics_contacts_qq获取内容 物流联系人qq
func (obj *_CCustomerCopy1Mgr) GetFromLogisticsContactsQq(logisticsContactsQq string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_qq` = ?", logisticsContactsQq).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsQq 批量查找 物流联系人qq
func (obj *_CCustomerCopy1Mgr) GetBatchFromLogisticsContactsQq(logisticsContactsQqs []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_qq` IN (?)", logisticsContactsQqs).Find(&results).Error

	return
}

// GetFromLogisticsContactsEmail 通过logistics_contacts_email获取内容 物流联系人邮箱
func (obj *_CCustomerCopy1Mgr) GetFromLogisticsContactsEmail(logisticsContactsEmail string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_email` = ?", logisticsContactsEmail).Find(&results).Error

	return
}

// GetBatchFromLogisticsContactsEmail 批量查找 物流联系人邮箱
func (obj *_CCustomerCopy1Mgr) GetBatchFromLogisticsContactsEmail(logisticsContactsEmails []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`logistics_contacts_email` IN (?)", logisticsContactsEmails).Find(&results).Error

	return
}

// GetFromFinanceContactsName 通过finance_contacts_name获取内容 财务联系人名字
func (obj *_CCustomerCopy1Mgr) GetFromFinanceContactsName(financeContactsName string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_name` = ?", financeContactsName).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsName 批量查找 财务联系人名字
func (obj *_CCustomerCopy1Mgr) GetBatchFromFinanceContactsName(financeContactsNames []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_name` IN (?)", financeContactsNames).Find(&results).Error

	return
}

// GetFromFinanceContactsTel 通过finance_contacts_tel获取内容 财务联系人电话
func (obj *_CCustomerCopy1Mgr) GetFromFinanceContactsTel(financeContactsTel string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_tel` = ?", financeContactsTel).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsTel 批量查找 财务联系人电话
func (obj *_CCustomerCopy1Mgr) GetBatchFromFinanceContactsTel(financeContactsTels []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_tel` IN (?)", financeContactsTels).Find(&results).Error

	return
}

// GetFromFinanceContactsPhone 通过finance_contacts_phone获取内容 财务联系人手机
func (obj *_CCustomerCopy1Mgr) GetFromFinanceContactsPhone(financeContactsPhone string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_phone` = ?", financeContactsPhone).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsPhone 批量查找 财务联系人手机
func (obj *_CCustomerCopy1Mgr) GetBatchFromFinanceContactsPhone(financeContactsPhones []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_phone` IN (?)", financeContactsPhones).Find(&results).Error

	return
}

// GetFromFinanceContactsQq 通过finance_contacts_qq获取内容 财务联系人qq
func (obj *_CCustomerCopy1Mgr) GetFromFinanceContactsQq(financeContactsQq string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_qq` = ?", financeContactsQq).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsQq 批量查找 财务联系人qq
func (obj *_CCustomerCopy1Mgr) GetBatchFromFinanceContactsQq(financeContactsQqs []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_qq` IN (?)", financeContactsQqs).Find(&results).Error

	return
}

// GetFromFinanceContactsEmail 通过finance_contacts_email获取内容 财务联系人邮箱
func (obj *_CCustomerCopy1Mgr) GetFromFinanceContactsEmail(financeContactsEmail string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_email` = ?", financeContactsEmail).Find(&results).Error

	return
}

// GetBatchFromFinanceContactsEmail 批量查找 财务联系人邮箱
func (obj *_CCustomerCopy1Mgr) GetBatchFromFinanceContactsEmail(financeContactsEmails []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`finance_contacts_email` IN (?)", financeContactsEmails).Find(&results).Error

	return
}

// GetFromCustomServiceUserID 通过custom_service_user_id获取内容 客服id(oa系统中的用户id)
func (obj *_CCustomerCopy1Mgr) GetFromCustomServiceUserID(customServiceUserID int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`custom_service_user_id` = ?", customServiceUserID).Find(&results).Error

	return
}

// GetBatchFromCustomServiceUserID 批量查找 客服id(oa系统中的用户id)
func (obj *_CCustomerCopy1Mgr) GetBatchFromCustomServiceUserID(customServiceUserIDs []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`custom_service_user_id` IN (?)", customServiceUserIDs).Find(&results).Error

	return
}

// GetFromLockState 通过lock_state获取内容 账号锁定，0：正常，1：锁定
func (obj *_CCustomerCopy1Mgr) GetFromLockState(lockState int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`lock_state` = ?", lockState).Find(&results).Error

	return
}

// GetBatchFromLockState 批量查找 账号锁定，0：正常，1：锁定
func (obj *_CCustomerCopy1Mgr) GetBatchFromLockState(lockStates []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`lock_state` IN (?)", lockStates).Find(&results).Error

	return
}

// GetFromAuditState 通过audit_state获取内容 账户审核，0：未审核，1：已审核
func (obj *_CCustomerCopy1Mgr) GetFromAuditState(auditState int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`audit_state` = ?", auditState).Find(&results).Error

	return
}

// GetBatchFromAuditState 批量查找 账户审核，0：未审核，1：已审核
func (obj *_CCustomerCopy1Mgr) GetBatchFromAuditState(auditStates []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`audit_state` IN (?)", auditStates).Find(&results).Error

	return
}

// GetFromAccountStatus 通过account_status获取内容 账户状态 0：正常，1：欠费，2：禁用
func (obj *_CCustomerCopy1Mgr) GetFromAccountStatus(accountStatus int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`account_status` = ?", accountStatus).Find(&results).Error

	return
}

// GetBatchFromAccountStatus 批量查找 账户状态 0：正常，1：欠费，2：禁用
func (obj *_CCustomerCopy1Mgr) GetBatchFromAccountStatus(accountStatuss []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`account_status` IN (?)", accountStatuss).Find(&results).Error

	return
}

// GetFromBalance 通过balance获取内容 账户余额
func (obj *_CCustomerCopy1Mgr) GetFromBalance(balance float64) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`balance` = ?", balance).Find(&results).Error

	return
}

// GetBatchFromBalance 批量查找 账户余额
func (obj *_CCustomerCopy1Mgr) GetBatchFromBalance(balances []float64) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`balance` IN (?)", balances).Find(&results).Error

	return
}

// GetFromInitialCreditLimit 通过initial_credit_limit获取内容 信用额度
func (obj *_CCustomerCopy1Mgr) GetFromInitialCreditLimit(initialCreditLimit float64) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`initial_credit_limit` = ?", initialCreditLimit).Find(&results).Error

	return
}

// GetBatchFromInitialCreditLimit 批量查找 信用额度
func (obj *_CCustomerCopy1Mgr) GetBatchFromInitialCreditLimit(initialCreditLimits []float64) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`initial_credit_limit` IN (?)", initialCreditLimits).Find(&results).Error

	return
}

// GetFromOutCumulativeFee 通过out_cumulative_fee获取内容 出库累加费用
// Note:距离 上次结算 到 当前时间 累计出库余额 变动
func (obj *_CCustomerCopy1Mgr) GetFromOutCumulativeFee(outCumulativeFee float64) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`out_cumulative_fee` = ?", outCumulativeFee).Find(&results).Error

	return
}

// GetBatchFromOutCumulativeFee 批量查找 出库累加费用
// Note:距离 上次结算 到 当前时间 累计出库余额 变动
func (obj *_CCustomerCopy1Mgr) GetBatchFromOutCumulativeFee(outCumulativeFees []float64) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`out_cumulative_fee` IN (?)", outCumulativeFees).Find(&results).Error

	return
}

// GetFromLastCollectionDate 通过last_collection_date获取内容 最后收款时间
func (obj *_CCustomerCopy1Mgr) GetFromLastCollectionDate(lastCollectionDate time.Time) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`last_collection_date` = ?", lastCollectionDate).Find(&results).Error

	return
}

// GetBatchFromLastCollectionDate 批量查找 最后收款时间
func (obj *_CCustomerCopy1Mgr) GetBatchFromLastCollectionDate(lastCollectionDates []time.Time) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`last_collection_date` IN (?)", lastCollectionDates).Find(&results).Error

	return
}

// GetFromPaymentCycle 通过payment_cycle获取内容 收款周期
func (obj *_CCustomerCopy1Mgr) GetFromPaymentCycle(paymentCycle int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`payment_cycle` = ?", paymentCycle).Find(&results).Error

	return
}

// GetBatchFromPaymentCycle 批量查找 收款周期
func (obj *_CCustomerCopy1Mgr) GetBatchFromPaymentCycle(paymentCycles []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`payment_cycle` IN (?)", paymentCycles).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_CCustomerCopy1Mgr) GetFromCreateTime(createTime time.Time) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_CCustomerCopy1Mgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_CCustomerCopy1Mgr) GetFromCreateUser(createUser int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_CCustomerCopy1Mgr) GetBatchFromCreateUser(createUsers []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_CCustomerCopy1Mgr) GetFromUpdateTime(updateTime time.Time) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_CCustomerCopy1Mgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_CCustomerCopy1Mgr) GetFromUpdateUser(updateUser int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_CCustomerCopy1Mgr) GetBatchFromUpdateUser(updateUsers []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_CCustomerCopy1Mgr) GetFromVersion(version int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_CCustomerCopy1Mgr) GetBatchFromVersion(versions []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_CCustomerCopy1Mgr) GetFromDeleted(deleted int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_CCustomerCopy1Mgr) GetBatchFromDeleted(deleteds []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromAuthorizationCode 通过authorization_code获取内容 授权码
func (obj *_CCustomerCopy1Mgr) GetFromAuthorizationCode(authorizationCode string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`authorization_code` = ?", authorizationCode).Find(&results).Error

	return
}

// GetBatchFromAuthorizationCode 批量查找 授权码
func (obj *_CCustomerCopy1Mgr) GetBatchFromAuthorizationCode(authorizationCodes []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`authorization_code` IN (?)", authorizationCodes).Find(&results).Error

	return
}

// GetFromHandlerName 通过handler_name获取内容 经办人名称
func (obj *_CCustomerCopy1Mgr) GetFromHandlerName(handlerName string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`handler_name` = ?", handlerName).Find(&results).Error

	return
}

// GetBatchFromHandlerName 批量查找 经办人名称
func (obj *_CCustomerCopy1Mgr) GetBatchFromHandlerName(handlerNames []string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`handler_name` IN (?)", handlerNames).Find(&results).Error

	return
}

// GetFromHandlerID 通过handler_id获取内容 经办人id
func (obj *_CCustomerCopy1Mgr) GetFromHandlerID(handlerID int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`handler_id` = ?", handlerID).Find(&results).Error

	return
}

// GetBatchFromHandlerID 批量查找 经办人id
func (obj *_CCustomerCopy1Mgr) GetBatchFromHandlerID(handlerIDs []int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`handler_id` IN (?)", handlerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_CCustomerCopy1Mgr) FetchByPrimaryKey(id int) (result CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexCode primary or index 获取唯一内容
func (obj *_CCustomerCopy1Mgr) FetchUniqueByIndexCode(code string) (result CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`code` = ?", code).First(&result).Error

	return
}

// FetchUniqueByIndexUsername primary or index 获取唯一内容
func (obj *_CCustomerCopy1Mgr) FetchUniqueByIndexUsername(username string) (result CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`username` = ?", username).First(&result).Error

	return
}

// FetchIndexByCCustomerCustomerAliasIndex  获取多个内容
func (obj *_CCustomerCopy1Mgr) FetchIndexByCCustomerCustomerAliasIndex(customerAlias string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`customer_alias` = ?", customerAlias).Find(&results).Error

	return
}

// FetchIndexByCCustomerCompanyIndex  获取多个内容
func (obj *_CCustomerCopy1Mgr) FetchIndexByCCustomerCompanyIndex(company string) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// FetchIndexByCCustomerLockStateIndex  获取多个内容
func (obj *_CCustomerCopy1Mgr) FetchIndexByCCustomerLockStateIndex(lockState int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`lock_state` = ?", lockState).Find(&results).Error

	return
}

// FetchIndexByCCustomerAuditStateIndex  获取多个内容
func (obj *_CCustomerCopy1Mgr) FetchIndexByCCustomerAuditStateIndex(auditState int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`audit_state` = ?", auditState).Find(&results).Error

	return
}

// FetchIndexByCCustomerCreateTimeIndex  获取多个内容
func (obj *_CCustomerCopy1Mgr) FetchIndexByCCustomerCreateTimeIndex(createTime time.Time) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByCCustomerDeletedIndex  获取多个内容
func (obj *_CCustomerCopy1Mgr) FetchIndexByCCustomerDeletedIndex(deleted int) (results []*CCustomerCopy1, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(CCustomerCopy1{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}
