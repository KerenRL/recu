package application

import (
	"actividad/src/perfumes/domain"
)

type CreatePerfume struct {
	db      domain.IPerfume
	encrypt domain.IEncryptService
}

func NewCreatePerfume(db domain.IPerfume, encrypt domain.IEncryptService) *CreatePerfume {
	return &CreatePerfume{db: db, encrypt: encrypt}
}

func (cp *CreatePerfume) Execute(marca string, modelo string, precio float32) error {
	encryptedPrice, err := cp.encrypt.Encrypt(precio)
	if err != nil {
		return err
	}
	return cp.db.SavePerfume(marca, modelo, encryptedPrice)
}
