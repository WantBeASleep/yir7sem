package grpc

import (
	"gateway/internal/adapters/grpc/auth"
	"gateway/internal/adapters/grpc/med"
	"gateway/internal/adapters/grpc/uzi"
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
