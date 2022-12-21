package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _PdGoodsRiskControlMgr struct {
	*_BaseMgr
}

// PdGoodsRiskControlMgr open func
func PdGoodsRiskControlMgr(db *gorm.DB) *_PdGoodsRiskControlMgr {
	if db == nil {
		panic(fmt.Errorf("PdGoodsRiskControlMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_PdGoodsRiskControlMgr{_BaseMgr: &_BaseMgr{DB: db.Table("pd_goods_risk_control"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PdGoodsRiskControlMgr) GetTableName() string {
	return "pd_goods_risk_control"
}

// Reset 重置gorm会话
func (obj *_PdGoodsRiskControlMgr) Reset() *_PdGoodsRiskControlMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PdGoodsRiskControlMgr) Get() (result PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PdGoodsRiskControlMgr) Gets() (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_PdGoodsRiskControlMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_PdGoodsRiskControlMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGoodsName goods_name获取 商品英文名称
func (obj *_PdGoodsRiskControlMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsNameCn goods_name_cn获取 商品中文名称
func (obj *_PdGoodsRiskControlMgr) WithGoodsNameCn(goodsNameCn string) Option {
	return optionFunc(func(o *options) { o.query["goods_name_cn"] = goodsNameCn })
}

// WithShortGoodsName short_goods_name获取 商品英文名称(简称)
func (obj *_PdGoodsRiskControlMgr) WithShortGoodsName(shortGoodsName string) Option {
	return optionFunc(func(o *options) { o.query["short_goods_name"] = shortGoodsName })
}

// WithPrice price获取 商品价值/单价
func (obj *_PdGoodsRiskControlMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithMinPrice min_price获取 最低价值/单价
func (obj *_PdGoodsRiskControlMgr) WithMinPrice(minPrice float64) Option {
	return optionFunc(func(o *options) { o.query["min_price"] = minPrice })
}

// WithLimitOut limit_out获取 限制出库 默认否(不限制)
func (obj *_PdGoodsRiskControlMgr) WithLimitOut(limitOut int) Option {
	return optionFunc(func(o *options) { o.query["limit_out"] = limitOut })
}

// WithGoodsURL goods_url获取 产品url
func (obj *_PdGoodsRiskControlMgr) WithGoodsURL(goodsURL string) Option {
	return optionFunc(func(o *options) { o.query["goods_url"] = goodsURL })
}

// WithHsCode hs_code获取 海关编码
func (obj *_PdGoodsRiskControlMgr) WithHsCode(hsCode string) Option {
	return optionFunc(func(o *options) { o.query["hs_code"] = hsCode })
}

// WithRemark remark获取 备注
func (obj *_PdGoodsRiskControlMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_PdGoodsRiskControlMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_PdGoodsRiskControlMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_PdGoodsRiskControlMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新时间
func (obj *_PdGoodsRiskControlMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_PdGoodsRiskControlMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_PdGoodsRiskControlMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_PdGoodsRiskControlMgr) GetByOption(opts ...Option) (result PdGoodsRiskControl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_PdGoodsRiskControlMgr) GetByOptions(opts ...Option) (results []*PdGoodsRiskControl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_PdGoodsRiskControlMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]PdGoodsRiskControl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 主键
func (obj *_PdGoodsRiskControlMgr) GetFromID(id int) (result PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_PdGoodsRiskControlMgr) GetBatchFromID(ids []int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品英文名称
func (obj *_PdGoodsRiskControlMgr) GetFromGoodsName(goodsName string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品英文名称
func (obj *_PdGoodsRiskControlMgr) GetBatchFromGoodsName(goodsNames []string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsNameCn 通过goods_name_cn获取内容 商品中文名称
func (obj *_PdGoodsRiskControlMgr) GetFromGoodsNameCn(goodsNameCn string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`goods_name_cn` = ?", goodsNameCn).Find(&results).Error

	return
}

// GetBatchFromGoodsNameCn 批量查找 商品中文名称
func (obj *_PdGoodsRiskControlMgr) GetBatchFromGoodsNameCn(goodsNameCns []string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`goods_name_cn` IN (?)", goodsNameCns).Find(&results).Error

	return
}

// GetFromShortGoodsName 通过short_goods_name获取内容 商品英文名称(简称)
func (obj *_PdGoodsRiskControlMgr) GetFromShortGoodsName(shortGoodsName string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`short_goods_name` = ?", shortGoodsName).Find(&results).Error

	return
}

// GetBatchFromShortGoodsName 批量查找 商品英文名称(简称)
func (obj *_PdGoodsRiskControlMgr) GetBatchFromShortGoodsName(shortGoodsNames []string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`short_goods_name` IN (?)", shortGoodsNames).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品价值/单价
func (obj *_PdGoodsRiskControlMgr) GetFromPrice(price float64) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品价值/单价
func (obj *_PdGoodsRiskControlMgr) GetBatchFromPrice(prices []float64) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromMinPrice 通过min_price获取内容 最低价值/单价
func (obj *_PdGoodsRiskControlMgr) GetFromMinPrice(minPrice float64) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`min_price` = ?", minPrice).Find(&results).Error

	return
}

// GetBatchFromMinPrice 批量查找 最低价值/单价
func (obj *_PdGoodsRiskControlMgr) GetBatchFromMinPrice(minPrices []float64) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`min_price` IN (?)", minPrices).Find(&results).Error

	return
}

// GetFromLimitOut 通过limit_out获取内容 限制出库 默认否(不限制)
func (obj *_PdGoodsRiskControlMgr) GetFromLimitOut(limitOut int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`limit_out` = ?", limitOut).Find(&results).Error

	return
}

// GetBatchFromLimitOut 批量查找 限制出库 默认否(不限制)
func (obj *_PdGoodsRiskControlMgr) GetBatchFromLimitOut(limitOuts []int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`limit_out` IN (?)", limitOuts).Find(&results).Error

	return
}

// GetFromGoodsURL 通过goods_url获取内容 产品url
func (obj *_PdGoodsRiskControlMgr) GetFromGoodsURL(goodsURL string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`goods_url` = ?", goodsURL).Find(&results).Error

	return
}

// GetBatchFromGoodsURL 批量查找 产品url
func (obj *_PdGoodsRiskControlMgr) GetBatchFromGoodsURL(goodsURLs []string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`goods_url` IN (?)", goodsURLs).Find(&results).Error

	return
}

// GetFromHsCode 通过hs_code获取内容 海关编码
func (obj *_PdGoodsRiskControlMgr) GetFromHsCode(hsCode string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`hs_code` = ?", hsCode).Find(&results).Error

	return
}

// GetBatchFromHsCode 批量查找 海关编码
func (obj *_PdGoodsRiskControlMgr) GetBatchFromHsCode(hsCodes []string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`hs_code` IN (?)", hsCodes).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_PdGoodsRiskControlMgr) GetFromRemark(remark string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_PdGoodsRiskControlMgr) GetBatchFromRemark(remarks []string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_PdGoodsRiskControlMgr) GetFromCreateTime(createTime time.Time) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_PdGoodsRiskControlMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_PdGoodsRiskControlMgr) GetFromCreateUser(createUser int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_PdGoodsRiskControlMgr) GetBatchFromCreateUser(createUsers []int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_PdGoodsRiskControlMgr) GetFromUpdateTime(updateTime time.Time) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_PdGoodsRiskControlMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新时间
func (obj *_PdGoodsRiskControlMgr) GetFromUpdateUser(updateUser int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新时间
func (obj *_PdGoodsRiskControlMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_PdGoodsRiskControlMgr) GetFromVersion(version int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_PdGoodsRiskControlMgr) GetBatchFromVersion(versions []int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_PdGoodsRiskControlMgr) GetFromDeleted(deleted int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_PdGoodsRiskControlMgr) GetBatchFromDeleted(deleteds []int) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_PdGoodsRiskControlMgr) FetchByPrimaryKey(id int) (result PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByUniqueGoodsName primary or index 获取唯一内容
func (obj *_PdGoodsRiskControlMgr) FetchUniqueByUniqueGoodsName(goodsName string) (result PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`goods_name` = ?", goodsName).First(&result).Error

	return
}

// FetchIndexByIndexGoodsName  获取多个内容
func (obj *_PdGoodsRiskControlMgr) FetchIndexByIndexGoodsName(goodsName string) (results []*PdGoodsRiskControl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PdGoodsRiskControl{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}
