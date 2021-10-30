package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/controllers/v1"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/middleware"
)

type TikiNowRouter struct {
	GinEngie   *gin.Engine
	Controller *v1.AuthControllers
}

func NewTikiNowRouter(g *gin.Engine, c *v1.AuthControllers) *TikiNowRouter {
	return &TikiNowRouter{
		GinEngie:   g,
		Controller: c,
	}
}

func (r *TikiNowRouter) SetUpRouter() {
	router := r.GinEngie.Group("v1")
	{
		router.POST("auth/users/register", r.Controller.RegisterHandler)
		router.POST("auth/users/login", r.Controller.LoginHandler)
	}

	routerWithAuth := r.GinEngie.Group("v1").Use(middleware.Authentication())
	{
		routerWithAuth.GET("info", r.Controller.GetUserInfoHandler)
	}

	routerWithoutAuth := r.GinEngie.Group("v1")
	{
		routerWithoutAuth.GET("public", r.Controller.GetPublicResourceHandler)
	}
}
