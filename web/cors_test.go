package web_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink/internal/cltest"
)

func TestCors_DefaultOrigins(t *testing.T) {
	t.Parallel()

	config, _ := cltest.NewConfig()
	app, appCleanup := cltest.NewApplicationWithConfig(config)
	defer appCleanup()

	headers := map[string]string{"Origin": "http://localhost:3000"}
	resp := cltest.BasicAuthGet(app.Server.URL+"/v2/config", headers)
	cltest.AssertServerResponse(t, resp, 200)

	headers = map[string]string{"Origin": "http://localhost:6689"}
	resp = cltest.BasicAuthGet(app.Server.URL+"/v2/config", headers)
	cltest.AssertServerResponse(t, resp, 200)

	headers = map[string]string{"Origin": "http://localhost:1234"}
	resp = cltest.BasicAuthGet(app.Server.URL+"/v2/config", headers)
	cltest.AssertServerResponse(t, resp, 403)
}

func TestCors_OverrideOrigins(t *testing.T) {
	t.Parallel()

	config, _ := cltest.NewConfig()
	config.AllowOrigins = "http://chainlink.com"
	app, appCleanup := cltest.NewApplicationWithConfig(config)
	defer appCleanup()

	headers := map[string]string{"Origin": "http://chainlink.com"}
	resp := cltest.BasicAuthGet(app.Server.URL+"/v2/config", headers)
	cltest.AssertServerResponse(t, resp, 200)

	headers = map[string]string{"Origin": "http://localhost:3000"}
	resp = cltest.BasicAuthGet(app.Server.URL+"/v2/config", headers)
	cltest.AssertServerResponse(t, resp, 403)
}
