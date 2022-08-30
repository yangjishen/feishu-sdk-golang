package vo

type GetScopesResp struct {
	CommonVo
	Data GetScopesRespData `json:"data"`
}

type GetScopesRespData struct {
	Scopes []GetScopesRespDataScopes `json:"scopes"`
}

type GetScopesRespDataScopes struct {
	ScopeName   string `json:"scope_name"`
	GrantStatus int    `json:"grant_status"`
}

type ApplyScopesResp struct {
	CommonVo
	Data ApplyScopesRespData `json:"data"`
}

type ApplyScopesRespData struct {
}

type GetContactScopesResp struct {
	CommonVo
	Data *GetContactScopesRespData `json:"data"`
}

type GetContactScopesRespData struct {
	DepartmentIds []string `json:"department_ids"`
	UserIds       []string `json:"user_ids"`
	GroupIds      []string `json:"group_ids"`
	HasMore       bool     `json:"has_more"`
	PageToken     string   `json:"page_token"`
}
