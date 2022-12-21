package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _WmsSorterGridRuleMgr struct {
	*_BaseMgr
}

// WmsSorterGridRuleMgr open func
func WmsSorterGridRuleMgr(db *gorm.DB) *_WmsSorterGridRuleMgr {
	if db == nil {
		panic(fmt.Errorf("WmsSorterGridRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_WmsSorterGridRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("wms_sorter_grid_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_WmsSorterGridRuleMgr) GetTableName() string {
	return "wms_sorter_grid_rule"
}

// Reset 重置gorm会话
func (obj *_WmsSorterGridRuleMgr) Reset() *_WmsSorterGridRuleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_WmsSorterGridRuleMgr) Get() (result WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_WmsSorterGridRuleMgr) Gets() (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_WmsSorterGridRuleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_WmsSorterGridRuleMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithRuleName rule_name获取 规则名称
func (obj *_WmsSorterGridRuleMgr) WithRuleName(ruleName string) Option {
	return optionFunc(func(o *options) { o.query["rule_name"] = ruleName })
}

// WithGridNums grid_nums获取 格口，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) WithGridNums(gridNums string) Option {
	return optionFunc(func(o *options) { o.query["grid_nums"] = gridNums })
}

// WithSubRuleType sub_rule_type获取 子规则类型，weight:重量，country:国家，customer:客户，no_sub_rule:不存在子规则，not_found_order:未找到订单，problem_order:问题订单
func (obj *_WmsSorterGridRuleMgr) WithSubRuleType(subRuleType string) Option {
	return optionFunc(func(o *options) { o.query["sub_rule_type"] = subRuleType })
}

// WithSpecialRuleFlag special_rule_flag获取 是否特殊规则，0:否，1:是
func (obj *_WmsSorterGridRuleMgr) WithSpecialRuleFlag(specialRuleFlag []uint8) Option {
	return optionFunc(func(o *options) { o.query["special_rule_flag"] = specialRuleFlag })
}

// WithCustomerChannelIDs customer_channel_ids获取 客户渠道id，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) WithCustomerChannelIDs(customerChannelIDs string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_ids"] = customerChannelIDs })
}

// WithCustomerChannelNames customer_channel_names获取 客户渠道名称，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) WithCustomerChannelNames(customerChannelNames string) Option {
	return optionFunc(func(o *options) { o.query["customer_channel_names"] = customerChannelNames })
}

// WithDownstreamCode downstream_code获取 下家code
func (obj *_WmsSorterGridRuleMgr) WithDownstreamCode(downstreamCode string) Option {
	return optionFunc(func(o *options) { o.query["downstream_code"] = downstreamCode })
}

// WithOverFlag over_flag获取 是否终结
func (obj *_WmsSorterGridRuleMgr) WithOverFlag(overFlag bool) Option {
	return optionFunc(func(o *options) { o.query["over_flag"] = overFlag })
}

// WithRemark remark获取 备注
func (obj *_WmsSorterGridRuleMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_WmsSorterGridRuleMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_WmsSorterGridRuleMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_WmsSorterGridRuleMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_WmsSorterGridRuleMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_WmsSorterGridRuleMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_WmsSorterGridRuleMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_WmsSorterGridRuleMgr) GetByOption(opts ...Option) (result WmsSorterGridRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_WmsSorterGridRuleMgr) GetByOptions(opts ...Option) (results []*WmsSorterGridRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_WmsSorterGridRuleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]WmsSorterGridRule, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_WmsSorterGridRuleMgr) GetFromID(id int) (result WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_WmsSorterGridRuleMgr) GetBatchFromID(ids []int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromRuleName 通过rule_name获取内容 规则名称
func (obj *_WmsSorterGridRuleMgr) GetFromRuleName(ruleName string) (result WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`rule_name` = ?", ruleName).First(&result).Error

	return
}

// GetBatchFromRuleName 批量查找 规则名称
func (obj *_WmsSorterGridRuleMgr) GetBatchFromRuleName(ruleNames []string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`rule_name` IN (?)", ruleNames).Find(&results).Error

	return
}

// GetFromGridNums 通过grid_nums获取内容 格口，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) GetFromGridNums(gridNums string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`grid_nums` = ?", gridNums).Find(&results).Error

	return
}

// GetBatchFromGridNums 批量查找 格口，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) GetBatchFromGridNums(gridNumss []string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`grid_nums` IN (?)", gridNumss).Find(&results).Error

	return
}

// GetFromSubRuleType 通过sub_rule_type获取内容 子规则类型，weight:重量，country:国家，customer:客户，no_sub_rule:不存在子规则，not_found_order:未找到订单，problem_order:问题订单
func (obj *_WmsSorterGridRuleMgr) GetFromSubRuleType(subRuleType string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`sub_rule_type` = ?", subRuleType).Find(&results).Error

	return
}

// GetBatchFromSubRuleType 批量查找 子规则类型，weight:重量，country:国家，customer:客户，no_sub_rule:不存在子规则，not_found_order:未找到订单，problem_order:问题订单
func (obj *_WmsSorterGridRuleMgr) GetBatchFromSubRuleType(subRuleTypes []string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`sub_rule_type` IN (?)", subRuleTypes).Find(&results).Error

	return
}

// GetFromSpecialRuleFlag 通过special_rule_flag获取内容 是否特殊规则，0:否，1:是
func (obj *_WmsSorterGridRuleMgr) GetFromSpecialRuleFlag(specialRuleFlag []uint8) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`special_rule_flag` = ?", specialRuleFlag).Find(&results).Error

	return
}

// GetBatchFromSpecialRuleFlag 批量查找 是否特殊规则，0:否，1:是
func (obj *_WmsSorterGridRuleMgr) GetBatchFromSpecialRuleFlag(specialRuleFlags [][]uint8) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`special_rule_flag` IN (?)", specialRuleFlags).Find(&results).Error

	return
}

// GetFromCustomerChannelIDs 通过customer_channel_ids获取内容 客户渠道id，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) GetFromCustomerChannelIDs(customerChannelIDs string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`customer_channel_ids` = ?", customerChannelIDs).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelIDs 批量查找 客户渠道id，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) GetBatchFromCustomerChannelIDs(customerChannelIDss []string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`customer_channel_ids` IN (?)", customerChannelIDss).Find(&results).Error

	return
}

// GetFromCustomerChannelNames 通过customer_channel_names获取内容 客户渠道名称，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) GetFromCustomerChannelNames(customerChannelNames string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`customer_channel_names` = ?", customerChannelNames).Find(&results).Error

	return
}

// GetBatchFromCustomerChannelNames 批量查找 客户渠道名称，多个以逗号隔开
func (obj *_WmsSorterGridRuleMgr) GetBatchFromCustomerChannelNames(customerChannelNamess []string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`customer_channel_names` IN (?)", customerChannelNamess).Find(&results).Error

	return
}

// GetFromDownstreamCode 通过downstream_code获取内容 下家code
func (obj *_WmsSorterGridRuleMgr) GetFromDownstreamCode(downstreamCode string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`downstream_code` = ?", downstreamCode).Find(&results).Error

	return
}

// GetBatchFromDownstreamCode 批量查找 下家code
func (obj *_WmsSorterGridRuleMgr) GetBatchFromDownstreamCode(downstreamCodes []string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`downstream_code` IN (?)", downstreamCodes).Find(&results).Error

	return
}

// GetFromOverFlag 通过over_flag获取内容 是否终结
func (obj *_WmsSorterGridRuleMgr) GetFromOverFlag(overFlag bool) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`over_flag` = ?", overFlag).Find(&results).Error

	return
}

// GetBatchFromOverFlag 批量查找 是否终结
func (obj *_WmsSorterGridRuleMgr) GetBatchFromOverFlag(overFlags []bool) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`over_flag` IN (?)", overFlags).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_WmsSorterGridRuleMgr) GetFromRemark(remark string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_WmsSorterGridRuleMgr) GetBatchFromRemark(remarks []string) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_WmsSorterGridRuleMgr) GetFromCreateTime(createTime time.Time) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_WmsSorterGridRuleMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_WmsSorterGridRuleMgr) GetFromCreateUser(createUser int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`create_user` = ?", createUser).Find(&results).Error

	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_WmsSorterGridRuleMgr) GetBatchFromCreateUser(createUsers []int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_WmsSorterGridRuleMgr) GetFromUpdateTime(updateTime time.Time) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_WmsSorterGridRuleMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_WmsSorterGridRuleMgr) GetFromUpdateUser(updateUser int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`update_user` = ?", updateUser).Find(&results).Error

	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_WmsSorterGridRuleMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_WmsSorterGridRuleMgr) GetFromVersion(version int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_WmsSorterGridRuleMgr) GetBatchFromVersion(versions []int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_WmsSorterGridRuleMgr) GetFromDeleted(deleted int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`deleted` = ?", deleted).Find(&results).Error

	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_WmsSorterGridRuleMgr) GetBatchFromDeleted(deleteds []int) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_WmsSorterGridRuleMgr) FetchByPrimaryKey(id int) (result WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchUniqueByRuleName primary or index 获取唯一内容
func (obj *_WmsSorterGridRuleMgr) FetchUniqueByRuleName(ruleName string) (result WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`rule_name` = ?", ruleName).First(&result).Error

	return
}

// FetchIndexBySpecialRuleFlagIndex  获取多个内容
func (obj *_WmsSorterGridRuleMgr) FetchIndexBySpecialRuleFlagIndex(specialRuleFlag []uint8) (results []*WmsSorterGridRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(WmsSorterGridRule{}).Where("`special_rule_flag` = ?", specialRuleFlag).Find(&results).Error

	return
}
