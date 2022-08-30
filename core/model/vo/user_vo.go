package vo

type GetDepartmentUserListRespVo struct {
	CommonVo
	Data *GetDepartmentUserListRespVoData `json:"data"`
}

type GetDepartmentUserListV2RespVo struct {
	CommonVo
	Data *GetDepartmentUserListV2RespVoData `json:"data"`
}

type GetDepartmentUserListV3RespVo struct {
	CommonVo
	Data *GetDepartmentUserListV3RespVoData `json:"data"`
}

type GetDepartmentUserDetailListRespVo struct {
	CommonVo
	Data *GetDepartmentUserDetailListRespVoData `json:"data"`
}

type GetDepartmentUserDetailListRespVoData struct {
	HasMore   bool             `json:"has_more"`
	UserInfos []UserDetailInfo `json:"user_infos"`
}

type GetDepartmentUserListRespVoData struct {
	HasMore  bool             `json:"has_more"`
	UserList []UserRestInfoVo `json:"user_list"`
}

type GetDepartmentUserListV2RespVoData struct {
	HasMore   bool             `json:"has_more"`
	PageToken string           `json:"page_token"`
	UserList  []UserRestInfoVo `json:"users"`
}

type GetDepartmentUserListV3RespVoData struct {
	HasMore   bool               `json:"has_more"`
	PageToken string             `json:"page_token"`
	Items     []UserRestV3InfoVo `json:"items"`
}

type UserRestV3InfoVo struct {
	Avatar          UserRestV3Avatar `json:"avatar"`
	City            string           `json:"city"`
	Country         string           `json:"country"`
	DepartmentIds   []string         `json:"department_ids"`
	Description     string           `json:"description"`
	Email           string           `json:"email"`
	EmployeeNo      string           `json:"employee_no"`
	EmployeeType    int              `json:"employee_type"`
	EnName          string           `json:"en_name"`
	Gender          int              `json:"gender"`
	IsTenantManager bool             `json:"is_tenant_manager"`
	JobTitle        string           `json:"job_title"`
	UnionId         string           `json:"union_id"`
	OpenId          string           `json:"open_id"`
	UserId          string           `json:"user_id"`
	Name            string           `json:"name"`
	Mobile          string           `json:"mobile"`
	MobileVisible   bool             `json:"mobile_visible"`
	Status          UserStatus       `json:"status"`
	LeaderUserId    string           `json:"leader_user_id"`
	WorkStation     string           `json:"work_station"`
	JoinTime        int              `json:"join_time"`
	Orders          []Order          `json:"orders"`
	//CustomAttrs     string           `json:"custom_attrs"`
	EnterpriseEmail string `json:"enterprise_email"`
	IsFrozen        bool   `json:"is_frozen"`
}

type UserRestV3Avatar struct {
	Avatar240    string `json:"avatar_240"`
	Avatar640    string `json:"avatar_640"`
	Avatar72     string `json:"avatar_72"`
	AvatarOrigin string `json:"avatar_origin"`
	UserId       string `json:"user_id"`
	Name         string `json:"name"`
	EmployeeNo   string `json:"employee_no"`
}

type UserRestInfoVo struct {
	EmployeeId string `json:"employee_id"`
	OpenId     string `json:"open_id"`
	UserId     string `json:"user_id"`
	Name       string `json:"name"`
	EmployeeNo string `json:"employee_no"`
}

type GetDepartmentSimpleListRespVo struct {
	CommonVo
	Data *GetDepartmentSimpleListRespVoData `json:"data"`
}

type GetDepartmentSimpleListV2RespVo struct {
	CommonVo
	Data *GetDepartmentSimpleListV2RespVoData `json:"data"`
}

type GetDepartmentSimpleListV3RespVo struct {
	CommonVo
	Data *GetDepartmentSimpleListV3RespVoData `json:"data"`
}

type GetDepartmentInfoV3RespVo struct {
	CommonVo
	Data *GetDepartmentInfoV3RespVoData `json:"data"`
}

type GetDepartmentInfoV3RespVoData struct {
	Department *DepartmentRestV3InfoVo `json:"department"`
}

type GetScopeRespVo struct {
	CommonVo
	Data *GetScopeRespData `json:"data"`
}

type GetScopeRespV2Vo struct {
	CommonVo
	Data *GetScopeRespV2Data `json:"data"`
}

type GetScopeRespData struct {
	AuthedDepartments []string `json:"authed_departments"`
	AuthedEmployeeIds []string `json:"authed_employee_ids"`
	AuthedOpenIds     []string `json:"authed_open_ids"`
}

type GetScopeRespV2Data struct {
	AuthedDepartments []string    `json:"authed_departments"`
	AuthedUsers       []ScopeUser `json:"authed_users"`
}

type ScopeUser struct {
	OpenId string `json:"open_id"`
	UserId string `json:"user_id"`
}

type GetDepartmentSimpleListRespVoData struct {
	HasMore         bool                   `json:"has_more"`
	PageToken       string                 `json:"page_token"`
	DepartmentInfos []DepartmentRestInfoVo `json:"department_infos"`
}

type GetDepartmentSimpleListV2RespVoData struct {
	HasMore         bool                   `json:"has_more"`
	PageToken       string                 `json:"page_token"`
	DepartmentInfos []DepartmentRestInfoVo `json:"departments"`
}

type GetDepartmentSimpleListV3RespVoData struct {
	HasMore   bool                     `json:"has_more"`
	PageToken string                   `json:"page_token"`
	Items     []DepartmentRestV3InfoVo `json:"items"`
}

type DepartmentRestV3InfoVo struct {
	Name               string             `json:"name"`
	ParentDepartmentId string             `json:"parent_department_id"`
	DepartmentId       string             `json:"department_id"`
	OpenDepartmentId   string             `json:"open_department_id"`
	LeaderUserId       string             `json:"leader_user_id"`
	ChatId             string             `json:"chat_id"`
	Order              string             `json:"order"`
	UnitIds            []string           `json:"unit_ids"`
	MemberCount        int                `json:"member_count"`
	Status             DepartmentV3Status `json:"status"`
	CreateGroupChat    bool               `json:"create_group_chat"`
}

type DepartmentV3Status struct {
	IsDeleted bool `json:"is_deleted"`
}

type DepartmentRestInfoVo struct {
	Id                     string `json:"id"`
	Name                   string `json:"name"`
	ParentId               string `json:"parent_id"`
	OpenDepartmentID       string `json:"open_department_id"`
	ParentOpenDepartmentID string `json:"parent_open_department_id"`
}

type GetDepartmentInfoRespVo struct {
	CommonVo
	Data *GetDepartmentInfoRespVoData `json:"data"`
}

type GetDepartmentInfoBatchRespVo struct {
	CommonVo
	Data GetDepartmentInfoBatchRespData `json:"data"`
}

type GetDepartmentInfoBatchRespData struct {
	Departments []DepartmentDetailInfo `json:"departments"`
	Errors      []GetUserBatchError    `json:"errors"`
}

type DepartmentDetailInfo struct {
	ChatId      string `json:"chat_id"`
	HasChild    bool   `json:"has_child"`
	Id          string `json:"id"`
	Leader      Leader `json:"leader"`
	MemberCount int    `json:"member_count"`
	Name        string `json:"name"`
	ParentId    string `json:"parent_id"`
	Status      int    `json:"status"`
}

type GetDepartmentInfoRespVoData struct {
	DepartmentInfo *DepartmentDetailInfoVo `json:"department_info"`
}

type DepartmentDetailInfoVo struct {
	Id               string `json:"id"`
	LeaderEmployeeId string `json:"leader_employee_id"`
	LeaderOpenId     string `json:"leader_open_id"`
	ChatId           string `json:"chat_id"`
	MemberCount      int    `json:"member_count"`
	Name             string `json:"name"`
	ParentId         string `json:"parent_id"`
	Status           int    `json:"status"`
}

type GetUserBatchGetRespVo struct {
	CommonVo
	Data *GetUserBatchGetRespVoData `json:"data"`
}

type GetUserBatchGetRespVoData struct {
	UserInfos []UserDetailInfo `json:"user_infos"`
}

type UserDetailInfo struct {
	Name             string           `json:"name"`
	NamePy           string           `json:"name_py"`
	EnName           string           `json:"en_name"`
	EmployeeId       string           `json:"employee_id"`
	EmployeeNo       string           `json:"employee_no"`
	OpenId           string           `json:"open_id"`
	Status           int              `json:"status"`
	EmployeeType     int              `json:"employee_type"`
	Avatar71         string           `json:"avatar_71"`
	Avatar240        string           `json:"avatar_240"`
	Avatar640        string           `json:"avatar_640"`
	AvatarUrl        string           `json:"avatar_url"`
	Gender           int              `json:"gender"`
	Email            string           `json:"email"`
	Mobile           string           `json:"mobile"`
	Description      string           `json:"description"`
	Country          string           `json:"country"`
	City             string           `json:"city"`
	WorkStation      string           `json:"work_station"`
	IsTenantManager  bool             `json:"is_tenant_manager"`
	JoinTime         int64            `json:"join_time"`
	UpdateTime       int64            `json:"update_time"`
	LeaderEmployeeId string           `json:"leader_employee_id"`
	LeaderOpenId     string           `json:"leader_open_id"`
	Departments      []string         `json:"departments"`
	CustomAttr       map[string]Entry `json:"custom_attr"`
}

type UserDetailInfoV2 struct {
	Name            string       `json:"name"`
	NamePy          string       `json:"name_py"`
	EnName          string       `json:"en_name"`
	EmployeeId      string       `json:"employee_id"`
	EmployeeNo      string       `json:"employee_no"`
	OpenId          string       `json:"open_id"`
	UserId          string       `json:"user_id"`
	UnionId         string       `json:"union_id"`
	Status          UserStatus   `json:"status"`
	EmployeeType    int          `json:"employee_type"`
	Avatar          UserAvatar   `json:"avatar"`
	Gender          int          `json:"gender"`
	Email           string       `json:"email"`
	Mobile          string       `json:"mobile"`
	Description     string       `json:"description"`
	Country         string       `json:"country"`
	City            string       `json:"city"`
	WorkStation     string       `json:"work_station"`
	IsTenantManager bool         `json:"is_tenant_manager"`
	JoinTime        int64        `json:"join_time"`
	UpdateTime      int64        `json:"update_time"`
	Leader          Leader       `json:"leader"`
	Departments     []string     `json:"departments"`
	Positions       []Position   `json:"positions"`
	Orders          []Order      `json:"orders"`
	CustomAttrs     []CustomAttr `json:"custom_attrs"`
}

type UserDetailInfoV3 struct {
	Name            string       `json:"name"`
	NamePy          string       `json:"name_py"`
	EnName          string       `json:"en_name"`
	NickName        string       `json:"nickname"`
	MobileVisible   bool         `json:"mobile_visible"`
	EmployeeId      string       `json:"employee_id"`
	EmployeeNo      string       `json:"employee_no"`
	OpenId          string       `json:"open_id"`
	UserId          string       `json:"user_id"`
	UnionId         string       `json:"union_id"`
	Status          UserStatus   `json:"status"`
	EmployeeType    int          `json:"employee_type"`
	Avatar          UserAvatar   `json:"avatar"`
	Gender          int          `json:"gender"`
	Email           string       `json:"email"`
	Mobile          string       `json:"mobile"`
	Description     string       `json:"description"`
	Country         string       `json:"country"`
	City            string       `json:"city"`
	WorkStation     string       `json:"work_station"`
	IsTenantManager bool         `json:"is_tenant_manager"`
	JoinTime        int64        `json:"join_time"`
	UpdateTime      int64        `json:"update_time"`
	LeaderUserId    string       `json:"leader_user_id"`
	DepartmentIds   []string     `json:"department_ids"`
	Positions       []Position   `json:"positions"`
	Orders          []Order      `json:"orders"`
	CustomAttrs     []CustomAttr `json:"custom_attrs"`
	EnterpriseEmail string       `json:"enterprise_email"`
	JobTitle        string       `json:"job_title"`
}

type CustomAttr struct {
	Id    string          `json:"id"`
	Type  string          `json:"type"`
	Value CustomAttrValue `json:"value"`
}

type GenericUserValue struct {
	Id   string `json:"id"`
	Type int    `json:"type"`
}

type CustomAttrValue struct {
	Url         string           `json:"url"`
	PcUrl       string           `json:"pc_url"`
	Text        string           `json:"text"`
	OptionValue string           `json:"option_value"`
	Name        string           `json:"name"`
	PictureUrl  string           `json:"picture_url"`
	GenericUser GenericUserValue `json:"generic_user"`
}

type Order struct {
	DepartmentId    string `json:"department_id"`
	UserOrder       int    `json:"user_order"`
	DepartmentOrder int    `json:"department_order"`
}

type Position struct {
	PositionCode string `json:"position_code"`
	PositionName string `json:"position_name"`
	DepartmentId string `json:"department_id"`
	Leader       Leader `json:"leader"`
	IsMajor      bool   `json:"is_major"`
}

type Leader struct {
	OpenId       string `json:"open_id"`
	UserId       string `json:"user_id"`
	PositionCode string `json:"position_code"`
}

type UserStatus struct {
	IsFrozen    bool `json:"is_frozen"`
	IsResigned  bool `json:"is_resigned"`
	IsActivated bool `json:"is_activated"`
	IsExited    bool `json:"is_exited"`
	IsUnjoin    bool `json:"is_unjoin"`
}

type Entry struct {
	Value string `json:"value"`
}

type SearchUserResp struct {
	CommonVo
	Data *SearchUserRespData `json:"data"`
}

type SearchUserRespData struct {
	HasMore   bool             `json:"has_more"`
	PageToken string           `json:"page_token"`
	Users     []SearchUserInfo `json:"users"`
}

type SearchUserInfo struct {
	Avatar        UserAvatar `json:"avatar"`
	DepartmentIds []string   `json:"department_ids"`
	Name          string     `json:"name"`
	OpenId        string     `json:"open_id"`
	UserId        string     `json:"user_id"`
}

type UserAvatar struct {
	Avatar72     string `json:"avatar_72"`
	Avatar240    string `json:"avatar_240"`
	Avatar640    string `json:"avatar_640"`
	AvatarOrigin string `json:"avatar_origin"`
}

type GetUserBatchGetV2Resp struct {
	CommonVo
	Data GetUserBatchGetV2RespData `json:"data"`
}

type GetUserBatchGetV2RespData struct {
	HasMore   bool                `json:"has_more"`
	PageToken string              `json:"page_token"`
	Users     []UserDetailInfoV2  `json:"users"`
	Errors    []GetUserBatchError `json:"errors"`
}

type GetUserBatchError struct {
	Id   string `json:"id"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type GetUserDetailV3RespData struct {
	User UserDetailInfoV3 `json:"user"`
}

type GetUserDetailV3Resp struct {
	CommonVo
	Data GetUserDetailV3RespData `json:"data"`
}

type GetUsersV3Resp struct {
	CommonVo
	Data GetUsersV3RespData `json:"data"`
}

type GetUsersV3RespData struct {
	HasMore   bool               `json:"has_more"`
	PageToken string             `json:"page_token"`
	Items     []UserDetailInfoV3 `json:"items"`
}

type BatchGetIdResp struct {
	CommonVo
	Data BatchGetIdData `json:"data"`
}

type BatchGetIdData struct {
	EmailUsers      map[string][]SimpleIdInfo `json:"email_users"`
	EmailsNotExist  []string                  `json:"emails_not_exist"`
	MobileUsers     map[string][]SimpleIdInfo `json:"mobile_users"`
	MobilesNotExist []string                  `json:"mobiles_not_exist"`
}

type SimpleIdInfo struct {
	OpenId string `json:"open_Id"`
	UserId string `json:"user_id"`
}

type AuthenAccessTokenResp struct {
	CommonVo
	Data AuthenAccessTokenData `json:"data"`
}

type AuthenAccessTokenData struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int64  `json:"expires_in"`
	Name             string `json:"name"`
	EnName           string `json:"en_name"`
	AvatarUrl        string `json:"avatar_url"`
	AvatarThumb      string `json:"avatar_thumb"`
	AvatarMiddle     string `json:"avatar_middle"`
	AvatarBig        string `json:"avatar_big"`
	OpenId           string `json:"open_id"`
	UnionId          string `json:"union_id"`
	Email            string `json:"email"`
	UserId           string `json:"user_id"`
	Mobile           string `json:"mobile"`
	TenantKey        string `json:"tenant_key"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
}
