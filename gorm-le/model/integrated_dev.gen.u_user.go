package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _UUserMgr struct {
	*_BaseMgr
}

// UUserMgr open func
func UUserMgr(db *gorm.DB) *_UUserMgr {
	if db == nil {
		panic(fmt.Errorf("UUserMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UUserMgr{_BaseMgr: &_BaseMgr{DB: db.Table("u_user"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UUserMgr) GetTableName() string {
	return "u_user"
}

// Reset 重置gorm会话
func (obj *_UUserMgr) Reset() *_UUserMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UUserMgr) Get() (result UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_position").Where("name = ?", result.PositionName).Find(&result.UPosition).Error; err != nil { // 职务维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_UUserMgr) Gets() (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(UUser{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_UUserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCode code获取 员工编号
func (obj *_UUserMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithUsername username获取 账户名称
func (obj *_UUserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithEnglishName english_name获取 英文名称
func (obj *_UUserMgr) WithEnglishName(englishName string) Option {
	return optionFunc(func(o *options) { o.query["english_name"] = englishName })
}

// WithPassword password获取 密码/不得明文
func (obj *_UUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithRealName real_name获取 真实姓名
func (obj *_UUserMgr) WithRealName(realName string) Option {
	return optionFunc(func(o *options) { o.query["real_name"] = realName })
}

// WithPhone phone获取 手机/联系方式
func (obj *_UUserMgr) WithPhone(phone string) Option {
	return optionFunc(func(o *options) { o.query["phone"] = phone })
}

// WithRoles roles获取 角色组逗号分割
func (obj *_UUserMgr) WithRoles(roles string) Option {
	return optionFunc(func(o *options) { o.query["roles"] = roles })
}

// WithOrgID org_id获取 组织架构ID
func (obj *_UUserMgr) WithOrgID(orgID int) Option {
	return optionFunc(func(o *options) { o.query["org_id"] = orgID })
}

// WithOrgCode org_code获取 组织机构
func (obj *_UUserMgr) WithOrgCode(orgCode string) Option {
	return optionFunc(func(o *options) { o.query["org_code"] = orgCode })
}

// WithOrgCodes org_codes获取 组织机构父ID
func (obj *_UUserMgr) WithOrgCodes(orgCodes string) Option {
	return optionFunc(func(o *options) { o.query["org_codes"] = orgCodes })
}

// WithPlatformCode platform_code获取 平台代码(绑定平台)
func (obj *_UUserMgr) WithPlatformCode(platformCode string) Option {
	return optionFunc(func(o *options) { o.query["platform_code"] = platformCode })
}

// WithPlatformName platform_name获取 平台名称
func (obj *_UUserMgr) WithPlatformName(platformName string) Option {
	return optionFunc(func(o *options) { o.query["platform_name"] = platformName })
}

// WithHasAudit has_audit获取 是否有审核权限(组织机构) 0:没有权限 (默认) 1:有权限
func (obj *_UUserMgr) WithHasAudit(hasAudit int) Option {
	return optionFunc(func(o *options) { o.query["has_audit"] = hasAudit })
}

// WithUserState user_state获取 用户状态 1:在职 2:离职 3:休假
func (obj *_UUserMgr) WithUserState(userState int) Option {
	return optionFunc(func(o *options) { o.query["user_state"] = userState })
}

// WithAccountState account_state获取 账户状态  0:停用  1:启用
func (obj *_UUserMgr) WithAccountState(accountState int) Option {
	return optionFunc(func(o *options) { o.query["account_state"] = accountState })
}

// WithPermissionType permission_type获取 用户权限类型  1:普通用户 2:个性化  3:未授权用户
func (obj *_UUserMgr) WithPermissionType(permissionType int) Option {
	return optionFunc(func(o *options) { o.query["permission_type"] = permissionType })
}

// WithIDcard idcard获取 身份证
func (obj *_UUserMgr) WithIDcard(idcard string) Option {
	return optionFunc(func(o *options) { o.query["idcard"] = idcard })
}

// WithBirthday birthday获取 生日
func (obj *_UUserMgr) WithBirthday(birthday time.Time) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithSex sex获取 性别 0:男 1:女
func (obj *_UUserMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithWechat wechat获取 微信
func (obj *_UUserMgr) WithWechat(wechat string) Option {
	return optionFunc(func(o *options) { o.query["wechat"] = wechat })
}

// WithQq qq获取 QQ
func (obj *_UUserMgr) WithQq(qq string) Option {
	return optionFunc(func(o *options) { o.query["qq"] = qq })
}

// WithSocialSecurity social_security获取 社保 0:否 1:是
func (obj *_UUserMgr) WithSocialSecurity(socialSecurity int) Option {
	return optionFunc(func(o *options) { o.query["social_security"] = socialSecurity })
}

// WithAccumulationFund accumulation_fund获取 公积金 0:否 1:是
func (obj *_UUserMgr) WithAccumulationFund(accumulationFund int) Option {
	return optionFunc(func(o *options) { o.query["accumulation_fund"] = accumulationFund })
}

// WithEntryTime entry_time获取 入职日期
func (obj *_UUserMgr) WithEntryTime(entryTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["entry_time"] = entryTime })
}

// WithDepartureTime departure_time获取 离职日期
func (obj *_UUserMgr) WithDepartureTime(departureTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["departure_time"] = departureTime })
}

// WithEmail email获取 邮箱
func (obj *_UUserMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithNailing nailing获取 钉钉
func (obj *_UUserMgr) WithNailing(nailing string) Option {
	return optionFunc(func(o *options) { o.query["nailing"] = nailing })
}

// WithWeibo weibo获取 微博
func (obj *_UUserMgr) WithWeibo(weibo string) Option {
	return optionFunc(func(o *options) { o.query["weibo"] = weibo })
}

// WithEducation education获取 学历 0:小学 1:初中 2:高中 3:大专 4:本科 5: 硕士 6:博士
func (obj *_UUserMgr) WithEducation(education int) Option {
	return optionFunc(func(o *options) { o.query["education"] = education })
}

// WithResume resume获取 简历
func (obj *_UUserMgr) WithResume(resume string) Option {
	return optionFunc(func(o *options) { o.query["resume"] = resume })
}

// WithAvatar avatar获取 头像
func (obj *_UUserMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithAddress address获取 住址
func (obj *_UUserMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithPosition position获取 职位
func (obj *_UUserMgr) WithPosition(position string) Option {
	return optionFunc(func(o *options) { o.query["position"] = position })
}

// WithEmergencyContact emergency_contact获取 紧急联系人
func (obj *_UUserMgr) WithEmergencyContact(emergencyContact string) Option {
	return optionFunc(func(o *options) { o.query["emergency_contact"] = emergencyContact })
}

// WithEmergencyContactPhone emergency_contact_phone获取 紧急联系人电话
func (obj *_UUserMgr) WithEmergencyContactPhone(emergencyContactPhone string) Option {
	return optionFunc(func(o *options) { o.query["emergency_contact_phone"] = emergencyContactPhone })
}

// WithPositionName position_name获取 职位名称
func (obj *_UUserMgr) WithPositionName(positionName string) Option {
	return optionFunc(func(o *options) { o.query["position_name"] = positionName })
}

// WithSiteNumber site_number获取 站点编号
func (obj *_UUserMgr) WithSiteNumber(siteNumber string) Option {
	return optionFunc(func(o *options) { o.query["site_number"] = siteNumber })
}

// WithSiteName site_name获取 站点名称
func (obj *_UUserMgr) WithSiteName(siteName string) Option {
	return optionFunc(func(o *options) { o.query["site_name"] = siteName })
}

// WithRemark remark获取 备注
func (obj *_UUserMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithCreateTime create_time获取 创建时间
func (obj *_UUserMgr) WithCreateTime(createTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithCreateUser create_user获取 创建用户
func (obj *_UUserMgr) WithCreateUser(createUser int) Option {
	return optionFunc(func(o *options) { o.query["create_user"] = createUser })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_UUserMgr) WithUpdateTime(updateTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// WithUpdateUser update_user获取 更新用户
func (obj *_UUserMgr) WithUpdateUser(updateUser int) Option {
	return optionFunc(func(o *options) { o.query["update_user"] = updateUser })
}

// WithVersion version获取 乐观锁
func (obj *_UUserMgr) WithVersion(version int) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithDeleted deleted获取 逻辑删除
func (obj *_UUserMgr) WithDeleted(deleted int) Option {
	return optionFunc(func(o *options) { o.query["deleted"] = deleted })
}

// GetByOption 功能选项模式获取
func (obj *_UUserMgr) GetByOption(opts ...Option) (result UUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_position").Where("name = ?", result.PositionName).Find(&result.UPosition).Error; err != nil { // 职务维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UUserMgr) GetByOptions(opts ...Option) (results []*UUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// SelectPage 分页查询
func (obj *_UUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]UUser, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(UUser{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
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
func (obj *_UUserMgr) GetFromID(id int) (result UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_position").Where("name = ?", result.PositionName).Find(&result.UPosition).Error; err != nil { // 职务维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_UUserMgr) GetBatchFromID(ids []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCode 通过code获取内容 员工编号
func (obj *_UUserMgr) GetFromCode(code string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`code` = ?", code).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCode 批量查找 员工编号
func (obj *_UUserMgr) GetBatchFromCode(codes []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`code` IN (?)", codes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUsername 通过username获取内容 账户名称
func (obj *_UUserMgr) GetFromUsername(username string) (result UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`username` = ?", username).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_position").Where("name = ?", result.PositionName).Find(&result.UPosition).Error; err != nil { // 职务维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromUsername 批量查找 账户名称
func (obj *_UUserMgr) GetBatchFromUsername(usernames []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`username` IN (?)", usernames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEnglishName 通过english_name获取内容 英文名称
func (obj *_UUserMgr) GetFromEnglishName(englishName string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`english_name` = ?", englishName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEnglishName 批量查找 英文名称
func (obj *_UUserMgr) GetBatchFromEnglishName(englishNames []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`english_name` IN (?)", englishNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPassword 通过password获取内容 密码/不得明文
func (obj *_UUserMgr) GetFromPassword(password string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`password` = ?", password).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPassword 批量查找 密码/不得明文
func (obj *_UUserMgr) GetBatchFromPassword(passwords []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`password` IN (?)", passwords).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRealName 通过real_name获取内容 真实姓名
func (obj *_UUserMgr) GetFromRealName(realName string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`real_name` = ?", realName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRealName 批量查找 真实姓名
func (obj *_UUserMgr) GetBatchFromRealName(realNames []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`real_name` IN (?)", realNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPhone 通过phone获取内容 手机/联系方式
func (obj *_UUserMgr) GetFromPhone(phone string) (result UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`phone` = ?", phone).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_position").Where("name = ?", result.PositionName).Find(&result.UPosition).Error; err != nil { // 职务维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromPhone 批量查找 手机/联系方式
func (obj *_UUserMgr) GetBatchFromPhone(phones []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`phone` IN (?)", phones).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRoles 通过roles获取内容 角色组逗号分割
func (obj *_UUserMgr) GetFromRoles(roles string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`roles` = ?", roles).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRoles 批量查找 角色组逗号分割
func (obj *_UUserMgr) GetBatchFromRoles(roless []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`roles` IN (?)", roless).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrgID 通过org_id获取内容 组织架构ID
func (obj *_UUserMgr) GetFromOrgID(orgID int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_id` = ?", orgID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrgID 批量查找 组织架构ID
func (obj *_UUserMgr) GetBatchFromOrgID(orgIDs []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_id` IN (?)", orgIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrgCode 通过org_code获取内容 组织机构
func (obj *_UUserMgr) GetFromOrgCode(orgCode string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_code` = ?", orgCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrgCode 批量查找 组织机构
func (obj *_UUserMgr) GetBatchFromOrgCode(orgCodes []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_code` IN (?)", orgCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromOrgCodes 通过org_codes获取内容 组织机构父ID
func (obj *_UUserMgr) GetFromOrgCodes(orgCodes string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_codes` = ?", orgCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromOrgCodes 批量查找 组织机构父ID
func (obj *_UUserMgr) GetBatchFromOrgCodes(orgCodess []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_codes` IN (?)", orgCodess).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPlatformCode 通过platform_code获取内容 平台代码(绑定平台)
func (obj *_UUserMgr) GetFromPlatformCode(platformCode string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`platform_code` = ?", platformCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPlatformCode 批量查找 平台代码(绑定平台)
func (obj *_UUserMgr) GetBatchFromPlatformCode(platformCodes []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`platform_code` IN (?)", platformCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPlatformName 通过platform_name获取内容 平台名称
func (obj *_UUserMgr) GetFromPlatformName(platformName string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`platform_name` = ?", platformName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPlatformName 批量查找 平台名称
func (obj *_UUserMgr) GetBatchFromPlatformName(platformNames []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`platform_name` IN (?)", platformNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromHasAudit 通过has_audit获取内容 是否有审核权限(组织机构) 0:没有权限 (默认) 1:有权限
func (obj *_UUserMgr) GetFromHasAudit(hasAudit int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`has_audit` = ?", hasAudit).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromHasAudit 批量查找 是否有审核权限(组织机构) 0:没有权限 (默认) 1:有权限
func (obj *_UUserMgr) GetBatchFromHasAudit(hasAudits []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`has_audit` IN (?)", hasAudits).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUserState 通过user_state获取内容 用户状态 1:在职 2:离职 3:休假
func (obj *_UUserMgr) GetFromUserState(userState int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`user_state` = ?", userState).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUserState 批量查找 用户状态 1:在职 2:离职 3:休假
func (obj *_UUserMgr) GetBatchFromUserState(userStates []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`user_state` IN (?)", userStates).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAccountState 通过account_state获取内容 账户状态  0:停用  1:启用
func (obj *_UUserMgr) GetFromAccountState(accountState int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`account_state` = ?", accountState).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAccountState 批量查找 账户状态  0:停用  1:启用
func (obj *_UUserMgr) GetBatchFromAccountState(accountStates []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`account_state` IN (?)", accountStates).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPermissionType 通过permission_type获取内容 用户权限类型  1:普通用户 2:个性化  3:未授权用户
func (obj *_UUserMgr) GetFromPermissionType(permissionType int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`permission_type` = ?", permissionType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPermissionType 批量查找 用户权限类型  1:普通用户 2:个性化  3:未授权用户
func (obj *_UUserMgr) GetBatchFromPermissionType(permissionTypes []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`permission_type` IN (?)", permissionTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromIDcard 通过idcard获取内容 身份证
func (obj *_UUserMgr) GetFromIDcard(idcard string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`idcard` = ?", idcard).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromIDcard 批量查找 身份证
func (obj *_UUserMgr) GetBatchFromIDcard(idcards []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`idcard` IN (?)", idcards).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBirthday 通过birthday获取内容 生日
func (obj *_UUserMgr) GetFromBirthday(birthday time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`birthday` = ?", birthday).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBirthday 批量查找 生日
func (obj *_UUserMgr) GetBatchFromBirthday(birthdays []time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`birthday` IN (?)", birthdays).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSex 通过sex获取内容 性别 0:男 1:女
func (obj *_UUserMgr) GetFromSex(sex int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`sex` = ?", sex).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSex 批量查找 性别 0:男 1:女
func (obj *_UUserMgr) GetBatchFromSex(sexs []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`sex` IN (?)", sexs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromWechat 通过wechat获取内容 微信
func (obj *_UUserMgr) GetFromWechat(wechat string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`wechat` = ?", wechat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromWechat 批量查找 微信
func (obj *_UUserMgr) GetBatchFromWechat(wechats []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`wechat` IN (?)", wechats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromQq 通过qq获取内容 QQ
func (obj *_UUserMgr) GetFromQq(qq string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`qq` = ?", qq).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromQq 批量查找 QQ
func (obj *_UUserMgr) GetBatchFromQq(qqs []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`qq` IN (?)", qqs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSocialSecurity 通过social_security获取内容 社保 0:否 1:是
func (obj *_UUserMgr) GetFromSocialSecurity(socialSecurity int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`social_security` = ?", socialSecurity).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSocialSecurity 批量查找 社保 0:否 1:是
func (obj *_UUserMgr) GetBatchFromSocialSecurity(socialSecuritys []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`social_security` IN (?)", socialSecuritys).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAccumulationFund 通过accumulation_fund获取内容 公积金 0:否 1:是
func (obj *_UUserMgr) GetFromAccumulationFund(accumulationFund int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`accumulation_fund` = ?", accumulationFund).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAccumulationFund 批量查找 公积金 0:否 1:是
func (obj *_UUserMgr) GetBatchFromAccumulationFund(accumulationFunds []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`accumulation_fund` IN (?)", accumulationFunds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEntryTime 通过entry_time获取内容 入职日期
func (obj *_UUserMgr) GetFromEntryTime(entryTime time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`entry_time` = ?", entryTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEntryTime 批量查找 入职日期
func (obj *_UUserMgr) GetBatchFromEntryTime(entryTimes []time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`entry_time` IN (?)", entryTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDepartureTime 通过departure_time获取内容 离职日期
func (obj *_UUserMgr) GetFromDepartureTime(departureTime time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`departure_time` = ?", departureTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDepartureTime 批量查找 离职日期
func (obj *_UUserMgr) GetBatchFromDepartureTime(departureTimes []time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`departure_time` IN (?)", departureTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *_UUserMgr) GetFromEmail(email string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`email` = ?", email).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEmail 批量查找 邮箱
func (obj *_UUserMgr) GetBatchFromEmail(emails []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`email` IN (?)", emails).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromNailing 通过nailing获取内容 钉钉
func (obj *_UUserMgr) GetFromNailing(nailing string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`nailing` = ?", nailing).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromNailing 批量查找 钉钉
func (obj *_UUserMgr) GetBatchFromNailing(nailings []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`nailing` IN (?)", nailings).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromWeibo 通过weibo获取内容 微博
func (obj *_UUserMgr) GetFromWeibo(weibo string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`weibo` = ?", weibo).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromWeibo 批量查找 微博
func (obj *_UUserMgr) GetBatchFromWeibo(weibos []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`weibo` IN (?)", weibos).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEducation 通过education获取内容 学历 0:小学 1:初中 2:高中 3:大专 4:本科 5: 硕士 6:博士
func (obj *_UUserMgr) GetFromEducation(education int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`education` = ?", education).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEducation 批量查找 学历 0:小学 1:初中 2:高中 3:大专 4:本科 5: 硕士 6:博士
func (obj *_UUserMgr) GetBatchFromEducation(educations []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`education` IN (?)", educations).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromResume 通过resume获取内容 简历
func (obj *_UUserMgr) GetFromResume(resume string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`resume` = ?", resume).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromResume 批量查找 简历
func (obj *_UUserMgr) GetBatchFromResume(resumes []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`resume` IN (?)", resumes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAvatar 通过avatar获取内容 头像
func (obj *_UUserMgr) GetFromAvatar(avatar string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`avatar` = ?", avatar).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAvatar 批量查找 头像
func (obj *_UUserMgr) GetBatchFromAvatar(avatars []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`avatar` IN (?)", avatars).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAddress 通过address获取内容 住址
func (obj *_UUserMgr) GetFromAddress(address string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`address` = ?", address).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAddress 批量查找 住址
func (obj *_UUserMgr) GetBatchFromAddress(addresss []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`address` IN (?)", addresss).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPosition 通过position获取内容 职位
func (obj *_UUserMgr) GetFromPosition(position string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`position` = ?", position).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPosition 批量查找 职位
func (obj *_UUserMgr) GetBatchFromPosition(positions []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`position` IN (?)", positions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEmergencyContact 通过emergency_contact获取内容 紧急联系人
func (obj *_UUserMgr) GetFromEmergencyContact(emergencyContact string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`emergency_contact` = ?", emergencyContact).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEmergencyContact 批量查找 紧急联系人
func (obj *_UUserMgr) GetBatchFromEmergencyContact(emergencyContacts []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`emergency_contact` IN (?)", emergencyContacts).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEmergencyContactPhone 通过emergency_contact_phone获取内容 紧急联系人电话
func (obj *_UUserMgr) GetFromEmergencyContactPhone(emergencyContactPhone string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`emergency_contact_phone` = ?", emergencyContactPhone).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEmergencyContactPhone 批量查找 紧急联系人电话
func (obj *_UUserMgr) GetBatchFromEmergencyContactPhone(emergencyContactPhones []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`emergency_contact_phone` IN (?)", emergencyContactPhones).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPositionName 通过position_name获取内容 职位名称
func (obj *_UUserMgr) GetFromPositionName(positionName string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`position_name` = ?", positionName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPositionName 批量查找 职位名称
func (obj *_UUserMgr) GetBatchFromPositionName(positionNames []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`position_name` IN (?)", positionNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSiteNumber 通过site_number获取内容 站点编号
func (obj *_UUserMgr) GetFromSiteNumber(siteNumber string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`site_number` = ?", siteNumber).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSiteNumber 批量查找 站点编号
func (obj *_UUserMgr) GetBatchFromSiteNumber(siteNumbers []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`site_number` IN (?)", siteNumbers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromSiteName 通过site_name获取内容 站点名称
func (obj *_UUserMgr) GetFromSiteName(siteName string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`site_name` = ?", siteName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromSiteName 批量查找 站点名称
func (obj *_UUserMgr) GetBatchFromSiteName(siteNames []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`site_name` IN (?)", siteNames).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *_UUserMgr) GetFromRemark(remark string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`remark` = ?", remark).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *_UUserMgr) GetBatchFromRemark(remarks []string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`remark` IN (?)", remarks).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_UUserMgr) GetFromCreateTime(createTime time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`create_time` = ?", createTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_UUserMgr) GetBatchFromCreateTime(createTimes []time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromCreateUser 通过create_user获取内容 创建用户
func (obj *_UUserMgr) GetFromCreateUser(createUser int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`create_user` = ?", createUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromCreateUser 批量查找 创建用户
func (obj *_UUserMgr) GetBatchFromCreateUser(createUsers []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`create_user` IN (?)", createUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_UUserMgr) GetFromUpdateTime(updateTime time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`update_time` = ?", updateTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_UUserMgr) GetBatchFromUpdateTime(updateTimes []time.Time) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromUpdateUser 通过update_user获取内容 更新用户
func (obj *_UUserMgr) GetFromUpdateUser(updateUser int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`update_user` = ?", updateUser).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromUpdateUser 批量查找 更新用户
func (obj *_UUserMgr) GetBatchFromUpdateUser(updateUsers []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`update_user` IN (?)", updateUsers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromVersion 通过version获取内容 乐观锁
func (obj *_UUserMgr) GetFromVersion(version int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`version` = ?", version).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromVersion 批量查找 乐观锁
func (obj *_UUserMgr) GetBatchFromVersion(versions []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`version` IN (?)", versions).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromDeleted 通过deleted获取内容 逻辑删除
func (obj *_UUserMgr) GetFromDeleted(deleted int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`deleted` = ?", deleted).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromDeleted 批量查找 逻辑删除
func (obj *_UUserMgr) GetBatchFromDeleted(deleteds []int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`deleted` IN (?)", deleteds).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
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
func (obj *_UUserMgr) FetchByPrimaryKey(id int) (result UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`id` = ?", id).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_position").Where("name = ?", result.PositionName).Find(&result.UPosition).Error; err != nil { // 职务维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByIndexUsername primary or index 获取唯一内容
func (obj *_UUserMgr) FetchUniqueByIndexUsername(username string) (result UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`username` = ?", username).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_position").Where("name = ?", result.PositionName).Find(&result.UPosition).Error; err != nil { // 职务维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueByIndexPhone primary or index 获取唯一内容
func (obj *_UUserMgr) FetchUniqueByIndexPhone(phone string) (result UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`phone` = ?", phone).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("u_position").Where("name = ?", result.PositionName).Find(&result.UPosition).Error; err != nil { // 职务维护
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByUUserOrgIDIndex  获取多个内容
func (obj *_UUserMgr) FetchIndexByUUserOrgIDIndex(orgID int) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_id` = ?", orgID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByUUserOrgCodeIndex  获取多个内容
func (obj *_UUserMgr) FetchIndexByUUserOrgCodeIndex(orgCode string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_code` = ?", orgCode).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByUUserOrgCodesIndex  获取多个内容
func (obj *_UUserMgr) FetchIndexByUUserOrgCodesIndex(orgCodes string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`org_codes` = ?", orgCodes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByUUserUPositionNameFk  获取多个内容
func (obj *_UUserMgr) FetchIndexByUUserUPositionNameFk(positionName string) (results []*UUser, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(UUser{}).Where("`position_name` = ?", positionName).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("u_position").Where("name = ?", results[i].PositionName).Find(&results[i].UPosition).Error; err != nil { // 职务维护
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
