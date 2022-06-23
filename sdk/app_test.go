package sdk

import (
	"github.com/yangjishen/feishu-sdk-golang/core/consts"
	"github.com/yangjishen/feishu-sdk-golang/core/util/json"
	"testing"
)

func TestTenant_IsUserAdmin(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.IsUserAdmin("ou_b7d861c94cea4d316a6bfc5e8421994c", "")
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_AdminUserList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.AdminUserList()
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_OrgInfo(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "116542e77eced75d")
	t.Log(e)

	resp, err := tenant.OrgInfo()
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_GetJsTicket(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "116542e77eced75d")
	t.Log(e)

	resp, err := tenant.GetJsTicket()
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))
}
