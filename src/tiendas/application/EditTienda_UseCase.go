package application

import (
	"actividad/src/tiendas/domain"
	"fmt"
)

type EditTienda struct {
	db domain.ITienda
}

func NewEditTienda(db domain.ITienda) *EditTienda {
	return &EditTienda{db: db}
}

func (ep *EditTienda) Execute(id int32, nombre string, ubicacion string) error {
	if id <= 0 {
		return fmt.Errorf("ID de tienda inválido")
	}
	return ep.db.UpdateTienda(id, nombre, ubicacion)
}
