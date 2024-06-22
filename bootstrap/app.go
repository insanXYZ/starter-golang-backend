package bootstrap

import (
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

	// Services

	// Controller

}
