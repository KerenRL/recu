package domain

type ITienda interface {
	SaveTienda(nombre string, direccion string) error
	GetAll() ([]Tienda, error)
	UpdateTienda(id int32, nombre string, direccion string) error
	DeleteTienda(id int32) error
}

type Tienda struct {
	ID        int32  `json:"id"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"`
}

func NewTienda(nombre string, direccion string) *Tienda {
	return &Tienda{Nombre: nombre, Direccion: direccion}
}

func (t *Tienda) SetNombre(nombre string) {
	t.Nombre = nombre
}
