// todo https://popovza.kaiten.ru/space/420777/card/37360398
package repositories

import (
	"context"
	"service/auth/internal/entity"
)

type MedRepo interface {
	AddMed(ctx context.Context, createData *entity.RequestRegister) (int, error)
}
