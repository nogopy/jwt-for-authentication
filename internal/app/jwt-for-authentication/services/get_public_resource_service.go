package services

import "github.com/gin-gonic/gin"

type GetPublicResourceService struct {
}

func NewGetPublicResourceService() *GetPublicResourceService {
	return &GetPublicResourceService{}
}

func (service *GetPublicResourceService) Get(ctx *gin.Context) (string, error) {
	return "every one can access this resource cause it's public", nil
}
