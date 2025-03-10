package application

import (
	"actividad/src/tiendas/domain"
)

type ViewTienda struct {
	db domain.ITienda
}

func NewViewTienda(db domain.ITienda) *ViewTienda {
	return &ViewTienda{db: db}
}

func (vp *ViewTienda) Execute() ([]domain.Tienda, error) {
	return vp.db.GetAll()
}
