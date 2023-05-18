package repository

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
	_ "github.com/RafaelRNunes/verify-my-age_backend-test/domain/repository"
	"github.com/google/uuid"
)

type UserMemoryRepository struct {
	users map[uuid.UUID]entity.User
}

func NewUserDAO() *UserMemoryRepository {
	return &UserMemoryRepository{}
}

func (this *UserMemoryRepository) Create(user entity.User) bool {
	user.Id = uuid.New()
	this.users[user.Id] = user
	return true
}

func (this *UserMemoryRepository) FindAll() map[uuid.UUID]entity.User {
	return this.users
}

func (this *UserMemoryRepository) FindById(userId uuid.UUID) entity.User {
	return this.users[userId]
}

func (this *UserMemoryRepository) Update(userId uuid.UUID, user entity.User) bool {
	this.users[userId] = user
	return true
}

func (this *UserMemoryRepository) Delete(userId uuid.UUID) bool {
	delete(this.users, userId)
	return true
}
