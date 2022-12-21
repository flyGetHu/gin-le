package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _WmsSorterGridSorterGridRuleMgr struct {
	*_BaseMgr
}

// WmsSorterGridSorterGridRuleMgr open func
func WmsSorterGridSorterGridRuleMgr(db *gorm.DB) *_WmsSorterGridSorterGridRuleMgr {
	if db == nil {
		panic(fmt.Errorf("WmsSorterGridSorterGridRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsSorterGridSorterGridRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_sorter_grid_sorter_grid_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsSorterGridSorterGridRuleMgr) GetTableName() string {
	return "wms_sorter_grid_sorter_grid_rule"
}

// Reset 重置gorm会话
func (obj *_WmsSorterGridSorterGridRuleMgr) Reset() *_WmsSorterGridSorterGridRuleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsSorterGridSorterGridRuleMgr) Get() (result WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", result.GridID).Find(&result.WmsSorterGrid).Error; err != nil { // 分拣机-格口
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WmsSorterGridSorterGridRuleMgr) Gets() (results []*WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsSorterGridSorterGridRuleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsSorterGridSorterGridRuleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGridID grid_id获取 格口id
func (obj *_WmsSorterGridSorterGridRuleMgr) WithGridID(gridID int) Option {
	return optionFunc(func(o *options) { o.query["grid_id"] = gridID })
}

// WithGridRuleID grid_rule_id获取 格口规则id
func (obj *_WmsSorterGridSorterGridRuleMgr) WithGridRuleID(gridRuleID int) Option {
	return optionFunc(func(o *options) { o.query["grid_rule_id"] = gridRuleID })
}

// GetByOption 功能选项模式获取
func (obj *_WmsSorterGridSorterGridRuleMgr) GetByOption(opts ...Option) (result WmsSorterGridSorterGridRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", result.GridID).Find(&result.WmsSorterGrid).Error; err != nil { // 分拣机-格口
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsSorterGridSorterGridRuleMgr) GetByOptions(opts ...Option) (results []*WmsSorterGridSorterGridRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_WmsSorterGridSorterGridRuleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsSorterGridSorterGridRule, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
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
func (obj *_WmsSorterGridSorterGridRuleMgr) GetFromID(id int) (result WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", result.GridID).Find(&result.WmsSorterGrid).Error; err != nil { // 分拣机-格口
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsSorterGridSorterGridRuleMgr) GetBatchFromID(ids []int) (results []*WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromGridID 通过grid_id获取内容 格口id
func (obj *_WmsSorterGridSorterGridRuleMgr) GetFromGridID(gridID int) (results []*WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`grid_id` = ?", gridID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromGridID 批量查找 格口id
func (obj *_WmsSorterGridSorterGridRuleMgr) GetBatchFromGridID(gridIDs []int) (results []*WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`grid_id` IN (?)", gridIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromGridRuleID 通过grid_rule_id获取内容 格口规则id
func (obj *_WmsSorterGridSorterGridRuleMgr) GetFromGridRuleID(gridRuleID int) (results []*WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`grid_rule_id` = ?", gridRuleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromGridRuleID 批量查找 格口规则id
func (obj *_WmsSorterGridSorterGridRuleMgr) GetBatchFromGridRuleID(gridRuleIDs []int) (results []*WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`grid_rule_id` IN (?)", gridRuleIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
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
func (obj *_WmsSorterGridSorterGridRuleMgr) FetchByPrimaryKey(id int) (result WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", result.GridID).Find(&result.WmsSorterGrid).Error; err != nil { // 分拣机-格口
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueIndexByGridIDGridRuleIDIndex primary or index 获取唯一内容
func (obj *_WmsSorterGridSorterGridRuleMgr) FetchUniqueIndexByGridIDGridRuleIDIndex(gridID int, gridRuleID int) (result WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`grid_id` = ? AND `grid_rule_id` = ?", gridID, gridRuleID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", result.GridID).Find(&result.WmsSorterGrid).Error; err != nil { // 分拣机-格口
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByWmsSorterGridSorterGridRuleWmsSorterGridRuleIDFk  获取多个内容
func (obj *_WmsSorterGridSorterGridRuleMgr) FetchIndexByWmsSorterGridSorterGridRuleWmsSorterGridRuleIDFk(gridRuleID int) (results []*WmsSorterGridSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSorterGridRule{}).Where("`grid_rule_id` = ?", gridRuleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid").Where("id = ?", results[i].GridID).Find(&results[i].WmsSorterGrid).Error; err != nil { // 分拣机-格口
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
