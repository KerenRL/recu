package infraestructure

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type EncryptService struct{}

func NewEncryptService() *EncryptService {
	return &EncryptService{}
}

func (es *EncryptService) Encrypt(value float32) (string, error) {
	priceStr := fmt.Sprintf("%.2f", value)
	hash, err := bcrypt.GenerateFromPassword([]byte(priceStr), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}