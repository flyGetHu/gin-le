package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgCustomerChannelCountryMgr struct {
	*_BaseMgr
}

// LgCustomerChannelCountryMgr open func
func LgCustomerChannelCountryMgr(db *gorm.DB) *_LgCustomerChannelCountryMgr {
	if db == nil {
		panic(fmt.Errorf("LgCustomerChannelCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgCustomerChannelCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_customer_channel_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgCustomerChannelCountryMgr) GetTableName() string {
	return "lg_customer_channel_country"
}

// Reset 重置gorm会话
func (obj *_LgCustomerChannelCountryMgr) Reset() *_LgCustomerChannelCountryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgCustomerChannelCountryMgr) Get() (result LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_label_template").Where("code = ?", result.LabelCode).Find(&result.LgLabelTemplate).Error; err != nil { // 物流面单模板维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_LgCustomerChannelCountryMgr) Gets() (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgCustomerChannelCountryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgCustomerChannelCountryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgCustomerChannelCountryMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCountryTwoCode country_two_code获取 国家二字代码
func (obj *_LgCustomerChannelCountryMgr) WithCountryTwoCode(countryTwoCode string) Option {
	return optionFunc(func(o *options) { o.query["country_two_code"] = countryTwoCode })
}

// WithProviderChannelCode provider_channel_code获取 服务商渠道Code
func (obj *_LgCustomerChannelCountryMgr) WithProviderChannelCode(providerChannelCode string) Option {
	return optionFunc(func(o *options) { o.query["provider_channel_code"] = providerChannelCode })
}

// WithLabelSupport label_support获取 面单支持
func (obj *_LgCustomerChannelCountryMgr) WithLabelSupport(labelSupport int) Option {
	return optionFunc(func(o *options) { o.query["label_support"] = labelSupport })
}

// WithLabelCode label_code获取 制作面单code
func (obj *_LgCustomerChannelCountryMgr) WithLabelCode(labelCode string) Option {
	return optionFunc(func(o *options) { o.query["label_code"] = labelCode })
}

// WithMaxWeight max_weight获取 最大出库重量
func (obj *_LgCustomerChannelCountryMgr) WithMaxWeight(maxWeight float64) Option {
	return optionFunc(func(o *options) { o.query["max_weight"] = maxWeight })
}

// WithMinWeight min_weight获取 最小出库重量
func (obj *_LgCustomerChannelCountryMgr) WithMinWeight(minWeight float64) Option {
	return optionFunc(func(o *options) { o.query["min_weight"] = minWeight })
}

// WithChannelCountryAddress channel_country_address获取 出/入库地址
func (obj *_LgCustomerChannelCountryMgr) WithChannelCountryAddress(channelCountryAddress string) Option {
	return optionFunc(func(o *options) { o.query["channel_country_address"] = channelCountryAddress })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgCustomerChannelCountryMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgCustomerChannelCountryMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgCustomerChannelCountryMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgCustomerChannelCountryMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgCustomerChannelCountryMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithShipmentCode shipment_code获取 轨迹标识CODE(APICODE)
func (obj *_LgCustomerChannelCountryMgr) WithShipmentCode(shipmentCode string) Option {
	return optionFunc(func(o *options) { o.query["shipment_code"] = shipmentCode })
}

// WithCarrierCode carrier_code获取 运输商代码(例如china-post)
func (obj *_LgCustomerChannelCountryMgr) WithCarrierCode(carrierCode string) Option {
	return optionFunc(func(o *options) { o.query["carrier_code"] = carrierCode })
}

// GetByOption 功能选项模式获取
func (obj *_LgCustomerChannelCountryMgr) GetByOption(opts ...Option) (result LgCustomerChannelCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_label_template").Where("code = ?", result.LabelCode).Find(&result.LgLabelTemplate).Error; err != nil { // 物流面单模板维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgCustomerChannelCountryMgr) GetByOptions(opts ...Option) (results []*LgCustomerChannelCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_LgCustomerChannelCountryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgCustomerChannelCountry, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
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

// GetFromID 通过id获取内容
func (obj *_LgCustomerChannelCountryMgr) GetFromID(id int) (result LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_label_template").Where("code = ?", result.LabelCode).Find(&result.LgLabelTemplate).Error; err != nil { // 物流面单模板维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromID(ids []int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgCustomerChannelCountryMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCountryTwoCode 通过country_two_code获取内容 国家二字代码
func (obj *_LgCustomerChannelCountryMgr) GetFromCountryTwoCode(countryTwoCode string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`country_two_code` = ?", countryTwoCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCountryTwoCode 批量查找 国家二字代码
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromCountryTwoCode(countryTwoCodes []string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`country_two_code` IN (?)", countryTwoCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProviderChannelCode 通过provider_channel_code获取内容 服务商渠道Code
func (obj *_LgCustomerChannelCountryMgr) GetFromProviderChannelCode(providerChannelCode string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`provider_channel_code` = ?", providerChannelCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProviderChannelCode 批量查找 服务商渠道Code
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromProviderChannelCode(providerChannelCodes []string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`provider_channel_code` IN (?)", providerChannelCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromLabelSupport 通过label_support获取内容 面单支持
func (obj *_LgCustomerChannelCountryMgr) GetFromLabelSupport(labelSupport int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`label_support` = ?", labelSupport).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromLabelSupport 批量查找 面单支持
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromLabelSupport(labelSupports []int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`label_support` IN (?)", labelSupports).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromLabelCode 通过label_code获取内容 制作面单code
func (obj *_LgCustomerChannelCountryMgr) GetFromLabelCode(labelCode string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`label_code` = ?", labelCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromLabelCode 批量查找 制作面单code
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromLabelCode(labelCodes []string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`label_code` IN (?)", labelCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMaxWeight 通过max_weight获取内容 最大出库重量
func (obj *_LgCustomerChannelCountryMgr) GetFromMaxWeight(maxWeight float64) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`max_weight` = ?", maxWeight).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMaxWeight 批量查找 最大出库重量
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromMaxWeight(maxWeights []float64) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`max_weight` IN (?)", maxWeights).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromMinWeight 通过min_weight获取内容 最小出库重量
func (obj *_LgCustomerChannelCountryMgr) GetFromMinWeight(minWeight float64) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`min_weight` = ?", minWeight).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromMinWeight 批量查找 最小出库重量
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromMinWeight(minWeights []float64) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`min_weight` IN (?)", minWeights).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromChannelCountryAddress 通过channel_country_address获取内容 出/入库地址
func (obj *_LgCustomerChannelCountryMgr) GetFromChannelCountryAddress(channelCountryAddress string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`channel_country_address` = ?", channelCountryAddress).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromChannelCountryAddress 批量查找 出/入库地址
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromChannelCountryAddress(channelCountryAddresss []string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`channel_country_address` IN (?)", channelCountryAddresss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgCustomerChannelCountryMgr) GetFromCreateTime(createTime time.Time) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgCustomerChannelCountryMgr) GetFromCreateUser(createUser int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgCustomerChannelCountryMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgCustomerChannelCountryMgr) GetFromUpdateUser(updateUser int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgCustomerChannelCountryMgr) GetFromVersion(version int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromVersion(versions []int) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromShipmentCode 通过shipment_code获取内容 轨迹标识CODE(APICODE)
func (obj *_LgCustomerChannelCountryMgr) GetFromShipmentCode(shipmentCode string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`shipment_code` = ?", shipmentCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromShipmentCode 批量查找 轨迹标识CODE(APICODE)
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromShipmentCode(shipmentCodes []string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`shipment_code` IN (?)", shipmentCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCarrierCode 通过carrier_code获取内容 运输商代码(例如china-post)
func (obj *_LgCustomerChannelCountryMgr) GetFromCarrierCode(carrierCode string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`carrier_code` = ?", carrierCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCarrierCode 批量查找 运输商代码(例如china-post)
func (obj *_LgCustomerChannelCountryMgr) GetBatchFromCarrierCode(carrierCodes []string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`carrier_code` IN (?)", carrierCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
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
func (obj *_LgCustomerChannelCountryMgr) FetchByPrimaryKey(id int) (result LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_label_template").Where("code = ?", result.LabelCode).Find(&result.LgLabelTemplate).Error; err != nil { // 物流面单模板维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByLgCustomerChannelCountryLgLabelTemplateCodeFk  获取多个内容
func (obj *_LgCustomerChannelCountryMgr) FetchIndexByLgCustomerChannelCountryLgLabelTemplateCodeFk(labelCode string) (results []*LgCustomerChannelCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountry{}).Where("`label_code` = ?", labelCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_label_template").Where("code = ?", results[i].LabelCode).Find(&results[i].LgLabelTemplate).Error; err != nil { // 物流面单模板维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
