package entity

import "io"

type File struct {
	Path string
	Data io.Reader
}
