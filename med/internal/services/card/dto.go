package card

import (
	"med/internal/domain"
)

type UpdateCard struct {
	Diagnosis *string
}

func (u UpdateCard) Update(d *domain.Card) {
	if u.Diagnosis != nil {
		d.Diagnosis = u.Diagnosis
	}
}
