package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsExpressSourceMgr struct {
	*_BaseMgr
}

// WmsExpressSourceMgr open func
func WmsExpressSourceMgr(db *gorm.DB) *_WmsExpressSourceMgr {
	if db == nil {
		panic(fmt.Errorf("WmsExpressSourceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsExpressSourceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_express_source"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsExpressSourceMgr) GetTableName() string {
	return "wms_express_source"
}

// Reset 重置gorm会话
func (obj *_WmsExpressSourceMgr) Reset() *_WmsExpressSourceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsExpressSourceMgr) Get() (result WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsExpressSourceMgr) Gets() (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsExpressSourceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_WmsExpressSourceMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithExpressCode express_code获取 快递代码
func (obj *_WmsExpressSourceMgr) WithExpressCode(expressCode string) Option {
	return optionFunc(func(o *options) { o.query["express_code"] = expressCode })
}

// WithExpressName express_name获取 快递名称
func (obj *_WmsExpressSourceMgr) WithExpressName(expressName string) Option {
	return optionFunc(func(o *options) { o.query["express_name"] = expressName })
}

// WithRemark remark获取 备注
func (obj *_WmsExpressSourceMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsExpressSourceMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建人
func (obj *_WmsExpressSourceMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsExpressSourceMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新人
func (obj *_WmsExpressSourceMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsExpressSourceMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsExpressSourceMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsExpressSourceMgr) GetByOption(opts ...Option) (result WmsExpressSource, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsExpressSourceMgr) GetByOptions(opts ...Option) (results []*WmsExpressSource, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsExpressSourceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsExpressSource, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where(options.query)
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
func (obj *_WmsExpressSourceMgr) GetFromID(id int) (result WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_WmsExpressSourceMgr) GetBatchFromID(ids []int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromExpressCode 通过express_code获取内容 快递代码
func (obj *_WmsExpressSourceMgr) GetFromExpressCode(expressCode string) (result WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`express_code` = ?", expressCode).First(&result).Error

	return
}

// GetBatchFromExpressCode 批量查找 快递代码
func (obj *_WmsExpressSourceMgr) GetBatchFromExpressCode(expressCodes []string) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`express_code` IN (?)", expressCodes).Find(&results).Error

	return
}

// GetFromExpressName 通过express_name获取内容 快递名称
func (obj *_WmsExpressSourceMgr) GetFromExpressName(expressName string) (result WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`express_name` = ?", expressName).First(&result).Error

	return
}

// GetBatchFromExpressName 批量查找 快递名称
func (obj *_WmsExpressSourceMgr) GetBatchFromExpressName(expressNames []string) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`express_name` IN (?)", expressNames).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WmsExpressSourceMgr) GetFromRemark(remark string) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WmsExpressSourceMgr) GetBatchFromRemark(remarks []string) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsExpressSourceMgr) GetFromCreateTime(createTime time.Time) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsExpressSourceMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建人
func (obj *_WmsExpressSourceMgr) GetFromCreateUser(createUser int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建人
func (obj *_WmsExpressSourceMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsExpressSourceMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsExpressSourceMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新人
func (obj *_WmsExpressSourceMgr) GetFromUpdateUser(updateUser int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新人
func (obj *_WmsExpressSourceMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsExpressSourceMgr) GetFromVersion(version int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsExpressSourceMgr) GetBatchFromVersion(versions []int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsExpressSourceMgr) GetFromDeleted(deleted int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsExpressSourceMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsExpressSourceMgr) FetchByPrimaryKey(id int) (result WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByWmsExpressSourceIDUIndex primary or index 获取唯一内容
func (obj *_WmsExpressSourceMgr) FetchUniqueByWmsExpressSourceIDUIndex(id int) (result WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByWmsExpressSourceExpressCodeUIndex primary or index 获取唯一内容
func (obj *_WmsExpressSourceMgr) FetchUniqueByWmsExpressSourceExpressCodeUIndex(expressCode string) (result WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`express_code` = ?", expressCode).First(&result).Error

	return
}

// FetchUniqueByWmsExpressSourceExpressNameUIndex primary or index 获取唯一内容
func (obj *_WmsExpressSourceMgr) FetchUniqueByWmsExpressSourceExpressNameUIndex(expressName string) (result WmsExpressSource, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsExpressSource{}).Where("`express_name` = ?", expressName).First(&result).Error

	return
}
