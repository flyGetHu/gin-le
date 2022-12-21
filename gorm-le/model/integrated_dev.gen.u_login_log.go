package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _ULoginLogMgr struct {
	*_BaseMgr
}

// ULoginLogMgr open func
func ULoginLogMgr(db *gorm.DB) *_ULoginLogMgr {
	if db == nil {
		panic(fmt.Errorf("ULoginLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ULoginLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_login_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ULoginLogMgr) GetTableName() string {
	return "u_login_log"
}

// Reset 重置gorm会话
func (obj *_ULoginLogMgr) Reset() *_ULoginLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ULoginLogMgr) Get() (result ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ULoginLogMgr) Gets() (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ULoginLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_ULoginLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 用户名
func (obj *_ULoginLogMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithIP ip获取 登录IP
func (obj *_ULoginLogMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// WithIPAddress ip_address获取 ip地址所在地区
func (obj *_ULoginLogMgr) WithIPAddress(ipAddress string) Option {
	return optionFunc(func(o *options) { o.query["ip_address"] = ipAddress })
}

// WithSource source获取 来源
func (obj *_ULoginLogMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithBrowserInfo browser_info获取 浏览器信息
func (obj *_ULoginLogMgr) WithBrowserInfo(browserInfo string) Option {
	return optionFunc(func(o *options) { o.query["browser_info"] = browserInfo })
}

// WithOsInfo os_info获取 操作系统信息
func (obj *_ULoginLogMgr) WithOsInfo(osInfo string) Option {
	return optionFunc(func(o *options) { o.query["os_info"] = osInfo })
}

// WithCreateTime create_time获取 创建时间
func (obj *_ULoginLogMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_ULoginLogMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_ULoginLogMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_ULoginLogMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_ULoginLogMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_ULoginLogMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_ULoginLogMgr) GetByOption(opts ...Option) (result ULoginLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ULoginLogMgr) GetByOptions(opts ...Option) (results []*ULoginLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ULoginLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ULoginLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where(options.query)
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

// GetFromID 通过id获取内容 ID
func (obj *_ULoginLogMgr) GetFromID(id int) (result ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_ULoginLogMgr) GetBatchFromID(ids []int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_ULoginLogMgr) GetFromUsername(username string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 用户名
func (obj *_ULoginLogMgr) GetBatchFromUsername(usernames []string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 登录IP
func (obj *_ULoginLogMgr) GetFromIP(ip string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找 登录IP
func (obj *_ULoginLogMgr) GetBatchFromIP(ips []string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

// GetFromIPAddress 通过ip_address获取内容 ip地址所在地区
func (obj *_ULoginLogMgr) GetFromIPAddress(ipAddress string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`ip_address` = ?", ipAddress).Find(&results).Error

	return
}

// GetBatchFromIPAddress 批量查找 ip地址所在地区
func (obj *_ULoginLogMgr) GetBatchFromIPAddress(ipAddresss []string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`ip_address` IN (?)", ipAddresss).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容 来源
func (obj *_ULoginLogMgr) GetFromSource(source string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找 来源
func (obj *_ULoginLogMgr) GetBatchFromSource(sources []string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromBrowserInfo 通过browser_info获取内容 浏览器信息
func (obj *_ULoginLogMgr) GetFromBrowserInfo(browserInfo string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`browser_info` = ?", browserInfo).Find(&results).Error

	return
}

// GetBatchFromBrowserInfo 批量查找 浏览器信息
func (obj *_ULoginLogMgr) GetBatchFromBrowserInfo(browserInfos []string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`browser_info` IN (?)", browserInfos).Find(&results).Error

	return
}

// GetFromOsInfo 通过os_info获取内容 操作系统信息
func (obj *_ULoginLogMgr) GetFromOsInfo(osInfo string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`os_info` = ?", osInfo).Find(&results).Error

	return
}

// GetBatchFromOsInfo 批量查找 操作系统信息
func (obj *_ULoginLogMgr) GetBatchFromOsInfo(osInfos []string) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`os_info` IN (?)", osInfos).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_ULoginLogMgr) GetFromCreateTime(createTime time.Time) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_ULoginLogMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_ULoginLogMgr) GetFromCreateUser(createUser int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_ULoginLogMgr) GetBatchFromCreateUser(createUsers []int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_ULoginLogMgr) GetFromUpdateTime(updateTime time.Time) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_ULoginLogMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_ULoginLogMgr) GetFromUpdateUser(updateUser int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_ULoginLogMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_ULoginLogMgr) GetFromVersion(version int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_ULoginLogMgr) GetBatchFromVersion(versions []int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_ULoginLogMgr) GetFromDeleted(deleted int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_ULoginLogMgr) GetBatchFromDeleted(deleteds []int) (results []*ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ULoginLogMgr) FetchByPrimaryKey(id int) (result ULoginLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ULoginLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
