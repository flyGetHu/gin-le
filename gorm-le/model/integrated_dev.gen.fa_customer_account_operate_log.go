package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaCustomerAccountOperateLogMgr struct {
	*_BaseMgr
}

// FaCustomerAccountOperateLogMgr open func
func FaCustomerAccountOperateLogMgr(db *gorm.DB) *_FaCustomerAccountOperateLogMgr {
	if db == nil {
		panic(fmt.Errorf("FaCustomerAccountOperateLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaCustomerAccountOperateLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_customer_account_operate_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaCustomerAccountOperateLogMgr) GetTableName() string {
	return "fa_customer_account_operate_log"
}

// Reset 重置gorm会话
func (obj *_FaCustomerAccountOperateLogMgr) Reset() *_FaCustomerAccountOperateLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaCustomerAccountOperateLogMgr) Get() (result FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaCustomerAccountOperateLogMgr) Gets() (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaCustomerAccountOperateLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_FaCustomerAccountOperateLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOperateModule operate_module获取 操作模块
func (obj *_FaCustomerAccountOperateLogMgr) WithOperateModule(operateModule string) Option {
	return optionFunc(func(o *options) { o.query["operate_module"] = operateModule })
}

// WithCustomerID customer_id获取 客户id
func (obj *_FaCustomerAccountOperateLogMgr) WithCustomerID(customerID int) Option {
	return optionFunc(func(o *options) { o.query["customer_id"] = customerID })
}

// WithOperateType operate_type获取 操作类型
func (obj *_FaCustomerAccountOperateLogMgr) WithOperateType(operateType string) Option {
	return optionFunc(func(o *options) { o.query["operate_type"] = operateType })
}

// WithOperateDescription operate_description获取 操作描述
func (obj *_FaCustomerAccountOperateLogMgr) WithOperateDescription(operateDescription string) Option {
	return optionFunc(func(o *options) { o.query["operate_description"] = operateDescription })
}

// WithUpdateParams update_params获取 更新参数
func (obj *_FaCustomerAccountOperateLogMgr) WithUpdateParams(updateParams string) Option {
	return optionFunc(func(o *options) { o.query["update_params"] = updateParams })
}

// WithPaymentMethod payment_method获取 收款方式
func (obj *_FaCustomerAccountOperateLogMgr) WithPaymentMethod(paymentMethod string) Option {
	return optionFunc(func(o *options) { o.query["payment_method"] = paymentMethod })
}

// WithAmount amount获取 金额
func (obj *_FaCustomerAccountOperateLogMgr) WithAmount(amount float64) Option {
	return optionFunc(func(o *options) { o.query["amount"] = amount })
}

// WithRechargeDate recharge_date获取 充值时间
func (obj *_FaCustomerAccountOperateLogMgr) WithRechargeDate(rechargeDate time.Time) Option {
	return optionFunc(func(o *options) { o.query["recharge_date"] = rechargeDate })
}

// WithAccount account获取 账户
func (obj *_FaCustomerAccountOperateLogMgr) WithAccount(account string) Option {
	return optionFunc(func(o *options) { o.query["account"] = account })
}

// WithTicketNumber ticket_number获取 票证号
func (obj *_FaCustomerAccountOperateLogMgr) WithTicketNumber(ticketNumber string) Option {
	return optionFunc(func(o *options) { o.query["ticket_number"] = ticketNumber })
}

// WithAttachment attachment获取 附件
func (obj *_FaCustomerAccountOperateLogMgr) WithAttachment(attachment string) Option {
	return optionFunc(func(o *options) { o.query["attachment"] = attachment })
}

// WithRemark remark获取 备注
func (obj *_FaCustomerAccountOperateLogMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithAgent agent获取 经办人
func (obj *_FaCustomerAccountOperateLogMgr) WithAgent(agent string) Option {
	return optionFunc(func(o *options) { o.query["agent"] = agent })
}

// WithOperateIP operate_ip获取 操作IP
func (obj *_FaCustomerAccountOperateLogMgr) WithOperateIP(operateIP string) Option {
	return optionFunc(func(o *options) { o.query["operate_ip"] = operateIP })
}

// WithCreateUser create_user获取 操作人id
func (obj *_FaCustomerAccountOperateLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 操作时间
func (obj *_FaCustomerAccountOperateLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaCustomerAccountOperateLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaCustomerAccountOperateLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaCustomerAccountOperateLogMgr) GetByOption(opts ...Option) (result FaCustomerAccountOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaCustomerAccountOperateLogMgr) GetByOptions(opts ...Option) (results []*FaCustomerAccountOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaCustomerAccountOperateLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaCustomerAccountOperateLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where(options.query)
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
func (obj *_FaCustomerAccountOperateLogMgr) GetFromID(id int) (result FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromID(ids []int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOperateModule 通过operate_module获取内容 操作模块
func (obj *_FaCustomerAccountOperateLogMgr) GetFromOperateModule(operateModule string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`operate_module` = ?", operateModule).Find(&results).Error

	return
}

// GetBatchFromOperateModule 批量查找 操作模块
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromOperateModule(operateModules []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`operate_module` IN (?)", operateModules).Find(&results).Error

	return
}

// GetFromCustomerID 通过customer_id获取内容 客户id
func (obj *_FaCustomerAccountOperateLogMgr) GetFromCustomerID(customerID int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}

// GetBatchFromCustomerID 批量查找 客户id
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromCustomerID(customerIDs []int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`customer_id` IN (?)", customerIDs).Find(&results).Error

	return
}

// GetFromOperateType 通过operate_type获取内容 操作类型
func (obj *_FaCustomerAccountOperateLogMgr) GetFromOperateType(operateType string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`operate_type` = ?", operateType).Find(&results).Error

	return
}

// GetBatchFromOperateType 批量查找 操作类型
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromOperateType(operateTypes []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`operate_type` IN (?)", operateTypes).Find(&results).Error

	return
}

// GetFromOperateDescription 通过operate_description获取内容 操作描述
func (obj *_FaCustomerAccountOperateLogMgr) GetFromOperateDescription(operateDescription string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`operate_description` = ?", operateDescription).Find(&results).Error

	return
}

// GetBatchFromOperateDescription 批量查找 操作描述
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromOperateDescription(operateDescriptions []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`operate_description` IN (?)", operateDescriptions).Find(&results).Error

	return
}

// GetFromUpdateParams 通过update_params获取内容 更新参数
func (obj *_FaCustomerAccountOperateLogMgr) GetFromUpdateParams(updateParams string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`update_params` = ?", updateParams).Find(&results).Error

	return
}

// GetBatchFromUpdateParams 批量查找 更新参数
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromUpdateParams(updateParamss []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`update_params` IN (?)", updateParamss).Find(&results).Error

	return
}

// GetFromPaymentMethod 通过payment_method获取内容 收款方式
func (obj *_FaCustomerAccountOperateLogMgr) GetFromPaymentMethod(paymentMethod string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`payment_method` = ?", paymentMethod).Find(&results).Error

	return
}

// GetBatchFromPaymentMethod 批量查找 收款方式
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromPaymentMethod(paymentMethods []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`payment_method` IN (?)", paymentMethods).Find(&results).Error

	return
}

// GetFromAmount 通过amount获取内容 金额
func (obj *_FaCustomerAccountOperateLogMgr) GetFromAmount(amount float64) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`amount` = ?", amount).Find(&results).Error

	return
}

// GetBatchFromAmount 批量查找 金额
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromAmount(amounts []float64) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`amount` IN (?)", amounts).Find(&results).Error

	return
}

// GetFromRechargeDate 通过recharge_date获取内容 充值时间
func (obj *_FaCustomerAccountOperateLogMgr) GetFromRechargeDate(rechargeDate time.Time) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`recharge_date` = ?", rechargeDate).Find(&results).Error

	return
}

// GetBatchFromRechargeDate 批量查找 充值时间
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromRechargeDate(rechargeDates []time.Time) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`recharge_date` IN (?)", rechargeDates).Find(&results).Error

	return
}

// GetFromAccount 通过account获取内容 账户
func (obj *_FaCustomerAccountOperateLogMgr) GetFromAccount(account string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`account` = ?", account).Find(&results).Error

	return
}

// GetBatchFromAccount 批量查找 账户
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromAccount(accounts []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`account` IN (?)", accounts).Find(&results).Error

	return
}

// GetFromTicketNumber 通过ticket_number获取内容 票证号
func (obj *_FaCustomerAccountOperateLogMgr) GetFromTicketNumber(ticketNumber string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`ticket_number` = ?", ticketNumber).Find(&results).Error

	return
}

// GetBatchFromTicketNumber 批量查找 票证号
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromTicketNumber(ticketNumbers []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`ticket_number` IN (?)", ticketNumbers).Find(&results).Error

	return
}

// GetFromAttachment 通过attachment获取内容 附件
func (obj *_FaCustomerAccountOperateLogMgr) GetFromAttachment(attachment string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`attachment` = ?", attachment).Find(&results).Error

	return
}

// GetBatchFromAttachment 批量查找 附件
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromAttachment(attachments []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`attachment` IN (?)", attachments).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaCustomerAccountOperateLogMgr) GetFromRemark(remark string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromRemark(remarks []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromAgent 通过agent获取内容 经办人
func (obj *_FaCustomerAccountOperateLogMgr) GetFromAgent(agent string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`agent` = ?", agent).Find(&results).Error

	return
}

// GetBatchFromAgent 批量查找 经办人
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromAgent(agents []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`agent` IN (?)", agents).Find(&results).Error

	return
}

// GetFromOperateIP 通过operate_ip获取内容 操作IP
func (obj *_FaCustomerAccountOperateLogMgr) GetFromOperateIP(operateIP string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`operate_ip` = ?", operateIP).Find(&results).Error

	return
}

// GetBatchFromOperateIP 批量查找 操作IP
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromOperateIP(operateIPs []string) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`operate_ip` IN (?)", operateIPs).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 操作人id
func (obj *_FaCustomerAccountOperateLogMgr) GetFromCreateUser(createUser int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 操作人id
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 操作时间
func (obj *_FaCustomerAccountOperateLogMgr) GetFromCreateTime(createTime time.Time) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 操作时间
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaCustomerAccountOperateLogMgr) GetFromVersion(version int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromVersion(versions []int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaCustomerAccountOperateLogMgr) GetFromDeleted(deleted int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaCustomerAccountOperateLogMgr) GetBatchFromDeleted(deleteds []int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaCustomerAccountOperateLogMgr) FetchByPrimaryKey(id int) (result FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgChannelOperateLogIDUIndex primary or index 获取唯一内容
func (obj *_FaCustomerAccountOperateLogMgr) FetchUniqueByLgChannelOperateLogIDUIndex(id int) (result FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByLgChannelOperateLogCustomerIDIndex  获取多个内容
func (obj *_FaCustomerAccountOperateLogMgr) FetchIndexByLgChannelOperateLogCustomerIDIndex(customerID int) (results []*FaCustomerAccountOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaCustomerAccountOperateLog{}).Where("`customer_id` = ?", customerID).Find(&results).Error

	return
}
