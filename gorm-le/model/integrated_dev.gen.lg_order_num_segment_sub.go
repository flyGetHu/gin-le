package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _LgOrderNumSegmentSubMgr struct {
	*_BaseMgr
}

// LgOrderNumSegmentSubMgr open func
func LgOrderNumSegmentSubMgr(db *gorm.DB) *_LgOrderNumSegmentSubMgr {
	if db == nil {
		panic(fmt.Errorf("LgOrderNumSegmentSubMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_LgOrderNumSegmentSubMgr{_BaseMgr: &_BaseMgr{DB: db.Table("lg_order_num_segment_sub"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_LgOrderNumSegmentSubMgr) GetTableName() string {
	return "lg_order_num_segment_sub"
}

// Reset 重置gorm会话
func (obj *_LgOrderNumSegmentSubMgr) Reset() *_LgOrderNumSegmentSubMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_LgOrderNumSegmentSubMgr) Get() (result LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", result.OrderNumSegmentID).Find(&result.LgOrderNumSegment).Error; err != nil { // 号码段主表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_LgOrderNumSegmentSubMgr) Gets() (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_LgOrderNumSegmentSubMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 号码段附表ID
func (obj *_LgOrderNumSegmentSubMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderNumSegmentID order_num_segment_id获取 号段主表关联ID
func (obj *_LgOrderNumSegmentSubMgr) WithOrderNumSegmentID(orderNumSegmentID int64) Option {
	return optionFunc(func(o *options) { o.query["order_num_segment_id"] = orderNumSegmentID })
}

// WithStartNum start_num获取 号段提取开始位置,数字部分
func (obj *_LgOrderNumSegmentSubMgr) WithStartNum(startNum int64) Option {
	return optionFunc(func(o *options) { o.query["start_num"] = startNum })
}

// WithEndNum end_num获取 号段提取结束位置,数字部分,包头包尾
func (obj *_LgOrderNumSegmentSubMgr) WithEndNum(endNum int64) Option {
	return optionFunc(func(o *options) { o.query["end_num"] = endNum })
}

// WithType type获取 提取类型:self:自用 customer:客户使用
func (obj *_LgOrderNumSegmentSubMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithRemark remark获取 备注
func (obj *_LgOrderNumSegmentSubMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_LgOrderNumSegmentSubMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_LgOrderNumSegmentSubMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_LgOrderNumSegmentSubMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_LgOrderNumSegmentSubMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_LgOrderNumSegmentSubMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_LgOrderNumSegmentSubMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_LgOrderNumSegmentSubMgr) GetByOption(opts ...Option) (result LgOrderNumSegmentSub, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", result.OrderNumSegmentID).Find(&result.LgOrderNumSegment).Error; err != nil { // 号码段主表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_LgOrderNumSegmentSubMgr) GetByOptions(opts ...Option) (results []*LgOrderNumSegmentSub, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_LgOrderNumSegmentSubMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]LgOrderNumSegmentSub, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
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

// GetFromID 通过id获取内容 号码段附表ID
func (obj *_LgOrderNumSegmentSubMgr) GetFromID(id int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`id` = ?", id).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromID 批量查找 号码段附表ID
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromID(ids []int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrderNumSegmentID 通过order_num_segment_id获取内容 号段主表关联ID
func (obj *_LgOrderNumSegmentSubMgr) GetFromOrderNumSegmentID(orderNumSegmentID int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`order_num_segment_id` = ?", orderNumSegmentID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrderNumSegmentID 批量查找 号段主表关联ID
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromOrderNumSegmentID(orderNumSegmentIDs []int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`order_num_segment_id` IN (?)", orderNumSegmentIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStartNum 通过start_num获取内容 号段提取开始位置,数字部分
func (obj *_LgOrderNumSegmentSubMgr) GetFromStartNum(startNum int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`start_num` = ?", startNum).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStartNum 批量查找 号段提取开始位置,数字部分
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromStartNum(startNums []int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`start_num` IN (?)", startNums).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEndNum 通过end_num获取内容 号段提取结束位置,数字部分,包头包尾
func (obj *_LgOrderNumSegmentSubMgr) GetFromEndNum(endNum int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`end_num` = ?", endNum).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEndNum 批量查找 号段提取结束位置,数字部分,包头包尾
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromEndNum(endNums []int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`end_num` IN (?)", endNums).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromType 通过type获取内容 提取类型:self:自用 customer:客户使用
func (obj *_LgOrderNumSegmentSubMgr) GetFromType(_type string) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`type` = ?", _type).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromType 批量查找 提取类型:self:自用 customer:客户使用
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromType(_types []string) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`type` IN (?)", _types).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_LgOrderNumSegmentSubMgr) GetFromRemark(remark string) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`remark` = ?", remark).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromRemark(remarks []string) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`remark` IN (?)", remarks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_LgOrderNumSegmentSubMgr) GetFromCreateTime(createTime time.Time) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_LgOrderNumSegmentSubMgr) GetFromCreateUser(createUser int) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromCreateUser(createUsers []int) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_LgOrderNumSegmentSubMgr) GetFromUpdateTime(updateTime time.Time) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_LgOrderNumSegmentSubMgr) GetFromUpdateUser(updateUser int) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_LgOrderNumSegmentSubMgr) GetFromVersion(version int) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromVersion(versions []int) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_LgOrderNumSegmentSubMgr) GetFromDeleted(deleted int) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_LgOrderNumSegmentSubMgr) GetBatchFromDeleted(deleteds []int) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchIndexByLgOrderNumSegmentSubLgOrderNumSegmentIDFk  获取多个内容
func (obj *_LgOrderNumSegmentSubMgr) FetchIndexByLgOrderNumSegmentSubLgOrderNumSegmentIDFk(orderNumSegmentID int64) (results []*LgOrderNumSegmentSub, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(LgOrderNumSegmentSub{}).Where("`order_num_segment_id` = ?", orderNumSegmentID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("lg_order_num_segment").Where("id = ?", results[i].OrderNumSegmentID).Find(&results[i].LgOrderNumSegment).Error; err != nil { // 号码段主表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
