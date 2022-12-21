package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UFileExportMgr struct {
	*_BaseMgr
}

// UFileExportMgr open func
func UFileExportMgr(db *gorm.DB) *_UFileExportMgr {
	if db == nil {
		panic(fmt.Errorf("UFileExportMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UFileExportMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_file_export"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UFileExportMgr) GetTableName() string {
	return "u_file_export"
}

// Reset 重置gorm会话
func (obj *_UFileExportMgr) Reset() *_UFileExportMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UFileExportMgr) Get() (result UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UFileExportMgr) Gets() (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UFileExportMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_UFileExportMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOperator operator获取 操作人
func (obj *_UFileExportMgr) WithOperator(operator string) Option {
	return optionFunc(func(o *options) { o.query["operator"] = operator })
}

// WithOperationModule operation_module获取 操作模块
func (obj *_UFileExportMgr) WithOperationModule(operationModule string) Option {
	return optionFunc(func(o *options) { o.query["operation_module"] = operationModule })
}

// WithFileType file_type获取 文件类型
func (obj *_UFileExportMgr) WithFileType(fileType string) Option {
	return optionFunc(func(o *options) { o.query["file_type"] = fileType })
}

// WithFileURL file_url获取 访问路径
func (obj *_UFileExportMgr) WithFileURL(fileURL string) Option {
	return optionFunc(func(o *options) { o.query["file_url"] = fileURL })
}

// WithOssPath oss_path获取 上传oos路径
func (obj *_UFileExportMgr) WithOssPath(ossPath string) Option {
	return optionFunc(func(o *options) { o.query["oss_path"] = ossPath })
}

// WithVerifyName verify_name获取 审核人
func (obj *_UFileExportMgr) WithVerifyName(verifyName string) Option {
	return optionFunc(func(o *options) { o.query["verify_name"] = verifyName })
}

// WithVerifyID verify_id获取 审核人id
func (obj *_UFileExportMgr) WithVerifyID(verifyID int) Option {
	return optionFunc(func(o *options) { o.query["verify_id"] = verifyID })
}

// WithIsVerify is_verify获取 是否需要审核：0-否，1-是
func (obj *_UFileExportMgr) WithIsVerify(isVerify int) Option {
	return optionFunc(func(o *options) { o.query["is_verify"] = isVerify })
}

// WithVerifyStatus verify_status获取 审核状态：0-未审核，1-已审核
func (obj *_UFileExportMgr) WithVerifyStatus(verifyStatus int) Option {
	return optionFunc(func(o *options) { o.query["verify_status"] = verifyStatus })
}

// WithExportResult export_result获取 导出结果
func (obj *_UFileExportMgr) WithExportResult(exportResult string) Option {
	return optionFunc(func(o *options) { o.query["export_result"] = exportResult })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UFileExportMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 操作人
func (obj *_UFileExportMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UFileExportMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithVersion version获取 乐观锁
func (obj *_UFileExportMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_UFileExportMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithDownloadsNumber downloads_number获取 下载次数
func (obj *_UFileExportMgr) WithDownloadsNumber(downloadsNumber int) Option {
	return optionFunc(func(o *options) { o.query["downloads_number"] = downloadsNumber })
}

// WithFileName file_name获取 文件名
func (obj *_UFileExportMgr) WithFileName(fileName string) Option {
	return optionFunc(func(o *options) { o.query["file_name"] = fileName })
}

// WithGenerateTime generate_time获取 生成时间
func (obj *_UFileExportMgr) WithGenerateTime(generateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["generate_time"] = generateTime })
}

// GetByOption 功能选项模式获取
func (obj *_UFileExportMgr) GetByOption(opts ...Option) (result UFileExport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UFileExportMgr) GetByOptions(opts ...Option) (results []*UFileExport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UFileExportMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UFileExport, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where(options.query)
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
func (obj *_UFileExportMgr) GetFromID(id int) (result UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_UFileExportMgr) GetBatchFromID(ids []int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOperator 通过operator获取内容 操作人
func (obj *_UFileExportMgr) GetFromOperator(operator string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`operator` = ?", operator).Find(&results).Error

	return
}

// GetBatchFromOperator 批量查找 操作人
func (obj *_UFileExportMgr) GetBatchFromOperator(operators []string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`operator` IN (?)", operators).Find(&results).Error

	return
}

// GetFromOperationModule 通过operation_module获取内容 操作模块
func (obj *_UFileExportMgr) GetFromOperationModule(operationModule string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`operation_module` = ?", operationModule).Find(&results).Error

	return
}

// GetBatchFromOperationModule 批量查找 操作模块
func (obj *_UFileExportMgr) GetBatchFromOperationModule(operationModules []string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`operation_module` IN (?)", operationModules).Find(&results).Error

	return
}

// GetFromFileType 通过file_type获取内容 文件类型
func (obj *_UFileExportMgr) GetFromFileType(fileType string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`file_type` = ?", fileType).Find(&results).Error

	return
}

// GetBatchFromFileType 批量查找 文件类型
func (obj *_UFileExportMgr) GetBatchFromFileType(fileTypes []string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`file_type` IN (?)", fileTypes).Find(&results).Error

	return
}

// GetFromFileURL 通过file_url获取内容 访问路径
func (obj *_UFileExportMgr) GetFromFileURL(fileURL string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`file_url` = ?", fileURL).Find(&results).Error

	return
}

// GetBatchFromFileURL 批量查找 访问路径
func (obj *_UFileExportMgr) GetBatchFromFileURL(fileURLs []string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`file_url` IN (?)", fileURLs).Find(&results).Error

	return
}

// GetFromOssPath 通过oss_path获取内容 上传oos路径
func (obj *_UFileExportMgr) GetFromOssPath(ossPath string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`oss_path` = ?", ossPath).Find(&results).Error

	return
}

// GetBatchFromOssPath 批量查找 上传oos路径
func (obj *_UFileExportMgr) GetBatchFromOssPath(ossPaths []string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`oss_path` IN (?)", ossPaths).Find(&results).Error

	return
}

// GetFromVerifyName 通过verify_name获取内容 审核人
func (obj *_UFileExportMgr) GetFromVerifyName(verifyName string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`verify_name` = ?", verifyName).Find(&results).Error

	return
}

// GetBatchFromVerifyName 批量查找 审核人
func (obj *_UFileExportMgr) GetBatchFromVerifyName(verifyNames []string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`verify_name` IN (?)", verifyNames).Find(&results).Error

	return
}

// GetFromVerifyID 通过verify_id获取内容 审核人id
func (obj *_UFileExportMgr) GetFromVerifyID(verifyID int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`verify_id` = ?", verifyID).Find(&results).Error

	return
}

// GetBatchFromVerifyID 批量查找 审核人id
func (obj *_UFileExportMgr) GetBatchFromVerifyID(verifyIDs []int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`verify_id` IN (?)", verifyIDs).Find(&results).Error

	return
}

// GetFromIsVerify 通过is_verify获取内容 是否需要审核：0-否，1-是
func (obj *_UFileExportMgr) GetFromIsVerify(isVerify int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`is_verify` = ?", isVerify).Find(&results).Error

	return
}

// GetBatchFromIsVerify 批量查找 是否需要审核：0-否，1-是
func (obj *_UFileExportMgr) GetBatchFromIsVerify(isVerifys []int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`is_verify` IN (?)", isVerifys).Find(&results).Error

	return
}

// GetFromVerifyStatus 通过verify_status获取内容 审核状态：0-未审核，1-已审核
func (obj *_UFileExportMgr) GetFromVerifyStatus(verifyStatus int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`verify_status` = ?", verifyStatus).Find(&results).Error

	return
}

// GetBatchFromVerifyStatus 批量查找 审核状态：0-未审核，1-已审核
func (obj *_UFileExportMgr) GetBatchFromVerifyStatus(verifyStatuss []int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`verify_status` IN (?)", verifyStatuss).Find(&results).Error

	return
}

// GetFromExportResult 通过export_result获取内容 导出结果
func (obj *_UFileExportMgr) GetFromExportResult(exportResult string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`export_result` = ?", exportResult).Find(&results).Error

	return
}

// GetBatchFromExportResult 批量查找 导出结果
func (obj *_UFileExportMgr) GetBatchFromExportResult(exportResults []string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`export_result` IN (?)", exportResults).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UFileExportMgr) GetFromCreateTime(createTime time.Time) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UFileExportMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 操作人
func (obj *_UFileExportMgr) GetFromCreateUser(createUser int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 操作人
func (obj *_UFileExportMgr) GetBatchFromCreateUser(createUsers []int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UFileExportMgr) GetFromUpdateTime(updateTime time.Time) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UFileExportMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UFileExportMgr) GetFromVersion(version int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UFileExportMgr) GetBatchFromVersion(versions []int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_UFileExportMgr) GetFromDeleted(deleted int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_UFileExportMgr) GetBatchFromDeleted(deleteds []int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromDownloadsNumber 通过downloads_number获取内容 下载次数
func (obj *_UFileExportMgr) GetFromDownloadsNumber(downloadsNumber int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`downloads_number` = ?", downloadsNumber).Find(&results).Error

	return
}

// GetBatchFromDownloadsNumber 批量查找 下载次数
func (obj *_UFileExportMgr) GetBatchFromDownloadsNumber(downloadsNumbers []int) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`downloads_number` IN (?)", downloadsNumbers).Find(&results).Error

	return
}

// GetFromFileName 通过file_name获取内容 文件名
func (obj *_UFileExportMgr) GetFromFileName(fileName string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`file_name` = ?", fileName).Find(&results).Error

	return
}

// GetBatchFromFileName 批量查找 文件名
func (obj *_UFileExportMgr) GetBatchFromFileName(fileNames []string) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`file_name` IN (?)", fileNames).Find(&results).Error

	return
}

// GetFromGenerateTime 通过generate_time获取内容 生成时间
func (obj *_UFileExportMgr) GetFromGenerateTime(generateTime time.Time) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`generate_time` = ?", generateTime).Find(&results).Error

	return
}

// GetBatchFromGenerateTime 批量查找 生成时间
func (obj *_UFileExportMgr) GetBatchFromGenerateTime(generateTimes []time.Time) (results []*UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`generate_time` IN (?)", generateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UFileExportMgr) FetchByPrimaryKey(id int) (result UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUExportIDUIndex primary or index 获取唯一内容
func (obj *_UFileExportMgr) FetchUniqueByUExportIDUIndex(id int) (result UFileExport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UFileExport{}).Where("`id` = ?", id).First(&result).Error

	return
}
