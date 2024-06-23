package service

import (
	"backend/entity"
	"backend/model"
	"backend/repository"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	DB        *gorm.DB
	Viper     *viper.Viper
	Validator *validator.Validate
	UserRepo  *repository.UserRepository
}

func NewUserService(DB *gorm.DB, viper *viper.Viper, validator *validator.Validate, userRepo *repository.UserRepository) *UserService {
	return &UserService{DB: DB, Viper: viper, Validator: validator, UserRepo: userRepo}
}

func (service *UserService) Register(req *model.RegisterUser) error {
	err := service.Validator.Struct(req)
	if err != nil {
		return err
	}

	if count := service.UserRepo.CountWhere(service.DB, "email = ?", req.Email); count == 0 {
		fmt.Println("masuk")
		passwordByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
		if err != nil {
			return err
		}

		user := entity.User{
			ID:       uuid.New().String(),
			Name:     req.Name,
			Email:    req.Email,
			Password: string(passwordByte),
		}

		err = service.UserRepo.Create(service.DB, &user)
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("email already taken")

}

func (service *UserService) Login(req *model.LoginUser) (*string, error) {
	err := service.Validator.Struct(req)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email: req.Email,
	}

	err = service.UserRepo.Where(service.DB, user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Duration(service.Viper.GetInt("JWT_EXP")) * time.Minute).Unix(),
	})

	signedString, err := claims.SignedString([]byte(service.Viper.GetString("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &signedString, nil
}

func (service *UserService) Refresh(claims jwt.MapClaims) (*string, error) {
	newClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  claims["sub"],
		"name": claims["name"],
		"exp":  time.Now().Add(time.Duration(service.Viper.GetInt("JWT_EXP")) * time.Minute).Unix(),
	})

	signedString, err := newClaims.SignedString([]byte(service.Viper.GetString("JWT_SECRET_KEY")))
	if err != nil {
		return nil, err
	}

	return &signedString, nil
}
