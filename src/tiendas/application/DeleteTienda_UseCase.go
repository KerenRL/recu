package application

import "actividad/src/tiendas/domain"

type DeleteTienda struct {
	db domain.ITienda
}

func NewDeleteTienda(db domain.ITienda) *DeleteTienda {
	return &DeleteTienda{db: db}
}

func (dp *DeleteTienda) Execute(id int32) error {
	return dp.db.DeleteTienda(id)
}
