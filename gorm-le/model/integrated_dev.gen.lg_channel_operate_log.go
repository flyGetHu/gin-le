package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgChannelOperateLogMgr struct {
	*_BaseMgr
}

// LgChannelOperateLogMgr open func
func LgChannelOperateLogMgr(db *gorm.DB) *_LgChannelOperateLogMgr {
	if db == nil {
		panic(fmt.Errorf("LgChannelOperateLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgChannelOperateLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_channel_operate_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgChannelOperateLogMgr) GetTableName() string {
	return "lg_channel_operate_log"
}

// Reset 重置gorm会话
func (obj *_LgChannelOperateLogMgr) Reset() *_LgChannelOperateLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgChannelOperateLogMgr) Get() (result LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgChannelOperateLogMgr) Gets() (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgChannelOperateLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgChannelOperateLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOperateModule operate_module获取 操作模块
func (obj *_LgChannelOperateLogMgr) WithOperateModule(operateModule string) Option {
	return optionFunc(func(o *options) { o.query["operate_module"] = operateModule })
}

// WithCheckCode check_code获取 渠道代码/系统服务商代码/渠道id
func (obj *_LgChannelOperateLogMgr) WithCheckCode(checkCode string) Option {
	return optionFunc(func(o *options) { o.query["check_code"] = checkCode })
}

// WithOperateType operate_type获取 操作类型
func (obj *_LgChannelOperateLogMgr) WithOperateType(operateType string) Option {
	return optionFunc(func(o *options) { o.query["operate_type"] = operateType })
}

// WithOperateDescription operate_description获取 操作描述
func (obj *_LgChannelOperateLogMgr) WithOperateDescription(operateDescription string) Option {
	return optionFunc(func(o *options) { o.query["operate_description"] = operateDescription })
}

// WithUpdateParams update_params获取 更新参数
func (obj *_LgChannelOperateLogMgr) WithUpdateParams(updateParams string) Option {
	return optionFunc(func(o *options) { o.query["update_params"] = updateParams })
}

// WithOperateIP operate_ip获取 操作IP
func (obj *_LgChannelOperateLogMgr) WithOperateIP(operateIP string) Option {
	return optionFunc(func(o *options) { o.query["operate_ip"] = operateIP })
}

// WithCreateUser create_user获取 操作人id
func (obj *_LgChannelOperateLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 操作时间
func (obj *_LgChannelOperateLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithVersion version获取 乐观锁
func (obj *_LgChannelOperateLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgChannelOperateLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgChannelOperateLogMgr) GetByOption(opts ...Option) (result LgChannelOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgChannelOperateLogMgr) GetByOptions(opts ...Option) (results []*LgChannelOperateLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgChannelOperateLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgChannelOperateLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where(options.query)
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
func (obj *_LgChannelOperateLogMgr) GetFromID(id int) (result LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgChannelOperateLogMgr) GetBatchFromID(ids []int) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOperateModule 通过operate_module获取内容 操作模块
func (obj *_LgChannelOperateLogMgr) GetFromOperateModule(operateModule string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`operate_module` = ?", operateModule).Find(&results).Error

	return
}

// GetBatchFromOperateModule 批量查找 操作模块
func (obj *_LgChannelOperateLogMgr) GetBatchFromOperateModule(operateModules []string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`operate_module` IN (?)", operateModules).Find(&results).Error

	return
}

// GetFromCheckCode 通过check_code获取内容 渠道代码/系统服务商代码/渠道id
func (obj *_LgChannelOperateLogMgr) GetFromCheckCode(checkCode string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`check_code` = ?", checkCode).Find(&results).Error

	return
}

// GetBatchFromCheckCode 批量查找 渠道代码/系统服务商代码/渠道id
func (obj *_LgChannelOperateLogMgr) GetBatchFromCheckCode(checkCodes []string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`check_code` IN (?)", checkCodes).Find(&results).Error

	return
}

// GetFromOperateType 通过operate_type获取内容 操作类型
func (obj *_LgChannelOperateLogMgr) GetFromOperateType(operateType string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`operate_type` = ?", operateType).Find(&results).Error

	return
}

// GetBatchFromOperateType 批量查找 操作类型
func (obj *_LgChannelOperateLogMgr) GetBatchFromOperateType(operateTypes []string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`operate_type` IN (?)", operateTypes).Find(&results).Error

	return
}

// GetFromOperateDescription 通过operate_description获取内容 操作描述
func (obj *_LgChannelOperateLogMgr) GetFromOperateDescription(operateDescription string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`operate_description` = ?", operateDescription).Find(&results).Error

	return
}

// GetBatchFromOperateDescription 批量查找 操作描述
func (obj *_LgChannelOperateLogMgr) GetBatchFromOperateDescription(operateDescriptions []string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`operate_description` IN (?)", operateDescriptions).Find(&results).Error

	return
}

// GetFromUpdateParams 通过update_params获取内容 更新参数
func (obj *_LgChannelOperateLogMgr) GetFromUpdateParams(updateParams string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`update_params` = ?", updateParams).Find(&results).Error

	return
}

// GetBatchFromUpdateParams 批量查找 更新参数
func (obj *_LgChannelOperateLogMgr) GetBatchFromUpdateParams(updateParamss []string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`update_params` IN (?)", updateParamss).Find(&results).Error

	return
}

// GetFromOperateIP 通过operate_ip获取内容 操作IP
func (obj *_LgChannelOperateLogMgr) GetFromOperateIP(operateIP string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`operate_ip` = ?", operateIP).Find(&results).Error

	return
}

// GetBatchFromOperateIP 批量查找 操作IP
func (obj *_LgChannelOperateLogMgr) GetBatchFromOperateIP(operateIPs []string) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`operate_ip` IN (?)", operateIPs).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 操作人id
func (obj *_LgChannelOperateLogMgr) GetFromCreateUser(createUser int) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 操作人id
func (obj *_LgChannelOperateLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 操作时间
func (obj *_LgChannelOperateLogMgr) GetFromCreateTime(createTime time.Time) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 操作时间
func (obj *_LgChannelOperateLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgChannelOperateLogMgr) GetFromVersion(version int) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgChannelOperateLogMgr) GetBatchFromVersion(versions []int) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgChannelOperateLogMgr) GetFromDeleted(deleted int) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgChannelOperateLogMgr) GetBatchFromDeleted(deleteds []int) (results []*LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgChannelOperateLogMgr) FetchByPrimaryKey(id int) (result LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgChannelOperateLogIDUIndex primary or index 获取唯一内容
func (obj *_LgChannelOperateLogMgr) FetchUniqueByLgChannelOperateLogIDUIndex(id int) (result LgChannelOperateLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgChannelOperateLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
