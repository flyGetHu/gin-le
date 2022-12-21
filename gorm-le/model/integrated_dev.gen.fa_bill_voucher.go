package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaBillVoucherMgr struct {
	*_BaseMgr
}

// FaBillVoucherMgr open func
func FaBillVoucherMgr(db *gorm.DB) *_FaBillVoucherMgr {
	if db == nil {
		panic(fmt.Errorf("FaBillVoucherMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaBillVoucherMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_bill_voucher"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaBillVoucherMgr) GetTableName() string {
	return "fa_bill_voucher"
}

// Reset 重置gorm会话
func (obj *_FaBillVoucherMgr) Reset() *_FaBillVoucherMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaBillVoucherMgr) Get() (result FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaBillVoucherMgr) Gets() (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaBillVoucherMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaBillVoucherMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 凭证标题
func (obj *_FaBillVoucherMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithCompanyID company_id获取 凭证公司id(客户组织id)
func (obj *_FaBillVoucherMgr) WithCompanyID(companyID int) Option {
	return optionFunc(func(o *options) { o.query["company_id"] = companyID })
}

// WithCompany company获取 凭证公司
func (obj *_FaBillVoucherMgr) WithCompany(company string) Option {
	return optionFunc(func(o *options) { o.query["company"] = company })
}

// WithCode code获取 凭证编码
func (obj *_FaBillVoucherMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithVoucherTime voucher_time获取 凭证年月
func (obj *_FaBillVoucherMgr) WithVoucherTime(voucherTime string) Option {
	return optionFunc(func(o *options) { o.query["voucher_time"] = voucherTime })
}

// WithAttachment attachment获取 附件：附件上传工具，必填
func (obj *_FaBillVoucherMgr) WithAttachment(attachment string) Option {
	return optionFunc(func(o *options) { o.query["attachment"] = attachment })
}

// WithAttachmentCount attachment_count获取 附件数量
func (obj *_FaBillVoucherMgr) WithAttachmentCount(attachmentCount int) Option {
	return optionFunc(func(o *options) { o.query["attachment_count"] = attachmentCount })
}

// WithRemark remark获取 备注
func (obj *_FaBillVoucherMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateUser create_user获取 创建人
func (obj *_FaBillVoucherMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaBillVoucherMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_FaBillVoucherMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_FaBillVoucherMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_FaBillVoucherMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaBillVoucherMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaBillVoucherMgr) GetByOption(opts ...Option) (result FaBillVoucher, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaBillVoucherMgr) GetByOptions(opts ...Option) (results []*FaBillVoucher, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaBillVoucherMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaBillVoucher, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_FaBillVoucherMgr) GetFromID(id int64) (result FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaBillVoucherMgr) GetBatchFromID(ids []int64) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 凭证标题
func (obj *_FaBillVoucherMgr) GetFromTitle(title string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 凭证标题
func (obj *_FaBillVoucherMgr) GetBatchFromTitle(titles []string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromCompanyID 通过company_id获取内容 凭证公司id(客户组织id)
func (obj *_FaBillVoucherMgr) GetFromCompanyID(companyID int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`company_id` = ?", companyID).Find(&results).Error

	return
}

// GetBatchFromCompanyID 批量查找 凭证公司id(客户组织id)
func (obj *_FaBillVoucherMgr) GetBatchFromCompanyID(companyIDs []int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`company_id` IN (?)", companyIDs).Find(&results).Error

	return
}

// GetFromCompany 通过company获取内容 凭证公司
func (obj *_FaBillVoucherMgr) GetFromCompany(company string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// GetBatchFromCompany 批量查找 凭证公司
func (obj *_FaBillVoucherMgr) GetBatchFromCompany(companys []string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`company` IN (?)", companys).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 凭证编码
func (obj *_FaBillVoucherMgr) GetFromCode(code string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 凭证编码
func (obj *_FaBillVoucherMgr) GetBatchFromCode(codes []string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromVoucherTime 通过voucher_time获取内容 凭证年月
func (obj *_FaBillVoucherMgr) GetFromVoucherTime(voucherTime string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`voucher_time` = ?", voucherTime).Find(&results).Error

	return
}

// GetBatchFromVoucherTime 批量查找 凭证年月
func (obj *_FaBillVoucherMgr) GetBatchFromVoucherTime(voucherTimes []string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`voucher_time` IN (?)", voucherTimes).Find(&results).Error

	return
}

// GetFromAttachment 通过attachment获取内容 附件：附件上传工具，必填
func (obj *_FaBillVoucherMgr) GetFromAttachment(attachment string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`attachment` = ?", attachment).Find(&results).Error

	return
}

// GetBatchFromAttachment 批量查找 附件：附件上传工具，必填
func (obj *_FaBillVoucherMgr) GetBatchFromAttachment(attachments []string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`attachment` IN (?)", attachments).Find(&results).Error

	return
}

// GetFromAttachmentCount 通过attachment_count获取内容 附件数量
func (obj *_FaBillVoucherMgr) GetFromAttachmentCount(attachmentCount int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`attachment_count` = ?", attachmentCount).Find(&results).Error

	return
}

// GetBatchFromAttachmentCount 批量查找 附件数量
func (obj *_FaBillVoucherMgr) GetBatchFromAttachmentCount(attachmentCounts []int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`attachment_count` IN (?)", attachmentCounts).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaBillVoucherMgr) GetFromRemark(remark string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaBillVoucherMgr) GetBatchFromRemark(remarks []string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_FaBillVoucherMgr) GetFromCreateUser(createUser int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_FaBillVoucherMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaBillVoucherMgr) GetFromCreateTime(createTime time.Time) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaBillVoucherMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_FaBillVoucherMgr) GetFromUpdateUser(updateUser int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_FaBillVoucherMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_FaBillVoucherMgr) GetFromUpdateTime(updateTime time.Time) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_FaBillVoucherMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaBillVoucherMgr) GetFromVersion(version int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaBillVoucherMgr) GetBatchFromVersion(versions []int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaBillVoucherMgr) GetFromDeleted(deleted int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaBillVoucherMgr) GetBatchFromDeleted(deleteds []int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaBillVoucherMgr) FetchByPrimaryKey(id int64) (result FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByIndexUniqueTitleCompany primary or index 获取唯一内容
func (obj *_FaBillVoucherMgr) FetchUniqueIndexByIndexUniqueTitleCompany(title string, company string) (result FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`title` = ? AND `company` = ?", title, company).First(&result).Error

	return
}

// FetchUniqueIndexByIndexUniqueCodeCompany primary or index 获取唯一内容
func (obj *_FaBillVoucherMgr) FetchUniqueIndexByIndexUniqueCodeCompany(company string, code string) (result FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`company` = ? AND `code` = ?", company, code).First(&result).Error

	return
}

// FetchIndexByIndexTitle  获取多个内容
func (obj *_FaBillVoucherMgr) FetchIndexByIndexTitle(title string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// FetchIndexByIndexCompanyID  获取多个内容
func (obj *_FaBillVoucherMgr) FetchIndexByIndexCompanyID(companyID int) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`company_id` = ?", companyID).Find(&results).Error

	return
}

// FetchIndexByIndexCompany  获取多个内容
func (obj *_FaBillVoucherMgr) FetchIndexByIndexCompany(company string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`company` = ?", company).Find(&results).Error

	return
}

// FetchIndexByIndexCode  获取多个内容
func (obj *_FaBillVoucherMgr) FetchIndexByIndexCode(code string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// FetchIndexByIndexVoucherTime  获取多个内容
func (obj *_FaBillVoucherMgr) FetchIndexByIndexVoucherTime(voucherTime string) (results []*FaBillVoucher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaBillVoucher{}).Where("`voucher_time` = ?", voucherTime).Find(&results).Error

	return
}
