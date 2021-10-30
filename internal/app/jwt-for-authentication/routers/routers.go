package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/controllers/v1"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/middleware"
)

type Router struct {
	GinEngie   *gin.Engine
	Controller *v1.AuthControllers
}

func NewRouter(g *gin.Engine, c *v1.AuthControllers) *Router {
	return &Router{
		GinEngie:   g,
		Controller: c,
	}
}

func (r *Router) SetUpRouter() {
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
