package repository

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(user entity.User) bool
	FindAll() map[uuid.UUID]entity.User
	FindById(userId uuid.UUID) entity.User
	Update(userId uuid.UUID, user entity.User) bool
	Delete(userId uuid.UUID) bool
}
