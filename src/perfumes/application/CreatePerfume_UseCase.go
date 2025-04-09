package application

import (
	"actividad/src/perfumes/domain"
)

type CreatePerfume struct {
	db domain.IPerfume
}

func NewCreatePerfume(db domain.IPerfume) *CreatePerfume {
	return &CreatePerfume{db: db}
}

func (cp *CreatePerfume) Execute(marca string, modelo string, precio float32) (int32, error) {
	id, err := cp.db.SavePerfume(marca, modelo, precio)
	if err != nil {
		return 0, err
	}
	return id, nil
}