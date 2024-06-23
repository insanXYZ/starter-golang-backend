package bootstrap

import (
	"backend/controller"
	"backend/middleware"
	"backend/repository"
	"backend/route"
	"backend/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Configs struct {
	Viper     *viper.Viper
	Gorm      *gorm.DB
	Echo      *echo.Echo
	Validator *validator.Validate
}

func (c *Configs) Run() {

	// Repositories
	userRepository := repository.NewUserRepository(c.Gorm)

	// Services
	userService := service.NewUserService(c.Gorm, c.Viper, c.Validator, userRepository)

	// Controllers

	userController := controller.NewUserController(userService)

	// Middlewares

	middlewareConfig := middleware.NewMiddleware(c.Viper)

	routeconfigs := route.RouteConfig{
		Echo:           c.Echo,
		Middlewares:    middlewareConfig,
		UserController: userController,
	}

	routeconfigs.Setup()
}
