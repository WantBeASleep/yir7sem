package entity

import "io"

type FileMeta struct {
	Path        string
	ContentType string
}

type File struct {
	Meta *FileMeta
	Data io.Reader
}
