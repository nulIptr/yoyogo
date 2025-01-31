package Endpoints

import (
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
)

func UseHealth(router Router.IRouterBuilder) {
	XLog.GetXLogger("Endpoint").Debug("loaded health endpoint.")

	router.GET("/actuator/health", func(ctx *Context.HttpContext) {
		ctx.JSON(200, Context.H{
			"status": "UP",
		})
	})
}
