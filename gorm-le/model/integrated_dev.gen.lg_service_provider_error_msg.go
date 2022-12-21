package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgServiceProviderErrorMsgMgr struct {
	*_BaseMgr
}

// LgServiceProviderErrorMsgMgr open func
func LgServiceProviderErrorMsgMgr(db *gorm.DB) *_LgServiceProviderErrorMsgMgr {
	if db == nil {
		panic(fmt.Errorf("LgServiceProviderErrorMsgMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgServiceProviderErrorMsgMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_service_provider_error_msg"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgServiceProviderErrorMsgMgr) GetTableName() string {
	return "lg_service_provider_error_msg"
}

// Reset 重置gorm会话
func (obj *_LgServiceProviderErrorMsgMgr) Reset() *_LgServiceProviderErrorMsgMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgServiceProviderErrorMsgMgr) Get() (result LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_LgServiceProviderErrorMsgMgr) Gets() (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgServiceProviderErrorMsgMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_LgServiceProviderErrorMsgMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithServiceSystemCode service_system_code获取 系统服务商代码
func (obj *_LgServiceProviderErrorMsgMgr) WithServiceSystemCode(serviceSystemCode string) Option {
	return optionFunc(func(o *options) { o.query["service_system_code"] = serviceSystemCode })
}

// WithServiceSystemName service_system_name获取 系统服务商名称
func (obj *_LgServiceProviderErrorMsgMgr) WithServiceSystemName(serviceSystemName string) Option {
	return optionFunc(func(o *options) { o.query["service_system_name"] = serviceSystemName })
}

// WithServiceErrorCode service_error_code获取 服务商错误
func (obj *_LgServiceProviderErrorMsgMgr) WithServiceErrorCode(serviceErrorCode string) Option {
	return optionFunc(func(o *options) { o.query["service_error_code"] = serviceErrorCode })
}

// WithTranslateInfo translate_info获取 翻译信息
func (obj *_LgServiceProviderErrorMsgMgr) WithTranslateInfo(translateInfo string) Option {
	return optionFunc(func(o *options) { o.query["translate_info"] = translateInfo })
}

// WithRemark remark获取 备注
func (obj *_LgServiceProviderErrorMsgMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateUser create_user获取 创建人id
func (obj *_LgServiceProviderErrorMsgMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgServiceProviderErrorMsgMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_LgServiceProviderErrorMsgMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgServiceProviderErrorMsgMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_LgServiceProviderErrorMsgMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgServiceProviderErrorMsgMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgServiceProviderErrorMsgMgr) GetByOption(opts ...Option) (result LgServiceProviderErrorMsg, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgServiceProviderErrorMsgMgr) GetByOptions(opts ...Option) (results []*LgServiceProviderErrorMsg, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_LgServiceProviderErrorMsgMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgServiceProviderErrorMsg, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where(options.query)
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
func (obj *_LgServiceProviderErrorMsgMgr) GetFromID(id int) (result LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromID(ids []int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromServiceSystemCode 通过service_system_code获取内容 系统服务商代码
func (obj *_LgServiceProviderErrorMsgMgr) GetFromServiceSystemCode(serviceSystemCode string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_system_code` = ?", serviceSystemCode).Find(&results).Error

	return
}

// GetBatchFromServiceSystemCode 批量查找 系统服务商代码
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromServiceSystemCode(serviceSystemCodes []string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_system_code` IN (?)", serviceSystemCodes).Find(&results).Error

	return
}

// GetFromServiceSystemName 通过service_system_name获取内容 系统服务商名称
func (obj *_LgServiceProviderErrorMsgMgr) GetFromServiceSystemName(serviceSystemName string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_system_name` = ?", serviceSystemName).Find(&results).Error

	return
}

// GetBatchFromServiceSystemName 批量查找 系统服务商名称
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromServiceSystemName(serviceSystemNames []string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_system_name` IN (?)", serviceSystemNames).Find(&results).Error

	return
}

// GetFromServiceErrorCode 通过service_error_code获取内容 服务商错误
func (obj *_LgServiceProviderErrorMsgMgr) GetFromServiceErrorCode(serviceErrorCode string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_error_code` = ?", serviceErrorCode).Find(&results).Error

	return
}

// GetBatchFromServiceErrorCode 批量查找 服务商错误
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromServiceErrorCode(serviceErrorCodes []string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_error_code` IN (?)", serviceErrorCodes).Find(&results).Error

	return
}

// GetFromTranslateInfo 通过translate_info获取内容 翻译信息
func (obj *_LgServiceProviderErrorMsgMgr) GetFromTranslateInfo(translateInfo string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`translate_info` = ?", translateInfo).Find(&results).Error

	return
}

// GetBatchFromTranslateInfo 批量查找 翻译信息
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromTranslateInfo(translateInfos []string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`translate_info` IN (?)", translateInfos).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_LgServiceProviderErrorMsgMgr) GetFromRemark(remark string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromRemark(remarks []string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人id
func (obj *_LgServiceProviderErrorMsgMgr) GetFromCreateUser(createUser int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人id
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgServiceProviderErrorMsgMgr) GetFromCreateTime(createTime time.Time) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_LgServiceProviderErrorMsgMgr) GetFromUpdateUser(updateUser int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgServiceProviderErrorMsgMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgServiceProviderErrorMsgMgr) GetFromVersion(version int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromVersion(versions []int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgServiceProviderErrorMsgMgr) GetFromDeleted(deleted int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgServiceProviderErrorMsgMgr) GetBatchFromDeleted(deleteds []int) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgServiceProviderErrorMsgMgr) FetchByPrimaryKey(id int) (result LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByLgServiceProviderErrorMsgIDUIndex primary or index 获取唯一内容
func (obj *_LgServiceProviderErrorMsgMgr) FetchUniqueByLgServiceProviderErrorMsgIDUIndex(id int) (result LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueIndexByLgServiceErrorCodeServiceSystemCodeUIndex primary or index 获取唯一内容
func (obj *_LgServiceProviderErrorMsgMgr) FetchUniqueIndexByLgServiceErrorCodeServiceSystemCodeUIndex(serviceSystemCode string, serviceErrorCode string) (result LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_system_code` = ? AND `service_error_code` = ?", serviceSystemCode, serviceErrorCode).First(&result).Error

	return
}

// FetchIndexByLgServiceProviderErrorMsgServiceSystemNameIndex  获取多个内容
func (obj *_LgServiceProviderErrorMsgMgr) FetchIndexByLgServiceProviderErrorMsgServiceSystemNameIndex(serviceSystemCode string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_system_code` = ?", serviceSystemCode).Find(&results).Error

	return
}

// FetchIndexByLgServiceProviderErrorMsgServiceSystemNameIndex2  获取多个内容
func (obj *_LgServiceProviderErrorMsgMgr) FetchIndexByLgServiceProviderErrorMsgServiceSystemNameIndex2(serviceSystemName string) (results []*LgServiceProviderErrorMsg, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgServiceProviderErrorMsg{}).Where("`service_system_name` = ?", serviceSystemName).Find(&results).Error

	return
}
