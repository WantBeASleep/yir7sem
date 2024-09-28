package entity

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrNotFound = errors.New("not found")
)

// DB

type ImagesNotFoundError struct {
	Ids uuid.UUIDs
}

func (e *ImagesNotFoundError) Error() string {
	ids := make([]string, 0, len(e.Ids))
	for _, v := range e.Ids {
		ids = append(ids, v.String())
	}

	return fmt.Sprintf("not found images ids: %s", strings.Join(ids, ", "))
}
