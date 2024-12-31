package domain

import (
	"errors"
	"io"
)

type File struct {
	Format string
	Buf    io.Reader
}

// TODO: исправить, или сделать отдельный тип Format (мб лучше взять библу) и в uzi Тоже исправить
func ParseFormatFromExt(s string) (string, error) {
	switch s {
	case ".png":
		return "image/png", nil
	case ".tiff", ".tif":
		return "image/tiff", nil
	}
	return "", errors.New("unsupport") // TODO: вынести в ошибку и обработать сверху
}
