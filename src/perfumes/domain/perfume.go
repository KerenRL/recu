package domain

type IPerfume interface {
	SavePerfume(marca string, modelo string, precio float32) error
	GetAll() ([]Perfume, error)
	UpdatePerfume(id int32, marca string, modelo string, precio float32) error
	DeletePerfume(id int32) error
}

type Perfume struct {
	ID     int32   `json:"id"`
	Marca  string  `json:"marca"`
	Modelo string  `json:"modelo"`
	Precio float32 `json:"precio"`
}

func NewPerfume(marca string, modelo string, precio float32) *Perfume {
	return &Perfume{Marca: marca, Modelo: modelo, Precio: precio}
}

func (p *Perfume) SetMarca(marca string) {
	p.Marca = marca
}
