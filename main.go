package main

import (
	"github.com/insanXYZ/starter-golang-backend/bootstrap"
	"github.com/insanXYZ/starter-golang-backend/config"
)

func main() {
	viper := config.NewViper()
	validator := config.NewValidator()
	gorm := config.NewGorm(viper)
	echo := config.NewEcho(viper)

	bootstrapInit := bootstrap.Configs{
		Viper:     viper,
		Gorm:      gorm,
		Echo:      echo,
		Validator: validator,
	}

	bootstrapInit.Run()

	echo.Logger.Fatal(echo.Start("WEB_PORT"))
}
