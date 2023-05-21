package usecase

type Encryptor interface {
	Encrypt(value string) (string, error)
}
