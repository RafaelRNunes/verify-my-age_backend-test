package config

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/application"
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/repository"
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/usecase"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/encryption"
	dao "github.com/RafaelRNunes/verify-my-age_backend-test/infra/repository"
)

var (
	encryptor      usecase.Encryptor         = encryption.NewEncryptorBcrypt()
	userRepository repository.UserRepository = dao.NewUserMySqlRepository()
	userService                              = application.NewUserService(userRepository, encryptor)
)

func NewUserService() *application.UserService {
	return userService
}
