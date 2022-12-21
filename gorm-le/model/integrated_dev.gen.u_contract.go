package model

import (
	"context"
	"fmt"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type _UContractMgr struct {
	*_BaseMgr
}

// UContractMgr open func
func UContractMgr(db *gorm.DB) *_UContractMgr {
	if db == nil {
		panic(fmt.Errorf("UContractMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UContractMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_contract"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UContractMgr) GetTableName() string {
	return "u_contract"
}

// Reset 重置gorm会话
func (obj *_UContractMgr) Reset() *_UContractMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UContractMgr) Get() (result UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UContractMgr) Gets() (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UContractMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UContract{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UContractMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 合同编号
func (obj *_UContractMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithTitle title获取 合同标题
func (obj *_UContractMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 合同内容
func (obj *_UContractMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithType type获取 合同类型
func (obj *_UContractMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSignDate sign_date获取 签订日期
func (obj *_UContractMgr) WithSignDate(signDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["sign_date"] = signDate })
}

// WithExpireDate expire_date获取 到期日期
func (obj *_UContractMgr) WithExpireDate(expireDate datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["expire_date"] = expireDate })
}

// WithDrafterUserID drafter_user_id获取 拟稿人用户id
func (obj *_UContractMgr) WithDrafterUserID(drafterUserID int) Option {
	return optionFunc(func(o *options) { o.query["drafter_user_id"] = drafterUserID })
}

// WithPartyA party_a获取 甲方
func (obj *_UContractMgr) WithPartyA(partyA string) Option {
	return optionFunc(func(o *options) { o.query["party_a"] = partyA })
}

// WithPartyB party_b获取 乙方
func (obj *_UContractMgr) WithPartyB(partyB string) Option {
	return optionFunc(func(o *options) { o.query["party_b"] = partyB })
}

// WithSignStatus sign_status获取 合同状态(履行中、合同到期、合同作废)
func (obj *_UContractMgr) WithSignStatus(signStatus string) Option {
	return optionFunc(func(o *options) { o.query["sign_status"] = signStatus })
}

// WithTarget target获取 合同用途名称对应u_contract_use_type（合同用途类型表） 表的type_name字段
func (obj *_UContractMgr) WithTarget(target string) Option {
	return optionFunc(func(o *options) { o.query["target"] = target })
}

// WithAttachment attachment获取 附件
func (obj *_UContractMgr) WithAttachment(attachment string) Option {
	return optionFunc(func(o *options) { o.query["attachment"] = attachment })
}

// WithESignature e_signature获取 电子签名资料
func (obj *_UContractMgr) WithESignature(eSignature string) Option {
	return optionFunc(func(o *options) { o.query["e_signature"] = eSignature })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UContractMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UContractMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UContractMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_UContractMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UContractMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_UContractMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithContractTypeID contract_type_id获取 合同用途类型表(u_contract_use_type)的id
func (obj *_UContractMgr) WithContractTypeID(contractTypeID int) Option {
	return optionFunc(func(o *options) { o.query["contract_type_id"] = contractTypeID })
}

// WithIsDeposit is_deposit获取 是否需要押金 0 否 1 是
func (obj *_UContractMgr) WithIsDeposit(isDeposit []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_deposit"] = isDeposit })
}

// WithDepositAmount deposit_amount获取 押金金额
func (obj *_UContractMgr) WithDepositAmount(depositAmount float64) Option {
	return optionFunc(func(o *options) { o.query["deposit_amount"] = depositAmount })
}

// WithCustomEncoding custom_encoding获取 自定义编码
func (obj *_UContractMgr) WithCustomEncoding(customEncoding string) Option {
	return optionFunc(func(o *options) { o.query["custom_encoding"] = customEncoding })
}

// WithPartyACustomerID party_a_customer_id获取 客户id
func (obj *_UContractMgr) WithPartyACustomerID(partyACustomerID int) Option {
	return optionFunc(func(o *options) { o.query["party_a_customer_id"] = partyACustomerID })
}

// WithExtendedData extended_data获取 扩展数据
func (obj *_UContractMgr) WithExtendedData(extendedData datatypes.JSON) Option {
	return optionFunc(func(o *options) { o.query["extended_data"] = extendedData })
}

// WithCustomerName customer_name获取 客户名称
func (obj *_UContractMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// GetByOption 功能选项模式获取
func (obj *_UContractMgr) GetByOption(opts ...Option) (result UContract, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UContractMgr) GetByOptions(opts ...Option) (results []*UContract, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UContractMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UContract, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UContract{}).Where(options.query)
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
func (obj *_UContractMgr) GetFromID(id int) (result UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UContractMgr) GetBatchFromID(ids []int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 合同编号
func (obj *_UContractMgr) GetFromCode(code string) (result UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`code` = ?", code).First(&result).Error

	return
}

// GetBatchFromCode 批量查找 合同编号
func (obj *_UContractMgr) GetBatchFromCode(codes []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 合同标题
func (obj *_UContractMgr) GetFromTitle(title string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 合同标题
func (obj *_UContractMgr) GetBatchFromTitle(titles []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 合同内容
func (obj *_UContractMgr) GetFromContent(content string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 合同内容
func (obj *_UContractMgr) GetBatchFromContent(contents []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 合同类型
func (obj *_UContractMgr) GetFromType(_type string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 合同类型
func (obj *_UContractMgr) GetBatchFromType(_types []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromSignDate 通过sign_date获取内容 签订日期
func (obj *_UContractMgr) GetFromSignDate(signDate datatypes.Date) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`sign_date` = ?", signDate).Find(&results).Error

	return
}

// GetBatchFromSignDate 批量查找 签订日期
func (obj *_UContractMgr) GetBatchFromSignDate(signDates []datatypes.Date) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`sign_date` IN (?)", signDates).Find(&results).Error

	return
}

// GetFromExpireDate 通过expire_date获取内容 到期日期
func (obj *_UContractMgr) GetFromExpireDate(expireDate datatypes.Date) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`expire_date` = ?", expireDate).Find(&results).Error

	return
}

// GetBatchFromExpireDate 批量查找 到期日期
func (obj *_UContractMgr) GetBatchFromExpireDate(expireDates []datatypes.Date) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`expire_date` IN (?)", expireDates).Find(&results).Error

	return
}

// GetFromDrafterUserID 通过drafter_user_id获取内容 拟稿人用户id
func (obj *_UContractMgr) GetFromDrafterUserID(drafterUserID int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`drafter_user_id` = ?", drafterUserID).Find(&results).Error

	return
}

// GetBatchFromDrafterUserID 批量查找 拟稿人用户id
func (obj *_UContractMgr) GetBatchFromDrafterUserID(drafterUserIDs []int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`drafter_user_id` IN (?)", drafterUserIDs).Find(&results).Error

	return
}

// GetFromPartyA 通过party_a获取内容 甲方
func (obj *_UContractMgr) GetFromPartyA(partyA string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`party_a` = ?", partyA).Find(&results).Error

	return
}

// GetBatchFromPartyA 批量查找 甲方
func (obj *_UContractMgr) GetBatchFromPartyA(partyAs []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`party_a` IN (?)", partyAs).Find(&results).Error

	return
}

// GetFromPartyB 通过party_b获取内容 乙方
func (obj *_UContractMgr) GetFromPartyB(partyB string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`party_b` = ?", partyB).Find(&results).Error

	return
}

// GetBatchFromPartyB 批量查找 乙方
func (obj *_UContractMgr) GetBatchFromPartyB(partyBs []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`party_b` IN (?)", partyBs).Find(&results).Error

	return
}

// GetFromSignStatus 通过sign_status获取内容 合同状态(履行中、合同到期、合同作废)
func (obj *_UContractMgr) GetFromSignStatus(signStatus string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`sign_status` = ?", signStatus).Find(&results).Error

	return
}

// GetBatchFromSignStatus 批量查找 合同状态(履行中、合同到期、合同作废)
func (obj *_UContractMgr) GetBatchFromSignStatus(signStatuss []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`sign_status` IN (?)", signStatuss).Find(&results).Error

	return
}

// GetFromTarget 通过target获取内容 合同用途名称对应u_contract_use_type（合同用途类型表） 表的type_name字段
func (obj *_UContractMgr) GetFromTarget(target string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`target` = ?", target).Find(&results).Error

	return
}

// GetBatchFromTarget 批量查找 合同用途名称对应u_contract_use_type（合同用途类型表） 表的type_name字段
func (obj *_UContractMgr) GetBatchFromTarget(targets []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`target` IN (?)", targets).Find(&results).Error

	return
}

// GetFromAttachment 通过attachment获取内容 附件
func (obj *_UContractMgr) GetFromAttachment(attachment string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`attachment` = ?", attachment).Find(&results).Error

	return
}

// GetBatchFromAttachment 批量查找 附件
func (obj *_UContractMgr) GetBatchFromAttachment(attachments []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`attachment` IN (?)", attachments).Find(&results).Error

	return
}

// GetFromESignature 通过e_signature获取内容 电子签名资料
func (obj *_UContractMgr) GetFromESignature(eSignature string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`e_signature` = ?", eSignature).Find(&results).Error

	return
}

// GetBatchFromESignature 批量查找 电子签名资料
func (obj *_UContractMgr) GetBatchFromESignature(eSignatures []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`e_signature` IN (?)", eSignatures).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UContractMgr) GetFromCreateTime(createTime time.Time) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UContractMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UContractMgr) GetFromCreateUser(createUser int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UContractMgr) GetBatchFromCreateUser(createUsers []int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UContractMgr) GetFromUpdateTime(updateTime time.Time) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UContractMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_UContractMgr) GetFromUpdateUser(updateUser int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_UContractMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UContractMgr) GetFromVersion(version int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UContractMgr) GetBatchFromVersion(versions []int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_UContractMgr) GetFromDeleted(deleted int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_UContractMgr) GetBatchFromDeleted(deleteds []int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromContractTypeID 通过contract_type_id获取内容 合同用途类型表(u_contract_use_type)的id
func (obj *_UContractMgr) GetFromContractTypeID(contractTypeID int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`contract_type_id` = ?", contractTypeID).Find(&results).Error

	return
}

// GetBatchFromContractTypeID 批量查找 合同用途类型表(u_contract_use_type)的id
func (obj *_UContractMgr) GetBatchFromContractTypeID(contractTypeIDs []int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`contract_type_id` IN (?)", contractTypeIDs).Find(&results).Error

	return
}

// GetFromIsDeposit 通过is_deposit获取内容 是否需要押金 0 否 1 是
func (obj *_UContractMgr) GetFromIsDeposit(isDeposit []uint8) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`is_deposit` = ?", isDeposit).Find(&results).Error

	return
}

// GetBatchFromIsDeposit 批量查找 是否需要押金 0 否 1 是
func (obj *_UContractMgr) GetBatchFromIsDeposit(isDeposits [][]uint8) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`is_deposit` IN (?)", isDeposits).Find(&results).Error

	return
}

// GetFromDepositAmount 通过deposit_amount获取内容 押金金额
func (obj *_UContractMgr) GetFromDepositAmount(depositAmount float64) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`deposit_amount` = ?", depositAmount).Find(&results).Error

	return
}

// GetBatchFromDepositAmount 批量查找 押金金额
func (obj *_UContractMgr) GetBatchFromDepositAmount(depositAmounts []float64) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`deposit_amount` IN (?)", depositAmounts).Find(&results).Error

	return
}

// GetFromCustomEncoding 通过custom_encoding获取内容 自定义编码
func (obj *_UContractMgr) GetFromCustomEncoding(customEncoding string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`custom_encoding` = ?", customEncoding).Find(&results).Error

	return
}

// GetBatchFromCustomEncoding 批量查找 自定义编码
func (obj *_UContractMgr) GetBatchFromCustomEncoding(customEncodings []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`custom_encoding` IN (?)", customEncodings).Find(&results).Error

	return
}

// GetFromPartyACustomerID 通过party_a_customer_id获取内容 客户id
func (obj *_UContractMgr) GetFromPartyACustomerID(partyACustomerID int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`party_a_customer_id` = ?", partyACustomerID).Find(&results).Error

	return
}

// GetBatchFromPartyACustomerID 批量查找 客户id
func (obj *_UContractMgr) GetBatchFromPartyACustomerID(partyACustomerIDs []int) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`party_a_customer_id` IN (?)", partyACustomerIDs).Find(&results).Error

	return
}

// GetFromExtendedData 通过extended_data获取内容 扩展数据
func (obj *_UContractMgr) GetFromExtendedData(extendedData datatypes.JSON) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`extended_data` = ?", extendedData).Find(&results).Error

	return
}

// GetBatchFromExtendedData 批量查找 扩展数据
func (obj *_UContractMgr) GetBatchFromExtendedData(extendedDatas []datatypes.JSON) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`extended_data` IN (?)", extendedDatas).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 客户名称
func (obj *_UContractMgr) GetFromCustomerName(customerName string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 客户名称
func (obj *_UContractMgr) GetBatchFromCustomerName(customerNames []string) (results []*UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UContractMgr) FetchByPrimaryKey(id int) (result UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByIndexCode primary or index 获取唯一内容
func (obj *_UContractMgr) FetchUniqueByIndexCode(code string) (result UContract, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UContract{}).Where("`code` = ?", code).First(&result).Error

	return
}
