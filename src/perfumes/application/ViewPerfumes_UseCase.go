package application

import (
	"actividad/src/perfumes/domain"
)

type ViewPerfumes struct {
	db domain.IPerfume
}

func NewViewPerfumes(db domain.IPerfume) *ViewPerfumes {
	return &ViewPerfumes{db: db}
}

func (vp *ViewPerfumes) Execute() ([]domain.Perfume, error) {
	return vp.db.GetAll()
}
