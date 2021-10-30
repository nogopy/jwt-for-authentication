package repositories

import "github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/models"

type UserRepositoryInterface interface {
	FindByUsername(userName string) (models.User, error)
	CreateOne(user models.User) error
}
