package unit

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/usecase"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/encryption"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestEncryptBcrypt(t *testing.T) {
	var encrypter usecase.Encryptor = encryption.NewEncryptorBcrypt()
	response, err := encrypter.Encrypt("123Mudar")
	assert.Equal(t, err, nil)
	assert.NotNil(t, response)
	err = bcrypt.CompareHashAndPassword([]byte(response), []byte("123Mudar"))
	assert.Nil(t, err)
}
