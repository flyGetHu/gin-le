package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgAreaMgr struct {
	*_BaseMgr
}

// LgAreaMgr open func
func LgAreaMgr(db *gorm.DB) *_LgAreaMgr {
	if db == nil {
		panic(fmt.Errorf("LgAreaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgAreaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_area"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgAreaMgr) GetTableName() string {
	return "lg_area"
}

// Reset 重置gorm会话
func (obj *_LgAreaMgr) Reset() *_LgAreaMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgAreaMgr) Get() (result LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_area").Where("id = ?", result.ParentID).Find(&result.LgArea).Error; err != nil { // 地区总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_LgAreaMgr) Gets() (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgAreaMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgArea{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_LgAreaMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 父ID
func (obj *_LgAreaMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithRegionalID regional_id获取 大区id
func (obj *_LgAreaMgr) WithRegionalID(regionalID string) Option {
	return optionFunc(func(o *options) { o.query["regional_id"] = regionalID })
}

// WithRegionalNameCn regional_name_cn获取 大区中文名称
func (obj *_LgAreaMgr) WithRegionalNameCn(regionalNameCn string) Option {
	return optionFunc(func(o *options) { o.query["regional_name_cn"] = regionalNameCn })
}

// WithLevel level获取 等级，0: 国家，1：省份，2：城市
func (obj *_LgAreaMgr) WithLevel(level int) Option {
	return optionFunc(func(o *options) { o.query["level"] = level })
}

// WithTwoCode two_code获取 二字码
func (obj *_LgAreaMgr) WithTwoCode(twoCode string) Option {
	return optionFunc(func(o *options) { o.query["two_code"] = twoCode })
}

// WithThreeCode three_code获取 三字码
func (obj *_LgAreaMgr) WithThreeCode(threeCode string) Option {
	return optionFunc(func(o *options) { o.query["three_code"] = threeCode })
}

// WithNameEn name_en获取 英文名
func (obj *_LgAreaMgr) WithNameEn(nameEn string) Option {
	return optionFunc(func(o *options) { o.query["name_en"] = nameEn })
}

// WithNameCn name_cn获取 中文名
func (obj *_LgAreaMgr) WithNameCn(nameCn string) Option {
	return optionFunc(func(o *options) { o.query["name_cn"] = nameCn })
}

// WithZipCode zip_code获取 邮政编码 level等级为2时 邮编不能为空
func (obj *_LgAreaMgr) WithZipCode(zipCode string) Option {
	return optionFunc(func(o *options) { o.query["zip_code"] = zipCode })
}

// WithExtra extra获取 扩展字段(匹配正确省市)
func (obj *_LgAreaMgr) WithExtra(extra string) Option {
	return optionFunc(func(o *options) { o.query["extra"] = extra })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgAreaMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgAreaMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgAreaMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgAreaMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgAreaMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgAreaMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgAreaMgr) GetByOption(opts ...Option) (result LgArea, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_area").Where("id = ?", result.ParentID).Find(&result.LgArea).Error; err != nil { // 地区总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgAreaMgr) GetByOptions(opts ...Option) (results []*LgArea, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_LgAreaMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgArea, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
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
func (obj *_LgAreaMgr) GetFromID(id int) (result LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_area").Where("id = ?", result.ParentID).Find(&result.LgArea).Error; err != nil { // 地区总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_LgAreaMgr) GetBatchFromID(ids []int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromParentID 通过parent_id获取内容 父ID
func (obj *_LgAreaMgr) GetFromParentID(parentID int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`parent_id` = ?", parentID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromParentID 批量查找 父ID
func (obj *_LgAreaMgr) GetBatchFromParentID(parentIDs []int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRegionalID 通过regional_id获取内容 大区id
func (obj *_LgAreaMgr) GetFromRegionalID(regionalID string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`regional_id` = ?", regionalID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRegionalID 批量查找 大区id
func (obj *_LgAreaMgr) GetBatchFromRegionalID(regionalIDs []string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`regional_id` IN (?)", regionalIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRegionalNameCn 通过regional_name_cn获取内容 大区中文名称
func (obj *_LgAreaMgr) GetFromRegionalNameCn(regionalNameCn string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`regional_name_cn` = ?", regionalNameCn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRegionalNameCn 批量查找 大区中文名称
func (obj *_LgAreaMgr) GetBatchFromRegionalNameCn(regionalNameCns []string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`regional_name_cn` IN (?)", regionalNameCns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromLevel 通过level获取内容 等级，0: 国家，1：省份，2：城市
func (obj *_LgAreaMgr) GetFromLevel(level int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`level` = ?", level).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromLevel 批量查找 等级，0: 国家，1：省份，2：城市
func (obj *_LgAreaMgr) GetBatchFromLevel(levels []int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`level` IN (?)", levels).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTwoCode 通过two_code获取内容 二字码
func (obj *_LgAreaMgr) GetFromTwoCode(twoCode string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`two_code` = ?", twoCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTwoCode 批量查找 二字码
func (obj *_LgAreaMgr) GetBatchFromTwoCode(twoCodes []string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`two_code` IN (?)", twoCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromThreeCode 通过three_code获取内容 三字码
func (obj *_LgAreaMgr) GetFromThreeCode(threeCode string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`three_code` = ?", threeCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromThreeCode 批量查找 三字码
func (obj *_LgAreaMgr) GetBatchFromThreeCode(threeCodes []string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`three_code` IN (?)", threeCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNameEn 通过name_en获取内容 英文名
func (obj *_LgAreaMgr) GetFromNameEn(nameEn string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`name_en` = ?", nameEn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNameEn 批量查找 英文名
func (obj *_LgAreaMgr) GetBatchFromNameEn(nameEns []string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`name_en` IN (?)", nameEns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNameCn 通过name_cn获取内容 中文名
func (obj *_LgAreaMgr) GetFromNameCn(nameCn string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`name_cn` = ?", nameCn).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNameCn 批量查找 中文名
func (obj *_LgAreaMgr) GetBatchFromNameCn(nameCns []string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`name_cn` IN (?)", nameCns).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromZipCode 通过zip_code获取内容 邮政编码 level等级为2时 邮编不能为空
func (obj *_LgAreaMgr) GetFromZipCode(zipCode string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`zip_code` = ?", zipCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromZipCode 批量查找 邮政编码 level等级为2时 邮编不能为空
func (obj *_LgAreaMgr) GetBatchFromZipCode(zipCodes []string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`zip_code` IN (?)", zipCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromExtra 通过extra获取内容 扩展字段(匹配正确省市)
func (obj *_LgAreaMgr) GetFromExtra(extra string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`extra` = ?", extra).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromExtra 批量查找 扩展字段(匹配正确省市)
func (obj *_LgAreaMgr) GetBatchFromExtra(extras []string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`extra` IN (?)", extras).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgAreaMgr) GetFromCreateTime(createTime time.Time) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgAreaMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgAreaMgr) GetFromCreateUser(createUser int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgAreaMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgAreaMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgAreaMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgAreaMgr) GetFromUpdateUser(updateUser int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgAreaMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgAreaMgr) GetFromVersion(version int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgAreaMgr) GetBatchFromVersion(versions []int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgAreaMgr) GetFromDeleted(deleted int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgAreaMgr) GetBatchFromDeleted(deleteds []int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
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
func (obj *_LgAreaMgr) FetchByPrimaryKey(id int) (result LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_area").Where("id = ?", result.ParentID).Find(&result.LgArea).Error; err != nil { // 地区总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByLgAreaParentIDFk  获取多个内容
func (obj *_LgAreaMgr) FetchIndexByLgAreaParentIDFk(parentID int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`parent_id` = ?", parentID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByLgRegionalIDIndex  获取多个内容
func (obj *_LgAreaMgr) FetchIndexByLgRegionalIDIndex(regionalID string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`regional_id` = ?", regionalID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByLgAreaLevelIndex  获取多个内容
func (obj *_LgAreaMgr) FetchIndexByLgAreaLevelIndex(level int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`level` = ?", level).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByLgAreaTwoCodeLevelIndex  获取多个内容
func (obj *_LgAreaMgr) FetchIndexByLgAreaTwoCodeLevelIndex(level int, twoCode string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`level` = ? AND `two_code` = ?", level, twoCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByIndexTwoCode  获取多个内容
func (obj *_LgAreaMgr) FetchIndexByIndexTwoCode(twoCode string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`two_code` = ?", twoCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByIndexThreeCode  获取多个内容
func (obj *_LgAreaMgr) FetchIndexByIndexThreeCode(threeCode string) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`three_code` = ?", threeCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByLgAreaDeletedIndex  获取多个内容
func (obj *_LgAreaMgr) FetchIndexByLgAreaDeletedIndex(deleted int) (results []*LgArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgArea{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_area").Where("id = ?", results[i].ParentID).Find(&results[i].LgArea).Error; err != nil { // 地区总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
