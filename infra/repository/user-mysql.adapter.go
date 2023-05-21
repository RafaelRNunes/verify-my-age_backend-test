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
	userModel := models.User{}
	userModel.MapUserToModel(user)
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
	var user models.User
	database.DB.Preload("Address").First(&user, userId)
	return *user.MapUserToEntity()
}

func (this *UserMySqlRepository) Update(userId int, user entity.User) entity.User {
	var userModel models.User
	database.DB.First(&userModel, userId)

	if userModel.ID == 0 {
		return entity.User{}
	}

	userModel.MapUserToModel(user)
	database.DB.Model(&userModel).Updates(userModel)
	database.DB.First(&userModel, userId)

	return *userModel.MapUserToEntity()
}

func (this *UserMySqlRepository) Delete(userId int) bool {
	var userModel models.User
	database.DB.First(&userModel, userId)

	if userModel.ID == 0 {
		return false
	}

	database.DB.Delete(&userModel, userId)
	return true
}
