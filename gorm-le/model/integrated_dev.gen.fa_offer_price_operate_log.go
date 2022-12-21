package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _FaOfferPriceOperateLogMgr struct {
	*_BaseMgr
}

// FaOfferPriceOperateLogMgr open func
func FaOfferPriceOperateLogMgr(db *gorm.DB) *_FaOfferPriceOperateLogMgr {
	if db == nil {
		panic(fmt.Errorf("FaOfferPriceOperateLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_FaOfferPriceOperateLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("fa_offer_price_operate_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FaOfferPriceOperateLogMgr) GetTableName() string {
	return "fa_offer_price_operate_log"
}

// Reset 重置gorm会话
func (obj *_FaOfferPriceOperateLogMgr) Reset() *_FaOfferPriceOperateLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FaOfferPriceOperateLogMgr) Get() (result FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FaOfferPriceOperateLogMgr) Gets() (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_FaOfferPriceOperateLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_FaOfferPriceOperateLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithIP ip获取 ip
func (obj *_FaOfferPriceOperateLogMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithOfferPriceID offer_price_id获取 报价表ID
func (obj *_FaOfferPriceOperateLogMgr) WithOfferPriceID(offerPriceID int) Option {
	return optionFunc(func(o *options) { o.query["offer_price_id"] = offerPriceID })
}

// WithOperationContent operation_content获取 操作内容
func (obj *_FaOfferPriceOperateLogMgr) WithOperationContent(operationContent string) Option {
	return optionFunc(func(o *options) { o.query["operation_content"] = operationContent })
}

// WithRemark remark获取 备注
func (obj *_FaOfferPriceOperateLogMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_FaOfferPriceOperateLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_FaOfferPriceOperateLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithOperatorName operator_name获取 操作人名称
func (obj *_FaOfferPriceOperateLogMgr) WithOperatorName(operatorName string) Option {
	return optionFunc(func(o *options) { o.query["operator_name"] = operatorName })
}

// WithVersion version获取 乐观锁
func (obj *_FaOfferPriceOperateLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_FaOfferPriceOperateLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_FaOfferPriceOperateLogMgr) GetByOption(opts ...Option) (result FaOfferPriceOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_FaOfferPriceOperateLogMgr) GetByOptions(opts ...Option) (results []*FaOfferPriceOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FaOfferPriceOperateLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FaOfferPriceOperateLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where(options.query)
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
func (obj *_FaOfferPriceOperateLogMgr) GetFromID(id int) (result FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromID(ids []int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 ip
func (obj *_FaOfferPriceOperateLogMgr) GetFromIP(ip string) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找 ip
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromIP(ips []string) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

// GetFromOfferPriceID 通过offer_price_id获取内容 报价表ID
func (obj *_FaOfferPriceOperateLogMgr) GetFromOfferPriceID(offerPriceID int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`offer_price_id` = ?", offerPriceID).Find(&results).Error

	return
}

// GetBatchFromOfferPriceID 批量查找 报价表ID
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromOfferPriceID(offerPriceIDs []int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`offer_price_id` IN (?)", offerPriceIDs).Find(&results).Error

	return
}

// GetFromOperationContent 通过operation_content获取内容 操作内容
func (obj *_FaOfferPriceOperateLogMgr) GetFromOperationContent(operationContent string) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`operation_content` = ?", operationContent).Find(&results).Error

	return
}

// GetBatchFromOperationContent 批量查找 操作内容
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromOperationContent(operationContents []string) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`operation_content` IN (?)", operationContents).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_FaOfferPriceOperateLogMgr) GetFromRemark(remark string) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromRemark(remarks []string) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_FaOfferPriceOperateLogMgr) GetFromCreateTime(createTime time.Time) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_FaOfferPriceOperateLogMgr) GetFromCreateUser(createUser int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromOperatorName 通过operator_name获取内容 操作人名称
func (obj *_FaOfferPriceOperateLogMgr) GetFromOperatorName(operatorName string) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`operator_name` = ?", operatorName).Find(&results).Error

	return
}

// GetBatchFromOperatorName 批量查找 操作人名称
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromOperatorName(operatorNames []string) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`operator_name` IN (?)", operatorNames).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_FaOfferPriceOperateLogMgr) GetFromVersion(version int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromVersion(versions []int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_FaOfferPriceOperateLogMgr) GetFromDeleted(deleted int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_FaOfferPriceOperateLogMgr) GetBatchFromDeleted(deleteds []int) (results []*FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_FaOfferPriceOperateLogMgr) FetchByPrimaryKey(id int) (result FaOfferPriceOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FaOfferPriceOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
