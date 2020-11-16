package context

import (
	"github.com/micro-plat/lib4go/tgo"
)

//GetTGOModules 获取tgo的模块配置信息
func GetTGOModules() []*tgo.Module {
	request := tgo.NewModule("request").
		Add("getString", tgo.FuncASRS(func(input string) string { ctx := Current(); return ctx.Request().GetString(input) })).
		Add("getPath", tgo.FuncARS(func() string { ctx := Current(); return ctx.Request().Path().GetRequestPath() })).
		Add("getHeader", tgo.FuncASRS(func(input string) string { ctx := Current(); return ctx.Request().Path().GetHeader(input) })).
		Add("getCookie", tgo.FuncASRS(func(input string) string { ctx := Current(); c, _ := ctx.Request().Path().GetCookie(input); return c })).
		Add("getClientIP", tgo.FuncARS(func() string { ctx := Current(); return ctx.User().GetClientIP() })).
		Add("getRequestID", tgo.FuncARS(func() string { ctx := Current(); return ctx.User().GetRequestID() }))

	response := tgo.NewModule("response").
		Add("getStatus", tgo.FuncARI(func() int { ctx := Current(); s, _ := ctx.Response().GetFinalResponse(); return s })).
		Add("getContent", tgo.FuncARS(func() string { ctx := Current(); _, s := ctx.Response().GetFinalResponse(); return s }))

	app := tgo.NewModule("app").
		Add("getServerID", tgo.FuncARS(func() string { ctx := Current(); return ctx.APPConf().GetServerConf().GetServerID() })).
		Add("getPlatName", tgo.FuncARS(func() string { ctx := Current(); return ctx.APPConf().GetServerConf().GetPlatName() })).
		Add("getSysName", tgo.FuncARS(func() string { ctx := Current(); return ctx.APPConf().GetServerConf().GetSysName() })).
		Add("getServerType", tgo.FuncARS(func() string { ctx := Current(); return ctx.APPConf().GetServerConf().GetServerType() })).
		Add("getClusterName", tgo.FuncARS(func() string { ctx := Current(); return ctx.APPConf().GetServerConf().GetClusterName() })).
		Add("getServerName", tgo.FuncARS(func() string { ctx := Current(); return ctx.APPConf().GetServerConf().GetServerName() })).
		Add("getServerPath", tgo.FuncARS(func() string { ctx := Current(); return ctx.APPConf().GetServerConf().GetServerPath() }))

	return []*tgo.Module{request, response, app}
}
