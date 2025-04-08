package domain

type IEncryptService interface {
	Encrypt(value float32) (string, error)
}
