package application

import (
    "actividad/src/tiendas/domain"
)

type CreateTienda struct {
    db domain.ITienda
}

func NewCreateTienda(db domain.ITienda) *CreateTienda {
    return &CreateTienda{db: db}
}

func (cp *CreateTienda) Execute(nombre string, ubicacion string) error {
    return cp.db.SaveTienda(nombre, ubicacion)
}
