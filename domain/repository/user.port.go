package repository

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
)

type UserRepository interface {
	Create(user entity.User) entity.User
	FindAll() []entity.User
	FindById(userId int) entity.User
	Update(userId int, user entity.User) entity.User
	Delete(userId int) bool
}
