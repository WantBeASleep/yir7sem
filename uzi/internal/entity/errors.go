package entity

import (
	"errors"
)

var (
	ErrNotFound          = errors.New("not found")
	ErrUnsupportedFormat = errors.New("unsupported format, use .png or .tiff")
)
