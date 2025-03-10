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

func (cp *CreatePerfume) Execute(marca string, modelo string, precio float32) error {
    return cp.db.SavePerfume(marca, modelo, precio)
}