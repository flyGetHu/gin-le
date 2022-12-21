package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgProviderChannelMgr struct {
	*_BaseMgr
}

// LgProviderChannelMgr open func
func LgProviderChannelMgr(db *gorm.DB) *_LgProviderChannelMgr {
	if db == nil {
		panic(fmt.Errorf("LgProviderChannelMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgProviderChannelMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_provider_channel"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgProviderChannelMgr) GetTableName() string {
	return "lg_provider_channel"
}

// Reset 重置gorm会话
func (obj *_LgProviderChannelMgr) Reset() *_LgProviderChannelMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgProviderChannelMgr) Get() (result LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", result.OrderNumPoolID).Find(&result.LgOrderNumPool).Error; err != nil { // 单号池主表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_LgProviderChannelMgr) Gets() (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgProviderChannelMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_LgProviderChannelMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithProviderID provider_id获取 系统服务商id
func (obj *_LgProviderChannelMgr) WithProviderID(providerID string) Option {
	return optionFunc(func(o *options) { o.query["provider_id"] = providerID })
}

// WithProviderChannelCode provider_channel_code获取 渠道代码
func (obj *_LgProviderChannelMgr) WithProviderChannelCode(providerChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_code"] = providerChannelCode })
}

// WithProviderChannelName provider_channel_name获取 服务商渠道名称
func (obj *_LgProviderChannelMgr) WithProviderChannelName(providerChannelName string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_name"] = providerChannelName })
}

// WithOrderNumPoolID order_num_pool_id获取 单号池id
func (obj *_LgProviderChannelMgr) WithOrderNumPoolID(orderNumPoolID int64) Option {
	return optionFunc(func(o *options) { o.query["order_num_pool_id"] = orderNumPoolID })
}

// WithOrderNumPoolName order_num_pool_name获取 单号池名称
func (obj *_LgProviderChannelMgr) WithOrderNumPoolName(orderNumPoolName string) Option {
	return optionFunc(func(o *options) { o.query["order_num_pool_name"] = orderNumPoolName })
}

// WithBusinessEntityName business_entity_name获取 业务主体代码code
func (obj *_LgProviderChannelMgr) WithBusinessEntityName(businessEntityName string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_name"] = businessEntityName })
}

// WithBusinessEntityCode business_entity_code获取 业务主体名称
func (obj *_LgProviderChannelMgr) WithBusinessEntityCode(businessEntityCode string) Option {
	return optionFunc(func(o *options) { o.query["business_entity_code"] = businessEntityCode })
}

// WithSyncLabel sync_label获取
func (obj *_LgProviderChannelMgr) WithSyncLabel(syncLabel []uint8) Option {
	return optionFunc(func(o *options) { o.query["sync_label"] = syncLabel })
}

// WithVariableList variable_list获取 变量列表
func (obj *_LgProviderChannelMgr) WithVariableList(variableList string) Option {
	return optionFunc(func(o *options) { o.query["variable_list"] = variableList })
}

// WithCheckRule check_rule获取 字段校验规则(正则表达式)
func (obj *_LgProviderChannelMgr) WithCheckRule(checkRule string) Option {
	return optionFunc(func(o *options) { o.query["check_rule"] = checkRule })
}

// WithStatus status获取 0:无效，1:有效
func (obj *_LgProviderChannelMgr) WithStatus(status int) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithTakeNoMode take_no_mode获取 取号方式
//
//1:api 2:单号池
func (obj *_LgProviderChannelMgr) WithTakeNoMode(takeNoMode int) Option {
	return optionFunc(func(o *options) { o.query["take_no_mode"] = takeNoMode })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgProviderChannelMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgProviderChannelMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgProviderChannelMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgProviderChannelMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgProviderChannelMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgProviderChannelMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgProviderChannelMgr) GetByOption(opts ...Option) (result LgProviderChannel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", result.OrderNumPoolID).Find(&result.LgOrderNumPool).Error; err != nil { // 单号池主表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgProviderChannelMgr) GetByOptions(opts ...Option) (results []*LgProviderChannel, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_LgProviderChannelMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgProviderChannel, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_LgProviderChannelMgr) GetFromID(id int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`id` = ?", id).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromID 批量查找 主键
func (obj *_LgProviderChannelMgr) GetBatchFromID(ids []int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProviderID 通过provider_id获取内容 系统服务商id
func (obj *_LgProviderChannelMgr) GetFromProviderID(providerID string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`provider_id` = ?", providerID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProviderID 批量查找 系统服务商id
func (obj *_LgProviderChannelMgr) GetBatchFromProviderID(providerIDs []string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`provider_id` IN (?)", providerIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProviderChannelCode 通过provider_channel_code获取内容 渠道代码
func (obj *_LgProviderChannelMgr) GetFromProviderChannelCode(providerChannelCode string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProviderChannelCode 批量查找 渠道代码
func (obj *_LgProviderChannelMgr) GetBatchFromProviderChannelCode(providerChannelCodes []string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`provider_channel_code` IN (?)", providerChannelCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProviderChannelName 通过provider_channel_name获取内容 服务商渠道名称
func (obj *_LgProviderChannelMgr) GetFromProviderChannelName(providerChannelName string) (result LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`provider_channel_name` = ?", providerChannelName).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", result.OrderNumPoolID).Find(&result.LgOrderNumPool).Error; err != nil { // 单号池主表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromProviderChannelName 批量查找 服务商渠道名称
func (obj *_LgProviderChannelMgr) GetBatchFromProviderChannelName(providerChannelNames []string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`provider_channel_name` IN (?)", providerChannelNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrderNumPoolID 通过order_num_pool_id获取内容 单号池id
func (obj *_LgProviderChannelMgr) GetFromOrderNumPoolID(orderNumPoolID int64) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`order_num_pool_id` = ?", orderNumPoolID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrderNumPoolID 批量查找 单号池id
func (obj *_LgProviderChannelMgr) GetBatchFromOrderNumPoolID(orderNumPoolIDs []int64) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`order_num_pool_id` IN (?)", orderNumPoolIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrderNumPoolName 通过order_num_pool_name获取内容 单号池名称
func (obj *_LgProviderChannelMgr) GetFromOrderNumPoolName(orderNumPoolName string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`order_num_pool_name` = ?", orderNumPoolName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrderNumPoolName 批量查找 单号池名称
func (obj *_LgProviderChannelMgr) GetBatchFromOrderNumPoolName(orderNumPoolNames []string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`order_num_pool_name` IN (?)", orderNumPoolNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBusinessEntityName 通过business_entity_name获取内容 业务主体代码code
func (obj *_LgProviderChannelMgr) GetFromBusinessEntityName(businessEntityName string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`business_entity_name` = ?", businessEntityName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBusinessEntityName 批量查找 业务主体代码code
func (obj *_LgProviderChannelMgr) GetBatchFromBusinessEntityName(businessEntityNames []string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`business_entity_name` IN (?)", businessEntityNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBusinessEntityCode 通过business_entity_code获取内容 业务主体名称
func (obj *_LgProviderChannelMgr) GetFromBusinessEntityCode(businessEntityCode string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`business_entity_code` = ?", businessEntityCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBusinessEntityCode 批量查找 业务主体名称
func (obj *_LgProviderChannelMgr) GetBatchFromBusinessEntityCode(businessEntityCodes []string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`business_entity_code` IN (?)", businessEntityCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSyncLabel 通过sync_label获取内容
func (obj *_LgProviderChannelMgr) GetFromSyncLabel(syncLabel []uint8) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`sync_label` = ?", syncLabel).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSyncLabel 批量查找
func (obj *_LgProviderChannelMgr) GetBatchFromSyncLabel(syncLabels [][]uint8) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`sync_label` IN (?)", syncLabels).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVariableList 通过variable_list获取内容 变量列表
func (obj *_LgProviderChannelMgr) GetFromVariableList(variableList string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`variable_list` = ?", variableList).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVariableList 批量查找 变量列表
func (obj *_LgProviderChannelMgr) GetBatchFromVariableList(variableLists []string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`variable_list` IN (?)", variableLists).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCheckRule 通过check_rule获取内容 字段校验规则(正则表达式)
func (obj *_LgProviderChannelMgr) GetFromCheckRule(checkRule string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`check_rule` = ?", checkRule).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCheckRule 批量查找 字段校验规则(正则表达式)
func (obj *_LgProviderChannelMgr) GetBatchFromCheckRule(checkRules []string) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`check_rule` IN (?)", checkRules).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStatus 通过status获取内容 0:无效，1:有效
func (obj *_LgProviderChannelMgr) GetFromStatus(status int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`status` = ?", status).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStatus 批量查找 0:无效，1:有效
func (obj *_LgProviderChannelMgr) GetBatchFromStatus(statuss []int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`status` IN (?)", statuss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTakeNoMode 通过take_no_mode获取内容 取号方式
//
//1:api 2:单号池
func (obj *_LgProviderChannelMgr) GetFromTakeNoMode(takeNoMode int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`take_no_mode` = ?", takeNoMode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTakeNoMode 批量查找 取号方式
//
//1:api 2:单号池
func (obj *_LgProviderChannelMgr) GetBatchFromTakeNoMode(takeNoModes []int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`take_no_mode` IN (?)", takeNoModes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgProviderChannelMgr) GetFromCreateTime(createTime time.Time) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgProviderChannelMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgProviderChannelMgr) GetFromCreateUser(createUser int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgProviderChannelMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgProviderChannelMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgProviderChannelMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgProviderChannelMgr) GetFromUpdateUser(updateUser int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgProviderChannelMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgProviderChannelMgr) GetFromVersion(version int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgProviderChannelMgr) GetBatchFromVersion(versions []int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgProviderChannelMgr) GetFromDeleted(deleted int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgProviderChannelMgr) GetBatchFromDeleted(deleteds []int) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_LgProviderChannelMgr) FetchByPrimaryKey(id int, providerID string, providerChannelCode string) (result LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`id` = ? AND `provider_id` = ? AND `provider_channel_code` = ?", id, providerID, providerChannelCode).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", result.OrderNumPoolID).Find(&result.LgOrderNumPool).Error; err != nil { // 单号池主表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByLgProviderChannelCodeUIndex primary or index 获取唯一内容
func (obj *_LgProviderChannelMgr) FetchUniqueByLgProviderChannelCodeUIndex(providerChannelCode string) (result LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`provider_channel_code` = ?", providerChannelCode).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", result.OrderNumPoolID).Find(&result.LgOrderNumPool).Error; err != nil { // 单号池主表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByLgProviderChannelChannelNameUIndex primary or index 获取唯一内容
func (obj *_LgProviderChannelMgr) FetchUniqueByLgProviderChannelChannelNameUIndex(providerChannelName string) (result LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`provider_channel_name` = ?", providerChannelName).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", result.OrderNumPoolID).Find(&result.LgOrderNumPool).Error; err != nil { // 单号池主表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByLgProviderChannelLgOrderNumPoolIDFk  获取多个内容
func (obj *_LgProviderChannelMgr) FetchIndexByLgProviderChannelLgOrderNumPoolIDFk(orderNumPoolID int64) (results []*LgProviderChannel, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgProviderChannel{}).Where("`order_num_pool_id` = ?", orderNumPoolID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_pool").Where("id = ?", results[i].OrderNumPoolID).Find(&results[i].LgOrderNumPool).Error; err != nil { // 单号池主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
