package main

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/application"
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/repository"
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/usecase"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/encryption"
	dao "github.com/RafaelRNunes/verify-my-age_backend-test/infra/repository"
)

func main() {
	var encryptor usecase.Encryptor = encryption.NewEncryptorBcrypt()
	var userRepository repository.UserRepository = dao.NewUserDAO()
	var userService application.UserService = application.NewUserService(userRepository, encryptor)

	userService.Create()
}
