package application

import (
	"actividad/src/tiendas/domain"
)

type EditTienda struct {
	db domain.ITienda
}

func NewEditTienda(db domain.ITienda) *EditTienda {
	return &EditTienda{db: db}
}

func (ep *EditTienda) Execute(id int32, nombre string, ubicacion string) error {
	return ep.db.UpdateTienda(id, nombre, ubicacion)
}
