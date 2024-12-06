package doctor

import (
	"med/internal/domain"
)

type UpdateDoctor struct {
	Org  *string
	Job  *string
	Desc *string
}

func (u UpdateDoctor) Update(d *domain.Doctor) {
	if u.Org != nil {
		d.Org = *u.Org
	}
	if u.Job != nil {
		d.Job = *u.Job
	}
	if u.Desc != nil {
		d.Desc = u.Desc
	}
}
