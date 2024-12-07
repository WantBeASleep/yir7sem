package adapters

import (
	"uzi/internal/adapters/broker"
)

type Adapter struct {
	BrokerAdapter broker.BrokerAdapter
}

func New(
	BrokerAdapter broker.BrokerAdapter,
) Adapter {
	return Adapter{
		BrokerAdapter: BrokerAdapter,
	}
}
