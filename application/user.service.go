package application

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/application/dto"
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/repository"
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/usecase"
)

type UserService struct {
	repository repository.UserRepository
	encryptor  usecase.Encryptor
}

func NewUserService(repository repository.UserRepository, encryptor usecase.Encryptor) *UserService {
	return &UserService{repository: repository, encryptor: encryptor}
}

func (this *UserService) Create(userInput dto.UserInput) bool {
	user := userInput.MapUser()
	user.Password, _ = this.encryptor.Encrypt(user.Password)

	return this.repository.Create(*user)
}
