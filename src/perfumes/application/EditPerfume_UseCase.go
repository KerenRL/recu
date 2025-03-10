package application

import (
    "actividad/src/perfumes/domain"
)

type EditPerfume struct {
    db domain.IPerfume
}

func NewEditPerfume(db domain.IPerfume) *EditPerfume {
    return &EditPerfume{db: db}
}

func (ep *EditPerfume) Execute(id int32, marca string, modelo string, precio float32) error {
    return ep.db.UpdatePerfume(id, marca, modelo, precio)
}