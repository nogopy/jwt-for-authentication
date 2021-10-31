package repositories

import (
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) FindByUsername(userName string) (models.User, error) {
	var user models.User

	err := repo.db.Where("username = ?", userName).Find(&user).Error
	//query := fmt.Sprintf("SELECT * FROM user where username = %s", userName)
	//repo.db.Raw(query).Scan(&user)

	return user, err
}

func (repo *UserRepository) CreateOne(user models.User) error {
	return repo.db.Create(&user).Error
}
