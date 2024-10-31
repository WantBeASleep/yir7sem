package usecases

import (
	"context"
	"yir/uzi/internal/entity"
	"yir/uzi/internal/entity/dto"

	"github.com/google/uuid"
)

type Uzi interface {
	CreateUzi(ctx context.Context, req *entity.Uzi) (uuid.UUID, error)

}
