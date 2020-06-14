package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/shiniao/gtodo/docs"
	"github.com/shiniao/gtodo/handler"
	"github.com/shiniao/gtodo/router/middleware"
)

func Load(g *gin.Engine) *gin.Engine {

	// * 全局 middleware
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	// * 404 handler
	g.NoRoute(handler.RouteNotFound)
	g.NoMethod(handler.RouteNotFound)

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 认证相关路由
	g.POST("/v1/register", handler.Register)
	g.POST("/v1/login", handler.Login)
	g.POST("/v1/login/phone", handler.PhoneLogin)
	g.GET("/v1/vcode", handler.VCode)

	// 账户相关
	g.GET("/v1/account", handler.Account)

	// todos 路由组
	todo := g.Group("/v1/todos")
	todo.Use(middleware.AuthMiddleware())
	{
		todo.GET("/", handler.FetchAllTodo)
		todo.GET("/:id", handler.FetchSingleTodo)
		todo.POST("/", handler.AddTodo)
		todo.PUT("/:id", handler.UpdateTodo)
		todo.DELETE("/:id", handler.DeleteTodo)
	}

	// api server 健康状况
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", handler.HealthCheck)
		svcd.GET("/disk", handler.DiskCheck)
		svcd.GET("/cpu", handler.CPUCheck)
		svcd.GET("/ram", handler.RAMCheck)
	}
	return g
}
