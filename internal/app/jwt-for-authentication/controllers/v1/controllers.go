package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/services"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/services/auth"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/utils"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/utils/exception"
)

type AuthControllers struct {
	registerService          *auth.RegisterService
	loginService             *auth.LoginService
	getUserInfoService       *services.GetUserInfoService
	getPublicResourceService *services.GetPublicResourceService
}

func NewAuthControllers(
	registerService *auth.RegisterService,
	loginService *auth.LoginService,
	getUserInfoService *services.GetUserInfoService,
	getPublicResourceService *services.GetPublicResourceService,
) *AuthControllers {
	return &AuthControllers{
		registerService:          registerService,
		loginService:             loginService,
		getUserInfoService:       getUserInfoService,
		getPublicResourceService: getPublicResourceService,
	}
}

func (ctrl *AuthControllers) RegisterHandler(ctx *gin.Context) {
	userName, err := ctrl.registerService.Register(ctx)
	if err != nil {
		if errors.Is(err, exception.UsernameExists) {
			utils.RespondBadRequest(ctx, err.Error())
			return
		}

		utils.RespondInternalError(ctx)
		return
	}

	utils.RespondSuccessful(ctx, gin.H{
		"username": userName,
		"greeting": fmt.Sprintf("hello %s, welcome to nogopy", userName),
	})
}

func (ctrl *AuthControllers) LoginHandler(ctx *gin.Context) {
	loginResult, err := ctrl.loginService.Login(ctx)

	if err != nil {
		if errors.Is(err, exception.WrongPassword) {
			utils.RespondBadRequest(ctx, err.Error())
			return
		}

		utils.RespondInternalError(ctx)
		return
	}

	utils.RespondSuccessful(ctx, loginResult)
}

func (ctrl *AuthControllers) GetUserInfoHandler(ctx *gin.Context) {
	info, err := ctrl.getUserInfoService.Get(ctx)

	if err != nil {

		utils.RespondInternalError(ctx)
		return
	}

	utils.RespondSuccessful(ctx, info)
}

func (ctrl *AuthControllers) GetPublicResourceHandler(ctx *gin.Context) {
	info, err := ctrl.getPublicResourceService.Get(ctx)

	if err != nil {

		utils.RespondInternalError(ctx)
		return
	}

	utils.RespondSuccessful(ctx, info)
}
