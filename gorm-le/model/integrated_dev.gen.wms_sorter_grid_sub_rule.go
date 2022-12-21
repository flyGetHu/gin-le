package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsSorterGridSubRuleMgr struct {
	*_BaseMgr
}

// WmsSorterGridSubRuleMgr open func
func WmsSorterGridSubRuleMgr(db *gorm.DB) *_WmsSorterGridSubRuleMgr {
	if db == nil {
		panic(fmt.Errorf("WmsSorterGridSubRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsSorterGridSubRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_sorter_grid_sub_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsSorterGridSubRuleMgr) GetTableName() string {
	return "wms_sorter_grid_sub_rule"
}

// Reset 重置gorm会话
func (obj *_WmsSorterGridSubRuleMgr) Reset() *_WmsSorterGridSubRuleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsSorterGridSubRuleMgr) Get() (result WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WmsSorterGridSubRuleMgr) Gets() (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_WmsSorterGridSubRuleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsSorterGridSubRuleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGridRuleID grid_rule_id获取 格口规则id
func (obj *_WmsSorterGridSubRuleMgr) WithGridRuleID(gridRuleID int) Option {
	return optionFunc(func(o *options) { o.query["grid_rule_id"] = gridRuleID })
}

// WithGridNums grid_nums获取 格口，多个以逗号隔开
func (obj *_WmsSorterGridSubRuleMgr) WithGridNums(gridNums string) Option {
	return optionFunc(func(o *options) { o.query["grid_nums"] = gridNums })
}

// WithDownstreamCode downstream_code获取 下家code
func (obj *_WmsSorterGridSubRuleMgr) WithDownstreamCode(downstreamCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_code"] = downstreamCode })
}

// WithRuleType rule_type获取 规则类型，weight:重量，country:国家，customer:客户，no_sub_rule:不存在子规则，not_found_order:未找到订单，problem_order:问题订单
func (obj *_WmsSorterGridSubRuleMgr) WithRuleType(ruleType string) Option {
	return optionFunc(func(o *options) { o.query["rule_type"] = ruleType })
}

// WithRuleValue rule_value获取
func (obj *_WmsSorterGridSubRuleMgr) WithRuleValue(ruleValue string) Option {
	return optionFunc(func(o *options) { o.query["rule_value"] = ruleValue })
}

// WithOverFlag over_flag获取 是否终结
func (obj *_WmsSorterGridSubRuleMgr) WithOverFlag(overFlag bool) Option {
	return optionFunc(func(o *options) { o.query["over_flag"] = overFlag })
}

// WithRemark remark获取 备注
func (obj *_WmsSorterGridSubRuleMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsSorterGridSubRuleMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsSorterGridSubRuleMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsSorterGridSubRuleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_WmsSorterGridSubRuleMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsSorterGridSubRuleMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsSorterGridSubRuleMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// WithUseState use_state获取 0使用中 1未使用
func (obj *_WmsSorterGridSubRuleMgr) WithUseState(useState int) Option {
	return optionFunc(func(o *options) { o.query["use_state"] = useState })
}

// GetByOption 功能选项模式获取
func (obj *_WmsSorterGridSubRuleMgr) GetByOption(opts ...Option) (result WmsSorterGridSubRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsSorterGridSubRuleMgr) GetByOptions(opts ...Option) (results []*WmsSorterGridSubRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_WmsSorterGridSubRuleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsSorterGridSubRule, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_WmsSorterGridSubRuleMgr) GetFromID(id int) (result WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromID(ids []int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_WmsSorterGridSubRuleMgr) GetFromGridRuleID(gridRuleID int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`grid_rule_id` = ?", gridRuleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromGridRuleID(gridRuleIDs []int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`grid_rule_id` IN (?)", gridRuleIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromGridNums 通过grid_nums获取内容 格口，多个以逗号隔开
func (obj *_WmsSorterGridSubRuleMgr) GetFromGridNums(gridNums string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`grid_nums` = ?", gridNums).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromGridNums 批量查找 格口，多个以逗号隔开
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromGridNums(gridNumss []string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`grid_nums` IN (?)", gridNumss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDownstreamCode 通过downstream_code获取内容 下家code
func (obj *_WmsSorterGridSubRuleMgr) GetFromDownstreamCode(downstreamCode string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`downstream_code` = ?", downstreamCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDownstreamCode 批量查找 下家code
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromDownstreamCode(downstreamCodes []string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`downstream_code` IN (?)", downstreamCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRuleType 通过rule_type获取内容 规则类型，weight:重量，country:国家，customer:客户，no_sub_rule:不存在子规则，not_found_order:未找到订单，problem_order:问题订单
func (obj *_WmsSorterGridSubRuleMgr) GetFromRuleType(ruleType string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`rule_type` = ?", ruleType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRuleType 批量查找 规则类型，weight:重量，country:国家，customer:客户，no_sub_rule:不存在子规则，not_found_order:未找到订单，problem_order:问题订单
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromRuleType(ruleTypes []string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`rule_type` IN (?)", ruleTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRuleValue 通过rule_value获取内容
func (obj *_WmsSorterGridSubRuleMgr) GetFromRuleValue(ruleValue string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`rule_value` = ?", ruleValue).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRuleValue 批量查找
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromRuleValue(ruleValues []string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`rule_value` IN (?)", ruleValues).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOverFlag 通过over_flag获取内容 是否终结
func (obj *_WmsSorterGridSubRuleMgr) GetFromOverFlag(overFlag bool) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`over_flag` = ?", overFlag).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOverFlag 批量查找 是否终结
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromOverFlag(overFlags []bool) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`over_flag` IN (?)", overFlags).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WmsSorterGridSubRuleMgr) GetFromRemark(remark string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`remark` = ?", remark).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromRemark(remarks []string) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`remark` IN (?)", remarks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsSorterGridSubRuleMgr) GetFromCreateTime(createTime time.Time) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsSorterGridSubRuleMgr) GetFromCreateUser(createUser int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsSorterGridSubRuleMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_WmsSorterGridSubRuleMgr) GetFromUpdateUser(updateUser int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsSorterGridSubRuleMgr) GetFromVersion(version int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromVersion(versions []int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsSorterGridSubRuleMgr) GetFromDeleted(deleted int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUseState 通过use_state获取内容 0使用中 1未使用
func (obj *_WmsSorterGridSubRuleMgr) GetFromUseState(useState int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`use_state` = ?", useState).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUseState 批量查找 0使用中 1未使用
func (obj *_WmsSorterGridSubRuleMgr) GetBatchFromUseState(useStates []int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`use_state` IN (?)", useStates).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
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
func (obj *_WmsSorterGridSubRuleMgr) FetchByPrimaryKey(id int) (result WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", result.GridRuleID).Find(&result.WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByGridRuleIDFk  获取多个内容
func (obj *_WmsSorterGridSubRuleMgr) FetchIndexByGridRuleIDFk(gridRuleID int) (results []*WmsSorterGridSubRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridSubRule{}).Where("`grid_rule_id` = ?", gridRuleID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("wms_sorter_grid_rule").Where("id = ?", results[i].GridRuleID).Find(&results[i].WmsSorterGridRule).Error; err != nil { // 分拣机-格口-规则
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
