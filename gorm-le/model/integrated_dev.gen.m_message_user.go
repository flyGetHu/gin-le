package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _MMessageUserMgr struct {
	*_BaseMgr
}

// MMessageUserMgr open func
func MMessageUserMgr(db *gorm.DB) *_MMessageUserMgr {
	if db == nil {
		panic(fmt.Errorf("MMessageUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_MMessageUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("m_message_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_MMessageUserMgr) GetTableName() string {
	return "m_message_user"
}

// Reset 重置gorm会话
func (obj *_MMessageUserMgr) Reset() *_MMessageUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_MMessageUserMgr) Get() (result MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_MMessageUserMgr) Gets() (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_MMessageUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMessageID message_id获取
func (obj *_MMessageUserMgr) WithMessageID(messageID int) Option {
	return optionFunc(func(o *options) { o.query["message_id"] = messageID })
}

// WithTitle title获取 消息标题
func (obj *_MMessageUserMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithMsgData msg_data获取 消息内容
func (obj *_MMessageUserMgr) WithMsgData(msgData string) Option {
	return optionFunc(func(o *options) { o.query["msg_data"] = msgData })
}

// WithMsgType msg_type获取 消息类型
func (obj *_MMessageUserMgr) WithMsgType(msgType int) Option {
	return optionFunc(func(o *options) { o.query["msg_type"] = msgType })
}

// WithMsgTypeDesc msg_type_desc获取 消息类型中文
func (obj *_MMessageUserMgr) WithMsgTypeDesc(msgTypeDesc string) Option {
	return optionFunc(func(o *options) { o.query["msg_type_desc"] = msgTypeDesc })
}

// WithIsPop is_pop获取 是否弹窗
func (obj *_MMessageUserMgr) WithIsPop(isPop int) Option {
	return optionFunc(func(o *options) { o.query["is_pop"] = isPop })
}

// WithExpirationTime expiration_time获取 过期时间
func (obj *_MMessageUserMgr) WithExpirationTime(expirationTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["expiration_time"] = expirationTime })
}

// WithReceiveID receive_id获取 接收人
func (obj *_MMessageUserMgr) WithReceiveID(receiveID int) Option {
	return optionFunc(func(o *options) { o.query["receive_id"] = receiveID })
}

// WithReceiveType receive_type获取 接收人类型
func (obj *_MMessageUserMgr) WithReceiveType(receiveType int) Option {
	return optionFunc(func(o *options) { o.query["receive_type"] = receiveType })
}

// WithReceiveTypeChinese receive_type_chinese获取 接收人中文
func (obj *_MMessageUserMgr) WithReceiveTypeChinese(receiveTypeChinese string) Option {
	return optionFunc(func(o *options) { o.query["receive_type_chinese"] = receiveTypeChinese })
}

// WithIsRead is_read获取 是否已读
func (obj *_MMessageUserMgr) WithIsRead(isRead int) Option {
	return optionFunc(func(o *options) { o.query["is_read"] = isRead })
}

// WithReadTime read_time获取 已读时间
func (obj *_MMessageUserMgr) WithReadTime(readTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["read_time"] = readTime })
}

// WithCreateTime create_time获取 创建时间
func (obj *_MMessageUserMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_MMessageUserMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithVersion version获取 乐观锁
func (obj *_MMessageUserMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_MMessageUserMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithBusinessNumber business_number获取 业务单号
func (obj *_MMessageUserMgr) WithBusinessNumber(businessNumber string) Option {
	return optionFunc(func(o *options) { o.query["business_number"] = businessNumber })
}

// GetByOption 功能选项模式获取
func (obj *_MMessageUserMgr) GetByOption(opts ...Option) (result MMessageUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_MMessageUserMgr) GetByOptions(opts ...Option) (results []*MMessageUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_MMessageUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]MMessageUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where(options.query)
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

// GetFromMessageID 通过message_id获取内容
func (obj *_MMessageUserMgr) GetFromMessageID(messageID int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`message_id` = ?", messageID).Find(&results).Error

	return
}

// GetBatchFromMessageID 批量查找
func (obj *_MMessageUserMgr) GetBatchFromMessageID(messageIDs []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`message_id` IN (?)", messageIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 消息标题
func (obj *_MMessageUserMgr) GetFromTitle(title string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 消息标题
func (obj *_MMessageUserMgr) GetBatchFromTitle(titles []string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromMsgData 通过msg_data获取内容 消息内容
func (obj *_MMessageUserMgr) GetFromMsgData(msgData string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`msg_data` = ?", msgData).Find(&results).Error

	return
}

// GetBatchFromMsgData 批量查找 消息内容
func (obj *_MMessageUserMgr) GetBatchFromMsgData(msgDatas []string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`msg_data` IN (?)", msgDatas).Find(&results).Error

	return
}

// GetFromMsgType 通过msg_type获取内容 消息类型
func (obj *_MMessageUserMgr) GetFromMsgType(msgType int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`msg_type` = ?", msgType).Find(&results).Error

	return
}

// GetBatchFromMsgType 批量查找 消息类型
func (obj *_MMessageUserMgr) GetBatchFromMsgType(msgTypes []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`msg_type` IN (?)", msgTypes).Find(&results).Error

	return
}

// GetFromMsgTypeDesc 通过msg_type_desc获取内容 消息类型中文
func (obj *_MMessageUserMgr) GetFromMsgTypeDesc(msgTypeDesc string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`msg_type_desc` = ?", msgTypeDesc).Find(&results).Error

	return
}

// GetBatchFromMsgTypeDesc 批量查找 消息类型中文
func (obj *_MMessageUserMgr) GetBatchFromMsgTypeDesc(msgTypeDescs []string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`msg_type_desc` IN (?)", msgTypeDescs).Find(&results).Error

	return
}

// GetFromIsPop 通过is_pop获取内容 是否弹窗
func (obj *_MMessageUserMgr) GetFromIsPop(isPop int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`is_pop` = ?", isPop).Find(&results).Error

	return
}

// GetBatchFromIsPop 批量查找 是否弹窗
func (obj *_MMessageUserMgr) GetBatchFromIsPop(isPops []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`is_pop` IN (?)", isPops).Find(&results).Error

	return
}

// GetFromExpirationTime 通过expiration_time获取内容 过期时间
func (obj *_MMessageUserMgr) GetFromExpirationTime(expirationTime time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`expiration_time` = ?", expirationTime).Find(&results).Error

	return
}

// GetBatchFromExpirationTime 批量查找 过期时间
func (obj *_MMessageUserMgr) GetBatchFromExpirationTime(expirationTimes []time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`expiration_time` IN (?)", expirationTimes).Find(&results).Error

	return
}

// GetFromReceiveID 通过receive_id获取内容 接收人
func (obj *_MMessageUserMgr) GetFromReceiveID(receiveID int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`receive_id` = ?", receiveID).Find(&results).Error

	return
}

// GetBatchFromReceiveID 批量查找 接收人
func (obj *_MMessageUserMgr) GetBatchFromReceiveID(receiveIDs []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`receive_id` IN (?)", receiveIDs).Find(&results).Error

	return
}

// GetFromReceiveType 通过receive_type获取内容 接收人类型
func (obj *_MMessageUserMgr) GetFromReceiveType(receiveType int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`receive_type` = ?", receiveType).Find(&results).Error

	return
}

// GetBatchFromReceiveType 批量查找 接收人类型
func (obj *_MMessageUserMgr) GetBatchFromReceiveType(receiveTypes []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`receive_type` IN (?)", receiveTypes).Find(&results).Error

	return
}

// GetFromReceiveTypeChinese 通过receive_type_chinese获取内容 接收人中文
func (obj *_MMessageUserMgr) GetFromReceiveTypeChinese(receiveTypeChinese string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`receive_type_chinese` = ?", receiveTypeChinese).Find(&results).Error

	return
}

// GetBatchFromReceiveTypeChinese 批量查找 接收人中文
func (obj *_MMessageUserMgr) GetBatchFromReceiveTypeChinese(receiveTypeChineses []string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`receive_type_chinese` IN (?)", receiveTypeChineses).Find(&results).Error

	return
}

// GetFromIsRead 通过is_read获取内容 是否已读
func (obj *_MMessageUserMgr) GetFromIsRead(isRead int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`is_read` = ?", isRead).Find(&results).Error

	return
}

// GetBatchFromIsRead 批量查找 是否已读
func (obj *_MMessageUserMgr) GetBatchFromIsRead(isReads []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`is_read` IN (?)", isReads).Find(&results).Error

	return
}

// GetFromReadTime 通过read_time获取内容 已读时间
func (obj *_MMessageUserMgr) GetFromReadTime(readTime time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`read_time` = ?", readTime).Find(&results).Error

	return
}

// GetBatchFromReadTime 批量查找 已读时间
func (obj *_MMessageUserMgr) GetBatchFromReadTime(readTimes []time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`read_time` IN (?)", readTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_MMessageUserMgr) GetFromCreateTime(createTime time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_MMessageUserMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_MMessageUserMgr) GetFromCreateUser(createUser int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_MMessageUserMgr) GetBatchFromCreateUser(createUsers []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_MMessageUserMgr) GetFromVersion(version int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_MMessageUserMgr) GetBatchFromVersion(versions []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_MMessageUserMgr) GetFromDeleted(deleted int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_MMessageUserMgr) GetBatchFromDeleted(deleteds []int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

// GetFromBusinessNumber 通过business_number获取内容 业务单号
func (obj *_MMessageUserMgr) GetFromBusinessNumber(businessNumber string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`business_number` = ?", businessNumber).Find(&results).Error

	return
}

// GetBatchFromBusinessNumber 批量查找 业务单号
func (obj *_MMessageUserMgr) GetBatchFromBusinessNumber(businessNumbers []string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`business_number` IN (?)", businessNumbers).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByMMessageUserMessageIDIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserMessageIDIndex(messageID int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`message_id` = ?", messageID).Find(&results).Error

	return
}

// FetchIndexByMMessageUserTitleIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserTitleIndex(title string) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// FetchIndexByMMessageUserMsgTypeIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserMsgTypeIndex(msgType int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`msg_type` = ?", msgType).Find(&results).Error

	return
}

// FetchIndexByMMessageUserIsPopIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserIsPopIndex(isPop int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`is_pop` = ?", isPop).Find(&results).Error

	return
}

// FetchIndexByMMessageUserExpirationTimeIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserExpirationTimeIndex(expirationTime time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`expiration_time` = ?", expirationTime).Find(&results).Error

	return
}

// FetchIndexByMMessageUserReceiveIDIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserReceiveIDIndex(receiveID int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`receive_id` = ?", receiveID).Find(&results).Error

	return
}

// FetchIndexByMMessageUserReceiveTypeIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserReceiveTypeIndex(receiveType int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`receive_type` = ?", receiveType).Find(&results).Error

	return
}

// FetchIndexByMMessageUserIsReadIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserIsReadIndex(isRead int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`is_read` = ?", isRead).Find(&results).Error

	return
}

// FetchIndexByMMessageUserReadTimeIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserReadTimeIndex(readTime time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`read_time` = ?", readTime).Find(&results).Error

	return
}

// FetchIndexByMMessageUserCreateTimeIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserCreateTimeIndex(createTime time.Time) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// FetchIndexByMMessageUserVersionIndex  获取多个内容
func (obj *_MMessageUserMgr) FetchIndexByMMessageUserVersionIndex(version int) (results []*MMessageUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(MMessageUser{}).Where("`version` = ?", version).Find(&results).Error

	return
}
