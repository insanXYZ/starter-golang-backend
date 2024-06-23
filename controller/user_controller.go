package controller

import (
	"backend/model"
	"backend/model/converter"
	"backend/service"
	"backend/utils/httpresponse"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (controller *UserController) Register(c echo.Context) error {
	req := new(model.RegisterUser)
	err := c.Bind(req)
	if err != nil {
		return httpresponse.Error(c, err.Error(), nil)
	}
	err = controller.UserService.Register(req)
	if err != nil {
		return httpresponse.Error(c, err.Error(), nil)
	}

	return httpresponse.Success(c, "success register user", nil)
}

func (controller *UserController) Login(c echo.Context) error {
	req := new(model.LoginUser)
	err := c.Bind(req)
	if err != nil {
		return httpresponse.Error(c, err.Error(), nil)
	}

	token, err := controller.UserService.Login(req)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = *token
	cookie.Path = "/"

	c.SetCookie(cookie)

	return httpresponse.Success(c, "success login", converter.UserToToken(*token))

}

func (controller *UserController) Refresh(c echo.Context) error {
	claims := c.Get("user").(jwt.MapClaims)

	token, err := controller.UserService.Refresh(claims)
	if err != nil {
		return err
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = *token
	cookie.Path = "/"
	c.SetCookie(cookie)

	return httpresponse.Success(c, "success refresh token", converter.UserToToken(*token))
}
