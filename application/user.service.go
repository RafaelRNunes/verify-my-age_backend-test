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

func (this *UserService) Create(userInputDto dto.UserInput) *dto.UserOutput {
	userEntity := userInputDto.MapUserToEntity()

	userEntity.Password, _ = this.encryptor.Encrypt(userEntity.Password)

	newUser := this.repository.Create(*userEntity)

	return dto.MapUserToDto(&newUser)
}

func (this *UserService) FindAll() []dto.UserOutput {
	users := this.repository.FindAll()
	usersDto := make([]dto.UserOutput, 0, len(users))

	for _, user := range users {
		usersDto = append(usersDto, *dto.MapUserToDto(&user))
	}

	return usersDto
}

func (this *UserService) FindById(userId int) dto.UserOutput {
	user := this.repository.FindById(userId)
	return *dto.MapUserToDto(&user)
}

func (this *UserService) Update(userId int, user dto.UserInput) dto.UserOutput {
	userEntityUpdated := this.repository.Update(userId, *user.MapUserToEntity())
	return *dto.MapUserToDto(&userEntityUpdated)
}

func (this *UserService) Delete(userId int) bool {
	user := this.repository.FindById(userId)

	if user.Id == 0 {
		return false
	}

	return this.repository.Delete(userId)
}
