package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MMessageMgr struct {
	*_BaseMgr
}

// MMessageMgr open func
func MMessageMgr(db *gorm.DB) *_MMessageMgr {
	if db == nil {
		panic(fmt.Errorf("MMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("m_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MMessageMgr) GetTableName() string {
	return "m_message"
}

// Reset 重置gorm会话
func (obj *_MMessageMgr) Reset() *_MMessageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MMessageMgr) Get() (result MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MMessageMgr) Gets() (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MMessageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MMessage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_MMessageMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 标题
func (obj *_MMessageMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithMsgType msg_type获取 消息类型
func (obj *_MMessageMgr) WithMsgType(msgType int) Option {
	return optionFunc(func(o *options) { o.query["msg_type"] = msgType })
}

// WithMsgTypeDesc msg_type_desc获取 类型描述
func (obj *_MMessageMgr) WithMsgTypeDesc(msgTypeDesc string) Option {
	return optionFunc(func(o *options) { o.query["msg_type_desc"] = msgTypeDesc })
}

// WithMsgData msg_data获取 消息内容
func (obj *_MMessageMgr) WithMsgData(msgData string) Option {
	return optionFunc(func(o *options) { o.query["msg_data"] = msgData })
}

// WithIsPop is_pop获取 是否弹窗
func (obj *_MMessageMgr) WithIsPop(isPop int) Option {
	return optionFunc(func(o *options) { o.query["is_pop"] = isPop })
}

// WithExpirationTime expiration_time获取 过期时间
func (obj *_MMessageMgr) WithExpirationTime(expirationTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expiration_time"] = expirationTime })
}

// WithRemark remark获取 备注
func (obj *_MMessageMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_MMessageMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_MMessageMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_MMessageMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_MMessageMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_MMessageMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_MMessageMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithBusinessNumber business_number获取 业务单号
func (obj *_MMessageMgr) WithBusinessNumber(businessNumber string) Option {
	return optionFunc(func(o *options) { o.query["business_number"] = businessNumber })
}

// GetByOption 功能选项模式获取
func (obj *_MMessageMgr) GetByOption(opts ...Option) (result MMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MMessageMgr) GetByOptions(opts ...Option) (results []*MMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MMessageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MMessage, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where(options.query)
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
func (obj *_MMessageMgr) GetFromID(id int) (result MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_MMessageMgr) GetBatchFromID(ids []int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_MMessageMgr) GetFromTitle(title string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_MMessageMgr) GetBatchFromTitle(titles []string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromMsgType 通过msg_type获取内容 消息类型
func (obj *_MMessageMgr) GetFromMsgType(msgType int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`msg_type` = ?", msgType).Find(&results).Error

	return
}

// GetBatchFromMsgType 批量查找 消息类型
func (obj *_MMessageMgr) GetBatchFromMsgType(msgTypes []int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`msg_type` IN (?)", msgTypes).Find(&results).Error

	return
}

// GetFromMsgTypeDesc 通过msg_type_desc获取内容 类型描述
func (obj *_MMessageMgr) GetFromMsgTypeDesc(msgTypeDesc string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`msg_type_desc` = ?", msgTypeDesc).Find(&results).Error

	return
}

// GetBatchFromMsgTypeDesc 批量查找 类型描述
func (obj *_MMessageMgr) GetBatchFromMsgTypeDesc(msgTypeDescs []string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`msg_type_desc` IN (?)", msgTypeDescs).Find(&results).Error

	return
}

// GetFromMsgData 通过msg_data获取内容 消息内容
func (obj *_MMessageMgr) GetFromMsgData(msgData string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`msg_data` = ?", msgData).Find(&results).Error

	return
}

// GetBatchFromMsgData 批量查找 消息内容
func (obj *_MMessageMgr) GetBatchFromMsgData(msgDatas []string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`msg_data` IN (?)", msgDatas).Find(&results).Error

	return
}

// GetFromIsPop 通过is_pop获取内容 是否弹窗
func (obj *_MMessageMgr) GetFromIsPop(isPop int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`is_pop` = ?", isPop).Find(&results).Error

	return
}

// GetBatchFromIsPop 批量查找 是否弹窗
func (obj *_MMessageMgr) GetBatchFromIsPop(isPops []int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`is_pop` IN (?)", isPops).Find(&results).Error

	return
}

// GetFromExpirationTime 通过expiration_time获取内容 过期时间
func (obj *_MMessageMgr) GetFromExpirationTime(expirationTime time.Time) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`expiration_time` = ?", expirationTime).Find(&results).Error

	return
}

// GetBatchFromExpirationTime 批量查找 过期时间
func (obj *_MMessageMgr) GetBatchFromExpirationTime(expirationTimes []time.Time) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`expiration_time` IN (?)", expirationTimes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_MMessageMgr) GetFromRemark(remark string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_MMessageMgr) GetBatchFromRemark(remarks []string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_MMessageMgr) GetFromCreateTime(createTime time.Time) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_MMessageMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_MMessageMgr) GetFromCreateUser(createUser int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_MMessageMgr) GetBatchFromCreateUser(createUsers []int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_MMessageMgr) GetFromUpdateTime(updateTime time.Time) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_MMessageMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_MMessageMgr) GetFromUpdateUser(updateUser int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_MMessageMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_MMessageMgr) GetFromVersion(version int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_MMessageMgr) GetBatchFromVersion(versions []int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_MMessageMgr) GetFromDeleted(deleted int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_MMessageMgr) GetBatchFromDeleted(deleteds []int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromBusinessNumber 通过business_number获取内容 业务单号
func (obj *_MMessageMgr) GetFromBusinessNumber(businessNumber string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`business_number` = ?", businessNumber).Find(&results).Error

	return
}

// GetBatchFromBusinessNumber 批量查找 业务单号
func (obj *_MMessageMgr) GetBatchFromBusinessNumber(businessNumbers []string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`business_number` IN (?)", businessNumbers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_MMessageMgr) FetchByPrimaryKey(id int) (result MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByMMessageTitleIndex  获取多个内容
func (obj *_MMessageMgr) FetchIndexByMMessageTitleIndex(title string) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// FetchIndexByMMessageIsPopIndex  获取多个内容
func (obj *_MMessageMgr) FetchIndexByMMessageIsPopIndex(isPop int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`is_pop` = ?", isPop).Find(&results).Error

	return
}

// FetchIndexByMMessageExpirationTimeIndex  获取多个内容
func (obj *_MMessageMgr) FetchIndexByMMessageExpirationTimeIndex(expirationTime time.Time) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`expiration_time` = ?", expirationTime).Find(&results).Error

	return
}

// FetchIndexByMMessageCreateTimeIndex  获取多个内容
func (obj *_MMessageMgr) FetchIndexByMMessageCreateTimeIndex(createTime time.Time) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByMMessageCreateUserIndex  获取多个内容
func (obj *_MMessageMgr) FetchIndexByMMessageCreateUserIndex(createUser int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// FetchIndexByMMessageDeletedIndex  获取多个内容
func (obj *_MMessageMgr) FetchIndexByMMessageDeletedIndex(deleted int) (results []*MMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessage{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}
