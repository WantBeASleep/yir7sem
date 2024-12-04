package domain

import (
	"io"
)

type File struct {
	Format string
	Size   int64
	Buf    io.Reader
}
