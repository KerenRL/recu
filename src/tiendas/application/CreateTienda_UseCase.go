package application

import "actividad/src/tiendas/domain"

type CreateTienda struct {
	db domain.ITienda
}

func NewCreateTienda(db domain.ITienda) *CreateTienda {
	return &CreateTienda{db: db}
}

func (ct *CreateTienda) Execute(nombre string, ubicacion string) (int32, error) {
	id, err := ct.db.SaveTienda(nombre, ubicacion)
	if err != nil {
		return 0, err
	}
	return id, nil
}
