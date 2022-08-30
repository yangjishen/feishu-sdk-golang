package sdk

import (
	"github.com/yangjishen/feishu-sdk-golang/core/consts"
	"github.com/yangjishen/feishu-sdk-golang/core/model/vo"
	"github.com/yangjishen/feishu-sdk-golang/core/util/http"
	"github.com/yangjishen/feishu-sdk-golang/core/util/json"
	"github.com/yangjishen/feishu-sdk-golang/core/util/log"
)

//1.查询租户授权状态 https://bytedance.feishu.cn/docs/doccnHJx2UbLZh5kiWjNawICyNd#dCNL6V
func (t Tenant) GetScopes() (*vo.GetScopesResp, error) {
	respBody, err := http.Get(consts.ApiGetScopes, nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetScopesResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 申请授权 https://bytedance.feishu.cn/docs/doccnHJx2UbLZh5kiWjNawICyNd#kHHiAa
func (t Tenant) ApplyScopes() (*vo.ApplyScopesResp, error) {
	respBody, err := http.Post(consts.ApiApplyScopes, nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.ApplyScopesResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 获取通讯录授权范围 https://open.feishu.cn/open-apis/contact/v3/scopes
func (t Tenant) GetContactScopes(userIdType string, departmentIdType string, pageSize int, pageToken string) (*vo.GetContactScopesResp, error) {
	queryParams := map[string]interface{}{}
	if userIdType != "" {
		queryParams["user_id_type"] = userIdType
	}
	if departmentIdType != "" {
		queryParams["department_id_type"] = departmentIdType
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}

	respBody, err := http.Get(consts.ApiContactScopes, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetContactScopesResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
