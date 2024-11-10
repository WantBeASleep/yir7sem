package uzimap

import (
	"yir/gateway/internal/entity/uzimodel"
	"yir/gateway/internal/pb/uzipb"
)

func TiradsToPBTirads(tirads *uzimodel.Tirads) *uzipb.Tirads {
	return &uzipb.Tirads{
		Tirads_23: tirads.Tirads23,
		Tirads_4:  tirads.Tirads4,
		Tirads_5:  tirads.Tirads5,
	}
}

func PBTiradsToTirads(tirads *uzipb.Tirads) *uzimodel.Tirads {
	return &uzimodel.Tirads{
		Tirads23: tirads.Tirads_23,
		Tirads4:  tirads.Tirads_4,
		Tirads5:  tirads.Tirads_5,
	}
}
