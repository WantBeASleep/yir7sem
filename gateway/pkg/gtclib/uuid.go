package gtclib

import "github.com/google/uuid"

type _uuid struct{}

var Uuid _uuid

func (_uuid) StringPToP(p *string) *uuid.UUID {
	if p == nil {
		return nil
	}

	v := uuid.MustParse(*p)
	return &v
}
