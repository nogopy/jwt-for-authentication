package infrastructure

import (
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	env "github.com/nogopy/jwt-for-authentication/config/environment"
	v1 "github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/controllers/v1"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/repositories"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/routers"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/services"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/services/auth"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/utils"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func NewGinEngine() *gin.Engine {
	var router *gin.Engine

	if env.GetConfiguration().AppEnv != "local" {
		router = gin.New()
		router.Use(gin.Recovery())
	} else {
		router = gin.Default()
	}

	router.Use(nice.Recovery(recoveryHandler))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"pong": "Server is still alive  and running ... " + time.Now().String()})
	})

	router.Use(cors.Default())

	return router
}

func recoveryHandler(ctx *gin.Context, err interface{}) {
	utils.RespondInternalError(ctx)
}

func SetupHttpServer(db *gorm.DB, ginEngie *gin.Engine) {
	userRepo := repositories.NewUserRepository(db)

	loginService := auth.NewLoginService(userRepo)
	registerService := auth.NewRegisterService(userRepo)
	getUserInfoService := services.NewGetUserInfoService()
	getPublicResourceService := services.NewGetPublicResourceService()

	controller := v1.NewAuthControllers(registerService, loginService, getUserInfoService, getPublicResourceService)

	router := routers.NewTikiNowRouter(ginEngie, controller)
	router.SetUpRouter()
}
