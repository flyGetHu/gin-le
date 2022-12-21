package model

import (
	"gorm.io/datatypes"
	"time"
)

// AAttendance 考勤记录
type AAttendance struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`
	UserID           int       `gorm:"column:user_id" json:"userId"`                      // 用户ID
	OrgID            int       `gorm:"column:org_id" json:"orgId"`                        // 部门ID
	WorkDate         string    `gorm:"column:work_date" json:"workDate"`                  // 日期
	ShiftName        string    `gorm:"column:shift_name" json:"shiftName"`                // 考勤班次名称
	UpWorkDatetime   time.Time `gorm:"column:up_work_datetime" json:"upWorkDatetime"`     // 上班时间
	DownWorkDatatime time.Time `gorm:"column:down_work_datatime" json:"downWorkDatatime"` // 下班时间
	UpWorkLate       int       `gorm:"column:up_work_late" json:"upWorkLate"`             // 迟到时间
	DownWorkEarly    int       `gorm:"column:down_work_early" json:"downWorkEarly"`       // 早退时间
	Remark           string    `gorm:"column:remark" json:"remark"`                       // 备注
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`              // 创建时间
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`              // 创建用户
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`              // 更新时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`              // 更新时间
	Version          int       `gorm:"column:version" json:"version"`                     // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                     // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *AAttendance) TableName() string {
	return "a_attendance"
}

// AAttendanceColumns get sql column name.获取数据库列名
var AAttendanceColumns = struct {
	ID               string
	UserID           string
	OrgID            string
	WorkDate         string
	ShiftName        string
	UpWorkDatetime   string
	DownWorkDatatime string
	UpWorkLate       string
	DownWorkEarly    string
	Remark           string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	UserID:           "user_id",
	OrgID:            "org_id",
	WorkDate:         "work_date",
	ShiftName:        "shift_name",
	UpWorkDatetime:   "up_work_datetime",
	DownWorkDatatime: "down_work_datatime",
	UpWorkLate:       "up_work_late",
	DownWorkEarly:    "down_work_early",
	Remark:           "remark",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
}

// AAttendanceLocation 打卡地点
type AAttendanceLocation struct {
	OrgID      int       `gorm:"primaryKey;column:org_id" json:"-"`
	Longitude  float64   `gorm:"column:longitude" json:"longitude"`    // 经度
	Dimension  float64   `gorm:"column:dimension" json:"dimension"`    // 纬度
	Radius     int       `gorm:"column:radius" json:"radius"`          // 打卡半径
	Location   string    `gorm:"column:location" json:"location"`      // 位置
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *AAttendanceLocation) TableName() string {
	return "a_attendance_location"
}

// AAttendanceLocationColumns get sql column name.获取数据库列名
var AAttendanceLocationColumns = struct {
	OrgID      string
	Longitude  string
	Dimension  string
	Radius     string
	Location   string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	OrgID:      "org_id",
	Longitude:  "longitude",
	Dimension:  "dimension",
	Radius:     "radius",
	Location:   "location",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// AShift 考勤-班次
type AShift struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`
	ShiftName        string    `gorm:"column:shift_name" json:"shiftName"`                // 班次名称
	UpWorkTime       string    `gorm:"column:up_work_time" json:"upWorkTime"`             // 上班打卡
	DownWorkTime     string    `gorm:"column:down_work_time" json:"downWorkTime"`         // 下班时间
	UpWorkEarliest   int       `gorm:"column:up_work_earliest" json:"upWorkEarliest"`     // 上班最早打卡时间
	DownWorkEarliest int       `gorm:"column:down_work_earliest" json:"downWorkEarliest"` // 下班最早打卡时间
	UpWorkLatest     int       `gorm:"column:up_work_latest" json:"upWorkLatest"`         // 上班最晚打卡时间
	DownWorkLatest   int       `gorm:"column:down_work_latest" json:"downWorkLatest"`     // 下班最晚打卡时间
	IsPhoto          int       `gorm:"column:is_photo" json:"isPhoto"`                    // 是否拍照
	IsLockRange      int       `gorm:"column:is_lock_range" json:"isLockRange"`           // 是否锁定打卡范围
	IsRest           int       `gorm:"column:is_rest" json:"isRest"`                      // 是否休息/考勤开关
	Remark           string    `gorm:"column:remark" json:"remark"`                       // 备注
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`              // 创建时间
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`              // 创建用户
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`              // 更新时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`              // 更新时间
	Version          int       `gorm:"column:version" json:"version"`                     // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                     // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *AShift) TableName() string {
	return "a_shift"
}

// AShiftColumns get sql column name.获取数据库列名
var AShiftColumns = struct {
	ID               string
	ShiftName        string
	UpWorkTime       string
	DownWorkTime     string
	UpWorkEarliest   string
	DownWorkEarliest string
	UpWorkLatest     string
	DownWorkLatest   string
	IsPhoto          string
	IsLockRange      string
	IsRest           string
	Remark           string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	ShiftName:        "shift_name",
	UpWorkTime:       "up_work_time",
	DownWorkTime:     "down_work_time",
	UpWorkEarliest:   "up_work_earliest",
	DownWorkEarliest: "down_work_earliest",
	UpWorkLatest:     "up_work_latest",
	DownWorkLatest:   "down_work_latest",
	IsPhoto:          "is_photo",
	IsLockRange:      "is_lock_range",
	IsRest:           "is_rest",
	Remark:           "remark",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
}

// AUserArrangeWork 员工排班表
type AUserArrangeWork struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`
	UserID       int       `gorm:"column:user_id" json:"userId"`              // 用户ID
	PlanMonth    string    `gorm:"column:plan_month" json:"planMonth"`        // 排班月份
	ShiftNames   string    `gorm:"column:shift_names" json:"shiftNames"`      // 班次名称字符串集
	ShiftNameIDs string    `gorm:"column:shift_name_ids" json:"shiftNameIds"` // 班次名称ID字符串集
	Remark       string    `gorm:"column:remark" json:"remark"`               // 备注
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`      // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`      // 创建用户
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`      // 更新时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`      // 更新时间
	Version      int       `gorm:"column:version" json:"version"`             // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`             // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *AUserArrangeWork) TableName() string {
	return "a_user_arrange_work"
}

// AUserArrangeWorkColumns get sql column name.获取数据库列名
var AUserArrangeWorkColumns = struct {
	ID           string
	UserID       string
	PlanMonth    string
	ShiftNames   string
	ShiftNameIDs string
	Remark       string
	CreateTime   string
	CreateUser   string
	UpdateTime   string
	UpdateUser   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	UserID:       "user_id",
	PlanMonth:    "plan_month",
	ShiftNames:   "shift_names",
	ShiftNameIDs: "shift_name_ids",
	Remark:       "remark",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateTime:   "update_time",
	UpdateUser:   "update_user",
	Version:      "version",
	Deleted:      "deleted",
}

// CConsumeLevel 客户消费等级
type CConsumeLevel struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`
	Level     int    `gorm:"column:level" json:"level"`          // 消费等级
	LevelDesc string `gorm:"column:level_desc" json:"levelDesc"` // 等级描述
}

// TableName get sql table name.获取数据库表名
func (m *CConsumeLevel) TableName() string {
	return "c_consume_level"
}

// CConsumeLevelColumns get sql column name.获取数据库列名
var CConsumeLevelColumns = struct {
	ID        string
	Level     string
	LevelDesc string
}{
	ID:        "id",
	Level:     "level",
	LevelDesc: "level_desc",
}

// CConsumeRule 客户消费规则
type CConsumeRule struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`
	RuleType        int       `gorm:"column:rule_type" json:"ruleType"`               // 规则类型，0:累计金额达标，1:周期内消费总金额达标
	RuleDesc        string    `gorm:"column:rule_desc" json:"ruleDesc"`               // 规则描述
	CalculatePeriod int       `gorm:"column:calculate_period" json:"calculatePeriod"` // 计费周期(单位:天)
	RuleAmount      float64   `gorm:"column:rule_amount" json:"ruleAmount"`           // 规则达标所需金额
	StartTime       time.Time `gorm:"column:start_time" json:"startTime"`             // 规则开始时间
	EndTime         time.Time `gorm:"column:end_time" json:"endTime"`                 // 规则结束时间
}

// TableName get sql table name.获取数据库表名
func (m *CConsumeRule) TableName() string {
	return "c_consume_rule"
}

// CConsumeRuleColumns get sql column name.获取数据库列名
var CConsumeRuleColumns = struct {
	ID              string
	RuleType        string
	RuleDesc        string
	CalculatePeriod string
	RuleAmount      string
	StartTime       string
	EndTime         string
}{
	ID:              "id",
	RuleType:        "rule_type",
	RuleDesc:        "rule_desc",
	CalculatePeriod: "calculate_period",
	RuleAmount:      "rule_amount",
	StartTime:       "start_time",
	EndTime:         "end_time",
}

// CCustomer 客户
type CCustomer struct {
	ID                     int       `gorm:"primaryKey;column:id" json:"-"`
	Code                   string    `gorm:"column:code" json:"code"`                                       // 客户编号
	Username               string    `gorm:"column:username" json:"username"`                               // 账户名称
	Password               string    `gorm:"column:password" json:"password"`                               // 密码/不得明文
	CustomerAlias          string    `gorm:"column:customer_alias" json:"customerAlias"`                    // 客户别名
	Level                  int       `gorm:"column:level" json:"level"`                                     // 等级，2:一级账户，3:一级账户下面的级别
	ParentID               int       `gorm:"column:parent_id" json:"parentId"`                              // 上级id(客户公司的上下级关系)
	Integral               int       `gorm:"column:integral" json:"integral"`                               // 积分
	Source                 string    `gorm:"column:source" json:"source"`                                   // 客户来源，corporation:企业，person:个人
	Company                string    `gorm:"column:company" json:"company"`                                 // 公司名称
	Address                string    `gorm:"column:address" json:"address"`                                 // 公司(个人)地址
	IDentityCard           string    `gorm:"column:identity_card" json:"identityCard"`                      // 营业执照或身份证号
	IDentityCardAttachment string    `gorm:"column:identity_card_attachment" json:"identityCardAttachment"` // 证件照附件
	CoPrincipalName        string    `gorm:"column:co_principal_name" json:"coPrincipalName"`               // 企业负责人名字
	CoPrincipalTel         string    `gorm:"column:co_principal_tel" json:"coPrincipalTel"`                 // 企业负责人电话
	CoPrincipalPhone       string    `gorm:"column:co_principal_phone" json:"coPrincipalPhone"`             // 企业负责人手机
	CoPrincipalQq          string    `gorm:"column:co_principal_qq" json:"coPrincipalQq"`                   // 企业负责人qq
	CoPrincipalEmail       string    `gorm:"column:co_principal_email" json:"coPrincipalEmail"`             // 企业负责人邮箱
	LogisticsContactsName  string    `gorm:"column:logistics_contacts_name" json:"logisticsContactsName"`   // 物流联系人名字
	LogisticsContactsTel   string    `gorm:"column:logistics_contacts_tel" json:"logisticsContactsTel"`     // 物流联系人电话
	LogisticsContactsPhone string    `gorm:"column:logistics_contacts_phone" json:"logisticsContactsPhone"` // 物流联系人手机
	LogisticsContactsQq    string    `gorm:"column:logistics_contacts_qq" json:"logisticsContactsQq"`       // 物流联系人qq
	LogisticsContactsEmail string    `gorm:"column:logistics_contacts_email" json:"logisticsContactsEmail"` // 物流联系人邮箱
	FinanceContactsName    string    `gorm:"column:finance_contacts_name" json:"financeContactsName"`       // 财务联系人名字
	FinanceContactsTel     string    `gorm:"column:finance_contacts_tel" json:"financeContactsTel"`         // 财务联系人电话
	FinanceContactsPhone   string    `gorm:"column:finance_contacts_phone" json:"financeContactsPhone"`     // 财务联系人手机
	FinanceContactsQq      string    `gorm:"column:finance_contacts_qq" json:"financeContactsQq"`           // 财务联系人qq
	FinanceContactsEmail   string    `gorm:"column:finance_contacts_email" json:"financeContactsEmail"`     // 财务联系人邮箱
	CustomServiceUserID    int       `gorm:"column:custom_service_user_id" json:"customServiceUserId"`      // 客服id(oa系统中的用户id)
	LockState              int       `gorm:"column:lock_state" json:"lockState"`                            // 账号锁定，0：正常，1：锁定
	AuditState             int       `gorm:"column:audit_state" json:"auditState"`                          // 账户审核，0：未审核，1：已审核
	AccountStatus          int       `gorm:"column:account_status" json:"accountStatus"`                    // 账户状态 0：正常，1：欠费，2：禁用
	Balance                float64   `gorm:"column:balance" json:"balance"`                                 // 账户余额(账面余额)
	InitialCreditLimit     float64   `gorm:"column:initial_credit_limit" json:"initialCreditLimit"`         // 信用额度
	TempCreditLimit        float64   `gorm:"column:temp_credit_limit" json:"tempCreditLimit"`               // 临时信用额度
	TotalBalance           float64   `gorm:"column:total_balance" json:"totalBalance"`                      // 总计余额,总计余额是把未出货的那些一起预扣了
	CreditBalance          float64   `gorm:"column:credit_balance" json:"creditBalance"`                    // 信用余额
	TempCreditExpireTime   time.Time `gorm:"column:temp_credit_expire_time" json:"tempCreditExpireTime"`    // 临时信用额度过期时间
	OutCumulativeFee       float64   `gorm:"column:out_cumulative_fee" json:"outCumulativeFee"`             // 出库累加费用,Note:距离 上次结算 到 当前时间 累计出库余额 变动
	InputCumulativeFee     float64   `gorm:"column:input_cumulative_fee" json:"inputCumulativeFee"`         // 入库累计费用
	LastCollectionDate     time.Time `gorm:"column:last_collection_date" json:"lastCollectionDate"`         // 最后收款时间
	PaymentCycle           int       `gorm:"column:payment_cycle" json:"paymentCycle"`                      // 收款周期
	CreateTime             time.Time `gorm:"column:create_time" json:"createTime"`                          // 创建时间
	CreateUser             int       `gorm:"column:create_user" json:"createUser"`                          // 创建用户
	UpdateTime             time.Time `gorm:"column:update_time" json:"updateTime"`                          // 更新时间
	UpdateUser             int       `gorm:"column:update_user" json:"updateUser"`                          // 更新时间
	Version                int       `gorm:"column:version" json:"version"`                                 // 乐观锁
	Deleted                int       `gorm:"column:deleted" json:"deleted"`                                 // 逻辑删除
	AuthorizationCode      string    `gorm:"column:authorization_code" json:"authorizationCode"`            // 授权码
	HandlerName            string    `gorm:"column:handler_name" json:"handlerName"`                        // 经办人名称
	HandlerID              int       `gorm:"column:handler_id" json:"handlerId"`                            // 经办人id
	SalesLink              string    `gorm:"column:sales_link" json:"salesLink"`                            // 销售链接
	AttributeState         int       `gorm:"column:attribute_state" json:"attributeState"`                  // 长宽高是否必填
	PlatformCode           string    `gorm:"column:platform_code" json:"platformCode"`                      // 平台代码
	OrganizationCode       string    `gorm:"column:organization_code" json:"organizationCode"`              // 客户组织编码
	OrganizationName       string    `gorm:"column:organization_name" json:"organizationName"`              // 客户组织名称
	KingdeeID              string    `gorm:"column:kingdee_id" json:"kingdeeId"`                            // 金蝶客户内码
}

// TableName get sql table name.获取数据库表名
func (m *CCustomer) TableName() string {
	return "c_customer"
}

// CCustomerColumns get sql column name.获取数据库列名
var CCustomerColumns = struct {
	ID                     string
	Code                   string
	Username               string
	Password               string
	CustomerAlias          string
	Level                  string
	ParentID               string
	Integral               string
	Source                 string
	Company                string
	Address                string
	IDentityCard           string
	IDentityCardAttachment string
	CoPrincipalName        string
	CoPrincipalTel         string
	CoPrincipalPhone       string
	CoPrincipalQq          string
	CoPrincipalEmail       string
	LogisticsContactsName  string
	LogisticsContactsTel   string
	LogisticsContactsPhone string
	LogisticsContactsQq    string
	LogisticsContactsEmail string
	FinanceContactsName    string
	FinanceContactsTel     string
	FinanceContactsPhone   string
	FinanceContactsQq      string
	FinanceContactsEmail   string
	CustomServiceUserID    string
	LockState              string
	AuditState             string
	AccountStatus          string
	Balance                string
	InitialCreditLimit     string
	TempCreditLimit        string
	TotalBalance           string
	CreditBalance          string
	TempCreditExpireTime   string
	OutCumulativeFee       string
	InputCumulativeFee     string
	LastCollectionDate     string
	PaymentCycle           string
	CreateTime             string
	CreateUser             string
	UpdateTime             string
	UpdateUser             string
	Version                string
	Deleted                string
	AuthorizationCode      string
	HandlerName            string
	HandlerID              string
	SalesLink              string
	AttributeState         string
	PlatformCode           string
	OrganizationCode       string
	OrganizationName       string
	KingdeeID              string
}{
	ID:                     "id",
	Code:                   "code",
	Username:               "username",
	Password:               "password",
	CustomerAlias:          "customer_alias",
	Level:                  "level",
	ParentID:               "parent_id",
	Integral:               "integral",
	Source:                 "source",
	Company:                "company",
	Address:                "address",
	IDentityCard:           "identity_card",
	IDentityCardAttachment: "identity_card_attachment",
	CoPrincipalName:        "co_principal_name",
	CoPrincipalTel:         "co_principal_tel",
	CoPrincipalPhone:       "co_principal_phone",
	CoPrincipalQq:          "co_principal_qq",
	CoPrincipalEmail:       "co_principal_email",
	LogisticsContactsName:  "logistics_contacts_name",
	LogisticsContactsTel:   "logistics_contacts_tel",
	LogisticsContactsPhone: "logistics_contacts_phone",
	LogisticsContactsQq:    "logistics_contacts_qq",
	LogisticsContactsEmail: "logistics_contacts_email",
	FinanceContactsName:    "finance_contacts_name",
	FinanceContactsTel:     "finance_contacts_tel",
	FinanceContactsPhone:   "finance_contacts_phone",
	FinanceContactsQq:      "finance_contacts_qq",
	FinanceContactsEmail:   "finance_contacts_email",
	CustomServiceUserID:    "custom_service_user_id",
	LockState:              "lock_state",
	AuditState:             "audit_state",
	AccountStatus:          "account_status",
	Balance:                "balance",
	InitialCreditLimit:     "initial_credit_limit",
	TempCreditLimit:        "temp_credit_limit",
	TotalBalance:           "total_balance",
	CreditBalance:          "credit_balance",
	TempCreditExpireTime:   "temp_credit_expire_time",
	OutCumulativeFee:       "out_cumulative_fee",
	InputCumulativeFee:     "input_cumulative_fee",
	LastCollectionDate:     "last_collection_date",
	PaymentCycle:           "payment_cycle",
	CreateTime:             "create_time",
	CreateUser:             "create_user",
	UpdateTime:             "update_time",
	UpdateUser:             "update_user",
	Version:                "version",
	Deleted:                "deleted",
	AuthorizationCode:      "authorization_code",
	HandlerName:            "handler_name",
	HandlerID:              "handler_id",
	SalesLink:              "sales_link",
	AttributeState:         "attribute_state",
	PlatformCode:           "platform_code",
	OrganizationCode:       "organization_code",
	OrganizationName:       "organization_name",
	KingdeeID:              "kingdee_id",
}

// CCustomerBillFlow 客户账单流水
type CCustomerBillFlow struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`                  // 主键id
	CustomerID      int       `gorm:"column:customer_id" json:"customerId"`           // 客户id
	CustomerName    string    `gorm:"column:customer_name" json:"customerName"`       // 客户名称
	BillType        string    `gorm:"column:bill_type" json:"billType"`               // 账单类型
	DeductionAmount float64   `gorm:"column:deduction_amount" json:"deductionAmount"` // 操作金额
	BeforeBalance   float64   `gorm:"column:before_balance" json:"beforeBalance"`     // 操作前余额
	AfterBalance    float64   `gorm:"column:after_balance" json:"afterBalance"`       // 操作后余额
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`           // 创建用户
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	Version         int       `gorm:"column:version" json:"version"`                  // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                  // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *CCustomerBillFlow) TableName() string {
	return "c_customer_bill_flow"
}

// CCustomerBillFlowColumns get sql column name.获取数据库列名
var CCustomerBillFlowColumns = struct {
	ID              string
	CustomerID      string
	CustomerName    string
	BillType        string
	DeductionAmount string
	BeforeBalance   string
	AfterBalance    string
	CreateUser      string
	CreateTime      string
	Version         string
	Deleted         string
}{
	ID:              "id",
	CustomerID:      "customer_id",
	CustomerName:    "customer_name",
	BillType:        "bill_type",
	DeductionAmount: "deduction_amount",
	BeforeBalance:   "before_balance",
	AfterBalance:    "after_balance",
	CreateUser:      "create_user",
	CreateTime:      "create_time",
	Version:         "version",
	Deleted:         "deleted",
}

// CCustomerChannelConfig 客户渠道配置
type CCustomerChannelConfig struct {
	ID                          int       `gorm:"primaryKey;column:id" json:"-"`
	CustomerID                  int       `gorm:"column:customer_id" json:"customerId"`                                     // 客户id
	CustomerChannelID           int       `gorm:"column:customer_channel_id" json:"customerChannelId"`                      // 客户渠道id
	OrderCountMax               int       `gorm:"column:order_count_max" json:"orderCountMax"`                              // 下单量/天
	AmountMax                   float64   `gorm:"column:amount_max" json:"amountMax"`                                       // 限额/元
	LabelCode                   string    `gorm:"column:label_code" json:"labelCode"`                                       // 客户配置此渠道使用的面单模板code
	SenderAddressBookID         int       `gorm:"column:sender_address_book_id" json:"senderAddressBookId"`                 // 发件人地址簿id
	WeightingWeightDiff         float64   `gorm:"column:weighting_weight_diff" json:"weightingWeightDiff"`                  // 称重重量差kg
	WeightingWeightNegativeDiff float64   `gorm:"column:weighting_weight_negative_diff" json:"weightingWeightNegativeDiff"` // 称重重量差kg(负)
	Proportion                  int       `gorm:"column:proportion" json:"proportion"`                                      // 抛比
	CreateTime                  time.Time `gorm:"column:create_time" json:"createTime"`                                     // 创建时间
	CreateUser                  int       `gorm:"column:create_user" json:"createUser"`                                     // 创建用户
	UpdateTime                  time.Time `gorm:"column:update_time" json:"updateTime"`                                     // 更新时间
	UpdateUser                  int       `gorm:"column:update_user" json:"updateUser"`                                     // 更新时间
	Version                     int       `gorm:"column:version" json:"version"`                                            // 乐观锁
	Deleted                     int       `gorm:"column:deleted" json:"deleted"`                                            // 逻辑删除
	SalesLink                   string    `gorm:"column:sales_link" json:"salesLink"`                                       // 销售链接
}

// TableName get sql table name.获取数据库表名
func (m *CCustomerChannelConfig) TableName() string {
	return "c_customer_channel_config"
}

// CCustomerChannelConfigColumns get sql column name.获取数据库列名
var CCustomerChannelConfigColumns = struct {
	ID                          string
	CustomerID                  string
	CustomerChannelID           string
	OrderCountMax               string
	AmountMax                   string
	LabelCode                   string
	SenderAddressBookID         string
	WeightingWeightDiff         string
	WeightingWeightNegativeDiff string
	Proportion                  string
	CreateTime                  string
	CreateUser                  string
	UpdateTime                  string
	UpdateUser                  string
	Version                     string
	Deleted                     string
	SalesLink                   string
}{
	ID:                          "id",
	CustomerID:                  "customer_id",
	CustomerChannelID:           "customer_channel_id",
	OrderCountMax:               "order_count_max",
	AmountMax:                   "amount_max",
	LabelCode:                   "label_code",
	SenderAddressBookID:         "sender_address_book_id",
	WeightingWeightDiff:         "weighting_weight_diff",
	WeightingWeightNegativeDiff: "weighting_weight_negative_diff",
	Proportion:                  "proportion",
	CreateTime:                  "create_time",
	CreateUser:                  "create_user",
	UpdateTime:                  "update_time",
	UpdateUser:                  "update_user",
	Version:                     "version",
	Deleted:                     "deleted",
	SalesLink:                   "sales_link",
}

// CCustomerCopy1 客户
type CCustomerCopy1 struct {
	ID                     int       `gorm:"primaryKey;column:id" json:"-"`
	Code                   string    `gorm:"column:code" json:"code"`                                       // 客户编号
	Username               string    `gorm:"column:username" json:"username"`                               // 账户名称
	Password               string    `gorm:"column:password" json:"password"`                               // 密码/不得明文
	CustomerAlias          string    `gorm:"column:customer_alias" json:"customerAlias"`                    // 客户别名
	Level                  int       `gorm:"column:level" json:"level"`                                     // 等级，2:一级账户，3:一级账户下面的级别
	ParentID               int       `gorm:"column:parent_id" json:"parentId"`                              // 上级id(客户公司的上下级关系)
	Integral               int       `gorm:"column:integral" json:"integral"`                               // 积分
	Source                 string    `gorm:"column:source" json:"source"`                                   // 客户来源，corporation:企业，person:个人
	Company                string    `gorm:"column:company" json:"company"`                                 // 公司名称
	Address                string    `gorm:"column:address" json:"address"`                                 // 公司(个人)地址
	IDentityCard           string    `gorm:"column:identity_card" json:"identityCard"`                      // 营业执照或身份证号
	IDentityCardAttachment string    `gorm:"column:identity_card_attachment" json:"identityCardAttachment"` // 证件照附件
	CoPrincipalName        string    `gorm:"column:co_principal_name" json:"coPrincipalName"`               // 企业负责人名字
	CoPrincipalTel         string    `gorm:"column:co_principal_tel" json:"coPrincipalTel"`                 // 企业负责人电话
	CoPrincipalPhone       string    `gorm:"column:co_principal_phone" json:"coPrincipalPhone"`             // 企业负责人手机
	CoPrincipalQq          string    `gorm:"column:co_principal_qq" json:"coPrincipalQq"`                   // 企业负责人qq
	CoPrincipalEmail       string    `gorm:"column:co_principal_email" json:"coPrincipalEmail"`             // 企业负责人邮箱
	LogisticsContactsName  string    `gorm:"column:logistics_contacts_name" json:"logisticsContactsName"`   // 物流联系人名字
	LogisticsContactsTel   string    `gorm:"column:logistics_contacts_tel" json:"logisticsContactsTel"`     // 物流联系人电话
	LogisticsContactsPhone string    `gorm:"column:logistics_contacts_phone" json:"logisticsContactsPhone"` // 物流联系人手机
	LogisticsContactsQq    string    `gorm:"column:logistics_contacts_qq" json:"logisticsContactsQq"`       // 物流联系人qq
	LogisticsContactsEmail string    `gorm:"column:logistics_contacts_email" json:"logisticsContactsEmail"` // 物流联系人邮箱
	FinanceContactsName    string    `gorm:"column:finance_contacts_name" json:"financeContactsName"`       // 财务联系人名字
	FinanceContactsTel     string    `gorm:"column:finance_contacts_tel" json:"financeContactsTel"`         // 财务联系人电话
	FinanceContactsPhone   string    `gorm:"column:finance_contacts_phone" json:"financeContactsPhone"`     // 财务联系人手机
	FinanceContactsQq      string    `gorm:"column:finance_contacts_qq" json:"financeContactsQq"`           // 财务联系人qq
	FinanceContactsEmail   string    `gorm:"column:finance_contacts_email" json:"financeContactsEmail"`     // 财务联系人邮箱
	CustomServiceUserID    int       `gorm:"column:custom_service_user_id" json:"customServiceUserId"`      // 客服id(oa系统中的用户id)
	LockState              int       `gorm:"column:lock_state" json:"lockState"`                            // 账号锁定，0：正常，1：锁定
	AuditState             int       `gorm:"column:audit_state" json:"auditState"`                          // 账户审核，0：未审核，1：已审核
	AccountStatus          int       `gorm:"column:account_status" json:"accountStatus"`                    // 账户状态 0：正常，1：欠费，2：禁用
	Balance                float64   `gorm:"column:balance" json:"balance"`                                 // 账户余额
	InitialCreditLimit     float64   `gorm:"column:initial_credit_limit" json:"initialCreditLimit"`         // 信用额度
	OutCumulativeFee       float64   `gorm:"column:out_cumulative_fee" json:"outCumulativeFee"`             // 出库累加费用,Note:距离 上次结算 到 当前时间 累计出库余额 变动
	LastCollectionDate     time.Time `gorm:"column:last_collection_date" json:"lastCollectionDate"`         // 最后收款时间
	PaymentCycle           int       `gorm:"column:payment_cycle" json:"paymentCycle"`                      // 收款周期
	CreateTime             time.Time `gorm:"column:create_time" json:"createTime"`                          // 创建时间
	CreateUser             int       `gorm:"column:create_user" json:"createUser"`                          // 创建用户
	UpdateTime             time.Time `gorm:"column:update_time" json:"updateTime"`                          // 更新时间
	UpdateUser             int       `gorm:"column:update_user" json:"updateUser"`                          // 更新时间
	Version                int       `gorm:"column:version" json:"version"`                                 // 乐观锁
	Deleted                int       `gorm:"column:deleted" json:"deleted"`                                 // 逻辑删除
	AuthorizationCode      string    `gorm:"column:authorization_code" json:"authorizationCode"`            // 授权码
	HandlerName            string    `gorm:"column:handler_name" json:"handlerName"`                        // 经办人名称
	HandlerID              int       `gorm:"column:handler_id" json:"handlerId"`                            // 经办人id
}

// TableName get sql table name.获取数据库表名
func (m *CCustomerCopy1) TableName() string {
	return "c_customer_copy1"
}

// CCustomerCopy1Columns get sql column name.获取数据库列名
var CCustomerCopy1Columns = struct {
	ID                     string
	Code                   string
	Username               string
	Password               string
	CustomerAlias          string
	Level                  string
	ParentID               string
	Integral               string
	Source                 string
	Company                string
	Address                string
	IDentityCard           string
	IDentityCardAttachment string
	CoPrincipalName        string
	CoPrincipalTel         string
	CoPrincipalPhone       string
	CoPrincipalQq          string
	CoPrincipalEmail       string
	LogisticsContactsName  string
	LogisticsContactsTel   string
	LogisticsContactsPhone string
	LogisticsContactsQq    string
	LogisticsContactsEmail string
	FinanceContactsName    string
	FinanceContactsTel     string
	FinanceContactsPhone   string
	FinanceContactsQq      string
	FinanceContactsEmail   string
	CustomServiceUserID    string
	LockState              string
	AuditState             string
	AccountStatus          string
	Balance                string
	InitialCreditLimit     string
	OutCumulativeFee       string
	LastCollectionDate     string
	PaymentCycle           string
	CreateTime             string
	CreateUser             string
	UpdateTime             string
	UpdateUser             string
	Version                string
	Deleted                string
	AuthorizationCode      string
	HandlerName            string
	HandlerID              string
}{
	ID:                     "id",
	Code:                   "code",
	Username:               "username",
	Password:               "password",
	CustomerAlias:          "customer_alias",
	Level:                  "level",
	ParentID:               "parent_id",
	Integral:               "integral",
	Source:                 "source",
	Company:                "company",
	Address:                "address",
	IDentityCard:           "identity_card",
	IDentityCardAttachment: "identity_card_attachment",
	CoPrincipalName:        "co_principal_name",
	CoPrincipalTel:         "co_principal_tel",
	CoPrincipalPhone:       "co_principal_phone",
	CoPrincipalQq:          "co_principal_qq",
	CoPrincipalEmail:       "co_principal_email",
	LogisticsContactsName:  "logistics_contacts_name",
	LogisticsContactsTel:   "logistics_contacts_tel",
	LogisticsContactsPhone: "logistics_contacts_phone",
	LogisticsContactsQq:    "logistics_contacts_qq",
	LogisticsContactsEmail: "logistics_contacts_email",
	FinanceContactsName:    "finance_contacts_name",
	FinanceContactsTel:     "finance_contacts_tel",
	FinanceContactsPhone:   "finance_contacts_phone",
	FinanceContactsQq:      "finance_contacts_qq",
	FinanceContactsEmail:   "finance_contacts_email",
	CustomServiceUserID:    "custom_service_user_id",
	LockState:              "lock_state",
	AuditState:             "audit_state",
	AccountStatus:          "account_status",
	Balance:                "balance",
	InitialCreditLimit:     "initial_credit_limit",
	OutCumulativeFee:       "out_cumulative_fee",
	LastCollectionDate:     "last_collection_date",
	PaymentCycle:           "payment_cycle",
	CreateTime:             "create_time",
	CreateUser:             "create_user",
	UpdateTime:             "update_time",
	UpdateUser:             "update_user",
	Version:                "version",
	Deleted:                "deleted",
	AuthorizationCode:      "authorization_code",
	HandlerName:            "handler_name",
	HandlerID:              "handler_id",
}

// CCustomerFinance 客户财务
type CCustomerFinance struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`            // 主键id
	CustomerID   int       `gorm:"column:customer_id" json:"customerId"`     // 客户id
	CustomerName string    `gorm:"column:customer_name" json:"customerName"` // 客户名称
	Balance      float64   `gorm:"column:balance" json:"balance"`            // 账户余额
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建用户
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新人
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *CCustomerFinance) TableName() string {
	return "c_customer_finance"
}

// CCustomerFinanceColumns get sql column name.获取数据库列名
var CCustomerFinanceColumns = struct {
	ID           string
	CustomerID   string
	CustomerName string
	Balance      string
	CreateUser   string
	CreateTime   string
	UpdateUser   string
	UpdateTime   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	CustomerID:   "customer_id",
	CustomerName: "customer_name",
	Balance:      "balance",
	CreateUser:   "create_user",
	CreateTime:   "create_time",
	UpdateUser:   "update_user",
	UpdateTime:   "update_time",
	Version:      "version",
	Deleted:      "deleted",
}

// CCustomerFollow 客户跟进
type CCustomerFollow struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	CustomerID int       `gorm:"column:customer_id" json:"customerId"` // 客户id
	Content    string    `gorm:"column:content" json:"content"`        // 跟进内容
	Status     int       `gorm:"column:status" json:"status"`          // 状态，0:无效，1:有效
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *CCustomerFollow) TableName() string {
	return "c_customer_follow"
}

// CCustomerFollowColumns get sql column name.获取数据库列名
var CCustomerFollowColumns = struct {
	ID         string
	CustomerID string
	Content    string
	Status     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	CustomerID: "customer_id",
	Content:    "content",
	Status:     "status",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// CCustomerOrganization 客户组织
type CCustomerOrganization struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`                        // ID
	OrganizationCode   string    `gorm:"column:organization_code" json:"organizationCode"`     // 编码
	OrganizationName   string    `gorm:"column:organization_name" json:"organizationName"`     // 名称
	OrganizationStatus int       `gorm:"column:organization_status" json:"organizationStatus"` // 是否启用 0：启用，1：禁用
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                 // 创建时间
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                 // 创建用户
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`                 // 更新时间
	UpdateUser         int       `gorm:"column:update_user" json:"updateUser"`                 // 更新时间
	Version            int       `gorm:"column:version" json:"version"`                        // 乐观锁
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *CCustomerOrganization) TableName() string {
	return "c_customer_organization"
}

// CCustomerOrganizationColumns get sql column name.获取数据库列名
var CCustomerOrganizationColumns = struct {
	ID                 string
	OrganizationCode   string
	OrganizationName   string
	OrganizationStatus string
	CreateTime         string
	CreateUser         string
	UpdateTime         string
	UpdateUser         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	OrganizationCode:   "organization_code",
	OrganizationName:   "organization_name",
	OrganizationStatus: "organization_status",
	CreateTime:         "create_time",
	CreateUser:         "create_user",
	UpdateTime:         "update_time",
	UpdateUser:         "update_user",
	Version:            "version",
	Deleted:            "deleted",
}

// ExcelData 读取excel数据
type ExcelData struct {
	ID  int    `gorm:"primaryKey;column:id" json:"-"` // 主键id
	A1  string `gorm:"column:a1" json:"a1"`
	A2  string `gorm:"column:a2" json:"a2"`
	A3  string `gorm:"column:a3" json:"a3"`
	A4  string `gorm:"column:a4" json:"a4"`
	A5  string `gorm:"column:a5" json:"a5"`
	A6  string `gorm:"column:a6" json:"a6"`
	A7  string `gorm:"column:a7" json:"a7"`
	A8  string `gorm:"column:a8" json:"a8"`
	A9  string `gorm:"column:a9" json:"a9"`
	A10 string `gorm:"column:a10" json:"a10"`
	A11 string `gorm:"column:a11" json:"a11"`
	A12 string `gorm:"column:a12" json:"a12"`
	A13 string `gorm:"column:a13" json:"a13"`
	A14 string `gorm:"column:a14" json:"a14"`
	A15 string `gorm:"column:a15" json:"a15"`
	A16 string `gorm:"column:a16" json:"a16"`
	A17 string `gorm:"column:a17" json:"a17"`
	A18 string `gorm:"column:a18" json:"a18"`
	A19 string `gorm:"column:a19" json:"a19"`
	A20 string `gorm:"column:a20" json:"a20"`
}

// TableName get sql table name.获取数据库表名
func (m *ExcelData) TableName() string {
	return "excel_data"
}

// ExcelDataColumns get sql column name.获取数据库列名
var ExcelDataColumns = struct {
	ID  string
	A1  string
	A2  string
	A3  string
	A4  string
	A5  string
	A6  string
	A7  string
	A8  string
	A9  string
	A10 string
	A11 string
	A12 string
	A13 string
	A14 string
	A15 string
	A16 string
	A17 string
	A18 string
	A19 string
	A20 string
}{
	ID:  "id",
	A1:  "a1",
	A2:  "a2",
	A3:  "a3",
	A4:  "a4",
	A5:  "a5",
	A6:  "a6",
	A7:  "a7",
	A8:  "a8",
	A9:  "a9",
	A10: "a10",
	A11: "a11",
	A12: "a12",
	A13: "a13",
	A14: "a14",
	A15: "a15",
	A16: "a16",
	A17: "a17",
	A18: "a18",
	A19: "a19",
	A20: "a20",
}

// FaAccount 账户管理
type FaAccount struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`              // ID
	AccountAlias  string    `gorm:"column:account_alias" json:"accountAlias"`   // 账户别称
	Status        int       `gorm:"column:status" json:"status"`                // 状态 启用/禁用
	Bank          string    `gorm:"column:bank" json:"bank"`                    // 开户行
	AccountNumber string    `gorm:"column:account_number" json:"accountNumber"` // 账户号码
	AccountName   string    `gorm:"column:account_name" json:"accountName"`     // 账户人名称
	Remark        string    `gorm:"column:remark" json:"remark"`                // 备注
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	CreateUser    int       `gorm:"column:create_user" json:"createUser"`       // 创建用户
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
	UpdateUser    int       `gorm:"column:update_user" json:"updateUser"`       // 更新时间
	Version       int       `gorm:"column:version" json:"version"`              // 乐观锁
	Deleted       int       `gorm:"column:deleted" json:"deleted"`              // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaAccount) TableName() string {
	return "fa_account"
}

// FaAccountColumns get sql column name.获取数据库列名
var FaAccountColumns = struct {
	ID            string
	AccountAlias  string
	Status        string
	Bank          string
	AccountNumber string
	AccountName   string
	Remark        string
	CreateTime    string
	CreateUser    string
	UpdateTime    string
	UpdateUser    string
	Version       string
	Deleted       string
}{
	ID:            "id",
	AccountAlias:  "account_alias",
	Status:        "status",
	Bank:          "bank",
	AccountNumber: "account_number",
	AccountName:   "account_name",
	Remark:        "remark",
	CreateTime:    "create_time",
	CreateUser:    "create_user",
	UpdateTime:    "update_time",
	UpdateUser:    "update_user",
	Version:       "version",
	Deleted:       "deleted",
}

// FaBillVoucher 账单凭证
type FaBillVoucher struct {
	ID              int64     `gorm:"primaryKey;column:id" json:"-"`                  // 主键
	Title           string    `gorm:"column:title" json:"title"`                      // 凭证标题
	CompanyID       int       `gorm:"column:company_id" json:"companyId"`             // 凭证公司id(客户组织id)
	Company         string    `gorm:"column:company" json:"company"`                  // 凭证公司
	Code            string    `gorm:"column:code" json:"code"`                        // 凭证编码
	VoucherTime     string    `gorm:"column:voucher_time" json:"voucherTime"`         // 凭证年月
	Attachment      string    `gorm:"column:attachment" json:"attachment"`            // 附件：附件上传工具，必填
	AttachmentCount int       `gorm:"column:attachment_count" json:"attachmentCount"` // 附件数量
	Remark          string    `gorm:"column:remark" json:"remark"`                    // 备注
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`           // 创建人
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	UpdateUser      int       `gorm:"column:update_user" json:"updateUser"`           // 更新人
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`           // 更新时间
	Version         int       `gorm:"column:version" json:"version"`                  // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                  // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaBillVoucher) TableName() string {
	return "fa_bill_voucher"
}

// FaBillVoucherColumns get sql column name.获取数据库列名
var FaBillVoucherColumns = struct {
	ID              string
	Title           string
	CompanyID       string
	Company         string
	Code            string
	VoucherTime     string
	Attachment      string
	AttachmentCount string
	Remark          string
	CreateUser      string
	CreateTime      string
	UpdateUser      string
	UpdateTime      string
	Version         string
	Deleted         string
}{
	ID:              "id",
	Title:           "title",
	CompanyID:       "company_id",
	Company:         "company",
	Code:            "code",
	VoucherTime:     "voucher_time",
	Attachment:      "attachment",
	AttachmentCount: "attachment_count",
	Remark:          "remark",
	CreateUser:      "create_user",
	CreateTime:      "create_time",
	UpdateUser:      "update_user",
	UpdateTime:      "update_time",
	Version:         "version",
	Deleted:         "deleted",
}

// FaBusinessEntity 业务主体
type FaBusinessEntity struct {
	ID                 int64     `gorm:"primaryKey;column:id" json:"-"`                         // 主键
	Code               string    `gorm:"column:code" json:"code"`                               // 业务主体代码
	Name               string    `gorm:"column:name" json:"name"`                               // 业务主体名称
	Type               bool      `gorm:"column:type" json:"type"`                               // 业务主体类型，0:服务商，1:航司
	Remark             string    `gorm:"column:remark" json:"remark"`                           // 备注
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                  // 创建时间
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                  // 创建人
	UpdateUser         int       `gorm:"column:update_user" json:"updateUser"`                  // 更新人
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`                  // 更新时间
	Version            int       `gorm:"column:version" json:"version"`                         // 乐观锁
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                         // 逻辑删除
	KingdeeID          string    `gorm:"column:kingdee_id" json:"kingdeeId"`                    // 金蝶客户内码
	InputCumulativeFee float64   `gorm:"column:input_cumulative_fee" json:"inputCumulativeFee"` // 入库预扣款
	OutCumulativeFee   float64   `gorm:"column:out_cumulative_fee" json:"outCumulativeFee"`     // 出库未结算
}

// TableName get sql table name.获取数据库表名
func (m *FaBusinessEntity) TableName() string {
	return "fa_business_entity"
}

// FaBusinessEntityColumns get sql column name.获取数据库列名
var FaBusinessEntityColumns = struct {
	ID                 string
	Code               string
	Name               string
	Type               string
	Remark             string
	CreateTime         string
	CreateUser         string
	UpdateUser         string
	UpdateTime         string
	Version            string
	Deleted            string
	KingdeeID          string
	InputCumulativeFee string
	OutCumulativeFee   string
}{
	ID:                 "id",
	Code:               "code",
	Name:               "name",
	Type:               "type",
	Remark:             "remark",
	CreateTime:         "create_time",
	CreateUser:         "create_user",
	UpdateUser:         "update_user",
	UpdateTime:         "update_time",
	Version:            "version",
	Deleted:            "deleted",
	KingdeeID:          "kingdee_id",
	InputCumulativeFee: "input_cumulative_fee",
	OutCumulativeFee:   "out_cumulative_fee",
}

// FaBusinessEntityPayment 业务主体付款单
type FaBusinessEntityPayment struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`                         // 主键
	PaymentNumber      string    `gorm:"column:payment_number" json:"paymentNumber"`            // 付款单号
	ParentID           int       `gorm:"column:parent_id" json:"parentId"`                      // 父级id
	IsParent           []uint8   `gorm:"column:is_parent" json:"isParent"`                      // 是否为父级；0-否，1-是
	InvoiceNumber      string    `gorm:"column:invoice_number" json:"invoiceNumber"`            // 发票号
	VoucherNumber      string    `gorm:"column:voucher_number" json:"voucherNumber"`            // 凭证号
	TicketNumber       string    `gorm:"column:ticket_number" json:"ticketNumber"`              // 票证号
	PaymentStatus      int       `gorm:"column:payment_status" json:"paymentStatus"`            // 付款状态:0未付,  1已付清, 2未付清
	IsPayment          []uint8   `gorm:"column:is_payment" json:"isPayment"`                    // 是否已付款
	BusinessEntityCode string    `gorm:"column:business_entity_code" json:"businessEntityCode"` // 业务主体代码
	BusinessEntityName string    `gorm:"column:business_entity_name" json:"businessEntityName"` // 业务实体名称
	BusinessEntityType bool      `gorm:"column:business_entity_type" json:"businessEntityType"` // 业务主体类型，0:服务商，1:航司
	BillFee            float64   `gorm:"column:bill_fee" json:"billFee"`                        // 票据金额
	ActualFee          float64   `gorm:"column:actual_fee" json:"actualFee"`                    // 实际付款金额
	CurrencyCode       string    `gorm:"column:currency_code" json:"currencyCode"`              // 币种
	CurrencyName       string    `gorm:"column:currency_name" json:"currencyName"`              // 币种中文名称
	ExchangeRate       float64   `gorm:"column:exchange_rate" json:"exchangeRate"`              // 汇率
	ActualFeeRmb       float64   `gorm:"column:actual_fee_rmb" json:"actualFeeRmb"`             // 实际收款金额rmb
	DealWithUser       string    `gorm:"column:deal_with_user" json:"dealWithUser"`             // 经办人
	BillInputUser      string    `gorm:"column:bill_input_user" json:"billInputUser"`           // 票据录入人员
	BillInputUserID    int       `gorm:"column:bill_input_user_id" json:"billInputUserId"`      // 票据录入人员ID
	PaymentType        int       `gorm:"column:payment_type" json:"paymentType"`                // 付款类别 0:冲销，1:核付
	PaymentMethod      string    `gorm:"column:payment_method" json:"paymentMethod"`            // 付款方式
	PaymentTime        time.Time `gorm:"column:payment_time" json:"paymentTime"`                // 支付时间
	Mark               string    `gorm:"column:mark" json:"mark"`                               // 标签/标记
	AccountAlias       string    `gorm:"column:account_alias" json:"accountAlias"`              // 账户别称/走账
	AttachmentURL      string    `gorm:"column:attachment_url" json:"attachmentUrl"`            // 附件
	Remark             string    `gorm:"column:remark" json:"remark"`                           // 备注
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                  // 创建时间
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                  // 生成用户,生成应付人员
	UpdateUser         int       `gorm:"column:update_user" json:"updateUser"`                  // 更新人
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`                  // 更新时间
	Version            int       `gorm:"column:version" json:"version"`                         // 乐观锁
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                         // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaBusinessEntityPayment) TableName() string {
	return "fa_business_entity_payment"
}

// FaBusinessEntityPaymentColumns get sql column name.获取数据库列名
var FaBusinessEntityPaymentColumns = struct {
	ID                 string
	PaymentNumber      string
	ParentID           string
	IsParent           string
	InvoiceNumber      string
	VoucherNumber      string
	TicketNumber       string
	PaymentStatus      string
	IsPayment          string
	BusinessEntityCode string
	BusinessEntityName string
	BusinessEntityType string
	BillFee            string
	ActualFee          string
	CurrencyCode       string
	CurrencyName       string
	ExchangeRate       string
	ActualFeeRmb       string
	DealWithUser       string
	BillInputUser      string
	BillInputUserID    string
	PaymentType        string
	PaymentMethod      string
	PaymentTime        string
	Mark               string
	AccountAlias       string
	AttachmentURL      string
	Remark             string
	CreateTime         string
	CreateUser         string
	UpdateUser         string
	UpdateTime         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	PaymentNumber:      "payment_number",
	ParentID:           "parent_id",
	IsParent:           "is_parent",
	InvoiceNumber:      "invoice_number",
	VoucherNumber:      "voucher_number",
	TicketNumber:       "ticket_number",
	PaymentStatus:      "payment_status",
	IsPayment:          "is_payment",
	BusinessEntityCode: "business_entity_code",
	BusinessEntityName: "business_entity_name",
	BusinessEntityType: "business_entity_type",
	BillFee:            "bill_fee",
	ActualFee:          "actual_fee",
	CurrencyCode:       "currency_code",
	CurrencyName:       "currency_name",
	ExchangeRate:       "exchange_rate",
	ActualFeeRmb:       "actual_fee_rmb",
	DealWithUser:       "deal_with_user",
	BillInputUser:      "bill_input_user",
	BillInputUserID:    "bill_input_user_id",
	PaymentType:        "payment_type",
	PaymentMethod:      "payment_method",
	PaymentTime:        "payment_time",
	Mark:               "mark",
	AccountAlias:       "account_alias",
	AttachmentURL:      "attachment_url",
	Remark:             "remark",
	CreateTime:         "create_time",
	CreateUser:         "create_user",
	UpdateUser:         "update_user",
	UpdateTime:         "update_time",
	Version:            "version",
	Deleted:            "deleted",
}

// FaBusinessEntityPaymentDetails 业务主体付款单明细
type FaBusinessEntityPaymentDetails struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`                             // 主键
	PaymentNumber        string    `gorm:"column:payment_number" json:"paymentNumber"`                // 付款单号
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	BusinessEntityCode   string    `gorm:"column:business_entity_code" json:"businessEntityCode"`     // 业务主体代码
	BusinessEntityName   string    `gorm:"column:business_entity_name" json:"businessEntityName"`     // 业务实体名称
	BusinessEntityType   bool      `gorm:"column:business_entity_type" json:"businessEntityType"`     // 业务主体类型，0:服务商，1:航司
	IsAccept             []uint8   `gorm:"column:is_accept" json:"isAccept"`                          // 已核付(关联付款单状态)
	AcceptTime           time.Time `gorm:"column:accept_time" json:"acceptTime"`                      // 核付时间
	AcceptUser           string    `gorm:"column:accept_user" json:"acceptUser"`                      // 核付用户
	AcceptUserID         int       `gorm:"column:accept_user_id" json:"acceptUserId"`                 // 核付用户ID
	Weight               float64   `gorm:"column:weight" json:"weight"`                               // 计费总重量
	PayableRmb           float64   `gorm:"column:payable_rmb" json:"payableRmb"`                      // 应付总费用(RMB)
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 备注
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 生成用户,生成应收人员
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新人
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaBusinessEntityPaymentDetails) TableName() string {
	return "fa_business_entity_payment_details"
}

// FaBusinessEntityPaymentDetailsColumns get sql column name.获取数据库列名
var FaBusinessEntityPaymentDetailsColumns = struct {
	ID                   string
	PaymentNumber        string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	BusinessEntityCode   string
	BusinessEntityName   string
	BusinessEntityType   string
	IsAccept             string
	AcceptTime           string
	AcceptUser           string
	AcceptUserID         string
	Weight               string
	PayableRmb           string
	Remark               string
	CreateTime           string
	CreateUser           string
	UpdateUser           string
	UpdateTime           string
	Version              string
}{
	ID:                   "id",
	PaymentNumber:        "payment_number",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	BusinessEntityCode:   "business_entity_code",
	BusinessEntityName:   "business_entity_name",
	BusinessEntityType:   "business_entity_type",
	IsAccept:             "is_accept",
	AcceptTime:           "accept_time",
	AcceptUser:           "accept_user",
	AcceptUserID:         "accept_user_id",
	Weight:               "weight",
	PayableRmb:           "payable_rmb",
	Remark:               "remark",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateUser:           "update_user",
	UpdateTime:           "update_time",
	Version:              "version",
}

// FaBusinessEntityPaymentRecord 业务主体付款流水
type FaBusinessEntityPaymentRecord struct {
	ID                  int       `gorm:"primaryKey;column:id" json:"-"`                            // 主键
	PaymentNumber       string    `gorm:"column:payment_number" json:"paymentNumber"`               // 付款单号
	InvoiceNumber       string    `gorm:"column:invoice_number" json:"invoiceNumber"`               // 发票号
	VoucherNumber       string    `gorm:"column:voucher_number" json:"voucherNumber"`               // 凭证号
	TicketNumber        string    `gorm:"column:ticket_number" json:"ticketNumber"`                 // 票证号
	BusinessEntityCode  string    `gorm:"column:business_entity_code" json:"businessEntityCode"`    // 业务主体代码
	BusinessEntityName  string    `gorm:"column:business_entity_name" json:"businessEntityName"`    // 业务实体名称
	BusinessEntityType  bool      `gorm:"column:business_entity_type" json:"businessEntityType"`    // 业务主体类型，0:服务商，1:航司
	PayableFeeRmb       float64   `gorm:"column:payable_fee_rmb" json:"payableFeeRmb"`              // 应付金额
	ActualPaymentFee    float64   `gorm:"column:actual_payment_fee" json:"actualPaymentFee"`        // 实付金额
	CurrencyCode        string    `gorm:"column:currency_code" json:"currencyCode"`                 // 币种
	CurrencyName        string    `gorm:"column:currency_name" json:"currencyName"`                 // 币种中文名称
	ExchangeRate        float64   `gorm:"column:exchange_rate" json:"exchangeRate"`                 // 汇率
	ActualPaymentFeeRmb float64   `gorm:"column:actual_payment_fee_rmb" json:"actualPaymentFeeRmb"` // 实付金额(人民币)
	IsPaid              []uint8   `gorm:"column:is_paid" json:"isPaid"`                             // 已付清
	PaymentMethod       string    `gorm:"column:payment_method" json:"paymentMethod"`               // 付款方式
	PaymentAccount      string    `gorm:"column:payment_account" json:"paymentAccount"`             // 付款账户
	PaymentAccountAlias string    `gorm:"column:payment_account_alias" json:"paymentAccountAlias"`  // 付款账户别称
	PaymentTime         time.Time `gorm:"column:payment_time" json:"paymentTime"`                   // 支付时间
	DealWithUser        string    `gorm:"column:deal_with_user" json:"dealWithUser"`                // 经办人
	AttachmentURL       string    `gorm:"column:attachment_url" json:"attachmentUrl"`               // 附件
	Remark              string    `gorm:"column:remark" json:"remark"`                              // 备注
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`                     // 创建时间
	CreateUser          int       `gorm:"column:create_user" json:"createUser"`                     // 流水生成用户
	UpdateUser          int       `gorm:"column:update_user" json:"updateUser"`                     // 更新人
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`                     // 更新时间
	Version             int       `gorm:"column:version" json:"version"`                            // 乐观锁
	Deleted             int       `gorm:"column:deleted" json:"deleted"`                            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaBusinessEntityPaymentRecord) TableName() string {
	return "fa_business_entity_payment_record"
}

// FaBusinessEntityPaymentRecordColumns get sql column name.获取数据库列名
var FaBusinessEntityPaymentRecordColumns = struct {
	ID                  string
	PaymentNumber       string
	InvoiceNumber       string
	VoucherNumber       string
	TicketNumber        string
	BusinessEntityCode  string
	BusinessEntityName  string
	BusinessEntityType  string
	PayableFeeRmb       string
	ActualPaymentFee    string
	CurrencyCode        string
	CurrencyName        string
	ExchangeRate        string
	ActualPaymentFeeRmb string
	IsPaid              string
	PaymentMethod       string
	PaymentAccount      string
	PaymentAccountAlias string
	PaymentTime         string
	DealWithUser        string
	AttachmentURL       string
	Remark              string
	CreateTime          string
	CreateUser          string
	UpdateUser          string
	UpdateTime          string
	Version             string
	Deleted             string
}{
	ID:                  "id",
	PaymentNumber:       "payment_number",
	InvoiceNumber:       "invoice_number",
	VoucherNumber:       "voucher_number",
	TicketNumber:        "ticket_number",
	BusinessEntityCode:  "business_entity_code",
	BusinessEntityName:  "business_entity_name",
	BusinessEntityType:  "business_entity_type",
	PayableFeeRmb:       "payable_fee_rmb",
	ActualPaymentFee:    "actual_payment_fee",
	CurrencyCode:        "currency_code",
	CurrencyName:        "currency_name",
	ExchangeRate:        "exchange_rate",
	ActualPaymentFeeRmb: "actual_payment_fee_rmb",
	IsPaid:              "is_paid",
	PaymentMethod:       "payment_method",
	PaymentAccount:      "payment_account",
	PaymentAccountAlias: "payment_account_alias",
	PaymentTime:         "payment_time",
	DealWithUser:        "deal_with_user",
	AttachmentURL:       "attachment_url",
	Remark:              "remark",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateUser:          "update_user",
	UpdateTime:          "update_time",
	Version:             "version",
	Deleted:             "deleted",
}

// FaCurrencyType 货币类型表
type FaCurrencyType struct {
	ID           int64     `gorm:"primaryKey;column:id" json:"-"`
	Code         string    `gorm:"column:code" json:"code"`                  // 货币编码
	ChineseName  string    `gorm:"column:chinese_name" json:"chineseName"`   // 中文名称
	EnglishName  string    `gorm:"column:english_name" json:"englishName"`   // 英文名称
	ExchangeRate float64   `gorm:"column:exchange_rate" json:"exchangeRate"` // 汇率
	Status       int       `gorm:"column:status" json:"status"`              // 状态： 1 启用 0 禁用
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 修改时间
	CreateUser   string    `gorm:"column:create_user" json:"createUser"`     // 创建用户
	UpdateUser   string    `gorm:"column:update_user" json:"updateUser"`     // 修改用户
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Remark       string    `gorm:"column:remark" json:"remark"`              // 备注
}

// TableName get sql table name.获取数据库表名
func (m *FaCurrencyType) TableName() string {
	return "fa_currency_type"
}

// FaCurrencyTypeColumns get sql column name.获取数据库列名
var FaCurrencyTypeColumns = struct {
	ID           string
	Code         string
	ChineseName  string
	EnglishName  string
	ExchangeRate string
	Status       string
	CreateTime   string
	UpdateTime   string
	CreateUser   string
	UpdateUser   string
	Version      string
	Remark       string
}{
	ID:           "id",
	Code:         "code",
	ChineseName:  "chinese_name",
	EnglishName:  "english_name",
	ExchangeRate: "exchange_rate",
	Status:       "status",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	CreateUser:   "create_user",
	UpdateUser:   "update_user",
	Version:      "version",
	Remark:       "remark",
}

// FaCustomerAccountOperateLog 客户账户操作日志
type FaCustomerAccountOperateLog struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`                        // 主键id
	OperateModule      string    `gorm:"column:operate_module" json:"operateModule"`           // 操作模块
	CustomerID         int       `gorm:"column:customer_id" json:"customerId"`                 // 客户id
	OperateType        string    `gorm:"column:operate_type" json:"operateType"`               // 操作类型
	OperateDescription string    `gorm:"column:operate_description" json:"operateDescription"` // 操作描述
	UpdateParams       string    `gorm:"column:update_params" json:"updateParams"`             // 更新参数
	PaymentMethod      string    `gorm:"column:payment_method" json:"paymentMethod"`           // 收款方式
	Amount             float64   `gorm:"column:amount" json:"amount"`                          // 金额
	RechargeDate       time.Time `gorm:"column:recharge_date" json:"rechargeDate"`             // 充值时间
	Account            string    `gorm:"column:account" json:"account"`                        // 账户
	TicketNumber       string    `gorm:"column:ticket_number" json:"ticketNumber"`             // 票证号
	Attachment         string    `gorm:"column:attachment" json:"attachment"`                  // 附件
	Remark             string    `gorm:"column:remark" json:"remark"`                          // 备注
	Agent              string    `gorm:"column:agent" json:"agent"`                            // 经办人
	OperateIP          string    `gorm:"column:operate_ip" json:"operateIp"`                   // 操作IP
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                 // 操作人id
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                 // 操作时间
	Version            int       `gorm:"column:version" json:"version"`                        // 乐观锁
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaCustomerAccountOperateLog) TableName() string {
	return "fa_customer_account_operate_log"
}

// FaCustomerAccountOperateLogColumns get sql column name.获取数据库列名
var FaCustomerAccountOperateLogColumns = struct {
	ID                 string
	OperateModule      string
	CustomerID         string
	OperateType        string
	OperateDescription string
	UpdateParams       string
	PaymentMethod      string
	Amount             string
	RechargeDate       string
	Account            string
	TicketNumber       string
	Attachment         string
	Remark             string
	Agent              string
	OperateIP          string
	CreateUser         string
	CreateTime         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	OperateModule:      "operate_module",
	CustomerID:         "customer_id",
	OperateType:        "operate_type",
	OperateDescription: "operate_description",
	UpdateParams:       "update_params",
	PaymentMethod:      "payment_method",
	Amount:             "amount",
	RechargeDate:       "recharge_date",
	Account:            "account",
	TicketNumber:       "ticket_number",
	Attachment:         "attachment",
	Remark:             "remark",
	Agent:              "agent",
	OperateIP:          "operate_ip",
	CreateUser:         "create_user",
	CreateTime:         "create_time",
	Version:            "version",
	Deleted:            "deleted",
}

// FaCustomerBill 客户对账单/应收账单
type FaCustomerBill struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`                             // 主键
	BillBatchNumber      string    `gorm:"column:bill_batch_number" json:"billBatchNumber"`           // 账单批次号
	BillBatchStatus      int       `gorm:"column:bill_batch_status" json:"billBatchStatus"`           // 账单状态
	ReceiptNumber        string    `gorm:"column:receipt_number" json:"receiptNumber"`                // 收款单号
	CustomerID           int       `gorm:"column:customer_id" json:"customerId"`                      // 客户ID
	CustomerAlias        string    `gorm:"column:customer_alias" json:"customerAlias"`                // 客户别名
	CustomerChannelNames string    `gorm:"column:customer_channel_names" json:"customerChannelNames"` // 客户渠道名称
	OrderQuantity        int       `gorm:"column:order_quantity" json:"orderQuantity"`                // 订单数量
	FeeQuantity          int       `gorm:"column:fee_quantity" json:"feeQuantity"`                    // 费用条数
	Weight               float64   `gorm:"column:weight" json:"weight"`                               // 计费总重量
	ReceivablesRmb       float64   `gorm:"column:receivables_rmb" json:"receivablesRmb"`              // 应收总费用(RMB)
	ActualFeeRmb         float64   `gorm:"column:actual_fee_rmb" json:"actualFeeRmb"`                 // 实际收款金额rmb
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 生成时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 生成用户,生成应收人员
	IsAudit              []uint8   `gorm:"column:is_audit" json:"isAudit"`                            // 审核状态(关联对账单批次号状态)
	IsAuditRollback      []uint8   `gorm:"column:is_audit_rollback" json:"isAuditRollback"`           // 是否反审标识 默认false ,被反审后true,重新审核后false
	AuditRollbackUser    string    `gorm:"column:audit_rollback_user" json:"auditRollbackUser"`       // 反审核人员
	AuditRollbackUserID  int       `gorm:"column:audit_rollback_user_id" json:"auditRollbackUserId"`  // 反审核人员ID
	CancelUser           string    `gorm:"column:cancel_user" json:"cancelUser"`                      // 取消应收用户
	CancelUserID         int       `gorm:"column:cancel_user_id" json:"cancelUserId"`                 // 取消应收用户ID
	AuditUser            string    `gorm:"column:audit_user" json:"auditUser"`                        // 审核用户
	AuditUserID          int       `gorm:"column:audit_user_id" json:"auditUserId"`                   // 审核用户
	AuditTime            time.Time `gorm:"column:audit_time" json:"auditTime"`                        // 审核时间
	IsReceipt            []uint8   `gorm:"column:is_receipt" json:"isReceipt"`                        // 已生成收款单
	ReceiptNumberUser    string    `gorm:"column:receipt_number_user" json:"receiptNumberUser"`       // 收款单生成用户
	ReceiptNumberUserID  int       `gorm:"column:receipt_number_user_id" json:"receiptNumberUserId"`  // 收款单生成用户ID
	ReceiptNumberTime    time.Time `gorm:"column:receipt_number_time" json:"receiptNumberTime"`       // 收款单生成用户生成时间
	IsAccept             []uint8   `gorm:"column:is_accept" json:"isAccept"`                          // 已核收(关联收款单状态)
	AcceptTime           time.Time `gorm:"column:accept_time" json:"acceptTime"`                      // 核收时间
	AcceptUser           string    `gorm:"column:accept_user" json:"acceptUser"`                      // 核收用户
	AcceptUserID         int       `gorm:"column:accept_user_id" json:"acceptUserId"`                 // 核收用户ID
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 备注
	FinancialRemark      string    `gorm:"column:financial_remark" json:"financialRemark"`            // 财务备注
	AcceptRemark         string    `gorm:"column:accept_remark" json:"acceptRemark"`                  // 核收备注
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新人
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	Deleted              int       `gorm:"column:deleted" json:"deleted"`                             // 逻辑删除
	ChannelHaul          string    `gorm:"column:channel_haul" json:"channelHaul"`                    // 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
	KingdeeID            string    `gorm:"column:kingdee_id" json:"kingdeeId"`                        // 金蝶客户内码
	LastSyncTime         time.Time `gorm:"column:last_sync_time" json:"lastSyncTime"`                 // 最后同步时间
	LastSyncUser         int       `gorm:"column:last_sync_user" json:"lastSyncUser"`                 // 最后同步人
}

// TableName get sql table name.获取数据库表名
func (m *FaCustomerBill) TableName() string {
	return "fa_customer_bill"
}

// FaCustomerBillColumns get sql column name.获取数据库列名
var FaCustomerBillColumns = struct {
	ID                   string
	BillBatchNumber      string
	BillBatchStatus      string
	ReceiptNumber        string
	CustomerID           string
	CustomerAlias        string
	CustomerChannelNames string
	OrderQuantity        string
	FeeQuantity          string
	Weight               string
	ReceivablesRmb       string
	ActualFeeRmb         string
	CreateTime           string
	CreateUser           string
	IsAudit              string
	IsAuditRollback      string
	AuditRollbackUser    string
	AuditRollbackUserID  string
	CancelUser           string
	CancelUserID         string
	AuditUser            string
	AuditUserID          string
	AuditTime            string
	IsReceipt            string
	ReceiptNumberUser    string
	ReceiptNumberUserID  string
	ReceiptNumberTime    string
	IsAccept             string
	AcceptTime           string
	AcceptUser           string
	AcceptUserID         string
	Remark               string
	FinancialRemark      string
	AcceptRemark         string
	UpdateUser           string
	UpdateTime           string
	Version              string
	Deleted              string
	ChannelHaul          string
	KingdeeID            string
	LastSyncTime         string
	LastSyncUser         string
}{
	ID:                   "id",
	BillBatchNumber:      "bill_batch_number",
	BillBatchStatus:      "bill_batch_status",
	ReceiptNumber:        "receipt_number",
	CustomerID:           "customer_id",
	CustomerAlias:        "customer_alias",
	CustomerChannelNames: "customer_channel_names",
	OrderQuantity:        "order_quantity",
	FeeQuantity:          "fee_quantity",
	Weight:               "weight",
	ReceivablesRmb:       "receivables_rmb",
	ActualFeeRmb:         "actual_fee_rmb",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	IsAudit:              "is_audit",
	IsAuditRollback:      "is_audit_rollback",
	AuditRollbackUser:    "audit_rollback_user",
	AuditRollbackUserID:  "audit_rollback_user_id",
	CancelUser:           "cancel_user",
	CancelUserID:         "cancel_user_id",
	AuditUser:            "audit_user",
	AuditUserID:          "audit_user_id",
	AuditTime:            "audit_time",
	IsReceipt:            "is_receipt",
	ReceiptNumberUser:    "receipt_number_user",
	ReceiptNumberUserID:  "receipt_number_user_id",
	ReceiptNumberTime:    "receipt_number_time",
	IsAccept:             "is_accept",
	AcceptTime:           "accept_time",
	AcceptUser:           "accept_user",
	AcceptUserID:         "accept_user_id",
	Remark:               "remark",
	FinancialRemark:      "financial_remark",
	AcceptRemark:         "accept_remark",
	UpdateUser:           "update_user",
	UpdateTime:           "update_time",
	Version:              "version",
	Deleted:              "deleted",
	ChannelHaul:          "channel_haul",
	KingdeeID:            "kingdee_id",
	LastSyncTime:         "last_sync_time",
	LastSyncUser:         "last_sync_user",
}

// FaCustomerBillFee 客户对账单费用 主表
type FaCustomerBillFee struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`                             // 主键
	OrderID              int       `gorm:"column:order_id" json:"orderId"`                            // 订单ID
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	CustomerID           int       `gorm:"column:customer_id" json:"customerId"`                      // 客户ID
	CustomerAlias        string    `gorm:"column:customer_alias" json:"customerAlias"`                // 客户别名
	CustomerChannelID    int       `gorm:"column:customer_channel_id" json:"customerChannelId"`       // 客户渠道ID
	CustomerChannelName  string    `gorm:"column:customer_channel_name" json:"customerChannelName"`   // 客户渠道名称
	ReceiptTime          time.Time `gorm:"column:receipt_time" json:"receiptTime"`                    // 入库时间
	SendTime             time.Time `gorm:"column:send_time" json:"sendTime"`                          // 出库时间
	IsSend               []uint8   `gorm:"column:is_send" json:"isSend"`                              // 是否出库
	Weight               float64   `gorm:"column:weight" json:"weight"`                               // 预报重量
	WeighingWeight       float64   `gorm:"column:weighing_weight" json:"weighingWeight"`              // 称重重量
	FoamWeight           float64   `gorm:"column:foam_weight" json:"foamWeight"`                      // 泡重
	ChargeWeight         float64   `gorm:"column:charge_weight" json:"chargeWeight"`                  // 计费重量
	FeeQty               int       `gorm:"column:fee_qty" json:"feeQty"`                              // 费用数量
	ReceivablesRmb       float64   `gorm:"column:receivables_rmb" json:"receivablesRmb"`              // 应收费用(RMB)
	BillBatchNumbers     string    `gorm:"column:bill_batch_numbers" json:"billBatchNumbers"`         // 客户对账单批次号
	ReceiptNumbers       string    `gorm:"column:receipt_numbers" json:"receiptNumbers"`              // 客户收款单批次号
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 备注
	FinancialRemark      string    `gorm:"column:financial_remark" json:"financialRemark"`            // 财务备注
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 生成用户,生成应收人员
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新人
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaCustomerBillFee) TableName() string {
	return "fa_customer_bill_fee"
}

// FaCustomerBillFeeColumns get sql column name.获取数据库列名
var FaCustomerBillFeeColumns = struct {
	ID                   string
	OrderID              string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	CustomerID           string
	CustomerAlias        string
	CustomerChannelID    string
	CustomerChannelName  string
	ReceiptTime          string
	SendTime             string
	IsSend               string
	Weight               string
	WeighingWeight       string
	FoamWeight           string
	ChargeWeight         string
	FeeQty               string
	ReceivablesRmb       string
	BillBatchNumbers     string
	ReceiptNumbers       string
	Remark               string
	FinancialRemark      string
	CreateTime           string
	CreateUser           string
	UpdateUser           string
	UpdateTime           string
	Version              string
}{
	ID:                   "id",
	OrderID:              "order_id",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	CustomerID:           "customer_id",
	CustomerAlias:        "customer_alias",
	CustomerChannelID:    "customer_channel_id",
	CustomerChannelName:  "customer_channel_name",
	ReceiptTime:          "receipt_time",
	SendTime:             "send_time",
	IsSend:               "is_send",
	Weight:               "weight",
	WeighingWeight:       "weighing_weight",
	FoamWeight:           "foam_weight",
	ChargeWeight:         "charge_weight",
	FeeQty:               "fee_qty",
	ReceivablesRmb:       "receivables_rmb",
	BillBatchNumbers:     "bill_batch_numbers",
	ReceiptNumbers:       "receipt_numbers",
	Remark:               "remark",
	FinancialRemark:      "financial_remark",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateUser:           "update_user",
	UpdateTime:           "update_time",
	Version:              "version",
}

// FaCustomerBillFeeDetails 客户对账单费用明细
type FaCustomerBillFeeDetails struct {
	ID                        int64     `gorm:"primaryKey;column:id" json:"-"`                                         // 主键
	OrderNumber               string    `gorm:"column:order_number" json:"orderNumber"`                                // 订单号
	ReferenceNumber           string    `gorm:"column:reference_number" json:"referenceNumber"`                        // 参考号
	LogisticsNumber           string    `gorm:"column:logistics_number" json:"logisticsNumber"`                        // 物流单号
	LogisticsNumberFinal      string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"`             // 最终物流单号
	CustomerID                int       `gorm:"column:customer_id" json:"customerId"`                                  // 客户ID
	CustomerAlias             string    `gorm:"column:customer_alias" json:"customerAlias"`                            // 客户别名
	CustomerChannelID         int       `gorm:"column:customer_channel_id" json:"customerChannelId"`                   // 客户渠道ID
	CustomerChannelName       string    `gorm:"column:customer_channel_name" json:"customerChannelName"`               // 客户渠道名称
	ReceiptTime               time.Time `gorm:"column:receipt_time" json:"receiptTime"`                                // 入库时间
	SendTime                  time.Time `gorm:"column:send_time" json:"sendTime"`                                      // 出库时间
	IsSend                    []uint8   `gorm:"column:is_send" json:"isSend"`                                          // 是否出库
	Weight                    float64   `gorm:"column:weight" json:"weight"`                                           // 预报重量
	WeighingWeight            float64   `gorm:"column:weighing_weight" json:"weighingWeight"`                          // 称重重量
	FoamWeight                float64   `gorm:"column:foam_weight" json:"foamWeight"`                                  // 泡重
	ChargeWeight              float64   `gorm:"column:charge_weight" json:"chargeWeight"`                              // 计费重量
	FeeTypeName               string    `gorm:"column:fee_type_name" json:"feeTypeName"`                               // 费用类型名称
	Receivables               float64   `gorm:"column:receivables" json:"receivables"`                                 // 应收费用
	CurrencyCode              string    `gorm:"column:currency_code" json:"currencyCode"`                              // 币种
	CurrencyName              string    `gorm:"column:currency_name" json:"currencyName"`                              // 币种中文名称
	ExchangeRate              float64   `gorm:"column:exchange_rate" json:"exchangeRate"`                              // 汇率
	ReceivablesRmb            float64   `gorm:"column:receivables_rmb" json:"receivablesRmb"`                          // 应收费用(RMB)
	FeeSource                 int       `gorm:"column:fee_source" json:"feeSource"`                                    // 费用来源  0:系统生成  1:手工添加
	BillBatchNumber           string    `gorm:"column:bill_batch_number" json:"billBatchNumber"`                       // 对账单批次号
	BillBatchNumberUser       string    `gorm:"column:bill_batch_number_user" json:"billBatchNumberUser"`              // 生成对账用户
	BillBatchNumberUserID     int       `gorm:"column:bill_batch_number_user_id" json:"billBatchNumberUserId"`         // 生成对账用户
	BillBatchNumberTime       time.Time `gorm:"column:bill_batch_number_time" json:"billBatchNumberTime"`              // 生成对账时间
	IsGenerateBillBatchNumber []uint8   `gorm:"column:is_generate_bill_batch_number" json:"isGenerateBillBatchNumber"` // 已生成对账单
	IsInvalid                 []uint8   `gorm:"column:is_invalid" json:"isInvalid"`                                    // 作废 0:正常  1:作废
	IsAudit                   []uint8   `gorm:"column:is_audit" json:"isAudit"`                                        // 审核状态(关联对账单批次号状态)
	AuditUser                 string    `gorm:"column:audit_user" json:"auditUser"`                                    // 审核用户
	AuditUserID               int       `gorm:"column:audit_user_id" json:"auditUserId"`                               // 审核用户ID
	AuditTime                 time.Time `gorm:"column:audit_time" json:"auditTime"`                                    // 审核时间
	IsAccept                  []uint8   `gorm:"column:is_accept" json:"isAccept"`                                      // 已核收(关联收款单状态)
	ReceiptNumber             string    `gorm:"column:receipt_number" json:"receiptNumber"`                            // 收款单号
	AcceptUser                string    `gorm:"column:accept_user" json:"acceptUser"`                                  // 核收用户
	AcceptUserID              int       `gorm:"column:accept_user_id" json:"acceptUserId"`                             // 核收人员id
	AcceptTime                time.Time `gorm:"column:accept_time" json:"acceptTime"`                                  // 核收时间
	IsSyncedOutCumulativeFee  []uint8   `gorm:"column:is_synced_out_cumulative_fee" json:"isSyncedOutCumulativeFee"`   // 是否同步了出库累加金额
	Remark                    string    `gorm:"column:remark" json:"remark"`                                           // 备注
	FinancialRemark           string    `gorm:"column:financial_remark" json:"financialRemark"`                        // 财务备注
	CreateTime                time.Time `gorm:"column:create_time" json:"createTime"`                                  // 创建时间
	CreateUser                int       `gorm:"column:create_user" json:"createUser"`                                  // 生成用户,生成应收人员
	UpdateUser                int       `gorm:"column:update_user" json:"updateUser"`                                  // 更新人
	UpdateTime                time.Time `gorm:"column:update_time" json:"updateTime"`                                  // 更新时间
	Version                   int       `gorm:"column:version" json:"version"`                                         // 乐观锁
	IsDiffAmount              []uint8   `gorm:"column:is_diff_amount" json:"isDiffAmount"`                             // 是否差额
}

// TableName get sql table name.获取数据库表名
func (m *FaCustomerBillFeeDetails) TableName() string {
	return "fa_customer_bill_fee_details"
}

// FaCustomerBillFeeDetailsColumns get sql column name.获取数据库列名
var FaCustomerBillFeeDetailsColumns = struct {
	ID                        string
	OrderNumber               string
	ReferenceNumber           string
	LogisticsNumber           string
	LogisticsNumberFinal      string
	CustomerID                string
	CustomerAlias             string
	CustomerChannelID         string
	CustomerChannelName       string
	ReceiptTime               string
	SendTime                  string
	IsSend                    string
	Weight                    string
	WeighingWeight            string
	FoamWeight                string
	ChargeWeight              string
	FeeTypeName               string
	Receivables               string
	CurrencyCode              string
	CurrencyName              string
	ExchangeRate              string
	ReceivablesRmb            string
	FeeSource                 string
	BillBatchNumber           string
	BillBatchNumberUser       string
	BillBatchNumberUserID     string
	BillBatchNumberTime       string
	IsGenerateBillBatchNumber string
	IsInvalid                 string
	IsAudit                   string
	AuditUser                 string
	AuditUserID               string
	AuditTime                 string
	IsAccept                  string
	ReceiptNumber             string
	AcceptUser                string
	AcceptUserID              string
	AcceptTime                string
	IsSyncedOutCumulativeFee  string
	Remark                    string
	FinancialRemark           string
	CreateTime                string
	CreateUser                string
	UpdateUser                string
	UpdateTime                string
	Version                   string
	IsDiffAmount              string
}{
	ID:                        "id",
	OrderNumber:               "order_number",
	ReferenceNumber:           "reference_number",
	LogisticsNumber:           "logistics_number",
	LogisticsNumberFinal:      "logistics_number_final",
	CustomerID:                "customer_id",
	CustomerAlias:             "customer_alias",
	CustomerChannelID:         "customer_channel_id",
	CustomerChannelName:       "customer_channel_name",
	ReceiptTime:               "receipt_time",
	SendTime:                  "send_time",
	IsSend:                    "is_send",
	Weight:                    "weight",
	WeighingWeight:            "weighing_weight",
	FoamWeight:                "foam_weight",
	ChargeWeight:              "charge_weight",
	FeeTypeName:               "fee_type_name",
	Receivables:               "receivables",
	CurrencyCode:              "currency_code",
	CurrencyName:              "currency_name",
	ExchangeRate:              "exchange_rate",
	ReceivablesRmb:            "receivables_rmb",
	FeeSource:                 "fee_source",
	BillBatchNumber:           "bill_batch_number",
	BillBatchNumberUser:       "bill_batch_number_user",
	BillBatchNumberUserID:     "bill_batch_number_user_id",
	BillBatchNumberTime:       "bill_batch_number_time",
	IsGenerateBillBatchNumber: "is_generate_bill_batch_number",
	IsInvalid:                 "is_invalid",
	IsAudit:                   "is_audit",
	AuditUser:                 "audit_user",
	AuditUserID:               "audit_user_id",
	AuditTime:                 "audit_time",
	IsAccept:                  "is_accept",
	ReceiptNumber:             "receipt_number",
	AcceptUser:                "accept_user",
	AcceptUserID:              "accept_user_id",
	AcceptTime:                "accept_time",
	IsSyncedOutCumulativeFee:  "is_synced_out_cumulative_fee",
	Remark:                    "remark",
	FinancialRemark:           "financial_remark",
	CreateTime:                "create_time",
	CreateUser:                "create_user",
	UpdateUser:                "update_user",
	UpdateTime:                "update_time",
	Version:                   "version",
	IsDiffAmount:              "is_diff_amount",
}

// FaCustomerReceipt 客户收款单
type FaCustomerReceipt struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`                    // 主键
	ReceiptNumber   string    `gorm:"column:receipt_number" json:"receiptNumber"`       // 收款单号
	ParentID        int       `gorm:"column:parent_id" json:"parentId"`                 // 父级iD
	IsParent        []uint8   `gorm:"column:is_parent" json:"isParent"`                 // 是否是父级节点
	InvoiceNumber   string    `gorm:"column:invoice_number" json:"invoiceNumber"`       // 发票号
	VoucherNumber   string    `gorm:"column:voucher_number" json:"voucherNumber"`       // 凭证号
	TicketNumber    string    `gorm:"column:ticket_number" json:"ticketNumber"`         // 票证号
	ReceiptStatus   int       `gorm:"column:receipt_status" json:"receiptStatus"`       // 收款状态
	IsAccept        []uint8   `gorm:"column:is_accept" json:"isAccept"`                 // 是否已收款
	CustomerID      int       `gorm:"column:customer_id" json:"customerId"`             // 客户ID
	CustomerAlias   string    `gorm:"column:customer_alias" json:"customerAlias"`       // 客户别名
	BillFee         float64   `gorm:"column:bill_fee" json:"billFee"`                   // 票据金额
	ActualFee       float64   `gorm:"column:actual_fee" json:"actualFee"`               // 实际收款金额
	CurrencyCode    string    `gorm:"column:currency_code" json:"currencyCode"`         // 币种
	CurrencyName    string    `gorm:"column:currency_name" json:"currencyName"`         // 币种中文名称
	ExchangeRate    float64   `gorm:"column:exchange_rate" json:"exchangeRate"`         // 汇率
	ActualFeeRmb    float64   `gorm:"column:actual_fee_rmb" json:"actualFeeRmb"`        // 实际收款金额rmb
	DealWithUser    string    `gorm:"column:deal_with_user" json:"dealWithUser"`        // 经办人
	BillInputUser   string    `gorm:"column:bill_input_user" json:"billInputUser"`      // 票据录入人员
	BillInputUserID int       `gorm:"column:bill_input_user_id" json:"billInputUserId"` // 票据录入人员ID
	ReceiptType     int       `gorm:"column:receipt_type" json:"receiptType"`           // 收款类别 0:核收 1:清缴
	ReceiptMethod   string    `gorm:"column:receipt_method" json:"receiptMethod"`       // 收款方式
	PaymentTime     time.Time `gorm:"column:payment_time" json:"paymentTime"`           // 收款时间/支付时间
	Mark            string    `gorm:"column:mark" json:"mark"`                          // 标签/标记
	AccountAlias    string    `gorm:"column:account_alias" json:"accountAlias"`         // 账户别称/走账
	AttachmentURL   string    `gorm:"column:attachment_url" json:"attachmentUrl"`       // 附件
	Remark          string    `gorm:"column:remark" json:"remark"`                      // 备注
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`             // 生成用户,生成应收人员
	UpdateUser      int       `gorm:"column:update_user" json:"updateUser"`             // 更新人
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
	Version         int       `gorm:"column:version" json:"version"`                    // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                    // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaCustomerReceipt) TableName() string {
	return "fa_customer_receipt"
}

// FaCustomerReceiptColumns get sql column name.获取数据库列名
var FaCustomerReceiptColumns = struct {
	ID              string
	ReceiptNumber   string
	ParentID        string
	IsParent        string
	InvoiceNumber   string
	VoucherNumber   string
	TicketNumber    string
	ReceiptStatus   string
	IsAccept        string
	CustomerID      string
	CustomerAlias   string
	BillFee         string
	ActualFee       string
	CurrencyCode    string
	CurrencyName    string
	ExchangeRate    string
	ActualFeeRmb    string
	DealWithUser    string
	BillInputUser   string
	BillInputUserID string
	ReceiptType     string
	ReceiptMethod   string
	PaymentTime     string
	Mark            string
	AccountAlias    string
	AttachmentURL   string
	Remark          string
	CreateTime      string
	CreateUser      string
	UpdateUser      string
	UpdateTime      string
	Version         string
	Deleted         string
}{
	ID:              "id",
	ReceiptNumber:   "receipt_number",
	ParentID:        "parent_id",
	IsParent:        "is_parent",
	InvoiceNumber:   "invoice_number",
	VoucherNumber:   "voucher_number",
	TicketNumber:    "ticket_number",
	ReceiptStatus:   "receipt_status",
	IsAccept:        "is_accept",
	CustomerID:      "customer_id",
	CustomerAlias:   "customer_alias",
	BillFee:         "bill_fee",
	ActualFee:       "actual_fee",
	CurrencyCode:    "currency_code",
	CurrencyName:    "currency_name",
	ExchangeRate:    "exchange_rate",
	ActualFeeRmb:    "actual_fee_rmb",
	DealWithUser:    "deal_with_user",
	BillInputUser:   "bill_input_user",
	BillInputUserID: "bill_input_user_id",
	ReceiptType:     "receipt_type",
	ReceiptMethod:   "receipt_method",
	PaymentTime:     "payment_time",
	Mark:            "mark",
	AccountAlias:    "account_alias",
	AttachmentURL:   "attachment_url",
	Remark:          "remark",
	CreateTime:      "create_time",
	CreateUser:      "create_user",
	UpdateUser:      "update_user",
	UpdateTime:      "update_time",
	Version:         "version",
	Deleted:         "deleted",
}

// FaCustomerReceiptDetails 客户收款单明细/流水
type FaCustomerReceiptDetails struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`                             // 主键
	ReceiptNumber        string    `gorm:"column:receipt_number" json:"receiptNumber"`                // 收款单号
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	CustomerID           int       `gorm:"column:customer_id" json:"customerId"`                      // 客户ID
	CustomerAlias        string    `gorm:"column:customer_alias" json:"customerAlias"`                // 客户别名
	CustomerChannelID    int       `gorm:"column:customer_channel_id" json:"customerChannelId"`       // 客户渠道ID
	CustomerChannelName  string    `gorm:"column:customer_channel_name" json:"customerChannelName"`   // 客户渠道名称
	IsAccept             []uint8   `gorm:"column:is_accept" json:"isAccept"`                          // 已核收(关联收款单状态)
	AcceptTime           time.Time `gorm:"column:accept_time" json:"acceptTime"`                      // 核收时间
	AcceptUser           string    `gorm:"column:accept_user" json:"acceptUser"`                      // 核收用户
	AcceptUserID         int       `gorm:"column:accept_user_id" json:"acceptUserId"`                 // 核收用户ID
	Weight               float64   `gorm:"column:weight" json:"weight"`                               // 计费总重量
	ReceivablesRmb       float64   `gorm:"column:receivables_rmb" json:"receivablesRmb"`              // 应收总费用(RMB)
	BeforeBalance        float64   `gorm:"column:before_balance" json:"beforeBalance"`                // 扣款前余额
	AfterBalance         float64   `gorm:"column:after_balance" json:"afterBalance"`                  // 扣款后余额
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 备注
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 生成用户,生成应收人员
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新人
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaCustomerReceiptDetails) TableName() string {
	return "fa_customer_receipt_details"
}

// FaCustomerReceiptDetailsColumns get sql column name.获取数据库列名
var FaCustomerReceiptDetailsColumns = struct {
	ID                   string
	ReceiptNumber        string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	CustomerID           string
	CustomerAlias        string
	CustomerChannelID    string
	CustomerChannelName  string
	IsAccept             string
	AcceptTime           string
	AcceptUser           string
	AcceptUserID         string
	Weight               string
	ReceivablesRmb       string
	BeforeBalance        string
	AfterBalance         string
	Remark               string
	CreateTime           string
	CreateUser           string
	UpdateUser           string
	UpdateTime           string
	Version              string
}{
	ID:                   "id",
	ReceiptNumber:        "receipt_number",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	CustomerID:           "customer_id",
	CustomerAlias:        "customer_alias",
	CustomerChannelID:    "customer_channel_id",
	CustomerChannelName:  "customer_channel_name",
	IsAccept:             "is_accept",
	AcceptTime:           "accept_time",
	AcceptUser:           "accept_user",
	AcceptUserID:         "accept_user_id",
	Weight:               "weight",
	ReceivablesRmb:       "receivables_rmb",
	BeforeBalance:        "before_balance",
	AfterBalance:         "after_balance",
	Remark:               "remark",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateUser:           "update_user",
	UpdateTime:           "update_time",
	Version:              "version",
}

// FaFeeType 费用类型
type FaFeeType struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // 主键
	FeeName    string    `gorm:"column:fee_name" json:"feeName"`       // 费用名称
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	State      []uint8   `gorm:"column:state" json:"state"`            // 状态(0:禁用1:启用)
	FeeType    int       `gorm:"column:fee_type" json:"feeType"`       // 费用种类(1:应收费用,2:应付费用,3:即是应收也是应付)
	FeeState   int       `gorm:"column:fee_state" json:"feeState"`     // 0:收款 1:退件退款 2:部分退款
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建人
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新人
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaFeeType) TableName() string {
	return "fa_fee_type"
}

// FaFeeTypeColumns get sql column name.获取数据库列名
var FaFeeTypeColumns = struct {
	ID         string
	FeeName    string
	Remark     string
	State      string
	FeeType    string
	FeeState   string
	CreateUser string
	CreateTime string
	UpdateUser string
	UpdateTime string
	Version    string
}{
	ID:         "id",
	FeeName:    "fee_name",
	Remark:     "remark",
	State:      "state",
	FeeType:    "fee_type",
	FeeState:   "fee_state",
	CreateUser: "create_user",
	CreateTime: "create_time",
	UpdateUser: "update_user",
	UpdateTime: "update_time",
	Version:    "version",
}

// FaOfferPrice 报价主表
type FaOfferPrice struct {
	ID                  int       `gorm:"primaryKey;column:id" json:"-"`                           // 主键
	Name                string    `gorm:"column:name" json:"name"`                                 // 名称
	OfferType           string    `gorm:"column:offer_type" json:"offerType"`                      // 报价类型 客户报价 渠道报价 成本价等
	PriceType           string    `gorm:"column:price_type" json:"priceType"`                      // 价格类型 公开报价等
	OfferAttr           int       `gorm:"column:offer_attr" json:"offerAttr"`                      // 报价属性 1:应收  2:应付
	Status              []uint8   `gorm:"column:status" json:"status"`                             // 报价表状态:启用/禁用
	CustomerChannelID   int64     `gorm:"column:customer_channel_id" json:"customerChannelId"`     // 关联的客户渠道ID
	CustomerChannelName string    `gorm:"column:customer_channel_name" json:"customerChannelName"` // 客户渠道名称
	CustomerID          int64     `gorm:"column:customer_id" json:"customerId"`                    // 客户ID
	CustomerAlias       string    `gorm:"column:customer_alias" json:"customerAlias"`              // 客户别名
	ProviderChannelCode string    `gorm:"column:provider_channel_code" json:"providerChannelCode"` // 服务商渠道代码
	ProviderChannelName string    `gorm:"column:provider_channel_name" json:"providerChannelName"` // 服务商渠道名称
	Currency            string    `gorm:"column:currency" json:"currency"`                         // 币种
	ExtraCurrency       string    `gorm:"column:extra_currency" json:"extraCurrency"`              // 额外加价币种
	PartitionID         int64     `gorm:"column:partition_id" json:"partitionId"`                  // 关联的分区ID
	PartitionName       string    `gorm:"column:partition_name" json:"partitionName"`              // 关联的分区名称
	HaveSpecialPrice    []uint8   `gorm:"column:have_special_price" json:"haveSpecialPrice"`       // 是否有特殊价格
	BeginTime           time.Time `gorm:"column:begin_time" json:"beginTime"`                      // 开始时间
	EndTime             time.Time `gorm:"column:end_time" json:"endTime"`                          // 结束时间
	Remark              string    `gorm:"column:remark" json:"remark"`                             // 备注
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`                    // 创建时间
	CreateUser          int       `gorm:"column:create_user" json:"createUser"`                    // 创建用户
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`                    // 更新时间
	UpdateUser          int       `gorm:"column:update_user" json:"updateUser"`                    // 更新时间
	Version             int       `gorm:"column:version" json:"version"`                           // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaOfferPrice) TableName() string {
	return "fa_offer_price"
}

// FaOfferPriceColumns get sql column name.获取数据库列名
var FaOfferPriceColumns = struct {
	ID                  string
	Name                string
	OfferType           string
	PriceType           string
	OfferAttr           string
	Status              string
	CustomerChannelID   string
	CustomerChannelName string
	CustomerID          string
	CustomerAlias       string
	ProviderChannelCode string
	ProviderChannelName string
	Currency            string
	ExtraCurrency       string
	PartitionID         string
	PartitionName       string
	HaveSpecialPrice    string
	BeginTime           string
	EndTime             string
	Remark              string
	CreateTime          string
	CreateUser          string
	UpdateTime          string
	UpdateUser          string
	Version             string
}{
	ID:                  "id",
	Name:                "name",
	OfferType:           "offer_type",
	PriceType:           "price_type",
	OfferAttr:           "offer_attr",
	Status:              "status",
	CustomerChannelID:   "customer_channel_id",
	CustomerChannelName: "customer_channel_name",
	CustomerID:          "customer_id",
	CustomerAlias:       "customer_alias",
	ProviderChannelCode: "provider_channel_code",
	ProviderChannelName: "provider_channel_name",
	Currency:            "currency",
	ExtraCurrency:       "extra_currency",
	PartitionID:         "partition_id",
	PartitionName:       "partition_name",
	HaveSpecialPrice:    "have_special_price",
	BeginTime:           "begin_time",
	EndTime:             "end_time",
	Remark:              "remark",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateTime:          "update_time",
	UpdateUser:          "update_user",
	Version:             "version",
}

// FaOfferPriceAffiliate 报价表附表/报价详情
type FaOfferPriceAffiliate struct {
	ID                  int64     `gorm:"primaryKey;column:id" json:"-"`
	OfferPriceID        int64     `gorm:"column:offer_price_id" json:"offerPriceId"`              // 关联主表ID
	PackageType         string    `gorm:"column:package_type" json:"packageType"`                 // 包裹类型
	PartitionIDentifier int       `gorm:"column:partition_identifier" json:"partitionIdentifier"` // 分区标识
	BeginWeight         float64   `gorm:"column:begin_weight" json:"beginWeight"`                 // 重量开始不包含
	EndWeight           float64   `gorm:"column:end_weight" json:"endWeight"`                     // 重量结束:包含
	PriceType           int       `gorm:"column:price_type" json:"priceType"`                     // 价格类型:,1.速查：只要重量在某1个重量段范围内。价格固定不变。比如“首重”就是其中一种例子。,2.总乘：重量在某1个重量段范围，没有起重，直接把公斤数字分割成多少段。这个段数与单价相乘得出的结果。,比如，快递行业常见的“大货价格算法”。按照1公斤为1段。,举例：20公斤到50公斤之间，按照1公斤10元来计算。这里的就是1公斤为1段。,20公斤的价格就是 20x10=200元,21公斤的价格就是 21x10=210元,20.5公斤的其中0.5也要单独分割为1段。不管是20.1公斤，还是20.9公斤，都被进位视为21公斤。价格就是21x10=210元。只要重量在就是20公斤到50公斤之间这个算法。比如，现在FEDEX、DHL等国际大件都是这个算法。,3.增量：以某个公斤数为起重，每增加多少公斤，对应的增加运费。比如“续重”其中1个例子。,比如国际EMS和国内EMS有这个算法。,4.直乘：没有首重。在某个公斤段范围内。把公斤数字和价格单价，直接进行数字相乘得出价格，就是直乘算法。,也没有小数点进位。算出价格是多少就是多少，小数点也算。比如现在的航空小包就是这个算法。
	WeightUnit          float64   `gorm:"column:weight_unit" json:"weightUnit"`                   // 重量单位
	UnitPrice           float64   `gorm:"column:unit_price" json:"unitPrice"`                     // 单价
	ExtraPrice          float64   `gorm:"column:extra_price" json:"extraPrice"`                   // 额外加价
	SumPrice            float64   `gorm:"column:sum_price" json:"sumPrice"`                       // 总和价格
	FirstWeight         float64   `gorm:"column:first_weight" json:"firstWeight"`                 // 首重重量
	FirstWeightPrice    float64   `gorm:"column:first_weight_price" json:"firstWeightPrice"`      // 首重价格
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`                   // 创建时间
	CreateUser          int       `gorm:"column:create_user" json:"createUser"`                   // 创建用户
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`                   // 更新时间
	UpdateUser          int       `gorm:"column:update_user" json:"updateUser"`                   // 更新时间
	Version             int       `gorm:"column:version" json:"version"`                          // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaOfferPriceAffiliate) TableName() string {
	return "fa_offer_price_affiliate"
}

// FaOfferPriceAffiliateColumns get sql column name.获取数据库列名
var FaOfferPriceAffiliateColumns = struct {
	ID                  string
	OfferPriceID        string
	PackageType         string
	PartitionIDentifier string
	BeginWeight         string
	EndWeight           string
	PriceType           string
	WeightUnit          string
	UnitPrice           string
	ExtraPrice          string
	SumPrice            string
	FirstWeight         string
	FirstWeightPrice    string
	CreateTime          string
	CreateUser          string
	UpdateTime          string
	UpdateUser          string
	Version             string
}{
	ID:                  "id",
	OfferPriceID:        "offer_price_id",
	PackageType:         "package_type",
	PartitionIDentifier: "partition_identifier",
	BeginWeight:         "begin_weight",
	EndWeight:           "end_weight",
	PriceType:           "price_type",
	WeightUnit:          "weight_unit",
	UnitPrice:           "unit_price",
	ExtraPrice:          "extra_price",
	SumPrice:            "sum_price",
	FirstWeight:         "first_weight",
	FirstWeightPrice:    "first_weight_price",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateTime:          "update_time",
	UpdateUser:          "update_user",
	Version:             "version",
}

// FaOfferPriceOperateLog 报价表操作日志
type FaOfferPriceOperateLog struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`                    // 主键
	IP               string    `gorm:"column:ip" json:"ip"`                              // ip
	OfferPriceID     int       `gorm:"column:offer_price_id" json:"offerPriceId"`        // 报价表ID
	OperationContent string    `gorm:"column:operation_content" json:"operationContent"` // 操作内容
	Remark           string    `gorm:"column:remark" json:"remark"`                      // 备注
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`             // 创建用户
	OperatorName     string    `gorm:"column:operator_name" json:"operatorName"`         // 操作人名称
	Version          int       `gorm:"column:version" json:"version"`                    // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                    // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaOfferPriceOperateLog) TableName() string {
	return "fa_offer_price_operate_log"
}

// FaOfferPriceOperateLogColumns get sql column name.获取数据库列名
var FaOfferPriceOperateLogColumns = struct {
	ID               string
	IP               string
	OfferPriceID     string
	OperationContent string
	Remark           string
	CreateTime       string
	CreateUser       string
	OperatorName     string
	Version          string
	Deleted          string
}{
	ID:               "id",
	IP:               "ip",
	OfferPriceID:     "offer_price_id",
	OperationContent: "operation_content",
	Remark:           "remark",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	OperatorName:     "operator_name",
	Version:          "version",
	Deleted:          "deleted",
}

// FaOfferPricePartitionAffiliate 报价分区扩展
type FaOfferPricePartitionAffiliate struct {
	ID                  int       `gorm:"primaryKey;column:id" json:"-"`                          // 主键
	OfferPriceID        int       `gorm:"column:offer_price_id" json:"offerPriceId"`              // 关联报价表ID
	PartitionID         int64     `gorm:"column:partition_id" json:"partitionId"`                 // 分区id
	PartitionName       string    `gorm:"column:partition_name" json:"partitionName"`             // 分区名称
	PartitionIDentifier string    `gorm:"column:partition_identifier" json:"partitionIdentifier"` // 分区标识:1区,2区
	Country             string    `gorm:"column:country" json:"country"`                          // 国家
	Province            string    `gorm:"column:province" json:"province"`                        // 省份
	ProvinceExtra       string    `gorm:"column:province_extra" json:"provinceExtra"`             // 省份扩展
	City                string    `gorm:"column:city" json:"city"`                                // 城市
	CityExtra           string    `gorm:"column:city_extra" json:"cityExtra"`                     // 城市扩展
	PostalCode          string    `gorm:"column:postal_code" json:"postalCode"`                   // 邮编
	Remarks             string    `gorm:"column:remarks" json:"remarks"`                          // 备注
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`                   // 创建时间
	CreateUser          int       `gorm:"column:create_user" json:"createUser"`                   // 创建用户
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`                   // 更新时间
	UpdateUser          int       `gorm:"column:update_user" json:"updateUser"`                   // 更新时间
	Version             int       `gorm:"column:version" json:"version"`                          // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaOfferPricePartitionAffiliate) TableName() string {
	return "fa_offer_price_partition_affiliate"
}

// FaOfferPricePartitionAffiliateColumns get sql column name.获取数据库列名
var FaOfferPricePartitionAffiliateColumns = struct {
	ID                  string
	OfferPriceID        string
	PartitionID         string
	PartitionName       string
	PartitionIDentifier string
	Country             string
	Province            string
	ProvinceExtra       string
	City                string
	CityExtra           string
	PostalCode          string
	Remarks             string
	CreateTime          string
	CreateUser          string
	UpdateTime          string
	UpdateUser          string
	Version             string
}{
	ID:                  "id",
	OfferPriceID:        "offer_price_id",
	PartitionID:         "partition_id",
	PartitionName:       "partition_name",
	PartitionIDentifier: "partition_identifier",
	Country:             "country",
	Province:            "province",
	ProvinceExtra:       "province_extra",
	City:                "city",
	CityExtra:           "city_extra",
	PostalCode:          "postal_code",
	Remarks:             "remarks",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateTime:          "update_time",
	UpdateUser:          "update_user",
	Version:             "version",
}

// FaPartition 分区维护主表
type FaPartition struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // 主键
	Name       string    `gorm:"column:name" json:"name"`              // 名称
	State      int       `gorm:"column:state" json:"state"`            // 是否启用 0:未启用 1:启用
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新用户
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaPartition) TableName() string {
	return "fa_partition"
}

// FaPartitionColumns get sql column name.获取数据库列名
var FaPartitionColumns = struct {
	ID         string
	Name       string
	State      string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
}{
	ID:         "id",
	Name:       "name",
	State:      "state",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
}

// FaPartitionAffiliate 分区维护附表
type FaPartitionAffiliate struct {
	ID                  int       `gorm:"primaryKey;column:id" json:"-"`                          // 主键
	PartitionID         int       `gorm:"column:partition_id" json:"partitionId"`                 // 主表分区ID
	PartitionIDentifier string    `gorm:"column:partition_identifier" json:"partitionIdentifier"` // 分区标识:1区,2区
	Country             string    `gorm:"column:country" json:"country"`                          // 国家
	CountryExtra        string    `gorm:"column:country_extra" json:"countryExtra"`               // 国家扩展
	Province            string    `gorm:"column:province" json:"province"`                        // 省份
	ProvinceExtra       string    `gorm:"column:province_extra" json:"provinceExtra"`             // 省份扩展
	City                string    `gorm:"column:city" json:"city"`                                // 城市
	CityExtra           string    `gorm:"column:city_extra" json:"cityExtra"`                     // 城市扩展
	PostalCode          string    `gorm:"column:postal_code" json:"postalCode"`                   // 邮编
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`                   // 创建时间
	CreateUser          int       `gorm:"column:create_user" json:"createUser"`                   // 创建用户
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`                   // 更新时间
	UpdateUser          int       `gorm:"column:update_user" json:"updateUser"`                   // 更新时间
	Version             int       `gorm:"column:version" json:"version"`                          // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *FaPartitionAffiliate) TableName() string {
	return "fa_partition_affiliate"
}

// FaPartitionAffiliateColumns get sql column name.获取数据库列名
var FaPartitionAffiliateColumns = struct {
	ID                  string
	PartitionID         string
	PartitionIDentifier string
	Country             string
	CountryExtra        string
	Province            string
	ProvinceExtra       string
	City                string
	CityExtra           string
	PostalCode          string
	CreateTime          string
	CreateUser          string
	UpdateTime          string
	UpdateUser          string
	Version             string
}{
	ID:                  "id",
	PartitionID:         "partition_id",
	PartitionIDentifier: "partition_identifier",
	Country:             "country",
	CountryExtra:        "country_extra",
	Province:            "province",
	ProvinceExtra:       "province_extra",
	City:                "city",
	CityExtra:           "city_extra",
	PostalCode:          "postal_code",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateTime:          "update_time",
	UpdateUser:          "update_user",
	Version:             "version",
}

// FaPayableBill 应付账单
type FaPayableBill struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`                             // 主键
	BillBatchNumber      string    `gorm:"column:bill_batch_number" json:"billBatchNumber"`           // 账单批次号
	BillBatchStatus      int       `gorm:"column:bill_batch_status" json:"billBatchStatus"`           // 账单状态
	PaymentNumber        string    `gorm:"column:payment_number" json:"paymentNumber"`                // 付款单号
	BusinessEntityCode   string    `gorm:"column:business_entity_code" json:"businessEntityCode"`     // 业务主体代码
	BusinessEntityName   string    `gorm:"column:business_entity_name" json:"businessEntityName"`     // 业务实体名称
	BusinessEntityType   bool      `gorm:"column:business_entity_type" json:"businessEntityType"`     // 业务主体类型，0:服务商，1:航司
	OrderQuantity        int       `gorm:"column:order_quantity" json:"orderQuantity"`                // 订单数量
	FeeQuantity          int       `gorm:"column:fee_quantity" json:"feeQuantity"`                    // 费用条数
	Weight               float64   `gorm:"column:weight" json:"weight"`                               // 计费总重量
	PayableRmb           float64   `gorm:"column:payable_rmb" json:"payableRmb"`                      // 应付总费用(RMB)
	ActualPaymentRmb     float64   `gorm:"column:actual_payment_rmb" json:"actualPaymentRmb"`         // 实际付款金额(RMB)
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 生成时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 生成用户,生成应付人员
	IsAudit              []uint8   `gorm:"column:is_audit" json:"isAudit"`                            // 审核状态(关联对账单批次号状态)
	IsAuditRollback      []uint8   `gorm:"column:is_audit_rollback" json:"isAuditRollback"`           // 是否反审标识 默认false ,被反审后true,重新审核后false
	AuditRollbackUser    string    `gorm:"column:audit_rollback_user" json:"auditRollbackUser"`       // 反审核人员
	AuditRollbackUserID  int       `gorm:"column:audit_rollback_user_id" json:"auditRollbackUserId"`  // 反审核人员ID
	CancelUser           string    `gorm:"column:cancel_user" json:"cancelUser"`                      // 取消应付用户
	CancelUserID         int       `gorm:"column:cancel_user_id" json:"cancelUserId"`                 // 取消应付用户ID
	AuditUser            string    `gorm:"column:audit_user" json:"auditUser"`                        // 审核用户
	AuditUserID          int       `gorm:"column:audit_user_id" json:"auditUserId"`                   // 审核用户
	AuditTime            time.Time `gorm:"column:audit_time" json:"auditTime"`                        // 审核时间
	IsPayment            []uint8   `gorm:"column:is_payment" json:"isPayment"`                        // 已生成付款单
	PaymentNumberUser    string    `gorm:"column:payment_number_user" json:"paymentNumberUser"`       // 付款单生成用户
	PaymentNumberUserID  int       `gorm:"column:payment_number_user_id" json:"paymentNumberUserId"`  // 付款单生成用户ID
	PaymentNumberTime    time.Time `gorm:"column:payment_number_time" json:"paymentNumberTime"`       // 付款单生成用户生成时间
	IsAccept             []uint8   `gorm:"column:is_accept" json:"isAccept"`                          // 已核收(关联付款单状态)
	AcceptTime           time.Time `gorm:"column:accept_time" json:"acceptTime"`                      // 核收时间
	AcceptUser           string    `gorm:"column:accept_user" json:"acceptUser"`                      // 核收用户
	AcceptUserID         int       `gorm:"column:accept_user_id" json:"acceptUserId"`                 // 核收用户ID
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 备注
	FinancialRemark      string    `gorm:"column:financial_remark" json:"financialRemark"`            // 财务备注
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新人
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	Deleted              int       `gorm:"column:deleted" json:"deleted"`                             // 逻辑删除
	PayableRemark        string    `gorm:"column:payable_remark" json:"payableRemark"`                // 核付备注
	ChannelHaul          string    `gorm:"column:channel_haul" json:"channelHaul"`                    // 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
	CustomerChannelNames string    `gorm:"column:customer_channel_names" json:"customerChannelNames"` // 客户渠道名称，多个用,隔开
	KingdeeID            string    `gorm:"column:kingdee_id" json:"kingdeeId"`                        // 金蝶客户内码
	LastSyncTime         time.Time `gorm:"column:last_sync_time" json:"lastSyncTime"`                 // 最后同步时间
	LastSyncUser         int       `gorm:"column:last_sync_user" json:"lastSyncUser"`                 // 最后同步人
}

// TableName get sql table name.获取数据库表名
func (m *FaPayableBill) TableName() string {
	return "fa_payable_bill"
}

// FaPayableBillColumns get sql column name.获取数据库列名
var FaPayableBillColumns = struct {
	ID                   string
	BillBatchNumber      string
	BillBatchStatus      string
	PaymentNumber        string
	BusinessEntityCode   string
	BusinessEntityName   string
	BusinessEntityType   string
	OrderQuantity        string
	FeeQuantity          string
	Weight               string
	PayableRmb           string
	ActualPaymentRmb     string
	CreateTime           string
	CreateUser           string
	IsAudit              string
	IsAuditRollback      string
	AuditRollbackUser    string
	AuditRollbackUserID  string
	CancelUser           string
	CancelUserID         string
	AuditUser            string
	AuditUserID          string
	AuditTime            string
	IsPayment            string
	PaymentNumberUser    string
	PaymentNumberUserID  string
	PaymentNumberTime    string
	IsAccept             string
	AcceptTime           string
	AcceptUser           string
	AcceptUserID         string
	Remark               string
	FinancialRemark      string
	UpdateUser           string
	UpdateTime           string
	Version              string
	Deleted              string
	PayableRemark        string
	ChannelHaul          string
	CustomerChannelNames string
	KingdeeID            string
	LastSyncTime         string
	LastSyncUser         string
}{
	ID:                   "id",
	BillBatchNumber:      "bill_batch_number",
	BillBatchStatus:      "bill_batch_status",
	PaymentNumber:        "payment_number",
	BusinessEntityCode:   "business_entity_code",
	BusinessEntityName:   "business_entity_name",
	BusinessEntityType:   "business_entity_type",
	OrderQuantity:        "order_quantity",
	FeeQuantity:          "fee_quantity",
	Weight:               "weight",
	PayableRmb:           "payable_rmb",
	ActualPaymentRmb:     "actual_payment_rmb",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	IsAudit:              "is_audit",
	IsAuditRollback:      "is_audit_rollback",
	AuditRollbackUser:    "audit_rollback_user",
	AuditRollbackUserID:  "audit_rollback_user_id",
	CancelUser:           "cancel_user",
	CancelUserID:         "cancel_user_id",
	AuditUser:            "audit_user",
	AuditUserID:          "audit_user_id",
	AuditTime:            "audit_time",
	IsPayment:            "is_payment",
	PaymentNumberUser:    "payment_number_user",
	PaymentNumberUserID:  "payment_number_user_id",
	PaymentNumberTime:    "payment_number_time",
	IsAccept:             "is_accept",
	AcceptTime:           "accept_time",
	AcceptUser:           "accept_user",
	AcceptUserID:         "accept_user_id",
	Remark:               "remark",
	FinancialRemark:      "financial_remark",
	UpdateUser:           "update_user",
	UpdateTime:           "update_time",
	Version:              "version",
	Deleted:              "deleted",
	PayableRemark:        "payable_remark",
	ChannelHaul:          "channel_haul",
	CustomerChannelNames: "customer_channel_names",
	KingdeeID:            "kingdee_id",
	LastSyncTime:         "last_sync_time",
	LastSyncUser:         "last_sync_user",
}

// FaPayableBillFee 应付账单费用明细-主表
type FaPayableBillFee struct {
	ID                   int64     `gorm:"primaryKey;column:id" json:"-"`                             // 主键
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	BusinessEntityCode   string    `gorm:"column:business_entity_code" json:"businessEntityCode"`     // 业务主体代码
	BusinessEntityName   string    `gorm:"column:business_entity_name" json:"businessEntityName"`     // 业务主体名称
	BusinessEntityType   bool      `gorm:"column:business_entity_type" json:"businessEntityType"`     // 业务主体类型，0:服务商，1:航司
	ProviderChannelCode  string    `gorm:"column:provider_channel_code" json:"providerChannelCode"`   // 服务商渠道代码
	ReceiptTime          time.Time `gorm:"column:receipt_time" json:"receiptTime"`                    // 入库时间
	SendTime             time.Time `gorm:"column:send_time" json:"sendTime"`                          // 出库时间
	IsSend               []uint8   `gorm:"column:is_send" json:"isSend"`                              // 是否出库
	Weight               float64   `gorm:"column:weight" json:"weight"`                               // 预报重量
	WeighingWeight       float64   `gorm:"column:weighing_weight" json:"weighingWeight"`              // 称重重量
	FoamWeight           float64   `gorm:"column:foam_weight" json:"foamWeight"`                      // 泡重
	ChargeWeight         float64   `gorm:"column:charge_weight" json:"chargeWeight"`                  // 计费重量
	FeeQty               int       `gorm:"column:fee_qty" json:"feeQty"`                              // 费用数量
	PayableRmb           float64   `gorm:"column:payable_rmb" json:"payableRmb"`                      // 应付费用(RMB),折前应付RMB
	BillBatchNumbers     string    `gorm:"column:bill_batch_numbers" json:"billBatchNumbers"`         // 对账单批次号，多个以逗号隔开
	PaymentNumbers       string    `gorm:"column:payment_numbers" json:"paymentNumbers"`              // 付款单号，多个以逗号隔开
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 备注
	FinancialRemark      string    `gorm:"column:financial_remark" json:"financialRemark"`            // 财务备注
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 创建人
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新人
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	CustomerChannelID    int       `gorm:"column:customer_channel_id" json:"customerChannelId"`       // 客户渠道ID
	CustomerChannelName  string    `gorm:"column:customer_channel_name" json:"customerChannelName"`   // 客户渠道名称
	CustomerID           int       `gorm:"column:customer_id" json:"customerId"`                      // 客户ID
	CustomerAlias        string    `gorm:"column:customer_alias" json:"customerAlias"`                // 客户别名
}

// TableName get sql table name.获取数据库表名
func (m *FaPayableBillFee) TableName() string {
	return "fa_payable_bill_fee"
}

// FaPayableBillFeeColumns get sql column name.获取数据库列名
var FaPayableBillFeeColumns = struct {
	ID                   string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	BusinessEntityCode   string
	BusinessEntityName   string
	BusinessEntityType   string
	ProviderChannelCode  string
	ReceiptTime          string
	SendTime             string
	IsSend               string
	Weight               string
	WeighingWeight       string
	FoamWeight           string
	ChargeWeight         string
	FeeQty               string
	PayableRmb           string
	BillBatchNumbers     string
	PaymentNumbers       string
	Remark               string
	FinancialRemark      string
	CreateUser           string
	CreateTime           string
	UpdateUser           string
	UpdateTime           string
	Version              string
	CustomerChannelID    string
	CustomerChannelName  string
	CustomerID           string
	CustomerAlias        string
}{
	ID:                   "id",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	BusinessEntityCode:   "business_entity_code",
	BusinessEntityName:   "business_entity_name",
	BusinessEntityType:   "business_entity_type",
	ProviderChannelCode:  "provider_channel_code",
	ReceiptTime:          "receipt_time",
	SendTime:             "send_time",
	IsSend:               "is_send",
	Weight:               "weight",
	WeighingWeight:       "weighing_weight",
	FoamWeight:           "foam_weight",
	ChargeWeight:         "charge_weight",
	FeeQty:               "fee_qty",
	PayableRmb:           "payable_rmb",
	BillBatchNumbers:     "bill_batch_numbers",
	PaymentNumbers:       "payment_numbers",
	Remark:               "remark",
	FinancialRemark:      "financial_remark",
	CreateUser:           "create_user",
	CreateTime:           "create_time",
	UpdateUser:           "update_user",
	UpdateTime:           "update_time",
	Version:              "version",
	CustomerChannelID:    "customer_channel_id",
	CustomerChannelName:  "customer_channel_name",
	CustomerID:           "customer_id",
	CustomerAlias:        "customer_alias",
}

// FaPayableBillFeeDetails 应付账单费用明细
type FaPayableBillFeeDetails struct {
	ID                        int64     `gorm:"primaryKey;column:id" json:"-"`                                         // 主键
	OrderNumber               string    `gorm:"column:order_number" json:"orderNumber"`                                // 订单号
	ReferenceNumber           string    `gorm:"column:reference_number" json:"referenceNumber"`                        // 参考号
	LogisticsNumber           string    `gorm:"column:logistics_number" json:"logisticsNumber"`                        // 物流单号
	LogisticsNumberFinal      string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"`             // 最终物流单号
	BusinessEntityCode        string    `gorm:"column:business_entity_code" json:"businessEntityCode"`                 // 业务主体代码
	BusinessEntityName        string    `gorm:"column:business_entity_name" json:"businessEntityName"`                 // 业务主体名称
	ProviderChannelCode       string    `gorm:"column:provider_channel_code" json:"providerChannelCode"`               // 服务商渠道代码
	BusinessEntityType        bool      `gorm:"column:business_entity_type" json:"businessEntityType"`                 // 业务主体类型，0:服务商，1:航司
	ReceiptTime               time.Time `gorm:"column:receipt_time" json:"receiptTime"`                                // 入库时间
	SendTime                  time.Time `gorm:"column:send_time" json:"sendTime"`                                      // 出库时间
	IsSend                    []uint8   `gorm:"column:is_send" json:"isSend"`                                          // 是否出库
	Weight                    float64   `gorm:"column:weight" json:"weight"`                                           // 预报重量
	WeighingWeight            float64   `gorm:"column:weighing_weight" json:"weighingWeight"`                          // 称重重量
	FoamWeight                float64   `gorm:"column:foam_weight" json:"foamWeight"`                                  // 泡重
	ChargeWeight              float64   `gorm:"column:charge_weight" json:"chargeWeight"`                              // 计费重量
	FeeTypeName               string    `gorm:"column:fee_type_name" json:"feeTypeName"`                               // 费用类型名称
	Payable                   float64   `gorm:"column:payable" json:"payable"`                                         // 应付费用
	CurrencyCode              string    `gorm:"column:currency_code" json:"currencyCode"`                              // 币种
	CurrencyName              string    `gorm:"column:currency_name" json:"currencyName"`                              // 币种中文名称
	ExchangeRate              float64   `gorm:"column:exchange_rate" json:"exchangeRate"`                              // 汇率
	PayableRmb                float64   `gorm:"column:payable_rmb" json:"payableRmb"`                                  // 应付费用(RMB),折扣前
	FeeSource                 int       `gorm:"column:fee_source" json:"feeSource"`                                    // 费用来源  0:系统生成  1:手工添加
	BillBatchNumber           string    `gorm:"column:bill_batch_number" json:"billBatchNumber"`                       // 对账单批次号
	BillBatchNumberUser       string    `gorm:"column:bill_batch_number_user" json:"billBatchNumberUser"`              // 生成对账用户
	BillBatchNumberUserID     int       `gorm:"column:bill_batch_number_user_id" json:"billBatchNumberUserId"`         // 生成对账用户
	BillBatchNumberTime       time.Time `gorm:"column:bill_batch_number_time" json:"billBatchNumberTime"`              // 生成对账时间
	IsGenerateBillBatchNumber []uint8   `gorm:"column:is_generate_bill_batch_number" json:"isGenerateBillBatchNumber"` // 已生成对账单
	IsInvalid                 []uint8   `gorm:"column:is_invalid" json:"isInvalid"`                                    // 作废 0:正常  1:作废
	IsAudit                   []uint8   `gorm:"column:is_audit" json:"isAudit"`                                        // 审核状态(关联对账单批次号状态)
	AuditUser                 string    `gorm:"column:audit_user" json:"auditUser"`                                    // 审核用户
	AuditUserID               int       `gorm:"column:audit_user_id" json:"auditUserId"`                               // 审核用户ID
	AuditTime                 time.Time `gorm:"column:audit_time" json:"auditTime"`                                    // 审核时间
	IsAccept                  []uint8   `gorm:"column:is_accept" json:"isAccept"`                                      // 已核付(关联付款单状态)
	PaymentNumber             string    `gorm:"column:payment_number" json:"paymentNumber"`                            // 付款单号
	AcceptUser                string    `gorm:"column:accept_user" json:"acceptUser"`                                  // 核付用户
	AcceptUserID              int       `gorm:"column:accept_user_id" json:"acceptUserId"`                             // 核付人员id
	AcceptTime                time.Time `gorm:"column:accept_time" json:"acceptTime"`                                  // 核付时间
	Remark                    string    `gorm:"column:remark" json:"remark"`                                           // 备注
	FinancialRemark           string    `gorm:"column:financial_remark" json:"financialRemark"`                        // 财务备注
	CreateTime                time.Time `gorm:"column:create_time" json:"createTime"`                                  // 创建时间
	CreateUser                int       `gorm:"column:create_user" json:"createUser"`                                  // 生成用户,生成应付人员
	UpdateUser                int       `gorm:"column:update_user" json:"updateUser"`                                  // 更新人
	UpdateTime                time.Time `gorm:"column:update_time" json:"updateTime"`                                  // 更新时间
	Version                   int       `gorm:"column:version" json:"version"`                                         // 乐观锁
	CustomerChannelID         int       `gorm:"column:customer_channel_id" json:"customerChannelId"`                   // 客户渠道ID
	CustomerChannelName       string    `gorm:"column:customer_channel_name" json:"customerChannelName"`               // 客户渠道名称
	DiscountRate              float64   `gorm:"column:discount_rate" json:"discountRate"`                              // 折扣率(小数),默认1 原币种
	ActualPayable             float64   `gorm:"column:actual_payable" json:"actualPayable"`                            // 实际原币种金额
	IsDiffAmount              []uint8   `gorm:"column:is_diff_amount" json:"isDiffAmount"`                             // 是否差额
	CustomerID                int       `gorm:"column:customer_id" json:"customerId"`                                  // 客户ID
	CustomerAlias             string    `gorm:"column:customer_alias" json:"customerAlias"`                            // 客户别名
}

// TableName get sql table name.获取数据库表名
func (m *FaPayableBillFeeDetails) TableName() string {
	return "fa_payable_bill_fee_details"
}

// FaPayableBillFeeDetailsColumns get sql column name.获取数据库列名
var FaPayableBillFeeDetailsColumns = struct {
	ID                        string
	OrderNumber               string
	ReferenceNumber           string
	LogisticsNumber           string
	LogisticsNumberFinal      string
	BusinessEntityCode        string
	BusinessEntityName        string
	ProviderChannelCode       string
	BusinessEntityType        string
	ReceiptTime               string
	SendTime                  string
	IsSend                    string
	Weight                    string
	WeighingWeight            string
	FoamWeight                string
	ChargeWeight              string
	FeeTypeName               string
	Payable                   string
	CurrencyCode              string
	CurrencyName              string
	ExchangeRate              string
	PayableRmb                string
	FeeSource                 string
	BillBatchNumber           string
	BillBatchNumberUser       string
	BillBatchNumberUserID     string
	BillBatchNumberTime       string
	IsGenerateBillBatchNumber string
	IsInvalid                 string
	IsAudit                   string
	AuditUser                 string
	AuditUserID               string
	AuditTime                 string
	IsAccept                  string
	PaymentNumber             string
	AcceptUser                string
	AcceptUserID              string
	AcceptTime                string
	Remark                    string
	FinancialRemark           string
	CreateTime                string
	CreateUser                string
	UpdateUser                string
	UpdateTime                string
	Version                   string
	CustomerChannelID         string
	CustomerChannelName       string
	DiscountRate              string
	ActualPayable             string
	IsDiffAmount              string
	CustomerID                string
	CustomerAlias             string
}{
	ID:                        "id",
	OrderNumber:               "order_number",
	ReferenceNumber:           "reference_number",
	LogisticsNumber:           "logistics_number",
	LogisticsNumberFinal:      "logistics_number_final",
	BusinessEntityCode:        "business_entity_code",
	BusinessEntityName:        "business_entity_name",
	ProviderChannelCode:       "provider_channel_code",
	BusinessEntityType:        "business_entity_type",
	ReceiptTime:               "receipt_time",
	SendTime:                  "send_time",
	IsSend:                    "is_send",
	Weight:                    "weight",
	WeighingWeight:            "weighing_weight",
	FoamWeight:                "foam_weight",
	ChargeWeight:              "charge_weight",
	FeeTypeName:               "fee_type_name",
	Payable:                   "payable",
	CurrencyCode:              "currency_code",
	CurrencyName:              "currency_name",
	ExchangeRate:              "exchange_rate",
	PayableRmb:                "payable_rmb",
	FeeSource:                 "fee_source",
	BillBatchNumber:           "bill_batch_number",
	BillBatchNumberUser:       "bill_batch_number_user",
	BillBatchNumberUserID:     "bill_batch_number_user_id",
	BillBatchNumberTime:       "bill_batch_number_time",
	IsGenerateBillBatchNumber: "is_generate_bill_batch_number",
	IsInvalid:                 "is_invalid",
	IsAudit:                   "is_audit",
	AuditUser:                 "audit_user",
	AuditUserID:               "audit_user_id",
	AuditTime:                 "audit_time",
	IsAccept:                  "is_accept",
	PaymentNumber:             "payment_number",
	AcceptUser:                "accept_user",
	AcceptUserID:              "accept_user_id",
	AcceptTime:                "accept_time",
	Remark:                    "remark",
	FinancialRemark:           "financial_remark",
	CreateTime:                "create_time",
	CreateUser:                "create_user",
	UpdateUser:                "update_user",
	UpdateTime:                "update_time",
	Version:                   "version",
	CustomerChannelID:         "customer_channel_id",
	CustomerChannelName:       "customer_channel_name",
	DiscountRate:              "discount_rate",
	ActualPayable:             "actual_payable",
	IsDiffAmount:              "is_diff_amount",
	CustomerID:                "customer_id",
	CustomerAlias:             "customer_alias",
}

// FaPayableCalculatePriceRecord 应付重新核算价格记录
type FaPayableCalculatePriceRecord struct {
	ID                   int64     `gorm:"primaryKey;column:id" json:"-"`                             // 主键
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	UpdateField          string    `gorm:"column:update_field" json:"updateField"`                    // 修改字段
	OldValue             string    `gorm:"column:old_value" json:"oldValue"`                          // 操作前值
	NewValue             string    `gorm:"column:new_value" json:"newValue"`                          // 操作后值
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 操作时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 操作人
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新人
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	Deleted              int       `gorm:"column:deleted" json:"deleted"`                             // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaPayableCalculatePriceRecord) TableName() string {
	return "fa_payable_calculate_price_record"
}

// FaPayableCalculatePriceRecordColumns get sql column name.获取数据库列名
var FaPayableCalculatePriceRecordColumns = struct {
	ID                   string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	UpdateField          string
	OldValue             string
	NewValue             string
	CreateTime           string
	CreateUser           string
	UpdateUser           string
	UpdateTime           string
	Version              string
	Deleted              string
}{
	ID:                   "id",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	UpdateField:          "update_field",
	OldValue:             "old_value",
	NewValue:             "new_value",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateUser:           "update_user",
	UpdateTime:           "update_time",
	Version:              "version",
	Deleted:              "deleted",
}

// FaPaymentOrderInfo 支付订单流水表
type FaPaymentOrderInfo struct {
	ID                  int            `gorm:"primaryKey;column:id" json:"-"`                           // ID
	SystemFlowingNumber string         `gorm:"column:system_flowing_number" json:"systemFlowingNumber"` // 系统流水号
	OrderFlowingNumber  string         `gorm:"column:order_flowing_number" json:"orderFlowingNumber"`   // 订单流水号
	PayAmount           float64        `gorm:"column:pay_amount" json:"payAmount"`                      // 支付金额
	PayType             int            `gorm:"column:pay_type" json:"payType"`                          // 支付方式（0-微信支付，1-支付宝支付，2-线下支付）
	PayStatus           int            `gorm:"column:pay_status" json:"payStatus"`                      // 支付状态（0-支付中，1-支付完成，2-转入退款，3-未支付，4-已关闭，5-已撤销（付款码支付），6-用户支付中（付款码支付），7-支付失败(其他原因，如银行返回失败)）
	PayTime             time.Time      `gorm:"column:pay_time" json:"payTime"`                          // 支付完成时间
	ThirdPartyNumber    string         `gorm:"column:third_party_number" json:"thirdPartyNumber"`       // 第三方单号
	MerchantNumber      string         `gorm:"column:merchant_number" json:"merchantNumber"`            // 商户订单号
	CustomerID          int            `gorm:"column:customer_id" json:"customerId"`                    // 客户id
	CustomerName        string         `gorm:"column:customer_name" json:"customerName"`                // 客户名称
	RequestSource       string         `gorm:"column:request_source" json:"requestSource"`              // 请求来源（0：内部，1：外部）
	CallbackStatus      int            `gorm:"column:callback_status" json:"callbackStatus"`            // 回调状态（0：未回调，1：已回调）
	Extra               datatypes.JSON `gorm:"column:extra" json:"extra"`                               // 扩展字段
	Remarks             string         `gorm:"column:remarks" json:"remarks"`                           // 备注信息
	CreateUser          int            `gorm:"column:create_user" json:"createUser"`                    // 创建用户
	CreateTime          time.Time      `gorm:"column:create_time" json:"createTime"`                    // 创建时间
	UpdateUser          int            `gorm:"column:update_user" json:"updateUser"`                    // 更新时间
	UpdateTime          time.Time      `gorm:"column:update_time" json:"updateTime"`                    // 更新时间
	Version             int            `gorm:"column:version" json:"version"`                           // 乐观锁
	IsDelete            int            `gorm:"column:is_delete" json:"isDelete"`                        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaPaymentOrderInfo) TableName() string {
	return "fa_payment_order_info"
}

// FaPaymentOrderInfoColumns get sql column name.获取数据库列名
var FaPaymentOrderInfoColumns = struct {
	ID                  string
	SystemFlowingNumber string
	OrderFlowingNumber  string
	PayAmount           string
	PayType             string
	PayStatus           string
	PayTime             string
	ThirdPartyNumber    string
	MerchantNumber      string
	CustomerID          string
	CustomerName        string
	RequestSource       string
	CallbackStatus      string
	Extra               string
	Remarks             string
	CreateUser          string
	CreateTime          string
	UpdateUser          string
	UpdateTime          string
	Version             string
	IsDelete            string
}{
	ID:                  "id",
	SystemFlowingNumber: "system_flowing_number",
	OrderFlowingNumber:  "order_flowing_number",
	PayAmount:           "pay_amount",
	PayType:             "pay_type",
	PayStatus:           "pay_status",
	PayTime:             "pay_time",
	ThirdPartyNumber:    "third_party_number",
	MerchantNumber:      "merchant_number",
	CustomerID:          "customer_id",
	CustomerName:        "customer_name",
	RequestSource:       "request_source",
	CallbackStatus:      "callback_status",
	Extra:               "extra",
	Remarks:             "remarks",
	CreateUser:          "create_user",
	CreateTime:          "create_time",
	UpdateUser:          "update_user",
	UpdateTime:          "update_time",
	Version:             "version",
	IsDelete:            "is_delete",
}

// FaReCalculatePriceRecord 财务重新核算价格记录
type FaReCalculatePriceRecord struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`                             // 主键
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	UpdateField          string    `gorm:"column:update_field" json:"updateField"`                    // 修改字段
	OldValue             string    `gorm:"column:old_value" json:"oldValue"`                          // 操作前值
	NewValue             string    `gorm:"column:new_value" json:"newValue"`                          // 操作后值
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 操作时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 操作人
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新人
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	Deleted              int       `gorm:"column:deleted" json:"deleted"`                             // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaReCalculatePriceRecord) TableName() string {
	return "fa_re_calculate_price_record"
}

// FaReCalculatePriceRecordColumns get sql column name.获取数据库列名
var FaReCalculatePriceRecordColumns = struct {
	ID                   string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	UpdateField          string
	OldValue             string
	NewValue             string
	CreateTime           string
	CreateUser           string
	UpdateUser           string
	UpdateTime           string
	Version              string
	Deleted              string
}{
	ID:                   "id",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	UpdateField:          "update_field",
	OldValue:             "old_value",
	NewValue:             "new_value",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateUser:           "update_user",
	UpdateTime:           "update_time",
	Version:              "version",
	Deleted:              "deleted",
}

// FaSettlementMethod 结算方式
type FaSettlementMethod struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // 主键id
	Method     string    `gorm:"column:method" json:"method"`          // 结算方式
	State      []uint8   `gorm:"column:state" json:"state"`            // 状态（0:禁用1:启用）
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建人
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新人
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *FaSettlementMethod) TableName() string {
	return "fa_settlement_method"
}

// FaSettlementMethodColumns get sql column name.获取数据库列名
var FaSettlementMethodColumns = struct {
	ID         string
	Method     string
	State      string
	Remark     string
	CreateUser string
	CreateTime string
	UpdateUser string
	UpdateTime string
	Version    string
	Deleted    string
}{
	ID:         "id",
	Method:     "method",
	State:      "state",
	Remark:     "remark",
	CreateUser: "create_user",
	CreateTime: "create_time",
	UpdateUser: "update_user",
	UpdateTime: "update_time",
	Version:    "version",
	Deleted:    "deleted",
}

// FaStatistical 财务统计报表
type FaStatistical struct {
	Type       int            `gorm:"column:type" json:"type"`              // 报表类型，1:客户渠道毛利计算报表
	DateType   int            `gorm:"column:date_type" json:"dateType"`     // 日期类型，1:订单创建时间、2:应收费用明细创建时间、3:应付费用明细创建时间、4:入库时间、5:出库时间
	Date       datatypes.Date `gorm:"column:date" json:"date"`              // 日期
	Data       datatypes.JSON `gorm:"column:data" json:"data"`              // 统计数据
	CreateTime time.Time      `gorm:"column:create_time" json:"createTime"` // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *FaStatistical) TableName() string {
	return "fa_statistical"
}

// FaStatisticalColumns get sql column name.获取数据库列名
var FaStatisticalColumns = struct {
	Type       string
	DateType   string
	Date       string
	Data       string
	CreateTime string
}{
	Type:       "type",
	DateType:   "date_type",
	Date:       "date",
	Data:       "data",
	CreateTime: "create_time",
}

// LgAddressBook 收寄地址管理
type LgAddressBook struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`                  // ID
	CustomerID      int       `gorm:"column:customer_id" json:"customerId"`           // 客户id
	AddressName     string    `gorm:"column:address_name" json:"addressName"`         // 地址簿名称
	ApplicationType string    `gorm:"column:application_type" json:"applicationType"` // 应用类型，customer:客户，provider_channel:服务商渠道
	Type            string    `gorm:"column:type" json:"type"`                        // 类型 收件人/发件人
	Country         string    `gorm:"column:country" json:"country"`                  // 国家二字码
	AreaID          int       `gorm:"column:area_id" json:"areaId"`                   // 地区id
	Province        string    `gorm:"column:province" json:"province"`                // 省份
	City            string    `gorm:"column:city" json:"city"`                        // 城市
	Phone           string    `gorm:"column:phone" json:"phone"`                      // 电话
	Name            string    `gorm:"column:name" json:"name"`                        // 姓名
	Zipcode         string    `gorm:"column:zipcode" json:"zipcode"`                  // 邮编
	Area            string    `gorm:"column:area" json:"area"`                        // 区/县
	Street          string    `gorm:"column:street" json:"street"`                    // 街道
	HouseNumber     string    `gorm:"column:house_number" json:"houseNumber"`         // 门牌号
	CertificateType string    `gorm:"column:certificate_type" json:"certificateType"` // 证件类型
	CertificateCode string    `gorm:"column:certificate_code" json:"certificateCode"` // 证件号
	Mail            string    `gorm:"column:mail" json:"mail"`                        // 邮箱
	Remark          string    `gorm:"column:remark" json:"remark"`                    // 备注
	Mobile          string    `gorm:"column:mobile" json:"mobile"`                    // 手机
	Tax             string    `gorm:"column:tax" json:"tax"`                          // 税号
	Address         string    `gorm:"column:address" json:"address"`                  // 地址
	Address2        string    `gorm:"column:address_2" json:"address2"`               // 地址2
	Company         string    `gorm:"column:company" json:"company"`                  // 公司
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`           // 创建用户
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`           // 更新时间
	UpdateUser      int       `gorm:"column:update_user" json:"updateUser"`           // 更新时间
	Version         int       `gorm:"column:version" json:"version"`                  // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                  // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgAddressBook) TableName() string {
	return "lg_address_book"
}

// LgAddressBookColumns get sql column name.获取数据库列名
var LgAddressBookColumns = struct {
	ID              string
	CustomerID      string
	AddressName     string
	ApplicationType string
	Type            string
	Country         string
	AreaID          string
	Province        string
	City            string
	Phone           string
	Name            string
	Zipcode         string
	Area            string
	Street          string
	HouseNumber     string
	CertificateType string
	CertificateCode string
	Mail            string
	Remark          string
	Mobile          string
	Tax             string
	Address         string
	Address2        string
	Company         string
	CreateTime      string
	CreateUser      string
	UpdateTime      string
	UpdateUser      string
	Version         string
	Deleted         string
}{
	ID:              "id",
	CustomerID:      "customer_id",
	AddressName:     "address_name",
	ApplicationType: "application_type",
	Type:            "type",
	Country:         "country",
	AreaID:          "area_id",
	Province:        "province",
	City:            "city",
	Phone:           "phone",
	Name:            "name",
	Zipcode:         "zipcode",
	Area:            "area",
	Street:          "street",
	HouseNumber:     "house_number",
	CertificateType: "certificate_type",
	CertificateCode: "certificate_code",
	Mail:            "mail",
	Remark:          "remark",
	Mobile:          "mobile",
	Tax:             "tax",
	Address:         "address",
	Address2:        "address_2",
	Company:         "company",
	CreateTime:      "create_time",
	CreateUser:      "create_user",
	UpdateTime:      "update_time",
	UpdateUser:      "update_user",
	Version:         "version",
	Deleted:         "deleted",
}

// LgArea 地区总表
type LgArea struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`                                                // ID
	ParentID       int       `gorm:"column:parent_id" json:"parentId"`                                             // 父ID
	LgArea         *LgArea   `gorm:"joinForeignKey:parent_id;foreignKey:id;references:ParentID" json:"lgAreaList"` // 地区总表
	RegionalID     string    `gorm:"column:regional_id" json:"regionalId"`                                         // 大区id
	RegionalNameCn string    `gorm:"column:regional_name_cn" json:"regionalNameCn"`                                // 大区中文名称
	Level          int       `gorm:"column:level" json:"level"`                                                    // 等级，0: 国家，1：省份，2：城市
	TwoCode        string    `gorm:"column:two_code" json:"twoCode"`                                               // 二字码
	ThreeCode      string    `gorm:"column:three_code" json:"threeCode"`                                           // 三字码
	NameEn         string    `gorm:"column:name_en" json:"nameEn"`                                                 // 英文名
	NameCn         string    `gorm:"column:name_cn" json:"nameCn"`                                                 // 中文名
	ZipCode        string    `gorm:"column:zip_code" json:"zipCode"`                                               // 邮政编码 level等级为2时 邮编不能为空
	Extra          string    `gorm:"column:extra" json:"extra"`                                                    // 扩展字段(匹配正确省市)
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`                                         // 创建时间
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`                                         // 创建用户
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`                                         // 更新时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`                                         // 更新时间
	Version        int       `gorm:"column:version" json:"version"`                                                // 乐观锁
	Deleted        int       `gorm:"column:deleted" json:"deleted"`                                                // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgArea) TableName() string {
	return "lg_area"
}

// LgAreaColumns get sql column name.获取数据库列名
var LgAreaColumns = struct {
	ID             string
	ParentID       string
	RegionalID     string
	RegionalNameCn string
	Level          string
	TwoCode        string
	ThreeCode      string
	NameEn         string
	NameCn         string
	ZipCode        string
	Extra          string
	CreateTime     string
	CreateUser     string
	UpdateTime     string
	UpdateUser     string
	Version        string
	Deleted        string
}{
	ID:             "id",
	ParentID:       "parent_id",
	RegionalID:     "regional_id",
	RegionalNameCn: "regional_name_cn",
	Level:          "level",
	TwoCode:        "two_code",
	ThreeCode:      "three_code",
	NameEn:         "name_en",
	NameCn:         "name_cn",
	ZipCode:        "zip_code",
	Extra:          "extra",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	Version:        "version",
	Deleted:        "deleted",
}

// LgChannel 渠道表(废弃)
type LgChannel struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`             // ID
	ChannelID    int       `gorm:"column:channel_id" json:"channelId"`        // 渠道ID
	ChannelNameA string    `gorm:"column:channel_name_a" json:"channelNameA"` // 渠道名称
	ChannelNameB string    `gorm:"column:channel_name_b" json:"channelNameB"` // 客户渠道名称
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`      // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`      // 创建用户
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`      // 更新时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`      // 更新时间
	Version      int       `gorm:"column:version" json:"version"`             // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`             // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgChannel) TableName() string {
	return "lg_channel"
}

// LgChannelColumns get sql column name.获取数据库列名
var LgChannelColumns = struct {
	ID           string
	ChannelID    string
	ChannelNameA string
	ChannelNameB string
	CreateTime   string
	CreateUser   string
	UpdateTime   string
	UpdateUser   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	ChannelID:    "channel_id",
	ChannelNameA: "channel_name_a",
	ChannelNameB: "channel_name_b",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateTime:   "update_time",
	UpdateUser:   "update_user",
	Version:      "version",
	Deleted:      "deleted",
}

// LgChannelCountry 渠道_国家
type LgChannelCountry struct {
	ChannelID int `gorm:"column:channel_id" json:"channelId"` // 渠道ID
	CountryID int `gorm:"column:country_id" json:"countryId"` // 国家表ID
}

// TableName get sql table name.获取数据库表名
func (m *LgChannelCountry) TableName() string {
	return "lg_channel_country"
}

// LgChannelCountryColumns get sql column name.获取数据库列名
var LgChannelCountryColumns = struct {
	ChannelID string
	CountryID string
}{
	ChannelID: "channel_id",
	CountryID: "country_id",
}

// LgChannelFollow 渠道跟进
type LgChannelFollow struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`          // ID
	CustomerID  int       `gorm:"column:customer_id" json:"customerId"`   // 客户ID
	ChannelType string    `gorm:"column:channel_type" json:"channelType"` // 渠道类型
	ChannelID   int       `gorm:"column:channel_id" json:"channelId"`     // 渠道id
	Content     string    `gorm:"column:content" json:"content"`          // 留言内容
	Status      int       `gorm:"column:status" json:"status"`            // 状态，0:无效，1:有效
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`   // 创建时间
	CreateUser  int       `gorm:"column:create_user" json:"createUser"`   // 创建用户
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`   // 更新时间
	UpdateUser  int       `gorm:"column:update_user" json:"updateUser"`   // 更新时间
	Version     int       `gorm:"column:version" json:"version"`          // 乐观锁
	Deleted     int       `gorm:"column:deleted" json:"deleted"`          // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgChannelFollow) TableName() string {
	return "lg_channel_follow"
}

// LgChannelFollowColumns get sql column name.获取数据库列名
var LgChannelFollowColumns = struct {
	ID          string
	CustomerID  string
	ChannelType string
	ChannelID   string
	Content     string
	Status      string
	CreateTime  string
	CreateUser  string
	UpdateTime  string
	UpdateUser  string
	Version     string
	Deleted     string
}{
	ID:          "id",
	CustomerID:  "customer_id",
	ChannelType: "channel_type",
	ChannelID:   "channel_id",
	Content:     "content",
	Status:      "status",
	CreateTime:  "create_time",
	CreateUser:  "create_user",
	UpdateTime:  "update_time",
	UpdateUser:  "update_user",
	Version:     "version",
	Deleted:     "deleted",
}

// LgChannelNumberRules 渠道_号码_规则
type LgChannelNumberRules struct {
	ChannelID   int    `gorm:"primaryKey;column:channel_id" json:"-"`  // 渠道ID
	NumberRules string `gorm:"column:number_rules" json:"numberRules"` // 单号规则 | 隔开
}

// TableName get sql table name.获取数据库表名
func (m *LgChannelNumberRules) TableName() string {
	return "lg_channel_number_rules"
}

// LgChannelNumberRulesColumns get sql column name.获取数据库列名
var LgChannelNumberRulesColumns = struct {
	ChannelID   string
	NumberRules string
}{
	ChannelID:   "channel_id",
	NumberRules: "number_rules",
}

// LgChannelOperateLog 渠道管理操作日志
type LgChannelOperateLog struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`                        // 主键id
	OperateModule      string    `gorm:"column:operate_module" json:"operateModule"`           // 操作模块
	CheckCode          string    `gorm:"column:check_code" json:"checkCode"`                   // 渠道代码/系统服务商代码/渠道id
	OperateType        string    `gorm:"column:operate_type" json:"operateType"`               // 操作类型
	OperateDescription string    `gorm:"column:operate_description" json:"operateDescription"` // 操作描述
	UpdateParams       string    `gorm:"column:update_params" json:"updateParams"`             // 更新参数
	OperateIP          string    `gorm:"column:operate_ip" json:"operateIp"`                   // 操作IP
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                 // 操作人id
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                 // 操作时间
	Version            int       `gorm:"column:version" json:"version"`                        // 乐观锁
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgChannelOperateLog) TableName() string {
	return "lg_channel_operate_log"
}

// LgChannelOperateLogColumns get sql column name.获取数据库列名
var LgChannelOperateLogColumns = struct {
	ID                 string
	OperateModule      string
	CheckCode          string
	OperateType        string
	OperateDescription string
	UpdateParams       string
	OperateIP          string
	CreateUser         string
	CreateTime         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	OperateModule:      "operate_module",
	CheckCode:          "check_code",
	OperateType:        "operate_type",
	OperateDescription: "operate_description",
	UpdateParams:       "update_params",
	OperateIP:          "operate_ip",
	CreateUser:         "create_user",
	CreateTime:         "create_time",
	Version:            "version",
	Deleted:            "deleted",
}

// LgCountry 国家表(废弃)
type LgCountry struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // ID
	NameEn     string    `gorm:"column:name_en" json:"nameEn"`         // 国家中文
	NameCn     string    `gorm:"column:name_cn" json:"nameCn"`         // 国家英文
	TwoCode    string    `gorm:"column:two_code" json:"twoCode"`       // 二字代码
	ThreeCode  string    `gorm:"column:three_code" json:"threeCode"`   // 三字代码
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgCountry) TableName() string {
	return "lg_country"
}

// LgCountryColumns get sql column name.获取数据库列名
var LgCountryColumns = struct {
	ID         string
	NameEn     string
	NameCn     string
	TwoCode    string
	ThreeCode  string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	NameEn:     "name_en",
	NameCn:     "name_cn",
	TwoCode:    "two_code",
	ThreeCode:  "three_code",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// LgCreateOrderLog 下单日志
type LgCreateOrderLog struct {
	ID                    int       `gorm:"primaryKey;column:id" json:"-"`                               // 主键id
	ClientID              int       `gorm:"column:client_id" json:"clientId"`                            // 客户id
	ClientName            string    `gorm:"column:client_name" json:"clientName"`                        // 客户名
	OrderNum              string    `gorm:"column:order_num" json:"orderNum"`                            // 订单号
	Success               []uint8   `gorm:"column:success" json:"success"`                               // 是否下单成功 0 失败，1 成功
	RequestParams         string    `gorm:"column:request_params" json:"requestParams"`                  // 请求报文
	OriginalRequestParams string    `gorm:"column:original_request_params" json:"originalRequestParams"` // 原始请求报文
	ResponseParams        string    `gorm:"column:response_params" json:"responseParams"`                // 响应报文
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`                        // 创建时间
	CreateUser            int       `gorm:"column:create_user" json:"createUser"`                        // 操作人
	Version               int       `gorm:"column:version" json:"version"`                               // 乐观锁
	Deleted               int       `gorm:"column:deleted" json:"deleted"`                               // 逻辑删除
	LogisticsNumber       string    `gorm:"column:logistics_number" json:"logisticsNumber"`              // 物流单号
	CustomerChannelID     int       `gorm:"column:customer_channel_id" json:"customerChannelId"`         // 客户渠道ID
	FailureMessage        string    `gorm:"column:failure_message" json:"failureMessage"`                // 失败原因
}

// TableName get sql table name.获取数据库表名
func (m *LgCreateOrderLog) TableName() string {
	return "lg_create_order_log"
}

// LgCreateOrderLogColumns get sql column name.获取数据库列名
var LgCreateOrderLogColumns = struct {
	ID                    string
	ClientID              string
	ClientName            string
	OrderNum              string
	Success               string
	RequestParams         string
	OriginalRequestParams string
	ResponseParams        string
	CreateTime            string
	CreateUser            string
	Version               string
	Deleted               string
	LogisticsNumber       string
	CustomerChannelID     string
	FailureMessage        string
}{
	ID:                    "id",
	ClientID:              "client_id",
	ClientName:            "client_name",
	OrderNum:              "order_num",
	Success:               "success",
	RequestParams:         "request_params",
	OriginalRequestParams: "original_request_params",
	ResponseParams:        "response_params",
	CreateTime:            "create_time",
	CreateUser:            "create_user",
	Version:               "version",
	Deleted:               "deleted",
	LogisticsNumber:       "logistics_number",
	CustomerChannelID:     "customer_channel_id",
	FailureMessage:        "failure_message",
}

// LgCustomerChannel 客户渠道
type LgCustomerChannel struct {
	ID                          int       `gorm:"primaryKey;column:id" json:"-"`                                              // 渠道ID
	CustomerChannelType         int       `gorm:"column:customer_channel_type" json:"customerChannelType"`                    // 渠道类型
	CustomerChannelName         string    `gorm:"column:customer_channel_name" json:"customerChannelName"`                    // 渠道名称
	MinDeclare                  float64   `gorm:"column:min_declare" json:"minDeclare"`                                       // 最低申报价值
	MaxDeclare                  float64   `gorm:"column:max_declare" json:"maxDeclare"`                                       // 最高申报价值
	MinWeight                   float64   `gorm:"column:min_weight" json:"minWeight"`                                         // 最低下单重量
	MaxWeight                   float64   `gorm:"column:max_weight" json:"maxWeight"`                                         // 最高下单重量
	MinSendWeight               float64   `gorm:"column:min_send_weight" json:"minSendWeight"`                                // 最小出库重量
	MaxSendWeight               float64   `gorm:"column:max_send_weight" json:"maxSendWeight"`                                // 最大出库重量
	InStorageUseOrderWeightFlag bool      `gorm:"column:in_storage_use_order_weight_flag" json:"inStorageUseOrderWeightFlag"` // 入库是否使用订单预报重量，0:false，1:true，默认0
	InStorageOrderWeightMin     float64   `gorm:"column:in_storage_order_weight_min" json:"inStorageOrderWeightMin"`          // 入库最小预报重量，单位:KG
	InStorageOrderWeightMax     float64   `gorm:"column:in_storage_order_weight_max" json:"inStorageOrderWeightMax"`          // 入库最大预报重量，单位:KG
	SenderBlackList             string    `gorm:"column:sender_black_list" json:"senderBlackList"`                            // 寄件人黑名单
	BanGoodsList                string    `gorm:"column:ban_goods_list" json:"banGoodsList"`                                  // 禁止商品
	NotePublic                  string    `gorm:"column:note_public" json:"notePublic"`                                       // 客户渠道说明_公开
	NotePrivate                 string    `gorm:"column:note_private" json:"notePrivate"`                                     // 客户渠道说明_私有
	OrderNumPoolID              int64     `gorm:"column:order_num_pool_id" json:"orderNumPoolId"`                             // 绑定的单号池ID
	OrderNumPoolName            string    `gorm:"column:order_num_pool_name" json:"orderNumPoolName"`                         // 绑定的单号池名称
	ExchangeOrder               []uint8   `gorm:"column:exchange_order" json:"exchangeOrder"`                                 // 是否需要换单0 不需要1 需要
	ExchangeType                int       `gorm:"column:exchange_type" json:"exchangeType"`                                   // 换单类型 1:只换面单不换号 2:即换单又换号
	LabelCode                   string    `gorm:"column:label_code" json:"labelCode"`                                         // 渠道为换单渠道则选取制定的面单code
	DataPushAPI                 string    `gorm:"column:data_push_api" json:"dataPushApi"`                                    // 订单数据推送地址
	Proportion                  int       `gorm:"column:proportion" json:"proportion"`                                        // 抛比
	CreateTime                  time.Time `gorm:"column:create_time" json:"createTime"`                                       // 创建时间
	CreateUser                  int       `gorm:"column:create_user" json:"createUser"`                                       // 创建用户
	UpdateTime                  time.Time `gorm:"column:update_time" json:"updateTime"`                                       // 更新时间
	UpdateUser                  int       `gorm:"column:update_user" json:"updateUser"`                                       // 更新时间
	Version                     int       `gorm:"column:version" json:"version"`                                              // 乐观锁
	Deleted                     int       `gorm:"column:deleted" json:"deleted"`                                              // 逻辑删除
	Status                      int       `gorm:"column:status" json:"status"`                                                // 状态，0:无效，1:有效
	RequiredPlatformNo          []uint8   `gorm:"column:required_platform_no" json:"requiredPlatformNo"`                      // 是否必填平台订单号，0-不必填，1-必填
	ChannelHaul                 string    `gorm:"column:channel_haul" json:"channelHaul"`                                     // 渠道运程（YC001：全程；YC002：尾程；YC003：其他）
	KingdeeID                   string    `gorm:"column:kingdee_id" json:"kingdeeId"`                                         // 金蝶客户内码
	CustomerChannelTypeID       string    `gorm:"column:customer_channel_type_id" json:"customerChannelTypeId"`               // 客户渠道类型id
	ChannelCountryAffiliationID int       `gorm:"column:channel_country_affiliation_id" json:"channelCountryAffiliationId"`   // 渠道所属国家id,lg_area，对应lg_area表的id（注意：只有level 为0 的才是国家）
}

// TableName get sql table name.获取数据库表名
func (m *LgCustomerChannel) TableName() string {
	return "lg_customer_channel"
}

// LgCustomerChannelColumns get sql column name.获取数据库列名
var LgCustomerChannelColumns = struct {
	ID                          string
	CustomerChannelType         string
	CustomerChannelName         string
	MinDeclare                  string
	MaxDeclare                  string
	MinWeight                   string
	MaxWeight                   string
	MinSendWeight               string
	MaxSendWeight               string
	InStorageUseOrderWeightFlag string
	InStorageOrderWeightMin     string
	InStorageOrderWeightMax     string
	SenderBlackList             string
	BanGoodsList                string
	NotePublic                  string
	NotePrivate                 string
	OrderNumPoolID              string
	OrderNumPoolName            string
	ExchangeOrder               string
	ExchangeType                string
	LabelCode                   string
	DataPushAPI                 string
	Proportion                  string
	CreateTime                  string
	CreateUser                  string
	UpdateTime                  string
	UpdateUser                  string
	Version                     string
	Deleted                     string
	Status                      string
	RequiredPlatformNo          string
	ChannelHaul                 string
	KingdeeID                   string
	CustomerChannelTypeID       string
	ChannelCountryAffiliationID string
}{
	ID:                          "id",
	CustomerChannelType:         "customer_channel_type",
	CustomerChannelName:         "customer_channel_name",
	MinDeclare:                  "min_declare",
	MaxDeclare:                  "max_declare",
	MinWeight:                   "min_weight",
	MaxWeight:                   "max_weight",
	MinSendWeight:               "min_send_weight",
	MaxSendWeight:               "max_send_weight",
	InStorageUseOrderWeightFlag: "in_storage_use_order_weight_flag",
	InStorageOrderWeightMin:     "in_storage_order_weight_min",
	InStorageOrderWeightMax:     "in_storage_order_weight_max",
	SenderBlackList:             "sender_black_list",
	BanGoodsList:                "ban_goods_list",
	NotePublic:                  "note_public",
	NotePrivate:                 "note_private",
	OrderNumPoolID:              "order_num_pool_id",
	OrderNumPoolName:            "order_num_pool_name",
	ExchangeOrder:               "exchange_order",
	ExchangeType:                "exchange_type",
	LabelCode:                   "label_code",
	DataPushAPI:                 "data_push_api",
	Proportion:                  "proportion",
	CreateTime:                  "create_time",
	CreateUser:                  "create_user",
	UpdateTime:                  "update_time",
	UpdateUser:                  "update_user",
	Version:                     "version",
	Deleted:                     "deleted",
	Status:                      "status",
	RequiredPlatformNo:          "required_platform_no",
	ChannelHaul:                 "channel_haul",
	KingdeeID:                   "kingdee_id",
	CustomerChannelTypeID:       "customer_channel_type_id",
	ChannelCountryAffiliationID: "channel_country_affiliation_id",
}

// LgCustomerChannelCountry 客户渠道_主持国家
type LgCustomerChannelCountry struct {
	ID                    int             `gorm:"primaryKey;column:id" json:"-"`
	CustomerChannelID     int             `gorm:"column:customer_channel_id" json:"customerChannelId"`                                       // 客户渠道ID
	CountryTwoCode        string          `gorm:"column:country_two_code" json:"countryTwoCode"`                                             // 国家二字代码
	ProviderChannelCode   string          `gorm:"column:provider_channel_code" json:"providerChannelCode"`                                   // 服务商渠道Code
	LabelSupport          int             `gorm:"column:label_support" json:"labelSupport"`                                                  // 面单支持
	LabelCode             string          `gorm:"column:label_code" json:"labelCode"`                                                        // 制作面单code
	LgLabelTemplate       LgLabelTemplate `gorm:"joinForeignKey:label_code;foreignKey:code;references:LabelCode" json:"lgLabelTemplateList"` // 物流面单模板维护
	MaxWeight             float64         `gorm:"column:max_weight" json:"maxWeight"`                                                        // 最大出库重量
	MinWeight             float64         `gorm:"column:min_weight" json:"minWeight"`                                                        // 最小出库重量
	ChannelCountryAddress string          `gorm:"column:channel_country_address" json:"channelCountryAddress"`                               // 出/入库地址
	CreateTime            time.Time       `gorm:"column:create_time" json:"createTime"`                                                      // 创建时间
	CreateUser            int             `gorm:"column:create_user" json:"createUser"`                                                      // 创建用户
	UpdateTime            time.Time       `gorm:"column:update_time" json:"updateTime"`                                                      // 更新时间
	UpdateUser            int             `gorm:"column:update_user" json:"updateUser"`                                                      // 更新时间
	Version               int             `gorm:"column:version" json:"version"`                                                             // 乐观锁
	ShipmentCode          string          `gorm:"column:shipment_code" json:"shipmentCode"`                                                  // 轨迹标识CODE(APICODE)
	CarrierCode           string          `gorm:"column:carrier_code" json:"carrierCode"`                                                    // 运输商代码(例如china-post)
}

// TableName get sql table name.获取数据库表名
func (m *LgCustomerChannelCountry) TableName() string {
	return "lg_customer_channel_country"
}

// LgCustomerChannelCountryColumns get sql column name.获取数据库列名
var LgCustomerChannelCountryColumns = struct {
	ID                    string
	CustomerChannelID     string
	CountryTwoCode        string
	ProviderChannelCode   string
	LabelSupport          string
	LabelCode             string
	MaxWeight             string
	MinWeight             string
	ChannelCountryAddress string
	CreateTime            string
	CreateUser            string
	UpdateTime            string
	UpdateUser            string
	Version               string
	ShipmentCode          string
	CarrierCode           string
}{
	ID:                    "id",
	CustomerChannelID:     "customer_channel_id",
	CountryTwoCode:        "country_two_code",
	ProviderChannelCode:   "provider_channel_code",
	LabelSupport:          "label_support",
	LabelCode:             "label_code",
	MaxWeight:             "max_weight",
	MinWeight:             "min_weight",
	ChannelCountryAddress: "channel_country_address",
	CreateTime:            "create_time",
	CreateUser:            "create_user",
	UpdateTime:            "update_time",
	UpdateUser:            "update_user",
	Version:               "version",
	ShipmentCode:          "shipment_code",
	CarrierCode:           "carrier_code",
}

// LgCustomerChannelCountryDownstream 客户渠道_支持国家_下家
type LgCustomerChannelCountryDownstream struct {
	ID                       int          `gorm:"primaryKey;column:id" json:"-"`
	CustomerChannelID        int          `gorm:"column:customer_channel_id" json:"customerChannelId"`                                        // 客户渠道ID
	CountryTwoCode           string       `gorm:"column:country_two_code" json:"countryTwoCode"`                                              // 国家二字码
	CustomerChannelCountryID int          `gorm:"column:customer_channel_country_id" json:"customerChannelCountryId"`                         // 客户渠道支持国家id
	DownstreamID             int          `gorm:"column:downstream_id" json:"downstreamId"`                                                   // 下家id
	LgDownstream             LgDownstream `gorm:"joinForeignKey:downstream_id;foreignKey:id;references:DownstreamID" json:"lgDownstreamList"` // 下家
	DownstreamName           string       `gorm:"column:downstream_name" json:"downstreamName"`                                               // 下家名称
	DownstreamCode           string       `gorm:"column:downstream_code" json:"downstreamCode"`                                               // 下家code
	CreateTime               time.Time    `gorm:"column:create_time" json:"createTime"`                                                       // 创建时间
	CreateUser               int          `gorm:"column:create_user" json:"createUser"`                                                       // 创建用户
	UpdateTime               time.Time    `gorm:"column:update_time" json:"updateTime"`                                                       // 更新时间
	UpdateUser               int          `gorm:"column:update_user" json:"updateUser"`                                                       // 更新时间
	Version                  int          `gorm:"column:version" json:"version"`                                                              // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *LgCustomerChannelCountryDownstream) TableName() string {
	return "lg_customer_channel_country_downstream"
}

// LgCustomerChannelCountryDownstreamColumns get sql column name.获取数据库列名
var LgCustomerChannelCountryDownstreamColumns = struct {
	ID                       string
	CustomerChannelID        string
	CountryTwoCode           string
	CustomerChannelCountryID string
	DownstreamID             string
	DownstreamName           string
	DownstreamCode           string
	CreateTime               string
	CreateUser               string
	UpdateTime               string
	UpdateUser               string
	Version                  string
}{
	ID:                       "id",
	CustomerChannelID:        "customer_channel_id",
	CountryTwoCode:           "country_two_code",
	CustomerChannelCountryID: "customer_channel_country_id",
	DownstreamID:             "downstream_id",
	DownstreamName:           "downstream_name",
	DownstreamCode:           "downstream_code",
	CreateTime:               "create_time",
	CreateUser:               "create_user",
	UpdateTime:               "update_time",
	UpdateUser:               "update_user",
	Version:                  "version",
}

// LgDistrict [...]
type LgDistrict struct {
	ID               int         `gorm:"primaryKey;column:id" json:"-"`                                                                      // ID
	CountryNameEn    string      `gorm:"column:country_name_en" json:"countryNameEn"`                                                        // 国家英文名
	CountryNameCn    string      `gorm:"column:country_name_cn" json:"countryNameCn"`                                                        // 国家中文名
	CountryTwoCode   string      `gorm:"column:country_two_code" json:"countryTwoCode"`                                                      // 国家二字代码
	LgCountryList    []LgCountry `gorm:"joinForeignKey:country_two_code;foreignKey:two_code;references:CountryTwoCode" json:"lgCountryList"` // 国家表(废弃)
	CountryThreeCode string      `gorm:"column:country_three_code" json:"countryThreeCode"`                                                  // 国家三字代码
	ProvinceNameEn   string      `gorm:"column:province_name_en" json:"provinceNameEn"`                                                      // 省份英文名
	ProvinceNameCn   string      `gorm:"column:province_name_cn" json:"provinceNameCn"`                                                      // 省份中文名
	ProvinceCode     string      `gorm:"column:province_code" json:"provinceCode"`                                                           // 省份二字代码
	CityNameEn       string      `gorm:"column:city_name_en" json:"cityNameEn"`                                                              // 城市英文名
	CityNameCn       string      `gorm:"column:city_name_cn" json:"cityNameCn"`                                                              // 城市中文名
	CityCode         string      `gorm:"column:city_code" json:"cityCode"`                                                                   // 城市二字代码
	Zipcode          string      `gorm:"column:zipcode" json:"zipcode"`                                                                      // 邮编
	Address          string      `gorm:"column:address" json:"address"`                                                                      // 地址
	IsVague          int         `gorm:"column:is_vague" json:"isVague"`                                                                     // 是否模糊匹配
	CreateTime       time.Time   `gorm:"column:create_time" json:"createTime"`                                                               // 创建时间
	CreateUser       int         `gorm:"column:create_user" json:"createUser"`                                                               // 创建用户
	UpdateTime       time.Time   `gorm:"column:update_time" json:"updateTime"`                                                               // 更新时间
	UpdateUser       int         `gorm:"column:update_user" json:"updateUser"`                                                               // 更新时间
	Version          int         `gorm:"column:version" json:"version"`                                                                      // 乐观锁
	Deleted          int         `gorm:"column:deleted" json:"deleted"`                                                                      // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgDistrict) TableName() string {
	return "lg_district"
}

// LgDistrictColumns get sql column name.获取数据库列名
var LgDistrictColumns = struct {
	ID               string
	CountryNameEn    string
	CountryNameCn    string
	CountryTwoCode   string
	CountryThreeCode string
	ProvinceNameEn   string
	ProvinceNameCn   string
	ProvinceCode     string
	CityNameEn       string
	CityNameCn       string
	CityCode         string
	Zipcode          string
	Address          string
	IsVague          string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	CountryNameEn:    "country_name_en",
	CountryNameCn:    "country_name_cn",
	CountryTwoCode:   "country_two_code",
	CountryThreeCode: "country_three_code",
	ProvinceNameEn:   "province_name_en",
	ProvinceNameCn:   "province_name_cn",
	ProvinceCode:     "province_code",
	CityNameEn:       "city_name_en",
	CityNameCn:       "city_name_cn",
	CityCode:         "city_code",
	Zipcode:          "zipcode",
	Address:          "address",
	IsVague:          "is_vague",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
}

// LgDownstream 下家
type LgDownstream struct {
	ID                     int       `gorm:"primaryKey;column:id" json:"-"`
	DownstreamCode         string    `gorm:"column:downstream_code" json:"downstreamCode"`                  // 下家code
	DownstreamName         string    `gorm:"column:downstream_name" json:"downstreamName"`                  // 下家名称
	DownstreamProviderCode string    `gorm:"column:downstream_provider_code" json:"downstreamProviderCode"` // 下家服务商代码
	ForecastTemplateCode   string    `gorm:"column:forecast_template_code" json:"forecastTemplateCode"`     // 预报模板code
	VariableList           string    `gorm:"column:variable_list" json:"variableList"`                      // 变量列表
	VariableKeyValueList   string    `gorm:"column:variable_key_value_list" json:"variableKeyValueList"`    // 变量key-value列表
	CreateTime             time.Time `gorm:"column:create_time" json:"createTime"`                          // 创建时间
	CreateUser             int       `gorm:"column:create_user" json:"createUser"`                          // 创建用户
	UpdateTime             time.Time `gorm:"column:update_time" json:"updateTime"`                          // 更新时间
	UpdateUser             int       `gorm:"column:update_user" json:"updateUser"`                          // 更新时间
	Version                int       `gorm:"column:version" json:"version"`                                 // 乐观锁
	Deleted                int       `gorm:"column:deleted" json:"deleted"`                                 // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgDownstream) TableName() string {
	return "lg_downstream"
}

// LgDownstreamColumns get sql column name.获取数据库列名
var LgDownstreamColumns = struct {
	ID                     string
	DownstreamCode         string
	DownstreamName         string
	DownstreamProviderCode string
	ForecastTemplateCode   string
	VariableList           string
	VariableKeyValueList   string
	CreateTime             string
	CreateUser             string
	UpdateTime             string
	UpdateUser             string
	Version                string
	Deleted                string
}{
	ID:                     "id",
	DownstreamCode:         "downstream_code",
	DownstreamName:         "downstream_name",
	DownstreamProviderCode: "downstream_provider_code",
	ForecastTemplateCode:   "forecast_template_code",
	VariableList:           "variable_list",
	VariableKeyValueList:   "variable_key_value_list",
	CreateTime:             "create_time",
	CreateUser:             "create_user",
	UpdateTime:             "update_time",
	UpdateUser:             "update_user",
	Version:                "version",
	Deleted:                "deleted",
}

// LgLabelTemplate 物流面单模板维护
type LgLabelTemplate struct {
	ID                   int64              `gorm:"primaryKey;column:id" json:"-"`
	Name                 string             `gorm:"column:name" json:"name"`                                                                                       // 面单名称
	Code                 string             `gorm:"column:code" json:"code"`                                                                                       // 面单代码
	Type                 int                `gorm:"column:type" json:"type"`                                                                                       // 面单类型1：自制 2：服务商
	ProviderSystemID     int                `gorm:"column:provider_system_id" json:"providerSystemId"`                                                             // 绑定的服务商系统id
	LgProviderSystemList []LgProviderSystem `gorm:"joinForeignKey:provider_system_id;foreignKey:id;references:ProviderSystemID" json:"lgProviderSystemList"`       // 系统服务商
	ProviderSystemName   string             `gorm:"column:provider_system_name" json:"providerSystemName"`                                                         // 绑定的服务商系统名称
	ProviderSystemCode   string             `gorm:"column:provider_system_code" json:"providerSystemCode"`                                                         // 绑定服务商系统code
	LgProviderSystem     LgProviderSystem   `gorm:"joinForeignKey:provider_system_code;foreignKey:code;references:ProviderSystemCode" json:"lgProviderSystemList"` // 系统服务商
	Remark               string             `gorm:"column:remark" json:"remark"`                                                                                   // 备注
	CreateTime           time.Time          `gorm:"column:create_time" json:"createTime"`                                                                          // 创建时间
	CreateUser           int                `gorm:"column:create_user" json:"createUser"`                                                                          // 创建用户
	UpdateTime           time.Time          `gorm:"column:update_time" json:"updateTime"`                                                                          // 更新时间
	UpdateUser           int                `gorm:"column:update_user" json:"updateUser"`                                                                          // 更新时间
	Version              int                `gorm:"column:version" json:"version"`                                                                                 // 乐观锁
	Deleted              int                `gorm:"column:deleted" json:"deleted"`                                                                                 // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgLabelTemplate) TableName() string {
	return "lg_label_template"
}

// LgLabelTemplateColumns get sql column name.获取数据库列名
var LgLabelTemplateColumns = struct {
	ID                 string
	Name               string
	Code               string
	Type               string
	ProviderSystemID   string
	ProviderSystemName string
	ProviderSystemCode string
	Remark             string
	CreateTime         string
	CreateUser         string
	UpdateTime         string
	UpdateUser         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	Name:               "name",
	Code:               "code",
	Type:               "type",
	ProviderSystemID:   "provider_system_id",
	ProviderSystemName: "provider_system_name",
	ProviderSystemCode: "provider_system_code",
	Remark:             "remark",
	CreateTime:         "create_time",
	CreateUser:         "create_user",
	UpdateTime:         "update_time",
	UpdateUser:         "update_user",
	Version:            "version",
	Deleted:            "deleted",
}

// LgMercadoOrder 美卡多订单
type LgMercadoOrder struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`                        // 主键id
	ChannelCode        string    `gorm:"column:channel_code" json:"channelCode"`               // 渠道代码
	MercadoID          string    `gorm:"column:mercado_id" json:"mercadoId"`                   // 美卡多id
	OrderNum           string    `gorm:"column:order_num" json:"orderNum"`                     // 订单号
	Account            string    `gorm:"column:account" json:"account"`                        // 下单账户
	AuthorizationCode  string    `gorm:"column:authorization_code" json:"authorizationCode"`   // 账户授权码
	OrderStatus        int       `gorm:"column:order_status" json:"orderStatus"`               // 订单状态
	PostOfficeStatus   int       `gorm:"column:post_office_status" json:"postOfficeStatus"`    // 取号状态
	TrackStatus        int       `gorm:"column:track_status" json:"trackStatus"`               // 轨迹状态
	TrackNumber        string    `gorm:"column:track_number" json:"trackNumber"`               // 跟踪号
	ReceiveName        string    `gorm:"column:receive_name" json:"receiveName"`               // 收件人名称
	ReceiveTax         string    `gorm:"column:receive_tax" json:"receiveTax"`                 // 收件人税号
	ReceivePhone       string    `gorm:"column:receive_phone" json:"receivePhone"`             // 收件人电话
	ReceiveZipCode     string    `gorm:"column:receive_zip_code" json:"receiveZipCode"`        // 收件人邮编
	ReceiveCountry     string    `gorm:"column:receive_country" json:"receiveCountry"`         // 收件人国家
	ReceiveProvince    string    `gorm:"column:receive_province" json:"receiveProvince"`       // 收件人省份
	ReceiveCity        string    `gorm:"column:receive_city" json:"receiveCity"`               // 收件人城市
	ReceiveAddress     string    `gorm:"column:receive_address" json:"receiveAddress"`         // 收件人地址
	DeliveryPreference string    `gorm:"column:delivery_preference" json:"deliveryPreference"` // 地址类型
	SenderID           int       `gorm:"column:sender_id" json:"senderId"`                     // 发件人id
	SenderName         string    `gorm:"column:sender_name" json:"senderName"`                 // 发件人名称
	SenderPhone        string    `gorm:"column:sender_phone" json:"senderPhone"`               // 发件人电话
	SenderTax          string    `gorm:"column:sender_tax" json:"senderTax"`                   // 发件人税号
	SenderAddress      string    `gorm:"column:sender_address" json:"senderAddress"`           // 发件人地址
	ItemInfo           string    `gorm:"column:item_info" json:"itemInfo"`                     // 商品信息
	Quantity           int       `gorm:"column:quantity" json:"quantity"`                      // 申报数量
	DeclaredValue      float64   `gorm:"column:declared_value" json:"declaredValue"`           // 申报金额
	DeclaredWeight     float64   `gorm:"column:declared_weight" json:"declaredWeight"`         // 申报重量
	RequestParams      string    `gorm:"column:request_params" json:"requestParams"`           // 请求参数
	ResponseParams     string    `gorm:"column:response_params" json:"responseParams"`         // 响应参数
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                 // 创建时间
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`                 // 更新时间
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                        // 逻辑删除
	Version            int       `gorm:"column:version" json:"version"`                        // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *LgMercadoOrder) TableName() string {
	return "lg_mercado_order"
}

// LgMercadoOrderColumns get sql column name.获取数据库列名
var LgMercadoOrderColumns = struct {
	ID                 string
	ChannelCode        string
	MercadoID          string
	OrderNum           string
	Account            string
	AuthorizationCode  string
	OrderStatus        string
	PostOfficeStatus   string
	TrackStatus        string
	TrackNumber        string
	ReceiveName        string
	ReceiveTax         string
	ReceivePhone       string
	ReceiveZipCode     string
	ReceiveCountry     string
	ReceiveProvince    string
	ReceiveCity        string
	ReceiveAddress     string
	DeliveryPreference string
	SenderID           string
	SenderName         string
	SenderPhone        string
	SenderTax          string
	SenderAddress      string
	ItemInfo           string
	Quantity           string
	DeclaredValue      string
	DeclaredWeight     string
	RequestParams      string
	ResponseParams     string
	CreateTime         string
	UpdateTime         string
	Deleted            string
	Version            string
}{
	ID:                 "id",
	ChannelCode:        "channel_code",
	MercadoID:          "mercado_id",
	OrderNum:           "order_num",
	Account:            "account",
	AuthorizationCode:  "authorization_code",
	OrderStatus:        "order_status",
	PostOfficeStatus:   "post_office_status",
	TrackStatus:        "track_status",
	TrackNumber:        "track_number",
	ReceiveName:        "receive_name",
	ReceiveTax:         "receive_tax",
	ReceivePhone:       "receive_phone",
	ReceiveZipCode:     "receive_zip_code",
	ReceiveCountry:     "receive_country",
	ReceiveProvince:    "receive_province",
	ReceiveCity:        "receive_city",
	ReceiveAddress:     "receive_address",
	DeliveryPreference: "delivery_preference",
	SenderID:           "sender_id",
	SenderName:         "sender_name",
	SenderPhone:        "sender_phone",
	SenderTax:          "sender_tax",
	SenderAddress:      "sender_address",
	ItemInfo:           "item_info",
	Quantity:           "quantity",
	DeclaredValue:      "declared_value",
	DeclaredWeight:     "declared_weight",
	RequestParams:      "request_params",
	ResponseParams:     "response_params",
	CreateTime:         "create_time",
	UpdateTime:         "update_time",
	Deleted:            "deleted",
	Version:            "version",
}

// LgNumPool 单号池
type LgNumPool struct {
	ID               int64     `gorm:"primaryKey;column:id" json:"-"`                    // 单号池ID
	Name             string    `gorm:"column:name" json:"name"`                          // 单号池名称
	Prefix           string    `gorm:"column:prefix" json:"prefix"`                      // 单号池前缀
	Suffix           string    `gorm:"column:suffix" json:"suffix"`                      // 单号池后缀
	StartNum         int64     `gorm:"column:start_num" json:"startNum"`                 // 开始号码
	EndNum           int64     `gorm:"column:end_num" json:"endNum"`                     // 结束号码
	VerificationMode string    `gorm:"column:verification_mode" json:"verificationMode"` // 验证号码合法方式:保存验证方式code
	UsedNum          int64     `gorm:"column:used_num" json:"usedNum"`                   // 已用到单号池位置:默认喂号码开始位置
	LowestThreshold  int64     `gorm:"column:lowest_threshold" json:"lowestThreshold"`   // 最低阈值
	NoticeUserID     int64     `gorm:"column:notice_user_id" json:"noticeUserId"`        // 低于阈值通知人ID
	Remark           string    `gorm:"column:remark" json:"remark"`                      // 单号池备注
	Type             int       `gorm:"column:type" json:"type"`                          // 单号池类型,1:默认类型,号段,2:倒入单号类型,所有单号需要导入
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`             // 创建用户
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`             // 更新用户
	Version          int       `gorm:"column:version" json:"version"`                    // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                    // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgNumPool) TableName() string {
	return "lg_num_pool"
}

// LgNumPoolColumns get sql column name.获取数据库列名
var LgNumPoolColumns = struct {
	ID               string
	Name             string
	Prefix           string
	Suffix           string
	StartNum         string
	EndNum           string
	VerificationMode string
	UsedNum          string
	LowestThreshold  string
	NoticeUserID     string
	Remark           string
	Type             string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	Name:             "name",
	Prefix:           "prefix",
	Suffix:           "suffix",
	StartNum:         "start_num",
	EndNum:           "end_num",
	VerificationMode: "verification_mode",
	UsedNum:          "used_num",
	LowestThreshold:  "lowest_threshold",
	NoticeUserID:     "notice_user_id",
	Remark:           "remark",
	Type:             "type",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
}

// LgNumPoolFetchRecord 单号池附表:主要记录提取记录
type LgNumPoolFetchRecord struct {
	ID           int64     `gorm:"primaryKey;column:id" json:"-"`            // 单号池提取记录ID
	OrderNumID   int64     `gorm:"column:order_num_id" json:"orderNumId"`    // 单号池关联ID
	Name         string    `gorm:"column:name" json:"name"`                  // 单号池名称
	Prefix       string    `gorm:"column:prefix" json:"prefix"`              // 单号池前缀
	Suffix       string    `gorm:"column:suffix" json:"suffix"`              // 单号池后缀
	StartNum     int64     `gorm:"column:start_num" json:"startNum"`         // 单号池提取开始位置,数字部分
	EndNum       int64     `gorm:"column:end_num" json:"endNum"`             // 单号池提取结束位置,数字部分
	Type         string    `gorm:"column:type" json:"type"`                  // 提取类型:self:自用 customer:客户使用
	CustomerID   int       `gorm:"column:customer_id" json:"customerId"`     // 客户ID
	CustomerName string    `gorm:"column:customer_name" json:"customerName"` // 客户名称
	Remark       string    `gorm:"column:remark" json:"remark"`              // 备注
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建用户
}

// TableName get sql table name.获取数据库表名
func (m *LgNumPoolFetchRecord) TableName() string {
	return "lg_num_pool_fetch_record"
}

// LgNumPoolFetchRecordColumns get sql column name.获取数据库列名
var LgNumPoolFetchRecordColumns = struct {
	ID           string
	OrderNumID   string
	Name         string
	Prefix       string
	Suffix       string
	StartNum     string
	EndNum       string
	Type         string
	CustomerID   string
	CustomerName string
	Remark       string
	CreateTime   string
	CreateUser   string
}{
	ID:           "id",
	OrderNumID:   "order_num_id",
	Name:         "name",
	Prefix:       "prefix",
	Suffix:       "suffix",
	StartNum:     "start_num",
	EndNum:       "end_num",
	Type:         "type",
	CustomerID:   "customer_id",
	CustomerName: "customer_name",
	Remark:       "remark",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
}

// LgNumPoolSub 单号池附属表:主要记录具体的可分配单号
type LgNumPoolSub struct {
	ID                     int64     `gorm:"primaryKey;column:id" json:"-"`                                 // 单号池副表主键ID
	OrderNumPoolID         int64     `gorm:"column:order_num_pool_id" json:"orderNumPoolId"`                // 绑定的单号池主单ID
	TrackNo                string    `gorm:"column:track_no" json:"trackNo"`                                // 物流单号
	ExtractionOrderNum     string    `gorm:"column:extraction_order_num" json:"extractionOrderNum"`         // 提取单号:取号后回填客户单号
	ExtractionCustomerID   int64     `gorm:"column:extraction_customer_id" json:"extractionCustomerId"`     // 提取客户ID
	ExtractionCustomerName string    `gorm:"column:extraction_customer_name" json:"extractionCustomerName"` // 提取客户名称
	CreateTime             time.Time `gorm:"column:create_time" json:"createTime"`                          // 创建时间
	CreateUser             int       `gorm:"column:create_user" json:"createUser"`                          // 创建用户
}

// TableName get sql table name.获取数据库表名
func (m *LgNumPoolSub) TableName() string {
	return "lg_num_pool_sub"
}

// LgNumPoolSubColumns get sql column name.获取数据库列名
var LgNumPoolSubColumns = struct {
	ID                     string
	OrderNumPoolID         string
	TrackNo                string
	ExtractionOrderNum     string
	ExtractionCustomerID   string
	ExtractionCustomerName string
	CreateTime             string
	CreateUser             string
}{
	ID:                     "id",
	OrderNumPoolID:         "order_num_pool_id",
	TrackNo:                "track_no",
	ExtractionOrderNum:     "extraction_order_num",
	ExtractionCustomerID:   "extraction_customer_id",
	ExtractionCustomerName: "extraction_customer_name",
	CreateTime:             "create_time",
	CreateUser:             "create_user",
}

// LgOperateOrderLog 订单操作日志
type LgOperateOrderLog struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`                        // 主键id
	ReferenceNumber    string    `gorm:"column:reference_number" json:"referenceNumber"`       // 参考号(安骏单号)
	OperateType        string    `gorm:"column:operate_type" json:"operateType"`               // 操作类型
	OperateDescription string    `gorm:"column:operate_description" json:"operateDescription"` // 操作描述
	OperateIP          string    `gorm:"column:operate_ip" json:"operateIp"`                   // 操作IP
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                 // 操作人id
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                 // 操作时间
	Version            int       `gorm:"column:version" json:"version"`                        // 乐观锁
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOperateOrderLog) TableName() string {
	return "lg_operate_order_log"
}

// LgOperateOrderLogColumns get sql column name.获取数据库列名
var LgOperateOrderLogColumns = struct {
	ID                 string
	ReferenceNumber    string
	OperateType        string
	OperateDescription string
	OperateIP          string
	CreateUser         string
	CreateTime         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	ReferenceNumber:    "reference_number",
	OperateType:        "operate_type",
	OperateDescription: "operate_description",
	OperateIP:          "operate_ip",
	CreateUser:         "create_user",
	CreateTime:         "create_time",
	Version:            "version",
	Deleted:            "deleted",
}

// LgOrder 物流订单
type LgOrder struct {
	ID                      int            `gorm:"primaryKey;column:id" json:"-"`                                  // ID
	OrderNumber             string         `gorm:"column:order_number" json:"orderNumber"`                         // 订单号
	ReferenceNumber         string         `gorm:"column:reference_number" json:"referenceNumber"`                 // 参考号
	LogisticsNumber         string         `gorm:"column:logistics_number" json:"logisticsNumber"`                 // 物流单号
	LogisticsNumberFinal    string         `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"`      // 最终物流单号
	ExpressNum              string         `gorm:"column:express_num" json:"expressNum"`                           // 订单快递单号
	PlatformNumber          string         `gorm:"column:platform_number" json:"platformNumber"`                   // 平台订单
	ProviderOrderID         string         `gorm:"column:provider_order_id" json:"providerOrderId"`                // 服务商系统订单ID
	LabelURL                string         `gorm:"column:label_url" json:"labelUrl"`                               // 面单地址
	LabelURLFinal           string         `gorm:"column:label_url_final" json:"labelUrlFinal"`                    // 最终面单地址
	OnlineNumber            string         `gorm:"column:online_number" json:"onlineNumber"`                       // 上网单号(17单号)
	BagNumber               string         `gorm:"column:bag_number" json:"bagNumber"`                             // 袋号
	TotalListNo             string         `gorm:"column:total_list_no" json:"totalListNo"`                        // 总单号
	AirBillNumber           string         `gorm:"column:air_bill_number" json:"airBillNumber"`                    // 航空提单号
	CustomerChannelID       int            `gorm:"column:customer_channel_id" json:"customerChannelId"`            // 客户渠道ID
	CustomerChannelName     string         `gorm:"column:customer_channel_name" json:"customerChannelName"`        // 客户渠道名称
	ProviderCode            string         `gorm:"column:provider_code" json:"providerCode"`                       // 服务商系统code
	ProviderChannelCode     string         `gorm:"column:provider_channel_code" json:"providerChannelCode"`        // 服务商系统渠道Code
	ProviderChannelName     string         `gorm:"column:provider_Channel_Name" json:"providerChannelName"`        // 服务商渠道名称
	ReceiveCountry          string         `gorm:"column:receive_country" json:"receiveCountry"`                   // 收件人国家
	ReceiveName             string         `gorm:"column:receive_name" json:"receiveName"`                         // 收件人姓名
	ReceivePhone            string         `gorm:"column:receive_phone" json:"receivePhone"`                       // 收件人电话
	ReceiveMobile           string         `gorm:"column:receive_mobile" json:"receiveMobile"`                     // 收件人手机
	ReceiveTax              string         `gorm:"column:receive_tax" json:"receiveTax"`                           // 收件人税号
	ReceivePassport         string         `gorm:"column:receive_passport" json:"receivePassport"`                 // 收件人护照号
	ReceiveMail             string         `gorm:"column:receive_mail" json:"receiveMail"`                         // 收件人邮箱
	ReceiveCompany          string         `gorm:"column:receive_company" json:"receiveCompany"`                   // 收件人公司
	ReceiveZipcode          string         `gorm:"column:receive_zipcode" json:"receiveZipcode"`                   // 收件人邮编
	ReceiveProvince         string         `gorm:"column:receive_province" json:"receiveProvince"`                 // 收件人省份
	ReceiveHouseNumber      string         `gorm:"column:receive_house_number" json:"receiveHouseNumber"`          // 收件人门牌号
	ReceiveArea             string         `gorm:"column:receive_area" json:"receiveArea"`                         // 收件人区
	ReceiveStreet           string         `gorm:"column:receive_street" json:"receiveStreet"`                     // 收件人街道
	ReceiveCity             string         `gorm:"column:receive_city" json:"receiveCity"`                         // 收件人城市
	ReceiveAddress1         string         `gorm:"column:receive_address1" json:"receiveAddress1"`                 // 收件人地址1
	ReceiveAddress2         string         `gorm:"column:receive_address2" json:"receiveAddress2"`                 // 收件人地址2
	SenderCountry           string         `gorm:"column:sender_country" json:"senderCountry"`                     // 发件人国家
	SenderName              string         `gorm:"column:sender_name" json:"senderName"`                           // 发件人姓名
	SenderPhone             string         `gorm:"column:sender_phone" json:"senderPhone"`                         // 发件人电话
	SenderMobile            string         `gorm:"column:sender_mobile" json:"senderMobile"`                       // 发件人手机
	SenderTax               string         `gorm:"column:sender_tax" json:"senderTax"`                             // 发件人税号
	SenderPassport          string         `gorm:"column:sender_passport" json:"senderPassport"`                   // 发件人护照号
	SenderMail              string         `gorm:"column:sender_mail" json:"senderMail"`                           // 发件人邮箱
	SenderCompany           string         `gorm:"column:sender_company" json:"senderCompany"`                     // 发件人公司
	SenderZipcode           string         `gorm:"column:sender_zipcode" json:"senderZipcode"`                     // 发件人邮编
	SenderProvince          string         `gorm:"column:sender_province" json:"senderProvince"`                   // 发件人省份
	SenderCity              string         `gorm:"column:sender_city" json:"senderCity"`                           // 发件人城市
	SenderArea              string         `gorm:"column:sender_area" json:"senderArea"`                           // 发件人区
	SenderStreet            string         `gorm:"column:sender_street" json:"senderStreet"`                       // 发件人街道
	SenderHouseNumber       string         `gorm:"column:sender_house_number" json:"senderHouseNumber"`            // 发件人门牌号
	SenderAddress           string         `gorm:"column:sender_address" json:"senderAddress"`                     // 发件人地址
	PackageType             string         `gorm:"column:package_type" json:"packageType"`                         // 包裹类型
	HasBattery              int            `gorm:"column:has_battery" json:"hasBattery"`                           // 是否带电
	BatteryType             string         `gorm:"column:battery_type" json:"batteryType"`                         // 电池类型
	HasBack                 int            `gorm:"column:has_back" json:"hasBack"`                                 // 是否退回
	Remarks                 string         `gorm:"column:remarks" json:"remarks"`                                  // 备注信息
	DistributionInformation string         `gorm:"column:distribution_information" json:"distributionInformation"` // 配货信息
	Insurance               string         `gorm:"column:insurance" json:"insurance"`                              // 保险
	SalesPlatform           string         `gorm:"column:sales_platform" json:"salesPlatform"`                     // 销售平台
	Weight                  float64        `gorm:"column:weight" json:"weight"`                                    // 预报重量
	Quantity                int            `gorm:"column:quantity" json:"quantity"`                                // 包裹数量
	DeclaredValue           float64        `gorm:"column:declared_value" json:"declaredValue"`                     // 申报价值
	Status                  string         `gorm:"column:status" json:"status"`                                    // 订单状态
	TransportStatus         string         `gorm:"column:transport_status" json:"transportStatus"`                 // 转运状态
	LastTrackStatus         string         `gorm:"column:last_track_status" json:"lastTrackStatus"`                // 最后一条轨迹状态
	LastTrackAddress        string         `gorm:"column:last_track_address" json:"lastTrackAddress"`              // 最后一条轨迹
	LastTrackTime           time.Time      `gorm:"column:last_track_time" json:"lastTrackTime"`                    // 最后一条轨迹的时间
	ProblemType             string         `gorm:"column:problem_type" json:"problemType"`                         // 问题件类型
	ProblemReason           string         `gorm:"column:problem_reason" json:"problemReason"`                     // 问题件原因
	WeighingWeight          float64        `gorm:"column:weighing_weight" json:"weighingWeight"`                   // 称重重量
	VolumeWeight            float64        `gorm:"column:volume_weight" json:"volumeWeight"`                       // 材积重量
	ChargeWeight            float64        `gorm:"column:charge_weight" json:"chargeWeight"`                       // 收费重量
	StandardFee             float64        `gorm:"column:standard_fee" json:"standardFee"`                         // 标准费用
	Currency                string         `gorm:"column:currency" json:"currency"`                                // 币种
	BuyersFee               float64        `gorm:"column:buyers_fee" json:"buyersFee"`                             // 下家费用
	BuyersWeight            float64        `gorm:"column:buyers_weight" json:"buyersWeight"`                       // 下家重量
	OtherFee                float64        `gorm:"column:other_fee" json:"otherFee"`                               // 其他费用
	CustomerID              int64          `gorm:"column:customer_id" json:"customerId"`                           // 归属客户ID
	CustomerName            string         `gorm:"column:customer_name" json:"customerName"`                       // 归属客户名称
	CustomerAlias           string         `gorm:"column:customer_alias" json:"customerAlias"`                     // 客户别名
	PrepaymentVat           string         `gorm:"column:prepayment_vat" json:"prepaymentVat"`                     // "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
	Vat                     string         `gorm:"column:vat" json:"vat"`                                          // vat号(英国方向必填)
	ReceiveCertificateType  string         `gorm:"column:receive_certificate_type" json:"receiveCertificateType"`  // 收件人证件类型
	ReceiveCertificateCode  string         `gorm:"column:receive_certificate_code" json:"receiveCertificateCode"`  // 收件人证件号码
	SenderCertificateType   string         `gorm:"column:sender_certificate_type" json:"senderCertificateType"`    // 发件人证件类型
	SenderCertificateCode   string         `gorm:"column:sender_certificate_code" json:"senderCertificateCode"`    // 发件人证件号码
	ServiceChannelCode      string         `gorm:"column:service_channel_code" json:"serviceChannelCode"`          // 服务商渠道代码
	DeliveryTerms           string         `gorm:"column:delivery_terms" json:"deliveryTerms"`                     // 贸易条款-(部分渠道巴西方向必填)- DDU DDP
	ReceiveCountryName      string         `gorm:"column:receive_country_name" json:"receiveCountryName"`          // 收件人国家中文
	SenderCountryName       string         `gorm:"column:sender_country_name" json:"senderCountryName"`            // 发件国家中文
	BillNumber              string         `gorm:"column:bill_number" json:"billNumber"`                           // 账单编号
	ApproximateCost         float64        `gorm:"column:approximate_cost" json:"approximateCost"`                 // 预报费用
	WeighingCost            float64        `gorm:"column:weighing_cost" json:"weighingCost"`                       // 称重费用
	PrepaidAmount           float64        `gorm:"column:prepaid_amount" json:"prepaidAmount"`                     // 预扣款金额
	ActualCost              float64        `gorm:"column:actual_cost" json:"actualCost"`                           // 实收费用
	PrimeCost               float64        `gorm:"column:prime_cost" json:"primeCost"`                             // 成本
	OrderType               int            `gorm:"column:order_type" json:"orderType"`                             // 订单类型(0-预报订单，1-物流订单)
	ExchangeOrder           []uint8        `gorm:"column:exchange_order" json:"exchangeOrder"`                     // 是否需要换单 0 不需要1 需要
	ReceiptTime             time.Time      `gorm:"column:receipt_time" json:"receiptTime"`                         // 入库时间
	ReceiptUser             int            `gorm:"column:receipt_user" json:"receiptUser"`                         // 入库人
	SendTime                time.Time      `gorm:"column:send_time" json:"sendTime"`                               // 出库时间
	SendUser                int            `gorm:"column:send_user" json:"sendUser"`                               // 出库人
	PicOssLink              string         `gorm:"column:pic_oss_link" json:"picOssLink"`                          // 包裹图片oss地址
	Test                    []uint8        `gorm:"column:test" json:"test"`                                        // 是否为测试订单
	OrderSource             int            `gorm:"column:order_source" json:"orderSource"`                         // 订单来源：1-默认，2-平台订单
	PlatformCode            string         `gorm:"column:platform_code" json:"platformCode"`                       // 平台代码
	PlatformName            string         `gorm:"column:platform_name" json:"platformName"`                       // 平台名称
	PlatformSellerName      string         `gorm:"column:platform_seller_name" json:"platformSellerName"`          // 平台卖家名称
	PlatformOrderTime       time.Time      `gorm:"column:platform_order_time" json:"platformOrderTime"`            // 平台订单时间
	PostOfficeBranchName    string         `gorm:"column:post_office_branch_name" json:"postOfficeBranchName"`     // 邮政局分行名称
	ExpressSendTime         time.Time      `gorm:"column:express_send_time" json:"expressSendTime"`                // 快递寄送时间
	ExpressReceiptTime      time.Time      `gorm:"column:express_receipt_time" json:"expressReceiptTime"`          // 快递签收时间
	FlightDeparturesTime    time.Time      `gorm:"column:flight_departures_time" json:"flightDeparturesTime"`      // 航班起飞时间
	CreateTime              time.Time      `gorm:"column:create_time" json:"createTime"`                           // 创建时间
	DeliveredTime           time.Time      `gorm:"column:delivered_time" json:"deliveredTime"`                     // 签收时间
	FlightArrivalTime       time.Time      `gorm:"column:flight_arrival_time" json:"flightArrivalTime"`            // 航班落地时间
	Extra                   datatypes.JSON `gorm:"column:extra" json:"extra"`                                      // 额外信息
	CreateUser              int            `gorm:"column:create_user" json:"createUser"`                           // 创建用户
	Deleted                 int            `gorm:"column:deleted" json:"deleted"`                                  // 逻辑删除
	UpdateTime              time.Time      `gorm:"column:update_time" json:"updateTime"`                           // 更新时间
	Version                 int            `gorm:"column:version" json:"version"`                                  // 乐观锁
	UpdateUser              int            `gorm:"column:update_user" json:"updateUser"`                           // 更新时间
	ChannelRemarks          string         `gorm:"column:channel_remarks" json:"channelRemarks"`                   // 渠道备注
	FinanceRemarks          string         `gorm:"column:finance_remarks" json:"financeRemarks"`                   // 财务备注
}

// TableName get sql table name.获取数据库表名
func (m *LgOrder) TableName() string {
	return "lg_order"
}

// LgOrderColumns get sql column name.获取数据库列名
var LgOrderColumns = struct {
	ID                      string
	OrderNumber             string
	ReferenceNumber         string
	LogisticsNumber         string
	LogisticsNumberFinal    string
	ExpressNum              string
	PlatformNumber          string
	ProviderOrderID         string
	LabelURL                string
	LabelURLFinal           string
	OnlineNumber            string
	BagNumber               string
	TotalListNo             string
	AirBillNumber           string
	CustomerChannelID       string
	CustomerChannelName     string
	ProviderCode            string
	ProviderChannelCode     string
	ProviderChannelName     string
	ReceiveCountry          string
	ReceiveName             string
	ReceivePhone            string
	ReceiveMobile           string
	ReceiveTax              string
	ReceivePassport         string
	ReceiveMail             string
	ReceiveCompany          string
	ReceiveZipcode          string
	ReceiveProvince         string
	ReceiveHouseNumber      string
	ReceiveArea             string
	ReceiveStreet           string
	ReceiveCity             string
	ReceiveAddress1         string
	ReceiveAddress2         string
	SenderCountry           string
	SenderName              string
	SenderPhone             string
	SenderMobile            string
	SenderTax               string
	SenderPassport          string
	SenderMail              string
	SenderCompany           string
	SenderZipcode           string
	SenderProvince          string
	SenderCity              string
	SenderArea              string
	SenderStreet            string
	SenderHouseNumber       string
	SenderAddress           string
	PackageType             string
	HasBattery              string
	BatteryType             string
	HasBack                 string
	Remarks                 string
	DistributionInformation string
	Insurance               string
	SalesPlatform           string
	Weight                  string
	Quantity                string
	DeclaredValue           string
	Status                  string
	TransportStatus         string
	LastTrackStatus         string
	LastTrackAddress        string
	LastTrackTime           string
	ProblemType             string
	ProblemReason           string
	WeighingWeight          string
	VolumeWeight            string
	ChargeWeight            string
	StandardFee             string
	Currency                string
	BuyersFee               string
	BuyersWeight            string
	OtherFee                string
	CustomerID              string
	CustomerName            string
	CustomerAlias           string
	PrepaymentVat           string
	Vat                     string
	ReceiveCertificateType  string
	ReceiveCertificateCode  string
	SenderCertificateType   string
	SenderCertificateCode   string
	ServiceChannelCode      string
	DeliveryTerms           string
	ReceiveCountryName      string
	SenderCountryName       string
	BillNumber              string
	ApproximateCost         string
	WeighingCost            string
	PrepaidAmount           string
	ActualCost              string
	PrimeCost               string
	OrderType               string
	ExchangeOrder           string
	ReceiptTime             string
	ReceiptUser             string
	SendTime                string
	SendUser                string
	PicOssLink              string
	Test                    string
	OrderSource             string
	PlatformCode            string
	PlatformName            string
	PlatformSellerName      string
	PlatformOrderTime       string
	PostOfficeBranchName    string
	ExpressSendTime         string
	ExpressReceiptTime      string
	FlightDeparturesTime    string
	CreateTime              string
	DeliveredTime           string
	FlightArrivalTime       string
	Extra                   string
	CreateUser              string
	Deleted                 string
	UpdateTime              string
	Version                 string
	UpdateUser              string
	ChannelRemarks          string
	FinanceRemarks          string
}{
	ID:                      "id",
	OrderNumber:             "order_number",
	ReferenceNumber:         "reference_number",
	LogisticsNumber:         "logistics_number",
	LogisticsNumberFinal:    "logistics_number_final",
	ExpressNum:              "express_num",
	PlatformNumber:          "platform_number",
	ProviderOrderID:         "provider_order_id",
	LabelURL:                "label_url",
	LabelURLFinal:           "label_url_final",
	OnlineNumber:            "online_number",
	BagNumber:               "bag_number",
	TotalListNo:             "total_list_no",
	AirBillNumber:           "air_bill_number",
	CustomerChannelID:       "customer_channel_id",
	CustomerChannelName:     "customer_channel_name",
	ProviderCode:            "provider_code",
	ProviderChannelCode:     "provider_channel_code",
	ProviderChannelName:     "provider_Channel_Name",
	ReceiveCountry:          "receive_country",
	ReceiveName:             "receive_name",
	ReceivePhone:            "receive_phone",
	ReceiveMobile:           "receive_mobile",
	ReceiveTax:              "receive_tax",
	ReceivePassport:         "receive_passport",
	ReceiveMail:             "receive_mail",
	ReceiveCompany:          "receive_company",
	ReceiveZipcode:          "receive_zipcode",
	ReceiveProvince:         "receive_province",
	ReceiveHouseNumber:      "receive_house_number",
	ReceiveArea:             "receive_area",
	ReceiveStreet:           "receive_street",
	ReceiveCity:             "receive_city",
	ReceiveAddress1:         "receive_address1",
	ReceiveAddress2:         "receive_address2",
	SenderCountry:           "sender_country",
	SenderName:              "sender_name",
	SenderPhone:             "sender_phone",
	SenderMobile:            "sender_mobile",
	SenderTax:               "sender_tax",
	SenderPassport:          "sender_passport",
	SenderMail:              "sender_mail",
	SenderCompany:           "sender_company",
	SenderZipcode:           "sender_zipcode",
	SenderProvince:          "sender_province",
	SenderCity:              "sender_city",
	SenderArea:              "sender_area",
	SenderStreet:            "sender_street",
	SenderHouseNumber:       "sender_house_number",
	SenderAddress:           "sender_address",
	PackageType:             "package_type",
	HasBattery:              "has_battery",
	BatteryType:             "battery_type",
	HasBack:                 "has_back",
	Remarks:                 "remarks",
	DistributionInformation: "distribution_information",
	Insurance:               "insurance",
	SalesPlatform:           "sales_platform",
	Weight:                  "weight",
	Quantity:                "quantity",
	DeclaredValue:           "declared_value",
	Status:                  "status",
	TransportStatus:         "transport_status",
	LastTrackStatus:         "last_track_status",
	LastTrackAddress:        "last_track_address",
	LastTrackTime:           "last_track_time",
	ProblemType:             "problem_type",
	ProblemReason:           "problem_reason",
	WeighingWeight:          "weighing_weight",
	VolumeWeight:            "volume_weight",
	ChargeWeight:            "charge_weight",
	StandardFee:             "standard_fee",
	Currency:                "currency",
	BuyersFee:               "buyers_fee",
	BuyersWeight:            "buyers_weight",
	OtherFee:                "other_fee",
	CustomerID:              "customer_id",
	CustomerName:            "customer_name",
	CustomerAlias:           "customer_alias",
	PrepaymentVat:           "prepayment_vat",
	Vat:                     "vat",
	ReceiveCertificateType:  "receive_certificate_type",
	ReceiveCertificateCode:  "receive_certificate_code",
	SenderCertificateType:   "sender_certificate_type",
	SenderCertificateCode:   "sender_certificate_code",
	ServiceChannelCode:      "service_channel_code",
	DeliveryTerms:           "delivery_terms",
	ReceiveCountryName:      "receive_country_name",
	SenderCountryName:       "sender_country_name",
	BillNumber:              "bill_number",
	ApproximateCost:         "approximate_cost",
	WeighingCost:            "weighing_cost",
	PrepaidAmount:           "prepaid_amount",
	ActualCost:              "actual_cost",
	PrimeCost:               "prime_cost",
	OrderType:               "order_type",
	ExchangeOrder:           "exchange_order",
	ReceiptTime:             "receipt_time",
	ReceiptUser:             "receipt_user",
	SendTime:                "send_time",
	SendUser:                "send_user",
	PicOssLink:              "pic_oss_link",
	Test:                    "test",
	OrderSource:             "order_source",
	PlatformCode:            "platform_code",
	PlatformName:            "platform_name",
	PlatformSellerName:      "platform_seller_name",
	PlatformOrderTime:       "platform_order_time",
	PostOfficeBranchName:    "post_office_branch_name",
	ExpressSendTime:         "express_send_time",
	ExpressReceiptTime:      "express_receipt_time",
	FlightDeparturesTime:    "flight_departures_time",
	CreateTime:              "create_time",
	DeliveredTime:           "delivered_time",
	FlightArrivalTime:       "flight_arrival_time",
	Extra:                   "extra",
	CreateUser:              "create_user",
	Deleted:                 "deleted",
	UpdateTime:              "update_time",
	Version:                 "version",
	UpdateUser:              "update_user",
	ChannelRemarks:          "channel_remarks",
	FinanceRemarks:          "finance_remarks",
}

// LgOrderAddressErrorCorrection 地址纠错
type LgOrderAddressErrorCorrection struct {
	ID                  int       `gorm:"primaryKey;column:id" json:"-"`
	ErrorZipCode        string    `gorm:"column:error_zip_code" json:"errorZipCode"`               // 错误邮编
	CorrectZipCode      string    `gorm:"column:correct_zip_code" json:"correctZipCode"`           // 正确邮编
	ReceiveProvince     string    `gorm:"column:receive_province" json:"receiveProvince"`          // 收件人省份
	ReceiveCity         string    `gorm:"column:receive_city" json:"receiveCity"`                  // 收件人城市
	ReceiveCountry      string    `gorm:"column:receive_country" json:"receiveCountry"`            // 收件人国家
	CustomerChannelID   int       `gorm:"column:customer_channel_id" json:"customerChannelId"`     // 客户渠道ID
	CustomerChannelName string    `gorm:"column:customer_channel_name" json:"customerChannelName"` // 客户渠道名称
	State               []uint8   `gorm:"column:state" json:"state"`                               // 状态(0:禁用，1:启用)
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`                    // 创建时间
	CreateUser          int       `gorm:"column:create_user" json:"createUser"`                    // 创建用户
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`                    // 修改时间
	UpdateUser          int       `gorm:"column:update_user" json:"updateUser"`                    // 更新时间
	Version             int       `gorm:"column:version" json:"version"`                           // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderAddressErrorCorrection) TableName() string {
	return "lg_order_address_error_correction"
}

// LgOrderAddressErrorCorrectionColumns get sql column name.获取数据库列名
var LgOrderAddressErrorCorrectionColumns = struct {
	ID                  string
	ErrorZipCode        string
	CorrectZipCode      string
	ReceiveProvince     string
	ReceiveCity         string
	ReceiveCountry      string
	CustomerChannelID   string
	CustomerChannelName string
	State               string
	CreateTime          string
	CreateUser          string
	UpdateTime          string
	UpdateUser          string
	Version             string
}{
	ID:                  "id",
	ErrorZipCode:        "error_zip_code",
	CorrectZipCode:      "correct_zip_code",
	ReceiveProvince:     "receive_province",
	ReceiveCity:         "receive_city",
	ReceiveCountry:      "receive_country",
	CustomerChannelID:   "customer_channel_id",
	CustomerChannelName: "customer_channel_name",
	State:               "state",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateTime:          "update_time",
	UpdateUser:          "update_user",
	Version:             "version",
}

// LgOrderAddressErrorTemp 地址纠错临时表
type LgOrderAddressErrorTemp struct {
	ID                  int       `gorm:"primaryKey;column:id" json:"-"`
	OrderNumber         string    `gorm:"column:order_number" json:"orderNumber"`                  // 订单号
	ReceiveCity         string    `gorm:"column:receive_city" json:"receiveCity"`                  // 收件人城市
	ReceiveProvince     string    `gorm:"column:receive_province" json:"receiveProvince"`          // 收件人省份
	ReceiveCountry      string    `gorm:"column:receive_country" json:"receiveCountry"`            // 收件人国家
	ReceiveZipCode      string    `gorm:"column:receive_zip_code" json:"receiveZipCode"`           // 收件人邮编
	CustomerChannelID   int       `gorm:"column:customer_channel_id" json:"customerChannelId"`     // 客户渠道ID
	CustomerChannelName string    `gorm:"column:customer_channel_name" json:"customerChannelName"` // 客户渠道名称
	CustomerID          int64     `gorm:"column:customer_id" json:"customerId"`                    // 归属客户ID
	CustomerAlias       string    `gorm:"column:customer_alias" json:"customerAlias"`              // 客户别名
	ErrorInfo           string    `gorm:"column:error_info" json:"errorInfo"`                      // 错误信息
	Status              int       `gorm:"column:status" json:"status"`                             // 状态 1 已处理 0 未处理
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`                    // 创建时间
	CreateUser          int       `gorm:"column:create_user" json:"createUser"`                    // 创建用户
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`                    // 修改时间
	UpdateUser          int       `gorm:"column:update_user" json:"updateUser"`                    // 更新时间
	Version             int       `gorm:"column:version" json:"version"`                           // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderAddressErrorTemp) TableName() string {
	return "lg_order_address_error_temp"
}

// LgOrderAddressErrorTempColumns get sql column name.获取数据库列名
var LgOrderAddressErrorTempColumns = struct {
	ID                  string
	OrderNumber         string
	ReceiveCity         string
	ReceiveProvince     string
	ReceiveCountry      string
	ReceiveZipCode      string
	CustomerChannelID   string
	CustomerChannelName string
	CustomerID          string
	CustomerAlias       string
	ErrorInfo           string
	Status              string
	CreateTime          string
	CreateUser          string
	UpdateTime          string
	UpdateUser          string
	Version             string
}{
	ID:                  "id",
	OrderNumber:         "order_number",
	ReceiveCity:         "receive_city",
	ReceiveProvince:     "receive_province",
	ReceiveCountry:      "receive_country",
	ReceiveZipCode:      "receive_zip_code",
	CustomerChannelID:   "customer_channel_id",
	CustomerChannelName: "customer_channel_name",
	CustomerID:          "customer_id",
	CustomerAlias:       "customer_alias",
	ErrorInfo:           "error_info",
	Status:              "status",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateTime:          "update_time",
	UpdateUser:          "update_user",
	Version:             "version",
}

// LgOrderCustomTemplateExport 订单自定义参数模板表
type LgOrderCustomTemplateExport struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`              // ID
	UserID        int       `gorm:"column:user_id" json:"userId"`               // 用户id
	CustomerID    int       `gorm:"column:customer_id" json:"customerId"`       // 客户id
	TemplateName  string    `gorm:"column:template_name" json:"templateName"`   // 模板名称
	TemplateField string    `gorm:"column:template_field" json:"templateField"` // 自定义字段模板
	Type          int       `gorm:"column:type" json:"type"`                    // 模板类型（0：页面模板，1：导出模板，3：客户顶带你自定义模板）
	Remarks       string    `gorm:"column:remarks" json:"remarks"`              // 备注信息
	CreateUser    int       `gorm:"column:create_user" json:"createUser"`       // 创建用户
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	UpdateUser    int       `gorm:"column:update_user" json:"updateUser"`       // 更新时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
	Version       int       `gorm:"column:version" json:"version"`              // 乐观锁
	IsDelete      int       `gorm:"column:is_delete" json:"isDelete"`           // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderCustomTemplateExport) TableName() string {
	return "lg_order_custom_template_export"
}

// LgOrderCustomTemplateExportColumns get sql column name.获取数据库列名
var LgOrderCustomTemplateExportColumns = struct {
	ID            string
	UserID        string
	CustomerID    string
	TemplateName  string
	TemplateField string
	Type          string
	Remarks       string
	CreateUser    string
	CreateTime    string
	UpdateUser    string
	UpdateTime    string
	Version       string
	IsDelete      string
}{
	ID:            "id",
	UserID:        "user_id",
	CustomerID:    "customer_id",
	TemplateName:  "template_name",
	TemplateField: "template_field",
	Type:          "type",
	Remarks:       "remarks",
	CreateUser:    "create_user",
	CreateTime:    "create_time",
	UpdateUser:    "update_user",
	UpdateTime:    "update_time",
	Version:       "version",
	IsDelete:      "is_delete",
}

// LgOrderDataPushLog 订单数据推送日志
type LgOrderDataPushLog struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`                             // 主键id
	ClientID             int       `gorm:"column:client_id" json:"clientId"`                          // 客户id
	ClientName           string    `gorm:"column:client_name" json:"clientName"`                      // 客户名称(客户别名)
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	LabelURL             string    `gorm:"column:label_url" json:"labelUrl"`                          // 面单url
	LabelURLFinal        string    `gorm:"column:label_url_final" json:"labelUrlFinal"`               // 最终面单地址
	ReceiptTime          time.Time `gorm:"column:receipt_time" json:"receiptTime"`                    // 入库时间
	SendTime             time.Time `gorm:"column:send_time" json:"sendTime"`                          // 出库时间
	CustomerChannelName  string    `gorm:"column:customer_channel_name" json:"customerChannelName"`   // 渠道名称
	Weight               float64   `gorm:"column:weight" json:"weight"`                               // 称重重量
	Node                 string    `gorm:"column:node" json:"node"`                                   // 推送节点(入库，复重，换单，出库)
	Success              []uint8   `gorm:"column:success" json:"success"`                             // 是否推送成功：0-失败，1-成功
	Frequency            int       `gorm:"column:frequency" json:"frequency"`                         // 推送次数
	PushResult           string    `gorm:"column:push_result" json:"pushResult"`                      // 推送结果
	RequestIP            string    `gorm:"column:request_ip" json:"requestIp"`                        // 请求地址
	RequestParams        string    `gorm:"column:request_params" json:"requestParams"`                // 请求报文
	ResponseParams       string    `gorm:"column:response_params" json:"responseParams"`              // 响应报文
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderDataPushLog) TableName() string {
	return "lg_order_data_push_log"
}

// LgOrderDataPushLogColumns get sql column name.获取数据库列名
var LgOrderDataPushLogColumns = struct {
	ID                   string
	ClientID             string
	ClientName           string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	LabelURL             string
	LabelURLFinal        string
	ReceiptTime          string
	SendTime             string
	CustomerChannelName  string
	Weight               string
	Node                 string
	Success              string
	Frequency            string
	PushResult           string
	RequestIP            string
	RequestParams        string
	ResponseParams       string
	CreateTime           string
	UpdateTime           string
}{
	ID:                   "id",
	ClientID:             "client_id",
	ClientName:           "client_name",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	LabelURL:             "label_url",
	LabelURLFinal:        "label_url_final",
	ReceiptTime:          "receipt_time",
	SendTime:             "send_time",
	CustomerChannelName:  "customer_channel_name",
	Weight:               "weight",
	Node:                 "node",
	Success:              "success",
	Frequency:            "frequency",
	PushResult:           "push_result",
	RequestIP:            "request_ip",
	RequestParams:        "request_params",
	ResponseParams:       "response_params",
	CreateTime:           "create_time",
	UpdateTime:           "update_time",
}

// LgOrderFinance 订单财务扩展表
type LgOrderFinance struct {
	OrderID                  int    `gorm:"primaryKey;column:order_id" json:"-"`                                // 订单ID
	CustomerBillBatchNumbers string `gorm:"column:customer_bill_batch_numbers" json:"customerBillBatchNumbers"` // 客户对账单批次号,多个批次号逗号分割
	CustomerReceiptNumbers   string `gorm:"column:customer_receipt_numbers" json:"customerReceiptNumbers"`      // 客户收款单批次号,多个批次号逗号分割
	PayableBillBatchNumbers  string `gorm:"column:payable_bill_batch_numbers" json:"payableBillBatchNumbers"`   // 应付对账单批次号,多个批次逗号分割
	PayablePaymentNumber     string `gorm:"column:payable_payment_number" json:"payablePaymentNumber"`          // 应付付款单批次号,多个批次逗号分割
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderFinance) TableName() string {
	return "lg_order_finance"
}

// LgOrderFinanceColumns get sql column name.获取数据库列名
var LgOrderFinanceColumns = struct {
	OrderID                  string
	CustomerBillBatchNumbers string
	CustomerReceiptNumbers   string
	PayableBillBatchNumbers  string
	PayablePaymentNumber     string
}{
	OrderID:                  "order_id",
	CustomerBillBatchNumbers: "customer_bill_batch_numbers",
	CustomerReceiptNumbers:   "customer_receipt_numbers",
	PayableBillBatchNumbers:  "payable_bill_batch_numbers",
	PayablePaymentNumber:     "payable_payment_number",
}

// LgOrderItem 订单详情
type LgOrderItem struct {
	ID                int       `gorm:"primaryKey;column:id" json:"-"`                       // ID
	OrderID           int       `gorm:"column:order_id" json:"orderId"`                      // 订单ID
	NameEn            string    `gorm:"column:name_en" json:"nameEn"`                        // 英文品名
	NameCn            string    `gorm:"column:name_cn" json:"nameCn"`                        // 中文品名
	DeclaredWeight    float64   `gorm:"column:declared_weight" json:"declaredWeight"`        // 申报重量
	DeclaredValue     float64   `gorm:"column:declared_value" json:"declaredValue"`          // 申报价值
	Quantity          int       `gorm:"column:quantity" json:"quantity"`                     // 数量
	Hscode            string    `gorm:"column:hscode" json:"hscode"`                         // HSCODE
	GoodsURL          string    `gorm:"column:goods_url" json:"goodsUrl"`                    // 商品链接
	Sku               string    `gorm:"column:sku" json:"sku"`                               // SKU
	Length            string    `gorm:"column:length" json:"length"`                         // 长
	Width             string    `gorm:"column:width" json:"width"`                           // 宽
	Height            string    `gorm:"column:height" json:"height"`                         // 高
	Material          int       `gorm:"column:material" json:"material"`                     // 材质
	Purpose           string    `gorm:"column:purpose" json:"purpose"`                       // 用途
	CaseNumber        string    `gorm:"column:case_number" json:"caseNumber"`                // 箱号
	PlatformProductID string    `gorm:"column:platform_product_id" json:"platformProductId"` // 平台产品id
	Specs             string    `gorm:"column:specs" json:"specs"`                           // 规格
	PlaceOrigin       string    `gorm:"column:place_origin" json:"placeOrigin"`              // 产地
	Currency          string    `gorm:"column:currency" json:"currency"`                     // 币种
	Company           string    `gorm:"column:company" json:"company"`                       // 单位
	PicturesURL       string    `gorm:"column:pictures_url" json:"picturesUrl"`              // 图片链接
	Brand             string    `gorm:"column:brand" json:"brand"`                           // 品牌
	HasBattery        int       `gorm:"column:has_battery" json:"hasBattery"`                // 是否带电
	Liquid            string    `gorm:"column:liquid" json:"liquid"`                         // 是否液体,1-液体，0-非液体
	Powder            string    `gorm:"column:powder" json:"powder"`                         // 是否粉末,1-粉末，0-非粉末
	Magnetic          string    `gorm:"column:magnetic" json:"magnetic"`                     // 是否带磁,1-带磁，0-不带磁
	DistributionInfo  string    `gorm:"column:distribution_info" json:"distributionInfo"`    // 配货信息
	BatteryType       string    `gorm:"column:battery_type" json:"batteryType"`              // 电池类型
	SalesPlatform     string    `gorm:"column:sales_platform" json:"salesPlatform"`          // 销售平台
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`                // 创建时间
	CreateUser        int       `gorm:"column:create_user" json:"createUser"`                // 创建用户
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`                // 更新时间
	UpdateUser        int       `gorm:"column:update_user" json:"updateUser"`                // 更新时间
	Version           int       `gorm:"column:version" json:"version"`                       // 乐观锁
	Deleted           int       `gorm:"column:deleted" json:"deleted"`                       // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderItem) TableName() string {
	return "lg_order_item"
}

// LgOrderItemColumns get sql column name.获取数据库列名
var LgOrderItemColumns = struct {
	ID                string
	OrderID           string
	NameEn            string
	NameCn            string
	DeclaredWeight    string
	DeclaredValue     string
	Quantity          string
	Hscode            string
	GoodsURL          string
	Sku               string
	Length            string
	Width             string
	Height            string
	Material          string
	Purpose           string
	CaseNumber        string
	PlatformProductID string
	Specs             string
	PlaceOrigin       string
	Currency          string
	Company           string
	PicturesURL       string
	Brand             string
	HasBattery        string
	Liquid            string
	Powder            string
	Magnetic          string
	DistributionInfo  string
	BatteryType       string
	SalesPlatform     string
	CreateTime        string
	CreateUser        string
	UpdateTime        string
	UpdateUser        string
	Version           string
	Deleted           string
}{
	ID:                "id",
	OrderID:           "order_id",
	NameEn:            "name_en",
	NameCn:            "name_cn",
	DeclaredWeight:    "declared_weight",
	DeclaredValue:     "declared_value",
	Quantity:          "quantity",
	Hscode:            "hscode",
	GoodsURL:          "goods_url",
	Sku:               "sku",
	Length:            "length",
	Width:             "width",
	Height:            "height",
	Material:          "material",
	Purpose:           "purpose",
	CaseNumber:        "case_number",
	PlatformProductID: "platform_product_id",
	Specs:             "specs",
	PlaceOrigin:       "place_origin",
	Currency:          "currency",
	Company:           "company",
	PicturesURL:       "pictures_url",
	Brand:             "brand",
	HasBattery:        "has_battery",
	Liquid:            "liquid",
	Powder:            "powder",
	Magnetic:          "magnetic",
	DistributionInfo:  "distribution_info",
	BatteryType:       "battery_type",
	SalesPlatform:     "sales_platform",
	CreateTime:        "create_time",
	CreateUser:        "create_user",
	UpdateTime:        "update_time",
	UpdateUser:        "update_user",
	Version:           "version",
	Deleted:           "deleted",
}

// LgOrderNumPool 单号池主表
type LgOrderNumPool struct {
	ID              int64     `gorm:"primaryKey;column:id" json:"-"`                  // 单号池ID
	Name            string    `gorm:"column:name" json:"name"`                        // 单号池名称
	LowestThreshold int64     `gorm:"column:lowest_threshold" json:"lowestThreshold"` // 最低阈值
	NoticeUserID    int64     `gorm:"column:notice_user_id" json:"noticeUserId"`      // 低于阈值通知人ID
	Remark          string    `gorm:"column:remark" json:"remark"`                    // 单号池备注
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`           // 创建用户
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`           // 更新时间
	UpdateUser      int       `gorm:"column:update_user" json:"updateUser"`           // 更新时间
	Version         int       `gorm:"column:version" json:"version"`                  // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                  // 逻辑删除
	Type            string    `gorm:"column:type" json:"type"`                        // 类型(系统,外部)
	Prefix          string    `gorm:"column:prefix" json:"prefix"`                    // 号段前缀
	Suffix          string    `gorm:"column:suffix" json:"suffix"`                    // 号段后缀
	InitNum         int64     `gorm:"column:init_num" json:"initNum"`                 // 初始化值
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderNumPool) TableName() string {
	return "lg_order_num_pool"
}

// LgOrderNumPoolColumns get sql column name.获取数据库列名
var LgOrderNumPoolColumns = struct {
	ID              string
	Name            string
	LowestThreshold string
	NoticeUserID    string
	Remark          string
	CreateTime      string
	CreateUser      string
	UpdateTime      string
	UpdateUser      string
	Version         string
	Deleted         string
	Type            string
	Prefix          string
	Suffix          string
	InitNum         string
}{
	ID:              "id",
	Name:            "name",
	LowestThreshold: "lowest_threshold",
	NoticeUserID:    "notice_user_id",
	Remark:          "remark",
	CreateTime:      "create_time",
	CreateUser:      "create_user",
	UpdateTime:      "update_time",
	UpdateUser:      "update_user",
	Version:         "version",
	Deleted:         "deleted",
	Type:            "type",
	Prefix:          "prefix",
	Suffix:          "suffix",
	InitNum:         "init_num",
}

// LgOrderNumPoolSub 单号池附属表:主要记录具体的可分配单号
type LgOrderNumPoolSub struct {
	ID                     int64     `gorm:"primaryKey;column:id" json:"-"`                                 // 单号池副表主键ID
	OrderNumPoolID         int64     `gorm:"column:order_num_pool_id" json:"orderNumPoolId"`                // 绑定的单号池主单ID
	OrderNumPoolName       string    `gorm:"column:order_num_pool_name" json:"orderNumPoolName"`            // 单号池名称关联主单名称
	TrackNo                string    `gorm:"column:track_no" json:"trackNo"`                                // 物流单号
	ExtractionOrderNum     string    `gorm:"column:extraction_order_num" json:"extractionOrderNum"`         // 提取单号:取号后回填客户单号
	ExtractionCustomerID   int64     `gorm:"column:extraction_customer_id" json:"extractionCustomerId"`     // 提取客户ID
	ExtractionCustomerName string    `gorm:"column:extraction_customer_name" json:"extractionCustomerName"` // 提取客户名称
	CreateTime             time.Time `gorm:"column:create_time" json:"createTime"`                          // 创建时间
	CreateUser             int       `gorm:"column:create_user" json:"createUser"`                          // 创建用户
	UpdateTime             time.Time `gorm:"column:update_time" json:"updateTime"`                          // 更新时间
	UpdateUser             int       `gorm:"column:update_user" json:"updateUser"`                          // 更新时间
	Version                int       `gorm:"column:version" json:"version"`                                 // 乐观锁
	Deleted                int       `gorm:"column:deleted" json:"deleted"`                                 // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderNumPoolSub) TableName() string {
	return "lg_order_num_pool_sub"
}

// LgOrderNumPoolSubColumns get sql column name.获取数据库列名
var LgOrderNumPoolSubColumns = struct {
	ID                     string
	OrderNumPoolID         string
	OrderNumPoolName       string
	TrackNo                string
	ExtractionOrderNum     string
	ExtractionCustomerID   string
	ExtractionCustomerName string
	CreateTime             string
	CreateUser             string
	UpdateTime             string
	UpdateUser             string
	Version                string
	Deleted                string
}{
	ID:                     "id",
	OrderNumPoolID:         "order_num_pool_id",
	OrderNumPoolName:       "order_num_pool_name",
	TrackNo:                "track_no",
	ExtractionOrderNum:     "extraction_order_num",
	ExtractionCustomerID:   "extraction_customer_id",
	ExtractionCustomerName: "extraction_customer_name",
	CreateTime:             "create_time",
	CreateUser:             "create_user",
	UpdateTime:             "update_time",
	UpdateUser:             "update_user",
	Version:                "version",
	Deleted:                "deleted",
}

// LgOrderNumSegment 号码段主表
type LgOrderNumSegment struct {
	ID               int64     `gorm:"primaryKey;column:id" json:"-"`                    // 号段主表ID
	Name             string    `gorm:"column:name" json:"name"`                          // 号段名称
	Prefix           string    `gorm:"column:prefix" json:"prefix"`                      // 号段前缀
	Suffix           string    `gorm:"column:suffix" json:"suffix"`                      // 号段后缀
	StartNum         int64     `gorm:"column:start_num" json:"startNum"`                 // 开始号码
	EndNum           int64     `gorm:"column:end_num" json:"endNum"`                     // 结束号码
	VerificationMode string    `gorm:"column:verification_mode" json:"verificationMode"` // 验证号码合法方式:保存验证方式code
	UsedNum          int64     `gorm:"column:used_num" json:"usedNum"`                   // 已用到号码段位置:默认喂号码开始位置
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`             // 创建用户
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`             // 更新时间
	Version          int       `gorm:"column:version" json:"version"`                    // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                    // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderNumSegment) TableName() string {
	return "lg_order_num_segment"
}

// LgOrderNumSegmentColumns get sql column name.获取数据库列名
var LgOrderNumSegmentColumns = struct {
	ID               string
	Name             string
	Prefix           string
	Suffix           string
	StartNum         string
	EndNum           string
	VerificationMode string
	UsedNum          string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	Name:             "name",
	Prefix:           "prefix",
	Suffix:           "suffix",
	StartNum:         "start_num",
	EndNum:           "end_num",
	VerificationMode: "verification_mode",
	UsedNum:          "used_num",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
}

// LgOrderNumSegmentSub 号码段附表:主要记录提取记录
type LgOrderNumSegmentSub struct {
	ID                int64             `gorm:"column:id" json:"id"`                                                                                         // 号码段附表ID
	OrderNumSegmentID int64             `gorm:"column:order_num_segment_id" json:"orderNumSegmentId"`                                                        // 号段主表关联ID
	LgOrderNumSegment LgOrderNumSegment `gorm:"joinForeignKey:order_num_segment_id;foreignKey:id;references:OrderNumSegmentID" json:"lgOrderNumSegmentList"` // 号码段主表
	StartNum          int64             `gorm:"column:start_num" json:"startNum"`                                                                            // 号段提取开始位置,数字部分
	EndNum            int64             `gorm:"column:end_num" json:"endNum"`                                                                                // 号段提取结束位置,数字部分,包头包尾
	Type              string            `gorm:"column:type" json:"type"`                                                                                     // 提取类型:self:自用 customer:客户使用
	Remark            string            `gorm:"column:remark" json:"remark"`                                                                                 // 备注
	CreateTime        time.Time         `gorm:"column:create_time" json:"createTime"`                                                                        // 创建时间
	CreateUser        int               `gorm:"column:create_user" json:"createUser"`                                                                        // 创建用户
	UpdateTime        time.Time         `gorm:"column:update_time" json:"updateTime"`                                                                        // 更新时间
	UpdateUser        int               `gorm:"column:update_user" json:"updateUser"`                                                                        // 更新时间
	Version           int               `gorm:"column:version" json:"version"`                                                                               // 乐观锁
	Deleted           int               `gorm:"column:deleted" json:"deleted"`                                                                               // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderNumSegmentSub) TableName() string {
	return "lg_order_num_segment_sub"
}

// LgOrderNumSegmentSubColumns get sql column name.获取数据库列名
var LgOrderNumSegmentSubColumns = struct {
	ID                string
	OrderNumSegmentID string
	StartNum          string
	EndNum            string
	Type              string
	Remark            string
	CreateTime        string
	CreateUser        string
	UpdateTime        string
	UpdateUser        string
	Version           string
	Deleted           string
}{
	ID:                "id",
	OrderNumSegmentID: "order_num_segment_id",
	StartNum:          "start_num",
	EndNum:            "end_num",
	Type:              "type",
	Remark:            "remark",
	CreateTime:        "create_time",
	CreateUser:        "create_user",
	UpdateTime:        "update_time",
	UpdateUser:        "update_user",
	Version:           "version",
	Deleted:           "deleted",
}

// LgOrderPlatformExpand 平台拓展订单表
type LgOrderPlatformExpand struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`                         // 主键id
	OrderID            int       `gorm:"primaryKey;column:order_id" json:"-"`                   //  订单id
	PlatformNumber     string    `gorm:"column:platform_number" json:"platformNumber"`          // 平台订单号
	PlatformCode       string    `gorm:"column:platform_code" json:"platformCode"`              // 平台代码
	PlatformName       string    `gorm:"column:platform_name" json:"platformName"`              // 平台名称
	PlatformSellerName string    `gorm:"column:platform_seller_name" json:"platformSellerName"` // 平台卖家名称
	PlatformOrderTime  time.Time `gorm:"column:platform_order_time" json:"platformOrderTime"`   // 平台订单时间
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                  //  创建时间
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                  //  创建用户
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`                  //  更新时间
	UpdateUser         int       `gorm:"column:update_user" json:"updateUser"`                  //  更新用户
	Version            uint      `gorm:"column:version" json:"version"`                         //  乐观锁
	Deleted            uint      `gorm:"column:deleted" json:"deleted"`                         //  逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderPlatformExpand) TableName() string {
	return "lg_order_platform_expand"
}

// LgOrderPlatformExpandColumns get sql column name.获取数据库列名
var LgOrderPlatformExpandColumns = struct {
	ID                 string
	OrderID            string
	PlatformNumber     string
	PlatformCode       string
	PlatformName       string
	PlatformSellerName string
	PlatformOrderTime  string
	CreateTime         string
	CreateUser         string
	UpdateTime         string
	UpdateUser         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	OrderID:            "order_id",
	PlatformNumber:     "platform_number",
	PlatformCode:       "platform_code",
	PlatformName:       "platform_name",
	PlatformSellerName: "platform_seller_name",
	PlatformOrderTime:  "platform_order_time",
	CreateTime:         "create_time",
	CreateUser:         "create_user",
	UpdateTime:         "update_time",
	UpdateUser:         "update_user",
	Version:            "version",
	Deleted:            "deleted",
}

// LgOrderProblemLog 订单问题件附表
type LgOrderProblemLog struct {
	ID                   int64     `gorm:"primaryKey;column:id" json:"-"`
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	CustomerID           int64     `gorm:"column:customer_id" json:"customerId"`                      // 归属客户ID
	CustomerAlias        string    `gorm:"column:customer_alias" json:"customerAlias"`                // 客户别名
	CustomerChannelID    int       `gorm:"column:customer_channel_id" json:"customerChannelId"`       // 客户渠道ID
	CustomerChannelName  string    `gorm:"column:customer_channel_name" json:"customerChannelName"`   // 渠道名称
	ProblemType          string    `gorm:"column:problem_type" json:"problemType"`                    // 问题件类型
	ProblemReason        string    `gorm:"column:problem_reason" json:"problemReason"`                // 问题件原因
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 创建用户
	CreateUserName       string    `gorm:"column:create_user_name" json:"createUserName"`             // 创建用户名称
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新用户
	UpdateUserName       string    `gorm:"column:update_user_name" json:"updateUserName"`             // 更新用户名称
	HandlerID            int       `gorm:"column:handler_id" json:"handlerId"`                        // 经办人id
	HandlerName          string    `gorm:"column:handler_name" json:"handlerName"`                    // 经办人名称
	HandleStatus         int       `gorm:"column:handle_status" json:"handleStatus"`                  // 处理状态 1:待处理 2:已处理
	OperationType        string    `gorm:"column:operation_type" json:"operationType"`                // 操作类型
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 说明
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	Deleted              int       `gorm:"column:deleted" json:"deleted"`                             // 逻辑删除
	HandleTime           time.Time `gorm:"column:handle_time" json:"handleTime"`                      // 处理时间
	HandleUser           int       `gorm:"column:handle_user" json:"handleUser"`                      // 处理用户
	HandleUserName       string    `gorm:"column:handle_user_name" json:"handleUserName"`             // 处理人名称
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderProblemLog) TableName() string {
	return "lg_order_problem_log"
}

// LgOrderProblemLogColumns get sql column name.获取数据库列名
var LgOrderProblemLogColumns = struct {
	ID                   string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	CustomerID           string
	CustomerAlias        string
	CustomerChannelID    string
	CustomerChannelName  string
	ProblemType          string
	ProblemReason        string
	CreateTime           string
	CreateUser           string
	CreateUserName       string
	UpdateTime           string
	UpdateUser           string
	UpdateUserName       string
	HandlerID            string
	HandlerName          string
	HandleStatus         string
	OperationType        string
	Remark               string
	Version              string
	Deleted              string
	HandleTime           string
	HandleUser           string
	HandleUserName       string
}{
	ID:                   "id",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	CustomerID:           "customer_id",
	CustomerAlias:        "customer_alias",
	CustomerChannelID:    "customer_channel_id",
	CustomerChannelName:  "customer_channel_name",
	ProblemType:          "problem_type",
	ProblemReason:        "problem_reason",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	CreateUserName:       "create_user_name",
	UpdateTime:           "update_time",
	UpdateUser:           "update_user",
	UpdateUserName:       "update_user_name",
	HandlerID:            "handler_id",
	HandlerName:          "handler_name",
	HandleStatus:         "handle_status",
	OperationType:        "operation_type",
	Remark:               "remark",
	Version:              "version",
	Deleted:              "deleted",
	HandleTime:           "handle_time",
	HandleUser:           "handle_user",
	HandleUserName:       "handle_user_name",
}

// LgOrderReturnStepLog 订单退件记录表
type LgOrderReturnStepLog struct {
	ID                   int64     `gorm:"primaryKey;column:id" json:"-"`
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	CustomerID           int64     `gorm:"column:customer_id" json:"customerId"`                      // 归属客户ID
	CustomerAlias        string    `gorm:"column:customer_alias" json:"customerAlias"`                // 客户别名
	CustomerChannelID    int       `gorm:"column:customer_channel_id" json:"customerChannelId"`       // 客户渠道ID
	CustomerChannelName  string    `gorm:"column:customer_channel_name" json:"customerChannelName"`   // 渠道名称
	ReceiveCountry       string    `gorm:"column:receive_country" json:"receiveCountry"`              // 收件人国家
	ReturnStepLocation   string    `gorm:"column:return_step_location" json:"returnStepLocation"`     // 退件地点
	ReturnStepReason     string    `gorm:"column:return_step_reason" json:"returnStepReason"`         // 退件原因
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 创建用户
	CreateUserName       string    `gorm:"column:create_user_name" json:"createUserName"`             // 创建用户名称
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新用户
	UpdateUserName       string    `gorm:"column:update_user_name" json:"updateUserName"`             // 更新用户名称
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	Deleted              int       `gorm:"column:deleted" json:"deleted"`                             // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderReturnStepLog) TableName() string {
	return "lg_order_return_step_log"
}

// LgOrderReturnStepLogColumns get sql column name.获取数据库列名
var LgOrderReturnStepLogColumns = struct {
	ID                   string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	CustomerID           string
	CustomerAlias        string
	CustomerChannelID    string
	CustomerChannelName  string
	ReceiveCountry       string
	ReturnStepLocation   string
	ReturnStepReason     string
	CreateTime           string
	CreateUser           string
	CreateUserName       string
	UpdateTime           string
	UpdateUser           string
	UpdateUserName       string
	Version              string
	Deleted              string
}{
	ID:                   "id",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	CustomerID:           "customer_id",
	CustomerAlias:        "customer_alias",
	CustomerChannelID:    "customer_channel_id",
	CustomerChannelName:  "customer_channel_name",
	ReceiveCountry:       "receive_country",
	ReturnStepLocation:   "return_step_location",
	ReturnStepReason:     "return_step_reason",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	CreateUserName:       "create_user_name",
	UpdateTime:           "update_time",
	UpdateUser:           "update_user",
	UpdateUserName:       "update_user_name",
	Version:              "version",
	Deleted:              "deleted",
}

// LgOrderStatistical 订单统计表
type LgOrderStatistical struct {
	Type       int            `gorm:"column:type" json:"type"`              // 统计类型
	Date       datatypes.Date `gorm:"column:date" json:"date"`              // 日期
	DateType   int            `gorm:"column:date_type" json:"dateType"`     // 时间类型
	Data       datatypes.JSON `gorm:"column:data" json:"data"`              // 订单统计数据
	CreateTime time.Time      `gorm:"column:create_time" json:"createTime"` // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderStatistical) TableName() string {
	return "lg_order_statistical"
}

// LgOrderStatisticalColumns get sql column name.获取数据库列名
var LgOrderStatisticalColumns = struct {
	Type       string
	Date       string
	DateType   string
	Data       string
	CreateTime string
}{
	Type:       "type",
	Date:       "date",
	DateType:   "date_type",
	Data:       "data",
	CreateTime: "create_time",
}

// LgOrderStatisticalCopy1 订单统计表
type LgOrderStatisticalCopy1 struct {
	Type              int            `gorm:"column:type" json:"type"`                             // 统计类型
	Date              datatypes.Date `gorm:"column:date" json:"date"`                             // 日期
	DateType          int            `gorm:"column:date_type" json:"dateType"`                    // 时间类型
	Data              datatypes.JSON `gorm:"column:data" json:"data"`                             // 订单统计数据
	CreateTime        time.Time      `gorm:"column:create_time" json:"createTime"`                // 创建时间
	CustomerChannelID int            `gorm:"column:customer_channel_id" json:"customerChannelId"` // 渠道id
	PlatformCode      string         `gorm:"column:platform_code" json:"platformCode"`            // 平台代码
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderStatisticalCopy1) TableName() string {
	return "lg_order_statistical_copy1"
}

// LgOrderStatisticalCopy1Columns get sql column name.获取数据库列名
var LgOrderStatisticalCopy1Columns = struct {
	Type              string
	Date              string
	DateType          string
	Data              string
	CreateTime        string
	CustomerChannelID string
	PlatformCode      string
}{
	Type:              "type",
	Date:              "date",
	DateType:          "date_type",
	Data:              "data",
	CreateTime:        "create_time",
	CustomerChannelID: "customer_channel_id",
	PlatformCode:      "platform_code",
}

// LgOrderTemp 物流订单
type LgOrderTemp struct {
	ID                       int            `gorm:"primaryKey;column:id" json:"-"`                                      // ID
	OrderNumber              string         `gorm:"column:order_number" json:"orderNumber"`                             // 订单号
	ReferenceNumber          string         `gorm:"column:reference_number" json:"referenceNumber"`                     // 参考号
	LogisticsNumber          string         `gorm:"column:logistics_number" json:"logisticsNumber"`                     // 物流单号
	LogisticsNumberFinal     string         `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"`          // 最终物流单号
	ExpressNum               string         `gorm:"column:express_num" json:"expressNum"`                               // 订单快递单号
	PlatformNumber           string         `gorm:"column:platform_number" json:"platformNumber"`                       // 平台订单
	ProviderOrderID          string         `gorm:"column:provider_order_id" json:"providerOrderId"`                    // 服务商系统订单ID
	LabelURL                 string         `gorm:"column:label_url" json:"labelUrl"`                                   // 面单地址
	LabelURLFinal            string         `gorm:"column:label_url_final" json:"labelUrlFinal"`                        // 最终面单地址
	OnlineNumber             string         `gorm:"column:online_number" json:"onlineNumber"`                           // 上网单号(17单号)
	BagNumber                string         `gorm:"column:bag_number" json:"bagNumber"`                                 // 袋号
	TotalListNo              string         `gorm:"column:total_list_no" json:"totalListNo"`                            // 总单号
	AirBillNumber            string         `gorm:"column:air_bill_number" json:"airBillNumber"`                        // 航空提单号
	CustomerChannelID        int            `gorm:"column:customer_channel_id" json:"customerChannelId"`                // 客户渠道ID
	CustomerChannelName      string         `gorm:"column:customer_channel_name" json:"customerChannelName"`            // 客户渠道名称
	ProviderCode             string         `gorm:"column:provider_code" json:"providerCode"`                           // 服务商系统code
	ProviderChannelCode      string         `gorm:"column:provider_channel_code" json:"providerChannelCode"`            // 服务商系统渠道Code
	ProviderChannelName      string         `gorm:"column:provider_Channel_Name" json:"providerChannelName"`            // 服务商渠道名称
	ReceiveCountry           string         `gorm:"column:receive_country" json:"receiveCountry"`                       // 收件人国家
	ReceiveName              string         `gorm:"column:receive_name" json:"receiveName"`                             // 收件人姓名
	ReceivePhone             string         `gorm:"column:receive_phone" json:"receivePhone"`                           // 收件人电话
	ReceiveMobile            string         `gorm:"column:receive_mobile" json:"receiveMobile"`                         // 收件人手机
	ReceiveTax               string         `gorm:"column:receive_tax" json:"receiveTax"`                               // 收件人税号
	ReceivePassport          string         `gorm:"column:receive_passport" json:"receivePassport"`                     // 收件人护照号
	ReceiveMail              string         `gorm:"column:receive_mail" json:"receiveMail"`                             // 收件人邮箱
	ReceiveCompany           string         `gorm:"column:receive_company" json:"receiveCompany"`                       // 收件人公司
	ReceiveZipcode           string         `gorm:"column:receive_zipcode" json:"receiveZipcode"`                       // 收件人邮编
	ReceiveProvince          string         `gorm:"column:receive_province" json:"receiveProvince"`                     // 收件人省份
	ReceiveHouseNumber       string         `gorm:"column:receive_house_number" json:"receiveHouseNumber"`              // 收件人门牌号
	ReceiveArea              string         `gorm:"column:receive_area" json:"receiveArea"`                             // 收件人区
	ReceiveStreet            string         `gorm:"column:receive_street" json:"receiveStreet"`                         // 收件人街道
	ReceiveCity              string         `gorm:"column:receive_city" json:"receiveCity"`                             // 收件人城市
	ReceiveAddress1          string         `gorm:"column:receive_address1" json:"receiveAddress1"`                     // 收件人地址1
	ReceiveAddress2          string         `gorm:"column:receive_address2" json:"receiveAddress2"`                     // 收件人地址2
	SenderCountry            string         `gorm:"column:sender_country" json:"senderCountry"`                         // 发件人国家
	SenderName               string         `gorm:"column:sender_name" json:"senderName"`                               // 发件人姓名
	SenderPhone              string         `gorm:"column:sender_phone" json:"senderPhone"`                             // 发件人电话
	SenderMobile             string         `gorm:"column:sender_mobile" json:"senderMobile"`                           // 发件人手机
	SenderTax                string         `gorm:"column:sender_tax" json:"senderTax"`                                 // 发件人税号
	SenderPassport           string         `gorm:"column:sender_passport" json:"senderPassport"`                       // 发件人护照号
	SenderMail               string         `gorm:"column:sender_mail" json:"senderMail"`                               // 发件人邮箱
	SenderCompany            string         `gorm:"column:sender_company" json:"senderCompany"`                         // 发件人公司
	SenderZipcode            string         `gorm:"column:sender_zipcode" json:"senderZipcode"`                         // 发件人邮编
	SenderProvince           string         `gorm:"column:sender_province" json:"senderProvince"`                       // 发件人省份
	SenderCity               string         `gorm:"column:sender_city" json:"senderCity"`                               // 发件人城市
	SenderArea               string         `gorm:"column:sender_area" json:"senderArea"`                               // 发件人区
	SenderStreet             string         `gorm:"column:sender_street" json:"senderStreet"`                           // 发件人街道
	SenderHouseNumber        string         `gorm:"column:sender_house_number" json:"senderHouseNumber"`                // 发件人门牌号
	SenderAddress            string         `gorm:"column:sender_address" json:"senderAddress"`                         // 发件人地址
	PackageType              string         `gorm:"column:package_type" json:"packageType"`                             // 包裹类型
	HasBattery               int            `gorm:"column:has_battery" json:"hasBattery"`                               // 是否带电
	BatteryType              string         `gorm:"column:battery_type" json:"batteryType"`                             // 电池类型
	HasBack                  int            `gorm:"column:has_back" json:"hasBack"`                                     // 是否退回
	Remarks                  string         `gorm:"column:remarks" json:"remarks"`                                      // 备注信息
	DistributionInformation  string         `gorm:"column:distribution_information" json:"distributionInformation"`     // 配货信息
	Insurance                string         `gorm:"column:insurance" json:"insurance"`                                  // 保险
	SalesPlatform            string         `gorm:"column:sales_platform" json:"salesPlatform"`                         // 销售平台
	Weight                   float64        `gorm:"column:weight" json:"weight"`                                        // 预报重量
	Quantity                 int            `gorm:"column:quantity" json:"quantity"`                                    // 包裹数量
	DeclaredValue            float64        `gorm:"column:declared_value" json:"declaredValue"`                         // 申报价值
	Status                   string         `gorm:"column:status" json:"status"`                                        // 订单状态
	TransportStatus          string         `gorm:"column:transport_status" json:"transportStatus"`                     // 转运状态
	LastTrackStatus          string         `gorm:"column:last_track_status" json:"lastTrackStatus"`                    // 最后一条轨迹状态
	LastTrackAddress         string         `gorm:"column:last_track_address" json:"lastTrackAddress"`                  // 最后一条轨迹
	LastTrackTime            time.Time      `gorm:"column:last_track_time" json:"lastTrackTime"`                        // 最后一条轨迹的时间
	ProblemType              string         `gorm:"column:problem_type" json:"problemType"`                             // 问题件类型
	ProblemReason            string         `gorm:"column:problem_reason" json:"problemReason"`                         // 问题件原因
	WeighingWeight           float64        `gorm:"column:weighing_weight" json:"weighingWeight"`                       // 称重重量
	VolumeWeight             float64        `gorm:"column:volume_weight" json:"volumeWeight"`                           // 材积重量
	ChargeWeight             float64        `gorm:"column:charge_weight" json:"chargeWeight"`                           // 收费重量
	StandardFee              float64        `gorm:"column:standard_fee" json:"standardFee"`                             // 标准费用
	Currency                 string         `gorm:"column:currency" json:"currency"`                                    // 币种
	BuyersFee                float64        `gorm:"column:buyers_fee" json:"buyersFee"`                                 // 下家费用
	BuyersWeight             float64        `gorm:"column:buyers_weight" json:"buyersWeight"`                           // 下家重量
	OtherFee                 float64        `gorm:"column:other_fee" json:"otherFee"`                                   // 其他费用
	CustomerID               int64          `gorm:"column:customer_id" json:"customerId"`                               // 归属客户ID
	CustomerName             string         `gorm:"column:customer_name" json:"customerName"`                           // 归属客户名称
	CustomerAlias            string         `gorm:"column:customer_alias" json:"customerAlias"`                         // 客户别名
	PrepaymentVat            string         `gorm:"column:prepayment_vat" json:"prepaymentVat"`                         // "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
	Vat                      string         `gorm:"column:vat" json:"vat"`                                              // vat号(英国方向必填)
	ReceiveCertificateType   string         `gorm:"column:receive_certificate_type" json:"receiveCertificateType"`      // 收件人证件类型
	ReceiveCertificateCode   string         `gorm:"column:receive_certificate_code" json:"receiveCertificateCode"`      // 收件人证件号码
	SenderCertificateType    string         `gorm:"column:sender_certificate_type" json:"senderCertificateType"`        // 发件人证件类型
	SenderCertificateCode    string         `gorm:"column:sender_certificate_code" json:"senderCertificateCode"`        // 发件人证件号码
	ServiceChannelCode       string         `gorm:"column:service_channel_code" json:"serviceChannelCode"`              // 服务商渠道代码
	DeliveryTerms            string         `gorm:"column:delivery_terms" json:"deliveryTerms"`                         // 贸易条款-(部分渠道巴西方向必填)- DDU DDP
	ReceiveCountryName       string         `gorm:"column:receive_country_name" json:"receiveCountryName"`              // 收件人国家中文
	SenderCountryName        string         `gorm:"column:sender_country_name" json:"senderCountryName"`                // 发件国家中文
	BillNumber               string         `gorm:"column:bill_number" json:"billNumber"`                               // 账单编号
	ApproximateCost          float64        `gorm:"column:approximate_cost" json:"approximateCost"`                     // 预报费用
	WeighingCost             float64        `gorm:"column:weighing_cost" json:"weighingCost"`                           // 称重费用
	PrepaidAmount            float64        `gorm:"column:prepaid_amount" json:"prepaidAmount"`                         // 预扣款金额
	ActualCost               float64        `gorm:"column:actual_cost" json:"actualCost"`                               // 实收费用
	PrimeCost                float64        `gorm:"column:prime_cost" json:"primeCost"`                                 // 成本
	OrderType                int            `gorm:"column:order_type" json:"orderType"`                                 // 订单类型(0-预报订单，1-物流订单)
	ExchangeOrder            []uint8        `gorm:"column:exchange_order" json:"exchangeOrder"`                         // 是否需要换单 0 不需要1 需要
	ReceiptTime              time.Time      `gorm:"column:receipt_time" json:"receiptTime"`                             // 入库时间
	ReceiptUser              int            `gorm:"column:receipt_user" json:"receiptUser"`                             // 入库人
	SendTime                 time.Time      `gorm:"column:send_time" json:"sendTime"`                                   // 出库时间
	SendUser                 int            `gorm:"column:send_user" json:"sendUser"`                                   // 出库人
	PicOssLink               string         `gorm:"column:pic_oss_link" json:"picOssLink"`                              // 包裹图片oss地址
	Test                     []uint8        `gorm:"column:test" json:"test"`                                            // 是否为测试订单
	OrderSource              int            `gorm:"column:order_source" json:"orderSource"`                             // 订单来源：1-默认，2-平台订单
	PlatformCode             string         `gorm:"column:platform_code" json:"platformCode"`                           // 平台代码
	PlatformName             string         `gorm:"column:platform_name" json:"platformName"`                           // 平台名称
	PlatformSellerName       string         `gorm:"column:platform_seller_name" json:"platformSellerName"`              // 平台卖家名称
	PlatformOrderTime        time.Time      `gorm:"column:platform_order_time" json:"platformOrderTime"`                // 平台订单时间
	PostOfficeBranchName     string         `gorm:"column:post_office_branch_name" json:"postOfficeBranchName"`         // 邮政局分行名称
	ExpressSendTime          time.Time      `gorm:"column:express_send_time" json:"expressSendTime"`                    // 快递寄送时间
	ExpressReceiptTime       time.Time      `gorm:"column:express_receipt_time" json:"expressReceiptTime"`              // 快递签收时间
	FlightDeparturesTime     time.Time      `gorm:"column:flight_departures_time" json:"flightDeparturesTime"`          // 航班起飞时间
	CreateTime               time.Time      `gorm:"column:create_time" json:"createTime"`                               // 创建时间
	DeliveredTime            time.Time      `gorm:"column:delivered_time" json:"deliveredTime"`                         // 签收时间
	FlightArrivalTime        time.Time      `gorm:"column:flight_arrival_time" json:"flightArrivalTime"`                // 航班落地时间
	Extra                    datatypes.JSON `gorm:"column:extra" json:"extra"`                                          // 额外信息
	CreateUser               int            `gorm:"column:create_user" json:"createUser"`                               // 创建用户
	Deleted                  int            `gorm:"column:deleted" json:"deleted"`                                      // 逻辑删除
	UpdateTime               time.Time      `gorm:"column:update_time" json:"updateTime"`                               // 更新时间
	Version                  int            `gorm:"column:version" json:"version"`                                      // 乐观锁
	UpdateUser               int            `gorm:"column:update_user" json:"updateUser"`                               // 更新时间
	ChannelRemarks           string         `gorm:"column:channel_remarks" json:"channelRemarks"`                       // 渠道备注
	FinanceRemarks           string         `gorm:"column:finance_remarks" json:"financeRemarks"`                       // 财务备注
	CustomerBillBatchNumbers string         `gorm:"column:customer_bill_batch_numbers" json:"customerBillBatchNumbers"` // 客户对账单批次号,多个批次号逗号分割
	CustomerReceiptNumbers   string         `gorm:"column:customer_receipt_numbers" json:"customerReceiptNumbers"`      // 客户收款单批次号,多个批次号逗号分割
	PayableBillBatchNumbers  string         `gorm:"column:payable_bill_batch_numbers" json:"payableBillBatchNumbers"`   // 应付对账单批次号,多个批次逗号分割
	PayablePaymentNumber     string         `gorm:"column:payable_payment_number" json:"payablePaymentNumber"`          // 应付付款单批次号,多个批次逗号分割
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderTemp) TableName() string {
	return "lg_order_temp"
}

// LgOrderTempColumns get sql column name.获取数据库列名
var LgOrderTempColumns = struct {
	ID                       string
	OrderNumber              string
	ReferenceNumber          string
	LogisticsNumber          string
	LogisticsNumberFinal     string
	ExpressNum               string
	PlatformNumber           string
	ProviderOrderID          string
	LabelURL                 string
	LabelURLFinal            string
	OnlineNumber             string
	BagNumber                string
	TotalListNo              string
	AirBillNumber            string
	CustomerChannelID        string
	CustomerChannelName      string
	ProviderCode             string
	ProviderChannelCode      string
	ProviderChannelName      string
	ReceiveCountry           string
	ReceiveName              string
	ReceivePhone             string
	ReceiveMobile            string
	ReceiveTax               string
	ReceivePassport          string
	ReceiveMail              string
	ReceiveCompany           string
	ReceiveZipcode           string
	ReceiveProvince          string
	ReceiveHouseNumber       string
	ReceiveArea              string
	ReceiveStreet            string
	ReceiveCity              string
	ReceiveAddress1          string
	ReceiveAddress2          string
	SenderCountry            string
	SenderName               string
	SenderPhone              string
	SenderMobile             string
	SenderTax                string
	SenderPassport           string
	SenderMail               string
	SenderCompany            string
	SenderZipcode            string
	SenderProvince           string
	SenderCity               string
	SenderArea               string
	SenderStreet             string
	SenderHouseNumber        string
	SenderAddress            string
	PackageType              string
	HasBattery               string
	BatteryType              string
	HasBack                  string
	Remarks                  string
	DistributionInformation  string
	Insurance                string
	SalesPlatform            string
	Weight                   string
	Quantity                 string
	DeclaredValue            string
	Status                   string
	TransportStatus          string
	LastTrackStatus          string
	LastTrackAddress         string
	LastTrackTime            string
	ProblemType              string
	ProblemReason            string
	WeighingWeight           string
	VolumeWeight             string
	ChargeWeight             string
	StandardFee              string
	Currency                 string
	BuyersFee                string
	BuyersWeight             string
	OtherFee                 string
	CustomerID               string
	CustomerName             string
	CustomerAlias            string
	PrepaymentVat            string
	Vat                      string
	ReceiveCertificateType   string
	ReceiveCertificateCode   string
	SenderCertificateType    string
	SenderCertificateCode    string
	ServiceChannelCode       string
	DeliveryTerms            string
	ReceiveCountryName       string
	SenderCountryName        string
	BillNumber               string
	ApproximateCost          string
	WeighingCost             string
	PrepaidAmount            string
	ActualCost               string
	PrimeCost                string
	OrderType                string
	ExchangeOrder            string
	ReceiptTime              string
	ReceiptUser              string
	SendTime                 string
	SendUser                 string
	PicOssLink               string
	Test                     string
	OrderSource              string
	PlatformCode             string
	PlatformName             string
	PlatformSellerName       string
	PlatformOrderTime        string
	PostOfficeBranchName     string
	ExpressSendTime          string
	ExpressReceiptTime       string
	FlightDeparturesTime     string
	CreateTime               string
	DeliveredTime            string
	FlightArrivalTime        string
	Extra                    string
	CreateUser               string
	Deleted                  string
	UpdateTime               string
	Version                  string
	UpdateUser               string
	ChannelRemarks           string
	FinanceRemarks           string
	CustomerBillBatchNumbers string
	CustomerReceiptNumbers   string
	PayableBillBatchNumbers  string
	PayablePaymentNumber     string
}{
	ID:                       "id",
	OrderNumber:              "order_number",
	ReferenceNumber:          "reference_number",
	LogisticsNumber:          "logistics_number",
	LogisticsNumberFinal:     "logistics_number_final",
	ExpressNum:               "express_num",
	PlatformNumber:           "platform_number",
	ProviderOrderID:          "provider_order_id",
	LabelURL:                 "label_url",
	LabelURLFinal:            "label_url_final",
	OnlineNumber:             "online_number",
	BagNumber:                "bag_number",
	TotalListNo:              "total_list_no",
	AirBillNumber:            "air_bill_number",
	CustomerChannelID:        "customer_channel_id",
	CustomerChannelName:      "customer_channel_name",
	ProviderCode:             "provider_code",
	ProviderChannelCode:      "provider_channel_code",
	ProviderChannelName:      "provider_Channel_Name",
	ReceiveCountry:           "receive_country",
	ReceiveName:              "receive_name",
	ReceivePhone:             "receive_phone",
	ReceiveMobile:            "receive_mobile",
	ReceiveTax:               "receive_tax",
	ReceivePassport:          "receive_passport",
	ReceiveMail:              "receive_mail",
	ReceiveCompany:           "receive_company",
	ReceiveZipcode:           "receive_zipcode",
	ReceiveProvince:          "receive_province",
	ReceiveHouseNumber:       "receive_house_number",
	ReceiveArea:              "receive_area",
	ReceiveStreet:            "receive_street",
	ReceiveCity:              "receive_city",
	ReceiveAddress1:          "receive_address1",
	ReceiveAddress2:          "receive_address2",
	SenderCountry:            "sender_country",
	SenderName:               "sender_name",
	SenderPhone:              "sender_phone",
	SenderMobile:             "sender_mobile",
	SenderTax:                "sender_tax",
	SenderPassport:           "sender_passport",
	SenderMail:               "sender_mail",
	SenderCompany:            "sender_company",
	SenderZipcode:            "sender_zipcode",
	SenderProvince:           "sender_province",
	SenderCity:               "sender_city",
	SenderArea:               "sender_area",
	SenderStreet:             "sender_street",
	SenderHouseNumber:        "sender_house_number",
	SenderAddress:            "sender_address",
	PackageType:              "package_type",
	HasBattery:               "has_battery",
	BatteryType:              "battery_type",
	HasBack:                  "has_back",
	Remarks:                  "remarks",
	DistributionInformation:  "distribution_information",
	Insurance:                "insurance",
	SalesPlatform:            "sales_platform",
	Weight:                   "weight",
	Quantity:                 "quantity",
	DeclaredValue:            "declared_value",
	Status:                   "status",
	TransportStatus:          "transport_status",
	LastTrackStatus:          "last_track_status",
	LastTrackAddress:         "last_track_address",
	LastTrackTime:            "last_track_time",
	ProblemType:              "problem_type",
	ProblemReason:            "problem_reason",
	WeighingWeight:           "weighing_weight",
	VolumeWeight:             "volume_weight",
	ChargeWeight:             "charge_weight",
	StandardFee:              "standard_fee",
	Currency:                 "currency",
	BuyersFee:                "buyers_fee",
	BuyersWeight:             "buyers_weight",
	OtherFee:                 "other_fee",
	CustomerID:               "customer_id",
	CustomerName:             "customer_name",
	CustomerAlias:            "customer_alias",
	PrepaymentVat:            "prepayment_vat",
	Vat:                      "vat",
	ReceiveCertificateType:   "receive_certificate_type",
	ReceiveCertificateCode:   "receive_certificate_code",
	SenderCertificateType:    "sender_certificate_type",
	SenderCertificateCode:    "sender_certificate_code",
	ServiceChannelCode:       "service_channel_code",
	DeliveryTerms:            "delivery_terms",
	ReceiveCountryName:       "receive_country_name",
	SenderCountryName:        "sender_country_name",
	BillNumber:               "bill_number",
	ApproximateCost:          "approximate_cost",
	WeighingCost:             "weighing_cost",
	PrepaidAmount:            "prepaid_amount",
	ActualCost:               "actual_cost",
	PrimeCost:                "prime_cost",
	OrderType:                "order_type",
	ExchangeOrder:            "exchange_order",
	ReceiptTime:              "receipt_time",
	ReceiptUser:              "receipt_user",
	SendTime:                 "send_time",
	SendUser:                 "send_user",
	PicOssLink:               "pic_oss_link",
	Test:                     "test",
	OrderSource:              "order_source",
	PlatformCode:             "platform_code",
	PlatformName:             "platform_name",
	PlatformSellerName:       "platform_seller_name",
	PlatformOrderTime:        "platform_order_time",
	PostOfficeBranchName:     "post_office_branch_name",
	ExpressSendTime:          "express_send_time",
	ExpressReceiptTime:       "express_receipt_time",
	FlightDeparturesTime:     "flight_departures_time",
	CreateTime:               "create_time",
	DeliveredTime:            "delivered_time",
	FlightArrivalTime:        "flight_arrival_time",
	Extra:                    "extra",
	CreateUser:               "create_user",
	Deleted:                  "deleted",
	UpdateTime:               "update_time",
	Version:                  "version",
	UpdateUser:               "update_user",
	ChannelRemarks:           "channel_remarks",
	FinanceRemarks:           "finance_remarks",
	CustomerBillBatchNumbers: "customer_bill_batch_numbers",
	CustomerReceiptNumbers:   "customer_receipt_numbers",
	PayableBillBatchNumbers:  "payable_bill_batch_numbers",
	PayablePaymentNumber:     "payable_payment_number",
}

// LgOrderVolumeWeight 材积重
type LgOrderVolumeWeight struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`                    // 主键id
	LogisticsNumber  string    `gorm:"column:logistics_number" json:"logisticsNumber"`   // 物流单号
	Length           float64   `gorm:"column:length" json:"length"`                      // 长
	Width            float64   `gorm:"column:width" json:"width"`                        // 宽
	Height           float64   `gorm:"column:height" json:"height"`                      // 高
	Weight           float64   `gorm:"column:weight" json:"weight"`                      // 称重重量
	VolumeWeight     float64   `gorm:"column:volume_weight" json:"volumeWeight"`         // 材积重
	Proportion       int       `gorm:"column:proportion" json:"proportion"`              // 抛比
	WeightDifference float64   `gorm:"column:weight_difference" json:"weightDifference"` // 重量差
	ProcessingStatus string    `gorm:"column:processing_status" json:"processingStatus"` // 处理状态，initial:待处理，done:已处理
	ProcessingUser   int       `gorm:"column:processing_user" json:"processingUser"`     // 处理状态人
	ProcessingTime   time.Time `gorm:"column:processing_time" json:"processingTime"`     // 处理时间
	Remark           string    `gorm:"column:remark" json:"remark"`                      // 说明
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`             // 创建人
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`             // 更新人
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
	Version          int       `gorm:"column:version" json:"version"`                    // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                    // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderVolumeWeight) TableName() string {
	return "lg_order_volume_weight"
}

// LgOrderVolumeWeightColumns get sql column name.获取数据库列名
var LgOrderVolumeWeightColumns = struct {
	ID               string
	LogisticsNumber  string
	Length           string
	Width            string
	Height           string
	Weight           string
	VolumeWeight     string
	Proportion       string
	WeightDifference string
	ProcessingStatus string
	ProcessingUser   string
	ProcessingTime   string
	Remark           string
	CreateUser       string
	CreateTime       string
	UpdateUser       string
	UpdateTime       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	LogisticsNumber:  "logistics_number",
	Length:           "length",
	Width:            "width",
	Height:           "height",
	Weight:           "weight",
	VolumeWeight:     "volume_weight",
	Proportion:       "proportion",
	WeightDifference: "weight_difference",
	ProcessingStatus: "processing_status",
	ProcessingUser:   "processing_user",
	ProcessingTime:   "processing_time",
	Remark:           "remark",
	CreateUser:       "create_user",
	CreateTime:       "create_time",
	UpdateUser:       "update_user",
	UpdateTime:       "update_time",
	Version:          "version",
	Deleted:          "deleted",
}

// LgOrderWmsOperationExpand 仓库订单操作表
type LgOrderWmsOperationExpand struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`              // 主键id
	OrderID       int       `gorm:"primaryKey;column:order_id" json:"-"`        //  订单id
	WarehouseID   int       `gorm:"column:warehouse_id" json:"warehouseId"`     //  仓库id
	WarehouseCode string    `gorm:"column:warehouse_code" json:"warehouseCode"` //  仓库code 默认code是没有
	WarehouseName string    `gorm:"column:warehouse_name" json:"warehouseName"` //  仓库名称 （默认仓库是没有仓库）
	ReceiptTime   time.Time `gorm:"column:receipt_time" json:"receiptTime"`     //  入库时间
	SendTime      time.Time `gorm:"column:send_time" json:"sendTime"`           //  出库时间
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       //  创建时间
	CreateUser    int       `gorm:"column:create_user" json:"createUser"`       //  创建用户
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       //  更新时间
	UpdateUser    int       `gorm:"column:update_user" json:"updateUser"`       //  更新用户
	Version       uint      `gorm:"column:version" json:"version"`              //  乐观锁
	Deleted       uint      `gorm:"column:deleted" json:"deleted"`              //  逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgOrderWmsOperationExpand) TableName() string {
	return "lg_order_wms_operation_expand"
}

// LgOrderWmsOperationExpandColumns get sql column name.获取数据库列名
var LgOrderWmsOperationExpandColumns = struct {
	ID            string
	OrderID       string
	WarehouseID   string
	WarehouseCode string
	WarehouseName string
	ReceiptTime   string
	SendTime      string
	CreateTime    string
	CreateUser    string
	UpdateTime    string
	UpdateUser    string
	Version       string
	Deleted       string
}{
	ID:            "id",
	OrderID:       "order_id",
	WarehouseID:   "warehouse_id",
	WarehouseCode: "warehouse_code",
	WarehouseName: "warehouse_name",
	ReceiptTime:   "receipt_time",
	SendTime:      "send_time",
	CreateTime:    "create_time",
	CreateUser:    "create_user",
	UpdateTime:    "update_time",
	UpdateUser:    "update_user",
	Version:       "version",
	Deleted:       "deleted",
}

// LgPlatformInfo 平台信息表
type LgPlatformInfo struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`            // 主键id
	PlatformCode string    `gorm:"column:platform_code" json:"platformCode"` // 平台代码
	PlatformName string    `gorm:"column:platform_name" json:"platformName"` // 平台名称
	Remark       string    `gorm:"column:remark" json:"remark"`              // 平台备注
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建人
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新人
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgPlatformInfo) TableName() string {
	return "lg_platform_info"
}

// LgPlatformInfoColumns get sql column name.获取数据库列名
var LgPlatformInfoColumns = struct {
	ID           string
	PlatformCode string
	PlatformName string
	Remark       string
	CreateTime   string
	CreateUser   string
	UpdateTime   string
	UpdateUser   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	PlatformCode: "platform_code",
	PlatformName: "platform_name",
	Remark:       "remark",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateTime:   "update_time",
	UpdateUser:   "update_user",
	Version:      "version",
	Deleted:      "deleted",
}

// LgProvider 服务商-下家
type LgProvider struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`            // 主键
	ProviderID   int       `gorm:"column:provider_id" json:"providerId"`     // 服务商code
	ProviderName string    `gorm:"column:provider_name" json:"providerName"` // 服务商名称
	Address      string    `gorm:"column:address" json:"address"`            // 地址
	Website      string    `gorm:"column:website" json:"website"`            // 主页
	Contact      string    `gorm:"column:contact" json:"contact"`            // 联系人
	Phone        string    `gorm:"column:phone" json:"phone"`                // 电话
	Email        string    `gorm:"column:email" json:"email"`                // 邮箱
	ClearingBank string    `gorm:"column:clearing_bank" json:"clearingBank"` // 结算银行
	BankAccount  string    `gorm:"column:bank_account" json:"bankAccount"`   // 银行账户
	State        int       `gorm:"column:state" json:"state"`                // 状态
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建用户
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新时间
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgProvider) TableName() string {
	return "lg_provider"
}

// LgProviderColumns get sql column name.获取数据库列名
var LgProviderColumns = struct {
	ID           string
	ProviderID   string
	ProviderName string
	Address      string
	Website      string
	Contact      string
	Phone        string
	Email        string
	ClearingBank string
	BankAccount  string
	State        string
	CreateTime   string
	CreateUser   string
	UpdateTime   string
	UpdateUser   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	ProviderID:   "provider_id",
	ProviderName: "provider_name",
	Address:      "address",
	Website:      "website",
	Contact:      "contact",
	Phone:        "phone",
	Email:        "email",
	ClearingBank: "clearing_bank",
	BankAccount:  "bank_account",
	State:        "state",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateTime:   "update_time",
	UpdateUser:   "update_user",
	Version:      "version",
	Deleted:      "deleted",
}

// LgProviderChannel 服务商-渠道
type LgProviderChannel struct {
	ID                  int            `gorm:"primaryKey;column:id" json:"-"`                                                                      // 主键
	ProviderID          string         `gorm:"primaryKey;column:provider_id" json:"-"`                                                             // 系统服务商id
	ProviderChannelCode string         `gorm:"primaryKey;column:provider_channel_code" json:"-"`                                                   // 渠道代码
	ProviderChannelName string         `gorm:"column:provider_channel_name" json:"providerChannelName"`                                            // 服务商渠道名称
	OrderNumPoolID      int64          `gorm:"column:order_num_pool_id" json:"orderNumPoolId"`                                                     // 单号池id
	LgOrderNumPool      LgOrderNumPool `gorm:"joinForeignKey:order_num_pool_id;foreignKey:id;references:OrderNumPoolID" json:"lgOrderNumPoolList"` // 单号池主表
	OrderNumPoolName    string         `gorm:"column:order_num_pool_name" json:"orderNumPoolName"`                                                 // 单号池名称
	BusinessEntityName  string         `gorm:"column:business_entity_name" json:"businessEntityName"`                                              // 业务主体代码code
	BusinessEntityCode  string         `gorm:"column:business_entity_code" json:"businessEntityCode"`                                              // 业务主体名称
	SyncLabel           []uint8        `gorm:"column:sync_label" json:"syncLabel"`
	VariableList        string         `gorm:"column:variable_list" json:"variableList"` // 变量列表
	CheckRule           string         `gorm:"column:check_rule" json:"checkRule"`       // 字段校验规则(正则表达式)
	Status              int            `gorm:"column:status" json:"status"`              // 0:无效，1:有效
	TakeNoMode          int            `gorm:"column:take_no_mode" json:"takeNoMode"`    // 取号方式,1:api 2:单号池
	CreateTime          time.Time      `gorm:"column:create_time" json:"createTime"`     // 创建时间
	CreateUser          int            `gorm:"column:create_user" json:"createUser"`     // 创建用户
	UpdateTime          time.Time      `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	UpdateUser          int            `gorm:"column:update_user" json:"updateUser"`     // 更新时间
	Version             int            `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted             int            `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgProviderChannel) TableName() string {
	return "lg_provider_channel"
}

// LgProviderChannelColumns get sql column name.获取数据库列名
var LgProviderChannelColumns = struct {
	ID                  string
	ProviderID          string
	ProviderChannelCode string
	ProviderChannelName string
	OrderNumPoolID      string
	OrderNumPoolName    string
	BusinessEntityName  string
	BusinessEntityCode  string
	SyncLabel           string
	VariableList        string
	CheckRule           string
	Status              string
	TakeNoMode          string
	CreateTime          string
	CreateUser          string
	UpdateTime          string
	UpdateUser          string
	Version             string
	Deleted             string
}{
	ID:                  "id",
	ProviderID:          "provider_id",
	ProviderChannelCode: "provider_channel_code",
	ProviderChannelName: "provider_channel_name",
	OrderNumPoolID:      "order_num_pool_id",
	OrderNumPoolName:    "order_num_pool_name",
	BusinessEntityName:  "business_entity_name",
	BusinessEntityCode:  "business_entity_code",
	SyncLabel:           "sync_label",
	VariableList:        "variable_list",
	CheckRule:           "check_rule",
	Status:              "status",
	TakeNoMode:          "take_no_mode",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateTime:          "update_time",
	UpdateUser:          "update_user",
	Version:             "version",
	Deleted:             "deleted",
}

// LgProviderChannelCheck 服务商渠道规则校验
type LgProviderChannelCheck struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`                 // ID
	ChannelFieldID int       `gorm:"column:channel_field_id" json:"channelFieldId"` // 渠道字段id
	IsRequired     int       `gorm:"column:is_required" json:"isRequired"`          // 是否必填
	Rule           string    `gorm:"column:rule" json:"rule"`                       // 限制规则
	ErrorMessage   string    `gorm:"column:error_message" json:"errorMessage"`      // 错误消息提示
	Remarks        string    `gorm:"column:remarks" json:"remarks"`                 // 备注
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`          // 创建时间
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`          // 创建用户
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`          // 更新时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`          // 更新时间
	Version        int       `gorm:"column:version" json:"version"`                 // 乐观锁
	Deleted        int       `gorm:"column:deleted" json:"deleted"`                 // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgProviderChannelCheck) TableName() string {
	return "lg_provider_channel_check"
}

// LgProviderChannelCheckColumns get sql column name.获取数据库列名
var LgProviderChannelCheckColumns = struct {
	ID             string
	ChannelFieldID string
	IsRequired     string
	Rule           string
	ErrorMessage   string
	Remarks        string
	CreateTime     string
	CreateUser     string
	UpdateTime     string
	UpdateUser     string
	Version        string
	Deleted        string
}{
	ID:             "id",
	ChannelFieldID: "channel_field_id",
	IsRequired:     "is_required",
	Rule:           "rule",
	ErrorMessage:   "error_message",
	Remarks:        "remarks",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	Version:        "version",
	Deleted:        "deleted",
}

// LgProviderChannelCountry 服务商支持国家
type LgProviderChannelCountry struct {
	ID                  int       `gorm:"primaryKey;column:id" json:"-"`                            // ID
	ProviderChannelID   int       `gorm:"column:provider_channel_id" json:"providerChannelId"`      // 服务商渠道ID
	ProviderChannelCode string    `gorm:"column:provider_channel_code" json:"providerChannelCode"`  // 服务商渠道code
	CountryCode         string    `gorm:"column:country_code" json:"countryCode"`                   // 国家二字代码
	Zipcodes            string    `gorm:"column:zipcodes" json:"zipcodes"`                          // 黑名单邮编
	SenderAddressBookID int       `gorm:"column:sender_address_book_id" json:"senderAddressBookId"` // 发件人地址簿id
	CreateTime          time.Time `gorm:"column:create_time" json:"createTime"`                     // 创建时间
	CreateUser          int       `gorm:"column:create_user" json:"createUser"`                     // 创建用户
	UpdateTime          time.Time `gorm:"column:update_time" json:"updateTime"`                     // 更新时间
	UpdateUser          int       `gorm:"column:update_user" json:"updateUser"`                     // 更新时间
	Version             int       `gorm:"column:version" json:"version"`                            // 乐观锁
	Deleted             int       `gorm:"column:deleted" json:"deleted"`                            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgProviderChannelCountry) TableName() string {
	return "lg_provider_channel_country"
}

// LgProviderChannelCountryColumns get sql column name.获取数据库列名
var LgProviderChannelCountryColumns = struct {
	ID                  string
	ProviderChannelID   string
	ProviderChannelCode string
	CountryCode         string
	Zipcodes            string
	SenderAddressBookID string
	CreateTime          string
	CreateUser          string
	UpdateTime          string
	UpdateUser          string
	Version             string
	Deleted             string
}{
	ID:                  "id",
	ProviderChannelID:   "provider_channel_id",
	ProviderChannelCode: "provider_channel_code",
	CountryCode:         "country_code",
	Zipcodes:            "zipcodes",
	SenderAddressBookID: "sender_address_book_id",
	CreateTime:          "create_time",
	CreateUser:          "create_user",
	UpdateTime:          "update_time",
	UpdateUser:          "update_user",
	Version:             "version",
	Deleted:             "deleted",
}

// LgProviderChannelField 服务商渠道公共字段
type LgProviderChannelField struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // ID
	Type       string    `gorm:"column:type" json:"type"`              // 所属类型
	Name       string    `gorm:"column:name" json:"name"`              // 字段名
	NameCn     string    `gorm:"column:name_cn" json:"nameCn"`         // 字段描述
	Remarks    string    `gorm:"column:remarks" json:"remarks"`        // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgProviderChannelField) TableName() string {
	return "lg_provider_channel_field"
}

// LgProviderChannelFieldColumns get sql column name.获取数据库列名
var LgProviderChannelFieldColumns = struct {
	ID         string
	Type       string
	Name       string
	NameCn     string
	Remarks    string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	Type:       "type",
	Name:       "name",
	NameCn:     "name_cn",
	Remarks:    "remarks",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// LgProviderChannelRule 服务商渠道规则
type LgProviderChannelRule struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`                 // ID
	ChannelID      int       `gorm:"column:channel_id" json:"channelId"`            // 服务商渠道ID
	CountryTwoCode string    `gorm:"column:country_two_code" json:"countryTwoCode"` // 国家二字代码
	ChannelCode    string    `gorm:"column:channel_code" json:"channelCode"`        // 服务商渠道Code
	LabelCode      string    `gorm:"column:label_code" json:"labelCode"`            // 面单代码
	StartZipcode   int       `gorm:"column:start_zipcode" json:"startZipcode"`      // 开始邮编
	EndZipcode     int       `gorm:"column:end_zipcode" json:"endZipcode"`          // 结束邮编
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`          // 创建时间
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`          // 创建用户
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`          // 更新时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`          // 更新时间
	Version        int       `gorm:"column:version" json:"version"`                 // 乐观锁
	Deleted        int       `gorm:"column:deleted" json:"deleted"`                 // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgProviderChannelRule) TableName() string {
	return "lg_provider_channel_rule"
}

// LgProviderChannelRuleColumns get sql column name.获取数据库列名
var LgProviderChannelRuleColumns = struct {
	ID             string
	ChannelID      string
	CountryTwoCode string
	ChannelCode    string
	LabelCode      string
	StartZipcode   string
	EndZipcode     string
	CreateTime     string
	CreateUser     string
	UpdateTime     string
	UpdateUser     string
	Version        string
	Deleted        string
}{
	ID:             "id",
	ChannelID:      "channel_id",
	CountryTwoCode: "country_two_code",
	ChannelCode:    "channel_code",
	LabelCode:      "label_code",
	StartZipcode:   "start_zipcode",
	EndZipcode:     "end_zipcode",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	Version:        "version",
	Deleted:        "deleted",
}

// LgProviderSystem 系统服务商
type LgProviderSystem struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`            // ID
	Name         string    `gorm:"column:name" json:"name"`                  // 系统名称
	Code         string    `gorm:"primaryKey;column:code" json:"-"`          // 系统CODE
	CheckRule    string    `gorm:"column:check_rule" json:"checkRule"`       // 字段校验规则(正则表达式)
	VariableList string    `gorm:"column:variable_list" json:"variableList"` // 下单字段列表
	Content      string    `gorm:"column:content" json:"content"`            // 附件及留言内容
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建用户
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新时间
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgProviderSystem) TableName() string {
	return "lg_provider_system"
}

// LgProviderSystemColumns get sql column name.获取数据库列名
var LgProviderSystemColumns = struct {
	ID           string
	Name         string
	Code         string
	CheckRule    string
	VariableList string
	Content      string
	CreateTime   string
	CreateUser   string
	UpdateTime   string
	UpdateUser   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	Name:         "name",
	Code:         "code",
	CheckRule:    "check_rule",
	VariableList: "variable_list",
	Content:      "content",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateTime:   "update_time",
	UpdateUser:   "update_user",
	Version:      "version",
	Deleted:      "deleted",
}

// LgRegional 大区表
type LgRegional struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`
	AreaID      int       `gorm:"column:area_id" json:"areaId"`            // 地区id
	AreaNameCn  string    `gorm:"column:area_name_cn" json:"areaNameCn"`   // 地区中文名
	AreaTwoCode string    `gorm:"column:area_two_code" json:"areaTwoCode"` // 地区二字码
	NameCn      string    `gorm:"column:name_cn" json:"nameCn"`            // 中文名
	NameEn      string    `gorm:"column:name_en" json:"nameEn"`            // 英文名
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`    // 创建时间
	CreateUser  int       `gorm:"column:create_user" json:"createUser"`    // 创建用户
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`    // 更新时间
	UpdateUser  int       `gorm:"column:update_user" json:"updateUser"`    // 更新用户
	Version     int       `gorm:"column:version" json:"version"`           // 乐观锁
	Deleted     int       `gorm:"column:deleted" json:"deleted"`           // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgRegional) TableName() string {
	return "lg_regional"
}

// LgRegionalColumns get sql column name.获取数据库列名
var LgRegionalColumns = struct {
	ID          string
	AreaID      string
	AreaNameCn  string
	AreaTwoCode string
	NameCn      string
	NameEn      string
	CreateTime  string
	CreateUser  string
	UpdateTime  string
	UpdateUser  string
	Version     string
	Deleted     string
}{
	ID:          "id",
	AreaID:      "area_id",
	AreaNameCn:  "area_name_cn",
	AreaTwoCode: "area_two_code",
	NameCn:      "name_cn",
	NameEn:      "name_en",
	CreateTime:  "create_time",
	CreateUser:  "create_user",
	UpdateTime:  "update_time",
	UpdateUser:  "update_user",
	Version:     "version",
	Deleted:     "deleted",
}

// LgRiskControlOrder 风控物流订单
type LgRiskControlOrder struct {
	ID                      int       `gorm:"primaryKey;column:id" json:"-"`                                  // ID
	OrderNumber             string    `gorm:"column:order_number" json:"orderNumber"`                         // 订单号
	ReferenceNumber         string    `gorm:"column:reference_number" json:"referenceNumber"`                 // 参考号
	LogisticsNumber         string    `gorm:"column:logistics_number" json:"logisticsNumber"`                 // 物流单号
	LabelURL                string    `gorm:"column:label_url" json:"labelUrl"`                               // 面单地址
	OnlineNumber            string    `gorm:"column:online_number" json:"onlineNumber"`                       // 上网单号(17单号)
	BagNumber               string    `gorm:"column:bag_number" json:"bagNumber"`                             // 袋号
	AirBillNumber           string    `gorm:"column:air_bill_number" json:"airBillNumber"`                    // 航空提单号
	CustomerChannelID       int       `gorm:"column:customer_channel_id" json:"customerChannelId"`            // 客户渠道ID
	CustomerChannelName     string    `gorm:"column:customer_channel_name" json:"customerChannelName"`        // 客户渠道名称
	ReceiveCountry          string    `gorm:"column:receive_country" json:"receiveCountry"`                   // 收件人国家
	ReceiveName             string    `gorm:"column:receive_name" json:"receiveName"`                         // 收件人姓名
	ReceivePhone            string    `gorm:"column:receive_phone" json:"receivePhone"`                       // 收件人电话
	ReceiveMobile           string    `gorm:"column:receive_mobile" json:"receiveMobile"`                     // 收件人手机
	ReceiveTax              string    `gorm:"column:receive_tax" json:"receiveTax"`                           // 收件人税号
	ReceivePassport         string    `gorm:"column:receive_passport" json:"receivePassport"`                 // 收件人护照号
	ReceiveMail             string    `gorm:"column:receive_mail" json:"receiveMail"`                         // 收件人邮箱
	ReceiveCompany          string    `gorm:"column:receive_company" json:"receiveCompany"`                   // 收件人公司
	ReceiveZipcode          string    `gorm:"column:receive_zipcode" json:"receiveZipcode"`                   // 收件人邮编
	ReceiveProvince         string    `gorm:"column:receive_province" json:"receiveProvince"`                 // 收件人省份
	ReceiveCity             string    `gorm:"column:receive_city" json:"receiveCity"`                         // 收件人城市
	ReceiveArea             string    `gorm:"column:receive_area" json:"receiveArea"`                         // 收件人区
	ReceiveStreet           string    `gorm:"column:receive_street" json:"receiveStreet"`                     // 收件人街道
	ReceiveHouseNumber      string    `gorm:"column:receive_house_number" json:"receiveHouseNumber"`          // 收件人门牌号
	ReceiveAddress1         string    `gorm:"column:receive_address1" json:"receiveAddress1"`                 // 收件人地址1
	ReceiveAddress2         string    `gorm:"column:receive_address2" json:"receiveAddress2"`                 // 收件人地址2
	SenderCountry           string    `gorm:"column:sender_country" json:"senderCountry"`                     // 发件人国家
	SenderName              string    `gorm:"column:sender_name" json:"senderName"`                           // 发件人姓名
	SenderPhone             string    `gorm:"column:sender_phone" json:"senderPhone"`                         // 发件人电话
	SenderMobile            string    `gorm:"column:sender_mobile" json:"senderMobile"`                       // 发件人手机
	SenderTax               string    `gorm:"column:sender_tax" json:"senderTax"`                             // 发件人税号
	SenderPassport          string    `gorm:"column:sender_passport" json:"senderPassport"`                   // 发件人护照号
	SenderMail              string    `gorm:"column:sender_mail" json:"senderMail"`                           // 发件人邮箱
	SenderCompany           string    `gorm:"column:sender_company" json:"senderCompany"`                     // 发件人公司
	SenderZipcode           string    `gorm:"column:sender_zipcode" json:"senderZipcode"`                     // 发件人邮编
	SenderProvince          string    `gorm:"column:sender_province" json:"senderProvince"`                   // 发件人省份
	SenderCity              string    `gorm:"column:sender_city" json:"senderCity"`                           // 发件人城市
	SenderArea              string    `gorm:"column:sender_area" json:"senderArea"`                           // 发件人区
	SenderStreet            string    `gorm:"column:sender_street" json:"senderStreet"`                       // 发件人街道
	SenderHouseNumber       string    `gorm:"column:sender_house_number" json:"senderHouseNumber"`            // 发件人门牌号
	SenderAddress           string    `gorm:"column:sender_address" json:"senderAddress"`                     // 发件人地址
	PackageType             string    `gorm:"column:package_type" json:"packageType"`                         // 包裹类型
	HasBattery              int       `gorm:"column:has_battery" json:"hasBattery"`                           // 是否带电
	BatteryType             string    `gorm:"column:battery_type" json:"batteryType"`                         // 电池类型
	HasBack                 int       `gorm:"column:has_back" json:"hasBack"`                                 // 是否退回
	Remarks                 string    `gorm:"column:remarks" json:"remarks"`                                  // 备注信息
	DistributionInformation string    `gorm:"column:distribution_information" json:"distributionInformation"` // 配货信息
	Insurance               string    `gorm:"column:insurance" json:"insurance"`                              // 保险
	SalesPlatform           string    `gorm:"column:sales_platform" json:"salesPlatform"`                     // 销售平台
	Weight                  float64   `gorm:"column:weight" json:"weight"`                                    // 预报重量
	Quantity                int       `gorm:"column:quantity" json:"quantity"`                                // 包裹数量
	DeclaredValue           float64   `gorm:"column:declared_value" json:"declaredValue"`                     // 申报价值
	Status                  string    `gorm:"column:status" json:"status"`                                    // 订单状态
	CreateTime              time.Time `gorm:"column:create_time" json:"createTime"`                           // 创建时间
	WeighingWeight          float64   `gorm:"column:weighing_weight" json:"weighingWeight"`                   // 称重重量
	VolumeWeight            float64   `gorm:"column:volume_weight" json:"volumeWeight"`                       // 材积重量
	ChargeWeight            float64   `gorm:"column:charge_weight" json:"chargeWeight"`                       // 收费重量
	StandardFee             float64   `gorm:"column:standard_fee" json:"standardFee"`                         // 标准费用
	Currency                string    `gorm:"column:currency" json:"currency"`                                // 币种
	BuyersFee               float64   `gorm:"column:buyers_fee" json:"buyersFee"`                             // 下家费用
	BuyersWeight            float64   `gorm:"column:buyers_weight" json:"buyersWeight"`                       // 下家重量
	OtherFee                float64   `gorm:"column:other_fee" json:"otherFee"`                               // 其他费用
	CustomerID              int64     `gorm:"column:customer_id" json:"customerId"`                           // 归属客户ID
	CustomerName            string    `gorm:"column:customer_name" json:"customerName"`                       // 归属客户名称
	PrepaymentVat           string    `gorm:"column:prepayment_vat" json:"prepaymentVat"`                     // "预缴税方式(0: IOSS 1: no-IOSS 2: other)-欧洲区国家必填
	Vat                     string    `gorm:"column:vat" json:"vat"`                                          // vat号(英国方向必填)
	ReceiveCertificateType  string    `gorm:"column:receive_certificate_type" json:"receiveCertificateType"`  // 收件人证件类型
	ReceiveCertificateCode  string    `gorm:"column:receive_certificate_code" json:"receiveCertificateCode"`  // 收件人证件号码
	SenderCertificateType   string    `gorm:"column:sender_certificate_type" json:"senderCertificateType"`    // 发件人证件类型
	SenderCertificateCode   string    `gorm:"column:sender_certificate_code" json:"senderCertificateCode"`    // 发件人证件号码
	ServiceChannelCode      string    `gorm:"column:service_channel_code" json:"serviceChannelCode"`          // api渠道代码
	DeliveryTerms           string    `gorm:"column:delivery_terms" json:"deliveryTerms"`                     // 贸易条款- DDU DDP
	ReceiveCountryName      string    `gorm:"column:receive_country_name" json:"receiveCountryName"`          // 收件人国家中文
	SenderCountryName       string    `gorm:"column:sender_country_name" json:"senderCountryName"`            // 发件国家中文
	OrderType               int       `gorm:"column:order_type" json:"orderType"`                             // 订单类型(0-预报订单，1-物流订单)
	ReceiptTime             time.Time `gorm:"column:receipt_time" json:"receiptTime"`                         // 入库时间
	ReceiptUser             int       `gorm:"column:receipt_user" json:"receiptUser"`                         // 入库人
	SendTime                time.Time `gorm:"column:send_time" json:"sendTime"`                               // 出库时间
	SendUser                int       `gorm:"column:send_user" json:"sendUser"`                               // 出库人
	PicOssLink              string    `gorm:"column:pic_oss_link" json:"picOssLink"`                          // 包裹图片oss地址
	RiskMessage             string    `gorm:"column:risk_message" json:"riskMessage"`                         // 风控信息
	StatusRisk              string    `gorm:"column:status_risk" json:"statusRisk"`                           // 风控订单状态:normal:正常,abnormal:异常 warn:警告
	CreateUser              int       `gorm:"column:create_user" json:"createUser"`                           // 创建用户
	UpdateTime              time.Time `gorm:"column:update_time" json:"updateTime"`                           // 更新时间
	UpdateUser              int       `gorm:"column:update_user" json:"updateUser"`                           // 更新时间
	Version                 int       `gorm:"column:version" json:"version"`                                  // 乐观锁
	Deleted                 int       `gorm:"column:deleted" json:"deleted"`                                  // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgRiskControlOrder) TableName() string {
	return "lg_risk_control_order"
}

// LgRiskControlOrderColumns get sql column name.获取数据库列名
var LgRiskControlOrderColumns = struct {
	ID                      string
	OrderNumber             string
	ReferenceNumber         string
	LogisticsNumber         string
	LabelURL                string
	OnlineNumber            string
	BagNumber               string
	AirBillNumber           string
	CustomerChannelID       string
	CustomerChannelName     string
	ReceiveCountry          string
	ReceiveName             string
	ReceivePhone            string
	ReceiveMobile           string
	ReceiveTax              string
	ReceivePassport         string
	ReceiveMail             string
	ReceiveCompany          string
	ReceiveZipcode          string
	ReceiveProvince         string
	ReceiveCity             string
	ReceiveArea             string
	ReceiveStreet           string
	ReceiveHouseNumber      string
	ReceiveAddress1         string
	ReceiveAddress2         string
	SenderCountry           string
	SenderName              string
	SenderPhone             string
	SenderMobile            string
	SenderTax               string
	SenderPassport          string
	SenderMail              string
	SenderCompany           string
	SenderZipcode           string
	SenderProvince          string
	SenderCity              string
	SenderArea              string
	SenderStreet            string
	SenderHouseNumber       string
	SenderAddress           string
	PackageType             string
	HasBattery              string
	BatteryType             string
	HasBack                 string
	Remarks                 string
	DistributionInformation string
	Insurance               string
	SalesPlatform           string
	Weight                  string
	Quantity                string
	DeclaredValue           string
	Status                  string
	CreateTime              string
	WeighingWeight          string
	VolumeWeight            string
	ChargeWeight            string
	StandardFee             string
	Currency                string
	BuyersFee               string
	BuyersWeight            string
	OtherFee                string
	CustomerID              string
	CustomerName            string
	PrepaymentVat           string
	Vat                     string
	ReceiveCertificateType  string
	ReceiveCertificateCode  string
	SenderCertificateType   string
	SenderCertificateCode   string
	ServiceChannelCode      string
	DeliveryTerms           string
	ReceiveCountryName      string
	SenderCountryName       string
	OrderType               string
	ReceiptTime             string
	ReceiptUser             string
	SendTime                string
	SendUser                string
	PicOssLink              string
	RiskMessage             string
	StatusRisk              string
	CreateUser              string
	UpdateTime              string
	UpdateUser              string
	Version                 string
	Deleted                 string
}{
	ID:                      "id",
	OrderNumber:             "order_number",
	ReferenceNumber:         "reference_number",
	LogisticsNumber:         "logistics_number",
	LabelURL:                "label_url",
	OnlineNumber:            "online_number",
	BagNumber:               "bag_number",
	AirBillNumber:           "air_bill_number",
	CustomerChannelID:       "customer_channel_id",
	CustomerChannelName:     "customer_channel_name",
	ReceiveCountry:          "receive_country",
	ReceiveName:             "receive_name",
	ReceivePhone:            "receive_phone",
	ReceiveMobile:           "receive_mobile",
	ReceiveTax:              "receive_tax",
	ReceivePassport:         "receive_passport",
	ReceiveMail:             "receive_mail",
	ReceiveCompany:          "receive_company",
	ReceiveZipcode:          "receive_zipcode",
	ReceiveProvince:         "receive_province",
	ReceiveCity:             "receive_city",
	ReceiveArea:             "receive_area",
	ReceiveStreet:           "receive_street",
	ReceiveHouseNumber:      "receive_house_number",
	ReceiveAddress1:         "receive_address1",
	ReceiveAddress2:         "receive_address2",
	SenderCountry:           "sender_country",
	SenderName:              "sender_name",
	SenderPhone:             "sender_phone",
	SenderMobile:            "sender_mobile",
	SenderTax:               "sender_tax",
	SenderPassport:          "sender_passport",
	SenderMail:              "sender_mail",
	SenderCompany:           "sender_company",
	SenderZipcode:           "sender_zipcode",
	SenderProvince:          "sender_province",
	SenderCity:              "sender_city",
	SenderArea:              "sender_area",
	SenderStreet:            "sender_street",
	SenderHouseNumber:       "sender_house_number",
	SenderAddress:           "sender_address",
	PackageType:             "package_type",
	HasBattery:              "has_battery",
	BatteryType:             "battery_type",
	HasBack:                 "has_back",
	Remarks:                 "remarks",
	DistributionInformation: "distribution_information",
	Insurance:               "insurance",
	SalesPlatform:           "sales_platform",
	Weight:                  "weight",
	Quantity:                "quantity",
	DeclaredValue:           "declared_value",
	Status:                  "status",
	CreateTime:              "create_time",
	WeighingWeight:          "weighing_weight",
	VolumeWeight:            "volume_weight",
	ChargeWeight:            "charge_weight",
	StandardFee:             "standard_fee",
	Currency:                "currency",
	BuyersFee:               "buyers_fee",
	BuyersWeight:            "buyers_weight",
	OtherFee:                "other_fee",
	CustomerID:              "customer_id",
	CustomerName:            "customer_name",
	PrepaymentVat:           "prepayment_vat",
	Vat:                     "vat",
	ReceiveCertificateType:  "receive_certificate_type",
	ReceiveCertificateCode:  "receive_certificate_code",
	SenderCertificateType:   "sender_certificate_type",
	SenderCertificateCode:   "sender_certificate_code",
	ServiceChannelCode:      "service_channel_code",
	DeliveryTerms:           "delivery_terms",
	ReceiveCountryName:      "receive_country_name",
	SenderCountryName:       "sender_country_name",
	OrderType:               "order_type",
	ReceiptTime:             "receipt_time",
	ReceiptUser:             "receipt_user",
	SendTime:                "send_time",
	SendUser:                "send_user",
	PicOssLink:              "pic_oss_link",
	RiskMessage:             "risk_message",
	StatusRisk:              "status_risk",
	CreateUser:              "create_user",
	UpdateTime:              "update_time",
	UpdateUser:              "update_user",
	Version:                 "version",
	Deleted:                 "deleted",
}

// LgRiskControlOrderItem 风控物流订单详情
type LgRiskControlOrderItem struct {
	ID                 int       `gorm:"primaryKey;column:id" json:"-"`                          // ID
	RiskControlOrderID int64     `gorm:"column:risk_control_order_id" json:"riskControlOrderId"` // 订单编号
	NameEn             string    `gorm:"column:name_en" json:"nameEn"`                           // 英文品名
	NameCn             string    `gorm:"column:name_cn" json:"nameCn"`                           // 中文品名
	DeclaredWeight     float64   `gorm:"column:declared_weight" json:"declaredWeight"`           // 申报重量
	DeclaredValue      float64   `gorm:"column:declared_value" json:"declaredValue"`             // 申报价值
	Quantity           int       `gorm:"column:quantity" json:"quantity"`                        // 数量
	Hscode             string    `gorm:"column:hscode" json:"hscode"`                            // HSCODE
	GoodsURL           string    `gorm:"column:goods_url" json:"goodsUrl"`                       // 商品链接
	Sku                string    `gorm:"column:sku" json:"sku"`                                  // SKU
	Length             string    `gorm:"column:length" json:"length"`                            // 长
	Width              string    `gorm:"column:width" json:"width"`                              // 宽
	Height             string    `gorm:"column:height" json:"height"`                            // 高
	Material           int       `gorm:"column:material" json:"material"`                        // 材质
	Purpose            string    `gorm:"column:purpose" json:"purpose"`                          // 用途
	CaseNumber         string    `gorm:"column:case_number" json:"caseNumber"`                   // 箱号
	PlatformProductID  string    `gorm:"column:platform_product_id" json:"platformProductId"`    // 平台产品id
	Specs              string    `gorm:"column:specs" json:"specs"`                              // 规格
	PlaceOrigin        string    `gorm:"column:place_origin" json:"placeOrigin"`                 // 产地
	Currency           string    `gorm:"column:currency" json:"currency"`                        // 币种
	Company            string    `gorm:"column:company" json:"company"`                          // 单位
	PicturesURL        string    `gorm:"column:pictures_url" json:"picturesUrl"`                 // 图片链接
	Brand              string    `gorm:"column:brand" json:"brand"`                              // 品牌
	HasBattery         int       `gorm:"column:has_battery" json:"hasBattery"`                   // 是否带电
	Liquid             string    `gorm:"column:liquid" json:"liquid"`                            // 是否液体,1-液体，0-非液体
	Powder             string    `gorm:"column:powder" json:"powder"`                            // 是否粉末,1-粉末，0-非粉末
	Magnetic           string    `gorm:"column:magnetic" json:"magnetic"`                        // 是否带磁,1-带磁，0-不带磁
	DistributionInfo   string    `gorm:"column:distribution_info" json:"distributionInfo"`       // 配货信息
	BatteryType        string    `gorm:"column:battery_type" json:"batteryType"`                 // 电池类型
	SalesPlatform      string    `gorm:"column:sales_platform" json:"salesPlatform"`             // 销售平台
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                   // 创建时间
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                   // 创建用户
	UpdateTime         time.Time `gorm:"column:update_time" json:"updateTime"`                   // 更新时间
	UpdateUser         int       `gorm:"column:update_user" json:"updateUser"`                   // 更新时间
	Version            int       `gorm:"column:version" json:"version"`                          // 乐观锁
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                          // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgRiskControlOrderItem) TableName() string {
	return "lg_risk_control_order_item"
}

// LgRiskControlOrderItemColumns get sql column name.获取数据库列名
var LgRiskControlOrderItemColumns = struct {
	ID                 string
	RiskControlOrderID string
	NameEn             string
	NameCn             string
	DeclaredWeight     string
	DeclaredValue      string
	Quantity           string
	Hscode             string
	GoodsURL           string
	Sku                string
	Length             string
	Width              string
	Height             string
	Material           string
	Purpose            string
	CaseNumber         string
	PlatformProductID  string
	Specs              string
	PlaceOrigin        string
	Currency           string
	Company            string
	PicturesURL        string
	Brand              string
	HasBattery         string
	Liquid             string
	Powder             string
	Magnetic           string
	DistributionInfo   string
	BatteryType        string
	SalesPlatform      string
	CreateTime         string
	CreateUser         string
	UpdateTime         string
	UpdateUser         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	RiskControlOrderID: "risk_control_order_id",
	NameEn:             "name_en",
	NameCn:             "name_cn",
	DeclaredWeight:     "declared_weight",
	DeclaredValue:      "declared_value",
	Quantity:           "quantity",
	Hscode:             "hscode",
	GoodsURL:           "goods_url",
	Sku:                "sku",
	Length:             "length",
	Width:              "width",
	Height:             "height",
	Material:           "material",
	Purpose:            "purpose",
	CaseNumber:         "case_number",
	PlatformProductID:  "platform_product_id",
	Specs:              "specs",
	PlaceOrigin:        "place_origin",
	Currency:           "currency",
	Company:            "company",
	PicturesURL:        "pictures_url",
	Brand:              "brand",
	HasBattery:         "has_battery",
	Liquid:             "liquid",
	Powder:             "powder",
	Magnetic:           "magnetic",
	DistributionInfo:   "distribution_info",
	BatteryType:        "battery_type",
	SalesPlatform:      "sales_platform",
	CreateTime:         "create_time",
	CreateUser:         "create_user",
	UpdateTime:         "update_time",
	UpdateUser:         "update_user",
	Version:            "version",
	Deleted:            "deleted",
}

// LgServiceProviderErrorMsg 服务商错误信息对照表
type LgServiceProviderErrorMsg struct {
	ID                int       `gorm:"primaryKey;column:id" json:"-"`                       // 主键id
	ServiceSystemCode string    `gorm:"column:service_system_code" json:"serviceSystemCode"` // 系统服务商代码
	ServiceSystemName string    `gorm:"column:service_system_name" json:"serviceSystemName"` // 系统服务商名称
	ServiceErrorCode  string    `gorm:"column:service_error_code" json:"serviceErrorCode"`   // 服务商错误
	TranslateInfo     string    `gorm:"column:translate_info" json:"translateInfo"`          // 翻译信息
	Remark            string    `gorm:"column:remark" json:"remark"`                         // 备注
	CreateUser        int       `gorm:"column:create_user" json:"createUser"`                // 创建人id
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`                // 创建时间
	UpdateUser        int       `gorm:"column:update_user" json:"updateUser"`                // 更新人
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`                // 更新时间
	Version           int       `gorm:"column:version" json:"version"`                       // 乐观锁
	Deleted           int       `gorm:"column:deleted" json:"deleted"`                       // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *LgServiceProviderErrorMsg) TableName() string {
	return "lg_service_provider_error_msg"
}

// LgServiceProviderErrorMsgColumns get sql column name.获取数据库列名
var LgServiceProviderErrorMsgColumns = struct {
	ID                string
	ServiceSystemCode string
	ServiceSystemName string
	ServiceErrorCode  string
	TranslateInfo     string
	Remark            string
	CreateUser        string
	CreateTime        string
	UpdateUser        string
	UpdateTime        string
	Version           string
	Deleted           string
}{
	ID:                "id",
	ServiceSystemCode: "service_system_code",
	ServiceSystemName: "service_system_name",
	ServiceErrorCode:  "service_error_code",
	TranslateInfo:     "translate_info",
	Remark:            "remark",
	CreateUser:        "create_user",
	CreateTime:        "create_time",
	UpdateUser:        "update_user",
	UpdateTime:        "update_time",
	Version:           "version",
	Deleted:           "deleted",
}

// LgTrackGradStatistical 轨迹抓取统计
type LgTrackGradStatistical struct {
	ID           int64     `gorm:"primaryKey;column:id" json:"-"`
	ShipmentCode string    `gorm:"column:shipment_code" json:"shipmentCode"` // 物流渠道代码
	StateCode    string    `gorm:"column:state_code" json:"stateCode"`       // 轨迹状态
	Count        int       `gorm:"column:count" json:"count"`                // 抓取数量
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *LgTrackGradStatistical) TableName() string {
	return "lg_track_grad_statistical"
}

// LgTrackGradStatisticalColumns get sql column name.获取数据库列名
var LgTrackGradStatisticalColumns = struct {
	ID           string
	ShipmentCode string
	StateCode    string
	Count        string
	CreateTime   string
}{
	ID:           "id",
	ShipmentCode: "shipment_code",
	StateCode:    "state_code",
	Count:        "count",
	CreateTime:   "create_time",
}

// LgTrackPushStatistical 轨迹推送统计表
type LgTrackPushStatistical struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Count      int64     `gorm:"column:count" json:"count"`            // 本次推送数量
	Type       string    `gorm:"column:type" json:"type"`              // 业务类型:比如巴西邮政菜鸟推送
	Country    string    `gorm:"column:country" json:"country"`        // 业务国家
	TrackState string    `gorm:"column:track_state" json:"trackState"` // 推送的节点状态码
	PushDate   string    `gorm:"column:push_date" json:"pushDate"`     // 推送日期 yyyy-MM-dd
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 推送时间
}

// TableName get sql table name.获取数据库表名
func (m *LgTrackPushStatistical) TableName() string {
	return "lg_track_push_statistical"
}

// LgTrackPushStatisticalColumns get sql column name.获取数据库列名
var LgTrackPushStatisticalColumns = struct {
	ID         string
	Count      string
	Type       string
	Country    string
	TrackState string
	PushDate   string
	CreateTime string
}{
	ID:         "id",
	Count:      "count",
	Type:       "type",
	Country:    "country",
	TrackState: "track_state",
	PushDate:   "push_date",
	CreateTime: "create_time",
}

// MInternalMessageNotice 站内通知表
type MInternalMessageNotice struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`
	NoticeName      string    `gorm:"column:notice_name" json:"noticeName"`           // 通知名称
	NoticeType      int       `gorm:"column:notice_type" json:"noticeType"`           // 通知类型 0 超过24小时未打大包 1 超过24小时未生成总单
	TemplateContent string    `gorm:"column:template_content" json:"templateContent"` // 通知模板（消息内容)
	IsEnabled       int       `gorm:"column:is_enabled" json:"isEnabled"`             // 是否启用 0 未启用 1 已启用
	Recipient       string    `gorm:"column:recipient" json:"recipient"`              // 消息接收人id 用逗号分割
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`           // 创建用户
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`           // 更新时间
	UpdateUser      int       `gorm:"column:update_user" json:"updateUser"`           // 更新时间
	Version         int       `gorm:"column:version" json:"version"`                  // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                  // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *MInternalMessageNotice) TableName() string {
	return "m_internal_message_notice"
}

// MInternalMessageNoticeColumns get sql column name.获取数据库列名
var MInternalMessageNoticeColumns = struct {
	ID              string
	NoticeName      string
	NoticeType      string
	TemplateContent string
	IsEnabled       string
	Recipient       string
	CreateTime      string
	CreateUser      string
	UpdateTime      string
	UpdateUser      string
	Version         string
	Deleted         string
}{
	ID:              "id",
	NoticeName:      "notice_name",
	NoticeType:      "notice_type",
	TemplateContent: "template_content",
	IsEnabled:       "is_enabled",
	Recipient:       "recipient",
	CreateTime:      "create_time",
	CreateUser:      "create_user",
	UpdateTime:      "update_time",
	UpdateUser:      "update_user",
	Version:         "version",
	Deleted:         "deleted",
}

// MMessage 消息
type MMessage struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`
	Title          string    `gorm:"column:title" json:"title"`                    // 标题
	MsgType        int       `gorm:"column:msg_type" json:"msgType"`               // 消息类型
	MsgTypeDesc    string    `gorm:"column:msg_type_desc" json:"msgTypeDesc"`      // 类型描述
	MsgData        string    `gorm:"column:msg_data" json:"msgData"`               // 消息内容
	IsPop          int       `gorm:"column:is_pop" json:"isPop"`                   // 是否弹窗
	ExpirationTime time.Time `gorm:"column:expiration_time" json:"expirationTime"` // 过期时间
	Remark         string    `gorm:"column:remark" json:"remark"`                  // 备注
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`         // 创建时间
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`         // 创建用户
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`         // 更新时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`         // 更新时间
	Version        int       `gorm:"column:version" json:"version"`                // 乐观锁
	Deleted        int       `gorm:"column:deleted" json:"deleted"`                // 逻辑删除
	BusinessNumber string    `gorm:"column:business_number" json:"businessNumber"` // 业务单号
}

// TableName get sql table name.获取数据库表名
func (m *MMessage) TableName() string {
	return "m_message"
}

// MMessageColumns get sql column name.获取数据库列名
var MMessageColumns = struct {
	ID             string
	Title          string
	MsgType        string
	MsgTypeDesc    string
	MsgData        string
	IsPop          string
	ExpirationTime string
	Remark         string
	CreateTime     string
	CreateUser     string
	UpdateTime     string
	UpdateUser     string
	Version        string
	Deleted        string
	BusinessNumber string
}{
	ID:             "id",
	Title:          "title",
	MsgType:        "msg_type",
	MsgTypeDesc:    "msg_type_desc",
	MsgData:        "msg_data",
	IsPop:          "is_pop",
	ExpirationTime: "expiration_time",
	Remark:         "remark",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	Version:        "version",
	Deleted:        "deleted",
	BusinessNumber: "business_number",
}

// MMessageTemplate 消息模板
type MMessageTemplate struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Name       string    `gorm:"column:name" json:"name"`              // 模板名称
	ModuleName string    `gorm:"column:module_name" json:"moduleName"` // 所属模块
	Template   string    `gorm:"column:template" json:"template"`      // 模板
	IsPrivate  int       `gorm:"column:is_private" json:"isPrivate"`   // 是否私有
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *MMessageTemplate) TableName() string {
	return "m_message_template"
}

// MMessageTemplateColumns get sql column name.获取数据库列名
var MMessageTemplateColumns = struct {
	ID         string
	Name       string
	ModuleName string
	Template   string
	IsPrivate  string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
}{
	ID:         "id",
	Name:       "name",
	ModuleName: "module_name",
	Template:   "template",
	IsPrivate:  "is_private",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
}

// MMessageUser 消息用户
type MMessageUser struct {
	MessageID          int       `gorm:"column:message_id" json:"messageId"`
	Title              string    `gorm:"column:title" json:"title"`                             // 消息标题
	MsgData            string    `gorm:"column:msg_data" json:"msgData"`                        // 消息内容
	MsgType            int       `gorm:"column:msg_type" json:"msgType"`                        // 消息类型
	MsgTypeDesc        string    `gorm:"column:msg_type_desc" json:"msgTypeDesc"`               // 消息类型中文
	IsPop              int       `gorm:"column:is_pop" json:"isPop"`                            // 是否弹窗
	ExpirationTime     time.Time `gorm:"column:expiration_time" json:"expirationTime"`          // 过期时间
	ReceiveID          int       `gorm:"column:receive_id" json:"receiveId"`                    // 接收人
	ReceiveType        int       `gorm:"column:receive_type" json:"receiveType"`                // 接收人类型
	ReceiveTypeChinese string    `gorm:"column:receive_type_chinese" json:"receiveTypeChinese"` // 接收人中文
	IsRead             int       `gorm:"column:is_read" json:"isRead"`                          // 是否已读
	ReadTime           time.Time `gorm:"column:read_time" json:"readTime"`                      // 已读时间
	CreateTime         time.Time `gorm:"column:create_time" json:"createTime"`                  // 创建时间
	CreateUser         int       `gorm:"column:create_user" json:"createUser"`                  // 创建用户
	Version            int       `gorm:"column:version" json:"version"`                         // 乐观锁
	Deleted            int       `gorm:"column:deleted" json:"deleted"`                         // 逻辑删除
	BusinessNumber     string    `gorm:"column:business_number" json:"businessNumber"`          // 业务单号
}

// TableName get sql table name.获取数据库表名
func (m *MMessageUser) TableName() string {
	return "m_message_user"
}

// MMessageUserColumns get sql column name.获取数据库列名
var MMessageUserColumns = struct {
	MessageID          string
	Title              string
	MsgData            string
	MsgType            string
	MsgTypeDesc        string
	IsPop              string
	ExpirationTime     string
	ReceiveID          string
	ReceiveType        string
	ReceiveTypeChinese string
	IsRead             string
	ReadTime           string
	CreateTime         string
	CreateUser         string
	Version            string
	Deleted            string
	BusinessNumber     string
}{
	MessageID:          "message_id",
	Title:              "title",
	MsgData:            "msg_data",
	MsgType:            "msg_type",
	MsgTypeDesc:        "msg_type_desc",
	IsPop:              "is_pop",
	ExpirationTime:     "expiration_time",
	ReceiveID:          "receive_id",
	ReceiveType:        "receive_type",
	ReceiveTypeChinese: "receive_type_chinese",
	IsRead:             "is_read",
	ReadTime:           "read_time",
	CreateTime:         "create_time",
	CreateUser:         "create_user",
	Version:            "version",
	Deleted:            "deleted",
	BusinessNumber:     "business_number",
}

// MWebNotice 公告表
type MWebNotice struct {
	ID             int64     `gorm:"primaryKey;column:id" json:"-"`                 // 公告表ID
	Title          string    `gorm:"column:title" json:"title"`                     // 标题
	Content        string    `gorm:"column:content" json:"content"`                 // 内容
	WatchVolume    int       `gorm:"column:watch_volume" json:"watchVolume"`        // 查看量
	FileID         int64     `gorm:"column:file_id" json:"fileId"`                  // 文件表id
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`          // 创建时间
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`          // 创建用户
	CreateUserName string    `gorm:"column:create_user_name" json:"createUserName"` // 创建用户名称
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`          // 更新时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`          // 更新时间
	UpdateUserName string    `gorm:"column:update_user_name" json:"updateUserName"` // 更新用户名称
	Version        int       `gorm:"column:version" json:"version"`                 // 乐观锁
	Deleted        int       `gorm:"column:deleted" json:"deleted"`                 // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *MWebNotice) TableName() string {
	return "m_web_notice"
}

// MWebNoticeColumns get sql column name.获取数据库列名
var MWebNoticeColumns = struct {
	ID             string
	Title          string
	Content        string
	WatchVolume    string
	FileID         string
	CreateTime     string
	CreateUser     string
	CreateUserName string
	UpdateTime     string
	UpdateUser     string
	UpdateUserName string
	Version        string
	Deleted        string
}{
	ID:             "id",
	Title:          "title",
	Content:        "content",
	WatchVolume:    "watch_volume",
	FileID:         "file_id",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	CreateUserName: "create_user_name",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	UpdateUserName: "update_user_name",
	Version:        "version",
	Deleted:        "deleted",
}

// PdGoodsRiskControl [...]
type PdGoodsRiskControl struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`                 // 主键
	GoodsName      string    `gorm:"column:goods_name" json:"goodsName"`            // 商品英文名称
	GoodsNameCn    string    `gorm:"column:goods_name_cn" json:"goodsNameCn"`       // 商品中文名称
	ShortGoodsName string    `gorm:"column:short_goods_name" json:"shortGoodsName"` // 商品英文名称(简称)
	Price          float64   `gorm:"column:price" json:"price"`                     // 商品价值/单价
	MinPrice       float64   `gorm:"column:min_price" json:"minPrice"`              // 最低价值/单价
	LimitOut       int       `gorm:"column:limit_out" json:"limitOut"`              // 限制出库 默认否(不限制)
	GoodsURL       string    `gorm:"column:goods_url" json:"goodsUrl"`              // 产品url
	HsCode         string    `gorm:"column:hs_code" json:"hsCode"`                  // 海关编码
	Remark         string    `gorm:"column:remark" json:"remark"`                   // 备注
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`          // 创建时间
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`          // 创建用户
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`          // 更新时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`          // 更新时间
	Version        int       `gorm:"column:version" json:"version"`                 // 乐观锁
	Deleted        int       `gorm:"column:deleted" json:"deleted"`                 // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *PdGoodsRiskControl) TableName() string {
	return "pd_goods_risk_control"
}

// PdGoodsRiskControlColumns get sql column name.获取数据库列名
var PdGoodsRiskControlColumns = struct {
	ID             string
	GoodsName      string
	GoodsNameCn    string
	ShortGoodsName string
	Price          string
	MinPrice       string
	LimitOut       string
	GoodsURL       string
	HsCode         string
	Remark         string
	CreateTime     string
	CreateUser     string
	UpdateTime     string
	UpdateUser     string
	Version        string
	Deleted        string
}{
	ID:             "id",
	GoodsName:      "goods_name",
	GoodsNameCn:    "goods_name_cn",
	ShortGoodsName: "short_goods_name",
	Price:          "price",
	MinPrice:       "min_price",
	LimitOut:       "limit_out",
	GoodsURL:       "goods_url",
	HsCode:         "hs_code",
	Remark:         "remark",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	Version:        "version",
	Deleted:        "deleted",
}

// SysCheckTheInformation [...]
type SysCheckTheInformation struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`          // ID
	FilterKey   string    `gorm:"column:filter_key" json:"filterKey"`     // 校验代码
	FilterValue string    `gorm:"column:filter_value" json:"filterValue"` // 校验过滤值
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`   // 创建时间
	CreateUser  int       `gorm:"column:create_user" json:"createUser"`   // 创建用户
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`   // 更新时间
	UpdateUser  int       `gorm:"column:update_user" json:"updateUser"`   // 更新用户
	Version     int       `gorm:"column:version" json:"version"`          // 乐观锁
	Remarks     string    `gorm:"column:remarks" json:"remarks"`          // 备注
}

// TableName get sql table name.获取数据库表名
func (m *SysCheckTheInformation) TableName() string {
	return "sys_check_the_information"
}

// SysCheckTheInformationColumns get sql column name.获取数据库列名
var SysCheckTheInformationColumns = struct {
	ID          string
	FilterKey   string
	FilterValue string
	CreateTime  string
	CreateUser  string
	UpdateTime  string
	UpdateUser  string
	Version     string
	Remarks     string
}{
	ID:          "id",
	FilterKey:   "filter_key",
	FilterValue: "filter_value",
	CreateTime:  "create_time",
	CreateUser:  "create_user",
	UpdateTime:  "update_time",
	UpdateUser:  "update_user",
	Version:     "version",
	Remarks:     "remarks",
}

// SysConfig [...]
type SysConfig struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`            // ID
	ConfigKey    string    `gorm:"column:config_key" json:"configKey"`       // 配置key
	ConfigValue  string    `gorm:"column:config_value" json:"configValue"`   // 配置值
	DefaultValue string    `gorm:"column:default_value" json:"defaultValue"` // 默认值
	Remarks      string    `gorm:"column:remarks" json:"remarks"`            // 备注说明
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建用户
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新时间
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *SysConfig) TableName() string {
	return "sys_config"
}

// SysConfigColumns get sql column name.获取数据库列名
var SysConfigColumns = struct {
	ID           string
	ConfigKey    string
	ConfigValue  string
	DefaultValue string
	Remarks      string
	CreateTime   string
	CreateUser   string
	UpdateTime   string
	UpdateUser   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	ConfigKey:    "config_key",
	ConfigValue:  "config_value",
	DefaultValue: "default_value",
	Remarks:      "remarks",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateTime:   "update_time",
	UpdateUser:   "update_user",
	Version:      "version",
	Deleted:      "deleted",
}

// SysEmailTemplate 邮件模板
type SysEmailTemplate struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`                  // 主键id
	TemplateID      int       `gorm:"column:template_id" json:"templateId"`           // 模板id
	TemplateName    string    `gorm:"column:template_name" json:"templateName"`       // 模板名称
	SendNode        string    `gorm:"column:send_node" json:"sendNode"`               // 发送节点
	TemplateContent string    `gorm:"column:template_content" json:"templateContent"` // 模板内容
	Status          []uint8   `gorm:"column:status" json:"status"`                    // 是否启用(0-否，1-启用)
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`           // 创建人
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`           // 更新时间
	UpdateUser      int       `gorm:"column:update_user" json:"updateUser"`           // 更新人
	Version         int       `gorm:"column:version" json:"version"`                  // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                  // 逻辑删除
	ReviewStatus    int       `gorm:"column:review_status" json:"reviewStatus"`       // 审核状态 1-审核中|0-已通过|2-拒绝|其它-不可用
	ReviewReason    string    `gorm:"column:review_reason" json:"reviewReason"`       // 审核原因 审核的失败原因
}

// TableName get sql table name.获取数据库表名
func (m *SysEmailTemplate) TableName() string {
	return "sys_email_template"
}

// SysEmailTemplateColumns get sql column name.获取数据库列名
var SysEmailTemplateColumns = struct {
	ID              string
	TemplateID      string
	TemplateName    string
	SendNode        string
	TemplateContent string
	Status          string
	CreateTime      string
	CreateUser      string
	UpdateTime      string
	UpdateUser      string
	Version         string
	Deleted         string
	ReviewStatus    string
	ReviewReason    string
}{
	ID:              "id",
	TemplateID:      "template_id",
	TemplateName:    "template_name",
	SendNode:        "send_node",
	TemplateContent: "template_content",
	Status:          "status",
	CreateTime:      "create_time",
	CreateUser:      "create_user",
	UpdateTime:      "update_time",
	UpdateUser:      "update_user",
	Version:         "version",
	Deleted:         "deleted",
	ReviewStatus:    "review_status",
	ReviewReason:    "review_reason",
}

// SysMessageSendRecord 消息发送记录
type SysMessageSendRecord struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`
	OrderNumber          string    `gorm:"column:order_number" json:"orderNumber"`                    // 订单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	SendContent          string    `gorm:"column:send_content" json:"sendContent"`                    // 发送内容
	MessageType          string    `gorm:"column:message_type" json:"messageType"`                    // 消息类型，sms:短信，email:邮件
	SendNode             string    `gorm:"column:send_node" json:"sendNode"`                          // 发送节点
	SendStatus           int8      `gorm:"column:send_status" json:"sendStatus"`                      // 发送状态，0:发送中，1:发送成功，2:发送失败
	SendCount            int       `gorm:"column:send_count" json:"sendCount"`                        // 发送次数
	Number               string    `gorm:"column:number" json:"number"`                               // 发送号码
	SendTime             time.Time `gorm:"column:send_time" json:"sendTime"`                          // 发送时间
	ServiceProviderCode  string    `gorm:"column:service_provider_code" json:"serviceProviderCode"`   // 发送消息服务商code
	OperatorsID          string    `gorm:"column:operators_id" json:"operatorsId"`                    // 运营商ID
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 备注
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 创建用户
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新用户
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *SysMessageSendRecord) TableName() string {
	return "sys_message_send_record"
}

// SysMessageSendRecordColumns get sql column name.获取数据库列名
var SysMessageSendRecordColumns = struct {
	ID                   string
	OrderNumber          string
	ReferenceNumber      string
	LogisticsNumber      string
	LogisticsNumberFinal string
	SendContent          string
	MessageType          string
	SendNode             string
	SendStatus           string
	SendCount            string
	Number               string
	SendTime             string
	ServiceProviderCode  string
	OperatorsID          string
	Remark               string
	CreateTime           string
	CreateUser           string
	UpdateTime           string
	UpdateUser           string
	Version              string
}{
	ID:                   "id",
	OrderNumber:          "order_number",
	ReferenceNumber:      "reference_number",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	SendContent:          "send_content",
	MessageType:          "message_type",
	SendNode:             "send_node",
	SendStatus:           "send_status",
	SendCount:            "send_count",
	Number:               "number",
	SendTime:             "send_time",
	ServiceProviderCode:  "service_provider_code",
	OperatorsID:          "operators_id",
	Remark:               "remark",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateTime:           "update_time",
	UpdateUser:           "update_user",
	Version:              "version",
}

// SysRabbitMqErrorRecord RabbitMQ错误消息记录
type SysRabbitMqErrorRecord struct {
	ID           int64     `gorm:"primaryKey;column:id" json:"-"`            // 主键
	ExchangeName string    `gorm:"column:exchange_name" json:"exchangeName"` // 交换机名称
	ExchangeType string    `gorm:"column:exchange_type" json:"exchangeType"` // 交换机类型，direct:精确路由、fanout:扇出,广播、topic:模糊路由
	QueueName    string    `gorm:"column:queue_name" json:"queueName"`       // 队列名称
	ModelName    string    `gorm:"column:model_name" json:"modelName"`       // 模块名称
	MsgContent   string    `gorm:"column:msg_content" json:"msgContent"`     // 消息内容
	ErrorReason  string    `gorm:"column:error_reason" json:"errorReason"`   // 错误原因
	ErrorType    bool      `gorm:"column:error_type" json:"errorType"`       // 错误类型，1:生产失败、2:消费失败
	IsResolve    []uint8   `gorm:"column:is_resolve" json:"isResolve"`       // 已处理，0:否，1:是
	Remark       string    `gorm:"column:remark" json:"remark"`              // 备注
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 生成用户,生成应收人员
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新人
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *SysRabbitMqErrorRecord) TableName() string {
	return "sys_rabbit_mq_error_record"
}

// SysRabbitMqErrorRecordColumns get sql column name.获取数据库列名
var SysRabbitMqErrorRecordColumns = struct {
	ID           string
	ExchangeName string
	ExchangeType string
	QueueName    string
	ModelName    string
	MsgContent   string
	ErrorReason  string
	ErrorType    string
	IsResolve    string
	Remark       string
	CreateTime   string
	CreateUser   string
	UpdateUser   string
	UpdateTime   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	ExchangeName: "exchange_name",
	ExchangeType: "exchange_type",
	QueueName:    "queue_name",
	ModelName:    "model_name",
	MsgContent:   "msg_content",
	ErrorReason:  "error_reason",
	ErrorType:    "error_type",
	IsResolve:    "is_resolve",
	Remark:       "remark",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateUser:   "update_user",
	UpdateTime:   "update_time",
	Version:      "version",
	Deleted:      "deleted",
}

// SysServerHealth 服务健康检测记录
type SysServerHealth struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"-"`
	Name       string    `gorm:"column:name" json:"name"`              // 服务名称
	URL        string    `gorm:"column:url" json:"url"`                // 服务健康检查请求地址
	State      []uint8   `gorm:"column:state" json:"state"`            // 正常:1异常:0
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *SysServerHealth) TableName() string {
	return "sys_server_health"
}

// SysServerHealthColumns get sql column name.获取数据库列名
var SysServerHealthColumns = struct {
	ID         string
	Name       string
	URL        string
	State      string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	Name:       "name",
	URL:        "url",
	State:      "state",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// SysTemplateField 场景模板字段变量
type SysTemplateField struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`
	ModuleName   string    `gorm:"column:module_name" json:"moduleName"`     // 模块
	FieldName    string    `gorm:"column:field_name" json:"fieldName"`       // 字段名
	FieldChinese string    `gorm:"column:field_chinese" json:"fieldChinese"` // 中文
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *SysTemplateField) TableName() string {
	return "sys_template_field"
}

// SysTemplateFieldColumns get sql column name.获取数据库列名
var SysTemplateFieldColumns = struct {
	ID           string
	ModuleName   string
	FieldName    string
	FieldChinese string
	CreateTime   string
}{
	ID:           "id",
	ModuleName:   "module_name",
	FieldName:    "field_name",
	FieldChinese: "field_chinese",
	CreateTime:   "create_time",
}

// TUser [...]
type TUser struct {
	ID       int       `gorm:"primaryKey;column:id" json:"-"`    // 主键
	Username string    `gorm:"column:username" json:"username"`  // 账户名
	Password string    `gorm:"column:password" json:"password"`  // 密码(必须加密)
	RealName string    `gorm:"column:real_name" json:"realName"` // 真实姓名
	Phone    string    `gorm:"column:phone" json:"phone"`        // 手机
	Created  time.Time `gorm:"column:created" json:"created"`    // 创建时间
	Updated  time.Time `gorm:"column:updated" json:"updated"`    // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *TUser) TableName() string {
	return "t_user"
}

// TUserColumns get sql column name.获取数据库列名
var TUserColumns = struct {
	ID       string
	Username string
	Password string
	RealName string
	Phone    string
	Created  string
	Updated  string
}{
	ID:       "id",
	Username: "username",
	Password: "password",
	RealName: "real_name",
	Phone:    "phone",
	Created:  "created",
	Updated:  "updated",
}

// UContract 合同
type UContract struct {
	ID               int            `gorm:"primaryKey;column:id" json:"-"`
	Code             string         `gorm:"column:code" json:"code"`                            // 合同编号
	Title            string         `gorm:"column:title" json:"title"`                          // 合同标题
	Content          string         `gorm:"column:content" json:"content"`                      // 合同内容
	Type             string         `gorm:"column:type" json:"type"`                            // 合同类型
	SignDate         datatypes.Date `gorm:"column:sign_date" json:"signDate"`                   // 签订日期
	ExpireDate       datatypes.Date `gorm:"column:expire_date" json:"expireDate"`               // 到期日期
	DrafterUserID    int            `gorm:"column:drafter_user_id" json:"drafterUserId"`        // 拟稿人用户id
	PartyA           string         `gorm:"column:party_a" json:"partyA"`                       // 甲方
	PartyB           string         `gorm:"column:party_b" json:"partyB"`                       // 乙方
	SignStatus       string         `gorm:"column:sign_status" json:"signStatus"`               // 合同状态(履行中、合同到期、合同作废)
	Target           string         `gorm:"column:target" json:"target"`                        // 合同用途名称对应u_contract_use_type（合同用途类型表） 表的type_name字段
	Attachment       string         `gorm:"column:attachment" json:"attachment"`                // 附件
	ESignature       string         `gorm:"column:e_signature" json:"eSignature"`               // 电子签名资料
	CreateTime       time.Time      `gorm:"column:create_time" json:"createTime"`               // 创建时间
	CreateUser       int            `gorm:"column:create_user" json:"createUser"`               // 创建用户
	UpdateTime       time.Time      `gorm:"column:update_time" json:"updateTime"`               // 更新时间
	UpdateUser       int            `gorm:"column:update_user" json:"updateUser"`               // 更新时间
	Version          int            `gorm:"column:version" json:"version"`                      // 乐观锁
	Deleted          int            `gorm:"column:deleted" json:"deleted"`                      // 逻辑删除
	ContractTypeID   int            `gorm:"column:contract_type_id" json:"contractTypeId"`      // 合同用途类型表(u_contract_use_type)的id
	IsDeposit        []uint8        `gorm:"column:is_deposit" json:"isDeposit"`                 // 是否需要押金 0 否 1 是
	DepositAmount    float64        `gorm:"column:deposit_amount" json:"depositAmount"`         // 押金金额
	CustomEncoding   string         `gorm:"column:custom_encoding" json:"customEncoding"`       // 自定义编码
	PartyACustomerID int            `gorm:"column:party_a_customer_id" json:"partyACustomerId"` // 客户id
	ExtendedData     datatypes.JSON `gorm:"column:extended_data" json:"extendedData"`           // 扩展数据
	CustomerName     string         `gorm:"column:customer_name" json:"customerName"`           // 客户名称
}

// TableName get sql table name.获取数据库表名
func (m *UContract) TableName() string {
	return "u_contract"
}

// UContractColumns get sql column name.获取数据库列名
var UContractColumns = struct {
	ID               string
	Code             string
	Title            string
	Content          string
	Type             string
	SignDate         string
	ExpireDate       string
	DrafterUserID    string
	PartyA           string
	PartyB           string
	SignStatus       string
	Target           string
	Attachment       string
	ESignature       string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
	ContractTypeID   string
	IsDeposit        string
	DepositAmount    string
	CustomEncoding   string
	PartyACustomerID string
	ExtendedData     string
	CustomerName     string
}{
	ID:               "id",
	Code:             "code",
	Title:            "title",
	Content:          "content",
	Type:             "type",
	SignDate:         "sign_date",
	ExpireDate:       "expire_date",
	DrafterUserID:    "drafter_user_id",
	PartyA:           "party_a",
	PartyB:           "party_b",
	SignStatus:       "sign_status",
	Target:           "target",
	Attachment:       "attachment",
	ESignature:       "e_signature",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
	ContractTypeID:   "contract_type_id",
	IsDeposit:        "is_deposit",
	DepositAmount:    "deposit_amount",
	CustomEncoding:   "custom_encoding",
	PartyACustomerID: "party_a_customer_id",
	ExtendedData:     "extended_data",
	CustomerName:     "customer_name",
}

// UContractUseType 合同用途类型表（动态维护）
type UContractUseType struct {
	ID         int64     `gorm:"primaryKey;column:id" json:"-"`
	TypeName   string    `gorm:"column:type_name" json:"typeName"`     // 类型名称
	Code       string    `gorm:"column:code" json:"code"`              // 类型code
	UserIDs    string    `gorm:"column:user_ids" json:"userIds"`       // 部分可见的用户id（用","逗号隔开）
	IsDelete   int       `gorm:"column:is_delete" json:"isDelete"`     // 逻辑删除 0 未删除 1 已删除
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 修改时间
	CreateUser string    `gorm:"column:create_user" json:"createUser"` // 创建人
	UpdateUser string    `gorm:"column:update_user" json:"updateUser"` // 修改人
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Remark     string    `gorm:"column:remark" json:"remark"`          // 合同用途备注
}

// TableName get sql table name.获取数据库表名
func (m *UContractUseType) TableName() string {
	return "u_contract_use_type"
}

// UContractUseTypeColumns get sql column name.获取数据库列名
var UContractUseTypeColumns = struct {
	ID         string
	TypeName   string
	Code       string
	UserIDs    string
	IsDelete   string
	CreateTime string
	UpdateTime string
	CreateUser string
	UpdateUser string
	Version    string
	Remark     string
}{
	ID:         "id",
	TypeName:   "type_name",
	Code:       "code",
	UserIDs:    "user_ids",
	IsDelete:   "is_delete",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	CreateUser: "create_user",
	UpdateUser: "update_user",
	Version:    "version",
	Remark:     "remark",
}

// UFileExport oa文件导出
type UFileExport struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`                  // 主键
	Operator        string    `gorm:"column:operator" json:"operator"`                // 操作人
	OperationModule string    `gorm:"column:operation_module" json:"operationModule"` // 操作模块
	FileType        string    `gorm:"column:file_type" json:"fileType"`               // 文件类型
	FileURL         string    `gorm:"column:file_url" json:"fileUrl"`                 // 访问路径
	OssPath         string    `gorm:"column:oss_path" json:"ossPath"`                 // 上传oos路径
	VerifyName      string    `gorm:"column:verify_name" json:"verifyName"`           // 审核人
	VerifyID        int       `gorm:"column:verify_id" json:"verifyId"`               // 审核人id
	IsVerify        int       `gorm:"column:is_verify" json:"isVerify"`               // 是否需要审核：0-否，1-是
	VerifyStatus    int       `gorm:"column:verify_status" json:"verifyStatus"`       // 审核状态：0-未审核，1-已审核
	ExportResult    string    `gorm:"column:export_result" json:"exportResult"`       // 导出结果
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`           // 操作人
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`           // 更新时间
	Version         int       `gorm:"column:version" json:"version"`                  // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                  // 逻辑删除
	DownloadsNumber int       `gorm:"column:downloads_number" json:"downloadsNumber"` // 下载次数
	FileName        string    `gorm:"column:file_name" json:"fileName"`               // 文件名
	GenerateTime    time.Time `gorm:"column:generate_time" json:"generateTime"`       // 生成时间
}

// TableName get sql table name.获取数据库表名
func (m *UFileExport) TableName() string {
	return "u_file_export"
}

// UFileExportColumns get sql column name.获取数据库列名
var UFileExportColumns = struct {
	ID              string
	Operator        string
	OperationModule string
	FileType        string
	FileURL         string
	OssPath         string
	VerifyName      string
	VerifyID        string
	IsVerify        string
	VerifyStatus    string
	ExportResult    string
	CreateTime      string
	CreateUser      string
	UpdateTime      string
	Version         string
	Deleted         string
	DownloadsNumber string
	FileName        string
	GenerateTime    string
}{
	ID:              "id",
	Operator:        "operator",
	OperationModule: "operation_module",
	FileType:        "file_type",
	FileURL:         "file_url",
	OssPath:         "oss_path",
	VerifyName:      "verify_name",
	VerifyID:        "verify_id",
	IsVerify:        "is_verify",
	VerifyStatus:    "verify_status",
	ExportResult:    "export_result",
	CreateTime:      "create_time",
	CreateUser:      "create_user",
	UpdateTime:      "update_time",
	Version:         "version",
	Deleted:         "deleted",
	DownloadsNumber: "downloads_number",
	FileName:        "file_name",
	GenerateTime:    "generate_time",
}

// ULoginLog [...]
type ULoginLog struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`          // ID
	Username    string    `gorm:"column:username" json:"username"`        // 用户名
	IP          string    `gorm:"column:ip" json:"ip"`                    // 登录IP
	IPAddress   string    `gorm:"column:ip_address" json:"ipAddress"`     // ip地址所在地区
	Source      string    `gorm:"column:source" json:"source"`            // 来源
	BrowserInfo string    `gorm:"column:browser_info" json:"browserInfo"` // 浏览器信息
	OsInfo      string    `gorm:"column:os_info" json:"osInfo"`           // 操作系统信息
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`   // 创建时间
	CreateUser  int       `gorm:"column:create_user" json:"createUser"`   // 创建用户
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`   // 更新时间
	UpdateUser  int       `gorm:"column:update_user" json:"updateUser"`   // 更新时间
	Version     int       `gorm:"column:version" json:"version"`          // 乐观锁
	Deleted     int       `gorm:"column:deleted" json:"deleted"`          // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *ULoginLog) TableName() string {
	return "u_login_log"
}

// ULoginLogColumns get sql column name.获取数据库列名
var ULoginLogColumns = struct {
	ID          string
	Username    string
	IP          string
	IPAddress   string
	Source      string
	BrowserInfo string
	OsInfo      string
	CreateTime  string
	CreateUser  string
	UpdateTime  string
	UpdateUser  string
	Version     string
	Deleted     string
}{
	ID:          "id",
	Username:    "username",
	IP:          "ip",
	IPAddress:   "ip_address",
	Source:      "source",
	BrowserInfo: "browser_info",
	OsInfo:      "os_info",
	CreateTime:  "create_time",
	CreateUser:  "create_user",
	UpdateTime:  "update_time",
	UpdateUser:  "update_user",
	Version:     "version",
	Deleted:     "deleted",
}

// UMemorandum 用户随手记
type UMemorandum struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Content    string    `gorm:"column:content" json:"content"` // 随手记内容
	IsComplete []uint8   `gorm:"column:is_complete" json:"isComplete"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *UMemorandum) TableName() string {
	return "u_memorandum"
}

// UMemorandumColumns get sql column name.获取数据库列名
var UMemorandumColumns = struct {
	ID         string
	Content    string
	IsComplete string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	Content:    "content",
	IsComplete: "is_complete",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// UOrganization 组织架构
type UOrganization struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // 主键
	ParentID   int       `gorm:"column:parent_id" json:"parentId"`     // 父ID
	Name       string    `gorm:"column:name" json:"name"`              // 名称
	OrgCode    string    `gorm:"column:org_code" json:"orgCode"`       // 组织结构代码
	Sort       int       `gorm:"column:sort" json:"sort"`              // 排序 值越小越靠前
	IsParent   int8      `gorm:"column:is_parent" json:"isParent"`     // 该类目是否为父类目，1为true，0为false
	Level      int       `gorm:"column:level" json:"level"`            // 级别
	OrgType    int       `gorm:"column:org_type" json:"orgType"`       // 组织类型
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *UOrganization) TableName() string {
	return "u_organization"
}

// UOrganizationColumns get sql column name.获取数据库列名
var UOrganizationColumns = struct {
	ID         string
	ParentID   string
	Name       string
	OrgCode    string
	Sort       string
	IsParent   string
	Level      string
	OrgType    string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	ParentID:   "parent_id",
	Name:       "name",
	OrgCode:    "org_code",
	Sort:       "sort",
	IsParent:   "is_parent",
	Level:      "level",
	OrgType:    "org_type",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// UPermission 权限表
type UPermission struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`                // 主键
	ParentID       int       `gorm:"column:parent_id" json:"parentId"`             // 父节点ID
	PermissionName string    `gorm:"column:permission_name" json:"permissionName"` // 权限名称
	PermissionCode string    `gorm:"column:permission_code" json:"permissionCode"` // 权限标识符
	Sort           int       `gorm:"column:sort" json:"sort"`                      // 排序 值越小越靠前
	IsParent       int8      `gorm:"column:is_parent" json:"isParent"`             // 该类目是否为父类目，1为true，0为false
	Remark         string    `gorm:"column:remark" json:"remark"`                  // 备注
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`         // 创建时间
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`         // 创建用户
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`         // 更新时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`         // 更新时间
	Version        int       `gorm:"column:version" json:"version"`                // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *UPermission) TableName() string {
	return "u_permission"
}

// UPermissionColumns get sql column name.获取数据库列名
var UPermissionColumns = struct {
	ID             string
	ParentID       string
	PermissionName string
	PermissionCode string
	Sort           string
	IsParent       string
	Remark         string
	CreateTime     string
	CreateUser     string
	UpdateTime     string
	UpdateUser     string
	Version        string
}{
	ID:             "id",
	ParentID:       "parent_id",
	PermissionName: "permission_name",
	PermissionCode: "permission_code",
	Sort:           "sort",
	IsParent:       "is_parent",
	Remark:         "remark",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	Version:        "version",
}

// UPosition 职务维护
type UPosition struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Name       string    `gorm:"column:name" json:"name"`              // 职务名称
	OrgID      int       `gorm:"column:org_id" json:"orgId"`           // 职务绑定的部门ID
	OrgIDs     string    `gorm:"column:org_ids" json:"orgIds"`         // 职务关联组织架构ids
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *UPosition) TableName() string {
	return "u_position"
}

// UPositionColumns get sql column name.获取数据库列名
var UPositionColumns = struct {
	ID         string
	Name       string
	OrgID      string
	OrgIDs     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
}{
	ID:         "id",
	Name:       "name",
	OrgID:      "org_id",
	OrgIDs:     "org_ids",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
}

// UResource 资源管理表
type UResource struct {
	ID           int          `gorm:"primaryKey;column:id" json:"-"`                                             // ID
	Code         string       `gorm:"column:code" json:"code"`                                                   // 资源编号
	UserID       int          `gorm:"column:user_id" json:"userId"`                                              // 用户iD
	Type         string       `gorm:"column:type" json:"type"`                                                   // 文件标签类型(文件,图片等)
	Tag          string       `gorm:"column:tag" json:"tag"`                                                     // 自定义标签(例如IT部文件)
	UResourceTag UResourceTag `gorm:"joinForeignKey:tag;foreignKey:name;references:Tag" json:"uResourceTagList"` // 资源标签维护
	Name         string       `gorm:"column:name" json:"name"`                                                   // 文件名
	Resources    string       `gorm:"column:resources" json:"resources"`                                         // 资源列表(下载链接)
	IsOpen       int          `gorm:"column:is_open" json:"isOpen"`                                              // 是否公开 0:不公开(私人) 1:公开(公共)
	CreateTime   time.Time    `gorm:"column:create_time" json:"createTime"`                                      // 创建时间
	CreateUser   int          `gorm:"column:create_user" json:"createUser"`                                      // 创建用户
	UpdateTime   time.Time    `gorm:"column:update_time" json:"updateTime"`                                      // 更新时间
	UpdateUser   int          `gorm:"column:update_user" json:"updateUser"`                                      // 更新用户
	Version      int          `gorm:"column:version" json:"version"`                                             // 乐观锁
	Deleted      int          `gorm:"column:deleted" json:"deleted"`                                             // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *UResource) TableName() string {
	return "u_resource"
}

// UResourceColumns get sql column name.获取数据库列名
var UResourceColumns = struct {
	ID         string
	Code       string
	UserID     string
	Type       string
	Tag        string
	Name       string
	Resources  string
	IsOpen     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	Code:       "code",
	UserID:     "user_id",
	Type:       "type",
	Tag:        "tag",
	Name:       "name",
	Resources:  "resources",
	IsOpen:     "is_open",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// UResourceTag 资源标签维护
type UResourceTag struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Name       string    `gorm:"column:name" json:"name"`              // 资源标签名称
	Type       int       `gorm:"column:type" json:"type"`              // 类型:保留字段
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *UResourceTag) TableName() string {
	return "u_resource_tag"
}

// UResourceTagColumns get sql column name.获取数据库列名
var UResourceTagColumns = struct {
	ID         string
	Name       string
	Type       string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
}{
	ID:         "id",
	Name:       "name",
	Type:       "type",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
}

// URole 角色表
type URole struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`        // 主键
	RoleName   string    `gorm:"column:role_name" json:"roleName"`     // 角色名称
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新用户
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *URole) TableName() string {
	return "u_role"
}

// URoleColumns get sql column name.获取数据库列名
var URoleColumns = struct {
	ID         string
	RoleName   string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	RoleName:   "role_name",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// URolePermission 角色_权限
type URolePermission struct {
	ID           int64       `gorm:"column:id" json:"id"`
	RoleID       int         `gorm:"column:role_id" json:"roleId"`                                                              // 角色id
	URole        URole       `gorm:"joinForeignKey:role_id;foreignKey:id;references:RoleID" json:"uRoleList"`                   // 角色表
	PermissionID int         `gorm:"column:permission_id" json:"permissionId"`                                                  // 权限id
	UPermission  UPermission `gorm:"joinForeignKey:permission_id;foreignKey:id;references:PermissionID" json:"uPermissionList"` // 权限表
}

// TableName get sql table name.获取数据库表名
func (m *URolePermission) TableName() string {
	return "u_role_permission"
}

// URolePermissionColumns get sql column name.获取数据库列名
var URolePermissionColumns = struct {
	ID           string
	RoleID       string
	PermissionID string
}{
	ID:           "id",
	RoleID:       "role_id",
	PermissionID: "permission_id",
}

// USite 站点表
type USite struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`            // ID
	SiteNumber   string    `gorm:"column:site_number" json:"siteNumber"`     // 站点编码
	SiteName     string    `gorm:"column:site_name" json:"siteName"`         // 站点名称
	CountryCode  string    `gorm:"column:country_code" json:"countryCode"`   // 国家二字码
	TrackAddress string    `gorm:"column:track_address" json:"trackAddress"` // 轨迹地址
	Remarks      string    `gorm:"column:remarks" json:"remarks"`            // 备注信息
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建用户
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新时间
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	IsDelete     int       `gorm:"column:is_delete" json:"isDelete"`         // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *USite) TableName() string {
	return "u_site"
}

// USiteColumns get sql column name.获取数据库列名
var USiteColumns = struct {
	ID           string
	SiteNumber   string
	SiteName     string
	CountryCode  string
	TrackAddress string
	Remarks      string
	CreateUser   string
	CreateTime   string
	UpdateUser   string
	UpdateTime   string
	Version      string
	IsDelete     string
}{
	ID:           "id",
	SiteNumber:   "site_number",
	SiteName:     "site_name",
	CountryCode:  "country_code",
	TrackAddress: "track_address",
	Remarks:      "remarks",
	CreateUser:   "create_user",
	CreateTime:   "create_time",
	UpdateUser:   "update_user",
	UpdateTime:   "update_time",
	Version:      "version",
	IsDelete:     "is_delete",
}

// UUser 用户表
type UUser struct {
	ID                    int       `gorm:"primaryKey;column:id" json:"-"`                                                             // 主键
	Code                  string    `gorm:"column:code" json:"code"`                                                                   // 员工编号
	Username              string    `gorm:"column:username" json:"username"`                                                           // 账户名称
	EnglishName           string    `gorm:"column:english_name" json:"englishName"`                                                    // 英文名称
	Password              string    `gorm:"column:password" json:"password"`                                                           // 密码/不得明文
	RealName              string    `gorm:"column:real_name" json:"realName"`                                                          // 真实姓名
	Phone                 string    `gorm:"column:phone" json:"phone"`                                                                 // 手机/联系方式
	Roles                 string    `gorm:"column:roles" json:"roles"`                                                                 // 角色组逗号分割
	OrgID                 int       `gorm:"column:org_id" json:"orgId"`                                                                // 组织架构ID
	OrgCode               string    `gorm:"column:org_code" json:"orgCode"`                                                            // 组织机构
	OrgCodes              string    `gorm:"column:org_codes" json:"orgCodes"`                                                          // 组织机构父ID
	PlatformCode          string    `gorm:"column:platform_code" json:"platformCode"`                                                  // 平台代码(绑定平台)
	PlatformName          string    `gorm:"column:platform_name" json:"platformName"`                                                  // 平台名称
	HasAudit              int       `gorm:"column:has_audit" json:"hasAudit"`                                                          // 是否有审核权限(组织机构) 0:没有权限 (默认) 1:有权限
	UserState             int       `gorm:"column:user_state" json:"userState"`                                                        // 用户状态 1:在职 2:离职 3:休假
	AccountState          int       `gorm:"column:account_state" json:"accountState"`                                                  // 账户状态  0:停用  1:启用
	PermissionType        int       `gorm:"column:permission_type" json:"permissionType"`                                              // 用户权限类型  1:普通用户 2:个性化  3:未授权用户
	IDcard                string    `gorm:"column:idcard" json:"idcard"`                                                               // 身份证
	Birthday              time.Time `gorm:"column:birthday" json:"birthday"`                                                           // 生日
	Sex                   int       `gorm:"column:sex" json:"sex"`                                                                     // 性别 0:男 1:女
	Wechat                string    `gorm:"column:wechat" json:"wechat"`                                                               // 微信
	Qq                    string    `gorm:"column:qq" json:"qq"`                                                                       // QQ
	SocialSecurity        int       `gorm:"column:social_security" json:"socialSecurity"`                                              // 社保 0:否 1:是
	AccumulationFund      int       `gorm:"column:accumulation_fund" json:"accumulationFund"`                                          // 公积金 0:否 1:是
	EntryTime             time.Time `gorm:"column:entry_time" json:"entryTime"`                                                        // 入职日期
	DepartureTime         time.Time `gorm:"column:departure_time" json:"departureTime"`                                                // 离职日期
	Email                 string    `gorm:"column:email" json:"email"`                                                                 // 邮箱
	Nailing               string    `gorm:"column:nailing" json:"nailing"`                                                             // 钉钉
	Weibo                 string    `gorm:"column:weibo" json:"weibo"`                                                                 // 微博
	Education             int       `gorm:"column:education" json:"education"`                                                         // 学历 0:小学 1:初中 2:高中 3:大专 4:本科 5: 硕士 6:博士
	Resume                string    `gorm:"column:resume" json:"resume"`                                                               // 简历
	Avatar                string    `gorm:"column:avatar" json:"avatar"`                                                               // 头像
	Address               string    `gorm:"column:address" json:"address"`                                                             // 住址
	Position              string    `gorm:"column:position" json:"position"`                                                           // 职位
	EmergencyContact      string    `gorm:"column:emergency_contact" json:"emergencyContact"`                                          // 紧急联系人
	EmergencyContactPhone string    `gorm:"column:emergency_contact_phone" json:"emergencyContactPhone"`                               // 紧急联系人电话
	PositionName          string    `gorm:"column:position_name" json:"positionName"`                                                  // 职位名称
	UPosition             UPosition `gorm:"joinForeignKey:position_name;foreignKey:name;references:PositionName" json:"uPositionList"` // 职务维护
	SiteNumber            string    `gorm:"column:site_number" json:"siteNumber"`                                                      // 站点编号
	SiteName              string    `gorm:"column:site_name" json:"siteName"`                                                          // 站点名称
	Remark                string    `gorm:"column:remark" json:"remark"`                                                               // 备注
	CreateTime            time.Time `gorm:"column:create_time" json:"createTime"`                                                      // 创建时间
	CreateUser            int       `gorm:"column:create_user" json:"createUser"`                                                      // 创建用户
	UpdateTime            time.Time `gorm:"column:update_time" json:"updateTime"`                                                      // 更新时间
	UpdateUser            int       `gorm:"column:update_user" json:"updateUser"`                                                      // 更新用户
	Version               int       `gorm:"column:version" json:"version"`                                                             // 乐观锁
	Deleted               int       `gorm:"column:deleted" json:"deleted"`                                                             // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *UUser) TableName() string {
	return "u_user"
}

// UUserColumns get sql column name.获取数据库列名
var UUserColumns = struct {
	ID                    string
	Code                  string
	Username              string
	EnglishName           string
	Password              string
	RealName              string
	Phone                 string
	Roles                 string
	OrgID                 string
	OrgCode               string
	OrgCodes              string
	PlatformCode          string
	PlatformName          string
	HasAudit              string
	UserState             string
	AccountState          string
	PermissionType        string
	IDcard                string
	Birthday              string
	Sex                   string
	Wechat                string
	Qq                    string
	SocialSecurity        string
	AccumulationFund      string
	EntryTime             string
	DepartureTime         string
	Email                 string
	Nailing               string
	Weibo                 string
	Education             string
	Resume                string
	Avatar                string
	Address               string
	Position              string
	EmergencyContact      string
	EmergencyContactPhone string
	PositionName          string
	SiteNumber            string
	SiteName              string
	Remark                string
	CreateTime            string
	CreateUser            string
	UpdateTime            string
	UpdateUser            string
	Version               string
	Deleted               string
}{
	ID:                    "id",
	Code:                  "code",
	Username:              "username",
	EnglishName:           "english_name",
	Password:              "password",
	RealName:              "real_name",
	Phone:                 "phone",
	Roles:                 "roles",
	OrgID:                 "org_id",
	OrgCode:               "org_code",
	OrgCodes:              "org_codes",
	PlatformCode:          "platform_code",
	PlatformName:          "platform_name",
	HasAudit:              "has_audit",
	UserState:             "user_state",
	AccountState:          "account_state",
	PermissionType:        "permission_type",
	IDcard:                "idcard",
	Birthday:              "birthday",
	Sex:                   "sex",
	Wechat:                "wechat",
	Qq:                    "qq",
	SocialSecurity:        "social_security",
	AccumulationFund:      "accumulation_fund",
	EntryTime:             "entry_time",
	DepartureTime:         "departure_time",
	Email:                 "email",
	Nailing:               "nailing",
	Weibo:                 "weibo",
	Education:             "education",
	Resume:                "resume",
	Avatar:                "avatar",
	Address:               "address",
	Position:              "position",
	EmergencyContact:      "emergency_contact",
	EmergencyContactPhone: "emergency_contact_phone",
	PositionName:          "position_name",
	SiteNumber:            "site_number",
	SiteName:              "site_name",
	Remark:                "remark",
	CreateTime:            "create_time",
	CreateUser:            "create_user",
	UpdateTime:            "update_time",
	UpdateUser:            "update_user",
	Version:               "version",
	Deleted:               "deleted",
}

// UUserConfig [...]
type UUserConfig struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`            // ID
	UserID       int       `gorm:"column:user_id" json:"userId"`             // 用户id
	ConfigKey    string    `gorm:"column:config_key" json:"configKey"`       // 配置key
	ConfigValue  string    `gorm:"column:config_value" json:"configValue"`   // 配置值
	DefaultValue string    `gorm:"column:default_value" json:"defaultValue"` // 默认值
	Remarks      string    `gorm:"column:remarks" json:"remarks"`            // 备注说明
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建用户
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新时间
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *UUserConfig) TableName() string {
	return "u_user_config"
}

// UUserConfigColumns get sql column name.获取数据库列名
var UUserConfigColumns = struct {
	ID           string
	UserID       string
	ConfigKey    string
	ConfigValue  string
	DefaultValue string
	Remarks      string
	CreateTime   string
	CreateUser   string
	UpdateTime   string
	UpdateUser   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	UserID:       "user_id",
	ConfigKey:    "config_key",
	ConfigValue:  "config_value",
	DefaultValue: "default_value",
	Remarks:      "remarks",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateTime:   "update_time",
	UpdateUser:   "update_user",
	Version:      "version",
	Deleted:      "deleted",
}

// UUserPermission 用户_权限
type UUserPermission struct {
	UserID       int         `gorm:"column:user_id" json:"userId"`                                                              // 用户ID
	UUser        UUser       `gorm:"joinForeignKey:user_id;foreignKey:id;references:UserID" json:"uUserList"`                   // 用户表
	PermissionID int         `gorm:"column:permission_id" json:"permissionId"`                                                  // 权限ID
	UPermission  UPermission `gorm:"joinForeignKey:permission_id;foreignKey:id;references:PermissionID" json:"uPermissionList"` // 权限表
}

// TableName get sql table name.获取数据库表名
func (m *UUserPermission) TableName() string {
	return "u_user_permission"
}

// UUserPermissionColumns get sql column name.获取数据库列名
var UUserPermissionColumns = struct {
	UserID       string
	PermissionID string
}{
	UserID:       "user_id",
	PermissionID: "permission_id",
}

// WMicroTicket 微工单
type WMicroTicket struct {
	ID               int            `gorm:"primaryKey;column:id" json:"-"`                     // 微工单
	TicketName       string         `gorm:"column:ticket_name" json:"ticketName"`              // 工单名称
	TicketType       string         `gorm:"column:ticket_type" json:"ticketType"`              // 工单类型
	TicketStatus     string         `gorm:"column:ticket_status" json:"ticketStatus"`          // 工单状态
	Sort             int            `gorm:"column:sort" json:"sort"`                           // 优先级
	TicketDescribe   string         `gorm:"column:ticket_describe" json:"ticketDescribe"`      // 工单描述
	Appendix         string         `gorm:"column:appendix" json:"appendix"`                   // 工单附件
	LastFollower     int            `gorm:"column:last_follower" json:"lastFollower"`          // 最后跟进人
	LastFollowTime   time.Time      `gorm:"column:last_follow_time" json:"lastFollowTime"`     // 最后跟进时间
	ExecuteUser      datatypes.JSON `gorm:"column:execute_user" json:"executeUser"`            // 执行人
	PlanFinishTime   time.Time      `gorm:"column:plan_finish_time" json:"planFinishTime"`     // 计划完成时间
	ActualFinishTime time.Time      `gorm:"column:actual_finish_time" json:"actualFinishTime"` // 实际完成时间
	TimeConsuming    int            `gorm:"column:time_consuming" json:"timeConsuming"`        // 工时
	CreateUser       int            `gorm:"column:create_user" json:"createUser"`              // 创建人
	CreateTime       time.Time      `gorm:"column:create_time" json:"createTime"`              // 创建时间
	UpdateUser       int            `gorm:"column:update_user" json:"updateUser"`              // 更新人
	UpdateTime       time.Time      `gorm:"column:update_time" json:"updateTime"`              // 更新时间
	Version          int            `gorm:"column:version" json:"version"`                     // 乐观锁
	Deleted          int            `gorm:"column:deleted" json:"deleted"`                     // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WMicroTicket) TableName() string {
	return "w_micro_ticket"
}

// WMicroTicketColumns get sql column name.获取数据库列名
var WMicroTicketColumns = struct {
	ID               string
	TicketName       string
	TicketType       string
	TicketStatus     string
	Sort             string
	TicketDescribe   string
	Appendix         string
	LastFollower     string
	LastFollowTime   string
	ExecuteUser      string
	PlanFinishTime   string
	ActualFinishTime string
	TimeConsuming    string
	CreateUser       string
	CreateTime       string
	UpdateUser       string
	UpdateTime       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	TicketName:       "ticket_name",
	TicketType:       "ticket_type",
	TicketStatus:     "ticket_status",
	Sort:             "sort",
	TicketDescribe:   "ticket_describe",
	Appendix:         "appendix",
	LastFollower:     "last_follower",
	LastFollowTime:   "last_follow_time",
	ExecuteUser:      "execute_user",
	PlanFinishTime:   "plan_finish_time",
	ActualFinishTime: "actual_finish_time",
	TimeConsuming:    "time_consuming",
	CreateUser:       "create_user",
	CreateTime:       "create_time",
	UpdateUser:       "update_user",
	UpdateTime:       "update_time",
	Version:          "version",
	Deleted:          "deleted",
}

// WMicroTicketFollowUpInfo 微工单跟进信息
type WMicroTicketFollowUpInfo struct {
	ID               int            `gorm:"primaryKey;column:id" json:"-"`                     // 主键id
	TicketID         int            `gorm:"column:ticket_id" json:"ticketId"`                  // 微工单id
	FollowUpDescribe string         `gorm:"column:follow_up_describe" json:"followUpDescribe"` // 跟进描述
	Follower         datatypes.JSON `gorm:"column:follower" json:"follower"`                   // 跟进人
	FollowUpStatus   int            `gorm:"column:follow_up_status" json:"followUpStatus"`     // 跟进状态(0:正常跟进1:完结任务2:作废任务)
	CreateTime       time.Time      `gorm:"column:create_time" json:"createTime"`              // 创建时间
	CreateUser       int            `gorm:"column:create_user" json:"createUser"`              // 创建人
	UpdateTime       time.Time      `gorm:"column:update_time" json:"updateTime"`              // 更新时间
	UpdateUser       int            `gorm:"column:update_user" json:"updateUser"`              // 更新人
	Version          int            `gorm:"column:version" json:"version"`                     // 乐观锁
	Deleted          int            `gorm:"column:deleted" json:"deleted"`                     // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WMicroTicketFollowUpInfo) TableName() string {
	return "w_micro_ticket_follow_up_info"
}

// WMicroTicketFollowUpInfoColumns get sql column name.获取数据库列名
var WMicroTicketFollowUpInfoColumns = struct {
	ID               string
	TicketID         string
	FollowUpDescribe string
	Follower         string
	FollowUpStatus   string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	TicketID:         "ticket_id",
	FollowUpDescribe: "follow_up_describe",
	Follower:         "follower",
	FollowUpStatus:   "follow_up_status",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
}

// WMicroTicketType 工单类型
type WMicroTicketType struct {
	ID           int       `gorm:"column:id" json:"id"`                      // 主键id
	TypeName     string    `gorm:"column:type_name" json:"typeName"`         // 类型名称
	EnglishName  string    `gorm:"column:english_name" json:"englishName"`   // 英文名称
	TypeStatus   []uint8   `gorm:"column:type_status" json:"typeStatus"`     // 状态（0:禁用1:启用）
	TypeDescribe string    `gorm:"column:type_describe" json:"typeDescribe"` // 类型描述
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`     // 创建用户
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`     // 创建时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`     // 更新用户
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`     // 更新时间
	Version      int       `gorm:"column:version" json:"version"`            // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`            // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WMicroTicketType) TableName() string {
	return "w_micro_ticket_type"
}

// WMicroTicketTypeColumns get sql column name.获取数据库列名
var WMicroTicketTypeColumns = struct {
	ID           string
	TypeName     string
	EnglishName  string
	TypeStatus   string
	TypeDescribe string
	CreateUser   string
	CreateTime   string
	UpdateUser   string
	UpdateTime   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	TypeName:     "type_name",
	EnglishName:  "english_name",
	TypeStatus:   "type_status",
	TypeDescribe: "type_describe",
	CreateUser:   "create_user",
	CreateTime:   "create_time",
	UpdateUser:   "update_user",
	UpdateTime:   "update_time",
	Version:      "version",
	Deleted:      "deleted",
}

// WWorkOrder 工单表
type WWorkOrder struct {
	ID                 int           `gorm:"primaryKey;column:id" json:"-"`                                                       // 主键
	Code               string        `gorm:"column:code" json:"code"`                                                             // 工单编号
	Title              string        `gorm:"column:title" json:"title"`                                                           // 标题
	Content            string        `gorm:"column:content" json:"content"`                                                       // 内容
	HistoricalContent  string        `gorm:"column:historical_content" json:"historicalContent"`                                  // 历史内容
	OrganizationID     int           `gorm:"column:organization_id" json:"organizationId"`                                        // 组织机构ID
	OrganizationIDs    string        `gorm:"column:organization_ids" json:"organizationIds"`                                      // 组织机构ids
	ReviewerID         int           `gorm:"column:reviewer_id" json:"reviewerId"`                                                // 审核人ID
	Resources          string        `gorm:"column:resources" json:"resources"`                                                   // 资源列表
	State              string        `gorm:"column:state" json:"state"`                                                           // 工单状态,            草稿:新建但未指定部门一created,            审核中:新建并指定审核部门- reviewing,            审核通过:部门审核通过,但未拆分任务，指定实施人一reviewed,            审核驳回:部门审核人审核不通过,驳回一review_callback,            任务关闭:审核驳回后需求放关闭工单一closed,            实施中:拆分任务指定实施人员一implementing,            实施完成:所有任务完成一implemented,            审核人验收:部门审核人验收一reviewer_acceptance,            需求方验收-demander_ acceptance,            终结 over,
	Sort               int           `gorm:"column:sort" json:"sort"`                                                             // 优先级,            紧急，高，正常，低
	Type               int           `gorm:"column:type" json:"type"`                                                             // 工单类型
	EstimateFinishTime time.Time     `gorm:"column:estimate_finish_time" json:"estimateFinishTime"`                               // 预计完成时间
	ActualFinishTime   time.Time     `gorm:"column:actual_finish_time" json:"actualFinishTime"`                                   // 实际完成时间
	Remark             string        `gorm:"column:remark" json:"remark"`                                                         // 备注
	IsCooperation      int           `gorm:"column:is_cooperation" json:"isCooperation"`                                          // 是否为多部门协作
	CooperationOrgIDs  string        `gorm:"column:cooperation_org_ids" json:"cooperationOrgIds"`                                 // 多部门协作部门ids集合
	TagName            string        `gorm:"column:tag_name" json:"tagName"`                                                      // 标签名称
	WWorkOrderTag      WWorkOrderTag `gorm:"joinForeignKey:tag_name;foreignKey:name;references:TagName" json:"wWorkOrderTagList"` // 标签维护
	EvaluationGrade    int           `gorm:"column:evaluation_grade" json:"evaluationGrade"`                                      // 评价等级1-10
	EvaluationComment  string        `gorm:"column:evaluation_comment" json:"evaluationComment"`                                  // 评价留言
	CreateTime         time.Time     `gorm:"column:create_time" json:"createTime"`                                                // 创建时间
	CreateUser         int           `gorm:"column:create_user" json:"createUser"`                                                // 创建用户
	UpdateTime         time.Time     `gorm:"column:update_time" json:"updateTime"`                                                // 更新时间
	UpdateUser         int           `gorm:"column:update_user" json:"updateUser"`                                                // 更新时间
	Version            int           `gorm:"column:version" json:"version"`                                                       // 乐观锁
	Deleted            int           `gorm:"column:deleted" json:"deleted"`                                                       // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WWorkOrder) TableName() string {
	return "w_work_order"
}

// WWorkOrderColumns get sql column name.获取数据库列名
var WWorkOrderColumns = struct {
	ID                 string
	Code               string
	Title              string
	Content            string
	HistoricalContent  string
	OrganizationID     string
	OrganizationIDs    string
	ReviewerID         string
	Resources          string
	State              string
	Sort               string
	Type               string
	EstimateFinishTime string
	ActualFinishTime   string
	Remark             string
	IsCooperation      string
	CooperationOrgIDs  string
	TagName            string
	EvaluationGrade    string
	EvaluationComment  string
	CreateTime         string
	CreateUser         string
	UpdateTime         string
	UpdateUser         string
	Version            string
	Deleted            string
}{
	ID:                 "id",
	Code:               "code",
	Title:              "title",
	Content:            "content",
	HistoricalContent:  "historical_content",
	OrganizationID:     "organization_id",
	OrganizationIDs:    "organization_ids",
	ReviewerID:         "reviewer_id",
	Resources:          "resources",
	State:              "state",
	Sort:               "sort",
	Type:               "type",
	EstimateFinishTime: "estimate_finish_time",
	ActualFinishTime:   "actual_finish_time",
	Remark:             "remark",
	IsCooperation:      "is_cooperation",
	CooperationOrgIDs:  "cooperation_org_ids",
	TagName:            "tag_name",
	EvaluationGrade:    "evaluation_grade",
	EvaluationComment:  "evaluation_comment",
	CreateTime:         "create_time",
	CreateUser:         "create_user",
	UpdateTime:         "update_time",
	UpdateUser:         "update_user",
	Version:            "version",
	Deleted:            "deleted",
}

// WWorkOrderComment 评论表
type WWorkOrderComment struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	ParentID   int       `gorm:"column:parent_id" json:"parentId"`     // 工单ID
	Type       int       `gorm:"column:type" json:"type"`              // 类型 1:工单 2:任务
	Content    string    `gorm:"column:content" json:"content"`        // 留言内容
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WWorkOrderComment) TableName() string {
	return "w_work_order_comment"
}

// WWorkOrderCommentColumns get sql column name.获取数据库列名
var WWorkOrderCommentColumns = struct {
	ID         string
	ParentID   string
	Type       string
	Content    string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	ParentID:   "parent_id",
	Type:       "type",
	Content:    "content",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// WWorkOrderDailyReport [...]
type WWorkOrderDailyReport struct {
	ID           int       `gorm:"primaryKey;column:id" json:"-"`              // 主键
	CompleteTask string    `gorm:"column:complete_task" json:"completeTask"`   // 完成的任务
	TodoTask     string    `gorm:"column:todo_task" json:"todoTask"`           // 待办任务
	AssistTask   string    `gorm:"column:assist_task" json:"assistTask"`       // 需要协助任务
	State        string    `gorm:"column:state" json:"state"`                  // 日报状态
	Remark       string    `gorm:"column:remark" json:"remark"`                // 备注
	Resources    []byte    `gorm:"column:resources" json:"resources"`          // 附件
	CopyToUserID []byte    `gorm:"column:copy_to_user_id" json:"copyToUserId"` // 抄送人ID
	ReadTime     time.Time `gorm:"column:read_time" json:"readTime"`           // 阅读时间
	CreateTime   time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	CreateUser   int       `gorm:"column:create_user" json:"createUser"`       // 创建用户
	UpdateTime   time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
	UpdateUser   int       `gorm:"column:update_user" json:"updateUser"`       // 更新时间
	Version      int       `gorm:"column:version" json:"version"`              // 乐观锁
	Deleted      int       `gorm:"column:deleted" json:"deleted"`              // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WWorkOrderDailyReport) TableName() string {
	return "w_work_order_daily_report"
}

// WWorkOrderDailyReportColumns get sql column name.获取数据库列名
var WWorkOrderDailyReportColumns = struct {
	ID           string
	CompleteTask string
	TodoTask     string
	AssistTask   string
	State        string
	Remark       string
	Resources    string
	CopyToUserID string
	ReadTime     string
	CreateTime   string
	CreateUser   string
	UpdateTime   string
	UpdateUser   string
	Version      string
	Deleted      string
}{
	ID:           "id",
	CompleteTask: "complete_task",
	TodoTask:     "todo_task",
	AssistTask:   "assist_task",
	State:        "state",
	Remark:       "remark",
	Resources:    "resources",
	CopyToUserID: "copy_to_user_id",
	ReadTime:     "read_time",
	CreateTime:   "create_time",
	CreateUser:   "create_user",
	UpdateTime:   "update_time",
	UpdateUser:   "update_user",
	Version:      "version",
	Deleted:      "deleted",
}

// WWorkOrderOperationLog 工单日志
type WWorkOrderOperationLog struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`
	OrderID          int       `gorm:"column:order_id" json:"orderId"`                   // 工单id
	TaskID           int       `gorm:"column:task_id" json:"taskId"`                     // 任务id
	OperationContent string    `gorm:"column:operation_content" json:"operationContent"` // 操作内容
	Remark           string    `gorm:"column:remark" json:"remark"`                      // 备注
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`             // 创建用户
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`             // 更新时间
	Version          int       `gorm:"column:version" json:"version"`                    // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                    // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WWorkOrderOperationLog) TableName() string {
	return "w_work_order_operation_log"
}

// WWorkOrderOperationLogColumns get sql column name.获取数据库列名
var WWorkOrderOperationLogColumns = struct {
	ID               string
	OrderID          string
	TaskID           string
	OperationContent string
	Remark           string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
}{
	ID:               "id",
	OrderID:          "order_id",
	TaskID:           "task_id",
	OperationContent: "operation_content",
	Remark:           "remark",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
}

// WWorkOrderReviewing 工单_待审核人员
type WWorkOrderReviewing struct {
	OrderID    int `gorm:"column:order_id" json:"orderId"`
	ReviewerID int `gorm:"column:reviewer_id" json:"reviewerId"`
}

// TableName get sql table name.获取数据库表名
func (m *WWorkOrderReviewing) TableName() string {
	return "w_work_order_reviewing"
}

// WWorkOrderReviewingColumns get sql column name.获取数据库列名
var WWorkOrderReviewingColumns = struct {
	OrderID    string
	ReviewerID string
}{
	OrderID:    "order_id",
	ReviewerID: "reviewer_id",
}

// WWorkOrderTag 标签维护
type WWorkOrderTag struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	Name       string    `gorm:"column:name" json:"name"`              // 标签名称
	Type       int       `gorm:"column:type" json:"type"`              // 类型:保留字段
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
}

// TableName get sql table name.获取数据库表名
func (m *WWorkOrderTag) TableName() string {
	return "w_work_order_tag"
}

// WWorkOrderTagColumns get sql column name.获取数据库列名
var WWorkOrderTagColumns = struct {
	ID         string
	Name       string
	Type       string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
}{
	ID:         "id",
	Name:       "name",
	Type:       "type",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
}

// WWorkOrderTask 工单任务表
type WWorkOrderTask struct {
	ID                   int           `gorm:"primaryKey;column:id" json:"-"`                                                       // 主键
	Code                 string        `gorm:"column:code" json:"code"`                                                             // 任务编号
	OrderID              int           `gorm:"column:order_id" json:"orderId"`                                                      // 工单ID
	WWorkOrder           WWorkOrder    `gorm:"joinForeignKey:order_id;foreignKey:id;references:OrderID" json:"wWorkOrderList"`      // 工单表
	Title                string        `gorm:"column:title" json:"title"`                                                           // 标题
	Content              string        `gorm:"column:content" json:"content"`                                                       // 内容
	UserID               int           `gorm:"column:user_id" json:"userId"`                                                        // 实施人
	OrganizationID       int           `gorm:"column:organization_id" json:"organizationId"`                                        // 工单所属组织机构id
	State                string        `gorm:"column:state" json:"state"`                                                           // 任务状态
	Sort                 int           `gorm:"column:sort" json:"sort"`                                                             // 优先级,            紧急，高，正常，低
	EstimateFinishTime   time.Time     `gorm:"column:estimate_finish_time" json:"estimateFinishTime"`                               // 预计完成时间
	TaskBeginTime        time.Time     `gorm:"column:task_begin_time" json:"taskBeginTime"`                                         // 任务开始时间
	ActualFinishTime     time.Time     `gorm:"column:actual_finish_time" json:"actualFinishTime"`                                   // 实际完成时间
	EstimateWorkingHours string        `gorm:"column:estimate_working_hours" json:"estimateWorkingHours"`                           // 预计完成工时
	ActualWorkingHours   string        `gorm:"column:actual_working_hours" json:"actualWorkingHours"`                               // 实际完成使用工时
	Remark               string        `gorm:"column:remark" json:"remark"`                                                         // 备注
	TagName              string        `gorm:"column:tag_name" json:"tagName"`                                                      // 标签名称
	WWorkOrderTag        WWorkOrderTag `gorm:"joinForeignKey:tag_name;foreignKey:name;references:TagName" json:"wWorkOrderTagList"` // 标签维护
	EvaluationGrade      int           `gorm:"column:evaluation_grade" json:"evaluationGrade"`                                      // 评价等级
	EvaluationComment    string        `gorm:"column:evaluation_comment" json:"evaluationComment"`                                  // 评价内容
	CreateTime           time.Time     `gorm:"column:create_time" json:"createTime"`                                                // 创建时间
	CreateUser           int           `gorm:"column:create_user" json:"createUser"`                                                // 创建用户
	UpdateTime           time.Time     `gorm:"column:update_time" json:"updateTime"`                                                // 更新时间
	UpdateUser           int           `gorm:"column:update_user" json:"updateUser"`                                                // 更新时间
	Version              int           `gorm:"column:version" json:"version"`                                                       // 乐观锁
	Deleted              int           `gorm:"column:deleted" json:"deleted"`                                                       // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WWorkOrderTask) TableName() string {
	return "w_work_order_task"
}

// WWorkOrderTaskColumns get sql column name.获取数据库列名
var WWorkOrderTaskColumns = struct {
	ID                   string
	Code                 string
	OrderID              string
	Title                string
	Content              string
	UserID               string
	OrganizationID       string
	State                string
	Sort                 string
	EstimateFinishTime   string
	TaskBeginTime        string
	ActualFinishTime     string
	EstimateWorkingHours string
	ActualWorkingHours   string
	Remark               string
	TagName              string
	EvaluationGrade      string
	EvaluationComment    string
	CreateTime           string
	CreateUser           string
	UpdateTime           string
	UpdateUser           string
	Version              string
	Deleted              string
}{
	ID:                   "id",
	Code:                 "code",
	OrderID:              "order_id",
	Title:                "title",
	Content:              "content",
	UserID:               "user_id",
	OrganizationID:       "organization_id",
	State:                "state",
	Sort:                 "sort",
	EstimateFinishTime:   "estimate_finish_time",
	TaskBeginTime:        "task_begin_time",
	ActualFinishTime:     "actual_finish_time",
	EstimateWorkingHours: "estimate_working_hours",
	ActualWorkingHours:   "actual_working_hours",
	Remark:               "remark",
	TagName:              "tag_name",
	EvaluationGrade:      "evaluation_grade",
	EvaluationComment:    "evaluation_comment",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateTime:           "update_time",
	UpdateUser:           "update_user",
	Version:              "version",
	Deleted:              "deleted",
}

// WmsAirBill 航空提单
type WmsAirBill struct {
	ID              int       `gorm:"primaryKey;column:id" json:"-"`
	AirBillNumber   string    `gorm:"column:air_bill_number" json:"airBillNumber"`    // 航空提单号
	BigBagCount     int       `gorm:"column:big_bag_count" json:"bigBagCount"`        // 大包数量
	BigBagWeight    float64   `gorm:"column:big_bag_weight" json:"bigBagWeight"`      // 大包重量
	SmallBagCount   int       `gorm:"column:small_bag_count" json:"smallBagCount"`    // 小包数量
	DownstreamCodes string    `gorm:"column:downstream_codes" json:"downstreamCodes"` // 下家三字码，可能有多个
	CreateTime      time.Time `gorm:"column:create_time" json:"createTime"`           // 创建时间
	CreateUser      int       `gorm:"column:create_user" json:"createUser"`           // 创建用户
	UpdateTime      time.Time `gorm:"column:update_time" json:"updateTime"`           // 更新时间
	UpdateUser      int       `gorm:"column:update_user" json:"updateUser"`           // 更新时间
	Version         int       `gorm:"column:version" json:"version"`                  // 乐观锁
	Deleted         int       `gorm:"column:deleted" json:"deleted"`                  // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsAirBill) TableName() string {
	return "wms_air_bill"
}

// WmsAirBillColumns get sql column name.获取数据库列名
var WmsAirBillColumns = struct {
	ID              string
	AirBillNumber   string
	BigBagCount     string
	BigBagWeight    string
	SmallBagCount   string
	DownstreamCodes string
	CreateTime      string
	CreateUser      string
	UpdateTime      string
	UpdateUser      string
	Version         string
	Deleted         string
}{
	ID:              "id",
	AirBillNumber:   "air_bill_number",
	BigBagCount:     "big_bag_count",
	BigBagWeight:    "big_bag_weight",
	SmallBagCount:   "small_bag_count",
	DownstreamCodes: "downstream_codes",
	CreateTime:      "create_time",
	CreateUser:      "create_user",
	UpdateTime:      "update_time",
	UpdateUser:      "update_user",
	Version:         "version",
	Deleted:         "deleted",
}

// WmsCn35DispatchNumRecord cn35预报id表
type WmsCn35DispatchNumRecord struct {
	ID                 int     `gorm:"primaryKey;column:id" json:"-"`
	BigBagID           int     `gorm:"column:big_bag_id" json:"bigBagId"`                     // 大包id
	Cn35DispatchNumber float64 `gorm:"column:cn35_dispatch_number" json:"cn35DispatchNumber"` // cn35预报唯一标识(最大 999999)
}

// TableName get sql table name.获取数据库表名
func (m *WmsCn35DispatchNumRecord) TableName() string {
	return "wms_cn35_dispatch_num_record"
}

// WmsCn35DispatchNumRecordColumns get sql column name.获取数据库列名
var WmsCn35DispatchNumRecordColumns = struct {
	ID                 string
	BigBagID           string
	Cn35DispatchNumber string
}{
	ID:                 "id",
	BigBagID:           "big_bag_id",
	Cn35DispatchNumber: "cn35_dispatch_number",
}

// WmsExpressBigbagCollection 快递大包揽收
type WmsExpressBigbagCollection struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`                // 主键id
	BigBagNum      string    `gorm:"column:bigBag_num" json:"bigBagNum"`           // 大包号
	BigBagWeight   float64   `gorm:"column:bigBag_weight" json:"bigBagWeight"`     // 大包重量
	BigBagQuantity int       `gorm:"column:bigBag_quantity" json:"bigBagQuantity"` // 包裹数量
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`         // 创建人
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`         // 创建时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`         // 更新人
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`         // 更新时间
	Version        int       `gorm:"column:version" json:"version"`                // 乐观锁
	Deleted        int       `gorm:"column:deleted" json:"deleted"`                // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsExpressBigbagCollection) TableName() string {
	return "wms_express_bigbag_collection"
}

// WmsExpressBigbagCollectionColumns get sql column name.获取数据库列名
var WmsExpressBigbagCollectionColumns = struct {
	ID             string
	BigBagNum      string
	BigBagWeight   string
	BigBagQuantity string
	CreateUser     string
	CreateTime     string
	UpdateUser     string
	UpdateTime     string
	Version        string
	Deleted        string
}{
	ID:             "id",
	BigBagNum:      "bigBag_num",
	BigBagWeight:   "bigBag_weight",
	BigBagQuantity: "bigBag_quantity",
	CreateUser:     "create_user",
	CreateTime:     "create_time",
	UpdateUser:     "update_user",
	UpdateTime:     "update_time",
	Version:        "version",
	Deleted:        "deleted",
}

// WmsExpressCollection 快递包裹揽收
type WmsExpressCollection struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`                     // 主键id
	ExpressNum       string    `gorm:"column:express_num" json:"expressNum"`              // 快递单号
	BigBagNum        string    `gorm:"column:bigBag_num" json:"bigBagNum"`                // 大包号
	Weight           float64   `gorm:"column:weight" json:"weight"`                       // 称重重量
	ExpressResource  string    `gorm:"column:express_resource" json:"expressResource"`    // 快递来源
	ImageURL         string    `gorm:"column:image_url" json:"imageUrl"`                  // 快递图片路径
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`              // 操作人
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`              // 创建时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`              // 更新人
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`              // 更新时间
	Version          int       `gorm:"column:version" json:"version"`                     // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                     // 逻辑删除
	Status           int       `gorm:"column:status" json:"status"`                       // 状态 0 作废 1 正常 2 预录状态 默认为1
	IsLogisticsOrder int       `gorm:"column:is_logistics_order" json:"isLogisticsOrder"` // 是否是物流订单 1 是物流订单 0 不是物流订单
	PreEntryUserID   int       `gorm:"column:pre_entry_user_id" json:"preEntryUserId"`    // 预录人
	PreEntryTime     time.Time `gorm:"column:pre_entry_time" json:"preEntryTime"`         // 预录时间
}

// TableName get sql table name.获取数据库表名
func (m *WmsExpressCollection) TableName() string {
	return "wms_express_collection"
}

// WmsExpressCollectionColumns get sql column name.获取数据库列名
var WmsExpressCollectionColumns = struct {
	ID               string
	ExpressNum       string
	BigBagNum        string
	Weight           string
	ExpressResource  string
	ImageURL         string
	CreateUser       string
	CreateTime       string
	UpdateUser       string
	UpdateTime       string
	Version          string
	Deleted          string
	Status           string
	IsLogisticsOrder string
	PreEntryUserID   string
	PreEntryTime     string
}{
	ID:               "id",
	ExpressNum:       "express_num",
	BigBagNum:        "bigBag_num",
	Weight:           "weight",
	ExpressResource:  "express_resource",
	ImageURL:         "image_url",
	CreateUser:       "create_user",
	CreateTime:       "create_time",
	UpdateUser:       "update_user",
	UpdateTime:       "update_time",
	Version:          "version",
	Deleted:          "deleted",
	Status:           "status",
	IsLogisticsOrder: "is_logistics_order",
	PreEntryUserID:   "pre_entry_user_id",
	PreEntryTime:     "pre_entry_time",
}

// WmsExpressSource 快递来源
type WmsExpressSource struct {
	ID          int       `gorm:"primaryKey;column:id" json:"-"`          // 主键id
	ExpressCode string    `gorm:"column:express_code" json:"expressCode"` // 快递代码
	ExpressName string    `gorm:"column:express_name" json:"expressName"` // 快递名称
	Remark      string    `gorm:"column:remark" json:"remark"`            // 备注
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`   // 创建时间
	CreateUser  int       `gorm:"column:create_user" json:"createUser"`   // 创建人
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`   // 更新时间
	UpdateUser  int       `gorm:"column:update_user" json:"updateUser"`   // 更新人
	Version     int       `gorm:"column:version" json:"version"`          // 乐观锁
	Deleted     int       `gorm:"column:deleted" json:"deleted"`          // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsExpressSource) TableName() string {
	return "wms_express_source"
}

// WmsExpressSourceColumns get sql column name.获取数据库列名
var WmsExpressSourceColumns = struct {
	ID          string
	ExpressCode string
	ExpressName string
	Remark      string
	CreateTime  string
	CreateUser  string
	UpdateTime  string
	UpdateUser  string
	Version     string
	Deleted     string
}{
	ID:          "id",
	ExpressCode: "express_code",
	ExpressName: "express_name",
	Remark:      "remark",
	CreateTime:  "create_time",
	CreateUser:  "create_user",
	UpdateTime:  "update_time",
	UpdateUser:  "update_user",
	Version:     "version",
	Deleted:     "deleted",
}

// WmsOperationBigBagInfo 大包
type WmsOperationBigBagInfo struct {
	ID                    int            `gorm:"primaryKey;column:id" json:"-"`
	BigBagNo              string         `gorm:"column:big_bag_no" json:"bigBagNo"`                            // 大包号
	CustomerBigBagNo      string         `gorm:"column:customer_big_bag_no" json:"customerBigBagNo"`           // 客户大包号，来源为内部时，默认为大包号
	CustomerID            int            `gorm:"column:customer_id" json:"customerId"`                         // 客户id，来源为内部时，默认为0
	CustomerAlias         string         `gorm:"column:customer_alias" json:"customerAlias"`                   // 客户别名
	Source                int8           `gorm:"column:source" json:"source"`                                  // 来源，0:内部，1:外部 2自飞
	DownstreamCode        string         `gorm:"column:downstream_code" json:"downstreamCode"`                 // 下家code
	CustomerChannelIDs    string         `gorm:"column:customer_channel_ids" json:"customerChannelIds"`        // 客户渠道id，多个以逗号隔开
	APIBigBagNo           string         `gorm:"column:api_big_bag_no" json:"apiBigBagNo"`                     // api大包号
	ExpressWaybillOssLink string         `gorm:"column:express_waybill_oss_link" json:"expressWaybillOssLink"` // 大包面单OSS地址
	TotalListNo           string         `gorm:"column:total_list_no" json:"totalListNo"`                      // 总单号
	AirBillNumber         string         `gorm:"column:air_bill_number" json:"airBillNumber"`                  // 航空提单号
	Length                float64        `gorm:"column:length" json:"length"`                                  // 长(cm)
	Width                 float64        `gorm:"column:width" json:"width"`                                    // 宽(cm)
	Height                float64        `gorm:"column:height" json:"height"`                                  // 高(cm)
	Weight                float64        `gorm:"column:weight" json:"weight"`                                  // 重量
	SmallBagCount         int            `gorm:"column:small_bag_count" json:"smallBagCount"`                  // 小包数量
	BigBagStatus          string         `gorm:"column:big_bag_status" json:"bigBagStatus"`                    // 大包状态，packaged:已打包、unpacked:未打包
	PackageUserID         int            `gorm:"column:package_user_id" json:"packageUserId"`                  // 打包人用户id
	PackageUserName       string         `gorm:"column:package_user_name" json:"packageUserName"`              // 打包人名称
	BackPackageUserName   string         `gorm:"column:back_package_user_name" json:"backPackageUserName"`     // 上一个打包人名称
	BackPackageUserID     int            `gorm:"column:back_package_user_id" json:"backPackageUserId"`         // 上一个打包人id
	FitterUserID          int            `gorm:"column:fitter_user_id" json:"fitterUserId"`                    // 装配人用户id
	FitterUserName        string         `gorm:"column:fitter_user_name" json:"fitterUserName"`                // 装配人用户名称
	WarehouseCode         string         `gorm:"column:warehouse_code" json:"warehouseCode"`                   // 仓库code
	WarehouseName         string         `gorm:"column:warehouse_name" json:"warehouseName"`                   // 仓库name
	Extra                 datatypes.JSON `gorm:"column:extra" json:"extra"`                                    // 扩展
	CreateTime            time.Time      `gorm:"column:create_time" json:"createTime"`                         // 创建时间
	CreateUser            int            `gorm:"column:create_user" json:"createUser"`                         // 创建用户
	UpdateTime            time.Time      `gorm:"column:update_time" json:"updateTime"`                         // 更新时间
	UpdateUser            int            `gorm:"column:update_user" json:"updateUser"`                         // 更新时间
	Version               int            `gorm:"column:version" json:"version"`                                // 乐观锁
	Deleted               int            `gorm:"column:deleted" json:"deleted"`                                // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsOperationBigBagInfo) TableName() string {
	return "wms_operation_big_bag_info"
}

// WmsOperationBigBagInfoColumns get sql column name.获取数据库列名
var WmsOperationBigBagInfoColumns = struct {
	ID                    string
	BigBagNo              string
	CustomerBigBagNo      string
	CustomerID            string
	CustomerAlias         string
	Source                string
	DownstreamCode        string
	CustomerChannelIDs    string
	APIBigBagNo           string
	ExpressWaybillOssLink string
	TotalListNo           string
	AirBillNumber         string
	Length                string
	Width                 string
	Height                string
	Weight                string
	SmallBagCount         string
	BigBagStatus          string
	PackageUserID         string
	PackageUserName       string
	BackPackageUserName   string
	BackPackageUserID     string
	FitterUserID          string
	FitterUserName        string
	WarehouseCode         string
	WarehouseName         string
	Extra                 string
	CreateTime            string
	CreateUser            string
	UpdateTime            string
	UpdateUser            string
	Version               string
	Deleted               string
}{
	ID:                    "id",
	BigBagNo:              "big_bag_no",
	CustomerBigBagNo:      "customer_big_bag_no",
	CustomerID:            "customer_id",
	CustomerAlias:         "customer_alias",
	Source:                "source",
	DownstreamCode:        "downstream_code",
	CustomerChannelIDs:    "customer_channel_ids",
	APIBigBagNo:           "api_big_bag_no",
	ExpressWaybillOssLink: "express_waybill_oss_link",
	TotalListNo:           "total_list_no",
	AirBillNumber:         "air_bill_number",
	Length:                "length",
	Width:                 "width",
	Height:                "height",
	Weight:                "weight",
	SmallBagCount:         "small_bag_count",
	BigBagStatus:          "big_bag_status",
	PackageUserID:         "package_user_id",
	PackageUserName:       "package_user_name",
	BackPackageUserName:   "back_package_user_name",
	BackPackageUserID:     "back_package_user_id",
	FitterUserID:          "fitter_user_id",
	FitterUserName:        "fitter_user_name",
	WarehouseCode:         "warehouse_code",
	WarehouseName:         "warehouse_name",
	Extra:                 "extra",
	CreateTime:            "create_time",
	CreateUser:            "create_user",
	UpdateTime:            "update_time",
	UpdateUser:            "update_user",
	Version:               "version",
	Deleted:               "deleted",
}

// WmsOperationBigBagInfoAffiliated 大包附属表
type WmsOperationBigBagInfoAffiliated struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`              // ID
	BigBagNo      string    `gorm:"column:big_bag_no" json:"bigBagNo"`          // 大包号
	TotalListNo   string    `gorm:"column:total_list_no" json:"totalListNo"`    // 总单号
	SiteNumber    string    `gorm:"column:site_number" json:"siteNumber"`       // 站点编码
	FlightNo      string    `gorm:"column:flight_no" json:"flightNo"`           // 航班号
	Weight        float64   `gorm:"column:weight" json:"weight"`                // 重量
	ScannedStatus string    `gorm:"column:scanned_status" json:"scannedStatus"` // 扫描状态 done:已完成，unconfirmed:待确认
	ScannedUser   int       `gorm:"column:scanned_user" json:"scannedUser"`     // 扫描人
	ScannedTime   time.Time `gorm:"column:scanned_time" json:"scannedTime"`     // 扫描时间
	Remarks       string    `gorm:"column:remarks" json:"remarks"`              // 备注信息
	CreateUser    int       `gorm:"column:create_user" json:"createUser"`       // 创建用户
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	UpdateUser    int       `gorm:"column:update_user" json:"updateUser"`       // 更新时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
	Version       int       `gorm:"column:version" json:"version"`              // 乐观锁
	IsDelete      int       `gorm:"column:is_delete" json:"isDelete"`           // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsOperationBigBagInfoAffiliated) TableName() string {
	return "wms_operation_big_bag_info_affiliated"
}

// WmsOperationBigBagInfoAffiliatedColumns get sql column name.获取数据库列名
var WmsOperationBigBagInfoAffiliatedColumns = struct {
	ID            string
	BigBagNo      string
	TotalListNo   string
	SiteNumber    string
	FlightNo      string
	Weight        string
	ScannedStatus string
	ScannedUser   string
	ScannedTime   string
	Remarks       string
	CreateUser    string
	CreateTime    string
	UpdateUser    string
	UpdateTime    string
	Version       string
	IsDelete      string
}{
	ID:            "id",
	BigBagNo:      "big_bag_no",
	TotalListNo:   "total_list_no",
	SiteNumber:    "site_number",
	FlightNo:      "flight_no",
	Weight:        "weight",
	ScannedStatus: "scanned_status",
	ScannedUser:   "scanned_user",
	ScannedTime:   "scanned_time",
	Remarks:       "remarks",
	CreateUser:    "create_user",
	CreateTime:    "create_time",
	UpdateUser:    "update_user",
	UpdateTime:    "update_time",
	Version:       "version",
	IsDelete:      "is_delete",
}

// WmsOperationPackageInfo 包裹信息
type WmsOperationPackageInfo struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`
	TrackingNo           string    `gorm:"column:tracking_no" json:"trackingNo"`                      // 快递单号
	TrackingNoType       string    `gorm:"column:tracking_no_type" json:"trackingNoType"`             // 快递单号类型，reference:参考号，transfer:转单号
	LogisticsNumber      string    `gorm:"column:logistics_number" json:"logisticsNumber"`            // 物流单号
	LogisticsNumberFinal string    `gorm:"column:logistics_number_final" json:"logisticsNumberFinal"` // 最终物流单号
	ReferenceNumber      string    `gorm:"column:reference_number" json:"referenceNumber"`            // 参考号
	ReceiveCountry       string    `gorm:"column:receive_country" json:"receiveCountry"`              // 收件人国家
	OrderNo              string    `gorm:"column:order_no" json:"orderNo"`                            // 订单号
	CustomerID           int       `gorm:"column:customer_id" json:"customerId"`                      // 客户id
	Source               int8      `gorm:"column:source" json:"source"`                               // 来源，0:内部，1:外部
	CustomerChannelID    int       `gorm:"column:customer_channel_id" json:"customerChannelId"`       // 客户渠道id
	DownstreamCode       string    `gorm:"column:downstream_code" json:"downstreamCode"`              // 下家code
	APIChannelCode       string    `gorm:"column:api_channel_code" json:"apiChannelCode"`             // api渠道代码
	BigBagNo             string    `gorm:"column:big_bag_no" json:"bigBagNo"`                         // 大包号
	TotalListNo          string    `gorm:"column:total_list_no" json:"totalListNo"`                   // 总单号
	AirBillNumber        string    `gorm:"column:air_bill_number" json:"airBillNumber"`               // 航空提单号
	Weight               float64   `gorm:"column:weight" json:"weight"`                               // 包裹重量(KG),保留4位小数
	OrderWeight          float64   `gorm:"column:order_weight" json:"orderWeight"`                    // 订单预报重量
	WorkConsole          string    `gorm:"column:work_console" json:"workConsole"`                    // 操作台号
	PicOssLink           string    `gorm:"column:pic_oss_link" json:"picOssLink"`                     // 包裹图片oss地址
	ScanTime             time.Time `gorm:"column:scan_time" json:"scanTime"`                          // 扫描时间
	ScanUserID           int       `gorm:"column:scan_user_id" json:"scanUserId"`                     // 扫描人用户id
	ScanUserName         string    `gorm:"column:scan_user_name" json:"scanUserName"`                 // 扫描人用户名字
	PackageUserID        int       `gorm:"column:package_user_id" json:"packageUserId"`               // 打包人用户id
	PackageUserName      string    `gorm:"column:package_user_name" json:"packageUserName"`           // 打包人名称
	PackageTime          time.Time `gorm:"column:package_time" json:"packageTime"`                    // 打包时间
	FitterUserID         int       `gorm:"column:fitter_user_id" json:"fitterUserId"`                 // 装配人用户id
	FitterUserName       string    `gorm:"column:fitter_user_name" json:"fitterUserName"`             // 装配人用户名称
	WarehouseCode        string    `gorm:"column:warehouse_code" json:"warehouseCode"`                // 仓库code
	WarehouseName        string    `gorm:"column:warehouse_name" json:"warehouseName"`                // 仓库名字
	ScanStatus           string    `gorm:"column:scan_status" json:"scanStatus"`                      // 扫描状态 done:已完成，unconfirmed:待确认(包裹异常，如超重、欠费等)，initial：初始化(包裹信息填写错误时，重新录入)
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 创建用户
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新时间
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	Deleted              int       `gorm:"column:deleted" json:"deleted"`                             // 逻辑删除
	Length               float64   `gorm:"column:length" json:"length"`                               // 长 单位:cm
	Width                float64   `gorm:"column:width" json:"width"`                                 // 宽 单位:cm
	Height               float64   `gorm:"column:height" json:"height"`                               // 高 单位:cm
}

// TableName get sql table name.获取数据库表名
func (m *WmsOperationPackageInfo) TableName() string {
	return "wms_operation_package_info"
}

// WmsOperationPackageInfoColumns get sql column name.获取数据库列名
var WmsOperationPackageInfoColumns = struct {
	ID                   string
	TrackingNo           string
	TrackingNoType       string
	LogisticsNumber      string
	LogisticsNumberFinal string
	ReferenceNumber      string
	ReceiveCountry       string
	OrderNo              string
	CustomerID           string
	Source               string
	CustomerChannelID    string
	DownstreamCode       string
	APIChannelCode       string
	BigBagNo             string
	TotalListNo          string
	AirBillNumber        string
	Weight               string
	OrderWeight          string
	WorkConsole          string
	PicOssLink           string
	ScanTime             string
	ScanUserID           string
	ScanUserName         string
	PackageUserID        string
	PackageUserName      string
	PackageTime          string
	FitterUserID         string
	FitterUserName       string
	WarehouseCode        string
	WarehouseName        string
	ScanStatus           string
	CreateTime           string
	CreateUser           string
	UpdateTime           string
	UpdateUser           string
	Version              string
	Deleted              string
	Length               string
	Width                string
	Height               string
}{
	ID:                   "id",
	TrackingNo:           "tracking_no",
	TrackingNoType:       "tracking_no_type",
	LogisticsNumber:      "logistics_number",
	LogisticsNumberFinal: "logistics_number_final",
	ReferenceNumber:      "reference_number",
	ReceiveCountry:       "receive_country",
	OrderNo:              "order_no",
	CustomerID:           "customer_id",
	Source:               "source",
	CustomerChannelID:    "customer_channel_id",
	DownstreamCode:       "downstream_code",
	APIChannelCode:       "api_channel_code",
	BigBagNo:             "big_bag_no",
	TotalListNo:          "total_list_no",
	AirBillNumber:        "air_bill_number",
	Weight:               "weight",
	OrderWeight:          "order_weight",
	WorkConsole:          "work_console",
	PicOssLink:           "pic_oss_link",
	ScanTime:             "scan_time",
	ScanUserID:           "scan_user_id",
	ScanUserName:         "scan_user_name",
	PackageUserID:        "package_user_id",
	PackageUserName:      "package_user_name",
	PackageTime:          "package_time",
	FitterUserID:         "fitter_user_id",
	FitterUserName:       "fitter_user_name",
	WarehouseCode:        "warehouse_code",
	WarehouseName:        "warehouse_name",
	ScanStatus:           "scan_status",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateTime:           "update_time",
	UpdateUser:           "update_user",
	Version:              "version",
	Deleted:              "deleted",
	Length:               "length",
	Width:                "width",
	Height:               "height",
}

// WmsOperationTotalListInfo 总单
type WmsOperationTotalListInfo struct {
	ID                    int            `gorm:"primaryKey;column:id" json:"-"`
	TotalListNo           string         `gorm:"column:total_list_no" json:"totalListNo"`                      // 总单号
	TotalListType         string         `gorm:"column:total_list_type" json:"totalListType"`                  // 总单类型: normal-普货，battery-纯电，站点扫描
	APITotalListNo        string         `gorm:"column:api_total_list_no" json:"apiTotalListNo"`               // api总单号
	CustomerTotalListNo   string         `gorm:"column:customer_total_list_no" json:"customerTotalListNo"`     // 客户总单号
	ExpressWaybillOssLink string         `gorm:"column:express_waybill_oss_link" json:"expressWaybillOssLink"` // 总单面单OSS地址
	FlightNo              string         `gorm:"column:flight_no" json:"flightNo"`                             // 航班号
	Downstream            string         `gorm:"column:downstream" json:"downstream"`                          // 下家
	Weight                float64        `gorm:"column:weight" json:"weight"`                                  // 重量
	BigBagCount           int            `gorm:"column:big_bag_count" json:"bigBagCount"`                      // 大包数量
	SmallBagCount         int            `gorm:"column:small_bag_count" json:"smallBagCount"`                  // 小包数量
	FitterUserID          int            `gorm:"column:fitter_user_id" json:"fitterUserId"`                    // 装配人用户id
	FitterUserName        string         `gorm:"column:fitter_user_name" json:"fitterUserName"`                // 装配人用户名称
	FitStatus             string         `gorm:"column:fit_status" json:"fitStatus"`                           // 装配状态，initial:未装配，done:已装配
	FitOverTime           time.Time      `gorm:"column:fit_over_time" json:"fitOverTime"`                      // 装配完结时间
	WarehouseCode         string         `gorm:"column:warehouse_code" json:"warehouseCode"`                   // 仓库code
	WarehouseName         string         `gorm:"column:warehouse_name" json:"warehouseName"`                   // 仓库名字
	Extra                 datatypes.JSON `gorm:"column:extra" json:"extra"`                                    // 扩展字段
	Remark                string         `gorm:"column:remark" json:"remark"`                                  // 备注
	CreateTime            time.Time      `gorm:"column:create_time" json:"createTime"`                         // 创建时间
	CreateUser            int            `gorm:"column:create_user" json:"createUser"`                         // 创建用户
	UpdateTime            time.Time      `gorm:"column:update_time" json:"updateTime"`                         // 更新时间
	UpdateUser            uint32         `gorm:"column:update_user" json:"updateUser"`                         // 更新时间
	Version               int            `gorm:"column:version" json:"version"`                                // 乐观锁
	Deleted               int            `gorm:"column:deleted" json:"deleted"`                                // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsOperationTotalListInfo) TableName() string {
	return "wms_operation_total_list_info"
}

// WmsOperationTotalListInfoColumns get sql column name.获取数据库列名
var WmsOperationTotalListInfoColumns = struct {
	ID                    string
	TotalListNo           string
	TotalListType         string
	APITotalListNo        string
	CustomerTotalListNo   string
	ExpressWaybillOssLink string
	FlightNo              string
	Downstream            string
	Weight                string
	BigBagCount           string
	SmallBagCount         string
	FitterUserID          string
	FitterUserName        string
	FitStatus             string
	FitOverTime           string
	WarehouseCode         string
	WarehouseName         string
	Extra                 string
	Remark                string
	CreateTime            string
	CreateUser            string
	UpdateTime            string
	UpdateUser            string
	Version               string
	Deleted               string
}{
	ID:                    "id",
	TotalListNo:           "total_list_no",
	TotalListType:         "total_list_type",
	APITotalListNo:        "api_total_list_no",
	CustomerTotalListNo:   "customer_total_list_no",
	ExpressWaybillOssLink: "express_waybill_oss_link",
	FlightNo:              "flight_no",
	Downstream:            "downstream",
	Weight:                "weight",
	BigBagCount:           "big_bag_count",
	SmallBagCount:         "small_bag_count",
	FitterUserID:          "fitter_user_id",
	FitterUserName:        "fitter_user_name",
	FitStatus:             "fit_status",
	FitOverTime:           "fit_over_time",
	WarehouseCode:         "warehouse_code",
	WarehouseName:         "warehouse_name",
	Extra:                 "extra",
	Remark:                "remark",
	CreateTime:            "create_time",
	CreateUser:            "create_user",
	UpdateTime:            "update_time",
	UpdateUser:            "update_user",
	Version:               "version",
	Deleted:               "deleted",
}

// WmsOperationTotalListInfoAffiliated 总单附属表
type WmsOperationTotalListInfoAffiliated struct {
	ID            int       `gorm:"primaryKey;column:id" json:"-"`              // ID
	TotalListNo   string    `gorm:"column:total_list_no" json:"totalListNo"`    // 总单号
	SiteNumber    string    `gorm:"column:site_number" json:"siteNumber"`       // 站点编码
	CountryCode   string    `gorm:"column:country_code" json:"countryCode"`     // 国家二字码
	ScannedCount  int       `gorm:"column:scanned_count" json:"scannedCount"`   // 已扫袋数
	ScannedWeight float64   `gorm:"column:scanned_weight" json:"scannedWeight"` // 已扫重量
	ScannedUser   int       `gorm:"column:scanned_user" json:"scannedUser"`     // 扫描人
	ScannedTime   time.Time `gorm:"column:scanned_time" json:"scannedTime"`     // 扫描时间
	Remarks       string    `gorm:"column:remarks" json:"remarks"`              // 备注信息
	CreateUser    int       `gorm:"column:create_user" json:"createUser"`       // 创建用户
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	UpdateUser    int       `gorm:"column:update_user" json:"updateUser"`       // 更新时间
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
	Version       int       `gorm:"column:version" json:"version"`              // 乐观锁
	IsDelete      int       `gorm:"column:is_delete" json:"isDelete"`           // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsOperationTotalListInfoAffiliated) TableName() string {
	return "wms_operation_total_list_info_affiliated"
}

// WmsOperationTotalListInfoAffiliatedColumns get sql column name.获取数据库列名
var WmsOperationTotalListInfoAffiliatedColumns = struct {
	ID            string
	TotalListNo   string
	SiteNumber    string
	CountryCode   string
	ScannedCount  string
	ScannedWeight string
	ScannedUser   string
	ScannedTime   string
	Remarks       string
	CreateUser    string
	CreateTime    string
	UpdateUser    string
	UpdateTime    string
	Version       string
	IsDelete      string
}{
	ID:            "id",
	TotalListNo:   "total_list_no",
	SiteNumber:    "site_number",
	CountryCode:   "country_code",
	ScannedCount:  "scanned_count",
	ScannedWeight: "scanned_weight",
	ScannedUser:   "scanned_user",
	ScannedTime:   "scanned_time",
	Remarks:       "remarks",
	CreateUser:    "create_user",
	CreateTime:    "create_time",
	UpdateUser:    "update_user",
	UpdateTime:    "update_time",
	Version:       "version",
	IsDelete:      "is_delete",
}

// WmsPackageOperationLog 包裹操作日志
type WmsPackageOperationLog struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`
	PackageID        int       `gorm:"column:package_id" json:"packageId"`               // 包裹id
	OperationContent string    `gorm:"column:operation_content" json:"operationContent"` // 操作内容
	Remark           string    `gorm:"column:remark" json:"remark"`                      // 备注
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`             // 创建用户
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`             // 更新时间
	Version          int       `gorm:"column:version" json:"version"`                    // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                    // 逻辑删除
	IP               string    `gorm:"column:ip" json:"ip"`                              // ip地址
}

// TableName get sql table name.获取数据库表名
func (m *WmsPackageOperationLog) TableName() string {
	return "wms_package_operation_log"
}

// WmsPackageOperationLogColumns get sql column name.获取数据库列名
var WmsPackageOperationLogColumns = struct {
	ID               string
	PackageID        string
	OperationContent string
	Remark           string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
	IP               string
}{
	ID:               "id",
	PackageID:        "package_id",
	OperationContent: "operation_content",
	Remark:           "remark",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
	IP:               "ip",
}

// WmsPackageOperationLogExtend 包裹入库前操作日志
type WmsPackageOperationLogExtend struct {
	ID               int       `gorm:"primaryKey;column:id" json:"-"`
	OrderID          int       `gorm:"column:order_id" json:"orderId"`                   // 订单id
	OperationContent string    `gorm:"column:operation_content" json:"operationContent"` // 操作内容
	Remark           string    `gorm:"column:remark" json:"remark"`                      // 备注
	CreateTime       time.Time `gorm:"column:create_time" json:"createTime"`             // 创建时间
	CreateUser       int       `gorm:"column:create_user" json:"createUser"`             // 创建用户
	UpdateTime       time.Time `gorm:"column:update_time" json:"updateTime"`             // 更新时间
	UpdateUser       int       `gorm:"column:update_user" json:"updateUser"`             // 更新时间
	Version          int       `gorm:"column:version" json:"version"`                    // 乐观锁
	Deleted          int       `gorm:"column:deleted" json:"deleted"`                    // 逻辑删除
	IP               string    `gorm:"column:ip" json:"ip"`                              // ip地址
}

// TableName get sql table name.获取数据库表名
func (m *WmsPackageOperationLogExtend) TableName() string {
	return "wms_package_operation_log_extend"
}

// WmsPackageOperationLogExtendColumns get sql column name.获取数据库列名
var WmsPackageOperationLogExtendColumns = struct {
	ID               string
	OrderID          string
	OperationContent string
	Remark           string
	CreateTime       string
	CreateUser       string
	UpdateTime       string
	UpdateUser       string
	Version          string
	Deleted          string
	IP               string
}{
	ID:               "id",
	OrderID:          "order_id",
	OperationContent: "operation_content",
	Remark:           "remark",
	CreateTime:       "create_time",
	CreateUser:       "create_user",
	UpdateTime:       "update_time",
	UpdateUser:       "update_user",
	Version:          "version",
	Deleted:          "deleted",
	IP:               "ip",
}

// WmsSorterBigBag 分拣机-大包记录
type WmsSorterBigBag struct {
	ID             int       `gorm:"primaryKey;column:id" json:"-"`
	GridNum        int       `gorm:"column:grid_num" json:"gridNum"`               // 格口
	DownstreamCode string    `gorm:"column:downstream_code" json:"downstreamCode"` // 下家code
	SmallBagCount  int       `gorm:"column:small_bag_count" json:"smallBagCount"`  // 小包数量
	Weight         float64   `gorm:"column:weight" json:"weight"`                  // 重量
	BigBagNo       string    `gorm:"column:big_bag_no" json:"bigBagNo"`            // 大包号
	CreateTime     time.Time `gorm:"column:create_time" json:"createTime"`         // 创建时间
	CreateUser     int       `gorm:"column:create_user" json:"createUser"`         // 创建用户
	UpdateTime     time.Time `gorm:"column:update_time" json:"updateTime"`         // 更新时间
	UpdateUser     int       `gorm:"column:update_user" json:"updateUser"`         // 更新时间
	Version        int       `gorm:"column:version" json:"version"`                // 乐观锁
	Deleted        int       `gorm:"column:deleted" json:"deleted"`                // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsSorterBigBag) TableName() string {
	return "wms_sorter_big_bag"
}

// WmsSorterBigBagColumns get sql column name.获取数据库列名
var WmsSorterBigBagColumns = struct {
	ID             string
	GridNum        string
	DownstreamCode string
	SmallBagCount  string
	Weight         string
	BigBagNo       string
	CreateTime     string
	CreateUser     string
	UpdateTime     string
	UpdateUser     string
	Version        string
	Deleted        string
}{
	ID:             "id",
	GridNum:        "grid_num",
	DownstreamCode: "downstream_code",
	SmallBagCount:  "small_bag_count",
	Weight:         "weight",
	BigBagNo:       "big_bag_no",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	Version:        "version",
	Deleted:        "deleted",
}

// WmsSorterGrid 分拣机-格口
type WmsSorterGrid struct {
	ID         int       `gorm:"primaryKey;column:id" json:"-"`
	GridNum    int       `gorm:"column:grid_num" json:"gridNum"`       // 格口
	Remark     string    `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser int       `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser int       `gorm:"column:update_user" json:"updateUser"` // 更新时间
	Version    int       `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted    int       `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsSorterGrid) TableName() string {
	return "wms_sorter_grid"
}

// WmsSorterGridColumns get sql column name.获取数据库列名
var WmsSorterGridColumns = struct {
	ID         string
	GridNum    string
	Remark     string
	CreateTime string
	CreateUser string
	UpdateTime string
	UpdateUser string
	Version    string
	Deleted    string
}{
	ID:         "id",
	GridNum:    "grid_num",
	Remark:     "remark",
	CreateTime: "create_time",
	CreateUser: "create_user",
	UpdateTime: "update_time",
	UpdateUser: "update_user",
	Version:    "version",
	Deleted:    "deleted",
}

// WmsSorterGridRule 分拣机-格口-规则
type WmsSorterGridRule struct {
	ID                   int       `gorm:"primaryKey;column:id" json:"-"`
	RuleName             string    `gorm:"column:rule_name" json:"ruleName"`                          // 规则名称
	GridNums             string    `gorm:"column:grid_nums" json:"gridNums"`                          // 格口，多个以逗号隔开
	SubRuleType          string    `gorm:"column:sub_rule_type" json:"subRuleType"`                   // 子规则类型，weight:重量，country:国家，customer:客户，no_sub_rule:不存在子规则，not_found_order:未找到订单，problem_order:问题订单
	SpecialRuleFlag      []uint8   `gorm:"column:special_rule_flag" json:"specialRuleFlag"`           // 是否特殊规则，0:否，1:是
	CustomerChannelIDs   string    `gorm:"column:customer_channel_ids" json:"customerChannelIds"`     // 客户渠道id，多个以逗号隔开
	CustomerChannelNames string    `gorm:"column:customer_channel_names" json:"customerChannelNames"` // 客户渠道名称，多个以逗号隔开
	DownstreamCode       string    `gorm:"column:downstream_code" json:"downstreamCode"`              // 下家code
	OverFlag             bool      `gorm:"column:over_flag" json:"overFlag"`                          // 是否终结
	Remark               string    `gorm:"column:remark" json:"remark"`                               // 备注
	CreateTime           time.Time `gorm:"column:create_time" json:"createTime"`                      // 创建时间
	CreateUser           int       `gorm:"column:create_user" json:"createUser"`                      // 创建用户
	UpdateTime           time.Time `gorm:"column:update_time" json:"updateTime"`                      // 更新时间
	UpdateUser           int       `gorm:"column:update_user" json:"updateUser"`                      // 更新用户
	Version              int       `gorm:"column:version" json:"version"`                             // 乐观锁
	Deleted              int       `gorm:"column:deleted" json:"deleted"`                             // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsSorterGridRule) TableName() string {
	return "wms_sorter_grid_rule"
}

// WmsSorterGridRuleColumns get sql column name.获取数据库列名
var WmsSorterGridRuleColumns = struct {
	ID                   string
	RuleName             string
	GridNums             string
	SubRuleType          string
	SpecialRuleFlag      string
	CustomerChannelIDs   string
	CustomerChannelNames string
	DownstreamCode       string
	OverFlag             string
	Remark               string
	CreateTime           string
	CreateUser           string
	UpdateTime           string
	UpdateUser           string
	Version              string
	Deleted              string
}{
	ID:                   "id",
	RuleName:             "rule_name",
	GridNums:             "grid_nums",
	SubRuleType:          "sub_rule_type",
	SpecialRuleFlag:      "special_rule_flag",
	CustomerChannelIDs:   "customer_channel_ids",
	CustomerChannelNames: "customer_channel_names",
	DownstreamCode:       "downstream_code",
	OverFlag:             "over_flag",
	Remark:               "remark",
	CreateTime:           "create_time",
	CreateUser:           "create_user",
	UpdateTime:           "update_time",
	UpdateUser:           "update_user",
	Version:              "version",
	Deleted:              "deleted",
}

// WmsSorterGridSorterGridRule 分拣机-格口-规则中间表
type WmsSorterGridSorterGridRule struct {
	ID                int               `gorm:"primaryKey;column:id" json:"-"`
	GridID            int               `gorm:"column:grid_id" json:"gridId"`                                                                 // 格口id
	WmsSorterGrid     WmsSorterGrid     `gorm:"joinForeignKey:grid_id;foreignKey:id;references:GridID" json:"wmsSorterGridList"`              // 分拣机-格口
	GridRuleID        int               `gorm:"column:grid_rule_id" json:"gridRuleId"`                                                        // 格口规则id
	WmsSorterGridRule WmsSorterGridRule `gorm:"joinForeignKey:grid_rule_id;foreignKey:id;references:GridRuleID" json:"wmsSorterGridRuleList"` // 分拣机-格口-规则
}

// TableName get sql table name.获取数据库表名
func (m *WmsSorterGridSorterGridRule) TableName() string {
	return "wms_sorter_grid_sorter_grid_rule"
}

// WmsSorterGridSorterGridRuleColumns get sql column name.获取数据库列名
var WmsSorterGridSorterGridRuleColumns = struct {
	ID         string
	GridID     string
	GridRuleID string
}{
	ID:         "id",
	GridID:     "grid_id",
	GridRuleID: "grid_rule_id",
}

// WmsSorterGridSubRule 分拣机-格口-子规则
type WmsSorterGridSubRule struct {
	ID                int               `gorm:"primaryKey;column:id" json:"-"`
	GridRuleID        int               `gorm:"column:grid_rule_id" json:"gridRuleId"`                                                        // 格口规则id
	WmsSorterGridRule WmsSorterGridRule `gorm:"joinForeignKey:grid_rule_id;foreignKey:id;references:GridRuleID" json:"wmsSorterGridRuleList"` // 分拣机-格口-规则
	GridNums          string            `gorm:"column:grid_nums" json:"gridNums"`                                                             // 格口，多个以逗号隔开
	DownstreamCode    string            `gorm:"column:downstream_code" json:"downstreamCode"`                                                 // 下家code
	RuleType          string            `gorm:"column:rule_type" json:"ruleType"`                                                             // 规则类型，weight:重量，country:国家，customer:客户，no_sub_rule:不存在子规则，not_found_order:未找到订单，problem_order:问题订单
	RuleValue         string            `gorm:"column:rule_value" json:"ruleValue"`
	OverFlag          bool              `gorm:"column:over_flag" json:"overFlag"`     // 是否终结
	Remark            string            `gorm:"column:remark" json:"remark"`          // 备注
	CreateTime        time.Time         `gorm:"column:create_time" json:"createTime"` // 创建时间
	CreateUser        int               `gorm:"column:create_user" json:"createUser"` // 创建用户
	UpdateTime        time.Time         `gorm:"column:update_time" json:"updateTime"` // 更新时间
	UpdateUser        int               `gorm:"column:update_user" json:"updateUser"` // 更新用户
	Version           int               `gorm:"column:version" json:"version"`        // 乐观锁
	Deleted           int               `gorm:"column:deleted" json:"deleted"`        // 逻辑删除
	UseState          int               `gorm:"column:use_state" json:"useState"`     // 0使用中 1未使用
}

// TableName get sql table name.获取数据库表名
func (m *WmsSorterGridSubRule) TableName() string {
	return "wms_sorter_grid_sub_rule"
}

// WmsSorterGridSubRuleColumns get sql column name.获取数据库列名
var WmsSorterGridSubRuleColumns = struct {
	ID             string
	GridRuleID     string
	GridNums       string
	DownstreamCode string
	RuleType       string
	RuleValue      string
	OverFlag       string
	Remark         string
	CreateTime     string
	CreateUser     string
	UpdateTime     string
	UpdateUser     string
	Version        string
	Deleted        string
	UseState       string
}{
	ID:             "id",
	GridRuleID:     "grid_rule_id",
	GridNums:       "grid_nums",
	DownstreamCode: "downstream_code",
	RuleType:       "rule_type",
	RuleValue:      "rule_value",
	OverFlag:       "over_flag",
	Remark:         "remark",
	CreateTime:     "create_time",
	CreateUser:     "create_user",
	UpdateTime:     "update_time",
	UpdateUser:     "update_user",
	Version:        "version",
	Deleted:        "deleted",
	UseState:       "use_state",
}

// WmsSorterSmallBag 分拣机-小包记录
type WmsSorterSmallBag struct {
	ID                int       `gorm:"primaryKey;column:id" json:"-"`
	GridNum           int       `gorm:"column:grid_num" json:"gridNum"`                      // 格口
	OrderNumber       string    `gorm:"column:order_number" json:"orderNumber"`              // 订单号
	ReferenceNumber   string    `gorm:"column:reference_number" json:"referenceNumber"`      // 参考号
	LogisticsNumber   string    `gorm:"column:logistics_number" json:"logisticsNumber"`      // 物流单号
	Weight            float64   `gorm:"column:weight" json:"weight"`                         // 重量
	SorterBigBagNo    string    `gorm:"column:sorter_big_bag_no" json:"sorterBigBagNo"`      // 分拣机大包号
	DownstreamCode    string    `gorm:"column:downstream_code" json:"downstreamCode"`        // 下家code
	CustomerChannelID int       `gorm:"column:customer_channel_id" json:"customerChannelId"` // 客户渠道id
	CreateTime        time.Time `gorm:"column:create_time" json:"createTime"`                // 创建时间
	CreateUser        int       `gorm:"column:create_user" json:"createUser"`                // 创建用户
	UpdateTime        time.Time `gorm:"column:update_time" json:"updateTime"`                // 更新时间
	UpdateUser        int       `gorm:"column:update_user" json:"updateUser"`                // 更新时间
	Version           int       `gorm:"column:version" json:"version"`                       // 乐观锁
	Deleted           int       `gorm:"column:deleted" json:"deleted"`                       // 逻辑删除
}

// TableName get sql table name.获取数据库表名
func (m *WmsSorterSmallBag) TableName() string {
	return "wms_sorter_small_bag"
}

// WmsSorterSmallBagColumns get sql column name.获取数据库列名
var WmsSorterSmallBagColumns = struct {
	ID                string
	GridNum           string
	OrderNumber       string
	ReferenceNumber   string
	LogisticsNumber   string
	Weight            string
	SorterBigBagNo    string
	DownstreamCode    string
	CustomerChannelID string
	CreateTime        string
	CreateUser        string
	UpdateTime        string
	UpdateUser        string
	Version           string
	Deleted           string
}{
	ID:                "id",
	GridNum:           "grid_num",
	OrderNumber:       "order_number",
	ReferenceNumber:   "reference_number",
	LogisticsNumber:   "logistics_number",
	Weight:            "weight",
	SorterBigBagNo:    "sorter_big_bag_no",
	DownstreamCode:    "downstream_code",
	CustomerChannelID: "customer_channel_id",
	CreateTime:        "create_time",
	CreateUser:        "create_user",
	UpdateTime:        "update_time",
	UpdateUser:        "update_user",
	Version:           "version",
	Deleted:           "deleted",
}

// WmsWarehouse 仓库表
type WmsWarehouse struct {
	ID            uint      `gorm:"primaryKey;column:id" json:"-"`              // 主键id
	WarehouseCode string    `gorm:"column:warehouse_code" json:"warehouseCode"` // 仓库code
	WarehouseName string    `gorm:"column:warehouse_name" json:"warehouseName"` // 仓库名
	WarehouseType int       `gorm:"column:warehouse_type" json:"warehouseType"` // 仓库类型(0客户仓库,1安骏仓库)
	CustomerID    int       `gorm:"column:customer_id" json:"customerId"`       // 客户id(客户仓库必填)
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`       // 创建时间
	CreateUser    uint      `gorm:"column:create_user" json:"createUser"`       // 创建用户
	UpdateTime    time.Time `gorm:"column:update_time" json:"updateTime"`       // 更新时间
	UpdateUser    uint      `gorm:"column:update_user" json:"updateUser"`       // 更新用户
	Version       uint      `gorm:"column:version" json:"version"`              // 乐观锁
	Deleted       uint32    `gorm:"column:deleted" json:"deleted"`              // 逻辑删除
	Remark        string    `gorm:"column:remark" json:"remark"`                // 备注
	Status        uint32    `gorm:"column:status" json:"status"`                // 使用状态是否禁用0代表否1代表是
}

// TableName get sql table name.获取数据库表名
func (m *WmsWarehouse) TableName() string {
	return "wms_warehouse"
}

// WmsWarehouseColumns get sql column name.获取数据库列名
var WmsWarehouseColumns = struct {
	ID            string
	WarehouseCode string
	WarehouseName string
	WarehouseType string
	CustomerID    string
	CreateTime    string
	CreateUser    string
	UpdateTime    string
	UpdateUser    string
	Version       string
	Deleted       string
	Remark        string
	Status        string
}{
	ID:            "id",
	WarehouseCode: "warehouse_code",
	WarehouseName: "warehouse_name",
	WarehouseType: "warehouse_type",
	CustomerID:    "customer_id",
	CreateTime:    "create_time",
	CreateUser:    "create_user",
	UpdateTime:    "update_time",
	UpdateUser:    "update_user",
	Version:       "version",
	Deleted:       "deleted",
	Remark:        "remark",
	Status:        "status",
}
