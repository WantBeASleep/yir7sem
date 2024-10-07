package entity

import "fmt"

var (
	ErrUnsupportedFormat = fmt.Errorf("unsupported format, use .png or .tiff")
)