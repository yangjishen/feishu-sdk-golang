package sdk

import (
	"fmt"
	"github.com/yangjishen/feishu-sdk-golang/core/consts"
	"github.com/yangjishen/feishu-sdk-golang/core/model/vo"
	"github.com/yangjishen/feishu-sdk-golang/core/util/http"
	"github.com/yangjishen/feishu-sdk-golang/core/util/json"
	"github.com/yangjishen/feishu-sdk-golang/core/util/log"
)

//CreateCalendarV4 创建日历v4版本 https://open.feishu.cn/document/ukTMukTMukTM/uUDM3YjL1AzN24SNwcjN
func (t Tenant) CreateCalendarV4(bodyParams vo.CreateCalendarV4Req) (*vo.CommonCalendarV4Resp, error) {
	respBody, err := http.Post(consts.ApiCalendarCreateV4, nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//创建日程v4 https://open.feishu.cn/document/ukTMukTMukTM/uYTM3YjL2EzN24iNxcjN
func (t Tenant) CreateCalendarEventV4(calendarId string, bodyParams vo.CreateCalendarEventV4Req) (*vo.CalendarEventInfoV4Resp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarEventCreateV4, calendarId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CalendarEventInfoV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//删除日程v4 https://open.feishu.cn/document/ukTMukTMukTM/ucTM3YjL3EzN24yNxcjN
func (t Tenant) DeleteCalendarEventV4(calendarId string, eventId string) (*vo.CommonVo, error) {
	respBody, err := http.Delete(fmt.Sprintf(consts.ApiCalendarEventDeleteV4, calendarId, eventId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取日程v4 https://open.feishu.cn/document/ukTMukTMukTM/ugTM3YjL4EzN24COxcjN
func (t Tenant) GetCalendarEventV4(calendarId string, eventId string) (*vo.CalendarEventInfoV4Resp, error) {
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventGetV4, calendarId, eventId), nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CalendarEventInfoV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取日程列表v4 https://open.feishu.cn/document/ukTMukTMukTM/ukTM3YjL5EzN24SOxcjN
func (t Tenant) GetCalendarEventListV4(calendarId string, eventId string, pageSize int, pageToken *string, syncToken *string) (*vo.GetCalendarEventListV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_size":  pageSize,
		"page_token": pageToken,
		"sync_token": syncToken,
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventBatchGetV4, calendarId), queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarEventListV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//更新日程v4 https://open.feishu.cn/document/ukTMukTMukTM/uAjM3YjLwIzN24CMycjN
func (t Tenant) UpdateCalendarEventV4(calendarId string, eventId string, bodyParams vo.UpdateCalendarEventV4Req) (*vo.CalendarEventInfoV4Resp, error) {
	respBody, err := http.Patch(fmt.Sprintf(consts.ApiCalendarEventUpdateV4, calendarId, eventId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CalendarEventInfoV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//创建日程参与人v4 https://open.feishu.cn/document/ukTMukTMukTM/uAjM3YjLwIzN24CMycjN
func (t Tenant) AddCalendarEventAttendeesV4(calendarId string, eventId string, userIdType string, bodyParams vo.AddCalendarEventAttendeesV4Req) (*vo.AddCalendarEventAttendeesV4Resp, error) {
	params := map[string]interface{}{
		"user_id_type": userIdType,
	}
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarEventAttendeesAddV4, calendarId, eventId), params, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AddCalendarEventAttendeesV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//删除日程参与人v4 https://open.feishu.cn/document/ukTMukTMukTM/uAzM3YjLwMzN24CMzcjN
func (t Tenant) DeleteCalendarEventAttendeesV4(calendarId string, eventId string, userIdType string, bodyParams vo.DeleteCalendarEventAttendeesV4Req) (*vo.CommonVo, error) {
	params := map[string]interface{}{
		"user_id_type": userIdType,
	}
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarEventAttendeesDeleteV4, calendarId, eventId), params, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取日程参与人列表v4 https://open.feishu.cn/document/ukTMukTMukTM/uEzM3YjLxMzN24SMzcjN
func (t Tenant) GetCalendarEventAttendeesV4(calendarId string, eventId string, pageSize int, pageToken *string, userIdType string) (*vo.GetCalendarEventAttendeesV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_size":    pageSize,
		"page_token":   pageToken,
		"user_id_type": userIdType,
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventAttendeesGetV4, calendarId, eventId), queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarEventAttendeesV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取参与人群成员列表v4 https://open.feishu.cn/document/ukTMukTMukTM/uATN3YjLwUzN24CM1cjN
func (t Tenant) GetCalendarEventAttendeesChatMembersV4(calendarId string, eventId string, attendeeId string, pageSize int, pageToken *string, userIdType string) (*vo.GetCalendarEventAttendeesV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_size":    pageSize,
		"page_token":   pageToken,
		"user_id_type": userIdType,
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventAttendeesChatMembersGetV4, calendarId, eventId, attendeeId), queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarEventAttendeesV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//DeleteCalendarV4 删除日历 https://open.feishu.cn/document/ukTMukTMukTM/uYDM3YjL2AzN24iNwcjN
func (t Tenant) DeleteCalendarV4(calendarId string) (*vo.VoidV4Resp, error) {
	respBody, err := http.Delete(fmt.Sprintf(consts.ApiCalendarDeleteV4, calendarId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.VoidV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//GetCalendarV4 获取日历 https://open.feishu.cn/document/ukTMukTMukTM/ucDM3YjL3AzN24yNwcjN
func (t Tenant) GetCalendarV4(calendarId string) (*vo.GetCalendarV4Resp, error) {
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarGetV4, calendarId), nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 获取日历列表 https://open.feishu.cn/document/ukTMukTMukTM/ugDM3YjL4AzN24COwcjN  todo
func (t Tenant) GetCalendarListV4(pageSize int, pageToken, syncToken string) (*vo.GetCalendarListV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_size":  pageSize,
		"page_token": pageToken,
		"sync_token": syncToken,
	}
	respBody, err := http.Post(consts.ApiCalendarGetListV4, queryParams, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarListV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//UpdateCalendarV4 更新日历 https://open.feishu.cn/document/ukTMukTMukTM/ukDM3YjL5AzN24SOwcjN
func (t Tenant) UpdateCalendarV4(calendarId string, bodyParams vo.UpdateCalendarV4Req) (*vo.CommonCalendarV4Resp, error) {
	respBody, err := http.Patch(fmt.Sprintf(consts.ApiCalendarUpdateV4, calendarId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//SearchCalendarV4 搜索日历 https://open.feishu.cn/document/ukTMukTMukTM/uATM3YjLwEzN24CMxcjN
func (t Tenant) SearchCalendarV4(bodyParams vo.SearchCalendarV4Req, pageToken string, pageSize int) (*vo.SearchCalendarV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_token": pageToken,
		"page_size":  pageSize,
	}
	respBody, err := http.Post(consts.ApiCalendarSearchV4, queryParams, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.SearchCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//UnsubscribeCalendarV4 取消订阅日历 https://open.feishu.cn/document/ukTMukTMukTM/ugDO3YjL4gzN24CO4cjN
func (t Tenant) UnsubscribeCalendarV4(calendarId string) (*vo.VoidV4Resp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarUnsubscribeV4, calendarId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.VoidV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//SubscribeCalendarV4 订阅日历 https://open.feishu.cn/document/ukTMukTMukTM/ucDO3YjL3gzN24yN4cjN
func (t Tenant) SubscribeCalendarV4(calendarId string) (*vo.SubscribeCalendarV4Resp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarSubscribeV4, calendarId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.SubscribeCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//SubscriptionCalendarV4 订阅日历变更事件 https://open.feishu.cn/document/ukTMukTMukTM/ugTO2YjL4kjN24CO5YjN/subscribe%20calendar%20changed%20event
func (t Tenant) SubscriptionCalendarV4() (*vo.VoidV4Resp, error) {
	respBody, err := http.Post(consts.ApiCalendarSubscriptionV4, nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.VoidV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//创建访问控制 https://open.feishu.cn/document/ukTMukTMukTM/uUjM3YjL1IzN24SNycjN
func (t Tenant) AddCalendarAclV4(calendarId string, userIdType string, body vo.AddCalendarAclV4Req) (*vo.AddCalendarAclV4Resp, error) {
	params := map[string]interface{}{
		"user_id_type": userIdType,
	}
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarAddCalendarAclV4, calendarId), params, json.ToJsonIgnoreError(body), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AddCalendarAclV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//删除访问控制 https://open.feishu.cn/document/ukTMukTMukTM/uYjM3YjL2IzN24iNycjN
func (t Tenant) DeleteCalendarAclV4(calendarId string, aclId string) (*vo.CommonVo, error) {
	respBody, err := http.Delete(fmt.Sprintf(consts.ApiCalendarDeleteCalendarAclV4, calendarId, aclId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取访问控制列表 https://open.feishu.cn/document/ukTMukTMukTM/ucjM3YjL3IzN24yNycjN
func (t Tenant) GetCalendarAclList(calendarId string, pageSize int, pageToken string) (*vo.GetCalendarAclListV4Resp, error) {
	params := map[string]interface{}{}
	if pageSize > 0 {
		params["page_size"] = pageSize
	}
	if pageToken != "" {
		params["page_token"] = pageToken
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarCalendarAclGetV4, calendarId), params, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarAclListV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
