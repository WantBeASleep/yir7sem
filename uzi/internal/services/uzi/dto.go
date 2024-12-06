package uzi

import (
	"uzi/internal/domain"
)

type UpdateUzi struct {
	Projection *string
	Checked    *bool
}

func (u UpdateUzi) Update(d *domain.Uzi) {
	if u.Projection != nil {
		d.Projection = *u.Projection
	}
	if u.Checked != nil {
		d.Checked = *u.Checked
	}
}

type UpdateEchographic struct {
	Contors         *string
	LeftLobeLength  *float64
	LeftLobeWidth   *float64
	LeftLobeThick   *float64
	LeftLobeVolum   *float64
	RightLobeLength *float64
	RightLobeWidth  *float64
	RightLobeThick  *float64
	RightLobeVolum  *float64
	GlandVolum      *float64
	Isthmus         *float64
	Struct          *string
	Echogenicity    *string
	RegionalLymph   *string
	Vascularization *string
	Location        *string
	Additional      *string
	Conclusion      *string
}

// TODO: пойти написать рефлексию
func (u UpdateEchographic) Update(d *domain.Echographic) {
	if u.Contors != nil {
		d.Contors = u.Contors
	}
	if u.LeftLobeLength != nil {
		d.LeftLobeLength = u.LeftLobeLength
	}
	if u.LeftLobeWidth != nil {
		d.LeftLobeWidth = u.LeftLobeWidth
	}
	if u.LeftLobeThick != nil {
		d.LeftLobeThick = u.LeftLobeThick
	}
	if u.LeftLobeVolum != nil {
		d.LeftLobeVolum = u.LeftLobeVolum
	}
	if u.RightLobeLength != nil {
		d.RightLobeLength = u.RightLobeLength
	}
	if u.RightLobeWidth != nil {
		d.RightLobeWidth = u.RightLobeWidth
	}
	if u.RightLobeThick != nil {
		d.RightLobeThick = u.RightLobeThick
	}
	if u.RightLobeVolum != nil {
		d.RightLobeVolum = u.RightLobeVolum
	}
	if u.GlandVolum != nil {
		d.GlandVolum = u.GlandVolum
	}
	if u.Isthmus != nil {
		d.Isthmus = u.Isthmus
	}
	if u.Struct != nil {
		d.Struct = u.Struct
	}
	if u.Echogenicity != nil {
		d.Echogenicity = u.Echogenicity
	}
	if u.RegionalLymph != nil {
		d.RegionalLymph = u.RegionalLymph
	}
	if u.Vascularization != nil {
		d.Vascularization = u.Vascularization
	}
	if u.Location != nil {
		d.Location = u.Location
	}
	if u.Additional != nil {
		d.Additional = u.Additional
	}
	if u.Conclusion != nil {
		d.Conclusion = u.Conclusion
	}
}
