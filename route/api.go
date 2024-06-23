package route

import (
	"backend/controller"
	"backend/middleware"
	"github.com/labstack/echo/v4"
)

type RouteConfig struct {
	Echo           *echo.Echo
	Middlewares    *middleware.MiddlewareConfig
	UserController *controller.UserController
}

func (c *RouteConfig) Setup() {
	g := c.Echo.Group("/api")
	g.POST("/register", c.UserController.Register)
	g.POST("/login", c.UserController.Login)
	g.POST("/refresh", c.UserController.Refresh, c.Middlewares.Refresh)
}
