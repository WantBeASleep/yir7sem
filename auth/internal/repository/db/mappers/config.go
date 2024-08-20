// не уверен что вот такая реализация ОК, возможно стоит прописать эту логику поместу
package mappers

import (
	"fmt"
	"yir/auth/internal/config"
)

func GetDSN(cfg config.DBConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
	)
}
