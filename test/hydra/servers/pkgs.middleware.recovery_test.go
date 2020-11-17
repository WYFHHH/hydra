package servers

import (
	"net/http"
	"testing"

	"github.com/micro-plat/hydra/hydra/servers/pkg/middleware"
	"github.com/micro-plat/hydra/test/assert"
	"github.com/micro-plat/hydra/test/mocks"
)

func TestRecovery(t *testing.T) {
	tests := []struct {
		name           string
		requestURL     string
		clientIP       string
		requstID       string
		method         string
		responseStatus int
		wantStatus     int
		wantContent    string
		next           func()
	}{
		{name: "请求返回200", method: "OPTIONS", clientIP: "127.0.0.1", requstID: "06c6fb24c", responseStatus: 200, wantStatus: 200, wantContent: ""},
		{name: "请求返回400", method: "GET", clientIP: "127.0.0.1", requstID: "06c6fb24c", responseStatus: 400, wantStatus: 400, wantContent: ""},
		{name: "请求过程中出现的非预见的错误", method: "POST", clientIP: "127.0.0.1", requstID: "06c6fb24c", next: func() { panic("error") }, responseStatus: 400, wantStatus: http.StatusNotExtended, wantContent: "error"},
	}

	for _, tt := range tests {
		//初始化测试用例参数
		ctx := &mocks.MiddleContext{
			MockNext:     tt.next,
			MockUser:     &mocks.MockUser{MockClientIP: tt.clientIP, MockRequestID: tt.requstID},
			MockRequest:  &mocks.MockRequest{MockPath: &mocks.MockPath{MockURL: tt.requestURL, MockMethod: tt.method}},
			MockResponse: &mocks.MockResponse{MockStatus: tt.responseStatus},
			MockAPPConf:  mocks.NewConf().GetAPIConf(),
		}

		//调用中间件
		handler := middleware.Recovery()
		handler(ctx)

		gotStatus, gotContent := ctx.Response().GetFinalResponse()
		assert.Equalf(t, tt.wantStatus, gotStatus, tt.name)
		assert.Equalf(t, tt.wantContent, gotContent, tt.name)

	}
}