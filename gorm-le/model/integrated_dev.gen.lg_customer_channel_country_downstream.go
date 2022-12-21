package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgCustomerChannelCountryDownstreamMgr struct {
	*_BaseMgr
}

// LgCustomerChannelCountryDownstreamMgr open func
func LgCustomerChannelCountryDownstreamMgr(db *gorm.DB) *_LgCustomerChannelCountryDownstreamMgr {
	if db == nil {
		panic(fmt.Errorf("LgCustomerChannelCountryDownstreamMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgCustomerChannelCountryDownstreamMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_customer_channel_country_downstream"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetTableName() string {
	return "lg_customer_channel_country_downstream"
}

// Reset 重置gorm会话
func (obj *_LgCustomerChannelCountryDownstreamMgr) Reset() *_LgCustomerChannelCountryDownstreamMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgCustomerChannelCountryDownstreamMgr) Get() (result LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_downstream").Where("id = ?", result.DownstreamID).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", result.DownstreamName).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_LgCustomerChannelCountryDownstreamMgr) Gets() (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgCustomerChannelCountryDownstreamMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCustomerChannelID customer_channel_id获取 客户渠道ID
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithCustomerChannelID(customerChannelID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_id"] = customerChannelID })
}

// WithCountryTwoCode country_two_code获取 国家二字码
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithCountryTwoCode(countryTwoCode string) Option {
	return optionFunc(func(o *options) { o.query["country_two_code"] = countryTwoCode })
}

// WithCustomerChannelCountryID customer_channel_country_id获取 客户渠道支持国家id
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithCustomerChannelCountryID(customerChannelCountryID int) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_country_id"] = customerChannelCountryID })
}

// WithDownstreamID downstream_id获取 下家id
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithDownstreamID(downstreamID int) Option {
	return optionFunc(func(o *options) { o.query["downstream_id"] = downstreamID })
}

// WithDownstreamName downstream_name获取 下家名称
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithDownstreamName(downstreamName string) Option {
	return optionFunc(func(o *options) { o.query["downstream_name"] = downstreamName })
}

// WithDownstreamCode downstream_code获取 下家code
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithDownstreamCode(downstreamCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_code"] = downstreamCode })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgCustomerChannelCountryDownstreamMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// GetByOption 功能选项模式获取
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetByOption(opts ...Option) (result LgCustomerChannelCountryDownstream, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_downstream").Where("id = ?", result.DownstreamID).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", result.DownstreamName).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetByOptions(opts ...Option) (results []*LgCustomerChannelCountryDownstream, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_LgCustomerChannelCountryDownstreamMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgCustomerChannelCountryDownstream, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
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
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromID(id int) (result LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_downstream").Where("id = ?", result.DownstreamID).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", result.DownstreamName).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromID(ids []int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCustomerChannelID 通过customer_channel_id获取内容 客户渠道ID
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromCustomerChannelID(customerChannelID int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`customer_channel_id` = ?", customerChannelID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCustomerChannelID 批量查找 客户渠道ID
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromCustomerChannelID(customerChannelIDs []int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`customer_channel_id` IN (?)", customerChannelIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCountryTwoCode 通过country_two_code获取内容 国家二字码
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromCountryTwoCode(countryTwoCode string) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`country_two_code` = ?", countryTwoCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCountryTwoCode 批量查找 国家二字码
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromCountryTwoCode(countryTwoCodes []string) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`country_two_code` IN (?)", countryTwoCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCustomerChannelCountryID 通过customer_channel_country_id获取内容 客户渠道支持国家id
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromCustomerChannelCountryID(customerChannelCountryID int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`customer_channel_country_id` = ?", customerChannelCountryID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCustomerChannelCountryID 批量查找 客户渠道支持国家id
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromCustomerChannelCountryID(customerChannelCountryIDs []int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`customer_channel_country_id` IN (?)", customerChannelCountryIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDownstreamID 通过downstream_id获取内容 下家id
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromDownstreamID(downstreamID int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`downstream_id` = ?", downstreamID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDownstreamID 批量查找 下家id
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromDownstreamID(downstreamIDs []int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`downstream_id` IN (?)", downstreamIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDownstreamName 通过downstream_name获取内容 下家名称
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromDownstreamName(downstreamName string) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`downstream_name` = ?", downstreamName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDownstreamName 批量查找 下家名称
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromDownstreamName(downstreamNames []string) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`downstream_name` IN (?)", downstreamNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDownstreamCode 通过downstream_code获取内容 下家code
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromDownstreamCode(downstreamCode string) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`downstream_code` = ?", downstreamCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDownstreamCode 批量查找 下家code
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromDownstreamCode(downstreamCodes []string) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`downstream_code` IN (?)", downstreamCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromCreateTime(createTime time.Time) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromCreateUser(createUser int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromUpdateUser(updateUser int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetFromVersion(version int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgCustomerChannelCountryDownstreamMgr) GetBatchFromVersion(versions []int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
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
func (obj *_LgCustomerChannelCountryDownstreamMgr) FetchByPrimaryKey(id int) (result LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_downstream").Where("id = ?", result.DownstreamID).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", result.DownstreamName).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByIndexCustomerChannelCountryIDDownstreamID primary or index 获取唯一内容
func (obj *_LgCustomerChannelCountryDownstreamMgr) FetchUniqueIndexByIndexCustomerChannelCountryIDDownstreamID(customerChannelCountryID int, downstreamID int) (result LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`customer_channel_country_id` = ? AND `downstream_id` = ?", customerChannelCountryID, downstreamID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_downstream").Where("id = ?", result.DownstreamID).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", result.DownstreamName).Find(&result.LgDownstream).Error; err != nil { // 下家
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByDownstreamIDFk  获取多个内容
func (obj *_LgCustomerChannelCountryDownstreamMgr) FetchIndexByDownstreamIDFk(downstreamID int) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`downstream_id` = ?", downstreamID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByDownstreamNameFk  获取多个内容
func (obj *_LgCustomerChannelCountryDownstreamMgr) FetchIndexByDownstreamNameFk(downstreamName string) (results []*LgCustomerChannelCountryDownstream, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgCustomerChannelCountryDownstream{}).Where("`downstream_name` = ?", downstreamName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_downstream").Where("id = ?", results[i].DownstreamID).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("lg_downstream").Where("downstream_name = ?", results[i].DownstreamName).Find(&results[i].LgDownstream).Error; err != nil { // 下家
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
