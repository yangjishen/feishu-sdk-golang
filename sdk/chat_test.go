package sdk

import (
	"fmt"
	"github.com/yangjishen/feishu-sdk-golang/core/consts"
	"github.com/yangjishen/feishu-sdk-golang/core/util/json"
	"github.com/yangjishen/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestTenant_GroupList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GroupList("u-OVRWRbISgBZK6j9pu2ApJg", 0, "")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ChatMembers(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "1319b042200f575e")
	t.Log(e)

	resp, err := tenant.ChatMembers("u-mY81BLL4kurNqgbN0tsFka", "oc_a414a24110b73441307413f1bf8f2a33", 0, "")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ChatSearch(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.ChatSearch("u-OVRWRbISgBZK6j9pu2ApJg", "北极星测试企业。", 0, "")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ImChatList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "1279794b670f575f")
	t.Log(e)

	resp, err := tenant.ImChatList("", 0, "", nil)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ImMessage(t *testing.T) {
	resp, err := GetTenantAccessTokenInternal(consts.TestAppId, consts.TestAppSecret)
	log.Info(json.ToJsonIgnoreError(resp), err)
	tenant := Tenant{
		TenantAccessToken: resp.TenantAccessToken,
	}
	body := map[string]interface{}{"receive_id": "on_a23cccb0e391e54f8936a287e3ab0516", "content": "{\"config\": {\"wide_screen_mode\": true}, \"header\": {\"title\": {\"content\": \"zhouwei\\u5728\\u6a21\\u578b/\\u573a\\u666fV8_engine(V1) (1)\\u4e2d\\u56de\\u590d\\u4e86\\u4f60\", \"tag\": \"plain_text\"}}, \"elements\": [{\"tag\": \"div\", \"text\": {\"content\": \"你好\", \"tag\": \"lark_md\"}}, {\"tag\": \"div\", \"text\": {\"content\": \"2022-07-08 11:00:00\", \"tag\": \"lark_md\"}}, {\"tag\": \"hr\"}, {\"elements\": [{\"content\": \"[\\u67e5\\u770b\\u8be6\\u60c5](http://hub.stage.realibox.com/project/2328215637847965876/2340192109332529193/2340192109975306281)\", \"tag\": \"lark_md\"}], \"tag\": \"note\"}]}", "msg_type": "interactive"}
	fmt.Println(body["content"])

	res, err := tenant.ImMessageInfo("union_id", body)
	log.Info(json.ToJsonIgnoreError(res), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
