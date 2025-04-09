package domain

type ITienda interface {
	SaveTienda(nombre string, ubicacion string) (int32, error)
	GetAll() ([]Tienda, error)
	UpdateTienda(id int32, nombre string, ubicacion string) error
	DeleteTienda(id int32) error
}

type Tienda struct {
	ID        int32  `json:"id"`
	Nombre    string `json:"nombre"`
	Ubicacion string `json:"ubicacion"`
}

func NewTienda(nombre string, ubicacion string) *Tienda {
	return &Tienda{Nombre: nombre, Ubicacion: ubicacion}
}
