package segment

import "yir/uzi/internal/domain"

// TODO: починить баг при запросе со всеми полями nil
type UpdateSegment struct {
	Contor   *string
	Tirads23 *float64
	Tirads4  *float64
	Tirads5  *float64
}

func (u UpdateSegment) Update(d *domain.Segment) {
	if u.Contor != nil {
		d.Contor = *u.Contor
	}
	if u.Tirads23 != nil {
		d.Tirads23 = *u.Tirads23
	}
	if u.Tirads4 != nil {
		d.Tirads4 = *u.Tirads4
	}
	if u.Tirads5 != nil {
		d.Tirads5 = *u.Tirads5
	}
}
