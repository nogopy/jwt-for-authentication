package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/utils"
)

type GetUserInfoService struct {
}

func NewGetUserInfoService() *GetUserInfoService {
	return &GetUserInfoService{}
}

func (service *GetUserInfoService) Get(ctx *gin.Context) (string, error) {
	userId := utils.GetUserIdFromContext(ctx)
	return fmt.Sprintf("hello user_id = %v, I know that you've been authenticated by our service", userId), nil
}
