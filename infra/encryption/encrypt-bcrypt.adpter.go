package encryption

import (
	_ "github.com/RafaelRNunes/verify-my-age_backend-test/domain/usecase"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type EncrypterBcryptAdapter struct {
}

func NewEncryptorBcrypt() *EncrypterBcryptAdapter {
	return &EncrypterBcryptAdapter{}
}

func (this *EncrypterBcryptAdapter) Encrypt(value string) (string, error) {
	crypto, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return string(crypto), nil
}
