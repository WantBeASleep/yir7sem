package entity

import (
	"database/sql"
	// TODO: убрать gtclib и сделать pointer
	// например https://github.com/AlekSi/pointer
	"github.com/WantBeASleep/med_ml_lib/gtc"

	"uzi/internal/domain"

	"github.com/google/uuid"
)

type Echographic struct {
	Id              uuid.UUID       `db:"id"`
	Contors         sql.NullString  `db:"contors"`
	LeftLobeLength  sql.NullFloat64 `db:"left_lobe_length"`
	LeftLobeWidth   sql.NullFloat64 `db:"left_lobe_width"`
	LeftLobeThick   sql.NullFloat64 `db:"left_lobe_thick"`
	LeftLobeVolum   sql.NullFloat64 `db:"left_lobe_volum"`
	RightLobeLength sql.NullFloat64 `db:"right_lobe_length"`
	RightLobeWidth  sql.NullFloat64 `db:"right_lobe_width"`
	RightLobeThick  sql.NullFloat64 `db:"right_lobe_thick"`
	RightLobeVolum  sql.NullFloat64 `db:"right_lobe_volum"`
	GlandVolum      sql.NullFloat64 `db:"gland_volum"`
	Isthmus         sql.NullFloat64 `db:"isthmus"`
	Struct          sql.NullString  `db:"struct"`
	Echogenicity    sql.NullString  `db:"echogenicity"`
	RegionalLymph   sql.NullString  `db:"regional_lymph"`
	Vascularization sql.NullString  `db:"vascularization"`
	Location        sql.NullString  `db:"location"`
	Additional      sql.NullString  `db:"additional"`
	Conclusion      sql.NullString  `db:"conclusion"`
}

func (Echographic) FromDomain(d domain.Echographic) Echographic {
	return Echographic{
		Id:              d.Id,
		Contors:         gtc.String.PointerToSql(d.Contors),
		LeftLobeLength:  gtc.Float64.PointerToSql(d.LeftLobeLength),
		LeftLobeWidth:   gtc.Float64.PointerToSql(d.LeftLobeWidth),
		LeftLobeThick:   gtc.Float64.PointerToSql(d.LeftLobeThick),
		LeftLobeVolum:   gtc.Float64.PointerToSql(d.LeftLobeVolum),
		RightLobeLength: gtc.Float64.PointerToSql(d.RightLobeLength),
		RightLobeWidth:  gtc.Float64.PointerToSql(d.RightLobeWidth),
		RightLobeThick:  gtc.Float64.PointerToSql(d.RightLobeThick),
		RightLobeVolum:  gtc.Float64.PointerToSql(d.RightLobeVolum),
		GlandVolum:      gtc.Float64.PointerToSql(d.GlandVolum),
		Isthmus:         gtc.Float64.PointerToSql(d.Isthmus),
		Struct:          gtc.String.PointerToSql(d.Struct),
		Echogenicity:    gtc.String.PointerToSql(d.Echogenicity),
		RegionalLymph:   gtc.String.PointerToSql(d.RegionalLymph),
		Vascularization: gtc.String.PointerToSql(d.Vascularization),
		Location:        gtc.String.PointerToSql(d.Location),
		Additional:      gtc.String.PointerToSql(d.Additional),
		Conclusion:      gtc.String.PointerToSql(d.Conclusion),
	}
}

func (d Echographic) ToDomain() domain.Echographic {
	return domain.Echographic{
		Id:              d.Id,
		Contors:         gtc.String.SqlToPointer(d.Contors),
		LeftLobeLength:  gtc.Float64.SqlToPointer(d.LeftLobeLength),
		LeftLobeWidth:   gtc.Float64.SqlToPointer(d.LeftLobeWidth),
		LeftLobeThick:   gtc.Float64.SqlToPointer(d.LeftLobeThick),
		LeftLobeVolum:   gtc.Float64.SqlToPointer(d.LeftLobeVolum),
		RightLobeLength: gtc.Float64.SqlToPointer(d.RightLobeLength),
		RightLobeWidth:  gtc.Float64.SqlToPointer(d.RightLobeWidth),
		RightLobeThick:  gtc.Float64.SqlToPointer(d.RightLobeThick),
		RightLobeVolum:  gtc.Float64.SqlToPointer(d.RightLobeVolum),
		GlandVolum:      gtc.Float64.SqlToPointer(d.GlandVolum),
		Isthmus:         gtc.Float64.SqlToPointer(d.Isthmus),
		Struct:          gtc.String.SqlToPointer(d.Struct),
		Echogenicity:    gtc.String.SqlToPointer(d.Echogenicity),
		RegionalLymph:   gtc.String.SqlToPointer(d.RegionalLymph),
		Vascularization: gtc.String.SqlToPointer(d.Vascularization),
		Location:        gtc.String.SqlToPointer(d.Location),
		Additional:      gtc.String.SqlToPointer(d.Additional),
		Conclusion:      gtc.String.SqlToPointer(d.Conclusion),
	}
}
