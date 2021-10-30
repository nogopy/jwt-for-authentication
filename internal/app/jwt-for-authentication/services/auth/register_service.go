package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/models"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/repositories"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/utils"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/utils/exception"
	log "github.com/sirupsen/logrus"
)

type RegisterService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewRegisterService(
	userRepo repositories.UserRepositoryInterface,
) *RegisterService {
	return &RegisterService{
		userRepo: userRepo,
	}
}

func (service *RegisterService) Register(ctx *gin.Context) (string, error) {
	var body RegisterBodySchema

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Errorf("cannot bind request body, exception: %s", err.Error())
		return "", err
	}

	user, err := service.userRepo.FindByUsername(body.UserName)
	if err != nil {
		return "", err
	}

	if user.ID > 0 {
		return "", exception.UsernameExists
	}

	newUser := models.User{
		Username: body.UserName,
		Password: utils.Hash(body.Password),
	}
	if err2 := service.userRepo.CreateOne(newUser); err2 != nil {
		return "", err2
	}

	return newUser.Username, nil
}
