package sdk

import (
	"github.com/yangjishen/feishu-sdk-golang/core/consts"
	"github.com/yangjishen/feishu-sdk-golang/core/model/vo"
	"github.com/yangjishen/feishu-sdk-golang/core/util/json"
	"github.com/yangjishen/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestTenant_SearchDocs(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)
	//a := int(50)
	resp, err := tenant.SearchDocs("u-KaKuv3qHfriTjs2LI0wTah", vo.SearchDocsReqVo{
		SearchKey: nil,
		//Count:     &a,
		//Offset:    &a,
		OwnerIds: nil,
		ChatIds:  nil,
		//DocsTypes: &[]string{"doc", "file},
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetDocMeta(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)
	resp, err := tenant.GetDocMeta("u-b2pimsbPjM7n6Q68L2quma", "shtcnXthMbCHWH5SPd3GGOiMVpg")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
