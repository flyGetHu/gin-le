package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgDistrictMgr struct {
	*_BaseMgr
}

// LgDistrictMgr open func
func LgDistrictMgr(db *gorm.DB) *_LgDistrictMgr {
	if db == nil {
		panic(fmt.Errorf("LgDistrictMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgDistrictMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_district"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgDistrictMgr) GetTableName() string {
	return "lg_district"
}

// Reset 重置gorm会话
func (obj *_LgDistrictMgr) Reset() *_LgDistrictMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgDistrictMgr) Get() (result LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_country").Where("two_code = ?", result.CountryTwoCode).Find(&result.LgCountryList).Error; err != nil { // 国家表(废弃)
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_LgDistrictMgr) Gets() (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgDistrictMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgDistrictMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCountryNameEn country_name_en获取 国家英文名
func (obj *_LgDistrictMgr) WithCountryNameEn(countryNameEn string) Option {
	return optionFunc(func(o *options) { o.query["country_name_en"] = countryNameEn })
}

// WithCountryNameCn country_name_cn获取 国家中文名
func (obj *_LgDistrictMgr) WithCountryNameCn(countryNameCn string) Option {
	return optionFunc(func(o *options) { o.query["country_name_cn"] = countryNameCn })
}

// WithCountryTwoCode country_two_code获取 国家二字代码
func (obj *_LgDistrictMgr) WithCountryTwoCode(countryTwoCode string) Option {
	return optionFunc(func(o *options) { o.query["country_two_code"] = countryTwoCode })
}

// WithCountryThreeCode country_three_code获取 国家三字代码
func (obj *_LgDistrictMgr) WithCountryThreeCode(countryThreeCode string) Option {
	return optionFunc(func(o *options) { o.query["country_three_code"] = countryThreeCode })
}

// WithProvinceNameEn province_name_en获取 省份英文名
func (obj *_LgDistrictMgr) WithProvinceNameEn(provinceNameEn string) Option {
	return optionFunc(func(o *options) { o.query["province_name_en"] = provinceNameEn })
}

// WithProvinceNameCn province_name_cn获取 省份中文名
func (obj *_LgDistrictMgr) WithProvinceNameCn(provinceNameCn string) Option {
	return optionFunc(func(o *options) { o.query["province_name_cn"] = provinceNameCn })
}

// WithProvinceCode province_code获取 省份二字代码
func (obj *_LgDistrictMgr) WithProvinceCode(provinceCode string) Option {
	return optionFunc(func(o *options) { o.query["province_code"] = provinceCode })
}

// WithCityNameEn city_name_en获取 城市英文名
func (obj *_LgDistrictMgr) WithCityNameEn(cityNameEn string) Option {
	return optionFunc(func(o *options) { o.query["city_name_en"] = cityNameEn })
}

// WithCityNameCn city_name_cn获取 城市中文名
func (obj *_LgDistrictMgr) WithCityNameCn(cityNameCn string) Option {
	return optionFunc(func(o *options) { o.query["city_name_cn"] = cityNameCn })
}

// WithCityCode city_code获取 城市二字代码
func (obj *_LgDistrictMgr) WithCityCode(cityCode string) Option {
	return optionFunc(func(o *options) { o.query["city_code"] = cityCode })
}

// WithZipcode zipcode获取 邮编
func (obj *_LgDistrictMgr) WithZipcode(zipcode string) Option {
	return optionFunc(func(o *options) { o.query["zipcode"] = zipcode })
}

// WithAddress address获取 地址
func (obj *_LgDistrictMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithIsVague is_vague获取 是否模糊匹配
func (obj *_LgDistrictMgr) WithIsVague(isVague int) Option {
	return optionFunc(func(o *options) { o.query["is_vague"] = isVague })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgDistrictMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgDistrictMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgDistrictMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgDistrictMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgDistrictMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgDistrictMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgDistrictMgr) GetByOption(opts ...Option) (result LgDistrict, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_country").Where("two_code = ?", result.CountryTwoCode).Find(&result.LgCountryList).Error; err != nil { // 国家表(废弃)
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgDistrictMgr) GetByOptions(opts ...Option) (results []*LgDistrict, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_LgDistrictMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgDistrict, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
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

// GetFromID 通过id获取内容 ID
func (obj *_LgDistrictMgr) GetFromID(id int) (result LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_country").Where("two_code = ?", result.CountryTwoCode).Find(&result.LgCountryList).Error; err != nil { // 国家表(废弃)
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgDistrictMgr) GetBatchFromID(ids []int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCountryNameEn 通过country_name_en获取内容 国家英文名
func (obj *_LgDistrictMgr) GetFromCountryNameEn(countryNameEn string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_name_en` = ?", countryNameEn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCountryNameEn 批量查找 国家英文名
func (obj *_LgDistrictMgr) GetBatchFromCountryNameEn(countryNameEns []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_name_en` IN (?)", countryNameEns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCountryNameCn 通过country_name_cn获取内容 国家中文名
func (obj *_LgDistrictMgr) GetFromCountryNameCn(countryNameCn string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_name_cn` = ?", countryNameCn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCountryNameCn 批量查找 国家中文名
func (obj *_LgDistrictMgr) GetBatchFromCountryNameCn(countryNameCns []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_name_cn` IN (?)", countryNameCns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCountryTwoCode 通过country_two_code获取内容 国家二字代码
func (obj *_LgDistrictMgr) GetFromCountryTwoCode(countryTwoCode string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_two_code` = ?", countryTwoCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCountryTwoCode 批量查找 国家二字代码
func (obj *_LgDistrictMgr) GetBatchFromCountryTwoCode(countryTwoCodes []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_two_code` IN (?)", countryTwoCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCountryThreeCode 通过country_three_code获取内容 国家三字代码
func (obj *_LgDistrictMgr) GetFromCountryThreeCode(countryThreeCode string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_three_code` = ?", countryThreeCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCountryThreeCode 批量查找 国家三字代码
func (obj *_LgDistrictMgr) GetBatchFromCountryThreeCode(countryThreeCodes []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_three_code` IN (?)", countryThreeCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProvinceNameEn 通过province_name_en获取内容 省份英文名
func (obj *_LgDistrictMgr) GetFromProvinceNameEn(provinceNameEn string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`province_name_en` = ?", provinceNameEn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProvinceNameEn 批量查找 省份英文名
func (obj *_LgDistrictMgr) GetBatchFromProvinceNameEn(provinceNameEns []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`province_name_en` IN (?)", provinceNameEns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProvinceNameCn 通过province_name_cn获取内容 省份中文名
func (obj *_LgDistrictMgr) GetFromProvinceNameCn(provinceNameCn string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`province_name_cn` = ?", provinceNameCn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProvinceNameCn 批量查找 省份中文名
func (obj *_LgDistrictMgr) GetBatchFromProvinceNameCn(provinceNameCns []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`province_name_cn` IN (?)", provinceNameCns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromProvinceCode 通过province_code获取内容 省份二字代码
func (obj *_LgDistrictMgr) GetFromProvinceCode(provinceCode string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`province_code` = ?", provinceCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromProvinceCode 批量查找 省份二字代码
func (obj *_LgDistrictMgr) GetBatchFromProvinceCode(provinceCodes []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`province_code` IN (?)", provinceCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCityNameEn 通过city_name_en获取内容 城市英文名
func (obj *_LgDistrictMgr) GetFromCityNameEn(cityNameEn string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`city_name_en` = ?", cityNameEn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCityNameEn 批量查找 城市英文名
func (obj *_LgDistrictMgr) GetBatchFromCityNameEn(cityNameEns []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`city_name_en` IN (?)", cityNameEns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCityNameCn 通过city_name_cn获取内容 城市中文名
func (obj *_LgDistrictMgr) GetFromCityNameCn(cityNameCn string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`city_name_cn` = ?", cityNameCn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCityNameCn 批量查找 城市中文名
func (obj *_LgDistrictMgr) GetBatchFromCityNameCn(cityNameCns []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`city_name_cn` IN (?)", cityNameCns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCityCode 通过city_code获取内容 城市二字代码
func (obj *_LgDistrictMgr) GetFromCityCode(cityCode string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`city_code` = ?", cityCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCityCode 批量查找 城市二字代码
func (obj *_LgDistrictMgr) GetBatchFromCityCode(cityCodes []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`city_code` IN (?)", cityCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromZipcode 通过zipcode获取内容 邮编
func (obj *_LgDistrictMgr) GetFromZipcode(zipcode string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`zipcode` = ?", zipcode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromZipcode 批量查找 邮编
func (obj *_LgDistrictMgr) GetBatchFromZipcode(zipcodes []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`zipcode` IN (?)", zipcodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAddress 通过address获取内容 地址
func (obj *_LgDistrictMgr) GetFromAddress(address string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`address` = ?", address).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAddress 批量查找 地址
func (obj *_LgDistrictMgr) GetBatchFromAddress(addresss []string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`address` IN (?)", addresss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsVague 通过is_vague获取内容 是否模糊匹配
func (obj *_LgDistrictMgr) GetFromIsVague(isVague int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`is_vague` = ?", isVague).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsVague 批量查找 是否模糊匹配
func (obj *_LgDistrictMgr) GetBatchFromIsVague(isVagues []int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`is_vague` IN (?)", isVagues).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgDistrictMgr) GetFromCreateTime(createTime time.Time) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgDistrictMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgDistrictMgr) GetFromCreateUser(createUser int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgDistrictMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgDistrictMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgDistrictMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgDistrictMgr) GetFromUpdateUser(updateUser int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgDistrictMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgDistrictMgr) GetFromVersion(version int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgDistrictMgr) GetBatchFromVersion(versions []int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgDistrictMgr) GetFromDeleted(deleted int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgDistrictMgr) GetBatchFromDeleted(deleteds []int) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
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
func (obj *_LgDistrictMgr) FetchByPrimaryKey(id int) (result LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_country").Where("two_code = ?", result.CountryTwoCode).Find(&result.LgCountryList).Error; err != nil { // 国家表(废弃)
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByLgDistrictLgCountryTwoCodeFk  获取多个内容
func (obj *_LgDistrictMgr) FetchIndexByLgDistrictLgCountryTwoCodeFk(countryTwoCode string) (results []*LgDistrict, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgDistrict{}).Where("`country_two_code` = ?", countryTwoCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_country").Where("two_code = ?", results[i].CountryTwoCode).Find(&results[i].LgCountryList).Error; err != nil { // 国家表(废弃)
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
