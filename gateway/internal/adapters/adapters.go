package adapters

import (
	"yirv2/gateway/internal/adapters/auth"
	"yirv2/gateway/internal/adapters/med"
	"yirv2/gateway/internal/adapters/uzi"
)

type Adapter struct {
	AuthAdapter auth.AuthAdapter
	MedAdapter  med.MedAdapter
	UziAdapter  uzi.UziAdapter
}

func New(
	AuthAdapter auth.AuthAdapter,
	MedAdapter med.MedAdapter,
	UziAdapter uzi.UziAdapter,
) Adapter {
	return Adapter{
		AuthAdapter: AuthAdapter,
		MedAdapter:  MedAdapter,
		UziAdapter:  UziAdapter,
	}
}
