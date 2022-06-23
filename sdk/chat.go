package sdk

import (
	"fmt"
	"github.com/yangjishen/feishu-sdk-golang/core/consts"
	"github.com/yangjishen/feishu-sdk-golang/core/model/vo"
	"github.com/yangjishen/feishu-sdk-golang/core/util/http"
	"github.com/yangjishen/feishu-sdk-golang/core/util/json"
	"github.com/yangjishen/feishu-sdk-golang/core/util/log"
	"net/url"
)

//获取用户所在的群列表 https://open.feishu.cn/document/ukTMukTMukTM/uQzMwUjL0MDM14CNzATN
func (t Tenant) GroupList(userAccessToken string, pageSize int, pageToken string) (*vo.GroupListRespVo, error) {
	queryParams := map[string]interface{}{}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiUserGroupLIst, queryParams, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GroupListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取群成员列表 https://open.feishu.cn/document/ukTMukTMukTM/uUzMwUjL1MDM14SNzATN
func (t Tenant) ChatMembers(userAccessToken string, chatId string, pageSize int, pageToken string) (*vo.ChatMembersRespVo, error) {
	queryParams := map[string]interface{}{
		"chat_id": chatId,
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiChatMembers, queryParams, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.ChatMembersRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//搜索用户所在的群列表 https://open.feishu.cn/document/ukTMukTMukTM/uUTOyUjL1kjM14SN5ITN
func (t Tenant) ChatSearch(userAccessToken string, query string, pageSize int, pageToken string) (*vo.GroupListRespVo, error) {
	queryParams := map[string]interface{}{
		"query": url.QueryEscape(query),
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiChatSearch, queryParams, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GroupListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取用户或机器人所在的群列表 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/list
func (t Tenant) ImChatList(userAccessToken string, pageSize int, pageToken string, userIdType *string) (*vo.ImChatListRespVo, error) {
	queryParams := map[string]interface{}{}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	if userIdType != nil {
		queryParams["user_id_type"] = *userIdType
	}

	accessToken := t.TenantAccessToken
	if userAccessToken != "" {
		accessToken = userAccessToken
	}
	respBody, err := http.Get(consts.ApiImChatList, queryParams, http.BuildTokenHeaderOptions(accessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	respVo := &vo.ImChatListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取群信息 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat/get
func (t Tenant) ImChatInfo(userAccessToken string, chatId string) (*vo.ImChatInfoRespVo, error) {
	accessToken := t.TenantAccessToken
	if userAccessToken != "" {
		accessToken = userAccessToken
	}
	respBody, err := http.Get(fmt.Sprintf("%s%s", consts.ApiImChatInfo, chatId), nil, http.BuildTokenHeaderOptions(accessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	respVo := &vo.ImChatInfoRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
