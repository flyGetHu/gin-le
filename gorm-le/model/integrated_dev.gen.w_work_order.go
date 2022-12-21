package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WWorkOrderMgr struct {
	*_BaseMgr
}

// WWorkOrderMgr open func
func WWorkOrderMgr(db *gorm.DB) *_WWorkOrderMgr {
	if db == nil {
		panic(fmt.Errorf("WWorkOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WWorkOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_work_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WWorkOrderMgr) GetTableName() string {
	return "w_work_order"
}

// Reset 重置gorm会话
func (obj *_WWorkOrderMgr) Reset() *_WWorkOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WWorkOrderMgr) Get() (result WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", result.TagName).Find(&result.WWorkOrderTag).Error; err != nil { // 标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WWorkOrderMgr) Gets() (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WWorkOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_WWorkOrderMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 工单编号
func (obj *_WWorkOrderMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithTitle title获取 标题
func (obj *_WWorkOrderMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_WWorkOrderMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithHistoricalContent historical_content获取 历史内容
func (obj *_WWorkOrderMgr) WithHistoricalContent(historicalContent string) Option {
	return optionFunc(func(o *options) { o.query["historical_content"] = historicalContent })
}

// WithOrganizationID organization_id获取 组织机构ID
func (obj *_WWorkOrderMgr) WithOrganizationID(organizationID int) Option {
	return optionFunc(func(o *options) { o.query["organization_id"] = organizationID })
}

// WithOrganizationIDs organization_ids获取 组织机构ids
func (obj *_WWorkOrderMgr) WithOrganizationIDs(organizationIDs string) Option {
	return optionFunc(func(o *options) { o.query["organization_ids"] = organizationIDs })
}

// WithReviewerID reviewer_id获取 审核人ID
func (obj *_WWorkOrderMgr) WithReviewerID(reviewerID int) Option {
	return optionFunc(func(o *options) { o.query["reviewer_id"] = reviewerID })
}

// WithResources resources获取 资源列表
func (obj *_WWorkOrderMgr) WithResources(resources string) Option {
	return optionFunc(func(o *options) { o.query["resources"] = resources })
}

// WithState state获取 工单状态
//
//	草稿:新建但未指定部门一created
//	审核中:新建并指定审核部门- reviewing
//	审核通过:部门审核通过,但未拆分任务，指定实施人一reviewed
//	审核驳回:部门审核人审核不通过,驳回一review_callback
//	任务关闭:审核驳回后需求放关闭工单一closed
//	实施中:拆分任务指定实施人员一implementing
//	实施完成:所有任务完成一implemented
//	审核人验收:部门审核人验收一reviewer_acceptance
//	需求方验收-demander_ acceptance
//	终结 over
func (obj *_WWorkOrderMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithSort sort获取 优先级
//
//	紧急，高，正常，低
func (obj *_WWorkOrderMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithType type获取 工单类型
func (obj *_WWorkOrderMgr) WithType(_type int) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithEstimateFinishTime estimate_finish_time获取 预计完成时间
func (obj *_WWorkOrderMgr) WithEstimateFinishTime(estimateFinishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["estimate_finish_time"] = estimateFinishTime })
}

// WithActualFinishTime actual_finish_time获取 实际完成时间
func (obj *_WWorkOrderMgr) WithActualFinishTime(actualFinishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["actual_finish_time"] = actualFinishTime })
}

// WithRemark remark获取 备注
func (obj *_WWorkOrderMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithIsCooperation is_cooperation获取 是否为多部门协作
func (obj *_WWorkOrderMgr) WithIsCooperation(isCooperation int) Option {
	return optionFunc(func(o *options) { o.query["is_cooperation"] = isCooperation })
}

// WithCooperationOrgIDs cooperation_org_ids获取 多部门协作部门ids集合
func (obj *_WWorkOrderMgr) WithCooperationOrgIDs(cooperationOrgIDs string) Option {
	return optionFunc(func(o *options) { o.query["cooperation_org_ids"] = cooperationOrgIDs })
}

// WithTagName tag_name获取 标签名称
func (obj *_WWorkOrderMgr) WithTagName(tagName string) Option {
	return optionFunc(func(o *options) { o.query["tag_name"] = tagName })
}

// WithEvaluationGrade evaluation_grade获取 评价等级1-10
func (obj *_WWorkOrderMgr) WithEvaluationGrade(evaluationGrade int) Option {
	return optionFunc(func(o *options) { o.query["evaluation_grade"] = evaluationGrade })
}

// WithEvaluationComment evaluation_comment获取 评价留言
func (obj *_WWorkOrderMgr) WithEvaluationComment(evaluationComment string) Option {
	return optionFunc(func(o *options) { o.query["evaluation_comment"] = evaluationComment })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WWorkOrderMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WWorkOrderMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WWorkOrderMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WWorkOrderMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WWorkOrderMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WWorkOrderMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WWorkOrderMgr) GetByOption(opts ...Option) (result WWorkOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", result.TagName).Find(&result.WWorkOrderTag).Error; err != nil { // 标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WWorkOrderMgr) GetByOptions(opts ...Option) (results []*WWorkOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_WWorkOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WWorkOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
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
func (obj *_WWorkOrderMgr) GetFromID(id int) (result WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", result.TagName).Find(&result.WWorkOrderTag).Error; err != nil { // 标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_WWorkOrderMgr) GetBatchFromID(ids []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCode 通过code获取内容 工单编号
func (obj *_WWorkOrderMgr) GetFromCode(code string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`code` = ?", code).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCode 批量查找 工单编号
func (obj *_WWorkOrderMgr) GetBatchFromCode(codes []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`code` IN (?)", codes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *_WWorkOrderMgr) GetFromTitle(title string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`title` = ?", title).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *_WWorkOrderMgr) GetBatchFromTitle(titles []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`title` IN (?)", titles).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromContent 通过content获取内容 内容
func (obj *_WWorkOrderMgr) GetFromContent(content string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`content` = ?", content).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromContent 批量查找 内容
func (obj *_WWorkOrderMgr) GetBatchFromContent(contents []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`content` IN (?)", contents).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromHistoricalContent 通过historical_content获取内容 历史内容
func (obj *_WWorkOrderMgr) GetFromHistoricalContent(historicalContent string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`historical_content` = ?", historicalContent).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromHistoricalContent 批量查找 历史内容
func (obj *_WWorkOrderMgr) GetBatchFromHistoricalContent(historicalContents []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`historical_content` IN (?)", historicalContents).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrganizationID 通过organization_id获取内容 组织机构ID
func (obj *_WWorkOrderMgr) GetFromOrganizationID(organizationID int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`organization_id` = ?", organizationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrganizationID 批量查找 组织机构ID
func (obj *_WWorkOrderMgr) GetBatchFromOrganizationID(organizationIDs []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`organization_id` IN (?)", organizationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrganizationIDs 通过organization_ids获取内容 组织机构ids
func (obj *_WWorkOrderMgr) GetFromOrganizationIDs(organizationIDs string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`organization_ids` = ?", organizationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrganizationIDs 批量查找 组织机构ids
func (obj *_WWorkOrderMgr) GetBatchFromOrganizationIDs(organizationIDss []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`organization_ids` IN (?)", organizationIDss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromReviewerID 通过reviewer_id获取内容 审核人ID
func (obj *_WWorkOrderMgr) GetFromReviewerID(reviewerID int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`reviewer_id` = ?", reviewerID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromReviewerID 批量查找 审核人ID
func (obj *_WWorkOrderMgr) GetBatchFromReviewerID(reviewerIDs []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`reviewer_id` IN (?)", reviewerIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromResources 通过resources获取内容 资源列表
func (obj *_WWorkOrderMgr) GetFromResources(resources string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`resources` = ?", resources).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromResources 批量查找 资源列表
func (obj *_WWorkOrderMgr) GetBatchFromResources(resourcess []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`resources` IN (?)", resourcess).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromState 通过state获取内容 工单状态
//
//	草稿:新建但未指定部门一created
//	审核中:新建并指定审核部门- reviewing
//	审核通过:部门审核通过,但未拆分任务，指定实施人一reviewed
//	审核驳回:部门审核人审核不通过,驳回一review_callback
//	任务关闭:审核驳回后需求放关闭工单一closed
//	实施中:拆分任务指定实施人员一implementing
//	实施完成:所有任务完成一implemented
//	审核人验收:部门审核人验收一reviewer_acceptance
//	需求方验收-demander_ acceptance
//	终结 over
func (obj *_WWorkOrderMgr) GetFromState(state string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`state` = ?", state).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromState 批量查找 工单状态
//
//	草稿:新建但未指定部门一created
//	审核中:新建并指定审核部门- reviewing
//	审核通过:部门审核通过,但未拆分任务，指定实施人一reviewed
//	审核驳回:部门审核人审核不通过,驳回一review_callback
//	任务关闭:审核驳回后需求放关闭工单一closed
//	实施中:拆分任务指定实施人员一implementing
//	实施完成:所有任务完成一implemented
//	审核人验收:部门审核人验收一reviewer_acceptance
//	需求方验收-demander_ acceptance
//	终结 over
func (obj *_WWorkOrderMgr) GetBatchFromState(states []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`state` IN (?)", states).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSort 通过sort获取内容 优先级
//
//	紧急，高，正常，低
func (obj *_WWorkOrderMgr) GetFromSort(sort int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`sort` = ?", sort).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSort 批量查找 优先级
//
//	紧急，高，正常，低
func (obj *_WWorkOrderMgr) GetBatchFromSort(sorts []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`sort` IN (?)", sorts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromType 通过type获取内容 工单类型
func (obj *_WWorkOrderMgr) GetFromType(_type int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`type` = ?", _type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromType 批量查找 工单类型
func (obj *_WWorkOrderMgr) GetBatchFromType(_types []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`type` IN (?)", _types).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEstimateFinishTime 通过estimate_finish_time获取内容 预计完成时间
func (obj *_WWorkOrderMgr) GetFromEstimateFinishTime(estimateFinishTime time.Time) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`estimate_finish_time` = ?", estimateFinishTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEstimateFinishTime 批量查找 预计完成时间
func (obj *_WWorkOrderMgr) GetBatchFromEstimateFinishTime(estimateFinishTimes []time.Time) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`estimate_finish_time` IN (?)", estimateFinishTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromActualFinishTime 通过actual_finish_time获取内容 实际完成时间
func (obj *_WWorkOrderMgr) GetFromActualFinishTime(actualFinishTime time.Time) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`actual_finish_time` = ?", actualFinishTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromActualFinishTime 批量查找 实际完成时间
func (obj *_WWorkOrderMgr) GetBatchFromActualFinishTime(actualFinishTimes []time.Time) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`actual_finish_time` IN (?)", actualFinishTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WWorkOrderMgr) GetFromRemark(remark string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`remark` = ?", remark).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WWorkOrderMgr) GetBatchFromRemark(remarks []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`remark` IN (?)", remarks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIsCooperation 通过is_cooperation获取内容 是否为多部门协作
func (obj *_WWorkOrderMgr) GetFromIsCooperation(isCooperation int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`is_cooperation` = ?", isCooperation).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIsCooperation 批量查找 是否为多部门协作
func (obj *_WWorkOrderMgr) GetBatchFromIsCooperation(isCooperations []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`is_cooperation` IN (?)", isCooperations).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCooperationOrgIDs 通过cooperation_org_ids获取内容 多部门协作部门ids集合
func (obj *_WWorkOrderMgr) GetFromCooperationOrgIDs(cooperationOrgIDs string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`cooperation_org_ids` = ?", cooperationOrgIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCooperationOrgIDs 批量查找 多部门协作部门ids集合
func (obj *_WWorkOrderMgr) GetBatchFromCooperationOrgIDs(cooperationOrgIDss []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`cooperation_org_ids` IN (?)", cooperationOrgIDss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTagName 通过tag_name获取内容 标签名称
func (obj *_WWorkOrderMgr) GetFromTagName(tagName string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`tag_name` = ?", tagName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTagName 批量查找 标签名称
func (obj *_WWorkOrderMgr) GetBatchFromTagName(tagNames []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`tag_name` IN (?)", tagNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEvaluationGrade 通过evaluation_grade获取内容 评价等级1-10
func (obj *_WWorkOrderMgr) GetFromEvaluationGrade(evaluationGrade int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`evaluation_grade` = ?", evaluationGrade).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEvaluationGrade 批量查找 评价等级1-10
func (obj *_WWorkOrderMgr) GetBatchFromEvaluationGrade(evaluationGrades []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`evaluation_grade` IN (?)", evaluationGrades).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEvaluationComment 通过evaluation_comment获取内容 评价留言
func (obj *_WWorkOrderMgr) GetFromEvaluationComment(evaluationComment string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`evaluation_comment` = ?", evaluationComment).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEvaluationComment 批量查找 评价留言
func (obj *_WWorkOrderMgr) GetBatchFromEvaluationComment(evaluationComments []string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`evaluation_comment` IN (?)", evaluationComments).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WWorkOrderMgr) GetFromCreateTime(createTime time.Time) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WWorkOrderMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WWorkOrderMgr) GetFromCreateUser(createUser int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WWorkOrderMgr) GetBatchFromCreateUser(createUsers []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WWorkOrderMgr) GetFromUpdateTime(updateTime time.Time) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WWorkOrderMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_WWorkOrderMgr) GetFromUpdateUser(updateUser int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_WWorkOrderMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WWorkOrderMgr) GetFromVersion(version int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WWorkOrderMgr) GetBatchFromVersion(versions []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WWorkOrderMgr) GetFromDeleted(deleted int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WWorkOrderMgr) GetBatchFromDeleted(deleteds []int) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
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
func (obj *_WWorkOrderMgr) FetchByPrimaryKey(id int) (result WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", result.TagName).Find(&result.WWorkOrderTag).Error; err != nil { // 标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByWOrderWTagNameFk  获取多个内容
func (obj *_WWorkOrderMgr) FetchIndexByWOrderWTagNameFk(tagName string) (results []*WWorkOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrder{}).Where("`tag_name` = ?", tagName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
