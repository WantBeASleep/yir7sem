package adapters

import (
	"gateway/internal/adapters/broker"
	"gateway/internal/adapters/grpc/auth"
	"gateway/internal/adapters/grpc/med"
	"gateway/internal/adapters/grpc/uzi"
)

type Adapter struct {
	AuthAdapter   auth.AuthAdapter
	MedAdapter    med.MedAdapter
	UziAdapter    uzi.UziAdapter
	BrokerAdapter broker.BrokerAdapter
}

func New(
	AuthAdapter auth.AuthAdapter,
	MedAdapter med.MedAdapter,
	UziAdapter uzi.UziAdapter,
	BrokerAdapter broker.BrokerAdapter,
) Adapter {
	return Adapter{
		AuthAdapter:   AuthAdapter,
		MedAdapter:    MedAdapter,
		UziAdapter:    UziAdapter,
		BrokerAdapter: BrokerAdapter,
	}
}
