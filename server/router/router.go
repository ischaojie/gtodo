package router

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/shiniao/gtodo/docs"
	"github.com/shiniao/gtodo/handler"
	"github.com/shiniao/gtodo/middleware"
	"net/http"
)

func Load(g *gin.Engine) *gin.Engine {

	// * 全局middlewares

	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	// * 404 handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// swagger api docs
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// * token下发
	g.POST("/token", handler.Token)

	// * todos 路由组
	v1 := g.Group("/v1/todos")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/", handler.FetchAllTodo)
		v1.GET("/:id", handler.FetchSingleTodo)
		v1.POST("/", handler.AddTodo)
		v1.PUT("/:id", handler.UpdateTodo)
		v1.DELETE("/:id", handler.DeleteTodo)
	}

	// * api server 健康状况
	svcd := g.Group("/sd")

	{
		svcd.GET("/health", handler.HealthCheck)
		svcd.GET("/disk", handler.DiskCheck)
		svcd.GET("/cpu", handler.CPUCheck)
		svcd.GET("/ram", handler.RAMCheck)
	}
	return g
}
