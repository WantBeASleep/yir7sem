package entity

import (
	"database/sql"

	"github.com/WantBeASleep/goooool/gtclib"

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
		Contors:         gtclib.String.PointerToSql(d.Contors),
		LeftLobeLength:  gtclib.Float64.PointerToSql(d.LeftLobeLength),
		LeftLobeWidth:   gtclib.Float64.PointerToSql(d.LeftLobeWidth),
		LeftLobeThick:   gtclib.Float64.PointerToSql(d.LeftLobeThick),
		LeftLobeVolum:   gtclib.Float64.PointerToSql(d.LeftLobeVolum),
		RightLobeLength: gtclib.Float64.PointerToSql(d.RightLobeLength),
		RightLobeWidth:  gtclib.Float64.PointerToSql(d.RightLobeWidth),
		RightLobeThick:  gtclib.Float64.PointerToSql(d.RightLobeThick),
		RightLobeVolum:  gtclib.Float64.PointerToSql(d.RightLobeVolum),
		GlandVolum:      gtclib.Float64.PointerToSql(d.GlandVolum),
		Isthmus:         gtclib.Float64.PointerToSql(d.Isthmus),
		Struct:          gtclib.String.PointerToSql(d.Struct),
		Echogenicity:    gtclib.String.PointerToSql(d.Echogenicity),
		RegionalLymph:   gtclib.String.PointerToSql(d.RegionalLymph),
		Vascularization: gtclib.String.PointerToSql(d.Vascularization),
		Location:        gtclib.String.PointerToSql(d.Location),
		Additional:      gtclib.String.PointerToSql(d.Additional),
		Conclusion:      gtclib.String.PointerToSql(d.Conclusion),
	}
}

func (d Echographic) ToDomain() domain.Echographic {
	return domain.Echographic{
		Id:              d.Id,
		Contors:         gtclib.String.SqlToPointer(d.Contors),
		LeftLobeLength:  gtclib.Float64.SqlToPointer(d.LeftLobeLength),
		LeftLobeWidth:   gtclib.Float64.SqlToPointer(d.LeftLobeWidth),
		LeftLobeThick:   gtclib.Float64.SqlToPointer(d.LeftLobeThick),
		LeftLobeVolum:   gtclib.Float64.SqlToPointer(d.LeftLobeVolum),
		RightLobeLength: gtclib.Float64.SqlToPointer(d.RightLobeLength),
		RightLobeWidth:  gtclib.Float64.SqlToPointer(d.RightLobeWidth),
		RightLobeThick:  gtclib.Float64.SqlToPointer(d.RightLobeThick),
		RightLobeVolum:  gtclib.Float64.SqlToPointer(d.RightLobeVolum),
		GlandVolum:      gtclib.Float64.SqlToPointer(d.GlandVolum),
		Isthmus:         gtclib.Float64.SqlToPointer(d.Isthmus),
		Struct:          gtclib.String.SqlToPointer(d.Struct),
		Echogenicity:    gtclib.String.SqlToPointer(d.Echogenicity),
		RegionalLymph:   gtclib.String.SqlToPointer(d.RegionalLymph),
		Vascularization: gtclib.String.SqlToPointer(d.Vascularization),
		Location:        gtclib.String.SqlToPointer(d.Location),
		Additional:      gtclib.String.SqlToPointer(d.Additional),
		Conclusion:      gtclib.String.SqlToPointer(d.Conclusion),
	}
}
