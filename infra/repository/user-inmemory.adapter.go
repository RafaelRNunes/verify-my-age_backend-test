package repository

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
)

type UserMemoryRepository struct {
	users []entity.User
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{users: make([]entity.User, 0)}
}

func (this *UserMemoryRepository) Create(user entity.User) entity.User {
	user.Id = len(this.users) + 1
	user.Address.Id = len(this.users) + 1
	this.users = append(this.users, user)
	return user
}

func (this *UserMemoryRepository) FindAll() []entity.User {
	return this.users
}

func (this *UserMemoryRepository) FindById(userId int) entity.User {
	for _, userOld := range this.users {
		if userOld.Id == userId {
			return userOld
		}
	}

	return entity.User{}
}

func (this *UserMemoryRepository) Update(userId int, user entity.User) entity.User {
	user.Id = userId

	for i, userOld := range this.users {
		if userOld.Id == userId {
			this.users[i] = user
			break
		}
	}

	return user
}

func (this *UserMemoryRepository) Delete(userId int) bool {
	var index int
	for i, user := range this.users {
		if user.Id == userId {
			index = i
			break
		}
	}

	this.users = remove(this.users, index)
	return true
}

func remove(s []entity.User, i int) []entity.User {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
