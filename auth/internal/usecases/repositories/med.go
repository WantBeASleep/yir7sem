// todo https://popovza.kaiten.ru/space/420777/card/37360398
package repositories

import (
	"context"
	"yir/auth/internal/entity"

	"github.com/google/uuid"
)

type MedRepo interface {
	AddMed(ctx context.Context, createData *entity.RequestRegister) (uuid.UUID, error)
}
