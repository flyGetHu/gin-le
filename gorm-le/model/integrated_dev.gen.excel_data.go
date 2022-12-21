package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ExcelDataMgr struct {
	*_BaseMgr
}

// ExcelDataMgr open func
func ExcelDataMgr(db *gorm.DB) *_ExcelDataMgr {
	if db == nil {
		panic(fmt.Errorf("ExcelDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ExcelDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("excel_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ExcelDataMgr) GetTableName() string {
	return "excel_data"
}

// Reset 重置gorm会话
func (obj *_ExcelDataMgr) Reset() *_ExcelDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ExcelDataMgr) Get() (result ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ExcelDataMgr) Gets() (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ExcelDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_ExcelDataMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithA1 a1获取
func (obj *_ExcelDataMgr) WithA1(a1 string) Option {
	return optionFunc(func(o *options) { o.query["a1"] = a1 })
}

// WithA2 a2获取
func (obj *_ExcelDataMgr) WithA2(a2 string) Option {
	return optionFunc(func(o *options) { o.query["a2"] = a2 })
}

// WithA3 a3获取
func (obj *_ExcelDataMgr) WithA3(a3 string) Option {
	return optionFunc(func(o *options) { o.query["a3"] = a3 })
}

// WithA4 a4获取
func (obj *_ExcelDataMgr) WithA4(a4 string) Option {
	return optionFunc(func(o *options) { o.query["a4"] = a4 })
}

// WithA5 a5获取
func (obj *_ExcelDataMgr) WithA5(a5 string) Option {
	return optionFunc(func(o *options) { o.query["a5"] = a5 })
}

// WithA6 a6获取
func (obj *_ExcelDataMgr) WithA6(a6 string) Option {
	return optionFunc(func(o *options) { o.query["a6"] = a6 })
}

// WithA7 a7获取
func (obj *_ExcelDataMgr) WithA7(a7 string) Option {
	return optionFunc(func(o *options) { o.query["a7"] = a7 })
}

// WithA8 a8获取
func (obj *_ExcelDataMgr) WithA8(a8 string) Option {
	return optionFunc(func(o *options) { o.query["a8"] = a8 })
}

// WithA9 a9获取
func (obj *_ExcelDataMgr) WithA9(a9 string) Option {
	return optionFunc(func(o *options) { o.query["a9"] = a9 })
}

// WithA10 a10获取
func (obj *_ExcelDataMgr) WithA10(a10 string) Option {
	return optionFunc(func(o *options) { o.query["a10"] = a10 })
}

// WithA11 a11获取
func (obj *_ExcelDataMgr) WithA11(a11 string) Option {
	return optionFunc(func(o *options) { o.query["a11"] = a11 })
}

// WithA12 a12获取
func (obj *_ExcelDataMgr) WithA12(a12 string) Option {
	return optionFunc(func(o *options) { o.query["a12"] = a12 })
}

// WithA13 a13获取
func (obj *_ExcelDataMgr) WithA13(a13 string) Option {
	return optionFunc(func(o *options) { o.query["a13"] = a13 })
}

// WithA14 a14获取
func (obj *_ExcelDataMgr) WithA14(a14 string) Option {
	return optionFunc(func(o *options) { o.query["a14"] = a14 })
}

// WithA15 a15获取
func (obj *_ExcelDataMgr) WithA15(a15 string) Option {
	return optionFunc(func(o *options) { o.query["a15"] = a15 })
}

// WithA16 a16获取
func (obj *_ExcelDataMgr) WithA16(a16 string) Option {
	return optionFunc(func(o *options) { o.query["a16"] = a16 })
}

// WithA17 a17获取
func (obj *_ExcelDataMgr) WithA17(a17 string) Option {
	return optionFunc(func(o *options) { o.query["a17"] = a17 })
}

// WithA18 a18获取
func (obj *_ExcelDataMgr) WithA18(a18 string) Option {
	return optionFunc(func(o *options) { o.query["a18"] = a18 })
}

// WithA19 a19获取
func (obj *_ExcelDataMgr) WithA19(a19 string) Option {
	return optionFunc(func(o *options) { o.query["a19"] = a19 })
}

// WithA20 a20获取
func (obj *_ExcelDataMgr) WithA20(a20 string) Option {
	return optionFunc(func(o *options) { o.query["a20"] = a20 })
}

// GetByOption 功能选项模式获取
func (obj *_ExcelDataMgr) GetByOption(opts ...Option) (result ExcelData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ExcelDataMgr) GetByOptions(opts ...Option) (results []*ExcelData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ExcelDataMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ExcelData, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *_ExcelDataMgr) GetFromID(id int) (result ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_ExcelDataMgr) GetBatchFromID(ids []int) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromA1 通过a1获取内容
func (obj *_ExcelDataMgr) GetFromA1(a1 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a1` = ?", a1).Find(&results).Error

	return
}

// GetBatchFromA1 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA1(a1s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a1` IN (?)", a1s).Find(&results).Error

	return
}

// GetFromA2 通过a2获取内容
func (obj *_ExcelDataMgr) GetFromA2(a2 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a2` = ?", a2).Find(&results).Error

	return
}

// GetBatchFromA2 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA2(a2s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a2` IN (?)", a2s).Find(&results).Error

	return
}

// GetFromA3 通过a3获取内容
func (obj *_ExcelDataMgr) GetFromA3(a3 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a3` = ?", a3).Find(&results).Error

	return
}

// GetBatchFromA3 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA3(a3s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a3` IN (?)", a3s).Find(&results).Error

	return
}

// GetFromA4 通过a4获取内容
func (obj *_ExcelDataMgr) GetFromA4(a4 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a4` = ?", a4).Find(&results).Error

	return
}

// GetBatchFromA4 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA4(a4s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a4` IN (?)", a4s).Find(&results).Error

	return
}

// GetFromA5 通过a5获取内容
func (obj *_ExcelDataMgr) GetFromA5(a5 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a5` = ?", a5).Find(&results).Error

	return
}

// GetBatchFromA5 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA5(a5s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a5` IN (?)", a5s).Find(&results).Error

	return
}

// GetFromA6 通过a6获取内容
func (obj *_ExcelDataMgr) GetFromA6(a6 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a6` = ?", a6).Find(&results).Error

	return
}

// GetBatchFromA6 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA6(a6s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a6` IN (?)", a6s).Find(&results).Error

	return
}

// GetFromA7 通过a7获取内容
func (obj *_ExcelDataMgr) GetFromA7(a7 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a7` = ?", a7).Find(&results).Error

	return
}

// GetBatchFromA7 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA7(a7s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a7` IN (?)", a7s).Find(&results).Error

	return
}

// GetFromA8 通过a8获取内容
func (obj *_ExcelDataMgr) GetFromA8(a8 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a8` = ?", a8).Find(&results).Error

	return
}

// GetBatchFromA8 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA8(a8s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a8` IN (?)", a8s).Find(&results).Error

	return
}

// GetFromA9 通过a9获取内容
func (obj *_ExcelDataMgr) GetFromA9(a9 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a9` = ?", a9).Find(&results).Error

	return
}

// GetBatchFromA9 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA9(a9s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a9` IN (?)", a9s).Find(&results).Error

	return
}

// GetFromA10 通过a10获取内容
func (obj *_ExcelDataMgr) GetFromA10(a10 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a10` = ?", a10).Find(&results).Error

	return
}

// GetBatchFromA10 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA10(a10s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a10` IN (?)", a10s).Find(&results).Error

	return
}

// GetFromA11 通过a11获取内容
func (obj *_ExcelDataMgr) GetFromA11(a11 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a11` = ?", a11).Find(&results).Error

	return
}

// GetBatchFromA11 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA11(a11s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a11` IN (?)", a11s).Find(&results).Error

	return
}

// GetFromA12 通过a12获取内容
func (obj *_ExcelDataMgr) GetFromA12(a12 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a12` = ?", a12).Find(&results).Error

	return
}

// GetBatchFromA12 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA12(a12s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a12` IN (?)", a12s).Find(&results).Error

	return
}

// GetFromA13 通过a13获取内容
func (obj *_ExcelDataMgr) GetFromA13(a13 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a13` = ?", a13).Find(&results).Error

	return
}

// GetBatchFromA13 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA13(a13s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a13` IN (?)", a13s).Find(&results).Error

	return
}

// GetFromA14 通过a14获取内容
func (obj *_ExcelDataMgr) GetFromA14(a14 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a14` = ?", a14).Find(&results).Error

	return
}

// GetBatchFromA14 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA14(a14s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a14` IN (?)", a14s).Find(&results).Error

	return
}

// GetFromA15 通过a15获取内容
func (obj *_ExcelDataMgr) GetFromA15(a15 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a15` = ?", a15).Find(&results).Error

	return
}

// GetBatchFromA15 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA15(a15s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a15` IN (?)", a15s).Find(&results).Error

	return
}

// GetFromA16 通过a16获取内容
func (obj *_ExcelDataMgr) GetFromA16(a16 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a16` = ?", a16).Find(&results).Error

	return
}

// GetBatchFromA16 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA16(a16s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a16` IN (?)", a16s).Find(&results).Error

	return
}

// GetFromA17 通过a17获取内容
func (obj *_ExcelDataMgr) GetFromA17(a17 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a17` = ?", a17).Find(&results).Error

	return
}

// GetBatchFromA17 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA17(a17s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a17` IN (?)", a17s).Find(&results).Error

	return
}

// GetFromA18 通过a18获取内容
func (obj *_ExcelDataMgr) GetFromA18(a18 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a18` = ?", a18).Find(&results).Error

	return
}

// GetBatchFromA18 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA18(a18s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a18` IN (?)", a18s).Find(&results).Error

	return
}

// GetFromA19 通过a19获取内容
func (obj *_ExcelDataMgr) GetFromA19(a19 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a19` = ?", a19).Find(&results).Error

	return
}

// GetBatchFromA19 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA19(a19s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a19` IN (?)", a19s).Find(&results).Error

	return
}

// GetFromA20 通过a20获取内容
func (obj *_ExcelDataMgr) GetFromA20(a20 string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a20` = ?", a20).Find(&results).Error

	return
}

// GetBatchFromA20 批量查找
func (obj *_ExcelDataMgr) GetBatchFromA20(a20s []string) (results []*ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`a20` IN (?)", a20s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ExcelDataMgr) FetchByPrimaryKey(id int) (result ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByExcelDataIDUIndex primary or index 获取唯一内容
func (obj *_ExcelDataMgr) FetchUniqueByExcelDataIDUIndex(id int) (result ExcelData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExcelData{}).Where("`id` = ?", id).First(&result).Error

	return
}
