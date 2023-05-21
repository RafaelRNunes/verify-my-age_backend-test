package repository

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/database"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/database/models"
)

type UserMySqlRepository struct {
}

func NewUserMySqlRepository() *UserMySqlRepository {
	return &UserMySqlRepository{}
}

func (this *UserMySqlRepository) Create(user entity.User) entity.User {
	userModel := models.MapUserToModel(user)
	database.DB.Create(&userModel)
	return *userModel.MapUserToEntity()
}

func (this *UserMySqlRepository) FindAll() []entity.User {
	var usersModels []models.User
	database.DB.Preload("Address").Find(&usersModels)
	var users []entity.User

	for _, user := range usersModels {
		users = append(users, *user.MapUserToEntity())
	}

	return users
}

func (this *UserMySqlRepository) FindById(userId int) entity.User {
	return entity.User{}
}

func (this *UserMySqlRepository) Update(userId int, user entity.User) entity.User {
	return entity.User{}
}

func (this *UserMySqlRepository) Delete(userId int) bool {
	return true
}
