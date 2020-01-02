package rest

import (
	"ddbs-final/internal/middlewares"
	"ddbs-final/pkg/rest/common"
	"ddbs-final/pkg/rest/order"
	"ddbs-final/pkg/rest/trace"
	"ddbs-final/pkg/rest/user"
	"github.com/gin-gonic/gin"
)

func REST(engine *gin.Engine) {
	guestRoute := engine.Group("/user")
	{
		guestRoute.PUT("/", user.UserRegister)
		guestRoute.POST("/", user.UserLogin)
		guestRoute.GET("/area", common.GetAllAreas)
	}

	userRoute := engine.Group("/user")
	userRoute.Use(middlewares.UserAuth())
	{
		userRoute.DELETE("/", user.UserLogout)
	}

	adminRoute := engine.Group("/admin")
	adminRoute.Use(middlewares.AdminAuth())
	{
		adminRoute.GET("/order", order.GetAllOrders)
		adminRoute.PUT("/trace", trace.CreateTrace)
		adminRoute.POST("/trace", trace.UpdateTrace)
	}

	orderRoute := engine.Group("/order")
	orderRoute.Use(middlewares.UserAuth())
	{
		orderRoute.PUT("/", order.CreateOrder)
		orderRoute.GET("/", order.GetOrder)
		orderRoute.DELETE("/", order.TerminateOrder)
	}

	traceRoute := engine.Group("/trace")
	{
		traceRoute.GET("/", trace.GetTrace)
	}
}
