package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WWorkOrderTaskMgr struct {
	*_BaseMgr
}

// WWorkOrderTaskMgr open func
func WWorkOrderTaskMgr(db *gorm.DB) *_WWorkOrderTaskMgr {
	if db == nil {
		panic(fmt.Errorf("WWorkOrderTaskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WWorkOrderTaskMgr{_BaseMgr: &_BaseMgr{DB: db.Table("w_work_order_task"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WWorkOrderTaskMgr) GetTableName() string {
	return "w_work_order_task"
}

// Reset 重置gorm会话
func (obj *_WWorkOrderTaskMgr) Reset() *_WWorkOrderTaskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WWorkOrderTaskMgr) Get() (result WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("w_work_order").Where("id = ?", result.OrderID).Find(&result.WWorkOrder).Error; err != nil { // 工单表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", result.TagName).Find(&result.WWorkOrderTag).Error; err != nil { // 标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_WWorkOrderTaskMgr) Gets() (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_WWorkOrderTaskMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 任务编号
func (obj *_WWorkOrderTaskMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithOrderID order_id获取 工单ID
func (obj *_WWorkOrderTaskMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithTitle title获取 标题
func (obj *_WWorkOrderTaskMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 内容
func (obj *_WWorkOrderTaskMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithUserID user_id获取 实施人
func (obj *_WWorkOrderTaskMgr) WithUserID(userID int) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithOrganizationID organization_id获取 工单所属组织机构id
func (obj *_WWorkOrderTaskMgr) WithOrganizationID(organizationID int) Option {
	return optionFunc(func(o *options) { o.query["organization_id"] = organizationID })
}

// WithState state获取 任务状态
func (obj *_WWorkOrderTaskMgr) WithState(state string) Option {
	return optionFunc(func(o *options) { o.query["state"] = state })
}

// WithSort sort获取 优先级
//
//	紧急，高，正常，低
func (obj *_WWorkOrderTaskMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithEstimateFinishTime estimate_finish_time获取 预计完成时间
func (obj *_WWorkOrderTaskMgr) WithEstimateFinishTime(estimateFinishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["estimate_finish_time"] = estimateFinishTime })
}

// WithTaskBeginTime task_begin_time获取 任务开始时间
func (obj *_WWorkOrderTaskMgr) WithTaskBeginTime(taskBeginTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["task_begin_time"] = taskBeginTime })
}

// WithActualFinishTime actual_finish_time获取 实际完成时间
func (obj *_WWorkOrderTaskMgr) WithActualFinishTime(actualFinishTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["actual_finish_time"] = actualFinishTime })
}

// WithEstimateWorkingHours estimate_working_hours获取 预计完成工时
func (obj *_WWorkOrderTaskMgr) WithEstimateWorkingHours(estimateWorkingHours string) Option {
	return optionFunc(func(o *options) { o.query["estimate_working_hours"] = estimateWorkingHours })
}

// WithActualWorkingHours actual_working_hours获取 实际完成使用工时
func (obj *_WWorkOrderTaskMgr) WithActualWorkingHours(actualWorkingHours string) Option {
	return optionFunc(func(o *options) { o.query["actual_working_hours"] = actualWorkingHours })
}

// WithRemark remark获取 备注
func (obj *_WWorkOrderTaskMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithTagName tag_name获取 标签名称
func (obj *_WWorkOrderTaskMgr) WithTagName(tagName string) Option {
	return optionFunc(func(o *options) { o.query["tag_name"] = tagName })
}

// WithEvaluationGrade evaluation_grade获取 评价等级
func (obj *_WWorkOrderTaskMgr) WithEvaluationGrade(evaluationGrade int) Option {
	return optionFunc(func(o *options) { o.query["evaluation_grade"] = evaluationGrade })
}

// WithEvaluationComment evaluation_comment获取 评价内容
func (obj *_WWorkOrderTaskMgr) WithEvaluationComment(evaluationComment string) Option {
	return optionFunc(func(o *options) { o.query["evaluation_comment"] = evaluationComment })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WWorkOrderTaskMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WWorkOrderTaskMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WWorkOrderTaskMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_WWorkOrderTaskMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WWorkOrderTaskMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WWorkOrderTaskMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WWorkOrderTaskMgr) GetByOption(opts ...Option) (result WWorkOrderTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("w_work_order").Where("id = ?", result.OrderID).Find(&result.WWorkOrder).Error; err != nil { // 工单表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", result.TagName).Find(&result.WWorkOrderTag).Error; err != nil { // 标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WWorkOrderTaskMgr) GetByOptions(opts ...Option) (results []*WWorkOrderTask, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WWorkOrderTask, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromID(id int) (result WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("w_work_order").Where("id = ?", result.OrderID).Find(&result.WWorkOrder).Error; err != nil { // 工单表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", result.TagName).Find(&result.WWorkOrderTag).Error; err != nil { // 标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_WWorkOrderTaskMgr) GetBatchFromID(ids []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCode 通过code获取内容 任务编号
func (obj *_WWorkOrderTaskMgr) GetFromCode(code string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`code` = ?", code).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCode 批量查找 任务编号
func (obj *_WWorkOrderTaskMgr) GetBatchFromCode(codes []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`code` IN (?)", codes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrderID 通过order_id获取内容 工单ID
func (obj *_WWorkOrderTaskMgr) GetFromOrderID(orderID int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`order_id` = ?", orderID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrderID 批量查找 工单ID
func (obj *_WWorkOrderTaskMgr) GetBatchFromOrderID(orderIDs []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromTitle(title string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`title` = ?", title).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromTitle(titles []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`title` IN (?)", titles).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromContent(content string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`content` = ?", content).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromContent(contents []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`content` IN (?)", contents).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserID 通过user_id获取内容 实施人
func (obj *_WWorkOrderTaskMgr) GetFromUserID(userID int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`user_id` = ?", userID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserID 批量查找 实施人
func (obj *_WWorkOrderTaskMgr) GetBatchFromUserID(userIDs []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrganizationID 通过organization_id获取内容 工单所属组织机构id
func (obj *_WWorkOrderTaskMgr) GetFromOrganizationID(organizationID int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`organization_id` = ?", organizationID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrganizationID 批量查找 工单所属组织机构id
func (obj *_WWorkOrderTaskMgr) GetBatchFromOrganizationID(organizationIDs []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`organization_id` IN (?)", organizationIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromState 通过state获取内容 任务状态
func (obj *_WWorkOrderTaskMgr) GetFromState(state string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`state` = ?", state).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromState 批量查找 任务状态
func (obj *_WWorkOrderTaskMgr) GetBatchFromState(states []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`state` IN (?)", states).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromSort(sort int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`sort` = ?", sort).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromSort(sorts []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`sort` IN (?)", sorts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromEstimateFinishTime(estimateFinishTime time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`estimate_finish_time` = ?", estimateFinishTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromEstimateFinishTime(estimateFinishTimes []time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`estimate_finish_time` IN (?)", estimateFinishTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromTaskBeginTime 通过task_begin_time获取内容 任务开始时间
func (obj *_WWorkOrderTaskMgr) GetFromTaskBeginTime(taskBeginTime time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`task_begin_time` = ?", taskBeginTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromTaskBeginTime 批量查找 任务开始时间
func (obj *_WWorkOrderTaskMgr) GetBatchFromTaskBeginTime(taskBeginTimes []time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`task_begin_time` IN (?)", taskBeginTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromActualFinishTime(actualFinishTime time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`actual_finish_time` = ?", actualFinishTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromActualFinishTime(actualFinishTimes []time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`actual_finish_time` IN (?)", actualFinishTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEstimateWorkingHours 通过estimate_working_hours获取内容 预计完成工时
func (obj *_WWorkOrderTaskMgr) GetFromEstimateWorkingHours(estimateWorkingHours string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`estimate_working_hours` = ?", estimateWorkingHours).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEstimateWorkingHours 批量查找 预计完成工时
func (obj *_WWorkOrderTaskMgr) GetBatchFromEstimateWorkingHours(estimateWorkingHourss []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`estimate_working_hours` IN (?)", estimateWorkingHourss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromActualWorkingHours 通过actual_working_hours获取内容 实际完成使用工时
func (obj *_WWorkOrderTaskMgr) GetFromActualWorkingHours(actualWorkingHours string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`actual_working_hours` = ?", actualWorkingHours).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromActualWorkingHours 批量查找 实际完成使用工时
func (obj *_WWorkOrderTaskMgr) GetBatchFromActualWorkingHours(actualWorkingHourss []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`actual_working_hours` IN (?)", actualWorkingHourss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromRemark(remark string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`remark` = ?", remark).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromRemark(remarks []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`remark` IN (?)", remarks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromTagName(tagName string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`tag_name` = ?", tagName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromTagName(tagNames []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`tag_name` IN (?)", tagNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEvaluationGrade 通过evaluation_grade获取内容 评价等级
func (obj *_WWorkOrderTaskMgr) GetFromEvaluationGrade(evaluationGrade int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`evaluation_grade` = ?", evaluationGrade).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEvaluationGrade 批量查找 评价等级
func (obj *_WWorkOrderTaskMgr) GetBatchFromEvaluationGrade(evaluationGrades []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`evaluation_grade` IN (?)", evaluationGrades).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEvaluationComment 通过evaluation_comment获取内容 评价内容
func (obj *_WWorkOrderTaskMgr) GetFromEvaluationComment(evaluationComment string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`evaluation_comment` = ?", evaluationComment).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEvaluationComment 批量查找 评价内容
func (obj *_WWorkOrderTaskMgr) GetBatchFromEvaluationComment(evaluationComments []string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`evaluation_comment` IN (?)", evaluationComments).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromCreateTime(createTime time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromCreateUser(createUser int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromCreateUser(createUsers []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromUpdateTime(updateTime time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromUpdateUser(updateUser int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromVersion(version int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromVersion(versions []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetFromDeleted(deleted int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) GetBatchFromDeleted(deleteds []int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
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
func (obj *_WWorkOrderTaskMgr) FetchByPrimaryKey(id int) (result WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("w_work_order").Where("id = ?", result.OrderID).Find(&result.WWorkOrder).Error; err != nil { // 工单表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", result.TagName).Find(&result.WWorkOrderTag).Error; err != nil { // 标签维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByWTaskWOrderIDFk  获取多个内容
func (obj *_WWorkOrderTaskMgr) FetchIndexByWTaskWOrderIDFk(orderID int) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`order_id` = ?", orderID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByWTaskWTagNameFk  获取多个内容
func (obj *_WWorkOrderTaskMgr) FetchIndexByWTaskWTagNameFk(tagName string) (results []*WWorkOrderTask, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WWorkOrderTask{}).Where("`tag_name` = ?", tagName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("w_work_order").Where("id = ?", results[i].OrderID).Find(&results[i].WWorkOrder).Error; err != nil { // 工单表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("w_work_order_tag").Where("name = ?", results[i].TagName).Find(&results[i].WWorkOrderTag).Error; err != nil { // 标签维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
