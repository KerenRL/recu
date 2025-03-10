package application

import (
	"actividad/src/perfumes/domain"
)

type DeletePerfume struct {
	db domain.IPerfume
}

func NewDeletePerfume(db domain.IPerfume) *DeletePerfume {
	return &DeletePerfume{db: db}
}

func (dp *DeletePerfume) Execute(id int32) error {
	return dp.db.DeletePerfume(id)
}
